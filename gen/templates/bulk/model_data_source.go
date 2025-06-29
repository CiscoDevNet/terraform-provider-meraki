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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSource{{camelCase .BulkName}} struct {
{{- range .Attributes}}
{{- if .Reference}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
	Items []DataSource{{camelCase .BulkName}}Items `tfsdk:"items"`
}

type DataSource{{camelCase .BulkName}}Items struct {
	Id types.String `tfsdk:"id"`
{{- range .Attributes}}
{{- if and (not .Reference) (not .Value)}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []DataSource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]DataSource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSetMap .}}
type DataSource{{.GoTypeBulkName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []DataSource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]DataSource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{end}}

{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSetMap .}}
{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSetMap .}}
type DataSource{{.GoTypeBulkName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []DataSource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]DataSource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{end}}

{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSetMap .}}
{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSetMap .}}
{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSetMap .}}
type DataSource{{.GoTypeBulkName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{end}}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSource{{camelCase .BulkName}}) getPath() string {
	{{- if hasReference .Attributes}}
		return fmt.Sprintf("{{.RestEndpoint}}"{{range .Attributes}}{{if .Reference}}, url.QueryEscape(data.{{toGoName .TfName}}.Value{{.Type}}()){{end}}{{end}})
	{{- else}}
		return "{{.RestEndpoint}}"
	{{- end}}
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSource{{camelCase .BulkName}}) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSource{{camelCase .BulkName}}Items, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSource{{camelCase .BulkName}}Items{}
		data.Id = types.StringValue(res.Get("{{if hasId .Attributes}}{{(getId .Attributes).ModelName}}{{else if .IdName}}{{.IdName}}{{else}}id{{end}}").String())
		{{- define "fromBodyTemplate"}}
			{{- range .Attributes}}
			{{- if and (not .Value) (not .WriteOnly) (not .Reference) .ModelName}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
			if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && value.Value() != nil {
				data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
			} else {
				{{- if .DefaultValue}}
				data.{{toGoName .TfName}} = types.{{.Type}}Value({{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}})
				{{- else}}
				data.{{toGoName .TfName}} = types.{{.Type}}Null()
				{{- end}}
			}
			{{- else if isListSet .}}
			if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && value.Value() != nil {
				data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
			} else {
				data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
			}
			{{- else if isNestedListSetMap .}}
			if value := res{{if .ModelName}}.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"){{end}}; value.Exists() && value.Value() != nil {
				{{- if isNestedMap .}}
				data.{{toGoName .TfName}} = make(map[string]DataSource{{.GoTypeBulkName}})
				{{- else}}
				data.{{toGoName .TfName}} = make([]DataSource{{.GoTypeBulkName}}, 0)
				{{- end}}
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DataSource{{.GoTypeBulkName}}{}
					{{- template "fromBodyTemplate" .}}
					{{- if isNestedMap .}}
					(*parent).{{toGoName .TfName}}[k.String()] = data
					{{- else}}
					(*parent).{{toGoName .TfName}} = append((*parent).{{toGoName .TfName}}, data)
					{{- end}}
					return true
				})
			}
			{{- end}}
			{{- end}}
			{{- end}}
		{{- end}}
		{{- template "fromBodyTemplate" .}}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody

{{- range .Attributes}}
	{{- range .Attributes}}
		{{- if .OrderedList }}
			{{- errorf "ordered_list not yet implemented at this depth"}}
		{{- end}}

		{{- range .Attributes}}
			{{- if .OrderedList }}
				{{- errorf "ordered_list not yet implemented at this depth"}}
			{{- end}}

			{{- range .Attributes}}
				{{- if .OrderedList }}
					{{- errorf "ordered_list not yet implemented at this depth"}}
				{{- end}}
				{{- range .Attributes}}
					{{- errorf "attributes not yet implemented at this depth"}}
				{{- end}}
			{{- end}}
		{{- end}}
	{{- end}}
{{- end}}