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

type Admin struct {
	Id             types.String    `tfsdk:"id"`
	OrganizationId types.String    `tfsdk:"organization_id"`
	Email          types.String    `tfsdk:"email"`
	Name           types.String    `tfsdk:"name"`
	OrgAccess      types.String    `tfsdk:"org_access"`
	Networks       []AdminNetworks `tfsdk:"networks"`
	Tags           []AdminTags     `tfsdk:"tags"`
}

type AdminNetworks struct {
	Access types.String `tfsdk:"access"`
	Id     types.String `tfsdk:"id"`
}

type AdminTags struct {
	Access types.String `tfsdk:"access"`
	Tag    types.String `tfsdk:"tag"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data Admin) getPath() string {
	return fmt.Sprintf("/organizations/%v/admins", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data Admin) toBody(ctx context.Context, state Admin) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Email.IsNull() {
		body, _ = sjson.Set(body, "email", data.Email.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.OrgAccess.IsNull() {
		body, _ = sjson.Set(body, "orgAccess", data.OrgAccess.ValueString())
	}
	if len(data.Networks) > 0 {
		body, _ = sjson.Set(body, "networks", []interface{}{})
		for _, item := range data.Networks {
			itemBody := ""
			if !item.Access.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "access", item.Access.ValueString())
			}
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "networks.-1", itemBody)
		}
	}
	if len(data.Tags) > 0 {
		body, _ = sjson.Set(body, "tags", []interface{}{})
		for _, item := range data.Tags {
			itemBody := ""
			if !item.Access.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "access", item.Access.ValueString())
			}
			if !item.Tag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tag", item.Tag.ValueString())
			}
			body, _ = sjson.SetRaw(body, "tags.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *Admin) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("email"); value.Exists() {
		data.Email = types.StringValue(value.String())
	} else {
		data.Email = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("orgAccess"); value.Exists() {
		data.OrgAccess = types.StringValue(value.String())
	} else {
		data.OrgAccess = types.StringNull()
	}
	if value := res.Get("networks"); value.Exists() {
		data.Networks = make([]AdminNetworks, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AdminNetworks{}
			if value := res.Get("access"); value.Exists() {
				data.Access = types.StringValue(value.String())
			} else {
				data.Access = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Networks = append((*parent).Networks, data)
			return true
		})
	}
	if value := res.Get("tags"); value.Exists() {
		data.Tags = make([]AdminTags, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := AdminTags{}
			if value := res.Get("access"); value.Exists() {
				data.Access = types.StringValue(value.String())
			} else {
				data.Access = types.StringNull()
			}
			if value := res.Get("tag"); value.Exists() {
				data.Tag = types.StringValue(value.String())
			} else {
				data.Tag = types.StringNull()
			}
			(*parent).Tags = append((*parent).Tags, data)
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
func (data *Admin) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("email"); value.Exists() && !data.Email.IsNull() {
		data.Email = types.StringValue(value.String())
	} else {
		data.Email = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("orgAccess"); value.Exists() && !data.OrgAccess.IsNull() {
		data.OrgAccess = types.StringValue(value.String())
	} else {
		data.OrgAccess = types.StringNull()
	}
	for i := 0; i < len(data.Networks); i++ {
		keys := [...]string{"access", "id"}
		keyValues := [...]string{data.Networks[i].Access.ValueString(), data.Networks[i].Id.ValueString()}

		parent := &data
		data := (*parent).Networks[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("networks").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Networks[%d] = %+v",
				i,
				(*parent).Networks[i],
			))
			(*parent).Networks = slices.Delete((*parent).Networks, i, i+1)
			i--

			continue
		}
		if value := res.Get("access"); value.Exists() && !data.Access.IsNull() {
			data.Access = types.StringValue(value.String())
		} else {
			data.Access = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).Networks[i] = data
	}
	for i := 0; i < len(data.Tags); i++ {
		keys := [...]string{"access", "tag"}
		keyValues := [...]string{data.Tags[i].Access.ValueString(), data.Tags[i].Tag.ValueString()}

		parent := &data
		data := (*parent).Tags[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("tags").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Tags[%d] = %+v",
				i,
				(*parent).Tags[i],
			))
			(*parent).Tags = slices.Delete((*parent).Tags, i, i+1)
			i--

			continue
		}
		if value := res.Get("access"); value.Exists() && !data.Access.IsNull() {
			data.Access = types.StringValue(value.String())
		} else {
			data.Access = types.StringNull()
		}
		if value := res.Get("tag"); value.Exists() && !data.Tag.IsNull() {
			data.Tag = types.StringValue(value.String())
		} else {
			data.Tag = types.StringNull()
		}
		(*parent).Tags[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *Admin) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
