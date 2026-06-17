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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkWirelessRadioRRM struct {
	Id                             types.String `tfsdk:"id"`
	NetworkId                      types.String `tfsdk:"network_id"`
	AiEnabled                      types.Bool   `tfsdk:"ai_enabled"`
	BusyHourMinimizeChangesEnabled types.Bool   `tfsdk:"busy_hour_minimize_changes_enabled"`
	BusyHourScheduleMode           types.String `tfsdk:"busy_hour_schedule_mode"`
	BusyHourScheduleManualEnd      types.String `tfsdk:"busy_hour_schedule_manual_end"`
	BusyHourScheduleManualStart    types.String `tfsdk:"busy_hour_schedule_manual_start"`
	ChannelAvoidanceEnabled        types.Bool   `tfsdk:"channel_avoidance_enabled"`
	FraEnabled                     types.Bool   `tfsdk:"fra_enabled"`
}

type NetworkWirelessRadioRRMIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkWirelessRadioRRM) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/radio/rrm", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkWirelessRadioRRM) toBody(ctx context.Context, state NetworkWirelessRadioRRM) string {
	body := ""
	if !data.AiEnabled.IsNull() {
		body, _ = sjson.Set(body, "ai.enabled", data.AiEnabled.ValueBool())
	}
	if !data.BusyHourMinimizeChangesEnabled.IsNull() {
		body, _ = sjson.Set(body, "busyHour.minimizeChanges.enabled", data.BusyHourMinimizeChangesEnabled.ValueBool())
	}
	if !data.BusyHourScheduleMode.IsNull() {
		body, _ = sjson.Set(body, "busyHour.schedule.mode", data.BusyHourScheduleMode.ValueString())
	}
	if !data.BusyHourScheduleManualEnd.IsNull() {
		body, _ = sjson.Set(body, "busyHour.schedule.manual.end", data.BusyHourScheduleManualEnd.ValueString())
	}
	if !data.BusyHourScheduleManualStart.IsNull() {
		body, _ = sjson.Set(body, "busyHour.schedule.manual.start", data.BusyHourScheduleManualStart.ValueString())
	}
	if !data.ChannelAvoidanceEnabled.IsNull() {
		body, _ = sjson.Set(body, "channel.avoidance.enabled", data.ChannelAvoidanceEnabled.ValueBool())
	}
	if !data.FraEnabled.IsNull() {
		body, _ = sjson.Set(body, "fra.enabled", data.FraEnabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkWirelessRadioRRM) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("ai.enabled"); value.Exists() && value.Value() != nil {
		data.AiEnabled = types.BoolValue(value.Bool())
	} else {
		data.AiEnabled = types.BoolNull()
	}
	if value := res.Get("busyHour.minimizeChanges.enabled"); value.Exists() && value.Value() != nil {
		data.BusyHourMinimizeChangesEnabled = types.BoolValue(value.Bool())
	} else {
		data.BusyHourMinimizeChangesEnabled = types.BoolNull()
	}
	if value := res.Get("busyHour.schedule.mode"); value.Exists() && value.Value() != nil {
		data.BusyHourScheduleMode = types.StringValue(value.String())
	} else {
		data.BusyHourScheduleMode = types.StringNull()
	}
	if value := res.Get("busyHour.schedule.manual.end"); value.Exists() && value.Value() != nil {
		data.BusyHourScheduleManualEnd = types.StringValue(value.String())
	} else {
		data.BusyHourScheduleManualEnd = types.StringNull()
	}
	if value := res.Get("busyHour.schedule.manual.start"); value.Exists() && value.Value() != nil {
		data.BusyHourScheduleManualStart = types.StringValue(value.String())
	} else {
		data.BusyHourScheduleManualStart = types.StringNull()
	}
	if value := res.Get("channel.avoidance.enabled"); value.Exists() && value.Value() != nil {
		data.ChannelAvoidanceEnabled = types.BoolValue(value.Bool())
	} else {
		data.ChannelAvoidanceEnabled = types.BoolNull()
	}
	if value := res.Get("fra.enabled"); value.Exists() && value.Value() != nil {
		data.FraEnabled = types.BoolValue(value.Bool())
	} else {
		data.FraEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkWirelessRadioRRM) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("ai.enabled"); value.Exists() && !data.AiEnabled.IsNull() {
		data.AiEnabled = types.BoolValue(value.Bool())
	} else {
		data.AiEnabled = types.BoolNull()
	}
	if value := res.Get("busyHour.minimizeChanges.enabled"); value.Exists() && !data.BusyHourMinimizeChangesEnabled.IsNull() {
		data.BusyHourMinimizeChangesEnabled = types.BoolValue(value.Bool())
	} else {
		data.BusyHourMinimizeChangesEnabled = types.BoolNull()
	}
	if value := res.Get("busyHour.schedule.mode"); value.Exists() && !data.BusyHourScheduleMode.IsNull() {
		data.BusyHourScheduleMode = types.StringValue(value.String())
	} else {
		data.BusyHourScheduleMode = types.StringNull()
	}
	if value := res.Get("busyHour.schedule.manual.end"); value.Exists() && !data.BusyHourScheduleManualEnd.IsNull() {
		data.BusyHourScheduleManualEnd = types.StringValue(value.String())
	} else {
		data.BusyHourScheduleManualEnd = types.StringNull()
	}
	if value := res.Get("busyHour.schedule.manual.start"); value.Exists() && !data.BusyHourScheduleManualStart.IsNull() {
		data.BusyHourScheduleManualStart = types.StringValue(value.String())
	} else {
		data.BusyHourScheduleManualStart = types.StringNull()
	}
	if value := res.Get("channel.avoidance.enabled"); value.Exists() && !data.ChannelAvoidanceEnabled.IsNull() {
		data.ChannelAvoidanceEnabled = types.BoolValue(value.Bool())
	} else {
		data.ChannelAvoidanceEnabled = types.BoolNull()
	}
	if value := res.Get("fra.enabled"); value.Exists() && !data.FraEnabled.IsNull() {
		data.FraEnabled = types.BoolValue(value.Bool())
	} else {
		data.FraEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkWirelessRadioRRM) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *NetworkWirelessRadioRRMIdentity) toIdentity(ctx context.Context, plan *NetworkWirelessRadioRRM) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *NetworkWirelessRadioRRM) fromIdentity(ctx context.Context, identity *NetworkWirelessRadioRRMIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkWirelessRadioRRM) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
