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

type SwitchSettings struct {
	Id                             types.String                    `tfsdk:"id"`
	NetworkId                      types.String                    `tfsdk:"network_id"`
	UseCombinedPower               types.Bool                      `tfsdk:"use_combined_power"`
	Vlan                           types.Int64                     `tfsdk:"vlan"`
	MacBlocklistEnabled            types.Bool                      `tfsdk:"mac_blocklist_enabled"`
	UplinkClientSamplingEnabled    types.Bool                      `tfsdk:"uplink_client_sampling_enabled"`
	UplinkSelectionCandidates      types.String                    `tfsdk:"uplink_selection_candidates"`
	UplinkSelectionFailbackEnabled types.Bool                      `tfsdk:"uplink_selection_failback_enabled"`
	PowerExceptions                []SwitchSettingsPowerExceptions `tfsdk:"power_exceptions"`
}

type SwitchSettingsPowerExceptions struct {
	PowerType types.String `tfsdk:"power_type"`
	Serial    types.String `tfsdk:"serial"`
}

type SwitchSettingsIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/settings", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchSettings) toBody(ctx context.Context, state SwitchSettings) string {
	body := ""
	if !data.UseCombinedPower.IsNull() {
		body, _ = sjson.Set(body, "useCombinedPower", data.UseCombinedPower.ValueBool())
	}
	if !data.Vlan.IsNull() {
		body, _ = sjson.Set(body, "vlan", data.Vlan.ValueInt64())
	}
	if !data.MacBlocklistEnabled.IsNull() {
		body, _ = sjson.Set(body, "macBlocklist.enabled", data.MacBlocklistEnabled.ValueBool())
	}
	if !data.UplinkClientSamplingEnabled.IsNull() {
		body, _ = sjson.Set(body, "uplinkClientSampling.enabled", data.UplinkClientSamplingEnabled.ValueBool())
	}
	if !data.UplinkSelectionCandidates.IsNull() {
		body, _ = sjson.Set(body, "uplinkSelection.candidates", data.UplinkSelectionCandidates.ValueString())
	}
	if !data.UplinkSelectionFailbackEnabled.IsNull() {
		body, _ = sjson.Set(body, "uplinkSelection.failback.enabled", data.UplinkSelectionFailbackEnabled.ValueBool())
	}
	if data.PowerExceptions != nil {
		body, _ = sjson.Set(body, "powerExceptions", []interface{}{})
		for _, item := range data.PowerExceptions {
			itemBody := ""
			if !item.PowerType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "powerType", item.PowerType.ValueString())
			}
			if !item.Serial.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serial", item.Serial.ValueString())
			}
			body, _ = sjson.SetRaw(body, "powerExceptions.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchSettings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("useCombinedPower"); value.Exists() && value.Value() != nil {
		data.UseCombinedPower = types.BoolValue(value.Bool())
	} else {
		data.UseCombinedPower = types.BoolNull()
	}
	if value := res.Get("vlan"); value.Exists() && value.Value() != nil {
		data.Vlan = types.Int64Value(value.Int())
	} else {
		data.Vlan = types.Int64Null()
	}
	if value := res.Get("macBlocklist.enabled"); value.Exists() && value.Value() != nil {
		data.MacBlocklistEnabled = types.BoolValue(value.Bool())
	} else {
		data.MacBlocklistEnabled = types.BoolNull()
	}
	if value := res.Get("uplinkClientSampling.enabled"); value.Exists() && value.Value() != nil {
		data.UplinkClientSamplingEnabled = types.BoolValue(value.Bool())
	} else {
		data.UplinkClientSamplingEnabled = types.BoolNull()
	}
	if value := res.Get("uplinkSelection.candidates"); value.Exists() && value.Value() != nil {
		data.UplinkSelectionCandidates = types.StringValue(value.String())
	} else {
		data.UplinkSelectionCandidates = types.StringNull()
	}
	if value := res.Get("uplinkSelection.failback.enabled"); value.Exists() && value.Value() != nil {
		data.UplinkSelectionFailbackEnabled = types.BoolValue(value.Bool())
	} else {
		data.UplinkSelectionFailbackEnabled = types.BoolNull()
	}
	if value := res.Get("powerExceptions"); value.Exists() && value.Value() != nil {
		data.PowerExceptions = make([]SwitchSettingsPowerExceptions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchSettingsPowerExceptions{}
			if value := res.Get("powerType"); value.Exists() && value.Value() != nil {
				data.PowerType = types.StringValue(value.String())
			} else {
				data.PowerType = types.StringNull()
			}
			if value := res.Get("serial"); value.Exists() && value.Value() != nil {
				data.Serial = types.StringValue(value.String())
			} else {
				data.Serial = types.StringNull()
			}
			(*parent).PowerExceptions = append((*parent).PowerExceptions, data)
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
func (data *SwitchSettings) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("useCombinedPower"); value.Exists() && !data.UseCombinedPower.IsNull() {
		data.UseCombinedPower = types.BoolValue(value.Bool())
	} else {
		data.UseCombinedPower = types.BoolNull()
	}
	if value := res.Get("vlan"); value.Exists() && !data.Vlan.IsNull() {
		data.Vlan = types.Int64Value(value.Int())
	} else {
		data.Vlan = types.Int64Null()
	}
	if value := res.Get("macBlocklist.enabled"); value.Exists() && !data.MacBlocklistEnabled.IsNull() {
		data.MacBlocklistEnabled = types.BoolValue(value.Bool())
	} else {
		data.MacBlocklistEnabled = types.BoolNull()
	}
	if value := res.Get("uplinkClientSampling.enabled"); value.Exists() && !data.UplinkClientSamplingEnabled.IsNull() {
		data.UplinkClientSamplingEnabled = types.BoolValue(value.Bool())
	} else {
		data.UplinkClientSamplingEnabled = types.BoolNull()
	}
	if value := res.Get("uplinkSelection.candidates"); value.Exists() && !data.UplinkSelectionCandidates.IsNull() {
		data.UplinkSelectionCandidates = types.StringValue(value.String())
	} else {
		data.UplinkSelectionCandidates = types.StringNull()
	}
	if value := res.Get("uplinkSelection.failback.enabled"); value.Exists() && !data.UplinkSelectionFailbackEnabled.IsNull() {
		data.UplinkSelectionFailbackEnabled = types.BoolValue(value.Bool())
	} else {
		data.UplinkSelectionFailbackEnabled = types.BoolNull()
	}
	for i := 0; i < len(data.PowerExceptions); i++ {
		keys := [...]string{"serial"}
		keyValues := [...]string{data.PowerExceptions[i].Serial.ValueString()}

		parent := &data
		data := (*parent).PowerExceptions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("powerExceptions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing PowerExceptions[%d] = %+v",
				i,
				(*parent).PowerExceptions[i],
			))
			(*parent).PowerExceptions = slices.Delete((*parent).PowerExceptions, i, i+1)
			i--

			continue
		}
		if value := res.Get("powerType"); value.Exists() && !data.PowerType.IsNull() {
			data.PowerType = types.StringValue(value.String())
		} else {
			data.PowerType = types.StringNull()
		}
		if value := res.Get("serial"); value.Exists() && !data.Serial.IsNull() {
			data.Serial = types.StringValue(value.String())
		} else {
			data.Serial = types.StringNull()
		}
		(*parent).PowerExceptions[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SwitchSettings) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *SwitchSettingsIdentity) toIdentity(ctx context.Context, plan *SwitchSettings) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *SwitchSettings) fromIdentity(ctx context.Context, identity *SwitchSettingsIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data SwitchSettings) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
