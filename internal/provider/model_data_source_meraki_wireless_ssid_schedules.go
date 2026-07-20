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
// It emits a separate model - DataSourceWirelessSSIDSchedules - used only by the data source, always including every
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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessSSIDSchedules struct {
	Id              types.String                                     `tfsdk:"id"`
	NetworkId       types.String                                     `tfsdk:"network_id"`
	Number          types.String                                     `tfsdk:"number"`
	Enabled         types.Bool                                       `tfsdk:"enabled"`
	Ranges          []DataSourceWirelessSSIDSchedulesRanges          `tfsdk:"ranges"`
	RangesInSeconds []DataSourceWirelessSSIDSchedulesRangesInSeconds `tfsdk:"ranges_in_seconds"`
}

type DataSourceWirelessSSIDSchedulesRanges struct {
	EndDay    types.String `tfsdk:"end_day"`
	EndTime   types.String `tfsdk:"end_time"`
	StartDay  types.String `tfsdk:"start_day"`
	StartTime types.String `tfsdk:"start_time"`
}

type DataSourceWirelessSSIDSchedulesRangesInSeconds struct {
	End   types.Int64 `tfsdk:"end"`
	Start types.Int64 `tfsdk:"start"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessSSIDSchedules) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/schedules", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessSSIDSchedules) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("ranges"); value.Exists() && value.Value() != nil {
		data.Ranges = make([]DataSourceWirelessSSIDSchedulesRanges, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDSchedulesRanges{}
			if value := res.Get("endDay"); value.Exists() && value.Value() != nil {
				data.EndDay = types.StringValue(value.String())
			} else {
				data.EndDay = types.StringNull()
			}
			if value := res.Get("endTime"); value.Exists() && value.Value() != nil {
				data.EndTime = types.StringValue(value.String())
			} else {
				data.EndTime = types.StringNull()
			}
			if value := res.Get("startDay"); value.Exists() && value.Value() != nil {
				data.StartDay = types.StringValue(value.String())
			} else {
				data.StartDay = types.StringNull()
			}
			if value := res.Get("startTime"); value.Exists() && value.Value() != nil {
				data.StartTime = types.StringValue(value.String())
			} else {
				data.StartTime = types.StringNull()
			}
			(*parent).Ranges = append((*parent).Ranges, data)
			return true
		})
	}
	if value := res.Get("rangesInSeconds"); value.Exists() && value.Value() != nil {
		data.RangesInSeconds = make([]DataSourceWirelessSSIDSchedulesRangesInSeconds, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDSchedulesRangesInSeconds{}
			if value := res.Get("end"); value.Exists() && value.Value() != nil {
				data.End = types.Int64Value(value.Int())
			} else {
				data.End = types.Int64Null()
			}
			if value := res.Get("start"); value.Exists() && value.Value() != nil {
				data.Start = types.Int64Value(value.Int())
			} else {
				data.Start = types.Int64Null()
			}
			(*parent).RangesInSeconds = append((*parent).RangesInSeconds, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
