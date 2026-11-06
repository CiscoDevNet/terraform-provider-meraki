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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceUplinksSettings struct {
	Id                                                 types.String                                                      `tfsdk:"id"`
	Serial                                             types.String                                                      `tfsdk:"serial"`
	InterfacesWan1Enabled                              types.Bool                                                        `tfsdk:"interfaces_wan1_enabled"`
	InterfacesWan1PppoeEnabled                         types.Bool                                                        `tfsdk:"interfaces_wan1_pppoe_enabled"`
	InterfacesWan1PppoeAuthenticationEnabled           types.Bool                                                        `tfsdk:"interfaces_wan1_pppoe_authentication_enabled"`
	InterfacesWan1PppoeAuthenticationPassword          types.String                                                      `tfsdk:"interfaces_wan1_pppoe_authentication_password"`
	InterfacesWan1PppoeAuthenticationPasswordWo        types.String                                                      `tfsdk:"interfaces_wan1_pppoe_authentication_password_wo"`
	InterfacesWan1PppoeAuthenticationPasswordWoVersion types.Int64                                                       `tfsdk:"interfaces_wan1_pppoe_authentication_password_wo_version"`
	InterfacesWan1PppoeAuthenticationUsername          types.String                                                      `tfsdk:"interfaces_wan1_pppoe_authentication_username"`
	InterfacesWan1SvisIpv4Address                      types.String                                                      `tfsdk:"interfaces_wan1_svis_ipv4_address"`
	InterfacesWan1SvisIpv4AssignmentMode               types.String                                                      `tfsdk:"interfaces_wan1_svis_ipv4_assignment_mode"`
	InterfacesWan1SvisIpv4Gateway                      types.String                                                      `tfsdk:"interfaces_wan1_svis_ipv4_gateway"`
	InterfacesWan1SvisIpv4NameserversAddresses         types.List                                                        `tfsdk:"interfaces_wan1_svis_ipv4_nameservers_addresses"`
	InterfacesWan1SvisIpv6Address                      types.String                                                      `tfsdk:"interfaces_wan1_svis_ipv6_address"`
	InterfacesWan1SvisIpv6AssignmentMode               types.String                                                      `tfsdk:"interfaces_wan1_svis_ipv6_assignment_mode"`
	InterfacesWan1SvisIpv6Gateway                      types.String                                                      `tfsdk:"interfaces_wan1_svis_ipv6_gateway"`
	InterfacesWan1SvisIpv6NameserversAddresses         types.List                                                        `tfsdk:"interfaces_wan1_svis_ipv6_nameservers_addresses"`
	InterfacesWan1TransitionEnabled                    types.Bool                                                        `tfsdk:"interfaces_wan1_transition_enabled"`
	InterfacesWan1TransitionMode                       types.String                                                      `tfsdk:"interfaces_wan1_transition_mode"`
	InterfacesWan1TransitionAuthenticationEnabled      types.Bool                                                        `tfsdk:"interfaces_wan1_transition_authentication_enabled"`
	InterfacesWan1TransitionAuthenticationPassword     types.String                                                      `tfsdk:"interfaces_wan1_transition_authentication_password"`
	InterfacesWan1TransitionAuthenticationUsername     types.String                                                      `tfsdk:"interfaces_wan1_transition_authentication_username"`
	InterfacesWan1TransitionIpv4LocalAddress           types.String                                                      `tfsdk:"interfaces_wan1_transition_ipv4_local_address"`
	InterfacesWan1TransitionIpv6InterfaceId            types.String                                                      `tfsdk:"interfaces_wan1_transition_ipv6_interface_id"`
	InterfacesWan1TransitionIpv6BorderRelayAddress     types.String                                                      `tfsdk:"interfaces_wan1_transition_ipv6_border_relay_address"`
	InterfacesWan1TransitionIpv6BorderRelayMode        types.String                                                      `tfsdk:"interfaces_wan1_transition_ipv6_border_relay_mode"`
	InterfacesWan1TransitionIspName                    types.String                                                      `tfsdk:"interfaces_wan1_transition_isp_name"`
	InterfacesWan1TransitionMapDraftEnabled            types.Bool                                                        `tfsdk:"interfaces_wan1_transition_map_draft_enabled"`
	InterfacesWan1TransitionMapRulesSource             types.String                                                      `tfsdk:"interfaces_wan1_transition_map_rules_source"`
	InterfacesWan1TransitionMapRulesEntries            []ApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries `tfsdk:"interfaces_wan1_transition_map_rules_entries"`
	InterfacesWan1VlanTaggingEnabled                   types.Bool                                                        `tfsdk:"interfaces_wan1_vlan_tagging_enabled"`
	InterfacesWan1VlanTaggingVlanId                    types.Int64                                                       `tfsdk:"interfaces_wan1_vlan_tagging_vlan_id"`
	InterfacesWan2Enabled                              types.Bool                                                        `tfsdk:"interfaces_wan2_enabled"`
	InterfacesWan2PppoeEnabled                         types.Bool                                                        `tfsdk:"interfaces_wan2_pppoe_enabled"`
	InterfacesWan2PppoeAuthenticationEnabled           types.Bool                                                        `tfsdk:"interfaces_wan2_pppoe_authentication_enabled"`
	InterfacesWan2PppoeAuthenticationPassword          types.String                                                      `tfsdk:"interfaces_wan2_pppoe_authentication_password"`
	InterfacesWan2PppoeAuthenticationPasswordWo        types.String                                                      `tfsdk:"interfaces_wan2_pppoe_authentication_password_wo"`
	InterfacesWan2PppoeAuthenticationPasswordWoVersion types.Int64                                                       `tfsdk:"interfaces_wan2_pppoe_authentication_password_wo_version"`
	InterfacesWan2PppoeAuthenticationUsername          types.String                                                      `tfsdk:"interfaces_wan2_pppoe_authentication_username"`
	InterfacesWan2SvisIpv4Address                      types.String                                                      `tfsdk:"interfaces_wan2_svis_ipv4_address"`
	InterfacesWan2SvisIpv4AssignmentMode               types.String                                                      `tfsdk:"interfaces_wan2_svis_ipv4_assignment_mode"`
	InterfacesWan2SvisIpv4Gateway                      types.String                                                      `tfsdk:"interfaces_wan2_svis_ipv4_gateway"`
	InterfacesWan2SvisIpv4NameserversAddresses         types.List                                                        `tfsdk:"interfaces_wan2_svis_ipv4_nameservers_addresses"`
	InterfacesWan2SvisIpv6Address                      types.String                                                      `tfsdk:"interfaces_wan2_svis_ipv6_address"`
	InterfacesWan2SvisIpv6AssignmentMode               types.String                                                      `tfsdk:"interfaces_wan2_svis_ipv6_assignment_mode"`
	InterfacesWan2SvisIpv6Gateway                      types.String                                                      `tfsdk:"interfaces_wan2_svis_ipv6_gateway"`
	InterfacesWan2SvisIpv6NameserversAddresses         types.List                                                        `tfsdk:"interfaces_wan2_svis_ipv6_nameservers_addresses"`
	InterfacesWan2TransitionEnabled                    types.Bool                                                        `tfsdk:"interfaces_wan2_transition_enabled"`
	InterfacesWan2TransitionMode                       types.String                                                      `tfsdk:"interfaces_wan2_transition_mode"`
	InterfacesWan2TransitionAuthenticationEnabled      types.Bool                                                        `tfsdk:"interfaces_wan2_transition_authentication_enabled"`
	InterfacesWan2TransitionAuthenticationPassword     types.String                                                      `tfsdk:"interfaces_wan2_transition_authentication_password"`
	InterfacesWan2TransitionAuthenticationUsername     types.String                                                      `tfsdk:"interfaces_wan2_transition_authentication_username"`
	InterfacesWan2TransitionIpv4LocalAddress           types.String                                                      `tfsdk:"interfaces_wan2_transition_ipv4_local_address"`
	InterfacesWan2TransitionIpv6InterfaceId            types.String                                                      `tfsdk:"interfaces_wan2_transition_ipv6_interface_id"`
	InterfacesWan2TransitionIpv6BorderRelayAddress     types.String                                                      `tfsdk:"interfaces_wan2_transition_ipv6_border_relay_address"`
	InterfacesWan2TransitionIpv6BorderRelayMode        types.String                                                      `tfsdk:"interfaces_wan2_transition_ipv6_border_relay_mode"`
	InterfacesWan2TransitionIspName                    types.String                                                      `tfsdk:"interfaces_wan2_transition_isp_name"`
	InterfacesWan2TransitionMapDraftEnabled            types.Bool                                                        `tfsdk:"interfaces_wan2_transition_map_draft_enabled"`
	InterfacesWan2TransitionMapRulesSource             types.String                                                      `tfsdk:"interfaces_wan2_transition_map_rules_source"`
	InterfacesWan2TransitionMapRulesEntries            []ApplianceUplinksSettingsInterfacesWan2TransitionMapRulesEntries `tfsdk:"interfaces_wan2_transition_map_rules_entries"`
	InterfacesWan2VlanTaggingEnabled                   types.Bool                                                        `tfsdk:"interfaces_wan2_vlan_tagging_enabled"`
	InterfacesWan2VlanTaggingVlanId                    types.Int64                                                       `tfsdk:"interfaces_wan2_vlan_tagging_vlan_id"`
}

type ApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries struct {
	EaBitLength        types.Int64  `tfsdk:"ea_bit_length"`
	PsIdOffset         types.Int64  `tfsdk:"ps_id_offset"`
	BorderRelayAddress types.String `tfsdk:"border_relay_address"`
	BorderRelayPrefix  types.String `tfsdk:"border_relay_prefix"`
	Ipv4PrefixAddress  types.String `tfsdk:"ipv4_prefix_address"`
	Ipv4PrefixLength   types.Int64  `tfsdk:"ipv4_prefix_length"`
	Ipv6PrefixAddress  types.String `tfsdk:"ipv6_prefix_address"`
	Ipv6PrefixLength   types.Int64  `tfsdk:"ipv6_prefix_length"`
}

type ApplianceUplinksSettingsInterfacesWan2TransitionMapRulesEntries struct {
	EaBitLength        types.Int64  `tfsdk:"ea_bit_length"`
	PsIdOffset         types.Int64  `tfsdk:"ps_id_offset"`
	BorderRelayAddress types.String `tfsdk:"border_relay_address"`
	BorderRelayPrefix  types.String `tfsdk:"border_relay_prefix"`
	Ipv4PrefixAddress  types.String `tfsdk:"ipv4_prefix_address"`
	Ipv4PrefixLength   types.Int64  `tfsdk:"ipv4_prefix_length"`
	Ipv6PrefixAddress  types.String `tfsdk:"ipv6_prefix_address"`
	Ipv6PrefixLength   types.Int64  `tfsdk:"ipv6_prefix_length"`
}

type ApplianceUplinksSettingsIdentity struct {
	Serial types.String `tfsdk:"serial"`
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
	if !data.InterfacesWan1PppoeAuthenticationPasswordWo.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.pppoe.authentication.password", data.InterfacesWan1PppoeAuthenticationPasswordWo.ValueString())
	} else if !data.InterfacesWan1PppoeAuthenticationPassword.IsNull() {
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
	if !data.InterfacesWan1TransitionEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.enabled", data.InterfacesWan1TransitionEnabled.ValueBool())
	}
	if !data.InterfacesWan1TransitionMode.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.mode", data.InterfacesWan1TransitionMode.ValueString())
	}
	if !data.InterfacesWan1TransitionAuthenticationEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.authentication.enabled", data.InterfacesWan1TransitionAuthenticationEnabled.ValueBool())
	}
	if !data.InterfacesWan1TransitionAuthenticationPassword.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.authentication.password", data.InterfacesWan1TransitionAuthenticationPassword.ValueString())
	}
	if !data.InterfacesWan1TransitionAuthenticationUsername.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.authentication.username", data.InterfacesWan1TransitionAuthenticationUsername.ValueString())
	}
	if !data.InterfacesWan1TransitionIpv4LocalAddress.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.ipv4.local.address", data.InterfacesWan1TransitionIpv4LocalAddress.ValueString())
	}
	if !data.InterfacesWan1TransitionIpv6InterfaceId.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.ipv6.interfaceId", data.InterfacesWan1TransitionIpv6InterfaceId.ValueString())
	}
	if !data.InterfacesWan1TransitionIpv6BorderRelayAddress.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.ipv6.borderRelay.address", data.InterfacesWan1TransitionIpv6BorderRelayAddress.ValueString())
	}
	if !data.InterfacesWan1TransitionIpv6BorderRelayMode.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.ipv6.borderRelay.mode", data.InterfacesWan1TransitionIpv6BorderRelayMode.ValueString())
	}
	if !data.InterfacesWan1TransitionIspName.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.isp.name", data.InterfacesWan1TransitionIspName.ValueString())
	}
	if !data.InterfacesWan1TransitionMapDraftEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.map.draft.enabled", data.InterfacesWan1TransitionMapDraftEnabled.ValueBool())
	}
	if !data.InterfacesWan1TransitionMapRulesSource.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.map.rules.source", data.InterfacesWan1TransitionMapRulesSource.ValueString())
	}
	if len(data.InterfacesWan1TransitionMapRulesEntries) > 0 {
		body, _ = sjson.Set(body, "interfaces.wan1.transition.map.rules.entries", []interface{}{})
		for _, item := range data.InterfacesWan1TransitionMapRulesEntries {
			itemBody := ""
			if !item.EaBitLength.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "eaBitLength", item.EaBitLength.ValueInt64())
			}
			if !item.PsIdOffset.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "psIdOffset", item.PsIdOffset.ValueInt64())
			}
			if !item.BorderRelayAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "borderRelay.address", item.BorderRelayAddress.ValueString())
			}
			if !item.BorderRelayPrefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "borderRelay.prefix", item.BorderRelayPrefix.ValueString())
			}
			if !item.Ipv4PrefixAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4.prefix.address", item.Ipv4PrefixAddress.ValueString())
			}
			if !item.Ipv4PrefixLength.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4.prefix.length", item.Ipv4PrefixLength.ValueInt64())
			}
			if !item.Ipv6PrefixAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6.prefix.address", item.Ipv6PrefixAddress.ValueString())
			}
			if !item.Ipv6PrefixLength.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6.prefix.length", item.Ipv6PrefixLength.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.map.rules.entries.-1", itemBody)
		}
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
	if !data.InterfacesWan2PppoeAuthenticationPasswordWo.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.pppoe.authentication.password", data.InterfacesWan2PppoeAuthenticationPasswordWo.ValueString())
	} else if !data.InterfacesWan2PppoeAuthenticationPassword.IsNull() {
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
	if !data.InterfacesWan2TransitionEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.enabled", data.InterfacesWan2TransitionEnabled.ValueBool())
	}
	if !data.InterfacesWan2TransitionMode.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.mode", data.InterfacesWan2TransitionMode.ValueString())
	}
	if !data.InterfacesWan2TransitionAuthenticationEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.authentication.enabled", data.InterfacesWan2TransitionAuthenticationEnabled.ValueBool())
	}
	if !data.InterfacesWan2TransitionAuthenticationPassword.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.authentication.password", data.InterfacesWan2TransitionAuthenticationPassword.ValueString())
	}
	if !data.InterfacesWan2TransitionAuthenticationUsername.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.authentication.username", data.InterfacesWan2TransitionAuthenticationUsername.ValueString())
	}
	if !data.InterfacesWan2TransitionIpv4LocalAddress.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.ipv4.local.address", data.InterfacesWan2TransitionIpv4LocalAddress.ValueString())
	}
	if !data.InterfacesWan2TransitionIpv6InterfaceId.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.ipv6.interfaceId", data.InterfacesWan2TransitionIpv6InterfaceId.ValueString())
	}
	if !data.InterfacesWan2TransitionIpv6BorderRelayAddress.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.ipv6.borderRelay.address", data.InterfacesWan2TransitionIpv6BorderRelayAddress.ValueString())
	}
	if !data.InterfacesWan2TransitionIpv6BorderRelayMode.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.ipv6.borderRelay.mode", data.InterfacesWan2TransitionIpv6BorderRelayMode.ValueString())
	}
	if !data.InterfacesWan2TransitionIspName.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.isp.name", data.InterfacesWan2TransitionIspName.ValueString())
	}
	if !data.InterfacesWan2TransitionMapDraftEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.map.draft.enabled", data.InterfacesWan2TransitionMapDraftEnabled.ValueBool())
	}
	if !data.InterfacesWan2TransitionMapRulesSource.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.map.rules.source", data.InterfacesWan2TransitionMapRulesSource.ValueString())
	}
	if len(data.InterfacesWan2TransitionMapRulesEntries) > 0 {
		body, _ = sjson.Set(body, "interfaces.wan2.transition.map.rules.entries", []interface{}{})
		for _, item := range data.InterfacesWan2TransitionMapRulesEntries {
			itemBody := ""
			if !item.EaBitLength.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "eaBitLength", item.EaBitLength.ValueInt64())
			}
			if !item.PsIdOffset.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "psIdOffset", item.PsIdOffset.ValueInt64())
			}
			if !item.BorderRelayAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "borderRelay.address", item.BorderRelayAddress.ValueString())
			}
			if !item.BorderRelayPrefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "borderRelay.prefix", item.BorderRelayPrefix.ValueString())
			}
			if !item.Ipv4PrefixAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4.prefix.address", item.Ipv4PrefixAddress.ValueString())
			}
			if !item.Ipv4PrefixLength.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4.prefix.length", item.Ipv4PrefixLength.ValueInt64())
			}
			if !item.Ipv6PrefixAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6.prefix.address", item.Ipv6PrefixAddress.ValueString())
			}
			if !item.Ipv6PrefixLength.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6.prefix.length", item.Ipv6PrefixLength.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.map.rules.entries.-1", itemBody)
		}
	}
	if !data.InterfacesWan2VlanTaggingEnabled.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.vlanTagging.enabled", data.InterfacesWan2VlanTaggingEnabled.ValueBool())
	}
	if !data.InterfacesWan2VlanTaggingVlanId.IsNull() {
		body, _ = sjson.Set(body, "interfaces.wan2.vlanTagging.vlanId", data.InterfacesWan2VlanTaggingVlanId.ValueInt64())
	}
	return body
}

// setWoValues copies write-only attribute values from the configuration, as they are always null in the plan
func (data *ApplianceUplinksSettings) setWoValues(config ApplianceUplinksSettings) {
	data.InterfacesWan1PppoeAuthenticationPasswordWo = config.InterfacesWan1PppoeAuthenticationPasswordWo
	data.InterfacesWan2PppoeAuthenticationPasswordWo = config.InterfacesWan2PppoeAuthenticationPasswordWo
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPreservingNulls

// toBodyPreservingNulls walks the same writable-attribute schema as toBody but
// reads directly from the raw API response (gjson) instead of from the
// Terraform model. Unlike toBody, it preserves attributes that the API
// explicitly returned as `null` (emitting them as JSON `null` rather than
// dropping them). This is used by the singleton restoreOriginalStateOnDestroy
// path so that explicit-null fields captured during Create are restored on
// Delete. Keep this method in sync with toBody — both walk the same
// `.Attributes` schema and must agree on which fields are writable.
func (data ApplianceUplinksSettings) toBodyPreservingNulls(ctx context.Context, res meraki.Res) string {
	body := ""
	if value := res.Get("interfaces.wan1.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan1.pppoe.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.pppoe.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.pppoe.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan1.pppoe.authentication.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.pppoe.authentication.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.pppoe.authentication.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan1.pppoe.authentication.username"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.pppoe.authentication.username", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.pppoe.authentication.username", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.address"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.svis.ipv4.address", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv4.address", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.assignmentMode"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.svis.ipv4.assignmentMode", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv4.assignmentMode", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.gateway"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.svis.ipv4.gateway", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv4.gateway", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.svis.ipv4.nameservers.addresses"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.svis.ipv4.nameservers.addresses", "null")
		} else {
			var values []string
			for _, v := range value.Array() {
				values = append(values, v.String())
			}
			body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv4.nameservers.addresses", values)
		}
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.address"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.svis.ipv6.address", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv6.address", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.assignmentMode"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.svis.ipv6.assignmentMode", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv6.assignmentMode", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.gateway"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.svis.ipv6.gateway", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv6.gateway", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.svis.ipv6.nameservers.addresses"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.svis.ipv6.nameservers.addresses", "null")
		} else {
			var values []string
			for _, v := range value.Array() {
				values = append(values, v.String())
			}
			body, _ = sjson.Set(body, "interfaces.wan1.svis.ipv6.nameservers.addresses", values)
		}
	}
	if value := res.Get("interfaces.wan1.transition.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan1.transition.mode"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.mode", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.mode", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.transition.authentication.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.authentication.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.authentication.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan1.transition.authentication.password"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.authentication.password", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.authentication.password", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.transition.authentication.username"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.authentication.username", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.authentication.username", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.transition.ipv4.local.address"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.ipv4.local.address", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.ipv4.local.address", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.transition.ipv6.interfaceId"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.ipv6.interfaceId", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.ipv6.interfaceId", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.transition.ipv6.borderRelay.address"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.ipv6.borderRelay.address", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.ipv6.borderRelay.address", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.transition.ipv6.borderRelay.mode"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.ipv6.borderRelay.mode", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.ipv6.borderRelay.mode", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.transition.isp.name"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.isp.name", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.isp.name", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.transition.map.draft.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.map.draft.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.map.draft.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan1.transition.map.rules.source"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.map.rules.source", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.map.rules.source", value.String())
		}
	}
	if value := res.Get("interfaces.wan1.transition.map.rules.entries"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.transition.map.rules.entries", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.transition.map.rules.entries", []interface{}{})
			parent := &body
			value.ForEach(func(k, res gjson.Result) bool {
				body := ""
				if value := res.Get("eaBitLength"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "eaBitLength", "null")
					} else {
						body, _ = sjson.Set(body, "eaBitLength", value.Int())
					}
				}
				if value := res.Get("psIdOffset"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "psIdOffset", "null")
					} else {
						body, _ = sjson.Set(body, "psIdOffset", value.Int())
					}
				}
				if value := res.Get("borderRelay.address"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "borderRelay.address", "null")
					} else {
						body, _ = sjson.Set(body, "borderRelay.address", value.String())
					}
				}
				if value := res.Get("borderRelay.prefix"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "borderRelay.prefix", "null")
					} else {
						body, _ = sjson.Set(body, "borderRelay.prefix", value.String())
					}
				}
				if value := res.Get("ipv4.prefix.address"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "ipv4.prefix.address", "null")
					} else {
						body, _ = sjson.Set(body, "ipv4.prefix.address", value.String())
					}
				}
				if value := res.Get("ipv4.prefix.length"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "ipv4.prefix.length", "null")
					} else {
						body, _ = sjson.Set(body, "ipv4.prefix.length", value.Int())
					}
				}
				if value := res.Get("ipv6.prefix.address"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "ipv6.prefix.address", "null")
					} else {
						body, _ = sjson.Set(body, "ipv6.prefix.address", value.String())
					}
				}
				if value := res.Get("ipv6.prefix.length"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "ipv6.prefix.length", "null")
					} else {
						body, _ = sjson.Set(body, "ipv6.prefix.length", value.Int())
					}
				}
				*parent, _ = sjson.SetRaw(*parent, "interfaces.wan1.transition.map.rules.entries.-1", body)
				return true
			})
		}
	}
	if value := res.Get("interfaces.wan1.vlanTagging.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.vlanTagging.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.vlanTagging.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan1.vlanTagging.vlanId"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan1.vlanTagging.vlanId", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan1.vlanTagging.vlanId", value.Int())
		}
	}
	if value := res.Get("interfaces.wan2.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan2.pppoe.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.pppoe.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.pppoe.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan2.pppoe.authentication.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.pppoe.authentication.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.pppoe.authentication.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan2.pppoe.authentication.username"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.pppoe.authentication.username", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.pppoe.authentication.username", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.address"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.svis.ipv4.address", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv4.address", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.assignmentMode"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.svis.ipv4.assignmentMode", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv4.assignmentMode", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.gateway"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.svis.ipv4.gateway", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv4.gateway", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.svis.ipv4.nameservers.addresses"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.svis.ipv4.nameservers.addresses", "null")
		} else {
			var values []string
			for _, v := range value.Array() {
				values = append(values, v.String())
			}
			body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv4.nameservers.addresses", values)
		}
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.address"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.svis.ipv6.address", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv6.address", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.assignmentMode"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.svis.ipv6.assignmentMode", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv6.assignmentMode", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.gateway"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.svis.ipv6.gateway", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv6.gateway", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.svis.ipv6.nameservers.addresses"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.svis.ipv6.nameservers.addresses", "null")
		} else {
			var values []string
			for _, v := range value.Array() {
				values = append(values, v.String())
			}
			body, _ = sjson.Set(body, "interfaces.wan2.svis.ipv6.nameservers.addresses", values)
		}
	}
	if value := res.Get("interfaces.wan2.transition.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan2.transition.mode"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.mode", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.mode", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.transition.authentication.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.authentication.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.authentication.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan2.transition.authentication.password"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.authentication.password", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.authentication.password", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.transition.authentication.username"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.authentication.username", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.authentication.username", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.transition.ipv4.local.address"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.ipv4.local.address", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.ipv4.local.address", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.transition.ipv6.interfaceId"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.ipv6.interfaceId", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.ipv6.interfaceId", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.transition.ipv6.borderRelay.address"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.ipv6.borderRelay.address", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.ipv6.borderRelay.address", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.transition.ipv6.borderRelay.mode"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.ipv6.borderRelay.mode", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.ipv6.borderRelay.mode", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.transition.isp.name"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.isp.name", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.isp.name", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.transition.map.draft.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.map.draft.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.map.draft.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan2.transition.map.rules.source"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.map.rules.source", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.map.rules.source", value.String())
		}
	}
	if value := res.Get("interfaces.wan2.transition.map.rules.entries"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.transition.map.rules.entries", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.transition.map.rules.entries", []interface{}{})
			parent := &body
			value.ForEach(func(k, res gjson.Result) bool {
				body := ""
				if value := res.Get("eaBitLength"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "eaBitLength", "null")
					} else {
						body, _ = sjson.Set(body, "eaBitLength", value.Int())
					}
				}
				if value := res.Get("psIdOffset"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "psIdOffset", "null")
					} else {
						body, _ = sjson.Set(body, "psIdOffset", value.Int())
					}
				}
				if value := res.Get("borderRelay.address"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "borderRelay.address", "null")
					} else {
						body, _ = sjson.Set(body, "borderRelay.address", value.String())
					}
				}
				if value := res.Get("borderRelay.prefix"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "borderRelay.prefix", "null")
					} else {
						body, _ = sjson.Set(body, "borderRelay.prefix", value.String())
					}
				}
				if value := res.Get("ipv4.prefix.address"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "ipv4.prefix.address", "null")
					} else {
						body, _ = sjson.Set(body, "ipv4.prefix.address", value.String())
					}
				}
				if value := res.Get("ipv4.prefix.length"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "ipv4.prefix.length", "null")
					} else {
						body, _ = sjson.Set(body, "ipv4.prefix.length", value.Int())
					}
				}
				if value := res.Get("ipv6.prefix.address"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "ipv6.prefix.address", "null")
					} else {
						body, _ = sjson.Set(body, "ipv6.prefix.address", value.String())
					}
				}
				if value := res.Get("ipv6.prefix.length"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "ipv6.prefix.length", "null")
					} else {
						body, _ = sjson.Set(body, "ipv6.prefix.length", value.Int())
					}
				}
				*parent, _ = sjson.SetRaw(*parent, "interfaces.wan2.transition.map.rules.entries.-1", body)
				return true
			})
		}
	}
	if value := res.Get("interfaces.wan2.vlanTagging.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.vlanTagging.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.vlanTagging.enabled", value.Bool())
		}
	}
	if value := res.Get("interfaces.wan2.vlanTagging.vlanId"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "interfaces.wan2.vlanTagging.vlanId", "null")
		} else {
			body, _ = sjson.Set(body, "interfaces.wan2.vlanTagging.vlanId", value.Int())
		}
	}
	return body
}

// End of section. //template:end toBodyPreservingNulls

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
	if value := res.Get("interfaces.wan1.transition.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1TransitionEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.transition.mode"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.authentication.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionAuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1TransitionAuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.transition.authentication.password"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionAuthenticationPassword = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionAuthenticationPassword = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.authentication.username"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionAuthenticationUsername = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionAuthenticationUsername = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.ipv4.local.address"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionIpv4LocalAddress = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionIpv4LocalAddress = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.ipv6.interfaceId"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionIpv6InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionIpv6InterfaceId = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.ipv6.borderRelay.address"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionIpv6BorderRelayAddress = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionIpv6BorderRelayAddress = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.ipv6.borderRelay.mode"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionIpv6BorderRelayMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionIpv6BorderRelayMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.isp.name"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionIspName = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionIspName = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.map.draft.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionMapDraftEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1TransitionMapDraftEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.transition.map.rules.source"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionMapRulesSource = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionMapRulesSource = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.map.rules.entries"); value.Exists() && value.Value() != nil {
		data.InterfacesWan1TransitionMapRulesEntries = make([]ApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries{}
			if value := res.Get("eaBitLength"); value.Exists() && value.Value() != nil {
				data.EaBitLength = types.Int64Value(value.Int())
			} else {
				data.EaBitLength = types.Int64Null()
			}
			if value := res.Get("psIdOffset"); value.Exists() && value.Value() != nil {
				data.PsIdOffset = types.Int64Value(value.Int())
			} else {
				data.PsIdOffset = types.Int64Null()
			}
			if value := res.Get("borderRelay.address"); value.Exists() && value.Value() != nil {
				data.BorderRelayAddress = types.StringValue(value.String())
			} else {
				data.BorderRelayAddress = types.StringNull()
			}
			if value := res.Get("borderRelay.prefix"); value.Exists() && value.Value() != nil {
				data.BorderRelayPrefix = types.StringValue(value.String())
			} else {
				data.BorderRelayPrefix = types.StringNull()
			}
			if value := res.Get("ipv4.prefix.address"); value.Exists() && value.Value() != nil {
				data.Ipv4PrefixAddress = types.StringValue(value.String())
			} else {
				data.Ipv4PrefixAddress = types.StringNull()
			}
			if value := res.Get("ipv4.prefix.length"); value.Exists() && value.Value() != nil {
				data.Ipv4PrefixLength = types.Int64Value(value.Int())
			} else {
				data.Ipv4PrefixLength = types.Int64Null()
			}
			if value := res.Get("ipv6.prefix.address"); value.Exists() && value.Value() != nil {
				data.Ipv6PrefixAddress = types.StringValue(value.String())
			} else {
				data.Ipv6PrefixAddress = types.StringNull()
			}
			if value := res.Get("ipv6.prefix.length"); value.Exists() && value.Value() != nil {
				data.Ipv6PrefixLength = types.Int64Value(value.Int())
			} else {
				data.Ipv6PrefixLength = types.Int64Null()
			}
			(*parent).InterfacesWan1TransitionMapRulesEntries = append((*parent).InterfacesWan1TransitionMapRulesEntries, data)
			return true
		})
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
	if value := res.Get("interfaces.wan2.transition.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2TransitionEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.transition.mode"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.authentication.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionAuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2TransitionAuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.transition.authentication.password"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionAuthenticationPassword = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionAuthenticationPassword = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.authentication.username"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionAuthenticationUsername = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionAuthenticationUsername = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.ipv4.local.address"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionIpv4LocalAddress = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionIpv4LocalAddress = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.ipv6.interfaceId"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionIpv6InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionIpv6InterfaceId = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.ipv6.borderRelay.address"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionIpv6BorderRelayAddress = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionIpv6BorderRelayAddress = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.ipv6.borderRelay.mode"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionIpv6BorderRelayMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionIpv6BorderRelayMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.isp.name"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionIspName = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionIspName = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.map.draft.enabled"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionMapDraftEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2TransitionMapDraftEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.transition.map.rules.source"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionMapRulesSource = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionMapRulesSource = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.map.rules.entries"); value.Exists() && value.Value() != nil {
		data.InterfacesWan2TransitionMapRulesEntries = make([]ApplianceUplinksSettingsInterfacesWan2TransitionMapRulesEntries, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceUplinksSettingsInterfacesWan2TransitionMapRulesEntries{}
			if value := res.Get("eaBitLength"); value.Exists() && value.Value() != nil {
				data.EaBitLength = types.Int64Value(value.Int())
			} else {
				data.EaBitLength = types.Int64Null()
			}
			if value := res.Get("psIdOffset"); value.Exists() && value.Value() != nil {
				data.PsIdOffset = types.Int64Value(value.Int())
			} else {
				data.PsIdOffset = types.Int64Null()
			}
			if value := res.Get("borderRelay.address"); value.Exists() && value.Value() != nil {
				data.BorderRelayAddress = types.StringValue(value.String())
			} else {
				data.BorderRelayAddress = types.StringNull()
			}
			if value := res.Get("borderRelay.prefix"); value.Exists() && value.Value() != nil {
				data.BorderRelayPrefix = types.StringValue(value.String())
			} else {
				data.BorderRelayPrefix = types.StringNull()
			}
			if value := res.Get("ipv4.prefix.address"); value.Exists() && value.Value() != nil {
				data.Ipv4PrefixAddress = types.StringValue(value.String())
			} else {
				data.Ipv4PrefixAddress = types.StringNull()
			}
			if value := res.Get("ipv4.prefix.length"); value.Exists() && value.Value() != nil {
				data.Ipv4PrefixLength = types.Int64Value(value.Int())
			} else {
				data.Ipv4PrefixLength = types.Int64Null()
			}
			if value := res.Get("ipv6.prefix.address"); value.Exists() && value.Value() != nil {
				data.Ipv6PrefixAddress = types.StringValue(value.String())
			} else {
				data.Ipv6PrefixAddress = types.StringNull()
			}
			if value := res.Get("ipv6.prefix.length"); value.Exists() && value.Value() != nil {
				data.Ipv6PrefixLength = types.Int64Value(value.Int())
			} else {
				data.Ipv6PrefixLength = types.Int64Null()
			}
			(*parent).InterfacesWan2TransitionMapRulesEntries = append((*parent).InterfacesWan2TransitionMapRulesEntries, data)
			return true
		})
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
	if value := res.Get("interfaces.wan1.transition.enabled"); value.Exists() && !data.InterfacesWan1TransitionEnabled.IsNull() {
		data.InterfacesWan1TransitionEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1TransitionEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.transition.mode"); value.Exists() && !data.InterfacesWan1TransitionMode.IsNull() {
		data.InterfacesWan1TransitionMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.authentication.enabled"); value.Exists() && !data.InterfacesWan1TransitionAuthenticationEnabled.IsNull() {
		data.InterfacesWan1TransitionAuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1TransitionAuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.transition.authentication.password"); value.Exists() && !data.InterfacesWan1TransitionAuthenticationPassword.IsNull() {
		data.InterfacesWan1TransitionAuthenticationPassword = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionAuthenticationPassword = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.authentication.username"); value.Exists() && !data.InterfacesWan1TransitionAuthenticationUsername.IsNull() {
		data.InterfacesWan1TransitionAuthenticationUsername = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionAuthenticationUsername = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.ipv4.local.address"); value.Exists() && !data.InterfacesWan1TransitionIpv4LocalAddress.IsNull() {
		data.InterfacesWan1TransitionIpv4LocalAddress = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionIpv4LocalAddress = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.ipv6.interfaceId"); value.Exists() && !data.InterfacesWan1TransitionIpv6InterfaceId.IsNull() {
		data.InterfacesWan1TransitionIpv6InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionIpv6InterfaceId = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.ipv6.borderRelay.address"); value.Exists() && !data.InterfacesWan1TransitionIpv6BorderRelayAddress.IsNull() {
		data.InterfacesWan1TransitionIpv6BorderRelayAddress = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionIpv6BorderRelayAddress = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.ipv6.borderRelay.mode"); value.Exists() && !data.InterfacesWan1TransitionIpv6BorderRelayMode.IsNull() {
		data.InterfacesWan1TransitionIpv6BorderRelayMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionIpv6BorderRelayMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.isp.name"); value.Exists() && !data.InterfacesWan1TransitionIspName.IsNull() {
		data.InterfacesWan1TransitionIspName = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionIspName = types.StringNull()
	}
	if value := res.Get("interfaces.wan1.transition.map.draft.enabled"); value.Exists() && !data.InterfacesWan1TransitionMapDraftEnabled.IsNull() {
		data.InterfacesWan1TransitionMapDraftEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan1TransitionMapDraftEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan1.transition.map.rules.source"); value.Exists() && !data.InterfacesWan1TransitionMapRulesSource.IsNull() {
		data.InterfacesWan1TransitionMapRulesSource = types.StringValue(value.String())
	} else {
		data.InterfacesWan1TransitionMapRulesSource = types.StringNull()
	}
	for i := 0; i < len(data.InterfacesWan1TransitionMapRulesEntries); i++ {
		keys := [...]string{"eaBitLength", "psIdOffset", "borderRelay.address", "borderRelay.prefix", "ipv4.prefix.address", "ipv4.prefix.length", "ipv6.prefix.address", "ipv6.prefix.length"}
		keyValues := [...]string{strconv.FormatInt(data.InterfacesWan1TransitionMapRulesEntries[i].EaBitLength.ValueInt64(), 10), strconv.FormatInt(data.InterfacesWan1TransitionMapRulesEntries[i].PsIdOffset.ValueInt64(), 10), data.InterfacesWan1TransitionMapRulesEntries[i].BorderRelayAddress.ValueString(), data.InterfacesWan1TransitionMapRulesEntries[i].BorderRelayPrefix.ValueString(), data.InterfacesWan1TransitionMapRulesEntries[i].Ipv4PrefixAddress.ValueString(), strconv.FormatInt(data.InterfacesWan1TransitionMapRulesEntries[i].Ipv4PrefixLength.ValueInt64(), 10), data.InterfacesWan1TransitionMapRulesEntries[i].Ipv6PrefixAddress.ValueString(), strconv.FormatInt(data.InterfacesWan1TransitionMapRulesEntries[i].Ipv6PrefixLength.ValueInt64(), 10)}

		parent := &data
		data := (*parent).InterfacesWan1TransitionMapRulesEntries[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("interfaces.wan1.transition.map.rules.entries").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing InterfacesWan1TransitionMapRulesEntries[%d] = %+v",
				i,
				(*parent).InterfacesWan1TransitionMapRulesEntries[i],
			))
			(*parent).InterfacesWan1TransitionMapRulesEntries = slices.Delete((*parent).InterfacesWan1TransitionMapRulesEntries, i, i+1)
			i--

			continue
		}
		if value := res.Get("eaBitLength"); value.Exists() && !data.EaBitLength.IsNull() {
			data.EaBitLength = types.Int64Value(value.Int())
		} else {
			data.EaBitLength = types.Int64Null()
		}
		if value := res.Get("psIdOffset"); value.Exists() && !data.PsIdOffset.IsNull() {
			data.PsIdOffset = types.Int64Value(value.Int())
		} else {
			data.PsIdOffset = types.Int64Null()
		}
		if value := res.Get("borderRelay.address"); value.Exists() && !data.BorderRelayAddress.IsNull() {
			data.BorderRelayAddress = types.StringValue(value.String())
		} else {
			data.BorderRelayAddress = types.StringNull()
		}
		if value := res.Get("borderRelay.prefix"); value.Exists() && !data.BorderRelayPrefix.IsNull() {
			data.BorderRelayPrefix = types.StringValue(value.String())
		} else {
			data.BorderRelayPrefix = types.StringNull()
		}
		if value := res.Get("ipv4.prefix.address"); value.Exists() && !data.Ipv4PrefixAddress.IsNull() {
			data.Ipv4PrefixAddress = types.StringValue(value.String())
		} else {
			data.Ipv4PrefixAddress = types.StringNull()
		}
		if value := res.Get("ipv4.prefix.length"); value.Exists() && !data.Ipv4PrefixLength.IsNull() {
			data.Ipv4PrefixLength = types.Int64Value(value.Int())
		} else {
			data.Ipv4PrefixLength = types.Int64Null()
		}
		if value := res.Get("ipv6.prefix.address"); value.Exists() && !data.Ipv6PrefixAddress.IsNull() {
			data.Ipv6PrefixAddress = types.StringValue(value.String())
		} else {
			data.Ipv6PrefixAddress = types.StringNull()
		}
		if value := res.Get("ipv6.prefix.length"); value.Exists() && !data.Ipv6PrefixLength.IsNull() {
			data.Ipv6PrefixLength = types.Int64Value(value.Int())
		} else {
			data.Ipv6PrefixLength = types.Int64Null()
		}
		(*parent).InterfacesWan1TransitionMapRulesEntries[i] = data
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
	if value := res.Get("interfaces.wan2.transition.enabled"); value.Exists() && !data.InterfacesWan2TransitionEnabled.IsNull() {
		data.InterfacesWan2TransitionEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2TransitionEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.transition.mode"); value.Exists() && !data.InterfacesWan2TransitionMode.IsNull() {
		data.InterfacesWan2TransitionMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.authentication.enabled"); value.Exists() && !data.InterfacesWan2TransitionAuthenticationEnabled.IsNull() {
		data.InterfacesWan2TransitionAuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2TransitionAuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.transition.authentication.password"); value.Exists() && !data.InterfacesWan2TransitionAuthenticationPassword.IsNull() {
		data.InterfacesWan2TransitionAuthenticationPassword = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionAuthenticationPassword = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.authentication.username"); value.Exists() && !data.InterfacesWan2TransitionAuthenticationUsername.IsNull() {
		data.InterfacesWan2TransitionAuthenticationUsername = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionAuthenticationUsername = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.ipv4.local.address"); value.Exists() && !data.InterfacesWan2TransitionIpv4LocalAddress.IsNull() {
		data.InterfacesWan2TransitionIpv4LocalAddress = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionIpv4LocalAddress = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.ipv6.interfaceId"); value.Exists() && !data.InterfacesWan2TransitionIpv6InterfaceId.IsNull() {
		data.InterfacesWan2TransitionIpv6InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionIpv6InterfaceId = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.ipv6.borderRelay.address"); value.Exists() && !data.InterfacesWan2TransitionIpv6BorderRelayAddress.IsNull() {
		data.InterfacesWan2TransitionIpv6BorderRelayAddress = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionIpv6BorderRelayAddress = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.ipv6.borderRelay.mode"); value.Exists() && !data.InterfacesWan2TransitionIpv6BorderRelayMode.IsNull() {
		data.InterfacesWan2TransitionIpv6BorderRelayMode = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionIpv6BorderRelayMode = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.isp.name"); value.Exists() && !data.InterfacesWan2TransitionIspName.IsNull() {
		data.InterfacesWan2TransitionIspName = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionIspName = types.StringNull()
	}
	if value := res.Get("interfaces.wan2.transition.map.draft.enabled"); value.Exists() && !data.InterfacesWan2TransitionMapDraftEnabled.IsNull() {
		data.InterfacesWan2TransitionMapDraftEnabled = types.BoolValue(value.Bool())
	} else {
		data.InterfacesWan2TransitionMapDraftEnabled = types.BoolNull()
	}
	if value := res.Get("interfaces.wan2.transition.map.rules.source"); value.Exists() && !data.InterfacesWan2TransitionMapRulesSource.IsNull() {
		data.InterfacesWan2TransitionMapRulesSource = types.StringValue(value.String())
	} else {
		data.InterfacesWan2TransitionMapRulesSource = types.StringNull()
	}
	for i := 0; i < len(data.InterfacesWan2TransitionMapRulesEntries); i++ {
		keys := [...]string{"eaBitLength", "psIdOffset", "borderRelay.address", "borderRelay.prefix", "ipv4.prefix.address", "ipv4.prefix.length", "ipv6.prefix.address", "ipv6.prefix.length"}
		keyValues := [...]string{strconv.FormatInt(data.InterfacesWan2TransitionMapRulesEntries[i].EaBitLength.ValueInt64(), 10), strconv.FormatInt(data.InterfacesWan2TransitionMapRulesEntries[i].PsIdOffset.ValueInt64(), 10), data.InterfacesWan2TransitionMapRulesEntries[i].BorderRelayAddress.ValueString(), data.InterfacesWan2TransitionMapRulesEntries[i].BorderRelayPrefix.ValueString(), data.InterfacesWan2TransitionMapRulesEntries[i].Ipv4PrefixAddress.ValueString(), strconv.FormatInt(data.InterfacesWan2TransitionMapRulesEntries[i].Ipv4PrefixLength.ValueInt64(), 10), data.InterfacesWan2TransitionMapRulesEntries[i].Ipv6PrefixAddress.ValueString(), strconv.FormatInt(data.InterfacesWan2TransitionMapRulesEntries[i].Ipv6PrefixLength.ValueInt64(), 10)}

		parent := &data
		data := (*parent).InterfacesWan2TransitionMapRulesEntries[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("interfaces.wan2.transition.map.rules.entries").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing InterfacesWan2TransitionMapRulesEntries[%d] = %+v",
				i,
				(*parent).InterfacesWan2TransitionMapRulesEntries[i],
			))
			(*parent).InterfacesWan2TransitionMapRulesEntries = slices.Delete((*parent).InterfacesWan2TransitionMapRulesEntries, i, i+1)
			i--

			continue
		}
		if value := res.Get("eaBitLength"); value.Exists() && !data.EaBitLength.IsNull() {
			data.EaBitLength = types.Int64Value(value.Int())
		} else {
			data.EaBitLength = types.Int64Null()
		}
		if value := res.Get("psIdOffset"); value.Exists() && !data.PsIdOffset.IsNull() {
			data.PsIdOffset = types.Int64Value(value.Int())
		} else {
			data.PsIdOffset = types.Int64Null()
		}
		if value := res.Get("borderRelay.address"); value.Exists() && !data.BorderRelayAddress.IsNull() {
			data.BorderRelayAddress = types.StringValue(value.String())
		} else {
			data.BorderRelayAddress = types.StringNull()
		}
		if value := res.Get("borderRelay.prefix"); value.Exists() && !data.BorderRelayPrefix.IsNull() {
			data.BorderRelayPrefix = types.StringValue(value.String())
		} else {
			data.BorderRelayPrefix = types.StringNull()
		}
		if value := res.Get("ipv4.prefix.address"); value.Exists() && !data.Ipv4PrefixAddress.IsNull() {
			data.Ipv4PrefixAddress = types.StringValue(value.String())
		} else {
			data.Ipv4PrefixAddress = types.StringNull()
		}
		if value := res.Get("ipv4.prefix.length"); value.Exists() && !data.Ipv4PrefixLength.IsNull() {
			data.Ipv4PrefixLength = types.Int64Value(value.Int())
		} else {
			data.Ipv4PrefixLength = types.Int64Null()
		}
		if value := res.Get("ipv6.prefix.address"); value.Exists() && !data.Ipv6PrefixAddress.IsNull() {
			data.Ipv6PrefixAddress = types.StringValue(value.String())
		} else {
			data.Ipv6PrefixAddress = types.StringNull()
		}
		if value := res.Get("ipv6.prefix.length"); value.Exists() && !data.Ipv6PrefixLength.IsNull() {
			data.Ipv6PrefixLength = types.Int64Value(value.Int())
		} else {
			data.Ipv6PrefixLength = types.Int64Null()
		}
		(*parent).InterfacesWan2TransitionMapRulesEntries[i] = data
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceUplinksSettings) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceUplinksSettingsIdentity) toIdentity(ctx context.Context, plan *ApplianceUplinksSettings) {
	data.Serial = plan.Serial
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceUplinksSettings) fromIdentity(ctx context.Context, identity *ApplianceUplinksSettingsIdentity) {
	data.Serial = identity.Serial
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceUplinksSettings) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
