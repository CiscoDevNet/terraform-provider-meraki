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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource = &SwitchPortsResource{}
)

func NewSwitchPortsResource() resource.Resource {
	return &SwitchPortsResource{}
}

type SwitchPortsResource struct {
	client *meraki.Client
}

func (r *SwitchPortsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_ports"
}

func (r *SwitchPortsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Switch Ports` configuration.").String,

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
				Optional:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device serial").String,
				Required:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The list of items",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access_policy_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The number of a custom access policy to configure on the switch port. Only applicable when `accessPolicyType` is `Custom access policy`.").String,
							Optional:            true,
						},
						"access_policy_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of the access policy of the switch port. Only applicable to access ports. Can be one of `Open`, `Custom access policy`, `MAC allow list` or `Sticky MAC allow list`.").AddStringEnumDescription("Custom access policy", "MAC allow list", "Open", "Sticky MAC allow list").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Custom access policy", "MAC allow list", "Open", "Sticky MAC allow list"),
							},
						},
						"allowed_vlans": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The VLANs allowed on the switch port. Only applicable to trunk ports.").String,
							Optional:            true,
						},
						"dai_trusted": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.").String,
							Optional:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The status of the switch port.").String,
							Optional:            true,
						},
						"flexible_stacking_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.").String,
							Optional:            true,
						},
						"isolation_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The isolation status of the switch port.").String,
							Optional:            true,
						},
						"link_negotiation": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The link speed for the switch port.").String,
							Optional:            true,
						},
						"mac_whitelist_limit": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The maximum number of MAC addresses for regular MAC allow list. Only applicable when `accessPolicyType` is `MAC allow list`. Note: Config only supported on verions greater than ms18 only for classic switches.").String,
							Optional:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the switch port.").String,
							Optional:            true,
						},
						"peer_sgt_capable": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.").String,
							Optional:            true,
						},
						"poe_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The PoE status of the switch port.").String,
							Optional:            true,
						},
						"port_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The identifier of the switch port.").String,
							Optional:            true,
						},
						"port_schedule_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The ID of the port schedule. A value of null will clear the port schedule.").String,
							Optional:            true,
						},
						"rstp_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The rapid spanning tree protocol status.").String,
							Optional:            true,
						},
						"sticky_mac_allow_list_limit": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The maximum number of MAC addresses for sticky MAC allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.").String,
							Optional:            true,
						},
						"storm_control_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The storm control status of the switch port.").String,
							Optional:            true,
						},
						"stp_guard": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The state of the STP guard (`disabled`, `root guard`, `bpdu guard` or `loop guard`).").AddStringEnumDescription("bpdu guard", "disabled", "loop guard", "root guard").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("bpdu guard", "disabled", "loop guard", "root guard"),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of the switch port (`trunk`, `access` or `stack`).").AddStringEnumDescription("access", "stack", "trunk").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("access", "stack", "trunk"),
							},
						},
						"udld": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.").AddStringEnumDescription("Alert only", "Enforce").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Alert only", "Enforce"),
							},
						},
						"vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.").String,
							Optional:            true,
						},
						"voice_vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The voice VLAN of the switch port. Only applicable to access ports.").String,
							Optional:            true,
						},
						"adaptive_policy_group_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The ID of the adaptive policy group.").String,
							Optional:            true,
						},
						"adaptive_policy_group_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the adaptive policy group.").String,
							Optional:            true,
						},
						"dot3az_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The Energy Efficient Ethernet status of the switch port.").String,
							Optional:            true,
						},
						"mirror_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The port mirror mode. Can be one of (`Destination port`, `Source port` or `Not mirroring traffic`).").AddStringEnumDescription("Destination port", "Not mirroring traffic", "Source port").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Destination port", "Not mirroring traffic", "Source port"),
							},
						},
						"module_model": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The model of the expansion module.").String,
							Optional:            true,
						},
						"profile_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("When enabled, override this port`s configuration with a port profile.").String,
							Optional:            true,
						},
						"profile_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("When enabled, the ID of the port profile used to override the port`s configuration.").String,
							Optional:            true,
						},
						"profile_iname": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("When enabled, the IName of the profile.").String,
							Optional:            true,
						},
						"schedule_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The ID of the port schedule.").String,
							Optional:            true,
						},
						"schedule_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the port schedule.").String,
							Optional:            true,
						},
						"stackwise_virtual_is_dual_active_detector": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("For SVL devices, whether or not the port is used for Dual Active Detection.").String,
							Optional:            true,
						},
						"stackwise_virtual_is_stack_wise_virtual_link": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("For SVL devices, whether or not the port is used for StackWise Virtual Link.").String,
							Optional:            true,
						},
						"link_negotiation_capabilities": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Available link speeds for the switch port.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"mac_allow_list": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when `accessPolicyType` is `MAC allow list`.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"sticky_mac_allow_list": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The initial list of MAC addresses for sticky Mac allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The list of tags of the switch port.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *SwitchPortsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

func (r *SwitchPortsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SwitchPorts

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	actions := make([]meraki.ActionModel, 0, len(plan.Items))
	for _, item := range plan.Items {
		body := plan.toBody(ctx, SwitchPorts{}, item.PortId.ValueString())
		actions = append(actions, meraki.NewAction("update", "/networks/{networkId}/appliance/ports/{portId}", body))
	}

	res, err := r.client.Batch(plan.OrganizationId.String(), actions)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create objects (Batch), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.Serial

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *SwitchPortsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SwitchPorts

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

func (r *SwitchPortsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SwitchPorts

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *SwitchPortsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SwitchPorts

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
