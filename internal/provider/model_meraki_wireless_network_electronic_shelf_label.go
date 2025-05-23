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

type WirelessNetworkElectronicShelfLabel struct {
	Id        types.String `tfsdk:"id"`
	NetworkId types.String `tfsdk:"network_id"`
	Enabled   types.Bool   `tfsdk:"enabled"`
	Hostname  types.String `tfsdk:"hostname"`
	Mode      types.String `tfsdk:"mode"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessNetworkElectronicShelfLabel) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/electronicShelfLabel", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessNetworkElectronicShelfLabel) toBody(ctx context.Context, state WirelessNetworkElectronicShelfLabel) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.Hostname.IsNull() {
		body, _ = sjson.Set(body, "hostname", data.Hostname.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessNetworkElectronicShelfLabel) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("hostname"); value.Exists() && value.Value() != nil {
		data.Hostname = types.StringValue(value.String())
	} else {
		data.Hostname = types.StringNull()
	}
	if value := res.Get("mode"); value.Exists() && value.Value() != nil {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessNetworkElectronicShelfLabel) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("hostname"); value.Exists() && !data.Hostname.IsNull() {
		data.Hostname = types.StringValue(value.String())
	} else {
		data.Hostname = types.StringNull()
	}
	if value := res.Get("mode"); value.Exists() && !data.Mode.IsNull() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessNetworkElectronicShelfLabel) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data WirelessNetworkElectronicShelfLabel) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
