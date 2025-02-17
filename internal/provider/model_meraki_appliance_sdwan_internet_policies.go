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

type ApplianceSDWANInternetPolicies struct {
	Id                          types.String                                                `tfsdk:"id"`
	NetworkId                   types.String                                                `tfsdk:"network_id"`
	WanTrafficUplinkPreferences []ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferences `tfsdk:"wan_traffic_uplink_preferences"`
}

type ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferences struct {
	FailOverCriterion           types.String                                                              `tfsdk:"fail_over_criterion"`
	PreferredUplink             types.String                                                              `tfsdk:"preferred_uplink"`
	BuiltinPerformanceClassName types.String                                                              `tfsdk:"builtin_performance_class_name"`
	CustomPerformanceClassId    types.String                                                              `tfsdk:"custom_performance_class_id"`
	PerformanceClassType        types.String                                                              `tfsdk:"performance_class_type"`
	TrafficFilters              []ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters `tfsdk:"traffic_filters"`
}

type ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters struct {
	Type                    types.String                                                                                     `tfsdk:"type"`
	Protocol                types.String                                                                                     `tfsdk:"protocol"`
	DestinationCidr         types.String                                                                                     `tfsdk:"destination_cidr"`
	DestinationPort         types.String                                                                                     `tfsdk:"destination_port"`
	DestinationApplications []ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersDestinationApplications `tfsdk:"destination_applications"`
	SourceCidr              types.String                                                                                     `tfsdk:"source_cidr"`
	SourceHost              types.Int64                                                                                      `tfsdk:"source_host"`
	SourcePort              types.String                                                                                     `tfsdk:"source_port"`
	SourceVlan              types.Int64                                                                                      `tfsdk:"source_vlan"`
}

type ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersDestinationApplications struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceSDWANInternetPolicies) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/sdwan/internetPolicies", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceSDWANInternetPolicies) toBody(ctx context.Context, state ApplianceSDWANInternetPolicies) string {
	body := ""
	{
		body, _ = sjson.Set(body, "wanTrafficUplinkPreferences", []interface{}{})
		for _, item := range data.WanTrafficUplinkPreferences {
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
					if !childItem.Protocol.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.protocol", childItem.Protocol.ValueString())
					}
					if !childItem.DestinationCidr.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.destination.cidr", childItem.DestinationCidr.ValueString())
					}
					if !childItem.DestinationPort.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.destination.port", childItem.DestinationPort.ValueString())
					}
					if len(childItem.DestinationApplications) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "value.destination.applications", []interface{}{})
						for _, childChildItem := range childItem.DestinationApplications {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "id", childChildItem.Id.ValueString())
							}
							if !childChildItem.Name.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "name", childChildItem.Name.ValueString())
							}
							if !childChildItem.Type.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "type", childChildItem.Type.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "value.destination.applications.-1", itemChildChildBody)
						}
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceSDWANInternetPolicies) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("wanTrafficUplinkPreferences"); value.Exists() && value.Value() != nil {
		data.WanTrafficUplinkPreferences = make([]ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferences, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferences{}
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
				data.TrafficFilters = make([]ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferencesTrafficFilters{}
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
					if value := res.Get("value.destination.applications"); value.Exists() && value.Value() != nil {
						data.DestinationApplications = make([]ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersDestinationApplications, 0)
						value.ForEach(func(k, res gjson.Result) bool {
							parent := &data
							data := ApplianceSDWANInternetPoliciesWanTrafficUplinkPreferencesTrafficFiltersDestinationApplications{}
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
							if value := res.Get("type"); value.Exists() && value.Value() != nil {
								data.Type = types.StringValue(value.String())
							} else {
								data.Type = types.StringNull()
							}
							(*parent).DestinationApplications = append((*parent).DestinationApplications, data)
							return true
						})
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
func (data *ApplianceSDWANInternetPolicies) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
			for i := 0; i < len(data.DestinationApplications); i++ {
				keys := [...]string{"id", "name", "type"}
				keyValues := [...]string{data.DestinationApplications[i].Id.ValueString(), data.DestinationApplications[i].Name.ValueString(), data.DestinationApplications[i].Type.ValueString()}

				parent := &data
				data := (*parent).DestinationApplications[i]
				parentRes := &res
				var res gjson.Result

				parentRes.Get("value.destination.applications").ForEach(
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
					tflog.Debug(ctx, fmt.Sprintf("removing DestinationApplications[%d] = %+v",
						i,
						(*parent).DestinationApplications[i],
					))
					(*parent).DestinationApplications = slices.Delete((*parent).DestinationApplications, i, i+1)
					i--

					continue
				}
				if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
					data.Type = types.StringValue(value.String())
				} else {
					data.Type = types.StringNull()
				}
				(*parent).DestinationApplications[i] = data
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
func (data *ApplianceSDWANInternetPolicies) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceSDWANInternetPolicies) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
