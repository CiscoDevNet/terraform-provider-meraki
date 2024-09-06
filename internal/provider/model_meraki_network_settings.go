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

type NetworkSettings struct {
	Id                                    types.String `tfsdk:"id"`
	NetworkId                             types.String `tfsdk:"network_id"`
	LocalStatusPageEnabled                types.Bool   `tfsdk:"local_status_page_enabled"`
	RemoteStatusPageEnabled               types.Bool   `tfsdk:"remote_status_page_enabled"`
	LocalStatusPageAuthenticationEnabled  types.Bool   `tfsdk:"local_status_page_authentication_enabled"`
	LocalStatusPageAuthenticationPassword types.String `tfsdk:"local_status_page_authentication_password"`
	NamedVlansEnabled                     types.Bool   `tfsdk:"named_vlans_enabled"`
	SecurePortEnabled                     types.Bool   `tfsdk:"secure_port_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/settings", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkSettings) toBody(ctx context.Context, state NetworkSettings) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.LocalStatusPageEnabled.IsNull() {
		body, _ = sjson.Set(body, "localStatusPageEnabled", data.LocalStatusPageEnabled.ValueBool())
	}
	if !data.RemoteStatusPageEnabled.IsNull() {
		body, _ = sjson.Set(body, "remoteStatusPageEnabled", data.RemoteStatusPageEnabled.ValueBool())
	}
	if !data.LocalStatusPageAuthenticationEnabled.IsNull() {
		body, _ = sjson.Set(body, "localStatusPage.authentication.enabled", data.LocalStatusPageAuthenticationEnabled.ValueBool())
	}
	if !data.LocalStatusPageAuthenticationPassword.IsNull() {
		body, _ = sjson.Set(body, "localStatusPage.authentication.password", data.LocalStatusPageAuthenticationPassword.ValueString())
	}
	if !data.NamedVlansEnabled.IsNull() {
		body, _ = sjson.Set(body, "namedVlans.enabled", data.NamedVlansEnabled.ValueBool())
	}
	if !data.SecurePortEnabled.IsNull() {
		body, _ = sjson.Set(body, "securePort.enabled", data.SecurePortEnabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkSettings) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("localStatusPageEnabled"); value.Exists() {
		data.LocalStatusPageEnabled = types.BoolValue(value.Bool())
	} else {
		data.LocalStatusPageEnabled = types.BoolNull()
	}
	if value := res.Get("remoteStatusPageEnabled"); value.Exists() {
		data.RemoteStatusPageEnabled = types.BoolValue(value.Bool())
	} else {
		data.RemoteStatusPageEnabled = types.BoolNull()
	}
	if value := res.Get("localStatusPage.authentication.enabled"); value.Exists() {
		data.LocalStatusPageAuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.LocalStatusPageAuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("namedVlans.enabled"); value.Exists() {
		data.NamedVlansEnabled = types.BoolValue(value.Bool())
	} else {
		data.NamedVlansEnabled = types.BoolNull()
	}
	if value := res.Get("securePort.enabled"); value.Exists() {
		data.SecurePortEnabled = types.BoolValue(value.Bool())
	} else {
		data.SecurePortEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkSettings) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("localStatusPageEnabled"); value.Exists() && !data.LocalStatusPageEnabled.IsNull() {
		data.LocalStatusPageEnabled = types.BoolValue(value.Bool())
	} else {
		data.LocalStatusPageEnabled = types.BoolNull()
	}
	if value := res.Get("remoteStatusPageEnabled"); value.Exists() && !data.RemoteStatusPageEnabled.IsNull() {
		data.RemoteStatusPageEnabled = types.BoolValue(value.Bool())
	} else {
		data.RemoteStatusPageEnabled = types.BoolNull()
	}
	if value := res.Get("localStatusPage.authentication.enabled"); value.Exists() && !data.LocalStatusPageAuthenticationEnabled.IsNull() {
		data.LocalStatusPageAuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.LocalStatusPageAuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("namedVlans.enabled"); value.Exists() && !data.NamedVlansEnabled.IsNull() {
		data.NamedVlansEnabled = types.BoolValue(value.Bool())
	} else {
		data.NamedVlansEnabled = types.BoolNull()
	}
	if value := res.Get("securePort.enabled"); value.Exists() && !data.SecurePortEnabled.IsNull() {
		data.SecurePortEnabled = types.BoolValue(value.Bool())
	} else {
		data.SecurePortEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkSettings) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
