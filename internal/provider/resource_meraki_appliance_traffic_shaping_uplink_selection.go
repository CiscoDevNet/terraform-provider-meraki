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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	_ resource.Resource                = &ApplianceTrafficShapingUplinkSelectionResource{}
	_ resource.ResourceWithImportState = &ApplianceTrafficShapingUplinkSelectionResource{}
)

func NewApplianceTrafficShapingUplinkSelectionResource() resource.Resource {
	return &ApplianceTrafficShapingUplinkSelectionResource{}
}

type ApplianceTrafficShapingUplinkSelectionResource struct {
	client *meraki.Client
}

func (r *ApplianceTrafficShapingUplinkSelectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_traffic_shaping_uplink_selection"
}

func (r *ApplianceTrafficShapingUplinkSelectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance Traffic Shaping Uplink Selection` configuration.").String,

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
			"active_active_auto_vpn_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Toggle for enabling or disabling active-active AutoVPN").String,
				Optional:            true,
			},
			"default_uplink": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The default uplink. Must be one of: `wan1` or `wan2`").AddStringEnumDescription("wan1", "wan2").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("wan1", "wan2"),
				},
			},
			"load_balancing_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Toggle for enabling or disabling load balancing").String,
				Optional:            true,
			},
			"failover_and_failback_immediate_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Toggle for enabling or disabling immediate WAN failover and failback").String,
				Optional:            true,
			},
			"vpn_traffic_uplink_preferences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Array of uplink preference rules for VPN traffic").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fail_over_criterion": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Fail over criterion for this uplink preference rule. Must be one of: `poorPerformance` or `uplinkDown`").AddStringEnumDescription("poorPerformance", "uplinkDown").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("poorPerformance", "uplinkDown"),
							},
						},
						"preferred_uplink": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Preferred uplink for this uplink preference rule. Must be one of: `wan1`, `wan2`, `bestForVoIP`, `loadBalancing` or `defaultUplink`").AddStringEnumDescription("bestForVoIP", "defaultUplink", "loadBalancing", "wan1", "wan2").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("bestForVoIP", "defaultUplink", "loadBalancing", "wan1", "wan2"),
							},
						},
						"builtin_performance_class_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of builtin performance class, must be present when performanceClass type is `builtin`, and value must be one of: `VoIP`").AddStringEnumDescription("VoIP").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("VoIP"),
							},
						},
						"custom_performance_class_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of created custom performance class, must be present when performanceClass type is `custom`").String,
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
							MarkdownDescription: helpers.NewAttributeDescription("Array of traffic filters for this uplink preference rule").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of this traffic filter. Must be one of: `applicationCategory`, `application` or `custom`").AddStringEnumDescription("application", "applicationCategory", "custom").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("application", "applicationCategory", "custom"),
										},
									},
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ID of this applicationCategory or application type traffic filter. E.g.: 'meraki:layer7/category/1', 'meraki:layer7/application/4'").String,
										Optional:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Protocol of this custom type traffic filter. Must be one of: `tcp`, `udp`, `icmp`, `icmp6` or `any`").AddStringEnumDescription("any", "icmp", "icmp6", "tcp", "udp").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("any", "icmp", "icmp6", "tcp", "udp"),
										},
									},
									"destination_cidr": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')").String,
										Optional:            true,
									},
									"destination_fqdn": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("FQDN format address. Currently only availabe in `destination` of `vpnTrafficUplinkPreference` object. E.g.: `www.google.com`").String,
										Optional:            true,
									},
									"destination_host": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Host ID in the VLAN, should be used along with `vlan`, and not exceed the vlan subnet capacity. Currently only available under a template network.").String,
										Optional:            true,
									},
									"destination_network": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: 'L_12345678'.").String,
										Optional:            true,
									},
									"destination_port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'").String,
										Optional:            true,
									},
									"destination_vlan": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.").String,
										Optional:            true,
									},
									"source_cidr": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')").String,
										Optional:            true,
									},
									"source_host": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Host ID in the VLAN, should be used along with `vlan`, and not exceed the vlan subnet capacity. Currently only available under a template network.").String,
										Optional:            true,
									},
									"source_network": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: 'L_12345678'.").String,
										Optional:            true,
									},
									"source_port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'").String,
										Optional:            true,
									},
									"source_vlan": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.").String,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"wan_traffic_uplink_preferences": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Array of uplink preference rules for WAN traffic. Note: these preferences are shared (merged) with meraki_appliance_sdwan_internet_policies resource. It is recommended to only use one resource for these preferences, not both at the same time: Deleting this resource clears preferences created in meraki_appliance_sdwan_internet_policies, which isn't detected as a change by the provider. The same happens the other way around, but the change is detected in uplink_selection on a subsequent apply.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"preferred_uplink": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Preferred uplink for this uplink preference rule. Must be one of: `wan1` or `wan2`").AddStringEnumDescription("wan1", "wan2").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("wan1", "wan2"),
							},
						},
						"traffic_filters": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Array of traffic filters for this uplink preference rule").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of this traffic filter. Must be one of: `custom`").AddStringEnumDescription("custom").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("custom"),
										},
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Protocol of this custom type traffic filter. Must be one of: `tcp`, `udp`, `icmp6` or `any`").AddStringEnumDescription("any", "icmp6", "tcp", "udp").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("any", "icmp6", "tcp", "udp"),
										},
									},
									"destination_cidr": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')").String,
										Optional:            true,
									},
									"destination_port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'").String,
										Optional:            true,
									},
									"source_cidr": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')").String,
										Optional:            true,
									},
									"source_host": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Host ID in the VLAN, should be used along with `vlan`, and not exceed the vlan subnet capacity. Currently only available under a template network.").String,
										Optional:            true,
									},
									"source_port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'").String,
										Optional:            true,
									},
									"source_vlan": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.").String,
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

func (r *ApplianceTrafficShapingUplinkSelectionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ApplianceTrafficShapingUplinkSelectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplianceTrafficShapingUplinkSelection

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ApplianceTrafficShapingUplinkSelection{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.NetworkId
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *ApplianceTrafficShapingUplinkSelectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplianceTrafficShapingUplinkSelection

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
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
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *ApplianceTrafficShapingUplinkSelectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplianceTrafficShapingUplinkSelection

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

func (r *ApplianceTrafficShapingUplinkSelectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplianceTrafficShapingUplinkSelection

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	body := state.toDestroyBody(ctx)
	res, err := r.client.Put(state.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	// Clear the shared wanTrafficUplinkPreferences
	// by PUTting "wanTrafficUplinkPreferences": []
	// to appliance/sdwan/internetPolicies endpoint,
	// as PUTting "wanTrafficUplinkPreferences": []
	// to appliance/trafficShaping/uplinkSelection endpoint does nothing.
	sdwanInternetPolicies := ApplianceSDWANInternetPolicies{
		NetworkId: state.NetworkId,
	}
	body = sdwanInternetPolicies.toDestroyBody(ctx)
	res, err = r.client.Put(sdwanInternetPolicies.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to clear Appliance Traffic Shaping Uplink Preferences wan_traffic_uplink_preferences via Appliance SDWAN Internet Policies endpoint (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ApplianceTrafficShapingUplinkSelectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
