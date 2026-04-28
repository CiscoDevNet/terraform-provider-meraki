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

type CameraRole struct {
	Id                types.String                  `tfsdk:"id"`
	OrganizationId    types.String                  `tfsdk:"organization_id"`
	Name              types.String                  `tfsdk:"name"`
	AppliedOnDevices  []CameraRoleAppliedOnDevices  `tfsdk:"applied_on_devices"`
	AppliedOnNetworks []CameraRoleAppliedOnNetworks `tfsdk:"applied_on_networks"`
	AppliedOrgWide    []CameraRoleAppliedOrgWide    `tfsdk:"applied_org_wide"`
}

type CameraRoleAppliedOnDevices struct {
	Id                types.String `tfsdk:"id"`
	InNetworksWithId  types.String `tfsdk:"in_networks_with_id"`
	InNetworksWithTag types.String `tfsdk:"in_networks_with_tag"`
	PermissionScopeId types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

type CameraRoleAppliedOnNetworks struct {
	Id                types.String `tfsdk:"id"`
	PermissionScopeId types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

type CameraRoleAppliedOrgWide struct {
	PermissionScopeId types.String `tfsdk:"permission_scope_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CameraRole) getPath() string {
	return fmt.Sprintf("/organizations/%v/camera/roles", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CameraRole) toBody(ctx context.Context, state CameraRole) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if len(data.AppliedOnDevices) > 0 {
		body, _ = sjson.Set(body, "appliedOnDevices", []interface{}{})
		for _, item := range data.AppliedOnDevices {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.InNetworksWithId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "inNetworksWithId", item.InNetworksWithId.ValueString())
			}
			if !item.InNetworksWithTag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "inNetworksWithTag", item.InNetworksWithTag.ValueString())
			}
			if !item.PermissionScopeId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "permissionScopeId", item.PermissionScopeId.ValueString())
			}
			if !item.Tag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tag", item.Tag.ValueString())
			}
			body, _ = sjson.SetRaw(body, "appliedOnDevices.-1", itemBody)
		}
	}
	if len(data.AppliedOnNetworks) > 0 {
		body, _ = sjson.Set(body, "appliedOnNetworks", []interface{}{})
		for _, item := range data.AppliedOnNetworks {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.PermissionScopeId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "permissionScopeId", item.PermissionScopeId.ValueString())
			}
			if !item.Tag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tag", item.Tag.ValueString())
			}
			body, _ = sjson.SetRaw(body, "appliedOnNetworks.-1", itemBody)
		}
	}
	if len(data.AppliedOrgWide) > 0 {
		body, _ = sjson.Set(body, "appliedOrgWide", []interface{}{})
		for _, item := range data.AppliedOrgWide {
			itemBody := ""
			if !item.PermissionScopeId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "permissionScopeId", item.PermissionScopeId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "appliedOrgWide.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CameraRole) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("appliedOnDevices"); value.Exists() && value.Value() != nil {
		data.AppliedOnDevices = make([]CameraRoleAppliedOnDevices, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := CameraRoleAppliedOnDevices{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("inNetworksWithId"); value.Exists() && value.Value() != nil {
				data.InNetworksWithId = types.StringValue(value.String())
			} else {
				data.InNetworksWithId = types.StringNull()
			}
			if value := res.Get("inNetworksWithTag"); value.Exists() && value.Value() != nil {
				data.InNetworksWithTag = types.StringValue(value.String())
			} else {
				data.InNetworksWithTag = types.StringNull()
			}
			if value := res.Get("permissionScopeId"); value.Exists() && value.Value() != nil {
				data.PermissionScopeId = types.StringValue(value.String())
			} else {
				data.PermissionScopeId = types.StringNull()
			}
			if value := res.Get("tag"); value.Exists() && value.Value() != nil {
				data.Tag = types.StringValue(value.String())
			} else {
				data.Tag = types.StringNull()
			}
			(*parent).AppliedOnDevices = append((*parent).AppliedOnDevices, data)
			return true
		})
	}
	if value := res.Get("appliedOnNetworks"); value.Exists() && value.Value() != nil {
		data.AppliedOnNetworks = make([]CameraRoleAppliedOnNetworks, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := CameraRoleAppliedOnNetworks{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("permissionScopeId"); value.Exists() && value.Value() != nil {
				data.PermissionScopeId = types.StringValue(value.String())
			} else {
				data.PermissionScopeId = types.StringNull()
			}
			if value := res.Get("tag"); value.Exists() && value.Value() != nil {
				data.Tag = types.StringValue(value.String())
			} else {
				data.Tag = types.StringNull()
			}
			(*parent).AppliedOnNetworks = append((*parent).AppliedOnNetworks, data)
			return true
		})
	}
	if value := res.Get("appliedOrgWide"); value.Exists() && value.Value() != nil {
		data.AppliedOrgWide = make([]CameraRoleAppliedOrgWide, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := CameraRoleAppliedOrgWide{}
			if value := res.Get("permissionScopeId"); value.Exists() && value.Value() != nil {
				data.PermissionScopeId = types.StringValue(value.String())
			} else {
				data.PermissionScopeId = types.StringNull()
			}
			(*parent).AppliedOrgWide = append((*parent).AppliedOrgWide, data)
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
func (data *CameraRole) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	for i := 0; i < len(data.AppliedOnDevices); i++ {
		keys := [...]string{"id", "inNetworksWithId", "inNetworksWithTag", "permissionScopeId", "tag"}
		keyValues := [...]string{data.AppliedOnDevices[i].Id.ValueString(), data.AppliedOnDevices[i].InNetworksWithId.ValueString(), data.AppliedOnDevices[i].InNetworksWithTag.ValueString(), data.AppliedOnDevices[i].PermissionScopeId.ValueString(), data.AppliedOnDevices[i].Tag.ValueString()}

		parent := &data
		data := (*parent).AppliedOnDevices[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("appliedOnDevices").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AppliedOnDevices[%d] = %+v",
				i,
				(*parent).AppliedOnDevices[i],
			))
			(*parent).AppliedOnDevices = slices.Delete((*parent).AppliedOnDevices, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("inNetworksWithId"); value.Exists() && !data.InNetworksWithId.IsNull() {
			data.InNetworksWithId = types.StringValue(value.String())
		} else {
			data.InNetworksWithId = types.StringNull()
		}
		if value := res.Get("inNetworksWithTag"); value.Exists() && !data.InNetworksWithTag.IsNull() {
			data.InNetworksWithTag = types.StringValue(value.String())
		} else {
			data.InNetworksWithTag = types.StringNull()
		}
		if value := res.Get("permissionScopeId"); value.Exists() && !data.PermissionScopeId.IsNull() {
			data.PermissionScopeId = types.StringValue(value.String())
		} else {
			data.PermissionScopeId = types.StringNull()
		}
		if value := res.Get("tag"); value.Exists() && !data.Tag.IsNull() {
			data.Tag = types.StringValue(value.String())
		} else {
			data.Tag = types.StringNull()
		}
		(*parent).AppliedOnDevices[i] = data
	}
	for i := 0; i < len(data.AppliedOnNetworks); i++ {
		keys := [...]string{"id", "permissionScopeId", "tag"}
		keyValues := [...]string{data.AppliedOnNetworks[i].Id.ValueString(), data.AppliedOnNetworks[i].PermissionScopeId.ValueString(), data.AppliedOnNetworks[i].Tag.ValueString()}

		parent := &data
		data := (*parent).AppliedOnNetworks[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("appliedOnNetworks").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AppliedOnNetworks[%d] = %+v",
				i,
				(*parent).AppliedOnNetworks[i],
			))
			(*parent).AppliedOnNetworks = slices.Delete((*parent).AppliedOnNetworks, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("permissionScopeId"); value.Exists() && !data.PermissionScopeId.IsNull() {
			data.PermissionScopeId = types.StringValue(value.String())
		} else {
			data.PermissionScopeId = types.StringNull()
		}
		if value := res.Get("tag"); value.Exists() && !data.Tag.IsNull() {
			data.Tag = types.StringValue(value.String())
		} else {
			data.Tag = types.StringNull()
		}
		(*parent).AppliedOnNetworks[i] = data
	}
	for i := 0; i < len(data.AppliedOrgWide); i++ {
		keys := [...]string{"permissionScopeId"}
		keyValues := [...]string{data.AppliedOrgWide[i].PermissionScopeId.ValueString()}

		parent := &data
		data := (*parent).AppliedOrgWide[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("appliedOrgWide").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AppliedOrgWide[%d] = %+v",
				i,
				(*parent).AppliedOrgWide[i],
			))
			(*parent).AppliedOrgWide = slices.Delete((*parent).AppliedOrgWide, i, i+1)
			i--

			continue
		}
		if value := res.Get("permissionScopeId"); value.Exists() && !data.PermissionScopeId.IsNull() {
			data.PermissionScopeId = types.StringValue(value.String())
		} else {
			data.PermissionScopeId = types.StringNull()
		}
		(*parent).AppliedOrgWide[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CameraRole) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data CameraRole) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
