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
	_ datasource.DataSource              = &WirelessSSIDDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessSSIDDataSource{}
)

func NewWirelessSSIDDataSource() datasource.DataSource {
	return &WirelessSSIDDataSource{}
}

type WirelessSSIDDataSource struct {
	client *meraki.Client
}

func (d *WirelessSSIDDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssid"
}

func (d *WirelessSSIDDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Wireless SSID` configuration.").String,

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
			"adaptive_policy_group_id": schema.StringAttribute{
				MarkdownDescription: "Adaptive policy group ID this SSID is assigned to.",
				Computed:            true,
			},
			"adult_content_filtering_enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether or not adult content will be blocked",
				Computed:            true,
			},
			"auth_mode": schema.StringAttribute{
				MarkdownDescription: "The association control method for the SSID (`open`, `open-enhanced`, `psk`, `open-with-radius`, `open-with-nac`, `8021x-meraki`, `8021x-nac`, `8021x-radius`, `8021x-google`, `8021x-entra`, `8021x-localradius`, `ipsk-with-radius`, `ipsk-without-radius` or `ipsk-with-nac`)",
				Computed:            true,
			},
			"available_on_all_aps": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether all APs should broadcast the SSID or if it should be restricted to APs matching any availability tags. Can only be false if the SSID has availability tags.",
				Computed:            true,
			},
			"band_selection": schema.StringAttribute{
				MarkdownDescription: "The client-serving radio frequencies of this SSID in the default indoor RF profile. (`Dual band operation`, `5 GHz band only` or `Dual band operation with Band Steering`)",
				Computed:            true,
			},
			"concentrator_network_id": schema.StringAttribute{
				MarkdownDescription: "The concentrator to use when the ipAssignmentMode is `Layer 3 roaming with a concentrator` or `VPN`.",
				Computed:            true,
			},
			"default_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "The default VLAN ID used for `all other APs`. This param is only valid when the ipAssignmentMode is `Bridge mode` or `Layer 3 roaming`",
				Computed:            true,
			},
			"disassociate_clients_on_vpn_failover": schema.BoolAttribute{
				MarkdownDescription: "Disassociate clients when `VPN` concentrator failover occurs in order to trigger clients to re-associate and generate new DHCP requests. This param is only valid if ipAssignmentMode is `VPN`.",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the SSID is enabled",
				Computed:            true,
			},
			"encryption_mode": schema.StringAttribute{
				MarkdownDescription: "The psk encryption mode for the SSID (`wep` or `wpa`). This param is only valid if the authMode is `psk`",
				Computed:            true,
			},
			"enterprise_admin_access": schema.StringAttribute{
				MarkdownDescription: "Whether or not an SSID is accessible by `enterprise` administrators (`access disabled` or `access enabled`)",
				Computed:            true,
			},
			"ip_assignment_mode": schema.StringAttribute{
				MarkdownDescription: "The client IP assignment mode (`NAT mode`, `Bridge mode`, `Layer 3 roaming`, `Ethernet over GRE`, `Layer 3 roaming with a concentrator` or `VPN`)",
				Computed:            true,
			},
			"lan_isolation_enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether Layer 2 LAN isolation should be enabled or disabled. Only configurable when ipAssignmentMode is `Bridge mode`.",
				Computed:            true,
			},
			"mandatory_dhcp_enabled": schema.BoolAttribute{
				MarkdownDescription: "If true, Mandatory DHCP will enforce that clients connecting to this SSID must use the IP address assigned by the DHCP server. Clients who use a static IP address won`t be able to associate.",
				Computed:            true,
			},
			"min_bitrate": schema.Float64Attribute{
				MarkdownDescription: "The minimum bitrate in Mbps of this SSID in the default indoor RF profile. (`1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`)",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the SSID",
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
				MarkdownDescription: "The total download bandwidth limit in Kbps. (0 represents no limit.)",
				Computed:            true,
			},
			"per_ssid_bandwidth_limit_up": schema.Int64Attribute{
				MarkdownDescription: "The total upload bandwidth limit in Kbps. (0 represents no limit.)",
				Computed:            true,
			},
			"psk": schema.StringAttribute{
				MarkdownDescription: "The passkey for the SSID. This param is only valid if the authMode is `psk`",
				Computed:            true,
			},
			"radius_accounting_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not RADIUS accounting is enabled. This param is only valid if the authMode is `open-with-radius`, `8021x-radius` or `ipsk-with-radius`",
				Computed:            true,
			},
			"radius_accounting_interim_interval": schema.Int64Attribute{
				MarkdownDescription: "The interval (in seconds) in which accounting information is updated and sent to the RADIUS accounting server.",
				Computed:            true,
			},
			"radius_accounting_start_delay": schema.Int64Attribute{
				MarkdownDescription: "The delay (in seconds) before sending the first RADIUS accounting start message. Must be between 0 and 59 seconds.",
				Computed:            true,
			},
			"radius_attribute_for_group_policies": schema.StringAttribute{
				MarkdownDescription: "Specify the RADIUS attribute used to look up group policies (`Filter-Id`, `Reply-Message`, `Airespace-ACL-Name` or `Aruba-User-Role`). Access points must receive this attribute in the RADIUS Access-Accept message",
				Computed:            true,
			},
			"radius_authentication_nas_id": schema.StringAttribute{
				MarkdownDescription: "The template of the NAS identifier to be used for RADIUS authentication (ex. $NODE_MAC$:$VAP_NUM$).",
				Computed:            true,
			},
			"radius_called_station_id": schema.StringAttribute{
				MarkdownDescription: "The template of the called station identifier to be used for RADIUS (ex. $NODE_MAC$:$VAP_NUM$).",
				Computed:            true,
			},
			"radius_coa_enabled": schema.BoolAttribute{
				MarkdownDescription: "If true, Meraki devices will act as a RADIUS Dynamic Authorization Server and will respond to RADIUS Change-of-Authorization and Disconnect messages sent by the RADIUS server.",
				Computed:            true,
			},
			"radius_failover_policy": schema.StringAttribute{
				MarkdownDescription: "This policy determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable (`Deny access` or `Allow access`)",
				Computed:            true,
			},
			"radius_fallback_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not higher priority RADIUS servers should be retried after 60 seconds.",
				Computed:            true,
			},
			"radius_guest_vlan_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not RADIUS Guest VLAN is enabled. This param is only valid if the authMode is `open-with-radius` and addressing mode is not set to `isolated` or `nat` mode",
				Computed:            true,
			},
			"radius_guest_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "VLAN ID of the RADIUS Guest VLAN. This param is only valid if the authMode is `open-with-radius` and addressing mode is not set to `isolated` or `nat` mode",
				Computed:            true,
			},
			"radius_load_balancing_policy": schema.StringAttribute{
				MarkdownDescription: "This policy determines which RADIUS server will be contacted first in an authentication attempt and the ordering of any necessary retry attempts (`Strict priority order` or `Round robin`)",
				Computed:            true,
			},
			"radius_override": schema.BoolAttribute{
				MarkdownDescription: "If true, the RADIUS response can override VLAN tag. This is not valid when ipAssignmentMode is `NAT mode`.",
				Computed:            true,
			},
			"radius_proxy_enabled": schema.BoolAttribute{
				MarkdownDescription: "If true, Meraki devices will proxy RADIUS messages through the Meraki cloud to the configured RADIUS auth and accounting servers.",
				Computed:            true,
			},
			"radius_server_attempts_limit": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of transmit attempts after which a RADIUS server is failed over (must be between 1-5).",
				Computed:            true,
			},
			"radius_server_timeout": schema.Int64Attribute{
				MarkdownDescription: "The amount of time for which a RADIUS client waits for a reply from the RADIUS server (must be between 1-10 seconds).",
				Computed:            true,
			},
			"radius_testing_enabled": schema.BoolAttribute{
				MarkdownDescription: "If true, Meraki devices will periodically send Access-Request messages to configured RADIUS servers using identity `meraki_8021x_test` to ensure that the RADIUS servers are reachable.",
				Computed:            true,
			},
			"secondary_concentrator_network_id": schema.StringAttribute{
				MarkdownDescription: "The secondary concentrator to use when the ipAssignmentMode is `VPN`. If configured, the APs will switch to using this concentrator if the primary concentrator is unreachable. This param is optional. (`disabled` represents no secondary concentrator.)",
				Computed:            true,
			},
			"splash_page": schema.StringAttribute{
				MarkdownDescription: "The type of splash page for the SSID (`None`, `Click-through splash page`, `Billing`, `Password-protected with Meraki RADIUS`, `Password-protected with custom RADIUS`, `Password-protected with Active Directory`, `Password-protected with LDAP`, `SMS authentication`, `Systems Manager Sentry`, `Facebook Wi-Fi`, `Google OAuth`, `Microsoft Entra ID`, `Sponsored guest`, `Cisco ISE` or `Google Apps domain`). This attribute is not supported for template children.",
				Computed:            true,
			},
			"use_vlan_tagging": schema.BoolAttribute{
				MarkdownDescription: "Whether or not traffic should be directed to use specific VLANs. This param is only valid if the ipAssignmentMode is `Bridge mode` or `Layer 3 roaming`",
				Computed:            true,
			},
			"visible": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether APs should advertise or hide this SSID. APs will only broadcast this SSID if set to true",
				Computed:            true,
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: "The VLAN ID used for VLAN tagging. This param is only valid when the ipAssignmentMode is `Layer 3 roaming with a concentrator` or `VPN`",
				Computed:            true,
			},
			"walled_garden_enabled": schema.BoolAttribute{
				MarkdownDescription: "Allow access to a configurable list of IP ranges, which users may access prior to sign-on.",
				Computed:            true,
			},
			"wpa_encryption_mode": schema.StringAttribute{
				MarkdownDescription: "The types of WPA encryption. (`WPA1 only`, `WPA1 and WPA2`, `WPA2 only`, `WPA3 Transition Mode`, `WPA3 only` or `WPA3 192-bit Security`)",
				Computed:            true,
			},
			"active_directory_credentials_logon_name": schema.StringAttribute{
				MarkdownDescription: "The logon name of the Active Directory account.",
				Computed:            true,
			},
			"active_directory_credentials_password": schema.StringAttribute{
				MarkdownDescription: "The password to the Active Directory user account.",
				Computed:            true,
			},
			"active_directory_servers": schema.ListNestedAttribute{
				MarkdownDescription: "The Active Directory servers to be used for authentication.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host": schema.StringAttribute{
							MarkdownDescription: "IP address (or FQDN) of your Active Directory server.",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "(Optional) UDP port the Active Directory server listens on. By default, uses port 3268.",
							Computed:            true,
						},
					},
				},
			},
			"dns_rewrite_enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether or not DNS server rewrite is enabled. If disabled, upstream DNS will be used",
				Computed:            true,
			},
			"dns_rewrite_dns_custom_nameservers": schema.ListAttribute{
				MarkdownDescription: "User specified DNS servers (up to two servers)",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"dot11r_adaptive": schema.BoolAttribute{
				MarkdownDescription: "(Optional) Whether 802.11r is adaptive or not.",
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
				MarkdownDescription: "(Optional) Whether 802.11w is required or not.",
				Computed:            true,
			},
			"gre_key": schema.Int64Attribute{
				MarkdownDescription: "Optional numerical identifier that will add the GRE key field to the GRE header. Used to identify an individual traffic flow within a tunnel.",
				Computed:            true,
			},
			"gre_concentrator_host": schema.StringAttribute{
				MarkdownDescription: "The EoGRE concentrator's IP or FQDN. This param is required when ipAssignmentMode is `Ethernet over GRE`.",
				Computed:            true,
			},
			"ldap_base_distinguished_name": schema.StringAttribute{
				MarkdownDescription: "The base distinguished name of users on the LDAP server.",
				Computed:            true,
			},
			"ldap_credentials_distinguished_name": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the LDAP user account (example: cn=user,dc=meraki,dc=com).",
				Computed:            true,
			},
			"ldap_credentials_password": schema.StringAttribute{
				MarkdownDescription: "The password of the LDAP user account.",
				Computed:            true,
			},
			"ldap_server_ca_certificate_contents": schema.StringAttribute{
				MarkdownDescription: "The contents of the CA certificate. Must be in PEM or DER format.",
				Computed:            true,
			},
			"ldap_servers": schema.ListNestedAttribute{
				MarkdownDescription: "The LDAP servers to be used for authentication.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host": schema.StringAttribute{
							MarkdownDescription: "IP address (or FQDN) of your LDAP server.",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "UDP port the LDAP server listens on.",
							Computed:            true,
						},
					},
				},
			},
			"local_auth_fallback_cache_timeout": schema.Int64Attribute{
				MarkdownDescription: "The duration (in seconds) for which auths are cached. The timeout is measured from the user`s most recent non-cached authentication to the network. Between 3600 (1 hour) and 86400 (1 day)",
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
			"local_radius_cache_timeout": schema.Int64Attribute{
				MarkdownDescription: "The duration (in seconds) for which LDAP and OCSP lookups are cached.",
				Computed:            true,
			},
			"local_radius_certificate_authentication_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not to use EAP-TLS certificate-based authentication to validate wireless clients.",
				Computed:            true,
			},
			"local_radius_certificate_authentication_ocsp_responder_url": schema.StringAttribute{
				MarkdownDescription: "(Optional) The URL of the OCSP responder to verify client certificate status.",
				Computed:            true,
			},
			"local_radius_certificate_authentication_use_ldap": schema.BoolAttribute{
				MarkdownDescription: "Whether or not to verify the certificate with LDAP.",
				Computed:            true,
			},
			"local_radius_certificate_authentication_use_ocsp": schema.BoolAttribute{
				MarkdownDescription: "Whether or not to verify the certificate with OCSP.",
				Computed:            true,
			},
			"local_radius_certificate_authentication_client_root_ca_certificate_contents": schema.StringAttribute{
				MarkdownDescription: "The contents of the Client CA Certificate. Must be in PEM or DER format.",
				Computed:            true,
			},
			"local_radius_password_authentication_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not to use EAP-TTLS/PAP or PEAP-GTC password-based authentication via LDAP lookup.",
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
				MarkdownDescription: "The list of AP tags and VLAN names used for named VLAN tagging. If an AP has a tag matching one in the list, then traffic on this SSID will be directed to use the VLAN name associated to the tag.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vlan_name": schema.StringAttribute{
							MarkdownDescription: "VLAN name that will be used to tag traffic.",
							Computed:            true,
						},
						"tags": schema.SetAttribute{
							MarkdownDescription: "List of AP tags.",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
			"oauth_allowed_domains": schema.SetAttribute{
				MarkdownDescription: "(Optional) The list of domains allowed access to the network.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"radius_radsec_tls_tunnel_timeout": schema.Int64Attribute{
				MarkdownDescription: "The interval (in seconds) to determines how long a TLS session can remain idle for a RADSec server before it is automatically terminated",
				Computed:            true,
			},
			"speed_burst_enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether or not to allow users to temporarily exceed the bandwidth limit for short periods while still keeping them under the bandwidth limit over time.",
				Computed:            true,
			},
			"ap_tags_and_vlan_ids": schema.ListNestedAttribute{
				MarkdownDescription: "The list of tags and VLAN IDs used for VLAN tagging. This param is only valid when the ipAssignmentMode is `Bridge mode` or `Layer 3 roaming`",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: "Numerical identifier that is assigned to the VLAN",
							Computed:            true,
						},
						"tags": schema.SetAttribute{
							MarkdownDescription: "Array of AP tags",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
			"availability_tags": schema.SetAttribute{
				MarkdownDescription: "Accepts a list of tags for this SSID. If availableOnAllAps is false, then the SSID will only be broadcast by APs with tags matching any of the tags in this list.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"radius_accounting_servers": schema.ListNestedAttribute{
				MarkdownDescription: "The RADIUS accounting 802.1X servers to be used for authentication. This param is only valid if the authMode is `open-with-radius`, `8021x-radius` or `ipsk-with-radius` and radiusAccountingEnabled is `true`",
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
						"port": schema.Int64Attribute{
							MarkdownDescription: "Port on the RADIUS server that is listening for accounting messages",
							Computed:            true,
						},
						"radsec_enabled": schema.BoolAttribute{
							MarkdownDescription: "Use RADSEC (TLS over TCP) to connect to this RADIUS accounting server. Requires radiusProxyEnabled.",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "Shared key used to authenticate messages between the APs and RADIUS server",
							Computed:            true,
						},
					},
				},
			},
			"radius_servers": schema.ListNestedAttribute{
				MarkdownDescription: "The RADIUS 802.1X servers to be used for authentication. This param is only valid if the authMode is `open-with-radius`, `8021x-radius` or `ipsk-with-radius`",
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
							MarkdownDescription: "The ID of the Openroaming Certificate attached to radius server.",
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
						"secret": schema.StringAttribute{
							MarkdownDescription: "RADIUS client shared secret",
							Computed:            true,
						},
					},
				},
			},
			"splash_guest_sponsor_domains": schema.SetAttribute{
				MarkdownDescription: "Array of valid sponsor email domains for sponsored guest splash type.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"walled_garden_ranges": schema.SetAttribute{
				MarkdownDescription: "Specify your walled garden by entering an array of addresses, ranges using CIDR notation, domain names, and domain wildcards (e.g. `192.168.1.1/24`, `192.168.37.10/32`, `www.yahoo.com`, `*.google.com`]). Meraki`s splash page is automatically included in your walled garden.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"radius_das_clients_ips": schema.SetAttribute{
				MarkdownDescription: "List of DAS (Dynamic Authorization Server) IPs. This is an unsupported attribute and is subject to breaking changes without prior notice.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"radius_das_clients_shared_secret": schema.StringAttribute{
				MarkdownDescription: "Shared secret for DAS (Dynamic Authorization Server). This is an unsupported attribute and is subject to breaking changes without prior notice.",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessSSIDDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessSSIDDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessSSID

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
