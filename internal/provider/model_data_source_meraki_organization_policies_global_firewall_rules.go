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
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceOrganizationPoliciesGlobalFirewallRules struct {
	OrganizationId types.String                                             `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationPoliciesGlobalFirewallRulesItems `tfsdk:"items"`
}

type DataSourceOrganizationPoliciesGlobalFirewallRulesItems struct {
	Id                                        types.String                                                                                 `tfsdk:"id"`
	Description                               types.String                                                                                 `tfsdk:"description"`
	Enabled                                   types.Bool                                                                                   `tfsdk:"enabled"`
	Name                                      types.String                                                                                 `tfsdk:"name"`
	Policy                                    types.String                                                                                 `tfsdk:"policy"`
	Priority                                  types.Int64                                                                                  `tfsdk:"priority"`
	RulesetId                                 types.String                                                                                 `tfsdk:"ruleset_id"`
	DestinationsCriteriaAddressRanges         types.List                                                                                   `tfsdk:"destinations_criteria_address_ranges"`
	DestinationsCriteriaApplianceVlans        []DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplianceVlans        `tfsdk:"destinations_criteria_appliance_vlans"`
	DestinationsCriteriaApplicationCategories []DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplicationCategories `tfsdk:"destinations_criteria_application_categories"`
	DestinationsCriteriaApplications          []DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplications          `tfsdk:"destinations_criteria_applications"`
	DestinationsCriteriaPolicyObjectGroups    []DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaPolicyObjectGroups    `tfsdk:"destinations_criteria_policy_object_groups"`
	DestinationsCriteriaPolicyObjects         []DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaPolicyObjects         `tfsdk:"destinations_criteria_policy_objects"`
	DestinationsCriteriaPorts                 types.List                                                                                   `tfsdk:"destinations_criteria_ports"`
	DestinationsCriteriaServices              []DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaServices              `tfsdk:"destinations_criteria_services"`
	DestinationsMatchCriteria                 types.Set                                                                                    `tfsdk:"destinations_match_criteria"`
	SourcesCriteriaAddressRanges              types.List                                                                                   `tfsdk:"sources_criteria_address_ranges"`
	SourcesCriteriaApplianceVlans             []DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaApplianceVlans             `tfsdk:"sources_criteria_appliance_vlans"`
	SourcesCriteriaPolicyObjectGroups         []DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaPolicyObjectGroups         `tfsdk:"sources_criteria_policy_object_groups"`
	SourcesCriteriaPolicyObjects              []DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaPolicyObjects              `tfsdk:"sources_criteria_policy_objects"`
	SourcesCriteriaPorts                      types.List                                                                                   `tfsdk:"sources_criteria_ports"`
	SourcesMatchCriteria                      types.Set                                                                                    `tfsdk:"sources_match_criteria"`
}

type DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplianceVlans struct {
	InterfaceId types.String `tfsdk:"interface_id"`
}

type DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplicationCategories struct {
	Id           types.String                                                                                             `tfsdk:"id"`
	Name         types.String                                                                                             `tfsdk:"name"`
	Applications []DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplicationCategoriesApplications `tfsdk:"applications"`
}

type DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplications struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaPolicyObjectGroups struct {
	Id types.String `tfsdk:"id"`
}

type DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaPolicyObjects struct {
	Id types.String `tfsdk:"id"`
}

type DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaServices struct {
	Protocol types.String `tfsdk:"protocol"`
	Ports    types.List   `tfsdk:"ports"`
}

type DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaApplianceVlans struct {
	InterfaceId types.String `tfsdk:"interface_id"`
}

type DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaPolicyObjectGroups struct {
	Id types.String `tfsdk:"id"`
}

type DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaPolicyObjects struct {
	Id types.String `tfsdk:"id"`
}

type DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplicationCategoriesApplications struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationPoliciesGlobalFirewallRules) getPath() string {
	return fmt.Sprintf("/organizations/%v/policies/global/firewall/rulesets/rules", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationPoliciesGlobalFirewallRules) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationPoliciesGlobalFirewallRulesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationPoliciesGlobalFirewallRulesItems{}
		data.Id = types.StringValue(res.Get("ruleId").String())
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
			data.DestinationsCriteriaApplianceVlans = make([]DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplianceVlans, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplianceVlans{}
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
			data.DestinationsCriteriaApplicationCategories = make([]DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplicationCategories, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplicationCategories{}
				if value := res.Get("id"); value.Exists() && value.Value() != nil {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("applications"); value.Exists() && value.Value() != nil {
					data.Applications = make([]DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplicationCategoriesApplications, 0)
					value.ForEach(func(k, res gjson.Result) bool {
						parent := &data
						data := DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplicationCategoriesApplications{}
						if value := res.Get("id"); value.Exists() && value.Value() != nil {
							data.Id = types.StringValue(value.String())
						} else {
							data.Id = types.StringNull()
						}
						if value := res.Get("name"); value.Exists() && value.Value() != nil {
							data.Name = types.StringValue(value.String())
						} else {
							data.Name = types.StringNull()
						}
						(*parent).Applications = append((*parent).Applications, data)
						return true
					})
				}
				(*parent).DestinationsCriteriaApplicationCategories = append((*parent).DestinationsCriteriaApplicationCategories, data)
				return true
			})
		}
		if value := res.Get("destinations.criteria.applications"); value.Exists() && value.Value() != nil {
			data.DestinationsCriteriaApplications = make([]DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplications, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaApplications{}
				if value := res.Get("id"); value.Exists() && value.Value() != nil {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				(*parent).DestinationsCriteriaApplications = append((*parent).DestinationsCriteriaApplications, data)
				return true
			})
		}
		if value := res.Get("destinations.criteria.policyObjectGroups"); value.Exists() && value.Value() != nil {
			data.DestinationsCriteriaPolicyObjectGroups = make([]DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaPolicyObjectGroups, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaPolicyObjectGroups{}
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
			data.DestinationsCriteriaPolicyObjects = make([]DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaPolicyObjects, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaPolicyObjects{}
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
			data.DestinationsCriteriaServices = make([]DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaServices, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationPoliciesGlobalFirewallRulesDestinationsCriteriaServices{}
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
			data.SourcesCriteriaApplianceVlans = make([]DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaApplianceVlans, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaApplianceVlans{}
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
			data.SourcesCriteriaPolicyObjectGroups = make([]DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaPolicyObjectGroups, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaPolicyObjectGroups{}
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
			data.SourcesCriteriaPolicyObjects = make([]DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaPolicyObjects, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationPoliciesGlobalFirewallRulesSourcesCriteriaPolicyObjects{}
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
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
