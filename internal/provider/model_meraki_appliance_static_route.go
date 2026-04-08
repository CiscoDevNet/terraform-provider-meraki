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

type ApplianceStaticRoute struct {
	Id            types.String `tfsdk:"id"`
	NetworkId     types.String `tfsdk:"network_id"`
	GatewayIp     types.String `tfsdk:"gateway_ip"`
	GatewayVlanId types.String `tfsdk:"gateway_vlan_id"`
	Name          types.String `tfsdk:"name"`
	Subnet        types.String `tfsdk:"subnet"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceStaticRoute) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/staticRoutes", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceStaticRoute) toBody(ctx context.Context, state ApplianceStaticRoute) string {
	body := ""
	if !data.GatewayIp.IsNull() {
		body, _ = sjson.Set(body, "gatewayIp", data.GatewayIp.ValueString())
	}
	if !data.GatewayVlanId.IsNull() {
		body, _ = sjson.Set(body, "gatewayVlanId", data.GatewayVlanId.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Subnet.IsNull() {
		body, _ = sjson.Set(body, "subnet", data.Subnet.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceStaticRoute) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("gatewayIp"); value.Exists() && value.Value() != nil {
		data.GatewayIp = types.StringValue(value.String())
	} else {
		data.GatewayIp = types.StringNull()
	}
	if value := res.Get("gatewayVlanId"); value.Exists() && value.Value() != nil {
		data.GatewayVlanId = types.StringValue(value.String())
	} else {
		data.GatewayVlanId = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("subnet"); value.Exists() && value.Value() != nil {
		data.Subnet = types.StringValue(value.String())
	} else {
		data.Subnet = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceStaticRoute) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("gatewayIp"); value.Exists() && !data.GatewayIp.IsNull() {
		data.GatewayIp = types.StringValue(value.String())
	} else {
		data.GatewayIp = types.StringNull()
	}
	if value := res.Get("gatewayVlanId"); value.Exists() && !data.GatewayVlanId.IsNull() {
		data.GatewayVlanId = types.StringValue(value.String())
	} else {
		data.GatewayVlanId = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("subnet"); value.Exists() && !data.Subnet.IsNull() {
		data.Subnet = types.StringValue(value.String())
	} else {
		data.Subnet = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceStaticRoute) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns
