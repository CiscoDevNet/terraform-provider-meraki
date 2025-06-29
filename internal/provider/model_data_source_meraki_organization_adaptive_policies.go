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

type DataSourceOrganizationAdaptivePolicies struct {
	OrganizationId types.String                                  `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationAdaptivePoliciesItems `tfsdk:"items"`
}

type DataSourceOrganizationAdaptivePoliciesItems struct {
	Id                   types.String                                 `tfsdk:"id"`
	LastEntryRule        types.String                                 `tfsdk:"last_entry_rule"`
	DestinationGroupId   types.String                                 `tfsdk:"destination_group_id"`
	DestinationGroupName types.String                                 `tfsdk:"destination_group_name"`
	DestinationGroupSgt  types.Int64                                  `tfsdk:"destination_group_sgt"`
	SourceGroupId        types.String                                 `tfsdk:"source_group_id"`
	SourceGroupName      types.String                                 `tfsdk:"source_group_name"`
	SourceGroupSgt       types.Int64                                  `tfsdk:"source_group_sgt"`
	Acls                 []DataSourceOrganizationAdaptivePoliciesAcls `tfsdk:"acls"`
}

type DataSourceOrganizationAdaptivePoliciesAcls struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationAdaptivePolicies) getPath() string {
	return fmt.Sprintf("/organizations/%v/adaptivePolicy/policies", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationAdaptivePolicies) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationAdaptivePoliciesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationAdaptivePoliciesItems{}
		data.Id = types.StringValue(res.Get("adaptivePolicyId").String())
		if value := res.Get("lastEntryRule"); value.Exists() && value.Value() != nil {
			data.LastEntryRule = types.StringValue(value.String())
		} else {
			data.LastEntryRule = types.StringNull()
		}
		if value := res.Get("destinationGroup.id"); value.Exists() && value.Value() != nil {
			data.DestinationGroupId = types.StringValue(value.String())
		} else {
			data.DestinationGroupId = types.StringNull()
		}
		if value := res.Get("destinationGroup.name"); value.Exists() && value.Value() != nil {
			data.DestinationGroupName = types.StringValue(value.String())
		} else {
			data.DestinationGroupName = types.StringNull()
		}
		if value := res.Get("destinationGroup.sgt"); value.Exists() && value.Value() != nil {
			data.DestinationGroupSgt = types.Int64Value(value.Int())
		} else {
			data.DestinationGroupSgt = types.Int64Null()
		}
		if value := res.Get("sourceGroup.id"); value.Exists() && value.Value() != nil {
			data.SourceGroupId = types.StringValue(value.String())
		} else {
			data.SourceGroupId = types.StringNull()
		}
		if value := res.Get("sourceGroup.name"); value.Exists() && value.Value() != nil {
			data.SourceGroupName = types.StringValue(value.String())
		} else {
			data.SourceGroupName = types.StringNull()
		}
		if value := res.Get("sourceGroup.sgt"); value.Exists() && value.Value() != nil {
			data.SourceGroupSgt = types.Int64Value(value.Int())
		} else {
			data.SourceGroupSgt = types.Int64Null()
		}
		if value := res.Get("acls"); value.Exists() && value.Value() != nil {
			data.Acls = make([]DataSourceOrganizationAdaptivePoliciesAcls, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationAdaptivePoliciesAcls{}
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
				(*parent).Acls = append((*parent).Acls, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
