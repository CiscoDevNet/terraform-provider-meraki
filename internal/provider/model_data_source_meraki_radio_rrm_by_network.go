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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceRadioRRMByNetwork struct {
	OrganizationId types.String                       `tfsdk:"organization_id"`
	Items          []DataSourceRadioRRMByNetworkItems `tfsdk:"items"`
}

type DataSourceRadioRRMByNetworkItems struct {
	Id                             types.String                           `tfsdk:"id"`
	Name                           types.String                           `tfsdk:"name"`
	NetworkId                      types.String                           `tfsdk:"network_id"`
	TimeZone                       types.String                           `tfsdk:"time_zone"`
	AiRrmEnabled                   types.Bool                             `tfsdk:"ai_rrm_enabled"`
	AiRrmEnablementDate            types.String                           `tfsdk:"ai_rrm_enablement_date"`
	BusyHourMinimizeChangesEnabled types.Bool                             `tfsdk:"busy_hour_minimize_changes_enabled"`
	BusyHourScheduleMode           types.String                           `tfsdk:"busy_hour_schedule_mode"`
	BusyHourScheduleAutomaticEnd   types.String                           `tfsdk:"busy_hour_schedule_automatic_end"`
	BusyHourScheduleAutomaticStart types.String                           `tfsdk:"busy_hour_schedule_automatic_start"`
	BusyHourScheduleManualEnd      types.String                           `tfsdk:"busy_hour_schedule_manual_end"`
	BusyHourScheduleManualStart    types.String                           `tfsdk:"busy_hour_schedule_manual_start"`
	ChannelAvoidanceEnabled        types.Bool                             `tfsdk:"channel_avoidance_enabled"`
	FraEnabled                     types.Bool                             `tfsdk:"fra_enabled"`
	MetaCountsItemsRemaining       types.Int64                            `tfsdk:"meta_counts_items_remaining"`
	MetaCountsItemsTotal           types.Int64                            `tfsdk:"meta_counts_items_total"`
	Items                          []DataSourceRadioRRMByNetworkItemsItem `tfsdk:"items"`
}

type DataSourceRadioRRMByNetworkItemsItem struct {
	Name                           types.String `tfsdk:"name"`
	NetworkId                      types.String `tfsdk:"network_id"`
	TimeZone                       types.String `tfsdk:"time_zone"`
	AiEnabled                      types.Bool   `tfsdk:"ai_enabled"`
	AiLastEnabledAt                types.String `tfsdk:"ai_last_enabled_at"`
	BusyHourMinimizeChangesEnabled types.Bool   `tfsdk:"busy_hour_minimize_changes_enabled"`
	BusyHourScheduleMode           types.String `tfsdk:"busy_hour_schedule_mode"`
	BusyHourScheduleAutomaticEnd   types.String `tfsdk:"busy_hour_schedule_automatic_end"`
	BusyHourScheduleAutomaticStart types.String `tfsdk:"busy_hour_schedule_automatic_start"`
	BusyHourScheduleManualEnd      types.String `tfsdk:"busy_hour_schedule_manual_end"`
	BusyHourScheduleManualStart    types.String `tfsdk:"busy_hour_schedule_manual_start"`
	ChannelAvoidanceEnabled        types.Bool   `tfsdk:"channel_avoidance_enabled"`
	FraEnabled                     types.Bool   `tfsdk:"fra_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceRadioRRMByNetwork) getPath() string {
	return fmt.Sprintf("/organizations/%v/wireless/radio/rrm/byNetwork", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceRadioRRMByNetwork) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceRadioRRMByNetworkItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceRadioRRMByNetworkItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("networkId"); value.Exists() && value.Value() != nil {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("timeZone"); value.Exists() && value.Value() != nil {
			data.TimeZone = types.StringValue(value.String())
		} else {
			data.TimeZone = types.StringNull()
		}
		if value := res.Get("aiRrm.enabled"); value.Exists() && value.Value() != nil {
			data.AiRrmEnabled = types.BoolValue(value.Bool())
		} else {
			data.AiRrmEnabled = types.BoolNull()
		}
		if value := res.Get("aiRrm.enablementDate"); value.Exists() && value.Value() != nil {
			data.AiRrmEnablementDate = types.StringValue(value.String())
		} else {
			data.AiRrmEnablementDate = types.StringNull()
		}
		if value := res.Get("busyHour.minimizeChanges.enabled"); value.Exists() && value.Value() != nil {
			data.BusyHourMinimizeChangesEnabled = types.BoolValue(value.Bool())
		} else {
			data.BusyHourMinimizeChangesEnabled = types.BoolNull()
		}
		if value := res.Get("busyHour.schedule.mode"); value.Exists() && value.Value() != nil {
			data.BusyHourScheduleMode = types.StringValue(value.String())
		} else {
			data.BusyHourScheduleMode = types.StringNull()
		}
		if value := res.Get("busyHour.schedule.automatic.end"); value.Exists() && value.Value() != nil {
			data.BusyHourScheduleAutomaticEnd = types.StringValue(value.String())
		} else {
			data.BusyHourScheduleAutomaticEnd = types.StringNull()
		}
		if value := res.Get("busyHour.schedule.automatic.start"); value.Exists() && value.Value() != nil {
			data.BusyHourScheduleAutomaticStart = types.StringValue(value.String())
		} else {
			data.BusyHourScheduleAutomaticStart = types.StringNull()
		}
		if value := res.Get("busyHour.schedule.manual.end"); value.Exists() && value.Value() != nil {
			data.BusyHourScheduleManualEnd = types.StringValue(value.String())
		} else {
			data.BusyHourScheduleManualEnd = types.StringNull()
		}
		if value := res.Get("busyHour.schedule.manual.start"); value.Exists() && value.Value() != nil {
			data.BusyHourScheduleManualStart = types.StringValue(value.String())
		} else {
			data.BusyHourScheduleManualStart = types.StringNull()
		}
		if value := res.Get("channel.avoidance.enabled"); value.Exists() && value.Value() != nil {
			data.ChannelAvoidanceEnabled = types.BoolValue(value.Bool())
		} else {
			data.ChannelAvoidanceEnabled = types.BoolNull()
		}
		if value := res.Get("fra.enabled"); value.Exists() && value.Value() != nil {
			data.FraEnabled = types.BoolValue(value.Bool())
		} else {
			data.FraEnabled = types.BoolNull()
		}
		if value := res.Get("meta.counts.items.remaining"); value.Exists() && value.Value() != nil {
			data.MetaCountsItemsRemaining = types.Int64Value(value.Int())
		} else {
			data.MetaCountsItemsRemaining = types.Int64Null()
		}
		if value := res.Get("meta.counts.items.total"); value.Exists() && value.Value() != nil {
			data.MetaCountsItemsTotal = types.Int64Value(value.Int())
		} else {
			data.MetaCountsItemsTotal = types.Int64Null()
		}
		if value := res.Get("items"); value.Exists() && value.Value() != nil {
			data.Items = make([]DataSourceRadioRRMByNetworkItemsItem, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceRadioRRMByNetworkItemsItem{}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("networkId"); value.Exists() && value.Value() != nil {
					data.NetworkId = types.StringValue(value.String())
				} else {
					data.NetworkId = types.StringNull()
				}
				if value := res.Get("timeZone"); value.Exists() && value.Value() != nil {
					data.TimeZone = types.StringValue(value.String())
				} else {
					data.TimeZone = types.StringNull()
				}
				if value := res.Get("ai.enabled"); value.Exists() && value.Value() != nil {
					data.AiEnabled = types.BoolValue(value.Bool())
				} else {
					data.AiEnabled = types.BoolNull()
				}
				if value := res.Get("ai.lastEnabledAt"); value.Exists() && value.Value() != nil {
					data.AiLastEnabledAt = types.StringValue(value.String())
				} else {
					data.AiLastEnabledAt = types.StringNull()
				}
				if value := res.Get("busyHour.minimizeChanges.enabled"); value.Exists() && value.Value() != nil {
					data.BusyHourMinimizeChangesEnabled = types.BoolValue(value.Bool())
				} else {
					data.BusyHourMinimizeChangesEnabled = types.BoolNull()
				}
				if value := res.Get("busyHour.schedule.mode"); value.Exists() && value.Value() != nil {
					data.BusyHourScheduleMode = types.StringValue(value.String())
				} else {
					data.BusyHourScheduleMode = types.StringNull()
				}
				if value := res.Get("busyHour.schedule.automatic.end"); value.Exists() && value.Value() != nil {
					data.BusyHourScheduleAutomaticEnd = types.StringValue(value.String())
				} else {
					data.BusyHourScheduleAutomaticEnd = types.StringNull()
				}
				if value := res.Get("busyHour.schedule.automatic.start"); value.Exists() && value.Value() != nil {
					data.BusyHourScheduleAutomaticStart = types.StringValue(value.String())
				} else {
					data.BusyHourScheduleAutomaticStart = types.StringNull()
				}
				if value := res.Get("busyHour.schedule.manual.end"); value.Exists() && value.Value() != nil {
					data.BusyHourScheduleManualEnd = types.StringValue(value.String())
				} else {
					data.BusyHourScheduleManualEnd = types.StringNull()
				}
				if value := res.Get("busyHour.schedule.manual.start"); value.Exists() && value.Value() != nil {
					data.BusyHourScheduleManualStart = types.StringValue(value.String())
				} else {
					data.BusyHourScheduleManualStart = types.StringNull()
				}
				if value := res.Get("channel.avoidance.enabled"); value.Exists() && value.Value() != nil {
					data.ChannelAvoidanceEnabled = types.BoolValue(value.Bool())
				} else {
					data.ChannelAvoidanceEnabled = types.BoolNull()
				}
				if value := res.Get("fra.enabled"); value.Exists() && value.Value() != nil {
					data.FraEnabled = types.BoolValue(value.Bool())
				} else {
					data.FraEnabled = types.BoolNull()
				}
				(*parent).Items = append((*parent).Items, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
