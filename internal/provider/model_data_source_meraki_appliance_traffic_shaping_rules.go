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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSourceApplianceTrafficShapingRules - used only by the data source, always including every
// attribute (unlike model.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceApplianceTrafficShapingRules struct {
	Id                  types.String                                  `tfsdk:"id"`
	NetworkId           types.String                                  `tfsdk:"network_id"`
	DefaultRulesEnabled types.Bool                                    `tfsdk:"default_rules_enabled"`
	Rules               []DataSourceApplianceTrafficShapingRulesRules `tfsdk:"rules"`
}

type DataSourceApplianceTrafficShapingRulesRules struct {
	DscpTagValue                    types.Int64                                              `tfsdk:"dscp_tag_value"`
	Priority                        types.String                                             `tfsdk:"priority"`
	PerClientBandwidthLimitSettings types.String                                             `tfsdk:"per_client_bandwidth_limit_settings"`
	PerClientBandwidthLimitDown     types.Int64                                              `tfsdk:"per_client_bandwidth_limit_down"`
	PerClientBandwidthLimitUp       types.Int64                                              `tfsdk:"per_client_bandwidth_limit_up"`
	Definitions                     []DataSourceApplianceTrafficShapingRulesRulesDefinitions `tfsdk:"definitions"`
}

type DataSourceApplianceTrafficShapingRulesRulesDefinitions struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceTrafficShapingRules) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/trafficShaping/rules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// fromBody is de-generated (kept in sync manually with the resource-side model.go's custom fromBody) because
// `definitions[].type` discriminates whether `value` is read from the raw `value` path or `value.id`.
func (data *DataSourceApplianceTrafficShapingRules) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("defaultRulesEnabled"); value.Exists() && value.Value() != nil {
		data.DefaultRulesEnabled = types.BoolValue(value.Bool())
	} else {
		data.DefaultRulesEnabled = types.BoolNull()
	}
	if value := res.Get("rules"); value.Exists() && value.Value() != nil {
		data.Rules = make([]DataSourceApplianceTrafficShapingRulesRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceTrafficShapingRulesRules{}
			if value := res.Get("dscpTagValue"); value.Exists() && value.Value() != nil {
				data.DscpTagValue = types.Int64Value(value.Int())
			} else {
				data.DscpTagValue = types.Int64Null()
			}
			if value := res.Get("priority"); value.Exists() && value.Value() != nil {
				data.Priority = types.StringValue(value.String())
			} else {
				data.Priority = types.StringNull()
			}
			if value := res.Get("perClientBandwidthLimits.settings"); value.Exists() && value.Value() != nil {
				data.PerClientBandwidthLimitSettings = types.StringValue(value.String())
			} else {
				data.PerClientBandwidthLimitSettings = types.StringNull()
			}
			if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitDown"); value.Exists() && value.Value() != nil {
				data.PerClientBandwidthLimitDown = types.Int64Value(value.Int())
			} else {
				data.PerClientBandwidthLimitDown = types.Int64Null()
			}
			if value := res.Get("perClientBandwidthLimits.bandwidthLimits.limitUp"); value.Exists() && value.Value() != nil {
				data.PerClientBandwidthLimitUp = types.Int64Value(value.Int())
			} else {
				data.PerClientBandwidthLimitUp = types.Int64Null()
			}
			if value := res.Get("definitions"); value.Exists() && value.Value() != nil {
				data.Definitions = make([]DataSourceApplianceTrafficShapingRulesRulesDefinitions, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DataSourceApplianceTrafficShapingRulesRulesDefinitions{}
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
					if value := res.Get(valuePath); value.Exists() && value.Value() != nil {
						data.Value = types.StringValue(value.String())
					} else {
						data.Value = types.StringNull()
					}
					(*parent).Definitions = append((*parent).Definitions, data)
					return true
				})
			}
			(*parent).Rules = append((*parent).Rules, data)
			return true
		})
	}
}
