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

type DataSourceSwitchPortSchedules struct {
	NetworkId types.String                         `tfsdk:"network_id"`
	Items     []DataSourceSwitchPortSchedulesItems `tfsdk:"items"`
}

type DataSourceSwitchPortSchedulesItems struct {
	Id                          types.String `tfsdk:"id"`
	Name                        types.String `tfsdk:"name"`
	PortScheduleFridayActive    types.Bool   `tfsdk:"port_schedule_friday_active"`
	PortScheduleFridayFrom      types.String `tfsdk:"port_schedule_friday_from"`
	PortScheduleFridayTo        types.String `tfsdk:"port_schedule_friday_to"`
	PortScheduleMondayActive    types.Bool   `tfsdk:"port_schedule_monday_active"`
	PortScheduleMondayFrom      types.String `tfsdk:"port_schedule_monday_from"`
	PortScheduleMondayTo        types.String `tfsdk:"port_schedule_monday_to"`
	PortScheduleSaturdayActive  types.Bool   `tfsdk:"port_schedule_saturday_active"`
	PortScheduleSaturdayFrom    types.String `tfsdk:"port_schedule_saturday_from"`
	PortScheduleSaturdayTo      types.String `tfsdk:"port_schedule_saturday_to"`
	PortScheduleSundayActive    types.Bool   `tfsdk:"port_schedule_sunday_active"`
	PortScheduleSundayFrom      types.String `tfsdk:"port_schedule_sunday_from"`
	PortScheduleSundayTo        types.String `tfsdk:"port_schedule_sunday_to"`
	PortScheduleThursdayActive  types.Bool   `tfsdk:"port_schedule_thursday_active"`
	PortScheduleThursdayFrom    types.String `tfsdk:"port_schedule_thursday_from"`
	PortScheduleThursdayTo      types.String `tfsdk:"port_schedule_thursday_to"`
	PortScheduleTuesdayActive   types.Bool   `tfsdk:"port_schedule_tuesday_active"`
	PortScheduleTuesdayFrom     types.String `tfsdk:"port_schedule_tuesday_from"`
	PortScheduleTuesdayTo       types.String `tfsdk:"port_schedule_tuesday_to"`
	PortScheduleWednesdayActive types.Bool   `tfsdk:"port_schedule_wednesday_active"`
	PortScheduleWednesdayFrom   types.String `tfsdk:"port_schedule_wednesday_from"`
	PortScheduleWednesdayTo     types.String `tfsdk:"port_schedule_wednesday_to"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceSwitchPortSchedules) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/portSchedules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceSwitchPortSchedules) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceSwitchPortSchedulesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceSwitchPortSchedulesItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("portSchedule.friday.active"); value.Exists() && value.Value() != nil {
			data.PortScheduleFridayActive = types.BoolValue(value.Bool())
		} else {
			data.PortScheduleFridayActive = types.BoolNull()
		}
		if value := res.Get("portSchedule.friday.from"); value.Exists() && value.Value() != nil {
			data.PortScheduleFridayFrom = types.StringValue(value.String())
		} else {
			data.PortScheduleFridayFrom = types.StringNull()
		}
		if value := res.Get("portSchedule.friday.to"); value.Exists() && value.Value() != nil {
			data.PortScheduleFridayTo = types.StringValue(value.String())
		} else {
			data.PortScheduleFridayTo = types.StringNull()
		}
		if value := res.Get("portSchedule.monday.active"); value.Exists() && value.Value() != nil {
			data.PortScheduleMondayActive = types.BoolValue(value.Bool())
		} else {
			data.PortScheduleMondayActive = types.BoolNull()
		}
		if value := res.Get("portSchedule.monday.from"); value.Exists() && value.Value() != nil {
			data.PortScheduleMondayFrom = types.StringValue(value.String())
		} else {
			data.PortScheduleMondayFrom = types.StringNull()
		}
		if value := res.Get("portSchedule.monday.to"); value.Exists() && value.Value() != nil {
			data.PortScheduleMondayTo = types.StringValue(value.String())
		} else {
			data.PortScheduleMondayTo = types.StringNull()
		}
		if value := res.Get("portSchedule.saturday.active"); value.Exists() && value.Value() != nil {
			data.PortScheduleSaturdayActive = types.BoolValue(value.Bool())
		} else {
			data.PortScheduleSaturdayActive = types.BoolNull()
		}
		if value := res.Get("portSchedule.saturday.from"); value.Exists() && value.Value() != nil {
			data.PortScheduleSaturdayFrom = types.StringValue(value.String())
		} else {
			data.PortScheduleSaturdayFrom = types.StringNull()
		}
		if value := res.Get("portSchedule.saturday.to"); value.Exists() && value.Value() != nil {
			data.PortScheduleSaturdayTo = types.StringValue(value.String())
		} else {
			data.PortScheduleSaturdayTo = types.StringNull()
		}
		if value := res.Get("portSchedule.sunday.active"); value.Exists() && value.Value() != nil {
			data.PortScheduleSundayActive = types.BoolValue(value.Bool())
		} else {
			data.PortScheduleSundayActive = types.BoolNull()
		}
		if value := res.Get("portSchedule.sunday.from"); value.Exists() && value.Value() != nil {
			data.PortScheduleSundayFrom = types.StringValue(value.String())
		} else {
			data.PortScheduleSundayFrom = types.StringNull()
		}
		if value := res.Get("portSchedule.sunday.to"); value.Exists() && value.Value() != nil {
			data.PortScheduleSundayTo = types.StringValue(value.String())
		} else {
			data.PortScheduleSundayTo = types.StringNull()
		}
		if value := res.Get("portSchedule.thursday.active"); value.Exists() && value.Value() != nil {
			data.PortScheduleThursdayActive = types.BoolValue(value.Bool())
		} else {
			data.PortScheduleThursdayActive = types.BoolNull()
		}
		if value := res.Get("portSchedule.thursday.from"); value.Exists() && value.Value() != nil {
			data.PortScheduleThursdayFrom = types.StringValue(value.String())
		} else {
			data.PortScheduleThursdayFrom = types.StringNull()
		}
		if value := res.Get("portSchedule.thursday.to"); value.Exists() && value.Value() != nil {
			data.PortScheduleThursdayTo = types.StringValue(value.String())
		} else {
			data.PortScheduleThursdayTo = types.StringNull()
		}
		if value := res.Get("portSchedule.tuesday.active"); value.Exists() && value.Value() != nil {
			data.PortScheduleTuesdayActive = types.BoolValue(value.Bool())
		} else {
			data.PortScheduleTuesdayActive = types.BoolNull()
		}
		if value := res.Get("portSchedule.tuesday.from"); value.Exists() && value.Value() != nil {
			data.PortScheduleTuesdayFrom = types.StringValue(value.String())
		} else {
			data.PortScheduleTuesdayFrom = types.StringNull()
		}
		if value := res.Get("portSchedule.tuesday.to"); value.Exists() && value.Value() != nil {
			data.PortScheduleTuesdayTo = types.StringValue(value.String())
		} else {
			data.PortScheduleTuesdayTo = types.StringNull()
		}
		if value := res.Get("portSchedule.wednesday.active"); value.Exists() && value.Value() != nil {
			data.PortScheduleWednesdayActive = types.BoolValue(value.Bool())
		} else {
			data.PortScheduleWednesdayActive = types.BoolNull()
		}
		if value := res.Get("portSchedule.wednesday.from"); value.Exists() && value.Value() != nil {
			data.PortScheduleWednesdayFrom = types.StringValue(value.String())
		} else {
			data.PortScheduleWednesdayFrom = types.StringNull()
		}
		if value := res.Get("portSchedule.wednesday.to"); value.Exists() && value.Value() != nil {
			data.PortScheduleWednesdayTo = types.StringValue(value.String())
		} else {
			data.PortScheduleWednesdayTo = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
