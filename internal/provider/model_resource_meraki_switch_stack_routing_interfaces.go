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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceSwitchStackRoutingInterfaces struct {
	Id             types.String                                `tfsdk:"id"`
	OrganizationId types.String                                `tfsdk:"organization_id"`
	NetworkId      types.String                                `tfsdk:"network_id"`
	SwitchStackId  types.String                                `tfsdk:"switch_stack_id"`
	Items          []ResourceSwitchStackRoutingInterfacesItems `tfsdk:"items"`
}

type ResourceSwitchStackRoutingInterfacesItems struct {
	Id                           types.String `tfsdk:"id"`
	CandidateUplinkV4            types.Bool   `tfsdk:"candidate_uplink_v4"`
	DefaultGateway               types.String `tfsdk:"default_gateway"`
	InterfaceIp                  types.String `tfsdk:"interface_ip"`
	IsSwitchDefaultGateway       types.Bool   `tfsdk:"is_switch_default_gateway"`
	Mode                         types.String `tfsdk:"mode"`
	MulticastRouting             types.String `tfsdk:"multicast_routing"`
	Name                         types.String `tfsdk:"name"`
	StaticV4Dns1                 types.String `tfsdk:"static_v4_dns1"`
	StaticV4Dns2                 types.String `tfsdk:"static_v4_dns2"`
	Subnet                       types.String `tfsdk:"subnet"`
	SwitchPortId                 types.String `tfsdk:"switch_port_id"`
	UplinkV4                     types.Bool   `tfsdk:"uplink_v4"`
	UplinkV6                     types.Bool   `tfsdk:"uplink_v6"`
	VlanId                       types.Int64  `tfsdk:"vlan_id"`
	Ipv6Address                  types.String `tfsdk:"ipv6_address"`
	Ipv6AssignmentMode           types.String `tfsdk:"ipv6_assignment_mode"`
	Ipv6CandidateUplink          types.Bool   `tfsdk:"ipv6_candidate_uplink"`
	Ipv6Gateway                  types.String `tfsdk:"ipv6_gateway"`
	Ipv6IsSwitchDefaultGateway   types.Bool   `tfsdk:"ipv6_is_switch_default_gateway"`
	Ipv6Prefix                   types.String `tfsdk:"ipv6_prefix"`
	Ipv6StaticV6Dns1             types.String `tfsdk:"ipv6_static_v6_dns1"`
	Ipv6StaticV6Dns2             types.String `tfsdk:"ipv6_static_v6_dns2"`
	OspfSettingsArea             types.String `tfsdk:"ospf_settings_area"`
	OspfSettingsCost             types.Int64  `tfsdk:"ospf_settings_cost"`
	OspfSettingsIsPassiveEnabled types.Bool   `tfsdk:"ospf_settings_is_passive_enabled"`
	OspfSettingsNetworkType      types.String `tfsdk:"ospf_settings_network_type"`
	VrfName                      types.String `tfsdk:"vrf_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceSwitchStackRoutingInterfaces) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/stacks/%v/routing/interfaces", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.SwitchStackId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceSwitchStackRoutingInterfacesItems) toBody(ctx context.Context, state ResourceSwitchStackRoutingInterfacesItems) string {
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

func (data *ResourceSwitchStackRoutingInterfaces) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceSwitchStackRoutingInterfacesItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceSwitchStackRoutingInterfacesItems{}
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
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	index := 0
	res.ForEach(func(k, res gjson.Result) bool {
		data.Items[index].Id = types.StringValue(res.Get("interfaceId").String())
		index++
		return true
	})
	data.Id = data.OrganizationId
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceSwitchStackRoutingInterfaces) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("interfaceId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
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
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyPartial(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceSwitchStackRoutingInterfaces) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceSwitchStackRoutingInterfaces) fromBodyImport(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("interfaceId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
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
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyImport(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceSwitchStackRoutingInterfaces) hasChanges(ctx context.Context, state *ResourceSwitchStackRoutingInterfaces, id string) bool {
	hasChanges := false

	item := ResourceSwitchStackRoutingInterfacesItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceSwitchStackRoutingInterfacesItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.CandidateUplinkV4.Equal(stateItem.CandidateUplinkV4) {
		hasChanges = true
	}
	if !item.DefaultGateway.Equal(stateItem.DefaultGateway) {
		hasChanges = true
	}
	if !item.InterfaceIp.Equal(stateItem.InterfaceIp) {
		hasChanges = true
	}
	if !item.IsSwitchDefaultGateway.Equal(stateItem.IsSwitchDefaultGateway) {
		hasChanges = true
	}
	if !item.Mode.Equal(stateItem.Mode) {
		hasChanges = true
	}
	if !item.MulticastRouting.Equal(stateItem.MulticastRouting) {
		hasChanges = true
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	if !item.StaticV4Dns1.Equal(stateItem.StaticV4Dns1) {
		hasChanges = true
	}
	if !item.StaticV4Dns2.Equal(stateItem.StaticV4Dns2) {
		hasChanges = true
	}
	if !item.Subnet.Equal(stateItem.Subnet) {
		hasChanges = true
	}
	if !item.SwitchPortId.Equal(stateItem.SwitchPortId) {
		hasChanges = true
	}
	if !item.UplinkV4.Equal(stateItem.UplinkV4) {
		hasChanges = true
	}
	if !item.UplinkV6.Equal(stateItem.UplinkV6) {
		hasChanges = true
	}
	if !item.VlanId.Equal(stateItem.VlanId) {
		hasChanges = true
	}
	if !item.Ipv6Address.Equal(stateItem.Ipv6Address) {
		hasChanges = true
	}
	if !item.Ipv6AssignmentMode.Equal(stateItem.Ipv6AssignmentMode) {
		hasChanges = true
	}
	if !item.Ipv6CandidateUplink.Equal(stateItem.Ipv6CandidateUplink) {
		hasChanges = true
	}
	if !item.Ipv6Gateway.Equal(stateItem.Ipv6Gateway) {
		hasChanges = true
	}
	if !item.Ipv6IsSwitchDefaultGateway.Equal(stateItem.Ipv6IsSwitchDefaultGateway) {
		hasChanges = true
	}
	if !item.Ipv6Prefix.Equal(stateItem.Ipv6Prefix) {
		hasChanges = true
	}
	if !item.Ipv6StaticV6Dns1.Equal(stateItem.Ipv6StaticV6Dns1) {
		hasChanges = true
	}
	if !item.Ipv6StaticV6Dns2.Equal(stateItem.Ipv6StaticV6Dns2) {
		hasChanges = true
	}
	if !item.OspfSettingsArea.Equal(stateItem.OspfSettingsArea) {
		hasChanges = true
	}
	if !item.OspfSettingsCost.Equal(stateItem.OspfSettingsCost) {
		hasChanges = true
	}
	if !item.OspfSettingsIsPassiveEnabled.Equal(stateItem.OspfSettingsIsPassiveEnabled) {
		hasChanges = true
	}
	if !item.OspfSettingsNetworkType.Equal(stateItem.OspfSettingsNetworkType) {
		hasChanges = true
	}
	if !item.VrfName.Equal(stateItem.VrfName) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
