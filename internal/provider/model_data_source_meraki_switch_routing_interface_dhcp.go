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
// It emits a separate model - DataSourceSwitchRoutingInterfaceDHCP - used only by the data source, always including every
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

type DataSourceSwitchRoutingInterfaceDHCP struct {
	Id                   types.String                                             `tfsdk:"id"`
	Serial               types.String                                             `tfsdk:"serial"`
	InterfaceId          types.String                                             `tfsdk:"interface_id"`
	BootFileName         types.String                                             `tfsdk:"boot_file_name"`
	BootNextServer       types.String                                             `tfsdk:"boot_next_server"`
	BootOptionsEnabled   types.Bool                                               `tfsdk:"boot_options_enabled"`
	DhcpLeaseTime        types.String                                             `tfsdk:"dhcp_lease_time"`
	DhcpMode             types.String                                             `tfsdk:"dhcp_mode"`
	DnsNameserversOption types.String                                             `tfsdk:"dns_nameservers_option"`
	DhcpOptions          []DataSourceSwitchRoutingInterfaceDHCPDhcpOptions        `tfsdk:"dhcp_options"`
	DhcpRelayServerIps   types.Set                                                `tfsdk:"dhcp_relay_server_ips"`
	DnsCustomNameservers types.List                                               `tfsdk:"dns_custom_nameservers"`
	FixedIpAssignments   []DataSourceSwitchRoutingInterfaceDHCPFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	ReservedIpRanges     []DataSourceSwitchRoutingInterfaceDHCPReservedIpRanges   `tfsdk:"reserved_ip_ranges"`
}

type DataSourceSwitchRoutingInterfaceDHCPDhcpOptions struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type DataSourceSwitchRoutingInterfaceDHCPFixedIpAssignments struct {
	Ip   types.String `tfsdk:"ip"`
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
}

type DataSourceSwitchRoutingInterfaceDHCPReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceSwitchRoutingInterfaceDHCP) getPath() string {
	return fmt.Sprintf("/devices/%v/switch/routing/interfaces/%v/dhcp", url.QueryEscape(data.Serial.ValueString()), url.QueryEscape(data.InterfaceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceSwitchRoutingInterfaceDHCP) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("bootFileName"); value.Exists() && value.Value() != nil {
		data.BootFileName = types.StringValue(value.String())
	} else {
		data.BootFileName = types.StringNull()
	}
	if value := res.Get("bootNextServer"); value.Exists() && value.Value() != nil {
		data.BootNextServer = types.StringValue(value.String())
	} else {
		data.BootNextServer = types.StringNull()
	}
	if value := res.Get("bootOptionsEnabled"); value.Exists() && value.Value() != nil {
		data.BootOptionsEnabled = types.BoolValue(value.Bool())
	} else {
		data.BootOptionsEnabled = types.BoolNull()
	}
	if value := res.Get("dhcpLeaseTime"); value.Exists() && value.Value() != nil {
		data.DhcpLeaseTime = types.StringValue(value.String())
	} else {
		data.DhcpLeaseTime = types.StringNull()
	}
	if value := res.Get("dhcpMode"); value.Exists() && value.Value() != nil {
		data.DhcpMode = types.StringValue(value.String())
	} else {
		data.DhcpMode = types.StringNull()
	}
	if value := res.Get("dnsNameserversOption"); value.Exists() && value.Value() != nil {
		data.DnsNameserversOption = types.StringValue(value.String())
	} else {
		data.DnsNameserversOption = types.StringNull()
	}
	if value := res.Get("dhcpOptions"); value.Exists() && value.Value() != nil {
		data.DhcpOptions = make([]DataSourceSwitchRoutingInterfaceDHCPDhcpOptions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceSwitchRoutingInterfaceDHCPDhcpOptions{}
			if value := res.Get("code"); value.Exists() && value.Value() != nil {
				data.Code = types.StringValue(value.String())
			} else {
				data.Code = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && value.Value() != nil {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("value"); value.Exists() && value.Value() != nil {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).DhcpOptions = append((*parent).DhcpOptions, data)
			return true
		})
	}
	if value := res.Get("dhcpRelayServerIps"); value.Exists() && value.Value() != nil {
		data.DhcpRelayServerIps = helpers.GetStringSet(value.Array())
	} else {
		data.DhcpRelayServerIps = types.SetNull(types.StringType)
	}
	if value := res.Get("dnsCustomNameservers"); value.Exists() && value.Value() != nil {
		data.DnsCustomNameservers = helpers.GetStringList(value.Array())
	} else {
		data.DnsCustomNameservers = types.ListNull(types.StringType)
	}
	if value := res.Get("fixedIpAssignments"); value.Exists() && value.Value() != nil {
		data.FixedIpAssignments = make([]DataSourceSwitchRoutingInterfaceDHCPFixedIpAssignments, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceSwitchRoutingInterfaceDHCPFixedIpAssignments{}
			if value := res.Get("ip"); value.Exists() && value.Value() != nil {
				data.Ip = types.StringValue(value.String())
			} else {
				data.Ip = types.StringNull()
			}
			if value := res.Get("mac"); value.Exists() && value.Value() != nil {
				data.Mac = types.StringValue(value.String())
			} else {
				data.Mac = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).FixedIpAssignments = append((*parent).FixedIpAssignments, data)
			return true
		})
	}
	if value := res.Get("reservedIpRanges"); value.Exists() && value.Value() != nil {
		data.ReservedIpRanges = make([]DataSourceSwitchRoutingInterfaceDHCPReservedIpRanges, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceSwitchRoutingInterfaceDHCPReservedIpRanges{}
			if value := res.Get("comment"); value.Exists() && value.Value() != nil {
				data.Comment = types.StringValue(value.String())
			} else {
				data.Comment = types.StringNull()
			}
			if value := res.Get("end"); value.Exists() && value.Value() != nil {
				data.End = types.StringValue(value.String())
			} else {
				data.End = types.StringNull()
			}
			if value := res.Get("start"); value.Exists() && value.Value() != nil {
				data.Start = types.StringValue(value.String())
			} else {
				data.Start = types.StringNull()
			}
			(*parent).ReservedIpRanges = append((*parent).ReservedIpRanges, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
