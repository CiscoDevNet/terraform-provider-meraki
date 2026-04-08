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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceSwitchOrganizationPortsProfiles struct {
	OrganizationId types.String                                     `tfsdk:"organization_id"`
	Items          []DataSourceSwitchOrganizationPortsProfilesItems `tfsdk:"items"`
}

type DataSourceSwitchOrganizationPortsProfilesItems struct {
	Id                             types.String                                              `tfsdk:"id"`
	Description                    types.String                                              `tfsdk:"description"`
	IsOrganizationWide             types.Bool                                                `tfsdk:"is_organization_wide"`
	Name                           types.String                                              `tfsdk:"name"`
	NetworkId                      types.String                                              `tfsdk:"network_id"`
	NetworksType                   types.String                                              `tfsdk:"networks_type"`
	NetworksValues                 []DataSourceSwitchOrganizationPortsProfilesNetworksValues `tfsdk:"networks_values"`
	PortAccessPolicyNumber         types.Int64                                               `tfsdk:"port_access_policy_number"`
	PortAccessPolicyType           types.String                                              `tfsdk:"port_access_policy_type"`
	PortAdaptivePolicyGroupId      types.String                                              `tfsdk:"port_adaptive_policy_group_id"`
	PortAdaptivePolicyVoiceGroupId types.String                                              `tfsdk:"port_adaptive_policy_voice_group_id"`
	PortAllowedVlans               types.String                                              `tfsdk:"port_allowed_vlans"`
	PortDaiTrusted                 types.Bool                                                `tfsdk:"port_dai_trusted"`
	PortIsolationEnabled           types.Bool                                                `tfsdk:"port_isolation_enabled"`
	PortPeerSgtCapable             types.Bool                                                `tfsdk:"port_peer_sgt_capable"`
	PortPoeEnabled                 types.Bool                                                `tfsdk:"port_poe_enabled"`
	PortRstpEnabled                types.Bool                                                `tfsdk:"port_rstp_enabled"`
	PortStickyMacAllowListLimit    types.Int64                                               `tfsdk:"port_sticky_mac_allow_list_limit"`
	PortStormControlEnabled        types.Bool                                                `tfsdk:"port_storm_control_enabled"`
	PortStpGuard                   types.String                                              `tfsdk:"port_stp_guard"`
	PortType                       types.String                                              `tfsdk:"port_type"`
	PortUdld                       types.String                                              `tfsdk:"port_udld"`
	PortVlan                       types.Int64                                               `tfsdk:"port_vlan"`
	PortVoiceVlan                  types.Int64                                               `tfsdk:"port_voice_vlan"`
	PortMacAllowList               types.List                                                `tfsdk:"port_mac_allow_list"`
	PortStickyMacAllowList         types.List                                                `tfsdk:"port_sticky_mac_allow_list"`
	Tags                           types.List                                                `tfsdk:"tags"`
}

type DataSourceSwitchOrganizationPortsProfilesNetworksValues struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceSwitchOrganizationPortsProfiles) getPath() string {
	return fmt.Sprintf("/organizations/%v/switch/ports/profiles", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceSwitchOrganizationPortsProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceSwitchOrganizationPortsProfilesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceSwitchOrganizationPortsProfilesItems{}
		data.Id = types.StringValue(res.Get("profileId").String())
		if value := res.Get("description"); value.Exists() && value.Value() != nil {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("isOrganizationWide"); value.Exists() && value.Value() != nil {
			data.IsOrganizationWide = types.BoolValue(value.Bool())
		} else {
			data.IsOrganizationWide = types.BoolNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("networkId"); value.Exists() && value.Value() != nil {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("networks.type"); value.Exists() && value.Value() != nil {
			data.NetworksType = types.StringValue(value.String())
		} else {
			data.NetworksType = types.StringNull()
		}
		if value := res.Get("networks.values"); value.Exists() && value.Value() != nil {
			data.NetworksValues = make([]DataSourceSwitchOrganizationPortsProfilesNetworksValues, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceSwitchOrganizationPortsProfilesNetworksValues{}
				if value := res.Get("id"); value.Exists() && value.Value() != nil {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				(*parent).NetworksValues = append((*parent).NetworksValues, data)
				return true
			})
		}
		if value := res.Get("port.accessPolicyNumber"); value.Exists() && value.Value() != nil {
			data.PortAccessPolicyNumber = types.Int64Value(value.Int())
		} else {
			data.PortAccessPolicyNumber = types.Int64Null()
		}
		if value := res.Get("port.accessPolicyType"); value.Exists() && value.Value() != nil {
			data.PortAccessPolicyType = types.StringValue(value.String())
		} else {
			data.PortAccessPolicyType = types.StringNull()
		}
		if value := res.Get("port.adaptivePolicyGroupId"); value.Exists() && value.Value() != nil {
			data.PortAdaptivePolicyGroupId = types.StringValue(value.String())
		} else {
			data.PortAdaptivePolicyGroupId = types.StringNull()
		}
		if value := res.Get("port.adaptivePolicyVoiceGroupId"); value.Exists() && value.Value() != nil {
			data.PortAdaptivePolicyVoiceGroupId = types.StringValue(value.String())
		} else {
			data.PortAdaptivePolicyVoiceGroupId = types.StringNull()
		}
		if value := res.Get("port.allowedVlans"); value.Exists() && value.Value() != nil {
			data.PortAllowedVlans = types.StringValue(value.String())
		} else {
			data.PortAllowedVlans = types.StringNull()
		}
		if value := res.Get("port.daiTrusted"); value.Exists() && value.Value() != nil {
			data.PortDaiTrusted = types.BoolValue(value.Bool())
		} else {
			data.PortDaiTrusted = types.BoolNull()
		}
		if value := res.Get("port.isolationEnabled"); value.Exists() && value.Value() != nil {
			data.PortIsolationEnabled = types.BoolValue(value.Bool())
		} else {
			data.PortIsolationEnabled = types.BoolNull()
		}
		if value := res.Get("port.peerSgtCapable"); value.Exists() && value.Value() != nil {
			data.PortPeerSgtCapable = types.BoolValue(value.Bool())
		} else {
			data.PortPeerSgtCapable = types.BoolNull()
		}
		if value := res.Get("port.poeEnabled"); value.Exists() && value.Value() != nil {
			data.PortPoeEnabled = types.BoolValue(value.Bool())
		} else {
			data.PortPoeEnabled = types.BoolNull()
		}
		if value := res.Get("port.rstpEnabled"); value.Exists() && value.Value() != nil {
			data.PortRstpEnabled = types.BoolValue(value.Bool())
		} else {
			data.PortRstpEnabled = types.BoolNull()
		}
		if value := res.Get("port.stickyMacAllowListLimit"); value.Exists() && value.Value() != nil {
			data.PortStickyMacAllowListLimit = types.Int64Value(value.Int())
		} else {
			data.PortStickyMacAllowListLimit = types.Int64Null()
		}
		if value := res.Get("port.stormControlEnabled"); value.Exists() && value.Value() != nil {
			data.PortStormControlEnabled = types.BoolValue(value.Bool())
		} else {
			data.PortStormControlEnabled = types.BoolNull()
		}
		if value := res.Get("port.stpGuard"); value.Exists() && value.Value() != nil {
			data.PortStpGuard = types.StringValue(value.String())
		} else {
			data.PortStpGuard = types.StringNull()
		}
		if value := res.Get("port.type"); value.Exists() && value.Value() != nil {
			data.PortType = types.StringValue(value.String())
		} else {
			data.PortType = types.StringNull()
		}
		if value := res.Get("port.udld"); value.Exists() && value.Value() != nil {
			data.PortUdld = types.StringValue(value.String())
		} else {
			data.PortUdld = types.StringNull()
		}
		if value := res.Get("port.vlan"); value.Exists() && value.Value() != nil {
			data.PortVlan = types.Int64Value(value.Int())
		} else {
			data.PortVlan = types.Int64Null()
		}
		if value := res.Get("port.voiceVlan"); value.Exists() && value.Value() != nil {
			data.PortVoiceVlan = types.Int64Value(value.Int())
		} else {
			data.PortVoiceVlan = types.Int64Null()
		}
		if value := res.Get("port.macAllowList"); value.Exists() && value.Value() != nil {
			data.PortMacAllowList = helpers.GetStringList(value.Array())
		} else {
			data.PortMacAllowList = types.ListNull(types.StringType)
		}
		if value := res.Get("port.stickyMacAllowList"); value.Exists() && value.Value() != nil {
			data.PortStickyMacAllowList = helpers.GetStringList(value.Array())
		} else {
			data.PortStickyMacAllowList = types.ListNull(types.StringType)
		}
		if value := res.Get("tags"); value.Exists() && value.Value() != nil {
			data.Tags = helpers.GetStringList(value.Array())
		} else {
			data.Tags = types.ListNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
