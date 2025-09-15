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

type CameraDeviceWirelessProfiles struct {
	Id           types.String `tfsdk:"id"`
	Serial       types.String `tfsdk:"serial"`
	IdsBackup    types.String `tfsdk:"ids_backup"`
	IdsPrimary   types.String `tfsdk:"ids_primary"`
	IdsSecondary types.String `tfsdk:"ids_secondary"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CameraDeviceWirelessProfiles) getPath() string {
	return fmt.Sprintf("/devices/%v/camera/wirelessProfiles", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CameraDeviceWirelessProfiles) toBody(ctx context.Context, state CameraDeviceWirelessProfiles) string {
	body := ""
	if !data.IdsBackup.IsNull() {
		body, _ = sjson.Set(body, "ids.backup", data.IdsBackup.ValueString())
	}
	if !data.IdsPrimary.IsNull() {
		body, _ = sjson.Set(body, "ids.primary", data.IdsPrimary.ValueString())
	}
	if !data.IdsSecondary.IsNull() {
		body, _ = sjson.Set(body, "ids.secondary", data.IdsSecondary.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CameraDeviceWirelessProfiles) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("ids.backup"); value.Exists() && value.Value() != nil {
		data.IdsBackup = types.StringValue(value.String())
	} else {
		data.IdsBackup = types.StringNull()
	}
	if value := res.Get("ids.primary"); value.Exists() && value.Value() != nil {
		data.IdsPrimary = types.StringValue(value.String())
	} else {
		data.IdsPrimary = types.StringNull()
	}
	if value := res.Get("ids.secondary"); value.Exists() && value.Value() != nil {
		data.IdsSecondary = types.StringValue(value.String())
	} else {
		data.IdsSecondary = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *CameraDeviceWirelessProfiles) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("ids.backup"); value.Exists() && !data.IdsBackup.IsNull() {
		data.IdsBackup = types.StringValue(value.String())
	} else {
		data.IdsBackup = types.StringNull()
	}
	if value := res.Get("ids.primary"); value.Exists() && !data.IdsPrimary.IsNull() {
		data.IdsPrimary = types.StringValue(value.String())
	} else {
		data.IdsPrimary = types.StringNull()
	}
	if value := res.Get("ids.secondary"); value.Exists() && !data.IdsSecondary.IsNull() {
		data.IdsSecondary = types.StringValue(value.String())
	} else {
		data.IdsSecondary = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CameraDeviceWirelessProfiles) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data CameraDeviceWirelessProfiles) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
