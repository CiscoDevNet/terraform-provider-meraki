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

type WirelessLocationScanningReceiver struct {
	Id             types.String `tfsdk:"id"`
	OrganizationId types.String `tfsdk:"organization_id"`
	SharedSecret   types.String `tfsdk:"shared_secret"`
	Url            types.String `tfsdk:"url"`
	Version        types.String `tfsdk:"version"`
	NetworkId      types.String `tfsdk:"network_id"`
	RadioType      types.String `tfsdk:"radio_type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessLocationScanningReceiver) getPath() string {
	return fmt.Sprintf("/organizations/%v/wireless/location/scanning/receivers", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessLocationScanningReceiver) toBody(ctx context.Context, state WirelessLocationScanningReceiver) string {
	body := ""
	if !data.SharedSecret.IsNull() {
		body, _ = sjson.Set(body, "sharedSecret", data.SharedSecret.ValueString())
	}
	if !data.Url.IsNull() {
		body, _ = sjson.Set(body, "url", data.Url.ValueString())
	}
	if !data.Version.IsNull() {
		body, _ = sjson.Set(body, "version", data.Version.ValueString())
	}
	if !data.NetworkId.IsNull() {
		body, _ = sjson.Set(body, "network.id", data.NetworkId.ValueString())
	}
	if !data.RadioType.IsNull() {
		body, _ = sjson.Set(body, "radio.type", data.RadioType.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessLocationScanningReceiver) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("url"); value.Exists() && value.Value() != nil {
		data.Url = types.StringValue(value.String())
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get("version"); value.Exists() && value.Value() != nil {
		data.Version = types.StringValue(value.String())
	} else {
		data.Version = types.StringNull()
	}
	if value := res.Get("network.id"); value.Exists() && value.Value() != nil {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
	}
	if value := res.Get("radio.type"); value.Exists() && value.Value() != nil {
		data.RadioType = types.StringValue(value.String())
	} else {
		data.RadioType = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessLocationScanningReceiver) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("url"); value.Exists() && !data.Url.IsNull() {
		data.Url = types.StringValue(value.String())
	} else {
		data.Url = types.StringNull()
	}
	if value := res.Get("version"); value.Exists() && !data.Version.IsNull() {
		data.Version = types.StringValue(value.String())
	} else {
		data.Version = types.StringNull()
	}
	if value := res.Get("network.id"); value.Exists() && !data.NetworkId.IsNull() {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
	}
	if value := res.Get("radio.type"); value.Exists() && !data.RadioType.IsNull() {
		data.RadioType = types.StringValue(value.String())
	} else {
		data.RadioType = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessLocationScanningReceiver) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data WirelessLocationScanningReceiver) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
