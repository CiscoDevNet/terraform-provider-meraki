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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchAlternateManagementInterface struct {
	Id        types.String                                 `tfsdk:"id"`
	NetworkId types.String                                 `tfsdk:"network_id"`
	Enabled   types.Bool                                   `tfsdk:"enabled"`
	VlanId    types.Int64                                  `tfsdk:"vlan_id"`
	Protocols types.Set                                    `tfsdk:"protocols"`
	Switches  []SwitchAlternateManagementInterfaceSwitches `tfsdk:"switches"`
}

type SwitchAlternateManagementInterfaceSwitches struct {
	AlternateManagementIp types.String `tfsdk:"alternate_management_ip"`
	Gateway               types.String `tfsdk:"gateway"`
	Serial                types.String `tfsdk:"serial"`
	SubnetMask            types.String `tfsdk:"subnet_mask"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchAlternateManagementInterface) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/alternateManagementInterface", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchAlternateManagementInterface) toBody(ctx context.Context, state SwitchAlternateManagementInterface) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanId", data.VlanId.ValueInt64())
	}
	if !data.Protocols.IsNull() {
		var values []string
		data.Protocols.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "protocols", values)
	}
	if len(data.Switches) > 0 {
		body, _ = sjson.Set(body, "switches", []interface{}{})
		for _, item := range data.Switches {
			itemBody := ""
			if !item.AlternateManagementIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "alternateManagementIp", item.AlternateManagementIp.ValueString())
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
			body, _ = sjson.SetRaw(body, "switches.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchAlternateManagementInterface) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("switches"); value.Exists() && value.Value() != nil {
		data.Switches = make([]SwitchAlternateManagementInterfaceSwitches, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchAlternateManagementInterfaceSwitches{}
			if value := res.Get("alternateManagementIp"); value.Exists() && value.Value() != nil {
				data.AlternateManagementIp = types.StringValue(value.String())
			} else {
				data.AlternateManagementIp = types.StringNull()
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
			(*parent).Switches = append((*parent).Switches, data)
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
func (data *SwitchAlternateManagementInterface) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	for i := 0; i < len(data.Switches); i++ {
		keys := [...]string{"serial"}
		keyValues := [...]string{data.Switches[i].Serial.ValueString()}

		parent := &data
		data := (*parent).Switches[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("switches").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Switches[%d] = %+v",
				i,
				(*parent).Switches[i],
			))
			(*parent).Switches = slices.Delete((*parent).Switches, i, i+1)
			i--

			continue
		}
		if value := res.Get("alternateManagementIp"); value.Exists() && !data.AlternateManagementIp.IsNull() {
			data.AlternateManagementIp = types.StringValue(value.String())
		} else {
			data.AlternateManagementIp = types.StringNull()
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
		(*parent).Switches[i] = data
	}
}

// End of section. //template:end fromBodyPartial
