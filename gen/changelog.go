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
	"bytes"
	"log"
	"os"
	"text/template"
)

// gen-changelog renders CHANGELOG.md into the two guide copies that a full
// `make gen` would otherwise regenerate as a side effect of re-running codegen
// and tfplugindocs for every resource and data source. Use it to pick up a
// CHANGELOG.md edit without waiting on that full run.
const (
	changelogTemplate = "./gen/templates/changelog.md.tmpl"
	changelogOriginal = "./CHANGELOG.md"
)

var changelogLocations = []string{
	"./templates/guides/changelog.md.tmpl",
	"./docs/guides/changelog.md",
}

func main() {
	changelog, err := os.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading %q: %v", changelogOriginal, err)
	}

	tmplBytes, err := os.ReadFile(changelogTemplate)
	if err != nil {
		log.Fatalf("Error reading template %q: %v", changelogTemplate, err)
	}

	tmpl, err := template.New("changelog").Parse(string(tmplBytes))
	if err != nil {
		log.Fatalf("Error parsing template %q: %v", changelogTemplate, err)
	}

	output := new(bytes.Buffer)
	if err := tmpl.Execute(output, string(changelog)); err != nil {
		log.Fatalf("Error executing template %q: %v", changelogTemplate, err)
	}

	for _, location := range changelogLocations {
		if err := os.WriteFile(location, output.Bytes(), 0644); err != nil {
			log.Fatalf("Error writing %q: %v", location, err)
		}
	}
}
