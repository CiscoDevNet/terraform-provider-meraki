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
// It emits a separate model - DataSourceNetworkWirelessRadioRRM - used only by the data source, always including every
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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceNetworkWirelessRadioRRM struct {
	Id                             types.String `tfsdk:"id"`
	NetworkId                      types.String `tfsdk:"network_id"`
	AiEnabled                      types.Bool   `tfsdk:"ai_enabled"`
	BusyHourMinimizeChangesEnabled types.Bool   `tfsdk:"busy_hour_minimize_changes_enabled"`
	BusyHourScheduleMode           types.String `tfsdk:"busy_hour_schedule_mode"`
	BusyHourScheduleManualEnd      types.String `tfsdk:"busy_hour_schedule_manual_end"`
	BusyHourScheduleManualStart    types.String `tfsdk:"busy_hour_schedule_manual_start"`
	ChannelAvoidanceEnabled        types.Bool   `tfsdk:"channel_avoidance_enabled"`
	FraEnabled                     types.Bool   `tfsdk:"fra_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworkWirelessRadioRRM) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/radio/rrm", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkWirelessRadioRRM) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("ai.enabled"); value.Exists() && value.Value() != nil {
		data.AiEnabled = types.BoolValue(value.Bool())
	} else {
		data.AiEnabled = types.BoolNull()
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
}

// End of section. //template:end fromBody
