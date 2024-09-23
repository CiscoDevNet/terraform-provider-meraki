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

type WirelessEthernetPortProfileDefault struct {
	Id        types.String `tfsdk:"id"`
	NetworkId types.String `tfsdk:"network_id"`
	ProfileId types.String `tfsdk:"profile_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessEthernetPortProfileDefault) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ethernet/ports/profiles/setDefault", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

func (data WirelessEthernetPortProfileDefault) getGetPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ethernet/ports/profiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessEthernetPortProfileDefault) toBody(ctx context.Context, state WirelessEthernetPortProfileDefault) string {
	body := ""
	if !data.ProfileId.IsNull() {
		body, _ = sjson.Set(body, "profileId", data.ProfileId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessEthernetPortProfileDefault) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("profileId"); value.Exists() {
		data.ProfileId = types.StringValue(value.String())
	} else {
		data.ProfileId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessEthernetPortProfileDefault) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("profileId"); value.Exists() && !data.ProfileId.IsNull() {
		data.ProfileId = types.StringValue(value.String())
	} else {
		data.ProfileId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial