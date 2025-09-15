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

type SwitchRoutingInterfaceDHCP struct {
	Id                   types.String                                   `tfsdk:"id"`
	Serial               types.String                                   `tfsdk:"serial"`
	InterfaceId          types.String                                   `tfsdk:"interface_id"`
	BootFileName         types.String                                   `tfsdk:"boot_file_name"`
	BootNextServer       types.String                                   `tfsdk:"boot_next_server"`
	BootOptionsEnabled   types.Bool                                     `tfsdk:"boot_options_enabled"`
	DhcpLeaseTime        types.String                                   `tfsdk:"dhcp_lease_time"`
	DhcpMode             types.String                                   `tfsdk:"dhcp_mode"`
	DnsNameserversOption types.String                                   `tfsdk:"dns_nameservers_option"`
	DhcpOptions          []SwitchRoutingInterfaceDHCPDhcpOptions        `tfsdk:"dhcp_options"`
	DhcpRelayServerIps   types.Set                                      `tfsdk:"dhcp_relay_server_ips"`
	DnsCustomNameservers types.List                                     `tfsdk:"dns_custom_nameservers"`
	FixedIpAssignments   []SwitchRoutingInterfaceDHCPFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	ReservedIpRanges     []SwitchRoutingInterfaceDHCPReservedIpRanges   `tfsdk:"reserved_ip_ranges"`
}

type SwitchRoutingInterfaceDHCPDhcpOptions struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type SwitchRoutingInterfaceDHCPFixedIpAssignments struct {
	Ip   types.String `tfsdk:"ip"`
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
}

type SwitchRoutingInterfaceDHCPReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchRoutingInterfaceDHCP) getPath() string {
	return fmt.Sprintf("/devices/%v/switch/routing/interfaces/%v/dhcp", url.QueryEscape(data.Serial.ValueString()), url.QueryEscape(data.InterfaceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchRoutingInterfaceDHCP) toBody(ctx context.Context, state SwitchRoutingInterfaceDHCP) string {
	body := ""
	if !data.BootFileName.IsNull() {
		body, _ = sjson.Set(body, "bootFileName", data.BootFileName.ValueString())
	}
	if !data.BootNextServer.IsNull() {
		body, _ = sjson.Set(body, "bootNextServer", data.BootNextServer.ValueString())
	}
	if !data.BootOptionsEnabled.IsNull() {
		body, _ = sjson.Set(body, "bootOptionsEnabled", data.BootOptionsEnabled.ValueBool())
	}
	if !data.DhcpLeaseTime.IsNull() {
		body, _ = sjson.Set(body, "dhcpLeaseTime", data.DhcpLeaseTime.ValueString())
	}
	if !data.DhcpMode.IsNull() {
		body, _ = sjson.Set(body, "dhcpMode", data.DhcpMode.ValueString())
	}
	if !data.DnsNameserversOption.IsNull() {
		body, _ = sjson.Set(body, "dnsNameserversOption", data.DnsNameserversOption.ValueString())
	}
	if len(data.DhcpOptions) > 0 {
		body, _ = sjson.Set(body, "dhcpOptions", []interface{}{})
		for _, item := range data.DhcpOptions {
			itemBody := ""
			if !item.Code.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "code", item.Code.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Value.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dhcpOptions.-1", itemBody)
		}
	}
	if !data.DhcpRelayServerIps.IsNull() {
		var values []string
		data.DhcpRelayServerIps.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "dhcpRelayServerIps", values)
	}
	if !data.DnsCustomNameservers.IsNull() {
		var values []string
		data.DnsCustomNameservers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "dnsCustomNameservers", values)
	}
	if len(data.FixedIpAssignments) > 0 {
		body, _ = sjson.Set(body, "fixedIpAssignments", []interface{}{})
		for _, item := range data.FixedIpAssignments {
			itemBody := ""
			if !item.Ip.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ip", item.Ip.ValueString())
			}
			if !item.Mac.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "mac", item.Mac.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "fixedIpAssignments.-1", itemBody)
		}
	}
	if len(data.ReservedIpRanges) > 0 {
		body, _ = sjson.Set(body, "reservedIpRanges", []interface{}{})
		for _, item := range data.ReservedIpRanges {
			itemBody := ""
			if !item.Comment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "comment", item.Comment.ValueString())
			}
			if !item.End.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "end", item.End.ValueString())
			}
			if !item.Start.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "start", item.Start.ValueString())
			}
			body, _ = sjson.SetRaw(body, "reservedIpRanges.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchRoutingInterfaceDHCP) fromBody(ctx context.Context, res meraki.Res) {
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
		data.DhcpOptions = make([]SwitchRoutingInterfaceDHCPDhcpOptions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchRoutingInterfaceDHCPDhcpOptions{}
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
		data.FixedIpAssignments = make([]SwitchRoutingInterfaceDHCPFixedIpAssignments, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchRoutingInterfaceDHCPFixedIpAssignments{}
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
		data.ReservedIpRanges = make([]SwitchRoutingInterfaceDHCPReservedIpRanges, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchRoutingInterfaceDHCPReservedIpRanges{}
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchRoutingInterfaceDHCP) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("bootFileName"); value.Exists() && !data.BootFileName.IsNull() {
		data.BootFileName = types.StringValue(value.String())
	} else {
		data.BootFileName = types.StringNull()
	}
	if value := res.Get("bootNextServer"); value.Exists() && !data.BootNextServer.IsNull() {
		data.BootNextServer = types.StringValue(value.String())
	} else {
		data.BootNextServer = types.StringNull()
	}
	if value := res.Get("bootOptionsEnabled"); value.Exists() && !data.BootOptionsEnabled.IsNull() {
		data.BootOptionsEnabled = types.BoolValue(value.Bool())
	} else {
		data.BootOptionsEnabled = types.BoolNull()
	}
	if value := res.Get("dhcpLeaseTime"); value.Exists() && !data.DhcpLeaseTime.IsNull() {
		data.DhcpLeaseTime = types.StringValue(value.String())
	} else {
		data.DhcpLeaseTime = types.StringNull()
	}
	if value := res.Get("dhcpMode"); value.Exists() && !data.DhcpMode.IsNull() {
		data.DhcpMode = types.StringValue(value.String())
	} else {
		data.DhcpMode = types.StringNull()
	}
	if value := res.Get("dnsNameserversOption"); value.Exists() && !data.DnsNameserversOption.IsNull() {
		data.DnsNameserversOption = types.StringValue(value.String())
	} else {
		data.DnsNameserversOption = types.StringNull()
	}
	for i := 0; i < len(data.DhcpOptions); i++ {
		keys := [...]string{"code"}
		keyValues := [...]string{data.DhcpOptions[i].Code.ValueString()}

		parent := &data
		data := (*parent).DhcpOptions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dhcpOptions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DhcpOptions[%d] = %+v",
				i,
				(*parent).DhcpOptions[i],
			))
			(*parent).DhcpOptions = slices.Delete((*parent).DhcpOptions, i, i+1)
			i--

			continue
		}
		if value := res.Get("code"); value.Exists() && !data.Code.IsNull() {
			data.Code = types.StringValue(value.String())
		} else {
			data.Code = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
			data.Value = types.StringValue(value.String())
		} else {
			data.Value = types.StringNull()
		}
		(*parent).DhcpOptions[i] = data
	}
	if value := res.Get("dhcpRelayServerIps"); value.Exists() && !data.DhcpRelayServerIps.IsNull() {
		data.DhcpRelayServerIps = helpers.GetStringSet(value.Array())
	} else {
		data.DhcpRelayServerIps = types.SetNull(types.StringType)
	}
	if value := res.Get("dnsCustomNameservers"); value.Exists() && !data.DnsCustomNameservers.IsNull() {
		data.DnsCustomNameservers = helpers.GetStringList(value.Array())
	} else {
		data.DnsCustomNameservers = types.ListNull(types.StringType)
	}
	for i := 0; i < len(data.FixedIpAssignments); i++ {
		keys := [...]string{"ip"}
		keyValues := [...]string{data.FixedIpAssignments[i].Ip.ValueString()}

		parent := &data
		data := (*parent).FixedIpAssignments[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("fixedIpAssignments").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing FixedIpAssignments[%d] = %+v",
				i,
				(*parent).FixedIpAssignments[i],
			))
			(*parent).FixedIpAssignments = slices.Delete((*parent).FixedIpAssignments, i, i+1)
			i--

			continue
		}
		if value := res.Get("ip"); value.Exists() && !data.Ip.IsNull() {
			data.Ip = types.StringValue(value.String())
		} else {
			data.Ip = types.StringNull()
		}
		if value := res.Get("mac"); value.Exists() && !data.Mac.IsNull() {
			data.Mac = types.StringValue(value.String())
		} else {
			data.Mac = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).FixedIpAssignments[i] = data
	}
	for i := 0; i < len(data.ReservedIpRanges); i++ {
		keys := [...]string{"end", "start"}
		keyValues := [...]string{data.ReservedIpRanges[i].End.ValueString(), data.ReservedIpRanges[i].Start.ValueString()}

		parent := &data
		data := (*parent).ReservedIpRanges[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("reservedIpRanges").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ReservedIpRanges[%d] = %+v",
				i,
				(*parent).ReservedIpRanges[i],
			))
			(*parent).ReservedIpRanges = slices.Delete((*parent).ReservedIpRanges, i, i+1)
			i--

			continue
		}
		if value := res.Get("comment"); value.Exists() && !data.Comment.IsNull() {
			data.Comment = types.StringValue(value.String())
		} else {
			data.Comment = types.StringNull()
		}
		if value := res.Get("end"); value.Exists() && !data.End.IsNull() {
			data.End = types.StringValue(value.String())
		} else {
			data.End = types.StringNull()
		}
		if value := res.Get("start"); value.Exists() && !data.Start.IsNull() {
			data.Start = types.StringValue(value.String())
		} else {
			data.Start = types.StringNull()
		}
		(*parent).ReservedIpRanges[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SwitchRoutingInterfaceDHCP) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data SwitchRoutingInterfaceDHCP) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
