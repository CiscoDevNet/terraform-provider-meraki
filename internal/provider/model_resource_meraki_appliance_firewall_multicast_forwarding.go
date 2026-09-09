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

type ApplianceFirewallMulticastForwarding struct {
	Id        types.String                                `tfsdk:"id"`
	NetworkId types.String                                `tfsdk:"network_id"`
	Rules     []ApplianceFirewallMulticastForwardingRules `tfsdk:"rules"`
}

type ApplianceFirewallMulticastForwardingRules struct {
	Address     types.String `tfsdk:"address"`
	Description types.String `tfsdk:"description"`
	VlanIds     types.List   `tfsdk:"vlan_ids"`
}

type ApplianceFirewallMulticastForwardingIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceFirewallMulticastForwarding) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/firewall/multicastForwarding", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceFirewallMulticastForwarding) toBody(ctx context.Context, state ApplianceFirewallMulticastForwarding) string {
	body := ""
	{
		body, _ = sjson.Set(body, "rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "address", item.Address.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.VlanIds.IsNull() {
				var values []string
				item.VlanIds.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "vlanIds", values)
			}
			body, _ = sjson.SetRaw(body, "rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPreservingNulls

// toBodyPreservingNulls walks the same writable-attribute schema as toBody but
// reads directly from the raw API response (gjson) instead of from the
// Terraform model. Unlike toBody, it preserves attributes that the API
// explicitly returned as `null` (emitting them as JSON `null` rather than
// dropping them). This is used by the singleton restoreOriginalStateOnDestroy
// path so that explicit-null fields captured during Create are restored on
// Delete. Keep this method in sync with toBody — both walk the same
// `.Attributes` schema and must agree on which fields are writable.
func (data ApplianceFirewallMulticastForwarding) toBodyPreservingNulls(ctx context.Context, res meraki.Res) string {
	body := ""
	if value := res.Get("rules"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "rules", "null")
		} else {
			body, _ = sjson.Set(body, "rules", []interface{}{})
			parent := &body
			value.ForEach(func(k, res gjson.Result) bool {
				body := ""
				if value := res.Get("address"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "address", "null")
					} else {
						body, _ = sjson.Set(body, "address", value.String())
					}
				}
				if value := res.Get("description"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "description", "null")
					} else {
						body, _ = sjson.Set(body, "description", value.String())
					}
				}
				if value := res.Get("vlanIds"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "vlanIds", "null")
					} else {
						var values []string
						for _, v := range value.Array() {
							values = append(values, v.String())
						}
						body, _ = sjson.Set(body, "vlanIds", values)
					}
				}
				*parent, _ = sjson.SetRaw(*parent, "rules.-1", body)
				return true
			})
		}
	}
	return body
}

// End of section. //template:end toBodyPreservingNulls

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceFirewallMulticastForwarding) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("rules"); value.Exists() && value.Value() != nil {
		data.Rules = make([]ApplianceFirewallMulticastForwardingRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceFirewallMulticastForwardingRules{}
			if value := res.Get("address"); value.Exists() && value.Value() != nil {
				data.Address = types.StringValue(value.String())
			} else {
				data.Address = types.StringNull()
			}
			if value := res.Get("description"); value.Exists() && value.Value() != nil {
				data.Description = types.StringValue(value.String())
			} else {
				data.Description = types.StringNull()
			}
			if value := res.Get("vlanIds"); value.Exists() && value.Value() != nil {
				data.VlanIds = helpers.GetStringList(value.Array())
			} else {
				data.VlanIds = types.ListNull(types.StringType)
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
func (data *ApplianceFirewallMulticastForwarding) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Rules); i++ {
		keys := [...]string{"address", "description"}
		keyValues := [...]string{data.Rules[i].Address.ValueString(), data.Rules[i].Description.ValueString()}

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
		if value := res.Get("address"); value.Exists() && !data.Address.IsNull() {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
		}
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("vlanIds"); value.Exists() && !data.VlanIds.IsNull() {
			data.VlanIds = helpers.GetStringList(value.Array())
		} else {
			data.VlanIds = types.ListNull(types.StringType)
		}
		(*parent).Rules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceFirewallMulticastForwarding) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceFirewallMulticastForwardingIdentity) toIdentity(ctx context.Context, plan *ApplianceFirewallMulticastForwarding) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceFirewallMulticastForwarding) fromIdentity(ctx context.Context, identity *ApplianceFirewallMulticastForwardingIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceFirewallMulticastForwarding) toDestroyBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "rules", []interface{}{})
	return body
}

// End of section. //template:end toDestroyBody
