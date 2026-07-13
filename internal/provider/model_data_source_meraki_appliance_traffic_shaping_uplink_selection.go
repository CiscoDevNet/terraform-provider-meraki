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
// It emits a separate model - DataSourceApplianceTrafficShapingUplinkSelection - used only by the data source, always including every
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

type DataSourceApplianceTrafficShapingUplinkSelection struct {
	Id                                  types.String                                                                  `tfsdk:"id"`
	NetworkId                           types.String                                                                  `tfsdk:"network_id"`
	ActiveActiveAutoVpnEnabled          types.Bool                                                                    `tfsdk:"active_active_auto_vpn_enabled"`
	DefaultUplink                       types.String                                                                  `tfsdk:"default_uplink"`
	LoadBalancingEnabled                types.Bool                                                                    `tfsdk:"load_balancing_enabled"`
	FailoverAndFailbackImmediateEnabled types.Bool                                                                    `tfsdk:"failover_and_failback_immediate_enabled"`
	VpnTrafficUplinkPreferences         []DataSourceApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences `tfsdk:"vpn_traffic_uplink_preferences"`
	WanTrafficUplinkPreferences         []DataSourceApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences `tfsdk:"wan_traffic_uplink_preferences"`
}

type DataSourceApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences struct {
	FailOverCriterion           types.String                                                                                `tfsdk:"fail_over_criterion"`
	PreferredUplink             types.String                                                                                `tfsdk:"preferred_uplink"`
	BuiltinPerformanceClassName types.String                                                                                `tfsdk:"builtin_performance_class_name"`
	CustomPerformanceClassId    types.String                                                                                `tfsdk:"custom_performance_class_id"`
	PerformanceClassType        types.String                                                                                `tfsdk:"performance_class_type"`
	TrafficFilters              []DataSourceApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters `tfsdk:"traffic_filters"`
}

type DataSourceApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences struct {
	PreferredUplink types.String                                                                                `tfsdk:"preferred_uplink"`
	TrafficFilters  []DataSourceApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters `tfsdk:"traffic_filters"`
}

type DataSourceApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters struct {
	Type               types.String `tfsdk:"type"`
	Id                 types.String `tfsdk:"id"`
	Protocol           types.String `tfsdk:"protocol"`
	DestinationCidr    types.String `tfsdk:"destination_cidr"`
	DestinationFqdn    types.String `tfsdk:"destination_fqdn"`
	DestinationHost    types.Int64  `tfsdk:"destination_host"`
	DestinationNetwork types.String `tfsdk:"destination_network"`
	DestinationPort    types.String `tfsdk:"destination_port"`
	DestinationVlan    types.Int64  `tfsdk:"destination_vlan"`
	SourceCidr         types.String `tfsdk:"source_cidr"`
	SourceHost         types.Int64  `tfsdk:"source_host"`
	SourceNetwork      types.String `tfsdk:"source_network"`
	SourcePort         types.String `tfsdk:"source_port"`
	SourceVlan         types.Int64  `tfsdk:"source_vlan"`
}

type DataSourceApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters struct {
	Type            types.String `tfsdk:"type"`
	Protocol        types.String `tfsdk:"protocol"`
	DestinationCidr types.String `tfsdk:"destination_cidr"`
	DestinationPort types.String `tfsdk:"destination_port"`
	SourceCidr      types.String `tfsdk:"source_cidr"`
	SourceHost      types.Int64  `tfsdk:"source_host"`
	SourcePort      types.String `tfsdk:"source_port"`
	SourceVlan      types.Int64  `tfsdk:"source_vlan"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceTrafficShapingUplinkSelection) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/trafficShaping/uplinkSelection", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceTrafficShapingUplinkSelection) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("activeActiveAutoVpnEnabled"); value.Exists() && value.Value() != nil {
		data.ActiveActiveAutoVpnEnabled = types.BoolValue(value.Bool())
	} else {
		data.ActiveActiveAutoVpnEnabled = types.BoolNull()
	}
	if value := res.Get("defaultUplink"); value.Exists() && value.Value() != nil {
		data.DefaultUplink = types.StringValue(value.String())
	} else {
		data.DefaultUplink = types.StringNull()
	}
	if value := res.Get("loadBalancingEnabled"); value.Exists() && value.Value() != nil {
		data.LoadBalancingEnabled = types.BoolValue(value.Bool())
	} else {
		data.LoadBalancingEnabled = types.BoolNull()
	}
	if value := res.Get("failoverAndFailback.immediate.enabled"); value.Exists() && value.Value() != nil {
		data.FailoverAndFailbackImmediateEnabled = types.BoolValue(value.Bool())
	} else {
		data.FailoverAndFailbackImmediateEnabled = types.BoolNull()
	}
	if value := res.Get("vpnTrafficUplinkPreferences"); value.Exists() && value.Value() != nil {
		data.VpnTrafficUplinkPreferences = make([]DataSourceApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences{}
			if value := res.Get("failOverCriterion"); value.Exists() && value.Value() != nil {
				data.FailOverCriterion = types.StringValue(value.String())
			} else {
				data.FailOverCriterion = types.StringNull()
			}
			if value := res.Get("preferredUplink"); value.Exists() && value.Value() != nil {
				data.PreferredUplink = types.StringValue(value.String())
			} else {
				data.PreferredUplink = types.StringNull()
			}
			if value := res.Get("performanceClass.builtinPerformanceClassName"); value.Exists() && value.Value() != nil {
				data.BuiltinPerformanceClassName = types.StringValue(value.String())
			} else {
				data.BuiltinPerformanceClassName = types.StringNull()
			}
			if value := res.Get("performanceClass.customPerformanceClassId"); value.Exists() && value.Value() != nil {
				data.CustomPerformanceClassId = types.StringValue(value.String())
			} else {
				data.CustomPerformanceClassId = types.StringNull()
			}
			if value := res.Get("performanceClass.type"); value.Exists() && value.Value() != nil {
				data.PerformanceClassType = types.StringValue(value.String())
			} else {
				data.PerformanceClassType = types.StringNull()
			}
			if value := res.Get("trafficFilters"); value.Exists() && value.Value() != nil {
				data.TrafficFilters = make([]DataSourceApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DataSourceApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters{}
					if value := res.Get("type"); value.Exists() && value.Value() != nil {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					if value := res.Get("value.id"); value.Exists() && value.Value() != nil {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("value.protocol"); value.Exists() && value.Value() != nil {
						data.Protocol = types.StringValue(value.String())
					} else {
						data.Protocol = types.StringNull()
					}
					if value := res.Get("value.destination.cidr"); value.Exists() && value.Value() != nil {
						data.DestinationCidr = types.StringValue(value.String())
					} else {
						data.DestinationCidr = types.StringNull()
					}
					if value := res.Get("value.destination.fqdn"); value.Exists() && value.Value() != nil {
						data.DestinationFqdn = types.StringValue(value.String())
					} else {
						data.DestinationFqdn = types.StringNull()
					}
					if value := res.Get("value.destination.host"); value.Exists() && value.Value() != nil {
						data.DestinationHost = types.Int64Value(value.Int())
					} else {
						data.DestinationHost = types.Int64Null()
					}
					if value := res.Get("value.destination.network"); value.Exists() && value.Value() != nil {
						data.DestinationNetwork = types.StringValue(value.String())
					} else {
						data.DestinationNetwork = types.StringNull()
					}
					if value := res.Get("value.destination.port"); value.Exists() && value.Value() != nil {
						data.DestinationPort = types.StringValue(value.String())
					} else {
						data.DestinationPort = types.StringNull()
					}
					if value := res.Get("value.destination.vlan"); value.Exists() && value.Value() != nil {
						data.DestinationVlan = types.Int64Value(value.Int())
					} else {
						data.DestinationVlan = types.Int64Null()
					}
					if value := res.Get("value.source.cidr"); value.Exists() && value.Value() != nil {
						data.SourceCidr = types.StringValue(value.String())
					} else {
						data.SourceCidr = types.StringNull()
					}
					if value := res.Get("value.source.host"); value.Exists() && value.Value() != nil {
						data.SourceHost = types.Int64Value(value.Int())
					} else {
						data.SourceHost = types.Int64Null()
					}
					if value := res.Get("value.source.network"); value.Exists() && value.Value() != nil {
						data.SourceNetwork = types.StringValue(value.String())
					} else {
						data.SourceNetwork = types.StringNull()
					}
					if value := res.Get("value.source.port"); value.Exists() && value.Value() != nil {
						data.SourcePort = types.StringValue(value.String())
					} else {
						data.SourcePort = types.StringNull()
					}
					if value := res.Get("value.source.vlan"); value.Exists() && value.Value() != nil {
						data.SourceVlan = types.Int64Value(value.Int())
					} else {
						data.SourceVlan = types.Int64Null()
					}
					(*parent).TrafficFilters = append((*parent).TrafficFilters, data)
					return true
				})
			}
			(*parent).VpnTrafficUplinkPreferences = append((*parent).VpnTrafficUplinkPreferences, data)
			return true
		})
	}
	if value := res.Get("wanTrafficUplinkPreferences"); value.Exists() && value.Value() != nil {
		data.WanTrafficUplinkPreferences = make([]DataSourceApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences{}
			if value := res.Get("preferredUplink"); value.Exists() && value.Value() != nil {
				data.PreferredUplink = types.StringValue(value.String())
			} else {
				data.PreferredUplink = types.StringNull()
			}
			if value := res.Get("trafficFilters"); value.Exists() && value.Value() != nil {
				data.TrafficFilters = make([]DataSourceApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DataSourceApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters{}
					if value := res.Get("type"); value.Exists() && value.Value() != nil {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					if value := res.Get("value.protocol"); value.Exists() && value.Value() != nil {
						data.Protocol = types.StringValue(value.String())
					} else {
						data.Protocol = types.StringNull()
					}
					if value := res.Get("value.destination.cidr"); value.Exists() && value.Value() != nil {
						data.DestinationCidr = types.StringValue(value.String())
					} else {
						data.DestinationCidr = types.StringNull()
					}
					if value := res.Get("value.destination.port"); value.Exists() && value.Value() != nil {
						data.DestinationPort = types.StringValue(value.String())
					} else {
						data.DestinationPort = types.StringNull()
					}
					if value := res.Get("value.source.cidr"); value.Exists() && value.Value() != nil {
						data.SourceCidr = types.StringValue(value.String())
					} else {
						data.SourceCidr = types.StringNull()
					}
					if value := res.Get("value.source.host"); value.Exists() && value.Value() != nil {
						data.SourceHost = types.Int64Value(value.Int())
					} else {
						data.SourceHost = types.Int64Null()
					}
					if value := res.Get("value.source.port"); value.Exists() && value.Value() != nil {
						data.SourcePort = types.StringValue(value.String())
					} else {
						data.SourcePort = types.StringNull()
					}
					if value := res.Get("value.source.vlan"); value.Exists() && value.Value() != nil {
						data.SourceVlan = types.Int64Value(value.Int())
					} else {
						data.SourceVlan = types.Int64Null()
					}
					(*parent).TrafficFilters = append((*parent).TrafficFilters, data)
					return true
				})
			}
			(*parent).WanTrafficUplinkPreferences = append((*parent).WanTrafficUplinkPreferences, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
