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

type SwitchLinkAggregation struct {
	Id                 types.String                              `tfsdk:"id"`
	NetworkId          types.String                              `tfsdk:"network_id"`
	SwitchPorts        []SwitchLinkAggregationSwitchPorts        `tfsdk:"switch_ports"`
	SwitchProfilePorts []SwitchLinkAggregationSwitchProfilePorts `tfsdk:"switch_profile_ports"`
}

type SwitchLinkAggregationSwitchPorts struct {
	PortId types.String `tfsdk:"port_id"`
	Serial types.String `tfsdk:"serial"`
}

type SwitchLinkAggregationSwitchProfilePorts struct {
	PortId  types.String `tfsdk:"port_id"`
	Profile types.String `tfsdk:"profile"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchLinkAggregation) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/linkAggregations", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchLinkAggregation) toBody(ctx context.Context, state SwitchLinkAggregation) string {
	body := ""
	if len(data.SwitchPorts) > 0 {
		body, _ = sjson.Set(body, "switchPorts", []interface{}{})
		for _, item := range data.SwitchPorts {
			itemBody := ""
			if !item.PortId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "portId", item.PortId.ValueString())
			}
			if !item.Serial.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serial", item.Serial.ValueString())
			}
			body, _ = sjson.SetRaw(body, "switchPorts.-1", itemBody)
		}
	}
	if len(data.SwitchProfilePorts) > 0 {
		body, _ = sjson.Set(body, "switchProfilePorts", []interface{}{})
		for _, item := range data.SwitchProfilePorts {
			itemBody := ""
			if !item.PortId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "portId", item.PortId.ValueString())
			}
			if !item.Profile.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "profile", item.Profile.ValueString())
			}
			body, _ = sjson.SetRaw(body, "switchProfilePorts.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchLinkAggregation) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("switchPorts"); value.Exists() && value.Value() != nil {
		data.SwitchPorts = make([]SwitchLinkAggregationSwitchPorts, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchLinkAggregationSwitchPorts{}
			if value := res.Get("portId"); value.Exists() && value.Value() != nil {
				data.PortId = types.StringValue(value.String())
			} else {
				data.PortId = types.StringNull()
			}
			if value := res.Get("serial"); value.Exists() && value.Value() != nil {
				data.Serial = types.StringValue(value.String())
			} else {
				data.Serial = types.StringNull()
			}
			(*parent).SwitchPorts = append((*parent).SwitchPorts, data)
			return true
		})
	}
	if value := res.Get("switchProfilePorts"); value.Exists() && value.Value() != nil {
		data.SwitchProfilePorts = make([]SwitchLinkAggregationSwitchProfilePorts, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchLinkAggregationSwitchProfilePorts{}
			if value := res.Get("portId"); value.Exists() && value.Value() != nil {
				data.PortId = types.StringValue(value.String())
			} else {
				data.PortId = types.StringNull()
			}
			if value := res.Get("profile"); value.Exists() && value.Value() != nil {
				data.Profile = types.StringValue(value.String())
			} else {
				data.Profile = types.StringNull()
			}
			(*parent).SwitchProfilePorts = append((*parent).SwitchProfilePorts, data)
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
func (data *SwitchLinkAggregation) fromBodyPartial(ctx context.Context, res gjson.Result) {
	for i := 0; i < len(data.SwitchPorts); i++ {
		keys := [...]string{"portId", "serial"}
		keyValues := [...]string{data.SwitchPorts[i].PortId.ValueString(), data.SwitchPorts[i].Serial.ValueString()}

		parent := &data
		data := (*parent).SwitchPorts[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("switchPorts").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing SwitchPorts[%d] = %+v",
				i,
				(*parent).SwitchPorts[i],
			))
			(*parent).SwitchPorts = slices.Delete((*parent).SwitchPorts, i, i+1)
			i--

			continue
		}
		if value := res.Get("portId"); value.Exists() && !data.PortId.IsNull() {
			data.PortId = types.StringValue(value.String())
		} else {
			data.PortId = types.StringNull()
		}
		if value := res.Get("serial"); value.Exists() && !data.Serial.IsNull() {
			data.Serial = types.StringValue(value.String())
		} else {
			data.Serial = types.StringNull()
		}
		(*parent).SwitchPorts[i] = data
	}
	for i := 0; i < len(data.SwitchProfilePorts); i++ {
		keys := [...]string{"portId", "profile"}
		keyValues := [...]string{data.SwitchProfilePorts[i].PortId.ValueString(), data.SwitchProfilePorts[i].Profile.ValueString()}

		parent := &data
		data := (*parent).SwitchProfilePorts[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("switchProfilePorts").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing SwitchProfilePorts[%d] = %+v",
				i,
				(*parent).SwitchProfilePorts[i],
			))
			(*parent).SwitchProfilePorts = slices.Delete((*parent).SwitchProfilePorts, i, i+1)
			i--

			continue
		}
		if value := res.Get("portId"); value.Exists() && !data.PortId.IsNull() {
			data.PortId = types.StringValue(value.String())
		} else {
			data.PortId = types.StringNull()
		}
		if value := res.Get("profile"); value.Exists() && !data.Profile.IsNull() {
			data.Profile = types.StringValue(value.String())
		} else {
			data.Profile = types.StringNull()
		}
		(*parent).SwitchProfilePorts[i] = data
	}
}

// End of section. //template:end fromBodyPartial
