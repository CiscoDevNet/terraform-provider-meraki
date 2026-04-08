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
	_ datasource.DataSource              = &WirelessSSIDsDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessSSIDsDataSource{}
)

func NewWirelessSSIDsDataSource() datasource.DataSource {
	return &WirelessSSIDsDataSource{}
}

type WirelessSSIDsDataSource struct {
	client *meraki.Client
}

func (d *WirelessSSIDsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssids"
}

func (d *WirelessSSIDsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Wireless SSIDs` configuration in bulk.").String,

		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
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
						"admin_splash_url": schema.StringAttribute{
							MarkdownDescription: "URL for the admin splash page",
							Computed:            true,
						},
						"auth_mode": schema.StringAttribute{
							MarkdownDescription: "The association control method for the SSID",
							Computed:            true,
						},
						"available_on_all_aps": schema.BoolAttribute{
							MarkdownDescription: "Whether all APs broadcast the SSID or if it`s restricted to APs matching any availability tags",
							Computed:            true,
						},
						"band_selection": schema.StringAttribute{
							MarkdownDescription: "The client-serving radio frequencies of this SSID in the default indoor RF profile",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether or not the SSID is enabled",
							Computed:            true,
						},
						"encryption_mode": schema.StringAttribute{
							MarkdownDescription: "The psk encryption mode for the SSID",
							Computed:            true,
						},
						"ip_assignment_mode": schema.StringAttribute{
							MarkdownDescription: "The client IP assignment mode",
							Computed:            true,
						},
						"local_auth": schema.BoolAttribute{
							MarkdownDescription: "Extended local auth flag for Enterprise NAC",
							Computed:            true,
						},
						"mandatory_dhcp_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether clients connecting to this SSID must use the IP address assigned by the DHCP server",
							Computed:            true,
						},
						"min_bitrate": schema.Int64Attribute{
							MarkdownDescription: "The minimum bitrate in Mbps of this SSID in the default indoor RF profile",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the SSID",
							Computed:            true,
						},
						"number": schema.Int64Attribute{
							MarkdownDescription: "Unique identifier of the SSID",
							Computed:            true,
						},
						"per_client_bandwidth_limit_down": schema.Int64Attribute{
							MarkdownDescription: "The download bandwidth limit in Kbps. (0 represents no limit.)",
							Computed:            true,
						},
						"per_client_bandwidth_limit_up": schema.Int64Attribute{
							MarkdownDescription: "The upload bandwidth limit in Kbps. (0 represents no limit.)",
							Computed:            true,
						},
						"per_ssid_bandwidth_limit_down": schema.Int64Attribute{
							MarkdownDescription: "The total download bandwidth limit in Kbps (0 represents no limit)",
							Computed:            true,
						},
						"per_ssid_bandwidth_limit_up": schema.Int64Attribute{
							MarkdownDescription: "The total upload bandwidth limit in Kbps (0 represents no limit)",
							Computed:            true,
						},
						"radius_accounting_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether or not RADIUS accounting is enabled",
							Computed:            true,
						},
						"radius_attribute_for_group_policies": schema.StringAttribute{
							MarkdownDescription: "RADIUS attribute used to look up group policies",
							Computed:            true,
						},
						"radius_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether RADIUS authentication is enabled",
							Computed:            true,
						},
						"radius_failover_policy": schema.StringAttribute{
							MarkdownDescription: "Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable",
							Computed:            true,
						},
						"radius_load_balancing_policy": schema.StringAttribute{
							MarkdownDescription: "Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts",
							Computed:            true,
						},
						"splash_page": schema.StringAttribute{
							MarkdownDescription: "The type of splash page for the SSID",
							Computed:            true,
						},
						"splash_timeout": schema.StringAttribute{
							MarkdownDescription: "Splash page timeout",
							Computed:            true,
						},
						"ssid_admin_accessible": schema.BoolAttribute{
							MarkdownDescription: "SSID Administrator access status",
							Computed:            true,
						},
						"visible": schema.BoolAttribute{
							MarkdownDescription: "Whether the SSID is advertised or hidden by the AP",
							Computed:            true,
						},
						"walled_garden_enabled": schema.BoolAttribute{
							MarkdownDescription: "Allow users to access a configurable list of IP ranges prior to sign-on",
							Computed:            true,
						},
						"wpa_encryption_mode": schema.StringAttribute{
							MarkdownDescription: "The types of WPA encryption",
							Computed:            true,
						},
						"access_control_clients_blocked_from_using_lan": schema.BoolAttribute{
							MarkdownDescription: "Whether clients are blocked from using the LAN",
							Computed:            true,
						},
						"access_control_wired_clients_part_of_wifi_network": schema.BoolAttribute{
							MarkdownDescription: "Whether wired SSID clients are part of the Wi-Fi network",
							Computed:            true,
						},
						"access_control_bandwidth_limit": schema.StringAttribute{
							MarkdownDescription: "Defined SSID bandwidth limits",
							Computed:            true,
						},
						"access_control_client_ip_assignment_mode": schema.StringAttribute{
							MarkdownDescription: "Client IP assignment mode",
							Computed:            true,
						},
						"access_control_encryption_mode": schema.StringAttribute{
							MarkdownDescription: "Specifies the authentication and encryption type for the SSID",
							Computed:            true,
						},
						"access_control_splash_page_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether the splash page is enabled",
							Computed:            true,
						},
						"access_control_splash_page_theme": schema.StringAttribute{
							MarkdownDescription: "Name of the splash theme, or `n/a` if not applicable",
							Computed:            true,
						},
						"access_control_tunnel_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether tunneling is enabled",
							Computed:            true,
						},
						"access_control_tunnel_summary": schema.StringAttribute{
							MarkdownDescription: "Summary describing whether traffic is tunneled and to where",
							Computed:            true,
						},
						"access_control_vlan_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether VLAN tagging is enabled",
							Computed:            true,
						},
						"access_control_vlan_tag": schema.StringAttribute{
							MarkdownDescription: "VLAN tag applied to this SSID when tagging is enabled",
							Computed:            true,
						},
						"availability_tags": schema.ListAttribute{
							MarkdownDescription: "List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"radius_accounting_servers": schema.ListNestedAttribute{
							MarkdownDescription: "List of RADIUS accounting 802.1X servers to be used for authentication",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										MarkdownDescription: "Certificate used for authorization for the RADSEC Server",
										Computed:            true,
									},
									"host": schema.StringAttribute{
										MarkdownDescription: "IP address (or FQDN) to which the APs will send RADIUS accounting messages",
										Computed:            true,
									},
									"open_roaming_certificate_id": schema.Int64Attribute{
										MarkdownDescription: "The ID of the Openroaming Certificate attached to radius server",
										Computed:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: "Port on the RADIUS server that is listening for accounting messages",
										Computed:            true,
									},
								},
							},
						},
						"radius_servers": schema.ListNestedAttribute{
							MarkdownDescription: "List of RADIUS 802.1X servers to be used for authentication",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										MarkdownDescription: "Certificate used for authorization for the RADSEC Server",
										Computed:            true,
									},
									"host": schema.StringAttribute{
										MarkdownDescription: "IP address (or FQDN) of your RADIUS server",
										Computed:            true,
									},
									"open_roaming_certificate_id": schema.Int64Attribute{
										MarkdownDescription: "The ID of the Openroaming Certificate attached to radius server",
										Computed:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: "UDP port the RADIUS server listens on for Access-requests",
										Computed:            true,
									},
								},
							},
						},
						"walled_garden_ranges": schema.ListAttribute{
							MarkdownDescription: "Domain names and IP address ranges available in Walled Garden mode",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WirelessSSIDsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessSSIDsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceWirelessSSIDs

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "WirelessSSIDsDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "WirelessSSIDsDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
