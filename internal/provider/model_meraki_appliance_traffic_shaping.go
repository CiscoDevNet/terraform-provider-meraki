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

type ApplianceTrafficShaping struct {
	Id                       types.String `tfsdk:"id"`
	NetworkId                types.String `tfsdk:"network_id"`
	GlobalBandwidthLimitDown types.Int64  `tfsdk:"global_bandwidth_limit_down"`
	GlobalBandwidthLimitUp   types.Int64  `tfsdk:"global_bandwidth_limit_up"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceTrafficShaping) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/trafficShaping", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceTrafficShaping) toBody(ctx context.Context, state ApplianceTrafficShaping) string {
	body := ""
	if !data.GlobalBandwidthLimitDown.IsNull() {
		body, _ = sjson.Set(body, "globalBandwidthLimits.limitDown", data.GlobalBandwidthLimitDown.ValueInt64())
	}
	if !data.GlobalBandwidthLimitUp.IsNull() {
		body, _ = sjson.Set(body, "globalBandwidthLimits.limitUp", data.GlobalBandwidthLimitUp.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceTrafficShaping) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("globalBandwidthLimits.limitDown"); value.Exists() && value.Value() != nil {
		data.GlobalBandwidthLimitDown = types.Int64Value(value.Int())
	} else {
		data.GlobalBandwidthLimitDown = types.Int64Null()
	}
	if value := res.Get("globalBandwidthLimits.limitUp"); value.Exists() && value.Value() != nil {
		data.GlobalBandwidthLimitUp = types.Int64Value(value.Int())
	} else {
		data.GlobalBandwidthLimitUp = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceTrafficShaping) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("globalBandwidthLimits.limitDown"); value.Exists() && !data.GlobalBandwidthLimitDown.IsNull() {
		data.GlobalBandwidthLimitDown = types.Int64Value(value.Int())
	} else {
		data.GlobalBandwidthLimitDown = types.Int64Null()
	}
	if value := res.Get("globalBandwidthLimits.limitUp"); value.Exists() && !data.GlobalBandwidthLimitUp.IsNull() {
		data.GlobalBandwidthLimitUp = types.Int64Value(value.Int())
	} else {
		data.GlobalBandwidthLimitUp = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial