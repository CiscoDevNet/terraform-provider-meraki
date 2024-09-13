// Copyright © 2024 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-meraki/gen/yamlconfig"
	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v3"
)

const specPath = "./gen/models/spec3.json"

const usage = `
Usage: go run gen/definition.go <endpoint> <resource_name>

Arguments:
  endpoint      The specific endpoint that is to be converted to generator specification
  resource_name The name that will be given to the resource

Example:
  go run gen/definition.go "/networks/{networkId}/groupPolicies/{groupPolicyId}" "Network Group Policy"`

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Insufficient number of arguments")
		fmt.Println(usage)
		os.Exit(1)
	}

	endpointPath := os.Args[1]
	resourceName := os.Args[2]

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
	config := yamlconfig.YamlConfig{}
	urlResult := parseUrl(endpointPath)
	if urlResult.resultPath[len(urlResult.resultPath)-1] == '/' {
		urlResult.resultPath = urlResult.resultPath[:len(urlResult.resultPath)-1]
	}
	config.RestEndpoint = urlResult.resultPath
	if urlResult.idName != "" {
		config.IdName = urlResult.idName[1 : len(urlResult.idName)-1]
	}
	for i, r := range urlResult.references {
		attr := yamlconfig.YamlConfigAttribute{}
		attr.TfName = yamlconfig.CamelToSnake(r[1 : len(r)-1])
		attr.Type = "String"
		attr.Reference = true
		if urlResult.oneToOne && i == len(urlResult.references)-1 {
			attr.Id = true
		}
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
	oneToOne   bool
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
		ret.oneToOne = true
	}
	ret.category = parts[0][1 : len(parts[0])-1]
	ret.category = strings.ToUpper(string(ret.category[0])) + ret.category[1:]
	return ret
}

func traverseProperties(m map[string]interface{}, path []string, gjsonPath string, exampleStr string) []yamlconfig.YamlConfigAttribute {
	ret := []yamlconfig.YamlConfigAttribute{}
	for propName, v := range m {
		propMap := v.(map[string]interface{})
		if propMap["type"] == "object" {
			childPath := append(path, propName)
			childGjsonPath := gjsonPath + "." + propName
			children := traverseProperties(propMap["properties"].(map[string]interface{}), childPath, childGjsonPath, exampleStr)
			ret = append(ret, children...)
		} else if propMap["type"] == "array" {
			attr := yamlconfig.YamlConfigAttribute{}
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
			attr := yamlconfig.YamlConfigAttribute{}
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
