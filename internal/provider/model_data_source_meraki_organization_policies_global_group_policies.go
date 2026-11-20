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

type DataSourceOrganizationPoliciesGlobalGroupPolicies struct {
	OrganizationId types.String                                             `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationPoliciesGlobalGroupPoliciesItems `tfsdk:"items"`
}

type DataSourceOrganizationPoliciesGlobalGroupPoliciesItems struct {
	Id          types.String `tfsdk:"id"`
	Description types.String `tfsdk:"description"`
	Name        types.String `tfsdk:"name"`
	GroupNumber types.Int64  `tfsdk:"group_number"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationPoliciesGlobalGroupPolicies) getPath() string {
	return fmt.Sprintf("/organizations/%v/policies/global/group/policies", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationPoliciesGlobalGroupPolicies) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationPoliciesGlobalGroupPoliciesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationPoliciesGlobalGroupPoliciesItems{}
		data.Id = types.StringValue(res.Get("policyId").String())
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
		if value := res.Get("group.number"); value.Exists() && value.Value() != nil {
			data.GroupNumber = types.Int64Value(value.Int())
		} else {
			data.GroupNumber = types.Int64Null()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
