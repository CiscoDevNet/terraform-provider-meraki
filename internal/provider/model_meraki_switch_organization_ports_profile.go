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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchOrganizationPortsProfile struct {
	Id                             types.String                                   `tfsdk:"id"`
	OrganizationId                 types.String                                   `tfsdk:"organization_id"`
	Description                    types.String                                   `tfsdk:"description"`
	IsOrganizationWide             types.Bool                                     `tfsdk:"is_organization_wide"`
	Name                           types.String                                   `tfsdk:"name"`
	NetworkId                      types.String                                   `tfsdk:"network_id"`
	NetworksType                   types.String                                   `tfsdk:"networks_type"`
	NetworksValues                 []SwitchOrganizationPortsProfileNetworksValues `tfsdk:"networks_values"`
	PortAccessPolicyNumber         types.Int64                                    `tfsdk:"port_access_policy_number"`
	PortAccessPolicyType           types.String                                   `tfsdk:"port_access_policy_type"`
	PortAdaptivePolicyGroupId      types.String                                   `tfsdk:"port_adaptive_policy_group_id"`
	PortAdaptivePolicyVoiceGroupId types.String                                   `tfsdk:"port_adaptive_policy_voice_group_id"`
	PortAllowedVlans               types.String                                   `tfsdk:"port_allowed_vlans"`
	PortDaiTrusted                 types.Bool                                     `tfsdk:"port_dai_trusted"`
	PortIsolationEnabled           types.Bool                                     `tfsdk:"port_isolation_enabled"`
	PortPeerSgtCapable             types.Bool                                     `tfsdk:"port_peer_sgt_capable"`
	PortPoeEnabled                 types.Bool                                     `tfsdk:"port_poe_enabled"`
	PortRstpEnabled                types.Bool                                     `tfsdk:"port_rstp_enabled"`
	PortStickyMacAllowListLimit    types.Int64                                    `tfsdk:"port_sticky_mac_allow_list_limit"`
	PortStormControlEnabled        types.Bool                                     `tfsdk:"port_storm_control_enabled"`
	PortStpGuard                   types.String                                   `tfsdk:"port_stp_guard"`
	PortType                       types.String                                   `tfsdk:"port_type"`
	PortUdld                       types.String                                   `tfsdk:"port_udld"`
	PortVlan                       types.Int64                                    `tfsdk:"port_vlan"`
	PortVoiceVlan                  types.Int64                                    `tfsdk:"port_voice_vlan"`
	PortMacAllowList               types.List                                     `tfsdk:"port_mac_allow_list"`
	PortStickyMacAllowList         types.List                                     `tfsdk:"port_sticky_mac_allow_list"`
	Tags                           types.List                                     `tfsdk:"tags"`
}

type SwitchOrganizationPortsProfileNetworksValues struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type SwitchOrganizationPortsProfileIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchOrganizationPortsProfile) getPath() string {
	return fmt.Sprintf("/organizations/%v/switch/ports/profiles", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchOrganizationPortsProfile) toBody(ctx context.Context, state SwitchOrganizationPortsProfile) string {
	body := ""
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.IsOrganizationWide.IsNull() {
		body, _ = sjson.Set(body, "isOrganizationWide", data.IsOrganizationWide.ValueBool())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.NetworkId.IsNull() {
		body, _ = sjson.Set(body, "networkId", data.NetworkId.ValueString())
	}
	if !data.NetworksType.IsNull() {
		body, _ = sjson.Set(body, "networks.type", data.NetworksType.ValueString())
	}
	if data.NetworksValues != nil {
		body, _ = sjson.Set(body, "networks.values", []interface{}{})
		for _, item := range data.NetworksValues {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "networks.values.-1", itemBody)
		}
	}
	if !data.PortAccessPolicyNumber.IsNull() {
		body, _ = sjson.Set(body, "port.accessPolicyNumber", data.PortAccessPolicyNumber.ValueInt64())
	}
	if !data.PortAccessPolicyType.IsNull() {
		body, _ = sjson.Set(body, "port.accessPolicyType", data.PortAccessPolicyType.ValueString())
	}
	if !data.PortAdaptivePolicyGroupId.IsNull() {
		body, _ = sjson.Set(body, "port.adaptivePolicyGroupId", data.PortAdaptivePolicyGroupId.ValueString())
	}
	if !data.PortAdaptivePolicyVoiceGroupId.IsNull() {
		body, _ = sjson.Set(body, "port.adaptivePolicyVoiceGroupId", data.PortAdaptivePolicyVoiceGroupId.ValueString())
	}
	if !data.PortAllowedVlans.IsNull() {
		body, _ = sjson.Set(body, "port.allowedVlans", data.PortAllowedVlans.ValueString())
	}
	if !data.PortDaiTrusted.IsNull() {
		body, _ = sjson.Set(body, "port.daiTrusted", data.PortDaiTrusted.ValueBool())
	}
	if !data.PortIsolationEnabled.IsNull() {
		body, _ = sjson.Set(body, "port.isolationEnabled", data.PortIsolationEnabled.ValueBool())
	}
	if !data.PortPeerSgtCapable.IsNull() {
		body, _ = sjson.Set(body, "port.peerSgtCapable", data.PortPeerSgtCapable.ValueBool())
	}
	if !data.PortPoeEnabled.IsNull() {
		body, _ = sjson.Set(body, "port.poeEnabled", data.PortPoeEnabled.ValueBool())
	}
	if !data.PortRstpEnabled.IsNull() {
		body, _ = sjson.Set(body, "port.rstpEnabled", data.PortRstpEnabled.ValueBool())
	}
	if !data.PortStickyMacAllowListLimit.IsNull() {
		body, _ = sjson.Set(body, "port.stickyMacAllowListLimit", data.PortStickyMacAllowListLimit.ValueInt64())
	}
	if !data.PortStormControlEnabled.IsNull() {
		body, _ = sjson.Set(body, "port.stormControlEnabled", data.PortStormControlEnabled.ValueBool())
	}
	if !data.PortStpGuard.IsNull() {
		body, _ = sjson.Set(body, "port.stpGuard", data.PortStpGuard.ValueString())
	}
	if !data.PortType.IsNull() {
		body, _ = sjson.Set(body, "port.type", data.PortType.ValueString())
	}
	if !data.PortUdld.IsNull() {
		body, _ = sjson.Set(body, "port.udld", data.PortUdld.ValueString())
	}
	if !data.PortVlan.IsNull() {
		body, _ = sjson.Set(body, "port.vlan", data.PortVlan.ValueInt64())
	}
	if !data.PortVoiceVlan.IsNull() {
		body, _ = sjson.Set(body, "port.voiceVlan", data.PortVoiceVlan.ValueInt64())
	}
	if !data.PortMacAllowList.IsNull() {
		var values []string
		data.PortMacAllowList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "port.macAllowList", values)
	}
	if !data.PortStickyMacAllowList.IsNull() {
		var values []string
		data.PortStickyMacAllowList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "port.stickyMacAllowList", values)
	}
	if !data.Tags.IsNull() {
		var values []string
		data.Tags.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "tags", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchOrganizationPortsProfile) fromBody(ctx context.Context, res meraki.Res) {
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
		data.NetworksValues = make([]SwitchOrganizationPortsProfileNetworksValues, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchOrganizationPortsProfileNetworksValues{}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchOrganizationPortsProfile) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("isOrganizationWide"); value.Exists() && !data.IsOrganizationWide.IsNull() {
		data.IsOrganizationWide = types.BoolValue(value.Bool())
	} else {
		data.IsOrganizationWide = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("networkId"); value.Exists() && !data.NetworkId.IsNull() {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
	}
	if value := res.Get("networks.type"); value.Exists() && !data.NetworksType.IsNull() {
		data.NetworksType = types.StringValue(value.String())
	} else {
		data.NetworksType = types.StringNull()
	}
	for i := 0; i < len(data.NetworksValues); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.NetworksValues[i].Id.ValueString()}

		parent := &data
		data := (*parent).NetworksValues[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("networks.values").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing NetworksValues[%d] = %+v",
				i,
				(*parent).NetworksValues[i],
			))
			(*parent).NetworksValues = slices.Delete((*parent).NetworksValues, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).NetworksValues[i] = data
	}
	if value := res.Get("port.accessPolicyNumber"); value.Exists() && !data.PortAccessPolicyNumber.IsNull() {
		data.PortAccessPolicyNumber = types.Int64Value(value.Int())
	} else {
		data.PortAccessPolicyNumber = types.Int64Null()
	}
	if value := res.Get("port.accessPolicyType"); value.Exists() && !data.PortAccessPolicyType.IsNull() {
		data.PortAccessPolicyType = types.StringValue(value.String())
	} else {
		data.PortAccessPolicyType = types.StringNull()
	}
	if value := res.Get("port.adaptivePolicyGroupId"); value.Exists() && !data.PortAdaptivePolicyGroupId.IsNull() {
		data.PortAdaptivePolicyGroupId = types.StringValue(value.String())
	} else {
		data.PortAdaptivePolicyGroupId = types.StringNull()
	}
	if value := res.Get("port.adaptivePolicyVoiceGroupId"); value.Exists() && !data.PortAdaptivePolicyVoiceGroupId.IsNull() {
		data.PortAdaptivePolicyVoiceGroupId = types.StringValue(value.String())
	} else {
		data.PortAdaptivePolicyVoiceGroupId = types.StringNull()
	}
	if value := res.Get("port.allowedVlans"); value.Exists() && !data.PortAllowedVlans.IsNull() {
		data.PortAllowedVlans = types.StringValue(value.String())
	} else {
		data.PortAllowedVlans = types.StringNull()
	}
	if value := res.Get("port.daiTrusted"); value.Exists() && !data.PortDaiTrusted.IsNull() {
		data.PortDaiTrusted = types.BoolValue(value.Bool())
	} else {
		data.PortDaiTrusted = types.BoolNull()
	}
	if value := res.Get("port.isolationEnabled"); value.Exists() && !data.PortIsolationEnabled.IsNull() {
		data.PortIsolationEnabled = types.BoolValue(value.Bool())
	} else {
		data.PortIsolationEnabled = types.BoolNull()
	}
	if value := res.Get("port.peerSgtCapable"); value.Exists() && !data.PortPeerSgtCapable.IsNull() {
		data.PortPeerSgtCapable = types.BoolValue(value.Bool())
	} else {
		data.PortPeerSgtCapable = types.BoolNull()
	}
	if value := res.Get("port.poeEnabled"); value.Exists() && !data.PortPoeEnabled.IsNull() {
		data.PortPoeEnabled = types.BoolValue(value.Bool())
	} else {
		data.PortPoeEnabled = types.BoolNull()
	}
	if value := res.Get("port.rstpEnabled"); value.Exists() && !data.PortRstpEnabled.IsNull() {
		data.PortRstpEnabled = types.BoolValue(value.Bool())
	} else {
		data.PortRstpEnabled = types.BoolNull()
	}
	if value := res.Get("port.stickyMacAllowListLimit"); value.Exists() && !data.PortStickyMacAllowListLimit.IsNull() {
		data.PortStickyMacAllowListLimit = types.Int64Value(value.Int())
	} else {
		data.PortStickyMacAllowListLimit = types.Int64Null()
	}
	if value := res.Get("port.stormControlEnabled"); value.Exists() && !data.PortStormControlEnabled.IsNull() {
		data.PortStormControlEnabled = types.BoolValue(value.Bool())
	} else {
		data.PortStormControlEnabled = types.BoolNull()
	}
	if value := res.Get("port.stpGuard"); value.Exists() && !data.PortStpGuard.IsNull() {
		data.PortStpGuard = types.StringValue(value.String())
	} else {
		data.PortStpGuard = types.StringNull()
	}
	if value := res.Get("port.type"); value.Exists() && !data.PortType.IsNull() {
		data.PortType = types.StringValue(value.String())
	} else {
		data.PortType = types.StringNull()
	}
	if value := res.Get("port.udld"); value.Exists() && !data.PortUdld.IsNull() {
		data.PortUdld = types.StringValue(value.String())
	} else {
		data.PortUdld = types.StringNull()
	}
	if value := res.Get("port.vlan"); value.Exists() && !data.PortVlan.IsNull() {
		data.PortVlan = types.Int64Value(value.Int())
	} else {
		data.PortVlan = types.Int64Null()
	}
	if value := res.Get("port.voiceVlan"); value.Exists() && !data.PortVoiceVlan.IsNull() {
		data.PortVoiceVlan = types.Int64Value(value.Int())
	} else {
		data.PortVoiceVlan = types.Int64Null()
	}
	if value := res.Get("port.macAllowList"); value.Exists() && !data.PortMacAllowList.IsNull() {
		data.PortMacAllowList = helpers.GetStringList(value.Array())
	} else {
		data.PortMacAllowList = types.ListNull(types.StringType)
	}
	if value := res.Get("port.stickyMacAllowList"); value.Exists() && !data.PortStickyMacAllowList.IsNull() {
		data.PortStickyMacAllowList = helpers.GetStringList(value.Array())
	} else {
		data.PortStickyMacAllowList = types.ListNull(types.StringType)
	}
	if value := res.Get("tags"); value.Exists() && !data.Tags.IsNull() {
		data.Tags = helpers.GetStringList(value.Array())
	} else {
		data.Tags = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SwitchOrganizationPortsProfile) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *SwitchOrganizationPortsProfileIdentity) toIdentity(ctx context.Context, plan *SwitchOrganizationPortsProfile) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *SwitchOrganizationPortsProfile) fromIdentity(ctx context.Context, identity *SwitchOrganizationPortsProfileIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data SwitchOrganizationPortsProfile) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
