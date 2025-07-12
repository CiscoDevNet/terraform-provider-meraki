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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceApplianceRFProfiles struct {
	Id             types.String                       `tfsdk:"id"`
	OrganizationId types.String                       `tfsdk:"organization_id"`
	NetworkId      types.String                       `tfsdk:"network_id"`
	Items          []ResourceApplianceRFProfilesItems `tfsdk:"items"`
}

type ResourceApplianceRFProfilesItems struct {
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

func (data ResourceApplianceRFProfiles) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/rfProfiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceApplianceRFProfilesItems) toBody(ctx context.Context, state ResourceApplianceRFProfilesItems) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.FiveGhzSettingsAxEnabled.IsNull() {
		body, _ = sjson.Set(body, "fiveGhzSettings.axEnabled", data.FiveGhzSettingsAxEnabled.ValueBool())
	}
	if !data.FiveGhzSettingsMinBitrate.IsNull() {
		body, _ = sjson.Set(body, "fiveGhzSettings.minBitrate", data.FiveGhzSettingsMinBitrate.ValueInt64())
	}
	if !data.PerSsidSettings1BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:1.bandOperationMode", data.PerSsidSettings1BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings1BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:1.bandSteeringEnabled", data.PerSsidSettings1BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings2BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:2.bandOperationMode", data.PerSsidSettings2BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings2BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:2.bandSteeringEnabled", data.PerSsidSettings2BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings3BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:3.bandOperationMode", data.PerSsidSettings3BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings3BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:3.bandSteeringEnabled", data.PerSsidSettings3BandSteeringEnabled.ValueBool())
	}
	if !data.PerSsidSettings4BandOperationMode.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:4.bandOperationMode", data.PerSsidSettings4BandOperationMode.ValueString())
	}
	if !data.PerSsidSettings4BandSteeringEnabled.IsNull() {
		body, _ = sjson.Set(body, "perSsidSettings.:4.bandSteeringEnabled", data.PerSsidSettings4BandSteeringEnabled.ValueBool())
	}
	if !data.TwoFourGhzSettingsAxEnabled.IsNull() {
		body, _ = sjson.Set(body, "twoFourGhzSettings.axEnabled", data.TwoFourGhzSettingsAxEnabled.ValueBool())
	}
	if !data.TwoFourGhzSettingsMinBitrate.IsNull() {
		body, _ = sjson.Set(body, "twoFourGhzSettings.minBitrate", data.TwoFourGhzSettingsMinBitrate.ValueFloat64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceApplianceRFProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceApplianceRFProfilesItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceApplianceRFProfilesItems{}
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
func (data *ResourceApplianceRFProfiles) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("id").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("fiveGhzSettings.axEnabled"); value.Exists() && !data.FiveGhzSettingsAxEnabled.IsNull() {
			data.FiveGhzSettingsAxEnabled = types.BoolValue(value.Bool())
		} else {
			data.FiveGhzSettingsAxEnabled = types.BoolNull()
		}
		if value := res.Get("fiveGhzSettings.minBitrate"); value.Exists() && !data.FiveGhzSettingsMinBitrate.IsNull() {
			data.FiveGhzSettingsMinBitrate = types.Int64Value(value.Int())
		} else {
			data.FiveGhzSettingsMinBitrate = types.Int64Null()
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
		if value := res.Get("twoFourGhzSettings.axEnabled"); value.Exists() && !data.TwoFourGhzSettingsAxEnabled.IsNull() {
			data.TwoFourGhzSettingsAxEnabled = types.BoolValue(value.Bool())
		} else {
			data.TwoFourGhzSettingsAxEnabled = types.BoolNull()
		}
		if value := res.Get("twoFourGhzSettings.minBitrate"); value.Exists() && !data.TwoFourGhzSettingsMinBitrate.IsNull() {
			data.TwoFourGhzSettingsMinBitrate = types.Float64Value(value.Float())
		} else {
			data.TwoFourGhzSettingsMinBitrate = types.Float64Null()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceApplianceRFProfiles) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceApplianceRFProfiles) fromBodyImport(ctx context.Context, res meraki.Res) {
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
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("Import, removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceApplianceRFProfiles) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceApplianceRFProfiles) hasChanges(ctx context.Context, state *ResourceApplianceRFProfiles, id string) bool {
	hasChanges := false

	item := ResourceApplianceRFProfilesItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceApplianceRFProfilesItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	if !item.FiveGhzSettingsAxEnabled.Equal(stateItem.FiveGhzSettingsAxEnabled) {
		hasChanges = true
	}
	if !item.FiveGhzSettingsMinBitrate.Equal(stateItem.FiveGhzSettingsMinBitrate) {
		hasChanges = true
	}
	if !item.PerSsidSettings1BandOperationMode.Equal(stateItem.PerSsidSettings1BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings1BandSteeringEnabled.Equal(stateItem.PerSsidSettings1BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings2BandOperationMode.Equal(stateItem.PerSsidSettings2BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings2BandSteeringEnabled.Equal(stateItem.PerSsidSettings2BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings3BandOperationMode.Equal(stateItem.PerSsidSettings3BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings3BandSteeringEnabled.Equal(stateItem.PerSsidSettings3BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.PerSsidSettings4BandOperationMode.Equal(stateItem.PerSsidSettings4BandOperationMode) {
		hasChanges = true
	}
	if !item.PerSsidSettings4BandSteeringEnabled.Equal(stateItem.PerSsidSettings4BandSteeringEnabled) {
		hasChanges = true
	}
	if !item.TwoFourGhzSettingsAxEnabled.Equal(stateItem.TwoFourGhzSettingsAxEnabled) {
		hasChanges = true
	}
	if !item.TwoFourGhzSettingsMinBitrate.Equal(stateItem.TwoFourGhzSettingsMinBitrate) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
