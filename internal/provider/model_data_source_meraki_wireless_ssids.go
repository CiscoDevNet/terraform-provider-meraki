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
	Id                                         types.String                                     `tfsdk:"id"`
	AdminSplashUrl                             types.String                                     `tfsdk:"admin_splash_url"`
	AuthMode                                   types.String                                     `tfsdk:"auth_mode"`
	AvailableOnAllAps                          types.Bool                                       `tfsdk:"available_on_all_aps"`
	BandSelection                              types.String                                     `tfsdk:"band_selection"`
	Enabled                                    types.Bool                                       `tfsdk:"enabled"`
	EncryptionMode                             types.String                                     `tfsdk:"encryption_mode"`
	IpAssignmentMode                           types.String                                     `tfsdk:"ip_assignment_mode"`
	LocalAuth                                  types.Bool                                       `tfsdk:"local_auth"`
	MandatoryDhcpEnabled                       types.Bool                                       `tfsdk:"mandatory_dhcp_enabled"`
	MinBitrate                                 types.Int64                                      `tfsdk:"min_bitrate"`
	Name                                       types.String                                     `tfsdk:"name"`
	Number                                     types.Int64                                      `tfsdk:"number"`
	PerClientBandwidthLimitDown                types.Int64                                      `tfsdk:"per_client_bandwidth_limit_down"`
	PerClientBandwidthLimitUp                  types.Int64                                      `tfsdk:"per_client_bandwidth_limit_up"`
	PerSsidBandwidthLimitDown                  types.Int64                                      `tfsdk:"per_ssid_bandwidth_limit_down"`
	PerSsidBandwidthLimitUp                    types.Int64                                      `tfsdk:"per_ssid_bandwidth_limit_up"`
	RadiusAccountingEnabled                    types.Bool                                       `tfsdk:"radius_accounting_enabled"`
	RadiusAttributeForGroupPolicies            types.String                                     `tfsdk:"radius_attribute_for_group_policies"`
	RadiusEnabled                              types.Bool                                       `tfsdk:"radius_enabled"`
	RadiusFailoverPolicy                       types.String                                     `tfsdk:"radius_failover_policy"`
	RadiusLoadBalancingPolicy                  types.String                                     `tfsdk:"radius_load_balancing_policy"`
	SplashPage                                 types.String                                     `tfsdk:"splash_page"`
	SplashTimeout                              types.String                                     `tfsdk:"splash_timeout"`
	SsidAdminAccessible                        types.Bool                                       `tfsdk:"ssid_admin_accessible"`
	Visible                                    types.Bool                                       `tfsdk:"visible"`
	WalledGardenEnabled                        types.Bool                                       `tfsdk:"walled_garden_enabled"`
	WpaEncryptionMode                          types.String                                     `tfsdk:"wpa_encryption_mode"`
	AccessControlClientsBlockedFromUsingLan    types.Bool                                       `tfsdk:"access_control_clients_blocked_from_using_lan"`
	AccessControlWiredClientsPartOfWifiNetwork types.Bool                                       `tfsdk:"access_control_wired_clients_part_of_wifi_network"`
	AccessControlBandwidthLimit                types.String                                     `tfsdk:"access_control_bandwidth_limit"`
	AccessControlClientIpAssignmentMode        types.String                                     `tfsdk:"access_control_client_ip_assignment_mode"`
	AccessControlEncryptionMode                types.String                                     `tfsdk:"access_control_encryption_mode"`
	AccessControlSplashPageEnabled             types.Bool                                       `tfsdk:"access_control_splash_page_enabled"`
	AccessControlSplashPageTheme               types.String                                     `tfsdk:"access_control_splash_page_theme"`
	AccessControlTunnelEnabled                 types.Bool                                       `tfsdk:"access_control_tunnel_enabled"`
	AccessControlTunnelSummary                 types.String                                     `tfsdk:"access_control_tunnel_summary"`
	AccessControlVlanEnabled                   types.Bool                                       `tfsdk:"access_control_vlan_enabled"`
	AccessControlVlanTag                       types.String                                     `tfsdk:"access_control_vlan_tag"`
	AvailabilityTags                           types.List                                       `tfsdk:"availability_tags"`
	RadiusAccountingServers                    []DataSourceWirelessSSIDsRadiusAccountingServers `tfsdk:"radius_accounting_servers"`
	RadiusServers                              []DataSourceWirelessSSIDsRadiusServers           `tfsdk:"radius_servers"`
	WalledGardenRanges                         types.List                                       `tfsdk:"walled_garden_ranges"`
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
	OpenRoamingCertificateId types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
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
		if value := res.Get("adminSplashUrl"); value.Exists() && value.Value() != nil {
			data.AdminSplashUrl = types.StringValue(value.String())
		} else {
			data.AdminSplashUrl = types.StringNull()
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
		if value := res.Get("ipAssignmentMode"); value.Exists() && value.Value() != nil {
			data.IpAssignmentMode = types.StringValue(value.String())
		} else {
			data.IpAssignmentMode = types.StringNull()
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
			data.MinBitrate = types.Int64Value(value.Int())
		} else {
			data.MinBitrate = types.Int64Null()
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
		if value := res.Get("radiusLoadBalancingPolicy"); value.Exists() && value.Value() != nil {
			data.RadiusLoadBalancingPolicy = types.StringValue(value.String())
		} else {
			data.RadiusLoadBalancingPolicy = types.StringNull()
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
		if value := res.Get("wpaEncryptionMode"); value.Exists() && value.Value() != nil {
			data.WpaEncryptionMode = types.StringValue(value.String())
		} else {
			data.WpaEncryptionMode = types.StringNull()
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
