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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceSingleLAN struct {
	Id                    types.String                              `tfsdk:"id"`
	NetworkId             types.String                              `tfsdk:"network_id"`
	ApplianceIp           types.String                              `tfsdk:"appliance_ip"`
	Subnet                types.String                              `tfsdk:"subnet"`
	Ipv6Enabled           types.Bool                                `tfsdk:"ipv6_enabled"`
	Ipv6PrefixAssignments []ApplianceSingleLANIpv6PrefixAssignments `tfsdk:"ipv6_prefix_assignments"`
	MandatoryDhcpEnabled  types.Bool                                `tfsdk:"mandatory_dhcp_enabled"`
}

type ApplianceSingleLANIpv6PrefixAssignments struct {
	Autonomous         types.Bool   `tfsdk:"autonomous"`
	StaticApplianceIp6 types.String `tfsdk:"static_appliance_ip6"`
	StaticPrefix       types.String `tfsdk:"static_prefix"`
	OriginType         types.String `tfsdk:"origin_type"`
	OriginInterfaces   types.List   `tfsdk:"origin_interfaces"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceSingleLAN) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/singleLan", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceSingleLAN) toBody(ctx context.Context, state ApplianceSingleLAN) string {
	body := ""
	if !data.ApplianceIp.IsNull() {
		body, _ = sjson.Set(body, "applianceIp", data.ApplianceIp.ValueString())
	}
	if !data.Subnet.IsNull() {
		body, _ = sjson.Set(body, "subnet", data.Subnet.ValueString())
	}
	if !data.Ipv6Enabled.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enabled", data.Ipv6Enabled.ValueBool())
	}
	if data.Ipv6PrefixAssignments != nil {
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
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceSingleLAN) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("applianceIp"); value.Exists() && value.Value() != nil {
		data.ApplianceIp = types.StringValue(value.String())
	} else {
		data.ApplianceIp = types.StringNull()
	}
	if value := res.Get("subnet"); value.Exists() && value.Value() != nil {
		data.Subnet = types.StringValue(value.String())
	} else {
		data.Subnet = types.StringNull()
	}
	if value := res.Get("ipv6.enabled"); value.Exists() && value.Value() != nil {
		data.Ipv6Enabled = types.BoolValue(value.Bool())
	} else {
		data.Ipv6Enabled = types.BoolNull()
	}
	if value := res.Get("ipv6.prefixAssignments"); value.Exists() && value.Value() != nil {
		data.Ipv6PrefixAssignments = make([]ApplianceSingleLANIpv6PrefixAssignments, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceSingleLANIpv6PrefixAssignments{}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceSingleLAN) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("applianceIp"); value.Exists() && !data.ApplianceIp.IsNull() {
		data.ApplianceIp = types.StringValue(value.String())
	} else {
		data.ApplianceIp = types.StringNull()
	}
	if value := res.Get("subnet"); value.Exists() && !data.Subnet.IsNull() {
		data.Subnet = types.StringValue(value.String())
	} else {
		data.Subnet = types.StringNull()
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
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceSingleLAN) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceSingleLAN) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
