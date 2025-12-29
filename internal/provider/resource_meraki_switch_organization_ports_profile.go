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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
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
	_ resource.ResourceWithIdentity    = &SwitchOrganizationPortsProfileResource{}
	_ resource.ResourceWithImportState = &SwitchOrganizationPortsProfileResource{}
)

func NewSwitchOrganizationPortsProfileResource() resource.Resource {
	return &SwitchOrganizationPortsProfileResource{}
}

type SwitchOrganizationPortsProfileResource struct {
	client *meraki.Client
}

func (r *SwitchOrganizationPortsProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_organization_ports_profile"
}

func (r *SwitchOrganizationPortsProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Switch Organization Ports Profile` configuration.").AddEarlyAccessDescription().String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Organization ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Text describing the profile.").String,
				Optional:            true,
			},
			"is_organization_wide": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The scope of the profile whether it is organization level or network level").String,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the profile.").String,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The network identifier").String,
				Optional:            true,
			},
			"networks_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Flag to identify if the networks networks are excluded or included").String,
				Optional:            true,
			},
			"networks_values": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The network object containing name and id").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The network identifier").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The network name").String,
							Optional:            true,
						},
					},
				},
			},
			"port_access_policy_number": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The number of a custom access policy to configure on the port profile. Only applicable when `accessPolicyType` is `Custom access policy`.").String,
				Optional:            true,
			},
			"port_access_policy_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the access policy of the port profile. Only applicable to access ports.").AddStringEnumDescription("Custom access policy", "MAC allow list", "Open", "Sticky MAC allow list").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Custom access policy", "MAC allow list", "Open", "Sticky MAC allow list"),
				},
			},
			"port_adaptive_policy_group_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The adaptive policy group ID that will be used to tag traffic through this port profile. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API.").String,
				Optional:            true,
			},
			"port_adaptive_policy_voice_group_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The adaptive policy group ID that will be used to tag voice traffic through this port profile. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API.").String,
				Optional:            true,
			},
			"port_allowed_vlans": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The VLANs allowed on the port profile. Only applicable to trunk ports.").String,
				Optional:            true,
			},
			"port_dai_trusted": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.").String,
				Optional:            true,
			},
			"port_isolation_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The isolation status of the port profile.").String,
				Optional:            true,
			},
			"port_peer_sgt_capable": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If true, Peer SGT is enabled for traffic through this port profile. Applicable to trunk ports only.").String,
				Optional:            true,
			},
			"port_poe_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The PoE status of the port profile.").String,
				Optional:            true,
			},
			"port_rstp_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The rapid spanning tree protocol status.").String,
				Optional:            true,
			},
			"port_sticky_mac_allow_list_limit": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The maximum number of MAC addresses for sticky MAC allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.").String,
				Optional:            true,
			},
			"port_storm_control_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The storm control status of the port profile.").String,
				Optional:            true,
			},
			"port_stp_guard": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The state of the STP guard.").AddStringEnumDescription("bpdu guard", "disabled", "loop guard", "root guard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("bpdu guard", "disabled", "loop guard", "root guard"),
				},
			},
			"port_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The type of the port profile.").AddStringEnumDescription("access", "stack", "trunk").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("access", "stack", "trunk"),
				},
			},
			"port_udld": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The action to take when Unidirectional Link is detected. LinkDefault configuration is Alert only.").AddStringEnumDescription("Alert only", "Enforce").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Alert only", "Enforce"),
				},
			},
			"port_vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The VLAN of the port profile. A null value will clear the value set for trunk ports.").String,
				Optional:            true,
			},
			"port_voice_vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The voice VLAN of the port profile. Only applicable to access ports.").String,
				Optional:            true,
			},
			"port_mac_allow_list": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when `accessPolicyType` is `MAC allow list`.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"port_sticky_mac_allow_list": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The initial list of MAC addresses for sticky Mac allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Space-seperated list of tags").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *SwitchOrganizationPortsProfileResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"organization_id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("Organization ID").String,
				RequiredForImport: true,
			},
			"id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("").String,
				RequiredForImport: true,
			},
		},
	}
}

func (r *SwitchOrganizationPortsProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *SwitchOrganizationPortsProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SwitchOrganizationPortsProfile
	var identity SwitchOrganizationPortsProfileIdentity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, SwitchOrganizationPortsProfile{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("profileId").String())
	plan.fromBodyUnknowns(ctx, res)
	identity.toIdentity(ctx, &plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *SwitchOrganizationPortsProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SwitchOrganizationPortsProfile
	var identity SwitchOrganizationPortsProfileIdentity

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read identity
	diags = req.Identity.Get(ctx, &identity)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	state.fromIdentity(ctx, &identity)

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
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
	identity.toIdentity(ctx, &state)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *SwitchOrganizationPortsProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SwitchOrganizationPortsProfile

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
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
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

func (r *SwitchOrganizationPortsProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SwitchOrganizationPortsProfile

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *SwitchOrganizationPortsProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != "" {
		idParts := strings.Split(req.ID, ",")

		if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
			resp.Diagnostics.AddError(
				"Unexpected Import Identifier",
				fmt.Sprintf("Expected import identifier with format: <organization_id>,<id>. Got: %q", req.ID),
			)
			return
		}
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), idParts[1])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
	} else {
		var identity SwitchOrganizationPortsProfileIdentity
		diags := req.Identity.Get(ctx, &identity)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), identity.OrganizationId.ValueString())...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), identity.Id.ValueString())...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
