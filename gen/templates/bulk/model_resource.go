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
	"slices"
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type Resource{{camelCase .BulkName}} struct {
	Id types.String `tfsdk:"id"`
{{- if not (hasOrganizationId .)}}
	OrganizationId types.String `tfsdk:"organization_id"`
{{- end}}
{{- range getBulkParentAttributes .}}
{{- if (not .ModelName)}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
	Items []Resource{{camelCase .BulkName}}Items `tfsdk:"items"`
}

type Resource{{camelCase .BulkName}}Items struct {
{{- if not .PutCreate}}
	Id types.String `tfsdk:"id"`
{{- end}}
{{- range getBulkItemAttributes .}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []Resource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]Resource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSetMap .}}
type Resource{{.GoTypeBulkName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []Resource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]Resource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
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
type Resource{{.GoTypeBulkName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []Resource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]Resource{{.GoTypeBulkName}} `tfsdk:"{{.TfName}}"`
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
type Resource{{.GoTypeBulkName}} struct {
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

type Resource{{camelCase .BulkName}}Identity struct {
{{- range (getBulkImportAttributes .)}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
    ItemIds types.List `tfsdk:"item_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data Resource{{camelCase .BulkName}}) getPath() string {
	{{- if hasReference .Attributes}}
		return fmt.Sprintf("{{getBulkPath .RestEndpoint}}"{{range getBulkParentAttributes .}}, url.QueryEscape(data.{{toGoName .TfName}}.Value{{.Type}}()){{end}})
	{{- else}}
		return "{{.RestEndpoint}}"
	{{- end}}
}

{{if .PutCreate}}
func (data Resource{{camelCase .BulkName}}) getItemPath(id string) string {
	{{- if hasReference .Attributes}}
		return fmt.Sprintf("{{.RestEndpoint}}"{{range getBulkParentAttributes .}}, url.QueryEscape(data.{{toGoName .TfName}}.Value{{.Type}}()){{end}}, id)
	{{- else}}
		return "{{.RestEndpoint}}"
	{{- end}}
}
{{end}}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

{{if .BulkResource}}
func (data Resource{{camelCase .BulkName}}Items) toBody(ctx context.Context, state Resource{{camelCase .BulkName}}Items) string {
	body := ""
	{{- range getBulkItemAttributes .}}
	{{- if or .Computed (not .ModelName)}}{{- continue}}{{- end}}
	{{- if .Value}}
	body, _ = sjson.Set(body, "{{getFullModelName . true}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	{{- else if not .Reference}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if !data.{{toGoName .TfName}}.IsNull() {{if .WriteChangesOnly}}&& data.{{toGoName .TfName}} != state.{{toGoName .TfName}}{{end}} {
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

func (data *Resource{{camelCase .BulkName}}) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]Resource{{camelCase .BulkName}}Items, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := Resource{{camelCase .BulkName}}Items{}
		{{- define "fromBodyTemplate"}}
			{{- range .Attributes}}
			{{- if and (not .Value) (not .WriteOnly) .ModelName}}
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
				data.{{toGoName .TfName}} = make(map[string]Resource{{.GoTypeBulkName}})
				{{- else}}
				data.{{toGoName .TfName}} = make([]Resource{{.GoTypeBulkName}}, 0)
				{{- end}}
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := Resource{{.GoTypeBulkName}}{}
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
	{{- if not .PutCreate}}
	index := 0
	res.ForEach(func(k, res gjson.Result) bool {
		data.Items[index].Id = types.StringValue(res.Get("{{.IdName}}").String())
		index++
		return true
	})
	{{- end}}
	data.Id = data.OrganizationId
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *Resource{{camelCase .BulkName}}) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("{{.IdName}}").String() == (*parent).Items[i].{{getBulkItemId .}}.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
	{{- define "fromBodyPartialTemplate"}}
		{{- range .Attributes}}
		{{- if and (not .Value) (not .WriteOnly) (not .Reference) .ModelName}}
		{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
		if value := res.Get("{{getFullModelName . false}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
			data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
		} else {{if .DefaultValue}}if data.{{toGoName .TfName}}.Value{{.Type}}() != {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}} {{end}}{
			data.{{toGoName .TfName}} = types.{{.Type}}Null()
		}
		{{- else if isListSet .}}
		if value := res.Get("{{getFullModelName . false}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
			data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
		} else {
			data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
		}
		{{- else if isNestedListSetMap .}}
		{{- $list := (toGoName .TfName)}}
		{{- if .OrderedList }}
		{
			l := len(res.Get("{{getFullModelName . false}}").Array())
			tflog.Debug(ctx, fmt.Sprintf("{{getFullModelName . false}} array resizing from %d to %d", len(data.{{toGoName .TfName}}), l))
			if len(data.{{toGoName .TfName}}) > l {
				data.{{toGoName .TfName}} = data.{{toGoName .TfName}}[:l]
			}
		}
		for i := range data.{{toGoName .TfName}} {
			parent := &data
			data := (*parent).{{toGoName .TfName}}[i]
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("{{getFullModelName . false}}.%d", i))
		{{- else if isNestedMap .}}
		for i, item := range data.{{toGoName .TfName}} {
			parent := &data
			data := item
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("{{getFullModelName . false}}.%s", i))
		{{- else }}
		for i := 0; i < len(data.{{toGoName .TfName}}); i++ {
			keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value) (not .WriteOnly))}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{getFullModelName . false}}", {{end}}{{end}}{{end}} }
			keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value) (not .WriteOnly))}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

			parent := &data
			data := (*parent).{{toGoName .TfName}}[i]
			parentRes := &res
			var res gjson.Result

			parentRes.{{if .ModelName}}Get("{{getFullModelName . false}}").{{end}}ForEach(
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
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyPartial(), removing item with id: %s", data.Items[toBeDeleted[i]].{{getBulkItemId .}}.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *Resource{{camelCase .BulkName}}) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
	{{- if hasComputedAttributes .Attributes}}
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("{{.IdName}}").String() == (*parent).Items[i].{{getBulkItemId .}}.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
	{{- end}}
	{{- define "fromBodyUnknownsTemplate"}}
		{{- range .Attributes}}
		{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")) .Computed}}
		if data.{{toGoName .TfName}}.IsUnknown() {
			if value := res.Get("{{getFullModelName . false}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
				data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
			} else {
				data.{{toGoName .TfName}} = types.{{.Type}}Null()
			}
		}
		{{- else if and (isListSet .) .Computed}}
		if data.{{toGoName .TfName}}.IsUnknown() {
			if value := res.Get("{{getFullModelName . false}}"); value.Exists() && !data.{{toGoName .TfName}}.IsNull() {
				data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
			} else {
				data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
			}
		}
		{{- else if and (isNestedListSetMap .) (hasComputedAttributes .Attributes)}}
		{{- $list := (toGoName .TfName)}}
		{{- if .OrderedList }}
		{
			l := len(res.Get("{{getFullModelName . false}}").Array())
			tflog.Debug(ctx, fmt.Sprintf("{{getFullModelName . false}} array resizing from %d to %d", len(data.{{toGoName .TfName}}), l))
			if len(data.{{toGoName .TfName}}) > l {
				data.{{toGoName .TfName}} = data.{{toGoName .TfName}}[:l]
			}
		}
		for i := range data.{{toGoName .TfName}} {
			parent := &data
			data := (*parent).{{toGoName .TfName}}[i]
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("{{getFullModelName . false}}.%d", i))
		{{- else if isNestedMap .}}
		for i, item := range data.{{toGoName .TfName}} {
			parent := &data
			data := item
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("{{getFullModelName . false}}.%s", i))
		{{- else }}
		for i := 0; i < len(data.{{toGoName .TfName}}); i++ {
			keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value) (not .WriteOnly))}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{getFullModelName . false}}", {{end}}{{end}}{{end}} }
			keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id (and $noId (not .Value) (not .WriteOnly))}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }

			parent := &data
			data := (*parent).{{toGoName .TfName}}[i]
			parentRes := &res
			var res gjson.Result

			parentRes.{{if .ModelName}}Get("{{getFullModelName . false}}").{{end}}ForEach(
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

			{{- template "fromBodyUnknownsTemplate" .}}
			(*parent).{{toGoName .TfName}}[i] = data
		}
		{{- end}}
		{{- end}}
	{{- end}}
	{{- template "fromBodyUnknownsTemplate" .}}
	{{- if hasComputedAttributes .Attributes}}
	}
	{{- end}}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *Resource{{camelCase .BulkName}}) fromBodyImport(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("{{.IdName}}").String() == (*parent).Items[i].{{getBulkItemId .}}.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		{{- define "fromBodyImportTemplate"}}
			{{- range .Attributes}}
			{{- if and (not .Value) (not .WriteOnly) .ModelName}}
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
			if value := res.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
				data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
			} else {
				data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
			}
			{{- else if isNestedListSetMap .}}
			if value := res{{if .ModelName}}.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"){{end}}; value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
				{{- if isNestedMap .}}
				data.{{toGoName .TfName}} = make(map[string]Resource{{.GoTypeBulkName}})
				{{- else}}
				data.{{toGoName .TfName}} = make([]Resource{{.GoTypeBulkName}}, 0)
				{{- end}}
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := Resource{{.GoTypeBulkName}}{}
					{{- template "fromBodyImportTemplate" .}}
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
		{{- template "fromBodyImportTemplate" .}}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyImport(), removing item with id: %s", data.Items[toBeDeleted[i]].{{getBulkItemId .}}.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *Resource{{camelCase .BulkName}}Identity) toIdentity(ctx context.Context, plan *Resource{{camelCase .BulkName}}) {
	{{- range (getBulkImportAttributes .)}}
	data.{{toGoName .TfName}} = plan.{{toGoName .TfName}}
	{{- end}}
}
	
// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *Resource{{camelCase .BulkName}}) fromIdentity(ctx context.Context, identity *Resource{{camelCase .BulkName}}Identity) {
	{{- range (getBulkImportAttributes .)}}
	data.{{toGoName .TfName}} = identity.{{toGoName .TfName}}
	{{- end}}
}
	
// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data Resource{{camelCase .BulkName}}) toDestroyBody(ctx context.Context) string {
	body := ""
	{{- range getBulkItemAttributes .}}
	{{- if .DestroyValue}}
	body, _ = sjson.Set(body, "{{getFullModelName . true}}", {{.DestroyValue}})
	{{- end}}
	{{- end}}
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *Resource{{camelCase .BulkName}}) hasChanges(ctx context.Context, state *Resource{{camelCase .BulkName}}, id string) bool {
	hasChanges := false

	item := Resource{{camelCase .BulkName}}Items{}
	for i := range data.Items {
		if data.Items[i].{{getBulkItemId .}}.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := Resource{{camelCase .BulkName}}Items{}
	for i := range state.Items {
		if state.Items[i].{{getBulkItemId .}}.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	{{- range .Attributes}}
	{{- $name := toGoName .TfName}}
	{{- if and (not .Value) .ModelName}}
	{{- if isNestedMap .}}
	if !maps.Equal(item.{{toGoName .TfName}}, stateItem.{{toGoName .TfName}}) {
		hasChanges = true
	}
	{{- else if not (isNestedListSet .)}}
	if !item.{{toGoName .TfName}}.Equal(stateItem.{{toGoName .TfName}}) {
		hasChanges = true
	}
	{{- else}}
	if len(item.{{toGoName .TfName}}) != len(stateItem.{{toGoName .TfName}}) {
		hasChanges = true
	} else {
		for i := range item.{{toGoName .TfName}} {
			{{- range .Attributes}}
			{{- $cname := toGoName .TfName}}
			{{- if and (not .Value) .ModelName}}
			{{- if isNestedMap .}}
			if !maps.Equal(item.{{$name}}[i].{{toGoName .TfName}}, stateItem.{{$name}}[i].{{toGoName .TfName}}) {
				hasChanges = true
			}
			{{- else if not (isNestedListSet .)}}
			if !item.{{$name}}[i].{{toGoName .TfName}}.Equal(stateItem.{{$name}}[i].{{toGoName .TfName}}) {
				hasChanges = true
			}
			{{- else}}
			if len(item.{{$name}}[i].{{toGoName .TfName}}) != len(stateItem.{{$name}}[i].{{toGoName .TfName}}) {
				hasChanges = true
			} else {
				for ii := range item.{{$name}}[i].{{toGoName .TfName}} {
					{{- range .Attributes}}
					{{- $ccname := toGoName .TfName}}
					{{- if and (not .Value) .ModelName}}
					{{- if isNestedMap .}}
					if !maps.Equal(item.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}, stateItem.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) {
						hasChanges = true
					}
					{{- else if not (isNestedListSet .)}}
					if !item.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}.Equal(stateItem.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) {
						hasChanges = true
					}
					{{- else}}
					if len(item.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) != len(stateItem.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) {
						hasChanges = true
					} else {
						for iii := range item.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} {
							{{- range .Attributes}}
							{{- if and (not .Value) .ModelName}}
							{{- if isNestedMap .}}
							if !maps.Equal(item.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}, stateItem.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}) {
								hasChanges = true
							}
							{{- else}}
							if !item.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}.Equal(stateItem.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}) {
								hasChanges = true
							}
							{{- end}}
							{{- end}}
							{{- end}}
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return hasChanges
}

// End of section. //template:end hasChanges

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
