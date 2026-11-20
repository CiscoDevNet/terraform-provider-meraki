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

type OrganizationPoliciesGlobalFirewallRule struct {
	Id                                        types.String                                                                      `tfsdk:"id"`
	OrganizationId                            types.String                                                                      `tfsdk:"organization_id"`
	Description                               types.String                                                                      `tfsdk:"description"`
	Enabled                                   types.Bool                                                                        `tfsdk:"enabled"`
	Name                                      types.String                                                                      `tfsdk:"name"`
	Policy                                    types.String                                                                      `tfsdk:"policy"`
	Priority                                  types.Int64                                                                       `tfsdk:"priority"`
	RulesetId                                 types.String                                                                      `tfsdk:"ruleset_id"`
	DestinationsCriteriaAddressRanges         types.List                                                                        `tfsdk:"destinations_criteria_address_ranges"`
	DestinationsCriteriaApplianceVlans        []OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplianceVlans        `tfsdk:"destinations_criteria_appliance_vlans"`
	DestinationsCriteriaApplicationCategories []OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplicationCategories `tfsdk:"destinations_criteria_application_categories"`
	DestinationsCriteriaApplications          []OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplications          `tfsdk:"destinations_criteria_applications"`
	DestinationsCriteriaPolicyObjectGroups    []OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaPolicyObjectGroups    `tfsdk:"destinations_criteria_policy_object_groups"`
	DestinationsCriteriaPolicyObjects         []OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaPolicyObjects         `tfsdk:"destinations_criteria_policy_objects"`
	DestinationsCriteriaPorts                 types.List                                                                        `tfsdk:"destinations_criteria_ports"`
	DestinationsCriteriaServices              []OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaServices              `tfsdk:"destinations_criteria_services"`
	DestinationsMatchCriteria                 types.Set                                                                         `tfsdk:"destinations_match_criteria"`
	SourcesCriteriaAddressRanges              types.List                                                                        `tfsdk:"sources_criteria_address_ranges"`
	SourcesCriteriaApplianceVlans             []OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaApplianceVlans             `tfsdk:"sources_criteria_appliance_vlans"`
	SourcesCriteriaPolicyObjectGroups         []OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaPolicyObjectGroups         `tfsdk:"sources_criteria_policy_object_groups"`
	SourcesCriteriaPolicyObjects              []OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaPolicyObjects              `tfsdk:"sources_criteria_policy_objects"`
	SourcesCriteriaPorts                      types.List                                                                        `tfsdk:"sources_criteria_ports"`
	SourcesMatchCriteria                      types.Set                                                                         `tfsdk:"sources_match_criteria"`
}

type OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplianceVlans struct {
	InterfaceId types.String `tfsdk:"interface_id"`
}

type OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplicationCategories struct {
	Id types.String `tfsdk:"id"`
}

type OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplications struct {
	Id types.String `tfsdk:"id"`
}

type OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaPolicyObjectGroups struct {
	Id types.String `tfsdk:"id"`
}

type OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaPolicyObjects struct {
	Id types.String `tfsdk:"id"`
}

type OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaServices struct {
	Protocol types.String `tfsdk:"protocol"`
	Ports    types.List   `tfsdk:"ports"`
}

type OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaApplianceVlans struct {
	InterfaceId types.String `tfsdk:"interface_id"`
}

type OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaPolicyObjectGroups struct {
	Id types.String `tfsdk:"id"`
}

type OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaPolicyObjects struct {
	Id types.String `tfsdk:"id"`
}

type OrganizationPoliciesGlobalFirewallRuleIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationPoliciesGlobalFirewallRule) getPath() string {
	return fmt.Sprintf("/organizations/%v/policies/global/firewall/rulesets/rules", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationPoliciesGlobalFirewallRule) toBody(ctx context.Context, state OrganizationPoliciesGlobalFirewallRule) string {
	body := ""
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Policy.IsNull() {
		body, _ = sjson.Set(body, "policy", data.Policy.ValueString())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, "priority", data.Priority.ValueInt64())
	}
	if !data.RulesetId.IsNull() {
		body, _ = sjson.Set(body, "rulesetId", data.RulesetId.ValueString())
	}
	if !data.DestinationsCriteriaAddressRanges.IsNull() {
		var values []string
		data.DestinationsCriteriaAddressRanges.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "destinations.criteria.addressRanges", values)
	}
	if len(data.DestinationsCriteriaApplianceVlans) > 0 {
		body, _ = sjson.Set(body, "destinations.criteria.applianceVlans", []interface{}{})
		for _, item := range data.DestinationsCriteriaApplianceVlans {
			itemBody := ""
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceId", item.InterfaceId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "destinations.criteria.applianceVlans.-1", itemBody)
		}
	}
	if len(data.DestinationsCriteriaApplicationCategories) > 0 {
		body, _ = sjson.Set(body, "destinations.criteria.applicationCategories", []interface{}{})
		for _, item := range data.DestinationsCriteriaApplicationCategories {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "destinations.criteria.applicationCategories.-1", itemBody)
		}
	}
	if len(data.DestinationsCriteriaApplications) > 0 {
		body, _ = sjson.Set(body, "destinations.criteria.applications", []interface{}{})
		for _, item := range data.DestinationsCriteriaApplications {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "destinations.criteria.applications.-1", itemBody)
		}
	}
	if len(data.DestinationsCriteriaPolicyObjectGroups) > 0 {
		body, _ = sjson.Set(body, "destinations.criteria.policyObjectGroups", []interface{}{})
		for _, item := range data.DestinationsCriteriaPolicyObjectGroups {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "destinations.criteria.policyObjectGroups.-1", itemBody)
		}
	}
	if len(data.DestinationsCriteriaPolicyObjects) > 0 {
		body, _ = sjson.Set(body, "destinations.criteria.policyObjects", []interface{}{})
		for _, item := range data.DestinationsCriteriaPolicyObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "destinations.criteria.policyObjects.-1", itemBody)
		}
	}
	if !data.DestinationsCriteriaPorts.IsNull() {
		var values []string
		data.DestinationsCriteriaPorts.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "destinations.criteria.ports", values)
	}
	if len(data.DestinationsCriteriaServices) > 0 {
		body, _ = sjson.Set(body, "destinations.criteria.services", []interface{}{})
		for _, item := range data.DestinationsCriteriaServices {
			itemBody := ""
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			if !item.Ports.IsNull() {
				var values []string
				item.Ports.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ports", values)
			}
			body, _ = sjson.SetRaw(body, "destinations.criteria.services.-1", itemBody)
		}
	}
	if !data.DestinationsMatchCriteria.IsNull() {
		var values []string
		data.DestinationsMatchCriteria.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "destinations.matchCriteria", values)
	}
	if !data.SourcesCriteriaAddressRanges.IsNull() {
		var values []string
		data.SourcesCriteriaAddressRanges.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "sources.criteria.addressRanges", values)
	}
	if len(data.SourcesCriteriaApplianceVlans) > 0 {
		body, _ = sjson.Set(body, "sources.criteria.applianceVlans", []interface{}{})
		for _, item := range data.SourcesCriteriaApplianceVlans {
			itemBody := ""
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceId", item.InterfaceId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "sources.criteria.applianceVlans.-1", itemBody)
		}
	}
	if len(data.SourcesCriteriaPolicyObjectGroups) > 0 {
		body, _ = sjson.Set(body, "sources.criteria.policyObjectGroups", []interface{}{})
		for _, item := range data.SourcesCriteriaPolicyObjectGroups {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "sources.criteria.policyObjectGroups.-1", itemBody)
		}
	}
	if len(data.SourcesCriteriaPolicyObjects) > 0 {
		body, _ = sjson.Set(body, "sources.criteria.policyObjects", []interface{}{})
		for _, item := range data.SourcesCriteriaPolicyObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "sources.criteria.policyObjects.-1", itemBody)
		}
	}
	if !data.SourcesCriteriaPorts.IsNull() {
		var values []string
		data.SourcesCriteriaPorts.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "sources.criteria.ports", values)
	}
	if !data.SourcesMatchCriteria.IsNull() {
		var values []string
		data.SourcesMatchCriteria.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "sources.matchCriteria", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationPoliciesGlobalFirewallRule) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && value.Value() != nil {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("policy"); value.Exists() && value.Value() != nil {
		data.Policy = types.StringValue(value.String())
	} else {
		data.Policy = types.StringNull()
	}
	if value := res.Get("priority"); value.Exists() && value.Value() != nil {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("rulesetId"); value.Exists() && value.Value() != nil {
		data.RulesetId = types.StringValue(value.String())
	} else {
		data.RulesetId = types.StringNull()
	}
	if value := res.Get("destinations.criteria.addressRanges"); value.Exists() && value.Value() != nil {
		data.DestinationsCriteriaAddressRanges = helpers.GetStringList(value.Array())
	} else {
		data.DestinationsCriteriaAddressRanges = types.ListNull(types.StringType)
	}
	if value := res.Get("destinations.criteria.applianceVlans"); value.Exists() && value.Value() != nil {
		data.DestinationsCriteriaApplianceVlans = make([]OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplianceVlans, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplianceVlans{}
			if value := res.Get("interfaceId"); value.Exists() && value.Value() != nil {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			(*parent).DestinationsCriteriaApplianceVlans = append((*parent).DestinationsCriteriaApplianceVlans, data)
			return true
		})
	}
	if value := res.Get("destinations.criteria.applicationCategories"); value.Exists() && value.Value() != nil {
		data.DestinationsCriteriaApplicationCategories = make([]OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplicationCategories, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplicationCategories{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).DestinationsCriteriaApplicationCategories = append((*parent).DestinationsCriteriaApplicationCategories, data)
			return true
		})
	}
	if value := res.Get("destinations.criteria.applications"); value.Exists() && value.Value() != nil {
		data.DestinationsCriteriaApplications = make([]OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplications, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaApplications{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).DestinationsCriteriaApplications = append((*parent).DestinationsCriteriaApplications, data)
			return true
		})
	}
	if value := res.Get("destinations.criteria.policyObjectGroups"); value.Exists() && value.Value() != nil {
		data.DestinationsCriteriaPolicyObjectGroups = make([]OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaPolicyObjectGroups, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaPolicyObjectGroups{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).DestinationsCriteriaPolicyObjectGroups = append((*parent).DestinationsCriteriaPolicyObjectGroups, data)
			return true
		})
	}
	if value := res.Get("destinations.criteria.policyObjects"); value.Exists() && value.Value() != nil {
		data.DestinationsCriteriaPolicyObjects = make([]OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaPolicyObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaPolicyObjects{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).DestinationsCriteriaPolicyObjects = append((*parent).DestinationsCriteriaPolicyObjects, data)
			return true
		})
	}
	if value := res.Get("destinations.criteria.ports"); value.Exists() && value.Value() != nil {
		data.DestinationsCriteriaPorts = helpers.GetStringList(value.Array())
	} else {
		data.DestinationsCriteriaPorts = types.ListNull(types.StringType)
	}
	if value := res.Get("destinations.criteria.services"); value.Exists() && value.Value() != nil {
		data.DestinationsCriteriaServices = make([]OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaServices, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationPoliciesGlobalFirewallRuleDestinationsCriteriaServices{}
			if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("ports"); value.Exists() && value.Value() != nil {
				data.Ports = helpers.GetStringList(value.Array())
			} else {
				data.Ports = types.ListNull(types.StringType)
			}
			(*parent).DestinationsCriteriaServices = append((*parent).DestinationsCriteriaServices, data)
			return true
		})
	}
	if value := res.Get("destinations.matchCriteria"); value.Exists() && value.Value() != nil {
		data.DestinationsMatchCriteria = helpers.GetStringSet(value.Array())
	} else {
		data.DestinationsMatchCriteria = types.SetNull(types.StringType)
	}
	if value := res.Get("sources.criteria.addressRanges"); value.Exists() && value.Value() != nil {
		data.SourcesCriteriaAddressRanges = helpers.GetStringList(value.Array())
	} else {
		data.SourcesCriteriaAddressRanges = types.ListNull(types.StringType)
	}
	if value := res.Get("sources.criteria.applianceVlans"); value.Exists() && value.Value() != nil {
		data.SourcesCriteriaApplianceVlans = make([]OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaApplianceVlans, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaApplianceVlans{}
			if value := res.Get("interfaceId"); value.Exists() && value.Value() != nil {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			(*parent).SourcesCriteriaApplianceVlans = append((*parent).SourcesCriteriaApplianceVlans, data)
			return true
		})
	}
	if value := res.Get("sources.criteria.policyObjectGroups"); value.Exists() && value.Value() != nil {
		data.SourcesCriteriaPolicyObjectGroups = make([]OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaPolicyObjectGroups, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaPolicyObjectGroups{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SourcesCriteriaPolicyObjectGroups = append((*parent).SourcesCriteriaPolicyObjectGroups, data)
			return true
		})
	}
	if value := res.Get("sources.criteria.policyObjects"); value.Exists() && value.Value() != nil {
		data.SourcesCriteriaPolicyObjects = make([]OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaPolicyObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := OrganizationPoliciesGlobalFirewallRuleSourcesCriteriaPolicyObjects{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SourcesCriteriaPolicyObjects = append((*parent).SourcesCriteriaPolicyObjects, data)
			return true
		})
	}
	if value := res.Get("sources.criteria.ports"); value.Exists() && value.Value() != nil {
		data.SourcesCriteriaPorts = helpers.GetStringList(value.Array())
	} else {
		data.SourcesCriteriaPorts = types.ListNull(types.StringType)
	}
	if value := res.Get("sources.matchCriteria"); value.Exists() && value.Value() != nil {
		data.SourcesMatchCriteria = helpers.GetStringSet(value.Array())
	} else {
		data.SourcesMatchCriteria = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationPoliciesGlobalFirewallRule) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("policy"); value.Exists() && !data.Policy.IsNull() {
		data.Policy = types.StringValue(value.String())
	} else {
		data.Policy = types.StringNull()
	}
	if value := res.Get("priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("rulesetId"); value.Exists() && !data.RulesetId.IsNull() {
		data.RulesetId = types.StringValue(value.String())
	} else {
		data.RulesetId = types.StringNull()
	}
	if value := res.Get("destinations.criteria.addressRanges"); value.Exists() && !data.DestinationsCriteriaAddressRanges.IsNull() {
		data.DestinationsCriteriaAddressRanges = helpers.GetStringList(value.Array())
	} else {
		data.DestinationsCriteriaAddressRanges = types.ListNull(types.StringType)
	}
	for i := 0; i < len(data.DestinationsCriteriaApplianceVlans); i++ {
		keys := [...]string{"interfaceId"}
		keyValues := [...]string{data.DestinationsCriteriaApplianceVlans[i].InterfaceId.ValueString()}

		parent := &data
		data := (*parent).DestinationsCriteriaApplianceVlans[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinations.criteria.applianceVlans").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationsCriteriaApplianceVlans[%d] = %+v",
				i,
				(*parent).DestinationsCriteriaApplianceVlans[i],
			))
			(*parent).DestinationsCriteriaApplianceVlans = slices.Delete((*parent).DestinationsCriteriaApplianceVlans, i, i+1)
			i--

			continue
		}
		if value := res.Get("interfaceId"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		(*parent).DestinationsCriteriaApplianceVlans[i] = data
	}
	for i := 0; i < len(data.DestinationsCriteriaApplicationCategories); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.DestinationsCriteriaApplicationCategories[i].Id.ValueString()}

		parent := &data
		data := (*parent).DestinationsCriteriaApplicationCategories[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinations.criteria.applicationCategories").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationsCriteriaApplicationCategories[%d] = %+v",
				i,
				(*parent).DestinationsCriteriaApplicationCategories[i],
			))
			(*parent).DestinationsCriteriaApplicationCategories = slices.Delete((*parent).DestinationsCriteriaApplicationCategories, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).DestinationsCriteriaApplicationCategories[i] = data
	}
	for i := 0; i < len(data.DestinationsCriteriaApplications); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.DestinationsCriteriaApplications[i].Id.ValueString()}

		parent := &data
		data := (*parent).DestinationsCriteriaApplications[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinations.criteria.applications").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationsCriteriaApplications[%d] = %+v",
				i,
				(*parent).DestinationsCriteriaApplications[i],
			))
			(*parent).DestinationsCriteriaApplications = slices.Delete((*parent).DestinationsCriteriaApplications, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).DestinationsCriteriaApplications[i] = data
	}
	for i := 0; i < len(data.DestinationsCriteriaPolicyObjectGroups); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.DestinationsCriteriaPolicyObjectGroups[i].Id.ValueString()}

		parent := &data
		data := (*parent).DestinationsCriteriaPolicyObjectGroups[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinations.criteria.policyObjectGroups").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationsCriteriaPolicyObjectGroups[%d] = %+v",
				i,
				(*parent).DestinationsCriteriaPolicyObjectGroups[i],
			))
			(*parent).DestinationsCriteriaPolicyObjectGroups = slices.Delete((*parent).DestinationsCriteriaPolicyObjectGroups, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).DestinationsCriteriaPolicyObjectGroups[i] = data
	}
	for i := 0; i < len(data.DestinationsCriteriaPolicyObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.DestinationsCriteriaPolicyObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).DestinationsCriteriaPolicyObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinations.criteria.policyObjects").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationsCriteriaPolicyObjects[%d] = %+v",
				i,
				(*parent).DestinationsCriteriaPolicyObjects[i],
			))
			(*parent).DestinationsCriteriaPolicyObjects = slices.Delete((*parent).DestinationsCriteriaPolicyObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).DestinationsCriteriaPolicyObjects[i] = data
	}
	if value := res.Get("destinations.criteria.ports"); value.Exists() && !data.DestinationsCriteriaPorts.IsNull() {
		data.DestinationsCriteriaPorts = helpers.GetStringList(value.Array())
	} else {
		data.DestinationsCriteriaPorts = types.ListNull(types.StringType)
	}
	for i := 0; i < len(data.DestinationsCriteriaServices); i++ {
		keys := [...]string{"protocol"}
		keyValues := [...]string{data.DestinationsCriteriaServices[i].Protocol.ValueString()}

		parent := &data
		data := (*parent).DestinationsCriteriaServices[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinations.criteria.services").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DestinationsCriteriaServices[%d] = %+v",
				i,
				(*parent).DestinationsCriteriaServices[i],
			))
			(*parent).DestinationsCriteriaServices = slices.Delete((*parent).DestinationsCriteriaServices, i, i+1)
			i--

			continue
		}
		if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		if value := res.Get("ports"); value.Exists() && !data.Ports.IsNull() {
			data.Ports = helpers.GetStringList(value.Array())
		} else {
			data.Ports = types.ListNull(types.StringType)
		}
		(*parent).DestinationsCriteriaServices[i] = data
	}
	if value := res.Get("destinations.matchCriteria"); value.Exists() && !data.DestinationsMatchCriteria.IsNull() {
		data.DestinationsMatchCriteria = helpers.GetStringSet(value.Array())
	} else {
		data.DestinationsMatchCriteria = types.SetNull(types.StringType)
	}
	if value := res.Get("sources.criteria.addressRanges"); value.Exists() && !data.SourcesCriteriaAddressRanges.IsNull() {
		data.SourcesCriteriaAddressRanges = helpers.GetStringList(value.Array())
	} else {
		data.SourcesCriteriaAddressRanges = types.ListNull(types.StringType)
	}
	for i := 0; i < len(data.SourcesCriteriaApplianceVlans); i++ {
		keys := [...]string{"interfaceId"}
		keyValues := [...]string{data.SourcesCriteriaApplianceVlans[i].InterfaceId.ValueString()}

		parent := &data
		data := (*parent).SourcesCriteriaApplianceVlans[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("sources.criteria.applianceVlans").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing SourcesCriteriaApplianceVlans[%d] = %+v",
				i,
				(*parent).SourcesCriteriaApplianceVlans[i],
			))
			(*parent).SourcesCriteriaApplianceVlans = slices.Delete((*parent).SourcesCriteriaApplianceVlans, i, i+1)
			i--

			continue
		}
		if value := res.Get("interfaceId"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		(*parent).SourcesCriteriaApplianceVlans[i] = data
	}
	for i := 0; i < len(data.SourcesCriteriaPolicyObjectGroups); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.SourcesCriteriaPolicyObjectGroups[i].Id.ValueString()}

		parent := &data
		data := (*parent).SourcesCriteriaPolicyObjectGroups[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("sources.criteria.policyObjectGroups").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing SourcesCriteriaPolicyObjectGroups[%d] = %+v",
				i,
				(*parent).SourcesCriteriaPolicyObjectGroups[i],
			))
			(*parent).SourcesCriteriaPolicyObjectGroups = slices.Delete((*parent).SourcesCriteriaPolicyObjectGroups, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).SourcesCriteriaPolicyObjectGroups[i] = data
	}
	for i := 0; i < len(data.SourcesCriteriaPolicyObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.SourcesCriteriaPolicyObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).SourcesCriteriaPolicyObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("sources.criteria.policyObjects").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing SourcesCriteriaPolicyObjects[%d] = %+v",
				i,
				(*parent).SourcesCriteriaPolicyObjects[i],
			))
			(*parent).SourcesCriteriaPolicyObjects = slices.Delete((*parent).SourcesCriteriaPolicyObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).SourcesCriteriaPolicyObjects[i] = data
	}
	if value := res.Get("sources.criteria.ports"); value.Exists() && !data.SourcesCriteriaPorts.IsNull() {
		data.SourcesCriteriaPorts = helpers.GetStringList(value.Array())
	} else {
		data.SourcesCriteriaPorts = types.ListNull(types.StringType)
	}
	if value := res.Get("sources.matchCriteria"); value.Exists() && !data.SourcesMatchCriteria.IsNull() {
		data.SourcesMatchCriteria = helpers.GetStringSet(value.Array())
	} else {
		data.SourcesMatchCriteria = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationPoliciesGlobalFirewallRule) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *OrganizationPoliciesGlobalFirewallRuleIdentity) toIdentity(ctx context.Context, plan *OrganizationPoliciesGlobalFirewallRule) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *OrganizationPoliciesGlobalFirewallRule) fromIdentity(ctx context.Context, identity *OrganizationPoliciesGlobalFirewallRuleIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationPoliciesGlobalFirewallRule) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
