package yamlconfig

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"text/template"
)

const EndpointToken = "@endpoint:"

type YamlConfig struct {
	Name                string                `yaml:"name,omitempty"`
	TfName              string                `yaml:"tf_name,omitempty"`
	SpecEndpoint        string                `yaml:"spec_endpoint,omitempty"`
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
	TestTags            []string              `yaml:"test_tags,omitempty,flow"`
	TestVariables       []string              `yaml:"test_variables,omitempty,flow"`
	IgnoreAttributes    []string              `yaml:"ignore_attributes,omitempty,flow"`
	Attributes          []YamlConfigAttribute `yaml:"attributes,omitempty"`
	TestPrerequisites   string                `yaml:"test_prerequisites,omitempty"`
}

type YamlConfigP struct {
	Name                *string                 `yaml:"name,omitempty"`
	TfName              *string                 `yaml:"tf_name,omitempty"`
	SpecEndpoint        *string                 `yaml:"spec_endpoint,omitempty"`
	RestEndpoint        *string                 `yaml:"rest_endpoint,omitempty"`
	NoDataSource        *bool                   `yaml:"no_data_source,omitempty"`
	NoResource          *bool                   `yaml:"no_resource,omitempty"`
	PutCreate           *bool                   `yaml:"put_create,omitempty"`
	GetFromAll          *bool                   `yaml:"get_from_all,omitempty"`
	NoUpdate            *bool                   `yaml:"no_update,omitempty"`
	NoDelete            *bool                   `yaml:"no_delete,omitempty"`
	NoImport            *bool                   `yaml:"no_import,omitempty"`
	IdName              *string                 `yaml:"id_name,omitempty"`
	DataSourceNameQuery *bool                   `yaml:"data_source_name_query,omitempty"`
	MinimumVersion      *string                 `yaml:"minimum_version,omitempty"`
	DsDescription       *string                 `yaml:"ds_description,omitempty"`
	ResDescription      *string                 `yaml:"res_description,omitempty"`
	DocCategory         *string                 `yaml:"doc_category,omitempty"`
	ExcludeTest         *bool                   `yaml:"exclude_test,omitempty"`
	SkipMinimumTest     *bool                   `yaml:"skip_minimum_test,omitempty"`
	TestTags            *[]string               `yaml:"test_tags,omitempty,flow"`
	TestVariables       *[]string               `yaml:"test_variables,omitempty,flow"`
	IgnoreAttributes    *[]string               `yaml:"ignore_attributes,omitempty,flow"`
	Attributes          *[]YamlConfigAttributeP `yaml:"attributes,omitempty"`
	TestPrerequisites   *string                 `yaml:"test_prerequisites,omitempty"`
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

type YamlConfigAttributeP struct {
	ModelName        *string                 `yaml:"model_name,omitempty"`
	TfName           *string                 `yaml:"tf_name,omitempty"`
	Type             *string                 `yaml:"type,omitempty"`
	ElementType      *string                 `yaml:"element_type,omitempty"`
	DataPath         *[]string               `yaml:"data_path,omitempty,flow"`
	Id               *bool                   `yaml:"id,omitempty"`
	Reference        *bool                   `yaml:"reference,omitempty"`
	RequiresReplace  *bool                   `yaml:"requires_replace,omitempty"`
	Mandatory        *bool                   `yaml:"mandatory,omitempty"`
	WriteOnly        *bool                   `yaml:"write_only,omitempty"`
	WriteChangesOnly *bool                   `yaml:"write_changes_only,omitempty"`
	ExcludeTest      *bool                   `yaml:"exclude_test,omitempty"`
	ExcludeExample   *bool                   `yaml:"exclude_example,omitempty"`
	Description      *string                 `yaml:"description,omitempty"`
	Example          *string                 `yaml:"example,omitempty"`
	EnumValues       *[]string               `yaml:"enum_values,omitempty,flow"`
	MinList          *int64                  `yaml:"min_list,omitempty"`
	MaxList          *int64                  `yaml:"max_list,omitempty"`
	MinInt           *int64                  `yaml:"min_int,omitempty"`
	MaxInt           *int64                  `yaml:"max_int,omitempty"`
	MinFloat         *float64                `yaml:"min_float,omitempty"`
	MaxFloat         *float64                `yaml:"max_float,omitempty"`
	OrderedList      *bool                   `yaml:"ordered_list,omitempty"`
	StringPatterns   *[]string               `yaml:"string_patterns,omitempty,flow"`
	StringMinLength  *int64                  `yaml:"string_min_length,omitempty"`
	StringMaxLength  *int64                  `yaml:"string_max_length,omitempty"`
	DefaultValue     *string                 `yaml:"default_value,omitempty"`
	Value            *string                 `yaml:"value,omitempty"`
	TestValue        *string                 `yaml:"test_value,omitempty"`
	MinimumTestValue *string                 `yaml:"minimum_test_value,omitempty"`
	TestTags         *[]string               `yaml:"test_tags,omitempty,flow"`
	Attributes       *[]YamlConfigAttributeP `yaml:"attributes,omitempty"`
	GoTypeName       *string                 `yaml:"gotypename,omitempty"`
}

func P[T any](v T) *T {
	return &v
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
	s = strings.ReplaceAll(s, ".", "")
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
	s = strings.ReplaceAll(s, ".", "")
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
func ImportParts(config YamlConfig) int {
	parts := 0
	for _, attr := range config.Attributes {
		if attr.Reference {
			parts += 1
		} else if attr.Id {
			parts += 1
		}
	}
	if !HasId(config.Attributes) {
		parts += 1
	}

	return parts
}

// Templating helper function to subtract one number from another
func Subtract(a, b int) int {
	return a - b
}

// Templating helper function to iterate over a count
func Iterate(count int) []int {
	var i int
	var Items []int
	for i = 0; i < count; i++ {
		Items = append(Items, i)
	}
	return Items
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
	"iterate":         Iterate,
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

func MergeYamlConfig(existing *YamlConfigP, new *YamlConfigP) *YamlConfigP {
	// Merge existing into new
	if existing.Name != nil {
		new.Name = existing.Name
	}
	if existing.TfName != nil {
		new.TfName = existing.TfName
	}
	if existing.SpecEndpoint != nil {
		new.SpecEndpoint = existing.SpecEndpoint
	}
	if existing.RestEndpoint != nil {
		new.RestEndpoint = existing.RestEndpoint
	}
	if existing.NoDataSource != nil {
		new.NoDataSource = existing.NoDataSource
	}
	if existing.NoResource != nil {
		new.NoResource = existing.NoResource
	}
	if existing.PutCreate != nil {
		new.PutCreate = existing.PutCreate
	}
	if existing.GetFromAll != nil {
		new.GetFromAll = existing.GetFromAll
	}
	if existing.NoUpdate != nil {
		new.NoUpdate = existing.NoUpdate
	}
	if existing.NoDelete != nil {
		new.NoDelete = existing.NoDelete
	}
	if existing.NoImport != nil {
		new.NoImport = existing.NoImport
	}
	if existing.IdName != nil {
		new.IdName = existing.IdName
	}
	if existing.DataSourceNameQuery != nil {
		new.DataSourceNameQuery = existing.DataSourceNameQuery
	}
	if existing.MinimumVersion != nil {
		new.MinimumVersion = existing.MinimumVersion
	}
	if existing.DsDescription != nil {
		new.DsDescription = existing.DsDescription
	}
	if existing.ResDescription != nil {
		new.ResDescription = existing.ResDescription
	}
	if existing.DocCategory != nil {
		new.DocCategory = existing.DocCategory
	}
	if existing.ExcludeTest != nil {
		new.ExcludeTest = existing.ExcludeTest
	}
	if existing.SkipMinimumTest != nil {
		new.SkipMinimumTest = existing.SkipMinimumTest
	}
	if existing.TestTags != nil {
		new.TestTags = existing.TestTags
	}
	if existing.TestVariables != nil {
		new.TestVariables = existing.TestVariables
	}
	if existing.IgnoreAttributes != nil {
		new.IgnoreAttributes = existing.IgnoreAttributes
	}
	if existing.Attributes != nil {
		new.Attributes = MergeYamlConfigAttributes(existing.Attributes, new.Attributes)
	}
	if existing.TestPrerequisites != nil {
		new.TestPrerequisites = existing.TestPrerequisites
	}
	return new
}

func cutYamlConfigAttribute(i int, attributes []YamlConfigAttributeP) []YamlConfigAttributeP {
	cut := append(attributes[:i], attributes[i+1:]...)
	return cut
}

func MergeYamlConfigAttributes(existing *[]YamlConfigAttributeP, new *[]YamlConfigAttributeP) *[]YamlConfigAttributeP {
	// Merge existing into new
	var c []YamlConfigAttributeP
	for _, existingAttr := range *existing {
		found := -1
		for i, attr := range *new {
			if existingAttr.ModelName != nil && *existingAttr.ModelName != "" && attr.ModelName != nil && *attr.ModelName != "" && *existingAttr.ModelName == *attr.ModelName {
				if (existingAttr.DataPath == nil && attr.DataPath == nil) || (existingAttr.DataPath != nil && attr.DataPath != nil && slices.Equal(*existingAttr.DataPath, *attr.DataPath)) {
					c = append(c, *MergeYamlConfigAttribute(&existingAttr, &attr))
					found = i
					break
				}
			} else if existingAttr.TfName != nil && *existingAttr.TfName != "" && attr.TfName != nil && *attr.TfName != "" && *existingAttr.TfName == *attr.TfName {
				c = append(c, *MergeYamlConfigAttribute(&existingAttr, &attr))
				found = i
				break
			}
		}
		if found > -1 {
			// we have found a matching new attribute, so we remove it from the list
			new = P(cutYamlConfigAttribute(found, *new))
			// check if there are more elements from the new list and add them if there is no matching attribute in existing list
			foundCount := 0
			for _, attr := range *new {
				foundExisting := false
				for _, eAttr := range *existing {
					if eAttr.ModelName != nil && *eAttr.ModelName != "" && attr.ModelName != nil && *attr.ModelName != "" && *eAttr.ModelName == *attr.ModelName {
						if (eAttr.DataPath == nil && attr.DataPath == nil) || (eAttr.DataPath != nil && attr.DataPath != nil && slices.Equal(*eAttr.DataPath, *attr.DataPath)) {
							foundExisting = true
							break
						}
					} else if eAttr.TfName != nil && *eAttr.TfName != "" && attr.TfName != nil && *attr.TfName != "" && *eAttr.TfName == *attr.TfName {
						foundExisting = true
						break
					}
				}
				if foundExisting {
					// if there is a matching attribute in existing list, we stop as it will be added later
					break
				} else {
					// if there is no matching attribute in existing list, it is a new element and we add it to the list
					foundCount += 1
					c = append(c, attr)
				}
			}
			for i := 0; i < foundCount; i++ {
				// if we have found a new elements that were not in the existing list, we remove them from the new list
				new = P(cutYamlConfigAttribute(0, *new))
			}
		} else {
			// we haven't found a matching new attribute, so we add the existing to the list
			c = append(c, *MergeYamlConfigAttribute(&YamlConfigAttributeP{}, &existingAttr))
		}
	}
	// add remaining elements from existing list
	c = append(c, *new...)
	return &c
}

func MergeYamlConfigAttribute(existing *YamlConfigAttributeP, new *YamlConfigAttributeP) *YamlConfigAttributeP {
	// Merge existing into new
	if existing.ModelName != nil {
		new.ModelName = existing.ModelName
	}
	if existing.TfName != nil {
		new.TfName = existing.TfName
	}
	if existing.Type != nil {
		new.Type = existing.Type
	}
	if existing.ElementType != nil {
		new.ElementType = existing.ElementType
	}
	if existing.DataPath != nil {
		new.DataPath = existing.DataPath
	}
	if existing.Id != nil {
		new.Id = existing.Id
	}
	if existing.Reference != nil {
		new.Reference = existing.Reference
	}
	if existing.RequiresReplace != nil {
		new.RequiresReplace = existing.RequiresReplace
	}
	if existing.Mandatory != nil {
		new.Mandatory = existing.Mandatory
	}
	if existing.WriteOnly != nil {
		new.WriteOnly = existing.WriteOnly
	}
	if existing.WriteChangesOnly != nil {
		new.WriteChangesOnly = existing.WriteChangesOnly
	}
	if existing.ExcludeTest != nil {
		new.ExcludeTest = existing.ExcludeTest
	}
	if existing.ExcludeExample != nil {
		new.ExcludeExample = existing.ExcludeExample
	}
	if existing.Description != nil {
		new.Description = existing.Description
	}
	if existing.Example != nil {
		new.Example = existing.Example
	}
	if existing.EnumValues != nil {
		new.EnumValues = existing.EnumValues
	}
	if existing.MinList != nil {
		new.MinList = existing.MinList
	}
	if existing.MaxList != nil {
		new.MaxList = existing.MaxList
	}
	if existing.MinInt != nil {
		new.MinInt = existing.MinInt
	}
	if existing.MaxInt != nil {
		new.MaxInt = existing.MaxInt
	}
	if existing.MinFloat != nil {
		new.MinFloat = existing.MinFloat
	}
	if existing.MaxFloat != nil {
		new.MaxFloat = existing.MaxFloat
	}
	if existing.OrderedList != nil {
		new.OrderedList = existing.OrderedList
	}
	if existing.StringPatterns != nil {
		new.StringPatterns = existing.StringPatterns
	}
	if existing.StringMinLength != nil {
		new.StringMinLength = existing.StringMinLength
	}
	if existing.StringMaxLength != nil {
		new.StringMaxLength = existing.StringMaxLength
	}
	if existing.DefaultValue != nil {
		new.DefaultValue = existing.DefaultValue
	}
	if existing.Value != nil {
		new.Value = existing.Value
	}
	if existing.TestValue != nil {
		new.TestValue = existing.TestValue
	}
	if existing.MinimumTestValue != nil {
		new.MinimumTestValue = existing.MinimumTestValue
	}
	if existing.TestTags != nil {
		new.TestTags = existing.TestTags
	}
	if existing.Attributes != nil {
		new.Attributes = MergeYamlConfigAttributes(existing.Attributes, new.Attributes)
	}
	if existing.GoTypeName != nil {
		new.GoTypeName = existing.GoTypeName
	}
	return new
}
