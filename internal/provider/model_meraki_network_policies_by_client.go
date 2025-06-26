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

type NetworkPoliciesByClient struct {
	Id        types.String                      `tfsdk:"id"`
	NetworkId types.String                      `tfsdk:"network_id"`
	ClientId  types.String                      `tfsdk:"client_id"`
	Name      types.String                      `tfsdk:"name"`
	Assigned  []NetworkPoliciesByClientAssigned `tfsdk:"assigned"`
}

type NetworkPoliciesByClientAssigned struct {
	GroupPolicyId types.String                          `tfsdk:"group_policy_id"`
	Name          types.String                          `tfsdk:"name"`
	Type          types.String                          `tfsdk:"type"`
	Ssid          []NetworkPoliciesByClientAssignedSsid `tfsdk:"ssid"`
}

type NetworkPoliciesByClientAssignedSsid struct {
	SsidNumber types.Int64 `tfsdk:"ssid_number"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkPoliciesByClient) getPath() string {
	return fmt.Sprintf("/networks/%v/policies/byClient", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkPoliciesByClient) fromBody(ctx context.Context, res meraki.Res) {
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
		data.Assigned = make([]NetworkPoliciesByClientAssigned, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkPoliciesByClientAssigned{}
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
				data.Ssid = make([]NetworkPoliciesByClientAssignedSsid, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := NetworkPoliciesByClientAssignedSsid{}
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
}

// End of section. //template:end fromBody
