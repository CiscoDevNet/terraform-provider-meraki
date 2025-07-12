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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceAppliancePrefixDelegatedStatics struct {
	Id             types.String                                   `tfsdk:"id"`
	OrganizationId types.String                                   `tfsdk:"organization_id"`
	NetworkId      types.String                                   `tfsdk:"network_id"`
	Items          []ResourceAppliancePrefixDelegatedStaticsItems `tfsdk:"items"`
}

type ResourceAppliancePrefixDelegatedStaticsItems struct {
	Id               types.String `tfsdk:"id"`
	Description      types.String `tfsdk:"description"`
	Prefix           types.String `tfsdk:"prefix"`
	OriginType       types.String `tfsdk:"origin_type"`
	OriginInterfaces types.List   `tfsdk:"origin_interfaces"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceAppliancePrefixDelegatedStatics) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/prefixes/delegated/statics", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceAppliancePrefixDelegatedStaticsItems) toBody(ctx context.Context, state ResourceAppliancePrefixDelegatedStaticsItems) string {
	body := ""
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Prefix.IsNull() {
		body, _ = sjson.Set(body, "prefix", data.Prefix.ValueString())
	}
	if !data.OriginType.IsNull() {
		body, _ = sjson.Set(body, "origin.type", data.OriginType.ValueString())
	}
	if !data.OriginInterfaces.IsNull() {
		var values []string
		data.OriginInterfaces.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "origin.interfaces", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceAppliancePrefixDelegatedStatics) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceAppliancePrefixDelegatedStaticsItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceAppliancePrefixDelegatedStaticsItems{}
		if value := res.Get("description"); value.Exists() && value.Value() != nil {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("prefix"); value.Exists() && value.Value() != nil {
			data.Prefix = types.StringValue(value.String())
		} else {
			data.Prefix = types.StringNull()
		}
		if value := res.Get("origin.type"); value.Exists() && value.Value() != nil {
			data.OriginType = types.StringValue(value.String())
		} else {
			data.OriginType = types.StringNull()
		}
		if value := res.Get("origin.interfaces"); value.Exists() && value.Value() != nil {
			data.OriginInterfaces = helpers.GetStringList(value.Array())
		} else {
			data.OriginInterfaces = types.ListNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	index := 0
	res.ForEach(func(k, res gjson.Result) bool {
		data.Items[index].Id = types.StringValue(res.Get("staticDelegatedPrefixId").String())
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
func (data *ResourceAppliancePrefixDelegatedStatics) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("staticDelegatedPrefixId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("prefix"); value.Exists() && !data.Prefix.IsNull() {
			data.Prefix = types.StringValue(value.String())
		} else {
			data.Prefix = types.StringNull()
		}
		if value := res.Get("origin.type"); value.Exists() && !data.OriginType.IsNull() {
			data.OriginType = types.StringValue(value.String())
		} else {
			data.OriginType = types.StringNull()
		}
		if value := res.Get("origin.interfaces"); value.Exists() && !data.OriginInterfaces.IsNull() {
			data.OriginInterfaces = helpers.GetStringList(value.Array())
		} else {
			data.OriginInterfaces = types.ListNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceAppliancePrefixDelegatedStatics) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceAppliancePrefixDelegatedStatics) fromBodyImport(ctx context.Context, res meraki.Res) {
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
				if v.Get("staticDelegatedPrefixId").String() == (*parent).Items[i].Id.ValueString() {
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
		if value := res.Get("description"); value.Exists() && value.Value() != nil {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("prefix"); value.Exists() && value.Value() != nil {
			data.Prefix = types.StringValue(value.String())
		} else {
			data.Prefix = types.StringNull()
		}
		if value := res.Get("origin.type"); value.Exists() && value.Value() != nil {
			data.OriginType = types.StringValue(value.String())
		} else {
			data.OriginType = types.StringNull()
		}
		if value := res.Get("origin.interfaces"); value.Exists() && value.Value() != nil {
			data.OriginInterfaces = helpers.GetStringList(value.Array())
		} else {
			data.OriginInterfaces = types.ListNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("Import, removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceAppliancePrefixDelegatedStatics) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceAppliancePrefixDelegatedStatics) hasChanges(ctx context.Context, state *ResourceAppliancePrefixDelegatedStatics, id string) bool {
	hasChanges := false

	item := ResourceAppliancePrefixDelegatedStaticsItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceAppliancePrefixDelegatedStaticsItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.Description.Equal(stateItem.Description) {
		hasChanges = true
	}
	if !item.Prefix.Equal(stateItem.Prefix) {
		hasChanges = true
	}
	if !item.OriginType.Equal(stateItem.OriginType) {
		hasChanges = true
	}
	if !item.OriginInterfaces.Equal(stateItem.OriginInterfaces) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
