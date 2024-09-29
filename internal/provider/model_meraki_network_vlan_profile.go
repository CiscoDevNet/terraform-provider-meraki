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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkVLANProfile struct {
	Id         types.String                   `tfsdk:"id"`
	NetworkId  types.String                   `tfsdk:"network_id"`
	Iname      types.String                   `tfsdk:"iname"`
	Name       types.String                   `tfsdk:"name"`
	VlanGroups []NetworkVLANProfileVlanGroups `tfsdk:"vlan_groups"`
	VlanNames  []NetworkVLANProfileVlanNames  `tfsdk:"vlan_names"`
}

type NetworkVLANProfileVlanGroups struct {
	Name    types.String `tfsdk:"name"`
	VlanIds types.String `tfsdk:"vlan_ids"`
}

type NetworkVLANProfileVlanNames struct {
	Name                  types.String `tfsdk:"name"`
	VlanId                types.String `tfsdk:"vlan_id"`
	AdaptivePolicyGroupId types.String `tfsdk:"adaptive_policy_group_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkVLANProfile) getPath() string {
	return fmt.Sprintf("/networks/%v/vlanProfiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkVLANProfile) toBody(ctx context.Context, state NetworkVLANProfile) string {
	body := ""
	if !data.Iname.IsNull() {
		body, _ = sjson.Set(body, "iname", data.Iname.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	{
		body, _ = sjson.Set(body, "vlanGroups", []interface{}{})
		for _, item := range data.VlanGroups {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.VlanIds.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanIds", item.VlanIds.ValueString())
			}
			body, _ = sjson.SetRaw(body, "vlanGroups.-1", itemBody)
		}
	}
	{
		body, _ = sjson.Set(body, "vlanNames", []interface{}{})
		for _, item := range data.VlanNames {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.VlanId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanId", item.VlanId.ValueString())
			}
			if !item.AdaptivePolicyGroupId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "adaptivePolicyGroup.id", item.AdaptivePolicyGroupId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "vlanNames.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkVLANProfile) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("iname"); value.Exists() && value.Value() != nil {
		data.Iname = types.StringValue(value.String())
	} else {
		data.Iname = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("vlanGroups"); value.Exists() && value.Value() != nil {
		data.VlanGroups = make([]NetworkVLANProfileVlanGroups, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkVLANProfileVlanGroups{}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("vlanIds"); value.Exists() && value.Value() != nil {
				data.VlanIds = types.StringValue(value.String())
			} else {
				data.VlanIds = types.StringNull()
			}
			(*parent).VlanGroups = append((*parent).VlanGroups, data)
			return true
		})
	}
	if value := res.Get("vlanNames"); value.Exists() && value.Value() != nil {
		data.VlanNames = make([]NetworkVLANProfileVlanNames, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkVLANProfileVlanNames{}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("vlanId"); value.Exists() && value.Value() != nil {
				data.VlanId = types.StringValue(value.String())
			} else {
				data.VlanId = types.StringNull()
			}
			if value := res.Get("adaptivePolicyGroup.id"); value.Exists() && value.Value() != nil {
				data.AdaptivePolicyGroupId = types.StringValue(value.String())
			} else {
				data.AdaptivePolicyGroupId = types.StringNull()
			}
			(*parent).VlanNames = append((*parent).VlanNames, data)
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
func (data *NetworkVLANProfile) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("iname"); value.Exists() && !data.Iname.IsNull() {
		data.Iname = types.StringValue(value.String())
	} else {
		data.Iname = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	for i := 0; i < len(data.VlanGroups); i++ {
		keys := [...]string{"name"}
		keyValues := [...]string{data.VlanGroups[i].Name.ValueString()}

		parent := &data
		data := (*parent).VlanGroups[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("vlanGroups").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing VlanGroups[%d] = %+v",
				i,
				(*parent).VlanGroups[i],
			))
			(*parent).VlanGroups = slices.Delete((*parent).VlanGroups, i, i+1)
			i--

			continue
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("vlanIds"); value.Exists() && !data.VlanIds.IsNull() {
			data.VlanIds = types.StringValue(value.String())
		} else {
			data.VlanIds = types.StringNull()
		}
		(*parent).VlanGroups[i] = data
	}
	for i := 0; i < len(data.VlanNames); i++ {
		keys := [...]string{"name"}
		keyValues := [...]string{data.VlanNames[i].Name.ValueString()}

		parent := &data
		data := (*parent).VlanNames[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("vlanNames").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing VlanNames[%d] = %+v",
				i,
				(*parent).VlanNames[i],
			))
			(*parent).VlanNames = slices.Delete((*parent).VlanNames, i, i+1)
			i--

			continue
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
			data.VlanId = types.StringValue(value.String())
		} else {
			data.VlanId = types.StringNull()
		}
		if value := res.Get("adaptivePolicyGroup.id"); value.Exists() && !data.AdaptivePolicyGroupId.IsNull() {
			data.AdaptivePolicyGroupId = types.StringValue(value.String())
		} else {
			data.AdaptivePolicyGroupId = types.StringNull()
		}
		(*parent).VlanNames[i] = data
	}
}

// End of section. //template:end fromBodyPartial
