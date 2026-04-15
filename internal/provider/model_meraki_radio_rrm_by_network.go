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

type RadioRRMByNetwork struct {
	Id                       types.String             `tfsdk:"id"`
	OrganizationId           types.String             `tfsdk:"organization_id"`
	MetaCountsItemsRemaining types.Int64              `tfsdk:"meta_counts_items_remaining"`
	MetaCountsItemsTotal     types.Int64              `tfsdk:"meta_counts_items_total"`
	Items                    []RadioRRMByNetworkItems `tfsdk:"items"`
}

type RadioRRMByNetworkItems struct {
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

func (data RadioRRMByNetwork) getPath() string {
	return fmt.Sprintf("/organizations/%v/wireless/radio/rrm/byNetwork", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data RadioRRMByNetwork) toBody(ctx context.Context, state RadioRRMByNetwork) string {
	body := ""
	if !data.MetaCountsItemsRemaining.IsNull() {
		body, _ = sjson.Set(body, "meta.counts.items.remaining", data.MetaCountsItemsRemaining.ValueInt64())
	}
	if !data.MetaCountsItemsTotal.IsNull() {
		body, _ = sjson.Set(body, "meta.counts.items.total", data.MetaCountsItemsTotal.ValueInt64())
	}
	if data.Items != nil {
		body, _ = sjson.Set(body, "items", []interface{}{})
		for _, item := range data.Items {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.NetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "networkId", item.NetworkId.ValueString())
			}
			if !item.TimeZone.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "timeZone", item.TimeZone.ValueString())
			}
			if !item.AiEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ai.enabled", item.AiEnabled.ValueBool())
			}
			if !item.AiLastEnabledAt.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ai.lastEnabledAt", item.AiLastEnabledAt.ValueString())
			}
			if !item.BusyHourMinimizeChangesEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "busyHour.minimizeChanges.enabled", item.BusyHourMinimizeChangesEnabled.ValueBool())
			}
			if !item.BusyHourScheduleMode.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "busyHour.schedule.mode", item.BusyHourScheduleMode.ValueString())
			}
			if !item.BusyHourScheduleAutomaticEnd.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "busyHour.schedule.automatic.end", item.BusyHourScheduleAutomaticEnd.ValueString())
			}
			if !item.BusyHourScheduleAutomaticStart.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "busyHour.schedule.automatic.start", item.BusyHourScheduleAutomaticStart.ValueString())
			}
			if !item.BusyHourScheduleManualEnd.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "busyHour.schedule.manual.end", item.BusyHourScheduleManualEnd.ValueString())
			}
			if !item.BusyHourScheduleManualStart.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "busyHour.schedule.manual.start", item.BusyHourScheduleManualStart.ValueString())
			}
			if !item.ChannelAvoidanceEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "channel.avoidance.enabled", item.ChannelAvoidanceEnabled.ValueBool())
			}
			if !item.FraEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "fra.enabled", item.FraEnabled.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *RadioRRMByNetwork) fromBody(ctx context.Context, res meraki.Res) {
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
		data.Items = make([]RadioRRMByNetworkItems, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := RadioRRMByNetworkItems{}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *RadioRRMByNetwork) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("meta.counts.items.remaining"); value.Exists() && !data.MetaCountsItemsRemaining.IsNull() {
		data.MetaCountsItemsRemaining = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsRemaining = types.Int64Null()
	}
	if value := res.Get("meta.counts.items.total"); value.Exists() && !data.MetaCountsItemsTotal.IsNull() {
		data.MetaCountsItemsTotal = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsTotal = types.Int64Null()
	}
	for i := 0; i < len(data.Items); i++ {
		keys := [...]string{"name", "networkId", "timeZone", "ai.enabled", "ai.lastEnabledAt", "busyHour.minimizeChanges.enabled", "busyHour.schedule.mode", "busyHour.schedule.automatic.end", "busyHour.schedule.automatic.start", "busyHour.schedule.manual.end", "busyHour.schedule.manual.start", "channel.avoidance.enabled", "fra.enabled"}
		keyValues := [...]string{data.Items[i].Name.ValueString(), data.Items[i].NetworkId.ValueString(), data.Items[i].TimeZone.ValueString(), strconv.FormatBool(data.Items[i].AiEnabled.ValueBool()), data.Items[i].AiLastEnabledAt.ValueString(), strconv.FormatBool(data.Items[i].BusyHourMinimizeChangesEnabled.ValueBool()), data.Items[i].BusyHourScheduleMode.ValueString(), data.Items[i].BusyHourScheduleAutomaticEnd.ValueString(), data.Items[i].BusyHourScheduleAutomaticStart.ValueString(), data.Items[i].BusyHourScheduleManualEnd.ValueString(), data.Items[i].BusyHourScheduleManualStart.ValueString(), strconv.FormatBool(data.Items[i].ChannelAvoidanceEnabled.ValueBool()), strconv.FormatBool(data.Items[i].FraEnabled.ValueBool())}

		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Items[%d] = %+v",
				i,
				(*parent).Items[i],
			))
			(*parent).Items = slices.Delete((*parent).Items, i, i+1)
			i--

			continue
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("networkId"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("timeZone"); value.Exists() && !data.TimeZone.IsNull() {
			data.TimeZone = types.StringValue(value.String())
		} else {
			data.TimeZone = types.StringNull()
		}
		if value := res.Get("ai.enabled"); value.Exists() && !data.AiEnabled.IsNull() {
			data.AiEnabled = types.BoolValue(value.Bool())
		} else {
			data.AiEnabled = types.BoolNull()
		}
		if value := res.Get("ai.lastEnabledAt"); value.Exists() && !data.AiLastEnabledAt.IsNull() {
			data.AiLastEnabledAt = types.StringValue(value.String())
		} else {
			data.AiLastEnabledAt = types.StringNull()
		}
		if value := res.Get("busyHour.minimizeChanges.enabled"); value.Exists() && !data.BusyHourMinimizeChangesEnabled.IsNull() {
			data.BusyHourMinimizeChangesEnabled = types.BoolValue(value.Bool())
		} else {
			data.BusyHourMinimizeChangesEnabled = types.BoolNull()
		}
		if value := res.Get("busyHour.schedule.mode"); value.Exists() && !data.BusyHourScheduleMode.IsNull() {
			data.BusyHourScheduleMode = types.StringValue(value.String())
		} else {
			data.BusyHourScheduleMode = types.StringNull()
		}
		if value := res.Get("busyHour.schedule.automatic.end"); value.Exists() && !data.BusyHourScheduleAutomaticEnd.IsNull() {
			data.BusyHourScheduleAutomaticEnd = types.StringValue(value.String())
		} else {
			data.BusyHourScheduleAutomaticEnd = types.StringNull()
		}
		if value := res.Get("busyHour.schedule.automatic.start"); value.Exists() && !data.BusyHourScheduleAutomaticStart.IsNull() {
			data.BusyHourScheduleAutomaticStart = types.StringValue(value.String())
		} else {
			data.BusyHourScheduleAutomaticStart = types.StringNull()
		}
		if value := res.Get("busyHour.schedule.manual.end"); value.Exists() && !data.BusyHourScheduleManualEnd.IsNull() {
			data.BusyHourScheduleManualEnd = types.StringValue(value.String())
		} else {
			data.BusyHourScheduleManualEnd = types.StringNull()
		}
		if value := res.Get("busyHour.schedule.manual.start"); value.Exists() && !data.BusyHourScheduleManualStart.IsNull() {
			data.BusyHourScheduleManualStart = types.StringValue(value.String())
		} else {
			data.BusyHourScheduleManualStart = types.StringNull()
		}
		if value := res.Get("channel.avoidance.enabled"); value.Exists() && !data.ChannelAvoidanceEnabled.IsNull() {
			data.ChannelAvoidanceEnabled = types.BoolValue(value.Bool())
		} else {
			data.ChannelAvoidanceEnabled = types.BoolNull()
		}
		if value := res.Get("fra.enabled"); value.Exists() && !data.FraEnabled.IsNull() {
			data.FraEnabled = types.BoolValue(value.Bool())
		} else {
			data.FraEnabled = types.BoolNull()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *RadioRRMByNetwork) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data RadioRRMByNetwork) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
