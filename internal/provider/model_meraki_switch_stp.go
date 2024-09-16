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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchSTP struct {
	Id                types.String                 `tfsdk:"id"`
	NetworkId         types.String                 `tfsdk:"network_id"`
	RstpEnabled       types.Bool                   `tfsdk:"rstp_enabled"`
	StpBridgePriority []SwitchSTPStpBridgePriority `tfsdk:"stp_bridge_priority"`
}

type SwitchSTPStpBridgePriority struct {
	StpPriority    types.Int64 `tfsdk:"stp_priority"`
	Stacks         types.List  `tfsdk:"stacks"`
	SwitchProfiles types.List  `tfsdk:"switch_profiles"`
	Switches       types.List  `tfsdk:"switches"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchSTP) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/stp", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchSTP) toBody(ctx context.Context, state SwitchSTP) string {
	body := ""
	if !data.RstpEnabled.IsNull() {
		body, _ = sjson.Set(body, "rstpEnabled", data.RstpEnabled.ValueBool())
	}
	if len(data.StpBridgePriority) > 0 {
		body, _ = sjson.Set(body, "stpBridgePriority", []interface{}{})
		for _, item := range data.StpBridgePriority {
			itemBody := ""
			if !item.StpPriority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "stpPriority", item.StpPriority.ValueInt64())
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
			body, _ = sjson.SetRaw(body, "stpBridgePriority.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchSTP) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("rstpEnabled"); value.Exists() {
		data.RstpEnabled = types.BoolValue(value.Bool())
	} else {
		data.RstpEnabled = types.BoolNull()
	}
	if value := res.Get("stpBridgePriority"); value.Exists() {
		data.StpBridgePriority = make([]SwitchSTPStpBridgePriority, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchSTPStpBridgePriority{}
			if value := res.Get("stpPriority"); value.Exists() {
				data.StpPriority = types.Int64Value(value.Int())
			} else {
				data.StpPriority = types.Int64Null()
			}
			if value := res.Get("stacks"); value.Exists() {
				data.Stacks = helpers.GetStringList(value.Array())
			} else {
				data.Stacks = types.ListNull(types.StringType)
			}
			if value := res.Get("switchProfiles"); value.Exists() {
				data.SwitchProfiles = helpers.GetStringList(value.Array())
			} else {
				data.SwitchProfiles = types.ListNull(types.StringType)
			}
			if value := res.Get("switches"); value.Exists() {
				data.Switches = helpers.GetStringList(value.Array())
			} else {
				data.Switches = types.ListNull(types.StringType)
			}
			(*parent).StpBridgePriority = append((*parent).StpBridgePriority, data)
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
func (data *SwitchSTP) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("rstpEnabled"); value.Exists() && !data.RstpEnabled.IsNull() {
		data.RstpEnabled = types.BoolValue(value.Bool())
	} else {
		data.RstpEnabled = types.BoolNull()
	}
	{
		l := len(res.Get("stpBridgePriority").Array())
		tflog.Debug(ctx, fmt.Sprintf("stpBridgePriority array resizing from %d to %d", len(data.StpBridgePriority), l))
		if len(data.StpBridgePriority) > l {
			data.StpBridgePriority = data.StpBridgePriority[:l]
		}
	}
	for i := range data.StpBridgePriority {
		parent := &data
		data := (*parent).StpBridgePriority[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("stpBridgePriority.%d", i))
		if value := res.Get("stpPriority"); value.Exists() && !data.StpPriority.IsNull() {
			data.StpPriority = types.Int64Value(value.Int())
		} else {
			data.StpPriority = types.Int64Null()
		}
		if value := res.Get("stacks"); value.Exists() && !data.Stacks.IsNull() {
			data.Stacks = helpers.GetStringList(value.Array())
		} else {
			data.Stacks = types.ListNull(types.StringType)
		}
		if value := res.Get("switchProfiles"); value.Exists() && !data.SwitchProfiles.IsNull() {
			data.SwitchProfiles = helpers.GetStringList(value.Array())
		} else {
			data.SwitchProfiles = types.ListNull(types.StringType)
		}
		if value := res.Get("switches"); value.Exists() && !data.Switches.IsNull() {
			data.Switches = helpers.GetStringList(value.Array())
		} else {
			data.Switches = types.ListNull(types.StringType)
		}
		(*parent).StpBridgePriority[i] = data
	}
}

// End of section. //template:end fromBodyPartial
