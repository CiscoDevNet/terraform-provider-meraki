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

type ApplianceTrafficShapingRules struct {
	Id                  types.String                        `tfsdk:"id"`
	NetworkId           types.String                        `tfsdk:"network_id"`
	DefaultRulesEnabled types.Bool                          `tfsdk:"default_rules_enabled"`
	Rules               []ApplianceTrafficShapingRulesRules `tfsdk:"rules"`
}

type ApplianceTrafficShapingRulesRules struct {
	DscpTagValue                    types.Int64                                    `tfsdk:"dscp_tag_value"`
	Priority                        types.String                                   `tfsdk:"priority"`
	PerClientBandwidthLimitSettings types.String                                   `tfsdk:"per_client_bandwidth_limit_settings"`
	PerClientBandwidthLimitDown     types.Int64                                    `tfsdk:"per_client_bandwidth_limit_down"`
	PerClientBandwidthLimitUp       types.Int64                                    `tfsdk:"per_client_bandwidth_limit_up"`
	Definitions                     []ApplianceTrafficShapingRulesRulesDefinitions `tfsdk:"definitions"`
}

type ApplianceTrafficShapingRulesRulesDefinitions struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceTrafficShapingRules) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/trafficShaping/rules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceTrafficShapingRules) toBody(ctx context.Context, state ApplianceTrafficShapingRules) string {
	body := ""
	if !data.DefaultRulesEnabled.IsNull() {
		body, _ = sjson.Set(body, "defaultRulesEnabled", data.DefaultRulesEnabled.ValueBool())
	}
	if len(data.Rules) > 0 {
		body, _ = sjson.Set(body, "rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.DscpTagValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dscpTagValue", item.DscpTagValue.ValueInt64())
			}
			if !item.Priority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "priority", item.Priority.ValueString())
			}
			if !item.PerClientBandwidthLimitSettings.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.settings", item.PerClientBandwidthLimitSettings.ValueString())
			}
			if !item.PerClientBandwidthLimitDown.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.bandwidthLimits.limitDown", item.PerClientBandwidthLimitDown.ValueInt64())
			}
			if !item.PerClientBandwidthLimitUp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.bandwidthLimits.limitUp", item.PerClientBandwidthLimitUp.ValueInt64())
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

func (data *ApplianceTrafficShapingRules) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("defaultRulesEnabled"); value.Exists() && value.Value() != nil {
		data.DefaultRulesEnabled = types.BoolValue(value.Bool())
	} else {
		data.DefaultRulesEnabled = types.BoolNull()
	}
	if value := res.Get("rules"); value.Exists() && value.Value() != nil {
		data.Rules = make([]ApplianceTrafficShapingRulesRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceTrafficShapingRulesRules{}
			if value := res.Get("dscpTagValue"); value.Exists() && value.Value() != nil {
				data.DscpTagValue = types.Int64Value(value.Int())
			} else {
				data.DscpTagValue = types.Int64Null()
			}
			if value := res.Get("priority"); value.Exists() && value.Value() != nil {
				data.Priority = types.StringValue(value.String())
			} else {
				data.Priority = types.StringNull()
			}
			if value := res.Get("perClientBandwidthLimits.settings"); value.Exists() && value.Value() != nil {
				data.PerClientBandwidthLimitSettings = types.StringValue(value.String())
			} else {
				data.PerClientBandwidthLimitSettings = types.StringNull()
			}
			if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitDown"); value.Exists() && value.Value() != nil {
				data.PerClientBandwidthLimitDown = types.Int64Value(value.Int())
			} else {
				data.PerClientBandwidthLimitDown = types.Int64Null()
			}
			if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitUp"); value.Exists() && value.Value() != nil {
				data.PerClientBandwidthLimitUp = types.Int64Value(value.Int())
			} else {
				data.PerClientBandwidthLimitUp = types.Int64Null()
			}
			if value := res.Get("definitions"); value.Exists() && value.Value() != nil {
				data.Definitions = make([]ApplianceTrafficShapingRulesRulesDefinitions, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ApplianceTrafficShapingRulesRulesDefinitions{}
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
func (data *ApplianceTrafficShapingRules) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("defaultRulesEnabled"); value.Exists() && !data.DefaultRulesEnabled.IsNull() {
		data.DefaultRulesEnabled = types.BoolValue(value.Bool())
	} else {
		data.DefaultRulesEnabled = types.BoolNull()
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
		if value := res.Get("priority"); value.Exists() && !data.Priority.IsNull() {
			data.Priority = types.StringValue(value.String())
		} else {
			data.Priority = types.StringNull()
		}
		if value := res.Get("perClientBandwidthLimits.settings"); value.Exists() && !data.PerClientBandwidthLimitSettings.IsNull() {
			data.PerClientBandwidthLimitSettings = types.StringValue(value.String())
		} else {
			data.PerClientBandwidthLimitSettings = types.StringNull()
		}
		if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitDown"); value.Exists() && !data.PerClientBandwidthLimitDown.IsNull() {
			data.PerClientBandwidthLimitDown = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitDown = types.Int64Null()
		}
		if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitUp"); value.Exists() && !data.PerClientBandwidthLimitUp.IsNull() {
			data.PerClientBandwidthLimitUp = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitUp = types.Int64Null()
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
func (data *ApplianceTrafficShapingRules) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns
