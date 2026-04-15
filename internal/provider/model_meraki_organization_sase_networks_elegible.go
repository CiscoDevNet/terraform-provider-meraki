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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationSaseNetworksElegible struct {
	Id                       types.String                            `tfsdk:"id"`
	OrganizationId           types.String                            `tfsdk:"organization_id"`
	MetaCountsItemsRemaining types.Int64                             `tfsdk:"meta_counts_items_remaining"`
	MetaCountsItemsTotal     types.Int64                             `tfsdk:"meta_counts_items_total"`
	Items                    []OrganizationSaseNetworksElegibleItems `tfsdk:"items"`
}

type OrganizationSaseNetworksElegibleItems struct {
	Name                       types.String `tfsdk:"name"`
	NetworkId                  types.String `tfsdk:"network_id"`
	Type                       types.String `tfsdk:"type"`
	AddressStreet              types.String `tfsdk:"address_street"`
	DevicePrimaryModel         types.String `tfsdk:"device_primary_model"`
	RegionName                 types.String `tfsdk:"region_name"`
	RoutingDefaultRouteEnabled types.Bool   `tfsdk:"routing_default_route_enabled"`
	VpnType                    types.String `tfsdk:"vpn_type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationSaseNetworksElegible) getPath() string {
	return fmt.Sprintf("/organizations/%v/sase/networks/eligible", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationSaseNetworksElegible) toBody(ctx context.Context, state OrganizationSaseNetworksElegible) string {
	body := ""
	if !data.MetaCountsItemsRemaining.IsNull() {
		body, _ = sjson.Set(body, "meta.counts.items.remaining", data.MetaCountsItemsRemaining.ValueInt64())
	}
	if !data.MetaCountsItemsTotal.IsNull() {
		body, _ = sjson.Set(body, "meta.counts.items.total", data.MetaCountsItemsTotal.ValueInt64())
	}
	if data.Items != nil {
		body, _ = sjson.Set(body, "items", []interface{}{})
		for _, item := range data.Items {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.NetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "networkId", item.NetworkId.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.AddressStreet.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "address.street", item.AddressStreet.ValueString())
			}
			if !item.DevicePrimaryModel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "device.primary.model", item.DevicePrimaryModel.ValueString())
			}
			if !item.RegionName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "region.name", item.RegionName.ValueString())
			}
			if !item.RoutingDefaultRouteEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routing.defaultRoute.enabled", item.RoutingDefaultRouteEnabled.ValueBool())
			}
			if !item.VpnType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vpn.type", item.VpnType.ValueString())
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationSaseNetworksElegible) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("meta.counts.items.remaining"); value.Exists() && value.Value() != nil {
		data.MetaCountsItemsRemaining = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsRemaining = types.Int64Null()
	}
	if value := res.Get("meta.counts.items.total"); value.Exists() && value.Value() != nil {
		data.MetaCountsItemsTotal = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsTotal = types.Int64Null()
	}
	if value := res.Get("items"); value.Exists() && value.Value() != nil {
		data.Items = make([]OrganizationSaseNetworksElegibleItems, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationSaseNetworksElegibleItems{}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("networkId"); value.Exists() && value.Value() != nil {
				data.NetworkId = types.StringValue(value.String())
			} else {
				data.NetworkId = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && value.Value() != nil {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("address.street"); value.Exists() && value.Value() != nil {
				data.AddressStreet = types.StringValue(value.String())
			} else {
				data.AddressStreet = types.StringNull()
			}
			if value := res.Get("device.primary.model"); value.Exists() && value.Value() != nil {
				data.DevicePrimaryModel = types.StringValue(value.String())
			} else {
				data.DevicePrimaryModel = types.StringNull()
			}
			if value := res.Get("region.name"); value.Exists() && value.Value() != nil {
				data.RegionName = types.StringValue(value.String())
			} else {
				data.RegionName = types.StringNull()
			}
			if value := res.Get("routing.defaultRoute.enabled"); value.Exists() && value.Value() != nil {
				data.RoutingDefaultRouteEnabled = types.BoolValue(value.Bool())
			} else {
				data.RoutingDefaultRouteEnabled = types.BoolNull()
			}
			if value := res.Get("vpn.type"); value.Exists() && value.Value() != nil {
				data.VpnType = types.StringValue(value.String())
			} else {
				data.VpnType = types.StringNull()
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
func (data *OrganizationSaseNetworksElegible) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("meta.counts.items.remaining"); value.Exists() && !data.MetaCountsItemsRemaining.IsNull() {
		data.MetaCountsItemsRemaining = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsRemaining = types.Int64Null()
	}
	if value := res.Get("meta.counts.items.total"); value.Exists() && !data.MetaCountsItemsTotal.IsNull() {
		data.MetaCountsItemsTotal = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsTotal = types.Int64Null()
	}
	for i := 0; i < len(data.Items); i++ {
		keys := [...]string{"name", "networkId", "type", "address.street", "device.primary.model", "region.name", "routing.defaultRoute.enabled", "vpn.type"}
		keyValues := [...]string{data.Items[i].Name.ValueString(), data.Items[i].NetworkId.ValueString(), data.Items[i].Type.ValueString(), data.Items[i].AddressStreet.ValueString(), data.Items[i].DevicePrimaryModel.ValueString(), data.Items[i].RegionName.ValueString(), strconv.FormatBool(data.Items[i].RoutingDefaultRouteEnabled.ValueBool()), data.Items[i].VpnType.ValueString()}

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
		if value := res.Get("networkId"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("address.street"); value.Exists() && !data.AddressStreet.IsNull() {
			data.AddressStreet = types.StringValue(value.String())
		} else {
			data.AddressStreet = types.StringNull()
		}
		if value := res.Get("device.primary.model"); value.Exists() && !data.DevicePrimaryModel.IsNull() {
			data.DevicePrimaryModel = types.StringValue(value.String())
		} else {
			data.DevicePrimaryModel = types.StringNull()
		}
		if value := res.Get("region.name"); value.Exists() && !data.RegionName.IsNull() {
			data.RegionName = types.StringValue(value.String())
		} else {
			data.RegionName = types.StringNull()
		}
		if value := res.Get("routing.defaultRoute.enabled"); value.Exists() && !data.RoutingDefaultRouteEnabled.IsNull() {
			data.RoutingDefaultRouteEnabled = types.BoolValue(value.Bool())
		} else {
			data.RoutingDefaultRouteEnabled = types.BoolNull()
		}
		if value := res.Get("vpn.type"); value.Exists() && !data.VpnType.IsNull() {
			data.VpnType = types.StringValue(value.String())
		} else {
			data.VpnType = types.StringNull()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationSaseNetworksElegible) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationSaseNetworksElegible) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
