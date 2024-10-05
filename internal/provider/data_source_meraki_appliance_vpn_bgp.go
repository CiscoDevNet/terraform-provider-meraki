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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ApplianceVPNBGPDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceVPNBGPDataSource{}
)

func NewApplianceVPNBGPDataSource() datasource.DataSource {
	return &ApplianceVPNBGPDataSource{}
}

type ApplianceVPNBGPDataSource struct {
	client *meraki.Client
}

func (d *ApplianceVPNBGPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_vpn_bgp"
}

func (d *ApplianceVPNBGPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Appliance VPN BGP` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"as_number": schema.Int64Attribute{
				MarkdownDescription: "An Autonomous System Number (ASN) is required if you are to run BGP and peer with another BGP Speaker outside of the Auto VPN domain. This ASN will be applied to the entire Auto VPN domain. The entire 4-byte ASN range is supported. So, the ASN must be an integer between 1 and 4294967295. When absent, this field is not updated. If no value exists then it defaults to 64512.",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean value to enable or disable the BGP configuration. When BGP is enabled, the asNumber (ASN) will be autopopulated with the preconfigured ASN at other Hubs or a default value if there is no ASN configured.",
				Computed:            true,
			},
			"ibgp_hold_timer": schema.Int64Attribute{
				MarkdownDescription: "The iBGP holdtimer in seconds. The iBGP holdtimer must be an integer between 12 and 240. When absent, this field is not updated. If no value exists then it defaults to 240.",
				Computed:            true,
			},
			"neighbors": schema.ListNestedAttribute{
				MarkdownDescription: "List of BGP neighbors. This list replaces the existing set of neighbors. When absent, this field is not updated.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allow_transit": schema.BoolAttribute{
							MarkdownDescription: "When this feature is on, the Meraki device will advertise routes learned from other Autonomous Systems, thereby allowing traffic between Autonomous Systems to transit this AS. When absent, it defaults to false.",
							Computed:            true,
						},
						"ebgp_hold_timer": schema.Int64Attribute{
							MarkdownDescription: "The eBGP hold timer in seconds for each neighbor. The eBGP hold timer must be an integer between 12 and 240.",
							Computed:            true,
						},
						"ebgp_multihop": schema.Int64Attribute{
							MarkdownDescription: "Configure this if the neighbor is not adjacent. The eBGP multi-hop must be an integer between 1 and 255.",
							Computed:            true,
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: "The IPv4 address of the neighbor",
							Computed:            true,
						},
						"next_hop_ip": schema.StringAttribute{
							MarkdownDescription: "The IPv4 address of the remote BGP peer that will establish a TCP session with the local MX.",
							Computed:            true,
						},
						"receive_limit": schema.Int64Attribute{
							MarkdownDescription: "The receive limit is the maximum number of routes that can be received from any BGP peer. The receive limit must be an integer between 0 and 2147483647. When absent, it defaults to 0.",
							Computed:            true,
						},
						"remote_as_number": schema.Int64Attribute{
							MarkdownDescription: "Remote ASN of the neighbor. The remote ASN must be an integer between 1 and 4294967295.",
							Computed:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: "The output interface for peering with the remote BGP peer. Valid values are: `wan1`, `wan2` or `vlan{VLAN ID}`(e.g. `vlan123`).",
							Computed:            true,
						},
						"authentication_password": schema.StringAttribute{
							MarkdownDescription: "Password to configure MD5 authentication between BGP peers.",
							Computed:            true,
						},
						"ipv6_address": schema.StringAttribute{
							MarkdownDescription: "The IPv6 address of the neighbor.",
							Computed:            true,
						},
						"ttl_security_enabled": schema.BoolAttribute{
							MarkdownDescription: "Boolean value to enable or disable BGP TTL security.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ApplianceVPNBGPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceVPNBGPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplianceVPNBGP

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res meraki.Res
	var err error

	if !res.Exists() {
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)
	config.Id = config.NetworkId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
