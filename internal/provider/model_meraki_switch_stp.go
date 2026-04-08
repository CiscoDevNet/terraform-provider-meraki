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
	"sort"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
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
	Stacks         types.Set   `tfsdk:"stacks"`
	SwitchProfiles types.Set   `tfsdk:"switch_profiles"`
	Switches       types.Set   `tfsdk:"switches"`
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
	if data.StpBridgePriority != nil {
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

func (data *SwitchSTP) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("rstpEnabled"); value.Exists() && value.Value() != nil {
		data.RstpEnabled = types.BoolValue(value.Bool())
	} else {
		data.RstpEnabled = types.BoolNull()
	}
	if value := res.Get("stpBridgePriority"); value.Exists() && value.Value() != nil {
		data.StpBridgePriority = make([]SwitchSTPStpBridgePriority, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchSTPStpBridgePriority{}
			if value := res.Get("stpPriority"); value.Exists() && value.Value() != nil {
				data.StpPriority = types.Int64Value(value.Int())
			} else {
				data.StpPriority = types.Int64Null()
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
			(*parent).StpBridgePriority = append((*parent).StpBridgePriority, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchSTP) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
	priorities := map[int64]tempPriority{}
	for i := range res.Get("stpBridgePriority").Array() {
		priorityRes := res.Get(fmt.Sprintf("stpBridgePriority.%d", i))
		priority := priorityRes.Get("stpPriority").Int()
		t := tempPriority{
			stacks:         []string{},
			switchProfiles: []string{},
			switches:       []string{},
		}
		if p, ok := priorities[priority]; ok {
			t = p
		}
		if value := priorityRes.Get("stacks"); value.Exists() {
			for _, r := range value.Array() {
				t.stacks = append(t.stacks, r.String())
			}
		}
		if value := priorityRes.Get("switchProfiles"); value.Exists() {
			for _, r := range value.Array() {
				t.switchProfiles = append(t.switchProfiles, r.String())
			}
		}
		if value := priorityRes.Get("switches"); value.Exists() {
			for _, r := range value.Array() {
				t.switches = append(t.switches, r.String())
			}
		}
		priorities[priority] = t
	}

	keys := make([]int64, 0, len(priorities))
	for k := range priorities {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] > keys[j] })

	newStpBridgePriority := []SwitchSTPStpBridgePriority{}
	for _, k := range keys {
		found := false
		for _, i := range data.StpBridgePriority {
			if i.StpPriority.ValueInt64() == k {
				found = true
			}
		}
		if !found {
			continue
		}
		newStacks := []attr.Value{}
		for _, s := range priorities[k].stacks {
			newStacks = append(newStacks, types.StringValue(s))
		}
		newSwitches := []attr.Value{}
		for _, s := range priorities[k].switches {
			newSwitches = append(newSwitches, types.StringValue(s))
		}
		newSwitchProfiles := []attr.Value{}
		for _, s := range priorities[k].switchProfiles {
			newSwitchProfiles = append(newSwitchProfiles, types.StringValue(s))
		}
		newStpBridgePriority = append(newStpBridgePriority, SwitchSTPStpBridgePriority{
			StpPriority:    types.Int64Value(k),
			Stacks:         setOrNull(types.SetValueMust(types.StringType, newStacks)),
			Switches:       setOrNull(types.SetValueMust(types.StringType, newSwitches)),
			SwitchProfiles: setOrNull(types.SetValueMust(types.StringType, newSwitchProfiles)),
		})
	}
	if len(newStpBridgePriority) > 0 {
		data.StpBridgePriority = newStpBridgePriority
	}

}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SwitchSTP) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data SwitchSTP) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

func setOrNull(s types.Set) types.Set {
	if len(s.Elements()) == 0 {
		return types.SetNull(types.StringType)
	}
	return s
}

type tempPriority struct {
	stacks         []string
	switchProfiles []string
	switches       []string
}
