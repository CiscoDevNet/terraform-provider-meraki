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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type CameraRoles struct {
	OrganizationId types.String       `tfsdk:"organization_id"`
	Items          []CameraRolesItems `tfsdk:"items"`
}

type CameraRolesItems struct {
	Id                types.String                   `tfsdk:"id"`
	Name              types.String                   `tfsdk:"name"`
	AppliedOnDevices  []CameraRolesAppliedOnDevices  `tfsdk:"applied_on_devices"`
	AppliedOnNetworks []CameraRolesAppliedOnNetworks `tfsdk:"applied_on_networks"`
	AppliedOrgWide    []CameraRolesAppliedOrgWide    `tfsdk:"applied_org_wide"`
}

type CameraRolesAppliedOnDevices struct {
	Id                types.String `tfsdk:"id"`
	InNetworksWithId  types.String `tfsdk:"in_networks_with_id"`
	InNetworksWithTag types.String `tfsdk:"in_networks_with_tag"`
	PermissionScopeId types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

type CameraRolesAppliedOnNetworks struct {
	Id                types.String `tfsdk:"id"`
	PermissionScopeId types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

type CameraRolesAppliedOrgWide struct {
	PermissionScopeId types.String `tfsdk:"permission_scope_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CameraRoles) getPath() string {
	return fmt.Sprintf("/organizations/%v/camera/roles", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CameraRoles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]CameraRolesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := CameraRolesItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("appliedOnDevices"); value.Exists() && value.Value() != nil {
			data.AppliedOnDevices = make([]CameraRolesAppliedOnDevices, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := CameraRolesAppliedOnDevices{}
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
			data.AppliedOnNetworks = make([]CameraRolesAppliedOnNetworks, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := CameraRolesAppliedOnNetworks{}
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
			data.AppliedOrgWide = make([]CameraRolesAppliedOrgWide, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := CameraRolesAppliedOrgWide{}
				if value := res.Get("permissionScopeId"); value.Exists() && value.Value() != nil {
					data.PermissionScopeId = types.StringValue(value.String())
				} else {
					data.PermissionScopeId = types.StringNull()
				}
				(*parent).AppliedOrgWide = append((*parent).AppliedOrgWide, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
