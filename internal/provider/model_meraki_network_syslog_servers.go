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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkSyslogServers struct {
	Id        types.String                  `tfsdk:"id"`
	NetworkId types.String                  `tfsdk:"network_id"`
	Servers   []NetworkSyslogServersServers `tfsdk:"servers"`
}

type NetworkSyslogServersServers struct {
	Host  types.String `tfsdk:"host"`
	Port  types.Int64  `tfsdk:"port"`
	Roles types.Set    `tfsdk:"roles"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkSyslogServers) getPath() string {
	return fmt.Sprintf("/networks/%v/syslogServers", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkSyslogServers) toBody(ctx context.Context, state NetworkSyslogServers) string {
	body := ""
	{
		body, _ = sjson.Set(body, "servers", []interface{}{})
		for _, item := range data.Servers {
			itemBody := ""
			if !item.Host.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "host", item.Host.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			if !item.Roles.IsNull() {
				var values []string
				item.Roles.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "roles", values)
			}
			body, _ = sjson.SetRaw(body, "servers.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkSyslogServers) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("servers"); value.Exists() && value.Value() != nil {
		data.Servers = make([]NetworkSyslogServersServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkSyslogServersServers{}
			if value := res.Get("host"); value.Exists() && value.Value() != nil {
				data.Host = types.StringValue(value.String())
			} else {
				data.Host = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && value.Value() != nil {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			if value := res.Get("roles"); value.Exists() && value.Value() != nil {
				data.Roles = helpers.GetStringSet(value.Array())
			} else {
				data.Roles = types.SetNull(types.StringType)
			}
			(*parent).Servers = append((*parent).Servers, data)
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
func (data *NetworkSyslogServers) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Servers); i++ {
		keys := [...]string{"host", "port"}
		keyValues := [...]string{data.Servers[i].Host.ValueString(), strconv.FormatInt(data.Servers[i].Port.ValueInt64(), 10)}

		parent := &data
		data := (*parent).Servers[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("servers").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Servers[%d] = %+v",
				i,
				(*parent).Servers[i],
			))
			(*parent).Servers = slices.Delete((*parent).Servers, i, i+1)
			i--

			continue
		}
		if value := res.Get("host"); value.Exists() && !data.Host.IsNull() {
			data.Host = types.StringValue(value.String())
		} else {
			data.Host = types.StringNull()
		}
		if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
			data.Port = types.Int64Value(value.Int())
		} else {
			data.Port = types.Int64Null()
		}
		if value := res.Get("roles"); value.Exists() && !data.Roles.IsNull() {
			data.Roles = helpers.GetStringSet(value.Array())
		} else {
			data.Roles = types.SetNull(types.StringType)
		}
		(*parent).Servers[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkSyslogServers) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns
