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
	_ resource.Resource                = &ApplianceSSIDsResource{}
	_ resource.ResourceWithImportState = &ApplianceSSIDsResource{}
)

func NewApplianceSSIDsResource() resource.Resource {
	return &ApplianceSSIDsResource{}
}

type ApplianceSSIDsResource struct {
	client *meraki.Client
}

func (r *ApplianceSSIDsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_ssids"
}

func (r *ApplianceSSIDsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance SSID` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "The organization ID",
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network ID").String,
				Required:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The list of items",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"number": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Wireless SSID number").String,
							Required:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"auth_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The association control method for the SSID (`open`, `psk`, `8021x-meraki` or `8021x-radius`).").AddStringEnumDescription("8021x-meraki", "8021x-radius", "open", "psk").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("8021x-meraki", "8021x-radius", "open", "psk"),
							},
						},
						"default_vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The VLAN ID of the VLAN associated to this SSID. This parameter is only valid if the network is in routed mode.").String,
							Optional:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not the SSID is enabled.").String,
							Optional:            true,
						},
						"encryption_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The psk encryption mode for the SSID (`wep` or `wpa`). This param is only valid if the authMode is `psk`.").AddStringEnumDescription("wep", "wpa").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("wep", "wpa"),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the SSID.").String,
							Optional:            true,
						},
						"psk": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The passkey for the SSID. This param is only valid if the authMode is `psk`.").String,
							Optional:            true,
						},
						"visible": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether the MX should advertise or hide this SSID.").String,
							Optional:            true,
						},
						"wpa_encryption_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The types of WPA encryption. (`WPA1 and WPA2`, `WPA2 only`, `WPA3 Transition Mode` or `WPA3 only`). This param is only valid if (1) the authMode is `psk` & the encryptionMode is `wpa` OR (2) the authMode is `8021x-meraki` OR (3) the authMode is `8021x-radius`").AddStringEnumDescription("WPA1 and WPA2", "WPA2 only", "WPA3 Transition Mode", "WPA3 only").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("WPA1 and WPA2", "WPA2 only", "WPA3 Transition Mode", "WPA3 only"),
							},
						},
						"dhcp_enforced_deauthentication_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable DCHP Enforced Deauthentication on the SSID.").String,
							Optional:            true,
						},
						"dot11w_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether 802.11w is enabled or not.").String,
							Optional:            true,
						},
						"dot11w_required": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("(Optional) Whether 802.11w is required or not.").String,
							Optional:            true,
						},
						"radius_servers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The RADIUS 802.1x servers to be used for authentication. This param is only valid if the authMode is `8021x-radius`.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The IP address of your RADIUS server.").String,
										Optional:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("The UDP port your RADIUS servers listens on for Access-requests.").String,
										Optional:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The RADIUS client shared secret.").String,
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

func (r *ApplianceSSIDsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ApplianceSSIDsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ResourceApplianceSSIDs

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	actions := make([]meraki.ActionModel, len(plan.Items))

	for i, item := range plan.Items {
		actions[i] = meraki.ActionModel{
			Operation: "update",
			Resource:  plan.getItemPath(item.Number.ValueString()),
			Body:      plan.toBody(ctx, ResourceApplianceSSIDs{}, item.Number.ValueString()),
		}
	}
	res, err := r.client.Batch(plan.OrganizationId.ValueString(), actions)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.OrganizationId

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *ApplianceSSIDsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ResourceApplianceSSIDs

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
		state.Id = state.OrganizationId
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

func (r *ApplianceSSIDsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ResourceApplianceSSIDs

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
	var actions []meraki.ActionModel

	// Check for new and updated items
	for _, item := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			if item.Number.ValueString() == itemState.Number.ValueString() {
				found = true
				// If the item is present in both plan and state, we need to check if it has changes
				hasChanges := plan.hasChanges(ctx, &state, item.Number.ValueString())
				if hasChanges {
					actions = append(actions, meraki.ActionModel{
						Operation: "update",
						Resource:  plan.getItemPath(item.Number.ValueString()),
						Body:      plan.toBody(ctx, ResourceApplianceSSIDs{}, item.Number.ValueString()),
					})
				}
				break
			}
		}
		if !found {
			// If the item is present in plan, but not in state, we need to create it
			actions = append(actions, meraki.ActionModel{
				Operation: "update",
				Resource:  plan.getItemPath(item.Number.ValueString()),
				Body:      plan.toBody(ctx, ResourceApplianceSSIDs{}, item.Number.ValueString()),
			})
		}
	}

	res, err := r.client.Batch(plan.OrganizationId.ValueString(), actions)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *ApplianceSSIDsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ResourceApplianceSSIDs

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
func (r *ApplianceSSIDsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <organization_id>,<network_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
