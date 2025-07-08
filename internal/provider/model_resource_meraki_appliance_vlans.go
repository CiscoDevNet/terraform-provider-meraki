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
	"maps"
	"net/url"
	"slices"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceApplianceVLANs struct {
	Id             types.String                  `tfsdk:"id"`
	OrganizationId types.String                  `tfsdk:"organization_id"`
	NetworkId      types.String                  `tfsdk:"network_id"`
	Items          []ResourceApplianceVLANsItems `tfsdk:"items"`
}

type ResourceApplianceVLANsItems struct {
	Id                     types.String                                        `tfsdk:"id"`
	ApplianceIp            types.String                                        `tfsdk:"appliance_ip"`
	Cidr                   types.String                                        `tfsdk:"cidr"`
	DhcpBootFilename       types.String                                        `tfsdk:"dhcp_boot_filename"`
	DhcpBootNextServer     types.String                                        `tfsdk:"dhcp_boot_next_server"`
	DhcpBootOptionsEnabled types.Bool                                          `tfsdk:"dhcp_boot_options_enabled"`
	DhcpHandling           types.String                                        `tfsdk:"dhcp_handling"`
	DhcpLeaseTime          types.String                                        `tfsdk:"dhcp_lease_time"`
	DnsNameservers         types.String                                        `tfsdk:"dns_nameservers"`
	GroupPolicyId          types.String                                        `tfsdk:"group_policy_id"`
	VlanId                 types.String                                        `tfsdk:"vlan_id"`
	Mask                   types.Int64                                         `tfsdk:"mask"`
	Name                   types.String                                        `tfsdk:"name"`
	Subnet                 types.String                                        `tfsdk:"subnet"`
	TemplateVlanType       types.String                                        `tfsdk:"template_vlan_type"`
	VpnNatSubnet           types.String                                        `tfsdk:"vpn_nat_subnet"`
	FixedIpAssignments     map[string]ResourceApplianceVLANsFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	Ipv6Enabled            types.Bool                                          `tfsdk:"ipv6_enabled"`
	Ipv6PrefixAssignments  []ResourceApplianceVLANsIpv6PrefixAssignments       `tfsdk:"ipv6_prefix_assignments"`
	MandatoryDhcpEnabled   types.Bool                                          `tfsdk:"mandatory_dhcp_enabled"`
	DhcpOptions            []ResourceApplianceVLANsDhcpOptions                 `tfsdk:"dhcp_options"`
	DhcpRelayServerIps     types.List                                          `tfsdk:"dhcp_relay_server_ips"`
	ReservedIpRanges       []ResourceApplianceVLANsReservedIpRanges            `tfsdk:"reserved_ip_ranges"`
}

type ResourceApplianceVLANsFixedIpAssignments struct {
	Ip   types.String `tfsdk:"ip"`
	Name types.String `tfsdk:"name"`
}

type ResourceApplianceVLANsIpv6PrefixAssignments struct {
	Autonomous         types.Bool   `tfsdk:"autonomous"`
	StaticApplianceIp6 types.String `tfsdk:"static_appliance_ip6"`
	StaticPrefix       types.String `tfsdk:"static_prefix"`
	OriginType         types.String `tfsdk:"origin_type"`
	OriginInterfaces   types.List   `tfsdk:"origin_interfaces"`
}

type ResourceApplianceVLANsDhcpOptions struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResourceApplianceVLANsReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceApplianceVLANs) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/vlans", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceApplianceVLANsItems) toBody(ctx context.Context, state ResourceApplianceVLANsItems) string {
	body := ""
	if !data.ApplianceIp.IsNull() {
		body, _ = sjson.Set(body, "applianceIp", data.ApplianceIp.ValueString())
	}
	if !data.Cidr.IsNull() {
		body, _ = sjson.Set(body, "cidr", data.Cidr.ValueString())
	}
	if !data.DhcpBootFilename.IsNull() {
		body, _ = sjson.Set(body, "dhcpBootFilename", data.DhcpBootFilename.ValueString())
	}
	if !data.DhcpBootNextServer.IsNull() {
		body, _ = sjson.Set(body, "dhcpBootNextServer", data.DhcpBootNextServer.ValueString())
	}
	if !data.DhcpBootOptionsEnabled.IsNull() {
		body, _ = sjson.Set(body, "dhcpBootOptionsEnabled", data.DhcpBootOptionsEnabled.ValueBool())
	}
	if !data.DhcpHandling.IsNull() {
		body, _ = sjson.Set(body, "dhcpHandling", data.DhcpHandling.ValueString())
	}
	if !data.DhcpLeaseTime.IsNull() {
		body, _ = sjson.Set(body, "dhcpLeaseTime", data.DhcpLeaseTime.ValueString())
	}
	if !data.DnsNameservers.IsNull() {
		body, _ = sjson.Set(body, "dnsNameservers", data.DnsNameservers.ValueString())
	}
	if !data.GroupPolicyId.IsNull() {
		body, _ = sjson.Set(body, "groupPolicyId", data.GroupPolicyId.ValueString())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "id", data.VlanId.ValueString())
	}
	if !data.Mask.IsNull() {
		body, _ = sjson.Set(body, "mask", data.Mask.ValueInt64())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Subnet.IsNull() {
		body, _ = sjson.Set(body, "subnet", data.Subnet.ValueString())
	}
	if !data.TemplateVlanType.IsNull() {
		body, _ = sjson.Set(body, "templateVlanType", data.TemplateVlanType.ValueString())
	}
	if !data.VpnNatSubnet.IsNull() {
		body, _ = sjson.Set(body, "vpnNatSubnet", data.VpnNatSubnet.ValueString())
	}
	if len(data.FixedIpAssignments) > 0 {
		body, _ = sjson.Set(body, "fixedIpAssignments", map[string]interface{}{})
		for key, item := range data.FixedIpAssignments {
			itemBody := ""
			if !item.Ip.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ip", item.Ip.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "fixedIpAssignments."+key, itemBody)
		}
	}
	if !data.Ipv6Enabled.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enabled", data.Ipv6Enabled.ValueBool())
	}
	if len(data.Ipv6PrefixAssignments) > 0 {
		body, _ = sjson.Set(body, "ipv6.prefixAssignments", []interface{}{})
		for _, item := range data.Ipv6PrefixAssignments {
			itemBody := ""
			if !item.Autonomous.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "autonomous", item.Autonomous.ValueBool())
			}
			if !item.StaticApplianceIp6.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "staticApplianceIp6", item.StaticApplianceIp6.ValueString())
			}
			if !item.StaticPrefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "staticPrefix", item.StaticPrefix.ValueString())
			}
			if !item.OriginType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "origin.type", item.OriginType.ValueString())
			}
			if !item.OriginInterfaces.IsNull() {
				var values []string
				item.OriginInterfaces.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "origin.interfaces", values)
			}
			body, _ = sjson.SetRaw(body, "ipv6.prefixAssignments.-1", itemBody)
		}
	}
	if !data.MandatoryDhcpEnabled.IsNull() {
		body, _ = sjson.Set(body, "mandatoryDhcp.enabled", data.MandatoryDhcpEnabled.ValueBool())
	}
	if len(data.DhcpOptions) > 0 {
		body, _ = sjson.Set(body, "dhcpOptions", []interface{}{})
		for _, item := range data.DhcpOptions {
			itemBody := ""
			if !item.Code.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "code", item.Code.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Value.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dhcpOptions.-1", itemBody)
		}
	}
	if !data.DhcpRelayServerIps.IsNull() {
		var values []string
		data.DhcpRelayServerIps.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "dhcpRelayServerIps", values)
	}
	if len(data.ReservedIpRanges) > 0 {
		body, _ = sjson.Set(body, "reservedIpRanges", []interface{}{})
		for _, item := range data.ReservedIpRanges {
			itemBody := ""
			if !item.Comment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "comment", item.Comment.ValueString())
			}
			if !item.End.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "end", item.End.ValueString())
			}
			if !item.Start.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "start", item.Start.ValueString())
			}
			body, _ = sjson.SetRaw(body, "reservedIpRanges.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceApplianceVLANs) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceApplianceVLANsItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceApplianceVLANsItems{}
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
			data.FixedIpAssignments = make(map[string]ResourceApplianceVLANsFixedIpAssignments)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceApplianceVLANsFixedIpAssignments{}
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
			data.Ipv6PrefixAssignments = make([]ResourceApplianceVLANsIpv6PrefixAssignments, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceApplianceVLANsIpv6PrefixAssignments{}
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
			data.DhcpOptions = make([]ResourceApplianceVLANsDhcpOptions, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceApplianceVLANsDhcpOptions{}
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
			data.ReservedIpRanges = make([]ResourceApplianceVLANsReservedIpRanges, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceApplianceVLANsReservedIpRanges{}
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
	index := 0
	res.ForEach(func(k, res gjson.Result) bool {
		data.Items[index].Id = types.StringValue(res.Get("id").String())
		index++
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
func (data *ResourceApplianceVLANs) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
				if v.Get("id").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
		if value := res.Get("applianceIp"); value.Exists() && !data.ApplianceIp.IsNull() {
			data.ApplianceIp = types.StringValue(value.String())
		} else {
			data.ApplianceIp = types.StringNull()
		}
		if value := res.Get("cidr"); value.Exists() && !data.Cidr.IsNull() {
			data.Cidr = types.StringValue(value.String())
		} else {
			data.Cidr = types.StringNull()
		}
		if value := res.Get("dhcpBootFilename"); value.Exists() && !data.DhcpBootFilename.IsNull() {
			data.DhcpBootFilename = types.StringValue(value.String())
		} else {
			data.DhcpBootFilename = types.StringNull()
		}
		if value := res.Get("dhcpBootNextServer"); value.Exists() && !data.DhcpBootNextServer.IsNull() {
			data.DhcpBootNextServer = types.StringValue(value.String())
		} else {
			data.DhcpBootNextServer = types.StringNull()
		}
		if value := res.Get("dhcpBootOptionsEnabled"); value.Exists() && !data.DhcpBootOptionsEnabled.IsNull() {
			data.DhcpBootOptionsEnabled = types.BoolValue(value.Bool())
		} else {
			data.DhcpBootOptionsEnabled = types.BoolNull()
		}
		if value := res.Get("dhcpHandling"); value.Exists() && !data.DhcpHandling.IsNull() {
			data.DhcpHandling = types.StringValue(value.String())
		} else {
			data.DhcpHandling = types.StringNull()
		}
		if value := res.Get("dhcpLeaseTime"); value.Exists() && !data.DhcpLeaseTime.IsNull() {
			data.DhcpLeaseTime = types.StringValue(value.String())
		} else {
			data.DhcpLeaseTime = types.StringNull()
		}
		if value := res.Get("dnsNameservers"); value.Exists() && !data.DnsNameservers.IsNull() {
			data.DnsNameservers = types.StringValue(value.String())
		} else {
			data.DnsNameservers = types.StringNull()
		}
		if value := res.Get("groupPolicyId"); value.Exists() && !data.GroupPolicyId.IsNull() {
			data.GroupPolicyId = types.StringValue(value.String())
		} else {
			data.GroupPolicyId = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.VlanId.IsNull() {
			data.VlanId = types.StringValue(value.String())
		} else {
			data.VlanId = types.StringNull()
		}
		if value := res.Get("mask"); value.Exists() && !data.Mask.IsNull() {
			data.Mask = types.Int64Value(value.Int())
		} else {
			data.Mask = types.Int64Null()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("subnet"); value.Exists() && !data.Subnet.IsNull() {
			data.Subnet = types.StringValue(value.String())
		} else {
			data.Subnet = types.StringNull()
		}
		if value := res.Get("templateVlanType"); value.Exists() && !data.TemplateVlanType.IsNull() {
			data.TemplateVlanType = types.StringValue(value.String())
		} else {
			data.TemplateVlanType = types.StringNull()
		}
		if value := res.Get("vpnNatSubnet"); value.Exists() && !data.VpnNatSubnet.IsNull() {
			data.VpnNatSubnet = types.StringValue(value.String())
		} else {
			data.VpnNatSubnet = types.StringNull()
		}
		for i, item := range data.FixedIpAssignments {
			parent := &data
			data := item
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("fixedIpAssignments.%s", i))
			if value := res.Get("ip"); value.Exists() && !data.Ip.IsNull() {
				data.Ip = types.StringValue(value.String())
			} else {
				data.Ip = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).FixedIpAssignments[i] = data
		}
		if value := res.Get("ipv6.enabled"); value.Exists() && !data.Ipv6Enabled.IsNull() {
			data.Ipv6Enabled = types.BoolValue(value.Bool())
		} else {
			data.Ipv6Enabled = types.BoolNull()
		}
		{
			l := len(res.Get("ipv6.prefixAssignments").Array())
			tflog.Debug(ctx, fmt.Sprintf("ipv6.prefixAssignments array resizing from %d to %d", len(data.Ipv6PrefixAssignments), l))
			if len(data.Ipv6PrefixAssignments) > l {
				data.Ipv6PrefixAssignments = data.Ipv6PrefixAssignments[:l]
			}
		}
		for i := range data.Ipv6PrefixAssignments {
			parent := &data
			data := (*parent).Ipv6PrefixAssignments[i]
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("ipv6.prefixAssignments.%d", i))
			if value := res.Get("autonomous"); value.Exists() && !data.Autonomous.IsNull() {
				data.Autonomous = types.BoolValue(value.Bool())
			} else {
				data.Autonomous = types.BoolNull()
			}
			if value := res.Get("staticApplianceIp6"); value.Exists() && !data.StaticApplianceIp6.IsNull() {
				data.StaticApplianceIp6 = types.StringValue(value.String())
			} else {
				data.StaticApplianceIp6 = types.StringNull()
			}
			if value := res.Get("staticPrefix"); value.Exists() && !data.StaticPrefix.IsNull() {
				data.StaticPrefix = types.StringValue(value.String())
			} else {
				data.StaticPrefix = types.StringNull()
			}
			if value := res.Get("origin.type"); value.Exists() && !data.OriginType.IsNull() {
				data.OriginType = types.StringValue(value.String())
			} else {
				data.OriginType = types.StringNull()
			}
			if value := res.Get("origin.interfaces"); value.Exists() && !data.OriginInterfaces.IsNull() {
				data.OriginInterfaces = helpers.GetStringList(value.Array())
			} else {
				data.OriginInterfaces = types.ListNull(types.StringType)
			}
			(*parent).Ipv6PrefixAssignments[i] = data
		}
		if value := res.Get("mandatoryDhcp.enabled"); value.Exists() && !data.MandatoryDhcpEnabled.IsNull() {
			data.MandatoryDhcpEnabled = types.BoolValue(value.Bool())
		} else {
			data.MandatoryDhcpEnabled = types.BoolNull()
		}
		for i := 0; i < len(data.DhcpOptions); i++ {
			keys := [...]string{"code", "type", "value"}
			keyValues := [...]string{data.DhcpOptions[i].Code.ValueString(), data.DhcpOptions[i].Type.ValueString(), data.DhcpOptions[i].Value.ValueString()}

			parent := &data
			data := (*parent).DhcpOptions[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("dhcpOptions").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DhcpOptions[%d] = %+v",
					i,
					(*parent).DhcpOptions[i],
				))
				(*parent).DhcpOptions = slices.Delete((*parent).DhcpOptions, i, i+1)
				i--

				continue
			}
			if value := res.Get("code"); value.Exists() && !data.Code.IsNull() {
				data.Code = types.StringValue(value.String())
			} else {
				data.Code = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).DhcpOptions[i] = data
		}
		if value := res.Get("dhcpRelayServerIps"); value.Exists() && !data.DhcpRelayServerIps.IsNull() {
			data.DhcpRelayServerIps = helpers.GetStringList(value.Array())
		} else {
			data.DhcpRelayServerIps = types.ListNull(types.StringType)
		}
		for i := 0; i < len(data.ReservedIpRanges); i++ {
			keys := [...]string{"comment", "end", "start"}
			keyValues := [...]string{data.ReservedIpRanges[i].Comment.ValueString(), data.ReservedIpRanges[i].End.ValueString(), data.ReservedIpRanges[i].Start.ValueString()}

			parent := &data
			data := (*parent).ReservedIpRanges[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("reservedIpRanges").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing ReservedIpRanges[%d] = %+v",
					i,
					(*parent).ReservedIpRanges[i],
				))
				(*parent).ReservedIpRanges = slices.Delete((*parent).ReservedIpRanges, i, i+1)
				i--

				continue
			}
			if value := res.Get("comment"); value.Exists() && !data.Comment.IsNull() {
				data.Comment = types.StringValue(value.String())
			} else {
				data.Comment = types.StringNull()
			}
			if value := res.Get("end"); value.Exists() && !data.End.IsNull() {
				data.End = types.StringValue(value.String())
			} else {
				data.End = types.StringNull()
			}
			if value := res.Get("start"); value.Exists() && !data.Start.IsNull() {
				data.Start = types.StringValue(value.String())
			} else {
				data.Start = types.StringNull()
			}
			(*parent).ReservedIpRanges[i] = data
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceApplianceVLANs) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceApplianceVLANs) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceApplianceVLANs) hasChanges(ctx context.Context, state *ResourceApplianceVLANs, id string) bool {
	hasChanges := false

	item := ResourceApplianceVLANsItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceApplianceVLANsItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.ApplianceIp.Equal(stateItem.ApplianceIp) {
		hasChanges = true
	}
	if !item.Cidr.Equal(stateItem.Cidr) {
		hasChanges = true
	}
	if !item.DhcpBootFilename.Equal(stateItem.DhcpBootFilename) {
		hasChanges = true
	}
	if !item.DhcpBootNextServer.Equal(stateItem.DhcpBootNextServer) {
		hasChanges = true
	}
	if !item.DhcpBootOptionsEnabled.Equal(stateItem.DhcpBootOptionsEnabled) {
		hasChanges = true
	}
	if !item.DhcpHandling.Equal(stateItem.DhcpHandling) {
		hasChanges = true
	}
	if !item.DhcpLeaseTime.Equal(stateItem.DhcpLeaseTime) {
		hasChanges = true
	}
	if !item.DnsNameservers.Equal(stateItem.DnsNameservers) {
		hasChanges = true
	}
	if !item.GroupPolicyId.Equal(stateItem.GroupPolicyId) {
		hasChanges = true
	}
	if !item.VlanId.Equal(stateItem.VlanId) {
		hasChanges = true
	}
	if !item.Mask.Equal(stateItem.Mask) {
		hasChanges = true
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	if !item.Subnet.Equal(stateItem.Subnet) {
		hasChanges = true
	}
	if !item.TemplateVlanType.Equal(stateItem.TemplateVlanType) {
		hasChanges = true
	}
	if !item.VpnNatSubnet.Equal(stateItem.VpnNatSubnet) {
		hasChanges = true
	}
	if !maps.Equal(item.FixedIpAssignments, stateItem.FixedIpAssignments) {
		hasChanges = true
	}
	if !item.Ipv6Enabled.Equal(stateItem.Ipv6Enabled) {
		hasChanges = true
	}
	if len(item.Ipv6PrefixAssignments) != len(stateItem.Ipv6PrefixAssignments) {
		hasChanges = true
	} else {
		for i := range item.Ipv6PrefixAssignments {
			if !item.Ipv6PrefixAssignments[i].Autonomous.Equal(stateItem.Ipv6PrefixAssignments[i].Autonomous) {
				hasChanges = true
			}
			if !item.Ipv6PrefixAssignments[i].StaticApplianceIp6.Equal(stateItem.Ipv6PrefixAssignments[i].StaticApplianceIp6) {
				hasChanges = true
			}
			if !item.Ipv6PrefixAssignments[i].StaticPrefix.Equal(stateItem.Ipv6PrefixAssignments[i].StaticPrefix) {
				hasChanges = true
			}
			if !item.Ipv6PrefixAssignments[i].OriginType.Equal(stateItem.Ipv6PrefixAssignments[i].OriginType) {
				hasChanges = true
			}
			if !item.Ipv6PrefixAssignments[i].OriginInterfaces.Equal(stateItem.Ipv6PrefixAssignments[i].OriginInterfaces) {
				hasChanges = true
			}
		}
	}
	if !item.MandatoryDhcpEnabled.Equal(stateItem.MandatoryDhcpEnabled) {
		hasChanges = true
	}
	if len(item.DhcpOptions) != len(stateItem.DhcpOptions) {
		hasChanges = true
	} else {
		for i := range item.DhcpOptions {
			if !item.DhcpOptions[i].Code.Equal(stateItem.DhcpOptions[i].Code) {
				hasChanges = true
			}
			if !item.DhcpOptions[i].Type.Equal(stateItem.DhcpOptions[i].Type) {
				hasChanges = true
			}
			if !item.DhcpOptions[i].Value.Equal(stateItem.DhcpOptions[i].Value) {
				hasChanges = true
			}
		}
	}
	if !item.DhcpRelayServerIps.Equal(stateItem.DhcpRelayServerIps) {
		hasChanges = true
	}
	if len(item.ReservedIpRanges) != len(stateItem.ReservedIpRanges) {
		hasChanges = true
	} else {
		for i := range item.ReservedIpRanges {
			if !item.ReservedIpRanges[i].Comment.Equal(stateItem.ReservedIpRanges[i].Comment) {
				hasChanges = true
			}
			if !item.ReservedIpRanges[i].End.Equal(stateItem.ReservedIpRanges[i].End) {
				hasChanges = true
			}
			if !item.ReservedIpRanges[i].Start.Equal(stateItem.ReservedIpRanges[i].Start) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
