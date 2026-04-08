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

type ApplianceContentFiltering struct {
	Id                   types.String `tfsdk:"id"`
	NetworkId            types.String `tfsdk:"network_id"`
	UrlCategoryListSize  types.String `tfsdk:"url_category_list_size"`
	AllowedUrlPatterns   types.Set    `tfsdk:"allowed_url_patterns"`
	BlockedUrlCategories types.Set    `tfsdk:"blocked_url_categories"`
	BlockedUrlPatterns   types.Set    `tfsdk:"blocked_url_patterns"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceContentFiltering) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/contentFiltering", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceContentFiltering) toBody(ctx context.Context, state ApplianceContentFiltering) string {
	body := ""
	if !data.UrlCategoryListSize.IsNull() {
		body, _ = sjson.Set(body, "urlCategoryListSize", data.UrlCategoryListSize.ValueString())
	}
	if !data.AllowedUrlPatterns.IsNull() {
		var values []string
		data.AllowedUrlPatterns.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "allowedUrlPatterns", values)
	}
	if !data.BlockedUrlCategories.IsNull() {
		var values []string
		data.BlockedUrlCategories.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "blockedUrlCategories", values)
	}
	if !data.BlockedUrlPatterns.IsNull() {
		var values []string
		data.BlockedUrlPatterns.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "blockedUrlPatterns", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceContentFiltering) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("allowedUrlPatterns"); value.Exists() && value.Value() != nil {
		data.AllowedUrlPatterns = helpers.GetStringSet(value.Array())
	} else {
		data.AllowedUrlPatterns = types.SetNull(types.StringType)
	}
	if value := res.Get("blockedUrlPatterns"); value.Exists() && value.Value() != nil {
		data.BlockedUrlPatterns = helpers.GetStringSet(value.Array())
	} else {
		data.BlockedUrlPatterns = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceContentFiltering) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("allowedUrlPatterns"); value.Exists() && !data.AllowedUrlPatterns.IsNull() {
		data.AllowedUrlPatterns = helpers.GetStringSet(value.Array())
	} else {
		data.AllowedUrlPatterns = types.SetNull(types.StringType)
	}
	if value := res.Get("blockedUrlPatterns"); value.Exists() && !data.BlockedUrlPatterns.IsNull() {
		data.BlockedUrlPatterns = helpers.GetStringSet(value.Array())
	} else {
		data.BlockedUrlPatterns = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceContentFiltering) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceContentFiltering) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
