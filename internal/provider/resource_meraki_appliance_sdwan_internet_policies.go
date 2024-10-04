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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource = &ApplianceSDWANInternetPoliciesResource{}
)

func NewApplianceSDWANInternetPoliciesResource() resource.Resource {
	return &ApplianceSDWANInternetPoliciesResource{}
}

type ApplianceSDWANInternetPoliciesResource struct {
	client *meraki.Client
}

func (r *ApplianceSDWANInternetPoliciesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_sdwan_internet_policies"
}

func (r *ApplianceSDWANInternetPoliciesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance SDWAN Internet Policies` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"wan_traffic_uplink_preferences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("policies with respective traffic filters for an MX network").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fail_over_criterion": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("WAN failover and failback behavior").AddStringEnumDescription("poorPerformance", "uplinkDown").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("poorPerformance", "uplinkDown"),
							},
						},
						"preferred_uplink": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Preferred uplink for uplink preference rule. Must be one of: `wan1`, `wan2`, `bestForVoIP`, `loadBalancing` or `defaultUplink`").AddStringEnumDescription("bestForVoIP", "defaultUplink", "loadBalancing", "wan1", "wan2").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("bestForVoIP", "defaultUplink", "loadBalancing", "wan1", "wan2"),
							},
						},
						"builtin_performance_class_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of builtin performance class. Must be present when performanceClass type is `builtin` and value must be one of: `VoIP`").AddStringEnumDescription("VoIP").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("VoIP"),
							},
						},
						"custom_performance_class_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of created custom performance class, must be present when performanceClass type is 'custom'").String,
							Optional:            true,
						},
						"performance_class_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of this performance class. Must be one of: `builtin` or `custom`").AddStringEnumDescription("builtin", "custom").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("builtin", "custom"),
							},
						},
						"traffic_filters": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Traffic filters").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Traffic filter type. Must be `custom`, `major_application`, `application (NBAR)`, if type is `application`, you can pass either an NBAR App Category or Application").AddStringEnumDescription("application", "custom", "majorApplication").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("application", "custom", "majorApplication"),
										},
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Protocol of the traffic filter. Must be one of: `tcp`, `udp`, `icmp6` or `any`").AddStringEnumDescription("any", "icmp6", "tcp", "udp").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("any", "icmp6", "tcp", "udp"),
										},
									},
									"destination_cidr": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("CIDR format address (e.g.'192.168.10.1', which is the same as '192.168.10.1/32'), or 'any'").String,
										Optional:            true,
									},
									"destination_port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'").String,
										Optional:            true,
									},
									"destination_applications": schema.ListNestedAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("list of application objects (either majorApplication or nbar)").String,
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Id of the major application, or a list of NBAR Application Category or Application selections").String,
													Optional:            true,
												},
												"name": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("Name of the major application or application category selected").String,
													Optional:            true,
												},
												"type": schema.StringAttribute{
													MarkdownDescription: helpers.NewAttributeDescription("app type (major or nbar)").String,
													Optional:            true,
												},
											},
										},
									},
									"source_cidr": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("CIDR format address (e.g.'192.168.10.1', which is the same as '192.168.10.1/32'), or 'any'. Cannot be used in combination with the 'vlan' property").String,
										Optional:            true,
									},
									"source_host": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Host ID in the VLAN. Should not exceed the VLAN subnet capacity. Must be used along with the 'vlan' property and is currently only available under a template network.").String,
										Optional:            true,
									},
									"source_port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'").String,
										Optional:            true,
									},
									"source_vlan": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("VLAN ID of the configured VLAN in the Meraki network. Cannot be used in combination with the 'cidr' property and is currently only available under a template network.").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *ApplianceSDWANInternetPoliciesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ApplianceSDWANInternetPoliciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplianceSDWANInternetPolicies

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ApplianceSDWANInternetPolicies{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.NetworkId

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *ApplianceSDWANInternetPoliciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplianceSDWANInternetPolicies

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *ApplianceSDWANInternetPoliciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplianceSDWANInternetPolicies

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

	body := plan.toBody(ctx, state)
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *ApplianceSDWANInternetPoliciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplianceSDWANInternetPolicies

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import
