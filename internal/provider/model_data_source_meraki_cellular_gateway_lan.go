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
// It emits a separate model - DataSourceCellularGatewayLAN - used only by the data source, always including every
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

type DataSourceCellularGatewayLAN struct {
	Id                 types.String                                     `tfsdk:"id"`
	Serial             types.String                                     `tfsdk:"serial"`
	FixedIpAssignments []DataSourceCellularGatewayLANFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	ReservedIpRanges   []DataSourceCellularGatewayLANReservedIpRanges   `tfsdk:"reserved_ip_ranges"`
}

type DataSourceCellularGatewayLANFixedIpAssignments struct {
	Ip   types.String `tfsdk:"ip"`
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
}

type DataSourceCellularGatewayLANReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceCellularGatewayLAN) getPath() string {
	return fmt.Sprintf("/devices/%v/cellularGateway/lan", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceCellularGatewayLAN) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("fixedIpAssignments"); value.Exists() && value.Value() != nil {
		data.FixedIpAssignments = make([]DataSourceCellularGatewayLANFixedIpAssignments, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceCellularGatewayLANFixedIpAssignments{}
			if value := res.Get("ip"); value.Exists() && value.Value() != nil {
				data.Ip = types.StringValue(value.String())
			} else {
				data.Ip = types.StringNull()
			}
			if value := res.Get("mac"); value.Exists() && value.Value() != nil {
				data.Mac = types.StringValue(value.String())
			} else {
				data.Mac = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).FixedIpAssignments = append((*parent).FixedIpAssignments, data)
			return true
		})
	}
	if value := res.Get("reservedIpRanges"); value.Exists() && value.Value() != nil {
		data.ReservedIpRanges = make([]DataSourceCellularGatewayLANReservedIpRanges, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceCellularGatewayLANReservedIpRanges{}
			if value := res.Get("comment"); value.Exists() && value.Value() != nil {
				data.Comment = types.StringValue(value.String())
			} else {
				data.Comment = types.StringNull()
			}
			if value := res.Get("end"); value.Exists() && value.Value() != nil {
				data.End = types.StringValue(value.String())
			} else {
				data.End = types.StringNull()
			}
			if value := res.Get("start"); value.Exists() && value.Value() != nil {
				data.Start = types.StringValue(value.String())
			} else {
				data.Start = types.StringNull()
			}
			(*parent).ReservedIpRanges = append((*parent).ReservedIpRanges, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
