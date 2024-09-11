package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/netascode/terraform-provider-meraki/gen"
	"gopkg.in/yaml.v2"
)

const usage = `
Usage: openapiconverter <openapi_spec> <endpoint>

Arguments:
  openapi_spec  Path to the file containing the OpenAPI specification (YAML or JSON)
  endpoint      The specific endpoint that is to be converted to generator specification

Example:
  openapiconverter ./api-spec.yaml "/networks/{networkId}/groupPolicies/{groupPolicyId}"`

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Insufficient number of arguments")
		fmt.Println(usage)
		os.Exit(1)
	}

	specPath := os.Args[1]
	endpointPath := os.Args[2]

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

	attributes := traverseProperties(schema["schema"].(map[string]interface{})["properties"].(map[string]interface{}), []string{}, false)
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

func traverseProperties(m map[string]interface{}, path []string, isList bool) []gen.YamlConfigAttribute {
	ret := []gen.YamlConfigAttribute{}
	for propName, v := range m {
		propMap := v.(map[string]interface{})
		if propMap["type"] == "object" {
			childPath := append(path, propName)
			children := traverseProperties(propMap["properties"].(map[string]interface{}), childPath, isList)
			ret = append(ret, children...)
		} else if propMap["type"] == "array" {
			attr := gen.YamlConfigAttribute{}
			attr.DataPath = path
			attr.Type = "List"
			attr.ModelName = propName
			desc, ok := propMap["description"]
			if ok {
				attr.Description = desc.(string)
			}
			if isList {
				attr.ElementType = "String"
			} else {
				items := propMap["items"].(map[string]interface{})
				if items["type"].(string) == "object" {
					children := traverseProperties(items["properties"].(map[string]interface{}), []string{}, true)
					attr.Attributes = children
				} else {
					attr.ElementType = "String"
				}

			}
			ret = append(ret, attr)
		} else {
			// primitive value
			attr := gen.YamlConfigAttribute{}
			attr.DataPath = path
			attr.Type = jsonTypes[propMap["type"].(string)]
			attr.ModelName = propName
			desc, ok := propMap["description"]
			if ok {
				attr.Description = desc.(string)
			}
			ret = append(ret, attr)
		}
	}
	return ret
}
