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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceTrafficShapingCustomPerformanceClasses struct {
	NetworkId types.String                                           `tfsdk:"network_id"`
	Items     []ApplianceTrafficShapingCustomPerformanceClassesItems `tfsdk:"items"`
}

type ApplianceTrafficShapingCustomPerformanceClassesItems struct {
	Id                types.String `tfsdk:"id"`
	MaxJitter         types.Int64  `tfsdk:"max_jitter"`
	MaxLatency        types.Int64  `tfsdk:"max_latency"`
	MaxLossPercentage types.Int64  `tfsdk:"max_loss_percentage"`
	Name              types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceTrafficShapingCustomPerformanceClasses) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/trafficShaping/customPerformanceClasses", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceTrafficShapingCustomPerformanceClasses) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ApplianceTrafficShapingCustomPerformanceClassesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ApplianceTrafficShapingCustomPerformanceClassesItems{}
		data.Id = types.StringValue(res.Get("customPerformanceClassId").String())
		if value := res.Get("maxJitter"); value.Exists() && value.Value() != nil {
			data.MaxJitter = types.Int64Value(value.Int())
		} else {
			data.MaxJitter = types.Int64Null()
		}
		if value := res.Get("maxLatency"); value.Exists() && value.Value() != nil {
			data.MaxLatency = types.Int64Value(value.Int())
		} else {
			data.MaxLatency = types.Int64Null()
		}
		if value := res.Get("maxLossPercentage"); value.Exists() && value.Value() != nil {
			data.MaxLossPercentage = types.Int64Value(value.Int())
		} else {
			data.MaxLossPercentage = types.Int64Null()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
