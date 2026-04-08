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

type CameraWirelessProfile struct {
	Id                 types.String `tfsdk:"id"`
	NetworkId          types.String `tfsdk:"network_id"`
	Name               types.String `tfsdk:"name"`
	IdentityPassword   types.String `tfsdk:"identity_password"`
	IdentityUsername   types.String `tfsdk:"identity_username"`
	SsidAuthMode       types.String `tfsdk:"ssid_auth_mode"`
	SsidEncryptionMode types.String `tfsdk:"ssid_encryption_mode"`
	SsidName           types.String `tfsdk:"ssid_name"`
	SsidPsk            types.String `tfsdk:"ssid_psk"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CameraWirelessProfile) getPath() string {
	return fmt.Sprintf("/networks/%v/camera/wirelessProfiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CameraWirelessProfile) toBody(ctx context.Context, state CameraWirelessProfile) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.IdentityPassword.IsNull() {
		body, _ = sjson.Set(body, "identity.password", data.IdentityPassword.ValueString())
	}
	if !data.IdentityUsername.IsNull() {
		body, _ = sjson.Set(body, "identity.username", data.IdentityUsername.ValueString())
	}
	if !data.SsidAuthMode.IsNull() {
		body, _ = sjson.Set(body, "ssid.authMode", data.SsidAuthMode.ValueString())
	}
	if !data.SsidEncryptionMode.IsNull() {
		body, _ = sjson.Set(body, "ssid.encryptionMode", data.SsidEncryptionMode.ValueString())
	}
	if !data.SsidName.IsNull() {
		body, _ = sjson.Set(body, "ssid.name", data.SsidName.ValueString())
	}
	if !data.SsidPsk.IsNull() {
		body, _ = sjson.Set(body, "ssid.psk", data.SsidPsk.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CameraWirelessProfile) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("identity.password"); value.Exists() && value.Value() != nil {
		data.IdentityPassword = types.StringValue(value.String())
	} else {
		data.IdentityPassword = types.StringNull()
	}
	if value := res.Get("identity.username"); value.Exists() && value.Value() != nil {
		data.IdentityUsername = types.StringValue(value.String())
	} else {
		data.IdentityUsername = types.StringNull()
	}
	if value := res.Get("ssid.authMode"); value.Exists() && value.Value() != nil {
		data.SsidAuthMode = types.StringValue(value.String())
	} else {
		data.SsidAuthMode = types.StringNull()
	}
	if value := res.Get("ssid.encryptionMode"); value.Exists() && value.Value() != nil {
		data.SsidEncryptionMode = types.StringValue(value.String())
	} else {
		data.SsidEncryptionMode = types.StringNull()
	}
	if value := res.Get("ssid.name"); value.Exists() && value.Value() != nil {
		data.SsidName = types.StringValue(value.String())
	} else {
		data.SsidName = types.StringNull()
	}
	if value := res.Get("ssid.psk"); value.Exists() && value.Value() != nil {
		data.SsidPsk = types.StringValue(value.String())
	} else {
		data.SsidPsk = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *CameraWirelessProfile) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("identity.password"); value.Exists() && !data.IdentityPassword.IsNull() {
		data.IdentityPassword = types.StringValue(value.String())
	} else {
		data.IdentityPassword = types.StringNull()
	}
	if value := res.Get("identity.username"); value.Exists() && !data.IdentityUsername.IsNull() {
		data.IdentityUsername = types.StringValue(value.String())
	} else {
		data.IdentityUsername = types.StringNull()
	}
	if value := res.Get("ssid.authMode"); value.Exists() && !data.SsidAuthMode.IsNull() {
		data.SsidAuthMode = types.StringValue(value.String())
	} else {
		data.SsidAuthMode = types.StringNull()
	}
	if value := res.Get("ssid.encryptionMode"); value.Exists() && !data.SsidEncryptionMode.IsNull() {
		data.SsidEncryptionMode = types.StringValue(value.String())
	} else {
		data.SsidEncryptionMode = types.StringNull()
	}
	if value := res.Get("ssid.name"); value.Exists() && !data.SsidName.IsNull() {
		data.SsidName = types.StringValue(value.String())
	} else {
		data.SsidName = types.StringNull()
	}
	if value := res.Get("ssid.psk"); value.Exists() && !data.SsidPsk.IsNull() {
		data.SsidPsk = types.StringValue(value.String())
	} else {
		data.SsidPsk = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CameraWirelessProfile) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data CameraWirelessProfile) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
