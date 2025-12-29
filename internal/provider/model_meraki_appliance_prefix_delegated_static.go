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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type AppliancePrefixDelegatedStatic struct {
	Id               types.String `tfsdk:"id"`
	NetworkId        types.String `tfsdk:"network_id"`
	Description      types.String `tfsdk:"description"`
	Prefix           types.String `tfsdk:"prefix"`
	OriginType       types.String `tfsdk:"origin_type"`
	OriginInterfaces types.List   `tfsdk:"origin_interfaces"`
}

type AppliancePrefixDelegatedStaticIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
	Id        types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data AppliancePrefixDelegatedStatic) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/prefixes/delegated/statics", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data AppliancePrefixDelegatedStatic) toBody(ctx context.Context, state AppliancePrefixDelegatedStatic) string {
	body := ""
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Prefix.IsNull() {
		body, _ = sjson.Set(body, "prefix", data.Prefix.ValueString())
	}
	if !data.OriginType.IsNull() {
		body, _ = sjson.Set(body, "origin.type", data.OriginType.ValueString())
	}
	if !data.OriginInterfaces.IsNull() {
		var values []string
		data.OriginInterfaces.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "origin.interfaces", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *AppliancePrefixDelegatedStatic) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && value.Value() != nil {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("prefix"); value.Exists() && value.Value() != nil {
		data.Prefix = types.StringValue(value.String())
	} else {
		data.Prefix = types.StringNull()
	}
	if value := res.Get("origin.type"); value.Exists() && value.Value() != nil {
		data.OriginType = types.StringValue(value.String())
	} else {
		data.OriginType = types.StringNull()
	}
	if value := res.Get("origin.interfaces"); value.Exists() && value.Value() != nil {
		data.OriginInterfaces = helpers.GetStringList(value.Array())
	} else {
		data.OriginInterfaces = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *AppliancePrefixDelegatedStatic) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("prefix"); value.Exists() && !data.Prefix.IsNull() {
		data.Prefix = types.StringValue(value.String())
	} else {
		data.Prefix = types.StringNull()
	}
	if value := res.Get("origin.type"); value.Exists() && !data.OriginType.IsNull() {
		data.OriginType = types.StringValue(value.String())
	} else {
		data.OriginType = types.StringNull()
	}
	if value := res.Get("origin.interfaces"); value.Exists() && !data.OriginInterfaces.IsNull() {
		data.OriginInterfaces = helpers.GetStringList(value.Array())
	} else {
		data.OriginInterfaces = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *AppliancePrefixDelegatedStatic) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *AppliancePrefixDelegatedStaticIdentity) toIdentity(ctx context.Context, plan *AppliancePrefixDelegatedStatic) {
	data.NetworkId = plan.NetworkId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *AppliancePrefixDelegatedStatic) fromIdentity(ctx context.Context, identity *AppliancePrefixDelegatedStaticIdentity) {
	data.NetworkId = identity.NetworkId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data AppliancePrefixDelegatedStatic) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
