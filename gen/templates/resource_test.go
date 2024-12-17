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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMeraki{{camelCase .Name}}(t *testing.T) {
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
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue) (not .Computed)}}
	{{- if isNestedListSetMap .}}
	{{- $parent0 := .}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue) (not .Computed)}}
	{{- if isNestedListSetMap .}}
	{{- $parent1 := .}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue) (not .Computed)}}
	{{- if isNestedListSetMap .}}
	{{- $parent2 := .}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Value) (not .TestValue) (not .Computed) (not (isSet .))}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0 $parent1 $parent2}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0 $parent1 $parent2}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else if not (isSet .)}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0 $parent1}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0 $parent1}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
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
		checks = append(checks, resource.TestCheckResourceAttr("meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("meraki_{{snakeCase $name}}.test", "{{buildTestPath $parent0}}{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
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
		checks = append(checks, resource.TestCheckResourceAttr("meraki_{{snakeCase $name}}.test", "{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("meraki_{{snakeCase $name}}.test", "{{.TfName}}{{if isList .}}.0{{end}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}

	var steps []resource.TestStep
	{{- if not .SkipMinimumTest}}
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: {{if or .TestPrerequisites (len .TestVariables)}}testAccMeraki{{camelCase .Name}}PrerequisitesConfig+{{end}}testAccMeraki{{camelCase .Name}}Config_minimum(),
		})
	}
	{{- end}}
	steps = append(steps, resource.TestStep{
		Config: {{if or .TestPrerequisites (len .TestVariables)}}testAccMeraki{{camelCase .Name}}PrerequisitesConfig+{{end}}testAccMeraki{{camelCase .Name}}Config_all(),
		Check: resource.ComposeTestCheckFunc(checks...),
	})
	{{- if not .NoImport}}
	steps = append(steps, resource.TestStep{
		ResourceName: "meraki_{{snakeCase $name}}.test",
		ImportState: true,
		ImportStateVerify: true,
		ImportStateIdFunc: meraki{{camelCase .Name}}ImportStateIdFunc("meraki_{{snakeCase $name}}.test"),
		ImportStateVerifyIgnore: []string{ {{range getImportExcludes .Attributes}}"{{.}}",{{end}} },
		Check: resource.ComposeTestCheckFunc(checks...),
	})
	{{- end}}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func meraki{{camelCase .Name}}ImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary

		{{- range .Attributes}}
		{{- if or .Reference .Id}}
		{{toGoName .TfName}} := primary.Attributes["{{.TfName}}"]
		{{- end}}
		{{- end}}
		{{- if not (hasId .Attributes)}}
		id := primary.Attributes["id"]
		{{- end}}

		return fmt.Sprintf("{{range $i := (iterate (importParts .))}}{{if $i}},{{end}}%s{{end}}", {{$idRef := false}}{{range $i, $e := .Attributes}}{{if or .Reference .Id}}{{$idRef = true}}{{if $i}},{{end}}{{toGoName .TfName}}{{end}}{{end}}{{if not (hasId .Attributes)}}{{if $idRef}},{{end}}id{{end}}), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
{{- if or .TestPrerequisites (len .TestVariables)}}

const testAccMeraki{{camelCase .Name}}PrerequisitesConfig = `
{{- range .TestVariables}}
variable "{{.}}" {}
{{- end}}
{{.TestPrerequisites}}
`

{{- end}}
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMeraki{{camelCase .Name}}Config_minimum() string {
	config := `resource "meraki_{{snakeCase $name}}" "test" {` + "\n"
	{{- range  .Attributes}}
	{{- if or .Id .Reference .Mandatory .MinimumTestValue}}
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
		{{- if or .Id .Reference .Mandatory .MinimumTestValue}}
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
			{{- if or .Id .Reference .Mandatory .MinimumTestValue}}
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
				{{- if or .Id .Reference .Mandatory .MinimumTestValue}}
				{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `        {{.TfName}} = {{if .MinimumTestValue}}{{.MinimumTestValue}}{{else if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
				{{- else}}
	config += `        {{.TfName}} = {{if .MinimumTestValue}}{{.MinimumTestValue}}{{else if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
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
		config += `    {{.TfName}} = {{if .MinimumTestValue}}{{.MinimumTestValue}}{{else if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
			{{- else}}
	config += `    {{.TfName}} = {{if .MinimumTestValue}}{{.MinimumTestValue}}{{else if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
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
		config += `    {{.TfName}} = {{if .MinimumTestValue}}{{.MinimumTestValue}}{{else if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
			{{- else}}
	config += `    {{.TfName}} = {{if .MinimumTestValue}}{{.MinimumTestValue}}{{else if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
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
		config += `  {{.TfName}} = {{if .MinimumTestValue}}{{.MinimumTestValue}}{{else if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
	{{- else}}
	config += `  {{.TfName}} = {{if .MinimumTestValue}}{{.MinimumTestValue}}{{else if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMeraki{{camelCase .Name}}Config_all() string {
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
	return config
}

// End of section. //template:end testAccConfigAll
