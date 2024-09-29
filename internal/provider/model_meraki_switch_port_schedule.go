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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchPortSchedule struct {
	Id                          types.String `tfsdk:"id"`
	NetworkId                   types.String `tfsdk:"network_id"`
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

func (data SwitchPortSchedule) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/portSchedules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchPortSchedule) toBody(ctx context.Context, state SwitchPortSchedule) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.PortScheduleFridayActive.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.friday.active", data.PortScheduleFridayActive.ValueBool())
	}
	if !data.PortScheduleFridayFrom.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.friday.from", data.PortScheduleFridayFrom.ValueString())
	}
	if !data.PortScheduleFridayTo.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.friday.to", data.PortScheduleFridayTo.ValueString())
	}
	if !data.PortScheduleMondayActive.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.monday.active", data.PortScheduleMondayActive.ValueBool())
	}
	if !data.PortScheduleMondayFrom.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.monday.from", data.PortScheduleMondayFrom.ValueString())
	}
	if !data.PortScheduleMondayTo.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.monday.to", data.PortScheduleMondayTo.ValueString())
	}
	if !data.PortScheduleSaturdayActive.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.saturday.active", data.PortScheduleSaturdayActive.ValueBool())
	}
	if !data.PortScheduleSaturdayFrom.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.saturday.from", data.PortScheduleSaturdayFrom.ValueString())
	}
	if !data.PortScheduleSaturdayTo.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.saturday.to", data.PortScheduleSaturdayTo.ValueString())
	}
	if !data.PortScheduleSundayActive.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.sunday.active", data.PortScheduleSundayActive.ValueBool())
	}
	if !data.PortScheduleSundayFrom.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.sunday.from", data.PortScheduleSundayFrom.ValueString())
	}
	if !data.PortScheduleSundayTo.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.sunday.to", data.PortScheduleSundayTo.ValueString())
	}
	if !data.PortScheduleThursdayActive.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.thursday.active", data.PortScheduleThursdayActive.ValueBool())
	}
	if !data.PortScheduleThursdayFrom.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.thursday.from", data.PortScheduleThursdayFrom.ValueString())
	}
	if !data.PortScheduleThursdayTo.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.thursday.to", data.PortScheduleThursdayTo.ValueString())
	}
	if !data.PortScheduleTuesdayActive.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.tuesday.active", data.PortScheduleTuesdayActive.ValueBool())
	}
	if !data.PortScheduleTuesdayFrom.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.tuesday.from", data.PortScheduleTuesdayFrom.ValueString())
	}
	if !data.PortScheduleTuesdayTo.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.tuesday.to", data.PortScheduleTuesdayTo.ValueString())
	}
	if !data.PortScheduleWednesdayActive.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.wednesday.active", data.PortScheduleWednesdayActive.ValueBool())
	}
	if !data.PortScheduleWednesdayFrom.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.wednesday.from", data.PortScheduleWednesdayFrom.ValueString())
	}
	if !data.PortScheduleWednesdayTo.IsNull() {
		body, _ = sjson.Set(body, "portSchedule.wednesday.to", data.PortScheduleWednesdayTo.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchPortSchedule) fromBody(ctx context.Context, res meraki.Res) {
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchPortSchedule) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("portSchedule.friday.active"); value.Exists() && !data.PortScheduleFridayActive.IsNull() {
		data.PortScheduleFridayActive = types.BoolValue(value.Bool())
	} else {
		data.PortScheduleFridayActive = types.BoolNull()
	}
	if value := res.Get("portSchedule.friday.from"); value.Exists() && !data.PortScheduleFridayFrom.IsNull() {
		data.PortScheduleFridayFrom = types.StringValue(value.String())
	} else {
		data.PortScheduleFridayFrom = types.StringNull()
	}
	if value := res.Get("portSchedule.friday.to"); value.Exists() && !data.PortScheduleFridayTo.IsNull() {
		data.PortScheduleFridayTo = types.StringValue(value.String())
	} else {
		data.PortScheduleFridayTo = types.StringNull()
	}
	if value := res.Get("portSchedule.monday.active"); value.Exists() && !data.PortScheduleMondayActive.IsNull() {
		data.PortScheduleMondayActive = types.BoolValue(value.Bool())
	} else {
		data.PortScheduleMondayActive = types.BoolNull()
	}
	if value := res.Get("portSchedule.monday.from"); value.Exists() && !data.PortScheduleMondayFrom.IsNull() {
		data.PortScheduleMondayFrom = types.StringValue(value.String())
	} else {
		data.PortScheduleMondayFrom = types.StringNull()
	}
	if value := res.Get("portSchedule.monday.to"); value.Exists() && !data.PortScheduleMondayTo.IsNull() {
		data.PortScheduleMondayTo = types.StringValue(value.String())
	} else {
		data.PortScheduleMondayTo = types.StringNull()
	}
	if value := res.Get("portSchedule.saturday.active"); value.Exists() && !data.PortScheduleSaturdayActive.IsNull() {
		data.PortScheduleSaturdayActive = types.BoolValue(value.Bool())
	} else {
		data.PortScheduleSaturdayActive = types.BoolNull()
	}
	if value := res.Get("portSchedule.saturday.from"); value.Exists() && !data.PortScheduleSaturdayFrom.IsNull() {
		data.PortScheduleSaturdayFrom = types.StringValue(value.String())
	} else {
		data.PortScheduleSaturdayFrom = types.StringNull()
	}
	if value := res.Get("portSchedule.saturday.to"); value.Exists() && !data.PortScheduleSaturdayTo.IsNull() {
		data.PortScheduleSaturdayTo = types.StringValue(value.String())
	} else {
		data.PortScheduleSaturdayTo = types.StringNull()
	}
	if value := res.Get("portSchedule.sunday.active"); value.Exists() && !data.PortScheduleSundayActive.IsNull() {
		data.PortScheduleSundayActive = types.BoolValue(value.Bool())
	} else {
		data.PortScheduleSundayActive = types.BoolNull()
	}
	if value := res.Get("portSchedule.sunday.from"); value.Exists() && !data.PortScheduleSundayFrom.IsNull() {
		data.PortScheduleSundayFrom = types.StringValue(value.String())
	} else {
		data.PortScheduleSundayFrom = types.StringNull()
	}
	if value := res.Get("portSchedule.sunday.to"); value.Exists() && !data.PortScheduleSundayTo.IsNull() {
		data.PortScheduleSundayTo = types.StringValue(value.String())
	} else {
		data.PortScheduleSundayTo = types.StringNull()
	}
	if value := res.Get("portSchedule.thursday.active"); value.Exists() && !data.PortScheduleThursdayActive.IsNull() {
		data.PortScheduleThursdayActive = types.BoolValue(value.Bool())
	} else {
		data.PortScheduleThursdayActive = types.BoolNull()
	}
	if value := res.Get("portSchedule.thursday.from"); value.Exists() && !data.PortScheduleThursdayFrom.IsNull() {
		data.PortScheduleThursdayFrom = types.StringValue(value.String())
	} else {
		data.PortScheduleThursdayFrom = types.StringNull()
	}
	if value := res.Get("portSchedule.thursday.to"); value.Exists() && !data.PortScheduleThursdayTo.IsNull() {
		data.PortScheduleThursdayTo = types.StringValue(value.String())
	} else {
		data.PortScheduleThursdayTo = types.StringNull()
	}
	if value := res.Get("portSchedule.tuesday.active"); value.Exists() && !data.PortScheduleTuesdayActive.IsNull() {
		data.PortScheduleTuesdayActive = types.BoolValue(value.Bool())
	} else {
		data.PortScheduleTuesdayActive = types.BoolNull()
	}
	if value := res.Get("portSchedule.tuesday.from"); value.Exists() && !data.PortScheduleTuesdayFrom.IsNull() {
		data.PortScheduleTuesdayFrom = types.StringValue(value.String())
	} else {
		data.PortScheduleTuesdayFrom = types.StringNull()
	}
	if value := res.Get("portSchedule.tuesday.to"); value.Exists() && !data.PortScheduleTuesdayTo.IsNull() {
		data.PortScheduleTuesdayTo = types.StringValue(value.String())
	} else {
		data.PortScheduleTuesdayTo = types.StringNull()
	}
	if value := res.Get("portSchedule.wednesday.active"); value.Exists() && !data.PortScheduleWednesdayActive.IsNull() {
		data.PortScheduleWednesdayActive = types.BoolValue(value.Bool())
	} else {
		data.PortScheduleWednesdayActive = types.BoolNull()
	}
	if value := res.Get("portSchedule.wednesday.from"); value.Exists() && !data.PortScheduleWednesdayFrom.IsNull() {
		data.PortScheduleWednesdayFrom = types.StringValue(value.String())
	} else {
		data.PortScheduleWednesdayFrom = types.StringNull()
	}
	if value := res.Get("portSchedule.wednesday.to"); value.Exists() && !data.PortScheduleWednesdayTo.IsNull() {
		data.PortScheduleWednesdayTo = types.StringValue(value.String())
	} else {
		data.PortScheduleWednesdayTo = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial
