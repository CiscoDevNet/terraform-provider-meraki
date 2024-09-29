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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessRadioSettings struct {
	Id                            types.String `tfsdk:"id"`
	Serial                        types.String `tfsdk:"serial"`
	RfProfileId                   types.String `tfsdk:"rf_profile_id"`
	FiveGhzSettingsChannel        types.Int64  `tfsdk:"five_ghz_settings_channel"`
	FiveGhzSettingsChannelWidth   types.Int64  `tfsdk:"five_ghz_settings_channel_width"`
	FiveGhzSettingsTargetPower    types.Int64  `tfsdk:"five_ghz_settings_target_power"`
	TwoFourGhzSettingsChannel     types.Int64  `tfsdk:"two_four_ghz_settings_channel"`
	TwoFourGhzSettingsTargetPower types.Int64  `tfsdk:"two_four_ghz_settings_target_power"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessRadioSettings) getPath() string {
	return fmt.Sprintf("/devices/%v/wireless/radio/settings", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessRadioSettings) toBody(ctx context.Context, state WirelessRadioSettings) string {
	body := ""
	if !data.RfProfileId.IsNull() {
		body, _ = sjson.Set(body, "rfProfileId", data.RfProfileId.ValueString())
	}
	if !data.FiveGhzSettingsChannel.IsNull() {
		body, _ = sjson.Set(body, "fiveGhzSettings.channel", data.FiveGhzSettingsChannel.ValueInt64())
	}
	if !data.FiveGhzSettingsChannelWidth.IsNull() {
		body, _ = sjson.Set(body, "fiveGhzSettings.channelWidth", data.FiveGhzSettingsChannelWidth.ValueInt64())
	}
	if !data.FiveGhzSettingsTargetPower.IsNull() {
		body, _ = sjson.Set(body, "fiveGhzSettings.targetPower", data.FiveGhzSettingsTargetPower.ValueInt64())
	}
	if !data.TwoFourGhzSettingsChannel.IsNull() {
		body, _ = sjson.Set(body, "twoFourGhzSettings.channel", data.TwoFourGhzSettingsChannel.ValueInt64())
	}
	if !data.TwoFourGhzSettingsTargetPower.IsNull() {
		body, _ = sjson.Set(body, "twoFourGhzSettings.targetPower", data.TwoFourGhzSettingsTargetPower.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessRadioSettings) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("rfProfileId"); value.Exists() && value.Value() != nil {
		data.RfProfileId = types.StringValue(value.String())
	} else {
		data.RfProfileId = types.StringNull()
	}
	if value := res.Get("fiveGhzSettings.channel"); value.Exists() && value.Value() != nil {
		data.FiveGhzSettingsChannel = types.Int64Value(value.Int())
	} else {
		data.FiveGhzSettingsChannel = types.Int64Null()
	}
	if value := res.Get("fiveGhzSettings.channelWidth"); value.Exists() && value.Value() != nil {
		data.FiveGhzSettingsChannelWidth = types.Int64Value(value.Int())
	} else {
		data.FiveGhzSettingsChannelWidth = types.Int64Null()
	}
	if value := res.Get("fiveGhzSettings.targetPower"); value.Exists() && value.Value() != nil {
		data.FiveGhzSettingsTargetPower = types.Int64Value(value.Int())
	} else {
		data.FiveGhzSettingsTargetPower = types.Int64Null()
	}
	if value := res.Get("twoFourGhzSettings.channel"); value.Exists() && value.Value() != nil {
		data.TwoFourGhzSettingsChannel = types.Int64Value(value.Int())
	} else {
		data.TwoFourGhzSettingsChannel = types.Int64Null()
	}
	if value := res.Get("twoFourGhzSettings.targetPower"); value.Exists() && value.Value() != nil {
		data.TwoFourGhzSettingsTargetPower = types.Int64Value(value.Int())
	} else {
		data.TwoFourGhzSettingsTargetPower = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessRadioSettings) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("rfProfileId"); value.Exists() && !data.RfProfileId.IsNull() {
		data.RfProfileId = types.StringValue(value.String())
	} else {
		data.RfProfileId = types.StringNull()
	}
	if value := res.Get("fiveGhzSettings.channel"); value.Exists() && !data.FiveGhzSettingsChannel.IsNull() {
		data.FiveGhzSettingsChannel = types.Int64Value(value.Int())
	} else {
		data.FiveGhzSettingsChannel = types.Int64Null()
	}
	if value := res.Get("fiveGhzSettings.channelWidth"); value.Exists() && !data.FiveGhzSettingsChannelWidth.IsNull() {
		data.FiveGhzSettingsChannelWidth = types.Int64Value(value.Int())
	} else {
		data.FiveGhzSettingsChannelWidth = types.Int64Null()
	}
	if value := res.Get("fiveGhzSettings.targetPower"); value.Exists() && !data.FiveGhzSettingsTargetPower.IsNull() {
		data.FiveGhzSettingsTargetPower = types.Int64Value(value.Int())
	} else {
		data.FiveGhzSettingsTargetPower = types.Int64Null()
	}
	if value := res.Get("twoFourGhzSettings.channel"); value.Exists() && !data.TwoFourGhzSettingsChannel.IsNull() {
		data.TwoFourGhzSettingsChannel = types.Int64Value(value.Int())
	} else {
		data.TwoFourGhzSettingsChannel = types.Int64Null()
	}
	if value := res.Get("twoFourGhzSettings.targetPower"); value.Exists() && !data.TwoFourGhzSettingsTargetPower.IsNull() {
		data.TwoFourGhzSettingsTargetPower = types.Int64Value(value.Int())
	} else {
		data.TwoFourGhzSettingsTargetPower = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial
