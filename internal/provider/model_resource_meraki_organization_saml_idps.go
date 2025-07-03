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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceOrganizationSAMLIdPs struct {
	Id             types.String                        `tfsdk:"id"`
	OrganizationId types.String                        `tfsdk:"organization_id"`
	Items          []ResourceOrganizationSAMLIdPsItems `tfsdk:"items"`
}

type ResourceOrganizationSAMLIdPsItems struct {
	Id                      types.String `tfsdk:"id"`
	SloLogoutUrl            types.String `tfsdk:"slo_logout_url"`
	X509certSha1Fingerprint types.String `tfsdk:"x509cert_sha1_fingerprint"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceOrganizationSAMLIdPs) getPath() string {
	return fmt.Sprintf("/organizations/%v/saml/idps", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceOrganizationSAMLIdPsItems) toBody(ctx context.Context, state ResourceOrganizationSAMLIdPsItems) string {
	body := ""
	if !data.SloLogoutUrl.IsNull() {
		body, _ = sjson.Set(body, "sloLogoutUrl", data.SloLogoutUrl.ValueString())
	}
	if !data.X509certSha1Fingerprint.IsNull() {
		body, _ = sjson.Set(body, "x509certSha1Fingerprint", data.X509certSha1Fingerprint.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceOrganizationSAMLIdPs) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceOrganizationSAMLIdPsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceOrganizationSAMLIdPsItems{}
		if value := res.Get("sloLogoutUrl"); value.Exists() && value.Value() != nil {
			data.SloLogoutUrl = types.StringValue(value.String())
		} else {
			data.SloLogoutUrl = types.StringNull()
		}
		if value := res.Get("x509certSha1Fingerprint"); value.Exists() && value.Value() != nil {
			data.X509certSha1Fingerprint = types.StringValue(value.String())
		} else {
			data.X509certSha1Fingerprint = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceOrganizationSAMLIdPs) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("idpId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
		if value := res.Get("sloLogoutUrl"); value.Exists() && !data.SloLogoutUrl.IsNull() {
			data.SloLogoutUrl = types.StringValue(value.String())
		} else {
			data.SloLogoutUrl = types.StringNull()
		}
		if value := res.Get("x509certSha1Fingerprint"); value.Exists() && !data.X509certSha1Fingerprint.IsNull() {
			data.X509certSha1Fingerprint = types.StringValue(value.String())
		} else {
			data.X509certSha1Fingerprint = types.StringNull()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceOrganizationSAMLIdPs) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceOrganizationSAMLIdPs) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceOrganizationSAMLIdPs) hasChanges(ctx context.Context, state *ResourceOrganizationSAMLIdPs, id string) bool {
	hasChanges := false

	item := ResourceOrganizationSAMLIdPsItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceOrganizationSAMLIdPsItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.SloLogoutUrl.Equal(stateItem.SloLogoutUrl) {
		hasChanges = true
	}
	if !item.X509certSha1Fingerprint.Equal(stateItem.X509certSha1Fingerprint) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
