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
						"adaptive_policy_group_id": schema.StringAttribute{
							MarkdownDescription: "Adaptive policy group ID this SSID is assigned to.",
							Computed:            true,
						},
						"admin_splash_url": schema.StringAttribute{
							MarkdownDescription: "URL for the admin splash page",
							Computed:            true,
						},
						"adult_content_filtering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Boolean indicating whether or not adult content will be blocked. Only present when ipAssignmentMode is `NAT mode`.",
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
						"default_vlan_id": schema.Int64Attribute{
							MarkdownDescription: "The default VLAN ID used for `all other APs`. Only present when the ipAssignmentMode is `Bridge mode` or `Layer 3 roaming` and VLAN tagging is enabled.",
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
						"enterprise_admin_access": schema.StringAttribute{
							MarkdownDescription: "Whether or not an SSID is accessible by `enterprise` administrators. Only present when the organization supports enterprise admins.",
							Computed:            true,
						},
						"ip_assignment_mode": schema.StringAttribute{
							MarkdownDescription: "The client IP assignment mode",
							Computed:            true,
						},
						"lan_isolation_enabled": schema.BoolAttribute{
							MarkdownDescription: "Boolean indicating whether Layer 2 LAN isolation should be enabled or disabled. Only configurable when ipAssignmentMode is `Bridge mode`.",
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
						"min_bitrate": schema.Float64Attribute{
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
						"psk": schema.StringAttribute{
							MarkdownDescription: "The passkey for the SSID. This param is only valid if the authMode is `psk`. Read-only admins receive a masked value. Only present when authMode is `psk`.",
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
						"radius_authentication_nas_id": schema.StringAttribute{
							MarkdownDescription: "The configured NAS identifier used for RADIUS authentication. May include custom text and template variables. Only present when configured and the SSID uses RADIUS-based authentication.",
							Computed:            true,
						},
						"radius_called_station_id": schema.StringAttribute{
							MarkdownDescription: "The template of the called station identifier to be used for RADIUS (ex. $NODE_MAC$:$VAP_NUM$). Only present when configured and the SSID uses RADIUS-based auth.",
							Computed:            true,
						},
						"radius_coa_enabled": schema.BoolAttribute{
							MarkdownDescription: "If true, Meraki devices will act as a RADIUS Dynamic Authorization Server and will respond to RADIUS Change-of-Authorization and Disconnect messages sent by the RADIUS server. Only present when the SSID uses RADIUS-based auth.",
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
						"radius_fallback_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether or not higher priority RADIUS servers should be retried after 60 seconds. Only present when the SSID uses RADIUS-based auth.",
							Computed:            true,
						},
						"radius_load_balancing_policy": schema.StringAttribute{
							MarkdownDescription: "Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts",
							Computed:            true,
						},
						"radius_override": schema.BoolAttribute{
							MarkdownDescription: "If true, the RADIUS response can override VLAN tag. This is not valid when ipAssignmentMode is `NAT mode`. Only present when the SSID uses RADIUS-based auth and the network supports group policies.",
							Computed:            true,
						},
						"radius_proxy_enabled": schema.BoolAttribute{
							MarkdownDescription: "If true, Meraki devices will proxy RADIUS messages through the Meraki cloud to the configured RADIUS auth and accounting servers. Only present when the SSID uses RADIUS-based auth.",
							Computed:            true,
						},
						"radius_server_attempts_limit": schema.Int64Attribute{
							MarkdownDescription: "The maximum number of transmit attempts after which a RADIUS server is failed over (must be between 1-5). Only present when the SSID uses RADIUS-based auth.",
							Computed:            true,
						},
						"radius_server_timeout": schema.Int64Attribute{
							MarkdownDescription: "The amount of time for which a RADIUS client waits for a reply from the RADIUS server (must be between 1-10 seconds). Only present when the SSID uses RADIUS-based auth.",
							Computed:            true,
						},
						"radius_testing_enabled": schema.BoolAttribute{
							MarkdownDescription: "If true, Meraki devices will periodically send Access-Request messages to configured RADIUS servers using identity `meraki_8021x_test` to ensure that the RADIUS servers are reachable. Only present when the SSID is configured to use RADIUS.",
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
						"use_vlan_tagging": schema.BoolAttribute{
							MarkdownDescription: "Whether or not traffic should be directed to use specific VLANs. Only present when the ipAssignmentMode is `Bridge mode` or `Layer 3 roaming`.",
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
						"wifi_personal_network_enabled": schema.BoolAttribute{
							MarkdownDescription: "Boolean indicating whether or not Wi-Fi Personal Network is enabled for this SSID. Only present when the network supports Wi-Fi Personal Network.",
							Computed:            true,
						},
						"wlan_identifier": schema.Int64Attribute{
							MarkdownDescription: "WLAN identifier for the SSID. Only present for Meraki admin users.",
							Computed:            true,
						},
						"wpa_encryption_mode": schema.StringAttribute{
							MarkdownDescription: "The types of WPA encryption",
							Computed:            true,
						},
						"campus_gateway_cluster_id": schema.StringAttribute{
							MarkdownDescription: "The campus gateway cluster ID used when the ipAssignmentMode is `Campus Gateway`.",
							Computed:            true,
						},
						"dns_rewrite_enabled": schema.BoolAttribute{
							MarkdownDescription: "Boolean indicating whether or not DNS server rewrite is enabled. If disabled, upstream DNS will be used.",
							Computed:            true,
						},
						"dns_rewrite_dns_custom_nameservers": schema.ListAttribute{
							MarkdownDescription: "User specified DNS servers (up to two servers)",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"dot11k_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether 802.11k is enabled or not.",
							Computed:            true,
						},
						"dot11r_adaptive": schema.BoolAttribute{
							MarkdownDescription: "Whether 802.11r is adaptive or not.",
							Computed:            true,
						},
						"dot11r_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether 802.11r is enabled or not.",
							Computed:            true,
						},
						"dot11w_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether 802.11w is enabled or not.",
							Computed:            true,
						},
						"dot11w_required": schema.BoolAttribute{
							MarkdownDescription: "Whether 802.11w is required or not.",
							Computed:            true,
						},
						"gre_client_isolation": schema.BoolAttribute{
							MarkdownDescription: "Whether or not client isolation is enabled.",
							Computed:            true,
						},
						"gre_key": schema.Int64Attribute{
							MarkdownDescription: "Optional numerical identifier that will add the GRE key field to the GRE header. Used to identify an individual traffic flow within a tunnel.",
							Computed:            true,
						},
						"gre_concentrator_host": schema.StringAttribute{
							MarkdownDescription: "The EoGRE concentrator`s IP or FQDN.",
							Computed:            true,
						},
						"local_auth_fallback_cache_timeout": schema.Int64Attribute{
							MarkdownDescription: "The duration (in seconds) for which auths are cached.",
							Computed:            true,
						},
						"local_auth_fallback_enabled": schema.BoolAttribute{
							MarkdownDescription: "If true, MR devices will cache authentication credentials for EAP-TLS or for MAC based authentication.",
							Computed:            true,
						},
						"local_auth_fallback_server_ca_certificate_contents": schema.StringAttribute{
							MarkdownDescription: "The contents of the Server CA Certificate. Must be in PEM or DER format.",
							Computed:            true,
						},
						"named_vlans_radius_guest_vlan_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether or not RADIUS guest named VLAN is enabled.",
							Computed:            true,
						},
						"named_vlans_radius_guest_vlan_name": schema.StringAttribute{
							MarkdownDescription: "RADIUS guest VLAN name.",
							Computed:            true,
						},
						"named_vlans_tagging_default_vlan_name": schema.StringAttribute{
							MarkdownDescription: "The default VLAN name used to tag traffic in the absence of a matching AP tag.",
							Computed:            true,
						},
						"named_vlans_tagging_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether or not traffic should be directed to use specific VLAN names.",
							Computed:            true,
						},
						"named_vlans_tagging_by_ap_tags": schema.ListNestedAttribute{
							MarkdownDescription: "The list of AP tags and VLAN names used for named VLAN tagging.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"vlan_name": schema.StringAttribute{
										MarkdownDescription: "VLAN name that will be used to tag traffic.",
										Computed:            true,
									},
									"tags": schema.ListAttribute{
										MarkdownDescription: "List of AP tags.",
										ElementType:         types.StringType,
										Computed:            true,
									},
								},
							},
						},
						"oauth_allowed_domains": schema.ListAttribute{
							MarkdownDescription: "The list of domains allowed access to the network.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"security_encryption_akms": schema.ListAttribute{
							MarkdownDescription: "Authentication key management systems used by the SSID.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"security_encryption_ciphers": schema.ListAttribute{
							MarkdownDescription: "Encryption ciphers used by the SSID.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"speed_burst_enabled": schema.BoolAttribute{
							MarkdownDescription: "Boolean indicating whether or not to allow users to temporarily exceed the bandwidth limit for short periods while still keeping them under the bandwidth limit over time.",
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
									"id": schema.StringAttribute{
										MarkdownDescription: "Unique identifier of the RADIUS server",
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
									"radsec_enabled": schema.BoolAttribute{
										MarkdownDescription: "Use RADSEC (TLS over TCP) to connect to this RADIUS server. Requires radiusProxyEnabled.",
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
