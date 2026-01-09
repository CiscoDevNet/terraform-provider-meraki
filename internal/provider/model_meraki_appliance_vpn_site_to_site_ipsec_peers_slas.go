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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceVPNSiteToSiteIPsecPeersSLAs struct {
	Id             types.String                                `tfsdk:"id"`
	OrganizationId types.String                                `tfsdk:"organization_id"`
	Items          []ApplianceVPNSiteToSiteIPsecPeersSLAsItems `tfsdk:"items"`
}

type ApplianceVPNSiteToSiteIPsecPeersSLAsItems struct {
	Name types.String `tfsdk:"name"`
	Uri  types.String `tfsdk:"uri"`
}

type ApplianceVPNSiteToSiteIPsecPeersSLAsIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceVPNSiteToSiteIPsecPeersSLAs) getPath() string {
	return fmt.Sprintf("/organizations/%v/appliance/vpn/siteToSite/ipsec/peers/slas", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceVPNSiteToSiteIPsecPeersSLAs) toBody(ctx context.Context, state ApplianceVPNSiteToSiteIPsecPeersSLAs) string {
	body := ""
	{
		body, _ = sjson.Set(body, "items", []interface{}{})
		for _, item := range data.Items {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Uri.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "uri", item.Uri.ValueString())
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceVPNSiteToSiteIPsecPeersSLAs) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("items"); value.Exists() && value.Value() != nil {
		data.Items = make([]ApplianceVPNSiteToSiteIPsecPeersSLAsItems, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceVPNSiteToSiteIPsecPeersSLAsItems{}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("uri"); value.Exists() && value.Value() != nil {
				data.Uri = types.StringValue(value.String())
			} else {
				data.Uri = types.StringNull()
			}
			(*parent).Items = append((*parent).Items, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceVPNSiteToSiteIPsecPeersSLAs) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Items); i++ {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Items[i].Name.ValueString()}

		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Items[%d] = %+v",
				i,
				(*parent).Items[i],
			))
			(*parent).Items = slices.Delete((*parent).Items, i, i+1)
			i--

			continue
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("uri"); value.Exists() && !data.Uri.IsNull() {
			data.Uri = types.StringValue(value.String())
		} else {
			data.Uri = types.StringNull()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceVPNSiteToSiteIPsecPeersSLAs) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceVPNSiteToSiteIPsecPeersSLAsIdentity) toIdentity(ctx context.Context, plan *ApplianceVPNSiteToSiteIPsecPeersSLAs) {
	data.OrganizationId = plan.OrganizationId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceVPNSiteToSiteIPsecPeersSLAs) fromIdentity(ctx context.Context, identity *ApplianceVPNSiteToSiteIPsecPeersSLAsIdentity) {
	data.OrganizationId = identity.OrganizationId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceVPNSiteToSiteIPsecPeersSLAs) toDestroyBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "items", []interface{}{})
	return body
}

// End of section. //template:end toDestroyBody
