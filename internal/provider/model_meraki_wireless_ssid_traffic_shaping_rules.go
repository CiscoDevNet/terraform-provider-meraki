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
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessSSIDTrafficShapingRules struct {
	Id                    types.String                           `tfsdk:"id"`
	NetworkId             types.String                           `tfsdk:"network_id"`
	Number                types.String                           `tfsdk:"number"`
	DefaultRulesEnabled   types.Bool                             `tfsdk:"default_rules_enabled"`
	TrafficShapingEnabled types.Bool                             `tfsdk:"traffic_shaping_enabled"`
	Rules                 []WirelessSSIDTrafficShapingRulesRules `tfsdk:"rules"`
}

type WirelessSSIDTrafficShapingRulesRules struct {
	DscpTagValue                                     types.Int64                                       `tfsdk:"dscp_tag_value"`
	PcpTagValue                                      types.Int64                                       `tfsdk:"pcp_tag_value"`
	PerClientBandwidthLimitsSettings                 types.String                                      `tfsdk:"per_client_bandwidth_limits_settings"`
	PerClientBandwidthLimitsBandwidthLimitsLimitDown types.Int64                                       `tfsdk:"per_client_bandwidth_limits_bandwidth_limits_limit_down"`
	PerClientBandwidthLimitsBandwidthLimitsLimitUp   types.Int64                                       `tfsdk:"per_client_bandwidth_limits_bandwidth_limits_limit_up"`
	Definitions                                      []WirelessSSIDTrafficShapingRulesRulesDefinitions `tfsdk:"definitions"`
}

type WirelessSSIDTrafficShapingRulesRulesDefinitions struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessSSIDTrafficShapingRules) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/trafficShaping/rules", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSSIDTrafficShapingRules) toBody(ctx context.Context, state WirelessSSIDTrafficShapingRules) string {
	body := ""
	if !data.DefaultRulesEnabled.IsNull() {
		body, _ = sjson.Set(body, "defaultRulesEnabled", data.DefaultRulesEnabled.ValueBool())
	}
	if !data.TrafficShapingEnabled.IsNull() {
		body, _ = sjson.Set(body, "trafficShapingEnabled", data.TrafficShapingEnabled.ValueBool())
	}
	if len(data.Rules) > 0 {
		body, _ = sjson.Set(body, "rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.DscpTagValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dscpTagValue", item.DscpTagValue.ValueInt64())
			}
			if !item.PcpTagValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "pcpTagValue", item.PcpTagValue.ValueInt64())
			}
			if !item.PerClientBandwidthLimitsSettings.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.settings", item.PerClientBandwidthLimitsSettings.ValueString())
			}
			if !item.PerClientBandwidthLimitsBandwidthLimitsLimitDown.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.bandwidthLimits.limitDown", item.PerClientBandwidthLimitsBandwidthLimitsLimitDown.ValueInt64())
			}
			if !item.PerClientBandwidthLimitsBandwidthLimitsLimitUp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.bandwidthLimits.limitUp", item.PerClientBandwidthLimitsBandwidthLimitsLimitUp.ValueInt64())
			}
			{
				itemBody, _ = sjson.Set(itemBody, "definitions", []interface{}{})
				for _, childItem := range item.Definitions {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "definitions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSSIDTrafficShapingRules) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("defaultRulesEnabled"); value.Exists() && value.Value() != nil {
		data.DefaultRulesEnabled = types.BoolValue(value.Bool())
	} else {
		data.DefaultRulesEnabled = types.BoolNull()
	}
	if value := res.Get("trafficShapingEnabled"); value.Exists() && value.Value() != nil {
		data.TrafficShapingEnabled = types.BoolValue(value.Bool())
	} else {
		data.TrafficShapingEnabled = types.BoolNull()
	}
	if value := res.Get("rules"); value.Exists() && value.Value() != nil {
		data.Rules = make([]WirelessSSIDTrafficShapingRulesRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDTrafficShapingRulesRules{}
			if value := res.Get("dscpTagValue"); value.Exists() && value.Value() != nil {
				data.DscpTagValue = types.Int64Value(value.Int())
			} else {
				data.DscpTagValue = types.Int64Null()
			}
			if value := res.Get("pcpTagValue"); value.Exists() && value.Value() != nil {
				data.PcpTagValue = types.Int64Value(value.Int())
			} else {
				data.PcpTagValue = types.Int64Null()
			}
			if value := res.Get("perClientBandwidthLimits.settings"); value.Exists() && value.Value() != nil {
				data.PerClientBandwidthLimitsSettings = types.StringValue(value.String())
			} else {
				data.PerClientBandwidthLimitsSettings = types.StringNull()
			}
			if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitDown"); value.Exists() && value.Value() != nil {
				data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Value(value.Int())
			} else {
				data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Null()
			}
			if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitUp"); value.Exists() && value.Value() != nil {
				data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Value(value.Int())
			} else {
				data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Null()
			}
			if value := res.Get("definitions"); value.Exists() && value.Value() != nil {
				data.Definitions = make([]WirelessSSIDTrafficShapingRulesRulesDefinitions, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := WirelessSSIDTrafficShapingRulesRulesDefinitions{}
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
					(*parent).Definitions = append((*parent).Definitions, data)
					return true
				})
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
func (data *WirelessSSIDTrafficShapingRules) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("defaultRulesEnabled"); value.Exists() && !data.DefaultRulesEnabled.IsNull() {
		data.DefaultRulesEnabled = types.BoolValue(value.Bool())
	} else {
		data.DefaultRulesEnabled = types.BoolNull()
	}
	if value := res.Get("trafficShapingEnabled"); value.Exists() && !data.TrafficShapingEnabled.IsNull() {
		data.TrafficShapingEnabled = types.BoolValue(value.Bool())
	} else {
		data.TrafficShapingEnabled = types.BoolNull()
	}
	{
		l := len(res.Get("rules").Array())
		tflog.Debug(ctx, fmt.Sprintf("rules array resizing from %d to %d", len(data.Rules), l))
		if len(data.Rules) > l {
			data.Rules = data.Rules[:l]
		}
	}
	for i := range data.Rules {
		parent := &data
		data := (*parent).Rules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("rules.%d", i))
		if value := res.Get("dscpTagValue"); value.Exists() && !data.DscpTagValue.IsNull() {
			data.DscpTagValue = types.Int64Value(value.Int())
		} else {
			data.DscpTagValue = types.Int64Null()
		}
		if value := res.Get("pcpTagValue"); value.Exists() && !data.PcpTagValue.IsNull() {
			data.PcpTagValue = types.Int64Value(value.Int())
		} else {
			data.PcpTagValue = types.Int64Null()
		}
		if value := res.Get("perClientBandwidthLimits.settings"); value.Exists() && !data.PerClientBandwidthLimitsSettings.IsNull() {
			data.PerClientBandwidthLimitsSettings = types.StringValue(value.String())
		} else {
			data.PerClientBandwidthLimitsSettings = types.StringNull()
		}
		if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitDown"); value.Exists() && !data.PerClientBandwidthLimitsBandwidthLimitsLimitDown.IsNull() {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Null()
		}
		if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitUp"); value.Exists() && !data.PerClientBandwidthLimitsBandwidthLimitsLimitUp.IsNull() {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Null()
		}
		for i := 0; i < len(data.Definitions); i++ {
			keys := [...]string{"type", "value"}
			keyValues := [...]string{data.Definitions[i].Type.ValueString(), data.Definitions[i].Value.ValueString()}

			parent := &data
			data := (*parent).Definitions[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("definitions").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Definitions[%d] = %+v",
					i,
					(*parent).Definitions[i],
				))
				(*parent).Definitions = slices.Delete((*parent).Definitions, i, i+1)
				i--

				continue
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
			(*parent).Definitions[i] = data
		}
		(*parent).Rules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessSSIDTrafficShapingRules) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns
