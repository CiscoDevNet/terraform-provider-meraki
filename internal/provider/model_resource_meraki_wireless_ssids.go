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
	"slices"
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceWirelessSSIDs struct {
	Id             types.String                 `tfsdk:"id"`
	OrganizationId types.String                 `tfsdk:"organization_id"`
	NetworkId      types.String                 `tfsdk:"network_id"`
	Items          []ResourceWirelessSSIDsItems `tfsdk:"items"`
}

type ResourceWirelessSSIDsItems struct {
	Number                                                              types.String                                     `tfsdk:"number"`
	AdaptivePolicyGroupId                                               types.String                                     `tfsdk:"adaptive_policy_group_id"`
	AdultContentFilteringEnabled                                        types.Bool                                       `tfsdk:"adult_content_filtering_enabled"`
	AuthMode                                                            types.String                                     `tfsdk:"auth_mode"`
	AvailableOnAllAps                                                   types.Bool                                       `tfsdk:"available_on_all_aps"`
	BandSelection                                                       types.String                                     `tfsdk:"band_selection"`
	ConcentratorNetworkId                                               types.String                                     `tfsdk:"concentrator_network_id"`
	DefaultVlanId                                                       types.Int64                                      `tfsdk:"default_vlan_id"`
	DisassociateClientsOnVpnFailover                                    types.Bool                                       `tfsdk:"disassociate_clients_on_vpn_failover"`
	Enabled                                                             types.Bool                                       `tfsdk:"enabled"`
	EncryptionMode                                                      types.String                                     `tfsdk:"encryption_mode"`
	EnterpriseAdminAccess                                               types.String                                     `tfsdk:"enterprise_admin_access"`
	IpAssignmentMode                                                    types.String                                     `tfsdk:"ip_assignment_mode"`
	LanIsolationEnabled                                                 types.Bool                                       `tfsdk:"lan_isolation_enabled"`
	MandatoryDhcpEnabled                                                types.Bool                                       `tfsdk:"mandatory_dhcp_enabled"`
	MinBitrate                                                          types.Float64                                    `tfsdk:"min_bitrate"`
	Name                                                                types.String                                     `tfsdk:"name"`
	PerClientBandwidthLimitDown                                         types.Int64                                      `tfsdk:"per_client_bandwidth_limit_down"`
	PerClientBandwidthLimitUp                                           types.Int64                                      `tfsdk:"per_client_bandwidth_limit_up"`
	PerSsidBandwidthLimitDown                                           types.Int64                                      `tfsdk:"per_ssid_bandwidth_limit_down"`
	PerSsidBandwidthLimitUp                                             types.Int64                                      `tfsdk:"per_ssid_bandwidth_limit_up"`
	Psk                                                                 types.String                                     `tfsdk:"psk"`
	RadiusAccountingEnabled                                             types.Bool                                       `tfsdk:"radius_accounting_enabled"`
	RadiusAccountingInterimInterval                                     types.Int64                                      `tfsdk:"radius_accounting_interim_interval"`
	RadiusAttributeForGroupPolicies                                     types.String                                     `tfsdk:"radius_attribute_for_group_policies"`
	RadiusAuthenticationNasId                                           types.String                                     `tfsdk:"radius_authentication_nas_id"`
	RadiusCalledStationId                                               types.String                                     `tfsdk:"radius_called_station_id"`
	RadiusCoaEnabled                                                    types.Bool                                       `tfsdk:"radius_coa_enabled"`
	RadiusFailoverPolicy                                                types.String                                     `tfsdk:"radius_failover_policy"`
	RadiusFallbackEnabled                                               types.Bool                                       `tfsdk:"radius_fallback_enabled"`
	RadiusGuestVlanEnabled                                              types.Bool                                       `tfsdk:"radius_guest_vlan_enabled"`
	RadiusGuestVlanId                                                   types.Int64                                      `tfsdk:"radius_guest_vlan_id"`
	RadiusLoadBalancingPolicy                                           types.String                                     `tfsdk:"radius_load_balancing_policy"`
	RadiusOverride                                                      types.Bool                                       `tfsdk:"radius_override"`
	RadiusProxyEnabled                                                  types.Bool                                       `tfsdk:"radius_proxy_enabled"`
	RadiusServerAttemptsLimit                                           types.Int64                                      `tfsdk:"radius_server_attempts_limit"`
	RadiusServerTimeout                                                 types.Int64                                      `tfsdk:"radius_server_timeout"`
	RadiusTestingEnabled                                                types.Bool                                       `tfsdk:"radius_testing_enabled"`
	SecondaryConcentratorNetworkId                                      types.String                                     `tfsdk:"secondary_concentrator_network_id"`
	SplashPage                                                          types.String                                     `tfsdk:"splash_page"`
	UseVlanTagging                                                      types.Bool                                       `tfsdk:"use_vlan_tagging"`
	Visible                                                             types.Bool                                       `tfsdk:"visible"`
	VlanId                                                              types.Int64                                      `tfsdk:"vlan_id"`
	WalledGardenEnabled                                                 types.Bool                                       `tfsdk:"walled_garden_enabled"`
	WpaEncryptionMode                                                   types.String                                     `tfsdk:"wpa_encryption_mode"`
	ActiveDirectoryCredentialsLogonName                                 types.String                                     `tfsdk:"active_directory_credentials_logon_name"`
	ActiveDirectoryCredentialsPassword                                  types.String                                     `tfsdk:"active_directory_credentials_password"`
	ActiveDirectoryServers                                              []ResourceWirelessSSIDsActiveDirectoryServers    `tfsdk:"active_directory_servers"`
	DnsRewriteEnabled                                                   types.Bool                                       `tfsdk:"dns_rewrite_enabled"`
	DnsRewriteDnsCustomNameservers                                      types.List                                       `tfsdk:"dns_rewrite_dns_custom_nameservers"`
	Dot11rAdaptive                                                      types.Bool                                       `tfsdk:"dot11r_adaptive"`
	Dot11rEnabled                                                       types.Bool                                       `tfsdk:"dot11r_enabled"`
	Dot11wEnabled                                                       types.Bool                                       `tfsdk:"dot11w_enabled"`
	Dot11wRequired                                                      types.Bool                                       `tfsdk:"dot11w_required"`
	GreKey                                                              types.Int64                                      `tfsdk:"gre_key"`
	GreConcentratorHost                                                 types.String                                     `tfsdk:"gre_concentrator_host"`
	LdapBaseDistinguishedName                                           types.String                                     `tfsdk:"ldap_base_distinguished_name"`
	LdapCredentialsDistinguishedName                                    types.String                                     `tfsdk:"ldap_credentials_distinguished_name"`
	LdapCredentialsPassword                                             types.String                                     `tfsdk:"ldap_credentials_password"`
	LdapServerCaCertificateContents                                     types.String                                     `tfsdk:"ldap_server_ca_certificate_contents"`
	LdapServers                                                         []ResourceWirelessSSIDsLdapServers               `tfsdk:"ldap_servers"`
	LocalRadiusCacheTimeout                                             types.Int64                                      `tfsdk:"local_radius_cache_timeout"`
	LocalRadiusCertificateAuthenticationEnabled                         types.Bool                                       `tfsdk:"local_radius_certificate_authentication_enabled"`
	LocalRadiusCertificateAuthenticationOcspResponderUrl                types.String                                     `tfsdk:"local_radius_certificate_authentication_ocsp_responder_url"`
	LocalRadiusCertificateAuthenticationUseLdap                         types.Bool                                       `tfsdk:"local_radius_certificate_authentication_use_ldap"`
	LocalRadiusCertificateAuthenticationUseOcsp                         types.Bool                                       `tfsdk:"local_radius_certificate_authentication_use_ocsp"`
	LocalRadiusCertificateAuthenticationClientRootCaCertificateContents types.String                                     `tfsdk:"local_radius_certificate_authentication_client_root_ca_certificate_contents"`
	LocalRadiusPasswordAuthenticationEnabled                            types.Bool                                       `tfsdk:"local_radius_password_authentication_enabled"`
	NamedVlansRadiusGuestVlanEnabled                                    types.Bool                                       `tfsdk:"named_vlans_radius_guest_vlan_enabled"`
	NamedVlansRadiusGuestVlanName                                       types.String                                     `tfsdk:"named_vlans_radius_guest_vlan_name"`
	NamedVlansTaggingDefaultVlanName                                    types.String                                     `tfsdk:"named_vlans_tagging_default_vlan_name"`
	NamedVlansTaggingEnabled                                            types.Bool                                       `tfsdk:"named_vlans_tagging_enabled"`
	NamedVlansTaggingByApTags                                           []ResourceWirelessSSIDsNamedVlansTaggingByApTags `tfsdk:"named_vlans_tagging_by_ap_tags"`
	OauthAllowedDomains                                                 types.Set                                        `tfsdk:"oauth_allowed_domains"`
	RadiusRadsecTlsTunnelTimeout                                        types.Int64                                      `tfsdk:"radius_radsec_tls_tunnel_timeout"`
	SpeedBurstEnabled                                                   types.Bool                                       `tfsdk:"speed_burst_enabled"`
	ApTagsAndVlanIds                                                    []ResourceWirelessSSIDsApTagsAndVlanIds          `tfsdk:"ap_tags_and_vlan_ids"`
	AvailabilityTags                                                    types.Set                                        `tfsdk:"availability_tags"`
	RadiusAccountingServers                                             []ResourceWirelessSSIDsRadiusAccountingServers   `tfsdk:"radius_accounting_servers"`
	RadiusServers                                                       []ResourceWirelessSSIDsRadiusServers             `tfsdk:"radius_servers"`
	SplashGuestSponsorDomains                                           types.Set                                        `tfsdk:"splash_guest_sponsor_domains"`
	WalledGardenRanges                                                  types.Set                                        `tfsdk:"walled_garden_ranges"`
	RadiusDasClientsIps                                                 types.Set                                        `tfsdk:"radius_das_clients_ips"`
	RadiusDasClientsSharedSecret                                        types.String                                     `tfsdk:"radius_das_clients_shared_secret"`
}

type ResourceWirelessSSIDsActiveDirectoryServers struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type ResourceWirelessSSIDsLdapServers struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type ResourceWirelessSSIDsNamedVlansTaggingByApTags struct {
	VlanName types.String `tfsdk:"vlan_name"`
	Tags     types.Set    `tfsdk:"tags"`
}

type ResourceWirelessSSIDsApTagsAndVlanIds struct {
	VlanId types.Int64 `tfsdk:"vlan_id"`
	Tags   types.Set   `tfsdk:"tags"`
}

type ResourceWirelessSSIDsRadiusAccountingServers struct {
	CaCertificate types.String `tfsdk:"ca_certificate"`
	Host          types.String `tfsdk:"host"`
	Port          types.Int64  `tfsdk:"port"`
	RadsecEnabled types.Bool   `tfsdk:"radsec_enabled"`
	Secret        types.String `tfsdk:"secret"`
}

type ResourceWirelessSSIDsRadiusServers struct {
	CaCertificate            types.String `tfsdk:"ca_certificate"`
	Host                     types.String `tfsdk:"host"`
	OpenRoamingCertificateId types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
	RadsecEnabled            types.Bool   `tfsdk:"radsec_enabled"`
	Secret                   types.String `tfsdk:"secret"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceWirelessSSIDs) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids", url.QueryEscape(data.NetworkId.ValueString()))
}

func (data ResourceWirelessSSIDs) getItemPath(id string) string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v", url.QueryEscape(data.NetworkId.ValueString()), id)
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceWirelessSSIDsItems) toBody(ctx context.Context, state ResourceWirelessSSIDsItems) string {
	body := ""
	if !data.AdaptivePolicyGroupId.IsNull() {
		body, _ = sjson.Set(body, "adaptivePolicyGroupId", data.AdaptivePolicyGroupId.ValueString())
	}
	if !data.AdultContentFilteringEnabled.IsNull() {
		body, _ = sjson.Set(body, "adultContentFilteringEnabled", data.AdultContentFilteringEnabled.ValueBool())
	}
	if !data.AuthMode.IsNull() {
		body, _ = sjson.Set(body, "authMode", data.AuthMode.ValueString())
	}
	if !data.AvailableOnAllAps.IsNull() {
		body, _ = sjson.Set(body, "availableOnAllAps", data.AvailableOnAllAps.ValueBool())
	}
	if !data.BandSelection.IsNull() {
		body, _ = sjson.Set(body, "bandSelection", data.BandSelection.ValueString())
	}
	if !data.ConcentratorNetworkId.IsNull() {
		body, _ = sjson.Set(body, "concentratorNetworkId", data.ConcentratorNetworkId.ValueString())
	}
	if !data.DefaultVlanId.IsNull() {
		body, _ = sjson.Set(body, "defaultVlanId", data.DefaultVlanId.ValueInt64())
	}
	if !data.DisassociateClientsOnVpnFailover.IsNull() {
		body, _ = sjson.Set(body, "disassociateClientsOnVpnFailover", data.DisassociateClientsOnVpnFailover.ValueBool())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.EncryptionMode.IsNull() {
		body, _ = sjson.Set(body, "encryptionMode", data.EncryptionMode.ValueString())
	}
	if !data.EnterpriseAdminAccess.IsNull() {
		body, _ = sjson.Set(body, "enterpriseAdminAccess", data.EnterpriseAdminAccess.ValueString())
	}
	if !data.IpAssignmentMode.IsNull() {
		body, _ = sjson.Set(body, "ipAssignmentMode", data.IpAssignmentMode.ValueString())
	}
	if !data.LanIsolationEnabled.IsNull() {
		body, _ = sjson.Set(body, "lanIsolationEnabled", data.LanIsolationEnabled.ValueBool())
	}
	if !data.MandatoryDhcpEnabled.IsNull() {
		body, _ = sjson.Set(body, "mandatoryDhcpEnabled", data.MandatoryDhcpEnabled.ValueBool())
	}
	if !data.MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "minBitrate", data.MinBitrate.ValueFloat64())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.PerClientBandwidthLimitDown.IsNull() {
		body, _ = sjson.Set(body, "perClientBandwidthLimitDown", data.PerClientBandwidthLimitDown.ValueInt64())
	}
	if !data.PerClientBandwidthLimitUp.IsNull() {
		body, _ = sjson.Set(body, "perClientBandwidthLimitUp", data.PerClientBandwidthLimitUp.ValueInt64())
	}
	if !data.PerSsidBandwidthLimitDown.IsNull() {
		body, _ = sjson.Set(body, "perSsidBandwidthLimitDown", data.PerSsidBandwidthLimitDown.ValueInt64())
	}
	if !data.PerSsidBandwidthLimitUp.IsNull() {
		body, _ = sjson.Set(body, "perSsidBandwidthLimitUp", data.PerSsidBandwidthLimitUp.ValueInt64())
	}
	if !data.Psk.IsNull() {
		body, _ = sjson.Set(body, "psk", data.Psk.ValueString())
	}
	if !data.RadiusAccountingEnabled.IsNull() {
		body, _ = sjson.Set(body, "radiusAccountingEnabled", data.RadiusAccountingEnabled.ValueBool())
	}
	if !data.RadiusAccountingInterimInterval.IsNull() {
		body, _ = sjson.Set(body, "radiusAccountingInterimInterval", data.RadiusAccountingInterimInterval.ValueInt64())
	}
	if !data.RadiusAttributeForGroupPolicies.IsNull() {
		body, _ = sjson.Set(body, "radiusAttributeForGroupPolicies", data.RadiusAttributeForGroupPolicies.ValueString())
	}
	if !data.RadiusAuthenticationNasId.IsNull() {
		body, _ = sjson.Set(body, "radiusAuthenticationNasId", data.RadiusAuthenticationNasId.ValueString())
	}
	if !data.RadiusCalledStationId.IsNull() {
		body, _ = sjson.Set(body, "radiusCalledStationId", data.RadiusCalledStationId.ValueString())
	}
	if !data.RadiusCoaEnabled.IsNull() {
		body, _ = sjson.Set(body, "radiusCoaEnabled", data.RadiusCoaEnabled.ValueBool())
	}
	if !data.RadiusFailoverPolicy.IsNull() {
		body, _ = sjson.Set(body, "radiusFailoverPolicy", data.RadiusFailoverPolicy.ValueString())
	}
	if !data.RadiusFallbackEnabled.IsNull() {
		body, _ = sjson.Set(body, "radiusFallbackEnabled", data.RadiusFallbackEnabled.ValueBool())
	}
	if !data.RadiusGuestVlanEnabled.IsNull() {
		body, _ = sjson.Set(body, "radiusGuestVlanEnabled", data.RadiusGuestVlanEnabled.ValueBool())
	}
	if !data.RadiusGuestVlanId.IsNull() {
		body, _ = sjson.Set(body, "radiusGuestVlanId", data.RadiusGuestVlanId.ValueInt64())
	}
	if !data.RadiusLoadBalancingPolicy.IsNull() {
		body, _ = sjson.Set(body, "radiusLoadBalancingPolicy", data.RadiusLoadBalancingPolicy.ValueString())
	}
	if !data.RadiusOverride.IsNull() {
		body, _ = sjson.Set(body, "radiusOverride", data.RadiusOverride.ValueBool())
	}
	if !data.RadiusProxyEnabled.IsNull() {
		body, _ = sjson.Set(body, "radiusProxyEnabled", data.RadiusProxyEnabled.ValueBool())
	}
	if !data.RadiusServerAttemptsLimit.IsNull() {
		body, _ = sjson.Set(body, "radiusServerAttemptsLimit", data.RadiusServerAttemptsLimit.ValueInt64())
	}
	if !data.RadiusServerTimeout.IsNull() {
		body, _ = sjson.Set(body, "radiusServerTimeout", data.RadiusServerTimeout.ValueInt64())
	}
	if !data.RadiusTestingEnabled.IsNull() {
		body, _ = sjson.Set(body, "radiusTestingEnabled", data.RadiusTestingEnabled.ValueBool())
	}
	if !data.SecondaryConcentratorNetworkId.IsNull() {
		body, _ = sjson.Set(body, "secondaryConcentratorNetworkId", data.SecondaryConcentratorNetworkId.ValueString())
	}
	if !data.SplashPage.IsNull() {
		body, _ = sjson.Set(body, "splashPage", data.SplashPage.ValueString())
	}
	if !data.UseVlanTagging.IsNull() {
		body, _ = sjson.Set(body, "useVlanTagging", data.UseVlanTagging.ValueBool())
	}
	if !data.Visible.IsNull() {
		body, _ = sjson.Set(body, "visible", data.Visible.ValueBool())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanId", data.VlanId.ValueInt64())
	}
	if !data.WalledGardenEnabled.IsNull() {
		body, _ = sjson.Set(body, "walledGardenEnabled", data.WalledGardenEnabled.ValueBool())
	}
	if !data.WpaEncryptionMode.IsNull() {
		body, _ = sjson.Set(body, "wpaEncryptionMode", data.WpaEncryptionMode.ValueString())
	}
	if !data.ActiveDirectoryCredentialsLogonName.IsNull() {
		body, _ = sjson.Set(body, "activeDirectory.credentials.logonName", data.ActiveDirectoryCredentialsLogonName.ValueString())
	}
	if !data.ActiveDirectoryCredentialsPassword.IsNull() {
		body, _ = sjson.Set(body, "activeDirectory.credentials.password", data.ActiveDirectoryCredentialsPassword.ValueString())
	}
	if len(data.ActiveDirectoryServers) > 0 {
		body, _ = sjson.Set(body, "activeDirectory.servers", []interface{}{})
		for _, item := range data.ActiveDirectoryServers {
			itemBody := ""
			if !item.Host.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "host", item.Host.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "activeDirectory.servers.-1", itemBody)
		}
	}
	if !data.DnsRewriteEnabled.IsNull() {
		body, _ = sjson.Set(body, "dnsRewrite.enabled", data.DnsRewriteEnabled.ValueBool())
	}
	if !data.DnsRewriteDnsCustomNameservers.IsNull() {
		var values []string
		data.DnsRewriteDnsCustomNameservers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "dnsRewrite.dnsCustomNameservers", values)
	}
	if !data.Dot11rAdaptive.IsNull() {
		body, _ = sjson.Set(body, "dot11r.adaptive", data.Dot11rAdaptive.ValueBool())
	}
	if !data.Dot11rEnabled.IsNull() {
		body, _ = sjson.Set(body, "dot11r.enabled", data.Dot11rEnabled.ValueBool())
	}
	if !data.Dot11wEnabled.IsNull() {
		body, _ = sjson.Set(body, "dot11w.enabled", data.Dot11wEnabled.ValueBool())
	}
	if !data.Dot11wRequired.IsNull() {
		body, _ = sjson.Set(body, "dot11w.required", data.Dot11wRequired.ValueBool())
	}
	if !data.GreKey.IsNull() {
		body, _ = sjson.Set(body, "gre.key", data.GreKey.ValueInt64())
	}
	if !data.GreConcentratorHost.IsNull() {
		body, _ = sjson.Set(body, "gre.concentrator.host", data.GreConcentratorHost.ValueString())
	}
	if !data.LdapBaseDistinguishedName.IsNull() {
		body, _ = sjson.Set(body, "ldap.baseDistinguishedName", data.LdapBaseDistinguishedName.ValueString())
	}
	if !data.LdapCredentialsDistinguishedName.IsNull() {
		body, _ = sjson.Set(body, "ldap.credentials.distinguishedName", data.LdapCredentialsDistinguishedName.ValueString())
	}
	if !data.LdapCredentialsPassword.IsNull() {
		body, _ = sjson.Set(body, "ldap.credentials.password", data.LdapCredentialsPassword.ValueString())
	}
	if !data.LdapServerCaCertificateContents.IsNull() {
		body, _ = sjson.Set(body, "ldap.serverCaCertificate.contents", data.LdapServerCaCertificateContents.ValueString())
	}
	if len(data.LdapServers) > 0 {
		body, _ = sjson.Set(body, "ldap.servers", []interface{}{})
		for _, item := range data.LdapServers {
			itemBody := ""
			if !item.Host.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "host", item.Host.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "ldap.servers.-1", itemBody)
		}
	}
	if !data.LocalRadiusCacheTimeout.IsNull() {
		body, _ = sjson.Set(body, "localRadius.cacheTimeout", data.LocalRadiusCacheTimeout.ValueInt64())
	}
	if !data.LocalRadiusCertificateAuthenticationEnabled.IsNull() {
		body, _ = sjson.Set(body, "localRadius.certificateAuthentication.enabled", data.LocalRadiusCertificateAuthenticationEnabled.ValueBool())
	}
	if !data.LocalRadiusCertificateAuthenticationOcspResponderUrl.IsNull() {
		body, _ = sjson.Set(body, "localRadius.certificateAuthentication.ocspResponderUrl", data.LocalRadiusCertificateAuthenticationOcspResponderUrl.ValueString())
	}
	if !data.LocalRadiusCertificateAuthenticationUseLdap.IsNull() {
		body, _ = sjson.Set(body, "localRadius.certificateAuthentication.useLdap", data.LocalRadiusCertificateAuthenticationUseLdap.ValueBool())
	}
	if !data.LocalRadiusCertificateAuthenticationUseOcsp.IsNull() {
		body, _ = sjson.Set(body, "localRadius.certificateAuthentication.useOcsp", data.LocalRadiusCertificateAuthenticationUseOcsp.ValueBool())
	}
	if !data.LocalRadiusCertificateAuthenticationClientRootCaCertificateContents.IsNull() {
		body, _ = sjson.Set(body, "localRadius.certificateAuthentication.clientRootCaCertificate.contents", data.LocalRadiusCertificateAuthenticationClientRootCaCertificateContents.ValueString())
	}
	if !data.LocalRadiusPasswordAuthenticationEnabled.IsNull() {
		body, _ = sjson.Set(body, "localRadius.passwordAuthentication.enabled", data.LocalRadiusPasswordAuthenticationEnabled.ValueBool())
	}
	if !data.NamedVlansRadiusGuestVlanEnabled.IsNull() {
		body, _ = sjson.Set(body, "namedVlans.radius.guestVlan.enabled", data.NamedVlansRadiusGuestVlanEnabled.ValueBool())
	}
	if !data.NamedVlansRadiusGuestVlanName.IsNull() {
		body, _ = sjson.Set(body, "namedVlans.radius.guestVlan.name", data.NamedVlansRadiusGuestVlanName.ValueString())
	}
	if !data.NamedVlansTaggingDefaultVlanName.IsNull() {
		body, _ = sjson.Set(body, "namedVlans.tagging.defaultVlanName", data.NamedVlansTaggingDefaultVlanName.ValueString())
	}
	if !data.NamedVlansTaggingEnabled.IsNull() {
		body, _ = sjson.Set(body, "namedVlans.tagging.enabled", data.NamedVlansTaggingEnabled.ValueBool())
	}
	if len(data.NamedVlansTaggingByApTags) > 0 {
		body, _ = sjson.Set(body, "namedVlans.tagging.byApTags", []interface{}{})
		for _, item := range data.NamedVlansTaggingByApTags {
			itemBody := ""
			if !item.VlanName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanName", item.VlanName.ValueString())
			}
			if !item.Tags.IsNull() {
				var values []string
				item.Tags.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "tags", values)
			}
			body, _ = sjson.SetRaw(body, "namedVlans.tagging.byApTags.-1", itemBody)
		}
	}
	if !data.OauthAllowedDomains.IsNull() {
		var values []string
		data.OauthAllowedDomains.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "oauth.allowedDomains", values)
	}
	if !data.RadiusRadsecTlsTunnelTimeout.IsNull() {
		body, _ = sjson.Set(body, "radiusRadsec.tlsTunnel.timeout", data.RadiusRadsecTlsTunnelTimeout.ValueInt64())
	}
	if !data.SpeedBurstEnabled.IsNull() {
		body, _ = sjson.Set(body, "speedBurst.enabled", data.SpeedBurstEnabled.ValueBool())
	}
	if len(data.ApTagsAndVlanIds) > 0 {
		body, _ = sjson.Set(body, "apTagsAndVlanIds", []interface{}{})
		for _, item := range data.ApTagsAndVlanIds {
			itemBody := ""
			if !item.VlanId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanId", item.VlanId.ValueInt64())
			}
			if !item.Tags.IsNull() {
				var values []string
				item.Tags.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "tags", values)
			}
			body, _ = sjson.SetRaw(body, "apTagsAndVlanIds.-1", itemBody)
		}
	}
	if !data.AvailabilityTags.IsNull() {
		var values []string
		data.AvailabilityTags.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "availabilityTags", values)
	}
	if len(data.RadiusAccountingServers) > 0 {
		body, _ = sjson.Set(body, "radiusAccountingServers", []interface{}{})
		for _, item := range data.RadiusAccountingServers {
			itemBody := ""
			if !item.CaCertificate.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "caCertificate", item.CaCertificate.ValueString())
			}
			if !item.Host.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "host", item.Host.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			if !item.RadsecEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "radsecEnabled", item.RadsecEnabled.ValueBool())
			}
			if !item.Secret.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secret", item.Secret.ValueString())
			}
			body, _ = sjson.SetRaw(body, "radiusAccountingServers.-1", itemBody)
		}
	}
	if len(data.RadiusServers) > 0 {
		body, _ = sjson.Set(body, "radiusServers", []interface{}{})
		for _, item := range data.RadiusServers {
			itemBody := ""
			if !item.CaCertificate.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "caCertificate", item.CaCertificate.ValueString())
			}
			if !item.Host.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "host", item.Host.ValueString())
			}
			if !item.OpenRoamingCertificateId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "openRoamingCertificateId", item.OpenRoamingCertificateId.ValueInt64())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			if !item.RadsecEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "radsecEnabled", item.RadsecEnabled.ValueBool())
			}
			if !item.Secret.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secret", item.Secret.ValueString())
			}
			body, _ = sjson.SetRaw(body, "radiusServers.-1", itemBody)
		}
	}
	if !data.SplashGuestSponsorDomains.IsNull() {
		var values []string
		data.SplashGuestSponsorDomains.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "splashGuestSponsorDomains", values)
	}
	if !data.WalledGardenRanges.IsNull() {
		var values []string
		data.WalledGardenRanges.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "walledGardenRanges", values)
	}
	if !data.RadiusDasClientsIps.IsNull() {
		var values []string
		data.RadiusDasClientsIps.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "radiusDasClients.clientsIps", values)
	}
	if !data.RadiusDasClientsSharedSecret.IsNull() {
		body, _ = sjson.Set(body, "radiusDasClients.clientsSharedSecret", data.RadiusDasClientsSharedSecret.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceWirelessSSIDs) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceWirelessSSIDsItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceWirelessSSIDsItems{}
		if value := res.Get("number"); value.Exists() && value.Value() != nil {
			data.Number = types.StringValue(value.String())
		} else {
			data.Number = types.StringNull()
		}
		if value := res.Get("adaptivePolicyGroupId"); value.Exists() && value.Value() != nil {
			data.AdaptivePolicyGroupId = types.StringValue(value.String())
		} else {
			data.AdaptivePolicyGroupId = types.StringNull()
		}
		if value := res.Get("adultContentFilteringEnabled"); value.Exists() && value.Value() != nil {
			data.AdultContentFilteringEnabled = types.BoolValue(value.Bool())
		} else {
			data.AdultContentFilteringEnabled = types.BoolNull()
		}
		if value := res.Get("authMode"); value.Exists() && value.Value() != nil {
			data.AuthMode = types.StringValue(value.String())
		} else {
			data.AuthMode = types.StringNull()
		}
		if value := res.Get("availableOnAllAps"); value.Exists() && value.Value() != nil {
			data.AvailableOnAllAps = types.BoolValue(value.Bool())
		} else {
			data.AvailableOnAllAps = types.BoolNull()
		}
		if value := res.Get("bandSelection"); value.Exists() && value.Value() != nil {
			data.BandSelection = types.StringValue(value.String())
		} else {
			data.BandSelection = types.StringNull()
		}
		if value := res.Get("concentratorNetworkId"); value.Exists() && value.Value() != nil {
			data.ConcentratorNetworkId = types.StringValue(value.String())
		} else {
			data.ConcentratorNetworkId = types.StringNull()
		}
		if value := res.Get("defaultVlanId"); value.Exists() && value.Value() != nil {
			data.DefaultVlanId = types.Int64Value(value.Int())
		} else {
			data.DefaultVlanId = types.Int64Null()
		}
		if value := res.Get("disassociateClientsOnVpnFailover"); value.Exists() && value.Value() != nil {
			data.DisassociateClientsOnVpnFailover = types.BoolValue(value.Bool())
		} else {
			data.DisassociateClientsOnVpnFailover = types.BoolNull()
		}
		if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("encryptionMode"); value.Exists() && value.Value() != nil {
			data.EncryptionMode = types.StringValue(value.String())
		} else {
			data.EncryptionMode = types.StringNull()
		}
		if value := res.Get("enterpriseAdminAccess"); value.Exists() && value.Value() != nil {
			data.EnterpriseAdminAccess = types.StringValue(value.String())
		} else {
			data.EnterpriseAdminAccess = types.StringNull()
		}
		if value := res.Get("ipAssignmentMode"); value.Exists() && value.Value() != nil {
			data.IpAssignmentMode = types.StringValue(value.String())
		} else {
			data.IpAssignmentMode = types.StringNull()
		}
		if value := res.Get("lanIsolationEnabled"); value.Exists() && value.Value() != nil {
			data.LanIsolationEnabled = types.BoolValue(value.Bool())
		} else {
			data.LanIsolationEnabled = types.BoolNull()
		}
		if value := res.Get("mandatoryDhcpEnabled"); value.Exists() && value.Value() != nil {
			data.MandatoryDhcpEnabled = types.BoolValue(value.Bool())
		} else {
			data.MandatoryDhcpEnabled = types.BoolNull()
		}
		if value := res.Get("minBitrate"); value.Exists() && value.Value() != nil {
			data.MinBitrate = types.Float64Value(value.Float())
		} else {
			data.MinBitrate = types.Float64Null()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("perClientBandwidthLimitDown"); value.Exists() && value.Value() != nil {
			data.PerClientBandwidthLimitDown = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitDown = types.Int64Null()
		}
		if value := res.Get("perClientBandwidthLimitUp"); value.Exists() && value.Value() != nil {
			data.PerClientBandwidthLimitUp = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitUp = types.Int64Null()
		}
		if value := res.Get("perSsidBandwidthLimitDown"); value.Exists() && value.Value() != nil {
			data.PerSsidBandwidthLimitDown = types.Int64Value(value.Int())
		} else {
			data.PerSsidBandwidthLimitDown = types.Int64Null()
		}
		if value := res.Get("perSsidBandwidthLimitUp"); value.Exists() && value.Value() != nil {
			data.PerSsidBandwidthLimitUp = types.Int64Value(value.Int())
		} else {
			data.PerSsidBandwidthLimitUp = types.Int64Null()
		}
		if value := res.Get("psk"); value.Exists() && value.Value() != nil {
			data.Psk = types.StringValue(value.String())
		} else {
			data.Psk = types.StringNull()
		}
		if value := res.Get("radiusAccountingEnabled"); value.Exists() && value.Value() != nil {
			data.RadiusAccountingEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusAccountingEnabled = types.BoolNull()
		}
		if value := res.Get("radiusAccountingInterimInterval"); value.Exists() && value.Value() != nil {
			data.RadiusAccountingInterimInterval = types.Int64Value(value.Int())
		} else {
			data.RadiusAccountingInterimInterval = types.Int64Null()
		}
		if value := res.Get("radiusAttributeForGroupPolicies"); value.Exists() && value.Value() != nil {
			data.RadiusAttributeForGroupPolicies = types.StringValue(value.String())
		} else {
			data.RadiusAttributeForGroupPolicies = types.StringNull()
		}
		if value := res.Get("radiusAuthenticationNasId"); value.Exists() && value.Value() != nil {
			data.RadiusAuthenticationNasId = types.StringValue(value.String())
		} else {
			data.RadiusAuthenticationNasId = types.StringNull()
		}
		if value := res.Get("radiusCalledStationId"); value.Exists() && value.Value() != nil {
			data.RadiusCalledStationId = types.StringValue(value.String())
		} else {
			data.RadiusCalledStationId = types.StringNull()
		}
		if value := res.Get("radiusCoaEnabled"); value.Exists() && value.Value() != nil {
			data.RadiusCoaEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusCoaEnabled = types.BoolNull()
		}
		if value := res.Get("radiusFailoverPolicy"); value.Exists() && value.Value() != nil {
			data.RadiusFailoverPolicy = types.StringValue(value.String())
		} else {
			data.RadiusFailoverPolicy = types.StringNull()
		}
		if value := res.Get("radiusFallbackEnabled"); value.Exists() && value.Value() != nil {
			data.RadiusFallbackEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusFallbackEnabled = types.BoolNull()
		}
		if value := res.Get("radiusGuestVlanEnabled"); value.Exists() && value.Value() != nil {
			data.RadiusGuestVlanEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusGuestVlanEnabled = types.BoolNull()
		}
		if value := res.Get("radiusGuestVlanId"); value.Exists() && value.Value() != nil {
			data.RadiusGuestVlanId = types.Int64Value(value.Int())
		} else {
			data.RadiusGuestVlanId = types.Int64Null()
		}
		if value := res.Get("radiusLoadBalancingPolicy"); value.Exists() && value.Value() != nil {
			data.RadiusLoadBalancingPolicy = types.StringValue(value.String())
		} else {
			data.RadiusLoadBalancingPolicy = types.StringNull()
		}
		if value := res.Get("radiusOverride"); value.Exists() && value.Value() != nil {
			data.RadiusOverride = types.BoolValue(value.Bool())
		} else {
			data.RadiusOverride = types.BoolNull()
		}
		if value := res.Get("radiusProxyEnabled"); value.Exists() && value.Value() != nil {
			data.RadiusProxyEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusProxyEnabled = types.BoolNull()
		}
		if value := res.Get("radiusServerAttemptsLimit"); value.Exists() && value.Value() != nil {
			data.RadiusServerAttemptsLimit = types.Int64Value(value.Int())
		} else {
			data.RadiusServerAttemptsLimit = types.Int64Null()
		}
		if value := res.Get("radiusServerTimeout"); value.Exists() && value.Value() != nil {
			data.RadiusServerTimeout = types.Int64Value(value.Int())
		} else {
			data.RadiusServerTimeout = types.Int64Null()
		}
		if value := res.Get("radiusTestingEnabled"); value.Exists() && value.Value() != nil {
			data.RadiusTestingEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusTestingEnabled = types.BoolNull()
		}
		if value := res.Get("secondaryConcentratorNetworkId"); value.Exists() && value.Value() != nil {
			data.SecondaryConcentratorNetworkId = types.StringValue(value.String())
		} else {
			data.SecondaryConcentratorNetworkId = types.StringNull()
		}
		if value := res.Get("splashPage"); value.Exists() && value.Value() != nil {
			data.SplashPage = types.StringValue(value.String())
		} else {
			data.SplashPage = types.StringNull()
		}
		if value := res.Get("useVlanTagging"); value.Exists() && value.Value() != nil {
			data.UseVlanTagging = types.BoolValue(value.Bool())
		} else {
			data.UseVlanTagging = types.BoolNull()
		}
		if value := res.Get("visible"); value.Exists() && value.Value() != nil {
			data.Visible = types.BoolValue(value.Bool())
		} else {
			data.Visible = types.BoolNull()
		}
		if value := res.Get("vlanId"); value.Exists() && value.Value() != nil {
			data.VlanId = types.Int64Value(value.Int())
		} else {
			data.VlanId = types.Int64Null()
		}
		if value := res.Get("walledGardenEnabled"); value.Exists() && value.Value() != nil {
			data.WalledGardenEnabled = types.BoolValue(value.Bool())
		} else {
			data.WalledGardenEnabled = types.BoolNull()
		}
		if value := res.Get("wpaEncryptionMode"); value.Exists() && value.Value() != nil {
			data.WpaEncryptionMode = types.StringValue(value.String())
		} else {
			data.WpaEncryptionMode = types.StringNull()
		}
		if value := res.Get("activeDirectory.credentials.logonName"); value.Exists() && value.Value() != nil {
			data.ActiveDirectoryCredentialsLogonName = types.StringValue(value.String())
		} else {
			data.ActiveDirectoryCredentialsLogonName = types.StringNull()
		}
		if value := res.Get("activeDirectory.credentials.password"); value.Exists() && value.Value() != nil {
			data.ActiveDirectoryCredentialsPassword = types.StringValue(value.String())
		} else {
			data.ActiveDirectoryCredentialsPassword = types.StringNull()
		}
		if value := res.Get("activeDirectory.servers"); value.Exists() && value.Value() != nil {
			data.ActiveDirectoryServers = make([]ResourceWirelessSSIDsActiveDirectoryServers, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceWirelessSSIDsActiveDirectoryServers{}
				if value := res.Get("host"); value.Exists() && value.Value() != nil {
					data.Host = types.StringValue(value.String())
				} else {
					data.Host = types.StringNull()
				}
				if value := res.Get("port"); value.Exists() && value.Value() != nil {
					data.Port = types.Int64Value(value.Int())
				} else {
					data.Port = types.Int64Null()
				}
				(*parent).ActiveDirectoryServers = append((*parent).ActiveDirectoryServers, data)
				return true
			})
		}
		if value := res.Get("dnsRewrite.enabled"); value.Exists() && value.Value() != nil {
			data.DnsRewriteEnabled = types.BoolValue(value.Bool())
		} else {
			data.DnsRewriteEnabled = types.BoolNull()
		}
		if value := res.Get("dnsRewrite.dnsCustomNameservers"); value.Exists() && value.Value() != nil {
			data.DnsRewriteDnsCustomNameservers = helpers.GetStringList(value.Array())
		} else {
			data.DnsRewriteDnsCustomNameservers = types.ListNull(types.StringType)
		}
		if value := res.Get("dot11r.adaptive"); value.Exists() && value.Value() != nil {
			data.Dot11rAdaptive = types.BoolValue(value.Bool())
		} else {
			data.Dot11rAdaptive = types.BoolNull()
		}
		if value := res.Get("dot11r.enabled"); value.Exists() && value.Value() != nil {
			data.Dot11rEnabled = types.BoolValue(value.Bool())
		} else {
			data.Dot11rEnabled = types.BoolNull()
		}
		if value := res.Get("dot11w.enabled"); value.Exists() && value.Value() != nil {
			data.Dot11wEnabled = types.BoolValue(value.Bool())
		} else {
			data.Dot11wEnabled = types.BoolNull()
		}
		if value := res.Get("dot11w.required"); value.Exists() && value.Value() != nil {
			data.Dot11wRequired = types.BoolValue(value.Bool())
		} else {
			data.Dot11wRequired = types.BoolNull()
		}
		if value := res.Get("gre.key"); value.Exists() && value.Value() != nil {
			data.GreKey = types.Int64Value(value.Int())
		} else {
			data.GreKey = types.Int64Null()
		}
		if value := res.Get("gre.concentrator.host"); value.Exists() && value.Value() != nil {
			data.GreConcentratorHost = types.StringValue(value.String())
		} else {
			data.GreConcentratorHost = types.StringNull()
		}
		if value := res.Get("ldap.baseDistinguishedName"); value.Exists() && value.Value() != nil {
			data.LdapBaseDistinguishedName = types.StringValue(value.String())
		} else {
			data.LdapBaseDistinguishedName = types.StringNull()
		}
		if value := res.Get("ldap.credentials.distinguishedName"); value.Exists() && value.Value() != nil {
			data.LdapCredentialsDistinguishedName = types.StringValue(value.String())
		} else {
			data.LdapCredentialsDistinguishedName = types.StringNull()
		}
		if value := res.Get("ldap.credentials.password"); value.Exists() && value.Value() != nil {
			data.LdapCredentialsPassword = types.StringValue(value.String())
		} else {
			data.LdapCredentialsPassword = types.StringNull()
		}
		if value := res.Get("ldap.serverCaCertificate.contents"); value.Exists() && value.Value() != nil {
			data.LdapServerCaCertificateContents = types.StringValue(value.String())
		} else {
			data.LdapServerCaCertificateContents = types.StringNull()
		}
		if value := res.Get("ldap.servers"); value.Exists() && value.Value() != nil {
			data.LdapServers = make([]ResourceWirelessSSIDsLdapServers, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceWirelessSSIDsLdapServers{}
				if value := res.Get("host"); value.Exists() && value.Value() != nil {
					data.Host = types.StringValue(value.String())
				} else {
					data.Host = types.StringNull()
				}
				if value := res.Get("port"); value.Exists() && value.Value() != nil {
					data.Port = types.Int64Value(value.Int())
				} else {
					data.Port = types.Int64Null()
				}
				(*parent).LdapServers = append((*parent).LdapServers, data)
				return true
			})
		}
		if value := res.Get("localRadius.cacheTimeout"); value.Exists() && value.Value() != nil {
			data.LocalRadiusCacheTimeout = types.Int64Value(value.Int())
		} else {
			data.LocalRadiusCacheTimeout = types.Int64Null()
		}
		if value := res.Get("localRadius.certificateAuthentication.enabled"); value.Exists() && value.Value() != nil {
			data.LocalRadiusCertificateAuthenticationEnabled = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusCertificateAuthenticationEnabled = types.BoolNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.ocspResponderUrl"); value.Exists() && value.Value() != nil {
			data.LocalRadiusCertificateAuthenticationOcspResponderUrl = types.StringValue(value.String())
		} else {
			data.LocalRadiusCertificateAuthenticationOcspResponderUrl = types.StringNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.useLdap"); value.Exists() && value.Value() != nil {
			data.LocalRadiusCertificateAuthenticationUseLdap = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusCertificateAuthenticationUseLdap = types.BoolNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.useOcsp"); value.Exists() && value.Value() != nil {
			data.LocalRadiusCertificateAuthenticationUseOcsp = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusCertificateAuthenticationUseOcsp = types.BoolNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.clientRootCaCertificate.contents"); value.Exists() && value.Value() != nil {
			data.LocalRadiusCertificateAuthenticationClientRootCaCertificateContents = types.StringValue(value.String())
		} else {
			data.LocalRadiusCertificateAuthenticationClientRootCaCertificateContents = types.StringNull()
		}
		if value := res.Get("localRadius.passwordAuthentication.enabled"); value.Exists() && value.Value() != nil {
			data.LocalRadiusPasswordAuthenticationEnabled = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusPasswordAuthenticationEnabled = types.BoolNull()
		}
		if value := res.Get("namedVlans.radius.guestVlan.enabled"); value.Exists() && value.Value() != nil {
			data.NamedVlansRadiusGuestVlanEnabled = types.BoolValue(value.Bool())
		} else {
			data.NamedVlansRadiusGuestVlanEnabled = types.BoolNull()
		}
		if value := res.Get("namedVlans.radius.guestVlan.name"); value.Exists() && value.Value() != nil {
			data.NamedVlansRadiusGuestVlanName = types.StringValue(value.String())
		} else {
			data.NamedVlansRadiusGuestVlanName = types.StringNull()
		}
		if value := res.Get("namedVlans.tagging.defaultVlanName"); value.Exists() && value.Value() != nil {
			data.NamedVlansTaggingDefaultVlanName = types.StringValue(value.String())
		} else {
			data.NamedVlansTaggingDefaultVlanName = types.StringNull()
		}
		if value := res.Get("namedVlans.tagging.enabled"); value.Exists() && value.Value() != nil {
			data.NamedVlansTaggingEnabled = types.BoolValue(value.Bool())
		} else {
			data.NamedVlansTaggingEnabled = types.BoolNull()
		}
		if value := res.Get("namedVlans.tagging.byApTags"); value.Exists() && value.Value() != nil {
			data.NamedVlansTaggingByApTags = make([]ResourceWirelessSSIDsNamedVlansTaggingByApTags, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceWirelessSSIDsNamedVlansTaggingByApTags{}
				if value := res.Get("vlanName"); value.Exists() && value.Value() != nil {
					data.VlanName = types.StringValue(value.String())
				} else {
					data.VlanName = types.StringNull()
				}
				if value := res.Get("tags"); value.Exists() && value.Value() != nil {
					data.Tags = helpers.GetStringSet(value.Array())
				} else {
					data.Tags = types.SetNull(types.StringType)
				}
				(*parent).NamedVlansTaggingByApTags = append((*parent).NamedVlansTaggingByApTags, data)
				return true
			})
		}
		if value := res.Get("oauth.allowedDomains"); value.Exists() && value.Value() != nil {
			data.OauthAllowedDomains = helpers.GetStringSet(value.Array())
		} else {
			data.OauthAllowedDomains = types.SetNull(types.StringType)
		}
		if value := res.Get("radiusRadsec.tlsTunnel.timeout"); value.Exists() && value.Value() != nil {
			data.RadiusRadsecTlsTunnelTimeout = types.Int64Value(value.Int())
		} else {
			data.RadiusRadsecTlsTunnelTimeout = types.Int64Null()
		}
		if value := res.Get("speedBurst.enabled"); value.Exists() && value.Value() != nil {
			data.SpeedBurstEnabled = types.BoolValue(value.Bool())
		} else {
			data.SpeedBurstEnabled = types.BoolNull()
		}
		if value := res.Get("apTagsAndVlanIds"); value.Exists() && value.Value() != nil {
			data.ApTagsAndVlanIds = make([]ResourceWirelessSSIDsApTagsAndVlanIds, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceWirelessSSIDsApTagsAndVlanIds{}
				if value := res.Get("vlanId"); value.Exists() && value.Value() != nil {
					data.VlanId = types.Int64Value(value.Int())
				} else {
					data.VlanId = types.Int64Null()
				}
				if value := res.Get("tags"); value.Exists() && value.Value() != nil {
					data.Tags = helpers.GetStringSet(value.Array())
				} else {
					data.Tags = types.SetNull(types.StringType)
				}
				(*parent).ApTagsAndVlanIds = append((*parent).ApTagsAndVlanIds, data)
				return true
			})
		}
		if value := res.Get("availabilityTags"); value.Exists() && value.Value() != nil {
			data.AvailabilityTags = helpers.GetStringSet(value.Array())
		} else {
			data.AvailabilityTags = types.SetNull(types.StringType)
		}
		if value := res.Get("radiusAccountingServers"); value.Exists() && value.Value() != nil {
			data.RadiusAccountingServers = make([]ResourceWirelessSSIDsRadiusAccountingServers, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceWirelessSSIDsRadiusAccountingServers{}
				if value := res.Get("caCertificate"); value.Exists() && value.Value() != nil {
					data.CaCertificate = types.StringValue(value.String())
				} else {
					data.CaCertificate = types.StringNull()
				}
				if value := res.Get("host"); value.Exists() && value.Value() != nil {
					data.Host = types.StringValue(value.String())
				} else {
					data.Host = types.StringNull()
				}
				if value := res.Get("port"); value.Exists() && value.Value() != nil {
					data.Port = types.Int64Value(value.Int())
				} else {
					data.Port = types.Int64Null()
				}
				if value := res.Get("radsecEnabled"); value.Exists() && value.Value() != nil {
					data.RadsecEnabled = types.BoolValue(value.Bool())
				} else {
					data.RadsecEnabled = types.BoolNull()
				}
				(*parent).RadiusAccountingServers = append((*parent).RadiusAccountingServers, data)
				return true
			})
		}
		if value := res.Get("radiusServers"); value.Exists() && value.Value() != nil {
			data.RadiusServers = make([]ResourceWirelessSSIDsRadiusServers, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceWirelessSSIDsRadiusServers{}
				if value := res.Get("caCertificate"); value.Exists() && value.Value() != nil {
					data.CaCertificate = types.StringValue(value.String())
				} else {
					data.CaCertificate = types.StringNull()
				}
				if value := res.Get("host"); value.Exists() && value.Value() != nil {
					data.Host = types.StringValue(value.String())
				} else {
					data.Host = types.StringNull()
				}
				if value := res.Get("openRoamingCertificateId"); value.Exists() && value.Value() != nil {
					data.OpenRoamingCertificateId = types.Int64Value(value.Int())
				} else {
					data.OpenRoamingCertificateId = types.Int64Null()
				}
				if value := res.Get("port"); value.Exists() && value.Value() != nil {
					data.Port = types.Int64Value(value.Int())
				} else {
					data.Port = types.Int64Null()
				}
				if value := res.Get("radsecEnabled"); value.Exists() && value.Value() != nil {
					data.RadsecEnabled = types.BoolValue(value.Bool())
				} else {
					data.RadsecEnabled = types.BoolNull()
				}
				(*parent).RadiusServers = append((*parent).RadiusServers, data)
				return true
			})
		}
		if value := res.Get("splashGuestSponsorDomains"); value.Exists() && value.Value() != nil {
			data.SplashGuestSponsorDomains = helpers.GetStringSet(value.Array())
		} else {
			data.SplashGuestSponsorDomains = types.SetNull(types.StringType)
		}
		if value := res.Get("walledGardenRanges"); value.Exists() && value.Value() != nil {
			data.WalledGardenRanges = helpers.GetStringSet(value.Array())
		} else {
			data.WalledGardenRanges = types.SetNull(types.StringType)
		}
		if value := res.Get("radiusDasClients.clientsIps"); value.Exists() && value.Value() != nil {
			data.RadiusDasClientsIps = helpers.GetStringSet(value.Array())
		} else {
			data.RadiusDasClientsIps = types.SetNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	data.Id = data.OrganizationId
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceWirelessSSIDs) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("number").String() == (*parent).Items[i].Number.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
		if value := res.Get("adaptivePolicyGroupId"); value.Exists() && !data.AdaptivePolicyGroupId.IsNull() {
			data.AdaptivePolicyGroupId = types.StringValue(value.String())
		} else {
			data.AdaptivePolicyGroupId = types.StringNull()
		}
		if value := res.Get("adultContentFilteringEnabled"); value.Exists() && !data.AdultContentFilteringEnabled.IsNull() {
			data.AdultContentFilteringEnabled = types.BoolValue(value.Bool())
		} else {
			data.AdultContentFilteringEnabled = types.BoolNull()
		}
		if value := res.Get("authMode"); value.Exists() && !data.AuthMode.IsNull() {
			data.AuthMode = types.StringValue(value.String())
		} else {
			data.AuthMode = types.StringNull()
		}
		if value := res.Get("availableOnAllAps"); value.Exists() && !data.AvailableOnAllAps.IsNull() {
			data.AvailableOnAllAps = types.BoolValue(value.Bool())
		} else {
			data.AvailableOnAllAps = types.BoolNull()
		}
		if value := res.Get("bandSelection"); value.Exists() && !data.BandSelection.IsNull() {
			data.BandSelection = types.StringValue(value.String())
		} else {
			data.BandSelection = types.StringNull()
		}
		if value := res.Get("concentratorNetworkId"); value.Exists() && !data.ConcentratorNetworkId.IsNull() {
			data.ConcentratorNetworkId = types.StringValue(value.String())
		} else {
			data.ConcentratorNetworkId = types.StringNull()
		}
		if value := res.Get("defaultVlanId"); value.Exists() && !data.DefaultVlanId.IsNull() {
			data.DefaultVlanId = types.Int64Value(value.Int())
		} else {
			data.DefaultVlanId = types.Int64Null()
		}
		if value := res.Get("disassociateClientsOnVpnFailover"); value.Exists() && !data.DisassociateClientsOnVpnFailover.IsNull() {
			data.DisassociateClientsOnVpnFailover = types.BoolValue(value.Bool())
		} else {
			data.DisassociateClientsOnVpnFailover = types.BoolNull()
		}
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("encryptionMode"); value.Exists() && !data.EncryptionMode.IsNull() {
			data.EncryptionMode = types.StringValue(value.String())
		} else {
			data.EncryptionMode = types.StringNull()
		}
		if value := res.Get("enterpriseAdminAccess"); value.Exists() && !data.EnterpriseAdminAccess.IsNull() {
			data.EnterpriseAdminAccess = types.StringValue(value.String())
		} else {
			data.EnterpriseAdminAccess = types.StringNull()
		}
		if value := res.Get("ipAssignmentMode"); value.Exists() && !data.IpAssignmentMode.IsNull() {
			data.IpAssignmentMode = types.StringValue(value.String())
		} else {
			data.IpAssignmentMode = types.StringNull()
		}
		if value := res.Get("lanIsolationEnabled"); value.Exists() && !data.LanIsolationEnabled.IsNull() {
			data.LanIsolationEnabled = types.BoolValue(value.Bool())
		} else {
			data.LanIsolationEnabled = types.BoolNull()
		}
		if value := res.Get("mandatoryDhcpEnabled"); value.Exists() && !data.MandatoryDhcpEnabled.IsNull() {
			data.MandatoryDhcpEnabled = types.BoolValue(value.Bool())
		} else {
			data.MandatoryDhcpEnabled = types.BoolNull()
		}
		if value := res.Get("minBitrate"); value.Exists() && !data.MinBitrate.IsNull() {
			data.MinBitrate = types.Float64Value(value.Float())
		} else {
			data.MinBitrate = types.Float64Null()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("perClientBandwidthLimitDown"); value.Exists() && !data.PerClientBandwidthLimitDown.IsNull() {
			data.PerClientBandwidthLimitDown = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitDown = types.Int64Null()
		}
		if value := res.Get("perClientBandwidthLimitUp"); value.Exists() && !data.PerClientBandwidthLimitUp.IsNull() {
			data.PerClientBandwidthLimitUp = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitUp = types.Int64Null()
		}
		if value := res.Get("perSsidBandwidthLimitDown"); value.Exists() && !data.PerSsidBandwidthLimitDown.IsNull() {
			data.PerSsidBandwidthLimitDown = types.Int64Value(value.Int())
		} else {
			data.PerSsidBandwidthLimitDown = types.Int64Null()
		}
		if value := res.Get("perSsidBandwidthLimitUp"); value.Exists() && !data.PerSsidBandwidthLimitUp.IsNull() {
			data.PerSsidBandwidthLimitUp = types.Int64Value(value.Int())
		} else {
			data.PerSsidBandwidthLimitUp = types.Int64Null()
		}
		if value := res.Get("psk"); value.Exists() && !data.Psk.IsNull() {
			data.Psk = types.StringValue(value.String())
		} else {
			data.Psk = types.StringNull()
		}
		if value := res.Get("radiusAccountingEnabled"); value.Exists() && !data.RadiusAccountingEnabled.IsNull() {
			data.RadiusAccountingEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusAccountingEnabled = types.BoolNull()
		}
		if value := res.Get("radiusAccountingInterimInterval"); value.Exists() && !data.RadiusAccountingInterimInterval.IsNull() {
			data.RadiusAccountingInterimInterval = types.Int64Value(value.Int())
		} else {
			data.RadiusAccountingInterimInterval = types.Int64Null()
		}
		if value := res.Get("radiusAttributeForGroupPolicies"); value.Exists() && !data.RadiusAttributeForGroupPolicies.IsNull() {
			data.RadiusAttributeForGroupPolicies = types.StringValue(value.String())
		} else {
			data.RadiusAttributeForGroupPolicies = types.StringNull()
		}
		if value := res.Get("radiusAuthenticationNasId"); value.Exists() && !data.RadiusAuthenticationNasId.IsNull() {
			data.RadiusAuthenticationNasId = types.StringValue(value.String())
		} else {
			data.RadiusAuthenticationNasId = types.StringNull()
		}
		if value := res.Get("radiusCalledStationId"); value.Exists() && !data.RadiusCalledStationId.IsNull() {
			data.RadiusCalledStationId = types.StringValue(value.String())
		} else {
			data.RadiusCalledStationId = types.StringNull()
		}
		if value := res.Get("radiusCoaEnabled"); value.Exists() && !data.RadiusCoaEnabled.IsNull() {
			data.RadiusCoaEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusCoaEnabled = types.BoolNull()
		}
		if value := res.Get("radiusFailoverPolicy"); value.Exists() && !data.RadiusFailoverPolicy.IsNull() {
			data.RadiusFailoverPolicy = types.StringValue(value.String())
		} else {
			data.RadiusFailoverPolicy = types.StringNull()
		}
		if value := res.Get("radiusFallbackEnabled"); value.Exists() && !data.RadiusFallbackEnabled.IsNull() {
			data.RadiusFallbackEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusFallbackEnabled = types.BoolNull()
		}
		if value := res.Get("radiusGuestVlanEnabled"); value.Exists() && !data.RadiusGuestVlanEnabled.IsNull() {
			data.RadiusGuestVlanEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusGuestVlanEnabled = types.BoolNull()
		}
		if value := res.Get("radiusGuestVlanId"); value.Exists() && !data.RadiusGuestVlanId.IsNull() {
			data.RadiusGuestVlanId = types.Int64Value(value.Int())
		} else {
			data.RadiusGuestVlanId = types.Int64Null()
		}
		if value := res.Get("radiusLoadBalancingPolicy"); value.Exists() && !data.RadiusLoadBalancingPolicy.IsNull() {
			data.RadiusLoadBalancingPolicy = types.StringValue(value.String())
		} else {
			data.RadiusLoadBalancingPolicy = types.StringNull()
		}
		if value := res.Get("radiusOverride"); value.Exists() && !data.RadiusOverride.IsNull() {
			data.RadiusOverride = types.BoolValue(value.Bool())
		} else {
			data.RadiusOverride = types.BoolNull()
		}
		if value := res.Get("radiusProxyEnabled"); value.Exists() && !data.RadiusProxyEnabled.IsNull() {
			data.RadiusProxyEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusProxyEnabled = types.BoolNull()
		}
		if value := res.Get("radiusServerAttemptsLimit"); value.Exists() && !data.RadiusServerAttemptsLimit.IsNull() {
			data.RadiusServerAttemptsLimit = types.Int64Value(value.Int())
		} else {
			data.RadiusServerAttemptsLimit = types.Int64Null()
		}
		if value := res.Get("radiusServerTimeout"); value.Exists() && !data.RadiusServerTimeout.IsNull() {
			data.RadiusServerTimeout = types.Int64Value(value.Int())
		} else {
			data.RadiusServerTimeout = types.Int64Null()
		}
		if value := res.Get("radiusTestingEnabled"); value.Exists() && !data.RadiusTestingEnabled.IsNull() {
			data.RadiusTestingEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusTestingEnabled = types.BoolNull()
		}
		if value := res.Get("secondaryConcentratorNetworkId"); value.Exists() && !data.SecondaryConcentratorNetworkId.IsNull() {
			data.SecondaryConcentratorNetworkId = types.StringValue(value.String())
		} else {
			data.SecondaryConcentratorNetworkId = types.StringNull()
		}
		if value := res.Get("splashPage"); value.Exists() && !data.SplashPage.IsNull() {
			data.SplashPage = types.StringValue(value.String())
		} else {
			data.SplashPage = types.StringNull()
		}
		if value := res.Get("useVlanTagging"); value.Exists() && !data.UseVlanTagging.IsNull() {
			data.UseVlanTagging = types.BoolValue(value.Bool())
		} else {
			data.UseVlanTagging = types.BoolNull()
		}
		if value := res.Get("visible"); value.Exists() && !data.Visible.IsNull() {
			data.Visible = types.BoolValue(value.Bool())
		} else {
			data.Visible = types.BoolNull()
		}
		if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
			data.VlanId = types.Int64Value(value.Int())
		} else {
			data.VlanId = types.Int64Null()
		}
		if value := res.Get("walledGardenEnabled"); value.Exists() && !data.WalledGardenEnabled.IsNull() {
			data.WalledGardenEnabled = types.BoolValue(value.Bool())
		} else {
			data.WalledGardenEnabled = types.BoolNull()
		}
		if value := res.Get("wpaEncryptionMode"); value.Exists() && !data.WpaEncryptionMode.IsNull() {
			data.WpaEncryptionMode = types.StringValue(value.String())
		} else {
			data.WpaEncryptionMode = types.StringNull()
		}
		if value := res.Get("activeDirectory.credentials.logonName"); value.Exists() && !data.ActiveDirectoryCredentialsLogonName.IsNull() {
			data.ActiveDirectoryCredentialsLogonName = types.StringValue(value.String())
		} else {
			data.ActiveDirectoryCredentialsLogonName = types.StringNull()
		}
		if value := res.Get("activeDirectory.credentials.password"); value.Exists() && !data.ActiveDirectoryCredentialsPassword.IsNull() {
			data.ActiveDirectoryCredentialsPassword = types.StringValue(value.String())
		} else {
			data.ActiveDirectoryCredentialsPassword = types.StringNull()
		}
		for i := 0; i < len(data.ActiveDirectoryServers); i++ {
			keys := [...]string{"host"}
			keyValues := [...]string{data.ActiveDirectoryServers[i].Host.ValueString()}

			parent := &data
			data := (*parent).ActiveDirectoryServers[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("activeDirectory.servers").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing ActiveDirectoryServers[%d] = %+v",
					i,
					(*parent).ActiveDirectoryServers[i],
				))
				(*parent).ActiveDirectoryServers = slices.Delete((*parent).ActiveDirectoryServers, i, i+1)
				i--

				continue
			}
			if value := res.Get("host"); value.Exists() && !data.Host.IsNull() {
				data.Host = types.StringValue(value.String())
			} else {
				data.Host = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			(*parent).ActiveDirectoryServers[i] = data
		}
		if value := res.Get("dnsRewrite.enabled"); value.Exists() && !data.DnsRewriteEnabled.IsNull() {
			data.DnsRewriteEnabled = types.BoolValue(value.Bool())
		} else {
			data.DnsRewriteEnabled = types.BoolNull()
		}
		if value := res.Get("dnsRewrite.dnsCustomNameservers"); value.Exists() && !data.DnsRewriteDnsCustomNameservers.IsNull() {
			data.DnsRewriteDnsCustomNameservers = helpers.GetStringList(value.Array())
		} else {
			data.DnsRewriteDnsCustomNameservers = types.ListNull(types.StringType)
		}
		if value := res.Get("dot11r.adaptive"); value.Exists() && !data.Dot11rAdaptive.IsNull() {
			data.Dot11rAdaptive = types.BoolValue(value.Bool())
		} else {
			data.Dot11rAdaptive = types.BoolNull()
		}
		if value := res.Get("dot11r.enabled"); value.Exists() && !data.Dot11rEnabled.IsNull() {
			data.Dot11rEnabled = types.BoolValue(value.Bool())
		} else {
			data.Dot11rEnabled = types.BoolNull()
		}
		if value := res.Get("dot11w.enabled"); value.Exists() && !data.Dot11wEnabled.IsNull() {
			data.Dot11wEnabled = types.BoolValue(value.Bool())
		} else {
			data.Dot11wEnabled = types.BoolNull()
		}
		if value := res.Get("dot11w.required"); value.Exists() && !data.Dot11wRequired.IsNull() {
			data.Dot11wRequired = types.BoolValue(value.Bool())
		} else {
			data.Dot11wRequired = types.BoolNull()
		}
		if value := res.Get("gre.key"); value.Exists() && !data.GreKey.IsNull() {
			data.GreKey = types.Int64Value(value.Int())
		} else {
			data.GreKey = types.Int64Null()
		}
		if value := res.Get("gre.concentrator.host"); value.Exists() && !data.GreConcentratorHost.IsNull() {
			data.GreConcentratorHost = types.StringValue(value.String())
		} else {
			data.GreConcentratorHost = types.StringNull()
		}
		if value := res.Get("ldap.baseDistinguishedName"); value.Exists() && !data.LdapBaseDistinguishedName.IsNull() {
			data.LdapBaseDistinguishedName = types.StringValue(value.String())
		} else {
			data.LdapBaseDistinguishedName = types.StringNull()
		}
		if value := res.Get("ldap.credentials.distinguishedName"); value.Exists() && !data.LdapCredentialsDistinguishedName.IsNull() {
			data.LdapCredentialsDistinguishedName = types.StringValue(value.String())
		} else {
			data.LdapCredentialsDistinguishedName = types.StringNull()
		}
		if value := res.Get("ldap.credentials.password"); value.Exists() && !data.LdapCredentialsPassword.IsNull() {
			data.LdapCredentialsPassword = types.StringValue(value.String())
		} else {
			data.LdapCredentialsPassword = types.StringNull()
		}
		if value := res.Get("ldap.serverCaCertificate.contents"); value.Exists() && !data.LdapServerCaCertificateContents.IsNull() {
			data.LdapServerCaCertificateContents = types.StringValue(value.String())
		} else {
			data.LdapServerCaCertificateContents = types.StringNull()
		}
		{
			l := len(res.Get("ldap.servers").Array())
			tflog.Debug(ctx, fmt.Sprintf("ldap.servers array resizing from %d to %d", len(data.LdapServers), l))
			if len(data.LdapServers) > l {
				data.LdapServers = data.LdapServers[:l]
			}
		}
		for i := range data.LdapServers {
			parent := &data
			data := (*parent).LdapServers[i]
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("ldap.servers.%d", i))
			if value := res.Get("host"); value.Exists() && !data.Host.IsNull() {
				data.Host = types.StringValue(value.String())
			} else {
				data.Host = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			(*parent).LdapServers[i] = data
		}
		if value := res.Get("localRadius.cacheTimeout"); value.Exists() && !data.LocalRadiusCacheTimeout.IsNull() {
			data.LocalRadiusCacheTimeout = types.Int64Value(value.Int())
		} else {
			data.LocalRadiusCacheTimeout = types.Int64Null()
		}
		if value := res.Get("localRadius.certificateAuthentication.enabled"); value.Exists() && !data.LocalRadiusCertificateAuthenticationEnabled.IsNull() {
			data.LocalRadiusCertificateAuthenticationEnabled = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusCertificateAuthenticationEnabled = types.BoolNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.ocspResponderUrl"); value.Exists() && !data.LocalRadiusCertificateAuthenticationOcspResponderUrl.IsNull() {
			data.LocalRadiusCertificateAuthenticationOcspResponderUrl = types.StringValue(value.String())
		} else {
			data.LocalRadiusCertificateAuthenticationOcspResponderUrl = types.StringNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.useLdap"); value.Exists() && !data.LocalRadiusCertificateAuthenticationUseLdap.IsNull() {
			data.LocalRadiusCertificateAuthenticationUseLdap = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusCertificateAuthenticationUseLdap = types.BoolNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.useOcsp"); value.Exists() && !data.LocalRadiusCertificateAuthenticationUseOcsp.IsNull() {
			data.LocalRadiusCertificateAuthenticationUseOcsp = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusCertificateAuthenticationUseOcsp = types.BoolNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.clientRootCaCertificate.contents"); value.Exists() && !data.LocalRadiusCertificateAuthenticationClientRootCaCertificateContents.IsNull() {
			data.LocalRadiusCertificateAuthenticationClientRootCaCertificateContents = types.StringValue(value.String())
		} else {
			data.LocalRadiusCertificateAuthenticationClientRootCaCertificateContents = types.StringNull()
		}
		if value := res.Get("localRadius.passwordAuthentication.enabled"); value.Exists() && !data.LocalRadiusPasswordAuthenticationEnabled.IsNull() {
			data.LocalRadiusPasswordAuthenticationEnabled = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusPasswordAuthenticationEnabled = types.BoolNull()
		}
		if value := res.Get("namedVlans.radius.guestVlan.enabled"); value.Exists() && !data.NamedVlansRadiusGuestVlanEnabled.IsNull() {
			data.NamedVlansRadiusGuestVlanEnabled = types.BoolValue(value.Bool())
		} else {
			data.NamedVlansRadiusGuestVlanEnabled = types.BoolNull()
		}
		if value := res.Get("namedVlans.radius.guestVlan.name"); value.Exists() && !data.NamedVlansRadiusGuestVlanName.IsNull() {
			data.NamedVlansRadiusGuestVlanName = types.StringValue(value.String())
		} else {
			data.NamedVlansRadiusGuestVlanName = types.StringNull()
		}
		if value := res.Get("namedVlans.tagging.defaultVlanName"); value.Exists() && !data.NamedVlansTaggingDefaultVlanName.IsNull() {
			data.NamedVlansTaggingDefaultVlanName = types.StringValue(value.String())
		} else {
			data.NamedVlansTaggingDefaultVlanName = types.StringNull()
		}
		if value := res.Get("namedVlans.tagging.enabled"); value.Exists() && !data.NamedVlansTaggingEnabled.IsNull() {
			data.NamedVlansTaggingEnabled = types.BoolValue(value.Bool())
		} else {
			data.NamedVlansTaggingEnabled = types.BoolNull()
		}
		for i := 0; i < len(data.NamedVlansTaggingByApTags); i++ {
			keys := [...]string{"vlanName"}
			keyValues := [...]string{data.NamedVlansTaggingByApTags[i].VlanName.ValueString()}

			parent := &data
			data := (*parent).NamedVlansTaggingByApTags[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("namedVlans.tagging.byApTags").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing NamedVlansTaggingByApTags[%d] = %+v",
					i,
					(*parent).NamedVlansTaggingByApTags[i],
				))
				(*parent).NamedVlansTaggingByApTags = slices.Delete((*parent).NamedVlansTaggingByApTags, i, i+1)
				i--

				continue
			}
			if value := res.Get("vlanName"); value.Exists() && !data.VlanName.IsNull() {
				data.VlanName = types.StringValue(value.String())
			} else {
				data.VlanName = types.StringNull()
			}
			if value := res.Get("tags"); value.Exists() && !data.Tags.IsNull() {
				data.Tags = helpers.GetStringSet(value.Array())
			} else {
				data.Tags = types.SetNull(types.StringType)
			}
			(*parent).NamedVlansTaggingByApTags[i] = data
		}
		if value := res.Get("oauth.allowedDomains"); value.Exists() && !data.OauthAllowedDomains.IsNull() {
			data.OauthAllowedDomains = helpers.GetStringSet(value.Array())
		} else {
			data.OauthAllowedDomains = types.SetNull(types.StringType)
		}
		if value := res.Get("radiusRadsec.tlsTunnel.timeout"); value.Exists() && !data.RadiusRadsecTlsTunnelTimeout.IsNull() {
			data.RadiusRadsecTlsTunnelTimeout = types.Int64Value(value.Int())
		} else {
			data.RadiusRadsecTlsTunnelTimeout = types.Int64Null()
		}
		if value := res.Get("speedBurst.enabled"); value.Exists() && !data.SpeedBurstEnabled.IsNull() {
			data.SpeedBurstEnabled = types.BoolValue(value.Bool())
		} else {
			data.SpeedBurstEnabled = types.BoolNull()
		}
		for i := 0; i < len(data.ApTagsAndVlanIds); i++ {
			keys := [...]string{"vlanId"}
			keyValues := [...]string{strconv.FormatInt(data.ApTagsAndVlanIds[i].VlanId.ValueInt64(), 10)}

			parent := &data
			data := (*parent).ApTagsAndVlanIds[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("apTagsAndVlanIds").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing ApTagsAndVlanIds[%d] = %+v",
					i,
					(*parent).ApTagsAndVlanIds[i],
				))
				(*parent).ApTagsAndVlanIds = slices.Delete((*parent).ApTagsAndVlanIds, i, i+1)
				i--

				continue
			}
			if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
				data.VlanId = types.Int64Value(value.Int())
			} else {
				data.VlanId = types.Int64Null()
			}
			if value := res.Get("tags"); value.Exists() && !data.Tags.IsNull() {
				data.Tags = helpers.GetStringSet(value.Array())
			} else {
				data.Tags = types.SetNull(types.StringType)
			}
			(*parent).ApTagsAndVlanIds[i] = data
		}
		if value := res.Get("availabilityTags"); value.Exists() && !data.AvailabilityTags.IsNull() {
			data.AvailabilityTags = helpers.GetStringSet(value.Array())
		} else {
			data.AvailabilityTags = types.SetNull(types.StringType)
		}
		{
			l := len(res.Get("radiusAccountingServers").Array())
			tflog.Debug(ctx, fmt.Sprintf("radiusAccountingServers array resizing from %d to %d", len(data.RadiusAccountingServers), l))
			if len(data.RadiusAccountingServers) > l {
				data.RadiusAccountingServers = data.RadiusAccountingServers[:l]
			}
		}
		for i := range data.RadiusAccountingServers {
			parent := &data
			data := (*parent).RadiusAccountingServers[i]
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("radiusAccountingServers.%d", i))
			if value := res.Get("caCertificate"); value.Exists() && !data.CaCertificate.IsNull() {
				data.CaCertificate = types.StringValue(value.String())
			} else {
				data.CaCertificate = types.StringNull()
			}
			if value := res.Get("host"); value.Exists() && !data.Host.IsNull() {
				data.Host = types.StringValue(value.String())
			} else {
				data.Host = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			if value := res.Get("radsecEnabled"); value.Exists() && !data.RadsecEnabled.IsNull() {
				data.RadsecEnabled = types.BoolValue(value.Bool())
			} else {
				data.RadsecEnabled = types.BoolNull()
			}
			(*parent).RadiusAccountingServers[i] = data
		}
		{
			l := len(res.Get("radiusServers").Array())
			tflog.Debug(ctx, fmt.Sprintf("radiusServers array resizing from %d to %d", len(data.RadiusServers), l))
			if len(data.RadiusServers) > l {
				data.RadiusServers = data.RadiusServers[:l]
			}
		}
		for i := range data.RadiusServers {
			parent := &data
			data := (*parent).RadiusServers[i]
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("radiusServers.%d", i))
			if value := res.Get("caCertificate"); value.Exists() && !data.CaCertificate.IsNull() {
				data.CaCertificate = types.StringValue(value.String())
			} else {
				data.CaCertificate = types.StringNull()
			}
			if value := res.Get("host"); value.Exists() && !data.Host.IsNull() {
				data.Host = types.StringValue(value.String())
			} else {
				data.Host = types.StringNull()
			}
			if value := res.Get("openRoamingCertificateId"); value.Exists() && !data.OpenRoamingCertificateId.IsNull() {
				data.OpenRoamingCertificateId = types.Int64Value(value.Int())
			} else {
				data.OpenRoamingCertificateId = types.Int64Null()
			}
			if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			if value := res.Get("radsecEnabled"); value.Exists() && !data.RadsecEnabled.IsNull() {
				data.RadsecEnabled = types.BoolValue(value.Bool())
			} else {
				data.RadsecEnabled = types.BoolNull()
			}
			(*parent).RadiusServers[i] = data
		}
		if value := res.Get("splashGuestSponsorDomains"); value.Exists() && !data.SplashGuestSponsorDomains.IsNull() {
			data.SplashGuestSponsorDomains = helpers.GetStringSet(value.Array())
		} else {
			data.SplashGuestSponsorDomains = types.SetNull(types.StringType)
		}
		if value := res.Get("walledGardenRanges"); value.Exists() && !data.WalledGardenRanges.IsNull() {
			data.WalledGardenRanges = helpers.GetStringSet(value.Array())
		} else {
			data.WalledGardenRanges = types.SetNull(types.StringType)
		}
		if value := res.Get("radiusDasClients.clientsIps"); value.Exists() && !data.RadiusDasClientsIps.IsNull() {
			data.RadiusDasClientsIps = helpers.GetStringSet(value.Array())
		} else {
			data.RadiusDasClientsIps = types.SetNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceWirelessSSIDs) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceWirelessSSIDs) fromBodyImport(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("number").String() == (*parent).Items[i].Number.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
		if value := res.Get("adaptivePolicyGroupId"); value.Exists() {
			data.AdaptivePolicyGroupId = types.StringValue(value.String())
		} else {
			data.AdaptivePolicyGroupId = types.StringNull()
		}
		if value := res.Get("adultContentFilteringEnabled"); value.Exists() {
			data.AdultContentFilteringEnabled = types.BoolValue(value.Bool())
		} else {
			data.AdultContentFilteringEnabled = types.BoolNull()
		}
		if value := res.Get("authMode"); value.Exists() {
			data.AuthMode = types.StringValue(value.String())
		} else {
			data.AuthMode = types.StringNull()
		}
		if value := res.Get("availableOnAllAps"); value.Exists() {
			data.AvailableOnAllAps = types.BoolValue(value.Bool())
		} else {
			data.AvailableOnAllAps = types.BoolNull()
		}
		if value := res.Get("bandSelection"); value.Exists() {
			data.BandSelection = types.StringValue(value.String())
		} else {
			data.BandSelection = types.StringNull()
		}
		if value := res.Get("concentratorNetworkId"); value.Exists() {
			data.ConcentratorNetworkId = types.StringValue(value.String())
		} else {
			data.ConcentratorNetworkId = types.StringNull()
		}
		if value := res.Get("defaultVlanId"); value.Exists() {
			data.DefaultVlanId = types.Int64Value(value.Int())
		} else {
			data.DefaultVlanId = types.Int64Null()
		}
		if value := res.Get("disassociateClientsOnVpnFailover"); value.Exists() {
			data.DisassociateClientsOnVpnFailover = types.BoolValue(value.Bool())
		} else {
			data.DisassociateClientsOnVpnFailover = types.BoolNull()
		}
		if value := res.Get("enabled"); value.Exists() {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("encryptionMode"); value.Exists() {
			data.EncryptionMode = types.StringValue(value.String())
		} else {
			data.EncryptionMode = types.StringNull()
		}
		if value := res.Get("enterpriseAdminAccess"); value.Exists() {
			data.EnterpriseAdminAccess = types.StringValue(value.String())
		} else {
			data.EnterpriseAdminAccess = types.StringNull()
		}
		if value := res.Get("ipAssignmentMode"); value.Exists() {
			data.IpAssignmentMode = types.StringValue(value.String())
		} else {
			data.IpAssignmentMode = types.StringNull()
		}
		if value := res.Get("lanIsolationEnabled"); value.Exists() {
			data.LanIsolationEnabled = types.BoolValue(value.Bool())
		} else {
			data.LanIsolationEnabled = types.BoolNull()
		}
		if value := res.Get("mandatoryDhcpEnabled"); value.Exists() {
			data.MandatoryDhcpEnabled = types.BoolValue(value.Bool())
		} else {
			data.MandatoryDhcpEnabled = types.BoolNull()
		}
		if value := res.Get("minBitrate"); value.Exists() {
			data.MinBitrate = types.Float64Value(value.Float())
		} else {
			data.MinBitrate = types.Float64Null()
		}
		if value := res.Get("name"); value.Exists() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("perClientBandwidthLimitDown"); value.Exists() {
			data.PerClientBandwidthLimitDown = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitDown = types.Int64Null()
		}
		if value := res.Get("perClientBandwidthLimitUp"); value.Exists() {
			data.PerClientBandwidthLimitUp = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitUp = types.Int64Null()
		}
		if value := res.Get("perSsidBandwidthLimitDown"); value.Exists() {
			data.PerSsidBandwidthLimitDown = types.Int64Value(value.Int())
		} else {
			data.PerSsidBandwidthLimitDown = types.Int64Null()
		}
		if value := res.Get("perSsidBandwidthLimitUp"); value.Exists() {
			data.PerSsidBandwidthLimitUp = types.Int64Value(value.Int())
		} else {
			data.PerSsidBandwidthLimitUp = types.Int64Null()
		}
		if value := res.Get("psk"); value.Exists() {
			data.Psk = types.StringValue(value.String())
		} else {
			data.Psk = types.StringNull()
		}
		if value := res.Get("radiusAccountingEnabled"); value.Exists() {
			data.RadiusAccountingEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusAccountingEnabled = types.BoolNull()
		}
		if value := res.Get("radiusAccountingInterimInterval"); value.Exists() {
			data.RadiusAccountingInterimInterval = types.Int64Value(value.Int())
		} else {
			data.RadiusAccountingInterimInterval = types.Int64Null()
		}
		if value := res.Get("radiusAttributeForGroupPolicies"); value.Exists() {
			data.RadiusAttributeForGroupPolicies = types.StringValue(value.String())
		} else {
			data.RadiusAttributeForGroupPolicies = types.StringNull()
		}
		if value := res.Get("radiusAuthenticationNasId"); value.Exists() {
			data.RadiusAuthenticationNasId = types.StringValue(value.String())
		} else {
			data.RadiusAuthenticationNasId = types.StringNull()
		}
		if value := res.Get("radiusCalledStationId"); value.Exists() {
			data.RadiusCalledStationId = types.StringValue(value.String())
		} else {
			data.RadiusCalledStationId = types.StringNull()
		}
		if value := res.Get("radiusCoaEnabled"); value.Exists() {
			data.RadiusCoaEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusCoaEnabled = types.BoolNull()
		}
		if value := res.Get("radiusFailoverPolicy"); value.Exists() {
			data.RadiusFailoverPolicy = types.StringValue(value.String())
		} else {
			data.RadiusFailoverPolicy = types.StringNull()
		}
		if value := res.Get("radiusFallbackEnabled"); value.Exists() {
			data.RadiusFallbackEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusFallbackEnabled = types.BoolNull()
		}
		if value := res.Get("radiusGuestVlanEnabled"); value.Exists() {
			data.RadiusGuestVlanEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusGuestVlanEnabled = types.BoolNull()
		}
		if value := res.Get("radiusGuestVlanId"); value.Exists() {
			data.RadiusGuestVlanId = types.Int64Value(value.Int())
		} else {
			data.RadiusGuestVlanId = types.Int64Null()
		}
		if value := res.Get("radiusLoadBalancingPolicy"); value.Exists() {
			data.RadiusLoadBalancingPolicy = types.StringValue(value.String())
		} else {
			data.RadiusLoadBalancingPolicy = types.StringNull()
		}
		if value := res.Get("radiusOverride"); value.Exists() {
			data.RadiusOverride = types.BoolValue(value.Bool())
		} else {
			data.RadiusOverride = types.BoolNull()
		}
		if value := res.Get("radiusProxyEnabled"); value.Exists() {
			data.RadiusProxyEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusProxyEnabled = types.BoolNull()
		}
		if value := res.Get("radiusServerAttemptsLimit"); value.Exists() {
			data.RadiusServerAttemptsLimit = types.Int64Value(value.Int())
		} else {
			data.RadiusServerAttemptsLimit = types.Int64Null()
		}
		if value := res.Get("radiusServerTimeout"); value.Exists() {
			data.RadiusServerTimeout = types.Int64Value(value.Int())
		} else {
			data.RadiusServerTimeout = types.Int64Null()
		}
		if value := res.Get("radiusTestingEnabled"); value.Exists() {
			data.RadiusTestingEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusTestingEnabled = types.BoolNull()
		}
		if value := res.Get("secondaryConcentratorNetworkId"); value.Exists() {
			data.SecondaryConcentratorNetworkId = types.StringValue(value.String())
		} else {
			data.SecondaryConcentratorNetworkId = types.StringNull()
		}
		if value := res.Get("splashPage"); value.Exists() {
			data.SplashPage = types.StringValue(value.String())
		} else {
			data.SplashPage = types.StringNull()
		}
		if value := res.Get("useVlanTagging"); value.Exists() {
			data.UseVlanTagging = types.BoolValue(value.Bool())
		} else {
			data.UseVlanTagging = types.BoolNull()
		}
		if value := res.Get("visible"); value.Exists() {
			data.Visible = types.BoolValue(value.Bool())
		} else {
			data.Visible = types.BoolNull()
		}
		if value := res.Get("vlanId"); value.Exists() {
			data.VlanId = types.Int64Value(value.Int())
		} else {
			data.VlanId = types.Int64Null()
		}
		if value := res.Get("walledGardenEnabled"); value.Exists() {
			data.WalledGardenEnabled = types.BoolValue(value.Bool())
		} else {
			data.WalledGardenEnabled = types.BoolNull()
		}
		if value := res.Get("wpaEncryptionMode"); value.Exists() {
			data.WpaEncryptionMode = types.StringValue(value.String())
		} else {
			data.WpaEncryptionMode = types.StringNull()
		}
		if value := res.Get("activeDirectory.credentials.logonName"); value.Exists() {
			data.ActiveDirectoryCredentialsLogonName = types.StringValue(value.String())
		} else {
			data.ActiveDirectoryCredentialsLogonName = types.StringNull()
		}
		if value := res.Get("activeDirectory.credentials.password"); value.Exists() {
			data.ActiveDirectoryCredentialsPassword = types.StringValue(value.String())
		} else {
			data.ActiveDirectoryCredentialsPassword = types.StringNull()
		}
		for i := 0; i < len(data.ActiveDirectoryServers); i++ {
			keys := [...]string{"host"}
			keyValues := [...]string{data.ActiveDirectoryServers[i].Host.ValueString()}

			parent := &data
			data := (*parent).ActiveDirectoryServers[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("activeDirectory.servers").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing ActiveDirectoryServers[%d] = %+v",
					i,
					(*parent).ActiveDirectoryServers[i],
				))
				(*parent).ActiveDirectoryServers = slices.Delete((*parent).ActiveDirectoryServers, i, i+1)
				i--

				continue
			}
			if value := res.Get("host"); value.Exists() {
				data.Host = types.StringValue(value.String())
			} else {
				data.Host = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			(*parent).ActiveDirectoryServers[i] = data
		}
		if value := res.Get("dnsRewrite.enabled"); value.Exists() {
			data.DnsRewriteEnabled = types.BoolValue(value.Bool())
		} else {
			data.DnsRewriteEnabled = types.BoolNull()
		}
		if value := res.Get("dnsRewrite.dnsCustomNameservers"); value.Exists() {
			data.DnsRewriteDnsCustomNameservers = helpers.GetStringList(value.Array())
		} else {
			data.DnsRewriteDnsCustomNameservers = types.ListNull(types.StringType)
		}
		if value := res.Get("dot11r.adaptive"); value.Exists() {
			data.Dot11rAdaptive = types.BoolValue(value.Bool())
		} else {
			data.Dot11rAdaptive = types.BoolNull()
		}
		if value := res.Get("dot11r.enabled"); value.Exists() {
			data.Dot11rEnabled = types.BoolValue(value.Bool())
		} else {
			data.Dot11rEnabled = types.BoolNull()
		}
		if value := res.Get("dot11w.enabled"); value.Exists() {
			data.Dot11wEnabled = types.BoolValue(value.Bool())
		} else {
			data.Dot11wEnabled = types.BoolNull()
		}
		if value := res.Get("dot11w.required"); value.Exists() {
			data.Dot11wRequired = types.BoolValue(value.Bool())
		} else {
			data.Dot11wRequired = types.BoolNull()
		}
		if value := res.Get("gre.key"); value.Exists() {
			data.GreKey = types.Int64Value(value.Int())
		} else {
			data.GreKey = types.Int64Null()
		}
		if value := res.Get("gre.concentrator.host"); value.Exists() {
			data.GreConcentratorHost = types.StringValue(value.String())
		} else {
			data.GreConcentratorHost = types.StringNull()
		}
		if value := res.Get("ldap.baseDistinguishedName"); value.Exists() {
			data.LdapBaseDistinguishedName = types.StringValue(value.String())
		} else {
			data.LdapBaseDistinguishedName = types.StringNull()
		}
		if value := res.Get("ldap.credentials.distinguishedName"); value.Exists() {
			data.LdapCredentialsDistinguishedName = types.StringValue(value.String())
		} else {
			data.LdapCredentialsDistinguishedName = types.StringNull()
		}
		if value := res.Get("ldap.credentials.password"); value.Exists() {
			data.LdapCredentialsPassword = types.StringValue(value.String())
		} else {
			data.LdapCredentialsPassword = types.StringNull()
		}
		if value := res.Get("ldap.serverCaCertificate.contents"); value.Exists() {
			data.LdapServerCaCertificateContents = types.StringValue(value.String())
		} else {
			data.LdapServerCaCertificateContents = types.StringNull()
		}
		{
			l := len(res.Get("ldap.servers").Array())
			tflog.Debug(ctx, fmt.Sprintf("ldap.servers array resizing from %d to %d", len(data.LdapServers), l))
			if len(data.LdapServers) > l {
				data.LdapServers = data.LdapServers[:l]
			}
		}
		for i := range data.LdapServers {
			parent := &data
			data := (*parent).LdapServers[i]
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("ldap.servers.%d", i))
			if value := res.Get("host"); value.Exists() {
				data.Host = types.StringValue(value.String())
			} else {
				data.Host = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			(*parent).LdapServers[i] = data
		}
		if value := res.Get("localRadius.cacheTimeout"); value.Exists() {
			data.LocalRadiusCacheTimeout = types.Int64Value(value.Int())
		} else {
			data.LocalRadiusCacheTimeout = types.Int64Null()
		}
		if value := res.Get("localRadius.certificateAuthentication.enabled"); value.Exists() {
			data.LocalRadiusCertificateAuthenticationEnabled = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusCertificateAuthenticationEnabled = types.BoolNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.ocspResponderUrl"); value.Exists() {
			data.LocalRadiusCertificateAuthenticationOcspResponderUrl = types.StringValue(value.String())
		} else {
			data.LocalRadiusCertificateAuthenticationOcspResponderUrl = types.StringNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.useLdap"); value.Exists() {
			data.LocalRadiusCertificateAuthenticationUseLdap = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusCertificateAuthenticationUseLdap = types.BoolNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.useOcsp"); value.Exists() {
			data.LocalRadiusCertificateAuthenticationUseOcsp = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusCertificateAuthenticationUseOcsp = types.BoolNull()
		}
		if value := res.Get("localRadius.certificateAuthentication.clientRootCaCertificate.contents"); value.Exists() {
			data.LocalRadiusCertificateAuthenticationClientRootCaCertificateContents = types.StringValue(value.String())
		} else {
			data.LocalRadiusCertificateAuthenticationClientRootCaCertificateContents = types.StringNull()
		}
		if value := res.Get("localRadius.passwordAuthentication.enabled"); value.Exists() {
			data.LocalRadiusPasswordAuthenticationEnabled = types.BoolValue(value.Bool())
		} else {
			data.LocalRadiusPasswordAuthenticationEnabled = types.BoolNull()
		}
		if value := res.Get("namedVlans.radius.guestVlan.enabled"); value.Exists() {
			data.NamedVlansRadiusGuestVlanEnabled = types.BoolValue(value.Bool())
		} else {
			data.NamedVlansRadiusGuestVlanEnabled = types.BoolNull()
		}
		if value := res.Get("namedVlans.radius.guestVlan.name"); value.Exists() {
			data.NamedVlansRadiusGuestVlanName = types.StringValue(value.String())
		} else {
			data.NamedVlansRadiusGuestVlanName = types.StringNull()
		}
		if value := res.Get("namedVlans.tagging.defaultVlanName"); value.Exists() {
			data.NamedVlansTaggingDefaultVlanName = types.StringValue(value.String())
		} else {
			data.NamedVlansTaggingDefaultVlanName = types.StringNull()
		}
		if value := res.Get("namedVlans.tagging.enabled"); value.Exists() {
			data.NamedVlansTaggingEnabled = types.BoolValue(value.Bool())
		} else {
			data.NamedVlansTaggingEnabled = types.BoolNull()
		}
		for i := 0; i < len(data.NamedVlansTaggingByApTags); i++ {
			keys := [...]string{"vlanName"}
			keyValues := [...]string{data.NamedVlansTaggingByApTags[i].VlanName.ValueString()}

			parent := &data
			data := (*parent).NamedVlansTaggingByApTags[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("namedVlans.tagging.byApTags").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing NamedVlansTaggingByApTags[%d] = %+v",
					i,
					(*parent).NamedVlansTaggingByApTags[i],
				))
				(*parent).NamedVlansTaggingByApTags = slices.Delete((*parent).NamedVlansTaggingByApTags, i, i+1)
				i--

				continue
			}
			if value := res.Get("vlanName"); value.Exists() {
				data.VlanName = types.StringValue(value.String())
			} else {
				data.VlanName = types.StringNull()
			}
			if value := res.Get("tags"); value.Exists() {
				data.Tags = helpers.GetStringSet(value.Array())
			} else {
				data.Tags = types.SetNull(types.StringType)
			}
			(*parent).NamedVlansTaggingByApTags[i] = data
		}
		if value := res.Get("oauth.allowedDomains"); value.Exists() {
			data.OauthAllowedDomains = helpers.GetStringSet(value.Array())
		} else {
			data.OauthAllowedDomains = types.SetNull(types.StringType)
		}
		if value := res.Get("radiusRadsec.tlsTunnel.timeout"); value.Exists() {
			data.RadiusRadsecTlsTunnelTimeout = types.Int64Value(value.Int())
		} else {
			data.RadiusRadsecTlsTunnelTimeout = types.Int64Null()
		}
		if value := res.Get("speedBurst.enabled"); value.Exists() {
			data.SpeedBurstEnabled = types.BoolValue(value.Bool())
		} else {
			data.SpeedBurstEnabled = types.BoolNull()
		}
		for i := 0; i < len(data.ApTagsAndVlanIds); i++ {
			keys := [...]string{"vlanId"}
			keyValues := [...]string{strconv.FormatInt(data.ApTagsAndVlanIds[i].VlanId.ValueInt64(), 10)}

			parent := &data
			data := (*parent).ApTagsAndVlanIds[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("apTagsAndVlanIds").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing ApTagsAndVlanIds[%d] = %+v",
					i,
					(*parent).ApTagsAndVlanIds[i],
				))
				(*parent).ApTagsAndVlanIds = slices.Delete((*parent).ApTagsAndVlanIds, i, i+1)
				i--

				continue
			}
			if value := res.Get("vlanId"); value.Exists() {
				data.VlanId = types.Int64Value(value.Int())
			} else {
				data.VlanId = types.Int64Null()
			}
			if value := res.Get("tags"); value.Exists() {
				data.Tags = helpers.GetStringSet(value.Array())
			} else {
				data.Tags = types.SetNull(types.StringType)
			}
			(*parent).ApTagsAndVlanIds[i] = data
		}
		if value := res.Get("availabilityTags"); value.Exists() {
			data.AvailabilityTags = helpers.GetStringSet(value.Array())
		} else {
			data.AvailabilityTags = types.SetNull(types.StringType)
		}
		{
			l := len(res.Get("radiusAccountingServers").Array())
			tflog.Debug(ctx, fmt.Sprintf("radiusAccountingServers array resizing from %d to %d", len(data.RadiusAccountingServers), l))
			if len(data.RadiusAccountingServers) > l {
				data.RadiusAccountingServers = data.RadiusAccountingServers[:l]
			}
		}
		for i := range data.RadiusAccountingServers {
			parent := &data
			data := (*parent).RadiusAccountingServers[i]
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("radiusAccountingServers.%d", i))
			if value := res.Get("caCertificate"); value.Exists() {
				data.CaCertificate = types.StringValue(value.String())
			} else {
				data.CaCertificate = types.StringNull()
			}
			if value := res.Get("host"); value.Exists() {
				data.Host = types.StringValue(value.String())
			} else {
				data.Host = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			if value := res.Get("radsecEnabled"); value.Exists() {
				data.RadsecEnabled = types.BoolValue(value.Bool())
			} else {
				data.RadsecEnabled = types.BoolNull()
			}
			(*parent).RadiusAccountingServers[i] = data
		}
		{
			l := len(res.Get("radiusServers").Array())
			tflog.Debug(ctx, fmt.Sprintf("radiusServers array resizing from %d to %d", len(data.RadiusServers), l))
			if len(data.RadiusServers) > l {
				data.RadiusServers = data.RadiusServers[:l]
			}
		}
		for i := range data.RadiusServers {
			parent := &data
			data := (*parent).RadiusServers[i]
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("radiusServers.%d", i))
			if value := res.Get("caCertificate"); value.Exists() {
				data.CaCertificate = types.StringValue(value.String())
			} else {
				data.CaCertificate = types.StringNull()
			}
			if value := res.Get("host"); value.Exists() {
				data.Host = types.StringValue(value.String())
			} else {
				data.Host = types.StringNull()
			}
			if value := res.Get("openRoamingCertificateId"); value.Exists() {
				data.OpenRoamingCertificateId = types.Int64Value(value.Int())
			} else {
				data.OpenRoamingCertificateId = types.Int64Null()
			}
			if value := res.Get("port"); value.Exists() {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			if value := res.Get("radsecEnabled"); value.Exists() {
				data.RadsecEnabled = types.BoolValue(value.Bool())
			} else {
				data.RadsecEnabled = types.BoolNull()
			}
			(*parent).RadiusServers[i] = data
		}
		if value := res.Get("splashGuestSponsorDomains"); value.Exists() {
			data.SplashGuestSponsorDomains = helpers.GetStringSet(value.Array())
		} else {
			data.SplashGuestSponsorDomains = types.SetNull(types.StringType)
		}
		if value := res.Get("walledGardenRanges"); value.Exists() {
			data.WalledGardenRanges = helpers.GetStringSet(value.Array())
		} else {
			data.WalledGardenRanges = types.SetNull(types.StringType)
		}
		if value := res.Get("radiusDasClients.clientsIps"); value.Exists() {
			data.RadiusDasClientsIps = helpers.GetStringSet(value.Array())
		} else {
			data.RadiusDasClientsIps = types.SetNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceWirelessSSIDs) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceWirelessSSIDs) hasChanges(ctx context.Context, state *ResourceWirelessSSIDs, id string) bool {
	hasChanges := false

	item := ResourceWirelessSSIDsItems{}
	for i := range data.Items {
		if data.Items[i].Number.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceWirelessSSIDsItems{}
	for i := range state.Items {
		if state.Items[i].Number.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.Number.Equal(stateItem.Number) {
		hasChanges = true
	}
	if !item.AdaptivePolicyGroupId.Equal(stateItem.AdaptivePolicyGroupId) {
		hasChanges = true
	}
	if !item.AdultContentFilteringEnabled.Equal(stateItem.AdultContentFilteringEnabled) {
		hasChanges = true
	}
	if !item.AuthMode.Equal(stateItem.AuthMode) {
		hasChanges = true
	}
	if !item.AvailableOnAllAps.Equal(stateItem.AvailableOnAllAps) {
		hasChanges = true
	}
	if !item.BandSelection.Equal(stateItem.BandSelection) {
		hasChanges = true
	}
	if !item.ConcentratorNetworkId.Equal(stateItem.ConcentratorNetworkId) {
		hasChanges = true
	}
	if !item.DefaultVlanId.Equal(stateItem.DefaultVlanId) {
		hasChanges = true
	}
	if !item.DisassociateClientsOnVpnFailover.Equal(stateItem.DisassociateClientsOnVpnFailover) {
		hasChanges = true
	}
	if !item.Enabled.Equal(stateItem.Enabled) {
		hasChanges = true
	}
	if !item.EncryptionMode.Equal(stateItem.EncryptionMode) {
		hasChanges = true
	}
	if !item.EnterpriseAdminAccess.Equal(stateItem.EnterpriseAdminAccess) {
		hasChanges = true
	}
	if !item.IpAssignmentMode.Equal(stateItem.IpAssignmentMode) {
		hasChanges = true
	}
	if !item.LanIsolationEnabled.Equal(stateItem.LanIsolationEnabled) {
		hasChanges = true
	}
	if !item.MandatoryDhcpEnabled.Equal(stateItem.MandatoryDhcpEnabled) {
		hasChanges = true
	}
	if !item.MinBitrate.Equal(stateItem.MinBitrate) {
		hasChanges = true
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	if !item.PerClientBandwidthLimitDown.Equal(stateItem.PerClientBandwidthLimitDown) {
		hasChanges = true
	}
	if !item.PerClientBandwidthLimitUp.Equal(stateItem.PerClientBandwidthLimitUp) {
		hasChanges = true
	}
	if !item.PerSsidBandwidthLimitDown.Equal(stateItem.PerSsidBandwidthLimitDown) {
		hasChanges = true
	}
	if !item.PerSsidBandwidthLimitUp.Equal(stateItem.PerSsidBandwidthLimitUp) {
		hasChanges = true
	}
	if !item.Psk.Equal(stateItem.Psk) {
		hasChanges = true
	}
	if !item.RadiusAccountingEnabled.Equal(stateItem.RadiusAccountingEnabled) {
		hasChanges = true
	}
	if !item.RadiusAccountingInterimInterval.Equal(stateItem.RadiusAccountingInterimInterval) {
		hasChanges = true
	}
	if !item.RadiusAttributeForGroupPolicies.Equal(stateItem.RadiusAttributeForGroupPolicies) {
		hasChanges = true
	}
	if !item.RadiusAuthenticationNasId.Equal(stateItem.RadiusAuthenticationNasId) {
		hasChanges = true
	}
	if !item.RadiusCalledStationId.Equal(stateItem.RadiusCalledStationId) {
		hasChanges = true
	}
	if !item.RadiusCoaEnabled.Equal(stateItem.RadiusCoaEnabled) {
		hasChanges = true
	}
	if !item.RadiusFailoverPolicy.Equal(stateItem.RadiusFailoverPolicy) {
		hasChanges = true
	}
	if !item.RadiusFallbackEnabled.Equal(stateItem.RadiusFallbackEnabled) {
		hasChanges = true
	}
	if !item.RadiusGuestVlanEnabled.Equal(stateItem.RadiusGuestVlanEnabled) {
		hasChanges = true
	}
	if !item.RadiusGuestVlanId.Equal(stateItem.RadiusGuestVlanId) {
		hasChanges = true
	}
	if !item.RadiusLoadBalancingPolicy.Equal(stateItem.RadiusLoadBalancingPolicy) {
		hasChanges = true
	}
	if !item.RadiusOverride.Equal(stateItem.RadiusOverride) {
		hasChanges = true
	}
	if !item.RadiusProxyEnabled.Equal(stateItem.RadiusProxyEnabled) {
		hasChanges = true
	}
	if !item.RadiusServerAttemptsLimit.Equal(stateItem.RadiusServerAttemptsLimit) {
		hasChanges = true
	}
	if !item.RadiusServerTimeout.Equal(stateItem.RadiusServerTimeout) {
		hasChanges = true
	}
	if !item.RadiusTestingEnabled.Equal(stateItem.RadiusTestingEnabled) {
		hasChanges = true
	}
	if !item.SecondaryConcentratorNetworkId.Equal(stateItem.SecondaryConcentratorNetworkId) {
		hasChanges = true
	}
	if !item.SplashPage.Equal(stateItem.SplashPage) {
		hasChanges = true
	}
	if !item.UseVlanTagging.Equal(stateItem.UseVlanTagging) {
		hasChanges = true
	}
	if !item.Visible.Equal(stateItem.Visible) {
		hasChanges = true
	}
	if !item.VlanId.Equal(stateItem.VlanId) {
		hasChanges = true
	}
	if !item.WalledGardenEnabled.Equal(stateItem.WalledGardenEnabled) {
		hasChanges = true
	}
	if !item.WpaEncryptionMode.Equal(stateItem.WpaEncryptionMode) {
		hasChanges = true
	}
	if !item.ActiveDirectoryCredentialsLogonName.Equal(stateItem.ActiveDirectoryCredentialsLogonName) {
		hasChanges = true
	}
	if !item.ActiveDirectoryCredentialsPassword.Equal(stateItem.ActiveDirectoryCredentialsPassword) {
		hasChanges = true
	}
	if len(item.ActiveDirectoryServers) != len(stateItem.ActiveDirectoryServers) {
		hasChanges = true
	} else {
		for i := range item.ActiveDirectoryServers {
			if !item.ActiveDirectoryServers[i].Host.Equal(stateItem.ActiveDirectoryServers[i].Host) {
				hasChanges = true
			}
			if !item.ActiveDirectoryServers[i].Port.Equal(stateItem.ActiveDirectoryServers[i].Port) {
				hasChanges = true
			}
		}
	}
	if !item.DnsRewriteEnabled.Equal(stateItem.DnsRewriteEnabled) {
		hasChanges = true
	}
	if !item.DnsRewriteDnsCustomNameservers.Equal(stateItem.DnsRewriteDnsCustomNameservers) {
		hasChanges = true
	}
	if !item.Dot11rAdaptive.Equal(stateItem.Dot11rAdaptive) {
		hasChanges = true
	}
	if !item.Dot11rEnabled.Equal(stateItem.Dot11rEnabled) {
		hasChanges = true
	}
	if !item.Dot11wEnabled.Equal(stateItem.Dot11wEnabled) {
		hasChanges = true
	}
	if !item.Dot11wRequired.Equal(stateItem.Dot11wRequired) {
		hasChanges = true
	}
	if !item.GreKey.Equal(stateItem.GreKey) {
		hasChanges = true
	}
	if !item.GreConcentratorHost.Equal(stateItem.GreConcentratorHost) {
		hasChanges = true
	}
	if !item.LdapBaseDistinguishedName.Equal(stateItem.LdapBaseDistinguishedName) {
		hasChanges = true
	}
	if !item.LdapCredentialsDistinguishedName.Equal(stateItem.LdapCredentialsDistinguishedName) {
		hasChanges = true
	}
	if !item.LdapCredentialsPassword.Equal(stateItem.LdapCredentialsPassword) {
		hasChanges = true
	}
	if !item.LdapServerCaCertificateContents.Equal(stateItem.LdapServerCaCertificateContents) {
		hasChanges = true
	}
	if len(item.LdapServers) != len(stateItem.LdapServers) {
		hasChanges = true
	} else {
		for i := range item.LdapServers {
			if !item.LdapServers[i].Host.Equal(stateItem.LdapServers[i].Host) {
				hasChanges = true
			}
			if !item.LdapServers[i].Port.Equal(stateItem.LdapServers[i].Port) {
				hasChanges = true
			}
		}
	}
	if !item.LocalRadiusCacheTimeout.Equal(stateItem.LocalRadiusCacheTimeout) {
		hasChanges = true
	}
	if !item.LocalRadiusCertificateAuthenticationEnabled.Equal(stateItem.LocalRadiusCertificateAuthenticationEnabled) {
		hasChanges = true
	}
	if !item.LocalRadiusCertificateAuthenticationOcspResponderUrl.Equal(stateItem.LocalRadiusCertificateAuthenticationOcspResponderUrl) {
		hasChanges = true
	}
	if !item.LocalRadiusCertificateAuthenticationUseLdap.Equal(stateItem.LocalRadiusCertificateAuthenticationUseLdap) {
		hasChanges = true
	}
	if !item.LocalRadiusCertificateAuthenticationUseOcsp.Equal(stateItem.LocalRadiusCertificateAuthenticationUseOcsp) {
		hasChanges = true
	}
	if !item.LocalRadiusCertificateAuthenticationClientRootCaCertificateContents.Equal(stateItem.LocalRadiusCertificateAuthenticationClientRootCaCertificateContents) {
		hasChanges = true
	}
	if !item.LocalRadiusPasswordAuthenticationEnabled.Equal(stateItem.LocalRadiusPasswordAuthenticationEnabled) {
		hasChanges = true
	}
	if !item.NamedVlansRadiusGuestVlanEnabled.Equal(stateItem.NamedVlansRadiusGuestVlanEnabled) {
		hasChanges = true
	}
	if !item.NamedVlansRadiusGuestVlanName.Equal(stateItem.NamedVlansRadiusGuestVlanName) {
		hasChanges = true
	}
	if !item.NamedVlansTaggingDefaultVlanName.Equal(stateItem.NamedVlansTaggingDefaultVlanName) {
		hasChanges = true
	}
	if !item.NamedVlansTaggingEnabled.Equal(stateItem.NamedVlansTaggingEnabled) {
		hasChanges = true
	}
	if len(item.NamedVlansTaggingByApTags) != len(stateItem.NamedVlansTaggingByApTags) {
		hasChanges = true
	} else {
		for i := range item.NamedVlansTaggingByApTags {
			if !item.NamedVlansTaggingByApTags[i].VlanName.Equal(stateItem.NamedVlansTaggingByApTags[i].VlanName) {
				hasChanges = true
			}
			if !item.NamedVlansTaggingByApTags[i].Tags.Equal(stateItem.NamedVlansTaggingByApTags[i].Tags) {
				hasChanges = true
			}
		}
	}
	if !item.OauthAllowedDomains.Equal(stateItem.OauthAllowedDomains) {
		hasChanges = true
	}
	if !item.RadiusRadsecTlsTunnelTimeout.Equal(stateItem.RadiusRadsecTlsTunnelTimeout) {
		hasChanges = true
	}
	if !item.SpeedBurstEnabled.Equal(stateItem.SpeedBurstEnabled) {
		hasChanges = true
	}
	if len(item.ApTagsAndVlanIds) != len(stateItem.ApTagsAndVlanIds) {
		hasChanges = true
	} else {
		for i := range item.ApTagsAndVlanIds {
			if !item.ApTagsAndVlanIds[i].VlanId.Equal(stateItem.ApTagsAndVlanIds[i].VlanId) {
				hasChanges = true
			}
			if !item.ApTagsAndVlanIds[i].Tags.Equal(stateItem.ApTagsAndVlanIds[i].Tags) {
				hasChanges = true
			}
		}
	}
	if !item.AvailabilityTags.Equal(stateItem.AvailabilityTags) {
		hasChanges = true
	}
	if len(item.RadiusAccountingServers) != len(stateItem.RadiusAccountingServers) {
		hasChanges = true
	} else {
		for i := range item.RadiusAccountingServers {
			if !item.RadiusAccountingServers[i].CaCertificate.Equal(stateItem.RadiusAccountingServers[i].CaCertificate) {
				hasChanges = true
			}
			if !item.RadiusAccountingServers[i].Host.Equal(stateItem.RadiusAccountingServers[i].Host) {
				hasChanges = true
			}
			if !item.RadiusAccountingServers[i].Port.Equal(stateItem.RadiusAccountingServers[i].Port) {
				hasChanges = true
			}
			if !item.RadiusAccountingServers[i].RadsecEnabled.Equal(stateItem.RadiusAccountingServers[i].RadsecEnabled) {
				hasChanges = true
			}
			if !item.RadiusAccountingServers[i].Secret.Equal(stateItem.RadiusAccountingServers[i].Secret) {
				hasChanges = true
			}
		}
	}
	if len(item.RadiusServers) != len(stateItem.RadiusServers) {
		hasChanges = true
	} else {
		for i := range item.RadiusServers {
			if !item.RadiusServers[i].CaCertificate.Equal(stateItem.RadiusServers[i].CaCertificate) {
				hasChanges = true
			}
			if !item.RadiusServers[i].Host.Equal(stateItem.RadiusServers[i].Host) {
				hasChanges = true
			}
			if !item.RadiusServers[i].OpenRoamingCertificateId.Equal(stateItem.RadiusServers[i].OpenRoamingCertificateId) {
				hasChanges = true
			}
			if !item.RadiusServers[i].Port.Equal(stateItem.RadiusServers[i].Port) {
				hasChanges = true
			}
			if !item.RadiusServers[i].RadsecEnabled.Equal(stateItem.RadiusServers[i].RadsecEnabled) {
				hasChanges = true
			}
			if !item.RadiusServers[i].Secret.Equal(stateItem.RadiusServers[i].Secret) {
				hasChanges = true
			}
		}
	}
	if !item.SplashGuestSponsorDomains.Equal(stateItem.SplashGuestSponsorDomains) {
		hasChanges = true
	}
	if !item.WalledGardenRanges.Equal(stateItem.WalledGardenRanges) {
		hasChanges = true
	}
	if !item.RadiusDasClientsIps.Equal(stateItem.RadiusDasClientsIps) {
		hasChanges = true
	}
	if !item.RadiusDasClientsSharedSecret.Equal(stateItem.RadiusDasClientsSharedSecret) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
