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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkSNMP struct {
	Id              types.String       `tfsdk:"id"`
	NetworkId       types.String       `tfsdk:"network_id"`
	Access          types.String       `tfsdk:"access"`
	CommunityString types.String       `tfsdk:"community_string"`
	Users           []NetworkSNMPUsers `tfsdk:"users"`
}

type NetworkSNMPUsers struct {
	Passphrase types.String `tfsdk:"passphrase"`
	Username   types.String `tfsdk:"username"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkSNMP) getPath() string {
	return fmt.Sprintf("/networks/%v/snmp", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkSNMP) toBody(ctx context.Context, state NetworkSNMP) string {
	body := ""
	if !data.Access.IsNull() {
		body, _ = sjson.Set(body, "access", data.Access.ValueString())
	}
	if !data.CommunityString.IsNull() {
		body, _ = sjson.Set(body, "communityString", data.CommunityString.ValueString())
	}
	if data.Users != nil {
		body, _ = sjson.Set(body, "users", []interface{}{})
		for _, item := range data.Users {
			itemBody := ""
			if !item.Passphrase.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "passphrase", item.Passphrase.ValueString())
			}
			if !item.Username.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "username", item.Username.ValueString())
			}
			body, _ = sjson.SetRaw(body, "users.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkSNMP) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("access"); value.Exists() && value.Value() != nil {
		data.Access = types.StringValue(value.String())
	} else {
		data.Access = types.StringNull()
	}
	if value := res.Get("communityString"); value.Exists() && value.Value() != nil {
		data.CommunityString = types.StringValue(value.String())
	} else {
		data.CommunityString = types.StringNull()
	}
	if value := res.Get("users"); value.Exists() && value.Value() != nil {
		data.Users = make([]NetworkSNMPUsers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkSNMPUsers{}
			if value := res.Get("passphrase"); value.Exists() && value.Value() != nil {
				data.Passphrase = types.StringValue(value.String())
			} else {
				data.Passphrase = types.StringNull()
			}
			if value := res.Get("username"); value.Exists() && value.Value() != nil {
				data.Username = types.StringValue(value.String())
			} else {
				data.Username = types.StringNull()
			}
			(*parent).Users = append((*parent).Users, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkSNMP) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("access"); value.Exists() && !data.Access.IsNull() {
		data.Access = types.StringValue(value.String())
	} else {
		data.Access = types.StringNull()
	}
	if value := res.Get("communityString"); value.Exists() && !data.CommunityString.IsNull() {
		data.CommunityString = types.StringValue(value.String())
	} else {
		data.CommunityString = types.StringNull()
	}
	for i := 0; i < len(data.Users); i++ {
		keys := [...]string{"username"}
		keyValues := [...]string{data.Users[i].Username.ValueString()}

		parent := &data
		data := (*parent).Users[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("users").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Users[%d] = %+v",
				i,
				(*parent).Users[i],
			))
			(*parent).Users = slices.Delete((*parent).Users, i, i+1)
			i--

			continue
		}
		if value := res.Get("passphrase"); value.Exists() && !data.Passphrase.IsNull() {
			data.Passphrase = types.StringValue(value.String())
		} else {
			data.Passphrase = types.StringNull()
		}
		if value := res.Get("username"); value.Exists() && !data.Username.IsNull() {
			data.Username = types.StringValue(value.String())
		} else {
			data.Username = types.StringNull()
		}
		(*parent).Users[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkSNMP) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkSNMP) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
