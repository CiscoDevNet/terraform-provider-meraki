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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessSSIDBonjourForwarding struct {
	Id               types.String                         `tfsdk:"id"`
	NetworkId        types.String                         `tfsdk:"network_id"`
	Number           types.String                         `tfsdk:"number"`
	Enabled          types.Bool                           `tfsdk:"enabled"`
	ExceptionEnabled types.Bool                           `tfsdk:"exception_enabled"`
	Rules            []WirelessSSIDBonjourForwardingRules `tfsdk:"rules"`
}

type WirelessSSIDBonjourForwardingRules struct {
	Description types.String `tfsdk:"description"`
	VlanId      types.String `tfsdk:"vlan_id"`
	Services    types.List   `tfsdk:"services"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessSSIDBonjourForwarding) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/bonjourForwarding", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSSIDBonjourForwarding) toBody(ctx context.Context, state WirelessSSIDBonjourForwarding) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.ExceptionEnabled.IsNull() {
		body, _ = sjson.Set(body, "exception.enabled", data.ExceptionEnabled.ValueBool())
	}
	if len(data.Rules) > 0 {
		body, _ = sjson.Set(body, "rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.VlanId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanId", item.VlanId.ValueString())
			}
			if !item.Services.IsNull() {
				var values []string
				item.Services.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "services", values)
			}
			body, _ = sjson.SetRaw(body, "rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSSIDBonjourForwarding) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("exception.enabled"); value.Exists() {
		data.ExceptionEnabled = types.BoolValue(value.Bool())
	} else {
		data.ExceptionEnabled = types.BoolNull()
	}
	if value := res.Get("rules"); value.Exists() {
		data.Rules = make([]WirelessSSIDBonjourForwardingRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDBonjourForwardingRules{}
			if value := res.Get("description"); value.Exists() {
				data.Description = types.StringValue(value.String())
			} else {
				data.Description = types.StringNull()
			}
			if value := res.Get("vlanId"); value.Exists() {
				data.VlanId = types.StringValue(value.String())
			} else {
				data.VlanId = types.StringNull()
			}
			if value := res.Get("services"); value.Exists() {
				data.Services = helpers.GetStringList(value.Array())
			} else {
				data.Services = types.ListNull(types.StringType)
			}
			(*parent).Rules = append((*parent).Rules, data)
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
func (data *WirelessSSIDBonjourForwarding) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("exception.enabled"); value.Exists() && !data.ExceptionEnabled.IsNull() {
		data.ExceptionEnabled = types.BoolValue(value.Bool())
	} else {
		data.ExceptionEnabled = types.BoolNull()
	}
	for i := 0; i < len(data.Rules); i++ {
		keys := [...]string{"vlanId"}
		keyValues := [...]string{data.Rules[i].VlanId.ValueString()}

		parent := &data
		data := (*parent).Rules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("rules").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Rules[%d] = %+v",
				i,
				(*parent).Rules[i],
			))
			(*parent).Rules = slices.Delete((*parent).Rules, i, i+1)
			i--

			continue
		}
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
			data.VlanId = types.StringValue(value.String())
		} else {
			data.VlanId = types.StringNull()
		}
		if value := res.Get("services"); value.Exists() && !data.Services.IsNull() {
			data.Services = helpers.GetStringList(value.Array())
		} else {
			data.Services = types.ListNull(types.StringType)
		}
		(*parent).Rules[i] = data
	}
}

// End of section. //template:end fromBodyPartial
