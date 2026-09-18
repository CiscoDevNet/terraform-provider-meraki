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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

// Actions require Terraform >= 1.14.0 (the `action` block and the ProviderWithActions
// interface are not available in older releases), hence the TerraformVersionChecks gate below.
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

	resource.Test(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_14_0),
		},
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: {{if or .TestPrerequisites (len .TestVariables)}}testAccMeraki{{camelCase .Name}}PrerequisitesConfig+{{end}}testAccMeraki{{camelCase .Name}}Config(),
			},
		},
	})
}

// End of section. //template:end testAcc

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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfig

// testAccMeraki{{camelCase .Name}}Config only sets the id/reference/mandatory top-level
// attributes. Nested (List/Set/Map) attributes are not yet supported in the generated test
// config; add them by hand outside this marked section if a future action needs them.
func testAccMeraki{{camelCase .Name}}Config() string {
	config := `action "meraki_{{snakeCase .Name}}" "test" {` + "\n"
	config += `  config {` + "\n"
	{{- range .Attributes}}
	{{- if and (or .Id .Reference .Mandatory) (not (isNestedListSetMap .))}}
	config += `    {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	{{- end}}
	{{- end}}
	config += `  }` + "\n"
	config += `}`
	return config
}

// End of section. //template:end testAccConfig
