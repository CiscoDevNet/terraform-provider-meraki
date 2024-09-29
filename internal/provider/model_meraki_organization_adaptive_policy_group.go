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

type OrganizationAdaptivePolicyGroup struct {
	Id             types.String                                   `tfsdk:"id"`
	OrganizationId types.String                                   `tfsdk:"organization_id"`
	Description    types.String                                   `tfsdk:"description"`
	Name           types.String                                   `tfsdk:"name"`
	Sgt            types.Int64                                    `tfsdk:"sgt"`
	PolicyObjects  []OrganizationAdaptivePolicyGroupPolicyObjects `tfsdk:"policy_objects"`
}

type OrganizationAdaptivePolicyGroupPolicyObjects struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationAdaptivePolicyGroup) getPath() string {
	return fmt.Sprintf("/organizations/%v/adaptivePolicy/groups", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationAdaptivePolicyGroup) toBody(ctx context.Context, state OrganizationAdaptivePolicyGroup) string {
	body := ""
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Sgt.IsNull() {
		body, _ = sjson.Set(body, "sgt", data.Sgt.ValueInt64())
	}
	if len(data.PolicyObjects) > 0 {
		body, _ = sjson.Set(body, "policyObjects", []interface{}{})
		for _, item := range data.PolicyObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "policyObjects.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationAdaptivePolicyGroup) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && value.Value() != nil {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("sgt"); value.Exists() && value.Value() != nil {
		data.Sgt = types.Int64Value(value.Int())
	} else {
		data.Sgt = types.Int64Null()
	}
	if value := res.Get("policyObjects"); value.Exists() && value.Value() != nil {
		data.PolicyObjects = make([]OrganizationAdaptivePolicyGroupPolicyObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationAdaptivePolicyGroupPolicyObjects{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).PolicyObjects = append((*parent).PolicyObjects, data)
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
func (data *OrganizationAdaptivePolicyGroup) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("sgt"); value.Exists() && !data.Sgt.IsNull() {
		data.Sgt = types.Int64Value(value.Int())
	} else {
		data.Sgt = types.Int64Null()
	}
	for i := 0; i < len(data.PolicyObjects); i++ {
		keys := [...]string{"id", "name"}
		keyValues := [...]string{data.PolicyObjects[i].Id.ValueString(), data.PolicyObjects[i].Name.ValueString()}

		parent := &data
		data := (*parent).PolicyObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("policyObjects").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing PolicyObjects[%d] = %+v",
				i,
				(*parent).PolicyObjects[i],
			))
			(*parent).PolicyObjects = slices.Delete((*parent).PolicyObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).PolicyObjects[i] = data
	}
}

// End of section. //template:end fromBodyPartial
