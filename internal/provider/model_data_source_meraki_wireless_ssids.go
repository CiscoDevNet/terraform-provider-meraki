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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessSSIDs struct {
	NetworkId types.String                   `tfsdk:"network_id"`
	Items     []DataSourceWirelessSSIDsItems `tfsdk:"items"`
}

type DataSourceWirelessSSIDsItems struct {
	Id                                           types.String                                       `tfsdk:"id"`
	AdaptivePolicyGroupId                        types.String                                       `tfsdk:"adaptive_policy_group_id"`
	AdminSplashUrl                               types.String                                       `tfsdk:"admin_splash_url"`
	AdultContentFilteringEnabled                 types.Bool                                         `tfsdk:"adult_content_filtering_enabled"`
	AuthMode                                     types.String                                       `tfsdk:"auth_mode"`
	AvailableOnAllAps                            types.Bool                                         `tfsdk:"available_on_all_aps"`
	BandSelection                                types.String                                       `tfsdk:"band_selection"`
	DefaultVlanId                                types.Int64                                        `tfsdk:"default_vlan_id"`
	Enabled                                      types.Bool                                         `tfsdk:"enabled"`
	EncryptionMode                               types.String                                       `tfsdk:"encryption_mode"`
	EnterpriseAdminAccess                        types.String                                       `tfsdk:"enterprise_admin_access"`
	IpAssignmentMode                             types.String                                       `tfsdk:"ip_assignment_mode"`
	LanIsolationEnabled                          types.Bool                                         `tfsdk:"lan_isolation_enabled"`
	LocalAuth                                    types.Bool                                         `tfsdk:"local_auth"`
	MandatoryDhcpEnabled                         types.Bool                                         `tfsdk:"mandatory_dhcp_enabled"`
	MinBitrate                                   types.Float64                                      `tfsdk:"min_bitrate"`
	Name                                         types.String                                       `tfsdk:"name"`
	Number                                       types.Int64                                        `tfsdk:"number"`
	PerClientBandwidthLimitDown                  types.Int64                                        `tfsdk:"per_client_bandwidth_limit_down"`
	PerClientBandwidthLimitUp                    types.Int64                                        `tfsdk:"per_client_bandwidth_limit_up"`
	PerSsidBandwidthLimitDown                    types.Int64                                        `tfsdk:"per_ssid_bandwidth_limit_down"`
	PerSsidBandwidthLimitUp                      types.Int64                                        `tfsdk:"per_ssid_bandwidth_limit_up"`
	Psk                                          types.String                                       `tfsdk:"psk"`
	RadiusAccountingEnabled                      types.Bool                                         `tfsdk:"radius_accounting_enabled"`
	RadiusAttributeForGroupPolicies              types.String                                       `tfsdk:"radius_attribute_for_group_policies"`
	RadiusAuthenticationNasId                    types.String                                       `tfsdk:"radius_authentication_nas_id"`
	RadiusCalledStationId                        types.String                                       `tfsdk:"radius_called_station_id"`
	RadiusCoaEnabled                             types.Bool                                         `tfsdk:"radius_coa_enabled"`
	RadiusEnabled                                types.Bool                                         `tfsdk:"radius_enabled"`
	RadiusFailoverPolicy                         types.String                                       `tfsdk:"radius_failover_policy"`
	RadiusFallbackEnabled                        types.Bool                                         `tfsdk:"radius_fallback_enabled"`
	RadiusLoadBalancingPolicy                    types.String                                       `tfsdk:"radius_load_balancing_policy"`
	RadiusOverride                               types.Bool                                         `tfsdk:"radius_override"`
	RadiusProxyEnabled                           types.Bool                                         `tfsdk:"radius_proxy_enabled"`
	RadiusServerAttemptsLimit                    types.Int64                                        `tfsdk:"radius_server_attempts_limit"`
	RadiusServerTimeout                          types.Int64                                        `tfsdk:"radius_server_timeout"`
	RadiusTestingEnabled                         types.Bool                                         `tfsdk:"radius_testing_enabled"`
	SplashPage                                   types.String                                       `tfsdk:"splash_page"`
	SplashTimeout                                types.String                                       `tfsdk:"splash_timeout"`
	SsidAdminAccessible                          types.Bool                                         `tfsdk:"ssid_admin_accessible"`
	UseVlanTagging                               types.Bool                                         `tfsdk:"use_vlan_tagging"`
	Visible                                      types.Bool                                         `tfsdk:"visible"`
	WalledGardenEnabled                          types.Bool                                         `tfsdk:"walled_garden_enabled"`
	WifiPersonalNetworkEnabled                   types.Bool                                         `tfsdk:"wifi_personal_network_enabled"`
	WlanIdentifier                               types.Int64                                        `tfsdk:"wlan_identifier"`
	WpaEncryptionMode                            types.String                                       `tfsdk:"wpa_encryption_mode"`
	CampusGatewayClusterId                       types.String                                       `tfsdk:"campus_gateway_cluster_id"`
	DnsRewriteEnabled                            types.Bool                                         `tfsdk:"dns_rewrite_enabled"`
	DnsRewriteDnsCustomNameservers               types.List                                         `tfsdk:"dns_rewrite_dns_custom_nameservers"`
	Dot11kEnabled                                types.Bool                                         `tfsdk:"dot11k_enabled"`
	Dot11rAdaptive                               types.Bool                                         `tfsdk:"dot11r_adaptive"`
	Dot11rEnabled                                types.Bool                                         `tfsdk:"dot11r_enabled"`
	Dot11wEnabled                                types.Bool                                         `tfsdk:"dot11w_enabled"`
	Dot11wRequired                               types.Bool                                         `tfsdk:"dot11w_required"`
	GreClientIsolation                           types.Bool                                         `tfsdk:"gre_client_isolation"`
	GreKey                                       types.Int64                                        `tfsdk:"gre_key"`
	GreConcentratorHost                          types.String                                       `tfsdk:"gre_concentrator_host"`
	LocalAuthFallbackCacheTimeout                types.Int64                                        `tfsdk:"local_auth_fallback_cache_timeout"`
	LocalAuthFallbackEnabled                     types.Bool                                         `tfsdk:"local_auth_fallback_enabled"`
	LocalAuthFallbackServerCaCertificateContents types.String                                       `tfsdk:"local_auth_fallback_server_ca_certificate_contents"`
	NamedVlansRadiusGuestVlanEnabled             types.Bool                                         `tfsdk:"named_vlans_radius_guest_vlan_enabled"`
	NamedVlansRadiusGuestVlanName                types.String                                       `tfsdk:"named_vlans_radius_guest_vlan_name"`
	NamedVlansTaggingDefaultVlanName             types.String                                       `tfsdk:"named_vlans_tagging_default_vlan_name"`
	NamedVlansTaggingEnabled                     types.Bool                                         `tfsdk:"named_vlans_tagging_enabled"`
	NamedVlansTaggingByApTags                    []DataSourceWirelessSSIDsNamedVlansTaggingByApTags `tfsdk:"named_vlans_tagging_by_ap_tags"`
	OauthAllowedDomains                          types.List                                         `tfsdk:"oauth_allowed_domains"`
	SecurityEncryptionAkms                       types.List                                         `tfsdk:"security_encryption_akms"`
	SecurityEncryptionCiphers                    types.List                                         `tfsdk:"security_encryption_ciphers"`
	SpeedBurstEnabled                            types.Bool                                         `tfsdk:"speed_burst_enabled"`
	AccessControlClientsBlockedFromUsingLan      types.Bool                                         `tfsdk:"access_control_clients_blocked_from_using_lan"`
	AccessControlWiredClientsPartOfWifiNetwork   types.Bool                                         `tfsdk:"access_control_wired_clients_part_of_wifi_network"`
	AccessControlBandwidthLimit                  types.String                                       `tfsdk:"access_control_bandwidth_limit"`
	AccessControlClientIpAssignmentMode          types.String                                       `tfsdk:"access_control_client_ip_assignment_mode"`
	AccessControlEncryptionMode                  types.String                                       `tfsdk:"access_control_encryption_mode"`
	AccessControlSplashPageEnabled               types.Bool                                         `tfsdk:"access_control_splash_page_enabled"`
	AccessControlSplashPageTheme                 types.String                                       `tfsdk:"access_control_splash_page_theme"`
	AccessControlTunnelEnabled                   types.Bool                                         `tfsdk:"access_control_tunnel_enabled"`
	AccessControlTunnelSummary                   types.String                                       `tfsdk:"access_control_tunnel_summary"`
	AccessControlVlanEnabled                     types.Bool                                         `tfsdk:"access_control_vlan_enabled"`
	AccessControlVlanTag                         types.String                                       `tfsdk:"access_control_vlan_tag"`
	AvailabilityTags                             types.List                                         `tfsdk:"availability_tags"`
	RadiusAccountingServers                      []DataSourceWirelessSSIDsRadiusAccountingServers   `tfsdk:"radius_accounting_servers"`
	RadiusServers                                []DataSourceWirelessSSIDsRadiusServers             `tfsdk:"radius_servers"`
	WalledGardenRanges                           types.List                                         `tfsdk:"walled_garden_ranges"`
}

type DataSourceWirelessSSIDsNamedVlansTaggingByApTags struct {
	VlanName types.String `tfsdk:"vlan_name"`
	Tags     types.List   `tfsdk:"tags"`
}

type DataSourceWirelessSSIDsRadiusAccountingServers struct {
	CaCertificate            types.String `tfsdk:"ca_certificate"`
	Host                     types.String `tfsdk:"host"`
	OpenRoamingCertificateId types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
}

type DataSourceWirelessSSIDsRadiusServers struct {
	CaCertificate            types.String `tfsdk:"ca_certificate"`
	Host                     types.String `tfsdk:"host"`
	Id                       types.String `tfsdk:"id"`
	OpenRoamingCertificateId types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
	RadsecEnabled            types.Bool   `tfsdk:"radsec_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessSSIDs) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessSSIDs) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceWirelessSSIDsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceWirelessSSIDsItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("adaptivePolicyGroupId"); value.Exists() && value.Value() != nil {
			data.AdaptivePolicyGroupId = types.StringValue(value.String())
		} else {
			data.AdaptivePolicyGroupId = types.StringNull()
		}
		if value := res.Get("adminSplashUrl"); value.Exists() && value.Value() != nil {
			data.AdminSplashUrl = types.StringValue(value.String())
		} else {
			data.AdminSplashUrl = types.StringNull()
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
		if value := res.Get("defaultVlanId"); value.Exists() && value.Value() != nil {
			data.DefaultVlanId = types.Int64Value(value.Int())
		} else {
			data.DefaultVlanId = types.Int64Null()
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
		if value := res.Get("localAuth"); value.Exists() && value.Value() != nil {
			data.LocalAuth = types.BoolValue(value.Bool())
		} else {
			data.LocalAuth = types.BoolNull()
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
		if value := res.Get("number"); value.Exists() && value.Value() != nil {
			data.Number = types.Int64Value(value.Int())
		} else {
			data.Number = types.Int64Null()
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
		if value := res.Get("radiusEnabled"); value.Exists() && value.Value() != nil {
			data.RadiusEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusEnabled = types.BoolNull()
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
		if value := res.Get("splashPage"); value.Exists() && value.Value() != nil {
			data.SplashPage = types.StringValue(value.String())
		} else {
			data.SplashPage = types.StringNull()
		}
		if value := res.Get("splashTimeout"); value.Exists() && value.Value() != nil {
			data.SplashTimeout = types.StringValue(value.String())
		} else {
			data.SplashTimeout = types.StringNull()
		}
		if value := res.Get("ssidAdminAccessible"); value.Exists() && value.Value() != nil {
			data.SsidAdminAccessible = types.BoolValue(value.Bool())
		} else {
			data.SsidAdminAccessible = types.BoolNull()
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
		if value := res.Get("walledGardenEnabled"); value.Exists() && value.Value() != nil {
			data.WalledGardenEnabled = types.BoolValue(value.Bool())
		} else {
			data.WalledGardenEnabled = types.BoolNull()
		}
		if value := res.Get("wifiPersonalNetworkEnabled"); value.Exists() && value.Value() != nil {
			data.WifiPersonalNetworkEnabled = types.BoolValue(value.Bool())
		} else {
			data.WifiPersonalNetworkEnabled = types.BoolNull()
		}
		if value := res.Get("wlanIdentifier"); value.Exists() && value.Value() != nil {
			data.WlanIdentifier = types.Int64Value(value.Int())
		} else {
			data.WlanIdentifier = types.Int64Null()
		}
		if value := res.Get("wpaEncryptionMode"); value.Exists() && value.Value() != nil {
			data.WpaEncryptionMode = types.StringValue(value.String())
		} else {
			data.WpaEncryptionMode = types.StringNull()
		}
		if value := res.Get("campusGateway.cluster.id"); value.Exists() && value.Value() != nil {
			data.CampusGatewayClusterId = types.StringValue(value.String())
		} else {
			data.CampusGatewayClusterId = types.StringNull()
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
		if value := res.Get("dot11k.enabled"); value.Exists() && value.Value() != nil {
			data.Dot11kEnabled = types.BoolValue(value.Bool())
		} else {
			data.Dot11kEnabled = types.BoolNull()
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
		if value := res.Get("gre.clientIsolation"); value.Exists() && value.Value() != nil {
			data.GreClientIsolation = types.BoolValue(value.Bool())
		} else {
			data.GreClientIsolation = types.BoolNull()
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
			data.NamedVlansTaggingByApTags = make([]DataSourceWirelessSSIDsNamedVlansTaggingByApTags, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceWirelessSSIDsNamedVlansTaggingByApTags{}
				if value := res.Get("vlanName"); value.Exists() && value.Value() != nil {
					data.VlanName = types.StringValue(value.String())
				} else {
					data.VlanName = types.StringNull()
				}
				if value := res.Get("tags"); value.Exists() && value.Value() != nil {
					data.Tags = helpers.GetStringList(value.Array())
				} else {
					data.Tags = types.ListNull(types.StringType)
				}
				(*parent).NamedVlansTaggingByApTags = append((*parent).NamedVlansTaggingByApTags, data)
				return true
			})
		}
		if value := res.Get("oauth.allowedDomains"); value.Exists() && value.Value() != nil {
			data.OauthAllowedDomains = helpers.GetStringList(value.Array())
		} else {
			data.OauthAllowedDomains = types.ListNull(types.StringType)
		}
		if value := res.Get("security.encryption.akms"); value.Exists() && value.Value() != nil {
			data.SecurityEncryptionAkms = helpers.GetStringList(value.Array())
		} else {
			data.SecurityEncryptionAkms = types.ListNull(types.StringType)
		}
		if value := res.Get("security.encryption.ciphers"); value.Exists() && value.Value() != nil {
			data.SecurityEncryptionCiphers = helpers.GetStringList(value.Array())
		} else {
			data.SecurityEncryptionCiphers = types.ListNull(types.StringType)
		}
		if value := res.Get("speedBurst.enabled"); value.Exists() && value.Value() != nil {
			data.SpeedBurstEnabled = types.BoolValue(value.Bool())
		} else {
			data.SpeedBurstEnabled = types.BoolNull()
		}
		if value := res.Get("accessControl.clientsBlockedFromUsingLan"); value.Exists() && value.Value() != nil {
			data.AccessControlClientsBlockedFromUsingLan = types.BoolValue(value.Bool())
		} else {
			data.AccessControlClientsBlockedFromUsingLan = types.BoolNull()
		}
		if value := res.Get("accessControl.wiredClientsPartOfWifiNetwork"); value.Exists() && value.Value() != nil {
			data.AccessControlWiredClientsPartOfWifiNetwork = types.BoolValue(value.Bool())
		} else {
			data.AccessControlWiredClientsPartOfWifiNetwork = types.BoolNull()
		}
		if value := res.Get("accessControl.bandwidth.limit"); value.Exists() && value.Value() != nil {
			data.AccessControlBandwidthLimit = types.StringValue(value.String())
		} else {
			data.AccessControlBandwidthLimit = types.StringNull()
		}
		if value := res.Get("accessControl.clientIpAssignment.mode"); value.Exists() && value.Value() != nil {
			data.AccessControlClientIpAssignmentMode = types.StringValue(value.String())
		} else {
			data.AccessControlClientIpAssignmentMode = types.StringNull()
		}
		if value := res.Get("accessControl.encryption.mode"); value.Exists() && value.Value() != nil {
			data.AccessControlEncryptionMode = types.StringValue(value.String())
		} else {
			data.AccessControlEncryptionMode = types.StringNull()
		}
		if value := res.Get("accessControl.splashPage.enabled"); value.Exists() && value.Value() != nil {
			data.AccessControlSplashPageEnabled = types.BoolValue(value.Bool())
		} else {
			data.AccessControlSplashPageEnabled = types.BoolNull()
		}
		if value := res.Get("accessControl.splashPage.theme"); value.Exists() && value.Value() != nil {
			data.AccessControlSplashPageTheme = types.StringValue(value.String())
		} else {
			data.AccessControlSplashPageTheme = types.StringNull()
		}
		if value := res.Get("accessControl.tunnel.enabled"); value.Exists() && value.Value() != nil {
			data.AccessControlTunnelEnabled = types.BoolValue(value.Bool())
		} else {
			data.AccessControlTunnelEnabled = types.BoolNull()
		}
		if value := res.Get("accessControl.tunnel.summary"); value.Exists() && value.Value() != nil {
			data.AccessControlTunnelSummary = types.StringValue(value.String())
		} else {
			data.AccessControlTunnelSummary = types.StringNull()
		}
		if value := res.Get("accessControl.vlan.enabled"); value.Exists() && value.Value() != nil {
			data.AccessControlVlanEnabled = types.BoolValue(value.Bool())
		} else {
			data.AccessControlVlanEnabled = types.BoolNull()
		}
		if value := res.Get("accessControl.vlan.tag"); value.Exists() && value.Value() != nil {
			data.AccessControlVlanTag = types.StringValue(value.String())
		} else {
			data.AccessControlVlanTag = types.StringNull()
		}
		if value := res.Get("availabilityTags"); value.Exists() && value.Value() != nil {
			data.AvailabilityTags = helpers.GetStringList(value.Array())
		} else {
			data.AvailabilityTags = types.ListNull(types.StringType)
		}
		if value := res.Get("radiusAccountingServers"); value.Exists() && value.Value() != nil {
			data.RadiusAccountingServers = make([]DataSourceWirelessSSIDsRadiusAccountingServers, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceWirelessSSIDsRadiusAccountingServers{}
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
				(*parent).RadiusAccountingServers = append((*parent).RadiusAccountingServers, data)
				return true
			})
		}
		if value := res.Get("radiusServers"); value.Exists() && value.Value() != nil {
			data.RadiusServers = make([]DataSourceWirelessSSIDsRadiusServers, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceWirelessSSIDsRadiusServers{}
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
				if value := res.Get("id"); value.Exists() && value.Value() != nil {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
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
		if value := res.Get("walledGardenRanges"); value.Exists() && value.Value() != nil {
			data.WalledGardenRanges = helpers.GetStringList(value.Array())
		} else {
			data.WalledGardenRanges = types.ListNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
