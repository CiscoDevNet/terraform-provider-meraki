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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessSSIDSchedules struct {
	Id              types.String                           `tfsdk:"id"`
	NetworkId       types.String                           `tfsdk:"network_id"`
	Number          types.String                           `tfsdk:"number"`
	Enabled         types.Bool                             `tfsdk:"enabled"`
	Ranges          []WirelessSSIDSchedulesRanges          `tfsdk:"ranges"`
	RangesInSeconds []WirelessSSIDSchedulesRangesInSeconds `tfsdk:"ranges_in_seconds"`
}

type WirelessSSIDSchedulesRanges struct {
	EndDay    types.String `tfsdk:"end_day"`
	EndTime   types.String `tfsdk:"end_time"`
	StartDay  types.String `tfsdk:"start_day"`
	StartTime types.String `tfsdk:"start_time"`
}

type WirelessSSIDSchedulesRangesInSeconds struct {
	End   types.Int64 `tfsdk:"end"`
	Start types.Int64 `tfsdk:"start"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessSSIDSchedules) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/schedules", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSSIDSchedules) toBody(ctx context.Context, state WirelessSSIDSchedules) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if len(data.Ranges) > 0 {
		body, _ = sjson.Set(body, "ranges", []interface{}{})
		for _, item := range data.Ranges {
			itemBody := ""
			if !item.EndDay.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "endDay", item.EndDay.ValueString())
			}
			if !item.EndTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "endTime", item.EndTime.ValueString())
			}
			if !item.StartDay.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "startDay", item.StartDay.ValueString())
			}
			if !item.StartTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "startTime", item.StartTime.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ranges.-1", itemBody)
		}
	}
	if len(data.RangesInSeconds) > 0 {
		body, _ = sjson.Set(body, "rangesInSeconds", []interface{}{})
		for _, item := range data.RangesInSeconds {
			itemBody := ""
			if !item.End.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "end", item.End.ValueInt64())
			}
			if !item.Start.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "start", item.Start.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "rangesInSeconds.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSSIDSchedules) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("ranges"); value.Exists() && value.Value() != nil {
		data.Ranges = make([]WirelessSSIDSchedulesRanges, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDSchedulesRanges{}
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
		data.RangesInSeconds = make([]WirelessSSIDSchedulesRangesInSeconds, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDSchedulesRangesInSeconds{}
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessSSIDSchedules) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	for i := 0; i < len(data.Ranges); i++ {
		keys := [...]string{"endDay", "endTime", "startDay", "startTime"}
		keyValues := [...]string{data.Ranges[i].EndDay.ValueString(), data.Ranges[i].EndTime.ValueString(), data.Ranges[i].StartDay.ValueString(), data.Ranges[i].StartTime.ValueString()}

		parent := &data
		data := (*parent).Ranges[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ranges").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ranges[%d] = %+v",
				i,
				(*parent).Ranges[i],
			))
			(*parent).Ranges = slices.Delete((*parent).Ranges, i, i+1)
			i--

			continue
		}
		if value := res.Get("endDay"); value.Exists() && !data.EndDay.IsNull() {
			data.EndDay = types.StringValue(value.String())
		} else {
			data.EndDay = types.StringNull()
		}
		if value := res.Get("endTime"); value.Exists() && !data.EndTime.IsNull() {
			data.EndTime = types.StringValue(value.String())
		} else {
			data.EndTime = types.StringNull()
		}
		if value := res.Get("startDay"); value.Exists() && !data.StartDay.IsNull() {
			data.StartDay = types.StringValue(value.String())
		} else {
			data.StartDay = types.StringNull()
		}
		if value := res.Get("startTime"); value.Exists() && !data.StartTime.IsNull() {
			data.StartTime = types.StringValue(value.String())
		} else {
			data.StartTime = types.StringNull()
		}
		(*parent).Ranges[i] = data
	}
	for i := 0; i < len(data.RangesInSeconds); i++ {
		keys := [...]string{"end", "start"}
		keyValues := [...]string{strconv.FormatInt(data.RangesInSeconds[i].End.ValueInt64(), 10), strconv.FormatInt(data.RangesInSeconds[i].Start.ValueInt64(), 10)}

		parent := &data
		data := (*parent).RangesInSeconds[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("rangesInSeconds").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing RangesInSeconds[%d] = %+v",
				i,
				(*parent).RangesInSeconds[i],
			))
			(*parent).RangesInSeconds = slices.Delete((*parent).RangesInSeconds, i, i+1)
			i--

			continue
		}
		if value := res.Get("end"); value.Exists() && !data.End.IsNull() {
			data.End = types.Int64Value(value.Int())
		} else {
			data.End = types.Int64Null()
		}
		if value := res.Get("start"); value.Exists() && !data.Start.IsNull() {
			data.Start = types.Int64Value(value.Int())
		} else {
			data.Start = types.Int64Null()
		}
		(*parent).RangesInSeconds[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessSSIDSchedules) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data WirelessSSIDSchedules) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
