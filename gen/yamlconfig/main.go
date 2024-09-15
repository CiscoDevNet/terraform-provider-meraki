package yamlconfig

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"text/template"
)

type YamlConfig struct {
	Name                string                `yaml:"name,omitempty"`
	TfName              string                `yaml:"tf_name,omitempty"`
	RestEndpoint        string                `yaml:"rest_endpoint,omitempty"`
	NoDataSource        bool                  `yaml:"no_data_source,omitempty"`
	NoResource          bool                  `yaml:"no_resource,omitempty"`
	PutCreate           bool                  `yaml:"put_create,omitempty"`
	GetFromAll          bool                  `yaml:"get_from_all,omitempty"`
	NoUpdate            bool                  `yaml:"no_update,omitempty"`
	NoDelete            bool                  `yaml:"no_delete,omitempty"`
	NoImport            bool                  `yaml:"no_import,omitempty"`
	IdName              string                `yaml:"id_name,omitempty"`
	DataSourceNameQuery bool                  `yaml:"data_source_name_query,omitempty"`
	MinimumVersion      string                `yaml:"minimum_version,omitempty"`
	DsDescription       string                `yaml:"ds_description,omitempty"`
	ResDescription      string                `yaml:"res_description,omitempty"`
	DocCategory         string                `yaml:"doc_category,omitempty"`
	ExcludeTest         bool                  `yaml:"exclude_test,omitempty"`
	SkipMinimumTest     bool                  `yaml:"skip_minimum_test,omitempty"`
	TestTags            []string              `yaml:"test_tags,omitempty"`
	Attributes          []YamlConfigAttribute `yaml:"attributes,omitempty"`
	TestPrerequisites   string                `yaml:"test_prerequisites,omitempty"`
}

type YamlConfigAttribute struct {
	ModelName        string                `yaml:"model_name,omitempty"`
	TfName           string                `yaml:"tf_name,omitempty"`
	Type             string                `yaml:"type,omitempty"`
	ElementType      string                `yaml:"element_type,omitempty"`
	DataPath         []string              `yaml:"data_path,omitempty,flow"`
	Id               bool                  `yaml:"id,omitempty"`
	Reference        bool                  `yaml:"reference,omitempty"`
	RequiresReplace  bool                  `yaml:"requires_replace,omitempty"`
	Mandatory        bool                  `yaml:"mandatory,omitempty"`
	WriteOnly        bool                  `yaml:"write_only,omitempty"`
	WriteChangesOnly bool                  `yaml:"write_changes_only,omitempty"`
	ExcludeTest      bool                  `yaml:"exclude_test,omitempty"`
	ExcludeExample   bool                  `yaml:"exclude_example,omitempty"`
	Description      string                `yaml:"description,omitempty"`
	Example          string                `yaml:"example,omitempty"`
	EnumValues       []string              `yaml:"enum_values,omitempty,flow"`
	MinList          int64                 `yaml:"min_list,omitempty"`
	MaxList          int64                 `yaml:"max_list,omitempty"`
	MinInt           int64                 `yaml:"min_int,omitempty"`
	MaxInt           int64                 `yaml:"max_int,omitempty"`
	MinFloat         float64               `yaml:"min_float,omitempty"`
	MaxFloat         float64               `yaml:"max_float,omitempty"`
	OrderedList      bool                  `yaml:"ordered_list,omitempty"`
	StringPatterns   []string              `yaml:"string_patterns,omitempty,flow"`
	StringMinLength  int64                 `yaml:"string_min_length,omitempty"`
	StringMaxLength  int64                 `yaml:"string_max_length,omitempty"`
	DefaultValue     string                `yaml:"default_value,omitempty"`
	Value            string                `yaml:"value,omitempty"`
	TestValue        string                `yaml:"test_value,omitempty"`
	MinimumTestValue string                `yaml:"minimum_test_value,omitempty"`
	TestTags         []string              `yaml:"test_tags,omitempty,flow"`
	Attributes       []YamlConfigAttribute `yaml:"attributes,omitempty"`
	GoTypeName       string                `yaml:"gotypename,omitempty"`
}

// Templating helper function to convert TF name to GO name
func ToGoName(s string) string {
	var g []string

	p := strings.Split(s, "_")

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	return s
}

// Templating helper function to convert string to camel case
func CamelCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
}

// Templating helper function to convert string to snake case
func SnakeCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// Templating helper function to fail a template mid-way
func Errorf(s string, args ...any) (struct{}, error) {
	return struct{}{}, fmt.Errorf(s, args...)
}

// Templating helper function to build a SJSON path
func BuildPath(s []string) string {
	return strings.Join(s, ".")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Templating helper function to return true if id included in attributes
func HasId(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Id {
			return true
		}
	}
	return false
}

// Templating helper function to return the id
func GetId(attributes []YamlConfigAttribute) YamlConfigAttribute {
	for _, attr := range attributes {
		if attr.Id {
			return attr
		}
	}
	return YamlConfigAttribute{}
}

// Templating helper function to return true if reference included in attributes
func HasReference(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Reference {
			return true
		}
	}
	return false
}

// Templating helper function to return true if type is a list or set without nested elements
func IsListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list without nested elements
func IsList(attribute YamlConfigAttribute) bool {
	if attribute.Type == "List" && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a set without nested elements
func IsSet(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Set" && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set of strings without nested elements
func IsStringListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "String" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set of integers without nested elements
func IsInt64ListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "Int64" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set with nested elements
func IsNestedListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list with nested elements
func IsNestedList(attribute YamlConfigAttribute) bool {
	if attribute.Type == "List" && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a set with nested elements
func IsNestedSet(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Set" && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return number of import parts
func ImportParts(attributes []YamlConfigAttribute) int {
	parts := 1
	for _, attr := range attributes {
		if attr.Reference {
			parts += 1
		} else if attr.Id {
			parts += 1
		}
	}
	return parts
}

// Templating helper function to subtract one number from another
func Subtract(a, b int) int {
	return a - b
}

// Map of templating functions
var Functions = template.FuncMap{
	"toGoName":        ToGoName,
	"camelCase":       CamelCase,
	"snakeCase":       SnakeCase,
	"sprintf":         fmt.Sprintf,
	"errorf":          Errorf,
	"toLower":         strings.ToLower,
	"path":            BuildPath,
	"hasId":           HasId,
	"getId":           GetId,
	"hasReference":    HasReference,
	"isListSet":       IsListSet,
	"isList":          IsList,
	"isSet":           IsSet,
	"isStringListSet": IsStringListSet,
	"isInt64ListSet":  IsInt64ListSet,
	"isNestedListSet": IsNestedListSet,
	"isNestedList":    IsNestedList,
	"isNestedSet":     IsNestedSet,
	"importParts":     ImportParts,
	"subtract":        Subtract,
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func CamelToSnake(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (attr *YamlConfigAttribute) Init(parentGoTypeName string) error {
	// Augument
	if attr.TfName == "" {
		// var words []string
		fullString := ""
		for _, s := range attr.DataPath {
			fullString += strings.ToUpper(string(s[0])) + s[1:]
		}
		fullString += strings.ToUpper(string(attr.ModelName[0])) + attr.ModelName[1:]
		// l := 0
		// for s := fullString; s != ""; s = s[l:] {
		// 	l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		// 	if l <= 0 {
		// 		l = len(s)
		// 	}
		// 	words = append(words, strings.ToLower(s[:l]))
		// }
		// attr.TfName = strings.Join(words, "_")
		attr.TfName = CamelToSnake(fullString)
	}

	attr.GoTypeName = parentGoTypeName + ToGoName(attr.TfName)

	// Validate
	if len(attr.Attributes) > 0 && attr.Type != "List" && attr.Type != "Map" && attr.Type != "Set" {
		return fmt.Errorf("%q has type %q which cannot have `attributes`: instead use type List, Map, Set",
			attr.TfName, attr.Type)
	}

	if len(attr.Attributes) > 0 && attr.ElementType != "" {
		return fmt.Errorf("%q: either `attributes` or `element_type` can be specified, but not both", attr.TfName)
	}

	if attr.Type == "Map" && attr.ElementType != "" {
		return fmt.Errorf("%q: the `element_type` is not yet implemented for type Map", attr.TfName)
	}

	if attr.OrderedList {
		if attr.Type != "List" {
			return fmt.Errorf("%q has type %q which cannot use `ordered_list`: instead use type List",
				attr.TfName, attr.Type)
		}
		if HasId(attr.Attributes) {
			return fmt.Errorf("%q: the `ordered_list: true` conflicts with sub-attributes having `id: true`, as it treats list index ([i]) as the only unique id",
				attr.TfName)
		}
	}

	if attr.Type == "Map" && HasId(attr.Attributes) {
		return fmt.Errorf("Map %q cannot contain sub-attributes with `id: true`, as it treats map key ([k]) as the only unique id",
			attr.TfName)
	}

	// Recurse
	for i := range attr.Attributes {
		if err := attr.Attributes[i].Init(attr.GoTypeName); err != nil {
			return err
		}
	}

	return nil
}

func MergeYamlConfig(a, b YamlConfig) YamlConfig {
	// Merge b into a
	if b.Name != "" {
		a.Name = b.Name
	}
	if b.TfName != "" {
		a.TfName = b.TfName
	}
	if b.RestEndpoint != "" {
		a.RestEndpoint = b.RestEndpoint
	}
	if b.NoDataSource {
		a.NoDataSource = b.NoDataSource
	}
	if b.NoResource {
		a.NoResource = b.NoResource
	}
	if b.PutCreate {
		a.PutCreate = b.PutCreate
	}
	if b.GetFromAll {
		a.GetFromAll = b.GetFromAll
	}
	if b.NoUpdate {
		a.NoUpdate = b.NoUpdate
	}
	if b.NoDelete {
		a.NoDelete = b.NoDelete
	}
	if b.NoImport {
		a.NoImport = b.NoImport
	}
	if b.IdName != "" {
		a.IdName = b.IdName
	}
	if b.DataSourceNameQuery {
		a.DataSourceNameQuery = b.DataSourceNameQuery
	}
	if b.MinimumVersion != "" {
		a.MinimumVersion = b.MinimumVersion
	}
	if b.DsDescription != "" {
		a.DsDescription = b.DsDescription
	}
	if b.ResDescription != "" {
		a.ResDescription = b.ResDescription
	}
	if b.DocCategory != "" {
		a.DocCategory = b.DocCategory
	}
	if b.ExcludeTest {
		a.ExcludeTest = b.ExcludeTest
	}
	if b.SkipMinimumTest {
		a.SkipMinimumTest = b.SkipMinimumTest
	}
	if b.TestTags != nil {
		a.TestTags = b.TestTags
	}
	if b.Attributes != nil {
		a.Attributes = MergeYamlConfigAttributes(a.Attributes, b.Attributes)
	}
	if b.TestPrerequisites != "" {
		a.TestPrerequisites = b.TestPrerequisites
	}
	return a
}

func MergeYamlConfigAttributes(a, b []YamlConfigAttribute) []YamlConfigAttribute {
	// Merge b into a
	var c []YamlConfigAttribute
	for _, attr := range b {
		found := false
		for i := range a {
			if a[i].ModelName != "" && attr.ModelName != "" && slices.Equal(a[i].DataPath, attr.DataPath) {
				if a[i].ModelName == attr.ModelName {
					c = append(c, MergeYamlConfigAttribute(a[i], attr))
					found = true
					break
				}
			} else if a[i].TfName != "" && attr.TfName != "" {
				if a[i].TfName == attr.TfName {
					c = append(c, MergeYamlConfigAttribute(a[i], attr))
					found = true
					break
				}
			}
		}
		if !found {
			c = append(c, attr)
		}
	}
	return c
}

func MergeYamlConfigAttribute(a, b YamlConfigAttribute) YamlConfigAttribute {
	// Merge b into a
	if b.ModelName != "" {
		a.ModelName = b.ModelName
	}
	if b.TfName != "" {
		a.TfName = b.TfName
	}
	if b.Type != "" {
		a.Type = b.Type
	}
	if b.ElementType != "" {
		a.ElementType = b.ElementType
	}
	if b.DataPath != nil {
		a.DataPath = b.DataPath
	}
	if b.Id {
		a.Id = b.Id
	}
	if b.Reference {
		a.Reference = b.Reference
	}
	if b.RequiresReplace {
		a.RequiresReplace = b.RequiresReplace
	}
	if b.Mandatory {
		a.Mandatory = b.Mandatory
	}
	if b.WriteOnly {
		a.WriteOnly = b.WriteOnly
	}
	if b.WriteChangesOnly {
		a.WriteChangesOnly = b.WriteChangesOnly
	}
	if b.ExcludeTest {
		a.ExcludeTest = b.ExcludeTest
	}
	if b.ExcludeExample {
		a.ExcludeExample = b.ExcludeExample
	}
	if b.Description != "" {
		a.Description = b.Description
	}
	if b.Example != "" {
		a.Example = b.Example
	}
	if b.EnumValues != nil {
		a.EnumValues = b.EnumValues
	}
	if b.MinList != 0 {
		a.MinList = b.MinList
	}
	if b.MaxList != 0 {
		a.MaxList = b.MaxList
	}
	if b.MinInt != 0 {
		a.MinInt = b.MinInt
	}
	if b.MaxInt != 0 {
		a.MaxInt = b.MaxInt
	}
	if b.MinFloat != 0 {
		a.MinFloat = b.MinFloat
	}
	if b.MaxFloat != 0 {
		a.MaxFloat = b.MaxFloat
	}
	if b.OrderedList {
		a.OrderedList = b.OrderedList
	}
	if b.StringPatterns != nil {
		a.StringPatterns = b.StringPatterns
	}
	if b.StringMinLength != 0 {
		a.StringMinLength = b.StringMinLength
	}
	if b.StringMaxLength != 0 {
		a.StringMaxLength = b.StringMaxLength
	}
	if b.DefaultValue != "" {
		a.DefaultValue = b.DefaultValue
	}
	if b.Value != "" {
		a.Value = b.Value
	}
	if b.TestValue != "" {
		a.TestValue = b.TestValue
	}
	if b.MinimumTestValue != "" {
		a.MinimumTestValue = b.MinimumTestValue
	}
	if b.TestTags != nil {
		a.TestTags = b.TestTags
	}
	if b.Attributes != nil {
		a.Attributes = MergeYamlConfigAttributes(a.Attributes, b.Attributes)
	}
	if b.GoTypeName != "" {
		a.GoTypeName = b.GoTypeName
	}
	return a
}
