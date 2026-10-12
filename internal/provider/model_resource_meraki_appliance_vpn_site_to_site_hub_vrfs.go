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

type ApplianceVPNSiteToSiteHubVRFs struct {
	Id           types.String                        `tfsdk:"id"`
	NetworkId    types.String                        `tfsdk:"network_id"`
	HubNetworkId types.String                        `tfsdk:"hub_network_id"`
	Json         []ApplianceVPNSiteToSiteHubVRFsJson `tfsdk:"_json"`
}

type ApplianceVPNSiteToSiteHubVRFsJson struct {
	Id types.String `tfsdk:"id"`
}

type ApplianceVPNSiteToSiteHubVRFsIdentity struct {
	NetworkId    types.String `tfsdk:"network_id"`
	HubNetworkId types.String `tfsdk:"hub_network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceVPNSiteToSiteHubVRFs) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/vpn/siteToSite/hubs/%v/vrfs", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.HubNetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceVPNSiteToSiteHubVRFs) toBody(ctx context.Context, state ApplianceVPNSiteToSiteHubVRFs) string {
	body := ""
	{
		body, _ = sjson.Set(body, "_json", []interface{}{})
		for _, item := range data.Json {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "_json.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceVPNSiteToSiteHubVRFs) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("_json"); value.Exists() && value.Value() != nil {
		data.Json = make([]ApplianceVPNSiteToSiteHubVRFsJson, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceVPNSiteToSiteHubVRFsJson{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Json = append((*parent).Json, data)
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
func (data *ApplianceVPNSiteToSiteHubVRFs) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Json); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Json[i].Id.ValueString()}

		parent := &data
		data := (*parent).Json[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("_json").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Json[%d] = %+v",
				i,
				(*parent).Json[i],
			))
			(*parent).Json = slices.Delete((*parent).Json, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).Json[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceVPNSiteToSiteHubVRFs) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceVPNSiteToSiteHubVRFsIdentity) toIdentity(ctx context.Context, plan *ApplianceVPNSiteToSiteHubVRFs) {
	data.NetworkId = plan.NetworkId
	data.HubNetworkId = plan.HubNetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceVPNSiteToSiteHubVRFs) fromIdentity(ctx context.Context, identity *ApplianceVPNSiteToSiteHubVRFsIdentity) {
	data.NetworkId = identity.NetworkId
	data.HubNetworkId = identity.HubNetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceVPNSiteToSiteHubVRFs) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
