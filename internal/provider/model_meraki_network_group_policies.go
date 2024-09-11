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
	"github.com/netascode/terraform-provider-meraki/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkGroupPolicies struct {
	Id                                             types.String                                 `tfsdk:"id"`
	NetworkId                                      types.String                                 `tfsdk:"network_id"`
	Name                                           types.String                                 `tfsdk:"name"`
	BonjourForwardingSettings                      types.String                                 `tfsdk:"bonjour_forwarding_settings"`
	BonjourForwardingRules                         []NetworkGroupPoliciesBonjourForwardingRules `tfsdk:"bonjour_forwarding_rules"`
	BandwidthBandwidthLimitsLimitDown              types.Int64                                  `tfsdk:"bandwidth_bandwidth_limits_limit_down"`
	BandwidthBandwidthLimitsLimitUp                types.Int64                                  `tfsdk:"bandwidth_bandwidth_limits_limit_up"`
	BandwidthSettings                              types.String                                 `tfsdk:"bandwidth_settings"`
	ContentFilteringAllowedUrlPatternsSettings     types.String                                 `tfsdk:"content_filtering_allowed_url_patterns_settings"`
	ContentFilteringBlockedUrlCategoriesCategories types.List                                   `tfsdk:"content_filtering_blocked_url_categories_categories"`
	ContentFilteringBlockedUrlCategoriesSettings   types.String                                 `tfsdk:"content_filtering_blocked_url_categories_settings"`
	ContentFilteringBlockedUrlPatternsPatterns     types.List                                   `tfsdk:"content_filtering_blocked_url_patterns_patterns"`
	ContentFilteringBlockedUrlPatternsSettings     types.String                                 `tfsdk:"content_filtering_blocked_url_patterns_settings"`
}

type NetworkGroupPoliciesBonjourForwardingRules struct {
	Description types.String `tfsdk:"description"`
	Services    types.List   `tfsdk:"services"`
	VlanId      types.String `tfsdk:"vlan_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkGroupPolicies) getPath() string {
	return fmt.Sprintf("/networks/%v/groupPolicies", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkGroupPolicies) toBody(ctx context.Context, state NetworkGroupPolicies) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "groupPolicyId", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
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
			if !item.Services.IsNull() {
				var values []string
				item.Services.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "services", values)
			}
			if !item.VlanId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlanId", item.VlanId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "bonjourForwarding.rules.-1", itemBody)
		}
	}
	if !data.BandwidthBandwidthLimitsLimitDown.IsNull() {
		body, _ = sjson.Set(body, "bandwidth.bandwidthLimits.limitDown", data.BandwidthBandwidthLimitsLimitDown.ValueInt64())
	}
	if !data.BandwidthBandwidthLimitsLimitUp.IsNull() {
		body, _ = sjson.Set(body, "bandwidth.bandwidthLimits.limitUp", data.BandwidthBandwidthLimitsLimitUp.ValueInt64())
	}
	if !data.BandwidthSettings.IsNull() {
		body, _ = sjson.Set(body, "bandwidth.settings", data.BandwidthSettings.ValueString())
	}
	if !data.ContentFilteringAllowedUrlPatternsSettings.IsNull() {
		body, _ = sjson.Set(body, "contentFiltering.allowedUrlPatterns.settings", data.ContentFilteringAllowedUrlPatternsSettings.ValueString())
	}
	if !data.ContentFilteringBlockedUrlCategoriesCategories.IsNull() {
		var values []string
		data.ContentFilteringBlockedUrlCategoriesCategories.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlCategories.categories", values)
	}
	if !data.ContentFilteringBlockedUrlCategoriesSettings.IsNull() {
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlCategories.settings", data.ContentFilteringBlockedUrlCategoriesSettings.ValueString())
	}
	if !data.ContentFilteringBlockedUrlPatternsPatterns.IsNull() {
		var values []string
		data.ContentFilteringBlockedUrlPatternsPatterns.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlPatterns.patterns", values)
	}
	if !data.ContentFilteringBlockedUrlPatternsSettings.IsNull() {
		body, _ = sjson.Set(body, "contentFiltering.blockedUrlPatterns.settings", data.ContentFilteringBlockedUrlPatternsSettings.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkGroupPolicies) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("bonjourForwarding.settings"); value.Exists() {
		data.BonjourForwardingSettings = types.StringValue(value.String())
	} else {
		data.BonjourForwardingSettings = types.StringNull()
	}
	if value := res.Get("bonjourForwarding.rules"); value.Exists() {
		data.BonjourForwardingRules = make([]NetworkGroupPoliciesBonjourForwardingRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkGroupPoliciesBonjourForwardingRules{}
			if value := res.Get("description"); value.Exists() {
				data.Description = types.StringValue(value.String())
			} else {
				data.Description = types.StringNull()
			}
			if value := res.Get("services"); value.Exists() {
				data.Services = helpers.GetStringList(value.Array())
			} else {
				data.Services = types.ListNull(types.StringType)
			}
			if value := res.Get("vlanId"); value.Exists() {
				data.VlanId = types.StringValue(value.String())
			} else {
				data.VlanId = types.StringNull()
			}
			(*parent).BonjourForwardingRules = append((*parent).BonjourForwardingRules, data)
			return true
		})
	}
	if value := res.Get("bandwidth.bandwidthLimits.limitDown"); value.Exists() {
		data.BandwidthBandwidthLimitsLimitDown = types.Int64Value(value.Int())
	} else {
		data.BandwidthBandwidthLimitsLimitDown = types.Int64Null()
	}
	if value := res.Get("bandwidth.bandwidthLimits.limitUp"); value.Exists() {
		data.BandwidthBandwidthLimitsLimitUp = types.Int64Value(value.Int())
	} else {
		data.BandwidthBandwidthLimitsLimitUp = types.Int64Null()
	}
	if value := res.Get("bandwidth.settings"); value.Exists() {
		data.BandwidthSettings = types.StringValue(value.String())
	} else {
		data.BandwidthSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.allowedUrlPatterns.settings"); value.Exists() {
		data.ContentFilteringAllowedUrlPatternsSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringAllowedUrlPatternsSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.blockedUrlCategories.categories"); value.Exists() {
		data.ContentFilteringBlockedUrlCategoriesCategories = helpers.GetStringList(value.Array())
	} else {
		data.ContentFilteringBlockedUrlCategoriesCategories = types.ListNull(types.StringType)
	}
	if value := res.Get("contentFiltering.blockedUrlCategories.settings"); value.Exists() {
		data.ContentFilteringBlockedUrlCategoriesSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringBlockedUrlCategoriesSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.blockedUrlPatterns.patterns"); value.Exists() {
		data.ContentFilteringBlockedUrlPatternsPatterns = helpers.GetStringList(value.Array())
	} else {
		data.ContentFilteringBlockedUrlPatternsPatterns = types.ListNull(types.StringType)
	}
	if value := res.Get("contentFiltering.blockedUrlPatterns.settings"); value.Exists() {
		data.ContentFilteringBlockedUrlPatternsSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringBlockedUrlPatternsSettings = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkGroupPolicies) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("bonjourForwarding.settings"); value.Exists() && !data.BonjourForwardingSettings.IsNull() {
		data.BonjourForwardingSettings = types.StringValue(value.String())
	} else {
		data.BonjourForwardingSettings = types.StringNull()
	}
	for i := 0; i < len(data.BonjourForwardingRules); i++ {
		keys := [...]string{"description", "vlanId"}
		keyValues := [...]string{data.BonjourForwardingRules[i].Description.ValueString(), data.BonjourForwardingRules[i].VlanId.ValueString()}

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
		if value := res.Get("services"); value.Exists() && !data.Services.IsNull() {
			data.Services = helpers.GetStringList(value.Array())
		} else {
			data.Services = types.ListNull(types.StringType)
		}
		if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
			data.VlanId = types.StringValue(value.String())
		} else {
			data.VlanId = types.StringNull()
		}
		(*parent).BonjourForwardingRules[i] = data
	}
	if value := res.Get("bandwidth.bandwidthLimits.limitDown"); value.Exists() && !data.BandwidthBandwidthLimitsLimitDown.IsNull() {
		data.BandwidthBandwidthLimitsLimitDown = types.Int64Value(value.Int())
	} else {
		data.BandwidthBandwidthLimitsLimitDown = types.Int64Null()
	}
	if value := res.Get("bandwidth.bandwidthLimits.limitUp"); value.Exists() && !data.BandwidthBandwidthLimitsLimitUp.IsNull() {
		data.BandwidthBandwidthLimitsLimitUp = types.Int64Value(value.Int())
	} else {
		data.BandwidthBandwidthLimitsLimitUp = types.Int64Null()
	}
	if value := res.Get("bandwidth.settings"); value.Exists() && !data.BandwidthSettings.IsNull() {
		data.BandwidthSettings = types.StringValue(value.String())
	} else {
		data.BandwidthSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.allowedUrlPatterns.settings"); value.Exists() && !data.ContentFilteringAllowedUrlPatternsSettings.IsNull() {
		data.ContentFilteringAllowedUrlPatternsSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringAllowedUrlPatternsSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.blockedUrlCategories.categories"); value.Exists() && !data.ContentFilteringBlockedUrlCategoriesCategories.IsNull() {
		data.ContentFilteringBlockedUrlCategoriesCategories = helpers.GetStringList(value.Array())
	} else {
		data.ContentFilteringBlockedUrlCategoriesCategories = types.ListNull(types.StringType)
	}
	if value := res.Get("contentFiltering.blockedUrlCategories.settings"); value.Exists() && !data.ContentFilteringBlockedUrlCategoriesSettings.IsNull() {
		data.ContentFilteringBlockedUrlCategoriesSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringBlockedUrlCategoriesSettings = types.StringNull()
	}
	if value := res.Get("contentFiltering.blockedUrlPatterns.patterns"); value.Exists() && !data.ContentFilteringBlockedUrlPatternsPatterns.IsNull() {
		data.ContentFilteringBlockedUrlPatternsPatterns = helpers.GetStringList(value.Array())
	} else {
		data.ContentFilteringBlockedUrlPatternsPatterns = types.ListNull(types.StringType)
	}
	if value := res.Get("contentFiltering.blockedUrlPatterns.settings"); value.Exists() && !data.ContentFilteringBlockedUrlPatternsSettings.IsNull() {
		data.ContentFilteringBlockedUrlPatternsSettings = types.StringValue(value.String())
	} else {
		data.ContentFilteringBlockedUrlPatternsSettings = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial
