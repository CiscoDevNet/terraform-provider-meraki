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

type ResourceSwitchStackRoutingStaticRoutes struct {
	Id             types.String                                  `tfsdk:"id"`
	OrganizationId types.String                                  `tfsdk:"organization_id"`
	NetworkId      types.String                                  `tfsdk:"network_id"`
	SwitchStackId  types.String                                  `tfsdk:"switch_stack_id"`
	Items          []ResourceSwitchStackRoutingStaticRoutesItems `tfsdk:"items"`
}

type ResourceSwitchStackRoutingStaticRoutesItems struct {
	Id                          types.String `tfsdk:"id"`
	AdvertiseViaOspfEnabled     types.Bool   `tfsdk:"advertise_via_ospf_enabled"`
	Name                        types.String `tfsdk:"name"`
	NextHopIp                   types.String `tfsdk:"next_hop_ip"`
	PreferOverOspfRoutesEnabled types.Bool   `tfsdk:"prefer_over_ospf_routes_enabled"`
	Subnet                      types.String `tfsdk:"subnet"`
	VrfLeakRouteToDefaultVrf    types.Bool   `tfsdk:"vrf_leak_route_to_default_vrf"`
	VrfName                     types.String `tfsdk:"vrf_name"`
}

type ResourceSwitchStackRoutingStaticRoutesIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	NetworkId      types.String `tfsdk:"network_id"`
	SwitchStackId  types.String `tfsdk:"switch_stack_id"`
	ItemIds        types.List   `tfsdk:"item_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceSwitchStackRoutingStaticRoutes) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/stacks/%v/routing/staticRoutes", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.SwitchStackId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceSwitchStackRoutingStaticRoutesItems) toBody(ctx context.Context, state ResourceSwitchStackRoutingStaticRoutesItems) string {
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
	if !data.VrfLeakRouteToDefaultVrf.IsNull() {
		body, _ = sjson.Set(body, "vrf.leakRouteToDefaultVrf", data.VrfLeakRouteToDefaultVrf.ValueBool())
	}
	if !data.VrfName.IsNull() {
		body, _ = sjson.Set(body, "vrf.name", data.VrfName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceSwitchStackRoutingStaticRoutes) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceSwitchStackRoutingStaticRoutesItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceSwitchStackRoutingStaticRoutesItems{}
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
		if value := res.Get("vrf.leakRouteToDefaultVrf"); value.Exists() && value.Value() != nil {
			data.VrfLeakRouteToDefaultVrf = types.BoolValue(value.Bool())
		} else {
			data.VrfLeakRouteToDefaultVrf = types.BoolNull()
		}
		if value := res.Get("vrf.name"); value.Exists() && value.Value() != nil {
			data.VrfName = types.StringValue(value.String())
		} else {
			data.VrfName = types.StringNull()
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
func (data *ResourceSwitchStackRoutingStaticRoutes) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
		if value := res.Get("vrf.leakRouteToDefaultVrf"); value.Exists() && !data.VrfLeakRouteToDefaultVrf.IsNull() {
			data.VrfLeakRouteToDefaultVrf = types.BoolValue(value.Bool())
		} else {
			data.VrfLeakRouteToDefaultVrf = types.BoolNull()
		}
		if value := res.Get("vrf.name"); value.Exists() && !data.VrfName.IsNull() {
			data.VrfName = types.StringValue(value.String())
		} else {
			data.VrfName = types.StringNull()
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
func (data *ResourceSwitchStackRoutingStaticRoutes) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceSwitchStackRoutingStaticRoutes) fromBodyImport(ctx context.Context, res meraki.Res) {
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
		if value := res.Get("vrf.leakRouteToDefaultVrf"); value.Exists() && value.Value() != nil {
			data.VrfLeakRouteToDefaultVrf = types.BoolValue(value.Bool())
		} else {
			data.VrfLeakRouteToDefaultVrf = types.BoolNull()
		}
		if value := res.Get("vrf.name"); value.Exists() && value.Value() != nil {
			data.VrfName = types.StringValue(value.String())
		} else {
			data.VrfName = types.StringNull()
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyImport(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ResourceSwitchStackRoutingStaticRoutesIdentity) toIdentity(ctx context.Context, plan *ResourceSwitchStackRoutingStaticRoutes) {
	data.OrganizationId = plan.OrganizationId
	data.NetworkId = plan.NetworkId
	data.SwitchStackId = plan.SwitchStackId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ResourceSwitchStackRoutingStaticRoutes) fromIdentity(ctx context.Context, identity *ResourceSwitchStackRoutingStaticRoutesIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.NetworkId = identity.NetworkId
	data.SwitchStackId = identity.SwitchStackId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceSwitchStackRoutingStaticRoutes) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceSwitchStackRoutingStaticRoutes) hasChanges(ctx context.Context, state *ResourceSwitchStackRoutingStaticRoutes, id string) bool {
	hasChanges := false

	item := ResourceSwitchStackRoutingStaticRoutesItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceSwitchStackRoutingStaticRoutesItems{}
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
	if !item.VrfLeakRouteToDefaultVrf.Equal(stateItem.VrfLeakRouteToDefaultVrf) {
		hasChanges = true
	}
	if !item.VrfName.Equal(stateItem.VrfName) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
