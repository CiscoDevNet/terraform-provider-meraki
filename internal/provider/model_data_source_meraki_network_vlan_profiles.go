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

type DataSourceNetworkVLANProfiles struct {
	NetworkId types.String                         `tfsdk:"network_id"`
	Items     []DataSourceNetworkVLANProfilesItems `tfsdk:"items"`
}

type DataSourceNetworkVLANProfilesItems struct {
	Id         types.String                              `tfsdk:"id"`
	Iname      types.String                              `tfsdk:"iname"`
	Name       types.String                              `tfsdk:"name"`
	VlanGroups []DataSourceNetworkVLANProfilesVlanGroups `tfsdk:"vlan_groups"`
	VlanNames  []DataSourceNetworkVLANProfilesVlanNames  `tfsdk:"vlan_names"`
}

type DataSourceNetworkVLANProfilesVlanGroups struct {
	Name    types.String `tfsdk:"name"`
	VlanIds types.String `tfsdk:"vlan_ids"`
}

type DataSourceNetworkVLANProfilesVlanNames struct {
	Name                  types.String `tfsdk:"name"`
	VlanId                types.String `tfsdk:"vlan_id"`
	AdaptivePolicyGroupId types.String `tfsdk:"adaptive_policy_group_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworkVLANProfiles) getPath() string {
	return fmt.Sprintf("/networks/%v/vlanProfiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkVLANProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceNetworkVLANProfilesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceNetworkVLANProfilesItems{}
		data.Id = types.StringValue(res.Get("iname").String())
		if value := res.Get("iname"); value.Exists() && value.Value() != nil {
			data.Iname = types.StringValue(value.String())
		} else {
			data.Iname = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("vlanGroups"); value.Exists() && value.Value() != nil {
			data.VlanGroups = make([]DataSourceNetworkVLANProfilesVlanGroups, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceNetworkVLANProfilesVlanGroups{}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("vlanIds"); value.Exists() && value.Value() != nil {
					data.VlanIds = types.StringValue(value.String())
				} else {
					data.VlanIds = types.StringNull()
				}
				(*parent).VlanGroups = append((*parent).VlanGroups, data)
				return true
			})
		}
		if value := res.Get("vlanNames"); value.Exists() && value.Value() != nil {
			data.VlanNames = make([]DataSourceNetworkVLANProfilesVlanNames, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceNetworkVLANProfilesVlanNames{}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("vlanId"); value.Exists() && value.Value() != nil {
					data.VlanId = types.StringValue(value.String())
				} else {
					data.VlanId = types.StringNull()
				}
				if value := res.Get("adaptivePolicyGroup.id"); value.Exists() && value.Value() != nil {
					data.AdaptivePolicyGroupId = types.StringValue(value.String())
				} else {
					data.AdaptivePolicyGroupId = types.StringNull()
				}
				(*parent).VlanNames = append((*parent).VlanNames, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
