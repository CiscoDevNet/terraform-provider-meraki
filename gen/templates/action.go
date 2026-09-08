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
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ action.Action              = &{{camelCase .Name}}Action{}
	_ action.ActionWithConfigure = &{{camelCase .Name}}Action{}
)

func New{{camelCase .Name}}Action() action.Action {
	return &{{camelCase .Name}}Action{}
}

type {{camelCase .Name}}Action struct {
	client *meraki.Client
}

func (a *{{camelCase .Name}}Action) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{snakeCase .Name}}"
}

// Schema defines the `config` block accepted by this action. Unlike a resource/data-source
// schema, there are no Computed/Default attributes and no plan modifiers: an action has no
// state or plan to diff against, it only ever reads `Config` once during Invoke.
func (a *{{camelCase .Name}}Action) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("{{.ActionDescription}}"){{if .EarlyAccess}}.AddEarlyAccessDescription(){{end}}.String,

		Attributes: map[string]schema.Attribute{
			{{- range .Attributes}}
			{{- if and (not .Value) (not .Computed) (not .DataSourceOnly)}}
			"{{.TfName}}": schema.{{if isNestedListSetMap .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else}}{{.Type}}{{end}}Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
					{{- if len .EnumValues -}}
					.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
					{{- end -}}
					.String,
				{{- if isListSet .}}
				ElementType:         types.{{.ElementType}}Type,
				{{- end}}
				{{- if or .Reference .Mandatory}}
				Required:            true,
				{{- else}}
				Optional:            true,
				{{- end}}
				{{- if .Sensitive}}
				Sensitive:           true,
				{{- end}}
				{{- if len .EnumValues}}
				Validators: []validator.String{
					stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
				},
				{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
				Validators: []validator.String{
					{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
					stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
					{{- end}}
					{{- range .StringPatterns}}
					stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
					{{- end}}
				},
				{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
				Validators: []validator.Int64{
					int64validator.Between({{.MinInt}}, {{.MaxInt}}),
				},
				{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
				Validators: []validator.Float64{
					float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
				},
				{{- end}}
				{{- if isNestedListSetMap .}}
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- range .Attributes}}
						{{- if and (not .Value) (not .Computed) (not .DataSourceOnly)}}
						"{{.TfName}}": schema.{{.Type}}Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}").String,
							{{- if .Mandatory}}
							Required:            true,
							{{- else}}
							Optional:            true,
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

func (a *{{camelCase .Name}}Action) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	a.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin invoke

// Invoke is a generic default: build the request body from Config, issue the request, and
// report any error. It has no way to write outputs back to Terraform (action.InvokeResponse
// carries only Diagnostics and SendProgress) and assumes the operation completes synchronously.
// Actions whose endpoint is asynchronous (returns a job id that must be polled for completion)
// need custom polling logic added by hand, outside of this marked section, since regeneration
// would otherwise overwrite it.
func (a *{{camelCase .Name}}Action) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config Action{{camelCase .Name}}

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Beginning Invoke")

	body := config.toBody(ctx)
	_, err := a.client.Post(config.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to invoke action, got error: %s", err))
		return
	}

	tflog.Debug(ctx, "Invoke finished successfully")
}

// End of section. //template:end invoke
