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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceSwitchStackRoutingInterfaces struct {
	NetworkId     types.String                                  `tfsdk:"network_id"`
	SwitchStackId types.String                                  `tfsdk:"switch_stack_id"`
	Items         []DataSourceSwitchStackRoutingInterfacesItems `tfsdk:"items"`
}

type DataSourceSwitchStackRoutingInterfacesItems struct {
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

func (data DataSourceSwitchStackRoutingInterfaces) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/stacks/%v/routing/interfaces", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.SwitchStackId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceSwitchStackRoutingInterfaces) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceSwitchStackRoutingInterfacesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceSwitchStackRoutingInterfacesItems{}
		data.Id = types.StringValue(res.Get("interfaceId").String())
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
}

// End of section. //template:end fromBody
