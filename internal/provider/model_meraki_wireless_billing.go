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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessBilling struct {
	Id        types.String           `tfsdk:"id"`
	NetworkId types.String           `tfsdk:"network_id"`
	Currency  types.String           `tfsdk:"currency"`
	Plans     []WirelessBillingPlans `tfsdk:"plans"`
}

type WirelessBillingPlans struct {
	Price                    types.Float64 `tfsdk:"price"`
	TimeLimit                types.String  `tfsdk:"time_limit"`
	BandwidthLimitsLimitDown types.Int64   `tfsdk:"bandwidth_limits_limit_down"`
	BandwidthLimitsLimitUp   types.Int64   `tfsdk:"bandwidth_limits_limit_up"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessBilling) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/billing", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessBilling) toBody(ctx context.Context, state WirelessBilling) string {
	body := ""
	if !data.Currency.IsNull() {
		body, _ = sjson.Set(body, "currency", data.Currency.ValueString())
	}
	if len(data.Plans) > 0 {
		body, _ = sjson.Set(body, "plans", []interface{}{})
		for _, item := range data.Plans {
			itemBody := ""
			if !item.Price.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "price", item.Price.ValueFloat64())
			}
			if !item.TimeLimit.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "timeLimit", item.TimeLimit.ValueString())
			}
			if !item.BandwidthLimitsLimitDown.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "bandwidthLimits.limitDown", item.BandwidthLimitsLimitDown.ValueInt64())
			}
			if !item.BandwidthLimitsLimitUp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "bandwidthLimits.limitUp", item.BandwidthLimitsLimitUp.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "plans.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessBilling) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("currency"); value.Exists() && value.Value() != nil {
		data.Currency = types.StringValue(value.String())
	} else {
		data.Currency = types.StringNull()
	}
	if value := res.Get("plans"); value.Exists() && value.Value() != nil {
		data.Plans = make([]WirelessBillingPlans, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessBillingPlans{}
			if value := res.Get("price"); value.Exists() && value.Value() != nil {
				data.Price = types.Float64Value(value.Float())
			} else {
				data.Price = types.Float64Null()
			}
			if value := res.Get("timeLimit"); value.Exists() && value.Value() != nil {
				data.TimeLimit = types.StringValue(value.String())
			} else {
				data.TimeLimit = types.StringNull()
			}
			if value := res.Get("bandwidthLimits.limitDown"); value.Exists() && value.Value() != nil {
				data.BandwidthLimitsLimitDown = types.Int64Value(value.Int())
			} else {
				data.BandwidthLimitsLimitDown = types.Int64Null()
			}
			if value := res.Get("bandwidthLimits.limitUp"); value.Exists() && value.Value() != nil {
				data.BandwidthLimitsLimitUp = types.Int64Value(value.Int())
			} else {
				data.BandwidthLimitsLimitUp = types.Int64Null()
			}
			(*parent).Plans = append((*parent).Plans, data)
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
func (data *WirelessBilling) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("currency"); value.Exists() && !data.Currency.IsNull() {
		data.Currency = types.StringValue(value.String())
	} else {
		data.Currency = types.StringNull()
	}
	{
		l := len(res.Get("plans").Array())
		tflog.Debug(ctx, fmt.Sprintf("plans array resizing from %d to %d", len(data.Plans), l))
		if len(data.Plans) > l {
			data.Plans = data.Plans[:l]
		}
	}
	for i := range data.Plans {
		parent := &data
		data := (*parent).Plans[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("plans.%d", i))
		if value := res.Get("price"); value.Exists() && !data.Price.IsNull() {
			data.Price = types.Float64Value(value.Float())
		} else {
			data.Price = types.Float64Null()
		}
		if value := res.Get("timeLimit"); value.Exists() && !data.TimeLimit.IsNull() {
			data.TimeLimit = types.StringValue(value.String())
		} else {
			data.TimeLimit = types.StringNull()
		}
		if value := res.Get("bandwidthLimits.limitDown"); value.Exists() && !data.BandwidthLimitsLimitDown.IsNull() {
			data.BandwidthLimitsLimitDown = types.Int64Value(value.Int())
		} else {
			data.BandwidthLimitsLimitDown = types.Int64Null()
		}
		if value := res.Get("bandwidthLimits.limitUp"); value.Exists() && !data.BandwidthLimitsLimitUp.IsNull() {
			data.BandwidthLimitsLimitUp = types.Int64Value(value.Int())
		} else {
			data.BandwidthLimitsLimitUp = types.Int64Null()
		}
		(*parent).Plans[i] = data
	}
}

// End of section. //template:end fromBodyPartial
