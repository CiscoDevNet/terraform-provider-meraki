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
// It emits a separate model - DataSourceNetworkNetflow - used only by the data source, always including every
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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceNetworkNetflow struct {
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

func (data DataSourceNetworkNetflow) getPath() string {
	return fmt.Sprintf("/networks/%v/netflow", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkNetflow) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("collectorIp"); value.Exists() && value.Value() != nil {
		data.CollectorIp = types.StringValue(value.String())
	} else {
		data.CollectorIp = types.StringNull()
	}
	if value := res.Get("collectorPort"); value.Exists() && value.Value() != nil {
		data.CollectorPort = types.Int64Value(value.Int())
	} else {
		data.CollectorPort = types.Int64Null()
	}
	if value := res.Get("etaDstPort"); value.Exists() && value.Value() != nil {
		data.EtaDstPort = types.Int64Value(value.Int())
	} else {
		data.EtaDstPort = types.Int64Null()
	}
	if value := res.Get("etaEnabled"); value.Exists() && value.Value() != nil {
		data.EtaEnabled = types.BoolValue(value.Bool())
	} else {
		data.EtaEnabled = types.BoolNull()
	}
	if value := res.Get("reportingEnabled"); value.Exists() && value.Value() != nil {
		data.ReportingEnabled = types.BoolValue(value.Bool())
	} else {
		data.ReportingEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody
