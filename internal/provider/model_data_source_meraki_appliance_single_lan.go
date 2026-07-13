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
// It emits a separate model - DataSourceApplianceSingleLAN - used only by the data source, always including every
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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceApplianceSingleLAN struct {
	Id                    types.String                                        `tfsdk:"id"`
	NetworkId             types.String                                        `tfsdk:"network_id"`
	ApplianceIp           types.String                                        `tfsdk:"appliance_ip"`
	Subnet                types.String                                        `tfsdk:"subnet"`
	Ipv6Enabled           types.Bool                                          `tfsdk:"ipv6_enabled"`
	Ipv6PrefixAssignments []DataSourceApplianceSingleLANIpv6PrefixAssignments `tfsdk:"ipv6_prefix_assignments"`
	MandatoryDhcpEnabled  types.Bool                                          `tfsdk:"mandatory_dhcp_enabled"`
}

type DataSourceApplianceSingleLANIpv6PrefixAssignments struct {
	Autonomous         types.Bool   `tfsdk:"autonomous"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	StaticApplianceIp6 types.String `tfsdk:"static_appliance_ip6"`
	StaticPrefix       types.String `tfsdk:"static_prefix"`
	OriginType         types.String `tfsdk:"origin_type"`
	OriginInterfaces   types.List   `tfsdk:"origin_interfaces"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceSingleLAN) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/singleLan", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceSingleLAN) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("applianceIp"); value.Exists() && value.Value() != nil {
		data.ApplianceIp = types.StringValue(value.String())
	} else {
		data.ApplianceIp = types.StringNull()
	}
	if value := res.Get("subnet"); value.Exists() && value.Value() != nil {
		data.Subnet = types.StringValue(value.String())
	} else {
		data.Subnet = types.StringNull()
	}
	if value := res.Get("ipv6.enabled"); value.Exists() && value.Value() != nil {
		data.Ipv6Enabled = types.BoolValue(value.Bool())
	} else {
		data.Ipv6Enabled = types.BoolNull()
	}
	if value := res.Get("ipv6.prefixAssignments"); value.Exists() && value.Value() != nil {
		data.Ipv6PrefixAssignments = make([]DataSourceApplianceSingleLANIpv6PrefixAssignments, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceSingleLANIpv6PrefixAssignments{}
			if value := res.Get("autonomous"); value.Exists() && value.Value() != nil {
				data.Autonomous = types.BoolValue(value.Bool())
			} else {
				data.Autonomous = types.BoolNull()
			}
			if value := res.Get("disabled"); value.Exists() && value.Value() != nil {
				data.Disabled = types.BoolValue(value.Bool())
			} else {
				data.Disabled = types.BoolNull()
			}
			if value := res.Get("staticApplianceIp6"); value.Exists() && value.Value() != nil {
				data.StaticApplianceIp6 = types.StringValue(value.String())
			} else {
				data.StaticApplianceIp6 = types.StringNull()
			}
			if value := res.Get("staticPrefix"); value.Exists() && value.Value() != nil {
				data.StaticPrefix = types.StringValue(value.String())
			} else {
				data.StaticPrefix = types.StringNull()
			}
			if value := res.Get("origin.type"); value.Exists() && value.Value() != nil {
				data.OriginType = types.StringValue(value.String())
			} else {
				data.OriginType = types.StringNull()
			}
			if value := res.Get("origin.interfaces"); value.Exists() && value.Value() != nil {
				data.OriginInterfaces = helpers.GetStringList(value.Array())
			} else {
				data.OriginInterfaces = types.ListNull(types.StringType)
			}
			(*parent).Ipv6PrefixAssignments = append((*parent).Ipv6PrefixAssignments, data)
			return true
		})
	}
	if value := res.Get("mandatoryDhcp.enabled"); value.Exists() && value.Value() != nil {
		data.MandatoryDhcpEnabled = types.BoolValue(value.Bool())
	} else {
		data.MandatoryDhcpEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody
