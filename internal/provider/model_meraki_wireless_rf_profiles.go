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

type WirelessRFProfiles struct {
	NetworkId types.String              `tfsdk:"network_id"`
	Items     []WirelessRFProfilesItems `tfsdk:"items"`
}

type WirelessRFProfilesItems struct {
	Id                                   types.String                          `tfsdk:"id"`
	BandSelectionType                    types.String                          `tfsdk:"band_selection_type"`
	ClientBalancingEnabled               types.Bool                            `tfsdk:"client_balancing_enabled"`
	MinBitrateType                       types.String                          `tfsdk:"min_bitrate_type"`
	Name                                 types.String                          `tfsdk:"name"`
	ApBandSettingsBandOperationMode      types.String                          `tfsdk:"ap_band_settings_band_operation_mode"`
	ApBandSettingsBandSteeringEnabled    types.Bool                            `tfsdk:"ap_band_settings_band_steering_enabled"`
	ApBandSettingsBandsEnabled           types.Set                             `tfsdk:"ap_band_settings_bands_enabled"`
	FiveGhzSettingsChannelWidth          types.String                          `tfsdk:"five_ghz_settings_channel_width"`
	FiveGhzSettingsMaxPower              types.Int64                           `tfsdk:"five_ghz_settings_max_power"`
	FiveGhzSettingsMinBitrate            types.Int64                           `tfsdk:"five_ghz_settings_min_bitrate"`
	FiveGhzSettingsMinPower              types.Int64                           `tfsdk:"five_ghz_settings_min_power"`
	FiveGhzSettingsRxsop                 types.Int64                           `tfsdk:"five_ghz_settings_rxsop"`
	FiveGhzSettingsValidAutoChannels     types.Set                             `tfsdk:"five_ghz_settings_valid_auto_channels"`
	FlexRadiosByModel                    []WirelessRFProfilesFlexRadiosByModel `tfsdk:"flex_radios_by_model"`
	PerSsidSettings0BandOperationMode    types.String                          `tfsdk:"per_ssid_settings0_band_operation_mode"`
	PerSsidSettings0BandSteeringEnabled  types.Bool                            `tfsdk:"per_ssid_settings0_band_steering_enabled"`
	PerSsidSettings0MinBitrate           types.Float64                         `tfsdk:"per_ssid_settings0_min_bitrate"`
	PerSsidSettings0BandsEnabled         types.Set                             `tfsdk:"per_ssid_settings0_bands_enabled"`
	PerSsidSettings1BandOperationMode    types.String                          `tfsdk:"per_ssid_settings1_band_operation_mode"`
	PerSsidSettings1BandSteeringEnabled  types.Bool                            `tfsdk:"per_ssid_settings1_band_steering_enabled"`
	PerSsidSettings1MinBitrate           types.Float64                         `tfsdk:"per_ssid_settings1_min_bitrate"`
	PerSsidSettings1BandsEnabled         types.Set                             `tfsdk:"per_ssid_settings1_bands_enabled"`
	PerSsidSettings10BandOperationMode   types.String                          `tfsdk:"per_ssid_settings10_band_operation_mode"`
	PerSsidSettings10BandSteeringEnabled types.Bool                            `tfsdk:"per_ssid_settings10_band_steering_enabled"`
	PerSsidSettings10MinBitrate          types.Float64                         `tfsdk:"per_ssid_settings10_min_bitrate"`
	PerSsidSettings10BandsEnabled        types.Set                             `tfsdk:"per_ssid_settings10_bands_enabled"`
	PerSsidSettings11BandOperationMode   types.String                          `tfsdk:"per_ssid_settings11_band_operation_mode"`
	PerSsidSettings11BandSteeringEnabled types.Bool                            `tfsdk:"per_ssid_settings11_band_steering_enabled"`
	PerSsidSettings11MinBitrate          types.Float64                         `tfsdk:"per_ssid_settings11_min_bitrate"`
	PerSsidSettings11BandsEnabled        types.Set                             `tfsdk:"per_ssid_settings11_bands_enabled"`
	PerSsidSettings12BandOperationMode   types.String                          `tfsdk:"per_ssid_settings12_band_operation_mode"`
	PerSsidSettings12BandSteeringEnabled types.Bool                            `tfsdk:"per_ssid_settings12_band_steering_enabled"`
	PerSsidSettings12MinBitrate          types.Float64                         `tfsdk:"per_ssid_settings12_min_bitrate"`
	PerSsidSettings12BandsEnabled        types.Set                             `tfsdk:"per_ssid_settings12_bands_enabled"`
	PerSsidSettings13BandOperationMode   types.String                          `tfsdk:"per_ssid_settings13_band_operation_mode"`
	PerSsidSettings13BandSteeringEnabled types.Bool                            `tfsdk:"per_ssid_settings13_band_steering_enabled"`
	PerSsidSettings13MinBitrate          types.Float64                         `tfsdk:"per_ssid_settings13_min_bitrate"`
	PerSsidSettings13BandsEnabled        types.Set                             `tfsdk:"per_ssid_settings13_bands_enabled"`
	PerSsidSettings14BandOperationMode   types.String                          `tfsdk:"per_ssid_settings14_band_operation_mode"`
	PerSsidSettings14BandSteeringEnabled types.Bool                            `tfsdk:"per_ssid_settings14_band_steering_enabled"`
	PerSsidSettings14MinBitrate          types.Float64                         `tfsdk:"per_ssid_settings14_min_bitrate"`
	PerSsidSettings14BandsEnabled        types.Set                             `tfsdk:"per_ssid_settings14_bands_enabled"`
	PerSsidSettings2BandOperationMode    types.String                          `tfsdk:"per_ssid_settings2_band_operation_mode"`
	PerSsidSettings2BandSteeringEnabled  types.Bool                            `tfsdk:"per_ssid_settings2_band_steering_enabled"`
	PerSsidSettings2MinBitrate           types.Float64                         `tfsdk:"per_ssid_settings2_min_bitrate"`
	PerSsidSettings2BandsEnabled         types.Set                             `tfsdk:"per_ssid_settings2_bands_enabled"`
	PerSsidSettings3BandOperationMode    types.String                          `tfsdk:"per_ssid_settings3_band_operation_mode"`
	PerSsidSettings3BandSteeringEnabled  types.Bool                            `tfsdk:"per_ssid_settings3_band_steering_enabled"`
	PerSsidSettings3MinBitrate           types.Float64                         `tfsdk:"per_ssid_settings3_min_bitrate"`
	PerSsidSettings3BandsEnabled         types.Set                             `tfsdk:"per_ssid_settings3_bands_enabled"`
	PerSsidSettings4BandOperationMode    types.String                          `tfsdk:"per_ssid_settings4_band_operation_mode"`
	PerSsidSettings4BandSteeringEnabled  types.Bool                            `tfsdk:"per_ssid_settings4_band_steering_enabled"`
	PerSsidSettings4MinBitrate           types.Float64                         `tfsdk:"per_ssid_settings4_min_bitrate"`
	PerSsidSettings4BandsEnabled         types.Set                             `tfsdk:"per_ssid_settings4_bands_enabled"`
	PerSsidSettings5BandOperationMode    types.String                          `tfsdk:"per_ssid_settings5_band_operation_mode"`
	PerSsidSettings5BandSteeringEnabled  types.Bool                            `tfsdk:"per_ssid_settings5_band_steering_enabled"`
	PerSsidSettings5MinBitrate           types.Float64                         `tfsdk:"per_ssid_settings5_min_bitrate"`
	PerSsidSettings5BandsEnabled         types.Set                             `tfsdk:"per_ssid_settings5_bands_enabled"`
	PerSsidSettings6BandOperationMode    types.String                          `tfsdk:"per_ssid_settings6_band_operation_mode"`
	PerSsidSettings6BandSteeringEnabled  types.Bool                            `tfsdk:"per_ssid_settings6_band_steering_enabled"`
	PerSsidSettings6MinBitrate           types.Float64                         `tfsdk:"per_ssid_settings6_min_bitrate"`
	PerSsidSettings6BandsEnabled         types.Set                             `tfsdk:"per_ssid_settings6_bands_enabled"`
	PerSsidSettings7BandOperationMode    types.String                          `tfsdk:"per_ssid_settings7_band_operation_mode"`
	PerSsidSettings7BandSteeringEnabled  types.Bool                            `tfsdk:"per_ssid_settings7_band_steering_enabled"`
	PerSsidSettings7MinBitrate           types.Float64                         `tfsdk:"per_ssid_settings7_min_bitrate"`
	PerSsidSettings7BandsEnabled         types.Set                             `tfsdk:"per_ssid_settings7_bands_enabled"`
	PerSsidSettings8BandOperationMode    types.String                          `tfsdk:"per_ssid_settings8_band_operation_mode"`
	PerSsidSettings8BandSteeringEnabled  types.Bool                            `tfsdk:"per_ssid_settings8_band_steering_enabled"`
	PerSsidSettings8MinBitrate           types.Float64                         `tfsdk:"per_ssid_settings8_min_bitrate"`
	PerSsidSettings8BandsEnabled         types.Set                             `tfsdk:"per_ssid_settings8_bands_enabled"`
	PerSsidSettings9BandOperationMode    types.String                          `tfsdk:"per_ssid_settings9_band_operation_mode"`
	PerSsidSettings9BandSteeringEnabled  types.Bool                            `tfsdk:"per_ssid_settings9_band_steering_enabled"`
	PerSsidSettings9MinBitrate           types.Float64                         `tfsdk:"per_ssid_settings9_min_bitrate"`
	PerSsidSettings9BandsEnabled         types.Set                             `tfsdk:"per_ssid_settings9_bands_enabled"`
	SixGhzSettingsChannelWidth           types.String                          `tfsdk:"six_ghz_settings_channel_width"`
	SixGhzSettingsMaxPower               types.Int64                           `tfsdk:"six_ghz_settings_max_power"`
	SixGhzSettingsMinBitrate             types.Int64                           `tfsdk:"six_ghz_settings_min_bitrate"`
	SixGhzSettingsMinPower               types.Int64                           `tfsdk:"six_ghz_settings_min_power"`
	SixGhzSettingsRxsop                  types.Int64                           `tfsdk:"six_ghz_settings_rxsop"`
	SixGhzSettingsValidAutoChannels      types.Set                             `tfsdk:"six_ghz_settings_valid_auto_channels"`
	TransmissionEnabled                  types.Bool                            `tfsdk:"transmission_enabled"`
	TwoFourGhzSettingsAxEnabled          types.Bool                            `tfsdk:"two_four_ghz_settings_ax_enabled"`
	TwoFourGhzSettingsMaxPower           types.Int64                           `tfsdk:"two_four_ghz_settings_max_power"`
	TwoFourGhzSettingsMinBitrate         types.Float64                         `tfsdk:"two_four_ghz_settings_min_bitrate"`
	TwoFourGhzSettingsMinPower           types.Int64                           `tfsdk:"two_four_ghz_settings_min_power"`
	TwoFourGhzSettingsRxsop              types.Int64                           `tfsdk:"two_four_ghz_settings_rxsop"`
	TwoFourGhzSettingsValidAutoChannels  types.Set                             `tfsdk:"two_four_ghz_settings_valid_auto_channels"`
}

type WirelessRFProfilesFlexRadiosByModel struct {
	Model types.String `tfsdk:"model"`
	Bands types.Set    `tfsdk:"bands"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessRFProfiles) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/rfProfiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessRFProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]WirelessRFProfilesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := WirelessRFProfilesItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("bandSelectionType"); value.Exists() && value.Value() != nil {
			data.BandSelectionType = types.StringValue(value.String())
		} else {
			data.BandSelectionType = types.StringNull()
		}
		if value := res.Get("clientBalancingEnabled"); value.Exists() && value.Value() != nil {
			data.ClientBalancingEnabled = types.BoolValue(value.Bool())
		} else {
			data.ClientBalancingEnabled = types.BoolNull()
		}
		if value := res.Get("minBitrateType"); value.Exists() && value.Value() != nil {
			data.MinBitrateType = types.StringValue(value.String())
		} else {
			data.MinBitrateType = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("apBandSettings.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.ApBandSettingsBandOperationMode = types.StringValue(value.String())
		} else {
			data.ApBandSettingsBandOperationMode = types.StringNull()
		}
		if value := res.Get("apBandSettings.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.ApBandSettingsBandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.ApBandSettingsBandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("apBandSettings.bands.enabled"); value.Exists() && value.Value() != nil {
			data.ApBandSettingsBandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.ApBandSettingsBandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("fiveGhzSettings.channelWidth"); value.Exists() && value.Value() != nil {
			data.FiveGhzSettingsChannelWidth = types.StringValue(value.String())
		} else {
			data.FiveGhzSettingsChannelWidth = types.StringNull()
		}
		if value := res.Get("fiveGhzSettings.maxPower"); value.Exists() && value.Value() != nil {
			data.FiveGhzSettingsMaxPower = types.Int64Value(value.Int())
		} else {
			data.FiveGhzSettingsMaxPower = types.Int64Null()
		}
		if value := res.Get("fiveGhzSettings.minBitrate"); value.Exists() && value.Value() != nil {
			data.FiveGhzSettingsMinBitrate = types.Int64Value(value.Int())
		} else {
			data.FiveGhzSettingsMinBitrate = types.Int64Null()
		}
		if value := res.Get("fiveGhzSettings.minPower"); value.Exists() && value.Value() != nil {
			data.FiveGhzSettingsMinPower = types.Int64Value(value.Int())
		} else {
			data.FiveGhzSettingsMinPower = types.Int64Null()
		}
		if value := res.Get("fiveGhzSettings.rxsop"); value.Exists() && value.Value() != nil {
			data.FiveGhzSettingsRxsop = types.Int64Value(value.Int())
		} else {
			data.FiveGhzSettingsRxsop = types.Int64Null()
		}
		if value := res.Get("fiveGhzSettings.validAutoChannels"); value.Exists() && value.Value() != nil {
			data.FiveGhzSettingsValidAutoChannels = helpers.GetInt64Set(value.Array())
		} else {
			data.FiveGhzSettingsValidAutoChannels = types.SetNull(types.Int64Type)
		}
		if value := res.Get("flexRadios.byModel"); value.Exists() && value.Value() != nil {
			data.FlexRadiosByModel = make([]WirelessRFProfilesFlexRadiosByModel, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := WirelessRFProfilesFlexRadiosByModel{}
				if value := res.Get("model"); value.Exists() && value.Value() != nil {
					data.Model = types.StringValue(value.String())
				} else {
					data.Model = types.StringNull()
				}
				if value := res.Get("bands"); value.Exists() && value.Value() != nil {
					data.Bands = helpers.GetStringSet(value.Array())
				} else {
					data.Bands = types.SetNull(types.StringType)
				}
				(*parent).FlexRadiosByModel = append((*parent).FlexRadiosByModel, data)
				return true
			})
		}
		if value := res.Get("perSsidSettings.0.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings0BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings0BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.0.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings0BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings0BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.0.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings0MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings0MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.0.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings0BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings0BandsEnabled = types.SetNull(types.StringType)
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
		if value := res.Get("perSsidSettings.1.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings1MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings1MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.1.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings1BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings1BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.10.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings10BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings10BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.10.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings10BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings10BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.10.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings10MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings10MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.10.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings10BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings10BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.11.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings11BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings11BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.11.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings11BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings11BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.11.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings11MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings11MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.11.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings11BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings11BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.12.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings12BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings12BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.12.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings12BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings12BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.12.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings12MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings12MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.12.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings12BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings12BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.13.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings13BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings13BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.13.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings13BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings13BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.13.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings13MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings13MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.13.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings13BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings13BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.14.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings14BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings14BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.14.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings14BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings14BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.14.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings14MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings14MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.14.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings14BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings14BandsEnabled = types.SetNull(types.StringType)
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
		if value := res.Get("perSsidSettings.2.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings2MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings2MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.2.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings2BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings2BandsEnabled = types.SetNull(types.StringType)
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
		if value := res.Get("perSsidSettings.3.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings3MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings3MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.3.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings3BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings3BandsEnabled = types.SetNull(types.StringType)
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
		if value := res.Get("perSsidSettings.4.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings4MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings4MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.4.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings4BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings4BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.5.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings5BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings5BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.5.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings5BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings5BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.5.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings5MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings5MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.5.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings5BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings5BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.6.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings6BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings6BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.6.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings6BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings6BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.6.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings6MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings6MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.6.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings6BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings6BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.7.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings7BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings7BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.7.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings7BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings7BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.7.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings7MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings7MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.7.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings7BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings7BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.8.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings8BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings8BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.8.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings8BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings8BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.8.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings8MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings8MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.8.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings8BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings8BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.9.bandOperationMode"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings9BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings9BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.9.bandSteeringEnabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings9BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings9BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.9.minBitrate"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings9MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings9MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.9.bands.enabled"); value.Exists() && value.Value() != nil {
			data.PerSsidSettings9BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings9BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("sixGhzSettings.channelWidth"); value.Exists() && value.Value() != nil {
			data.SixGhzSettingsChannelWidth = types.StringValue(value.String())
		} else {
			data.SixGhzSettingsChannelWidth = types.StringNull()
		}
		if value := res.Get("sixGhzSettings.maxPower"); value.Exists() && value.Value() != nil {
			data.SixGhzSettingsMaxPower = types.Int64Value(value.Int())
		} else {
			data.SixGhzSettingsMaxPower = types.Int64Null()
		}
		if value := res.Get("sixGhzSettings.minBitrate"); value.Exists() && value.Value() != nil {
			data.SixGhzSettingsMinBitrate = types.Int64Value(value.Int())
		} else {
			data.SixGhzSettingsMinBitrate = types.Int64Null()
		}
		if value := res.Get("sixGhzSettings.minPower"); value.Exists() && value.Value() != nil {
			data.SixGhzSettingsMinPower = types.Int64Value(value.Int())
		} else {
			data.SixGhzSettingsMinPower = types.Int64Null()
		}
		if value := res.Get("sixGhzSettings.rxsop"); value.Exists() && value.Value() != nil {
			data.SixGhzSettingsRxsop = types.Int64Value(value.Int())
		} else {
			data.SixGhzSettingsRxsop = types.Int64Null()
		}
		if value := res.Get("sixGhzSettings.validAutoChannels"); value.Exists() && value.Value() != nil {
			data.SixGhzSettingsValidAutoChannels = helpers.GetInt64Set(value.Array())
		} else {
			data.SixGhzSettingsValidAutoChannels = types.SetNull(types.Int64Type)
		}
		if value := res.Get("transmission.enabled"); value.Exists() && value.Value() != nil {
			data.TransmissionEnabled = types.BoolValue(value.Bool())
		} else {
			data.TransmissionEnabled = types.BoolNull()
		}
		if value := res.Get("twoFourGhzSettings.axEnabled"); value.Exists() && value.Value() != nil {
			data.TwoFourGhzSettingsAxEnabled = types.BoolValue(value.Bool())
		} else {
			data.TwoFourGhzSettingsAxEnabled = types.BoolNull()
		}
		if value := res.Get("twoFourGhzSettings.maxPower"); value.Exists() && value.Value() != nil {
			data.TwoFourGhzSettingsMaxPower = types.Int64Value(value.Int())
		} else {
			data.TwoFourGhzSettingsMaxPower = types.Int64Null()
		}
		if value := res.Get("twoFourGhzSettings.minBitrate"); value.Exists() && value.Value() != nil {
			data.TwoFourGhzSettingsMinBitrate = types.Float64Value(value.Float())
		} else {
			data.TwoFourGhzSettingsMinBitrate = types.Float64Null()
		}
		if value := res.Get("twoFourGhzSettings.minPower"); value.Exists() && value.Value() != nil {
			data.TwoFourGhzSettingsMinPower = types.Int64Value(value.Int())
		} else {
			data.TwoFourGhzSettingsMinPower = types.Int64Null()
		}
		if value := res.Get("twoFourGhzSettings.rxsop"); value.Exists() && value.Value() != nil {
			data.TwoFourGhzSettingsRxsop = types.Int64Value(value.Int())
		} else {
			data.TwoFourGhzSettingsRxsop = types.Int64Null()
		}
		if value := res.Get("twoFourGhzSettings.validAutoChannels"); value.Exists() && value.Value() != nil {
			data.TwoFourGhzSettingsValidAutoChannels = helpers.GetInt64Set(value.Array())
		} else {
			data.TwoFourGhzSettingsValidAutoChannels = types.SetNull(types.Int64Type)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
