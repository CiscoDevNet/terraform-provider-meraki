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

type WirelessAlternateManagementInterface struct {
	Id           types.String                                       `tfsdk:"id"`
	NetworkId    types.String                                       `tfsdk:"network_id"`
	Enabled      types.Bool                                         `tfsdk:"enabled"`
	VlanId       types.Int64                                        `tfsdk:"vlan_id"`
	AccessPoints []WirelessAlternateManagementInterfaceAccessPoints `tfsdk:"access_points"`
	Protocols    types.Set                                          `tfsdk:"protocols"`
}

type WirelessAlternateManagementInterfaceAccessPoints struct {
	AlternateManagementIp types.String `tfsdk:"alternate_management_ip"`
	Dns1                  types.String `tfsdk:"dns1"`
	Dns2                  types.String `tfsdk:"dns2"`
	Gateway               types.String `tfsdk:"gateway"`
	Serial                types.String `tfsdk:"serial"`
	SubnetMask            types.String `tfsdk:"subnet_mask"`
}

type WirelessAlternateManagementInterfaceIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessAlternateManagementInterface) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/alternateManagementInterface", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessAlternateManagementInterface) toBody(ctx context.Context, state WirelessAlternateManagementInterface) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanId", data.VlanId.ValueInt64())
	}
	if data.AccessPoints != nil {
		body, _ = sjson.Set(body, "accessPoints", []interface{}{})
		for _, item := range data.AccessPoints {
			itemBody := ""
			if !item.AlternateManagementIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "alternateManagementIp", item.AlternateManagementIp.ValueString())
			}
			if !item.Dns1.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dns1", item.Dns1.ValueString())
			}
			if !item.Dns2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dns2", item.Dns2.ValueString())
			}
			if !item.Gateway.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "gateway", item.Gateway.ValueString())
			}
			if !item.Serial.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serial", item.Serial.ValueString())
			}
			if !item.SubnetMask.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "subnetMask", item.SubnetMask.ValueString())
			}
			body, _ = sjson.SetRaw(body, "accessPoints.-1", itemBody)
		}
	}
	if !data.Protocols.IsNull() {
		var values []string
		data.Protocols.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "protocols", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessAlternateManagementInterface) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("vlanId"); value.Exists() && value.Value() != nil {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("accessPoints"); value.Exists() && value.Value() != nil {
		data.AccessPoints = make([]WirelessAlternateManagementInterfaceAccessPoints, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessAlternateManagementInterfaceAccessPoints{}
			if value := res.Get("alternateManagementIp"); value.Exists() && value.Value() != nil {
				data.AlternateManagementIp = types.StringValue(value.String())
			} else {
				data.AlternateManagementIp = types.StringNull()
			}
			if value := res.Get("dns1"); value.Exists() && value.Value() != nil {
				data.Dns1 = types.StringValue(value.String())
			} else {
				data.Dns1 = types.StringNull()
			}
			if value := res.Get("dns2"); value.Exists() && value.Value() != nil {
				data.Dns2 = types.StringValue(value.String())
			} else {
				data.Dns2 = types.StringNull()
			}
			if value := res.Get("gateway"); value.Exists() && value.Value() != nil {
				data.Gateway = types.StringValue(value.String())
			} else {
				data.Gateway = types.StringNull()
			}
			if value := res.Get("serial"); value.Exists() && value.Value() != nil {
				data.Serial = types.StringValue(value.String())
			} else {
				data.Serial = types.StringNull()
			}
			if value := res.Get("subnetMask"); value.Exists() && value.Value() != nil {
				data.SubnetMask = types.StringValue(value.String())
			} else {
				data.SubnetMask = types.StringNull()
			}
			(*parent).AccessPoints = append((*parent).AccessPoints, data)
			return true
		})
	}
	if value := res.Get("protocols"); value.Exists() && value.Value() != nil {
		data.Protocols = helpers.GetStringSet(value.Array())
	} else {
		data.Protocols = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessAlternateManagementInterface) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	for i := 0; i < len(data.AccessPoints); i++ {
		keys := [...]string{"serial"}
		keyValues := [...]string{data.AccessPoints[i].Serial.ValueString()}

		parent := &data
		data := (*parent).AccessPoints[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("accessPoints").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AccessPoints[%d] = %+v",
				i,
				(*parent).AccessPoints[i],
			))
			(*parent).AccessPoints = slices.Delete((*parent).AccessPoints, i, i+1)
			i--

			continue
		}
		if value := res.Get("alternateManagementIp"); value.Exists() && !data.AlternateManagementIp.IsNull() {
			data.AlternateManagementIp = types.StringValue(value.String())
		} else {
			data.AlternateManagementIp = types.StringNull()
		}
		if value := res.Get("dns1"); value.Exists() && !data.Dns1.IsNull() {
			data.Dns1 = types.StringValue(value.String())
		} else {
			data.Dns1 = types.StringNull()
		}
		if value := res.Get("dns2"); value.Exists() && !data.Dns2.IsNull() {
			data.Dns2 = types.StringValue(value.String())
		} else {
			data.Dns2 = types.StringNull()
		}
		if value := res.Get("gateway"); value.Exists() && !data.Gateway.IsNull() {
			data.Gateway = types.StringValue(value.String())
		} else {
			data.Gateway = types.StringNull()
		}
		if value := res.Get("serial"); value.Exists() && !data.Serial.IsNull() {
			data.Serial = types.StringValue(value.String())
		} else {
			data.Serial = types.StringNull()
		}
		if value := res.Get("subnetMask"); value.Exists() && !data.SubnetMask.IsNull() {
			data.SubnetMask = types.StringValue(value.String())
		} else {
			data.SubnetMask = types.StringNull()
		}
		(*parent).AccessPoints[i] = data
	}
	if value := res.Get("protocols"); value.Exists() && !data.Protocols.IsNull() {
		data.Protocols = helpers.GetStringSet(value.Array())
	} else {
		data.Protocols = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessAlternateManagementInterface) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *WirelessAlternateManagementInterfaceIdentity) toIdentity(ctx context.Context, plan *WirelessAlternateManagementInterface) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *WirelessAlternateManagementInterface) fromIdentity(ctx context.Context, identity *WirelessAlternateManagementInterfaceIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data WirelessAlternateManagementInterface) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
