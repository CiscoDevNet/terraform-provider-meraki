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

type ResourceWirelessRFProfiles struct {
	Id             types.String                      `tfsdk:"id"`
	OrganizationId types.String                      `tfsdk:"organization_id"`
	NetworkId      types.String                      `tfsdk:"network_id"`
	Items          []ResourceWirelessRFProfilesItems `tfsdk:"items"`
}

type ResourceWirelessRFProfilesItems struct {
	Id                                   types.String                                  `tfsdk:"id"`
	BandSelectionType                    types.String                                  `tfsdk:"band_selection_type"`
	ClientBalancingEnabled               types.Bool                                    `tfsdk:"client_balancing_enabled"`
	MinBitrateType                       types.String                                  `tfsdk:"min_bitrate_type"`
	Name                                 types.String                                  `tfsdk:"name"`
	ApBandSettingsBandOperationMode      types.String                                  `tfsdk:"ap_band_settings_band_operation_mode"`
	ApBandSettingsBandSteeringEnabled    types.Bool                                    `tfsdk:"ap_band_settings_band_steering_enabled"`
	ApBandSettingsBandsEnabled           types.Set                                     `tfsdk:"ap_band_settings_bands_enabled"`
	FiveGhzSettingsChannelWidth          types.String                                  `tfsdk:"five_ghz_settings_channel_width"`
	FiveGhzSettingsMaxPower              types.Int64                                   `tfsdk:"five_ghz_settings_max_power"`
	FiveGhzSettingsMinBitrate            types.Int64                                   `tfsdk:"five_ghz_settings_min_bitrate"`
	FiveGhzSettingsMinPower              types.Int64                                   `tfsdk:"five_ghz_settings_min_power"`
	FiveGhzSettingsRxsop                 types.Int64                                   `tfsdk:"five_ghz_settings_rxsop"`
	FiveGhzSettingsValidAutoChannels     types.Set                                     `tfsdk:"five_ghz_settings_valid_auto_channels"`
	FlexRadiosByModel                    []ResourceWirelessRFProfilesFlexRadiosByModel `tfsdk:"flex_radios_by_model"`
	PerSsidSettings0BandOperationMode    types.String                                  `tfsdk:"per_ssid_settings_0_band_operation_mode"`
	PerSsidSettings0BandSteeringEnabled  types.Bool                                    `tfsdk:"per_ssid_settings_0_band_steering_enabled"`
	PerSsidSettings0MinBitrate           types.Float64                                 `tfsdk:"per_ssid_settings_0_min_bitrate"`
	PerSsidSettings0BandsEnabled         types.Set                                     `tfsdk:"per_ssid_settings_0_bands_enabled"`
	PerSsidSettings1BandOperationMode    types.String                                  `tfsdk:"per_ssid_settings_1_band_operation_mode"`
	PerSsidSettings1BandSteeringEnabled  types.Bool                                    `tfsdk:"per_ssid_settings_1_band_steering_enabled"`
	PerSsidSettings1MinBitrate           types.Float64                                 `tfsdk:"per_ssid_settings_1_min_bitrate"`
	PerSsidSettings1BandsEnabled         types.Set                                     `tfsdk:"per_ssid_settings_1_bands_enabled"`
	PerSsidSettings10BandOperationMode   types.String                                  `tfsdk:"per_ssid_settings_10_band_operation_mode"`
	PerSsidSettings10BandSteeringEnabled types.Bool                                    `tfsdk:"per_ssid_settings_10_band_steering_enabled"`
	PerSsidSettings10MinBitrate          types.Float64                                 `tfsdk:"per_ssid_settings_10_min_bitrate"`
	PerSsidSettings10BandsEnabled        types.Set                                     `tfsdk:"per_ssid_settings_10_bands_enabled"`
	PerSsidSettings11BandOperationMode   types.String                                  `tfsdk:"per_ssid_settings_11_band_operation_mode"`
	PerSsidSettings11BandSteeringEnabled types.Bool                                    `tfsdk:"per_ssid_settings_11_band_steering_enabled"`
	PerSsidSettings11MinBitrate          types.Float64                                 `tfsdk:"per_ssid_settings_11_min_bitrate"`
	PerSsidSettings11BandsEnabled        types.Set                                     `tfsdk:"per_ssid_settings_11_bands_enabled"`
	PerSsidSettings12BandOperationMode   types.String                                  `tfsdk:"per_ssid_settings_12_band_operation_mode"`
	PerSsidSettings12BandSteeringEnabled types.Bool                                    `tfsdk:"per_ssid_settings_12_band_steering_enabled"`
	PerSsidSettings12MinBitrate          types.Float64                                 `tfsdk:"per_ssid_settings_12_min_bitrate"`
	PerSsidSettings12BandsEnabled        types.Set                                     `tfsdk:"per_ssid_settings_12_bands_enabled"`
	PerSsidSettings13BandOperationMode   types.String                                  `tfsdk:"per_ssid_settings_13_band_operation_mode"`
	PerSsidSettings13BandSteeringEnabled types.Bool                                    `tfsdk:"per_ssid_settings_13_band_steering_enabled"`
	PerSsidSettings13MinBitrate          types.Float64                                 `tfsdk:"per_ssid_settings_13_min_bitrate"`
	PerSsidSettings13BandsEnabled        types.Set                                     `tfsdk:"per_ssid_settings_13_bands_enabled"`
	PerSsidSettings14BandOperationMode   types.String                                  `tfsdk:"per_ssid_settings_14_band_operation_mode"`
	PerSsidSettings14BandSteeringEnabled types.Bool                                    `tfsdk:"per_ssid_settings_14_band_steering_enabled"`
	PerSsidSettings14MinBitrate          types.Float64                                 `tfsdk:"per_ssid_settings_14_min_bitrate"`
	PerSsidSettings14BandsEnabled        types.Set                                     `tfsdk:"per_ssid_settings_14_bands_enabled"`
	PerSsidSettings2BandOperationMode    types.String                                  `tfsdk:"per_ssid_settings_2_band_operation_mode"`
	PerSsidSettings2BandSteeringEnabled  types.Bool                                    `tfsdk:"per_ssid_settings_2_band_steering_enabled"`
	PerSsidSettings2MinBitrate           types.Float64                                 `tfsdk:"per_ssid_settings_2_min_bitrate"`
	PerSsidSettings2BandsEnabled         types.Set                                     `tfsdk:"per_ssid_settings_2_bands_enabled"`
	PerSsidSettings3BandOperationMode    types.String                                  `tfsdk:"per_ssid_settings_3_band_operation_mode"`
	PerSsidSettings3BandSteeringEnabled  types.Bool                                    `tfsdk:"per_ssid_settings_3_band_steering_enabled"`
	PerSsidSettings3MinBitrate           types.Float64                                 `tfsdk:"per_ssid_settings_3_min_bitrate"`
	PerSsidSettings3BandsEnabled         types.Set                                     `tfsdk:"per_ssid_settings_3_bands_enabled"`
	PerSsidSettings4BandOperationMode    types.String                                  `tfsdk:"per_ssid_settings_4_band_operation_mode"`
	PerSsidSettings4BandSteeringEnabled  types.Bool                                    `tfsdk:"per_ssid_settings_4_band_steering_enabled"`
	PerSsidSettings4MinBitrate           types.Float64                                 `tfsdk:"per_ssid_settings_4_min_bitrate"`
	PerSsidSettings4BandsEnabled         types.Set                                     `tfsdk:"per_ssid_settings_4_bands_enabled"`
	PerSsidSettings5BandOperationMode    types.String                                  `tfsdk:"per_ssid_settings_5_band_operation_mode"`
	PerSsidSettings5BandSteeringEnabled  types.Bool                                    `tfsdk:"per_ssid_settings_5_band_steering_enabled"`
	PerSsidSettings5MinBitrate           types.Float64                                 `tfsdk:"per_ssid_settings_5_min_bitrate"`
	PerSsidSettings5BandsEnabled         types.Set                                     `tfsdk:"per_ssid_settings_5_bands_enabled"`
	PerSsidSettings6BandOperationMode    types.String                                  `tfsdk:"per_ssid_settings_6_band_operation_mode"`
	PerSsidSettings6BandSteeringEnabled  types.Bool                                    `tfsdk:"per_ssid_settings_6_band_steering_enabled"`
	PerSsidSettings6MinBitrate           types.Float64                                 `tfsdk:"per_ssid_settings_6_min_bitrate"`
	PerSsidSettings6BandsEnabled         types.Set                                     `tfsdk:"per_ssid_settings_6_bands_enabled"`
	PerSsidSettings7BandOperationMode    types.String                                  `tfsdk:"per_ssid_settings_7_band_operation_mode"`
	PerSsidSettings7BandSteeringEnabled  types.Bool                                    `tfsdk:"per_ssid_settings_7_band_steering_enabled"`
	PerSsidSettings7MinBitrate           types.Float64                                 `tfsdk:"per_ssid_settings_7_min_bitrate"`
	PerSsidSettings7BandsEnabled         types.Set                                     `tfsdk:"per_ssid_settings_7_bands_enabled"`
	PerSsidSettings8BandOperationMode    types.String                                  `tfsdk:"per_ssid_settings_8_band_operation_mode"`
	PerSsidSettings8BandSteeringEnabled  types.Bool                                    `tfsdk:"per_ssid_settings_8_band_steering_enabled"`
	PerSsidSettings8MinBitrate           types.Float64                                 `tfsdk:"per_ssid_settings_8_min_bitrate"`
	PerSsidSettings8BandsEnabled         types.Set                                     `tfsdk:"per_ssid_settings_8_bands_enabled"`
	PerSsidSettings9BandOperationMode    types.String                                  `tfsdk:"per_ssid_settings_9_band_operation_mode"`
	PerSsidSettings9BandSteeringEnabled  types.Bool                                    `tfsdk:"per_ssid_settings_9_band_steering_enabled"`
	PerSsidSettings9MinBitrate           types.Float64                                 `tfsdk:"per_ssid_settings_9_min_bitrate"`
	PerSsidSettings9BandsEnabled         types.Set                                     `tfsdk:"per_ssid_settings_9_bands_enabled"`
	SixGhzSettingsChannelWidth           types.String                                  `tfsdk:"six_ghz_settings_channel_width"`
	SixGhzSettingsMaxPower               types.Int64                                   `tfsdk:"six_ghz_settings_max_power"`
	SixGhzSettingsMinBitrate             types.Int64                                   `tfsdk:"six_ghz_settings_min_bitrate"`
	SixGhzSettingsMinPower               types.Int64                                   `tfsdk:"six_ghz_settings_min_power"`
	SixGhzSettingsRxsop                  types.Int64                                   `tfsdk:"six_ghz_settings_rxsop"`
	SixGhzSettingsValidAutoChannels      types.Set                                     `tfsdk:"six_ghz_settings_valid_auto_channels"`
	TransmissionEnabled                  types.Bool                                    `tfsdk:"transmission_enabled"`
	TwoFourGhzSettingsAxEnabled          types.Bool                                    `tfsdk:"two_four_ghz_settings_ax_enabled"`
	TwoFourGhzSettingsMaxPower           types.Int64                                   `tfsdk:"two_four_ghz_settings_max_power"`
	TwoFourGhzSettingsMinBitrate         types.Float64                                 `tfsdk:"two_four_ghz_settings_min_bitrate"`
	TwoFourGhzSettingsMinPower           types.Int64                                   `tfsdk:"two_four_ghz_settings_min_power"`
	TwoFourGhzSettingsRxsop              types.Int64                                   `tfsdk:"two_four_ghz_settings_rxsop"`
	TwoFourGhzSettingsValidAutoChannels  types.Set                                     `tfsdk:"two_four_ghz_settings_valid_auto_channels"`
}

type ResourceWirelessRFProfilesFlexRadiosByModel struct {
	Model types.String `tfsdk:"model"`
	Bands types.Set    `tfsdk:"bands"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceWirelessRFProfiles) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/rfProfiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceWirelessRFProfilesItems) toBody(ctx context.Context, state ResourceWirelessRFProfilesItems) string {
	body := ""
	if !data.BandSelectionType.IsNull() {
		body, _ = sjson.Set(body, "bandSelectionType", data.BandSelectionType.ValueString())
	}
	if !data.ClientBalancingEnabled.IsNull() {
		body, _ = sjson.Set(body, "clientBalancingEnabled", data.ClientBalancingEnabled.ValueBool())
	}
	if !data.MinBitrateType.IsNull() {
		body, _ = sjson.Set(body, "minBitrateType", data.MinBitrateType.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.ApBandSettingsBandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "apBandSettings.bandOperationMode", data.ApBandSettingsBandOperationMode.ValueString())
	}
	if !data.ApBandSettingsBandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "apBandSettings.bandSteeringEnabled", data.ApBandSettingsBandSteeringEnabled.ValueBool())
	}
	if !data.ApBandSettingsBandsEnabled.IsNull() {
		var values []string
		data.ApBandSettingsBandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "apBandSettings.bands.enabled", values)
	}
	if !data.FiveGhzSettingsChannelWidth.IsNull() {
		body, _ = sjson.Set(body, "fiveGhzSettings.channelWidth", data.FiveGhzSettingsChannelWidth.ValueString())
	}
	if !data.FiveGhzSettingsMaxPower.IsNull() {
		body, _ = sjson.Set(body, "fiveGhzSettings.maxPower", data.FiveGhzSettingsMaxPower.ValueInt64())
	}
	if !data.FiveGhzSettingsMinBitrate.IsNull() {
		body, _ = sjson.Set(body, "fiveGhzSettings.minBitrate", data.FiveGhzSettingsMinBitrate.ValueInt64())
	}
	if !data.FiveGhzSettingsMinPower.IsNull() {
		body, _ = sjson.Set(body, "fiveGhzSettings.minPower", data.FiveGhzSettingsMinPower.ValueInt64())
	}
	if !data.FiveGhzSettingsRxsop.IsNull() {
		body, _ = sjson.Set(body, "fiveGhzSettings.rxsop", data.FiveGhzSettingsRxsop.ValueInt64())
	}
	if !data.FiveGhzSettingsValidAutoChannels.IsNull() {
		var values []int64
		data.FiveGhzSettingsValidAutoChannels.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "fiveGhzSettings.validAutoChannels", values)
	}
	if len(data.FlexRadiosByModel) > 0 {
		body, _ = sjson.Set(body, "flexRadios.byModel", []interface{}{})
		for _, item := range data.FlexRadiosByModel {
			itemBody := ""
			if !item.Model.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "model", item.Model.ValueString())
			}
			if !item.Bands.IsNull() {
				var values []string
				item.Bands.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "bands", values)
			}
			body, _ = sjson.SetRaw(body, "flexRadios.byModel.-1", itemBody)
		}
	}
	if !data.PerSsidSettings0BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:0.bandOperationMode", data.PerSsidSettings0BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings0BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:0.bandSteeringEnabled", data.PerSsidSettings0BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings0MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:0.minBitrate", data.PerSsidSettings0MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings0BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings0BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:0.bands.enabled", values)
	}
	if !data.PerSsidSettings1BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:1.bandOperationMode", data.PerSsidSettings1BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings1BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:1.bandSteeringEnabled", data.PerSsidSettings1BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings1MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:1.minBitrate", data.PerSsidSettings1MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings1BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings1BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:1.bands.enabled", values)
	}
	if !data.PerSsidSettings10BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:10.bandOperationMode", data.PerSsidSettings10BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings10BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:10.bandSteeringEnabled", data.PerSsidSettings10BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings10MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:10.minBitrate", data.PerSsidSettings10MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings10BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings10BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:10.bands.enabled", values)
	}
	if !data.PerSsidSettings11BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:11.bandOperationMode", data.PerSsidSettings11BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings11BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:11.bandSteeringEnabled", data.PerSsidSettings11BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings11MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:11.minBitrate", data.PerSsidSettings11MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings11BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings11BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:11.bands.enabled", values)
	}
	if !data.PerSsidSettings12BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:12.bandOperationMode", data.PerSsidSettings12BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings12BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:12.bandSteeringEnabled", data.PerSsidSettings12BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings12MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:12.minBitrate", data.PerSsidSettings12MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings12BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings12BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:12.bands.enabled", values)
	}
	if !data.PerSsidSettings13BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:13.bandOperationMode", data.PerSsidSettings13BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings13BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:13.bandSteeringEnabled", data.PerSsidSettings13BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings13MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:13.minBitrate", data.PerSsidSettings13MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings13BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings13BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:13.bands.enabled", values)
	}
	if !data.PerSsidSettings14BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:14.bandOperationMode", data.PerSsidSettings14BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings14BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:14.bandSteeringEnabled", data.PerSsidSettings14BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings14MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:14.minBitrate", data.PerSsidSettings14MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings14BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings14BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:14.bands.enabled", values)
	}
	if !data.PerSsidSettings2BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:2.bandOperationMode", data.PerSsidSettings2BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings2BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:2.bandSteeringEnabled", data.PerSsidSettings2BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings2MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:2.minBitrate", data.PerSsidSettings2MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings2BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings2BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:2.bands.enabled", values)
	}
	if !data.PerSsidSettings3BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:3.bandOperationMode", data.PerSsidSettings3BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings3BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:3.bandSteeringEnabled", data.PerSsidSettings3BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings3MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:3.minBitrate", data.PerSsidSettings3MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings3BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings3BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:3.bands.enabled", values)
	}
	if !data.PerSsidSettings4BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:4.bandOperationMode", data.PerSsidSettings4BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings4BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:4.bandSteeringEnabled", data.PerSsidSettings4BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings4MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:4.minBitrate", data.PerSsidSettings4MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings4BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings4BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:4.bands.enabled", values)
	}
	if !data.PerSsidSettings5BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:5.bandOperationMode", data.PerSsidSettings5BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings5BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:5.bandSteeringEnabled", data.PerSsidSettings5BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings5MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:5.minBitrate", data.PerSsidSettings5MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings5BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings5BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:5.bands.enabled", values)
	}
	if !data.PerSsidSettings6BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:6.bandOperationMode", data.PerSsidSettings6BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings6BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:6.bandSteeringEnabled", data.PerSsidSettings6BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings6MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:6.minBitrate", data.PerSsidSettings6MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings6BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings6BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:6.bands.enabled", values)
	}
	if !data.PerSsidSettings7BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:7.bandOperationMode", data.PerSsidSettings7BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings7BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:7.bandSteeringEnabled", data.PerSsidSettings7BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings7MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:7.minBitrate", data.PerSsidSettings7MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings7BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings7BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:7.bands.enabled", values)
	}
	if !data.PerSsidSettings8BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:8.bandOperationMode", data.PerSsidSettings8BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings8BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:8.bandSteeringEnabled", data.PerSsidSettings8BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings8MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:8.minBitrate", data.PerSsidSettings8MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings8BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings8BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:8.bands.enabled", values)
	}
	if !data.PerSsidSettings9BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:9.bandOperationMode", data.PerSsidSettings9BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings9BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:9.bandSteeringEnabled", data.PerSsidSettings9BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings9MinBitrate.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:9.minBitrate", data.PerSsidSettings9MinBitrate.ValueFloat64())
	}
	if !data.PerSsidSettings9BandsEnabled.IsNull() {
		var values []string
		data.PerSsidSettings9BandsEnabled.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "perSsidSettings.:9.bands.enabled", values)
	}
	if !data.SixGhzSettingsChannelWidth.IsNull() {
		body, _ = sjson.Set(body, "sixGhzSettings.channelWidth", data.SixGhzSettingsChannelWidth.ValueString())
	}
	if !data.SixGhzSettingsMaxPower.IsNull() {
		body, _ = sjson.Set(body, "sixGhzSettings.maxPower", data.SixGhzSettingsMaxPower.ValueInt64())
	}
	if !data.SixGhzSettingsMinBitrate.IsNull() {
		body, _ = sjson.Set(body, "sixGhzSettings.minBitrate", data.SixGhzSettingsMinBitrate.ValueInt64())
	}
	if !data.SixGhzSettingsMinPower.IsNull() {
		body, _ = sjson.Set(body, "sixGhzSettings.minPower", data.SixGhzSettingsMinPower.ValueInt64())
	}
	if !data.SixGhzSettingsRxsop.IsNull() {
		body, _ = sjson.Set(body, "sixGhzSettings.rxsop", data.SixGhzSettingsRxsop.ValueInt64())
	}
	if !data.SixGhzSettingsValidAutoChannels.IsNull() {
		var values []int64
		data.SixGhzSettingsValidAutoChannels.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "sixGhzSettings.validAutoChannels", values)
	}
	if !data.TransmissionEnabled.IsNull() {
		body, _ = sjson.Set(body, "transmission.enabled", data.TransmissionEnabled.ValueBool())
	}
	if !data.TwoFourGhzSettingsAxEnabled.IsNull() {
		body, _ = sjson.Set(body, "twoFourGhzSettings.axEnabled", data.TwoFourGhzSettingsAxEnabled.ValueBool())
	}
	if !data.TwoFourGhzSettingsMaxPower.IsNull() {
		body, _ = sjson.Set(body, "twoFourGhzSettings.maxPower", data.TwoFourGhzSettingsMaxPower.ValueInt64())
	}
	if !data.TwoFourGhzSettingsMinBitrate.IsNull() {
		body, _ = sjson.Set(body, "twoFourGhzSettings.minBitrate", data.TwoFourGhzSettingsMinBitrate.ValueFloat64())
	}
	if !data.TwoFourGhzSettingsMinPower.IsNull() {
		body, _ = sjson.Set(body, "twoFourGhzSettings.minPower", data.TwoFourGhzSettingsMinPower.ValueInt64())
	}
	if !data.TwoFourGhzSettingsRxsop.IsNull() {
		body, _ = sjson.Set(body, "twoFourGhzSettings.rxsop", data.TwoFourGhzSettingsRxsop.ValueInt64())
	}
	if !data.TwoFourGhzSettingsValidAutoChannels.IsNull() {
		var values []int64
		data.TwoFourGhzSettingsValidAutoChannels.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "twoFourGhzSettings.validAutoChannels", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceWirelessRFProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceWirelessRFProfilesItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceWirelessRFProfilesItems{}
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
			data.FlexRadiosByModel = make([]ResourceWirelessRFProfilesFlexRadiosByModel, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceWirelessRFProfilesFlexRadiosByModel{}
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
	index := 0
	res.ForEach(func(k, res gjson.Result) bool {
		data.Items[index].Id = types.StringValue(res.Get("id").String())
		index++
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
func (data *ResourceWirelessRFProfiles) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
				if v.Get("id").String() == (*parent).Items[i].Id.ValueString() {
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
		if value := res.Get("bandSelectionType"); value.Exists() && !data.BandSelectionType.IsNull() {
			data.BandSelectionType = types.StringValue(value.String())
		} else {
			data.BandSelectionType = types.StringNull()
		}
		if value := res.Get("clientBalancingEnabled"); value.Exists() && !data.ClientBalancingEnabled.IsNull() {
			data.ClientBalancingEnabled = types.BoolValue(value.Bool())
		} else {
			data.ClientBalancingEnabled = types.BoolNull()
		}
		if value := res.Get("minBitrateType"); value.Exists() && !data.MinBitrateType.IsNull() {
			data.MinBitrateType = types.StringValue(value.String())
		} else {
			data.MinBitrateType = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("apBandSettings.bandOperationMode"); value.Exists() && !data.ApBandSettingsBandOperationMode.IsNull() {
			data.ApBandSettingsBandOperationMode = types.StringValue(value.String())
		} else {
			data.ApBandSettingsBandOperationMode = types.StringNull()
		}
		if value := res.Get("apBandSettings.bandSteeringEnabled"); value.Exists() && !data.ApBandSettingsBandSteeringEnabled.IsNull() {
			data.ApBandSettingsBandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.ApBandSettingsBandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("apBandSettings.bands.enabled"); value.Exists() && !data.ApBandSettingsBandsEnabled.IsNull() {
			data.ApBandSettingsBandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.ApBandSettingsBandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("fiveGhzSettings.channelWidth"); value.Exists() && !data.FiveGhzSettingsChannelWidth.IsNull() {
			data.FiveGhzSettingsChannelWidth = types.StringValue(value.String())
		} else {
			data.FiveGhzSettingsChannelWidth = types.StringNull()
		}
		if value := res.Get("fiveGhzSettings.maxPower"); value.Exists() && !data.FiveGhzSettingsMaxPower.IsNull() {
			data.FiveGhzSettingsMaxPower = types.Int64Value(value.Int())
		} else {
			data.FiveGhzSettingsMaxPower = types.Int64Null()
		}
		if value := res.Get("fiveGhzSettings.minBitrate"); value.Exists() && !data.FiveGhzSettingsMinBitrate.IsNull() {
			data.FiveGhzSettingsMinBitrate = types.Int64Value(value.Int())
		} else {
			data.FiveGhzSettingsMinBitrate = types.Int64Null()
		}
		if value := res.Get("fiveGhzSettings.minPower"); value.Exists() && !data.FiveGhzSettingsMinPower.IsNull() {
			data.FiveGhzSettingsMinPower = types.Int64Value(value.Int())
		} else {
			data.FiveGhzSettingsMinPower = types.Int64Null()
		}
		if value := res.Get("fiveGhzSettings.rxsop"); value.Exists() && !data.FiveGhzSettingsRxsop.IsNull() {
			data.FiveGhzSettingsRxsop = types.Int64Value(value.Int())
		} else {
			data.FiveGhzSettingsRxsop = types.Int64Null()
		}
		if value := res.Get("fiveGhzSettings.validAutoChannels"); value.Exists() && !data.FiveGhzSettingsValidAutoChannels.IsNull() {
			data.FiveGhzSettingsValidAutoChannels = helpers.GetInt64Set(value.Array())
		} else {
			data.FiveGhzSettingsValidAutoChannels = types.SetNull(types.Int64Type)
		}
		for i := 0; i < len(data.FlexRadiosByModel); i++ {
			keys := [...]string{"model"}
			keyValues := [...]string{data.FlexRadiosByModel[i].Model.ValueString()}

			parent := &data
			data := (*parent).FlexRadiosByModel[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("flexRadios.byModel").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing FlexRadiosByModel[%d] = %+v",
					i,
					(*parent).FlexRadiosByModel[i],
				))
				(*parent).FlexRadiosByModel = slices.Delete((*parent).FlexRadiosByModel, i, i+1)
				i--

				continue
			}
			if value := res.Get("model"); value.Exists() && !data.Model.IsNull() {
				data.Model = types.StringValue(value.String())
			} else {
				data.Model = types.StringNull()
			}
			if value := res.Get("bands"); value.Exists() && !data.Bands.IsNull() {
				data.Bands = helpers.GetStringSet(value.Array())
			} else {
				data.Bands = types.SetNull(types.StringType)
			}
			(*parent).FlexRadiosByModel[i] = data
		}
		if value := res.Get("perSsidSettings.0.bandOperationMode"); value.Exists() && !data.PerSsidSettings0BandOperationMode.IsNull() {
			data.PerSsidSettings0BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings0BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.0.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings0BandSteeringEnabled.IsNull() {
			data.PerSsidSettings0BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings0BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.0.minBitrate"); value.Exists() && !data.PerSsidSettings0MinBitrate.IsNull() {
			data.PerSsidSettings0MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings0MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.0.bands.enabled"); value.Exists() && !data.PerSsidSettings0BandsEnabled.IsNull() {
			data.PerSsidSettings0BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings0BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.1.bandOperationMode"); value.Exists() && !data.PerSsidSettings1BandOperationMode.IsNull() {
			data.PerSsidSettings1BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings1BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.1.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings1BandSteeringEnabled.IsNull() {
			data.PerSsidSettings1BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings1BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.1.minBitrate"); value.Exists() && !data.PerSsidSettings1MinBitrate.IsNull() {
			data.PerSsidSettings1MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings1MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.1.bands.enabled"); value.Exists() && !data.PerSsidSettings1BandsEnabled.IsNull() {
			data.PerSsidSettings1BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings1BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.10.bandOperationMode"); value.Exists() && !data.PerSsidSettings10BandOperationMode.IsNull() {
			data.PerSsidSettings10BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings10BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.10.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings10BandSteeringEnabled.IsNull() {
			data.PerSsidSettings10BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings10BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.10.minBitrate"); value.Exists() && !data.PerSsidSettings10MinBitrate.IsNull() {
			data.PerSsidSettings10MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings10MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.10.bands.enabled"); value.Exists() && !data.PerSsidSettings10BandsEnabled.IsNull() {
			data.PerSsidSettings10BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings10BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.11.bandOperationMode"); value.Exists() && !data.PerSsidSettings11BandOperationMode.IsNull() {
			data.PerSsidSettings11BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings11BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.11.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings11BandSteeringEnabled.IsNull() {
			data.PerSsidSettings11BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings11BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.11.minBitrate"); value.Exists() && !data.PerSsidSettings11MinBitrate.IsNull() {
			data.PerSsidSettings11MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings11MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.11.bands.enabled"); value.Exists() && !data.PerSsidSettings11BandsEnabled.IsNull() {
			data.PerSsidSettings11BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings11BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.12.bandOperationMode"); value.Exists() && !data.PerSsidSettings12BandOperationMode.IsNull() {
			data.PerSsidSettings12BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings12BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.12.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings12BandSteeringEnabled.IsNull() {
			data.PerSsidSettings12BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings12BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.12.minBitrate"); value.Exists() && !data.PerSsidSettings12MinBitrate.IsNull() {
			data.PerSsidSettings12MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings12MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.12.bands.enabled"); value.Exists() && !data.PerSsidSettings12BandsEnabled.IsNull() {
			data.PerSsidSettings12BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings12BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.13.bandOperationMode"); value.Exists() && !data.PerSsidSettings13BandOperationMode.IsNull() {
			data.PerSsidSettings13BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings13BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.13.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings13BandSteeringEnabled.IsNull() {
			data.PerSsidSettings13BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings13BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.13.minBitrate"); value.Exists() && !data.PerSsidSettings13MinBitrate.IsNull() {
			data.PerSsidSettings13MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings13MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.13.bands.enabled"); value.Exists() && !data.PerSsidSettings13BandsEnabled.IsNull() {
			data.PerSsidSettings13BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings13BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.14.bandOperationMode"); value.Exists() && !data.PerSsidSettings14BandOperationMode.IsNull() {
			data.PerSsidSettings14BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings14BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.14.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings14BandSteeringEnabled.IsNull() {
			data.PerSsidSettings14BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings14BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.14.minBitrate"); value.Exists() && !data.PerSsidSettings14MinBitrate.IsNull() {
			data.PerSsidSettings14MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings14MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.14.bands.enabled"); value.Exists() && !data.PerSsidSettings14BandsEnabled.IsNull() {
			data.PerSsidSettings14BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings14BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.2.bandOperationMode"); value.Exists() && !data.PerSsidSettings2BandOperationMode.IsNull() {
			data.PerSsidSettings2BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings2BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.2.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings2BandSteeringEnabled.IsNull() {
			data.PerSsidSettings2BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings2BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.2.minBitrate"); value.Exists() && !data.PerSsidSettings2MinBitrate.IsNull() {
			data.PerSsidSettings2MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings2MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.2.bands.enabled"); value.Exists() && !data.PerSsidSettings2BandsEnabled.IsNull() {
			data.PerSsidSettings2BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings2BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.3.bandOperationMode"); value.Exists() && !data.PerSsidSettings3BandOperationMode.IsNull() {
			data.PerSsidSettings3BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings3BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.3.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings3BandSteeringEnabled.IsNull() {
			data.PerSsidSettings3BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings3BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.3.minBitrate"); value.Exists() && !data.PerSsidSettings3MinBitrate.IsNull() {
			data.PerSsidSettings3MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings3MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.3.bands.enabled"); value.Exists() && !data.PerSsidSettings3BandsEnabled.IsNull() {
			data.PerSsidSettings3BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings3BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.4.bandOperationMode"); value.Exists() && !data.PerSsidSettings4BandOperationMode.IsNull() {
			data.PerSsidSettings4BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings4BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.4.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings4BandSteeringEnabled.IsNull() {
			data.PerSsidSettings4BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings4BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.4.minBitrate"); value.Exists() && !data.PerSsidSettings4MinBitrate.IsNull() {
			data.PerSsidSettings4MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings4MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.4.bands.enabled"); value.Exists() && !data.PerSsidSettings4BandsEnabled.IsNull() {
			data.PerSsidSettings4BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings4BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.5.bandOperationMode"); value.Exists() && !data.PerSsidSettings5BandOperationMode.IsNull() {
			data.PerSsidSettings5BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings5BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.5.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings5BandSteeringEnabled.IsNull() {
			data.PerSsidSettings5BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings5BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.5.minBitrate"); value.Exists() && !data.PerSsidSettings5MinBitrate.IsNull() {
			data.PerSsidSettings5MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings5MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.5.bands.enabled"); value.Exists() && !data.PerSsidSettings5BandsEnabled.IsNull() {
			data.PerSsidSettings5BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings5BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.6.bandOperationMode"); value.Exists() && !data.PerSsidSettings6BandOperationMode.IsNull() {
			data.PerSsidSettings6BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings6BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.6.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings6BandSteeringEnabled.IsNull() {
			data.PerSsidSettings6BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings6BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.6.minBitrate"); value.Exists() && !data.PerSsidSettings6MinBitrate.IsNull() {
			data.PerSsidSettings6MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings6MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.6.bands.enabled"); value.Exists() && !data.PerSsidSettings6BandsEnabled.IsNull() {
			data.PerSsidSettings6BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings6BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.7.bandOperationMode"); value.Exists() && !data.PerSsidSettings7BandOperationMode.IsNull() {
			data.PerSsidSettings7BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings7BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.7.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings7BandSteeringEnabled.IsNull() {
			data.PerSsidSettings7BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings7BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.7.minBitrate"); value.Exists() && !data.PerSsidSettings7MinBitrate.IsNull() {
			data.PerSsidSettings7MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings7MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.7.bands.enabled"); value.Exists() && !data.PerSsidSettings7BandsEnabled.IsNull() {
			data.PerSsidSettings7BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings7BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.8.bandOperationMode"); value.Exists() && !data.PerSsidSettings8BandOperationMode.IsNull() {
			data.PerSsidSettings8BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings8BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.8.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings8BandSteeringEnabled.IsNull() {
			data.PerSsidSettings8BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings8BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.8.minBitrate"); value.Exists() && !data.PerSsidSettings8MinBitrate.IsNull() {
			data.PerSsidSettings8MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings8MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.8.bands.enabled"); value.Exists() && !data.PerSsidSettings8BandsEnabled.IsNull() {
			data.PerSsidSettings8BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings8BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("perSsidSettings.9.bandOperationMode"); value.Exists() && !data.PerSsidSettings9BandOperationMode.IsNull() {
			data.PerSsidSettings9BandOperationMode = types.StringValue(value.String())
		} else {
			data.PerSsidSettings9BandOperationMode = types.StringNull()
		}
		if value := res.Get("perSsidSettings.9.bandSteeringEnabled"); value.Exists() && !data.PerSsidSettings9BandSteeringEnabled.IsNull() {
			data.PerSsidSettings9BandSteeringEnabled = types.BoolValue(value.Bool())
		} else {
			data.PerSsidSettings9BandSteeringEnabled = types.BoolNull()
		}
		if value := res.Get("perSsidSettings.9.minBitrate"); value.Exists() && !data.PerSsidSettings9MinBitrate.IsNull() {
			data.PerSsidSettings9MinBitrate = types.Float64Value(value.Float())
		} else {
			data.PerSsidSettings9MinBitrate = types.Float64Null()
		}
		if value := res.Get("perSsidSettings.9.bands.enabled"); value.Exists() && !data.PerSsidSettings9BandsEnabled.IsNull() {
			data.PerSsidSettings9BandsEnabled = helpers.GetStringSet(value.Array())
		} else {
			data.PerSsidSettings9BandsEnabled = types.SetNull(types.StringType)
		}
		if value := res.Get("sixGhzSettings.channelWidth"); value.Exists() && !data.SixGhzSettingsChannelWidth.IsNull() {
			data.SixGhzSettingsChannelWidth = types.StringValue(value.String())
		} else {
			data.SixGhzSettingsChannelWidth = types.StringNull()
		}
		if value := res.Get("sixGhzSettings.maxPower"); value.Exists() && !data.SixGhzSettingsMaxPower.IsNull() {
			data.SixGhzSettingsMaxPower = types.Int64Value(value.Int())
		} else {
			data.SixGhzSettingsMaxPower = types.Int64Null()
		}
		if value := res.Get("sixGhzSettings.minBitrate"); value.Exists() && !data.SixGhzSettingsMinBitrate.IsNull() {
			data.SixGhzSettingsMinBitrate = types.Int64Value(value.Int())
		} else {
			data.SixGhzSettingsMinBitrate = types.Int64Null()
		}
		if value := res.Get("sixGhzSettings.minPower"); value.Exists() && !data.SixGhzSettingsMinPower.IsNull() {
			data.SixGhzSettingsMinPower = types.Int64Value(value.Int())
		} else {
			data.SixGhzSettingsMinPower = types.Int64Null()
		}
		if value := res.Get("sixGhzSettings.rxsop"); value.Exists() && !data.SixGhzSettingsRxsop.IsNull() {
			data.SixGhzSettingsRxsop = types.Int64Value(value.Int())
		} else {
			data.SixGhzSettingsRxsop = types.Int64Null()
		}
		if value := res.Get("sixGhzSettings.validAutoChannels"); value.Exists() && !data.SixGhzSettingsValidAutoChannels.IsNull() {
			data.SixGhzSettingsValidAutoChannels = helpers.GetInt64Set(value.Array())
		} else {
			data.SixGhzSettingsValidAutoChannels = types.SetNull(types.Int64Type)
		}
		if value := res.Get("transmission.enabled"); value.Exists() && !data.TransmissionEnabled.IsNull() {
			data.TransmissionEnabled = types.BoolValue(value.Bool())
		} else {
			data.TransmissionEnabled = types.BoolNull()
		}
		if value := res.Get("twoFourGhzSettings.axEnabled"); value.Exists() && !data.TwoFourGhzSettingsAxEnabled.IsNull() {
			data.TwoFourGhzSettingsAxEnabled = types.BoolValue(value.Bool())
		} else {
			data.TwoFourGhzSettingsAxEnabled = types.BoolNull()
		}
		if value := res.Get("twoFourGhzSettings.maxPower"); value.Exists() && !data.TwoFourGhzSettingsMaxPower.IsNull() {
			data.TwoFourGhzSettingsMaxPower = types.Int64Value(value.Int())
		} else {
			data.TwoFourGhzSettingsMaxPower = types.Int64Null()
		}
		if value := res.Get("twoFourGhzSettings.minBitrate"); value.Exists() && !data.TwoFourGhzSettingsMinBitrate.IsNull() {
			data.TwoFourGhzSettingsMinBitrate = types.Float64Value(value.Float())
		} else {
			data.TwoFourGhzSettingsMinBitrate = types.Float64Null()
		}
		if value := res.Get("twoFourGhzSettings.minPower"); value.Exists() && !data.TwoFourGhzSettingsMinPower.IsNull() {
			data.TwoFourGhzSettingsMinPower = types.Int64Value(value.Int())
		} else {
			data.TwoFourGhzSettingsMinPower = types.Int64Null()
		}
		if value := res.Get("twoFourGhzSettings.rxsop"); value.Exists() && !data.TwoFourGhzSettingsRxsop.IsNull() {
			data.TwoFourGhzSettingsRxsop = types.Int64Value(value.Int())
		} else {
			data.TwoFourGhzSettingsRxsop = types.Int64Null()
		}
		if value := res.Get("twoFourGhzSettings.validAutoChannels"); value.Exists() && !data.TwoFourGhzSettingsValidAutoChannels.IsNull() {
			data.TwoFourGhzSettingsValidAutoChannels = helpers.GetInt64Set(value.Array())
		} else {
			data.TwoFourGhzSettingsValidAutoChannels = types.SetNull(types.Int64Type)
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyPartial(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceWirelessRFProfiles) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceWirelessRFProfiles) fromBodyImport(ctx context.Context, res meraki.Res) {
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
				if v.Get("id").String() == (*parent).Items[i].Id.ValueString() {
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
		if value := res.Get("apBandSettings.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("fiveGhzSettings.validAutoChannels"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.FiveGhzSettingsValidAutoChannels = helpers.GetInt64Set(value.Array())
		} else {
			data.FiveGhzSettingsValidAutoChannels = types.SetNull(types.Int64Type)
		}
		if value := res.Get("flexRadios.byModel"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.FlexRadiosByModel = make([]ResourceWirelessRFProfilesFlexRadiosByModel, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceWirelessRFProfilesFlexRadiosByModel{}
				if value := res.Get("model"); value.Exists() && value.Value() != nil {
					data.Model = types.StringValue(value.String())
				} else {
					data.Model = types.StringNull()
				}
				if value := res.Get("bands"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.0.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.1.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.10.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.11.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.12.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.13.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.14.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.2.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.3.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.4.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.5.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.6.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.7.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.8.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("perSsidSettings.9.bands.enabled"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("sixGhzSettings.validAutoChannels"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
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
		if value := res.Get("twoFourGhzSettings.validAutoChannels"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.TwoFourGhzSettingsValidAutoChannels = helpers.GetInt64Set(value.Array())
		} else {
			data.TwoFourGhzSettingsValidAutoChannels = types.SetNull(types.Int64Type)
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyImport(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceWirelessRFProfiles) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceWirelessRFProfiles) hasChanges(ctx context.Context, state *ResourceWirelessRFProfiles, id string) bool {
	hasChanges := false

	item := ResourceWirelessRFProfilesItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceWirelessRFProfilesItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.BandSelectionType.Equal(stateItem.BandSelectionType) {
		hasChanges = true
	}
	if !item.ClientBalancingEnabled.Equal(stateItem.ClientBalancingEnabled) {
		hasChanges = true
	}
	if !item.MinBitrateType.Equal(stateItem.MinBitrateType) {
		hasChanges = true
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	if !item.ApBandSettingsBandOperationMode.Equal(stateItem.ApBandSettingsBandOperationMode) {
		hasChanges = true
	}
	if !item.ApBandSettingsBandSteeringEnabled.Equal(stateItem.ApBandSettingsBandSteeringEnabled) {
		hasChanges = true
	}
	if !item.ApBandSettingsBandsEnabled.Equal(stateItem.ApBandSettingsBandsEnabled) {
		hasChanges = true
	}
	if !item.FiveGhzSettingsChannelWidth.Equal(stateItem.FiveGhzSettingsChannelWidth) {
		hasChanges = true
	}
	if !item.FiveGhzSettingsMaxPower.Equal(stateItem.FiveGhzSettingsMaxPower) {
		hasChanges = true
	}
	if !item.FiveGhzSettingsMinBitrate.Equal(stateItem.FiveGhzSettingsMinBitrate) {
		hasChanges = true
	}
	if !item.FiveGhzSettingsMinPower.Equal(stateItem.FiveGhzSettingsMinPower) {
		hasChanges = true
	}
	if !item.FiveGhzSettingsRxsop.Equal(stateItem.FiveGhzSettingsRxsop) {
		hasChanges = true
	}
	if !item.FiveGhzSettingsValidAutoChannels.Equal(stateItem.FiveGhzSettingsValidAutoChannels) {
		hasChanges = true
	}
	if len(item.FlexRadiosByModel) != len(stateItem.FlexRadiosByModel) {
		hasChanges = true
	} else {
		for i := range item.FlexRadiosByModel {
			if !item.FlexRadiosByModel[i].Model.Equal(stateItem.FlexRadiosByModel[i].Model) {
				hasChanges = true
			}
			if !item.FlexRadiosByModel[i].Bands.Equal(stateItem.FlexRadiosByModel[i].Bands) {
				hasChanges = true
			}
		}
	}
	if !item.PerSsidSettings0BandOperationMode.Equal(stateItem.PerSsidSettings0BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings0BandSteeringEnabled.Equal(stateItem.PerSsidSettings0BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings0MinBitrate.Equal(stateItem.PerSsidSettings0MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings0BandsEnabled.Equal(stateItem.PerSsidSettings0BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings1BandOperationMode.Equal(stateItem.PerSsidSettings1BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings1BandSteeringEnabled.Equal(stateItem.PerSsidSettings1BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings1MinBitrate.Equal(stateItem.PerSsidSettings1MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings1BandsEnabled.Equal(stateItem.PerSsidSettings1BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings10BandOperationMode.Equal(stateItem.PerSsidSettings10BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings10BandSteeringEnabled.Equal(stateItem.PerSsidSettings10BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings10MinBitrate.Equal(stateItem.PerSsidSettings10MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings10BandsEnabled.Equal(stateItem.PerSsidSettings10BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings11BandOperationMode.Equal(stateItem.PerSsidSettings11BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings11BandSteeringEnabled.Equal(stateItem.PerSsidSettings11BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings11MinBitrate.Equal(stateItem.PerSsidSettings11MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings11BandsEnabled.Equal(stateItem.PerSsidSettings11BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings12BandOperationMode.Equal(stateItem.PerSsidSettings12BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings12BandSteeringEnabled.Equal(stateItem.PerSsidSettings12BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings12MinBitrate.Equal(stateItem.PerSsidSettings12MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings12BandsEnabled.Equal(stateItem.PerSsidSettings12BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings13BandOperationMode.Equal(stateItem.PerSsidSettings13BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings13BandSteeringEnabled.Equal(stateItem.PerSsidSettings13BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings13MinBitrate.Equal(stateItem.PerSsidSettings13MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings13BandsEnabled.Equal(stateItem.PerSsidSettings13BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings14BandOperationMode.Equal(stateItem.PerSsidSettings14BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings14BandSteeringEnabled.Equal(stateItem.PerSsidSettings14BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings14MinBitrate.Equal(stateItem.PerSsidSettings14MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings14BandsEnabled.Equal(stateItem.PerSsidSettings14BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings2BandOperationMode.Equal(stateItem.PerSsidSettings2BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings2BandSteeringEnabled.Equal(stateItem.PerSsidSettings2BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings2MinBitrate.Equal(stateItem.PerSsidSettings2MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings2BandsEnabled.Equal(stateItem.PerSsidSettings2BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings3BandOperationMode.Equal(stateItem.PerSsidSettings3BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings3BandSteeringEnabled.Equal(stateItem.PerSsidSettings3BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings3MinBitrate.Equal(stateItem.PerSsidSettings3MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings3BandsEnabled.Equal(stateItem.PerSsidSettings3BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings4BandOperationMode.Equal(stateItem.PerSsidSettings4BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings4BandSteeringEnabled.Equal(stateItem.PerSsidSettings4BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings4MinBitrate.Equal(stateItem.PerSsidSettings4MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings4BandsEnabled.Equal(stateItem.PerSsidSettings4BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings5BandOperationMode.Equal(stateItem.PerSsidSettings5BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings5BandSteeringEnabled.Equal(stateItem.PerSsidSettings5BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings5MinBitrate.Equal(stateItem.PerSsidSettings5MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings5BandsEnabled.Equal(stateItem.PerSsidSettings5BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings6BandOperationMode.Equal(stateItem.PerSsidSettings6BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings6BandSteeringEnabled.Equal(stateItem.PerSsidSettings6BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings6MinBitrate.Equal(stateItem.PerSsidSettings6MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings6BandsEnabled.Equal(stateItem.PerSsidSettings6BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings7BandOperationMode.Equal(stateItem.PerSsidSettings7BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings7BandSteeringEnabled.Equal(stateItem.PerSsidSettings7BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings7MinBitrate.Equal(stateItem.PerSsidSettings7MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings7BandsEnabled.Equal(stateItem.PerSsidSettings7BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings8BandOperationMode.Equal(stateItem.PerSsidSettings8BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings8BandSteeringEnabled.Equal(stateItem.PerSsidSettings8BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings8MinBitrate.Equal(stateItem.PerSsidSettings8MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings8BandsEnabled.Equal(stateItem.PerSsidSettings8BandsEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings9BandOperationMode.Equal(stateItem.PerSsidSettings9BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings9BandSteeringEnabled.Equal(stateItem.PerSsidSettings9BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings9MinBitrate.Equal(stateItem.PerSsidSettings9MinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings9BandsEnabled.Equal(stateItem.PerSsidSettings9BandsEnabled) {
		hasChanges = true
	}
	if !item.SixGhzSettingsChannelWidth.Equal(stateItem.SixGhzSettingsChannelWidth) {
		hasChanges = true
	}
	if !item.SixGhzSettingsMaxPower.Equal(stateItem.SixGhzSettingsMaxPower) {
		hasChanges = true
	}
	if !item.SixGhzSettingsMinBitrate.Equal(stateItem.SixGhzSettingsMinBitrate) {
		hasChanges = true
	}
	if !item.SixGhzSettingsMinPower.Equal(stateItem.SixGhzSettingsMinPower) {
		hasChanges = true
	}
	if !item.SixGhzSettingsRxsop.Equal(stateItem.SixGhzSettingsRxsop) {
		hasChanges = true
	}
	if !item.SixGhzSettingsValidAutoChannels.Equal(stateItem.SixGhzSettingsValidAutoChannels) {
		hasChanges = true
	}
	if !item.TransmissionEnabled.Equal(stateItem.TransmissionEnabled) {
		hasChanges = true
	}
	if !item.TwoFourGhzSettingsAxEnabled.Equal(stateItem.TwoFourGhzSettingsAxEnabled) {
		hasChanges = true
	}
	if !item.TwoFourGhzSettingsMaxPower.Equal(stateItem.TwoFourGhzSettingsMaxPower) {
		hasChanges = true
	}
	if !item.TwoFourGhzSettingsMinBitrate.Equal(stateItem.TwoFourGhzSettingsMinBitrate) {
		hasChanges = true
	}
	if !item.TwoFourGhzSettingsMinPower.Equal(stateItem.TwoFourGhzSettingsMinPower) {
		hasChanges = true
	}
	if !item.TwoFourGhzSettingsRxsop.Equal(stateItem.TwoFourGhzSettingsRxsop) {
		hasChanges = true
	}
	if !item.TwoFourGhzSettingsValidAutoChannels.Equal(stateItem.TwoFourGhzSettingsValidAutoChannels) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
