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

type NetworkVLANProfiles struct {
	NetworkId types.String               `tfsdk:"network_id"`
	Items     []NetworkVLANProfilesItems `tfsdk:"items"`
}

type NetworkVLANProfilesItems struct {
	Id         types.String                    `tfsdk:"id"`
	Iname      types.String                    `tfsdk:"iname"`
	Name       types.String                    `tfsdk:"name"`
	VlanGroups []NetworkVLANProfilesVlanGroups `tfsdk:"vlan_groups"`
	VlanNames  []NetworkVLANProfilesVlanNames  `tfsdk:"vlan_names"`
}

type NetworkVLANProfilesVlanGroups struct {
	Name    types.String `tfsdk:"name"`
	VlanIds types.String `tfsdk:"vlan_ids"`
}

type NetworkVLANProfilesVlanNames struct {
	Name                  types.String `tfsdk:"name"`
	VlanId                types.String `tfsdk:"vlan_id"`
	AdaptivePolicyGroupId types.String `tfsdk:"adaptive_policy_group_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkVLANProfiles) getPath() string {
	return fmt.Sprintf("/networks/%v/vlanProfiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkVLANProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]NetworkVLANProfilesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := NetworkVLANProfilesItems{}
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
			data.VlanGroups = make([]NetworkVLANProfilesVlanGroups, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := NetworkVLANProfilesVlanGroups{}
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
			data.VlanNames = make([]NetworkVLANProfilesVlanNames, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := NetworkVLANProfilesVlanNames{}
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
