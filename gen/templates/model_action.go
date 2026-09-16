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
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

// Action{{camelCase .Name}} models the `config` block of an action invocation. Unlike a
// resource/data-source model, this is never persisted: it exists only for the duration of
// a single Invoke call, so it has no computed/read-only fields and no identity.
type Action{{camelCase .Name}} struct {
{{- range .Attributes}}
{{- if and (not .Value) (not .Computed) (not .DataSourceOnly)}}
{{- if isNestedListSetMap .}}
	{{toGoName .TfName}} []{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{range .Attributes}}
{{- if and (not .Value) (not .Computed) (not .DataSourceOnly)}}
{{- if isNestedListSetMap .}}
type {{.GoTypeName}} struct {
{{- range .Attributes}}
{{- if and (not .Value) (not .Computed) (not .DataSourceOnly)}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{end}}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data Action{{camelCase .Name}}) getPath() string {
	{{- if hasReference .Attributes}}
		return fmt.Sprintf("{{.RestEndpoint}}"{{range .Attributes}}{{if .Reference}}, url.QueryEscape(data.{{toGoName .TfName}}.Value{{.Type}}()){{end}}{{end}})
	{{- else}}
		return "{{.RestEndpoint}}"
	{{- end}}
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data Action{{camelCase .Name}}) toBody(ctx context.Context) string {
	body := ""
	{{- range .Attributes}}
	{{- if or .Computed .DataSourceOnly .Reference (not .ModelName)}}{{- continue}}{{- end}}
	{{- if .Value}}
	body, _ = sjson.Set(body, "{{getFullModelName . true}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	{{- else if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if !data.{{toGoName .TfName}}.IsNull() {
		body, _ = sjson.Set(body, "{{getFullModelName . true}}", data.{{toGoName .TfName}}.Value{{.Type}}())
	}
	{{- else if isListSet .}}
	if !data.{{toGoName .TfName}}.IsNull() {
		var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{getFullModelName . true}}", values)
	}
	{{- else if isNestedListSetMap .}}
	{{if not .Mandatory}}if len(data.{{toGoName .TfName}}) > 0 {{end}}{
		{{- if isNestedMap .}}
		body, _ = sjson.Set(body, "{{getFullModelName . true}}", map[string]interface{}{})
		for key, item := range data.{{toGoName .TfName}} {
		{{- else}}
		body, _ = sjson.Set(body, "{{getFullModelName . true}}", []interface{}{})
		for _, item := range data.{{toGoName .TfName}} {
		{{- end}}
			itemBody := ""
			{{- range .Attributes}}
			{{- if or .Computed .DataSourceOnly (not .ModelName)}}{{- continue}}{{- end}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
			if !item.{{toGoName .TfName}}.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "{{getFullModelName . true}}", item.{{toGoName .TfName}}.Value{{.Type}}())
			}
			{{- end}}
			{{- end}}
			{{- if isNestedMap .}}
			body, _ = sjson.SetRaw(body, "{{getFullModelName . true}}."+key, itemBody)
			{{- else}}
			body, _ = sjson.SetRaw(body, "{{getFullModelName . true}}.-1", itemBody)
			{{- end}}
		}
	}
	{{- end}}
	{{- end}}
	return body
}

// End of section. //template:end toBody

{{- range .Attributes}}
	{{- range .Attributes}}
		{{- range .Attributes}}
			{{- errorf "nested attributes deeper than one level are not yet implemented for actions"}}
		{{- end}}
	{{- end}}
{{- end}}
