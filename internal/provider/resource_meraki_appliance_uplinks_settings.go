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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &ApplianceUplinksSettingsResource{}
	_ resource.ResourceWithImportState = &ApplianceUplinksSettingsResource{}
)

func NewApplianceUplinksSettingsResource() resource.Resource {
	return &ApplianceUplinksSettingsResource{}
}

type ApplianceUplinksSettingsResource struct {
	client *meraki.Client
}

func (r *ApplianceUplinksSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_uplinks_settings"
}

func (r *ApplianceUplinksSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance Uplinks Settings` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device serial").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"interfaces_wan1_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable the interface.").String,
				Optional:            true,
			},
			"interfaces_wan1_pppoe_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether PPPoE is enabled.").String,
				Optional:            true,
			},
			"interfaces_wan1_pppoe_authentication_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether PPPoE authentication is enabled.").String,
				Optional:            true,
			},
			"interfaces_wan1_pppoe_authentication_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Password for PPPoE authentication. This parameter is not returned.").String,
				Optional:            true,
			},
			"interfaces_wan1_pppoe_authentication_username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Username for PPPoE authentication.").String,
				Optional:            true,
			},
			"interfaces_wan1_svis_ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address and subnet mask when in static mode.").String,
				Optional:            true,
			},
			"interfaces_wan1_svis_ipv4_assignment_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The assignment mode for this SVI. Applies only when PPPoE is disabled.").AddStringEnumDescription("dynamic", "static").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "static"),
				},
			},
			"interfaces_wan1_svis_ipv4_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Gateway IP address when in static mode.").String,
				Optional:            true,
			},
			"interfaces_wan1_svis_ipv4_nameservers_addresses": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"interfaces_wan1_svis_ipv6_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static address that will override the one(s) received by SLAAC.").String,
				Optional:            true,
			},
			"interfaces_wan1_svis_ipv6_assignment_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The assignment mode for this SVI. Applies only when PPPoE is disabled.").AddStringEnumDescription("dynamic", "static").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "static"),
				},
			},
			"interfaces_wan1_svis_ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static gateway that will override the one received by autoconf.").String,
				Optional:            true,
			},
			"interfaces_wan1_svis_ipv6_nameservers_addresses": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"interfaces_wan1_vlan_tagging_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether VLAN tagging is enabled.").String,
				Optional:            true,
			},
			"interfaces_wan1_vlan_tagging_vlan_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the VLAN to use for VLAN tagging.").String,
				Optional:            true,
			},
			"interfaces_wan2_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable the interface.").String,
				Optional:            true,
			},
			"interfaces_wan2_pppoe_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether PPPoE is enabled.").String,
				Optional:            true,
			},
			"interfaces_wan2_pppoe_authentication_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether PPPoE authentication is enabled.").String,
				Optional:            true,
			},
			"interfaces_wan2_pppoe_authentication_password": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Password for PPPoE authentication. This parameter is not returned.").String,
				Optional:            true,
			},
			"interfaces_wan2_pppoe_authentication_username": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Username for PPPoE authentication.").String,
				Optional:            true,
			},
			"interfaces_wan2_svis_ipv4_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("IP address and subnet mask when in static mode.").String,
				Optional:            true,
			},
			"interfaces_wan2_svis_ipv4_assignment_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The assignment mode for this SVI. Applies only when PPPoE is disabled.").AddStringEnumDescription("dynamic", "static").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "static"),
				},
			},
			"interfaces_wan2_svis_ipv4_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Gateway IP address when in static mode.").String,
				Optional:            true,
			},
			"interfaces_wan2_svis_ipv4_nameservers_addresses": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"interfaces_wan2_svis_ipv6_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static address that will override the one(s) received by SLAAC.").String,
				Optional:            true,
			},
			"interfaces_wan2_svis_ipv6_assignment_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The assignment mode for this SVI. Applies only when PPPoE is disabled.").AddStringEnumDescription("dynamic", "static").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dynamic", "static"),
				},
			},
			"interfaces_wan2_svis_ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Static gateway that will override the one received by autoconf.").String,
				Optional:            true,
			},
			"interfaces_wan2_svis_ipv6_nameservers_addresses": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"interfaces_wan2_vlan_tagging_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether VLAN tagging is enabled.").String,
				Optional:            true,
			},
			"interfaces_wan2_vlan_tagging_vlan_id": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The ID of the VLAN to use for VLAN tagging.").String,
				Optional:            true,
			},
		},
	}
}

func (r *ApplianceUplinksSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ApplianceUplinksSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplianceUplinksSettings

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ApplianceUplinksSettings{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.Serial
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *ApplianceUplinksSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplianceUplinksSettings

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath())
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
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

func (r *ApplianceUplinksSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplianceUplinksSettings

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

func (r *ApplianceUplinksSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplianceUplinksSettings

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
func (r *ApplianceUplinksSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <serial>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
