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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &WirelessSSIDsResource{}
	_ resource.ResourceWithImportState = &WirelessSSIDsResource{}
	_ resource.ResourceWithModifyPlan  = &WirelessSSIDsResource{}
)

func NewWirelessSSIDsResource() resource.Resource {
	return &WirelessSSIDsResource{}
}

type WirelessSSIDsResource struct {
	client *meraki.Client
}

func (r *WirelessSSIDsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssids"
}

func (r *WirelessSSIDsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Wireless SSID` configuration in bulk.").AddBulkResourceIds("number").String,

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
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The list of items",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"number": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Wireless SSID number").String,
							Required:            true,
						},
						"adaptive_policy_group_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Adaptive policy group ID this SSID is assigned to.").String,
							Optional:            true,
						},
						"adult_content_filtering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether or not adult content will be blocked").String,
							Optional:            true,
						},
						"auth_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The association control method for the SSID (`open`, `open-enhanced`, `psk`, `open-with-radius`, `open-with-nac`, `8021x-meraki`, `8021x-nac`, `8021x-radius`, `8021x-google`, `8021x-entra`, `8021x-localradius`, `ipsk-with-radius`, `ipsk-without-radius` or `ipsk-with-nac`)").AddStringEnumDescription("8021x-entra", "8021x-google", "8021x-localradius", "8021x-meraki", "8021x-nac", "8021x-radius", "ipsk-with-nac", "ipsk-with-radius", "ipsk-without-radius", "open", "open-enhanced", "open-with-nac", "open-with-radius", "psk").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("8021x-entra", "8021x-google", "8021x-localradius", "8021x-meraki", "8021x-nac", "8021x-radius", "ipsk-with-nac", "ipsk-with-radius", "ipsk-without-radius", "open", "open-enhanced", "open-with-nac", "open-with-radius", "psk"),
							},
						},
						"available_on_all_aps": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether all APs should broadcast the SSID or if it should be restricted to APs matching any availability tags. Can only be false if the SSID has availability tags.").String,
							Optional:            true,
						},
						"band_selection": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The client-serving radio frequencies of this SSID in the default indoor RF profile. (`Dual band operation`, `5 GHz band only` or `Dual band operation with Band Steering`)").String,
							Optional:            true,
						},
						"concentrator_network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The concentrator to use when the ipAssignmentMode is `Layer 3 roaming with a concentrator` or `VPN`.").String,
							Optional:            true,
						},
						"default_vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The default VLAN ID used for `all other APs`. This param is only valid when the ipAssignmentMode is `Bridge mode` or `Layer 3 roaming`").String,
							Optional:            true,
						},
						"disassociate_clients_on_vpn_failover": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Disassociate clients when `VPN` concentrator failover occurs in order to trigger clients to re-associate and generate new DHCP requests. This param is only valid if ipAssignmentMode is `VPN`.").String,
							Optional:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not the SSID is enabled").String,
							Optional:            true,
						},
						"encryption_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The psk encryption mode for the SSID (`wep` or `wpa`). This param is only valid if the authMode is `psk`").AddStringEnumDescription("open", "wep", "wpa", "wpa-eap").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("open", "wep", "wpa", "wpa-eap"),
							},
						},
						"enterprise_admin_access": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not an SSID is accessible by `enterprise` administrators (`access disabled` or `access enabled`)").AddStringEnumDescription("access disabled", "access enabled").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("access disabled", "access enabled"),
							},
						},
						"ip_assignment_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The client IP assignment mode (`NAT mode`, `Bridge mode`, `Layer 3 roaming`, `Ethernet over GRE`, `Layer 3 roaming with a concentrator` or `VPN`)").String,
							Optional:            true,
						},
						"lan_isolation_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether Layer 2 LAN isolation should be enabled or disabled. Only configurable when ipAssignmentMode is `Bridge mode`.").String,
							Optional:            true,
						},
						"mandatory_dhcp_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If true, Mandatory DHCP will enforce that clients connecting to this SSID must use the IP address assigned by the DHCP server. Clients who use a static IP address won`t be able to associate.").String,
							Optional:            true,
						},
						"min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The minimum bitrate in Mbps of this SSID in the default indoor RF profile. (`1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`)").String,
							Optional:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the SSID").String,
							Required:            true,
						},
						"per_client_bandwidth_limit_down": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The download bandwidth limit in Kbps. (0 represents no limit.)").String,
							Optional:            true,
						},
						"per_client_bandwidth_limit_up": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The upload bandwidth limit in Kbps. (0 represents no limit.)").String,
							Optional:            true,
						},
						"per_ssid_bandwidth_limit_down": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The total download bandwidth limit in Kbps. (0 represents no limit.)").String,
							Optional:            true,
						},
						"per_ssid_bandwidth_limit_up": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The total upload bandwidth limit in Kbps. (0 represents no limit.)").String,
							Optional:            true,
						},
						"psk": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The passkey for the SSID. This param is only valid if the authMode is `psk`").String,
							Optional:            true,
						},
						"radius_accounting_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not RADIUS accounting is enabled. This param is only valid if the authMode is `open-with-radius`, `8021x-radius` or `ipsk-with-radius`").String,
							Optional:            true,
						},
						"radius_accounting_interim_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The interval (in seconds) in which accounting information is updated and sent to the RADIUS accounting server.").String,
							Optional:            true,
						},
						"radius_attribute_for_group_policies": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the RADIUS attribute used to look up group policies (`Filter-Id`, `Reply-Message`, `Airespace-ACL-Name` or `Aruba-User-Role`). Access points must receive this attribute in the RADIUS Access-Accept message").AddStringEnumDescription("Airespace-ACL-Name", "Aruba-User-Role", "Filter-Id", "Reply-Message").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Airespace-ACL-Name", "Aruba-User-Role", "Filter-Id", "Reply-Message"),
							},
						},
						"radius_authentication_nas_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The template of the NAS identifier to be used for RADIUS authentication (ex. $NODE_MAC$:$VAP_NUM$).").String,
							Optional:            true,
						},
						"radius_called_station_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The template of the called station identifier to be used for RADIUS (ex. $NODE_MAC$:$VAP_NUM$).").String,
							Optional:            true,
						},
						"radius_coa_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If true, Meraki devices will act as a RADIUS Dynamic Authorization Server and will respond to RADIUS Change-of-Authorization and Disconnect messages sent by the RADIUS server.").String,
							Optional:            true,
						},
						"radius_failover_policy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("This policy determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable (`Deny access` or `Allow access`)").AddStringEnumDescription("Allow access", "Deny access").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Allow access", "Deny access"),
							},
						},
						"radius_fallback_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not higher priority RADIUS servers should be retried after 60 seconds.").String,
							Optional:            true,
						},
						"radius_guest_vlan_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not RADIUS Guest VLAN is enabled. This param is only valid if the authMode is `open-with-radius` and addressing mode is not set to `isolated` or `nat` mode").String,
							Optional:            true,
						},
						"radius_guest_vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VLAN ID of the RADIUS Guest VLAN. This param is only valid if the authMode is `open-with-radius` and addressing mode is not set to `isolated` or `nat` mode").String,
							Optional:            true,
						},
						"radius_load_balancing_policy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("This policy determines which RADIUS server will be contacted first in an authentication attempt and the ordering of any necessary retry attempts (`Strict priority order` or `Round robin`)").AddStringEnumDescription("Round robin", "Strict priority order").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Round robin", "Strict priority order"),
							},
						},
						"radius_override": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If true, the RADIUS response can override VLAN tag. This is not valid when ipAssignmentMode is `NAT mode`.").String,
							Optional:            true,
						},
						"radius_proxy_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If true, Meraki devices will proxy RADIUS messages through the Meraki cloud to the configured RADIUS auth and accounting servers.").String,
							Optional:            true,
						},
						"radius_server_attempts_limit": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The maximum number of transmit attempts after which a RADIUS server is failed over (must be between 1-5).").String,
							Optional:            true,
						},
						"radius_server_timeout": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The amount of time for which a RADIUS client waits for a reply from the RADIUS server (must be between 1-10 seconds).").String,
							Optional:            true,
						},
						"radius_testing_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If true, Meraki devices will periodically send Access-Request messages to configured RADIUS servers using identity `meraki_8021x_test` to ensure that the RADIUS servers are reachable.").String,
							Optional:            true,
						},
						"secondary_concentrator_network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The secondary concentrator to use when the ipAssignmentMode is `VPN`. If configured, the APs will switch to using this concentrator if the primary concentrator is unreachable. This param is optional. (`disabled` represents no secondary concentrator.)").String,
							Optional:            true,
						},
						"splash_page": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of splash page for the SSID (`None`, `Click-through splash page`, `Billing`, `Password-protected with Meraki RADIUS`, `Password-protected with custom RADIUS`, `Password-protected with Active Directory`, `Password-protected with LDAP`, `SMS authentication`, `Systems Manager Sentry`, `Facebook Wi-Fi`, `Google OAuth`, `Microsoft Entra ID`, `Sponsored guest`, `Cisco ISE` or `Google Apps domain`). This attribute is not supported for template children.").AddStringEnumDescription("Billing", "Cisco ISE", "Click-through splash page", "Facebook Wi-Fi", "Google Apps domain", "Google OAuth", "Microsoft Entra ID", "None", "Password-protected with Active Directory", "Password-protected with LDAP", "Password-protected with Meraki RADIUS", "Password-protected with custom RADIUS", "SMS authentication", "Sponsored guest", "Systems Manager Sentry").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Billing", "Cisco ISE", "Click-through splash page", "Facebook Wi-Fi", "Google Apps domain", "Google OAuth", "Microsoft Entra ID", "None", "Password-protected with Active Directory", "Password-protected with LDAP", "Password-protected with Meraki RADIUS", "Password-protected with custom RADIUS", "SMS authentication", "Sponsored guest", "Systems Manager Sentry"),
							},
						},
						"use_vlan_tagging": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not traffic should be directed to use specific VLANs. This param is only valid if the ipAssignmentMode is `Bridge mode` or `Layer 3 roaming`").String,
							Optional:            true,
						},
						"visible": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether APs should advertise or hide this SSID. APs will only broadcast this SSID if set to true").String,
							Optional:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The VLAN ID used for VLAN tagging. This param is only valid when the ipAssignmentMode is `Layer 3 roaming with a concentrator` or `VPN`").String,
							Optional:            true,
						},
						"walled_garden_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Allow access to a configurable list of IP ranges, which users may access prior to sign-on.").String,
							Optional:            true,
						},
						"wpa_encryption_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The types of WPA encryption. (`WPA1 only`, `WPA1 and WPA2`, `WPA2 only`, `WPA3 Transition Mode`, `WPA3 only` or `WPA3 192-bit Security`)").AddStringEnumDescription("WPA1 and WPA2", "WPA1 only", "WPA2 only", "WPA3 192-bit Security", "WPA3 Transition Mode", "WPA3 only").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("WPA1 and WPA2", "WPA1 only", "WPA2 only", "WPA3 192-bit Security", "WPA3 Transition Mode", "WPA3 only"),
							},
						},
						"active_directory_credentials_logon_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The logon name of the Active Directory account.").String,
							Optional:            true,
						},
						"active_directory_credentials_password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The password to the Active Directory user account.").String,
							Optional:            true,
						},
						"active_directory_servers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The Active Directory servers to be used for authentication.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP address (or FQDN) of your Active Directory server.").String,
										Required:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("(Optional) UDP port the Active Directory server listens on. By default, uses port 3268.").String,
										Optional:            true,
									},
								},
							},
						},
						"dns_rewrite_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether or not DNS server rewrite is enabled. If disabled, upstream DNS will be used").String,
							Optional:            true,
						},
						"dns_rewrite_dns_custom_nameservers": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("User specified DNS servers (up to two servers)").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"dot11r_adaptive": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("(Optional) Whether 802.11r is adaptive or not.").String,
							Optional:            true,
						},
						"dot11r_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether 802.11r is enabled or not.").String,
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
						"gre_key": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Optional numerical identifier that will add the GRE key field to the GRE header. Used to identify an individual traffic flow within a tunnel.").String,
							Optional:            true,
						},
						"gre_concentrator_host": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The EoGRE concentrator's IP or FQDN. This param is required when ipAssignmentMode is `Ethernet over GRE`.").String,
							Optional:            true,
						},
						"ldap_base_distinguished_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The base distinguished name of users on the LDAP server.").String,
							Optional:            true,
						},
						"ldap_credentials_distinguished_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The distinguished name of the LDAP user account (example: cn=user,dc=meraki,dc=com).").String,
							Optional:            true,
						},
						"ldap_credentials_password": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The password of the LDAP user account.").String,
							Optional:            true,
						},
						"ldap_server_ca_certificate_contents": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The contents of the CA certificate. Must be in PEM or DER format.").String,
							Optional:            true,
						},
						"ldap_servers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The LDAP servers to be used for authentication.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP address (or FQDN) of your LDAP server.").String,
										Required:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("UDP port the LDAP server listens on.").String,
										Required:            true,
									},
								},
							},
						},
						"local_radius_cache_timeout": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The duration (in seconds) for which LDAP and OCSP lookups are cached.").String,
							Optional:            true,
						},
						"local_radius_certificate_authentication_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not to use EAP-TLS certificate-based authentication to validate wireless clients.").String,
							Optional:            true,
						},
						"local_radius_certificate_authentication_ocsp_responder_url": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("(Optional) The URL of the OCSP responder to verify client certificate status.").String,
							Optional:            true,
						},
						"local_radius_certificate_authentication_use_ldap": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not to verify the certificate with LDAP.").String,
							Optional:            true,
						},
						"local_radius_certificate_authentication_use_ocsp": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not to verify the certificate with OCSP.").String,
							Optional:            true,
						},
						"local_radius_certificate_authentication_client_root_ca_certificate_contents": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The contents of the Client CA Certificate. Must be in PEM or DER format.").String,
							Optional:            true,
						},
						"local_radius_password_authentication_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not to use EAP-TTLS/PAP or PEAP-GTC password-based authentication via LDAP lookup.").String,
							Optional:            true,
						},
						"named_vlans_radius_guest_vlan_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not RADIUS guest named VLAN is enabled.").String,
							Optional:            true,
						},
						"named_vlans_radius_guest_vlan_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("RADIUS guest VLAN name.").String,
							Optional:            true,
						},
						"named_vlans_tagging_default_vlan_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The default VLAN name used to tag traffic in the absence of a matching AP tag.").String,
							Optional:            true,
						},
						"named_vlans_tagging_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not traffic should be directed to use specific VLAN names.").String,
							Optional:            true,
						},
						"named_vlans_tagging_by_ap_tags": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The list of AP tags and VLAN names used for named VLAN tagging. If an AP has a tag matching one in the list, then traffic on this SSID will be directed to use the VLAN name associated to the tag.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"vlan_name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("VLAN name that will be used to tag traffic.").String,
										Optional:            true,
									},
									"tags": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("List of AP tags.").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
								},
							},
						},
						"oauth_allowed_domains": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("(Optional) The list of domains allowed access to the network.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"radius_radsec_tls_tunnel_timeout": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The interval (in seconds) to determines how long a TLS session can remain idle for a RADSec server before it is automatically terminated").String,
							Optional:            true,
						},
						"speed_burst_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether or not to allow users to temporarily exceed the bandwidth limit for short periods while still keeping them under the bandwidth limit over time.").String,
							Optional:            true,
						},
						"ap_tags_and_vlan_ids": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The list of tags and VLAN IDs used for VLAN tagging. This param is only valid when the ipAssignmentMode is `Bridge mode` or `Layer 3 roaming`").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"vlan_id": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Numerical identifier that is assigned to the VLAN").String,
										Optional:            true,
									},
									"tags": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Array of AP tags").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
								},
							},
						},
						"availability_tags": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Accepts a list of tags for this SSID. If availableOnAllAps is false, then the SSID will only be broadcast by APs with tags matching any of the tags in this list.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"radius_accounting_servers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The RADIUS accounting 802.1X servers to be used for authentication. This param is only valid if the authMode is `open-with-radius`, `8021x-radius` or `ipsk-with-radius` and radiusAccountingEnabled is `true`").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Certificate used for authorization for the RADSEC Server").String,
										Optional:            true,
									},
									"host": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP address (or FQDN) to which the APs will send RADIUS accounting messages").String,
										Required:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Port on the RADIUS server that is listening for accounting messages").String,
										Optional:            true,
									},
									"radsec_enabled": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Use RADSEC (TLS over TCP) to connect to this RADIUS accounting server. Requires radiusProxyEnabled.").String,
										Optional:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Shared key used to authenticate messages between the APs and RADIUS server").String,
										Optional:            true,
									},
								},
							},
						},
						"radius_servers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The RADIUS 802.1X servers to be used for authentication. This param is only valid if the authMode is `open-with-radius`, `8021x-radius` or `ipsk-with-radius`").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Certificate used for authorization for the RADSEC Server").String,
										Optional:            true,
									},
									"host": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("IP address (or FQDN) of your RADIUS server").String,
										Required:            true,
									},
									"open_roaming_certificate_id": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("The ID of the Openroaming Certificate attached to radius server.").String,
										Optional:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("UDP port the RADIUS server listens on for Access-requests").String,
										Optional:            true,
									},
									"radsec_enabled": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Use RADSEC (TLS over TCP) to connect to this RADIUS server. Requires radiusProxyEnabled.").String,
										Optional:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("RADIUS client shared secret").String,
										Optional:            true,
									},
								},
							},
						},
						"splash_guest_sponsor_domains": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Array of valid sponsor email domains for sponsored guest splash type.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"walled_garden_ranges": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify your walled garden by entering an array of addresses, ranges using CIDR notation, domain names, and domain wildcards (e.g. `192.168.1.1/24`, `192.168.37.10/32`, `www.yahoo.com`, `*.google.com`]). Meraki`s splash page is automatically included in your walled garden.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"radius_das_clients_ips": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of DAS (Dynamic Authorization Server) IPs. This is an unsupported attribute and is subject to breaking changes without prior notice.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"radius_das_clients_shared_secret": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Shared secret for DAS (Dynamic Authorization Server). This is an unsupported attribute and is subject to breaking changes without prior notice.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *WirelessSSIDsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *WirelessSSIDsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ResourceWirelessSSIDs

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
			Body:      item.toBody(ctx, ResourceWirelessSSIDsItems{}),
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

func (r *WirelessSSIDsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ResourceWirelessSSIDs

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

func (r *WirelessSSIDsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ResourceWirelessSSIDs

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
	for i := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			if plan.Items[i].Number.ValueString() != itemState.Number.ValueString() {
				continue
			}

			found = true
			// If the item is present in both plan and state, we need to check if it has changes
			hasChanges := plan.hasChanges(ctx, &state, plan.Items[i].Number.ValueString())
			if hasChanges {
				actions = append(actions, meraki.ActionModel{
					Operation: "update",
					Resource:  plan.getItemPath(plan.Items[i].Number.ValueString()),
					Body:      plan.Items[i].toBody(ctx, itemState),
				})
			}
			break
		}
		if !found {
			// If the item is present in plan, but not in state, we need to create it
			actions = append(actions, meraki.ActionModel{
				Operation: "update",
				Resource:  plan.getItemPath(plan.Items[i].Number.ValueString()),
				Body:      plan.Items[i].toBody(ctx, ResourceWirelessSSIDsItems{}),
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

func (r *WirelessSSIDsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ResourceWirelessSSIDs

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
func (r *WirelessSSIDsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin modifyPlan
func (r *WirelessSSIDsResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan, state ResourceWirelessSSIDs

	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning ModifyPlan", plan.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: ModifyPlan finished successfully", plan.Id.ValueString()))

	diags = resp.Plan.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end modifyPlan
