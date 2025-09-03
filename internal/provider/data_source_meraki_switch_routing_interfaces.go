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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SwitchRoutingInterfacesDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchRoutingInterfacesDataSource{}
)

func NewSwitchRoutingInterfacesDataSource() datasource.DataSource {
	return &SwitchRoutingInterfacesDataSource{}
}

type SwitchRoutingInterfacesDataSource struct {
	client *meraki.Client
}

func (d *SwitchRoutingInterfacesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_routing_interfaces"
}

func (d *SwitchRoutingInterfacesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Switch Routing Interface` configuration in bulk.").String,

		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: "Switch serial",
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
						"default_gateway": schema.StringAttribute{
							MarkdownDescription: "The next hop for any traffic that isn`t going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface. Required if this is the first IPv4 interface.",
							Computed:            true,
						},
						"interface_ip": schema.StringAttribute{
							MarkdownDescription: "The IP address this switch will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch`s management IP.",
							Computed:            true,
						},
						"mode": schema.StringAttribute{
							MarkdownDescription: "L3 Interface mode, can be one of `vlan`, `routed` or `loopback`. Default is `vlan`. CS 17.18 or higher is required for `routed` mode.",
							Computed:            true,
						},
						"multicast_routing": schema.StringAttribute{
							MarkdownDescription: "Enable multicast support if, multicast routing between VLANs is required. Options are: `disabled`, `enabled` or `IGMP snooping querier`. Default is `disabled`.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "A friendly name or description for the interface or VLAN.",
							Computed:            true,
						},
						"subnet": schema.StringAttribute{
							MarkdownDescription: "The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).",
							Computed:            true,
						},
						"switch_port_id": schema.StringAttribute{
							MarkdownDescription: "Switch Port ID when in Routed mode (CS 17.18 or higher required)",
							Computed:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: "The VLAN this routed interface is on. VLAN must be between 1 and 4094.",
							Computed:            true,
						},
						"ipv6_address": schema.StringAttribute{
							MarkdownDescription: "The IPv6 address of the interface. Required if assignmentMode is `static`. Must not be included if assignmentMode is `eui-64`.",
							Computed:            true,
						},
						"ipv6_assignment_mode": schema.StringAttribute{
							MarkdownDescription: "The IPv6 assignment mode for the interface. Can be either `eui-64` or `static`.",
							Computed:            true,
						},
						"ipv6_gateway": schema.StringAttribute{
							MarkdownDescription: "The IPv6 default gateway of the interface. Required if prefix is defined and this is the first interface with IPv6 configured for the switch.",
							Computed:            true,
						},
						"ipv6_prefix": schema.StringAttribute{
							MarkdownDescription: "The IPv6 prefix of the interface. Required if IPv6 object is included.",
							Computed:            true,
						},
						"ospf_settings_area": schema.StringAttribute{
							MarkdownDescription: "The OSPF area to which this interface should belong. Can be either `ospfDisabled` or the identifier of an existing OSPF area. Defaults to `ospfDisabled`.",
							Computed:            true,
						},
						"ospf_settings_cost": schema.Int64Attribute{
							MarkdownDescription: "The path cost for this interface. Defaults to 1, but can be increased up to 65535 to give lower priority.",
							Computed:            true,
						},
						"ospf_settings_is_passive_enabled": schema.BoolAttribute{
							MarkdownDescription: "When enabled, OSPF will not run on the interface, but the subnet will still be advertised.",
							Computed:            true,
						},
						"ospf_settings_network_type": schema.StringAttribute{
							MarkdownDescription: "OSPF network type",
							Computed:            true,
						},
						"vrf_name": schema.StringAttribute{
							MarkdownDescription: "The name of the VRF this interface belongs to.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SwitchRoutingInterfacesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchRoutingInterfacesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceSwitchRoutingInterfaces

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "SwitchRoutingInterfacesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "SwitchRoutingInterfacesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
