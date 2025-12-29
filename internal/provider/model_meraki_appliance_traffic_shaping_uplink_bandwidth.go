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

type ApplianceTrafficShapingUplinkBandwidth struct {
	Id                types.String `tfsdk:"id"`
	NetworkId         types.String `tfsdk:"network_id"`
	CellularLimitDown types.Int64  `tfsdk:"cellular_limit_down"`
	CellularLimitUp   types.Int64  `tfsdk:"cellular_limit_up"`
	Wan1LimitDown     types.Int64  `tfsdk:"wan1_limit_down"`
	Wan1LimitUp       types.Int64  `tfsdk:"wan1_limit_up"`
	Wan2LimitDown     types.Int64  `tfsdk:"wan2_limit_down"`
	Wan2LimitUp       types.Int64  `tfsdk:"wan2_limit_up"`
}

type ApplianceTrafficShapingUplinkBandwidthIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceTrafficShapingUplinkBandwidth) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/trafficShaping/uplinkBandwidth", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceTrafficShapingUplinkBandwidth) toBody(ctx context.Context, state ApplianceTrafficShapingUplinkBandwidth) string {
	body := ""
	if !data.CellularLimitDown.IsNull() {
		body, _ = sjson.Set(body, "bandwidthLimits.cellular.limitDown", data.CellularLimitDown.ValueInt64())
	}
	if !data.CellularLimitUp.IsNull() {
		body, _ = sjson.Set(body, "bandwidthLimits.cellular.limitUp", data.CellularLimitUp.ValueInt64())
	}
	if !data.Wan1LimitDown.IsNull() {
		body, _ = sjson.Set(body, "bandwidthLimits.wan1.limitDown", data.Wan1LimitDown.ValueInt64())
	}
	if !data.Wan1LimitUp.IsNull() {
		body, _ = sjson.Set(body, "bandwidthLimits.wan1.limitUp", data.Wan1LimitUp.ValueInt64())
	}
	if !data.Wan2LimitDown.IsNull() {
		body, _ = sjson.Set(body, "bandwidthLimits.wan2.limitDown", data.Wan2LimitDown.ValueInt64())
	}
	if !data.Wan2LimitUp.IsNull() {
		body, _ = sjson.Set(body, "bandwidthLimits.wan2.limitUp", data.Wan2LimitUp.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceTrafficShapingUplinkBandwidth) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("bandwidthLimits.cellular.limitDown"); value.Exists() && value.Value() != nil {
		data.CellularLimitDown = types.Int64Value(value.Int())
	} else {
		data.CellularLimitDown = types.Int64Null()
	}
	if value := res.Get("bandwidthLimits.cellular.limitUp"); value.Exists() && value.Value() != nil {
		data.CellularLimitUp = types.Int64Value(value.Int())
	} else {
		data.CellularLimitUp = types.Int64Null()
	}
	if value := res.Get("bandwidthLimits.wan1.limitDown"); value.Exists() && value.Value() != nil {
		data.Wan1LimitDown = types.Int64Value(value.Int())
	} else {
		data.Wan1LimitDown = types.Int64Null()
	}
	if value := res.Get("bandwidthLimits.wan1.limitUp"); value.Exists() && value.Value() != nil {
		data.Wan1LimitUp = types.Int64Value(value.Int())
	} else {
		data.Wan1LimitUp = types.Int64Null()
	}
	if value := res.Get("bandwidthLimits.wan2.limitDown"); value.Exists() && value.Value() != nil {
		data.Wan2LimitDown = types.Int64Value(value.Int())
	} else {
		data.Wan2LimitDown = types.Int64Null()
	}
	if value := res.Get("bandwidthLimits.wan2.limitUp"); value.Exists() && value.Value() != nil {
		data.Wan2LimitUp = types.Int64Value(value.Int())
	} else {
		data.Wan2LimitUp = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceTrafficShapingUplinkBandwidth) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("bandwidthLimits.cellular.limitDown"); value.Exists() && !data.CellularLimitDown.IsNull() {
		data.CellularLimitDown = types.Int64Value(value.Int())
	} else {
		data.CellularLimitDown = types.Int64Null()
	}
	if value := res.Get("bandwidthLimits.cellular.limitUp"); value.Exists() && !data.CellularLimitUp.IsNull() {
		data.CellularLimitUp = types.Int64Value(value.Int())
	} else {
		data.CellularLimitUp = types.Int64Null()
	}
	if value := res.Get("bandwidthLimits.wan1.limitDown"); value.Exists() && !data.Wan1LimitDown.IsNull() {
		data.Wan1LimitDown = types.Int64Value(value.Int())
	} else {
		data.Wan1LimitDown = types.Int64Null()
	}
	if value := res.Get("bandwidthLimits.wan1.limitUp"); value.Exists() && !data.Wan1LimitUp.IsNull() {
		data.Wan1LimitUp = types.Int64Value(value.Int())
	} else {
		data.Wan1LimitUp = types.Int64Null()
	}
	if value := res.Get("bandwidthLimits.wan2.limitDown"); value.Exists() && !data.Wan2LimitDown.IsNull() {
		data.Wan2LimitDown = types.Int64Value(value.Int())
	} else {
		data.Wan2LimitDown = types.Int64Null()
	}
	if value := res.Get("bandwidthLimits.wan2.limitUp"); value.Exists() && !data.Wan2LimitUp.IsNull() {
		data.Wan2LimitUp = types.Int64Value(value.Int())
	} else {
		data.Wan2LimitUp = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceTrafficShapingUplinkBandwidth) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceTrafficShapingUplinkBandwidthIdentity) toIdentity(ctx context.Context, plan *ApplianceTrafficShapingUplinkBandwidth) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceTrafficShapingUplinkBandwidth) fromIdentity(ctx context.Context, identity *ApplianceTrafficShapingUplinkBandwidthIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceTrafficShapingUplinkBandwidth) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
