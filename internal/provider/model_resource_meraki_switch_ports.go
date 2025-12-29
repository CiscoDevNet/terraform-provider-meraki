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

type ResourceSwitchPorts struct {
	Id             types.String               `tfsdk:"id"`
	OrganizationId types.String               `tfsdk:"organization_id"`
	Serial         types.String               `tfsdk:"serial"`
	Items          []ResourceSwitchPortsItems `tfsdk:"items"`
}

type ResourceSwitchPortsItems struct {
	PortId                  types.String `tfsdk:"port_id"`
	AccessPolicyNumber      types.Int64  `tfsdk:"access_policy_number"`
	AccessPolicyType        types.String `tfsdk:"access_policy_type"`
	AdaptivePolicyGroupId   types.String `tfsdk:"adaptive_policy_group_id"`
	AllowedVlans            types.String `tfsdk:"allowed_vlans"`
	DaiTrusted              types.Bool   `tfsdk:"dai_trusted"`
	Enabled                 types.Bool   `tfsdk:"enabled"`
	FlexibleStackingEnabled types.Bool   `tfsdk:"flexible_stacking_enabled"`
	IsolationEnabled        types.Bool   `tfsdk:"isolation_enabled"`
	LinkNegotiation         types.String `tfsdk:"link_negotiation"`
	MacWhitelistLimit       types.Int64  `tfsdk:"mac_whitelist_limit"`
	Name                    types.String `tfsdk:"name"`
	PeerSgtCapable          types.Bool   `tfsdk:"peer_sgt_capable"`
	PoeEnabled              types.Bool   `tfsdk:"poe_enabled"`
	PortScheduleId          types.String `tfsdk:"port_schedule_id"`
	RstpEnabled             types.Bool   `tfsdk:"rstp_enabled"`
	StickyMacAllowListLimit types.Int64  `tfsdk:"sticky_mac_allow_list_limit"`
	StormControlEnabled     types.Bool   `tfsdk:"storm_control_enabled"`
	StpGuard                types.String `tfsdk:"stp_guard"`
	StpPortFastTrunk        types.Bool   `tfsdk:"stp_port_fast_trunk"`
	Type                    types.String `tfsdk:"type"`
	Udld                    types.String `tfsdk:"udld"`
	Vlan                    types.Int64  `tfsdk:"vlan"`
	VoiceVlan               types.Int64  `tfsdk:"voice_vlan"`
	Dot3azEnabled           types.Bool   `tfsdk:"dot3az_enabled"`
	HighSpeedEnabled        types.Bool   `tfsdk:"high_speed_enabled"`
	ProfileEnabled          types.Bool   `tfsdk:"profile_enabled"`
	ProfileId               types.String `tfsdk:"profile_id"`
	ProfileIname            types.String `tfsdk:"profile_iname"`
	MacAllowList            types.Set    `tfsdk:"mac_allow_list"`
	StickyMacAllowList      types.Set    `tfsdk:"sticky_mac_allow_list"`
	Tags                    types.Set    `tfsdk:"tags"`
}

type ResourceSwitchPortsIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Serial         types.String `tfsdk:"serial"`
	ItemIds        types.List   `tfsdk:"item_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceSwitchPorts) getPath() string {
	return fmt.Sprintf("/devices/%v/switch/ports", url.QueryEscape(data.Serial.ValueString()))
}

func (data ResourceSwitchPorts) getItemPath(id string) string {
	return fmt.Sprintf("/devices/%v/switch/ports/%v", url.QueryEscape(data.Serial.ValueString()), id)
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceSwitchPortsItems) toBody(ctx context.Context, state ResourceSwitchPortsItems) string {
	body := ""
	if !data.AccessPolicyNumber.IsNull() {
		body, _ = sjson.Set(body, "accessPolicyNumber", data.AccessPolicyNumber.ValueInt64())
	}
	if !data.AccessPolicyType.IsNull() {
		body, _ = sjson.Set(body, "accessPolicyType", data.AccessPolicyType.ValueString())
	}
	if !data.AdaptivePolicyGroupId.IsNull() {
		body, _ = sjson.Set(body, "adaptivePolicyGroupId", data.AdaptivePolicyGroupId.ValueString())
	}
	if !data.AllowedVlans.IsNull() {
		body, _ = sjson.Set(body, "allowedVlans", data.AllowedVlans.ValueString())
	}
	if !data.DaiTrusted.IsNull() {
		body, _ = sjson.Set(body, "daiTrusted", data.DaiTrusted.ValueBool())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.FlexibleStackingEnabled.IsNull() {
		body, _ = sjson.Set(body, "flexibleStackingEnabled", data.FlexibleStackingEnabled.ValueBool())
	}
	if !data.IsolationEnabled.IsNull() {
		body, _ = sjson.Set(body, "isolationEnabled", data.IsolationEnabled.ValueBool())
	}
	if !data.LinkNegotiation.IsNull() {
		body, _ = sjson.Set(body, "linkNegotiation", data.LinkNegotiation.ValueString())
	}
	if !data.MacWhitelistLimit.IsNull() {
		body, _ = sjson.Set(body, "macWhitelistLimit", data.MacWhitelistLimit.ValueInt64())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.PeerSgtCapable.IsNull() {
		body, _ = sjson.Set(body, "peerSgtCapable", data.PeerSgtCapable.ValueBool())
	}
	if !data.PoeEnabled.IsNull() {
		body, _ = sjson.Set(body, "poeEnabled", data.PoeEnabled.ValueBool())
	}
	if !data.PortScheduleId.IsNull() {
		body, _ = sjson.Set(body, "portScheduleId", data.PortScheduleId.ValueString())
	}
	if !data.RstpEnabled.IsNull() {
		body, _ = sjson.Set(body, "rstpEnabled", data.RstpEnabled.ValueBool())
	}
	if !data.StickyMacAllowListLimit.IsNull() {
		body, _ = sjson.Set(body, "stickyMacAllowListLimit", data.StickyMacAllowListLimit.ValueInt64())
	}
	if !data.StormControlEnabled.IsNull() {
		body, _ = sjson.Set(body, "stormControlEnabled", data.StormControlEnabled.ValueBool())
	}
	if !data.StpGuard.IsNull() {
		body, _ = sjson.Set(body, "stpGuard", data.StpGuard.ValueString())
	}
	if !data.StpPortFastTrunk.IsNull() {
		body, _ = sjson.Set(body, "stpPortFastTrunk", data.StpPortFastTrunk.ValueBool())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.Udld.IsNull() {
		body, _ = sjson.Set(body, "udld", data.Udld.ValueString())
	}
	if !data.Vlan.IsNull() {
		body, _ = sjson.Set(body, "vlan", data.Vlan.ValueInt64())
	}
	if !data.VoiceVlan.IsNull() {
		body, _ = sjson.Set(body, "voiceVlan", data.VoiceVlan.ValueInt64())
	}
	if !data.Dot3azEnabled.IsNull() {
		body, _ = sjson.Set(body, "dot3az.enabled", data.Dot3azEnabled.ValueBool())
	}
	if !data.HighSpeedEnabled.IsNull() {
		body, _ = sjson.Set(body, "highSpeed.enabled", data.HighSpeedEnabled.ValueBool())
	}
	if !data.ProfileEnabled.IsNull() {
		body, _ = sjson.Set(body, "profile.enabled", data.ProfileEnabled.ValueBool())
	}
	if !data.ProfileId.IsNull() {
		body, _ = sjson.Set(body, "profile.id", data.ProfileId.ValueString())
	}
	if !data.ProfileIname.IsNull() {
		body, _ = sjson.Set(body, "profile.iname", data.ProfileIname.ValueString())
	}
	if !data.MacAllowList.IsNull() {
		var values []string
		data.MacAllowList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "macAllowList", values)
	}
	if !data.StickyMacAllowList.IsNull() {
		var values []string
		data.StickyMacAllowList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "stickyMacAllowList", values)
	}
	if !data.Tags.IsNull() {
		var values []string
		data.Tags.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "tags", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceSwitchPorts) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceSwitchPortsItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceSwitchPortsItems{}
		if value := res.Get("portId"); value.Exists() && value.Value() != nil {
			data.PortId = types.StringValue(value.String())
		} else {
			data.PortId = types.StringNull()
		}
		if value := res.Get("accessPolicyNumber"); value.Exists() && value.Value() != nil {
			data.AccessPolicyNumber = types.Int64Value(value.Int())
		} else {
			data.AccessPolicyNumber = types.Int64Null()
		}
		if value := res.Get("accessPolicyType"); value.Exists() && value.Value() != nil {
			data.AccessPolicyType = types.StringValue(value.String())
		} else {
			data.AccessPolicyType = types.StringNull()
		}
		if value := res.Get("adaptivePolicyGroupId"); value.Exists() && value.Value() != nil {
			data.AdaptivePolicyGroupId = types.StringValue(value.String())
		} else {
			data.AdaptivePolicyGroupId = types.StringNull()
		}
		if value := res.Get("allowedVlans"); value.Exists() && value.Value() != nil {
			data.AllowedVlans = types.StringValue(value.String())
		} else {
			data.AllowedVlans = types.StringNull()
		}
		if value := res.Get("daiTrusted"); value.Exists() && value.Value() != nil {
			data.DaiTrusted = types.BoolValue(value.Bool())
		} else {
			data.DaiTrusted = types.BoolNull()
		}
		if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("flexibleStackingEnabled"); value.Exists() && value.Value() != nil {
			data.FlexibleStackingEnabled = types.BoolValue(value.Bool())
		} else {
			data.FlexibleStackingEnabled = types.BoolNull()
		}
		if value := res.Get("isolationEnabled"); value.Exists() && value.Value() != nil {
			data.IsolationEnabled = types.BoolValue(value.Bool())
		} else {
			data.IsolationEnabled = types.BoolNull()
		}
		if value := res.Get("linkNegotiation"); value.Exists() && value.Value() != nil {
			data.LinkNegotiation = types.StringValue(value.String())
		} else {
			data.LinkNegotiation = types.StringNull()
		}
		if value := res.Get("macWhitelistLimit"); value.Exists() && value.Value() != nil {
			data.MacWhitelistLimit = types.Int64Value(value.Int())
		} else {
			data.MacWhitelistLimit = types.Int64Null()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("peerSgtCapable"); value.Exists() && value.Value() != nil {
			data.PeerSgtCapable = types.BoolValue(value.Bool())
		} else {
			data.PeerSgtCapable = types.BoolNull()
		}
		if value := res.Get("poeEnabled"); value.Exists() && value.Value() != nil {
			data.PoeEnabled = types.BoolValue(value.Bool())
		} else {
			data.PoeEnabled = types.BoolNull()
		}
		if value := res.Get("portScheduleId"); value.Exists() && value.Value() != nil {
			data.PortScheduleId = types.StringValue(value.String())
		} else {
			data.PortScheduleId = types.StringNull()
		}
		if value := res.Get("rstpEnabled"); value.Exists() && value.Value() != nil {
			data.RstpEnabled = types.BoolValue(value.Bool())
		} else {
			data.RstpEnabled = types.BoolNull()
		}
		if value := res.Get("stickyMacAllowListLimit"); value.Exists() && value.Value() != nil {
			data.StickyMacAllowListLimit = types.Int64Value(value.Int())
		} else {
			data.StickyMacAllowListLimit = types.Int64Null()
		}
		if value := res.Get("stormControlEnabled"); value.Exists() && value.Value() != nil {
			data.StormControlEnabled = types.BoolValue(value.Bool())
		} else {
			data.StormControlEnabled = types.BoolNull()
		}
		if value := res.Get("stpGuard"); value.Exists() && value.Value() != nil {
			data.StpGuard = types.StringValue(value.String())
		} else {
			data.StpGuard = types.StringNull()
		}
		if value := res.Get("stpPortFastTrunk"); value.Exists() && value.Value() != nil {
			data.StpPortFastTrunk = types.BoolValue(value.Bool())
		} else {
			data.StpPortFastTrunk = types.BoolNull()
		}
		if value := res.Get("type"); value.Exists() && value.Value() != nil {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("udld"); value.Exists() && value.Value() != nil {
			data.Udld = types.StringValue(value.String())
		} else {
			data.Udld = types.StringNull()
		}
		if value := res.Get("vlan"); value.Exists() && value.Value() != nil {
			data.Vlan = types.Int64Value(value.Int())
		} else {
			data.Vlan = types.Int64Null()
		}
		if value := res.Get("voiceVlan"); value.Exists() && value.Value() != nil {
			data.VoiceVlan = types.Int64Value(value.Int())
		} else {
			data.VoiceVlan = types.Int64Null()
		}
		if value := res.Get("dot3az.enabled"); value.Exists() && value.Value() != nil {
			data.Dot3azEnabled = types.BoolValue(value.Bool())
		} else {
			data.Dot3azEnabled = types.BoolNull()
		}
		if value := res.Get("highSpeed.enabled"); value.Exists() && value.Value() != nil {
			data.HighSpeedEnabled = types.BoolValue(value.Bool())
		} else {
			data.HighSpeedEnabled = types.BoolNull()
		}
		if value := res.Get("profile.enabled"); value.Exists() && value.Value() != nil {
			data.ProfileEnabled = types.BoolValue(value.Bool())
		} else {
			data.ProfileEnabled = types.BoolNull()
		}
		if value := res.Get("profile.id"); value.Exists() && value.Value() != nil {
			data.ProfileId = types.StringValue(value.String())
		} else {
			data.ProfileId = types.StringNull()
		}
		if value := res.Get("profile.iname"); value.Exists() && value.Value() != nil {
			data.ProfileIname = types.StringValue(value.String())
		} else {
			data.ProfileIname = types.StringNull()
		}
		if value := res.Get("macAllowList"); value.Exists() && value.Value() != nil {
			data.MacAllowList = helpers.GetStringSet(value.Array())
		} else {
			data.MacAllowList = types.SetNull(types.StringType)
		}
		if value := res.Get("stickyMacAllowList"); value.Exists() && value.Value() != nil {
			data.StickyMacAllowList = helpers.GetStringSet(value.Array())
		} else {
			data.StickyMacAllowList = types.SetNull(types.StringType)
		}
		if value := res.Get("tags"); value.Exists() && value.Value() != nil {
			data.Tags = helpers.GetStringSet(value.Array())
		} else {
			data.Tags = types.SetNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	data.Id = data.OrganizationId
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceSwitchPorts) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("portId").String() == (*parent).Items[i].PortId.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		if value := res.Get("accessPolicyNumber"); value.Exists() && !data.AccessPolicyNumber.IsNull() {
			data.AccessPolicyNumber = types.Int64Value(value.Int())
		} else {
			data.AccessPolicyNumber = types.Int64Null()
		}
		if value := res.Get("accessPolicyType"); value.Exists() && !data.AccessPolicyType.IsNull() {
			data.AccessPolicyType = types.StringValue(value.String())
		} else {
			data.AccessPolicyType = types.StringNull()
		}
		if value := res.Get("adaptivePolicyGroupId"); value.Exists() && !data.AdaptivePolicyGroupId.IsNull() {
			data.AdaptivePolicyGroupId = types.StringValue(value.String())
		} else {
			data.AdaptivePolicyGroupId = types.StringNull()
		}
		if value := res.Get("allowedVlans"); value.Exists() && !data.AllowedVlans.IsNull() {
			data.AllowedVlans = types.StringValue(value.String())
		} else {
			data.AllowedVlans = types.StringNull()
		}
		if value := res.Get("daiTrusted"); value.Exists() && !data.DaiTrusted.IsNull() {
			data.DaiTrusted = types.BoolValue(value.Bool())
		} else {
			data.DaiTrusted = types.BoolNull()
		}
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("flexibleStackingEnabled"); value.Exists() && !data.FlexibleStackingEnabled.IsNull() {
			data.FlexibleStackingEnabled = types.BoolValue(value.Bool())
		} else {
			data.FlexibleStackingEnabled = types.BoolNull()
		}
		if value := res.Get("isolationEnabled"); value.Exists() && !data.IsolationEnabled.IsNull() {
			data.IsolationEnabled = types.BoolValue(value.Bool())
		} else {
			data.IsolationEnabled = types.BoolNull()
		}
		if value := res.Get("linkNegotiation"); value.Exists() && !data.LinkNegotiation.IsNull() {
			data.LinkNegotiation = types.StringValue(value.String())
		} else {
			data.LinkNegotiation = types.StringNull()
		}
		if value := res.Get("macWhitelistLimit"); value.Exists() && !data.MacWhitelistLimit.IsNull() {
			data.MacWhitelistLimit = types.Int64Value(value.Int())
		} else {
			data.MacWhitelistLimit = types.Int64Null()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("peerSgtCapable"); value.Exists() && !data.PeerSgtCapable.IsNull() {
			data.PeerSgtCapable = types.BoolValue(value.Bool())
		} else {
			data.PeerSgtCapable = types.BoolNull()
		}
		if value := res.Get("poeEnabled"); value.Exists() && !data.PoeEnabled.IsNull() {
			data.PoeEnabled = types.BoolValue(value.Bool())
		} else {
			data.PoeEnabled = types.BoolNull()
		}
		if value := res.Get("portScheduleId"); value.Exists() && !data.PortScheduleId.IsNull() {
			data.PortScheduleId = types.StringValue(value.String())
		} else {
			data.PortScheduleId = types.StringNull()
		}
		if value := res.Get("rstpEnabled"); value.Exists() && !data.RstpEnabled.IsNull() {
			data.RstpEnabled = types.BoolValue(value.Bool())
		} else {
			data.RstpEnabled = types.BoolNull()
		}
		if value := res.Get("stickyMacAllowListLimit"); value.Exists() && !data.StickyMacAllowListLimit.IsNull() {
			data.StickyMacAllowListLimit = types.Int64Value(value.Int())
		} else {
			data.StickyMacAllowListLimit = types.Int64Null()
		}
		if value := res.Get("stormControlEnabled"); value.Exists() && !data.StormControlEnabled.IsNull() {
			data.StormControlEnabled = types.BoolValue(value.Bool())
		} else {
			data.StormControlEnabled = types.BoolNull()
		}
		if value := res.Get("stpGuard"); value.Exists() && !data.StpGuard.IsNull() {
			data.StpGuard = types.StringValue(value.String())
		} else {
			data.StpGuard = types.StringNull()
		}
		if value := res.Get("stpPortFastTrunk"); value.Exists() && !data.StpPortFastTrunk.IsNull() {
			data.StpPortFastTrunk = types.BoolValue(value.Bool())
		} else {
			data.StpPortFastTrunk = types.BoolNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("udld"); value.Exists() && !data.Udld.IsNull() {
			data.Udld = types.StringValue(value.String())
		} else {
			data.Udld = types.StringNull()
		}
		if value := res.Get("vlan"); value.Exists() && !data.Vlan.IsNull() {
			data.Vlan = types.Int64Value(value.Int())
		} else {
			data.Vlan = types.Int64Null()
		}
		if value := res.Get("voiceVlan"); value.Exists() && !data.VoiceVlan.IsNull() {
			data.VoiceVlan = types.Int64Value(value.Int())
		} else {
			data.VoiceVlan = types.Int64Null()
		}
		if value := res.Get("dot3az.enabled"); value.Exists() && !data.Dot3azEnabled.IsNull() {
			data.Dot3azEnabled = types.BoolValue(value.Bool())
		} else {
			data.Dot3azEnabled = types.BoolNull()
		}
		if value := res.Get("highSpeed.enabled"); value.Exists() && !data.HighSpeedEnabled.IsNull() {
			data.HighSpeedEnabled = types.BoolValue(value.Bool())
		} else {
			data.HighSpeedEnabled = types.BoolNull()
		}
		if value := res.Get("profile.enabled"); value.Exists() && !data.ProfileEnabled.IsNull() {
			data.ProfileEnabled = types.BoolValue(value.Bool())
		} else {
			data.ProfileEnabled = types.BoolNull()
		}
		if value := res.Get("profile.id"); value.Exists() && !data.ProfileId.IsNull() {
			data.ProfileId = types.StringValue(value.String())
		} else {
			data.ProfileId = types.StringNull()
		}
		if value := res.Get("profile.iname"); value.Exists() && !data.ProfileIname.IsNull() {
			data.ProfileIname = types.StringValue(value.String())
		} else {
			data.ProfileIname = types.StringNull()
		}
		if value := res.Get("macAllowList"); value.Exists() && !data.MacAllowList.IsNull() {
			data.MacAllowList = helpers.GetStringSet(value.Array())
		} else {
			data.MacAllowList = types.SetNull(types.StringType)
		}
		if value := res.Get("stickyMacAllowList"); value.Exists() && !data.StickyMacAllowList.IsNull() {
			data.StickyMacAllowList = helpers.GetStringSet(value.Array())
		} else {
			data.StickyMacAllowList = types.SetNull(types.StringType)
		}
		if value := res.Get("tags"); value.Exists() && !data.Tags.IsNull() {
			data.Tags = helpers.GetStringSet(value.Array())
		} else {
			data.Tags = types.SetNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyPartial(), removing item with id: %s", data.Items[toBeDeleted[i]].PortId.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceSwitchPorts) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceSwitchPorts) fromBodyImport(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("portId").String() == (*parent).Items[i].PortId.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		if value := res.Get("portId"); value.Exists() && value.Value() != nil {
			data.PortId = types.StringValue(value.String())
		} else {
			data.PortId = types.StringNull()
		}
		if value := res.Get("accessPolicyNumber"); value.Exists() && value.Value() != nil {
			data.AccessPolicyNumber = types.Int64Value(value.Int())
		} else {
			data.AccessPolicyNumber = types.Int64Null()
		}
		if value := res.Get("accessPolicyType"); value.Exists() && value.Value() != nil {
			data.AccessPolicyType = types.StringValue(value.String())
		} else {
			data.AccessPolicyType = types.StringNull()
		}
		if value := res.Get("adaptivePolicyGroupId"); value.Exists() && value.Value() != nil {
			data.AdaptivePolicyGroupId = types.StringValue(value.String())
		} else {
			data.AdaptivePolicyGroupId = types.StringNull()
		}
		if value := res.Get("allowedVlans"); value.Exists() && value.Value() != nil {
			data.AllowedVlans = types.StringValue(value.String())
		} else {
			data.AllowedVlans = types.StringNull()
		}
		if value := res.Get("daiTrusted"); value.Exists() && value.Value() != nil {
			data.DaiTrusted = types.BoolValue(value.Bool())
		} else {
			data.DaiTrusted = types.BoolNull()
		}
		if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("flexibleStackingEnabled"); value.Exists() && value.Value() != nil {
			data.FlexibleStackingEnabled = types.BoolValue(value.Bool())
		} else {
			data.FlexibleStackingEnabled = types.BoolNull()
		}
		if value := res.Get("isolationEnabled"); value.Exists() && value.Value() != nil {
			data.IsolationEnabled = types.BoolValue(value.Bool())
		} else {
			data.IsolationEnabled = types.BoolNull()
		}
		if value := res.Get("linkNegotiation"); value.Exists() && value.Value() != nil {
			data.LinkNegotiation = types.StringValue(value.String())
		} else {
			data.LinkNegotiation = types.StringNull()
		}
		if value := res.Get("macWhitelistLimit"); value.Exists() && value.Value() != nil {
			data.MacWhitelistLimit = types.Int64Value(value.Int())
		} else {
			data.MacWhitelistLimit = types.Int64Null()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("peerSgtCapable"); value.Exists() && value.Value() != nil {
			data.PeerSgtCapable = types.BoolValue(value.Bool())
		} else {
			data.PeerSgtCapable = types.BoolNull()
		}
		if value := res.Get("poeEnabled"); value.Exists() && value.Value() != nil {
			data.PoeEnabled = types.BoolValue(value.Bool())
		} else {
			data.PoeEnabled = types.BoolNull()
		}
		if value := res.Get("portScheduleId"); value.Exists() && value.Value() != nil {
			data.PortScheduleId = types.StringValue(value.String())
		} else {
			data.PortScheduleId = types.StringNull()
		}
		if value := res.Get("rstpEnabled"); value.Exists() && value.Value() != nil {
			data.RstpEnabled = types.BoolValue(value.Bool())
		} else {
			data.RstpEnabled = types.BoolNull()
		}
		if value := res.Get("stickyMacAllowListLimit"); value.Exists() && value.Value() != nil {
			data.StickyMacAllowListLimit = types.Int64Value(value.Int())
		} else {
			data.StickyMacAllowListLimit = types.Int64Null()
		}
		if value := res.Get("stormControlEnabled"); value.Exists() && value.Value() != nil {
			data.StormControlEnabled = types.BoolValue(value.Bool())
		} else {
			data.StormControlEnabled = types.BoolNull()
		}
		if value := res.Get("stpGuard"); value.Exists() && value.Value() != nil {
			data.StpGuard = types.StringValue(value.String())
		} else {
			data.StpGuard = types.StringNull()
		}
		if value := res.Get("stpPortFastTrunk"); value.Exists() && value.Value() != nil {
			data.StpPortFastTrunk = types.BoolValue(value.Bool())
		} else {
			data.StpPortFastTrunk = types.BoolNull()
		}
		if value := res.Get("type"); value.Exists() && value.Value() != nil {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("udld"); value.Exists() && value.Value() != nil {
			data.Udld = types.StringValue(value.String())
		} else {
			data.Udld = types.StringNull()
		}
		if value := res.Get("vlan"); value.Exists() && value.Value() != nil {
			data.Vlan = types.Int64Value(value.Int())
		} else {
			data.Vlan = types.Int64Null()
		}
		if value := res.Get("voiceVlan"); value.Exists() && value.Value() != nil {
			data.VoiceVlan = types.Int64Value(value.Int())
		} else {
			data.VoiceVlan = types.Int64Null()
		}
		if value := res.Get("dot3az.enabled"); value.Exists() && value.Value() != nil {
			data.Dot3azEnabled = types.BoolValue(value.Bool())
		} else {
			data.Dot3azEnabled = types.BoolNull()
		}
		if value := res.Get("highSpeed.enabled"); value.Exists() && value.Value() != nil {
			data.HighSpeedEnabled = types.BoolValue(value.Bool())
		} else {
			data.HighSpeedEnabled = types.BoolNull()
		}
		if value := res.Get("profile.enabled"); value.Exists() && value.Value() != nil {
			data.ProfileEnabled = types.BoolValue(value.Bool())
		} else {
			data.ProfileEnabled = types.BoolNull()
		}
		if value := res.Get("profile.id"); value.Exists() && value.Value() != nil {
			data.ProfileId = types.StringValue(value.String())
		} else {
			data.ProfileId = types.StringNull()
		}
		if value := res.Get("profile.iname"); value.Exists() && value.Value() != nil {
			data.ProfileIname = types.StringValue(value.String())
		} else {
			data.ProfileIname = types.StringNull()
		}
		if value := res.Get("macAllowList"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.MacAllowList = helpers.GetStringSet(value.Array())
		} else {
			data.MacAllowList = types.SetNull(types.StringType)
		}
		if value := res.Get("stickyMacAllowList"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.StickyMacAllowList = helpers.GetStringSet(value.Array())
		} else {
			data.StickyMacAllowList = types.SetNull(types.StringType)
		}
		if value := res.Get("tags"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.Tags = helpers.GetStringSet(value.Array())
		} else {
			data.Tags = types.SetNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyImport(), removing item with id: %s", data.Items[toBeDeleted[i]].PortId.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ResourceSwitchPortsIdentity) toIdentity(ctx context.Context, plan *ResourceSwitchPorts) {
	data.OrganizationId = plan.OrganizationId
	data.Serial = plan.Serial
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ResourceSwitchPorts) fromIdentity(ctx context.Context, identity *ResourceSwitchPortsIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Serial = identity.Serial
}

// End of section. //template:end fromIdentity

func (data ResourceSwitchPortsItems) toDestroyBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "accessPolicyType", "Open")
	if !data.AdaptivePolicyGroupId.IsNull() {
		// This can fail with "Adaptive Policy is not enabled in this network"
		// if meraki_organization_adaptive_policy_settings resource
		// does not have adaptive policy enabled for this network.
		// Unset adaptive policy group ID only if it was already set
		// (which could only be done if adaptive policy *is* enabled on the network first).
		body, _ = sjson.Set(body, "adaptivePolicyGroupId", nil)
	}
	body, _ = sjson.Set(body, "profile.enabled", false)
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceSwitchPorts) hasChanges(ctx context.Context, state *ResourceSwitchPorts, id string) bool {
	hasChanges := false

	item := ResourceSwitchPortsItems{}
	for i := range data.Items {
		if data.Items[i].PortId.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceSwitchPortsItems{}
	for i := range state.Items {
		if state.Items[i].PortId.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.PortId.Equal(stateItem.PortId) {
		hasChanges = true
	}
	if !item.AccessPolicyNumber.Equal(stateItem.AccessPolicyNumber) {
		hasChanges = true
	}
	if !item.AccessPolicyType.Equal(stateItem.AccessPolicyType) {
		hasChanges = true
	}
	if !item.AdaptivePolicyGroupId.Equal(stateItem.AdaptivePolicyGroupId) {
		hasChanges = true
	}
	if !item.AllowedVlans.Equal(stateItem.AllowedVlans) {
		hasChanges = true
	}
	if !item.DaiTrusted.Equal(stateItem.DaiTrusted) {
		hasChanges = true
	}
	if !item.Enabled.Equal(stateItem.Enabled) {
		hasChanges = true
	}
	if !item.FlexibleStackingEnabled.Equal(stateItem.FlexibleStackingEnabled) {
		hasChanges = true
	}
	if !item.IsolationEnabled.Equal(stateItem.IsolationEnabled) {
		hasChanges = true
	}
	if !item.LinkNegotiation.Equal(stateItem.LinkNegotiation) {
		hasChanges = true
	}
	if !item.MacWhitelistLimit.Equal(stateItem.MacWhitelistLimit) {
		hasChanges = true
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	if !item.PeerSgtCapable.Equal(stateItem.PeerSgtCapable) {
		hasChanges = true
	}
	if !item.PoeEnabled.Equal(stateItem.PoeEnabled) {
		hasChanges = true
	}
	if !item.PortScheduleId.Equal(stateItem.PortScheduleId) {
		hasChanges = true
	}
	if !item.RstpEnabled.Equal(stateItem.RstpEnabled) {
		hasChanges = true
	}
	if !item.StickyMacAllowListLimit.Equal(stateItem.StickyMacAllowListLimit) {
		hasChanges = true
	}
	if !item.StormControlEnabled.Equal(stateItem.StormControlEnabled) {
		hasChanges = true
	}
	if !item.StpGuard.Equal(stateItem.StpGuard) {
		hasChanges = true
	}
	if !item.StpPortFastTrunk.Equal(stateItem.StpPortFastTrunk) {
		hasChanges = true
	}
	if !item.Type.Equal(stateItem.Type) {
		hasChanges = true
	}
	if !item.Udld.Equal(stateItem.Udld) {
		hasChanges = true
	}
	if !item.Vlan.Equal(stateItem.Vlan) {
		hasChanges = true
	}
	if !item.VoiceVlan.Equal(stateItem.VoiceVlan) {
		hasChanges = true
	}
	if !item.Dot3azEnabled.Equal(stateItem.Dot3azEnabled) {
		hasChanges = true
	}
	if !item.HighSpeedEnabled.Equal(stateItem.HighSpeedEnabled) {
		hasChanges = true
	}
	if !item.ProfileEnabled.Equal(stateItem.ProfileEnabled) {
		hasChanges = true
	}
	if !item.ProfileId.Equal(stateItem.ProfileId) {
		hasChanges = true
	}
	if !item.ProfileIname.Equal(stateItem.ProfileIname) {
		hasChanges = true
	}
	if !item.MacAllowList.Equal(stateItem.MacAllowList) {
		hasChanges = true
	}
	if !item.StickyMacAllowList.Equal(stateItem.StickyMacAllowList) {
		hasChanges = true
	}
	if !item.Tags.Equal(stateItem.Tags) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
