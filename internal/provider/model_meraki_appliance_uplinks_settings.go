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

type ApplianceUplinksSettings struct {
	Id                                         types.String `tfsdk:"id"`
	Serial                                     types.String `tfsdk:"serial"`
	InterfacesWan1Enabled                      types.Bool   `tfsdk:"interfaces_wan1_enabled"`
	InterfacesWan1PppoeEnabled                 types.Bool   `tfsdk:"interfaces_wan1_pppoe_enabled"`
	InterfacesWan1PppoeAuthenticationEnabled   types.Bool   `tfsdk:"interfaces_wan1_pppoe_authentication_enabled"`
	InterfacesWan1PppoeAuthenticationPassword  types.String `tfsdk:"interfaces_wan1_pppoe_authentication_password"`
	InterfacesWan1PppoeAuthenticationUsername  types.String `tfsdk:"interfaces_wan1_pppoe_authentication_username"`
	InterfacesWan1SvisIpv4Address              types.String `tfsdk:"interfaces_wan1_svis_ipv4_address"`
	InterfacesWan1SvisIpv4AssignmentMode       types.String `tfsdk:"interfaces_wan1_svis_ipv4_assignment_mode"`
	InterfacesWan1SvisIpv4Gateway              types.String `tfsdk:"interfaces_wan1_svis_ipv4_gateway"`
	InterfacesWan1SvisIpv4NameserversAddresses types.List   `tfsdk:"interfaces_wan1_svis_ipv4_nameservers_addresses"`
	InterfacesWan1SvisIpv6Address              types.String `tfsdk:"interfaces_wan1_svis_ipv6_address"`
	InterfacesWan1SvisIpv6AssignmentMode       types.String `tfsdk:"interfaces_wan1_svis_ipv6_assignment_mode"`
	InterfacesWan1SvisIpv6Gateway              types.String `tfsdk:"interfaces_wan1_svis_ipv6_gateway"`
	InterfacesWan1SvisIpv6NameserversAddresses types.List   `tfsdk:"interfaces_wan1_svis_ipv6_nameservers_addresses"`
	InterfacesWan1VlanTaggingEnabled           types.Bool   `tfsdk:"interfaces_wan1_vlan_tagging_enabled"`
	InterfacesWan1VlanTaggingVlanId            types.Int64  `tfsdk:"interfaces_wan1_vlan_tagging_vlan_id"`
	InterfacesWan2Enabled                      types.Bool   `tfsdk:"interfaces_wan2_enabled"`
	InterfacesWan2PppoeEnabled                 types.Bool   `tfsdk:"interfaces_wan2_pppoe_enabled"`
	InterfacesWan2PppoeAuthenticationEnabled   types.Bool   `tfsdk:"interfaces_wan2_pppoe_authentication_enabled"`
	InterfacesWan2PppoeAuthenticationPassword  types.String `tfsdk:"interfaces_wan2_pppoe_authentication_password"`
	InterfacesWan2PppoeAuthenticationUsername  types.String `tfsdk:"interfaces_wan2_pppoe_authentication_username"`
	InterfacesWan2SvisIpv4Address              types.String `tfsdk:"interfaces_wan2_svis_ipv4_address"`
	InterfacesWan2SvisIpv4AssignmentMode       types.String `tfsdk:"interfaces_wan2_svis_ipv4_assignment_mode"`
	InterfacesWan2SvisIpv4Gateway              types.String `tfsdk:"interfaces_wan2_svis_ipv4_gateway"`
	InterfacesWan2SvisIpv4NameserversAddresses types.List   `tfsdk:"interfaces_wan2_svis_ipv4_nameservers_addresses"`
	InterfacesWan2SvisIpv6Address              types.String `tfsdk:"interfaces_wan2_svis_ipv6_address"`
	InterfacesWan2SvisIpv6AssignmentMode       types.String `tfsdk:"interfaces_wan2_svis_ipv6_assignment_mode"`
	InterfacesWan2SvisIpv6Gateway              types.String `tfsdk:"interfaces_wan2_svis_ipv6_gateway"`
	InterfacesWan2SvisIpv6NameserversAddresses types.List   `tfsdk:"interfaces_wan2_svis_ipv6_nameservers_addresses"`
	InterfacesWan2VlanTaggingEnabled           types.Bool   `tfsdk:"interfaces_wan2_vlan_tagging_enabled"`
	InterfacesWan2VlanTaggingVlanId            types.Int64  `tfsdk:"interfaces_wan2_vlan_tagging_vlan_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceUplinksSettings) getPath() string {
	return fmt.Sprintf("/devices/%v/appliance/uplinks/settings", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceUplinksSettings) toBody(ctx context.Context, state ApplianceUplinksSettings) string {
	body := ""
	if !data.InterfacesWan1Enabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.enabled", data.InterfacesWan1Enabled.ValueBool())
	}
	if !data.InterfacesWan1PppoeEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.pppoe.enabled", data.InterfacesWan1PppoeEnabled.ValueBool())
	}
	if !data.InterfacesWan1PppoeAuthenticationEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.pppoe.authentication.enabled", data.InterfacesWan1PppoeAuthenticationEnabled.ValueBool())
	}
	if !data.InterfacesWan1PppoeAuthenticationPassword.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.pppoe.authentication.password", data.InterfacesWan1PppoeAuthenticationPassword.ValueString())
	}
	if !data.InterfacesWan1PppoeAuthenticationUsername.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.pppoe.authentication.username", data.InterfacesWan1PppoeAuthenticationUsername.ValueString())
	}
	if !data.InterfacesWan1SvisIpv4Address.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv4.address", data.InterfacesWan1SvisIpv4Address.ValueString())
	}
	if !data.InterfacesWan1SvisIpv4AssignmentMode.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv4.assignmentMode", data.InterfacesWan1SvisIpv4AssignmentMode.ValueString())
	}
	if !data.InterfacesWan1SvisIpv4Gateway.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv4.gateway", data.InterfacesWan1SvisIpv4Gateway.ValueString())
	}
	if !data.InterfacesWan1SvisIpv4NameserversAddresses.IsNull() {
		var values []string
		data.InterfacesWan1SvisIpv4NameserversAddresses.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv4.nameservers.addresses", values)
	}
	if !data.InterfacesWan1SvisIpv6Address.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv6.address", data.InterfacesWan1SvisIpv6Address.ValueString())
	}
	if !data.InterfacesWan1SvisIpv6AssignmentMode.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv6.assignmentMode", data.InterfacesWan1SvisIpv6AssignmentMode.ValueString())
	}
	if !data.InterfacesWan1SvisIpv6Gateway.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv6.gateway", data.InterfacesWan1SvisIpv6Gateway.ValueString())
	}
	if !data.InterfacesWan1SvisIpv6NameserversAddresses.IsNull() {
		var values []string
		data.InterfacesWan1SvisIpv6NameserversAddresses.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv6.nameservers.addresses", values)
	}
	if !data.InterfacesWan1VlanTaggingEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.vlanTagging.enabled", data.InterfacesWan1VlanTaggingEnabled.ValueBool())
	}
	if !data.InterfacesWan1VlanTaggingVlanId.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.vlanTagging.vlanId", data.InterfacesWan1VlanTaggingVlanId.ValueInt64())
	}
	if !data.InterfacesWan2Enabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.enabled", data.InterfacesWan2Enabled.ValueBool())
	}
	if !data.InterfacesWan2PppoeEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.pppoe.enabled", data.InterfacesWan2PppoeEnabled.ValueBool())
	}
	if !data.InterfacesWan2PppoeAuthenticationEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.pppoe.authentication.enabled", data.InterfacesWan2PppoeAuthenticationEnabled.ValueBool())
	}
	if !data.InterfacesWan2PppoeAuthenticationPassword.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.pppoe.authentication.password", data.InterfacesWan2PppoeAuthenticationPassword.ValueString())
	}
	if !data.InterfacesWan2PppoeAuthenticationUsername.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.pppoe.authentication.username", data.InterfacesWan2PppoeAuthenticationUsername.ValueString())
	}
	if !data.InterfacesWan2SvisIpv4Address.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv4.address", data.InterfacesWan2SvisIpv4Address.ValueString())
	}
	if !data.InterfacesWan2SvisIpv4AssignmentMode.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv4.assignmentMode", data.InterfacesWan2SvisIpv4AssignmentMode.ValueString())
	}
	if !data.InterfacesWan2SvisIpv4Gateway.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv4.gateway", data.InterfacesWan2SvisIpv4Gateway.ValueString())
	}
	if !data.InterfacesWan2SvisIpv4NameserversAddresses.IsNull() {
		var values []string
		data.InterfacesWan2SvisIpv4NameserversAddresses.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv4.nameservers.addresses", values)
	}
	if !data.InterfacesWan2SvisIpv6Address.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv6.address", data.InterfacesWan2SvisIpv6Address.ValueString())
	}
	if !data.InterfacesWan2SvisIpv6AssignmentMode.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv6.assignmentMode", data.InterfacesWan2SvisIpv6AssignmentMode.ValueString())
	}
	if !data.InterfacesWan2SvisIpv6Gateway.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv6.gateway", data.InterfacesWan2SvisIpv6Gateway.ValueString())
	}
	if !data.InterfacesWan2SvisIpv6NameserversAddresses.IsNull() {
		var values []string
		data.InterfacesWan2SvisIpv6NameserversAddresses.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv6.nameservers.addresses", values)
	}
	if !data.InterfacesWan2VlanTaggingEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.vlanTagging.enabled", data.InterfacesWan2VlanTaggingEnabled.ValueBool())
	}
	if !data.InterfacesWan2VlanTaggingVlanId.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.vlanTagging.vlanId", data.InterfacesWan2VlanTaggingVlanId.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceUplinksSettings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("interfaces.wan1.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1Enabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1Enabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.pppoe.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1PppoeEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1PppoeEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.pppoe.authentication.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1PppoeAuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1PppoeAuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.pppoe.authentication.username"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1PppoeAuthenticationUsername = types.StringValue(value.String())
	} else {
		data.InterfacesWan1PppoeAuthenticationUsername = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.address"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1SvisIpv4Address = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv4Address = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.assignmentMode"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1SvisIpv4AssignmentMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv4AssignmentMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.gateway"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1SvisIpv4Gateway = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv4Gateway = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.nameservers.addresses"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1SvisIpv4NameserversAddresses = helpers.GetStringList(value.Array())
	} else {
		data.InterfacesWan1SvisIpv4NameserversAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.address"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1SvisIpv6Address = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv6Address = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.assignmentMode"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1SvisIpv6AssignmentMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv6AssignmentMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.gateway"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1SvisIpv6Gateway = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv6Gateway = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.nameservers.addresses"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1SvisIpv6NameserversAddresses = helpers.GetStringList(value.Array())
	} else {
		data.InterfacesWan1SvisIpv6NameserversAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get("interfaces.wan1.vlanTagging.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1VlanTaggingEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1VlanTaggingEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.vlanTagging.vlanId"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1VlanTaggingVlanId = types.Int64Value(value.Int())
	} else {
		data.InterfacesWan1VlanTaggingVlanId = types.Int64Null()
	}
	if value := res.Get("interfaces.wan2.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2Enabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2Enabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.pppoe.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2PppoeEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2PppoeEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.pppoe.authentication.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2PppoeAuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2PppoeAuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.pppoe.authentication.username"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2PppoeAuthenticationUsername = types.StringValue(value.String())
	} else {
		data.InterfacesWan2PppoeAuthenticationUsername = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.address"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2SvisIpv4Address = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv4Address = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.assignmentMode"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2SvisIpv4AssignmentMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv4AssignmentMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.gateway"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2SvisIpv4Gateway = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv4Gateway = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.nameservers.addresses"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2SvisIpv4NameserversAddresses = helpers.GetStringList(value.Array())
	} else {
		data.InterfacesWan2SvisIpv4NameserversAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.address"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2SvisIpv6Address = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv6Address = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.assignmentMode"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2SvisIpv6AssignmentMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv6AssignmentMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.gateway"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2SvisIpv6Gateway = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv6Gateway = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.nameservers.addresses"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2SvisIpv6NameserversAddresses = helpers.GetStringList(value.Array())
	} else {
		data.InterfacesWan2SvisIpv6NameserversAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get("interfaces.wan2.vlanTagging.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2VlanTaggingEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2VlanTaggingEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.vlanTagging.vlanId"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2VlanTaggingVlanId = types.Int64Value(value.Int())
	} else {
		data.InterfacesWan2VlanTaggingVlanId = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceUplinksSettings) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("interfaces.wan1.enabled"); value.Exists() && !data.InterfacesWan1Enabled.IsNull() {
		data.InterfacesWan1Enabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1Enabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.pppoe.enabled"); value.Exists() && !data.InterfacesWan1PppoeEnabled.IsNull() {
		data.InterfacesWan1PppoeEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1PppoeEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.pppoe.authentication.enabled"); value.Exists() && !data.InterfacesWan1PppoeAuthenticationEnabled.IsNull() {
		data.InterfacesWan1PppoeAuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1PppoeAuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.pppoe.authentication.username"); value.Exists() && !data.InterfacesWan1PppoeAuthenticationUsername.IsNull() {
		data.InterfacesWan1PppoeAuthenticationUsername = types.StringValue(value.String())
	} else {
		data.InterfacesWan1PppoeAuthenticationUsername = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.address"); value.Exists() && !data.InterfacesWan1SvisIpv4Address.IsNull() {
		data.InterfacesWan1SvisIpv4Address = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv4Address = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.assignmentMode"); value.Exists() && !data.InterfacesWan1SvisIpv4AssignmentMode.IsNull() {
		data.InterfacesWan1SvisIpv4AssignmentMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv4AssignmentMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.gateway"); value.Exists() && !data.InterfacesWan1SvisIpv4Gateway.IsNull() {
		data.InterfacesWan1SvisIpv4Gateway = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv4Gateway = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.nameservers.addresses"); value.Exists() && !data.InterfacesWan1SvisIpv4NameserversAddresses.IsNull() {
		data.InterfacesWan1SvisIpv4NameserversAddresses = helpers.GetStringList(value.Array())
	} else {
		data.InterfacesWan1SvisIpv4NameserversAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.address"); value.Exists() && !data.InterfacesWan1SvisIpv6Address.IsNull() {
		data.InterfacesWan1SvisIpv6Address = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv6Address = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.assignmentMode"); value.Exists() && !data.InterfacesWan1SvisIpv6AssignmentMode.IsNull() {
		data.InterfacesWan1SvisIpv6AssignmentMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv6AssignmentMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.gateway"); value.Exists() && !data.InterfacesWan1SvisIpv6Gateway.IsNull() {
		data.InterfacesWan1SvisIpv6Gateway = types.StringValue(value.String())
	} else {
		data.InterfacesWan1SvisIpv6Gateway = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.nameservers.addresses"); value.Exists() && !data.InterfacesWan1SvisIpv6NameserversAddresses.IsNull() {
		data.InterfacesWan1SvisIpv6NameserversAddresses = helpers.GetStringList(value.Array())
	} else {
		data.InterfacesWan1SvisIpv6NameserversAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get("interfaces.wan1.vlanTagging.enabled"); value.Exists() && !data.InterfacesWan1VlanTaggingEnabled.IsNull() {
		data.InterfacesWan1VlanTaggingEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1VlanTaggingEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.vlanTagging.vlanId"); value.Exists() && !data.InterfacesWan1VlanTaggingVlanId.IsNull() {
		data.InterfacesWan1VlanTaggingVlanId = types.Int64Value(value.Int())
	} else {
		data.InterfacesWan1VlanTaggingVlanId = types.Int64Null()
	}
	if value := res.Get("interfaces.wan2.enabled"); value.Exists() && !data.InterfacesWan2Enabled.IsNull() {
		data.InterfacesWan2Enabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2Enabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.pppoe.enabled"); value.Exists() && !data.InterfacesWan2PppoeEnabled.IsNull() {
		data.InterfacesWan2PppoeEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2PppoeEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.pppoe.authentication.enabled"); value.Exists() && !data.InterfacesWan2PppoeAuthenticationEnabled.IsNull() {
		data.InterfacesWan2PppoeAuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2PppoeAuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.pppoe.authentication.username"); value.Exists() && !data.InterfacesWan2PppoeAuthenticationUsername.IsNull() {
		data.InterfacesWan2PppoeAuthenticationUsername = types.StringValue(value.String())
	} else {
		data.InterfacesWan2PppoeAuthenticationUsername = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.address"); value.Exists() && !data.InterfacesWan2SvisIpv4Address.IsNull() {
		data.InterfacesWan2SvisIpv4Address = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv4Address = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.assignmentMode"); value.Exists() && !data.InterfacesWan2SvisIpv4AssignmentMode.IsNull() {
		data.InterfacesWan2SvisIpv4AssignmentMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv4AssignmentMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.gateway"); value.Exists() && !data.InterfacesWan2SvisIpv4Gateway.IsNull() {
		data.InterfacesWan2SvisIpv4Gateway = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv4Gateway = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.nameservers.addresses"); value.Exists() && !data.InterfacesWan2SvisIpv4NameserversAddresses.IsNull() {
		data.InterfacesWan2SvisIpv4NameserversAddresses = helpers.GetStringList(value.Array())
	} else {
		data.InterfacesWan2SvisIpv4NameserversAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.address"); value.Exists() && !data.InterfacesWan2SvisIpv6Address.IsNull() {
		data.InterfacesWan2SvisIpv6Address = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv6Address = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.assignmentMode"); value.Exists() && !data.InterfacesWan2SvisIpv6AssignmentMode.IsNull() {
		data.InterfacesWan2SvisIpv6AssignmentMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv6AssignmentMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.gateway"); value.Exists() && !data.InterfacesWan2SvisIpv6Gateway.IsNull() {
		data.InterfacesWan2SvisIpv6Gateway = types.StringValue(value.String())
	} else {
		data.InterfacesWan2SvisIpv6Gateway = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.nameservers.addresses"); value.Exists() && !data.InterfacesWan2SvisIpv6NameserversAddresses.IsNull() {
		data.InterfacesWan2SvisIpv6NameserversAddresses = helpers.GetStringList(value.Array())
	} else {
		data.InterfacesWan2SvisIpv6NameserversAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get("interfaces.wan2.vlanTagging.enabled"); value.Exists() && !data.InterfacesWan2VlanTaggingEnabled.IsNull() {
		data.InterfacesWan2VlanTaggingEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2VlanTaggingEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.vlanTagging.vlanId"); value.Exists() && !data.InterfacesWan2VlanTaggingVlanId.IsNull() {
		data.InterfacesWan2VlanTaggingVlanId = types.Int64Value(value.Int())
	} else {
		data.InterfacesWan2VlanTaggingVlanId = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial
