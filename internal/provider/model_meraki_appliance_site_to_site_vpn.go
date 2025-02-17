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

type ApplianceSiteToSiteVPN struct {
	Id        types.String                    `tfsdk:"id"`
	NetworkId types.String                    `tfsdk:"network_id"`
	Mode      types.String                    `tfsdk:"mode"`
	Hubs      []ApplianceSiteToSiteVPNHubs    `tfsdk:"hubs"`
	Subnets   []ApplianceSiteToSiteVPNSubnets `tfsdk:"subnets"`
}

type ApplianceSiteToSiteVPNHubs struct {
	HubId           types.String `tfsdk:"hub_id"`
	UseDefaultRoute types.Bool   `tfsdk:"use_default_route"`
}

type ApplianceSiteToSiteVPNSubnets struct {
	LocalSubnet types.String `tfsdk:"local_subnet"`
	UseVpn      types.Bool   `tfsdk:"use_vpn"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceSiteToSiteVPN) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/vpn/siteToSiteVpn", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceSiteToSiteVPN) toBody(ctx context.Context, state ApplianceSiteToSiteVPN) string {
	body := ""
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	if len(data.Hubs) > 0 {
		body, _ = sjson.Set(body, "hubs", []interface{}{})
		for _, item := range data.Hubs {
			itemBody := ""
			if !item.HubId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "hubId", item.HubId.ValueString())
			}
			if !item.UseDefaultRoute.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "useDefaultRoute", item.UseDefaultRoute.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "hubs.-1", itemBody)
		}
	}
	if len(data.Subnets) > 0 {
		body, _ = sjson.Set(body, "subnets", []interface{}{})
		for _, item := range data.Subnets {
			itemBody := ""
			if !item.LocalSubnet.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "localSubnet", item.LocalSubnet.ValueString())
			}
			if !item.UseVpn.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "useVpn", item.UseVpn.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "subnets.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceSiteToSiteVPN) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("mode"); value.Exists() && value.Value() != nil {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("hubs"); value.Exists() && value.Value() != nil {
		data.Hubs = make([]ApplianceSiteToSiteVPNHubs, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceSiteToSiteVPNHubs{}
			if value := res.Get("hubId"); value.Exists() && value.Value() != nil {
				data.HubId = types.StringValue(value.String())
			} else {
				data.HubId = types.StringNull()
			}
			if value := res.Get("useDefaultRoute"); value.Exists() && value.Value() != nil {
				data.UseDefaultRoute = types.BoolValue(value.Bool())
			} else {
				data.UseDefaultRoute = types.BoolNull()
			}
			(*parent).Hubs = append((*parent).Hubs, data)
			return true
		})
	}
	if value := res.Get("subnets"); value.Exists() && value.Value() != nil {
		data.Subnets = make([]ApplianceSiteToSiteVPNSubnets, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceSiteToSiteVPNSubnets{}
			if value := res.Get("localSubnet"); value.Exists() && value.Value() != nil {
				data.LocalSubnet = types.StringValue(value.String())
			} else {
				data.LocalSubnet = types.StringNull()
			}
			if value := res.Get("useVpn"); value.Exists() && value.Value() != nil {
				data.UseVpn = types.BoolValue(value.Bool())
			} else {
				data.UseVpn = types.BoolNull()
			}
			(*parent).Subnets = append((*parent).Subnets, data)
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
func (data *ApplianceSiteToSiteVPN) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("mode"); value.Exists() && !data.Mode.IsNull() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	for i := 0; i < len(data.Hubs); i++ {
		keys := [...]string{"hubId"}
		keyValues := [...]string{data.Hubs[i].HubId.ValueString()}

		parent := &data
		data := (*parent).Hubs[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("hubs").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Hubs[%d] = %+v",
				i,
				(*parent).Hubs[i],
			))
			(*parent).Hubs = slices.Delete((*parent).Hubs, i, i+1)
			i--

			continue
		}
		if value := res.Get("hubId"); value.Exists() && !data.HubId.IsNull() {
			data.HubId = types.StringValue(value.String())
		} else {
			data.HubId = types.StringNull()
		}
		if value := res.Get("useDefaultRoute"); value.Exists() && !data.UseDefaultRoute.IsNull() {
			data.UseDefaultRoute = types.BoolValue(value.Bool())
		} else {
			data.UseDefaultRoute = types.BoolNull()
		}
		(*parent).Hubs[i] = data
	}
	for i := 0; i < len(data.Subnets); i++ {
		keys := [...]string{"localSubnet"}
		keyValues := [...]string{data.Subnets[i].LocalSubnet.ValueString()}

		parent := &data
		data := (*parent).Subnets[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("subnets").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Subnets[%d] = %+v",
				i,
				(*parent).Subnets[i],
			))
			(*parent).Subnets = slices.Delete((*parent).Subnets, i, i+1)
			i--

			continue
		}
		if value := res.Get("localSubnet"); value.Exists() && !data.LocalSubnet.IsNull() {
			data.LocalSubnet = types.StringValue(value.String())
		} else {
			data.LocalSubnet = types.StringNull()
		}
		if value := res.Get("useVpn"); value.Exists() && !data.UseVpn.IsNull() {
			data.UseVpn = types.BoolValue(value.Bool())
		} else {
			data.UseVpn = types.BoolNull()
		}
		(*parent).Subnets[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceSiteToSiteVPN) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceSiteToSiteVPN) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
