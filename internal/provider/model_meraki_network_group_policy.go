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

type NetworkGroupPolicy struct {
	Id                                           types.String                               `tfsdk:"id"`
	NetworkId                                    types.String                               `tfsdk:"network_id"`
	Name                                         types.String                               `tfsdk:"name"`
	SplashAuthSettings                           types.String                               `tfsdk:"splash_auth_settings"`
	BandwidthSettings                            types.String                               `tfsdk:"bandwidth_settings"`
	BandwidthLimitDown                           types.Int64                                `tfsdk:"bandwidth_limit_down"`
	BandwidthLimitUp                             types.Int64                                `tfsdk:"bandwidth_limit_up"`
	BonjourForwardingSettings                    types.String                               `tfsdk:"bonjour_forwarding_settings"`
	BonjourForwardingRules                       []NetworkGroupPolicyBonjourForwardingRules `tfsdk:"bonjour_forwarding_rules"`
	ContentFilteringAllowedUrlPatternsSettings   types.String                               `tfsdk:"content_filtering_allowed_url_patterns_settings"`
	ContentFilteringAllowedUrlPatterns           types.Set                                  `tfsdk:"content_filtering_allowed_url_patterns"`
	ContentFilteringBlockedUrlCategoriesSettings types.String                               `tfsdk:"content_filtering_blocked_url_categories_settings"`
	ContentFilteringBlockedUrlCategories         types.Set                                  `tfsdk:"content_filtering_blocked_url_categories"`
	ContentFilteringBlockedUrlPatternsSettings   types.String                               `tfsdk:"content_filtering_blocked_url_patterns_settings"`
	ContentFilteringBlockedUrlPatterns           types.Set                                  `tfsdk:"content_filtering_blocked_url_patterns"`
	FirewallAndTrafficShapingSettings            types.String                               `tfsdk:"firewall_and_traffic_shaping_settings"`
	L3FirewallRules                              []NetworkGroupPolicyL3FirewallRules        `tfsdk:"l3_firewall_rules"`
	L7FirewallRules                              []NetworkGroupPolicyL7FirewallRules        `tfsdk:"l7_firewall_rules"`
	TrafficShapingRules                          []NetworkGroupPolicyTrafficShapingRules    `tfsdk:"traffic_shaping_rules"`
	SchedulingEnabled                            types.Bool                                 `tfsdk:"scheduling_enabled"`
	SchedulingFridayActive                       types.Bool                                 `tfsdk:"scheduling_friday_active"`
	SchedulingFridayFrom                         types.String                               `tfsdk:"scheduling_friday_from"`
	SchedulingFridayTo                           types.String                               `tfsdk:"scheduling_friday_to"`
	SchedulingMondayActive                       types.Bool                                 `tfsdk:"scheduling_monday_active"`
	SchedulingMondayFrom                         types.String                               `tfsdk:"scheduling_monday_from"`
	SchedulingMondayTo                           types.String                               `tfsdk:"scheduling_monday_to"`
	SchedulingSaturdayActive                     types.Bool                                 `tfsdk:"scheduling_saturday_active"`
	SchedulingSaturdayFrom                       types.String                               `tfsdk:"scheduling_saturday_from"`
	SchedulingSaturdayTo                         types.String                               `tfsdk:"scheduling_saturday_to"`
	SchedulingSundayActive                       types.Bool                                 `tfsdk:"scheduling_sunday_active"`
	SchedulingSundayFrom                         types.String                               `tfsdk:"scheduling_sunday_from"`
	SchedulingSundayTo                           types.String                               `tfsdk:"scheduling_sunday_to"`
	SchedulingThursdayActive                     types.Bool                                 `tfsdk:"scheduling_thursday_active"`
	SchedulingThursdayFrom                       types.String                               `tfsdk:"scheduling_thursday_from"`
	SchedulingThursdayTo                         types.String                               `tfsdk:"scheduling_thursday_to"`
	SchedulingTuesdayActive                      types.Bool                                 `tfsdk:"scheduling_tuesday_active"`
	SchedulingTuesdayFrom                        types.String                               `tfsdk:"scheduling_tuesday_from"`
	SchedulingTuesdayTo                          types.String                               `tfsdk:"scheduling_tuesday_to"`
	SchedulingWednesdayActive                    types.Bool                                 `tfsdk:"scheduling_wednesday_active"`
	SchedulingWednesdayFrom                      types.String                               `tfsdk:"scheduling_wednesday_from"`
	SchedulingWednesdayTo                        types.String                               `tfsdk:"scheduling_wednesday_to"`
	VlanTaggingSettings                          types.String                               `tfsdk:"vlan_tagging_settings"`
	VlanTaggingVlanId                            types.String                               `tfsdk:"vlan_tagging_vlan_id"`
}

type NetworkGroupPolicyBonjourForwardingRules struct {
	Description types.String `tfsdk:"description"`
	VlanId      types.String `tfsdk:"vlan_id"`
	Services    types.Set    `tfsdk:"services"`
}

type NetworkGroupPolicyL3FirewallRules struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

type NetworkGroupPolicyL7FirewallRules struct {
	Policy types.String `tfsdk:"policy"`
	Type   types.String `tfsdk:"type"`
	Value  types.String `tfsdk:"value"`
}

type NetworkGroupPolicyTrafficShapingRules struct {
	DscpTagValue                                     types.Int64                                        `tfsdk:"dscp_tag_value"`
	PcpTagValue                                      types.Int64                                        `tfsdk:"pcp_tag_value"`
	Priority                                         types.String                                       `tfsdk:"priority"`
	PerClientBandwidthLimitsSettings                 types.String                                       `tfsdk:"per_client_bandwidth_limits_settings"`
	PerClientBandwidthLimitsBandwidthLimitsLimitDown types.Int64                                        `tfsdk:"per_client_bandwidth_limits_bandwidth_limits_limit_down"`
	PerClientBandwidthLimitsBandwidthLimitsLimitUp   types.Int64                                        `tfsdk:"per_client_bandwidth_limits_bandwidth_limits_limit_up"`
	Definitions                                      []NetworkGroupPolicyTrafficShapingRulesDefinitions `tfsdk:"definitions"`
}

type NetworkGroupPolicyTrafficShapingRulesDefinitions struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkGroupPolicy) getPath() string {
	return fmt.Sprintf("/networks/%v/groupPolicies", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkGroupPolicy) toBody(ctx context.Context, state NetworkGroupPolicy) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.SplashAuthSettings.IsNull() {
		body, _ = sjson.Set(body, "splashAuthSettings", data.SplashAuthSettings.ValueString())
	}
	if !data.BandwidthSettings.IsNull() {
		body, _ = sjson.Set(body, "bandwidth.settings", data.BandwidthSettings.ValueString())
	}
	if !data.BandwidthLimitDown.IsNull() {
		body, _ = sjson.Set(body, "bandwidth.bandwidthLimits.limitDown", data.BandwidthLimitDown.ValueInt64())
	}
	if !data.BandwidthLimitUp.IsNull() {
		body, _ = sjson.Set(body, "bandwidth.bandwidthLimits.limitUp", data.BandwidthLimitUp.ValueInt64())
	}
	if !data.BonjourForwardingSettings.IsNull() {
		body, _ = sjson.Set(body, "bonjourForwarding.settings", data.BonjourForwardingSettings.ValueString())
	}
	if len(data.BonjourForwardingRules) > 0 {
		body, _ = sjson.Set(body, "bonjourForwarding.rules", []interface{}{})
		for _, item := range data.BonjourForwardingRules {
			itemBody := ""
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.VlanId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanId", item.VlanId.ValueString())
			}
			if !item.Services.IsNull() {
				var values []string
				item.Services.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "services", values)
			}
			body, _ = sjson.SetRaw(body, "bonjourForwarding.rules.-1", itemBody)
		}
	}
	if !data.ContentFilteringAllowedUrlPatternsSettings.IsNull() {
		body, _ = sjson.Set(body, "contentFiltering.allowedUrlPatterns.settings", data.ContentFilteringAllowedUrlPatternsSettings.ValueString())
	}
	if !data.ContentFilteringAllowedUrlPatterns.IsNull() {
		var values []string
		data.ContentFilteringAllowedUrlPatterns.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "contentFiltering.allowedUrlPatterns.patterns", values)
	}
	if !data.ContentFilteringBlockedUrlCategoriesSettings.IsNull() {
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlCategories.settings", data.ContentFilteringBlockedUrlCategoriesSettings.ValueString())
	}
	if !data.ContentFilteringBlockedUrlCategories.IsNull() {
		var values []string
		data.ContentFilteringBlockedUrlCategories.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlCategories.categories", values)
	}
	if !data.ContentFilteringBlockedUrlPatternsSettings.IsNull() {
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlPatterns.settings", data.ContentFilteringBlockedUrlPatternsSettings.ValueString())
	}
	if !data.ContentFilteringBlockedUrlPatterns.IsNull() {
		var values []string
		data.ContentFilteringBlockedUrlPatterns.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlPatterns.patterns", values)
	}
	if !data.FirewallAndTrafficShapingSettings.IsNull() {
		body, _ = sjson.Set(body, "firewallAndTrafficShaping.settings", data.FirewallAndTrafficShapingSettings.ValueString())
	}
	if len(data.L3FirewallRules) > 0 {
		body, _ = sjson.Set(body, "firewallAndTrafficShaping.l3FirewallRules", []interface{}{})
		for _, item := range data.L3FirewallRules {
			itemBody := ""
			if !item.Comment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "comment", item.Comment.ValueString())
			}
			if !item.DestCidr.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destCidr", item.DestCidr.ValueString())
			}
			if !item.DestPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destPort", item.DestPort.ValueString())
			}
			if !item.Policy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "policy", item.Policy.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			body, _ = sjson.SetRaw(body, "firewallAndTrafficShaping.l3FirewallRules.-1", itemBody)
		}
	}
	if len(data.L7FirewallRules) > 0 {
		body, _ = sjson.Set(body, "firewallAndTrafficShaping.l7FirewallRules", []interface{}{})
		for _, item := range data.L7FirewallRules {
			itemBody := ""
			if !item.Policy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "policy", item.Policy.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Value.ValueString())
			}
			body, _ = sjson.SetRaw(body, "firewallAndTrafficShaping.l7FirewallRules.-1", itemBody)
		}
	}
	if len(data.TrafficShapingRules) > 0 {
		body, _ = sjson.Set(body, "firewallAndTrafficShaping.trafficShapingRules", []interface{}{})
		for _, item := range data.TrafficShapingRules {
			itemBody := ""
			if !item.DscpTagValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dscpTagValue", item.DscpTagValue.ValueInt64())
			}
			if !item.PcpTagValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "pcpTagValue", item.PcpTagValue.ValueInt64())
			}
			if !item.Priority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "priority", item.Priority.ValueString())
			}
			if !item.PerClientBandwidthLimitsSettings.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.settings", item.PerClientBandwidthLimitsSettings.ValueString())
			}
			if !item.PerClientBandwidthLimitsBandwidthLimitsLimitDown.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.bandwidthLimits.limitDown", item.PerClientBandwidthLimitsBandwidthLimitsLimitDown.ValueInt64())
			}
			if !item.PerClientBandwidthLimitsBandwidthLimitsLimitUp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "perClientBandwidthLimits.bandwidthLimits.limitUp", item.PerClientBandwidthLimitsBandwidthLimitsLimitUp.ValueInt64())
			}
			{
				itemBody, _ = sjson.Set(itemBody, "definitions", []interface{}{})
				for _, childItem := range item.Definitions {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "definitions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "firewallAndTrafficShaping.trafficShapingRules.-1", itemBody)
		}
	}
	if !data.SchedulingEnabled.IsNull() {
		body, _ = sjson.Set(body, "scheduling.enabled", data.SchedulingEnabled.ValueBool())
	}
	if !data.SchedulingFridayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.friday.active", data.SchedulingFridayActive.ValueBool())
	}
	if !data.SchedulingFridayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.friday.from", data.SchedulingFridayFrom.ValueString())
	}
	if !data.SchedulingFridayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.friday.to", data.SchedulingFridayTo.ValueString())
	}
	if !data.SchedulingMondayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.monday.active", data.SchedulingMondayActive.ValueBool())
	}
	if !data.SchedulingMondayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.monday.from", data.SchedulingMondayFrom.ValueString())
	}
	if !data.SchedulingMondayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.monday.to", data.SchedulingMondayTo.ValueString())
	}
	if !data.SchedulingSaturdayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.saturday.active", data.SchedulingSaturdayActive.ValueBool())
	}
	if !data.SchedulingSaturdayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.saturday.from", data.SchedulingSaturdayFrom.ValueString())
	}
	if !data.SchedulingSaturdayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.saturday.to", data.SchedulingSaturdayTo.ValueString())
	}
	if !data.SchedulingSundayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.sunday.active", data.SchedulingSundayActive.ValueBool())
	}
	if !data.SchedulingSundayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.sunday.from", data.SchedulingSundayFrom.ValueString())
	}
	if !data.SchedulingSundayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.sunday.to", data.SchedulingSundayTo.ValueString())
	}
	if !data.SchedulingThursdayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.thursday.active", data.SchedulingThursdayActive.ValueBool())
	}
	if !data.SchedulingThursdayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.thursday.from", data.SchedulingThursdayFrom.ValueString())
	}
	if !data.SchedulingThursdayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.thursday.to", data.SchedulingThursdayTo.ValueString())
	}
	if !data.SchedulingTuesdayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.tuesday.active", data.SchedulingTuesdayActive.ValueBool())
	}
	if !data.SchedulingTuesdayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.tuesday.from", data.SchedulingTuesdayFrom.ValueString())
	}
	if !data.SchedulingTuesdayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.tuesday.to", data.SchedulingTuesdayTo.ValueString())
	}
	if !data.SchedulingWednesdayActive.IsNull() {
		body, _ = sjson.Set(body, "scheduling.wednesday.active", data.SchedulingWednesdayActive.ValueBool())
	}
	if !data.SchedulingWednesdayFrom.IsNull() {
		body, _ = sjson.Set(body, "scheduling.wednesday.from", data.SchedulingWednesdayFrom.ValueString())
	}
	if !data.SchedulingWednesdayTo.IsNull() {
		body, _ = sjson.Set(body, "scheduling.wednesday.to", data.SchedulingWednesdayTo.ValueString())
	}
	if !data.VlanTaggingSettings.IsNull() {
		body, _ = sjson.Set(body, "vlanTagging.settings", data.VlanTaggingSettings.ValueString())
	}
	if !data.VlanTaggingVlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanTagging.vlanId", data.VlanTaggingVlanId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkGroupPolicy) fromBody(ctx context.Context, res meraki.Res) {
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
		data.BonjourForwardingRules = make([]NetworkGroupPolicyBonjourForwardingRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkGroupPolicyBonjourForwardingRules{}
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
		data.L3FirewallRules = make([]NetworkGroupPolicyL3FirewallRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkGroupPolicyL3FirewallRules{}
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
		data.L7FirewallRules = make([]NetworkGroupPolicyL7FirewallRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkGroupPolicyL7FirewallRules{}
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
		data.TrafficShapingRules = make([]NetworkGroupPolicyTrafficShapingRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkGroupPolicyTrafficShapingRules{}
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
				data.Definitions = make([]NetworkGroupPolicyTrafficShapingRulesDefinitions, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := NetworkGroupPolicyTrafficShapingRulesDefinitions{}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkGroupPolicy) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("splashAuthSettings"); value.Exists() && !data.SplashAuthSettings.IsNull() {
		data.SplashAuthSettings = types.StringValue(value.String())
	} else {
		data.SplashAuthSettings = types.StringNull()
	}
	if value := res.Get("bandwidth.settings"); value.Exists() && !data.BandwidthSettings.IsNull() {
		data.BandwidthSettings = types.StringValue(value.String())
	} else {
		data.BandwidthSettings = types.StringNull()
	}
	if value := res.Get("bandwidth.bandwidthLimits.limitDown"); value.Exists() && !data.BandwidthLimitDown.IsNull() {
		data.BandwidthLimitDown = types.Int64Value(value.Int())
	} else {
		data.BandwidthLimitDown = types.Int64Null()
	}
	if value := res.Get("bandwidth.bandwidthLimits.limitUp"); value.Exists() && !data.BandwidthLimitUp.IsNull() {
		data.BandwidthLimitUp = types.Int64Value(value.Int())
	} else {
		data.BandwidthLimitUp = types.Int64Null()
	}
	if value := res.Get("bonjourForwarding.settings"); value.Exists() && !data.BonjourForwardingSettings.IsNull() {
		data.BonjourForwardingSettings = types.StringValue(value.String())
	} else {
		data.BonjourForwardingSettings = types.StringNull()
	}
	for i := 0; i < len(data.BonjourForwardingRules); i++ {
		keys := [...]string{"vlanId"}
		keyValues := [...]string{data.BonjourForwardingRules[i].VlanId.ValueString()}

		parent := &data
		data := (*parent).BonjourForwardingRules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("bonjourForwarding.rules").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing BonjourForwardingRules[%d] = %+v",
				i,
				(*parent).BonjourForwardingRules[i],
			))
			(*parent).BonjourForwardingRules = slices.Delete((*parent).BonjourForwardingRules, i, i+1)
			i--

			continue
		}
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
			data.VlanId = types.StringValue(value.String())
		} else {
			data.VlanId = types.StringNull()
		}
		if value := res.Get("services"); value.Exists() && !data.Services.IsNull() {
			data.Services = helpers.GetStringSet(value.Array())
		} else {
			data.Services = types.SetNull(types.StringType)
		}
		(*parent).BonjourForwardingRules[i] = data
	}
	if value := res.Get("contentFiltering.allowedUrlPatterns.settings"); value.Exists() && !data.ContentFilteringAllowedUrlPatternsSettings.IsNull() {
		data.ContentFilteringAllowedUrlPatternsSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringAllowedUrlPatternsSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.allowedUrlPatterns.patterns"); value.Exists() && !data.ContentFilteringAllowedUrlPatterns.IsNull() {
		data.ContentFilteringAllowedUrlPatterns = helpers.GetStringSet(value.Array())
	} else {
		data.ContentFilteringAllowedUrlPatterns = types.SetNull(types.StringType)
	}
	if value := res.Get("contentFiltering.blockedUrlCategories.settings"); value.Exists() && !data.ContentFilteringBlockedUrlCategoriesSettings.IsNull() {
		data.ContentFilteringBlockedUrlCategoriesSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringBlockedUrlCategoriesSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.blockedUrlCategories.categories"); value.Exists() && !data.ContentFilteringBlockedUrlCategories.IsNull() {
		data.ContentFilteringBlockedUrlCategories = helpers.GetStringSet(value.Array())
	} else {
		data.ContentFilteringBlockedUrlCategories = types.SetNull(types.StringType)
	}
	if value := res.Get("contentFiltering.blockedUrlPatterns.settings"); value.Exists() && !data.ContentFilteringBlockedUrlPatternsSettings.IsNull() {
		data.ContentFilteringBlockedUrlPatternsSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringBlockedUrlPatternsSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.blockedUrlPatterns.patterns"); value.Exists() && !data.ContentFilteringBlockedUrlPatterns.IsNull() {
		data.ContentFilteringBlockedUrlPatterns = helpers.GetStringSet(value.Array())
	} else {
		data.ContentFilteringBlockedUrlPatterns = types.SetNull(types.StringType)
	}
	if value := res.Get("firewallAndTrafficShaping.settings"); value.Exists() && !data.FirewallAndTrafficShapingSettings.IsNull() {
		data.FirewallAndTrafficShapingSettings = types.StringValue(value.String())
	} else {
		data.FirewallAndTrafficShapingSettings = types.StringNull()
	}
	{
		l := len(res.Get("firewallAndTrafficShaping.l3FirewallRules").Array())
		tflog.Debug(ctx, fmt.Sprintf("firewallAndTrafficShaping.l3FirewallRules array resizing from %d to %d", len(data.L3FirewallRules), l))
		if len(data.L3FirewallRules) > l {
			data.L3FirewallRules = data.L3FirewallRules[:l]
		}
	}
	for i := range data.L3FirewallRules {
		parent := &data
		data := (*parent).L3FirewallRules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("firewallAndTrafficShaping.l3FirewallRules.%d", i))
		if value := res.Get("comment"); value.Exists() && !data.Comment.IsNull() {
			data.Comment = types.StringValue(value.String())
		} else {
			data.Comment = types.StringNull()
		}
		if value := res.Get("destCidr"); value.Exists() && !data.DestCidr.IsNull() {
			data.DestCidr = types.StringValue(value.String())
		} else {
			data.DestCidr = types.StringNull()
		}
		if value := res.Get("destPort"); value.Exists() && !data.DestPort.IsNull() {
			data.DestPort = types.StringValue(value.String())
		} else {
			data.DestPort = types.StringNull()
		}
		if value := res.Get("policy"); value.Exists() && !data.Policy.IsNull() {
			data.Policy = types.StringValue(value.String())
		} else {
			data.Policy = types.StringNull()
		}
		if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		(*parent).L3FirewallRules[i] = data
	}
	{
		l := len(res.Get("firewallAndTrafficShaping.l7FirewallRules").Array())
		tflog.Debug(ctx, fmt.Sprintf("firewallAndTrafficShaping.l7FirewallRules array resizing from %d to %d", len(data.L7FirewallRules), l))
		if len(data.L7FirewallRules) > l {
			data.L7FirewallRules = data.L7FirewallRules[:l]
		}
	}
	for i := range data.L7FirewallRules {
		parent := &data
		data := (*parent).L7FirewallRules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("firewallAndTrafficShaping.l7FirewallRules.%d", i))
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
		if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
			data.Value = types.StringValue(value.String())
		} else {
			data.Value = types.StringNull()
		}
		(*parent).L7FirewallRules[i] = data
	}
	{
		l := len(res.Get("firewallAndTrafficShaping.trafficShapingRules").Array())
		tflog.Debug(ctx, fmt.Sprintf("firewallAndTrafficShaping.trafficShapingRules array resizing from %d to %d", len(data.TrafficShapingRules), l))
		if len(data.TrafficShapingRules) > l {
			data.TrafficShapingRules = data.TrafficShapingRules[:l]
		}
	}
	for i := range data.TrafficShapingRules {
		parent := &data
		data := (*parent).TrafficShapingRules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("firewallAndTrafficShaping.trafficShapingRules.%d", i))
		if value := res.Get("dscpTagValue"); value.Exists() && !data.DscpTagValue.IsNull() {
			data.DscpTagValue = types.Int64Value(value.Int())
		} else {
			data.DscpTagValue = types.Int64Null()
		}
		if value := res.Get("pcpTagValue"); value.Exists() && !data.PcpTagValue.IsNull() {
			data.PcpTagValue = types.Int64Value(value.Int())
		} else {
			data.PcpTagValue = types.Int64Null()
		}
		if value := res.Get("perClientBandwidthLimits.settings"); value.Exists() && !data.PerClientBandwidthLimitsSettings.IsNull() {
			data.PerClientBandwidthLimitsSettings = types.StringValue(value.String())
		} else {
			data.PerClientBandwidthLimitsSettings = types.StringNull()
		}
		if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitDown"); value.Exists() && !data.PerClientBandwidthLimitsBandwidthLimitsLimitDown.IsNull() {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitDown = types.Int64Null()
		}
		if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitUp"); value.Exists() && !data.PerClientBandwidthLimitsBandwidthLimitsLimitUp.IsNull() {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Value(value.Int())
		} else {
			data.PerClientBandwidthLimitsBandwidthLimitsLimitUp = types.Int64Null()
		}
		for i := 0; i < len(data.Definitions); i++ {
			keys := [...]string{"type", "value"}
			keyValues := [...]string{data.Definitions[i].Type.ValueString(), data.Definitions[i].Value.ValueString()}

			parent := &data
			data := (*parent).Definitions[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("definitions").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Definitions[%d] = %+v",
					i,
					(*parent).Definitions[i],
				))
				(*parent).Definitions = slices.Delete((*parent).Definitions, i, i+1)
				i--

				continue
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).Definitions[i] = data
		}
		(*parent).TrafficShapingRules[i] = data
	}
	if value := res.Get("scheduling.enabled"); value.Exists() && !data.SchedulingEnabled.IsNull() {
		data.SchedulingEnabled = types.BoolValue(value.Bool())
	} else {
		data.SchedulingEnabled = types.BoolNull()
	}
	if value := res.Get("scheduling.friday.active"); value.Exists() && !data.SchedulingFridayActive.IsNull() {
		data.SchedulingFridayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingFridayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.friday.from"); value.Exists() && !data.SchedulingFridayFrom.IsNull() {
		data.SchedulingFridayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingFridayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.friday.to"); value.Exists() && !data.SchedulingFridayTo.IsNull() {
		data.SchedulingFridayTo = types.StringValue(value.String())
	} else {
		data.SchedulingFridayTo = types.StringNull()
	}
	if value := res.Get("scheduling.monday.active"); value.Exists() && !data.SchedulingMondayActive.IsNull() {
		data.SchedulingMondayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingMondayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.monday.from"); value.Exists() && !data.SchedulingMondayFrom.IsNull() {
		data.SchedulingMondayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingMondayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.monday.to"); value.Exists() && !data.SchedulingMondayTo.IsNull() {
		data.SchedulingMondayTo = types.StringValue(value.String())
	} else {
		data.SchedulingMondayTo = types.StringNull()
	}
	if value := res.Get("scheduling.saturday.active"); value.Exists() && !data.SchedulingSaturdayActive.IsNull() {
		data.SchedulingSaturdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingSaturdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.saturday.from"); value.Exists() && !data.SchedulingSaturdayFrom.IsNull() {
		data.SchedulingSaturdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingSaturdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.saturday.to"); value.Exists() && !data.SchedulingSaturdayTo.IsNull() {
		data.SchedulingSaturdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingSaturdayTo = types.StringNull()
	}
	if value := res.Get("scheduling.sunday.active"); value.Exists() && !data.SchedulingSundayActive.IsNull() {
		data.SchedulingSundayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingSundayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.sunday.from"); value.Exists() && !data.SchedulingSundayFrom.IsNull() {
		data.SchedulingSundayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingSundayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.sunday.to"); value.Exists() && !data.SchedulingSundayTo.IsNull() {
		data.SchedulingSundayTo = types.StringValue(value.String())
	} else {
		data.SchedulingSundayTo = types.StringNull()
	}
	if value := res.Get("scheduling.thursday.active"); value.Exists() && !data.SchedulingThursdayActive.IsNull() {
		data.SchedulingThursdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingThursdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.thursday.from"); value.Exists() && !data.SchedulingThursdayFrom.IsNull() {
		data.SchedulingThursdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingThursdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.thursday.to"); value.Exists() && !data.SchedulingThursdayTo.IsNull() {
		data.SchedulingThursdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingThursdayTo = types.StringNull()
	}
	if value := res.Get("scheduling.tuesday.active"); value.Exists() && !data.SchedulingTuesdayActive.IsNull() {
		data.SchedulingTuesdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingTuesdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.tuesday.from"); value.Exists() && !data.SchedulingTuesdayFrom.IsNull() {
		data.SchedulingTuesdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingTuesdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.tuesday.to"); value.Exists() && !data.SchedulingTuesdayTo.IsNull() {
		data.SchedulingTuesdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingTuesdayTo = types.StringNull()
	}
	if value := res.Get("scheduling.wednesday.active"); value.Exists() && !data.SchedulingWednesdayActive.IsNull() {
		data.SchedulingWednesdayActive = types.BoolValue(value.Bool())
	} else {
		data.SchedulingWednesdayActive = types.BoolNull()
	}
	if value := res.Get("scheduling.wednesday.from"); value.Exists() && !data.SchedulingWednesdayFrom.IsNull() {
		data.SchedulingWednesdayFrom = types.StringValue(value.String())
	} else {
		data.SchedulingWednesdayFrom = types.StringNull()
	}
	if value := res.Get("scheduling.wednesday.to"); value.Exists() && !data.SchedulingWednesdayTo.IsNull() {
		data.SchedulingWednesdayTo = types.StringValue(value.String())
	} else {
		data.SchedulingWednesdayTo = types.StringNull()
	}
	if value := res.Get("vlanTagging.settings"); value.Exists() && !data.VlanTaggingSettings.IsNull() {
		data.VlanTaggingSettings = types.StringValue(value.String())
	} else {
		data.VlanTaggingSettings = types.StringNull()
	}
	if value := res.Get("vlanTagging.vlanId"); value.Exists() && !data.VlanTaggingVlanId.IsNull() {
		data.VlanTaggingVlanId = types.StringValue(value.String())
	} else {
		data.VlanTaggingVlanId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkGroupPolicy) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkGroupPolicy) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
