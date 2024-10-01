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

type NetworkGroupPolicies struct {
	NetworkId types.String                `tfsdk:"network_id"`
	Items     []NetworkGroupPoliciesItems `tfsdk:"items"`
}

type NetworkGroupPoliciesItems struct {
	Id                                           types.String                                 `tfsdk:"id"`
	Name                                         types.String                                 `tfsdk:"name"`
	SplashAuthSettings                           types.String                                 `tfsdk:"splash_auth_settings"`
	BandwidthSettings                            types.String                                 `tfsdk:"bandwidth_settings"`
	BandwidthLimitDown                           types.Int64                                  `tfsdk:"bandwidth_limit_down"`
	BandwidthLimitUp                             types.Int64                                  `tfsdk:"bandwidth_limit_up"`
	BonjourForwardingSettings                    types.String                                 `tfsdk:"bonjour_forwarding_settings"`
	BonjourForwardingRules                       []NetworkGroupPoliciesBonjourForwardingRules `tfsdk:"bonjour_forwarding_rules"`
	ContentFilteringAllowedUrlPatternsSettings   types.String                                 `tfsdk:"content_filtering_allowed_url_patterns_settings"`
	ContentFilteringAllowedUrlPatterns           types.Set                                    `tfsdk:"content_filtering_allowed_url_patterns"`
	ContentFilteringBlockedUrlCategoriesSettings types.String                                 `tfsdk:"content_filtering_blocked_url_categories_settings"`
	ContentFilteringBlockedUrlCategories         types.Set                                    `tfsdk:"content_filtering_blocked_url_categories"`
	ContentFilteringBlockedUrlPatternsSettings   types.String                                 `tfsdk:"content_filtering_blocked_url_patterns_settings"`
	ContentFilteringBlockedUrlPatterns           types.Set                                    `tfsdk:"content_filtering_blocked_url_patterns"`
	FirewallAndTrafficShapingSettings            types.String                                 `tfsdk:"firewall_and_traffic_shaping_settings"`
	L3FirewallRules                              []NetworkGroupPoliciesL3FirewallRules        `tfsdk:"l3_firewall_rules"`
	L7FirewallRules                              []NetworkGroupPoliciesL7FirewallRules        `tfsdk:"l7_firewall_rules"`
	TrafficShapingRules                          []NetworkGroupPoliciesTrafficShapingRules    `tfsdk:"traffic_shaping_rules"`
	SchedulingEnabled                            types.Bool                                   `tfsdk:"scheduling_enabled"`
	SchedulingFridayActive                       types.Bool                                   `tfsdk:"scheduling_friday_active"`
	SchedulingFridayFrom                         types.String                                 `tfsdk:"scheduling_friday_from"`
	SchedulingFridayTo                           types.String                                 `tfsdk:"scheduling_friday_to"`
	SchedulingMondayActive                       types.Bool                                   `tfsdk:"scheduling_monday_active"`
	SchedulingMondayFrom                         types.String                                 `tfsdk:"scheduling_monday_from"`
	SchedulingMondayTo                           types.String                                 `tfsdk:"scheduling_monday_to"`
	SchedulingSaturdayActive                     types.Bool                                   `tfsdk:"scheduling_saturday_active"`
	SchedulingSaturdayFrom                       types.String                                 `tfsdk:"scheduling_saturday_from"`
	SchedulingSaturdayTo                         types.String                                 `tfsdk:"scheduling_saturday_to"`
	SchedulingSundayActive                       types.Bool                                   `tfsdk:"scheduling_sunday_active"`
	SchedulingSundayFrom                         types.String                                 `tfsdk:"scheduling_sunday_from"`
	SchedulingSundayTo                           types.String                                 `tfsdk:"scheduling_sunday_to"`
	SchedulingThursdayActive                     types.Bool                                   `tfsdk:"scheduling_thursday_active"`
	SchedulingThursdayFrom                       types.String                                 `tfsdk:"scheduling_thursday_from"`
	SchedulingThursdayTo                         types.String                                 `tfsdk:"scheduling_thursday_to"`
	SchedulingTuesdayActive                      types.Bool                                   `tfsdk:"scheduling_tuesday_active"`
	SchedulingTuesdayFrom                        types.String                                 `tfsdk:"scheduling_tuesday_from"`
	SchedulingTuesdayTo                          types.String                                 `tfsdk:"scheduling_tuesday_to"`
	SchedulingWednesdayActive                    types.Bool                                   `tfsdk:"scheduling_wednesday_active"`
	SchedulingWednesdayFrom                      types.String                                 `tfsdk:"scheduling_wednesday_from"`
	SchedulingWednesdayTo                        types.String                                 `tfsdk:"scheduling_wednesday_to"`
	VlanTaggingSettings                          types.String                                 `tfsdk:"vlan_tagging_settings"`
	VlanTaggingVlanId                            types.String                                 `tfsdk:"vlan_tagging_vlan_id"`
}

type NetworkGroupPoliciesBonjourForwardingRules struct {
	Description types.String `tfsdk:"description"`
	VlanId      types.String `tfsdk:"vlan_id"`
	Services    types.Set    `tfsdk:"services"`
}

type NetworkGroupPoliciesL3FirewallRules struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

type NetworkGroupPoliciesL7FirewallRules struct {
	Policy types.String `tfsdk:"policy"`
	Type   types.String `tfsdk:"type"`
	Value  types.String `tfsdk:"value"`
}

type NetworkGroupPoliciesTrafficShapingRules struct {
	DscpTagValue                                     types.Int64                                          `tfsdk:"dscp_tag_value"`
	PcpTagValue                                      types.Int64                                          `tfsdk:"pcp_tag_value"`
	Priority                                         types.String                                         `tfsdk:"priority"`
	PerClientBandwidthLimitsSettings                 types.String                                         `tfsdk:"per_client_bandwidth_limits_settings"`
	PerClientBandwidthLimitsBandwidthLimitsLimitDown types.Int64                                          `tfsdk:"per_client_bandwidth_limits_bandwidth_limits_limit_down"`
	PerClientBandwidthLimitsBandwidthLimitsLimitUp   types.Int64                                          `tfsdk:"per_client_bandwidth_limits_bandwidth_limits_limit_up"`
	Definitions                                      []NetworkGroupPoliciesTrafficShapingRulesDefinitions `tfsdk:"definitions"`
}

type NetworkGroupPoliciesTrafficShapingRulesDefinitions struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkGroupPolicies) getPath() string {
	return fmt.Sprintf("/networks/%v/groupPolicies", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkGroupPolicies) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]NetworkGroupPoliciesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := NetworkGroupPoliciesItems{}
		data.Id = types.StringValue(res.Get("groupPolicyId").String())
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("splashAuthSettings"); value.Exists() && value.Value() != nil {
			data.SplashAuthSettings = types.StringValue(value.String())
		} else {
			data.SplashAuthSettings = types.StringNull()
		}
		if value := res.Get("bandwidth.settings"); value.Exists() && value.Value() != nil {
			data.BandwidthSettings = types.StringValue(value.String())
		} else {
			data.BandwidthSettings = types.StringNull()
		}
		if value := res.Get("bandwidth.bandwidthLimits.limitDown"); value.Exists() && value.Value() != nil {
			data.BandwidthLimitDown = types.Int64Value(value.Int())
		} else {
			data.BandwidthLimitDown = types.Int64Null()
		}
		if value := res.Get("bandwidth.bandwidthLimits.limitUp"); value.Exists() && value.Value() != nil {
			data.BandwidthLimitUp = types.Int64Value(value.Int())
		} else {
			data.BandwidthLimitUp = types.Int64Null()
		}
		if value := res.Get("bonjourForwarding.settings"); value.Exists() && value.Value() != nil {
			data.BonjourForwardingSettings = types.StringValue(value.String())
		} else {
			data.BonjourForwardingSettings = types.StringNull()
		}
		if value := res.Get("bonjourForwarding.rules"); value.Exists() && value.Value() != nil {
			data.BonjourForwardingRules = make([]NetworkGroupPoliciesBonjourForwardingRules, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := NetworkGroupPoliciesBonjourForwardingRules{}
				if value := res.Get("description"); value.Exists() && value.Value() != nil {
					data.Description = types.StringValue(value.String())
				} else {
					data.Description = types.StringNull()
				}
				if value := res.Get("vlanId"); value.Exists() && value.Value() != nil {
					data.VlanId = types.StringValue(value.String())
				} else {
					data.VlanId = types.StringNull()
				}
				if value := res.Get("services"); value.Exists() && value.Value() != nil {
					data.Services = helpers.GetStringSet(value.Array())
				} else {
					data.Services = types.SetNull(types.StringType)
				}
				(*parent).BonjourForwardingRules = append((*parent).BonjourForwardingRules, data)
				return true
			})
		}
		if value := res.Get("contentFiltering.allowedUrlPatterns.settings"); value.Exists() && value.Value() != nil {
			data.ContentFilteringAllowedUrlPatternsSettings = types.StringValue(value.String())
		} else {
			data.ContentFilteringAllowedUrlPatternsSettings = types.StringNull()
		}
		if value := res.Get("contentFiltering.allowedUrlPatterns.patterns"); value.Exists() && value.Value() != nil {
			data.ContentFilteringAllowedUrlPatterns = helpers.GetStringSet(value.Array())
		} else {
			data.ContentFilteringAllowedUrlPatterns = types.SetNull(types.StringType)
		}
		if value := res.Get("contentFiltering.blockedUrlCategories.settings"); value.Exists() && value.Value() != nil {
			data.ContentFilteringBlockedUrlCategoriesSettings = types.StringValue(value.String())
		} else {
			data.ContentFilteringBlockedUrlCategoriesSettings = types.StringNull()
		}
		if value := res.Get("contentFiltering.blockedUrlCategories.categories"); value.Exists() && value.Value() != nil {
			data.ContentFilteringBlockedUrlCategories = helpers.GetStringSet(value.Array())
		} else {
			data.ContentFilteringBlockedUrlCategories = types.SetNull(types.StringType)
		}
		if value := res.Get("contentFiltering.blockedUrlPatterns.settings"); value.Exists() && value.Value() != nil {
			data.ContentFilteringBlockedUrlPatternsSettings = types.StringValue(value.String())
		} else {
			data.ContentFilteringBlockedUrlPatternsSettings = types.StringNull()
		}
		if value := res.Get("contentFiltering.blockedUrlPatterns.patterns"); value.Exists() && value.Value() != nil {
			data.ContentFilteringBlockedUrlPatterns = helpers.GetStringSet(value.Array())
		} else {
			data.ContentFilteringBlockedUrlPatterns = types.SetNull(types.StringType)
		}
		if value := res.Get("firewallAndTrafficShaping.settings"); value.Exists() && value.Value() != nil {
			data.FirewallAndTrafficShapingSettings = types.StringValue(value.String())
		} else {
			data.FirewallAndTrafficShapingSettings = types.StringNull()
		}
		if value := res.Get("firewallAndTrafficShaping.l3FirewallRules"); value.Exists() && value.Value() != nil {
			data.L3FirewallRules = make([]NetworkGroupPoliciesL3FirewallRules, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := NetworkGroupPoliciesL3FirewallRules{}
				if value := res.Get("comment"); value.Exists() && value.Value() != nil {
					data.Comment = types.StringValue(value.String())
				} else {
					data.Comment = types.StringNull()
				}
				if value := res.Get("destCidr"); value.Exists() && value.Value() != nil {
					data.DestCidr = types.StringValue(value.String())
				} else {
					data.DestCidr = types.StringNull()
				}
				if value := res.Get("destPort"); value.Exists() && value.Value() != nil {
					data.DestPort = types.StringValue(value.String())
				} else {
					data.DestPort = types.StringNull()
				}
				if value := res.Get("policy"); value.Exists() && value.Value() != nil {
					data.Policy = types.StringValue(value.String())
				} else {
					data.Policy = types.StringNull()
				}
				if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
					data.Protocol = types.StringValue(value.String())
				} else {
					data.Protocol = types.StringNull()
				}
				(*parent).L3FirewallRules = append((*parent).L3FirewallRules, data)
				return true
			})
		}
		if value := res.Get("firewallAndTrafficShaping.l7FirewallRules"); value.Exists() && value.Value() != nil {
			data.L7FirewallRules = make([]NetworkGroupPoliciesL7FirewallRules, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := NetworkGroupPoliciesL7FirewallRules{}
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
				if value := res.Get("value"); value.Exists() && value.Value() != nil {
					data.Value = types.StringValue(value.String())
				} else {
					data.Value = types.StringNull()
				}
				(*parent).L7FirewallRules = append((*parent).L7FirewallRules, data)
				return true
			})
		}
		if value := res.Get("firewallAndTrafficShaping.trafficShapingRules"); value.Exists() && value.Value() != nil {
			data.TrafficShapingRules = make([]NetworkGroupPoliciesTrafficShapingRules, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := NetworkGroupPoliciesTrafficShapingRules{}
				if value := res.Get("dscpTagValue"); value.Exists() && value.Value() != nil {
					data.DscpTagValue = types.Int64Value(value.Int())
				} else {
					data.DscpTagValue = types.Int64Null()
				}
				if value := res.Get("pcpTagValue"); value.Exists() && value.Value() != nil {
					data.PcpTagValue = types.Int64Value(value.Int())
				} else {
					data.PcpTagValue = types.Int64Null()
				}
				if value := res.Get("perClientBandwidthLimits.settings"); value.Exists() && value.Value() != nil {
					data.PerClientBandwidthLimitsSettings = types.StringValue(value.String())
				} else {
					data.PerClientBandwidthLimitsSettings = types.StringNull()
				}
				if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitDown"); value.Exists() && value.Value() != nil {
					data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Value(value.Int())
				} else {
					data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Null()
				}
				if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitUp"); value.Exists() && value.Value() != nil {
					data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Value(value.Int())
				} else {
					data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Null()
				}
				if value := res.Get("definitions"); value.Exists() && value.Value() != nil {
					data.Definitions = make([]NetworkGroupPoliciesTrafficShapingRulesDefinitions, 0)
					value.ForEach(func(k, res gjson.Result) bool {
						parent := &data
						data := NetworkGroupPoliciesTrafficShapingRulesDefinitions{}
						if value := res.Get("type"); value.Exists() && value.Value() != nil {
							data.Type = types.StringValue(value.String())
						} else {
							data.Type = types.StringNull()
						}
						if value := res.Get("value"); value.Exists() && value.Value() != nil {
							data.Value = types.StringValue(value.String())
						} else {
							data.Value = types.StringNull()
						}
						(*parent).Definitions = append((*parent).Definitions, data)
						return true
					})
				}
				(*parent).TrafficShapingRules = append((*parent).TrafficShapingRules, data)
				return true
			})
		}
		if value := res.Get("scheduling.enabled"); value.Exists() && value.Value() != nil {
			data.SchedulingEnabled = types.BoolValue(value.Bool())
		} else {
			data.SchedulingEnabled = types.BoolNull()
		}
		if value := res.Get("scheduling.friday.active"); value.Exists() && value.Value() != nil {
			data.SchedulingFridayActive = types.BoolValue(value.Bool())
		} else {
			data.SchedulingFridayActive = types.BoolNull()
		}
		if value := res.Get("scheduling.friday.from"); value.Exists() && value.Value() != nil {
			data.SchedulingFridayFrom = types.StringValue(value.String())
		} else {
			data.SchedulingFridayFrom = types.StringNull()
		}
		if value := res.Get("scheduling.friday.to"); value.Exists() && value.Value() != nil {
			data.SchedulingFridayTo = types.StringValue(value.String())
		} else {
			data.SchedulingFridayTo = types.StringNull()
		}
		if value := res.Get("scheduling.monday.active"); value.Exists() && value.Value() != nil {
			data.SchedulingMondayActive = types.BoolValue(value.Bool())
		} else {
			data.SchedulingMondayActive = types.BoolNull()
		}
		if value := res.Get("scheduling.monday.from"); value.Exists() && value.Value() != nil {
			data.SchedulingMondayFrom = types.StringValue(value.String())
		} else {
			data.SchedulingMondayFrom = types.StringNull()
		}
		if value := res.Get("scheduling.monday.to"); value.Exists() && value.Value() != nil {
			data.SchedulingMondayTo = types.StringValue(value.String())
		} else {
			data.SchedulingMondayTo = types.StringNull()
		}
		if value := res.Get("scheduling.saturday.active"); value.Exists() && value.Value() != nil {
			data.SchedulingSaturdayActive = types.BoolValue(value.Bool())
		} else {
			data.SchedulingSaturdayActive = types.BoolNull()
		}
		if value := res.Get("scheduling.saturday.from"); value.Exists() && value.Value() != nil {
			data.SchedulingSaturdayFrom = types.StringValue(value.String())
		} else {
			data.SchedulingSaturdayFrom = types.StringNull()
		}
		if value := res.Get("scheduling.saturday.to"); value.Exists() && value.Value() != nil {
			data.SchedulingSaturdayTo = types.StringValue(value.String())
		} else {
			data.SchedulingSaturdayTo = types.StringNull()
		}
		if value := res.Get("scheduling.sunday.active"); value.Exists() && value.Value() != nil {
			data.SchedulingSundayActive = types.BoolValue(value.Bool())
		} else {
			data.SchedulingSundayActive = types.BoolNull()
		}
		if value := res.Get("scheduling.sunday.from"); value.Exists() && value.Value() != nil {
			data.SchedulingSundayFrom = types.StringValue(value.String())
		} else {
			data.SchedulingSundayFrom = types.StringNull()
		}
		if value := res.Get("scheduling.sunday.to"); value.Exists() && value.Value() != nil {
			data.SchedulingSundayTo = types.StringValue(value.String())
		} else {
			data.SchedulingSundayTo = types.StringNull()
		}
		if value := res.Get("scheduling.thursday.active"); value.Exists() && value.Value() != nil {
			data.SchedulingThursdayActive = types.BoolValue(value.Bool())
		} else {
			data.SchedulingThursdayActive = types.BoolNull()
		}
		if value := res.Get("scheduling.thursday.from"); value.Exists() && value.Value() != nil {
			data.SchedulingThursdayFrom = types.StringValue(value.String())
		} else {
			data.SchedulingThursdayFrom = types.StringNull()
		}
		if value := res.Get("scheduling.thursday.to"); value.Exists() && value.Value() != nil {
			data.SchedulingThursdayTo = types.StringValue(value.String())
		} else {
			data.SchedulingThursdayTo = types.StringNull()
		}
		if value := res.Get("scheduling.tuesday.active"); value.Exists() && value.Value() != nil {
			data.SchedulingTuesdayActive = types.BoolValue(value.Bool())
		} else {
			data.SchedulingTuesdayActive = types.BoolNull()
		}
		if value := res.Get("scheduling.tuesday.from"); value.Exists() && value.Value() != nil {
			data.SchedulingTuesdayFrom = types.StringValue(value.String())
		} else {
			data.SchedulingTuesdayFrom = types.StringNull()
		}
		if value := res.Get("scheduling.tuesday.to"); value.Exists() && value.Value() != nil {
			data.SchedulingTuesdayTo = types.StringValue(value.String())
		} else {
			data.SchedulingTuesdayTo = types.StringNull()
		}
		if value := res.Get("scheduling.wednesday.active"); value.Exists() && value.Value() != nil {
			data.SchedulingWednesdayActive = types.BoolValue(value.Bool())
		} else {
			data.SchedulingWednesdayActive = types.BoolNull()
		}
		if value := res.Get("scheduling.wednesday.from"); value.Exists() && value.Value() != nil {
			data.SchedulingWednesdayFrom = types.StringValue(value.String())
		} else {
			data.SchedulingWednesdayFrom = types.StringNull()
		}
		if value := res.Get("scheduling.wednesday.to"); value.Exists() && value.Value() != nil {
			data.SchedulingWednesdayTo = types.StringValue(value.String())
		} else {
			data.SchedulingWednesdayTo = types.StringNull()
		}
		if value := res.Get("vlanTagging.settings"); value.Exists() && value.Value() != nil {
			data.VlanTaggingSettings = types.StringValue(value.String())
		} else {
			data.VlanTaggingSettings = types.StringNull()
		}
		if value := res.Get("vlanTagging.vlanId"); value.Exists() && value.Value() != nil {
			data.VlanTaggingVlanId = types.StringValue(value.String())
		} else {
			data.VlanTaggingVlanId = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
