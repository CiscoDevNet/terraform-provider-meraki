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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSourceWirelessSSID - used only by the data source, always including every
// attribute (unlike model.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessSSID struct {
	Id                                                                  types.String                                      `tfsdk:"id"`
	NetworkId                                                           types.String                                      `tfsdk:"network_id"`
	Number                                                              types.String                                      `tfsdk:"number"`
	AdaptivePolicyGroupId                                               types.String                                      `tfsdk:"adaptive_policy_group_id"`
	AdultContentFilteringEnabled                                        types.Bool                                        `tfsdk:"adult_content_filtering_enabled"`
	AuthMode                                                            types.String                                      `tfsdk:"auth_mode"`
	AvailableOnAllAps                                                   types.Bool                                        `tfsdk:"available_on_all_aps"`
	BandSelection                                                       types.String                                      `tfsdk:"band_selection"`
	ConcentratorNetworkId                                               types.String                                      `tfsdk:"concentrator_network_id"`
	DefaultVlanId                                                       types.Int64                                       `tfsdk:"default_vlan_id"`
	DisassociateClientsOnVpnFailover                                    types.Bool                                        `tfsdk:"disassociate_clients_on_vpn_failover"`
	Enabled                                                             types.Bool                                        `tfsdk:"enabled"`
	EncryptionMode                                                      types.String                                      `tfsdk:"encryption_mode"`
	EnterpriseAdminAccess                                               types.String                                      `tfsdk:"enterprise_admin_access"`
	IpAssignmentMode                                                    types.String                                      `tfsdk:"ip_assignment_mode"`
	LanIsolationEnabled                                                 types.Bool                                        `tfsdk:"lan_isolation_enabled"`
	MandatoryDhcpEnabled                                                types.Bool                                        `tfsdk:"mandatory_dhcp_enabled"`
	MinBitrate                                                          types.Float64                                     `tfsdk:"min_bitrate"`
	Name                                                                types.String                                      `tfsdk:"name"`
	PerClientBandwidthLimitDown                                         types.Int64                                       `tfsdk:"per_client_bandwidth_limit_down"`
	PerClientBandwidthLimitUp                                           types.Int64                                       `tfsdk:"per_client_bandwidth_limit_up"`
	PerSsidBandwidthLimitDown                                           types.Int64                                       `tfsdk:"per_ssid_bandwidth_limit_down"`
	PerSsidBandwidthLimitUp                                             types.Int64                                       `tfsdk:"per_ssid_bandwidth_limit_up"`
	Psk                                                                 types.String                                      `tfsdk:"psk"`
	PskWo                                                               types.String                                      `tfsdk:"psk_wo"`
	PskWoVersion                                                        types.Int64                                       `tfsdk:"psk_wo_version"`
	RadiusAccountingEnabled                                             types.Bool                                        `tfsdk:"radius_accounting_enabled"`
	RadiusAccountingInterimInterval                                     types.Int64                                       `tfsdk:"radius_accounting_interim_interval"`
	RadiusAccountingStartDelay                                          types.Int64                                       `tfsdk:"radius_accounting_start_delay"`
	RadiusAttributeForGroupPolicies                                     types.String                                      `tfsdk:"radius_attribute_for_group_policies"`
	RadiusAuthenticationNasId                                           types.String                                      `tfsdk:"radius_authentication_nas_id"`
	RadiusCalledStationId                                               types.String                                      `tfsdk:"radius_called_station_id"`
	RadiusCoaEnabled                                                    types.Bool                                        `tfsdk:"radius_coa_enabled"`
	RadiusFailoverPolicy                                                types.String                                      `tfsdk:"radius_failover_policy"`
	RadiusFallbackEnabled                                               types.Bool                                        `tfsdk:"radius_fallback_enabled"`
	RadiusGuestVlanEnabled                                              types.Bool                                        `tfsdk:"radius_guest_vlan_enabled"`
	RadiusGuestVlanId                                                   types.Int64                                       `tfsdk:"radius_guest_vlan_id"`
	RadiusLoadBalancingPolicy                                           types.String                                      `tfsdk:"radius_load_balancing_policy"`
	RadiusOverride                                                      types.Bool                                        `tfsdk:"radius_override"`
	RadiusProxyEnabled                                                  types.Bool                                        `tfsdk:"radius_proxy_enabled"`
	RadiusServerAttemptsLimit                                           types.Int64                                       `tfsdk:"radius_server_attempts_limit"`
	RadiusServerTimeout                                                 types.Int64                                       `tfsdk:"radius_server_timeout"`
	RadiusTestingEnabled                                                types.Bool                                        `tfsdk:"radius_testing_enabled"`
	SecondaryConcentratorNetworkId                                      types.String                                      `tfsdk:"secondary_concentrator_network_id"`
	SplashPage                                                          types.String                                      `tfsdk:"splash_page"`
	UseVlanTagging                                                      types.Bool                                        `tfsdk:"use_vlan_tagging"`
	Visible                                                             types.Bool                                        `tfsdk:"visible"`
	VlanId                                                              types.Int64                                       `tfsdk:"vlan_id"`
	WalledGardenEnabled                                                 types.Bool                                        `tfsdk:"walled_garden_enabled"`
	WpaEncryptionMode                                                   types.String                                      `tfsdk:"wpa_encryption_mode"`
	ActiveDirectoryCredentialsLogonName                                 types.String                                      `tfsdk:"active_directory_credentials_logon_name"`
	ActiveDirectoryCredentialsPassword                                  types.String                                      `tfsdk:"active_directory_credentials_password"`
	ActiveDirectoryCredentialsPasswordWo                                types.String                                      `tfsdk:"active_directory_credentials_password_wo"`
	ActiveDirectoryCredentialsPasswordWoVersion                         types.Int64                                       `tfsdk:"active_directory_credentials_password_wo_version"`
	ActiveDirectoryServers                                              []DataSourceWirelessSSIDActiveDirectoryServers    `tfsdk:"active_directory_servers"`
	DnsRewriteEnabled                                                   types.Bool                                        `tfsdk:"dns_rewrite_enabled"`
	DnsRewriteDnsCustomNameservers                                      types.List                                        `tfsdk:"dns_rewrite_dns_custom_nameservers"`
	Dot11rAdaptive                                                      types.Bool                                        `tfsdk:"dot11r_adaptive"`
	Dot11rEnabled                                                       types.Bool                                        `tfsdk:"dot11r_enabled"`
	Dot11wEnabled                                                       types.Bool                                        `tfsdk:"dot11w_enabled"`
	Dot11wRequired                                                      types.Bool                                        `tfsdk:"dot11w_required"`
	GreKey                                                              types.Int64                                       `tfsdk:"gre_key"`
	GreConcentratorHost                                                 types.String                                      `tfsdk:"gre_concentrator_host"`
	LdapBaseDistinguishedName                                           types.String                                      `tfsdk:"ldap_base_distinguished_name"`
	LdapCredentialsDistinguishedName                                    types.String                                      `tfsdk:"ldap_credentials_distinguished_name"`
	LdapCredentialsPassword                                             types.String                                      `tfsdk:"ldap_credentials_password"`
	LdapCredentialsPasswordWo                                           types.String                                      `tfsdk:"ldap_credentials_password_wo"`
	LdapCredentialsPasswordWoVersion                                    types.Int64                                       `tfsdk:"ldap_credentials_password_wo_version"`
	LdapServerCaCertificateContents                                     types.String                                      `tfsdk:"ldap_server_ca_certificate_contents"`
	LdapServers                                                         []DataSourceWirelessSSIDLdapServers               `tfsdk:"ldap_servers"`
	LocalAuthFallbackCacheTimeout                                       types.Int64                                       `tfsdk:"local_auth_fallback_cache_timeout"`
	LocalAuthFallbackEnabled                                            types.Bool                                        `tfsdk:"local_auth_fallback_enabled"`
	LocalAuthFallbackServerCaCertificateContents                        types.String                                      `tfsdk:"local_auth_fallback_server_ca_certificate_contents"`
	LocalRadiusCacheTimeout                                             types.Int64                                       `tfsdk:"local_radius_cache_timeout"`
	LocalRadiusCertificateAuthenticationEnabled                         types.Bool                                        `tfsdk:"local_radius_certificate_authentication_enabled"`
	LocalRadiusCertificateAuthenticationOcspResponderUrl                types.String                                      `tfsdk:"local_radius_certificate_authentication_ocsp_responder_url"`
	LocalRadiusCertificateAuthenticationUseLdap                         types.Bool                                        `tfsdk:"local_radius_certificate_authentication_use_ldap"`
	LocalRadiusCertificateAuthenticationUseOcsp                         types.Bool                                        `tfsdk:"local_radius_certificate_authentication_use_ocsp"`
	LocalRadiusCertificateAuthenticationClientRootCaCertificateContents types.String                                      `tfsdk:"local_radius_certificate_authentication_client_root_ca_certificate_contents"`
	LocalRadiusPasswordAuthenticationEnabled                            types.Bool                                        `tfsdk:"local_radius_password_authentication_enabled"`
	NamedVlansRadiusGuestVlanEnabled                                    types.Bool                                        `tfsdk:"named_vlans_radius_guest_vlan_enabled"`
	NamedVlansRadiusGuestVlanName                                       types.String                                      `tfsdk:"named_vlans_radius_guest_vlan_name"`
	NamedVlansTaggingDefaultVlanName                                    types.String                                      `tfsdk:"named_vlans_tagging_default_vlan_name"`
	NamedVlansTaggingEnabled                                            types.Bool                                        `tfsdk:"named_vlans_tagging_enabled"`
	NamedVlansTaggingByApTags                                           []DataSourceWirelessSSIDNamedVlansTaggingByApTags `tfsdk:"named_vlans_tagging_by_ap_tags"`
	OauthAllowedDomains                                                 types.Set                                         `tfsdk:"oauth_allowed_domains"`
	RadiusRadsecTlsTunnelTimeout                                        types.Int64                                       `tfsdk:"radius_radsec_tls_tunnel_timeout"`
	SpeedBurstEnabled                                                   types.Bool                                        `tfsdk:"speed_burst_enabled"`
	ApTagsAndVlanIds                                                    []DataSourceWirelessSSIDApTagsAndVlanIds          `tfsdk:"ap_tags_and_vlan_ids"`
	AvailabilityTags                                                    types.Set                                         `tfsdk:"availability_tags"`
	RadiusAccountingServers                                             []DataSourceWirelessSSIDRadiusAccountingServers   `tfsdk:"radius_accounting_servers"`
	RadiusServers                                                       []DataSourceWirelessSSIDRadiusServers             `tfsdk:"radius_servers"`
	SplashGuestSponsorDomains                                           types.Set                                         `tfsdk:"splash_guest_sponsor_domains"`
	WalledGardenRanges                                                  types.Set                                         `tfsdk:"walled_garden_ranges"`
	RadiusDasClientsIps                                                 types.Set                                         `tfsdk:"radius_das_clients_ips"`
	RadiusDasClientsSharedSecret                                        types.String                                      `tfsdk:"radius_das_clients_shared_secret"`
	RadiusDasClientsSharedSecretWo                                      types.String                                      `tfsdk:"radius_das_clients_shared_secret_wo"`
	RadiusDasClientsSharedSecretWoVersion                               types.Int64                                       `tfsdk:"radius_das_clients_shared_secret_wo_version"`
}

type DataSourceWirelessSSIDActiveDirectoryServers struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type DataSourceWirelessSSIDLdapServers struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type DataSourceWirelessSSIDNamedVlansTaggingByApTags struct {
	VlanName types.String `tfsdk:"vlan_name"`
	Tags     types.Set    `tfsdk:"tags"`
}

type DataSourceWirelessSSIDApTagsAndVlanIds struct {
	VlanId types.Int64 `tfsdk:"vlan_id"`
	Tags   types.Set   `tfsdk:"tags"`
}

type DataSourceWirelessSSIDRadiusAccountingServers struct {
	CaCertificate   types.String `tfsdk:"ca_certificate"`
	Host            types.String `tfsdk:"host"`
	Port            types.Int64  `tfsdk:"port"`
	RadsecEnabled   types.Bool   `tfsdk:"radsec_enabled"`
	Secret          types.String `tfsdk:"secret"`
	SecretWo        types.String `tfsdk:"secret_wo"`
	SecretWoVersion types.Int64  `tfsdk:"secret_wo_version"`
}

type DataSourceWirelessSSIDRadiusServers struct {
	CaCertificate            types.String `tfsdk:"ca_certificate"`
	Host                     types.String `tfsdk:"host"`
	OpenRoamingCertificateId types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
	RadsecEnabled            types.Bool   `tfsdk:"radsec_enabled"`
	Secret                   types.String `tfsdk:"secret"`
	SecretWo                 types.String `tfsdk:"secret_wo"`
	SecretWoVersion          types.Int64  `tfsdk:"secret_wo_version"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessSSID) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessSSID) fromBody(ctx context.Context, res meraki.Res) {
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
	if value := res.Get("radiusAccountingStartDelay"); value.Exists() && value.Value() != nil {
		data.RadiusAccountingStartDelay = types.Int64Value(value.Int())
	} else {
		data.RadiusAccountingStartDelay = types.Int64Null()
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
		data.ActiveDirectoryServers = make([]DataSourceWirelessSSIDActiveDirectoryServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDActiveDirectoryServers{}
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
		data.LdapServers = make([]DataSourceWirelessSSIDLdapServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDLdapServers{}
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
	if value := res.Get("localAuthFallback.cacheTimeout"); value.Exists() && value.Value() != nil {
		data.LocalAuthFallbackCacheTimeout = types.Int64Value(value.Int())
	} else {
		data.LocalAuthFallbackCacheTimeout = types.Int64Null()
	}
	if value := res.Get("localAuthFallback.enabled"); value.Exists() && value.Value() != nil {
		data.LocalAuthFallbackEnabled = types.BoolValue(value.Bool())
	} else {
		data.LocalAuthFallbackEnabled = types.BoolNull()
	}
	if value := res.Get("localAuthFallback.serverCaCertificate.contents"); value.Exists() && value.Value() != nil {
		data.LocalAuthFallbackServerCaCertificateContents = types.StringValue(value.String())
	} else {
		data.LocalAuthFallbackServerCaCertificateContents = types.StringNull()
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
		data.NamedVlansTaggingByApTags = make([]DataSourceWirelessSSIDNamedVlansTaggingByApTags, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDNamedVlansTaggingByApTags{}
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
		data.ApTagsAndVlanIds = make([]DataSourceWirelessSSIDApTagsAndVlanIds, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDApTagsAndVlanIds{}
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
		data.RadiusAccountingServers = make([]DataSourceWirelessSSIDRadiusAccountingServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDRadiusAccountingServers{}
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
		data.RadiusServers = make([]DataSourceWirelessSSIDRadiusServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDRadiusServers{}
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
}

// End of section. //template:end fromBody
