//go:build ignore
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

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceMeraki{{camelCase .BulkName}}(t *testing.T) {
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") == ""{{end}} {
        t.Skip("skipping test, set environment variable {{range $i, $e := .TestTags}}{{if $i}} and {{end}}{{$e}}{{end}}")
	}
	{{- end}}
	{{- if len .TestVariables}}
	if {{range $i, $e := .TestVariables}}{{if $i}} || {{end}}os.Getenv("TF_VAR_{{$e}}") == ""{{end}} {
        t.Skip("skipping test, set environment variable {{range $i, $e := .TestVariables}}{{if $i}} and {{end}}TF_VAR_{{$e}}{{end}}")
	}
	{{- end}}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: {{if or .TestPrerequisites (len .TestVariables)}}testAccDataSourceMeraki{{camelCase .BulkName}}PrerequisitesConfig+{{end}}testAccDataSourceMeraki{{camelCase .BulkName}}Config(),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
{{- if or .TestPrerequisites (len .TestVariables)}}

const testAccDataSourceMeraki{{camelCase .BulkName}}PrerequisitesConfig = `
{{- range .TestVariables}}
variable "{{.}}" {}
{{- end}}
{{.TestPrerequisites}}
`

{{- end}}
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMeraki{{camelCase .BulkName}}Config() string {
	config := `data "meraki_{{snakeCase .BulkName}}" "test" {
			{{- range  .Attributes}}
			{{- if .Reference}}
			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}
			{{- end}}
			{{- end}}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
