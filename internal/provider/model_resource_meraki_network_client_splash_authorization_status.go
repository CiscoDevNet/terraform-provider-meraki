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

type NetworkClientSplashAuthorizationStatus struct {
	Id                  types.String `tfsdk:"id"`
	NetworkId           types.String `tfsdk:"network_id"`
	ClientId            types.String `tfsdk:"client_id"`
	Ssids0IsAuthorized  types.Bool   `tfsdk:"ssids_0_is_authorized"`
	Ssids1IsAuthorized  types.Bool   `tfsdk:"ssids_1_is_authorized"`
	Ssids10IsAuthorized types.Bool   `tfsdk:"ssids_10_is_authorized"`
	Ssids11IsAuthorized types.Bool   `tfsdk:"ssids_11_is_authorized"`
	Ssids12IsAuthorized types.Bool   `tfsdk:"ssids_12_is_authorized"`
	Ssids13IsAuthorized types.Bool   `tfsdk:"ssids_13_is_authorized"`
	Ssids14IsAuthorized types.Bool   `tfsdk:"ssids_14_is_authorized"`
	Ssids2IsAuthorized  types.Bool   `tfsdk:"ssids_2_is_authorized"`
	Ssids3IsAuthorized  types.Bool   `tfsdk:"ssids_3_is_authorized"`
	Ssids4IsAuthorized  types.Bool   `tfsdk:"ssids_4_is_authorized"`
	Ssids5IsAuthorized  types.Bool   `tfsdk:"ssids_5_is_authorized"`
	Ssids6IsAuthorized  types.Bool   `tfsdk:"ssids_6_is_authorized"`
	Ssids7IsAuthorized  types.Bool   `tfsdk:"ssids_7_is_authorized"`
	Ssids8IsAuthorized  types.Bool   `tfsdk:"ssids_8_is_authorized"`
	Ssids9IsAuthorized  types.Bool   `tfsdk:"ssids_9_is_authorized"`
}

type NetworkClientSplashAuthorizationStatusIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
	ClientId  types.String `tfsdk:"client_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkClientSplashAuthorizationStatus) getPath() string {
	return fmt.Sprintf("/networks/%v/clients/%v/splashAuthorizationStatus", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.ClientId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkClientSplashAuthorizationStatus) toBody(ctx context.Context, state NetworkClientSplashAuthorizationStatus) string {
	body := ""
	if !data.Ssids0IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:0.isAuthorized", data.Ssids0IsAuthorized.ValueBool())
	}
	if !data.Ssids1IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:1.isAuthorized", data.Ssids1IsAuthorized.ValueBool())
	}
	if !data.Ssids10IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:10.isAuthorized", data.Ssids10IsAuthorized.ValueBool())
	}
	if !data.Ssids11IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:11.isAuthorized", data.Ssids11IsAuthorized.ValueBool())
	}
	if !data.Ssids12IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:12.isAuthorized", data.Ssids12IsAuthorized.ValueBool())
	}
	if !data.Ssids13IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:13.isAuthorized", data.Ssids13IsAuthorized.ValueBool())
	}
	if !data.Ssids14IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:14.isAuthorized", data.Ssids14IsAuthorized.ValueBool())
	}
	if !data.Ssids2IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:2.isAuthorized", data.Ssids2IsAuthorized.ValueBool())
	}
	if !data.Ssids3IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:3.isAuthorized", data.Ssids3IsAuthorized.ValueBool())
	}
	if !data.Ssids4IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:4.isAuthorized", data.Ssids4IsAuthorized.ValueBool())
	}
	if !data.Ssids5IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:5.isAuthorized", data.Ssids5IsAuthorized.ValueBool())
	}
	if !data.Ssids6IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:6.isAuthorized", data.Ssids6IsAuthorized.ValueBool())
	}
	if !data.Ssids7IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:7.isAuthorized", data.Ssids7IsAuthorized.ValueBool())
	}
	if !data.Ssids8IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:8.isAuthorized", data.Ssids8IsAuthorized.ValueBool())
	}
	if !data.Ssids9IsAuthorized.IsNull() {
		body, _ = sjson.Set(body, "ssids.:9.isAuthorized", data.Ssids9IsAuthorized.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkClientSplashAuthorizationStatus) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("ssids.0.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids0IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids0IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.1.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids1IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids1IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.10.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids10IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids10IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.11.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids11IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids11IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.12.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids12IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids12IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.13.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids13IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids13IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.14.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids14IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids14IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.2.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids2IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids2IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.3.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids3IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids3IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.4.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids4IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids4IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.5.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids5IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids5IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.6.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids6IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids6IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.7.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids7IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids7IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.8.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids8IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids8IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.9.isAuthorized"); value.Exists() && value.Value() != nil {
		data.Ssids9IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids9IsAuthorized = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkClientSplashAuthorizationStatus) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("ssids.0.isAuthorized"); value.Exists() && !data.Ssids0IsAuthorized.IsNull() {
		data.Ssids0IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids0IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.1.isAuthorized"); value.Exists() && !data.Ssids1IsAuthorized.IsNull() {
		data.Ssids1IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids1IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.10.isAuthorized"); value.Exists() && !data.Ssids10IsAuthorized.IsNull() {
		data.Ssids10IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids10IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.11.isAuthorized"); value.Exists() && !data.Ssids11IsAuthorized.IsNull() {
		data.Ssids11IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids11IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.12.isAuthorized"); value.Exists() && !data.Ssids12IsAuthorized.IsNull() {
		data.Ssids12IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids12IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.13.isAuthorized"); value.Exists() && !data.Ssids13IsAuthorized.IsNull() {
		data.Ssids13IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids13IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.14.isAuthorized"); value.Exists() && !data.Ssids14IsAuthorized.IsNull() {
		data.Ssids14IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids14IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.2.isAuthorized"); value.Exists() && !data.Ssids2IsAuthorized.IsNull() {
		data.Ssids2IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids2IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.3.isAuthorized"); value.Exists() && !data.Ssids3IsAuthorized.IsNull() {
		data.Ssids3IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids3IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.4.isAuthorized"); value.Exists() && !data.Ssids4IsAuthorized.IsNull() {
		data.Ssids4IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids4IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.5.isAuthorized"); value.Exists() && !data.Ssids5IsAuthorized.IsNull() {
		data.Ssids5IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids5IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.6.isAuthorized"); value.Exists() && !data.Ssids6IsAuthorized.IsNull() {
		data.Ssids6IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids6IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.7.isAuthorized"); value.Exists() && !data.Ssids7IsAuthorized.IsNull() {
		data.Ssids7IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids7IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.8.isAuthorized"); value.Exists() && !data.Ssids8IsAuthorized.IsNull() {
		data.Ssids8IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids8IsAuthorized = types.BoolNull()
	}
	if value := res.Get("ssids.9.isAuthorized"); value.Exists() && !data.Ssids9IsAuthorized.IsNull() {
		data.Ssids9IsAuthorized = types.BoolValue(value.Bool())
	} else {
		data.Ssids9IsAuthorized = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkClientSplashAuthorizationStatus) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *NetworkClientSplashAuthorizationStatusIdentity) toIdentity(ctx context.Context, plan *NetworkClientSplashAuthorizationStatus) {
	data.NetworkId = plan.NetworkId
	data.ClientId = plan.ClientId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *NetworkClientSplashAuthorizationStatus) fromIdentity(ctx context.Context, identity *NetworkClientSplashAuthorizationStatusIdentity) {
	data.NetworkId = identity.NetworkId
	data.ClientId = identity.ClientId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkClientSplashAuthorizationStatus) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
