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

type WirelessDeviceBluetoothSettings struct {
	Id     types.String `tfsdk:"id"`
	Serial types.String `tfsdk:"serial"`
	Major  types.Int64  `tfsdk:"major"`
	Minor  types.Int64  `tfsdk:"minor"`
	Uuid   types.String `tfsdk:"uuid"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessDeviceBluetoothSettings) getPath() string {
	return fmt.Sprintf("/devices/%v/wireless/bluetooth/settings", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessDeviceBluetoothSettings) toBody(ctx context.Context, state WirelessDeviceBluetoothSettings) string {
	body := ""
	if !data.Major.IsNull() {
		body, _ = sjson.Set(body, "major", data.Major.ValueInt64())
	}
	if !data.Minor.IsNull() {
		body, _ = sjson.Set(body, "minor", data.Minor.ValueInt64())
	}
	if !data.Uuid.IsNull() {
		body, _ = sjson.Set(body, "uuid", data.Uuid.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessDeviceBluetoothSettings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("major"); value.Exists() && value.Value() != nil {
		data.Major = types.Int64Value(value.Int())
	} else {
		data.Major = types.Int64Null()
	}
	if value := res.Get("minor"); value.Exists() && value.Value() != nil {
		data.Minor = types.Int64Value(value.Int())
	} else {
		data.Minor = types.Int64Null()
	}
	if value := res.Get("uuid"); value.Exists() && value.Value() != nil {
		data.Uuid = types.StringValue(value.String())
	} else {
		data.Uuid = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessDeviceBluetoothSettings) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("major"); value.Exists() && !data.Major.IsNull() {
		data.Major = types.Int64Value(value.Int())
	} else {
		data.Major = types.Int64Null()
	}
	if value := res.Get("minor"); value.Exists() && !data.Minor.IsNull() {
		data.Minor = types.Int64Value(value.Int())
	} else {
		data.Minor = types.Int64Null()
	}
	if value := res.Get("uuid"); value.Exists() && !data.Uuid.IsNull() {
		data.Uuid = types.StringValue(value.String())
	} else {
		data.Uuid = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessDeviceBluetoothSettings) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data WirelessDeviceBluetoothSettings) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
