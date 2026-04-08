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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SwitchPortsDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchPortsDataSource{}
)

func NewSwitchPortsDataSource() datasource.DataSource {
	return &SwitchPortsDataSource{}
}

type SwitchPortsDataSource struct {
	client *meraki.Client
}

func (d *SwitchPortsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_ports"
}

func (d *SwitchPortsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Switch Ports` configuration in bulk.").String,

		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: "Device serial",
				Required:            true,
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The list of items",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the object",
							Computed:            true,
						},
						"access_policy_number": schema.Int64Attribute{
							MarkdownDescription: "The number of a custom access policy to configure on the switch port. Only applicable when `accessPolicyType` is `Custom access policy`.",
							Computed:            true,
						},
						"access_policy_type": schema.StringAttribute{
							MarkdownDescription: "The type of the access policy of the switch port. Only applicable to access ports. Can be one of `Open`, `Custom access policy`, `MAC allow list` or `Sticky MAC allow list`.",
							Computed:            true,
						},
						"allowed_vlans": schema.StringAttribute{
							MarkdownDescription: "The VLANs allowed on the switch port. Only applicable to trunk ports.",
							Computed:            true,
						},
						"dai_trusted": schema.BoolAttribute{
							MarkdownDescription: "If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "The status of the switch port.",
							Computed:            true,
						},
						"flexible_stacking_enabled": schema.BoolAttribute{
							MarkdownDescription: "For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.",
							Computed:            true,
						},
						"isolation_enabled": schema.BoolAttribute{
							MarkdownDescription: "The isolation status of the switch port.",
							Computed:            true,
						},
						"link_negotiation": schema.StringAttribute{
							MarkdownDescription: "The link speed for the switch port.",
							Computed:            true,
						},
						"mac_whitelist_limit": schema.Int64Attribute{
							MarkdownDescription: "The maximum number of MAC addresses for regular MAC allow list. Only applicable when `accessPolicyType` is `MAC allow list`. Note: Config only supported on verions greater than ms18 only for classic switches.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the switch port.",
							Computed:            true,
						},
						"peer_sgt_capable": schema.BoolAttribute{
							MarkdownDescription: "If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.",
							Computed:            true,
						},
						"poe_enabled": schema.BoolAttribute{
							MarkdownDescription: "The PoE status of the switch port.",
							Computed:            true,
						},
						"port_id": schema.StringAttribute{
							MarkdownDescription: "The identifier of the switch port.",
							Computed:            true,
						},
						"port_schedule_id": schema.StringAttribute{
							MarkdownDescription: "The ID of the port schedule. A value of null will clear the port schedule.",
							Computed:            true,
						},
						"rstp_enabled": schema.BoolAttribute{
							MarkdownDescription: "The rapid spanning tree protocol status.",
							Computed:            true,
						},
						"sticky_mac_allow_list_limit": schema.Int64Attribute{
							MarkdownDescription: "The maximum number of MAC addresses for sticky MAC allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.",
							Computed:            true,
						},
						"storm_control_enabled": schema.BoolAttribute{
							MarkdownDescription: "The storm control status of the switch port.",
							Computed:            true,
						},
						"stp_guard": schema.StringAttribute{
							MarkdownDescription: "The state of the STP guard (`disabled`, `root guard`, `bpdu guard` or `loop guard`).",
							Computed:            true,
						},
						"stp_port_fast_trunk": schema.BoolAttribute{
							MarkdownDescription: "The state of STP PortFast Trunk on the switch port.",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "The type of the switch port (`trunk`, `access` or `stack`).",
							Computed:            true,
						},
						"udld": schema.StringAttribute{
							MarkdownDescription: "The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.",
							Computed:            true,
						},
						"vlan": schema.Int64Attribute{
							MarkdownDescription: "The VLAN of the switch port. For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.",
							Computed:            true,
						},
						"voice_vlan": schema.Int64Attribute{
							MarkdownDescription: "The voice VLAN of the switch port. Only applicable to access ports.",
							Computed:            true,
						},
						"adaptive_policy_group_id": schema.StringAttribute{
							MarkdownDescription: "The ID of the adaptive policy group.",
							Computed:            true,
						},
						"adaptive_policy_group_name": schema.StringAttribute{
							MarkdownDescription: "The name of the adaptive policy group.",
							Computed:            true,
						},
						"dot3az_enabled": schema.BoolAttribute{
							MarkdownDescription: "The Energy Efficient Ethernet status of the switch port.",
							Computed:            true,
						},
						"high_speed_enabled": schema.BoolAttribute{
							MarkdownDescription: "For C9500-32QC, whether or not the port is enabled for high speed.",
							Computed:            true,
						},
						"mirror_mode": schema.StringAttribute{
							MarkdownDescription: "The port mirror mode. Can be one of (`Destination port`, `Source port` or `Not mirroring traffic`).",
							Computed:            true,
						},
						"module_model": schema.StringAttribute{
							MarkdownDescription: "The model of the expansion module.",
							Computed:            true,
						},
						"module_serial": schema.StringAttribute{
							MarkdownDescription: "The serial of the module.",
							Computed:            true,
						},
						"module_slot": schema.Int64Attribute{
							MarkdownDescription: "The slot number of the module.",
							Computed:            true,
						},
						"profile_enabled": schema.BoolAttribute{
							MarkdownDescription: "When enabled, override this port`s configuration with a port profile.",
							Computed:            true,
						},
						"profile_id": schema.StringAttribute{
							MarkdownDescription: "When enabled, the ID of the port profile used to override the port`s configuration.",
							Computed:            true,
						},
						"profile_iname": schema.StringAttribute{
							MarkdownDescription: "When enabled, the IName of the profile.",
							Computed:            true,
						},
						"schedule_id": schema.StringAttribute{
							MarkdownDescription: "The ID of the port schedule.",
							Computed:            true,
						},
						"schedule_name": schema.StringAttribute{
							MarkdownDescription: "The name of the port schedule.",
							Computed:            true,
						},
						"stackwise_virtual_is_dual_active_detector": schema.BoolAttribute{
							MarkdownDescription: "For SVL devices, whether or not the port is used for Dual Active Detection.",
							Computed:            true,
						},
						"stackwise_virtual_is_stack_wise_virtual_link": schema.BoolAttribute{
							MarkdownDescription: "For SVL devices, whether or not the port is used for StackWise Virtual Link.",
							Computed:            true,
						},
						"link_negotiation_capabilities": schema.ListAttribute{
							MarkdownDescription: "Available link speeds for the switch port.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"mac_allow_list": schema.ListAttribute{
							MarkdownDescription: "Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when `accessPolicyType` is `MAC allow list`.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"sticky_mac_allow_list": schema.ListAttribute{
							MarkdownDescription: "The initial list of MAC addresses for sticky Mac allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: "The list of tags of the switch port.",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SwitchPortsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchPortsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceSwitchPorts

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "SwitchPortsDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "SwitchPortsDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
