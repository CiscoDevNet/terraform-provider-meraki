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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceRFProfiles struct {
	NetworkId types.String               `tfsdk:"network_id"`
	Items     []ApplianceRFProfilesItems `tfsdk:"items"`
}

type ApplianceRFProfilesItems struct {
	Id                                  types.String  `tfsdk:"id"`
	Name                                types.String  `tfsdk:"name"`
	FiveGhzSettingsAxEnabled            types.Bool    `tfsdk:"five_ghz_settings_ax_enabled"`
	FiveGhzSettingsMinBitrate           types.Int64   `tfsdk:"five_ghz_settings_min_bitrate"`
	PerSsidSettings1BandOperationMode   types.String  `tfsdk:"per_ssid_settings_1_band_operation_mode"`
	PerSsidSettings1BandSteeringEnabled types.Bool    `tfsdk:"per_ssid_settings_1_band_steering_enabled"`
	PerSsidSettings2BandOperationMode   types.String  `tfsdk:"per_ssid_settings_2_band_operation_mode"`
	PerSsidSettings2BandSteeringEnabled types.Bool    `tfsdk:"per_ssid_settings_2_band_steering_enabled"`
	PerSsidSettings3BandOperationMode   types.String  `tfsdk:"per_ssid_settings_3_band_operation_mode"`
	PerSsidSettings3BandSteeringEnabled types.Bool    `tfsdk:"per_ssid_settings_3_band_steering_enabled"`
	PerSsidSettings4BandOperationMode   types.String  `tfsdk:"per_ssid_settings_4_band_operation_mode"`
	PerSsidSettings4BandSteeringEnabled types.Bool    `tfsdk:"per_ssid_settings_4_band_steering_enabled"`
	TwoFourGhzSettingsAxEnabled         types.Bool    `tfsdk:"two_four_ghz_settings_ax_enabled"`
	TwoFourGhzSettingsMinBitrate        types.Float64 `tfsdk:"two_four_ghz_settings_min_bitrate"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceRFProfiles) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/rfProfiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceRFProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ApplianceRFProfilesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ApplianceRFProfilesItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("fiveGhzSettings.axEnabled"); value.Exists() && value.Value() != nil {
			data.FiveGhzSettingsAxEnabled = types.BoolValue(value.Bool())
		} else {
			data.FiveGhzSettingsAxEnabled = types.BoolNull()
		}
		if value := res.Get("fiveGhzSettings.minBitrate"); value.Exists() && value.Value() != nil {
			data.FiveGhzSettingsMinBitrate = types.Int64Value(value.Int())
		} else {
			data.FiveGhzSettingsMinBitrate = types.Int64Null()
		}
		if value := res.Get("perSsidSettings.1.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings1BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings1BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.1.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings1BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings1BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.2.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings2BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings2BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.2.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings2BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings2BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.3.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings3BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings3BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.3.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings3BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings3BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.4.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings4BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings4BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.4.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings4BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings4BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("twoFourGhzSettings.axEnabled"); value.Exists() && value.Value() != nil {
			data.TwoFourGhzSettingsAxEnabled = types.BoolValue(value.Bool())
		} else {
			data.TwoFourGhzSettingsAxEnabled = types.BoolNull()
		}
		if value := res.Get("twoFourGhzSettings.minBitrate"); value.Exists() && value.Value() != nil {
			data.TwoFourGhzSettingsMinBitrate = types.Float64Value(value.Float())
		} else {
			data.TwoFourGhzSettingsMinBitrate = types.Float64Null()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
