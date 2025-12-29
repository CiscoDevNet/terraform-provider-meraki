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

type ApplianceVLANsSettings struct {
	Id           types.String `tfsdk:"id"`
	NetworkId    types.String `tfsdk:"network_id"`
	VlansEnabled types.Bool   `tfsdk:"vlans_enabled"`
}

type ApplianceVLANsSettingsIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceVLANsSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/vlans/settings", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceVLANsSettings) toBody(ctx context.Context, state ApplianceVLANsSettings) string {
	body := ""
	if !data.VlansEnabled.IsNull() {
		body, _ = sjson.Set(body, "vlansEnabled", data.VlansEnabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceVLANsSettings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("vlansEnabled"); value.Exists() && value.Value() != nil {
		data.VlansEnabled = types.BoolValue(value.Bool())
	} else {
		data.VlansEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceVLANsSettings) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("vlansEnabled"); value.Exists() && !data.VlansEnabled.IsNull() {
		data.VlansEnabled = types.BoolValue(value.Bool())
	} else {
		data.VlansEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceVLANsSettings) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceVLANsSettingsIdentity) toIdentity(ctx context.Context, plan *ApplianceVLANsSettings) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceVLANsSettings) fromIdentity(ctx context.Context, identity *ApplianceVLANsSettingsIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceVLANsSettings) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
