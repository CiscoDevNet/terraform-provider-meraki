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
	Id                                                          types.String `tfsdk:"id"`
	Serial                                                      types.String `tfsdk:"serial"`
	CandidateUplinkV4                                           types.Bool   `tfsdk:"candidate_uplink_v4"`
	DefaultGateway                                              types.String `tfsdk:"default_gateway"`
	InterfaceIp                                                 types.String `tfsdk:"interface_ip"`
	IsSwitchDefaultGateway                                      types.Bool   `tfsdk:"is_switch_default_gateway"`
	Mode                                                        types.String `tfsdk:"mode"`
	Mtu                                                         types.Int64  `tfsdk:"mtu"`
	MulticastRouting                                            types.String `tfsdk:"multicast_routing"`
	Name                                                        types.String `tfsdk:"name"`
	StaticV4Dns1                                                types.String `tfsdk:"static_v4_dns1"`
	StaticV4Dns2                                                types.String `tfsdk:"static_v4_dns2"`
	Subnet                                                      types.String `tfsdk:"subnet"`
	SwitchPortId                                                types.String `tfsdk:"switch_port_id"`
	UplinkV4                                                    types.Bool   `tfsdk:"uplink_v4"`
	UplinkV6                                                    types.Bool   `tfsdk:"uplink_v6"`
	VlanId                                                      types.Int64  `tfsdk:"vlan_id"`
	BidirectionalForwardingDetectionEnabled                     types.Bool   `tfsdk:"bidirectional_forwarding_detection_enabled"`
	BidirectionalForwardingDetectionIntervalInterfaceMinRx      types.Int64  `tfsdk:"bidirectional_forwarding_detection_interval_interface_min_rx"`
	BidirectionalForwardingDetectionIntervalInterfaceMsecs      types.Int64  `tfsdk:"bidirectional_forwarding_detection_interval_interface_msecs"`
	BidirectionalForwardingDetectionIntervalInterfaceMultiplier types.Int64  `tfsdk:"bidirectional_forwarding_detection_interval_interface_multiplier"`
	CloudUplinkEnabled                                          types.Bool   `tfsdk:"cloud_uplink_enabled"`
	CloudUplinkIsDefaultGateway                                 types.Bool   `tfsdk:"cloud_uplink_is_default_gateway"`
	CloudUplinkIsEligible                                       types.Bool   `tfsdk:"cloud_uplink_is_eligible"`
	CloudUplinkNameserversPrimary                               types.String `tfsdk:"cloud_uplink_nameservers_primary"`
	CloudUplinkNameserversSecondary                             types.String `tfsdk:"cloud_uplink_nameservers_secondary"`
	Ipv6Address                                                 types.String `tfsdk:"ipv6_address"`
	Ipv6AssignmentMode                                          types.String `tfsdk:"ipv6_assignment_mode"`
	Ipv6CandidateUplink                                         types.Bool   `tfsdk:"ipv6_candidate_uplink"`
	Ipv6Gateway                                                 types.String `tfsdk:"ipv6_gateway"`
	Ipv6IsSwitchDefaultGateway                                  types.Bool   `tfsdk:"ipv6_is_switch_default_gateway"`
	Ipv6Prefix                                                  types.String `tfsdk:"ipv6_prefix"`
	Ipv6StaticV6Dns1                                            types.String `tfsdk:"ipv6_static_v6_dns1"`
	Ipv6StaticV6Dns2                                            types.String `tfsdk:"ipv6_static_v6_dns2"`
	Ipv6CloudUplinkEnabled                                      types.Bool   `tfsdk:"ipv6_cloud_uplink_enabled"`
	Ipv6CloudUplinkIsDefaultGateway                             types.Bool   `tfsdk:"ipv6_cloud_uplink_is_default_gateway"`
	Ipv6CloudUplinkIsEligible                                   types.Bool   `tfsdk:"ipv6_cloud_uplink_is_eligible"`
	Ipv6CloudUplinkNameserversPrimary                           types.String `tfsdk:"ipv6_cloud_uplink_nameservers_primary"`
	Ipv6CloudUplinkNameserversSecondary                         types.String `tfsdk:"ipv6_cloud_uplink_nameservers_secondary"`
	OspfSettingsArea                                            types.String `tfsdk:"ospf_settings_area"`
	OspfSettingsCost                                            types.Int64  `tfsdk:"ospf_settings_cost"`
	OspfSettingsIsPassiveEnabled                                types.Bool   `tfsdk:"ospf_settings_is_passive_enabled"`
	OspfSettingsNetworkType                                     types.String `tfsdk:"ospf_settings_network_type"`
	VrfName                                                     types.String `tfsdk:"vrf_name"`
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
	if !data.CandidateUplinkV4.IsNull() {
		body, _ = sjson.Set(body, "candidateUplinkV4", data.CandidateUplinkV4.ValueBool())
	}
	if !data.DefaultGateway.IsNull() {
		body, _ = sjson.Set(body, "defaultGateway", data.DefaultGateway.ValueString())
	}
	if !data.InterfaceIp.IsNull() {
		body, _ = sjson.Set(body, "interfaceIp", data.InterfaceIp.ValueString())
	}
	if !data.IsSwitchDefaultGateway.IsNull() {
		body, _ = sjson.Set(body, "isSwitchDefaultGateway", data.IsSwitchDefaultGateway.ValueBool())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	if !data.Mtu.IsNull() {
		body, _ = sjson.Set(body, "mtu", data.Mtu.ValueInt64())
	}
	if !data.MulticastRouting.IsNull() {
		body, _ = sjson.Set(body, "multicastRouting", data.MulticastRouting.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.StaticV4Dns1.IsNull() {
		body, _ = sjson.Set(body, "staticV4Dns1", data.StaticV4Dns1.ValueString())
	}
	if !data.StaticV4Dns2.IsNull() {
		body, _ = sjson.Set(body, "staticV4Dns2", data.StaticV4Dns2.ValueString())
	}
	if !data.Subnet.IsNull() {
		body, _ = sjson.Set(body, "subnet", data.Subnet.ValueString())
	}
	if !data.SwitchPortId.IsNull() {
		body, _ = sjson.Set(body, "switchPortId", data.SwitchPortId.ValueString())
	}
	if !data.UplinkV4.IsNull() {
		body, _ = sjson.Set(body, "uplinkV4", data.UplinkV4.ValueBool())
	}
	if !data.UplinkV6.IsNull() {
		body, _ = sjson.Set(body, "uplinkV6", data.UplinkV6.ValueBool())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanId", data.VlanId.ValueInt64())
	}
	if !data.BidirectionalForwardingDetectionEnabled.IsNull() {
		body, _ = sjson.Set(body, "bidirectionalForwardingDetection.enabled", data.BidirectionalForwardingDetectionEnabled.ValueBool())
	}
	if !data.BidirectionalForwardingDetectionIntervalInterfaceMinRx.IsNull() {
		body, _ = sjson.Set(body, "bidirectionalForwardingDetection.intervalInterface.minRx", data.BidirectionalForwardingDetectionIntervalInterfaceMinRx.ValueInt64())
	}
	if !data.BidirectionalForwardingDetectionIntervalInterfaceMsecs.IsNull() {
		body, _ = sjson.Set(body, "bidirectionalForwardingDetection.intervalInterface.msecs", data.BidirectionalForwardingDetectionIntervalInterfaceMsecs.ValueInt64())
	}
	if !data.BidirectionalForwardingDetectionIntervalInterfaceMultiplier.IsNull() {
		body, _ = sjson.Set(body, "bidirectionalForwardingDetection.intervalInterface.multiplier", data.BidirectionalForwardingDetectionIntervalInterfaceMultiplier.ValueInt64())
	}
	if !data.CloudUplinkEnabled.IsNull() {
		body, _ = sjson.Set(body, "cloud.uplink.enabled", data.CloudUplinkEnabled.ValueBool())
	}
	if !data.CloudUplinkIsDefaultGateway.IsNull() {
		body, _ = sjson.Set(body, "cloud.uplink.isDefaultGateway", data.CloudUplinkIsDefaultGateway.ValueBool())
	}
	if !data.CloudUplinkIsEligible.IsNull() {
		body, _ = sjson.Set(body, "cloud.uplink.isEligible", data.CloudUplinkIsEligible.ValueBool())
	}
	if !data.CloudUplinkNameserversPrimary.IsNull() {
		body, _ = sjson.Set(body, "cloud.uplink.nameservers.primary", data.CloudUplinkNameserversPrimary.ValueString())
	}
	if !data.CloudUplinkNameserversSecondary.IsNull() {
		body, _ = sjson.Set(body, "cloud.uplink.nameservers.secondary", data.CloudUplinkNameserversSecondary.ValueString())
	}
	if !data.Ipv6Address.IsNull() {
		body, _ = sjson.Set(body, "ipv6.address", data.Ipv6Address.ValueString())
	}
	if !data.Ipv6AssignmentMode.IsNull() {
		body, _ = sjson.Set(body, "ipv6.assignmentMode", data.Ipv6AssignmentMode.ValueString())
	}
	if !data.Ipv6CandidateUplink.IsNull() {
		body, _ = sjson.Set(body, "ipv6.candidateUplink", data.Ipv6CandidateUplink.ValueBool())
	}
	if !data.Ipv6Gateway.IsNull() {
		body, _ = sjson.Set(body, "ipv6.gateway", data.Ipv6Gateway.ValueString())
	}
	if !data.Ipv6IsSwitchDefaultGateway.IsNull() {
		body, _ = sjson.Set(body, "ipv6.isSwitchDefaultGateway", data.Ipv6IsSwitchDefaultGateway.ValueBool())
	}
	if !data.Ipv6Prefix.IsNull() {
		body, _ = sjson.Set(body, "ipv6.prefix", data.Ipv6Prefix.ValueString())
	}
	if !data.Ipv6StaticV6Dns1.IsNull() {
		body, _ = sjson.Set(body, "ipv6.staticV6Dns1", data.Ipv6StaticV6Dns1.ValueString())
	}
	if !data.Ipv6StaticV6Dns2.IsNull() {
		body, _ = sjson.Set(body, "ipv6.staticV6Dns2", data.Ipv6StaticV6Dns2.ValueString())
	}
	if !data.Ipv6CloudUplinkEnabled.IsNull() {
		body, _ = sjson.Set(body, "ipv6.cloud.uplink.enabled", data.Ipv6CloudUplinkEnabled.ValueBool())
	}
	if !data.Ipv6CloudUplinkIsDefaultGateway.IsNull() {
		body, _ = sjson.Set(body, "ipv6.cloud.uplink.isDefaultGateway", data.Ipv6CloudUplinkIsDefaultGateway.ValueBool())
	}
	if !data.Ipv6CloudUplinkIsEligible.IsNull() {
		body, _ = sjson.Set(body, "ipv6.cloud.uplink.isEligible", data.Ipv6CloudUplinkIsEligible.ValueBool())
	}
	if !data.Ipv6CloudUplinkNameserversPrimary.IsNull() {
		body, _ = sjson.Set(body, "ipv6.cloud.uplink.nameservers.primary", data.Ipv6CloudUplinkNameserversPrimary.ValueString())
	}
	if !data.Ipv6CloudUplinkNameserversSecondary.IsNull() {
		body, _ = sjson.Set(body, "ipv6.cloud.uplink.nameservers.secondary", data.Ipv6CloudUplinkNameserversSecondary.ValueString())
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
	if value := res.Get("candidateUplinkV4"); value.Exists() && value.Value() != nil {
		data.CandidateUplinkV4 = types.BoolValue(value.Bool())
	} else {
		data.CandidateUplinkV4 = types.BoolNull()
	}
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
	if value := res.Get("isSwitchDefaultGateway"); value.Exists() && value.Value() != nil {
		data.IsSwitchDefaultGateway = types.BoolValue(value.Bool())
	} else {
		data.IsSwitchDefaultGateway = types.BoolNull()
	}
	if value := res.Get("mode"); value.Exists() && value.Value() != nil {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("mtu"); value.Exists() && value.Value() != nil {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
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
	if value := res.Get("staticV4Dns1"); value.Exists() && value.Value() != nil {
		data.StaticV4Dns1 = types.StringValue(value.String())
	} else {
		data.StaticV4Dns1 = types.StringNull()
	}
	if value := res.Get("staticV4Dns2"); value.Exists() && value.Value() != nil {
		data.StaticV4Dns2 = types.StringValue(value.String())
	} else {
		data.StaticV4Dns2 = types.StringNull()
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
	if value := res.Get("uplinkV4"); value.Exists() && value.Value() != nil {
		data.UplinkV4 = types.BoolValue(value.Bool())
	} else {
		data.UplinkV4 = types.BoolNull()
	}
	if value := res.Get("uplinkV6"); value.Exists() && value.Value() != nil {
		data.UplinkV6 = types.BoolValue(value.Bool())
	} else {
		data.UplinkV6 = types.BoolNull()
	}
	if value := res.Get("vlanId"); value.Exists() && value.Value() != nil {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("bidirectionalForwardingDetection.enabled"); value.Exists() && value.Value() != nil {
		data.BidirectionalForwardingDetectionEnabled = types.BoolValue(value.Bool())
	} else {
		data.BidirectionalForwardingDetectionEnabled = types.BoolNull()
	}
	if value := res.Get("bidirectionalForwardingDetection.intervalInterface.minRx"); value.Exists() && value.Value() != nil {
		data.BidirectionalForwardingDetectionIntervalInterfaceMinRx = types.Int64Value(value.Int())
	} else {
		data.BidirectionalForwardingDetectionIntervalInterfaceMinRx = types.Int64Null()
	}
	if value := res.Get("bidirectionalForwardingDetection.intervalInterface.msecs"); value.Exists() && value.Value() != nil {
		data.BidirectionalForwardingDetectionIntervalInterfaceMsecs = types.Int64Value(value.Int())
	} else {
		data.BidirectionalForwardingDetectionIntervalInterfaceMsecs = types.Int64Null()
	}
	if value := res.Get("bidirectionalForwardingDetection.intervalInterface.multiplier"); value.Exists() && value.Value() != nil {
		data.BidirectionalForwardingDetectionIntervalInterfaceMultiplier = types.Int64Value(value.Int())
	} else {
		data.BidirectionalForwardingDetectionIntervalInterfaceMultiplier = types.Int64Null()
	}
	if value := res.Get("cloud.uplink.enabled"); value.Exists() && value.Value() != nil {
		data.CloudUplinkEnabled = types.BoolValue(value.Bool())
	} else {
		data.CloudUplinkEnabled = types.BoolNull()
	}
	if value := res.Get("cloud.uplink.isDefaultGateway"); value.Exists() && value.Value() != nil {
		data.CloudUplinkIsDefaultGateway = types.BoolValue(value.Bool())
	} else {
		data.CloudUplinkIsDefaultGateway = types.BoolNull()
	}
	if value := res.Get("cloud.uplink.isEligible"); value.Exists() && value.Value() != nil {
		data.CloudUplinkIsEligible = types.BoolValue(value.Bool())
	} else {
		data.CloudUplinkIsEligible = types.BoolNull()
	}
	if value := res.Get("cloud.uplink.nameservers.primary"); value.Exists() && value.Value() != nil {
		data.CloudUplinkNameserversPrimary = types.StringValue(value.String())
	} else {
		data.CloudUplinkNameserversPrimary = types.StringNull()
	}
	if value := res.Get("cloud.uplink.nameservers.secondary"); value.Exists() && value.Value() != nil {
		data.CloudUplinkNameserversSecondary = types.StringValue(value.String())
	} else {
		data.CloudUplinkNameserversSecondary = types.StringNull()
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
	if value := res.Get("ipv6.candidateUplink"); value.Exists() && value.Value() != nil {
		data.Ipv6CandidateUplink = types.BoolValue(value.Bool())
	} else {
		data.Ipv6CandidateUplink = types.BoolNull()
	}
	if value := res.Get("ipv6.gateway"); value.Exists() && value.Value() != nil {
		data.Ipv6Gateway = types.StringValue(value.String())
	} else {
		data.Ipv6Gateway = types.StringNull()
	}
	if value := res.Get("ipv6.isSwitchDefaultGateway"); value.Exists() && value.Value() != nil {
		data.Ipv6IsSwitchDefaultGateway = types.BoolValue(value.Bool())
	} else {
		data.Ipv6IsSwitchDefaultGateway = types.BoolNull()
	}
	if value := res.Get("ipv6.prefix"); value.Exists() && value.Value() != nil {
		data.Ipv6Prefix = types.StringValue(value.String())
	} else {
		data.Ipv6Prefix = types.StringNull()
	}
	if value := res.Get("ipv6.staticV6Dns1"); value.Exists() && value.Value() != nil {
		data.Ipv6StaticV6Dns1 = types.StringValue(value.String())
	} else {
		data.Ipv6StaticV6Dns1 = types.StringNull()
	}
	if value := res.Get("ipv6.staticV6Dns2"); value.Exists() && value.Value() != nil {
		data.Ipv6StaticV6Dns2 = types.StringValue(value.String())
	} else {
		data.Ipv6StaticV6Dns2 = types.StringNull()
	}
	if value := res.Get("ipv6.cloud.uplink.enabled"); value.Exists() && value.Value() != nil {
		data.Ipv6CloudUplinkEnabled = types.BoolValue(value.Bool())
	} else {
		data.Ipv6CloudUplinkEnabled = types.BoolNull()
	}
	if value := res.Get("ipv6.cloud.uplink.isDefaultGateway"); value.Exists() && value.Value() != nil {
		data.Ipv6CloudUplinkIsDefaultGateway = types.BoolValue(value.Bool())
	} else {
		data.Ipv6CloudUplinkIsDefaultGateway = types.BoolNull()
	}
	if value := res.Get("ipv6.cloud.uplink.isEligible"); value.Exists() && value.Value() != nil {
		data.Ipv6CloudUplinkIsEligible = types.BoolValue(value.Bool())
	} else {
		data.Ipv6CloudUplinkIsEligible = types.BoolNull()
	}
	if value := res.Get("ipv6.cloud.uplink.nameservers.primary"); value.Exists() && value.Value() != nil {
		data.Ipv6CloudUplinkNameserversPrimary = types.StringValue(value.String())
	} else {
		data.Ipv6CloudUplinkNameserversPrimary = types.StringNull()
	}
	if value := res.Get("ipv6.cloud.uplink.nameservers.secondary"); value.Exists() && value.Value() != nil {
		data.Ipv6CloudUplinkNameserversSecondary = types.StringValue(value.String())
	} else {
		data.Ipv6CloudUplinkNameserversSecondary = types.StringNull()
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
	if value := res.Get("candidateUplinkV4"); value.Exists() && !data.CandidateUplinkV4.IsNull() {
		data.CandidateUplinkV4 = types.BoolValue(value.Bool())
	} else {
		data.CandidateUplinkV4 = types.BoolNull()
	}
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
	if value := res.Get("isSwitchDefaultGateway"); value.Exists() && !data.IsSwitchDefaultGateway.IsNull() {
		data.IsSwitchDefaultGateway = types.BoolValue(value.Bool())
	} else {
		data.IsSwitchDefaultGateway = types.BoolNull()
	}
	if value := res.Get("mode"); value.Exists() && !data.Mode.IsNull() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("mtu"); value.Exists() && !data.Mtu.IsNull() {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Null()
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
	if value := res.Get("staticV4Dns1"); value.Exists() && !data.StaticV4Dns1.IsNull() {
		data.StaticV4Dns1 = types.StringValue(value.String())
	} else {
		data.StaticV4Dns1 = types.StringNull()
	}
	if value := res.Get("staticV4Dns2"); value.Exists() && !data.StaticV4Dns2.IsNull() {
		data.StaticV4Dns2 = types.StringValue(value.String())
	} else {
		data.StaticV4Dns2 = types.StringNull()
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
	if value := res.Get("uplinkV4"); value.Exists() && !data.UplinkV4.IsNull() {
		data.UplinkV4 = types.BoolValue(value.Bool())
	} else {
		data.UplinkV4 = types.BoolNull()
	}
	if value := res.Get("uplinkV6"); value.Exists() && !data.UplinkV6.IsNull() {
		data.UplinkV6 = types.BoolValue(value.Bool())
	} else {
		data.UplinkV6 = types.BoolNull()
	}
	if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("bidirectionalForwardingDetection.enabled"); value.Exists() && !data.BidirectionalForwardingDetectionEnabled.IsNull() {
		data.BidirectionalForwardingDetectionEnabled = types.BoolValue(value.Bool())
	} else {
		data.BidirectionalForwardingDetectionEnabled = types.BoolNull()
	}
	if value := res.Get("bidirectionalForwardingDetection.intervalInterface.minRx"); value.Exists() && !data.BidirectionalForwardingDetectionIntervalInterfaceMinRx.IsNull() {
		data.BidirectionalForwardingDetectionIntervalInterfaceMinRx = types.Int64Value(value.Int())
	} else {
		data.BidirectionalForwardingDetectionIntervalInterfaceMinRx = types.Int64Null()
	}
	if value := res.Get("bidirectionalForwardingDetection.intervalInterface.msecs"); value.Exists() && !data.BidirectionalForwardingDetectionIntervalInterfaceMsecs.IsNull() {
		data.BidirectionalForwardingDetectionIntervalInterfaceMsecs = types.Int64Value(value.Int())
	} else {
		data.BidirectionalForwardingDetectionIntervalInterfaceMsecs = types.Int64Null()
	}
	if value := res.Get("bidirectionalForwardingDetection.intervalInterface.multiplier"); value.Exists() && !data.BidirectionalForwardingDetectionIntervalInterfaceMultiplier.IsNull() {
		data.BidirectionalForwardingDetectionIntervalInterfaceMultiplier = types.Int64Value(value.Int())
	} else {
		data.BidirectionalForwardingDetectionIntervalInterfaceMultiplier = types.Int64Null()
	}
	if value := res.Get("cloud.uplink.enabled"); value.Exists() && !data.CloudUplinkEnabled.IsNull() {
		data.CloudUplinkEnabled = types.BoolValue(value.Bool())
	} else {
		data.CloudUplinkEnabled = types.BoolNull()
	}
	if value := res.Get("cloud.uplink.isDefaultGateway"); value.Exists() && !data.CloudUplinkIsDefaultGateway.IsNull() {
		data.CloudUplinkIsDefaultGateway = types.BoolValue(value.Bool())
	} else {
		data.CloudUplinkIsDefaultGateway = types.BoolNull()
	}
	if value := res.Get("cloud.uplink.isEligible"); value.Exists() && !data.CloudUplinkIsEligible.IsNull() {
		data.CloudUplinkIsEligible = types.BoolValue(value.Bool())
	} else {
		data.CloudUplinkIsEligible = types.BoolNull()
	}
	if value := res.Get("cloud.uplink.nameservers.primary"); value.Exists() && !data.CloudUplinkNameserversPrimary.IsNull() {
		data.CloudUplinkNameserversPrimary = types.StringValue(value.String())
	} else {
		data.CloudUplinkNameserversPrimary = types.StringNull()
	}
	if value := res.Get("cloud.uplink.nameservers.secondary"); value.Exists() && !data.CloudUplinkNameserversSecondary.IsNull() {
		data.CloudUplinkNameserversSecondary = types.StringValue(value.String())
	} else {
		data.CloudUplinkNameserversSecondary = types.StringNull()
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
	if value := res.Get("ipv6.candidateUplink"); value.Exists() && !data.Ipv6CandidateUplink.IsNull() {
		data.Ipv6CandidateUplink = types.BoolValue(value.Bool())
	} else {
		data.Ipv6CandidateUplink = types.BoolNull()
	}
	if value := res.Get("ipv6.gateway"); value.Exists() && !data.Ipv6Gateway.IsNull() {
		data.Ipv6Gateway = types.StringValue(value.String())
	} else {
		data.Ipv6Gateway = types.StringNull()
	}
	if value := res.Get("ipv6.isSwitchDefaultGateway"); value.Exists() && !data.Ipv6IsSwitchDefaultGateway.IsNull() {
		data.Ipv6IsSwitchDefaultGateway = types.BoolValue(value.Bool())
	} else {
		data.Ipv6IsSwitchDefaultGateway = types.BoolNull()
	}
	if value := res.Get("ipv6.prefix"); value.Exists() && !data.Ipv6Prefix.IsNull() {
		data.Ipv6Prefix = types.StringValue(value.String())
	} else {
		data.Ipv6Prefix = types.StringNull()
	}
	if value := res.Get("ipv6.staticV6Dns1"); value.Exists() && !data.Ipv6StaticV6Dns1.IsNull() {
		data.Ipv6StaticV6Dns1 = types.StringValue(value.String())
	} else {
		data.Ipv6StaticV6Dns1 = types.StringNull()
	}
	if value := res.Get("ipv6.staticV6Dns2"); value.Exists() && !data.Ipv6StaticV6Dns2.IsNull() {
		data.Ipv6StaticV6Dns2 = types.StringValue(value.String())
	} else {
		data.Ipv6StaticV6Dns2 = types.StringNull()
	}
	if value := res.Get("ipv6.cloud.uplink.enabled"); value.Exists() && !data.Ipv6CloudUplinkEnabled.IsNull() {
		data.Ipv6CloudUplinkEnabled = types.BoolValue(value.Bool())
	} else {
		data.Ipv6CloudUplinkEnabled = types.BoolNull()
	}
	if value := res.Get("ipv6.cloud.uplink.isDefaultGateway"); value.Exists() && !data.Ipv6CloudUplinkIsDefaultGateway.IsNull() {
		data.Ipv6CloudUplinkIsDefaultGateway = types.BoolValue(value.Bool())
	} else {
		data.Ipv6CloudUplinkIsDefaultGateway = types.BoolNull()
	}
	if value := res.Get("ipv6.cloud.uplink.isEligible"); value.Exists() && !data.Ipv6CloudUplinkIsEligible.IsNull() {
		data.Ipv6CloudUplinkIsEligible = types.BoolValue(value.Bool())
	} else {
		data.Ipv6CloudUplinkIsEligible = types.BoolNull()
	}
	if value := res.Get("ipv6.cloud.uplink.nameservers.primary"); value.Exists() && !data.Ipv6CloudUplinkNameserversPrimary.IsNull() {
		data.Ipv6CloudUplinkNameserversPrimary = types.StringValue(value.String())
	} else {
		data.Ipv6CloudUplinkNameserversPrimary = types.StringNull()
	}
	if value := res.Get("ipv6.cloud.uplink.nameservers.secondary"); value.Exists() && !data.Ipv6CloudUplinkNameserversSecondary.IsNull() {
		data.Ipv6CloudUplinkNameserversSecondary = types.StringValue(value.String())
	} else {
		data.Ipv6CloudUplinkNameserversSecondary = types.StringNull()
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
