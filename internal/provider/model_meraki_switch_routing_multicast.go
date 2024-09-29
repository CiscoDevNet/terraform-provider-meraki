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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchRoutingMulticast struct {
	Id                                                 types.String                      `tfsdk:"id"`
	NetworkId                                          types.String                      `tfsdk:"network_id"`
	DefaultSettingsFloodUnknownMulticastTrafficEnabled types.Bool                        `tfsdk:"default_settings_flood_unknown_multicast_traffic_enabled"`
	DefaultSettingsIgmpSnoopingEnabled                 types.Bool                        `tfsdk:"default_settings_igmp_snooping_enabled"`
	Overrides                                          []SwitchRoutingMulticastOverrides `tfsdk:"overrides"`
}

type SwitchRoutingMulticastOverrides struct {
	FloodUnknownMulticastTrafficEnabled types.Bool `tfsdk:"flood_unknown_multicast_traffic_enabled"`
	IgmpSnoopingEnabled                 types.Bool `tfsdk:"igmp_snooping_enabled"`
	Stacks                              types.Set  `tfsdk:"stacks"`
	SwitchProfiles                      types.Set  `tfsdk:"switch_profiles"`
	Switches                            types.Set  `tfsdk:"switches"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchRoutingMulticast) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/routing/multicast", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchRoutingMulticast) toBody(ctx context.Context, state SwitchRoutingMulticast) string {
	body := ""
	if !data.DefaultSettingsFloodUnknownMulticastTrafficEnabled.IsNull() {
		body, _ = sjson.Set(body, "defaultSettings.floodUnknownMulticastTrafficEnabled", data.DefaultSettingsFloodUnknownMulticastTrafficEnabled.ValueBool())
	}
	if !data.DefaultSettingsIgmpSnoopingEnabled.IsNull() {
		body, _ = sjson.Set(body, "defaultSettings.igmpSnoopingEnabled", data.DefaultSettingsIgmpSnoopingEnabled.ValueBool())
	}
	if len(data.Overrides) > 0 {
		body, _ = sjson.Set(body, "overrides", []interface{}{})
		for _, item := range data.Overrides {
			itemBody := ""
			if !item.FloodUnknownMulticastTrafficEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "floodUnknownMulticastTrafficEnabled", item.FloodUnknownMulticastTrafficEnabled.ValueBool())
			}
			if !item.IgmpSnoopingEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "igmpSnoopingEnabled", item.IgmpSnoopingEnabled.ValueBool())
			}
			if !item.Stacks.IsNull() {
				var values []string
				item.Stacks.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "stacks", values)
			}
			if !item.SwitchProfiles.IsNull() {
				var values []string
				item.SwitchProfiles.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "switchProfiles", values)
			}
			if !item.Switches.IsNull() {
				var values []string
				item.Switches.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "switches", values)
			}
			body, _ = sjson.SetRaw(body, "overrides.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchRoutingMulticast) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("defaultSettings.floodUnknownMulticastTrafficEnabled"); value.Exists() && value.Value() != nil {
		data.DefaultSettingsFloodUnknownMulticastTrafficEnabled = types.BoolValue(value.Bool())
	} else {
		data.DefaultSettingsFloodUnknownMulticastTrafficEnabled = types.BoolNull()
	}
	if value := res.Get("defaultSettings.igmpSnoopingEnabled"); value.Exists() && value.Value() != nil {
		data.DefaultSettingsIgmpSnoopingEnabled = types.BoolValue(value.Bool())
	} else {
		data.DefaultSettingsIgmpSnoopingEnabled = types.BoolNull()
	}
	if value := res.Get("overrides"); value.Exists() && value.Value() != nil {
		data.Overrides = make([]SwitchRoutingMulticastOverrides, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchRoutingMulticastOverrides{}
			if value := res.Get("floodUnknownMulticastTrafficEnabled"); value.Exists() && value.Value() != nil {
				data.FloodUnknownMulticastTrafficEnabled = types.BoolValue(value.Bool())
			} else {
				data.FloodUnknownMulticastTrafficEnabled = types.BoolNull()
			}
			if value := res.Get("igmpSnoopingEnabled"); value.Exists() && value.Value() != nil {
				data.IgmpSnoopingEnabled = types.BoolValue(value.Bool())
			} else {
				data.IgmpSnoopingEnabled = types.BoolNull()
			}
			if value := res.Get("stacks"); value.Exists() && value.Value() != nil {
				data.Stacks = helpers.GetStringSet(value.Array())
			} else {
				data.Stacks = types.SetNull(types.StringType)
			}
			if value := res.Get("switchProfiles"); value.Exists() && value.Value() != nil {
				data.SwitchProfiles = helpers.GetStringSet(value.Array())
			} else {
				data.SwitchProfiles = types.SetNull(types.StringType)
			}
			if value := res.Get("switches"); value.Exists() && value.Value() != nil {
				data.Switches = helpers.GetStringSet(value.Array())
			} else {
				data.Switches = types.SetNull(types.StringType)
			}
			(*parent).Overrides = append((*parent).Overrides, data)
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
func (data *SwitchRoutingMulticast) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("defaultSettings.floodUnknownMulticastTrafficEnabled"); value.Exists() && !data.DefaultSettingsFloodUnknownMulticastTrafficEnabled.IsNull() {
		data.DefaultSettingsFloodUnknownMulticastTrafficEnabled = types.BoolValue(value.Bool())
	} else {
		data.DefaultSettingsFloodUnknownMulticastTrafficEnabled = types.BoolNull()
	}
	if value := res.Get("defaultSettings.igmpSnoopingEnabled"); value.Exists() && !data.DefaultSettingsIgmpSnoopingEnabled.IsNull() {
		data.DefaultSettingsIgmpSnoopingEnabled = types.BoolValue(value.Bool())
	} else {
		data.DefaultSettingsIgmpSnoopingEnabled = types.BoolNull()
	}
	for i := 0; i < len(data.Overrides); i++ {
		keys := [...]string{"floodUnknownMulticastTrafficEnabled", "igmpSnoopingEnabled"}
		keyValues := [...]string{strconv.FormatBool(data.Overrides[i].FloodUnknownMulticastTrafficEnabled.ValueBool()), strconv.FormatBool(data.Overrides[i].IgmpSnoopingEnabled.ValueBool())}

		parent := &data
		data := (*parent).Overrides[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("overrides").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Overrides[%d] = %+v",
				i,
				(*parent).Overrides[i],
			))
			(*parent).Overrides = slices.Delete((*parent).Overrides, i, i+1)
			i--

			continue
		}
		if value := res.Get("floodUnknownMulticastTrafficEnabled"); value.Exists() && !data.FloodUnknownMulticastTrafficEnabled.IsNull() {
			data.FloodUnknownMulticastTrafficEnabled = types.BoolValue(value.Bool())
		} else {
			data.FloodUnknownMulticastTrafficEnabled = types.BoolNull()
		}
		if value := res.Get("igmpSnoopingEnabled"); value.Exists() && !data.IgmpSnoopingEnabled.IsNull() {
			data.IgmpSnoopingEnabled = types.BoolValue(value.Bool())
		} else {
			data.IgmpSnoopingEnabled = types.BoolNull()
		}
		if value := res.Get("stacks"); value.Exists() && !data.Stacks.IsNull() {
			data.Stacks = helpers.GetStringSet(value.Array())
		} else {
			data.Stacks = types.SetNull(types.StringType)
		}
		if value := res.Get("switchProfiles"); value.Exists() && !data.SwitchProfiles.IsNull() {
			data.SwitchProfiles = helpers.GetStringSet(value.Array())
		} else {
			data.SwitchProfiles = types.SetNull(types.StringType)
		}
		if value := res.Get("switches"); value.Exists() && !data.Switches.IsNull() {
			data.Switches = helpers.GetStringSet(value.Array())
		} else {
			data.Switches = types.SetNull(types.StringType)
		}
		(*parent).Overrides[i] = data
	}
}

// End of section. //template:end fromBodyPartial
