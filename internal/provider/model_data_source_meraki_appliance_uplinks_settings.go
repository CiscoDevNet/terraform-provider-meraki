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
// attribute (unlike model.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceApplianceUplinksSettings struct {
	Id                                                 types.String `tfsdk:"id"`
	Serial                                             types.String `tfsdk:"serial"`
	InterfacesWan1Enabled                              types.Bool   `tfsdk:"interfaces_wan1_enabled"`
	InterfacesWan1PppoeEnabled                         types.Bool   `tfsdk:"interfaces_wan1_pppoe_enabled"`
	InterfacesWan1PppoeAuthenticationEnabled           types.Bool   `tfsdk:"interfaces_wan1_pppoe_authentication_enabled"`
	InterfacesWan1PppoeAuthenticationPassword          types.String `tfsdk:"interfaces_wan1_pppoe_authentication_password"`
	InterfacesWan1PppoeAuthenticationPasswordWo        types.String `tfsdk:"interfaces_wan1_pppoe_authentication_password_wo"`
	InterfacesWan1PppoeAuthenticationPasswordWoVersion types.Int64  `tfsdk:"interfaces_wan1_pppoe_authentication_password_wo_version"`
	InterfacesWan1PppoeAuthenticationUsername          types.String `tfsdk:"interfaces_wan1_pppoe_authentication_username"`
	InterfacesWan1SvisIpv4Address                      types.String `tfsdk:"interfaces_wan1_svis_ipv4_address"`
	InterfacesWan1SvisIpv4AssignmentMode               types.String `tfsdk:"interfaces_wan1_svis_ipv4_assignment_mode"`
	InterfacesWan1SvisIpv4Gateway                      types.String `tfsdk:"interfaces_wan1_svis_ipv4_gateway"`
	InterfacesWan1SvisIpv4NameserversAddresses         types.List   `tfsdk:"interfaces_wan1_svis_ipv4_nameservers_addresses"`
	InterfacesWan1SvisIpv6Address                      types.String `tfsdk:"interfaces_wan1_svis_ipv6_address"`
	InterfacesWan1SvisIpv6AssignmentMode               types.String `tfsdk:"interfaces_wan1_svis_ipv6_assignment_mode"`
	InterfacesWan1SvisIpv6Gateway                      types.String `tfsdk:"interfaces_wan1_svis_ipv6_gateway"`
	InterfacesWan1SvisIpv6NameserversAddresses         types.List   `tfsdk:"interfaces_wan1_svis_ipv6_nameservers_addresses"`
	InterfacesWan1VlanTaggingEnabled                   types.Bool   `tfsdk:"interfaces_wan1_vlan_tagging_enabled"`
	InterfacesWan1VlanTaggingVlanId                    types.Int64  `tfsdk:"interfaces_wan1_vlan_tagging_vlan_id"`
	InterfacesWan2Enabled                              types.Bool   `tfsdk:"interfaces_wan2_enabled"`
	InterfacesWan2PppoeEnabled                         types.Bool   `tfsdk:"interfaces_wan2_pppoe_enabled"`
	InterfacesWan2PppoeAuthenticationEnabled           types.Bool   `tfsdk:"interfaces_wan2_pppoe_authentication_enabled"`
	InterfacesWan2PppoeAuthenticationPassword          types.String `tfsdk:"interfaces_wan2_pppoe_authentication_password"`
	InterfacesWan2PppoeAuthenticationPasswordWo        types.String `tfsdk:"interfaces_wan2_pppoe_authentication_password_wo"`
	InterfacesWan2PppoeAuthenticationPasswordWoVersion types.Int64  `tfsdk:"interfaces_wan2_pppoe_authentication_password_wo_version"`
	InterfacesWan2PppoeAuthenticationUsername          types.String `tfsdk:"interfaces_wan2_pppoe_authentication_username"`
	InterfacesWan2SvisIpv4Address                      types.String `tfsdk:"interfaces_wan2_svis_ipv4_address"`
	InterfacesWan2SvisIpv4AssignmentMode               types.String `tfsdk:"interfaces_wan2_svis_ipv4_assignment_mode"`
	InterfacesWan2SvisIpv4Gateway                      types.String `tfsdk:"interfaces_wan2_svis_ipv4_gateway"`
	InterfacesWan2SvisIpv4NameserversAddresses         types.List   `tfsdk:"interfaces_wan2_svis_ipv4_nameservers_addresses"`
	InterfacesWan2SvisIpv6Address                      types.String `tfsdk:"interfaces_wan2_svis_ipv6_address"`
	InterfacesWan2SvisIpv6AssignmentMode               types.String `tfsdk:"interfaces_wan2_svis_ipv6_assignment_mode"`
	InterfacesWan2SvisIpv6Gateway                      types.String `tfsdk:"interfaces_wan2_svis_ipv6_gateway"`
	InterfacesWan2SvisIpv6NameserversAddresses         types.List   `tfsdk:"interfaces_wan2_svis_ipv6_nameservers_addresses"`
	InterfacesWan2VlanTaggingEnabled                   types.Bool   `tfsdk:"interfaces_wan2_vlan_tagging_enabled"`
	InterfacesWan2VlanTaggingVlanId                    types.Int64  `tfsdk:"interfaces_wan2_vlan_tagging_vlan_id"`
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
