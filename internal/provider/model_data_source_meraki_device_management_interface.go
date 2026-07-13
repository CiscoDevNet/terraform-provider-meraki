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
// It emits a separate model - DataSourceDeviceManagementInterface - used only by the data source, always including every
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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceDeviceManagementInterface struct {
	Id                   types.String `tfsdk:"id"`
	Serial               types.String `tfsdk:"serial"`
	Wan1StaticGatewayIp  types.String `tfsdk:"wan1_static_gateway_ip"`
	Wan1StaticIp         types.String `tfsdk:"wan1_static_ip"`
	Wan1StaticSubnetMask types.String `tfsdk:"wan1_static_subnet_mask"`
	Wan1UsingStaticIp    types.Bool   `tfsdk:"wan1_using_static_ip"`
	Wan1Vlan             types.Int64  `tfsdk:"wan1_vlan"`
	Wan1WanEnabled       types.String `tfsdk:"wan1_wan_enabled"`
	Wan1StaticDns        types.List   `tfsdk:"wan1_static_dns"`
	Wan1VrfName          types.String `tfsdk:"wan1_vrf_name"`
	Wan2StaticGatewayIp  types.String `tfsdk:"wan2_static_gateway_ip"`
	Wan2StaticIp         types.String `tfsdk:"wan2_static_ip"`
	Wan2StaticSubnetMask types.String `tfsdk:"wan2_static_subnet_mask"`
	Wan2UsingStaticIp    types.Bool   `tfsdk:"wan2_using_static_ip"`
	Wan2Vlan             types.Int64  `tfsdk:"wan2_vlan"`
	Wan2WanEnabled       types.String `tfsdk:"wan2_wan_enabled"`
	Wan2StaticDns        types.List   `tfsdk:"wan2_static_dns"`
	Wan2VrfName          types.String `tfsdk:"wan2_vrf_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceDeviceManagementInterface) getPath() string {
	return fmt.Sprintf("/devices/%v/managementInterface", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceDeviceManagementInterface) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("wan1.staticGatewayIp"); value.Exists() && value.Value() != nil {
		data.Wan1StaticGatewayIp = types.StringValue(value.String())
	} else {
		data.Wan1StaticGatewayIp = types.StringNull()
	}
	if value := res.Get("wan1.staticIp"); value.Exists() && value.Value() != nil {
		data.Wan1StaticIp = types.StringValue(value.String())
	} else {
		data.Wan1StaticIp = types.StringNull()
	}
	if value := res.Get("wan1.staticSubnetMask"); value.Exists() && value.Value() != nil {
		data.Wan1StaticSubnetMask = types.StringValue(value.String())
	} else {
		data.Wan1StaticSubnetMask = types.StringNull()
	}
	if value := res.Get("wan1.usingStaticIp"); value.Exists() && value.Value() != nil {
		data.Wan1UsingStaticIp = types.BoolValue(value.Bool())
	} else {
		data.Wan1UsingStaticIp = types.BoolNull()
	}
	if value := res.Get("wan1.vlan"); value.Exists() && value.Value() != nil {
		data.Wan1Vlan = types.Int64Value(value.Int())
	} else {
		data.Wan1Vlan = types.Int64Null()
	}
	if value := res.Get("wan1.wanEnabled"); value.Exists() && value.Value() != nil {
		data.Wan1WanEnabled = types.StringValue(value.String())
	} else {
		data.Wan1WanEnabled = types.StringNull()
	}
	if value := res.Get("wan1.staticDns"); value.Exists() && value.Value() != nil {
		data.Wan1StaticDns = helpers.GetStringList(value.Array())
	} else {
		data.Wan1StaticDns = types.ListNull(types.StringType)
	}
	if value := res.Get("wan1.vrf.name"); value.Exists() && value.Value() != nil {
		data.Wan1VrfName = types.StringValue(value.String())
	} else {
		data.Wan1VrfName = types.StringNull()
	}
	if value := res.Get("wan2.staticGatewayIp"); value.Exists() && value.Value() != nil {
		data.Wan2StaticGatewayIp = types.StringValue(value.String())
	} else {
		data.Wan2StaticGatewayIp = types.StringNull()
	}
	if value := res.Get("wan2.staticIp"); value.Exists() && value.Value() != nil {
		data.Wan2StaticIp = types.StringValue(value.String())
	} else {
		data.Wan2StaticIp = types.StringNull()
	}
	if value := res.Get("wan2.staticSubnetMask"); value.Exists() && value.Value() != nil {
		data.Wan2StaticSubnetMask = types.StringValue(value.String())
	} else {
		data.Wan2StaticSubnetMask = types.StringNull()
	}
	if value := res.Get("wan2.usingStaticIp"); value.Exists() && value.Value() != nil {
		data.Wan2UsingStaticIp = types.BoolValue(value.Bool())
	} else {
		data.Wan2UsingStaticIp = types.BoolNull()
	}
	if value := res.Get("wan2.vlan"); value.Exists() && value.Value() != nil {
		data.Wan2Vlan = types.Int64Value(value.Int())
	} else {
		data.Wan2Vlan = types.Int64Null()
	}
	if value := res.Get("wan2.wanEnabled"); value.Exists() && value.Value() != nil {
		data.Wan2WanEnabled = types.StringValue(value.String())
	} else {
		data.Wan2WanEnabled = types.StringNull()
	}
	if value := res.Get("wan2.staticDns"); value.Exists() && value.Value() != nil {
		data.Wan2StaticDns = helpers.GetStringList(value.Array())
	} else {
		data.Wan2StaticDns = types.ListNull(types.StringType)
	}
	if value := res.Get("wan2.vrf.name"); value.Exists() && value.Value() != nil {
		data.Wan2VrfName = types.StringValue(value.String())
	} else {
		data.Wan2VrfName = types.StringNull()
	}
}

// End of section. //template:end fromBody
