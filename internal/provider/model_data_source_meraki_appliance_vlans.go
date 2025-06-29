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

type DataSourceApplianceVLANs struct {
	NetworkId types.String                    `tfsdk:"network_id"`
	Items     []DataSourceApplianceVLANsItems `tfsdk:"items"`
}

type DataSourceApplianceVLANsItems struct {
	Id                     types.String                                          `tfsdk:"id"`
	ApplianceIp            types.String                                          `tfsdk:"appliance_ip"`
	Cidr                   types.String                                          `tfsdk:"cidr"`
	DhcpBootFilename       types.String                                          `tfsdk:"dhcp_boot_filename"`
	DhcpBootNextServer     types.String                                          `tfsdk:"dhcp_boot_next_server"`
	DhcpBootOptionsEnabled types.Bool                                            `tfsdk:"dhcp_boot_options_enabled"`
	DhcpHandling           types.String                                          `tfsdk:"dhcp_handling"`
	DhcpLeaseTime          types.String                                          `tfsdk:"dhcp_lease_time"`
	DnsNameservers         types.String                                          `tfsdk:"dns_nameservers"`
	GroupPolicyId          types.String                                          `tfsdk:"group_policy_id"`
	VlanId                 types.String                                          `tfsdk:"vlan_id"`
	Mask                   types.Int64                                           `tfsdk:"mask"`
	Name                   types.String                                          `tfsdk:"name"`
	Subnet                 types.String                                          `tfsdk:"subnet"`
	TemplateVlanType       types.String                                          `tfsdk:"template_vlan_type"`
	VpnNatSubnet           types.String                                          `tfsdk:"vpn_nat_subnet"`
	FixedIpAssignments     map[string]DataSourceApplianceVLANsFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	Ipv6Enabled            types.Bool                                            `tfsdk:"ipv6_enabled"`
	Ipv6PrefixAssignments  []DataSourceApplianceVLANsIpv6PrefixAssignments       `tfsdk:"ipv6_prefix_assignments"`
	MandatoryDhcpEnabled   types.Bool                                            `tfsdk:"mandatory_dhcp_enabled"`
	DhcpOptions            []DataSourceApplianceVLANsDhcpOptions                 `tfsdk:"dhcp_options"`
	DhcpRelayServerIps     types.List                                            `tfsdk:"dhcp_relay_server_ips"`
	ReservedIpRanges       []DataSourceApplianceVLANsReservedIpRanges            `tfsdk:"reserved_ip_ranges"`
}

type DataSourceApplianceVLANsFixedIpAssignments struct {
	Ip   types.String `tfsdk:"ip"`
	Name types.String `tfsdk:"name"`
}

type DataSourceApplianceVLANsIpv6PrefixAssignments struct {
	Autonomous         types.Bool   `tfsdk:"autonomous"`
	StaticApplianceIp6 types.String `tfsdk:"static_appliance_ip6"`
	StaticPrefix       types.String `tfsdk:"static_prefix"`
	OriginType         types.String `tfsdk:"origin_type"`
	OriginInterfaces   types.List   `tfsdk:"origin_interfaces"`
}

type DataSourceApplianceVLANsDhcpOptions struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type DataSourceApplianceVLANsReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceVLANs) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/vlans", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceVLANs) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceApplianceVLANsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceApplianceVLANsItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("applianceIp"); value.Exists() && value.Value() != nil {
			data.ApplianceIp = types.StringValue(value.String())
		} else {
			data.ApplianceIp = types.StringNull()
		}
		if value := res.Get("cidr"); value.Exists() && value.Value() != nil {
			data.Cidr = types.StringValue(value.String())
		} else {
			data.Cidr = types.StringNull()
		}
		if value := res.Get("dhcpBootFilename"); value.Exists() && value.Value() != nil {
			data.DhcpBootFilename = types.StringValue(value.String())
		} else {
			data.DhcpBootFilename = types.StringNull()
		}
		if value := res.Get("dhcpBootNextServer"); value.Exists() && value.Value() != nil {
			data.DhcpBootNextServer = types.StringValue(value.String())
		} else {
			data.DhcpBootNextServer = types.StringNull()
		}
		if value := res.Get("dhcpBootOptionsEnabled"); value.Exists() && value.Value() != nil {
			data.DhcpBootOptionsEnabled = types.BoolValue(value.Bool())
		} else {
			data.DhcpBootOptionsEnabled = types.BoolNull()
		}
		if value := res.Get("dhcpHandling"); value.Exists() && value.Value() != nil {
			data.DhcpHandling = types.StringValue(value.String())
		} else {
			data.DhcpHandling = types.StringNull()
		}
		if value := res.Get("dhcpLeaseTime"); value.Exists() && value.Value() != nil {
			data.DhcpLeaseTime = types.StringValue(value.String())
		} else {
			data.DhcpLeaseTime = types.StringNull()
		}
		if value := res.Get("dnsNameservers"); value.Exists() && value.Value() != nil {
			data.DnsNameservers = types.StringValue(value.String())
		} else {
			data.DnsNameservers = types.StringNull()
		}
		if value := res.Get("groupPolicyId"); value.Exists() && value.Value() != nil {
			data.GroupPolicyId = types.StringValue(value.String())
		} else {
			data.GroupPolicyId = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && value.Value() != nil {
			data.VlanId = types.StringValue(value.String())
		} else {
			data.VlanId = types.StringNull()
		}
		if value := res.Get("mask"); value.Exists() && value.Value() != nil {
			data.Mask = types.Int64Value(value.Int())
		} else {
			data.Mask = types.Int64Null()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("subnet"); value.Exists() && value.Value() != nil {
			data.Subnet = types.StringValue(value.String())
		} else {
			data.Subnet = types.StringNull()
		}
		if value := res.Get("templateVlanType"); value.Exists() && value.Value() != nil {
			data.TemplateVlanType = types.StringValue(value.String())
		} else {
			data.TemplateVlanType = types.StringNull()
		}
		if value := res.Get("vpnNatSubnet"); value.Exists() && value.Value() != nil {
			data.VpnNatSubnet = types.StringValue(value.String())
		} else {
			data.VpnNatSubnet = types.StringNull()
		}
		if value := res.Get("fixedIpAssignments"); value.Exists() && value.Value() != nil {
			data.FixedIpAssignments = make(map[string]DataSourceApplianceVLANsFixedIpAssignments)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceApplianceVLANsFixedIpAssignments{}
				if value := res.Get("ip"); value.Exists() && value.Value() != nil {
					data.Ip = types.StringValue(value.String())
				} else {
					data.Ip = types.StringNull()
				}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				(*parent).FixedIpAssignments[k.String()] = data
				return true
			})
		}
		if value := res.Get("ipv6.enabled"); value.Exists() && value.Value() != nil {
			data.Ipv6Enabled = types.BoolValue(value.Bool())
		} else {
			data.Ipv6Enabled = types.BoolNull()
		}
		if value := res.Get("ipv6.prefixAssignments"); value.Exists() && value.Value() != nil {
			data.Ipv6PrefixAssignments = make([]DataSourceApplianceVLANsIpv6PrefixAssignments, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceApplianceVLANsIpv6PrefixAssignments{}
				if value := res.Get("autonomous"); value.Exists() && value.Value() != nil {
					data.Autonomous = types.BoolValue(value.Bool())
				} else {
					data.Autonomous = types.BoolNull()
				}
				if value := res.Get("staticApplianceIp6"); value.Exists() && value.Value() != nil {
					data.StaticApplianceIp6 = types.StringValue(value.String())
				} else {
					data.StaticApplianceIp6 = types.StringNull()
				}
				if value := res.Get("staticPrefix"); value.Exists() && value.Value() != nil {
					data.StaticPrefix = types.StringValue(value.String())
				} else {
					data.StaticPrefix = types.StringNull()
				}
				if value := res.Get("origin.type"); value.Exists() && value.Value() != nil {
					data.OriginType = types.StringValue(value.String())
				} else {
					data.OriginType = types.StringNull()
				}
				if value := res.Get("origin.interfaces"); value.Exists() && value.Value() != nil {
					data.OriginInterfaces = helpers.GetStringList(value.Array())
				} else {
					data.OriginInterfaces = types.ListNull(types.StringType)
				}
				(*parent).Ipv6PrefixAssignments = append((*parent).Ipv6PrefixAssignments, data)
				return true
			})
		}
		if value := res.Get("mandatoryDhcp.enabled"); value.Exists() && value.Value() != nil {
			data.MandatoryDhcpEnabled = types.BoolValue(value.Bool())
		} else {
			data.MandatoryDhcpEnabled = types.BoolNull()
		}
		if value := res.Get("dhcpOptions"); value.Exists() && value.Value() != nil {
			data.DhcpOptions = make([]DataSourceApplianceVLANsDhcpOptions, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceApplianceVLANsDhcpOptions{}
				if value := res.Get("code"); value.Exists() && value.Value() != nil {
					data.Code = types.StringValue(value.String())
				} else {
					data.Code = types.StringNull()
				}
				if value := res.Get("type"); value.Exists() && value.Value() != nil {
					data.Type = types.StringValue(value.String())
				} else {
					data.Type = types.StringNull()
				}
				if value := res.Get("value"); value.Exists() && value.Value() != nil {
					data.Value = types.StringValue(value.String())
				} else {
					data.Value = types.StringNull()
				}
				(*parent).DhcpOptions = append((*parent).DhcpOptions, data)
				return true
			})
		}
		if value := res.Get("dhcpRelayServerIps"); value.Exists() && value.Value() != nil {
			data.DhcpRelayServerIps = helpers.GetStringList(value.Array())
		} else {
			data.DhcpRelayServerIps = types.ListNull(types.StringType)
		}
		if value := res.Get("reservedIpRanges"); value.Exists() && value.Value() != nil {
			data.ReservedIpRanges = make([]DataSourceApplianceVLANsReservedIpRanges, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceApplianceVLANsReservedIpRanges{}
				if value := res.Get("comment"); value.Exists() && value.Value() != nil {
					data.Comment = types.StringValue(value.String())
				} else {
					data.Comment = types.StringNull()
				}
				if value := res.Get("end"); value.Exists() && value.Value() != nil {
					data.End = types.StringValue(value.String())
				} else {
					data.End = types.StringNull()
				}
				if value := res.Get("start"); value.Exists() && value.Value() != nil {
					data.Start = types.StringValue(value.String())
				} else {
					data.Start = types.StringNull()
				}
				(*parent).ReservedIpRanges = append((*parent).ReservedIpRanges, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
