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

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/netascode/terraform-provider-meraki/internal/provider/helpers"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &{{camelCase .Name}}DataSource{}
	_ datasource.DataSourceWithConfigure = &{{camelCase .Name}}DataSource{}
)

func New{{camelCase .Name}}DataSource() datasource.DataSource {
	return &{{camelCase .Name}}DataSource{}
}

type {{camelCase .Name}}DataSource struct {
	client *meraki.Client
}

func (d *{{camelCase .Name}}DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{snakeCase .Name}}"
}

{{- $nameQuery := .DataSourceNameQuery}}

func (d *{{camelCase .Name}}DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "{{.DsDescription}}",

		Attributes: map[string]schema.Attribute{
			"{{ .IdName }}": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				{{- if not .DataSourceNameQuery}}
				Required:            true,
				{{- else}}
				Optional:            true,
				Computed:            true,
				{{- end}}
			},
			{{- range  .Attributes}}
			{{- if not .Value}}
			"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
				MarkdownDescription: "{{.Description}}",
				{{- if isListSet .}}
				ElementType:         types.{{.ElementType}}Type,
				{{- end}}
				{{- if .Reference}}
				Required:            true,
				{{- else}}
				{{- if and (eq .ModelName "name") ($nameQuery)}}
				Optional:            true,
				{{- end}}
				Computed:            true,
				{{- end}}
				{{- if isNestedListMapSet .}}
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- range  .Attributes}}
						{{- if not .Value}}
						"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
							MarkdownDescription: "{{.Description}}",
							{{- if isListSet .}}
							ElementType:         types.{{.ElementType}}Type,
							{{- end}}
							Computed:            true,
							{{- if isNestedListMapSet .}}
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									{{- range  .Attributes}}
									{{- if not .Value}}
									"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
										MarkdownDescription: "{{.Description}}",
										{{- if isListSet .}}
										ElementType:         types.{{.ElementType}}Type,
										{{- end}}
										Computed:            true,
										{{- if isNestedListMapSet .}}
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												{{- range  .Attributes}}
												{{- if not .Value}}
												"{{.TfName}}": schema.{{if isNestedListMapSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
													MarkdownDescription: "{{.Description}}",
													{{- if isListSet .}}
													ElementType:         types.{{.ElementType}}Type,
													{{- end}}
													Computed:            true,
												},
												{{- end}}
												{{- end}}
											},
										},
										{{- end}}
									},
									{{- end}}
									{{- end}}
								},
							},
							{{- end}}
						},
						{{- end}}
						{{- end}}
					},
				},
				{{- end}}
			},
			{{- end}}
			{{- end}}
		},
	}
}

{{- if .DataSourceNameQuery}}
func (d *{{camelCase .Name}}DataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
    return []datasource.ConfigValidator{
        datasourcevalidator.ExactlyOneOf(
            path.MatchRoot("{{ .IdName }}"),
            path.MatchRoot("name"),
        ),
    }
}
{{- end}}

func (d *{{camelCase .Name}}DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *{{camelCase .Name}}DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config {{camelCase .Name}}

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.{{ toGoName .IdName }}.String()))

	var res gjson.Result
	var err error
	{{- if .DataSourceNameQuery}}
	if config.{{toGoName .IdName}}.IsNull() && !config.Name.IsNull() {
			res, err = d.client.Get(config.getPath())
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
				return
			}
			if len(res.Array()) > 0 {
				res.ForEach(func(k, v gjson.Result) bool {
					if config.Name.ValueString() == v.Get("name").String() {
						config.{{toGoName .IdName}} = types.StringValue(v.Get("{{dromedaryCase .IdName}}").String())
						tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.{{toGoName .IdName}}.String(), config.Name.ValueString(), config.{{toGoName .IdName}}.String()))
						res = v
						return false
					}
					return true
				})
			}

		if config.{{toGoName .IdName}}.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}
	{{- end}}

	if !res.Exists() {
		{{- if .GetFromAll}}
		res, err = d.client.Get(config.getPath())
		{{- else}}
		res, err = d.client.Get(config.getPath() + "/" + url.QueryEscape(config.{{toGoName .IdName}}.ValueString()))
		{{- end}}
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}
	{{- if .GetFromAll}}
	if len(res.Array()) > 0 {
		res.ForEach(func(k, v gjson.Result) bool {
			if config.{{toGoName .IdName}}.ValueString() == v.Get("{{dromedaryCase .IdName}}").String() {
				res = v
				return false
			}
			return true
		})
	}
	{{- end}}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.{{toGoName .IdName}}.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
