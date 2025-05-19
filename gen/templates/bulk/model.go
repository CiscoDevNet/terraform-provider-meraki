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

type {{camelCase .BulkName}} struct {
	Id types.String `tfsdk:"id"`
{{- range .Attributes}}
{{- if or .Reference (not .ModelName)}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
	Items []{{camelCase .BulkName}}Items `tfsdk:"items"`
}

type {{camelCase .BulkName}}Items struct {
	Id types.String `tfsdk:"id"`
{{- range .Attributes}}
{{- if and (not .Reference) (not .Value) .ModelName}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSetMap .}}
type {{.GoTypeBulkName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
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
type {{.GoTypeBulkName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
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
type {{.GoTypeBulkName}} struct {
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

func (data {{camelCase .BulkName}}) getPath() string {
	{{- if hasReference .Attributes}}
		return fmt.Sprintf("{{.RestEndpoint}}"{{range .Attributes}}{{if .Reference}}, url.QueryEscape(data.{{toGoName .TfName}}.Value{{.Type}}()){{end}}{{end}})
	{{- else}}
		return "{{.RestEndpoint}}"
	{{- end}}
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

{{if .BulkResource}}
func (data {{camelCase .Name}}) toBody(ctx context.Context, state {{camelCase .BulkName}}, id string) string {
	var item {{camelCase .BulkName}}Items
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	body := ""
	{{- range .Attributes}}
	{{- if or .Computed (not .ModelName)}}{{- continue}}{{- end}}
	{{- if .Value}}
	body, _ = sjson.Set(body, "{{getFullModelName . true}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	{{- else if not .Reference}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if !item.{{toGoName .TfName}}.IsNull() {{if .WriteChangesOnly}}&& item.{{toGoName .TfName}} != state.{{toGoName .TfName}}{{end}} {
		body, _ = sjson.Set(body, "{{getFullModelName . true}}", item.{{toGoName .TfName}}.Value{{.Type}}())
	}
	{{- else if isListSet .}}
	if !item.{{toGoName .TfName}}.IsNull() {
		var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
		item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{getFullModelName . true}}", values)
	}
	{{- else if isNestedListSetMap .}}
	{{if not .Mandatory}}if len(item.{{toGoName .TfName}}) > 0 {{end}}{
		{{- if isNestedMap .}}
		body, _ = sjson.Set(body, "{{getFullModelName . true}}", map[string]interface{}{})
		for key, item := range item.{{toGoName .TfName}} {
		{{- else}}
		body, _ = sjson.Set(body, "{{getFullModelName . true}}", []interface{}{})
		for _, item := range item.{{toGoName .TfName}} {
		{{- end}}
			itemBody := ""
			{{- range .Attributes}}
			{{- if or .Computed (not .ModelName)}}{{- continue}}{{- end}}
			{{- if .Value}}
			itemBody, _ = sjson.Set(itemBody, "{{getFullModelName . true}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
			{{- else if not .Reference}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
			if !item.{{toGoName .TfName}}.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "{{getFullModelName . true}}", item.{{toGoName .TfName}}.Value{{.Type}}())
			}
			{{- else if isListSet .}}
			if !item.{{toGoName .TfName}}.IsNull() {
				var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "{{getFullModelName . true}}", values)
			}
			{{- else if isNestedListSetMap .}}
			{{if not .Mandatory}}if len(item.{{toGoName .TfName}}) > 0 {{end}}{
				{{- if isNestedMap .}}
				itemBody, _ = sjson.Set(itemBody, "{{getFullModelName . true}}", map[string]interface{}{})
				for key, childItem := range item.{{toGoName .TfName}} {
				{{- else}}
				itemBody, _ = sjson.Set(itemBody, "{{getFullModelName . true}}", []interface{}{})
				for _, childItem := range item.{{toGoName .TfName}} {
				{{- end}}
					itemChildBody := ""
					{{- range .Attributes}}
					{{- if or .Computed (not .ModelName)}}{{- continue}}{{- end}}
					{{- if .Value}}
					itemChildBody, _ = sjson.Set(itemChildBody, "{{getFullModelName . true}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
					{{- else if not .Reference}}
					{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
					if !childItem.{{toGoName .TfName}}.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{getFullModelName . true}}", childItem.{{toGoName .TfName}}.Value{{.Type}}())
					}
					{{- else if isListSet .}}
					if !childItem.{{toGoName .TfName}}.IsNull() {
						var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
						childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "{{getFullModelName . true}}", values)
					}
					{{- else if isNestedListSetMap .}}
					{{if not .Mandatory}}if len(childItem.{{toGoName .TfName}}) > 0 {{end}}{
						{{- if isNestedMap .}}
						itemChildBody, _ = sjson.Set(itemChildBody, "{{getFullModelName . true}}", map[string]interface{}{})
						for key, childChildItem := range childItem.{{toGoName .TfName}} {
						{{- else}}
						itemChildBody, _ = sjson.Set(itemChildBody, "{{getFullModelName . true}}", []interface{}{})
						for _, childChildItem := range childItem.{{toGoName .TfName}} {
						{{- end}}
							itemChildChildBody := ""
							{{- range .Attributes}}
							{{- if or .Computed (not .ModelName)}}{{- continue}}{{- end}}
							{{- if .Value}}
							itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{getFullModelName . true}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
							{{- else if not .Reference}}
							{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{getFullModelName . true}}", childChildItem.{{toGoName .TfName}}.Value{{.Type}}())
							}
							{{- else if isListSet .}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() {
								var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
								childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{getFullModelName . true}}", values)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							{{- if isNestedMap .}}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "{{getFullModelName . true}}."+key, itemChildChildBody)
							{{- else}}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "{{getFullModelName . true}}.-1", itemChildChildBody)
							{{- end}}
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
					{{- if isNestedMap .}}
					itemBody, _ = sjson.SetRaw(itemBody, "{{getFullModelName . true}}."+key, itemChildBody)
					{{- else}}
					itemBody, _ = sjson.SetRaw(itemBody, "{{getFullModelName . true}}.-1", itemChildBody)
					{{- end}}
				}
			}
			{{- end}}
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
	{{- end}}
	return body
}
{{end}}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *{{camelCase .BulkName}}) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]{{camelCase .BulkName}}Items, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := {{camelCase .BulkName}}Items{}
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
				data.{{toGoName .TfName}} = make(map[string]{{.GoTypeBulkName}})
				{{- else}}
				data.{{toGoName .TfName}} = make([]{{.GoTypeBulkName}}, 0)
				{{- end}}
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := {{.GoTypeBulkName}}{}
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
