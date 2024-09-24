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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkNetflow struct {
	Id               types.String `tfsdk:"id"`
	NetworkId        types.String `tfsdk:"network_id"`
	CollectorIp      types.String `tfsdk:"collector_ip"`
	CollectorPort    types.Int64  `tfsdk:"collector_port"`
	EtaDstPort       types.Int64  `tfsdk:"eta_dst_port"`
	EtaEnabled       types.Bool   `tfsdk:"eta_enabled"`
	ReportingEnabled types.Bool   `tfsdk:"reporting_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkNetflow) getPath() string {
	return fmt.Sprintf("/networks/%v/netflow", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkNetflow) toBody(ctx context.Context, state NetworkNetflow) string {
	body := ""
	if !data.CollectorIp.IsNull() {
		body, _ = sjson.Set(body, "collectorIp", data.CollectorIp.ValueString())
	}
	if !data.CollectorPort.IsNull() {
		body, _ = sjson.Set(body, "collectorPort", data.CollectorPort.ValueInt64())
	}
	if !data.EtaDstPort.IsNull() {
		body, _ = sjson.Set(body, "etaDstPort", data.EtaDstPort.ValueInt64())
	}
	if !data.EtaEnabled.IsNull() {
		body, _ = sjson.Set(body, "etaEnabled", data.EtaEnabled.ValueBool())
	}
	if !data.ReportingEnabled.IsNull() {
		body, _ = sjson.Set(body, "reportingEnabled", data.ReportingEnabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkNetflow) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("collectorIp"); value.Exists() {
		data.CollectorIp = types.StringValue(value.String())
	} else {
		data.CollectorIp = types.StringNull()
	}
	if value := res.Get("collectorPort"); value.Exists() {
		data.CollectorPort = types.Int64Value(value.Int())
	} else {
		data.CollectorPort = types.Int64Null()
	}
	if value := res.Get("etaDstPort"); value.Exists() {
		data.EtaDstPort = types.Int64Value(value.Int())
	} else {
		data.EtaDstPort = types.Int64Null()
	}
	if value := res.Get("etaEnabled"); value.Exists() {
		data.EtaEnabled = types.BoolValue(value.Bool())
	} else {
		data.EtaEnabled = types.BoolNull()
	}
	if value := res.Get("reportingEnabled"); value.Exists() {
		data.ReportingEnabled = types.BoolValue(value.Bool())
	} else {
		data.ReportingEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkNetflow) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("collectorIp"); value.Exists() && !data.CollectorIp.IsNull() {
		data.CollectorIp = types.StringValue(value.String())
	} else {
		data.CollectorIp = types.StringNull()
	}
	if value := res.Get("collectorPort"); value.Exists() && !data.CollectorPort.IsNull() {
		data.CollectorPort = types.Int64Value(value.Int())
	} else {
		data.CollectorPort = types.Int64Null()
	}
	if value := res.Get("etaDstPort"); value.Exists() && !data.EtaDstPort.IsNull() {
		data.EtaDstPort = types.Int64Value(value.Int())
	} else {
		data.EtaDstPort = types.Int64Null()
	}
	if value := res.Get("etaEnabled"); value.Exists() && !data.EtaEnabled.IsNull() {
		data.EtaEnabled = types.BoolValue(value.Bool())
	} else {
		data.EtaEnabled = types.BoolNull()
	}
	if value := res.Get("reportingEnabled"); value.Exists() && !data.ReportingEnabled.IsNull() {
		data.ReportingEnabled = types.BoolValue(value.Bool())
	} else {
		data.ReportingEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial
