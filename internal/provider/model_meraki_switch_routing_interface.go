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

type SwitchRoutingInterface struct {
	Id                           types.String `tfsdk:"id"`
	Serial                       types.String `tfsdk:"serial"`
	DefaultGateway               types.String `tfsdk:"default_gateway"`
	InterfaceIp                  types.String `tfsdk:"interface_ip"`
	Mode                         types.String `tfsdk:"mode"`
	MulticastRouting             types.String `tfsdk:"multicast_routing"`
	Name                         types.String `tfsdk:"name"`
	Subnet                       types.String `tfsdk:"subnet"`
	SwitchPortId                 types.String `tfsdk:"switch_port_id"`
	VlanId                       types.Int64  `tfsdk:"vlan_id"`
	Ipv6Address                  types.String `tfsdk:"ipv6_address"`
	Ipv6AssignmentMode           types.String `tfsdk:"ipv6_assignment_mode"`
	Ipv6Gateway                  types.String `tfsdk:"ipv6_gateway"`
	Ipv6Prefix                   types.String `tfsdk:"ipv6_prefix"`
	OspfSettingsArea             types.String `tfsdk:"ospf_settings_area"`
	OspfSettingsCost             types.Int64  `tfsdk:"ospf_settings_cost"`
	OspfSettingsIsPassiveEnabled types.Bool   `tfsdk:"ospf_settings_is_passive_enabled"`
	OspfSettingsNetworkType      types.String `tfsdk:"ospf_settings_network_type"`
	VrfName                      types.String `tfsdk:"vrf_name"`
}

type SwitchRoutingInterfaceIdentity struct {
	Serial types.String `tfsdk:"serial"`
	Id     types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchRoutingInterface) getPath() string {
	return fmt.Sprintf("/devices/%v/switch/routing/interfaces", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchRoutingInterface) toBody(ctx context.Context, state SwitchRoutingInterface) string {
	body := ""
	if !data.DefaultGateway.IsNull() && data.DefaultGateway != state.DefaultGateway {
		body, _ = sjson.Set(body, "defaultGateway", data.DefaultGateway.ValueString())
	}
	if !data.InterfaceIp.IsNull() {
		body, _ = sjson.Set(body, "interfaceIp", data.InterfaceIp.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	if !data.MulticastRouting.IsNull() {
		body, _ = sjson.Set(body, "multicastRouting", data.MulticastRouting.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Subnet.IsNull() {
		body, _ = sjson.Set(body, "subnet", data.Subnet.ValueString())
	}
	if !data.SwitchPortId.IsNull() {
		body, _ = sjson.Set(body, "switchPortId", data.SwitchPortId.ValueString())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanId", data.VlanId.ValueInt64())
	}
	if !data.Ipv6Address.IsNull() {
		body, _ = sjson.Set(body, "ipv6.address", data.Ipv6Address.ValueString())
	}
	if !data.Ipv6AssignmentMode.IsNull() {
		body, _ = sjson.Set(body, "ipv6.assignmentMode", data.Ipv6AssignmentMode.ValueString())
	}
	if !data.Ipv6Gateway.IsNull() {
		body, _ = sjson.Set(body, "ipv6.gateway", data.Ipv6Gateway.ValueString())
	}
	if !data.Ipv6Prefix.IsNull() {
		body, _ = sjson.Set(body, "ipv6.prefix", data.Ipv6Prefix.ValueString())
	}
	if !data.OspfSettingsArea.IsNull() {
		body, _ = sjson.Set(body, "ospfSettings.area", data.OspfSettingsArea.ValueString())
	}
	if !data.OspfSettingsCost.IsNull() {
		body, _ = sjson.Set(body, "ospfSettings.cost", data.OspfSettingsCost.ValueInt64())
	}
	if !data.OspfSettingsIsPassiveEnabled.IsNull() {
		body, _ = sjson.Set(body, "ospfSettings.isPassiveEnabled", data.OspfSettingsIsPassiveEnabled.ValueBool())
	}
	if !data.OspfSettingsNetworkType.IsNull() {
		body, _ = sjson.Set(body, "ospfSettings.networkType", data.OspfSettingsNetworkType.ValueString())
	}
	if !data.VrfName.IsNull() {
		body, _ = sjson.Set(body, "vrf.name", data.VrfName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchRoutingInterface) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("defaultGateway"); value.Exists() && value.Value() != nil {
		data.DefaultGateway = types.StringValue(value.String())
	} else {
		data.DefaultGateway = types.StringNull()
	}
	if value := res.Get("interfaceIp"); value.Exists() && value.Value() != nil {
		data.InterfaceIp = types.StringValue(value.String())
	} else {
		data.InterfaceIp = types.StringNull()
	}
	if value := res.Get("mode"); value.Exists() && value.Value() != nil {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("multicastRouting"); value.Exists() && value.Value() != nil {
		data.MulticastRouting = types.StringValue(value.String())
	} else {
		data.MulticastRouting = types.StringNull()
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
	if value := res.Get("switchPortId"); value.Exists() && value.Value() != nil {
		data.SwitchPortId = types.StringValue(value.String())
	} else {
		data.SwitchPortId = types.StringNull()
	}
	if value := res.Get("vlanId"); value.Exists() && value.Value() != nil {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("ipv6.address"); value.Exists() && value.Value() != nil {
		data.Ipv6Address = types.StringValue(value.String())
	} else {
		data.Ipv6Address = types.StringNull()
	}
	if value := res.Get("ipv6.assignmentMode"); value.Exists() && value.Value() != nil {
		data.Ipv6AssignmentMode = types.StringValue(value.String())
	} else {
		data.Ipv6AssignmentMode = types.StringNull()
	}
	if value := res.Get("ipv6.gateway"); value.Exists() && value.Value() != nil {
		data.Ipv6Gateway = types.StringValue(value.String())
	} else {
		data.Ipv6Gateway = types.StringNull()
	}
	if value := res.Get("ipv6.prefix"); value.Exists() && value.Value() != nil {
		data.Ipv6Prefix = types.StringValue(value.String())
	} else {
		data.Ipv6Prefix = types.StringNull()
	}
	if value := res.Get("ospfSettings.area"); value.Exists() && value.Value() != nil {
		data.OspfSettingsArea = types.StringValue(value.String())
	} else {
		data.OspfSettingsArea = types.StringNull()
	}
	if value := res.Get("ospfSettings.cost"); value.Exists() && value.Value() != nil {
		data.OspfSettingsCost = types.Int64Value(value.Int())
	} else {
		data.OspfSettingsCost = types.Int64Null()
	}
	if value := res.Get("ospfSettings.isPassiveEnabled"); value.Exists() && value.Value() != nil {
		data.OspfSettingsIsPassiveEnabled = types.BoolValue(value.Bool())
	} else {
		data.OspfSettingsIsPassiveEnabled = types.BoolNull()
	}
	if value := res.Get("ospfSettings.networkType"); value.Exists() && value.Value() != nil {
		data.OspfSettingsNetworkType = types.StringValue(value.String())
	} else {
		data.OspfSettingsNetworkType = types.StringNull()
	}
	if value := res.Get("vrf.name"); value.Exists() && value.Value() != nil {
		data.VrfName = types.StringValue(value.String())
	} else {
		data.VrfName = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchRoutingInterface) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("defaultGateway"); value.Exists() && !data.DefaultGateway.IsNull() {
		data.DefaultGateway = types.StringValue(value.String())
	} else {
		data.DefaultGateway = types.StringNull()
	}
	if value := res.Get("interfaceIp"); value.Exists() && !data.InterfaceIp.IsNull() {
		data.InterfaceIp = types.StringValue(value.String())
	} else {
		data.InterfaceIp = types.StringNull()
	}
	if value := res.Get("mode"); value.Exists() && !data.Mode.IsNull() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("multicastRouting"); value.Exists() && !data.MulticastRouting.IsNull() {
		data.MulticastRouting = types.StringValue(value.String())
	} else {
		data.MulticastRouting = types.StringNull()
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
	if value := res.Get("switchPortId"); value.Exists() && !data.SwitchPortId.IsNull() {
		data.SwitchPortId = types.StringValue(value.String())
	} else {
		data.SwitchPortId = types.StringNull()
	}
	if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("ipv6.address"); value.Exists() && !data.Ipv6Address.IsNull() {
		data.Ipv6Address = types.StringValue(value.String())
	} else {
		data.Ipv6Address = types.StringNull()
	}
	if value := res.Get("ipv6.assignmentMode"); value.Exists() && !data.Ipv6AssignmentMode.IsNull() {
		data.Ipv6AssignmentMode = types.StringValue(value.String())
	} else {
		data.Ipv6AssignmentMode = types.StringNull()
	}
	if value := res.Get("ipv6.gateway"); value.Exists() && !data.Ipv6Gateway.IsNull() {
		data.Ipv6Gateway = types.StringValue(value.String())
	} else {
		data.Ipv6Gateway = types.StringNull()
	}
	if value := res.Get("ipv6.prefix"); value.Exists() && !data.Ipv6Prefix.IsNull() {
		data.Ipv6Prefix = types.StringValue(value.String())
	} else {
		data.Ipv6Prefix = types.StringNull()
	}
	if value := res.Get("ospfSettings.area"); value.Exists() && !data.OspfSettingsArea.IsNull() {
		data.OspfSettingsArea = types.StringValue(value.String())
	} else {
		data.OspfSettingsArea = types.StringNull()
	}
	if value := res.Get("ospfSettings.cost"); value.Exists() && !data.OspfSettingsCost.IsNull() {
		data.OspfSettingsCost = types.Int64Value(value.Int())
	} else {
		data.OspfSettingsCost = types.Int64Null()
	}
	if value := res.Get("ospfSettings.isPassiveEnabled"); value.Exists() && !data.OspfSettingsIsPassiveEnabled.IsNull() {
		data.OspfSettingsIsPassiveEnabled = types.BoolValue(value.Bool())
	} else {
		data.OspfSettingsIsPassiveEnabled = types.BoolNull()
	}
	if value := res.Get("ospfSettings.networkType"); value.Exists() && !data.OspfSettingsNetworkType.IsNull() {
		data.OspfSettingsNetworkType = types.StringValue(value.String())
	} else {
		data.OspfSettingsNetworkType = types.StringNull()
	}
	if value := res.Get("vrf.name"); value.Exists() && !data.VrfName.IsNull() {
		data.VrfName = types.StringValue(value.String())
	} else {
		data.VrfName = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SwitchRoutingInterface) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *SwitchRoutingInterfaceIdentity) toIdentity(ctx context.Context, plan *SwitchRoutingInterface) {
	data.Serial = plan.Serial
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *SwitchRoutingInterface) fromIdentity(ctx context.Context, identity *SwitchRoutingInterfaceIdentity) {
	data.Serial = identity.Serial
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data SwitchRoutingInterface) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
