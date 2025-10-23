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

type ApplianceVLANDHCP struct {
	Id                     types.String                        `tfsdk:"id"`
	NetworkId              types.String                        `tfsdk:"network_id"`
	VlanId                 types.String                        `tfsdk:"vlan_id"`
	DhcpBootFilename       types.String                        `tfsdk:"dhcp_boot_filename"`
	DhcpBootNextServer     types.String                        `tfsdk:"dhcp_boot_next_server"`
	DhcpBootOptionsEnabled types.Bool                          `tfsdk:"dhcp_boot_options_enabled"`
	DhcpHandling           types.String                        `tfsdk:"dhcp_handling"`
	DhcpLeaseTime          types.String                        `tfsdk:"dhcp_lease_time"`
	DnsNameservers         types.String                        `tfsdk:"dns_nameservers"`
	MandatoryDhcpEnabled   types.Bool                          `tfsdk:"mandatory_dhcp_enabled"`
	DhcpOptions            []ApplianceVLANDHCPDhcpOptions      `tfsdk:"dhcp_options"`
	DhcpRelayServerIps     types.List                          `tfsdk:"dhcp_relay_server_ips"`
	ReservedIpRanges       []ApplianceVLANDHCPReservedIpRanges `tfsdk:"reserved_ip_ranges"`
}

type ApplianceVLANDHCPDhcpOptions struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ApplianceVLANDHCPReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceVLANDHCP) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/vlans", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceVLANDHCP) toBody(ctx context.Context, state ApplianceVLANDHCP) string {
	body := ""
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "id", data.VlanId.ValueString())
	}
	if !data.DhcpBootFilename.IsNull() {
		body, _ = sjson.Set(body, "dhcpBootFilename", data.DhcpBootFilename.ValueString())
	}
	if !data.DhcpBootNextServer.IsNull() {
		body, _ = sjson.Set(body, "dhcpBootNextServer", data.DhcpBootNextServer.ValueString())
	}
	if !data.DhcpBootOptionsEnabled.IsNull() {
		body, _ = sjson.Set(body, "dhcpBootOptionsEnabled", data.DhcpBootOptionsEnabled.ValueBool())
	}
	if !data.DhcpHandling.IsNull() {
		body, _ = sjson.Set(body, "dhcpHandling", data.DhcpHandling.ValueString())
	}
	if !data.DhcpLeaseTime.IsNull() {
		body, _ = sjson.Set(body, "dhcpLeaseTime", data.DhcpLeaseTime.ValueString())
	}
	if !data.DnsNameservers.IsNull() {
		body, _ = sjson.Set(body, "dnsNameservers", data.DnsNameservers.ValueString())
	}
	if !data.MandatoryDhcpEnabled.IsNull() {
		body, _ = sjson.Set(body, "mandatoryDhcp.enabled", data.MandatoryDhcpEnabled.ValueBool())
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

func (data *ApplianceVLANDHCP) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("id"); value.Exists() && value.Value() != nil {
		data.VlanId = types.StringValue(value.String())
	} else {
		data.VlanId = types.StringNull()
	}
	if value := res.Get("dhcpBootFilename"); value.Exists() && value.Value() != nil {
		data.DhcpBootFilename = types.StringValue(value.String())
	} else {
		data.DhcpBootFilename = types.StringNull()
	}
	if value := res.Get("dhcpBootNextServer"); value.Exists() && value.Value() != nil {
		data.DhcpBootNextServer = types.StringValue(value.String())
	} else {
		data.DhcpBootNextServer = types.StringNull()
	}
	if value := res.Get("dhcpBootOptionsEnabled"); value.Exists() && value.Value() != nil {
		data.DhcpBootOptionsEnabled = types.BoolValue(value.Bool())
	} else {
		data.DhcpBootOptionsEnabled = types.BoolNull()
	}
	if value := res.Get("dhcpHandling"); value.Exists() && value.Value() != nil {
		data.DhcpHandling = types.StringValue(value.String())
	} else {
		data.DhcpHandling = types.StringNull()
	}
	if value := res.Get("dhcpLeaseTime"); value.Exists() && value.Value() != nil {
		data.DhcpLeaseTime = types.StringValue(value.String())
	} else {
		data.DhcpLeaseTime = types.StringNull()
	}
	if value := res.Get("dnsNameservers"); value.Exists() && value.Value() != nil {
		data.DnsNameservers = types.StringValue(value.String())
	} else {
		data.DnsNameservers = types.StringNull()
	}
	if value := res.Get("mandatoryDhcp.enabled"); value.Exists() && value.Value() != nil {
		data.MandatoryDhcpEnabled = types.BoolValue(value.Bool())
	} else {
		data.MandatoryDhcpEnabled = types.BoolNull()
	}
	if value := res.Get("dhcpOptions"); value.Exists() && value.Value() != nil {
		data.DhcpOptions = make([]ApplianceVLANDHCPDhcpOptions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceVLANDHCPDhcpOptions{}
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
		data.DhcpRelayServerIps = helpers.GetStringList(value.Array())
	} else {
		data.DhcpRelayServerIps = types.ListNull(types.StringType)
	}
	if value := res.Get("reservedIpRanges"); value.Exists() && value.Value() != nil {
		data.ReservedIpRanges = make([]ApplianceVLANDHCPReservedIpRanges, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceVLANDHCPReservedIpRanges{}
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
func (data *ApplianceVLANDHCP) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("id"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.StringValue(value.String())
	} else {
		data.VlanId = types.StringNull()
	}
	if value := res.Get("dhcpBootFilename"); value.Exists() && !data.DhcpBootFilename.IsNull() {
		data.DhcpBootFilename = types.StringValue(value.String())
	} else {
		data.DhcpBootFilename = types.StringNull()
	}
	if value := res.Get("dhcpBootNextServer"); value.Exists() && !data.DhcpBootNextServer.IsNull() {
		data.DhcpBootNextServer = types.StringValue(value.String())
	} else {
		data.DhcpBootNextServer = types.StringNull()
	}
	if value := res.Get("dhcpBootOptionsEnabled"); value.Exists() && !data.DhcpBootOptionsEnabled.IsNull() {
		data.DhcpBootOptionsEnabled = types.BoolValue(value.Bool())
	} else {
		data.DhcpBootOptionsEnabled = types.BoolNull()
	}
	if value := res.Get("dhcpHandling"); value.Exists() && !data.DhcpHandling.IsNull() {
		data.DhcpHandling = types.StringValue(value.String())
	} else {
		data.DhcpHandling = types.StringNull()
	}
	if value := res.Get("dhcpLeaseTime"); value.Exists() && !data.DhcpLeaseTime.IsNull() {
		data.DhcpLeaseTime = types.StringValue(value.String())
	} else {
		data.DhcpLeaseTime = types.StringNull()
	}
	if value := res.Get("dnsNameservers"); value.Exists() && !data.DnsNameservers.IsNull() {
		data.DnsNameservers = types.StringValue(value.String())
	} else {
		data.DnsNameservers = types.StringNull()
	}
	if value := res.Get("mandatoryDhcp.enabled"); value.Exists() && !data.MandatoryDhcpEnabled.IsNull() {
		data.MandatoryDhcpEnabled = types.BoolValue(value.Bool())
	} else {
		data.MandatoryDhcpEnabled = types.BoolNull()
	}
	for i := 0; i < len(data.DhcpOptions); i++ {
		keys := [...]string{"code", "type", "value"}
		keyValues := [...]string{data.DhcpOptions[i].Code.ValueString(), data.DhcpOptions[i].Type.ValueString(), data.DhcpOptions[i].Value.ValueString()}

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
		data.DhcpRelayServerIps = helpers.GetStringList(value.Array())
	} else {
		data.DhcpRelayServerIps = types.ListNull(types.StringType)
	}
	for i := 0; i < len(data.ReservedIpRanges); i++ {
		keys := [...]string{"comment", "end", "start"}
		keyValues := [...]string{data.ReservedIpRanges[i].Comment.ValueString(), data.ReservedIpRanges[i].End.ValueString(), data.ReservedIpRanges[i].Start.ValueString()}

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
func (data *ApplianceVLANDHCP) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceVLANDHCP) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
