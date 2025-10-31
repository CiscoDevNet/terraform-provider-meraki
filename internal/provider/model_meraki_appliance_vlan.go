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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceVLAN struct {
	Id                     types.String                         `tfsdk:"id"`
	NetworkId              types.String                         `tfsdk:"network_id"`
	ApplianceIp            types.String                         `tfsdk:"appliance_ip"`
	Cidr                   types.String                         `tfsdk:"cidr"`
	DhcpBootOptionsEnabled types.Bool                           `tfsdk:"dhcp_boot_options_enabled"`
	DhcpHandling           types.String                         `tfsdk:"dhcp_handling"`
	DhcpLeaseTime          types.String                         `tfsdk:"dhcp_lease_time"`
	GroupPolicyId          types.String                         `tfsdk:"group_policy_id"`
	VlanId                 types.String                         `tfsdk:"vlan_id"`
	Mask                   types.Int64                          `tfsdk:"mask"`
	Name                   types.String                         `tfsdk:"name"`
	Subnet                 types.String                         `tfsdk:"subnet"`
	TemplateVlanType       types.String                         `tfsdk:"template_vlan_type"`
	Ipv6Enabled            types.Bool                           `tfsdk:"ipv6_enabled"`
	Ipv6PrefixAssignments  []ApplianceVLANIpv6PrefixAssignments `tfsdk:"ipv6_prefix_assignments"`
	MandatoryDhcpEnabled   types.Bool                           `tfsdk:"mandatory_dhcp_enabled"`
	DhcpOptions            []ApplianceVLANDhcpOptions           `tfsdk:"dhcp_options"`
}

type ApplianceVLANIpv6PrefixAssignments struct {
	Autonomous         types.Bool   `tfsdk:"autonomous"`
	StaticApplianceIp6 types.String `tfsdk:"static_appliance_ip6"`
	StaticPrefix       types.String `tfsdk:"static_prefix"`
	OriginType         types.String `tfsdk:"origin_type"`
	OriginInterfaces   types.List   `tfsdk:"origin_interfaces"`
}

type ApplianceVLANDhcpOptions struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceVLAN) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/vlans", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceVLAN) toBody(ctx context.Context, state ApplianceVLAN) string {
	body := ""
	if !data.ApplianceIp.IsNull() {
		body, _ = sjson.Set(body, "applianceIp", data.ApplianceIp.ValueString())
	}
	if !data.Cidr.IsNull() {
		body, _ = sjson.Set(body, "cidr", data.Cidr.ValueString())
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
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceVLAN) fromBody(ctx context.Context, res meraki.Res) {
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
	if value := res.Get("ipv6.enabled"); value.Exists() && value.Value() != nil {
		data.Ipv6Enabled = types.BoolValue(value.Bool())
	} else {
		data.Ipv6Enabled = types.BoolNull()
	}
	if value := res.Get("ipv6.prefixAssignments"); value.Exists() && value.Value() != nil {
		data.Ipv6PrefixAssignments = make([]ApplianceVLANIpv6PrefixAssignments, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceVLANIpv6PrefixAssignments{}
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
		data.DhcpOptions = make([]ApplianceVLANDhcpOptions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceVLANDhcpOptions{}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceVLAN) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceVLAN) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceVLAN) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
