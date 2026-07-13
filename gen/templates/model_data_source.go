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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSource{{camelCase .Name}} - used only by the data source, always including every
// attribute (unlike model_resource.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model_resource.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model_resource.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSource{{camelCase .Name}} struct {
	Id types.String `tfsdk:"id"`
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []DataSource{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]DataSource{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- if and .Sensitive (eq .Type "String")}}
	{{toGoName .TfName}}Wo types.String `tfsdk:"{{.TfName}}_wo"`
	{{toGoName .TfName}}WoVersion types.Int64 `tfsdk:"{{.TfName}}_wo_version"`
{{- end}}
{{- end}}
{{- end}}
}

{{range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSetMap .}}
type DataSource{{.GoTypeName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []DataSource{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]DataSource{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- if and .Sensitive (eq .Type "String")}}
	{{toGoName .TfName}}Wo types.String `tfsdk:"{{.TfName}}_wo"`
	{{toGoName .TfName}}WoVersion types.Int64 `tfsdk:"{{.TfName}}_wo_version"`
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
type DataSource{{.GoTypeName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []DataSource{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else if isNestedMap .}}
	{{toGoName .TfName}} map[string]DataSource{{.GoTypeName}} `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- if and .Sensitive (eq .Type "String")}}
	{{toGoName .TfName}}Wo types.String `tfsdk:"{{.TfName}}_wo"`
	{{toGoName .TfName}}WoVersion types.Int64 `tfsdk:"{{.TfName}}_wo_version"`
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
type DataSource{{.GoTypeName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if and .Sensitive (eq .Type "String")}}
	{{toGoName .TfName}}Wo types.String `tfsdk:"{{.TfName}}_wo"`
	{{toGoName .TfName}}WoVersion types.Int64 `tfsdk:"{{.TfName}}_wo_version"`
{{- end}}
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

func (data DataSource{{camelCase .Name}}) getPath() string {
	{{- if hasReference .Attributes}}
		return fmt.Sprintf("{{.RestEndpoint}}"{{range .Attributes}}{{if .Reference}}, url.QueryEscape(data.{{toGoName .TfName}}.Value{{.Type}}()){{end}}{{end}})
	{{- else}}
		return "{{.RestEndpoint}}"
	{{- end}}
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSource{{camelCase .Name}}) fromBody(ctx context.Context, res meraki.Res) {
{{- define "dataSourceFromBodyTemplate"}}
	{{- range .Attributes}}
	{{- if and (not .Value) (not .WriteOnly) (not .Reference) .ModelName}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool")}}
	if value := res.Get("{{getFullModelName . false}}"); value.Exists() && value.Value() != nil {
		data.{{toGoName .TfName}} = types.{{.Type}}Value(value.{{if eq .Type "Int64"}}Int{{else if eq .Type "Float64"}}Float{{else}}{{.Type}}{{end}}())
	} else {
		{{- if .DefaultValue}}
		data.{{toGoName .TfName}} = types.{{.Type}}Value({{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}})
		{{- else}}
		data.{{toGoName .TfName}} = types.{{.Type}}Null()
		{{- end}}
	}
	{{- else if isListSet .}}
	if value := res.Get("{{getFullModelName . false}}"); value.Exists() && value.Value() != nil {
		data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
	}
	{{- else if isNestedListSetMap .}}
	if value := res{{if .ModelName}}.Get("{{getFullModelName . false}}"){{end}}; value.Exists() && value.Value() != nil {
		{{- if isNestedMap .}}
		data.{{toGoName .TfName}} = make(map[string]DataSource{{.GoTypeName}})
		{{- else}}
		data.{{toGoName .TfName}} = make([]DataSource{{.GoTypeName}}, 0)
		{{- end}}
		value.ForEach(func(k, res gjson.Result) bool {
			{{- range .Attributes}}
			{{- $fullModelName := getFullModelName . false}}
			{{- range .IgnoreImportValues}}
			if value := res.Get("{{$fullModelName}}"); value.String() == "{{.}}" {
				return true
			}
			{{- end}}
			{{- end}}
			parent := &data
			data := DataSource{{.GoTypeName}}{}
			{{- template "dataSourceFromBodyTemplate" .}}
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
{{- template "dataSourceFromBodyTemplate" .}}
}

// End of section. //template:end fromBody

{{- range .Attributes}}
	{{- range .Attributes}}
		{{- range .Attributes}}
			{{- range .Attributes}}
				{{- range .Attributes}}
					{{- errorf "attributes not yet implemented at this depth"}}
				{{- end}}
			{{- end}}
		{{- end}}
	{{- end}}
{{- end}}
