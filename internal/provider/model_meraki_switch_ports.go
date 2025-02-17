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
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchPorts struct {
	Serial types.String       `tfsdk:"serial"`
	Items  []SwitchPortsItems `tfsdk:"items"`
}

type SwitchPortsItems struct {
	Id                                     types.String `tfsdk:"id"`
	AccessPolicyNumber                     types.Int64  `tfsdk:"access_policy_number"`
	AccessPolicyType                       types.String `tfsdk:"access_policy_type"`
	AllowedVlans                           types.String `tfsdk:"allowed_vlans"`
	DaiTrusted                             types.Bool   `tfsdk:"dai_trusted"`
	Enabled                                types.Bool   `tfsdk:"enabled"`
	FlexibleStackingEnabled                types.Bool   `tfsdk:"flexible_stacking_enabled"`
	IsolationEnabled                       types.Bool   `tfsdk:"isolation_enabled"`
	LinkNegotiation                        types.String `tfsdk:"link_negotiation"`
	Name                                   types.String `tfsdk:"name"`
	PeerSgtCapable                         types.Bool   `tfsdk:"peer_sgt_capable"`
	PoeEnabled                             types.Bool   `tfsdk:"poe_enabled"`
	PortId                                 types.String `tfsdk:"port_id"`
	PortScheduleId                         types.String `tfsdk:"port_schedule_id"`
	RstpEnabled                            types.Bool   `tfsdk:"rstp_enabled"`
	StickyMacAllowListLimit                types.Int64  `tfsdk:"sticky_mac_allow_list_limit"`
	StormControlEnabled                    types.Bool   `tfsdk:"storm_control_enabled"`
	StpGuard                               types.String `tfsdk:"stp_guard"`
	Type                                   types.String `tfsdk:"type"`
	Udld                                   types.String `tfsdk:"udld"`
	Vlan                                   types.Int64  `tfsdk:"vlan"`
	VoiceVlan                              types.Int64  `tfsdk:"voice_vlan"`
	AdaptivePolicyGroupId                  types.String `tfsdk:"adaptive_policy_group_id"`
	AdaptivePolicyGroupName                types.String `tfsdk:"adaptive_policy_group_name"`
	Dot3azEnabled                          types.Bool   `tfsdk:"dot3az_enabled"`
	MirrorMode                             types.String `tfsdk:"mirror_mode"`
	ModuleModel                            types.String `tfsdk:"module_model"`
	ProfileEnabled                         types.Bool   `tfsdk:"profile_enabled"`
	ProfileId                              types.String `tfsdk:"profile_id"`
	ProfileIname                           types.String `tfsdk:"profile_iname"`
	ScheduleId                             types.String `tfsdk:"schedule_id"`
	ScheduleName                           types.String `tfsdk:"schedule_name"`
	StackwiseVirtualIsDualActiveDetector   types.Bool   `tfsdk:"stackwise_virtual_is_dual_active_detector"`
	StackwiseVirtualIsStackWiseVirtualLink types.Bool   `tfsdk:"stackwise_virtual_is_stack_wise_virtual_link"`
	LinkNegotiationCapabilities            types.List   `tfsdk:"link_negotiation_capabilities"`
	MacAllowList                           types.List   `tfsdk:"mac_allow_list"`
	StickyMacAllowList                     types.List   `tfsdk:"sticky_mac_allow_list"`
	Tags                                   types.List   `tfsdk:"tags"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchPorts) getPath() string {
	return fmt.Sprintf("/devices/%v/switch/ports", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchPorts) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]SwitchPortsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := SwitchPortsItems{}
		data.Id = types.StringValue(res.Get("").String())
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
		if value := res.Get("portId"); value.Exists() && value.Value() != nil {
			data.PortId = types.StringValue(value.String())
		} else {
			data.PortId = types.StringNull()
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
		if value := res.Get("adaptivePolicyGroup.id"); value.Exists() && value.Value() != nil {
			data.AdaptivePolicyGroupId = types.StringValue(value.String())
		} else {
			data.AdaptivePolicyGroupId = types.StringNull()
		}
		if value := res.Get("adaptivePolicyGroup.name"); value.Exists() && value.Value() != nil {
			data.AdaptivePolicyGroupName = types.StringValue(value.String())
		} else {
			data.AdaptivePolicyGroupName = types.StringNull()
		}
		if value := res.Get("dot3az.enabled"); value.Exists() && value.Value() != nil {
			data.Dot3azEnabled = types.BoolValue(value.Bool())
		} else {
			data.Dot3azEnabled = types.BoolNull()
		}
		if value := res.Get("mirror.mode"); value.Exists() && value.Value() != nil {
			data.MirrorMode = types.StringValue(value.String())
		} else {
			data.MirrorMode = types.StringNull()
		}
		if value := res.Get("module.model"); value.Exists() && value.Value() != nil {
			data.ModuleModel = types.StringValue(value.String())
		} else {
			data.ModuleModel = types.StringNull()
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
		if value := res.Get("schedule.id"); value.Exists() && value.Value() != nil {
			data.ScheduleId = types.StringValue(value.String())
		} else {
			data.ScheduleId = types.StringNull()
		}
		if value := res.Get("schedule.name"); value.Exists() && value.Value() != nil {
			data.ScheduleName = types.StringValue(value.String())
		} else {
			data.ScheduleName = types.StringNull()
		}
		if value := res.Get("stackwiseVirtual.isDualActiveDetector"); value.Exists() && value.Value() != nil {
			data.StackwiseVirtualIsDualActiveDetector = types.BoolValue(value.Bool())
		} else {
			data.StackwiseVirtualIsDualActiveDetector = types.BoolNull()
		}
		if value := res.Get("stackwiseVirtual.isStackWiseVirtualLink"); value.Exists() && value.Value() != nil {
			data.StackwiseVirtualIsStackWiseVirtualLink = types.BoolValue(value.Bool())
		} else {
			data.StackwiseVirtualIsStackWiseVirtualLink = types.BoolNull()
		}
		if value := res.Get("linkNegotiationCapabilities"); value.Exists() && value.Value() != nil {
			data.LinkNegotiationCapabilities = helpers.GetStringList(value.Array())
		} else {
			data.LinkNegotiationCapabilities = types.ListNull(types.StringType)
		}
		if value := res.Get("macAllowList"); value.Exists() && value.Value() != nil {
			data.MacAllowList = helpers.GetStringList(value.Array())
		} else {
			data.MacAllowList = types.ListNull(types.StringType)
		}
		if value := res.Get("stickyMacAllowList"); value.Exists() && value.Value() != nil {
			data.StickyMacAllowList = helpers.GetStringList(value.Array())
		} else {
			data.StickyMacAllowList = types.ListNull(types.StringType)
		}
		if value := res.Get("tags"); value.Exists() && value.Value() != nil {
			data.Tags = helpers.GetStringList(value.Array())
		} else {
			data.Tags = types.ListNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
