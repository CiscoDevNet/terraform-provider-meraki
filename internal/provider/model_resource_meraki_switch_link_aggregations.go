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

type ResourceSwitchLinkAggregations struct {
	Id             types.String                          `tfsdk:"id"`
	OrganizationId types.String                          `tfsdk:"organization_id"`
	NetworkId      types.String                          `tfsdk:"network_id"`
	Items          []ResourceSwitchLinkAggregationsItems `tfsdk:"items"`
}

type ResourceSwitchLinkAggregationsItems struct {
	Id                 types.String                                       `tfsdk:"id"`
	SwitchPorts        []ResourceSwitchLinkAggregationsSwitchPorts        `tfsdk:"switch_ports"`
	SwitchProfilePorts []ResourceSwitchLinkAggregationsSwitchProfilePorts `tfsdk:"switch_profile_ports"`
}

type ResourceSwitchLinkAggregationsSwitchPorts struct {
	PortId types.String `tfsdk:"port_id"`
	Serial types.String `tfsdk:"serial"`
}

type ResourceSwitchLinkAggregationsSwitchProfilePorts struct {
	PortId  types.String `tfsdk:"port_id"`
	Profile types.String `tfsdk:"profile"`
}

type ResourceSwitchLinkAggregationsIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	NetworkId      types.String `tfsdk:"network_id"`
	ItemIds        types.List   `tfsdk:"item_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceSwitchLinkAggregations) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/linkAggregations", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceSwitchLinkAggregationsItems) toBody(ctx context.Context, state ResourceSwitchLinkAggregationsItems) string {
	body := ""
	if len(data.SwitchPorts) > 0 {
		body, _ = sjson.Set(body, "switchPorts", []interface{}{})
		for _, item := range data.SwitchPorts {
			itemBody := ""
			if !item.PortId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "portId", item.PortId.ValueString())
			}
			if !item.Serial.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serial", item.Serial.ValueString())
			}
			body, _ = sjson.SetRaw(body, "switchPorts.-1", itemBody)
		}
	}
	if len(data.SwitchProfilePorts) > 0 {
		body, _ = sjson.Set(body, "switchProfilePorts", []interface{}{})
		for _, item := range data.SwitchProfilePorts {
			itemBody := ""
			if !item.PortId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "portId", item.PortId.ValueString())
			}
			if !item.Profile.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "profile", item.Profile.ValueString())
			}
			body, _ = sjson.SetRaw(body, "switchProfilePorts.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceSwitchLinkAggregations) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceSwitchLinkAggregationsItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceSwitchLinkAggregationsItems{}
		if value := res.Get("switchPorts"); value.Exists() && value.Value() != nil {
			data.SwitchPorts = make([]ResourceSwitchLinkAggregationsSwitchPorts, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceSwitchLinkAggregationsSwitchPorts{}
				if value := res.Get("portId"); value.Exists() && value.Value() != nil {
					data.PortId = types.StringValue(value.String())
				} else {
					data.PortId = types.StringNull()
				}
				if value := res.Get("serial"); value.Exists() && value.Value() != nil {
					data.Serial = types.StringValue(value.String())
				} else {
					data.Serial = types.StringNull()
				}
				(*parent).SwitchPorts = append((*parent).SwitchPorts, data)
				return true
			})
		}
		if value := res.Get("switchProfilePorts"); value.Exists() && value.Value() != nil {
			data.SwitchProfilePorts = make([]ResourceSwitchLinkAggregationsSwitchProfilePorts, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceSwitchLinkAggregationsSwitchProfilePorts{}
				if value := res.Get("portId"); value.Exists() && value.Value() != nil {
					data.PortId = types.StringValue(value.String())
				} else {
					data.PortId = types.StringNull()
				}
				if value := res.Get("profile"); value.Exists() && value.Value() != nil {
					data.Profile = types.StringValue(value.String())
				} else {
					data.Profile = types.StringNull()
				}
				(*parent).SwitchProfilePorts = append((*parent).SwitchProfilePorts, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	index := 0
	res.ForEach(func(k, res gjson.Result) bool {
		data.Items[index].Id = types.StringValue(res.Get("id").String())
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
func (data *ResourceSwitchLinkAggregations) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
				if v.Get("id").String() == (*parent).Items[i].Id.ValueString() {
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
		for i := 0; i < len(data.SwitchPorts); i++ {
			keys := [...]string{"portId", "serial"}
			keyValues := [...]string{data.SwitchPorts[i].PortId.ValueString(), data.SwitchPorts[i].Serial.ValueString()}

			parent := &data
			data := (*parent).SwitchPorts[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("switchPorts").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SwitchPorts[%d] = %+v",
					i,
					(*parent).SwitchPorts[i],
				))
				(*parent).SwitchPorts = slices.Delete((*parent).SwitchPorts, i, i+1)
				i--

				continue
			}
			if value := res.Get("portId"); value.Exists() && !data.PortId.IsNull() {
				data.PortId = types.StringValue(value.String())
			} else {
				data.PortId = types.StringNull()
			}
			if value := res.Get("serial"); value.Exists() && !data.Serial.IsNull() {
				data.Serial = types.StringValue(value.String())
			} else {
				data.Serial = types.StringNull()
			}
			(*parent).SwitchPorts[i] = data
		}
		for i := 0; i < len(data.SwitchProfilePorts); i++ {
			keys := [...]string{"portId", "profile"}
			keyValues := [...]string{data.SwitchProfilePorts[i].PortId.ValueString(), data.SwitchProfilePorts[i].Profile.ValueString()}

			parent := &data
			data := (*parent).SwitchProfilePorts[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("switchProfilePorts").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SwitchProfilePorts[%d] = %+v",
					i,
					(*parent).SwitchProfilePorts[i],
				))
				(*parent).SwitchProfilePorts = slices.Delete((*parent).SwitchProfilePorts, i, i+1)
				i--

				continue
			}
			if value := res.Get("portId"); value.Exists() && !data.PortId.IsNull() {
				data.PortId = types.StringValue(value.String())
			} else {
				data.PortId = types.StringNull()
			}
			if value := res.Get("profile"); value.Exists() && !data.Profile.IsNull() {
				data.Profile = types.StringValue(value.String())
			} else {
				data.Profile = types.StringNull()
			}
			(*parent).SwitchProfilePorts[i] = data
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
func (data *ResourceSwitchLinkAggregations) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceSwitchLinkAggregations) fromBodyImport(ctx context.Context, res meraki.Res) {
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
				if v.Get("id").String() == (*parent).Items[i].Id.ValueString() {
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
		if value := res.Get("switchPorts"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.SwitchPorts = make([]ResourceSwitchLinkAggregationsSwitchPorts, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceSwitchLinkAggregationsSwitchPorts{}
				if value := res.Get("portId"); value.Exists() && value.Value() != nil {
					data.PortId = types.StringValue(value.String())
				} else {
					data.PortId = types.StringNull()
				}
				if value := res.Get("serial"); value.Exists() && value.Value() != nil {
					data.Serial = types.StringValue(value.String())
				} else {
					data.Serial = types.StringNull()
				}
				(*parent).SwitchPorts = append((*parent).SwitchPorts, data)
				return true
			})
		}
		if value := res.Get("switchProfilePorts"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.SwitchProfilePorts = make([]ResourceSwitchLinkAggregationsSwitchProfilePorts, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceSwitchLinkAggregationsSwitchProfilePorts{}
				if value := res.Get("portId"); value.Exists() && value.Value() != nil {
					data.PortId = types.StringValue(value.String())
				} else {
					data.PortId = types.StringNull()
				}
				if value := res.Get("profile"); value.Exists() && value.Value() != nil {
					data.Profile = types.StringValue(value.String())
				} else {
					data.Profile = types.StringNull()
				}
				(*parent).SwitchProfilePorts = append((*parent).SwitchProfilePorts, data)
				return true
			})
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

func (data *ResourceSwitchLinkAggregationsIdentity) toIdentity(ctx context.Context, plan *ResourceSwitchLinkAggregations) {
	data.OrganizationId = plan.OrganizationId
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ResourceSwitchLinkAggregations) fromIdentity(ctx context.Context, identity *ResourceSwitchLinkAggregationsIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceSwitchLinkAggregations) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceSwitchLinkAggregations) hasChanges(ctx context.Context, state *ResourceSwitchLinkAggregations, id string) bool {
	hasChanges := false

	item := ResourceSwitchLinkAggregationsItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceSwitchLinkAggregationsItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if len(item.SwitchPorts) != len(stateItem.SwitchPorts) {
		hasChanges = true
	} else {
		for i := range item.SwitchPorts {
			if !item.SwitchPorts[i].PortId.Equal(stateItem.SwitchPorts[i].PortId) {
				hasChanges = true
			}
			if !item.SwitchPorts[i].Serial.Equal(stateItem.SwitchPorts[i].Serial) {
				hasChanges = true
			}
		}
	}
	if len(item.SwitchProfilePorts) != len(stateItem.SwitchProfilePorts) {
		hasChanges = true
	} else {
		for i := range item.SwitchProfilePorts {
			if !item.SwitchProfilePorts[i].PortId.Equal(stateItem.SwitchProfilePorts[i].PortId) {
				hasChanges = true
			}
			if !item.SwitchProfilePorts[i].Profile.Equal(stateItem.SwitchProfilePorts[i].Profile) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

func (data *ResourceSwitchLinkAggregationsItems) equalItem(ctx context.Context, state *ResourceSwitchLinkAggregationsItems) bool {
	if len(data.SwitchPorts) != len(state.SwitchPorts) {
		return false
	}
	for i := range data.SwitchPorts {
		if !data.SwitchPorts[i].PortId.Equal(state.SwitchPorts[i].PortId) {
			return false
		}
		if !data.SwitchPorts[i].Serial.Equal(state.SwitchPorts[i].Serial) {
			return false
		}
	}
	if len(data.SwitchProfilePorts) != len(state.SwitchProfilePorts) {
		return false
	}
	for i := range data.SwitchProfilePorts {
		if !data.SwitchProfilePorts[i].PortId.Equal(state.SwitchProfilePorts[i].PortId) {
			return false
		}
		if !data.SwitchProfilePorts[i].Profile.Equal(state.SwitchProfilePorts[i].Profile) {
			return false
		}
	}
	return true
}
