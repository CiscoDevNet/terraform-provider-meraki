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

type OrganizationEarlyAccessFeaturesOptIn struct {
	Id                   types.String `tfsdk:"id"`
	OrganizationId       types.String `tfsdk:"organization_id"`
	ShortName            types.String `tfsdk:"short_name"`
	LimitScopeToNetworks types.List   `tfsdk:"limit_scope_to_networks"`
}

type OrganizationEarlyAccessFeaturesOptInIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationEarlyAccessFeaturesOptIn) getPath() string {
	return fmt.Sprintf("/organizations/%v/earlyAccess/features/optIns", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationEarlyAccessFeaturesOptIn) toBody(ctx context.Context, state OrganizationEarlyAccessFeaturesOptIn) string {
	body := ""
	if !data.ShortName.IsNull() {
		body, _ = sjson.Set(body, "shortName", data.ShortName.ValueString())
	}
	if !data.LimitScopeToNetworks.IsNull() {
		var values []string
		data.LimitScopeToNetworks.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "limitScopeToNetworks", values)
	}
	return body
}

// End of section. //template:end toBody

func (data *OrganizationEarlyAccessFeaturesOptIn) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
		data.ShortName = types.StringValue(value.String())
	} else {
		data.ShortName = types.StringNull()
	}
	if value := res.Get("limitScopeToNetworks"); value.Exists() && value.Value() != nil {
		data.LimitScopeToNetworks = helpers.GetStringListFromMapList(value.Array(), "id")
	} else {
		data.LimitScopeToNetworks = types.ListNull(types.StringType)
	}
}

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationEarlyAccessFeaturesOptIn) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
		data.ShortName = types.StringValue(value.String())
	} else {
		data.ShortName = types.StringNull()
	}
	if value := res.Get("limitScopeToNetworks"); value.Exists() && !data.LimitScopeToNetworks.IsNull() {
		data.LimitScopeToNetworks = helpers.GetStringListFromMapList(value.Array(), "id")
	} else {
		data.LimitScopeToNetworks = types.ListNull(types.StringType)
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationEarlyAccessFeaturesOptIn) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *OrganizationEarlyAccessFeaturesOptInIdentity) toIdentity(ctx context.Context, plan *OrganizationEarlyAccessFeaturesOptIn) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *OrganizationEarlyAccessFeaturesOptIn) fromIdentity(ctx context.Context, identity *OrganizationEarlyAccessFeaturesOptInIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationEarlyAccessFeaturesOptIn) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
