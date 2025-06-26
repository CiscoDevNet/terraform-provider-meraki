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

type DataSourceOrganizationSAMLRoles struct {
	OrganizationId types.String                           `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationSAMLRolesItems `tfsdk:"items"`
}

type DataSourceOrganizationSAMLRolesItems struct {
	Id        types.String                              `tfsdk:"id"`
	OrgAccess types.String                              `tfsdk:"org_access"`
	Role      types.String                              `tfsdk:"role"`
	Networks  []DataSourceOrganizationSAMLRolesNetworks `tfsdk:"networks"`
	Tags      []DataSourceOrganizationSAMLRolesTags     `tfsdk:"tags"`
}

type DataSourceOrganizationSAMLRolesNetworks struct {
	Access types.String `tfsdk:"access"`
	Id     types.String `tfsdk:"id"`
}

type DataSourceOrganizationSAMLRolesTags struct {
	Access types.String `tfsdk:"access"`
	Tag    types.String `tfsdk:"tag"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationSAMLRoles) getPath() string {
	return fmt.Sprintf("/organizations/%v/samlRoles", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationSAMLRoles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationSAMLRolesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationSAMLRolesItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("orgAccess"); value.Exists() && value.Value() != nil {
			data.OrgAccess = types.StringValue(value.String())
		} else {
			data.OrgAccess = types.StringNull()
		}
		if value := res.Get("role"); value.Exists() && value.Value() != nil {
			data.Role = types.StringValue(value.String())
		} else {
			data.Role = types.StringNull()
		}
		if value := res.Get("networks"); value.Exists() && value.Value() != nil {
			data.Networks = make([]DataSourceOrganizationSAMLRolesNetworks, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationSAMLRolesNetworks{}
				if value := res.Get("access"); value.Exists() && value.Value() != nil {
					data.Access = types.StringValue(value.String())
				} else {
					data.Access = types.StringNull()
				}
				if value := res.Get("id"); value.Exists() && value.Value() != nil {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).Networks = append((*parent).Networks, data)
				return true
			})
		}
		if value := res.Get("tags"); value.Exists() && value.Value() != nil {
			data.Tags = make([]DataSourceOrganizationSAMLRolesTags, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationSAMLRolesTags{}
				if value := res.Get("access"); value.Exists() && value.Value() != nil {
					data.Access = types.StringValue(value.String())
				} else {
					data.Access = types.StringNull()
				}
				if value := res.Get("tag"); value.Exists() && value.Value() != nil {
					data.Tag = types.StringValue(value.String())
				} else {
					data.Tag = types.StringNull()
				}
				(*parent).Tags = append((*parent).Tags, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
