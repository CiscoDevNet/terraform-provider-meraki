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

type DataSourceOrganizationAdaptivePolicyGroups struct {
	OrganizationId types.String                                      `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationAdaptivePolicyGroupsItems `tfsdk:"items"`
}

type DataSourceOrganizationAdaptivePolicyGroupsItems struct {
	Id            types.String                                              `tfsdk:"id"`
	Description   types.String                                              `tfsdk:"description"`
	Name          types.String                                              `tfsdk:"name"`
	Sgt           types.Int64                                               `tfsdk:"sgt"`
	PolicyObjects []DataSourceOrganizationAdaptivePolicyGroupsPolicyObjects `tfsdk:"policy_objects"`
}

type DataSourceOrganizationAdaptivePolicyGroupsPolicyObjects struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationAdaptivePolicyGroups) getPath() string {
	return fmt.Sprintf("/organizations/%v/adaptivePolicy/groups", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationAdaptivePolicyGroups) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationAdaptivePolicyGroupsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationAdaptivePolicyGroupsItems{}
		data.Id = types.StringValue(res.Get("groupId").String())
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
			data.PolicyObjects = make([]DataSourceOrganizationAdaptivePolicyGroupsPolicyObjects, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationAdaptivePolicyGroupsPolicyObjects{}
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
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
