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

type ApplianceOrganizationSecurityIntrusion struct {
	Id             types.String                                         `tfsdk:"id"`
	OrganizationId types.String                                         `tfsdk:"organization_id"`
	AllowedRules   []ApplianceOrganizationSecurityIntrusionAllowedRules `tfsdk:"allowed_rules"`
}

type ApplianceOrganizationSecurityIntrusionAllowedRules struct {
	Message types.String `tfsdk:"message"`
	RuleId  types.String `tfsdk:"rule_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceOrganizationSecurityIntrusion) getPath() string {
	return fmt.Sprintf("/organizations/%v/appliance/security/intrusion", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceOrganizationSecurityIntrusion) toBody(ctx context.Context, state ApplianceOrganizationSecurityIntrusion) string {
	body := ""
	{
		body, _ = sjson.Set(body, "allowedRules", []interface{}{})
		for _, item := range data.AllowedRules {
			itemBody := ""
			if !item.Message.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "message", item.Message.ValueString())
			}
			if !item.RuleId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ruleId", item.RuleId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "allowedRules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceOrganizationSecurityIntrusion) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("allowedRules"); value.Exists() && value.Value() != nil {
		data.AllowedRules = make([]ApplianceOrganizationSecurityIntrusionAllowedRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceOrganizationSecurityIntrusionAllowedRules{}
			if value := res.Get("message"); value.Exists() && value.Value() != nil {
				data.Message = types.StringValue(value.String())
			} else {
				data.Message = types.StringNull()
			}
			if value := res.Get("ruleId"); value.Exists() && value.Value() != nil {
				data.RuleId = types.StringValue(value.String())
			} else {
				data.RuleId = types.StringNull()
			}
			(*parent).AllowedRules = append((*parent).AllowedRules, data)
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
func (data *ApplianceOrganizationSecurityIntrusion) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.AllowedRules); i++ {
		keys := [...]string{"ruleId"}
		keyValues := [...]string{data.AllowedRules[i].RuleId.ValueString()}

		parent := &data
		data := (*parent).AllowedRules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("allowedRules").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AllowedRules[%d] = %+v",
				i,
				(*parent).AllowedRules[i],
			))
			(*parent).AllowedRules = slices.Delete((*parent).AllowedRules, i, i+1)
			i--

			continue
		}
		if value := res.Get("message"); value.Exists() && !data.Message.IsNull() {
			data.Message = types.StringValue(value.String())
		} else {
			data.Message = types.StringNull()
		}
		if value := res.Get("ruleId"); value.Exists() && !data.RuleId.IsNull() {
			data.RuleId = types.StringValue(value.String())
		} else {
			data.RuleId = types.StringNull()
		}
		(*parent).AllowedRules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceOrganizationSecurityIntrusion) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceOrganizationSecurityIntrusion) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
