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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSourceApplianceUplinksSettings - used only by the data source, always including every
// attribute (unlike model_resource.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model_resource.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model_resource.go's resource-side type names.

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

type DataSourceApplianceUplinksSettings struct {
	Id                                                 types.String                                                                `tfsdk:"id"`
	Serial                                             types.String                                                                `tfsdk:"serial"`
	InterfacesWan1Enabled                              types.Bool                                                                  `tfsdk:"interfaces_wan1_enabled"`
	InterfacesWan1PppoeEnabled                         types.Bool                                                                  `tfsdk:"interfaces_wan1_pppoe_enabled"`
	InterfacesWan1PppoeAuthenticationEnabled           types.Bool                                                                  `tfsdk:"interfaces_wan1_pppoe_authentication_enabled"`
	InterfacesWan1PppoeAuthenticationPassword          types.String                                                                `tfsdk:"interfaces_wan1_pppoe_authentication_password"`
	InterfacesWan1PppoeAuthenticationPasswordWo        types.String                                                                `tfsdk:"interfaces_wan1_pppoe_authentication_password_wo"`
	InterfacesWan1PppoeAuthenticationPasswordWoVersion types.Int64                                                                 `tfsdk:"interfaces_wan1_pppoe_authentication_password_wo_version"`
	InterfacesWan1PppoeAuthenticationUsername          types.String                                                                `tfsdk:"interfaces_wan1_pppoe_authentication_username"`
	InterfacesWan1SvisIpv4Address                      types.String                                                                `tfsdk:"interfaces_wan1_svis_ipv4_address"`
	InterfacesWan1SvisIpv4AssignmentMode               types.String                                                                `tfsdk:"interfaces_wan1_svis_ipv4_assignment_mode"`
	InterfacesWan1SvisIpv4Gateway                      types.String                                                                `tfsdk:"interfaces_wan1_svis_ipv4_gateway"`
	InterfacesWan1SvisIpv4NameserversAddresses         types.List                                                                  `tfsdk:"interfaces_wan1_svis_ipv4_nameservers_addresses"`
	InterfacesWan1SvisIpv6Address                      types.String                                                                `tfsdk:"interfaces_wan1_svis_ipv6_address"`
	InterfacesWan1SvisIpv6AssignmentMode               types.String                                                                `tfsdk:"interfaces_wan1_svis_ipv6_assignment_mode"`
	InterfacesWan1SvisIpv6Gateway                      types.String                                                                `tfsdk:"interfaces_wan1_svis_ipv6_gateway"`
	InterfacesWan1SvisIpv6NameserversAddresses         types.List                                                                  `tfsdk:"interfaces_wan1_svis_ipv6_nameservers_addresses"`
	InterfacesWan1TransitionEnabled                    types.Bool                                                                  `tfsdk:"interfaces_wan1_transition_enabled"`
	InterfacesWan1TransitionMode                       types.String                                                                `tfsdk:"interfaces_wan1_transition_mode"`
	InterfacesWan1TransitionAuthenticationEnabled      types.Bool                                                                  `tfsdk:"interfaces_wan1_transition_authentication_enabled"`
	InterfacesWan1TransitionAuthenticationPassword     types.String                                                                `tfsdk:"interfaces_wan1_transition_authentication_password"`
	InterfacesWan1TransitionAuthenticationUsername     types.String                                                                `tfsdk:"interfaces_wan1_transition_authentication_username"`
	InterfacesWan1TransitionIpv4LocalAddress           types.String                                                                `tfsdk:"interfaces_wan1_transition_ipv4_local_address"`
	InterfacesWan1TransitionIpv6InterfaceId            types.String                                                                `tfsdk:"interfaces_wan1_transition_ipv6_interface_id"`
	InterfacesWan1TransitionIpv6BorderRelayAddress     types.String                                                                `tfsdk:"interfaces_wan1_transition_ipv6_border_relay_address"`
	InterfacesWan1TransitionIpv6BorderRelayMode        types.String                                                                `tfsdk:"interfaces_wan1_transition_ipv6_border_relay_mode"`
	InterfacesWan1TransitionIspName                    types.String                                                                `tfsdk:"interfaces_wan1_transition_isp_name"`
	InterfacesWan1TransitionMapDraftEnabled            types.Bool                                                                  `tfsdk:"interfaces_wan1_transition_map_draft_enabled"`
	InterfacesWan1TransitionMapRulesSource             types.String                                                                `tfsdk:"interfaces_wan1_transition_map_rules_source"`
	InterfacesWan1TransitionMapRulesEntries            []DataSourceApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries `tfsdk:"interfaces_wan1_transition_map_rules_entries"`
	InterfacesWan1VlanTaggingEnabled                   types.Bool                                                                  `tfsdk:"interfaces_wan1_vlan_tagging_enabled"`
	InterfacesWan1VlanTaggingVlanId                    types.Int64                                                                 `tfsdk:"interfaces_wan1_vlan_tagging_vlan_id"`
	InterfacesWan2Enabled                              types.Bool                                                                  `tfsdk:"interfaces_wan2_enabled"`
	InterfacesWan2PppoeEnabled                         types.Bool                                                                  `tfsdk:"interfaces_wan2_pppoe_enabled"`
	InterfacesWan2PppoeAuthenticationEnabled           types.Bool                                                                  `tfsdk:"interfaces_wan2_pppoe_authentication_enabled"`
	InterfacesWan2PppoeAuthenticationPassword          types.String                                                                `tfsdk:"interfaces_wan2_pppoe_authentication_password"`
	InterfacesWan2PppoeAuthenticationPasswordWo        types.String                                                                `tfsdk:"interfaces_wan2_pppoe_authentication_password_wo"`
	InterfacesWan2PppoeAuthenticationPasswordWoVersion types.Int64                                                                 `tfsdk:"interfaces_wan2_pppoe_authentication_password_wo_version"`
	InterfacesWan2PppoeAuthenticationUsername          types.String                                                                `tfsdk:"interfaces_wan2_pppoe_authentication_username"`
	InterfacesWan2SvisIpv4Address                      types.String                                                                `tfsdk:"interfaces_wan2_svis_ipv4_address"`
	InterfacesWan2SvisIpv4AssignmentMode               types.String                                                                `tfsdk:"interfaces_wan2_svis_ipv4_assignment_mode"`
	InterfacesWan2SvisIpv4Gateway                      types.String                                                                `tfsdk:"interfaces_wan2_svis_ipv4_gateway"`
	InterfacesWan2SvisIpv4NameserversAddresses         types.List                                                                  `tfsdk:"interfaces_wan2_svis_ipv4_nameservers_addresses"`
	InterfacesWan2SvisIpv6Address                      types.String                                                                `tfsdk:"interfaces_wan2_svis_ipv6_address"`
	InterfacesWan2SvisIpv6AssignmentMode               types.String                                                                `tfsdk:"interfaces_wan2_svis_ipv6_assignment_mode"`
	InterfacesWan2SvisIpv6Gateway                      types.String                                                                `tfsdk:"interfaces_wan2_svis_ipv6_gateway"`
	InterfacesWan2SvisIpv6NameserversAddresses         types.List                                                                  `tfsdk:"interfaces_wan2_svis_ipv6_nameservers_addresses"`
	InterfacesWan2TransitionEnabled                    types.Bool                                                                  `tfsdk:"interfaces_wan2_transition_enabled"`
	InterfacesWan2TransitionMode                       types.String                                                                `tfsdk:"interfaces_wan2_transition_mode"`
	InterfacesWan2TransitionAuthenticationEnabled      types.Bool                                                                  `tfsdk:"interfaces_wan2_transition_authentication_enabled"`
	InterfacesWan2TransitionAuthenticationPassword     types.String                                                                `tfsdk:"interfaces_wan2_transition_authentication_password"`
	InterfacesWan2TransitionAuthenticationUsername     types.String                                                                `tfsdk:"interfaces_wan2_transition_authentication_username"`
	InterfacesWan2TransitionIpv4LocalAddress           types.String                                                                `tfsdk:"interfaces_wan2_transition_ipv4_local_address"`
	InterfacesWan2TransitionIpv6InterfaceId            types.String                                                                `tfsdk:"interfaces_wan2_transition_ipv6_interface_id"`
	InterfacesWan2TransitionIpv6BorderRelayAddress     types.String                                                                `tfsdk:"interfaces_wan2_transition_ipv6_border_relay_address"`
	InterfacesWan2TransitionIpv6BorderRelayMode        types.String                                                                `tfsdk:"interfaces_wan2_transition_ipv6_border_relay_mode"`
	InterfacesWan2TransitionIspName                    types.String                                                                `tfsdk:"interfaces_wan2_transition_isp_name"`
	InterfacesWan2TransitionMapDraftEnabled            types.Bool                                                                  `tfsdk:"interfaces_wan2_transition_map_draft_enabled"`
	InterfacesWan2TransitionMapRulesSource             types.String                                                                `tfsdk:"interfaces_wan2_transition_map_rules_source"`
	InterfacesWan2TransitionMapRulesEntries            []DataSourceApplianceUplinksSettingsInterfacesWan2TransitionMapRulesEntries `tfsdk:"interfaces_wan2_transition_map_rules_entries"`
	InterfacesWan2VlanTaggingEnabled                   types.Bool                                                                  `tfsdk:"interfaces_wan2_vlan_tagging_enabled"`
	InterfacesWan2VlanTaggingVlanId                    types.Int64                                                                 `tfsdk:"interfaces_wan2_vlan_tagging_vlan_id"`
}

type DataSourceApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries struct {
	EaBitLength        types.Int64  `tfsdk:"ea_bit_length"`
	PsIdOffset         types.Int64  `tfsdk:"ps_id_offset"`
	BorderRelayAddress types.String `tfsdk:"border_relay_address"`
	BorderRelayPrefix  types.String `tfsdk:"border_relay_prefix"`
	Ipv4PrefixAddress  types.String `tfsdk:"ipv4_prefix_address"`
	Ipv4PrefixLength   types.Int64  `tfsdk:"ipv4_prefix_length"`
	Ipv6PrefixAddress  types.String `tfsdk:"ipv6_prefix_address"`
	Ipv6PrefixLength   types.Int64  `tfsdk:"ipv6_prefix_length"`
}

type DataSourceApplianceUplinksSettingsInterfacesWan2TransitionMapRulesEntries struct {
	EaBitLength        types.Int64  `tfsdk:"ea_bit_length"`
	PsIdOffset         types.Int64  `tfsdk:"ps_id_offset"`
	BorderRelayAddress types.String `tfsdk:"border_relay_address"`
	BorderRelayPrefix  types.String `tfsdk:"border_relay_prefix"`
	Ipv4PrefixAddress  types.String `tfsdk:"ipv4_prefix_address"`
	Ipv4PrefixLength   types.Int64  `tfsdk:"ipv4_prefix_length"`
	Ipv6PrefixAddress  types.String `tfsdk:"ipv6_prefix_address"`
	Ipv6PrefixLength   types.Int64  `tfsdk:"ipv6_prefix_length"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceUplinksSettings) getPath() string {
	return fmt.Sprintf("/devices/%v/appliance/uplinks/settings", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceUplinksSettings) fromBody(ctx context.Context, res meraki.Res) {
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
		data.InterfacesWan1TransitionMapRulesEntries = make([]DataSourceApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceUplinksSettingsInterfacesWan1TransitionMapRulesEntries{}
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
		data.InterfacesWan2TransitionMapRulesEntries = make([]DataSourceApplianceUplinksSettingsInterfacesWan2TransitionMapRulesEntries, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceUplinksSettingsInterfacesWan2TransitionMapRulesEntries{}
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
