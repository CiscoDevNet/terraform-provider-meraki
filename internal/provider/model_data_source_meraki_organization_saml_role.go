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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSourceOrganizationSAMLRole - used only by the data source, always including every
// attribute (unlike model_resource.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model_resource.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model_resource.go's resource-side type names.

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

type DataSourceOrganizationSAMLRole struct {
	Id             types.String                             `tfsdk:"id"`
	OrganizationId types.String                             `tfsdk:"organization_id"`
	OrgAccess      types.String                             `tfsdk:"org_access"`
	Role           types.String                             `tfsdk:"role"`
	Networks       []DataSourceOrganizationSAMLRoleNetworks `tfsdk:"networks"`
	Tags           []DataSourceOrganizationSAMLRoleTags     `tfsdk:"tags"`
}

type DataSourceOrganizationSAMLRoleNetworks struct {
	Access types.String `tfsdk:"access"`
	Id     types.String `tfsdk:"id"`
}

type DataSourceOrganizationSAMLRoleTags struct {
	Access types.String `tfsdk:"access"`
	Tag    types.String `tfsdk:"tag"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationSAMLRole) getPath() string {
	return fmt.Sprintf("/organizations/%v/samlRoles", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationSAMLRole) fromBody(ctx context.Context, res meraki.Res) {
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
		data.Networks = make([]DataSourceOrganizationSAMLRoleNetworks, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceOrganizationSAMLRoleNetworks{}
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
		data.Tags = make([]DataSourceOrganizationSAMLRoleTags, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceOrganizationSAMLRoleTags{}
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
}

// End of section. //template:end fromBody
