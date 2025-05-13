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

func TestAccDataSourceMeraki{{camelCase .Name}}(t *testing.T) {
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
	var checks []resource.TestCheckFunc
	{{- $name := .Name }}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue) (not .Computed) .ModelName}}
	{{- if isNestedListSetMap .}}
	{{- $parent0 := .}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue) (not .Computed) .ModelName}}
	{{- if isNestedListSetMap .}}
	{{- $parent1 := .}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue) (not .Computed) .ModelName}}
	{{- if isNestedListSetMap .}}
	{{- $parent2 := .}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue) (not .Computed) .ModelName (not (isSet .))}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0 $parent1 $parent2}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0 $parent1 $parent2}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else if not (isSet .)}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0 $parent1}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0 $parent1}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else if not (isSet .)}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else if not (isSet .)}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.meraki_{{snakeCase $name}}.test", "{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_{{snakeCase $name}}.test", "{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: {{if or .TestPrerequisites (len .TestVariables)}}testAccDataSourceMeraki{{camelCase .Name}}PrerequisitesConfig+{{end}}testAccDataSourceMeraki{{camelCase .Name}}Config(),
				Check: resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
{{- if or .TestPrerequisites (len .TestVariables)}}

const testAccDataSourceMeraki{{camelCase .Name}}PrerequisitesConfig = `
{{- range .TestVariables}}
variable "{{.}}" {}
{{- end}}
{{.TestPrerequisites}}
`

{{- end}}
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMeraki{{camelCase .Name}}Config() string {
	config := `resource "meraki_{{snakeCase $name}}" "test" {` + "\n"
	{{- range  .Attributes}}
	{{- if and (not .ExcludeTest) (not .Value) (not .Computed)}}
	{{- if isNestedListSetMap .}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- if isNestedMap .}}
	config += `  {{.TfName}} = {"{{.MapKeyExample}}" = {` + "\n"
	{{- else}}
	config += `  {{.TfName}} = [{` + "\n"
	{{- end}}
		{{- range  .Attributes}}
		{{- if and (not .ExcludeTest) (not .Value) (not .Computed)}}
		{{- if isNestedListSetMap .}}
		{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		{{- end}}
		{{- if isNestedMap .}}
	config += `    {{.TfName}} = {"{{.MapKeyExample}}" = {` + "\n"
		{{- else}}
	config += `    {{.TfName}} = [{` + "\n"
		{{- end}}
			{{- range  .Attributes}}
			{{- if and (not .ExcludeTest) (not .Value) (not .Computed)}}
			{{- if isNestedListSetMap .}}
			{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
			{{- end}}
			{{- if isNestedMap .}}
	config += `      {{.TfName}} = {"{{.MapKeyExample}}" = {` + "\n"
			{{- else}}
	config += `      {{.TfName}} = [{` + "\n"
			{{- end}}
				{{- range  .Attributes}}
				{{- if and (not .ExcludeTest) (not .Value) (not .Computed)}}
				{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `        {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
				{{- else}}
	config += `        {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
				{{- end}}
				{{- end}}
				{{- end}}
			{{- if isNestedMap .}}
	config += `      }}` + "\n"
			{{- else}}
	config += `      }]` + "\n"
			{{- end}}
			{{- if len .TestTags}}
	}
			{{- end}}
			{{- else}}
			{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `      {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
			{{- else}}
	config += `      {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
			{{- end}}
			{{- end}}
			{{- end}}
			{{- end}}
		{{- if isNestedMap .}}
	config += `    }}` + "\n"
		{{- else}}
	config += `    }]` + "\n"
		{{- end}}
		{{- if len .TestTags}}
	}
		{{- end}}
		{{- else}}
		{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `    {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
			{{- else}}
	config += `    {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
		{{- end}}
		{{- end}}
		{{- end}}
		{{- end}}
	{{- if isNestedMap .}}
	config += `  }}` + "\n"
	{{- else}}
	config += `  }]` + "\n"
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `  {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
	{{- else}}
	config += `  {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	config += `}` + "\n"

	config += `
		data "meraki_{{snakeCase .Name}}" "test" {
			{{- if not .PutCreate}}
			id = meraki_{{snakeCase $name}}.test.id
			{{- end}}
			{{- range  .Attributes}}
			{{- if .Reference}}
			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}
			{{- end}}
			{{- end}}
			depends_on = [meraki_{{snakeCase $name}}.test]
		}
	`
	return config
}

{{if .DataSourceNameQuery -}}
func testAccNamedDataSourceMeraki{{camelCase .Name}}Config() string {
	config := `resource "meraki_{{snakeCase $name}}" "test" {` + "\n"
	{{- range  .Attributes}}
	{{- if and (not .ExcludeTest) (not .Value) (not .Computed)}}
	{{- if isNestedListSetMap .}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- if isNestedMap .}}
	config += `  {{.TfName}} = {"{{.MapKeyExample}}" = {` + "\n"
	{{- else}}
	config += `  {{.TfName}} = [{` + "\n"
	{{- end}}
		{{- range  .Attributes}}
		{{- if and (not .ExcludeTest) (not .Value) (not .Computed)}}
		{{- if isNestedListSetMap .}}
		{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		{{- end}}
		{{- if isNestedMap .}}
	config += `    {{.TfName}} = {"{{.MapKeyExample}}" = {` + "\n"
		{{- else}}
	config += `    {{.TfName}} = [{` + "\n"
		{{- end}}
			{{- range  .Attributes}}
			{{- if and (not .ExcludeTest) (not .Value) (not .Computed)}}
			{{- if isNestedListSetMap .}}
			{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
			{{- end}}
			{{- if isNestedMap .}}
	config += `      {{.TfName}} = {"{{.MapKeyExample}}" = {` + "\n"
			{{- else}}
	config += `      {{.TfName}} = [{` + "\n"
			{{- end}}
				{{- range  .Attributes}}
				{{- if and (not .ExcludeTest) (not .Value) (not .Computed)}}
				{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `        {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
				{{- else}}
	config += `        {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
				{{- end}}
				{{- end}}
				{{- end}}
			{{- if isNestedMap .}}
	config += `      }}` + "\n"
			{{- else}}
	config += `      }]` + "\n"
			{{- end}}
			{{- if len .TestTags}}
	}
			{{- end}}
			{{- else}}
			{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `      {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
			{{- else}}
	config += `      {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
			{{- end}}
			{{- end}}
			{{- end}}
			{{- end}}
		{{- if isNestedMap .}}
	config += `    }}` + "\n"
		{{- else}}
	config += `    }]` + "\n"
		{{- end}}
		{{- if len .TestTags}}
	}
		{{- end}}
		{{- else}}
		{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `    {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
			{{- else}}
	config += `    {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
		{{- end}}
		{{- end}}
		{{- end}}
		{{- end}}
	{{- if isNestedMap .}}
	config += `  }}` + "\n"
	{{- else}}
	config += `  }]` + "\n"
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `  {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
	{{- else}}
	config += `  {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	config += `}` + "\n"

	config += `
		data "meraki_{{snakeCase .Name}}" "test" {
			name = meraki_{{snakeCase $name}}.test.name
			{{- range  .Attributes}}
			{{- if .Reference}}
			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}
			{{- end}}
			{{- end}}
			depends_on = [meraki_{{snakeCase $name}}.test]
		}
	`
	return config
}
{{- end}}

// End of section. //template:end testAccDataSourceConfig
