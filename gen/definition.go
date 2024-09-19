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
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-meraki/gen/yamlconfig"
	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v3"

	"golang.org/x/exp/maps"
)

const specPath = "./gen/models/spec3.json"
const definitionsPath = "./gen/definitions/"

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

	generateDefinition(endpointPath, resourceName)
}

func generateDefinition(endpointPath, resourceName string) {
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

	config := yamlconfig.YamlConfig{}

	shortEndpointPath := ""
	if endpointPath[len(endpointPath)-1] == '}' {
		parts := strings.Split(endpointPath, "/")
		if len(parts) > 0 {
			parts = parts[:len(parts)-1]
		}
		shortEndpointPath = strings.Join(parts, "/")
	} else {
		shortEndpointPath = endpointPath

	}
	var schema map[string]interface{}
	paths := spec.(map[string]interface{})["paths"].(map[string]interface{})
	// use POST schema if it exists, otherwise fall back to PUT schema
	if sep, ok := paths[shortEndpointPath]; ok {
		if endpoint, ok := sep.(map[string]interface{})["post"]; ok {
			schema = endpoint.(map[string]interface{})["requestBody"].(map[string]interface{})["content"].(map[string]interface{})["application/json"].(map[string]interface{})
		}
	} else if sep, ok := paths[endpointPath]; ok {
		if endpoint, ok := sep.(map[string]interface{})["post"]; ok {
			schema = endpoint.(map[string]interface{})["requestBody"].(map[string]interface{})["content"].(map[string]interface{})["application/json"].(map[string]interface{})
		}
	}
	if schema == nil {
		schema = paths[endpointPath].(map[string]interface{})["put"].(map[string]interface{})["requestBody"].(map[string]interface{})["content"].(map[string]interface{})["application/json"].(map[string]interface{})
		config.PutCreate = true
	}
	example := schema["schema"].(map[string]interface{})["example"].(map[string]interface{})
	exampleStr, err := json.Marshal(&example)
	if err != nil {
		panic(err)
	}

	urlResult := parseUrl(endpointPath)
	if urlResult.resultPath[len(urlResult.resultPath)-1] == '/' {
		urlResult.resultPath = urlResult.resultPath[:len(urlResult.resultPath)-1]
	}
	config.RestEndpoint = urlResult.resultPath
	config.DocCategory = urlResult.category
	config.Name = resourceName
	if urlResult.oneToOne && config.PutCreate {
		config.NoDelete = true
	}

	attributes := []yamlconfig.YamlConfigAttribute{}
	for i, r := range urlResult.references {
		attr := yamlconfig.YamlConfigAttribute{}
		attr.TfName = yamlconfig.CamelToSnake(r[1 : len(r)-1])
		attr.Type = "String"
		attr.Reference = true
		if urlResult.oneToOne && i == len(urlResult.references)-1 {
			attr.Id = true
		}
		attr.Description = "<<Description>>"
		attr.TestValue = "<<TestValue>>"
		attr.Example = "<<Example>>"
		attributes = append(attributes, attr)
	}
	required := []string{}
	if r, ok := schema["schema"].(map[string]interface{})["required"]; ok {
		required = toStringSlice(r.([]interface{}))
	}
	attributes = append(attributes, traverseProperties(schema["schema"].(map[string]interface{})["properties"].(map[string]interface{}), []string{}, "", string(exampleStr), required)...)
	config.Attributes = attributes

	dataSourceNameQuery := false
	for _, a := range config.Attributes {
		if a.ModelName == "name" && len(a.DataPath) == 0 && !config.PutCreate {
			dataSourceNameQuery = true
			break
		}
	}
	config.DataSourceNameQuery = dataSourceNameQuery

	outputPath := definitionsPath + yamlconfig.SnakeCase(resourceName) + ".yaml"

	existingConfig := yamlconfig.YamlConfig{}
	comments := ""
	commentsEndpoint := ""
	if yamlFile, err := os.ReadFile(outputPath); err == nil {
		// retain comments at the beginning of the definition file
		scanner := bufio.NewScanner(bytes.NewReader(yamlFile))
		for scanner.Scan() {
			line := scanner.Text()
			if line[0] == '#' {
				if strings.Contains(line, yamlconfig.EndpointToken) {
					commentsEndpoint = strings.Trim(strings.Split(line, yamlconfig.EndpointToken)[1], " ")
				} else {
					comments += line + "\n"
				}
			} else {
				break
			}
		}
		existingConfig = yamlconfig.YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &existingConfig)
		if err != nil {
			panic(err)
		}
	}
	if commentsEndpoint == "" {
		commentsEndpoint = endpointPath
	}

	newConfig := yamlconfig.MergeYamlConfig(config, existingConfig)

	var yamlBytes bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&yamlBytes)
	yamlEncoder.SetIndent(2)
	err = yamlEncoder.Encode(&newConfig)
	if err != nil {
		panic(err)
	}
	writeString := fmt.Sprintf("# %s %s\n%s%s", yamlconfig.EndpointToken, commentsEndpoint, comments, yamlBytes.String())
	os.WriteFile(outputPath, []byte(writeString), 0644)
}

func toStringSlice(i []interface{}) []string {
	ret := []string{}
	for _, v := range i {
		ret = append(ret, v.(string))
	}
	return ret
}

var jsonTypes = map[string]string{
	"integer": "Int64",
	"number":  "Float64",
	"boolean": "Bool",
	"string":  "String",
}

type parseUrlResult struct {
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
		ret.resultPath = strings.Join(parts[:len(parts)-1], "%v")
		ret.references = ids[:len(ids)-1]
	} else {
		// one to one
		ret.resultPath = strings.Join(parts, "%v")
		ret.references = ids
		ret.oneToOne = true
	}
	if strings.Contains(parts[0], "/organizations") {
		ret.category = "Organizations"
	} else if strings.Contains(parts[0], "/networks") {
		ret.category = "Networks"
	} else if strings.Contains(parts[0], "/devices") {
		ret.category = "Devices"
	}
	if len(parts) > 1 {
		if strings.Contains(parts[1], "/switch") {
			ret.category = "Switches"
		} else if strings.Contains(parts[1], "/wireless") {
			ret.category = "Wireless"
		}
	}
	return ret
}

func traverseProperties(m map[string]interface{}, path []string, gjsonPath string, exampleStr string, requiredProperties []string) []yamlconfig.YamlConfigAttribute {
	ret := []yamlconfig.YamlConfigAttribute{}

	keys := maps.Keys(m)
	sort.Strings(keys)

	for _, propName := range keys {
		propMap := m[propName].(map[string]interface{})
		if propMap["type"] != "object" && propMap["type"] != "array" {
			// primitive value
			attr := yamlconfig.YamlConfigAttribute{}
			attr.DataPath = path
			attr.Type = jsonTypes[propMap["type"].(string)]
			attr.ModelName = propName
			childGjsonPath := (gjsonPath + "." + propName)[1:]
			res := gjson.Get(exampleStr, childGjsonPath)
			attr.Example = res.String()
			if desc, ok := propMap["description"]; ok {
				attr.Description = sanitizeDescription(desc.(string))
			}
			if enums, ok := propMap["enum"]; ok {
				for _, r := range enums.([]interface{}) {
					attr.EnumValues = append(attr.EnumValues, r.(string))
				}
			}
			if min, ok := propMap["minimum"]; ok {
				if attr.Type == "Int64" {
					attr.MinInt = min.(int64)
				} else if attr.Type == "Float64" {
					attr.MinFloat = min.(float64)
				}
			}
			if max, ok := propMap["maximum"]; ok {
				if attr.Type == "Int64" {
					attr.MaxInt = max.(int64)
				} else if attr.Type == "Float64" {
					attr.MaxFloat = max.(float64)
				}
			}
			if slices.Contains(requiredProperties, propName) {
				attr.Mandatory = true
			}
			ret = append(ret, attr)
		}
	}
	for _, propName := range keys {
		propMap := m[propName].(map[string]interface{})
		if propMap["type"] == "object" {
			childPath := append(path, propName)
			childGjsonPath := gjsonPath + "." + propName
			childRequired := []string{}
			if rp, ok := propMap["required"]; ok {
				childRequired = toStringSlice(rp.([]interface{}))
			}
			children := traverseProperties(propMap["properties"].(map[string]interface{}), childPath, childGjsonPath, exampleStr, childRequired)
			ret = append(ret, children...)
		}
	}
	for _, propName := range keys {
		propMap := m[propName].(map[string]interface{})
		if propMap["type"] == "array" {
			attr := yamlconfig.YamlConfigAttribute{}
			attr.DataPath = path
			attr.Type = "List"
			attr.ModelName = propName
			if slices.Contains(requiredProperties, propName) {
				attr.Mandatory = true
			}
			items := propMap["items"].(map[string]interface{})
			if desc, ok := propMap["description"]; ok {
				attr.Description = sanitizeDescription(desc.(string))
			}
			if strings.Contains(attr.Description, "ordered array") {
				attr.OrderedList = true
			}
			if t, ok := jsonTypes[items["type"].(string)]; ok {
				attr.ElementType = t
				childGjsonPath := (gjsonPath + "." + propName + ".0")[1:]
				res := gjson.Get(exampleStr, childGjsonPath)
				attr.Example = res.String()
			} else if items["type"].(string) == "object" {
				childGjsonPath := gjsonPath + "." + propName + ".0"
				childRequired := []string{}
				if rp, ok := items["required"]; ok {
					childRequired = toStringSlice(rp.([]interface{}))
				}
				children := traverseProperties(items["properties"].(map[string]interface{}), []string{}, childGjsonPath, exampleStr, childRequired)
				attr.Attributes = children
			}
			ret = append(ret, attr)
		}
	}
	return ret
}

func sanitizeDescription(desc string) string {
	desc = strings.ReplaceAll(desc, "\n", " ")
	r := regexp.MustCompile("<.*?>")
	desc = r.ReplaceAllString(desc, "")
	desc = strings.ReplaceAll(desc, "'", "`")
	desc = strings.ReplaceAll(desc, "\"", "'")
	desc = strings.Join(strings.Fields(desc), " ") // Remove extra spaces
	return desc
}
