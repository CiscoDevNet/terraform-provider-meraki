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

type ApplianceTrafficShapingUplinkSelection struct {
	Id                                  types.String                                                        `tfsdk:"id"`
	NetworkId                           types.String                                                        `tfsdk:"network_id"`
	ActiveActiveAutoVpnEnabled          types.Bool                                                          `tfsdk:"active_active_auto_vpn_enabled"`
	DefaultUplink                       types.String                                                        `tfsdk:"default_uplink"`
	LoadBalancingEnabled                types.Bool                                                          `tfsdk:"load_balancing_enabled"`
	FailoverAndFailbackImmediateEnabled types.Bool                                                          `tfsdk:"failover_and_failback_immediate_enabled"`
	VpnTrafficUplinkPreferences         []ApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences `tfsdk:"vpn_traffic_uplink_preferences"`
	WanTrafficUplinkPreferences         []ApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences `tfsdk:"wan_traffic_uplink_preferences"`
}

type ApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences struct {
	FailOverCriterion           types.String                                                                      `tfsdk:"fail_over_criterion"`
	PreferredUplink             types.String                                                                      `tfsdk:"preferred_uplink"`
	BuiltinPerformanceClassName types.String                                                                      `tfsdk:"builtin_performance_class_name"`
	CustomPerformanceClassId    types.String                                                                      `tfsdk:"custom_performance_class_id"`
	PerformanceClassType        types.String                                                                      `tfsdk:"performance_class_type"`
	TrafficFilters              []ApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters `tfsdk:"traffic_filters"`
}

type ApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences struct {
	PreferredUplink types.String                                                                      `tfsdk:"preferred_uplink"`
	TrafficFilters  []ApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters `tfsdk:"traffic_filters"`
}

type ApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters struct {
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

type ApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters struct {
	Type            types.String `tfsdk:"type"`
	Protocol        types.String `tfsdk:"protocol"`
	DestinationCidr types.String `tfsdk:"destination_cidr"`
	DestinationPort types.String `tfsdk:"destination_port"`
	SourceCidr      types.String `tfsdk:"source_cidr"`
	SourceHost      types.Int64  `tfsdk:"source_host"`
	SourcePort      types.String `tfsdk:"source_port"`
	SourceVlan      types.Int64  `tfsdk:"source_vlan"`
}

type ApplianceTrafficShapingUplinkSelectionIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceTrafficShapingUplinkSelection) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/trafficShaping/uplinkSelection", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceTrafficShapingUplinkSelection) toBody(ctx context.Context, state ApplianceTrafficShapingUplinkSelection) string {
	body := ""
	if !data.ActiveActiveAutoVpnEnabled.IsNull() {
		body, _ = sjson.Set(body, "activeActiveAutoVpnEnabled", data.ActiveActiveAutoVpnEnabled.ValueBool())
	}
	if !data.DefaultUplink.IsNull() {
		body, _ = sjson.Set(body, "defaultUplink", data.DefaultUplink.ValueString())
	}
	if !data.LoadBalancingEnabled.IsNull() {
		body, _ = sjson.Set(body, "loadBalancingEnabled", data.LoadBalancingEnabled.ValueBool())
	}
	if !data.FailoverAndFailbackImmediateEnabled.IsNull() {
		body, _ = sjson.Set(body, "failoverAndFailback.immediate.enabled", data.FailoverAndFailbackImmediateEnabled.ValueBool())
	}
	if len(data.VpnTrafficUplinkPreferences) > 0 {
		body, _ = sjson.Set(body, "vpnTrafficUplinkPreferences", []interface{}{})
		for _, item := range data.VpnTrafficUplinkPreferences {
			itemBody := ""
			if !item.FailOverCriterion.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "failOverCriterion", item.FailOverCriterion.ValueString())
			}
			if !item.PreferredUplink.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "preferredUplink", item.PreferredUplink.ValueString())
			}
			if !item.BuiltinPerformanceClassName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "performanceClass.builtinPerformanceClassName", item.BuiltinPerformanceClassName.ValueString())
			}
			if !item.CustomPerformanceClassId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "performanceClass.customPerformanceClassId", item.CustomPerformanceClassId.ValueString())
			}
			if !item.PerformanceClassType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "performanceClass.type", item.PerformanceClassType.ValueString())
			}
			{
				itemBody, _ = sjson.Set(itemBody, "trafficFilters", []interface{}{})
				for _, childItem := range item.TrafficFilters {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.id", childItem.Id.ValueString())
					}
					if !childItem.Protocol.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.protocol", childItem.Protocol.ValueString())
					}
					if !childItem.DestinationCidr.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.destination.cidr", childItem.DestinationCidr.ValueString())
					}
					if !childItem.DestinationFqdn.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.destination.fqdn", childItem.DestinationFqdn.ValueString())
					}
					if !childItem.DestinationHost.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.destination.host", childItem.DestinationHost.ValueInt64())
					}
					if !childItem.DestinationNetwork.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.destination.network", childItem.DestinationNetwork.ValueString())
					}
					if !childItem.DestinationPort.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.destination.port", childItem.DestinationPort.ValueString())
					}
					if !childItem.DestinationVlan.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.destination.vlan", childItem.DestinationVlan.ValueInt64())
					}
					if !childItem.SourceCidr.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.source.cidr", childItem.SourceCidr.ValueString())
					}
					if !childItem.SourceHost.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.source.host", childItem.SourceHost.ValueInt64())
					}
					if !childItem.SourceNetwork.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.source.network", childItem.SourceNetwork.ValueString())
					}
					if !childItem.SourcePort.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.source.port", childItem.SourcePort.ValueString())
					}
					if !childItem.SourceVlan.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.source.vlan", childItem.SourceVlan.ValueInt64())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "trafficFilters.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "vpnTrafficUplinkPreferences.-1", itemBody)
		}
	}
	if len(data.WanTrafficUplinkPreferences) > 0 {
		body, _ = sjson.Set(body, "wanTrafficUplinkPreferences", []interface{}{})
		for _, item := range data.WanTrafficUplinkPreferences {
			itemBody := ""
			if !item.PreferredUplink.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "preferredUplink", item.PreferredUplink.ValueString())
			}
			{
				itemBody, _ = sjson.Set(itemBody, "trafficFilters", []interface{}{})
				for _, childItem := range item.TrafficFilters {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Protocol.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.protocol", childItem.Protocol.ValueString())
					}
					if !childItem.DestinationCidr.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.destination.cidr", childItem.DestinationCidr.ValueString())
					}
					if !childItem.DestinationPort.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.destination.port", childItem.DestinationPort.ValueString())
					}
					if !childItem.SourceCidr.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.source.cidr", childItem.SourceCidr.ValueString())
					}
					if !childItem.SourceHost.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.source.host", childItem.SourceHost.ValueInt64())
					}
					if !childItem.SourcePort.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.source.port", childItem.SourcePort.ValueString())
					}
					if !childItem.SourceVlan.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.source.vlan", childItem.SourceVlan.ValueInt64())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "trafficFilters.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "wanTrafficUplinkPreferences.-1", itemBody)
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
func (data ApplianceTrafficShapingUplinkSelection) toBodyPreservingNulls(ctx context.Context, res meraki.Res) string {
	body := ""
	if value := res.Get("activeActiveAutoVpnEnabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "activeActiveAutoVpnEnabled", "null")
		} else {
			body, _ = sjson.Set(body, "activeActiveAutoVpnEnabled", value.Bool())
		}
	}
	if value := res.Get("defaultUplink"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "defaultUplink", "null")
		} else {
			body, _ = sjson.Set(body, "defaultUplink", value.String())
		}
	}
	if value := res.Get("loadBalancingEnabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "loadBalancingEnabled", "null")
		} else {
			body, _ = sjson.Set(body, "loadBalancingEnabled", value.Bool())
		}
	}
	if value := res.Get("failoverAndFailback.immediate.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "failoverAndFailback.immediate.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "failoverAndFailback.immediate.enabled", value.Bool())
		}
	}
	if value := res.Get("vpnTrafficUplinkPreferences"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "vpnTrafficUplinkPreferences", "null")
		} else {
			body, _ = sjson.Set(body, "vpnTrafficUplinkPreferences", []interface{}{})
			parent := &body
			value.ForEach(func(k, res gjson.Result) bool {
				body := ""
				if value := res.Get("failOverCriterion"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "failOverCriterion", "null")
					} else {
						body, _ = sjson.Set(body, "failOverCriterion", value.String())
					}
				}
				if value := res.Get("preferredUplink"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "preferredUplink", "null")
					} else {
						body, _ = sjson.Set(body, "preferredUplink", value.String())
					}
				}
				if value := res.Get("performanceClass.builtinPerformanceClassName"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "performanceClass.builtinPerformanceClassName", "null")
					} else {
						body, _ = sjson.Set(body, "performanceClass.builtinPerformanceClassName", value.String())
					}
				}
				if value := res.Get("performanceClass.customPerformanceClassId"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "performanceClass.customPerformanceClassId", "null")
					} else {
						body, _ = sjson.Set(body, "performanceClass.customPerformanceClassId", value.String())
					}
				}
				if value := res.Get("performanceClass.type"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "performanceClass.type", "null")
					} else {
						body, _ = sjson.Set(body, "performanceClass.type", value.String())
					}
				}
				if value := res.Get("trafficFilters"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "trafficFilters", "null")
					} else {
						body, _ = sjson.Set(body, "trafficFilters", []interface{}{})
						parent := &body
						value.ForEach(func(k, res gjson.Result) bool {
							body := ""
							if value := res.Get("type"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "type", "null")
								} else {
									body, _ = sjson.Set(body, "type", value.String())
								}
							}
							if value := res.Get("value.id"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.id", "null")
								} else {
									body, _ = sjson.Set(body, "value.id", value.String())
								}
							}
							if value := res.Get("value.protocol"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.protocol", "null")
								} else {
									body, _ = sjson.Set(body, "value.protocol", value.String())
								}
							}
							if value := res.Get("value.destination.cidr"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.destination.cidr", "null")
								} else {
									body, _ = sjson.Set(body, "value.destination.cidr", value.String())
								}
							}
							if value := res.Get("value.destination.fqdn"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.destination.fqdn", "null")
								} else {
									body, _ = sjson.Set(body, "value.destination.fqdn", value.String())
								}
							}
							if value := res.Get("value.destination.host"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.destination.host", "null")
								} else {
									body, _ = sjson.Set(body, "value.destination.host", value.Int())
								}
							}
							if value := res.Get("value.destination.network"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.destination.network", "null")
								} else {
									body, _ = sjson.Set(body, "value.destination.network", value.String())
								}
							}
							if value := res.Get("value.destination.port"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.destination.port", "null")
								} else {
									body, _ = sjson.Set(body, "value.destination.port", value.String())
								}
							}
							if value := res.Get("value.destination.vlan"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.destination.vlan", "null")
								} else {
									body, _ = sjson.Set(body, "value.destination.vlan", value.Int())
								}
							}
							if value := res.Get("value.source.cidr"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.source.cidr", "null")
								} else {
									body, _ = sjson.Set(body, "value.source.cidr", value.String())
								}
							}
							if value := res.Get("value.source.host"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.source.host", "null")
								} else {
									body, _ = sjson.Set(body, "value.source.host", value.Int())
								}
							}
							if value := res.Get("value.source.network"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.source.network", "null")
								} else {
									body, _ = sjson.Set(body, "value.source.network", value.String())
								}
							}
							if value := res.Get("value.source.port"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.source.port", "null")
								} else {
									body, _ = sjson.Set(body, "value.source.port", value.String())
								}
							}
							if value := res.Get("value.source.vlan"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.source.vlan", "null")
								} else {
									body, _ = sjson.Set(body, "value.source.vlan", value.Int())
								}
							}
							*parent, _ = sjson.SetRaw(*parent, "trafficFilters.-1", body)
							return true
						})
					}
				}
				*parent, _ = sjson.SetRaw(*parent, "vpnTrafficUplinkPreferences.-1", body)
				return true
			})
		}
	}
	if value := res.Get("wanTrafficUplinkPreferences"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "wanTrafficUplinkPreferences", "null")
		} else {
			body, _ = sjson.Set(body, "wanTrafficUplinkPreferences", []interface{}{})
			parent := &body
			value.ForEach(func(k, res gjson.Result) bool {
				body := ""
				if value := res.Get("preferredUplink"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "preferredUplink", "null")
					} else {
						body, _ = sjson.Set(body, "preferredUplink", value.String())
					}
				}
				if value := res.Get("trafficFilters"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "trafficFilters", "null")
					} else {
						body, _ = sjson.Set(body, "trafficFilters", []interface{}{})
						parent := &body
						value.ForEach(func(k, res gjson.Result) bool {
							body := ""
							if value := res.Get("type"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "type", "null")
								} else {
									body, _ = sjson.Set(body, "type", value.String())
								}
							}
							if value := res.Get("value.protocol"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.protocol", "null")
								} else {
									body, _ = sjson.Set(body, "value.protocol", value.String())
								}
							}
							if value := res.Get("value.destination.cidr"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.destination.cidr", "null")
								} else {
									body, _ = sjson.Set(body, "value.destination.cidr", value.String())
								}
							}
							if value := res.Get("value.destination.port"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.destination.port", "null")
								} else {
									body, _ = sjson.Set(body, "value.destination.port", value.String())
								}
							}
							if value := res.Get("value.source.cidr"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.source.cidr", "null")
								} else {
									body, _ = sjson.Set(body, "value.source.cidr", value.String())
								}
							}
							if value := res.Get("value.source.host"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.source.host", "null")
								} else {
									body, _ = sjson.Set(body, "value.source.host", value.Int())
								}
							}
							if value := res.Get("value.source.port"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.source.port", "null")
								} else {
									body, _ = sjson.Set(body, "value.source.port", value.String())
								}
							}
							if value := res.Get("value.source.vlan"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "value.source.vlan", "null")
								} else {
									body, _ = sjson.Set(body, "value.source.vlan", value.Int())
								}
							}
							*parent, _ = sjson.SetRaw(*parent, "trafficFilters.-1", body)
							return true
						})
					}
				}
				*parent, _ = sjson.SetRaw(*parent, "wanTrafficUplinkPreferences.-1", body)
				return true
			})
		}
	}
	return body
}

// End of section. //template:end toBodyPreservingNulls

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceTrafficShapingUplinkSelection) fromBody(ctx context.Context, res meraki.Res) {
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
		data.VpnTrafficUplinkPreferences = make([]ApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferences{}
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
				data.TrafficFilters = make([]ApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ApplianceTrafficShapingUplinkSelectionVpnTrafficUplinkPreferencesTrafficFilters{}
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
		data.WanTrafficUplinkPreferences = make([]ApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferences{}
			if value := res.Get("preferredUplink"); value.Exists() && value.Value() != nil {
				data.PreferredUplink = types.StringValue(value.String())
			} else {
				data.PreferredUplink = types.StringNull()
			}
			if value := res.Get("trafficFilters"); value.Exists() && value.Value() != nil {
				data.TrafficFilters = make([]ApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ApplianceTrafficShapingUplinkSelectionWanTrafficUplinkPreferencesTrafficFilters{}
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceTrafficShapingUplinkSelection) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("activeActiveAutoVpnEnabled"); value.Exists() && !data.ActiveActiveAutoVpnEnabled.IsNull() {
		data.ActiveActiveAutoVpnEnabled = types.BoolValue(value.Bool())
	} else {
		data.ActiveActiveAutoVpnEnabled = types.BoolNull()
	}
	if value := res.Get("defaultUplink"); value.Exists() && !data.DefaultUplink.IsNull() {
		data.DefaultUplink = types.StringValue(value.String())
	} else {
		data.DefaultUplink = types.StringNull()
	}
	if value := res.Get("loadBalancingEnabled"); value.Exists() && !data.LoadBalancingEnabled.IsNull() {
		data.LoadBalancingEnabled = types.BoolValue(value.Bool())
	} else {
		data.LoadBalancingEnabled = types.BoolNull()
	}
	if value := res.Get("failoverAndFailback.immediate.enabled"); value.Exists() && !data.FailoverAndFailbackImmediateEnabled.IsNull() {
		data.FailoverAndFailbackImmediateEnabled = types.BoolValue(value.Bool())
	} else {
		data.FailoverAndFailbackImmediateEnabled = types.BoolNull()
	}
	{
		l := len(res.Get("vpnTrafficUplinkPreferences").Array())
		tflog.Debug(ctx, fmt.Sprintf("vpnTrafficUplinkPreferences array resizing from %d to %d", len(data.VpnTrafficUplinkPreferences), l))
		if len(data.VpnTrafficUplinkPreferences) > l {
			data.VpnTrafficUplinkPreferences = data.VpnTrafficUplinkPreferences[:l]
		}
	}
	for i := range data.VpnTrafficUplinkPreferences {
		parent := &data
		data := (*parent).VpnTrafficUplinkPreferences[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("vpnTrafficUplinkPreferences.%d", i))
		if value := res.Get("failOverCriterion"); value.Exists() && !data.FailOverCriterion.IsNull() {
			data.FailOverCriterion = types.StringValue(value.String())
		} else {
			data.FailOverCriterion = types.StringNull()
		}
		if value := res.Get("preferredUplink"); value.Exists() && !data.PreferredUplink.IsNull() {
			data.PreferredUplink = types.StringValue(value.String())
		} else {
			data.PreferredUplink = types.StringNull()
		}
		if value := res.Get("performanceClass.builtinPerformanceClassName"); value.Exists() && !data.BuiltinPerformanceClassName.IsNull() {
			data.BuiltinPerformanceClassName = types.StringValue(value.String())
		} else {
			data.BuiltinPerformanceClassName = types.StringNull()
		}
		if value := res.Get("performanceClass.customPerformanceClassId"); value.Exists() && !data.CustomPerformanceClassId.IsNull() {
			data.CustomPerformanceClassId = types.StringValue(value.String())
		} else {
			data.CustomPerformanceClassId = types.StringNull()
		}
		if value := res.Get("performanceClass.type"); value.Exists() && !data.PerformanceClassType.IsNull() {
			data.PerformanceClassType = types.StringValue(value.String())
		} else {
			data.PerformanceClassType = types.StringNull()
		}
		for i := 0; i < len(data.TrafficFilters); i++ {
			keys := [...]string{"type"}
			keyValues := [...]string{data.TrafficFilters[i].Type.ValueString()}

			parent := &data
			data := (*parent).TrafficFilters[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("trafficFilters").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing TrafficFilters[%d] = %+v",
					i,
					(*parent).TrafficFilters[i],
				))
				(*parent).TrafficFilters = slices.Delete((*parent).TrafficFilters, i, i+1)
				i--

				continue
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("value.id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("value.protocol"); value.Exists() && !data.Protocol.IsNull() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("value.destination.cidr"); value.Exists() && !data.DestinationCidr.IsNull() {
				data.DestinationCidr = types.StringValue(value.String())
			} else {
				data.DestinationCidr = types.StringNull()
			}
			if value := res.Get("value.destination.fqdn"); value.Exists() && !data.DestinationFqdn.IsNull() {
				data.DestinationFqdn = types.StringValue(value.String())
			} else {
				data.DestinationFqdn = types.StringNull()
			}
			if value := res.Get("value.destination.host"); value.Exists() && !data.DestinationHost.IsNull() {
				data.DestinationHost = types.Int64Value(value.Int())
			} else {
				data.DestinationHost = types.Int64Null()
			}
			if value := res.Get("value.destination.network"); value.Exists() && !data.DestinationNetwork.IsNull() {
				data.DestinationNetwork = types.StringValue(value.String())
			} else {
				data.DestinationNetwork = types.StringNull()
			}
			if value := res.Get("value.destination.port"); value.Exists() && !data.DestinationPort.IsNull() {
				data.DestinationPort = types.StringValue(value.String())
			} else {
				data.DestinationPort = types.StringNull()
			}
			if value := res.Get("value.destination.vlan"); value.Exists() && !data.DestinationVlan.IsNull() {
				data.DestinationVlan = types.Int64Value(value.Int())
			} else {
				data.DestinationVlan = types.Int64Null()
			}
			if value := res.Get("value.source.cidr"); value.Exists() && !data.SourceCidr.IsNull() {
				data.SourceCidr = types.StringValue(value.String())
			} else {
				data.SourceCidr = types.StringNull()
			}
			if value := res.Get("value.source.host"); value.Exists() && !data.SourceHost.IsNull() {
				data.SourceHost = types.Int64Value(value.Int())
			} else {
				data.SourceHost = types.Int64Null()
			}
			if value := res.Get("value.source.network"); value.Exists() && !data.SourceNetwork.IsNull() {
				data.SourceNetwork = types.StringValue(value.String())
			} else {
				data.SourceNetwork = types.StringNull()
			}
			if value := res.Get("value.source.port"); value.Exists() && !data.SourcePort.IsNull() {
				data.SourcePort = types.StringValue(value.String())
			} else {
				data.SourcePort = types.StringNull()
			}
			if value := res.Get("value.source.vlan"); value.Exists() && !data.SourceVlan.IsNull() {
				data.SourceVlan = types.Int64Value(value.Int())
			} else {
				data.SourceVlan = types.Int64Null()
			}
			(*parent).TrafficFilters[i] = data
		}
		(*parent).VpnTrafficUplinkPreferences[i] = data
	}
	{
		l := len(res.Get("wanTrafficUplinkPreferences").Array())
		tflog.Debug(ctx, fmt.Sprintf("wanTrafficUplinkPreferences array resizing from %d to %d", len(data.WanTrafficUplinkPreferences), l))
		if len(data.WanTrafficUplinkPreferences) > l {
			data.WanTrafficUplinkPreferences = data.WanTrafficUplinkPreferences[:l]
		}
	}
	for i := range data.WanTrafficUplinkPreferences {
		parent := &data
		data := (*parent).WanTrafficUplinkPreferences[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("wanTrafficUplinkPreferences.%d", i))
		if value := res.Get("preferredUplink"); value.Exists() && !data.PreferredUplink.IsNull() {
			data.PreferredUplink = types.StringValue(value.String())
		} else {
			data.PreferredUplink = types.StringNull()
		}
		for i := 0; i < len(data.TrafficFilters); i++ {
			keys := [...]string{"type"}
			keyValues := [...]string{data.TrafficFilters[i].Type.ValueString()}

			parent := &data
			data := (*parent).TrafficFilters[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("trafficFilters").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing TrafficFilters[%d] = %+v",
					i,
					(*parent).TrafficFilters[i],
				))
				(*parent).TrafficFilters = slices.Delete((*parent).TrafficFilters, i, i+1)
				i--

				continue
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("value.protocol"); value.Exists() && !data.Protocol.IsNull() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("value.destination.cidr"); value.Exists() && !data.DestinationCidr.IsNull() {
				data.DestinationCidr = types.StringValue(value.String())
			} else {
				data.DestinationCidr = types.StringNull()
			}
			if value := res.Get("value.destination.port"); value.Exists() && !data.DestinationPort.IsNull() {
				data.DestinationPort = types.StringValue(value.String())
			} else {
				data.DestinationPort = types.StringNull()
			}
			if value := res.Get("value.source.cidr"); value.Exists() && !data.SourceCidr.IsNull() {
				data.SourceCidr = types.StringValue(value.String())
			} else {
				data.SourceCidr = types.StringNull()
			}
			if value := res.Get("value.source.host"); value.Exists() && !data.SourceHost.IsNull() {
				data.SourceHost = types.Int64Value(value.Int())
			} else {
				data.SourceHost = types.Int64Null()
			}
			if value := res.Get("value.source.port"); value.Exists() && !data.SourcePort.IsNull() {
				data.SourcePort = types.StringValue(value.String())
			} else {
				data.SourcePort = types.StringNull()
			}
			if value := res.Get("value.source.vlan"); value.Exists() && !data.SourceVlan.IsNull() {
				data.SourceVlan = types.Int64Value(value.Int())
			} else {
				data.SourceVlan = types.Int64Null()
			}
			(*parent).TrafficFilters[i] = data
		}
		(*parent).WanTrafficUplinkPreferences[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceTrafficShapingUplinkSelection) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceTrafficShapingUplinkSelectionIdentity) toIdentity(ctx context.Context, plan *ApplianceTrafficShapingUplinkSelection) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceTrafficShapingUplinkSelection) fromIdentity(ctx context.Context, identity *ApplianceTrafficShapingUplinkSelectionIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceTrafficShapingUplinkSelection) toDestroyBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "vpnTrafficUplinkPreferences", []interface{}{})
	return body
}

// End of section. //template:end toDestroyBody
