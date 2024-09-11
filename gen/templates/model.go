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
	"github.com/netascode/terraform-provider-meraki/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type {{camelCase .Name}} struct {
	Id types.String `tfsdk:"id"`
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
type {{.GoTypeName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
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
{{- if isNestedListSet .}}
{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
type {{.GoTypeName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
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
{{- if isNestedListSet .}}
{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
type {{.GoTypeName}} struct {
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

func (data {{camelCase .Name}}) getPath() string {
	{{- if hasReference .Attributes}}
		return fmt.Sprintf("{{.RestEndpoint}}"{{range .Attributes}}{{if .Reference}}, url.QueryEscape(data.{{toGoName .TfName}}.Value{{.Type}}()){{end}}{{end}})
	{{- else}}
		return "{{.RestEndpoint}}"
	{{- end}}
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data {{camelCase .Name}}) toBody(ctx context.Context, state {{camelCase .Name}}) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	{{- range .Attributes}}
	{{- if .Value}}
	body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	{{- else if not .Reference}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if !data.{{toGoName .TfName}}.IsNull() {{if .WriteChangesOnly}}&& data.{{toGoName .TfName}} != state.{{toGoName .TfName}}{{end}} {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", data.{{toGoName .TfName}}.Value{{.Type}}())
	}
	{{- else if isListSet .}}
	if !data.{{toGoName .TfName}}.IsNull() {
		var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
	}
	{{- else if isNestedListSet .}}
	if len(data.{{toGoName .TfName}}) > 0 {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
		for _, item := range data.{{toGoName .TfName}} {
			itemBody := ""
			{{- range .Attributes}}
			{{- if .Value}}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
			{{- else if not .Reference}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
			if !item.{{toGoName .TfName}}.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", item.{{toGoName .TfName}}.Value{{.Type}}())
			}
			{{- else if isListSet .}}
			if !item.{{toGoName .TfName}}.IsNull() {
				var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
			}
			{{- else if isNestedListSet .}}
			if len(item.{{toGoName .TfName}}) > 0 {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
				for _, childItem := range item.{{toGoName .TfName}} {
					itemChildBody := ""
					{{- range .Attributes}}
					{{- if .Value}}
					itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
					{{- else if not .Reference}}
					{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
					if !childItem.{{toGoName .TfName}}.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", childItem.{{toGoName .TfName}}.Value{{.Type}}())
					}
					{{- else if isListSet .}}
					if !childItem.{{toGoName .TfName}}.IsNull() {
						var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
						childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
					}
					{{- else if isNestedListSet .}}
					if len(childItem.{{toGoName .TfName}}) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
						for _, childChildItem := range childItem.{{toGoName .TfName}} {
							itemChildChildBody := ""
							{{- range .Attributes}}
							{{- if .Value}}
							itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
							{{- else if not .Reference}}
							{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", childChildItem.{{toGoName .TfName}}.Value{{.Type}}())
							}
							{{- else if isListSet .}}
							if !childChildItem.{{toGoName .TfName}}.IsNull() {
								var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
								childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildChildBody)
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
					itemBody, _ = sjson.SetRaw(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildBody)
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
			body, _ = sjson.SetRaw(body, "{{range .DataPath}}{{.}}.{{end}}{{if .ModelName}}{{.ModelName}}.{{end}}-1", itemBody)
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
{{- define "fromBodyTemplate"}}
	{{- range .Attributes}}
	{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
	} else {
		{{- if .DefaultValue}}
		data.{{toGoName .TfName}} = types.{{.Type}}Value({{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}})
		{{- else}}
		data.{{toGoName .TfName}} = types.{{.Type}}Null()
		{{- end}}
	}
	{{- else if isListSet .}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
	}
	{{- else if isNestedListSet .}}
	if value := res{{if .ModelName}}.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"){{end}}; value.Exists() {
		data.{{toGoName .TfName}} = make([]{{.GoTypeName}}, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := {{.GoTypeName}}{}
			{{- template "fromBodyTemplate" .}}
			(*parent).{{toGoName .TfName}} = append((*parent).{{toGoName .TfName}}, data)
			return true
		})
	}
	{{- end}}
	{{- end}}
	{{- end}}
{{- end}}
{{- template "fromBodyTemplate" .}}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *{{camelCase .Name}}) fromBodyPartial(ctx context.Context, res gjson.Result) {
{{- define "fromBodyPartialTemplate"}}
	{{- range .Attributes}}
	{{- if and (not .Value) (not .WriteOnly) (not .Reference)}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
	} else {{if .DefaultValue}}if data.{{toGoName .TfName}}.Value{{.Type}}() != {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}} {{end}}{
		data.{{toGoName .TfName}} = types.{{.Type}}Null()
	}
	{{- else if isListSet .}}
	if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
		data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
	}
	{{- else if isNestedListSet .}}
	{{- $list := (toGoName .TfName)}}
	{{- if .OrderedList }}
	{
		l := len(res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").Array())
		tflog.Debug(ctx, fmt.Sprintf("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}} array resizing from %d to %d", len(data.{{toGoName .TfName}}), l))
		for i := len(data.{{toGoName .TfName}}); i < l; i++ {
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, {{.GoTypeName}}{})
		}
		if len(data.{{toGoName .TfName}}) > l {
			data.{{toGoName .TfName}} = data.{{toGoName .TfName}}[:l]
		}
	}
	for i := range data.{{toGoName .TfName}} {
		parent := &data
		data := (*parent).{{toGoName .TfName}}[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.%d", i))
	{{- else }}
	for i := 0; i < len(data.{{toGoName .TfName}}); i++ {
		keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value))}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
		keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value))}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

		parent := &data
		data := (*parent).{{toGoName .TfName}}[i]
		parentRes := &res
		var res gjson.Result

		parentRes.{{if .ModelName}}Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").{{end}}ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing {{toGoName .TfName}}[%d] = %+v",
				i,
				(*parent).{{toGoName .TfName}}[i],
			))
			(*parent).{{toGoName .TfName}} = slices.Delete((*parent).{{toGoName .TfName}}, i, i+1)
			i--

			continue
		}
	{{- end}}

		{{- template "fromBodyPartialTemplate" .}}
		(*parent).{{toGoName .TfName}}[i] = data
	}
	{{- end}}
	{{- end}}
	{{- end}}
{{- end}}
{{- template "fromBodyPartialTemplate" .}}
}

// End of section. //template:end fromBodyPartial

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
