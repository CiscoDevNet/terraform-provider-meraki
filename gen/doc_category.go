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
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-meraki/gen/yamlconfig"
)

const (
	definitionsPath = "./gen/definitions/"
)

var docPaths = []string{"./docs/data-sources/", "./docs/resources/"}

var extraDocs = map[string]string{}

func main() {
	var configs []yamlconfig.YamlConfig
	files, _ := os.ReadDir(definitionsPath)

	for _, filename := range files {
		path := filepath.Join(definitionsPath, filename.Name())
		bytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading file %q: %v", path, err)
		}

		config, err := yamlconfig.NewYamlConfig(bytes)
		if err != nil {
			log.Fatalf("Error parsing %q: %v", path, err)
		}
		configs = append(configs, config)
	}

	// Update doc category
	for i := range configs {
		for _, path := range docPaths {
			if (configs[i].NoDataSource && path == "./docs/data-sources/") ||
				(configs[i].NoResource && path == "./docs/resources/") {
				continue
			}
			filename := path + yamlconfig.SnakeCase(configs[i].Name) + ".md"
			content, err := os.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error opening documentation: %v", err)
			}

			cat := configs[i].DocCategory
			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+cat+`"`)

			os.WriteFile(filename, []byte(s), 0644)

			if configs[i].BulkDataSource && path == "./docs/data-sources/" {
				filename := path + yamlconfig.SnakeCase(configs[i].BulkName) + ".md"
				content, err := os.ReadFile(filename)
				if err != nil {
					log.Fatalf("Error opening documentation: %v", err)
				}

				cat := configs[i].DocCategory
				s := string(content)
				s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+cat+`"`)

				os.WriteFile(filename, []byte(s), 0644)
			}
		}
	}

	// Update extra doc categories
	for doc, cat := range extraDocs {
		for _, path := range docPaths {
			filename := path + doc + ".md"
			content, err := os.ReadFile(filename)
			if err == nil {
				s := string(content)
				s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+cat+`"`)

				os.WriteFile(filename, []byte(s), 0644)
			}
		}
	}
}
