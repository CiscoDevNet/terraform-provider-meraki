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
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &{{camelCase .BulkName}}Resource{}
	{{- if not .NoImport}}
	_ resource.ResourceWithImportState = &{{camelCase .BulkName}}Resource{}
	{{- end}}
	_ resource.ResourceWithModifyPlan  = &{{camelCase .BulkName}}Resource{}
)

func New{{camelCase .BulkName}}Resource() resource.Resource {
	return &{{camelCase .BulkName}}Resource{}
}

type {{camelCase .BulkName}}Resource struct {
	client *meraki.Client
}

func (r *{{camelCase .BulkName}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{snakeCase .BulkName}}"
}

func (r *{{camelCase .BulkName}}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("{{.ResBulkDescription}}").AddBulkResourceIds({{range .Attributes}}{{if .BulkId}}"{{.TfName}}", {{end}}{{end}}){{if .EarlyAccess}}.AddEarlyAccessDescription(){{end}}.String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			{{- if not (hasOrganizationId .)}}
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "The organization ID",
				Required:            true,
			},
			{{- end}}
			{{- range getBulkParentAttributes .}}
			{{- if not .ModelName}}
			"{{.TfName}}": schema.{{.Type}}Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
					{{- if len .EnumValues -}}
					.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
					{{- end -}}
					{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
					.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
					{{- end -}}
					{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
					.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
					{{- end -}}
					{{- if .DefaultValue -}}
					.AddDefaultValueDescription("{{.DefaultValue}}")
					{{- end -}}
					.String,
				{{- if or .Reference .Mandatory}}
				Required:            true,
				{{- else if not .Computed}}
				Optional:            true,
				{{- end}}
				{{- if or (len .DefaultValue) .Computed}}
				Computed:            true,
				{{- end}}
				{{- if len .EnumValues}}
				Validators: []validator.String{
					stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
				},
				{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
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
				{{- if and (len .DefaultValue) (eq .Type "Int64")}}
				Default:             int64default.StaticInt64({{.DefaultValue}}),
				{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
				Default:             booldefault.StaticBool({{.DefaultValue}}),
				{{- else if and (len .DefaultValue) (eq .Type "String")}}
				Default:             stringdefault.StaticString("{{.DefaultValue}}"),
				{{- end}}
				{{- if .Computed}}
				PlanModifiers: []planmodifier.{{.Type}}{
					{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
				},
				{{- end}}
			},
			{{- end}}
			{{- end}}
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The list of items",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- if not .PutCreate}}
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the item",
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						{{- end}}
						{{- range getBulkItemAttributes .}}
						{{- if not .Value}}
						"{{.TfName}}": schema.{{if isNestedListSetMap .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
								{{- if len .EnumValues -}}
								.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
								{{- end -}}
								{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
								.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
								{{- end -}}
								{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
								.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
								{{- end -}}
								{{- if .DefaultValue -}}
								.AddDefaultValueDescription("{{.DefaultValue}}")
								{{- end -}}
								.String,
							{{- if isListSet .}}
							ElementType:         types.{{.ElementType}}Type,
							{{- end}}
							{{- if or .Reference .Mandatory}}
							Required:            true,
							{{- else if not .Computed}}
							Optional:            true,
							{{- end}}
							{{- if or (len .DefaultValue) .Computed}}
							Computed:            true,
							{{- end}}
							{{- if len .EnumValues}}
							Validators: []validator.String{
								stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
							},
							{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
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
							{{- if and (len .DefaultValue) (eq .Type "Int64")}}
							Default:             int64default.StaticInt64({{.DefaultValue}}),
							{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
							Default:             booldefault.StaticBool({{.DefaultValue}}),
							{{- else if and (len .DefaultValue) (eq .Type "String")}}
							Default:             stringdefault.StaticString("{{.DefaultValue}}"),
							{{- end}}
							{{- if or .RequiresReplace .Computed}}
							PlanModifiers: []planmodifier.{{.Type}}{
								{{- if or .RequiresReplace}}
								{{snakeCase .Type}}planmodifier.RequiresReplace(),
								{{end}}
								{{- if .Computed}}
								{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
								{{end}}
							},
							{{- end}}
							{{- if isNestedListSetMap .}}
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									{{- range  .Attributes}}
									{{- if not .Value}}
									"{{.TfName}}": schema.{{if isNestedListSetMap .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
											{{- if len .EnumValues -}}
											.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
											{{- end -}}
											{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
											.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
											{{- end -}}
											{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
											.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
											{{- end -}}
											{{- if .DefaultValue -}}
											.AddDefaultValueDescription("{{.DefaultValue}}")
											{{- end -}}
											.String,
										{{- if isListSet .}}
										ElementType:         types.{{.ElementType}}Type,
										{{- end}}
										{{- if or .Reference .Mandatory}}
										Required:            true,
										{{- else if not .Computed}}
										Optional:            true,
										{{- end}}
										{{- if or (len .DefaultValue) .Computed}}
										Computed:            true,
										{{- end}}
										{{- if len .EnumValues}}
										Validators: []validator.String{
											stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
										},
										{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
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
										{{- if and (len .DefaultValue) (eq .Type "Int64")}}
										Default:             int64default.StaticInt64({{.DefaultValue}}),
										{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
										Default:             booldefault.StaticBool({{.DefaultValue}}),
										{{- else if and (len .DefaultValue) (eq .Type "String")}}
										Default:             stringdefault.StaticString("{{.DefaultValue}}"),
										{{- end}}
										{{- if .Computed}}
										PlanModifiers: []planmodifier.{{.Type}}{
											{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
										},
										{{- end}}
										{{- if isNestedListSetMap .}}
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												{{- range  .Attributes}}
												{{- if not .Value}}
												"{{.TfName}}": schema.{{if isNestedListSetMap .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
														{{- if len .EnumValues -}}
														.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
														{{- end -}}
														{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
														.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
														{{- end -}}
														{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
														.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
														{{- end -}}
														{{- if .DefaultValue -}}
														.AddDefaultValueDescription("{{.DefaultValue}}")
														{{- end -}}
														.String,
													{{- if isListSet .}}
													ElementType:         types.{{.ElementType}}Type,
													{{- end}}
													{{- if or .Reference .Mandatory}}
													Required:            true,
													{{- else if not .Computed}}
													Optional:            true,
													{{- end}}
													{{- if or (len .DefaultValue) .Computed}}
													Computed:            true,
													{{- end}}
													{{- if len .EnumValues}}
													Validators: []validator.String{
														stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
													},
													{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
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
													{{- if and (len .DefaultValue) (eq .Type "Int64")}}
													Default:             int64default.StaticInt64({{.DefaultValue}}),
													{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
													Default:             booldefault.StaticBool({{.DefaultValue}}),
													{{- else if and (len .DefaultValue) (eq .Type "String")}}
													Default:             stringdefault.StaticString("{{.DefaultValue}}"),
													{{- end}}
													{{- if .Computed}}
													PlanModifiers: []planmodifier.{{.Type}}{
														{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
													},
													{{- end}}
													{{- if isNestedListSetMap .}}
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															{{- range  .Attributes}}
															{{- if not .Value}}
															"{{.TfName}}": schema.{{if isNestedListSetMap .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
																MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}")
																	{{- if len .EnumValues -}}
																	.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
																	{{- end -}}
																	{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
																	.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
																	{{- end -}}
																	{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
																	.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
																	{{- end -}}
																	{{- if .DefaultValue -}}
																	.AddDefaultValueDescription("{{.DefaultValue}}")
																	{{- end -}}
																	.String,
																{{- if isListSet .}}
																ElementType:         types.{{.ElementType}}Type,
																{{- end}}
																{{- if or .Reference .Mandatory}}
																Required:            true,
																{{- else if not .Computed}}
																Optional:            true,
																{{- end}}
																{{- if or (len .DefaultValue) .Computed}}
																Computed:            true,
																{{- end}}
																{{- if len .EnumValues}}
																Validators: []validator.String{
																	stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
																},
																{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
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
																{{- if and (len .DefaultValue) (eq .Type "Int64")}}
																Default:             int64default.StaticInt64({{.DefaultValue}}),
																{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
																Default:             booldefault.StaticBool({{.DefaultValue}}),
																{{- else if and (len .DefaultValue) (eq .Type "String")}}
																Default:             stringdefault.StaticString("{{.DefaultValue}}"),
																{{- end}}
																{{- if .Computed}}
																PlanModifiers: []planmodifier.{{.Type}}{
																	{{snakeCase .Type}}planmodifier.UseStateForUnknown(),
																},
																{{- end}}
															},
															{{- end}}
															{{- end}}
														},
													},
													{{- if or (ne .MinList 0) (ne .MaxList 0)}}
													Validators: []validator.List{
														{{- if ne .MinList 0}}
														listvalidator.SizeAtLeast({{.MinList}}),
														{{- end}}
														{{- if ne .MaxList 0}}
														listvalidator.SizeAtMost({{.MaxList}}),
														{{- end}}
													},
													{{- end}}
													{{- end}}
												},
												{{- end}}
												{{- end}}
											},
										},
										{{- if or (ne .MinList 0) (ne .MaxList 0)}}
										Validators: []validator.List{
											{{- if ne .MinList 0}}
											listvalidator.SizeAtLeast({{.MinList}}),
											{{- end}}
											{{- if ne .MaxList 0}}
											listvalidator.SizeAtMost({{.MaxList}}),
											{{- end}}
										},
										{{- end}}
										{{- end}}
									},
									{{- end}}
									{{- end}}
								},
							},
							{{- if or (ne .MinList 0) (ne .MaxList 0)}}
							Validators: []validator.List{
								{{- if ne .MinList 0}}
								listvalidator.SizeAtLeast({{.MinList}}),
								{{- end}}
								{{- if ne .MaxList 0}}
								listvalidator.SizeAtMost({{.MaxList}}),
								{{- end}}
							},
							{{- end}}
							{{- end}}
						},
						{{- end}}
						{{- end}}
					},
				},
			},
		},
	}
}

func (r *{{camelCase .BulkName}}Resource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *{{camelCase .BulkName}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Resource{{camelCase .BulkName}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	actions := make([]meraki.ActionModel, len(plan.Items))

	for i, item := range plan.Items {
		{{- if .PutCreate}}
		actions[i] = meraki.ActionModel{
			Operation: "update",
			Resource:  plan.getItemPath(item.{{toGoName ((getId .Attributes).TfName)}}.ValueString()),
			Body:      item.toBody(ctx, Resource{{camelCase .BulkName}}Items{}),
		}
		{{- else}}
		actions[i] = meraki.ActionModel{
			Operation: "create",
			Resource:  plan.getPath(),
			Body:      item.toBody(ctx, Resource{{camelCase .BulkName}}Items{}),
		}
		{{- end}}
	}
	var res meraki.Res
	var err error
	if len(actions) > 0 {
		res, err = r.client.Batch(plan.OrganizationId.ValueString(), actions)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
			return
		}
	}
	plan.Id = plan.OrganizationId
	{{- if not .PutCreate}}
	for i := range plan.Items {
		plan.Items[i].Id = types.StringValue(res.Get("status.createdResources." + strconv.Itoa(i) + ".id").String())
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *{{camelCase .BulkName}}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Resource{{camelCase .BulkName}}

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	{{- if not .NoRead}}
	res, err := r.client.Get(state.getPath())
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		if len(state.Items) > 0 {
			state.fromBodyImport(ctx, res)
		} else {
			state.fromBody(ctx, res)
		}
	} else {
		state.fromBodyPartial(ctx, res)
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *{{camelCase .BulkName}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Resource{{camelCase .BulkName}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	{{- if not .NoUpdate}}
	var actions []meraki.ActionModel

	{{- if or (hasDestroyValues .Attributes) (not .NoDelete)}}
	// If there are destroy values, we need to compare the plan and state to determine what to delete
	for _, itemState := range state.Items {
		found := false
		for _, item := range plan.Items {
			{{- range .Attributes}}
			{{- if .BulkId}}
			if item.{{toGoName .TfName}}.Value{{.Type}}() != itemState.{{toGoName .TfName}}.Value{{.Type}}() {
				continue
			}
			{{- end}}
			{{- end}}

			// If the item is present in both plan and state, we can skip it
			found = true
			break
		}
		if !found {
			// If the item is present in state, but not in plan, we need to delete it
			{{- if hasDestroyValues .Attributes}}
			actions = append(actions, meraki.ActionModel{
				Operation: "update",
				Resource:  plan.getItemPath(itemState.{{toGoName ((getId .Attributes).TfName)}}.ValueString()),
				Body:      plan.toDestroyBody(ctx),
			})
			{{- else if not .NoDelete}}
			actions = append(actions, meraki.ActionModel{
				Operation: "destroy",
				Resource:  plan.getPath() + "/" + itemState.Id.ValueString(),
				Body:      "{}",
			})
			{{- end}}
		}
	}
	{{- end}}

	{{- if not .PutCreate}}
	newIdIndexes := make([]int, 0)
	{{- end}}
	// Check for new and updated items
	for i := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			{{- range .Attributes}}
			{{- if .BulkId}}
			if plan.Items[i].{{toGoName .TfName}}.Value{{.Type}}() != itemState.{{toGoName .TfName}}.Value{{.Type}}() {
				continue
			}
			{{- end}}
			{{- end}}

			found = true
			// If the item is present in both plan and state, we need to check if it has changes
			hasChanges := plan.hasChanges(ctx, &state, plan.Items[i].{{getBulkItemId .}}.ValueString())
			if hasChanges {
				actions = append(actions, meraki.ActionModel{
					Operation: "update",
					{{- if .PutCreate}}
					Resource:  plan.getItemPath(plan.Items[i].{{getBulkItemId .}}.ValueString()),
					{{- else}}
					Resource:  plan.getPath() + "/" + plan.Items[i].Id.ValueString(),
					{{- end}}
					Body:      plan.Items[i].toBody(ctx, itemState),
				})					
			}
			break
		}
		if !found {
			// If the item is present in plan, but not in state, we need to create it
			{{- if .PutCreate}}
			actions = append(actions, meraki.ActionModel{
				Operation: "update",
				Resource:  plan.getItemPath(plan.Items[i].{{getBulkItemId .}}.ValueString()),
				Body:      plan.Items[i].toBody(ctx, Resource{{camelCase .BulkName}}Items{}),
			})
			{{- else}}
			actions = append(actions, meraki.ActionModel{
				Operation: "create",
				Resource:  plan.getPath(),
				Body:      plan.Items[i].toBody(ctx, Resource{{camelCase .BulkName}}Items{}),
			})
			newIdIndexes = append(newIdIndexes, i)
			{{- end}}
		}
	}

	var res meraki.Res
	var err error
	if len(actions) > 0 {
		res, err = r.client.Batch(plan.OrganizationId.ValueString(), actions)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
			return
		}
	}
	{{- end}}
	{{- if not .PutCreate}}
	responseIndex := 0
	for i := range plan.Items {
		if slices.Contains(newIdIndexes, i) {
			plan.Items[i].Id = types.StringValue(res.Get("status.createdResources." + strconv.Itoa(responseIndex) + ".id").String())
			responseIndex++
		}
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *{{camelCase .BulkName}}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Resource{{camelCase .BulkName}}

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	{{- if hasDestroyValues .Attributes}}
	actions := make([]meraki.ActionModel, len(state.Items))

	for i, item := range state.Items {
		actions[i] = meraki.ActionModel{
			Operation: "update",
			Resource:  state.getItemPath(item.{{toGoName ((getId .Attributes).TfName)}}.ValueString()),
			Body:      state.toDestroyBody(ctx),
		}
	}
	if len(actions) > 0 {
		res, err := r.client.Batch(state.OrganizationId.ValueString(), actions)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
			return
		}
	}
	{{- else if not .NoDelete}}
	actions := make([]meraki.ActionModel, len(state.Items))

	for i, item := range state.Items {
		actions[i] = meraki.ActionModel{
			Operation: "destroy",
			Resource:  state.getPath() + "/" + item.Id.ValueString(),
			Body:      "{}",
		}
	}
	if len(actions) > 0 {
		res, err := r.client.Batch(state.OrganizationId.ValueString(), actions)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
			return
		}
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
{{- if not .NoImport}}
func (r *{{camelCase .BulkName}}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	itemIdParts := make([]string, 0)
	if strings.Contains(req.ID, ",[") {
		itemIdParts = strings.Split(strings.Split(strings.Split(req.ID, ",[")[1], "]")[0], ",")
	}
	idParts := strings.Split(strings.Split(req.ID, ",[")[0], ",")

	if len(idParts) != {{len (getBulkImportAttributes .)}}{{range $index, $attr := (getBulkImportAttributes .)}} || idParts[{{$index}}] == ""{{end}} {
		expectedIdentifier := "Expected import identifier with format: {{range $i, $e := (getBulkImportAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
		expectedIdentifier += " or {{range $i, $e := (getBulkImportAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}},[<{{getBulkItemIdTfName .}}>,...]"
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("%s. Got: %q",expectedIdentifier, req.ID),
		)
		return
	}

	{{- range $index, $attr := (getBulkImportAttributes .)}}
	{{- if eq .TfName "organization_id"}}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[{{$index}}])...)
	{{- end}}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("{{.TfName}}"), {{if eq .Type "Bool"}}helpers.Must(strconv.ParseBool(idParts[{{$index}}])){{else if eq .Type "Int64"}}helpers.Must(strconv.ParseInt(idParts[{{$index}}])){{else if eq .Type "Float64"}}helpers.Must(strconv.ParseFloat(idParts[{{$index}}])){{else}}idParts[{{$index}}]{{end}})...)
	{{- end}}

	if len(itemIdParts) > 0 {
		items := make([]Resource{{camelCase .BulkName}}Items, len(itemIdParts))
		for i, itemId := range itemIdParts {
			item := Resource{{camelCase .BulkName}}Items{}
			item.{{getBulkItemId .}} = types.StringValue(itemId)
			{{- range .Attributes}}
			{{- if isListSet .}}
			item.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
			{{- end}}
			{{- end}}
			items[i] = item
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("items"), items)...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
{{- end}}
// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin modifyPlan
func (r *{{camelCase .BulkName}}Resource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan, state Resource{{camelCase .BulkName}}

	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning ModifyPlan", plan.Id.ValueString()))

	{{- if not .PutCreate}}
	// Remove incorrectly set IDs in plan (https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#prior-state-under-lists-and-sets)
	for i, item := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			{{- range .Attributes}}
			{{- if .BulkId}}
			if item.{{toGoName .TfName}}.Value{{.Type}}() != itemState.{{toGoName .TfName}}.Value{{.Type}}() {
				continue
			}
			{{- end}}
			{{- end}}
			found = true
		}
		if !found {
			plan.Items[i].Id = types.StringUnknown()
		}
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: ModifyPlan finished successfully", plan.Id.ValueString()))

	diags = resp.Plan.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}
// End of section. //template:end modifyPlan
