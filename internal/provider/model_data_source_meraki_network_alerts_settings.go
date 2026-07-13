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
// It emits a separate model - DataSourceNetworkAlertsSettings - used only by the data source, always including every
// attribute (unlike model_resource.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model_resource.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model_resource.go's resource-side type names.

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

type DataSourceNetworkAlertsSettings struct {
	Id                               types.String                            `tfsdk:"id"`
	NetworkId                        types.String                            `tfsdk:"network_id"`
	DefaultDestinationsAllAdmins     types.Bool                              `tfsdk:"default_destinations_all_admins"`
	DefaultDestinationsSnmp          types.Bool                              `tfsdk:"default_destinations_snmp"`
	DefaultDestinationsEmails        types.List                              `tfsdk:"default_destinations_emails"`
	DefaultDestinationsHttpServerIds types.List                              `tfsdk:"default_destinations_http_server_ids"`
	MutingByPortSchedulesEnabled     types.Bool                              `tfsdk:"muting_by_port_schedules_enabled"`
	Alerts                           []DataSourceNetworkAlertsSettingsAlerts `tfsdk:"alerts"`
}

type DataSourceNetworkAlertsSettingsAlerts struct {
	Enabled                        types.Bool                                               `tfsdk:"enabled"`
	Type                           types.String                                             `tfsdk:"type"`
	AlertDestinationsAllAdmins     types.Bool                                               `tfsdk:"alert_destinations_all_admins"`
	AlertDestinationsSnmp          types.Bool                                               `tfsdk:"alert_destinations_snmp"`
	AlertDestinationsEmails        types.List                                               `tfsdk:"alert_destinations_emails"`
	AlertDestinationsHttpServerIds types.List                                               `tfsdk:"alert_destinations_http_server_ids"`
	AlertDestinationsSmsNumbers    types.List                                               `tfsdk:"alert_destinations_sms_numbers"`
	FiltersFailureType             types.String                                             `tfsdk:"filters_failure_type"`
	FiltersLookbackWindow          types.Int64                                              `tfsdk:"filters_lookback_window"`
	FiltersMinDuration             types.Int64                                              `tfsdk:"filters_min_duration"`
	FiltersName                    types.String                                             `tfsdk:"filters_name"`
	FiltersPeriod                  types.Int64                                              `tfsdk:"filters_period"`
	FiltersPriority                types.String                                             `tfsdk:"filters_priority"`
	FiltersRegex                   types.String                                             `tfsdk:"filters_regex"`
	FiltersSelector                types.String                                             `tfsdk:"filters_selector"`
	FiltersSsidNum                 types.Int64                                              `tfsdk:"filters_ssid_num"`
	FiltersTag                     types.String                                             `tfsdk:"filters_tag"`
	FiltersThreshold               types.Int64                                              `tfsdk:"filters_threshold"`
	FiltersTimeout                 types.Int64                                              `tfsdk:"filters_timeout"`
	FiltersConditions              []DataSourceNetworkAlertsSettingsAlertsFiltersConditions `tfsdk:"filters_conditions"`
	FiltersSerials                 types.Set                                                `tfsdk:"filters_serials"`
}

type DataSourceNetworkAlertsSettingsAlertsFiltersConditions struct {
	Direction types.String  `tfsdk:"direction"`
	Duration  types.Int64   `tfsdk:"duration"`
	Threshold types.Float64 `tfsdk:"threshold"`
	Type      types.String  `tfsdk:"type"`
	Unit      types.String  `tfsdk:"unit"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworkAlertsSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/alerts/settings", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkAlertsSettings) fromBody(ctx context.Context, res meraki.Res) {
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
		data.Alerts = make([]DataSourceNetworkAlertsSettingsAlerts, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceNetworkAlertsSettingsAlerts{}
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
				data.FiltersConditions = make([]DataSourceNetworkAlertsSettingsAlertsFiltersConditions, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DataSourceNetworkAlertsSettingsAlertsFiltersConditions{}
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
