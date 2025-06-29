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

type DataSourceNetworkPoliciesByClient struct {
	NetworkId types.String                             `tfsdk:"network_id"`
	Items     []DataSourceNetworkPoliciesByClientItems `tfsdk:"items"`
}

type DataSourceNetworkPoliciesByClientItems struct {
	Id       types.String                                `tfsdk:"id"`
	ClientId types.String                                `tfsdk:"client_id"`
	Name     types.String                                `tfsdk:"name"`
	Assigned []DataSourceNetworkPoliciesByClientAssigned `tfsdk:"assigned"`
}

type DataSourceNetworkPoliciesByClientAssigned struct {
	GroupPolicyId types.String                                    `tfsdk:"group_policy_id"`
	Name          types.String                                    `tfsdk:"name"`
	Type          types.String                                    `tfsdk:"type"`
	Ssid          []DataSourceNetworkPoliciesByClientAssignedSsid `tfsdk:"ssid"`
}

type DataSourceNetworkPoliciesByClientAssignedSsid struct {
	SsidNumber types.Int64 `tfsdk:"ssid_number"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworkPoliciesByClient) getPath() string {
	return fmt.Sprintf("/networks/%v/policies/byClient", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkPoliciesByClient) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceNetworkPoliciesByClientItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceNetworkPoliciesByClientItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("clientId"); value.Exists() && value.Value() != nil {
			data.ClientId = types.StringValue(value.String())
		} else {
			data.ClientId = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("assigned"); value.Exists() && value.Value() != nil {
			data.Assigned = make([]DataSourceNetworkPoliciesByClientAssigned, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceNetworkPoliciesByClientAssigned{}
				if value := res.Get("groupPolicyId"); value.Exists() && value.Value() != nil {
					data.GroupPolicyId = types.StringValue(value.String())
				} else {
					data.GroupPolicyId = types.StringNull()
				}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("type"); value.Exists() && value.Value() != nil {
					data.Type = types.StringValue(value.String())
				} else {
					data.Type = types.StringNull()
				}
				if value := res.Get("ssid"); value.Exists() && value.Value() != nil {
					data.Ssid = make([]DataSourceNetworkPoliciesByClientAssignedSsid, 0)
					value.ForEach(func(k, res gjson.Result) bool {
						parent := &data
						data := DataSourceNetworkPoliciesByClientAssignedSsid{}
						if value := res.Get("ssidNumber"); value.Exists() && value.Value() != nil {
							data.SsidNumber = types.Int64Value(value.Int())
						} else {
							data.SsidNumber = types.Int64Null()
						}
						(*parent).Ssid = append((*parent).Ssid, data)
						return true
					})
				}
				(*parent).Assigned = append((*parent).Assigned, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
