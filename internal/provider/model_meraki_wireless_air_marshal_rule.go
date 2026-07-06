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

type WirelessAirMarshalRule struct {
	Id          types.String `tfsdk:"id"`
	NetworkId   types.String `tfsdk:"network_id"`
	Type        types.String `tfsdk:"type"`
	MatchString types.String `tfsdk:"match_string"`
	MatchType   types.String `tfsdk:"match_type"`
}

type WirelessAirMarshalRuleIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
	Id        types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessAirMarshalRule) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/airMarshal/rules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessAirMarshalRule) toBody(ctx context.Context, state WirelessAirMarshalRule) string {
	body := ""
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.MatchString.IsNull() {
		body, _ = sjson.Set(body, "match.string", data.MatchString.ValueString())
	}
	if !data.MatchType.IsNull() {
		body, _ = sjson.Set(body, "match.type", data.MatchType.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessAirMarshalRule) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("type"); value.Exists() && value.Value() != nil {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("match.string"); value.Exists() && value.Value() != nil {
		data.MatchString = types.StringValue(value.String())
	} else {
		data.MatchString = types.StringNull()
	}
	if value := res.Get("match.type"); value.Exists() && value.Value() != nil {
		data.MatchType = types.StringValue(value.String())
	} else {
		data.MatchType = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessAirMarshalRule) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("match.string"); value.Exists() && !data.MatchString.IsNull() {
		data.MatchString = types.StringValue(value.String())
	} else {
		data.MatchString = types.StringNull()
	}
	if value := res.Get("match.type"); value.Exists() && !data.MatchType.IsNull() {
		data.MatchType = types.StringValue(value.String())
	} else {
		data.MatchType = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessAirMarshalRule) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *WirelessAirMarshalRuleIdentity) toIdentity(ctx context.Context, plan *WirelessAirMarshalRule) {
	data.NetworkId = plan.NetworkId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *WirelessAirMarshalRule) fromIdentity(ctx context.Context, identity *WirelessAirMarshalRuleIdentity) {
	data.NetworkId = identity.NetworkId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data WirelessAirMarshalRule) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
