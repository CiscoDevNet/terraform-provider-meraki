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

type ResourceSwitchRoutingStaticRoutes struct {
	Id             types.String                             `tfsdk:"id"`
	OrganizationId types.String                             `tfsdk:"organization_id"`
	Serial         types.String                             `tfsdk:"serial"`
	Items          []ResourceSwitchRoutingStaticRoutesItems `tfsdk:"items"`
}

type ResourceSwitchRoutingStaticRoutesItems struct {
	Id                          types.String `tfsdk:"id"`
	AdvertiseViaOspfEnabled     types.Bool   `tfsdk:"advertise_via_ospf_enabled"`
	Name                        types.String `tfsdk:"name"`
	NextHopIp                   types.String `tfsdk:"next_hop_ip"`
	PreferOverOspfRoutesEnabled types.Bool   `tfsdk:"prefer_over_ospf_routes_enabled"`
	Subnet                      types.String `tfsdk:"subnet"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceSwitchRoutingStaticRoutes) getPath() string {
	return fmt.Sprintf("/devices/%v/switch/routing/staticRoutes", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceSwitchRoutingStaticRoutesItems) toBody(ctx context.Context, state ResourceSwitchRoutingStaticRoutesItems) string {
	body := ""
	if !data.AdvertiseViaOspfEnabled.IsNull() {
		body, _ = sjson.Set(body, "advertiseViaOspfEnabled", data.AdvertiseViaOspfEnabled.ValueBool())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.NextHopIp.IsNull() {
		body, _ = sjson.Set(body, "nextHopIp", data.NextHopIp.ValueString())
	}
	if !data.PreferOverOspfRoutesEnabled.IsNull() {
		body, _ = sjson.Set(body, "preferOverOspfRoutesEnabled", data.PreferOverOspfRoutesEnabled.ValueBool())
	}
	if !data.Subnet.IsNull() {
		body, _ = sjson.Set(body, "subnet", data.Subnet.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceSwitchRoutingStaticRoutes) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceSwitchRoutingStaticRoutesItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceSwitchRoutingStaticRoutesItems{}
		if value := res.Get("advertiseViaOspfEnabled"); value.Exists() && value.Value() != nil {
			data.AdvertiseViaOspfEnabled = types.BoolValue(value.Bool())
		} else {
			data.AdvertiseViaOspfEnabled = types.BoolNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("nextHopIp"); value.Exists() && value.Value() != nil {
			data.NextHopIp = types.StringValue(value.String())
		} else {
			data.NextHopIp = types.StringNull()
		}
		if value := res.Get("preferOverOspfRoutesEnabled"); value.Exists() && value.Value() != nil {
			data.PreferOverOspfRoutesEnabled = types.BoolValue(value.Bool())
		} else {
			data.PreferOverOspfRoutesEnabled = types.BoolNull()
		}
		if value := res.Get("subnet"); value.Exists() && value.Value() != nil {
			data.Subnet = types.StringValue(value.String())
		} else {
			data.Subnet = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	index := 0
	res.ForEach(func(k, res gjson.Result) bool {
		data.Items[index].Id = types.StringValue(res.Get("staticRouteId").String())
		index++
		return true
	})
	data.Id = data.OrganizationId
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceSwitchRoutingStaticRoutes) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("staticRouteId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		if value := res.Get("advertiseViaOspfEnabled"); value.Exists() && !data.AdvertiseViaOspfEnabled.IsNull() {
			data.AdvertiseViaOspfEnabled = types.BoolValue(value.Bool())
		} else {
			data.AdvertiseViaOspfEnabled = types.BoolNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("nextHopIp"); value.Exists() && !data.NextHopIp.IsNull() {
			data.NextHopIp = types.StringValue(value.String())
		} else {
			data.NextHopIp = types.StringNull()
		}
		if value := res.Get("preferOverOspfRoutesEnabled"); value.Exists() && !data.PreferOverOspfRoutesEnabled.IsNull() {
			data.PreferOverOspfRoutesEnabled = types.BoolValue(value.Bool())
		} else {
			data.PreferOverOspfRoutesEnabled = types.BoolNull()
		}
		if value := res.Get("subnet"); value.Exists() && !data.Subnet.IsNull() {
			data.Subnet = types.StringValue(value.String())
		} else {
			data.Subnet = types.StringNull()
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyPartial(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceSwitchRoutingStaticRoutes) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceSwitchRoutingStaticRoutes) fromBodyImport(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("staticRouteId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		if value := res.Get("advertiseViaOspfEnabled"); value.Exists() && value.Value() != nil {
			data.AdvertiseViaOspfEnabled = types.BoolValue(value.Bool())
		} else {
			data.AdvertiseViaOspfEnabled = types.BoolNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("nextHopIp"); value.Exists() && value.Value() != nil {
			data.NextHopIp = types.StringValue(value.String())
		} else {
			data.NextHopIp = types.StringNull()
		}
		if value := res.Get("preferOverOspfRoutesEnabled"); value.Exists() && value.Value() != nil {
			data.PreferOverOspfRoutesEnabled = types.BoolValue(value.Bool())
		} else {
			data.PreferOverOspfRoutesEnabled = types.BoolNull()
		}
		if value := res.Get("subnet"); value.Exists() && value.Value() != nil {
			data.Subnet = types.StringValue(value.String())
		} else {
			data.Subnet = types.StringNull()
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyImport(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceSwitchRoutingStaticRoutes) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceSwitchRoutingStaticRoutes) hasChanges(ctx context.Context, state *ResourceSwitchRoutingStaticRoutes, id string) bool {
	hasChanges := false

	item := ResourceSwitchRoutingStaticRoutesItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceSwitchRoutingStaticRoutesItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.AdvertiseViaOspfEnabled.Equal(stateItem.AdvertiseViaOspfEnabled) {
		hasChanges = true
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	if !item.NextHopIp.Equal(stateItem.NextHopIp) {
		hasChanges = true
	}
	if !item.PreferOverOspfRoutesEnabled.Equal(stateItem.PreferOverOspfRoutesEnabled) {
		hasChanges = true
	}
	if !item.Subnet.Equal(stateItem.Subnet) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
