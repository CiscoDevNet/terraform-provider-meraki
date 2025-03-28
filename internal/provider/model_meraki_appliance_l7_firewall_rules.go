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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceL7FirewallRules struct {
	Id        types.String                    `tfsdk:"id"`
	NetworkId types.String                    `tfsdk:"network_id"`
	Rules     []ApplianceL7FirewallRulesRules `tfsdk:"rules"`
}

type ApplianceL7FirewallRulesRules struct {
	Policy         types.String `tfsdk:"policy"`
	Type           types.String `tfsdk:"type"`
	Value          types.String `tfsdk:"value"`
	ValueCountries types.List   `tfsdk:"value_countries"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceL7FirewallRules) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/firewall/l7FirewallRules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

func (data ApplianceL7FirewallRules) toBody(ctx context.Context, state ApplianceL7FirewallRules) string {
	body := ""
	{
		body, _ = sjson.Set(body, "rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.Policy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "policy", item.Policy.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			var valuePath string
			if item.Type.ValueString() == "application" || item.Type.ValueString() == "applicationCategory" {
				valuePath = "value.id"
			} else {
				valuePath = "value"
			}
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, valuePath, item.Value.ValueString())
			}
			if !item.ValueCountries.IsNull() && (item.Type.ValueString() == "blockedCountries" || item.Type.ValueString() == "allowedCountries") {
				var values []string
				item.ValueCountries.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "value", values)
			}
			body, _ = sjson.SetRaw(body, "rules.-1", itemBody)
		}
	}
	return body
}

func (data *ApplianceL7FirewallRules) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("rules"); value.Exists() && value.Value() != nil {
		data.Rules = make([]ApplianceL7FirewallRulesRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceL7FirewallRulesRules{}
			if value := res.Get("policy"); value.Exists() && value.Value() != nil {
				data.Policy = types.StringValue(value.String())
			} else {
				data.Policy = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && value.Value() != nil {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			var valuePath string
			if data.Type.ValueString() == "application" || data.Type.ValueString() == "applicationCategory" {
				valuePath = "value.id"
			} else {
				valuePath = "value"
			}
			if !(data.Type.ValueString() == "blockedCountries") && !(data.Type.ValueString() == "allowedCountries") {
				if value := res.Get(valuePath); value.Exists() && value.Value() != nil {
					data.Value = types.StringValue(value.String())
				} else {
					data.Value = types.StringNull()
				}
			} else {
				if value := res.Get("value"); value.Exists() && value.Value() != nil {
					data.ValueCountries = helpers.GetStringList(value.Array())
				} else {
					data.ValueCountries = types.ListNull(types.StringType)
				}
			}
			(*parent).Rules = append((*parent).Rules, data)
			return true
		})
	}
}

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceL7FirewallRules) fromBodyPartial(ctx context.Context, res meraki.Res) {
	{
		l := len(res.Get("rules").Array())
		tflog.Debug(ctx, fmt.Sprintf("rules array resizing from %d to %d", len(data.Rules), l))
		if len(data.Rules) > l {
			data.Rules = data.Rules[:l]
		}
	}
	for i := range data.Rules {
		parent := &data
		data := (*parent).Rules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("rules.%d", i))
		if value := res.Get("policy"); value.Exists() && !data.Policy.IsNull() {
			data.Policy = types.StringValue(value.String())
		} else {
			data.Policy = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		var valuePath string
		if data.Type.ValueString() == "application" || data.Type.ValueString() == "applicationCategory" {
			valuePath = "value.id"
		} else {
			valuePath = "value"
		}
		if !(data.Type.ValueString() == "blockedCountries") && !(data.Type.ValueString() == "allowedCountries") {
			if value := res.Get(valuePath); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
		} else {
			if value := res.Get("value"); value.Exists() && !data.ValueCountries.IsNull() {
				data.ValueCountries = helpers.GetStringList(value.Array())
			} else {
				data.ValueCountries = types.ListNull(types.StringType)
			}
		}
		(*parent).Rules[i] = data
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceL7FirewallRules) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceL7FirewallRules) toDestroyBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "rules", []interface{}{})
	return body
}

// End of section. //template:end toDestroyBody
