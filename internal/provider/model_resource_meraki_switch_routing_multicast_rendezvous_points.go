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

type ResourceSwitchRoutingMulticastRendezvousPoints struct {
	Id             types.String                                          `tfsdk:"id"`
	OrganizationId types.String                                          `tfsdk:"organization_id"`
	NetworkId      types.String                                          `tfsdk:"network_id"`
	Items          []ResourceSwitchRoutingMulticastRendezvousPointsItems `tfsdk:"items"`
}

type ResourceSwitchRoutingMulticastRendezvousPointsItems struct {
	Id             types.String `tfsdk:"id"`
	InterfaceIp    types.String `tfsdk:"interface_ip"`
	MulticastGroup types.String `tfsdk:"multicast_group"`
	VrfName        types.String `tfsdk:"vrf_name"`
}

type ResourceSwitchRoutingMulticastRendezvousPointsIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	NetworkId      types.String `tfsdk:"network_id"`
	ItemIds        types.List   `tfsdk:"item_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceSwitchRoutingMulticastRendezvousPoints) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/routing/multicast/rendezvousPoints", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceSwitchRoutingMulticastRendezvousPointsItems) toBody(ctx context.Context, state ResourceSwitchRoutingMulticastRendezvousPointsItems) string {
	body := ""
	if !data.InterfaceIp.IsNull() {
		body, _ = sjson.Set(body, "interfaceIp", data.InterfaceIp.ValueString())
	}
	if !data.MulticastGroup.IsNull() {
		body, _ = sjson.Set(body, "multicastGroup", data.MulticastGroup.ValueString())
	}
	if !data.VrfName.IsNull() {
		body, _ = sjson.Set(body, "vrf.name", data.VrfName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceSwitchRoutingMulticastRendezvousPoints) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceSwitchRoutingMulticastRendezvousPointsItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceSwitchRoutingMulticastRendezvousPointsItems{}
		if value := res.Get("interfaceIp"); value.Exists() && value.Value() != nil {
			data.InterfaceIp = types.StringValue(value.String())
		} else {
			data.InterfaceIp = types.StringNull()
		}
		if value := res.Get("multicastGroup"); value.Exists() && value.Value() != nil {
			data.MulticastGroup = types.StringValue(value.String())
		} else {
			data.MulticastGroup = types.StringNull()
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
		data.Items[index].Id = types.StringValue(res.Get("rendezvousPointId").String())
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
func (data *ResourceSwitchRoutingMulticastRendezvousPoints) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
				if v.Get("rendezvousPointId").String() == (*parent).Items[i].Id.ValueString() {
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
		if value := res.Get("interfaceIp"); value.Exists() && !data.InterfaceIp.IsNull() {
			data.InterfaceIp = types.StringValue(value.String())
		} else {
			data.InterfaceIp = types.StringNull()
		}
		if value := res.Get("multicastGroup"); value.Exists() && !data.MulticastGroup.IsNull() {
			data.MulticastGroup = types.StringValue(value.String())
		} else {
			data.MulticastGroup = types.StringNull()
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
func (data *ResourceSwitchRoutingMulticastRendezvousPoints) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceSwitchRoutingMulticastRendezvousPoints) fromBodyImport(ctx context.Context, res meraki.Res) {
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
				if v.Get("rendezvousPointId").String() == (*parent).Items[i].Id.ValueString() {
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
		if value := res.Get("interfaceIp"); value.Exists() && value.Value() != nil {
			data.InterfaceIp = types.StringValue(value.String())
		} else {
			data.InterfaceIp = types.StringNull()
		}
		if value := res.Get("multicastGroup"); value.Exists() && value.Value() != nil {
			data.MulticastGroup = types.StringValue(value.String())
		} else {
			data.MulticastGroup = types.StringNull()
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

func (data *ResourceSwitchRoutingMulticastRendezvousPointsIdentity) toIdentity(ctx context.Context, plan *ResourceSwitchRoutingMulticastRendezvousPoints) {
	data.OrganizationId = plan.OrganizationId
	data.NetworkId = plan.NetworkId
	if len(data.ItemIds.Elements()) == 0 {
		data.ItemIds = types.ListNull(types.StringType)
	}
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ResourceSwitchRoutingMulticastRendezvousPoints) fromIdentity(ctx context.Context, identity *ResourceSwitchRoutingMulticastRendezvousPointsIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceSwitchRoutingMulticastRendezvousPoints) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceSwitchRoutingMulticastRendezvousPoints) hasChanges(ctx context.Context, state *ResourceSwitchRoutingMulticastRendezvousPoints, id string) bool {
	hasChanges := false

	item := ResourceSwitchRoutingMulticastRendezvousPointsItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceSwitchRoutingMulticastRendezvousPointsItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.InterfaceIp.Equal(stateItem.InterfaceIp) {
		hasChanges = true
	}
	if !item.MulticastGroup.Equal(stateItem.MulticastGroup) {
		hasChanges = true
	}
	if !item.VrfName.Equal(stateItem.VrfName) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
