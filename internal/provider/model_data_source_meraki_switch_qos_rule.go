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
// It emits a separate model - DataSourceSwitchQoSRule - used only by the data source, always including every
// attribute (unlike model.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model.go's resource-side type names.

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

type DataSourceSwitchQoSRule struct {
	Id           types.String `tfsdk:"id"`
	NetworkId    types.String `tfsdk:"network_id"`
	Dscp         types.Int64  `tfsdk:"dscp"`
	DstPort      types.Int64  `tfsdk:"dst_port"`
	DstPortRange types.String `tfsdk:"dst_port_range"`
	Protocol     types.String `tfsdk:"protocol"`
	SrcPort      types.Int64  `tfsdk:"src_port"`
	SrcPortRange types.String `tfsdk:"src_port_range"`
	Vlan         types.Int64  `tfsdk:"vlan"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceSwitchQoSRule) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/qosRules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceSwitchQoSRule) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("dscp"); value.Exists() && value.Value() != nil {
		data.Dscp = types.Int64Value(value.Int())
	} else {
		data.Dscp = types.Int64Null()
	}
	if value := res.Get("dstPort"); value.Exists() && value.Value() != nil {
		data.DstPort = types.Int64Value(value.Int())
	} else {
		data.DstPort = types.Int64Null()
	}
	if value := res.Get("dstPortRange"); value.Exists() && value.Value() != nil {
		data.DstPortRange = types.StringValue(value.String())
	} else {
		data.DstPortRange = types.StringNull()
	}
	if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("srcPort"); value.Exists() && value.Value() != nil {
		data.SrcPort = types.Int64Value(value.Int())
	} else {
		data.SrcPort = types.Int64Null()
	}
	if value := res.Get("srcPortRange"); value.Exists() && value.Value() != nil {
		data.SrcPortRange = types.StringValue(value.String())
	} else {
		data.SrcPortRange = types.StringNull()
	}
	if value := res.Get("vlan"); value.Exists() && value.Value() != nil {
		data.Vlan = types.Int64Value(value.Int())
	} else {
		data.Vlan = types.Int64Null()
	}
}

// End of section. //template:end fromBody
