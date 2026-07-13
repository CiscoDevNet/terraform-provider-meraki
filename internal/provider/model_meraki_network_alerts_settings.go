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

type NetworkAlertsSettings struct {
	Id                               types.String                  `tfsdk:"id"`
	NetworkId                        types.String                  `tfsdk:"network_id"`
	DefaultDestinationsAllAdmins     types.Bool                    `tfsdk:"default_destinations_all_admins"`
	DefaultDestinationsSnmp          types.Bool                    `tfsdk:"default_destinations_snmp"`
	DefaultDestinationsEmails        types.List                    `tfsdk:"default_destinations_emails"`
	DefaultDestinationsHttpServerIds types.List                    `tfsdk:"default_destinations_http_server_ids"`
	MutingByPortSchedulesEnabled     types.Bool                    `tfsdk:"muting_by_port_schedules_enabled"`
	Alerts                           []NetworkAlertsSettingsAlerts `tfsdk:"alerts"`
}

type NetworkAlertsSettingsAlerts struct {
	Enabled                        types.Bool                                     `tfsdk:"enabled"`
	Type                           types.String                                   `tfsdk:"type"`
	AlertDestinationsAllAdmins     types.Bool                                     `tfsdk:"alert_destinations_all_admins"`
	AlertDestinationsSnmp          types.Bool                                     `tfsdk:"alert_destinations_snmp"`
	AlertDestinationsEmails        types.List                                     `tfsdk:"alert_destinations_emails"`
	AlertDestinationsHttpServerIds types.List                                     `tfsdk:"alert_destinations_http_server_ids"`
	AlertDestinationsSmsNumbers    types.List                                     `tfsdk:"alert_destinations_sms_numbers"`
	FiltersFailureType             types.String                                   `tfsdk:"filters_failure_type"`
	FiltersLookbackWindow          types.Int64                                    `tfsdk:"filters_lookback_window"`
	FiltersMinDuration             types.Int64                                    `tfsdk:"filters_min_duration"`
	FiltersName                    types.String                                   `tfsdk:"filters_name"`
	FiltersPeriod                  types.Int64                                    `tfsdk:"filters_period"`
	FiltersPriority                types.String                                   `tfsdk:"filters_priority"`
	FiltersRegex                   types.String                                   `tfsdk:"filters_regex"`
	FiltersSelector                types.String                                   `tfsdk:"filters_selector"`
	FiltersSsidNum                 types.Int64                                    `tfsdk:"filters_ssid_num"`
	FiltersTag                     types.String                                   `tfsdk:"filters_tag"`
	FiltersThreshold               types.Int64                                    `tfsdk:"filters_threshold"`
	FiltersTimeout                 types.Int64                                    `tfsdk:"filters_timeout"`
	FiltersConditions              []NetworkAlertsSettingsAlertsFiltersConditions `tfsdk:"filters_conditions"`
	FiltersSerials                 types.Set                                      `tfsdk:"filters_serials"`
}

type NetworkAlertsSettingsAlertsFiltersConditions struct {
	Direction types.String  `tfsdk:"direction"`
	Duration  types.Int64   `tfsdk:"duration"`
	Threshold types.Float64 `tfsdk:"threshold"`
	Type      types.String  `tfsdk:"type"`
	Unit      types.String  `tfsdk:"unit"`
}

type NetworkAlertsSettingsIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkAlertsSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/alerts/settings", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkAlertsSettings) toBody(ctx context.Context, state NetworkAlertsSettings) string {
	body := ""
	if !data.DefaultDestinationsAllAdmins.IsNull() {
		body, _ = sjson.Set(body, "defaultDestinations.allAdmins", data.DefaultDestinationsAllAdmins.ValueBool())
	}
	if !data.DefaultDestinationsSnmp.IsNull() {
		body, _ = sjson.Set(body, "defaultDestinations.snmp", data.DefaultDestinationsSnmp.ValueBool())
	}
	if !data.DefaultDestinationsEmails.IsNull() {
		var values []string
		data.DefaultDestinationsEmails.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "defaultDestinations.emails", values)
	}
	if !data.DefaultDestinationsHttpServerIds.IsNull() {
		var values []string
		data.DefaultDestinationsHttpServerIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "defaultDestinations.httpServerIds", values)
	}
	if !data.MutingByPortSchedulesEnabled.IsNull() {
		body, _ = sjson.Set(body, "muting.byPortSchedules.enabled", data.MutingByPortSchedulesEnabled.ValueBool())
	}
	if len(data.Alerts) > 0 {
		body, _ = sjson.Set(body, "alerts", []interface{}{})
		for _, item := range data.Alerts {
			itemBody := ""
			if !item.Enabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enabled", item.Enabled.ValueBool())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.AlertDestinationsAllAdmins.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "alertDestinations.allAdmins", item.AlertDestinationsAllAdmins.ValueBool())
			}
			if !item.AlertDestinationsSnmp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "alertDestinations.snmp", item.AlertDestinationsSnmp.ValueBool())
			}
			if !item.AlertDestinationsEmails.IsNull() {
				var values []string
				item.AlertDestinationsEmails.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "alertDestinations.emails", values)
			}
			if !item.AlertDestinationsHttpServerIds.IsNull() {
				var values []string
				item.AlertDestinationsHttpServerIds.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "alertDestinations.httpServerIds", values)
			}
			if !item.AlertDestinationsSmsNumbers.IsNull() {
				var values []string
				item.AlertDestinationsSmsNumbers.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "alertDestinations.smsNumbers", values)
			}
			if !item.FiltersFailureType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.failureType", item.FiltersFailureType.ValueString())
			}
			if !item.FiltersLookbackWindow.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.lookbackWindow", item.FiltersLookbackWindow.ValueInt64())
			}
			if !item.FiltersMinDuration.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.minDuration", item.FiltersMinDuration.ValueInt64())
			}
			if !item.FiltersName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.name", item.FiltersName.ValueString())
			}
			if !item.FiltersPeriod.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.period", item.FiltersPeriod.ValueInt64())
			}
			if !item.FiltersPriority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.priority", item.FiltersPriority.ValueString())
			}
			if !item.FiltersRegex.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.regex", item.FiltersRegex.ValueString())
			}
			if !item.FiltersSelector.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.selector", item.FiltersSelector.ValueString())
			}
			if !item.FiltersSsidNum.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.ssidNum", item.FiltersSsidNum.ValueInt64())
			}
			if !item.FiltersTag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.tag", item.FiltersTag.ValueString())
			}
			if !item.FiltersThreshold.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.threshold", item.FiltersThreshold.ValueInt64())
			}
			if !item.FiltersTimeout.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filters.timeout", item.FiltersTimeout.ValueInt64())
			}
			if len(item.FiltersConditions) > 0 {
				itemBody, _ = sjson.Set(itemBody, "filters.conditions", []interface{}{})
				for _, childItem := range item.FiltersConditions {
					itemChildBody := ""
					if !childItem.Direction.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "direction", childItem.Direction.ValueString())
					}
					if !childItem.Duration.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "duration", childItem.Duration.ValueInt64())
					}
					if !childItem.Threshold.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "threshold", childItem.Threshold.ValueFloat64())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Unit.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "unit", childItem.Unit.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "filters.conditions.-1", itemChildBody)
				}
			}
			if !item.FiltersSerials.IsNull() {
				var values []string
				item.FiltersSerials.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "filters.serials", values)
			}
			body, _ = sjson.SetRaw(body, "alerts.-1", itemBody)
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
func (data NetworkAlertsSettings) toBodyPreservingNulls(ctx context.Context, res meraki.Res) string {
	body := ""
	if value := res.Get("defaultDestinations.allAdmins"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "defaultDestinations.allAdmins", "null")
		} else {
			body, _ = sjson.Set(body, "defaultDestinations.allAdmins", value.Bool())
		}
	}
	if value := res.Get("defaultDestinations.snmp"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "defaultDestinations.snmp", "null")
		} else {
			body, _ = sjson.Set(body, "defaultDestinations.snmp", value.Bool())
		}
	}
	if value := res.Get("defaultDestinations.emails"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "defaultDestinations.emails", "null")
		} else {
			var values []string
			for _, v := range value.Array() {
				values = append(values, v.String())
			}
			body, _ = sjson.Set(body, "defaultDestinations.emails", values)
		}
	}
	if value := res.Get("defaultDestinations.httpServerIds"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "defaultDestinations.httpServerIds", "null")
		} else {
			var values []string
			for _, v := range value.Array() {
				values = append(values, v.String())
			}
			body, _ = sjson.Set(body, "defaultDestinations.httpServerIds", values)
		}
	}
	if value := res.Get("muting.byPortSchedules.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "muting.byPortSchedules.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "muting.byPortSchedules.enabled", value.Bool())
		}
	}
	if value := res.Get("alerts"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "alerts", "null")
		} else {
			body, _ = sjson.Set(body, "alerts", []interface{}{})
			parent := &body
			value.ForEach(func(k, res gjson.Result) bool {
				body := ""
				if value := res.Get("enabled"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "enabled", "null")
					} else {
						body, _ = sjson.Set(body, "enabled", value.Bool())
					}
				}
				if value := res.Get("type"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "type", "null")
					} else {
						body, _ = sjson.Set(body, "type", value.String())
					}
				}
				if value := res.Get("alertDestinations.allAdmins"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "alertDestinations.allAdmins", "null")
					} else {
						body, _ = sjson.Set(body, "alertDestinations.allAdmins", value.Bool())
					}
				}
				if value := res.Get("alertDestinations.snmp"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "alertDestinations.snmp", "null")
					} else {
						body, _ = sjson.Set(body, "alertDestinations.snmp", value.Bool())
					}
				}
				if value := res.Get("alertDestinations.emails"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "alertDestinations.emails", "null")
					} else {
						var values []string
						for _, v := range value.Array() {
							values = append(values, v.String())
						}
						body, _ = sjson.Set(body, "alertDestinations.emails", values)
					}
				}
				if value := res.Get("alertDestinations.httpServerIds"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "alertDestinations.httpServerIds", "null")
					} else {
						var values []string
						for _, v := range value.Array() {
							values = append(values, v.String())
						}
						body, _ = sjson.Set(body, "alertDestinations.httpServerIds", values)
					}
				}
				if value := res.Get("alertDestinations.smsNumbers"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "alertDestinations.smsNumbers", "null")
					} else {
						var values []string
						for _, v := range value.Array() {
							values = append(values, v.String())
						}
						body, _ = sjson.Set(body, "alertDestinations.smsNumbers", values)
					}
				}
				if value := res.Get("filters.failureType"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.failureType", "null")
					} else {
						body, _ = sjson.Set(body, "filters.failureType", value.String())
					}
				}
				if value := res.Get("filters.lookbackWindow"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.lookbackWindow", "null")
					} else {
						body, _ = sjson.Set(body, "filters.lookbackWindow", value.Int())
					}
				}
				if value := res.Get("filters.minDuration"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.minDuration", "null")
					} else {
						body, _ = sjson.Set(body, "filters.minDuration", value.Int())
					}
				}
				if value := res.Get("filters.name"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.name", "null")
					} else {
						body, _ = sjson.Set(body, "filters.name", value.String())
					}
				}
				if value := res.Get("filters.period"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.period", "null")
					} else {
						body, _ = sjson.Set(body, "filters.period", value.Int())
					}
				}
				if value := res.Get("filters.priority"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.priority", "null")
					} else {
						body, _ = sjson.Set(body, "filters.priority", value.String())
					}
				}
				if value := res.Get("filters.regex"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.regex", "null")
					} else {
						body, _ = sjson.Set(body, "filters.regex", value.String())
					}
				}
				if value := res.Get("filters.selector"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.selector", "null")
					} else {
						body, _ = sjson.Set(body, "filters.selector", value.String())
					}
				}
				if value := res.Get("filters.ssidNum"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.ssidNum", "null")
					} else {
						body, _ = sjson.Set(body, "filters.ssidNum", value.Int())
					}
				}
				if value := res.Get("filters.tag"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.tag", "null")
					} else {
						body, _ = sjson.Set(body, "filters.tag", value.String())
					}
				}
				if value := res.Get("filters.threshold"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.threshold", "null")
					} else {
						body, _ = sjson.Set(body, "filters.threshold", value.Int())
					}
				}
				if value := res.Get("filters.timeout"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.timeout", "null")
					} else {
						body, _ = sjson.Set(body, "filters.timeout", value.Int())
					}
				}
				if value := res.Get("filters.conditions"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.conditions", "null")
					} else {
						body, _ = sjson.Set(body, "filters.conditions", []interface{}{})
						parent := &body
						value.ForEach(func(k, res gjson.Result) bool {
							body := ""
							if value := res.Get("direction"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "direction", "null")
								} else {
									body, _ = sjson.Set(body, "direction", value.String())
								}
							}
							if value := res.Get("duration"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "duration", "null")
								} else {
									body, _ = sjson.Set(body, "duration", value.Int())
								}
							}
							if value := res.Get("threshold"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "threshold", "null")
								} else {
									body, _ = sjson.Set(body, "threshold", value.Float())
								}
							}
							if value := res.Get("type"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "type", "null")
								} else {
									body, _ = sjson.Set(body, "type", value.String())
								}
							}
							if value := res.Get("unit"); value.Exists() {
								if value.Value() == nil {
									body, _ = sjson.SetRaw(body, "unit", "null")
								} else {
									body, _ = sjson.Set(body, "unit", value.String())
								}
							}
							*parent, _ = sjson.SetRaw(*parent, "filters.conditions.-1", body)
							return true
						})
					}
				}
				if value := res.Get("filters.serials"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "filters.serials", "null")
					} else {
						var values []string
						for _, v := range value.Array() {
							values = append(values, v.String())
						}
						body, _ = sjson.Set(body, "filters.serials", values)
					}
				}
				*parent, _ = sjson.SetRaw(*parent, "alerts.-1", body)
				return true
			})
		}
	}
	return body
}

// End of section. //template:end toBodyPreservingNulls

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkAlertsSettings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("defaultDestinations.allAdmins"); value.Exists() && value.Value() != nil {
		data.DefaultDestinationsAllAdmins = types.BoolValue(value.Bool())
	} else {
		data.DefaultDestinationsAllAdmins = types.BoolNull()
	}
	if value := res.Get("defaultDestinations.snmp"); value.Exists() && value.Value() != nil {
		data.DefaultDestinationsSnmp = types.BoolValue(value.Bool())
	} else {
		data.DefaultDestinationsSnmp = types.BoolNull()
	}
	if value := res.Get("defaultDestinations.emails"); value.Exists() && value.Value() != nil {
		data.DefaultDestinationsEmails = helpers.GetStringList(value.Array())
	} else {
		data.DefaultDestinationsEmails = types.ListNull(types.StringType)
	}
	if value := res.Get("defaultDestinations.httpServerIds"); value.Exists() && value.Value() != nil {
		data.DefaultDestinationsHttpServerIds = helpers.GetStringList(value.Array())
	} else {
		data.DefaultDestinationsHttpServerIds = types.ListNull(types.StringType)
	}
	if value := res.Get("muting.byPortSchedules.enabled"); value.Exists() && value.Value() != nil {
		data.MutingByPortSchedulesEnabled = types.BoolValue(value.Bool())
	} else {
		data.MutingByPortSchedulesEnabled = types.BoolNull()
	}
	if value := res.Get("alerts"); value.Exists() && value.Value() != nil {
		data.Alerts = make([]NetworkAlertsSettingsAlerts, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkAlertsSettingsAlerts{}
			if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
				data.Enabled = types.BoolValue(value.Bool())
			} else {
				data.Enabled = types.BoolNull()
			}
			if value := res.Get("type"); value.Exists() && value.Value() != nil {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("alertDestinations.allAdmins"); value.Exists() && value.Value() != nil {
				data.AlertDestinationsAllAdmins = types.BoolValue(value.Bool())
			} else {
				data.AlertDestinationsAllAdmins = types.BoolNull()
			}
			if value := res.Get("alertDestinations.snmp"); value.Exists() && value.Value() != nil {
				data.AlertDestinationsSnmp = types.BoolValue(value.Bool())
			} else {
				data.AlertDestinationsSnmp = types.BoolNull()
			}
			if value := res.Get("alertDestinations.emails"); value.Exists() && value.Value() != nil {
				data.AlertDestinationsEmails = helpers.GetStringList(value.Array())
			} else {
				data.AlertDestinationsEmails = types.ListNull(types.StringType)
			}
			if value := res.Get("alertDestinations.httpServerIds"); value.Exists() && value.Value() != nil {
				data.AlertDestinationsHttpServerIds = helpers.GetStringList(value.Array())
			} else {
				data.AlertDestinationsHttpServerIds = types.ListNull(types.StringType)
			}
			if value := res.Get("alertDestinations.smsNumbers"); value.Exists() && value.Value() != nil {
				data.AlertDestinationsSmsNumbers = helpers.GetStringList(value.Array())
			} else {
				data.AlertDestinationsSmsNumbers = types.ListNull(types.StringType)
			}
			if value := res.Get("filters.failureType"); value.Exists() && value.Value() != nil {
				data.FiltersFailureType = types.StringValue(value.String())
			} else {
				data.FiltersFailureType = types.StringNull()
			}
			if value := res.Get("filters.lookbackWindow"); value.Exists() && value.Value() != nil {
				data.FiltersLookbackWindow = types.Int64Value(value.Int())
			} else {
				data.FiltersLookbackWindow = types.Int64Null()
			}
			if value := res.Get("filters.minDuration"); value.Exists() && value.Value() != nil {
				data.FiltersMinDuration = types.Int64Value(value.Int())
			} else {
				data.FiltersMinDuration = types.Int64Null()
			}
			if value := res.Get("filters.name"); value.Exists() && value.Value() != nil {
				data.FiltersName = types.StringValue(value.String())
			} else {
				data.FiltersName = types.StringNull()
			}
			if value := res.Get("filters.period"); value.Exists() && value.Value() != nil {
				data.FiltersPeriod = types.Int64Value(value.Int())
			} else {
				data.FiltersPeriod = types.Int64Null()
			}
			if value := res.Get("filters.priority"); value.Exists() && value.Value() != nil {
				data.FiltersPriority = types.StringValue(value.String())
			} else {
				data.FiltersPriority = types.StringNull()
			}
			if value := res.Get("filters.regex"); value.Exists() && value.Value() != nil {
				data.FiltersRegex = types.StringValue(value.String())
			} else {
				data.FiltersRegex = types.StringNull()
			}
			if value := res.Get("filters.selector"); value.Exists() && value.Value() != nil {
				data.FiltersSelector = types.StringValue(value.String())
			} else {
				data.FiltersSelector = types.StringNull()
			}
			if value := res.Get("filters.ssidNum"); value.Exists() && value.Value() != nil {
				data.FiltersSsidNum = types.Int64Value(value.Int())
			} else {
				data.FiltersSsidNum = types.Int64Null()
			}
			if value := res.Get("filters.tag"); value.Exists() && value.Value() != nil {
				data.FiltersTag = types.StringValue(value.String())
			} else {
				data.FiltersTag = types.StringNull()
			}
			if value := res.Get("filters.threshold"); value.Exists() && value.Value() != nil {
				data.FiltersThreshold = types.Int64Value(value.Int())
			} else {
				data.FiltersThreshold = types.Int64Null()
			}
			if value := res.Get("filters.timeout"); value.Exists() && value.Value() != nil {
				data.FiltersTimeout = types.Int64Value(value.Int())
			} else {
				data.FiltersTimeout = types.Int64Null()
			}
			if value := res.Get("filters.conditions"); value.Exists() && value.Value() != nil {
				data.FiltersConditions = make([]NetworkAlertsSettingsAlertsFiltersConditions, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := NetworkAlertsSettingsAlertsFiltersConditions{}
					if value := res.Get("direction"); value.Exists() && value.Value() != nil {
						data.Direction = types.StringValue(value.String())
					} else {
						data.Direction = types.StringNull()
					}
					if value := res.Get("duration"); value.Exists() && value.Value() != nil {
						data.Duration = types.Int64Value(value.Int())
					} else {
						data.Duration = types.Int64Null()
					}
					if value := res.Get("threshold"); value.Exists() && value.Value() != nil {
						data.Threshold = types.Float64Value(value.Float())
					} else {
						data.Threshold = types.Float64Null()
					}
					if value := res.Get("type"); value.Exists() && value.Value() != nil {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					if value := res.Get("unit"); value.Exists() && value.Value() != nil {
						data.Unit = types.StringValue(value.String())
					} else {
						data.Unit = types.StringNull()
					}
					(*parent).FiltersConditions = append((*parent).FiltersConditions, data)
					return true
				})
			}
			if value := res.Get("filters.serials"); value.Exists() && value.Value() != nil {
				data.FiltersSerials = helpers.GetStringSet(value.Array())
			} else {
				data.FiltersSerials = types.SetNull(types.StringType)
			}
			(*parent).Alerts = append((*parent).Alerts, data)
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
func (data *NetworkAlertsSettings) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("defaultDestinations.allAdmins"); value.Exists() && !data.DefaultDestinationsAllAdmins.IsNull() {
		data.DefaultDestinationsAllAdmins = types.BoolValue(value.Bool())
	} else {
		data.DefaultDestinationsAllAdmins = types.BoolNull()
	}
	if value := res.Get("defaultDestinations.snmp"); value.Exists() && !data.DefaultDestinationsSnmp.IsNull() {
		data.DefaultDestinationsSnmp = types.BoolValue(value.Bool())
	} else {
		data.DefaultDestinationsSnmp = types.BoolNull()
	}
	if value := res.Get("defaultDestinations.emails"); value.Exists() && !data.DefaultDestinationsEmails.IsNull() {
		data.DefaultDestinationsEmails = helpers.GetStringList(value.Array())
	} else {
		data.DefaultDestinationsEmails = types.ListNull(types.StringType)
	}
	if value := res.Get("defaultDestinations.httpServerIds"); value.Exists() && !data.DefaultDestinationsHttpServerIds.IsNull() {
		data.DefaultDestinationsHttpServerIds = helpers.GetStringList(value.Array())
	} else {
		data.DefaultDestinationsHttpServerIds = types.ListNull(types.StringType)
	}
	if value := res.Get("muting.byPortSchedules.enabled"); value.Exists() && !data.MutingByPortSchedulesEnabled.IsNull() {
		data.MutingByPortSchedulesEnabled = types.BoolValue(value.Bool())
	} else {
		data.MutingByPortSchedulesEnabled = types.BoolNull()
	}
	for i := 0; i < len(data.Alerts); i++ {
		keys := [...]string{"type"}
		keyValues := [...]string{data.Alerts[i].Type.ValueString()}

		parent := &data
		data := (*parent).Alerts[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("alerts").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Alerts[%d] = %+v",
				i,
				(*parent).Alerts[i],
			))
			(*parent).Alerts = slices.Delete((*parent).Alerts, i, i+1)
			i--

			continue
		}
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("alertDestinations.allAdmins"); value.Exists() && !data.AlertDestinationsAllAdmins.IsNull() {
			data.AlertDestinationsAllAdmins = types.BoolValue(value.Bool())
		} else {
			data.AlertDestinationsAllAdmins = types.BoolNull()
		}
		if value := res.Get("alertDestinations.snmp"); value.Exists() && !data.AlertDestinationsSnmp.IsNull() {
			data.AlertDestinationsSnmp = types.BoolValue(value.Bool())
		} else {
			data.AlertDestinationsSnmp = types.BoolNull()
		}
		if value := res.Get("alertDestinations.emails"); value.Exists() && !data.AlertDestinationsEmails.IsNull() {
			data.AlertDestinationsEmails = helpers.GetStringList(value.Array())
		} else {
			data.AlertDestinationsEmails = types.ListNull(types.StringType)
		}
		if value := res.Get("alertDestinations.httpServerIds"); value.Exists() && !data.AlertDestinationsHttpServerIds.IsNull() {
			data.AlertDestinationsHttpServerIds = helpers.GetStringList(value.Array())
		} else {
			data.AlertDestinationsHttpServerIds = types.ListNull(types.StringType)
		}
		if value := res.Get("alertDestinations.smsNumbers"); value.Exists() && !data.AlertDestinationsSmsNumbers.IsNull() {
			data.AlertDestinationsSmsNumbers = helpers.GetStringList(value.Array())
		} else {
			data.AlertDestinationsSmsNumbers = types.ListNull(types.StringType)
		}
		if value := res.Get("filters.failureType"); value.Exists() && !data.FiltersFailureType.IsNull() {
			data.FiltersFailureType = types.StringValue(value.String())
		} else {
			data.FiltersFailureType = types.StringNull()
		}
		if value := res.Get("filters.lookbackWindow"); value.Exists() && !data.FiltersLookbackWindow.IsNull() {
			data.FiltersLookbackWindow = types.Int64Value(value.Int())
		} else {
			data.FiltersLookbackWindow = types.Int64Null()
		}
		if value := res.Get("filters.minDuration"); value.Exists() && !data.FiltersMinDuration.IsNull() {
			data.FiltersMinDuration = types.Int64Value(value.Int())
		} else {
			data.FiltersMinDuration = types.Int64Null()
		}
		if value := res.Get("filters.name"); value.Exists() && !data.FiltersName.IsNull() {
			data.FiltersName = types.StringValue(value.String())
		} else {
			data.FiltersName = types.StringNull()
		}
		if value := res.Get("filters.period"); value.Exists() && !data.FiltersPeriod.IsNull() {
			data.FiltersPeriod = types.Int64Value(value.Int())
		} else {
			data.FiltersPeriod = types.Int64Null()
		}
		if value := res.Get("filters.priority"); value.Exists() && !data.FiltersPriority.IsNull() {
			data.FiltersPriority = types.StringValue(value.String())
		} else {
			data.FiltersPriority = types.StringNull()
		}
		if value := res.Get("filters.regex"); value.Exists() && !data.FiltersRegex.IsNull() {
			data.FiltersRegex = types.StringValue(value.String())
		} else {
			data.FiltersRegex = types.StringNull()
		}
		if value := res.Get("filters.selector"); value.Exists() && !data.FiltersSelector.IsNull() {
			data.FiltersSelector = types.StringValue(value.String())
		} else {
			data.FiltersSelector = types.StringNull()
		}
		if value := res.Get("filters.ssidNum"); value.Exists() && !data.FiltersSsidNum.IsNull() {
			data.FiltersSsidNum = types.Int64Value(value.Int())
		} else {
			data.FiltersSsidNum = types.Int64Null()
		}
		if value := res.Get("filters.tag"); value.Exists() && !data.FiltersTag.IsNull() {
			data.FiltersTag = types.StringValue(value.String())
		} else {
			data.FiltersTag = types.StringNull()
		}
		if value := res.Get("filters.threshold"); value.Exists() && !data.FiltersThreshold.IsNull() {
			data.FiltersThreshold = types.Int64Value(value.Int())
		} else {
			data.FiltersThreshold = types.Int64Null()
		}
		if value := res.Get("filters.timeout"); value.Exists() && !data.FiltersTimeout.IsNull() {
			data.FiltersTimeout = types.Int64Value(value.Int())
		} else {
			data.FiltersTimeout = types.Int64Null()
		}
		for i := 0; i < len(data.FiltersConditions); i++ {
			keys := [...]string{"type"}
			keyValues := [...]string{data.FiltersConditions[i].Type.ValueString()}

			parent := &data
			data := (*parent).FiltersConditions[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("filters.conditions").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing FiltersConditions[%d] = %+v",
					i,
					(*parent).FiltersConditions[i],
				))
				(*parent).FiltersConditions = slices.Delete((*parent).FiltersConditions, i, i+1)
				i--

				continue
			}
			if value := res.Get("direction"); value.Exists() && !data.Direction.IsNull() {
				data.Direction = types.StringValue(value.String())
			} else {
				data.Direction = types.StringNull()
			}
			if value := res.Get("duration"); value.Exists() && !data.Duration.IsNull() {
				data.Duration = types.Int64Value(value.Int())
			} else {
				data.Duration = types.Int64Null()
			}
			if value := res.Get("threshold"); value.Exists() && !data.Threshold.IsNull() {
				data.Threshold = types.Float64Value(value.Float())
			} else {
				data.Threshold = types.Float64Null()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("unit"); value.Exists() && !data.Unit.IsNull() {
				data.Unit = types.StringValue(value.String())
			} else {
				data.Unit = types.StringNull()
			}
			(*parent).FiltersConditions[i] = data
		}
		if value := res.Get("filters.serials"); value.Exists() && !data.FiltersSerials.IsNull() {
			data.FiltersSerials = helpers.GetStringSet(value.Array())
		} else {
			data.FiltersSerials = types.SetNull(types.StringType)
		}
		(*parent).Alerts[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkAlertsSettings) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *NetworkAlertsSettingsIdentity) toIdentity(ctx context.Context, plan *NetworkAlertsSettings) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *NetworkAlertsSettings) fromIdentity(ctx context.Context, identity *NetworkAlertsSettingsIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkAlertsSettings) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
