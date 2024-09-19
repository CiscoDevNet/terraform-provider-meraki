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
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"

	"github.com/CiscoDevNet/terraform-provider-meraki/gen/yamlconfig"
)

const (
	definitionsPath   = "./gen/definitions/"
	providerTemplate  = "./gen/templates/provider.go"
	providerLocation  = "./internal/provider/provider.go"
	changelogTemplate = "./gen/templates/changelog.md.tmpl"
	changelogLocation = "./templates/guides/changelog.md.tmpl"
	changelogOriginal = "./CHANGELOG.md"
)

type t struct {
	path   string
	prefix string
	suffix string
}

var templates = []t{
	{
		path:   "./gen/templates/model.go",
		prefix: "./internal/provider/model_meraki_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source.go",
		prefix: "./internal/provider/data_source_meraki_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source_test.go",
		prefix: "./internal/provider/data_source_meraki_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/resource.go",
		prefix: "./internal/provider/resource_meraki_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/resource_test.go",
		prefix: "./internal/provider/resource_meraki_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/data-source.tf",
		prefix: "./examples/data-sources/meraki_",
		suffix: "/data-source.tf",
	},
	{
		path:   "./gen/templates/resource.tf",
		prefix: "./examples/resources/meraki_",
		suffix: "/resource.tf",
	},
	{
		path:   "./gen/templates/import.sh",
		prefix: "./examples/resources/meraki_",
		suffix: "/import.sh",
	},
}

func NewYamlConfig(bytes []byte) (yamlconfig.YamlConfig, error) {
	var config yamlconfig.YamlConfig

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return config, err
	}

	for i := range config.Attributes {
		if err := config.Attributes[i].Init(yamlconfig.CamelCase(config.Name)); err != nil {
			return yamlconfig.YamlConfig{}, err
		}
	}
	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the `%s` configuration.", config.Name)
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage the `%s` configuration.", config.Name)
	}
	if config.TfName == "" {
		config.TfName = strings.Replace(config.Name, " ", "_", -1)
	}
	if config.IdName == "" {
		config.IdName = "id"
	}

	return config, nil
}

func getTemplateSection(content, name string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	result := ""
	foundSection := false
	beginRegex := regexp.MustCompile(`\/\/template:begin\s` + name + `$`)
	endRegex := regexp.MustCompile(`\/\/template:end\s` + name + `$`)
	for scanner.Scan() {
		line := scanner.Text()
		if !foundSection {
			match := beginRegex.MatchString(line)
			if match {
				foundSection = true
				result += line + "\n"
			}
		} else {
			result += line + "\n"
			match := endRegex.MatchString(line)
			if match {
				foundSection = false
			}
		}
	}
	return result
}

func renderTemplate(templatePath, outputPath string, config interface{}) {
	file, err := os.Open(templatePath)
	if err != nil {
		log.Fatalf("Error opening template: %v", err)
	}
	defer file.Close()

	// skip first line with 'build-ignore' directive for go files
	scanner := bufio.NewScanner(file)
	if strings.HasSuffix(templatePath, ".go") {
		scanner.Scan()
	}
	var temp string
	for scanner.Scan() {
		temp = temp + scanner.Text() + "\n"
	}

	template, err := template.New(path.Base(templatePath)).Funcs(yamlconfig.Functions).Parse(temp)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template for %s: %v", outputPath, err)
	}

	outputFile := filepath.Join(outputPath)
	existingFile, err := os.Open(outputPath)
	if err != nil {
		os.MkdirAll(filepath.Dir(outputFile), 0755)
	} else if strings.HasSuffix(templatePath, ".go") {
		existingScanner := bufio.NewScanner(existingFile)
		var newContent string
		currentSectionName := ""
		beginRegex := regexp.MustCompile(`\/\/template:begin\s(.*?)$`)
		endRegex := regexp.MustCompile(`\/\/template:end\s(.*?)$`)
		for existingScanner.Scan() {
			line := existingScanner.Text()
			if currentSectionName == "" {
				matches := beginRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] != "" {
					currentSectionName = matches[1]
				} else {
					newContent += line + "\n"
				}
			} else {
				matches := endRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] == currentSectionName {
					currentSectionName = ""
					newSection := getTemplateSection(string(output.Bytes()), matches[1])
					newContent += newSection
				}
			}
		}
		output = bytes.NewBufferString(newContent)
	}
	// write to output file
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	f.Write(output.Bytes())
}

func updateDefinitions() {
	files, _ := os.ReadDir(definitionsPath)

	for _, filename := range files {
		path := filepath.Join(definitionsPath, filename.Name())
		fileBytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading file %q: %v", path, err)
		}

		commentsEndpoint := ""
		scanner := bufio.NewScanner(bytes.NewReader(fileBytes))
		for scanner.Scan() {
			line := scanner.Text()
			if line[0] == '#' {
				if strings.Contains(line, yamlconfig.EndpointToken) {
					commentsEndpoint = strings.Trim(strings.Split(line, yamlconfig.EndpointToken)[1], " ")
				}
			} else {
				break
			}
		}
		if commentsEndpoint == "" {
			continue
		}

		config, err := NewYamlConfig(fileBytes)
		if err != nil {
			log.Fatalf("Error parsing %q: %v", path, err)
		}
		cmd := exec.Command("go", "run", "gen/definition.go", commentsEndpoint, config.Name)
		if out, err := cmd.Output(); err != nil {
			log.Fatal(out, err)
		}
	}
}

func main() {
	var allFlag bool
	flag.BoolVar(&allFlag, "a", false, "Update all existing definitions from OpenAPI spec")
	flag.Parse()

	if allFlag {
		updateDefinitions()
		return
	}

	// Load configs
	var configs []yamlconfig.YamlConfig
	files, _ := os.ReadDir(definitionsPath)

	for _, filename := range files {
		path := filepath.Join(definitionsPath, filename.Name())
		bytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading file %q: %v", path, err)
		}

		config, err := NewYamlConfig(bytes)
		if err != nil {
			log.Fatalf("Error parsing %q: %v", path, err)
		}
		configs = append(configs, config)
	}

	var providerConfig []string
	for i := range configs {
		// Iterate over templates and render files
		for _, t := range templates {
			if (configs[i].NoImport && t.path == "./gen/templates/import.sh") ||
				(configs[i].NoDataSource && t.path == "./gen/templates/data_source.go") ||
				(configs[i].NoDataSource && t.path == "./gen/templates/data_source_test.go") ||
				(configs[i].NoDataSource && t.path == "./gen/templates/data-source.tf") ||
				(configs[i].NoResource && t.path == "./gen/templates/resource.go") ||
				(configs[i].NoResource && t.path == "./gen/templates/resource_test.go") ||
				(configs[i].NoResource && t.path == "./gen/templates/resource.tf") ||
				(configs[i].NoResource && t.path == "./gen/templates/import.sh") {
				continue
			}
			renderTemplate(t.path, t.prefix+yamlconfig.SnakeCase(configs[i].Name)+t.suffix, configs[i])
		}
		providerConfig = append(providerConfig, configs[i].Name)
	}

	// render provider.go
	renderTemplate(providerTemplate, providerLocation, configs)

	changelog, err := os.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading %q: %v", changelogOriginal, err)
	}
	renderTemplate(changelogTemplate, changelogLocation, string(changelog))
}
