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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceManagementInterface struct {
	Id                   types.String `tfsdk:"id"`
	Serial               types.String `tfsdk:"serial"`
	Wan1StaticGatewayIp  types.String `tfsdk:"wan1_static_gateway_ip"`
	Wan1StaticIp         types.String `tfsdk:"wan1_static_ip"`
	Wan1StaticSubnetMask types.String `tfsdk:"wan1_static_subnet_mask"`
	Wan1UsingStaticIp    types.Bool   `tfsdk:"wan1_using_static_ip"`
	Wan1Vlan             types.Int64  `tfsdk:"wan1_vlan"`
	Wan1WanEnabled       types.String `tfsdk:"wan1_wan_enabled"`
	Wan1StaticDns        types.List   `tfsdk:"wan1_static_dns"`
	Wan2StaticGatewayIp  types.String `tfsdk:"wan2_static_gateway_ip"`
	Wan2StaticIp         types.String `tfsdk:"wan2_static_ip"`
	Wan2StaticSubnetMask types.String `tfsdk:"wan2_static_subnet_mask"`
	Wan2UsingStaticIp    types.Bool   `tfsdk:"wan2_using_static_ip"`
	Wan2Vlan             types.Int64  `tfsdk:"wan2_vlan"`
	Wan2WanEnabled       types.String `tfsdk:"wan2_wan_enabled"`
	Wan2StaticDns        types.List   `tfsdk:"wan2_static_dns"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceManagementInterface) getPath() string {
	return fmt.Sprintf("/devices/%v/managementInterface", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceManagementInterface) toBody(ctx context.Context, state DeviceManagementInterface) string {
	body := ""
	if !data.Wan1StaticGatewayIp.IsNull() {
		body, _ = sjson.Set(body, "wan1.staticGatewayIp", data.Wan1StaticGatewayIp.ValueString())
	}
	if !data.Wan1StaticIp.IsNull() {
		body, _ = sjson.Set(body, "wan1.staticIp", data.Wan1StaticIp.ValueString())
	}
	if !data.Wan1StaticSubnetMask.IsNull() {
		body, _ = sjson.Set(body, "wan1.staticSubnetMask", data.Wan1StaticSubnetMask.ValueString())
	}
	if !data.Wan1UsingStaticIp.IsNull() {
		body, _ = sjson.Set(body, "wan1.usingStaticIp", data.Wan1UsingStaticIp.ValueBool())
	}
	if !data.Wan1Vlan.IsNull() {
		body, _ = sjson.Set(body, "wan1.vlan", data.Wan1Vlan.ValueInt64())
	}
	if !data.Wan1WanEnabled.IsNull() {
		body, _ = sjson.Set(body, "wan1.wanEnabled", data.Wan1WanEnabled.ValueString())
	}
	if !data.Wan1StaticDns.IsNull() {
		var values []string
		data.Wan1StaticDns.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "wan1.staticDns", values)
	}
	if !data.Wan2StaticGatewayIp.IsNull() {
		body, _ = sjson.Set(body, "wan2.staticGatewayIp", data.Wan2StaticGatewayIp.ValueString())
	}
	if !data.Wan2StaticIp.IsNull() {
		body, _ = sjson.Set(body, "wan2.staticIp", data.Wan2StaticIp.ValueString())
	}
	if !data.Wan2StaticSubnetMask.IsNull() {
		body, _ = sjson.Set(body, "wan2.staticSubnetMask", data.Wan2StaticSubnetMask.ValueString())
	}
	if !data.Wan2UsingStaticIp.IsNull() {
		body, _ = sjson.Set(body, "wan2.usingStaticIp", data.Wan2UsingStaticIp.ValueBool())
	}
	if !data.Wan2Vlan.IsNull() {
		body, _ = sjson.Set(body, "wan2.vlan", data.Wan2Vlan.ValueInt64())
	}
	if !data.Wan2WanEnabled.IsNull() {
		body, _ = sjson.Set(body, "wan2.wanEnabled", data.Wan2WanEnabled.ValueString())
	}
	if !data.Wan2StaticDns.IsNull() {
		var values []string
		data.Wan2StaticDns.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "wan2.staticDns", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceManagementInterface) fromBody(ctx context.Context, res meraki.Res) {
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceManagementInterface) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("wan1.staticGatewayIp"); value.Exists() && !data.Wan1StaticGatewayIp.IsNull() {
		data.Wan1StaticGatewayIp = types.StringValue(value.String())
	} else {
		data.Wan1StaticGatewayIp = types.StringNull()
	}
	if value := res.Get("wan1.staticIp"); value.Exists() && !data.Wan1StaticIp.IsNull() {
		data.Wan1StaticIp = types.StringValue(value.String())
	} else {
		data.Wan1StaticIp = types.StringNull()
	}
	if value := res.Get("wan1.staticSubnetMask"); value.Exists() && !data.Wan1StaticSubnetMask.IsNull() {
		data.Wan1StaticSubnetMask = types.StringValue(value.String())
	} else {
		data.Wan1StaticSubnetMask = types.StringNull()
	}
	if value := res.Get("wan1.usingStaticIp"); value.Exists() && !data.Wan1UsingStaticIp.IsNull() {
		data.Wan1UsingStaticIp = types.BoolValue(value.Bool())
	} else {
		data.Wan1UsingStaticIp = types.BoolNull()
	}
	if value := res.Get("wan1.vlan"); value.Exists() && !data.Wan1Vlan.IsNull() {
		data.Wan1Vlan = types.Int64Value(value.Int())
	} else {
		data.Wan1Vlan = types.Int64Null()
	}
	if value := res.Get("wan1.wanEnabled"); value.Exists() && !data.Wan1WanEnabled.IsNull() {
		data.Wan1WanEnabled = types.StringValue(value.String())
	} else {
		data.Wan1WanEnabled = types.StringNull()
	}
	if value := res.Get("wan1.staticDns"); value.Exists() && !data.Wan1StaticDns.IsNull() {
		data.Wan1StaticDns = helpers.GetStringList(value.Array())
	} else {
		data.Wan1StaticDns = types.ListNull(types.StringType)
	}
	if value := res.Get("wan2.staticGatewayIp"); value.Exists() && !data.Wan2StaticGatewayIp.IsNull() {
		data.Wan2StaticGatewayIp = types.StringValue(value.String())
	} else {
		data.Wan2StaticGatewayIp = types.StringNull()
	}
	if value := res.Get("wan2.staticIp"); value.Exists() && !data.Wan2StaticIp.IsNull() {
		data.Wan2StaticIp = types.StringValue(value.String())
	} else {
		data.Wan2StaticIp = types.StringNull()
	}
	if value := res.Get("wan2.staticSubnetMask"); value.Exists() && !data.Wan2StaticSubnetMask.IsNull() {
		data.Wan2StaticSubnetMask = types.StringValue(value.String())
	} else {
		data.Wan2StaticSubnetMask = types.StringNull()
	}
	if value := res.Get("wan2.usingStaticIp"); value.Exists() && !data.Wan2UsingStaticIp.IsNull() {
		data.Wan2UsingStaticIp = types.BoolValue(value.Bool())
	} else {
		data.Wan2UsingStaticIp = types.BoolNull()
	}
	if value := res.Get("wan2.vlan"); value.Exists() && !data.Wan2Vlan.IsNull() {
		data.Wan2Vlan = types.Int64Value(value.Int())
	} else {
		data.Wan2Vlan = types.Int64Null()
	}
	if value := res.Get("wan2.wanEnabled"); value.Exists() && !data.Wan2WanEnabled.IsNull() {
		data.Wan2WanEnabled = types.StringValue(value.String())
	} else {
		data.Wan2WanEnabled = types.StringNull()
	}
	if value := res.Get("wan2.staticDns"); value.Exists() && !data.Wan2StaticDns.IsNull() {
		data.Wan2StaticDns = helpers.GetStringList(value.Array())
	} else {
		data.Wan2StaticDns = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceManagementInterface) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data DeviceManagementInterface) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
