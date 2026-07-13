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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSourceApplianceRadioSettings - used only by the data source, always including every
// attribute (unlike model.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceApplianceRadioSettings struct {
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

func (data DataSourceApplianceRadioSettings) getPath() string {
	return fmt.Sprintf("/devices/%v/appliance/radio/settings", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceRadioSettings) fromBody(ctx context.Context, res meraki.Res) {
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
