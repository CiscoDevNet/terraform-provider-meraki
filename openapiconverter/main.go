package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-meraki/gen"
	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v2"
)

const usage = `
Usage: openapiconverter <openapi_spec> <endpoint> <resource_name>

Arguments:
  openapi_spec  Path to the file containing the OpenAPI specification (YAML or JSON)
  endpoint      The specific endpoint that is to be converted to generator specification
  resource_name The name that will be given to the resource

Example:
  openapiconverter ./api-spec.yaml "/networks/{networkId}/groupPolicies/{groupPolicyId}" "Network Group Policy"`

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Error: Insufficient number of arguments")
		fmt.Println(usage)
		os.Exit(1)
	}

	specPath := os.Args[1]
	endpointPath := os.Args[2]
	resourceName := os.Args[3]

	specData, err := os.ReadFile(specPath)
	if err != nil {
		fmt.Printf("Error reading OpenAPI spec file: %v\n", err)
		os.Exit(1)
	}

	var spec interface{}
	if strings.HasSuffix(specPath, ".json") {
		err = json.Unmarshal(specData, &spec)
	} else {
		err = yaml.Unmarshal(specData, &spec)
	}
	if err != nil {
		fmt.Printf("Error parsing OpenAPI spec: %v\n", err)
		os.Exit(1)
	}

	endpoint := spec.(map[string]interface{})["paths"].(map[string]interface{})[endpointPath].(map[string]interface{})
	schema := endpoint["put"].(map[string]interface{})["requestBody"].(map[string]interface{})["content"].(map[string]interface{})["application/json"].(map[string]interface{})
	example := schema["schema"].(map[string]interface{})["example"].(map[string]interface{})
	exampleStr, err := json.Marshal(&example)
	if err != nil {
		panic(err)
	}

	attributes := traverseProperties(schema["schema"].(map[string]interface{})["properties"].(map[string]interface{}), []string{}, "", string(exampleStr))
	config := gen.YamlConfig{}
	urlResult := parseUrl(endpointPath)
	if urlResult.resultPath[len(urlResult.resultPath)-1] == '/' {
		urlResult.resultPath = urlResult.resultPath[:len(urlResult.resultPath)-1]
	}
	config.RestEndpoint = urlResult.resultPath
	if urlResult.idName != "" {
		config.IdName = urlResult.idName[1 : len(urlResult.idName)-1]
	}
	for _, r := range urlResult.references {
		attr := gen.YamlConfigAttribute{}
		attr.TfName = gen.CamelToSnake(r[1 : len(r)-1])
		attr.Type = "String"
		attr.Reference = true
		attributes = append(attributes, attr)
	}
	config.Attributes = attributes
	config.DocCategory = urlResult.category
	dataSourceNameQuery := false
	for _, a := range config.Attributes {
		if a.ModelName == "name" {
			dataSourceNameQuery = true
			break
		}
	}
	config.DataSourceNameQuery = dataSourceNameQuery
	config.Name = resourceName
	yamlStr, err := yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}
	fmt.Println("---")
	fmt.Println(string(yamlStr))
}

var jsonTypes = map[string]string{
	"integer": "Int64",
	"boolean": "Bool",
	"string":  "String",
}

type parseUrlResult struct {
	idName     string
	resultPath string
	references []string
	category   string
}

func parseUrl(url string) parseUrlResult {
	ret := parseUrlResult{}
	r := regexp.MustCompile(`{[a-zA-Z]+}`)
	parts := r.Split(url, -1)
	ids := r.FindAllString(url, -1)
	if url[len(url)-1] == '}' {
		// one to many
		ret.idName = ids[len(ids)-1]
		ret.resultPath = strings.Join(parts[:len(parts)-1], "%v")
		ret.references = ids[:len(ids)-1]
	} else {
		// one to one
		ret.resultPath = strings.Join(parts, "%v")
		ret.references = ids
	}
	ret.category = parts[0][1 : len(parts[0])-1]
	ret.category = strings.ToUpper(string(ret.category[0])) + ret.category[1:]
	return ret
}

func traverseProperties(m map[string]interface{}, path []string, gjsonPath string, exampleStr string) []gen.YamlConfigAttribute {
	ret := []gen.YamlConfigAttribute{}
	for propName, v := range m {
		propMap := v.(map[string]interface{})
		if propMap["type"] == "object" {
			childPath := append(path, propName)
			childGjsonPath := gjsonPath + "." + propName
			children := traverseProperties(propMap["properties"].(map[string]interface{}), childPath, childGjsonPath, exampleStr)
			ret = append(ret, children...)
		} else if propMap["type"] == "array" {
			attr := gen.YamlConfigAttribute{}
			attr.DataPath = path
			attr.Type = "List"
			attr.ModelName = propName
			items := propMap["items"].(map[string]interface{})
			desc, ok := propMap["description"]
			if ok {
				attr.Description = sanitizeDescription(desc.(string))
			}
			if items["type"].(string) == "string" {
				attr.ElementType = "String"
				childGjsonPath := (gjsonPath + "." + propName + ".0")[1:]
				res := gjson.Get(exampleStr, childGjsonPath)
				attr.Example = res.String()
			} else if items["type"].(string) == "object" {
				childGjsonPath := gjsonPath + "." + propName + ".0"
				children := traverseProperties(items["properties"].(map[string]interface{}), []string{}, childGjsonPath, exampleStr)
				attr.Attributes = children
			}
			ret = append(ret, attr)
		} else {
			// primitive value
			attr := gen.YamlConfigAttribute{}
			attr.DataPath = path
			attr.Type = jsonTypes[propMap["type"].(string)]
			attr.ModelName = propName
			childGjsonPath := (gjsonPath + "." + propName)[1:]
			res := gjson.Get(exampleStr, childGjsonPath)
			attr.Example = res.String()
			desc, ok := propMap["description"]
			if ok {
				attr.Description = sanitizeDescription(desc.(string))
			}
			ret = append(ret, attr)
		}
	}
	return ret
}

func sanitizeDescription(desc string) string {
	desc = strings.ReplaceAll(desc, "\n", " ")
	desc = strings.ReplaceAll(desc, "\"", "'")
	return desc
}
