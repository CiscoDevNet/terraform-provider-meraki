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

type InsightMonitoredMediaServer struct {
	Id                          types.String `tfsdk:"id"`
	OrganizationId              types.String `tfsdk:"organization_id"`
	Address                     types.String `tfsdk:"address"`
	BestEffortMonitoringEnabled types.Bool   `tfsdk:"best_effort_monitoring_enabled"`
	Name                        types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data InsightMonitoredMediaServer) getPath() string {
	return fmt.Sprintf("/organizations/%v/insight/monitoredMediaServers", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data InsightMonitoredMediaServer) toBody(ctx context.Context, state InsightMonitoredMediaServer) string {
	body := ""
	if !data.Address.IsNull() {
		body, _ = sjson.Set(body, "address", data.Address.ValueString())
	}
	if !data.BestEffortMonitoringEnabled.IsNull() {
		body, _ = sjson.Set(body, "bestEffortMonitoringEnabled", data.BestEffortMonitoringEnabled.ValueBool())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *InsightMonitoredMediaServer) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("address"); value.Exists() && value.Value() != nil {
		data.Address = types.StringValue(value.String())
	} else {
		data.Address = types.StringNull()
	}
	if value := res.Get("bestEffortMonitoringEnabled"); value.Exists() && value.Value() != nil {
		data.BestEffortMonitoringEnabled = types.BoolValue(value.Bool())
	} else {
		data.BestEffortMonitoringEnabled = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *InsightMonitoredMediaServer) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("address"); value.Exists() && !data.Address.IsNull() {
		data.Address = types.StringValue(value.String())
	} else {
		data.Address = types.StringNull()
	}
	if value := res.Get("bestEffortMonitoringEnabled"); value.Exists() && !data.BestEffortMonitoringEnabled.IsNull() {
		data.BestEffortMonitoringEnabled = types.BoolValue(value.Bool())
	} else {
		data.BestEffortMonitoringEnabled = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *InsightMonitoredMediaServer) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data InsightMonitoredMediaServer) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
