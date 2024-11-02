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
	_ datasource.DataSource              = &WirelessSSIDVPNDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessSSIDVPNDataSource{}
)

func NewWirelessSSIDVPNDataSource() datasource.DataSource {
	return &WirelessSSIDVPNDataSource{}
}

type WirelessSSIDVPNDataSource struct {
	client *meraki.Client
}

func (d *WirelessSSIDVPNDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssid_vpn"
}

func (d *WirelessSSIDVPNDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Wireless SSID VPN` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: "Wireless SSID number",
				Required:            true,
			},
			"concentrator_network_id": schema.StringAttribute{
				MarkdownDescription: "The NAT ID of the concentrator that should be set.",
				Computed:            true,
			},
			"concentrator_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "The VLAN that should be tagged for the concentrator.",
				Computed:            true,
			},
			"failover_heartbeat_interval": schema.Int64Attribute{
				MarkdownDescription: "Idle timer interval in seconds.",
				Computed:            true,
			},
			"failover_idle_timeout": schema.Int64Attribute{
				MarkdownDescription: "Idle timer timeout in seconds.",
				Computed:            true,
			},
			"failover_request_ip": schema.StringAttribute{
				MarkdownDescription: "IP addressed reserved on DHCP server where SSID will terminate.",
				Computed:            true,
			},
			"split_tunnel_enabled": schema.BoolAttribute{
				MarkdownDescription: "If true, VPN split tunnel is enabled.",
				Computed:            true,
			},
			"split_tunnel_rules": schema.ListNestedAttribute{
				MarkdownDescription: "List of VPN split tunnel rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"comment": schema.StringAttribute{
							MarkdownDescription: "Description for this split tunnel rule (optional).",
							Computed:            true,
						},
						"dest_cidr": schema.StringAttribute{
							MarkdownDescription: "Destination for this split tunnel rule. IP address, fully-qualified domain names (FQDN) or `any`.",
							Computed:            true,
						},
						"dest_port": schema.StringAttribute{
							MarkdownDescription: "Destination port for this split tunnel rule, (integer in the range 1-65535), or `any`.",
							Computed:            true,
						},
						"policy": schema.StringAttribute{
							MarkdownDescription: "Traffic policy specified for this split tunnel rule, `allow` or `deny`.",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol for this split tunnel rule.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WirelessSSIDVPNDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessSSIDVPNDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessSSIDVPN

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
	config.Id = config.Number

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
