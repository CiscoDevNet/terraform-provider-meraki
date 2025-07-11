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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceApplianceTrafficShapingCustomPerformanceClasses struct {
	Id             types.String                                                   `tfsdk:"id"`
	OrganizationId types.String                                                   `tfsdk:"organization_id"`
	NetworkId      types.String                                                   `tfsdk:"network_id"`
	Items          []ResourceApplianceTrafficShapingCustomPerformanceClassesItems `tfsdk:"items"`
}

type ResourceApplianceTrafficShapingCustomPerformanceClassesItems struct {
	Id                types.String `tfsdk:"id"`
	MaxJitter         types.Int64  `tfsdk:"max_jitter"`
	MaxLatency        types.Int64  `tfsdk:"max_latency"`
	MaxLossPercentage types.Int64  `tfsdk:"max_loss_percentage"`
	Name              types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceApplianceTrafficShapingCustomPerformanceClasses) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/trafficShaping/customPerformanceClasses", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceApplianceTrafficShapingCustomPerformanceClassesItems) toBody(ctx context.Context, state ResourceApplianceTrafficShapingCustomPerformanceClassesItems) string {
	body := ""
	if !data.MaxJitter.IsNull() {
		body, _ = sjson.Set(body, "maxJitter", data.MaxJitter.ValueInt64())
	}
	if !data.MaxLatency.IsNull() {
		body, _ = sjson.Set(body, "maxLatency", data.MaxLatency.ValueInt64())
	}
	if !data.MaxLossPercentage.IsNull() {
		body, _ = sjson.Set(body, "maxLossPercentage", data.MaxLossPercentage.ValueInt64())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceApplianceTrafficShapingCustomPerformanceClasses) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceApplianceTrafficShapingCustomPerformanceClassesItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceApplianceTrafficShapingCustomPerformanceClassesItems{}
		if value := res.Get("maxJitter"); value.Exists() && value.Value() != nil {
			data.MaxJitter = types.Int64Value(value.Int())
		} else {
			data.MaxJitter = types.Int64Null()
		}
		if value := res.Get("maxLatency"); value.Exists() && value.Value() != nil {
			data.MaxLatency = types.Int64Value(value.Int())
		} else {
			data.MaxLatency = types.Int64Null()
		}
		if value := res.Get("maxLossPercentage"); value.Exists() && value.Value() != nil {
			data.MaxLossPercentage = types.Int64Value(value.Int())
		} else {
			data.MaxLossPercentage = types.Int64Null()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	index := 0
	res.ForEach(func(k, res gjson.Result) bool {
		data.Items[index].Id = types.StringValue(res.Get("customPerformanceClassId").String())
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
func (data *ResourceApplianceTrafficShapingCustomPerformanceClasses) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
				if v.Get("customPerformanceClassId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
		if value := res.Get("maxJitter"); value.Exists() && !data.MaxJitter.IsNull() {
			data.MaxJitter = types.Int64Value(value.Int())
		} else {
			data.MaxJitter = types.Int64Null()
		}
		if value := res.Get("maxLatency"); value.Exists() && !data.MaxLatency.IsNull() {
			data.MaxLatency = types.Int64Value(value.Int())
		} else {
			data.MaxLatency = types.Int64Null()
		}
		if value := res.Get("maxLossPercentage"); value.Exists() && !data.MaxLossPercentage.IsNull() {
			data.MaxLossPercentage = types.Int64Value(value.Int())
		} else {
			data.MaxLossPercentage = types.Int64Null()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceApplianceTrafficShapingCustomPerformanceClasses) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceApplianceTrafficShapingCustomPerformanceClasses) fromBodyImport(ctx context.Context, res meraki.Res) {
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
				if v.Get("customPerformanceClassId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
		if value := res.Get("maxJitter"); value.Exists() && value.Value() != nil {
			data.MaxJitter = types.Int64Value(value.Int())
		} else {
			data.MaxJitter = types.Int64Null()
		}
		if value := res.Get("maxLatency"); value.Exists() && value.Value() != nil {
			data.MaxLatency = types.Int64Value(value.Int())
		} else {
			data.MaxLatency = types.Int64Null()
		}
		if value := res.Get("maxLossPercentage"); value.Exists() && value.Value() != nil {
			data.MaxLossPercentage = types.Int64Value(value.Int())
		} else {
			data.MaxLossPercentage = types.Int64Null()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceApplianceTrafficShapingCustomPerformanceClasses) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceApplianceTrafficShapingCustomPerformanceClasses) hasChanges(ctx context.Context, state *ResourceApplianceTrafficShapingCustomPerformanceClasses, id string) bool {
	hasChanges := false

	item := ResourceApplianceTrafficShapingCustomPerformanceClassesItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceApplianceTrafficShapingCustomPerformanceClassesItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.MaxJitter.Equal(stateItem.MaxJitter) {
		hasChanges = true
	}
	if !item.MaxLatency.Equal(stateItem.MaxLatency) {
		hasChanges = true
	}
	if !item.MaxLossPercentage.Equal(stateItem.MaxLossPercentage) {
		hasChanges = true
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
