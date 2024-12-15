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

type WirelessDeviceElectronicShelfLabel struct {
	Id      types.String `tfsdk:"id"`
	Serial  types.String `tfsdk:"serial"`
	Channel types.String `tfsdk:"channel"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessDeviceElectronicShelfLabel) getPath() string {
	return fmt.Sprintf("/devices/%v/wireless/electronicShelfLabel", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessDeviceElectronicShelfLabel) toBody(ctx context.Context, state WirelessDeviceElectronicShelfLabel) string {
	body := ""
	if !data.Channel.IsNull() {
		body, _ = sjson.Set(body, "channel", data.Channel.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessDeviceElectronicShelfLabel) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("channel"); value.Exists() && value.Value() != nil {
		data.Channel = types.StringValue(value.String())
	} else {
		data.Channel = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessDeviceElectronicShelfLabel) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("channel"); value.Exists() && !data.Channel.IsNull() {
		data.Channel = types.StringValue(value.String())
	} else {
		data.Channel = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessDeviceElectronicShelfLabel) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns
