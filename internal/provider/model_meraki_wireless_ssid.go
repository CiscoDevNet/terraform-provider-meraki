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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessSSID struct {
	Id                                                                  types.String                            `tfsdk:"id"`
	NetworkId                                                           types.String                            `tfsdk:"network_id"`
	Number                                                              types.String                            `tfsdk:"number"`
	AdultContentFilteringEnabled                                        types.Bool                              `tfsdk:"adult_content_filtering_enabled"`
	AuthMode                                                            types.String                            `tfsdk:"auth_mode"`
	AvailableOnAllAps                                                   types.Bool                              `tfsdk:"available_on_all_aps"`
	BandSelection                                                       types.String                            `tfsdk:"band_selection"`
	ConcentratorNetworkId                                               types.String                            `tfsdk:"concentrator_network_id"`
	DefaultVlanId                                                       types.Int64                             `tfsdk:"default_vlan_id"`
	DisassociateClientsOnVpnFailover                                    types.Bool                              `tfsdk:"disassociate_clients_on_vpn_failover"`
	Enabled                                                             types.Bool                              `tfsdk:"enabled"`
	EncryptionMode                                                      types.String                            `tfsdk:"encryption_mode"`
	EnterpriseAdminAccess                                               types.String                            `tfsdk:"enterprise_admin_access"`
	IpAssignmentMode                                                    types.String                            `tfsdk:"ip_assignment_mode"`
	LanIsolationEnabled                                                 types.Bool                              `tfsdk:"lan_isolation_enabled"`
	MandatoryDhcpEnabled                                                types.Bool                              `tfsdk:"mandatory_dhcp_enabled"`
	MinBitrate                                                          types.Float64                           `tfsdk:"min_bitrate"`
	Name                                                                types.String                            `tfsdk:"name"`
	PerClientBandwidthLimitDown                                         types.Int64                             `tfsdk:"per_client_bandwidth_limit_down"`
	PerClientBandwidthLimitUp                                           types.Int64                             `tfsdk:"per_client_bandwidth_limit_up"`
	PerSsidBandwidthLimitDown                                           types.Int64                             `tfsdk:"per_ssid_bandwidth_limit_down"`
	PerSsidBandwidthLimitUp                                             types.Int64                             `tfsdk:"per_ssid_bandwidth_limit_up"`
	Psk                                                                 types.String                            `tfsdk:"psk"`
	RadiusAccountingEnabled                                             types.Bool                              `tfsdk:"radius_accounting_enabled"`
	RadiusAccountingInterimInterval                                     types.Int64                             `tfsdk:"radius_accounting_interim_interval"`
	RadiusAttributeForGroupPolicies                                     types.String                            `tfsdk:"radius_attribute_for_group_policies"`
	RadiusAuthenticationNasId                                           types.String                            `tfsdk:"radius_authentication_nas_id"`
	RadiusCalledStationId                                               types.String                            `tfsdk:"radius_called_station_id"`
	RadiusCoaEnabled                                                    types.Bool                              `tfsdk:"radius_coa_enabled"`
	RadiusFailoverPolicy                                                types.String                            `tfsdk:"radius_failover_policy"`
	RadiusFallbackEnabled                                               types.Bool                              `tfsdk:"radius_fallback_enabled"`
	RadiusGuestVlanEnabled                                              types.Bool                              `tfsdk:"radius_guest_vlan_enabled"`
	RadiusGuestVlanId                                                   types.Int64                             `tfsdk:"radius_guest_vlan_id"`
	RadiusLoadBalancingPolicy                                           types.String                            `tfsdk:"radius_load_balancing_policy"`
	RadiusOverride                                                      types.Bool                              `tfsdk:"radius_override"`
	RadiusProxyEnabled                                                  types.Bool                              `tfsdk:"radius_proxy_enabled"`
	RadiusServerAttemptsLimit                                           types.Int64                             `tfsdk:"radius_server_attempts_limit"`
	RadiusServerTimeout                                                 types.Int64                             `tfsdk:"radius_server_timeout"`
	RadiusTestingEnabled                                                types.Bool                              `tfsdk:"radius_testing_enabled"`
	SecondaryConcentratorNetworkId                                      types.String                            `tfsdk:"secondary_concentrator_network_id"`
	SplashPage                                                          types.String                            `tfsdk:"splash_page"`
	UseVlanTagging                                                      types.Bool                              `tfsdk:"use_vlan_tagging"`
	Visible                                                             types.Bool                              `tfsdk:"visible"`
	VlanId                                                              types.Int64                             `tfsdk:"vlan_id"`
	WalledGardenEnabled                                                 types.Bool                              `tfsdk:"walled_garden_enabled"`
	WpaEncryptionMode                                                   types.String                            `tfsdk:"wpa_encryption_mode"`
	ActiveDirectoryCredentialsLogonName                                 types.String                            `tfsdk:"active_directory_credentials_logon_name"`
	ActiveDirectoryCredentialsPassword                                  types.String                            `tfsdk:"active_directory_credentials_password"`
	ActiveDirectoryServers                                              []WirelessSSIDActiveDirectoryServers    `tfsdk:"active_directory_servers"`
	DnsRewriteEnabled                                                   types.Bool                              `tfsdk:"dns_rewrite_enabled"`
	DnsRewriteDnsCustomNameservers                                      types.List                              `tfsdk:"dns_rewrite_dns_custom_nameservers"`
	Dot11rAdaptive                                                      types.Bool                              `tfsdk:"dot11r_adaptive"`
	Dot11rEnabled                                                       types.Bool                              `tfsdk:"dot11r_enabled"`
	Dot11wEnabled                                                       types.Bool                              `tfsdk:"dot11w_enabled"`
	Dot11wRequired                                                      types.Bool                              `tfsdk:"dot11w_required"`
	GreKey                                                              types.Int64                             `tfsdk:"gre_key"`
	GreConcentratorHost                                                 types.String                            `tfsdk:"gre_concentrator_host"`
	LdapBaseDistinguishedName                                           types.String                            `tfsdk:"ldap_base_distinguished_name"`
	LdapCredentialsDistinguishedName                                    types.String                            `tfsdk:"ldap_credentials_distinguished_name"`
	LdapCredentialsPassword                                             types.String                            `tfsdk:"ldap_credentials_password"`
	LdapServerCaCertificateContents                                     types.String                            `tfsdk:"ldap_server_ca_certificate_contents"`
	LdapServers                                                         []WirelessSSIDLdapServers               `tfsdk:"ldap_servers"`
	LocalRadiusCacheTimeout                                             types.Int64                             `tfsdk:"local_radius_cache_timeout"`
	LocalRadiusCertificateAuthenticationEnabled                         types.Bool                              `tfsdk:"local_radius_certificate_authentication_enabled"`
	LocalRadiusCertificateAuthenticationOcspResponderUrl                types.String                            `tfsdk:"local_radius_certificate_authentication_ocsp_responder_url"`
	LocalRadiusCertificateAuthenticationUseLdap                         types.Bool                              `tfsdk:"local_radius_certificate_authentication_use_ldap"`
	LocalRadiusCertificateAuthenticationUseOcsp                         types.Bool                              `tfsdk:"local_radius_certificate_authentication_use_ocsp"`
	LocalRadiusCertificateAuthenticationClientRootCaCertificateContents types.String                            `tfsdk:"local_radius_certificate_authentication_client_root_ca_certificate_contents"`
	LocalRadiusPasswordAuthenticationEnabled                            types.Bool                              `tfsdk:"local_radius_password_authentication_enabled"`
	NamedVlansRadiusGuestVlanEnabled                                    types.Bool                              `tfsdk:"named_vlans_radius_guest_vlan_enabled"`
	NamedVlansRadiusGuestVlanName                                       types.String                            `tfsdk:"named_vlans_radius_guest_vlan_name"`
	NamedVlansTaggingDefaultVlanName                                    types.String                            `tfsdk:"named_vlans_tagging_default_vlan_name"`
	NamedVlansTaggingEnabled                                            types.Bool                              `tfsdk:"named_vlans_tagging_enabled"`
	NamedVlansTaggingByApTags                                           []WirelessSSIDNamedVlansTaggingByApTags `tfsdk:"named_vlans_tagging_by_ap_tags"`
	OauthAllowedDomains                                                 types.List                              `tfsdk:"oauth_allowed_domains"`
	SpeedBurstEnabled                                                   types.Bool                              `tfsdk:"speed_burst_enabled"`
	ApTagsAndVlanIds                                                    []WirelessSSIDApTagsAndVlanIds          `tfsdk:"ap_tags_and_vlan_ids"`
	AvailabilityTags                                                    types.List                              `tfsdk:"availability_tags"`
	RadiusAccountingServers                                             []WirelessSSIDRadiusAccountingServers   `tfsdk:"radius_accounting_servers"`
	RadiusServers                                                       []WirelessSSIDRadiusServers             `tfsdk:"radius_servers"`
	SplashGuestSponsorDomains                                           types.List                              `tfsdk:"splash_guest_sponsor_domains"`
	WalledGardenRanges                                                  types.List                              `tfsdk:"walled_garden_ranges"`
}

type WirelessSSIDActiveDirectoryServers struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type WirelessSSIDLdapServers struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type WirelessSSIDNamedVlansTaggingByApTags struct {
	VlanName types.String `tfsdk:"vlan_name"`
	Tags     types.List   `tfsdk:"tags"`
}

type WirelessSSIDApTagsAndVlanIds struct {
	VlanId types.Int64 `tfsdk:"vlan_id"`
	Tags   types.List  `tfsdk:"tags"`
}

type WirelessSSIDRadiusAccountingServers struct {
	CaCertificate types.String `tfsdk:"ca_certificate"`
	Host          types.String `tfsdk:"host"`
	Port          types.Int64  `tfsdk:"port"`
	RadsecEnabled types.Bool   `tfsdk:"radsec_enabled"`
	Secret        types.String `tfsdk:"secret"`
}

type WirelessSSIDRadiusServers struct {
	CaCertificate            types.String `tfsdk:"ca_certificate"`
	Host                     types.String `tfsdk:"host"`
	OpenRoamingCertificateId types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
	RadsecEnabled            types.Bool   `tfsdk:"radsec_enabled"`
	Secret                   types.String `tfsdk:"secret"`
}

// End of section. //template:end types

func (data WirelessSSID) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids", url.QueryEscape(data.NetworkId.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSSID) toBody(ctx context.Context, state WirelessSSID) string {
	body := ""
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
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSSID) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("activeDirectory.servers"); value.Exists() {
		data.ActiveDirectoryServers = make([]WirelessSSIDActiveDirectoryServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDActiveDirectoryServers{}
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
			(*parent).ActiveDirectoryServers = append((*parent).ActiveDirectoryServers, data)
			return true
		})
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
	if value := res.Get("ldap.servers"); value.Exists() {
		data.LdapServers = make([]WirelessSSIDLdapServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDLdapServers{}
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
			(*parent).LdapServers = append((*parent).LdapServers, data)
			return true
		})
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
	if value := res.Get("namedVlans.tagging.byApTags"); value.Exists() {
		data.NamedVlansTaggingByApTags = make([]WirelessSSIDNamedVlansTaggingByApTags, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDNamedVlansTaggingByApTags{}
			if value := res.Get("vlanName"); value.Exists() {
				data.VlanName = types.StringValue(value.String())
			} else {
				data.VlanName = types.StringNull()
			}
			if value := res.Get("tags"); value.Exists() {
				data.Tags = helpers.GetStringList(value.Array())
			} else {
				data.Tags = types.ListNull(types.StringType)
			}
			(*parent).NamedVlansTaggingByApTags = append((*parent).NamedVlansTaggingByApTags, data)
			return true
		})
	}
	if value := res.Get("oauth.allowedDomains"); value.Exists() {
		data.OauthAllowedDomains = helpers.GetStringList(value.Array())
	} else {
		data.OauthAllowedDomains = types.ListNull(types.StringType)
	}
	if value := res.Get("speedBurst.enabled"); value.Exists() {
		data.SpeedBurstEnabled = types.BoolValue(value.Bool())
	} else {
		data.SpeedBurstEnabled = types.BoolNull()
	}
	if value := res.Get("apTagsAndVlanIds"); value.Exists() {
		data.ApTagsAndVlanIds = make([]WirelessSSIDApTagsAndVlanIds, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDApTagsAndVlanIds{}
			if value := res.Get("vlanId"); value.Exists() {
				data.VlanId = types.Int64Value(value.Int())
			} else {
				data.VlanId = types.Int64Null()
			}
			if value := res.Get("tags"); value.Exists() {
				data.Tags = helpers.GetStringList(value.Array())
			} else {
				data.Tags = types.ListNull(types.StringType)
			}
			(*parent).ApTagsAndVlanIds = append((*parent).ApTagsAndVlanIds, data)
			return true
		})
	}
	if value := res.Get("availabilityTags"); value.Exists() {
		data.AvailabilityTags = helpers.GetStringList(value.Array())
	} else {
		data.AvailabilityTags = types.ListNull(types.StringType)
	}
	if value := res.Get("radiusAccountingServers"); value.Exists() {
		data.RadiusAccountingServers = make([]WirelessSSIDRadiusAccountingServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDRadiusAccountingServers{}
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
			if value := res.Get("secret"); value.Exists() {
				data.Secret = types.StringValue(value.String())
			} else {
				data.Secret = types.StringNull()
			}
			(*parent).RadiusAccountingServers = append((*parent).RadiusAccountingServers, data)
			return true
		})
	}
	if value := res.Get("radiusServers"); value.Exists() {
		data.RadiusServers = make([]WirelessSSIDRadiusServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDRadiusServers{}
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
			if value := res.Get("secret"); value.Exists() {
				data.Secret = types.StringValue(value.String())
			} else {
				data.Secret = types.StringNull()
			}
			(*parent).RadiusServers = append((*parent).RadiusServers, data)
			return true
		})
	}
	if value := res.Get("splashGuestSponsorDomains"); value.Exists() {
		data.SplashGuestSponsorDomains = helpers.GetStringList(value.Array())
	} else {
		data.SplashGuestSponsorDomains = types.ListNull(types.StringType)
	}
	if value := res.Get("walledGardenRanges"); value.Exists() {
		data.WalledGardenRanges = helpers.GetStringList(value.Array())
	} else {
		data.WalledGardenRanges = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessSSID) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
		keys := [...]string{"host", "port"}
		keyValues := [...]string{data.ActiveDirectoryServers[i].Host.ValueString(), strconv.FormatInt(data.ActiveDirectoryServers[i].Port.ValueInt64(), 10)}

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
	for i := 0; i < len(data.LdapServers); i++ {
		keys := [...]string{"host", "port"}
		keyValues := [...]string{data.LdapServers[i].Host.ValueString(), strconv.FormatInt(data.LdapServers[i].Port.ValueInt64(), 10)}

		parent := &data
		data := (*parent).LdapServers[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ldap.servers").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing LdapServers[%d] = %+v",
				i,
				(*parent).LdapServers[i],
			))
			(*parent).LdapServers = slices.Delete((*parent).LdapServers, i, i+1)
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
			data.Tags = helpers.GetStringList(value.Array())
		} else {
			data.Tags = types.ListNull(types.StringType)
		}
		(*parent).NamedVlansTaggingByApTags[i] = data
	}
	if value := res.Get("oauth.allowedDomains"); value.Exists() && !data.OauthAllowedDomains.IsNull() {
		data.OauthAllowedDomains = helpers.GetStringList(value.Array())
	} else {
		data.OauthAllowedDomains = types.ListNull(types.StringType)
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
			data.Tags = helpers.GetStringList(value.Array())
		} else {
			data.Tags = types.ListNull(types.StringType)
		}
		(*parent).ApTagsAndVlanIds[i] = data
	}
	if value := res.Get("availabilityTags"); value.Exists() && !data.AvailabilityTags.IsNull() {
		data.AvailabilityTags = helpers.GetStringList(value.Array())
	} else {
		data.AvailabilityTags = types.ListNull(types.StringType)
	}
	for i := 0; i < len(data.RadiusAccountingServers); i++ {
		keys := [...]string{"caCertificate", "host", "port", "radsecEnabled", "secret"}
		keyValues := [...]string{data.RadiusAccountingServers[i].CaCertificate.ValueString(), data.RadiusAccountingServers[i].Host.ValueString(), strconv.FormatInt(data.RadiusAccountingServers[i].Port.ValueInt64(), 10), strconv.FormatBool(data.RadiusAccountingServers[i].RadsecEnabled.ValueBool()), data.RadiusAccountingServers[i].Secret.ValueString()}

		parent := &data
		data := (*parent).RadiusAccountingServers[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("radiusAccountingServers").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing RadiusAccountingServers[%d] = %+v",
				i,
				(*parent).RadiusAccountingServers[i],
			))
			(*parent).RadiusAccountingServers = slices.Delete((*parent).RadiusAccountingServers, i, i+1)
			i--

			continue
		}
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
		if value := res.Get("secret"); value.Exists() && !data.Secret.IsNull() {
			data.Secret = types.StringValue(value.String())
		} else {
			data.Secret = types.StringNull()
		}
		(*parent).RadiusAccountingServers[i] = data
	}
	for i := 0; i < len(data.RadiusServers); i++ {
		keys := [...]string{"caCertificate", "host", "openRoamingCertificateId", "port", "radsecEnabled", "secret"}
		keyValues := [...]string{data.RadiusServers[i].CaCertificate.ValueString(), data.RadiusServers[i].Host.ValueString(), strconv.FormatInt(data.RadiusServers[i].OpenRoamingCertificateId.ValueInt64(), 10), strconv.FormatInt(data.RadiusServers[i].Port.ValueInt64(), 10), strconv.FormatBool(data.RadiusServers[i].RadsecEnabled.ValueBool()), data.RadiusServers[i].Secret.ValueString()}

		parent := &data
		data := (*parent).RadiusServers[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("radiusServers").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing RadiusServers[%d] = %+v",
				i,
				(*parent).RadiusServers[i],
			))
			(*parent).RadiusServers = slices.Delete((*parent).RadiusServers, i, i+1)
			i--

			continue
		}
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
		if value := res.Get("secret"); value.Exists() && !data.Secret.IsNull() {
			data.Secret = types.StringValue(value.String())
		} else {
			data.Secret = types.StringNull()
		}
		(*parent).RadiusServers[i] = data
	}
	if value := res.Get("splashGuestSponsorDomains"); value.Exists() && !data.SplashGuestSponsorDomains.IsNull() {
		data.SplashGuestSponsorDomains = helpers.GetStringList(value.Array())
	} else {
		data.SplashGuestSponsorDomains = types.ListNull(types.StringType)
	}
	if value := res.Get("walledGardenRanges"); value.Exists() && !data.WalledGardenRanges.IsNull() {
		data.WalledGardenRanges = helpers.GetStringList(value.Array())
	} else {
		data.WalledGardenRanges = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial
