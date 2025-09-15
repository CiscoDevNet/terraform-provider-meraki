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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceOneToManyNATRules struct {
	Id        types.String                      `tfsdk:"id"`
	NetworkId types.String                      `tfsdk:"network_id"`
	Rules     []ApplianceOneToManyNATRulesRules `tfsdk:"rules"`
}

type ApplianceOneToManyNATRulesRules struct {
	PublicIp  types.String                               `tfsdk:"public_ip"`
	Uplink    types.String                               `tfsdk:"uplink"`
	PortRules []ApplianceOneToManyNATRulesRulesPortRules `tfsdk:"port_rules"`
}

type ApplianceOneToManyNATRulesRulesPortRules struct {
	LocalIp    types.String `tfsdk:"local_ip"`
	LocalPort  types.String `tfsdk:"local_port"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	PublicPort types.String `tfsdk:"public_port"`
	AllowedIps types.List   `tfsdk:"allowed_ips"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceOneToManyNATRules) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/firewall/oneToManyNatRules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceOneToManyNATRules) toBody(ctx context.Context, state ApplianceOneToManyNATRules) string {
	body := ""
	{
		body, _ = sjson.Set(body, "rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.PublicIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "publicIp", item.PublicIp.ValueString())
			}
			if !item.Uplink.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "uplink", item.Uplink.ValueString())
			}
			{
				itemBody, _ = sjson.Set(itemBody, "portRules", []interface{}{})
				for _, childItem := range item.PortRules {
					itemChildBody := ""
					if !childItem.LocalIp.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "localIp", childItem.LocalIp.ValueString())
					}
					if !childItem.LocalPort.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "localPort", childItem.LocalPort.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Protocol.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "protocol", childItem.Protocol.ValueString())
					}
					if !childItem.PublicPort.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "publicPort", childItem.PublicPort.ValueString())
					}
					if !childItem.AllowedIps.IsNull() {
						var values []string
						childItem.AllowedIps.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "allowedIps", values)
					}
					itemBody, _ = sjson.SetRaw(itemBody, "portRules.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceOneToManyNATRules) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("rules"); value.Exists() && value.Value() != nil {
		data.Rules = make([]ApplianceOneToManyNATRulesRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceOneToManyNATRulesRules{}
			if value := res.Get("publicIp"); value.Exists() && value.Value() != nil {
				data.PublicIp = types.StringValue(value.String())
			} else {
				data.PublicIp = types.StringNull()
			}
			if value := res.Get("uplink"); value.Exists() && value.Value() != nil {
				data.Uplink = types.StringValue(value.String())
			} else {
				data.Uplink = types.StringNull()
			}
			if value := res.Get("portRules"); value.Exists() && value.Value() != nil {
				data.PortRules = make([]ApplianceOneToManyNATRulesRulesPortRules, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ApplianceOneToManyNATRulesRulesPortRules{}
					if value := res.Get("localIp"); value.Exists() && value.Value() != nil {
						data.LocalIp = types.StringValue(value.String())
					} else {
						data.LocalIp = types.StringNull()
					}
					if value := res.Get("localPort"); value.Exists() && value.Value() != nil {
						data.LocalPort = types.StringValue(value.String())
					} else {
						data.LocalPort = types.StringNull()
					}
					if value := res.Get("name"); value.Exists() && value.Value() != nil {
						data.Name = types.StringValue(value.String())
					} else {
						data.Name = types.StringNull()
					}
					if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
						data.Protocol = types.StringValue(value.String())
					} else {
						data.Protocol = types.StringNull()
					}
					if value := res.Get("publicPort"); value.Exists() && value.Value() != nil {
						data.PublicPort = types.StringValue(value.String())
					} else {
						data.PublicPort = types.StringNull()
					}
					if value := res.Get("allowedIps"); value.Exists() && value.Value() != nil {
						data.AllowedIps = helpers.GetStringList(value.Array())
					} else {
						data.AllowedIps = types.ListNull(types.StringType)
					}
					(*parent).PortRules = append((*parent).PortRules, data)
					return true
				})
			}
			(*parent).Rules = append((*parent).Rules, data)
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
func (data *ApplianceOneToManyNATRules) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Rules); i++ {
		keys := [...]string{"publicIp", "uplink"}
		keyValues := [...]string{data.Rules[i].PublicIp.ValueString(), data.Rules[i].Uplink.ValueString()}

		parent := &data
		data := (*parent).Rules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("rules").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Rules[%d] = %+v",
				i,
				(*parent).Rules[i],
			))
			(*parent).Rules = slices.Delete((*parent).Rules, i, i+1)
			i--

			continue
		}
		if value := res.Get("publicIp"); value.Exists() && !data.PublicIp.IsNull() {
			data.PublicIp = types.StringValue(value.String())
		} else {
			data.PublicIp = types.StringNull()
		}
		if value := res.Get("uplink"); value.Exists() && !data.Uplink.IsNull() {
			data.Uplink = types.StringValue(value.String())
		} else {
			data.Uplink = types.StringNull()
		}
		for i := 0; i < len(data.PortRules); i++ {
			keys := [...]string{"localIp", "localPort", "protocol", "publicPort"}
			keyValues := [...]string{data.PortRules[i].LocalIp.ValueString(), data.PortRules[i].LocalPort.ValueString(), data.PortRules[i].Protocol.ValueString(), data.PortRules[i].PublicPort.ValueString()}

			parent := &data
			data := (*parent).PortRules[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("portRules").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing PortRules[%d] = %+v",
					i,
					(*parent).PortRules[i],
				))
				(*parent).PortRules = slices.Delete((*parent).PortRules, i, i+1)
				i--

				continue
			}
			if value := res.Get("localIp"); value.Exists() && !data.LocalIp.IsNull() {
				data.LocalIp = types.StringValue(value.String())
			} else {
				data.LocalIp = types.StringNull()
			}
			if value := res.Get("localPort"); value.Exists() && !data.LocalPort.IsNull() {
				data.LocalPort = types.StringValue(value.String())
			} else {
				data.LocalPort = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("publicPort"); value.Exists() && !data.PublicPort.IsNull() {
				data.PublicPort = types.StringValue(value.String())
			} else {
				data.PublicPort = types.StringNull()
			}
			if value := res.Get("allowedIps"); value.Exists() && !data.AllowedIps.IsNull() {
				data.AllowedIps = helpers.GetStringList(value.Array())
			} else {
				data.AllowedIps = types.ListNull(types.StringType)
			}
			(*parent).PortRules[i] = data
		}
		(*parent).Rules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceOneToManyNATRules) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data ApplianceOneToManyNATRules) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
