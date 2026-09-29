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

type ApplianceUmbrellaAccount struct {
	Id                     types.String `tfsdk:"id"`
	NetworkId              types.String `tfsdk:"network_id"`
	ApiKey                 types.String `tfsdk:"api_key"`
	ApiKeyWo               types.String `tfsdk:"api_key_wo"`
	ApiKeyWoVersion        types.Int64  `tfsdk:"api_key_wo_version"`
	ApiSecret              types.String `tfsdk:"api_secret"`
	ApiSecretWo            types.String `tfsdk:"api_secret_wo"`
	ApiSecretWoVersion     types.Int64  `tfsdk:"api_secret_wo_version"`
	UmbrellaOrganizationId types.String `tfsdk:"umbrella_organization_id"`
}

type ApplianceUmbrellaAccountIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceUmbrellaAccount) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/umbrella/account/connect", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceUmbrellaAccount) toBody(ctx context.Context, state ApplianceUmbrellaAccount) string {
	body := ""
	if !data.ApiKeyWo.IsNull() {
		body, _ = sjson.Set(body, "api.key", data.ApiKeyWo.ValueString())
	} else if !data.ApiKey.IsNull() {
		body, _ = sjson.Set(body, "api.key", data.ApiKey.ValueString())
	}
	if !data.ApiSecretWo.IsNull() {
		body, _ = sjson.Set(body, "api.secret", data.ApiSecretWo.ValueString())
	} else if !data.ApiSecret.IsNull() {
		body, _ = sjson.Set(body, "api.secret", data.ApiSecret.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceUmbrellaAccount) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("api.key"); value.Exists() && value.Value() != nil {
		data.ApiKey = types.StringValue(value.String())
	} else {
		data.ApiKey = types.StringNull()
	}
	if value := res.Get("api.secret"); value.Exists() && value.Value() != nil {
		data.ApiSecret = types.StringValue(value.String())
	} else {
		data.ApiSecret = types.StringNull()
	}
	if value := res.Get("umbrella.organization.id"); value.Exists() && value.Value() != nil {
		data.UmbrellaOrganizationId = types.StringValue(value.String())
	} else {
		data.UmbrellaOrganizationId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceUmbrellaAccount) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("api.key"); value.Exists() && !data.ApiKey.IsNull() {
		data.ApiKey = types.StringValue(value.String())
	} else {
		data.ApiKey = types.StringNull()
	}
	if value := res.Get("api.secret"); value.Exists() && !data.ApiSecret.IsNull() {
		data.ApiSecret = types.StringValue(value.String())
	} else {
		data.ApiSecret = types.StringNull()
	}
	if value := res.Get("umbrella.organization.id"); value.Exists() && !data.UmbrellaOrganizationId.IsNull() {
		data.UmbrellaOrganizationId = types.StringValue(value.String())
	} else {
		data.UmbrellaOrganizationId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceUmbrellaAccount) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
	if data.UmbrellaOrganizationId.IsUnknown() {
		if value := res.Get("umbrella.organization.id"); value.Exists() && !data.UmbrellaOrganizationId.IsNull() {
			data.UmbrellaOrganizationId = types.StringValue(value.String())
		} else {
			data.UmbrellaOrganizationId = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceUmbrellaAccountIdentity) toIdentity(ctx context.Context, plan *ApplianceUmbrellaAccount) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceUmbrellaAccount) fromIdentity(ctx context.Context, identity *ApplianceUmbrellaAccountIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceUmbrellaAccount) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
