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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessSSIDVPN struct {
	Id                        types.String                      `tfsdk:"id"`
	NetworkId                 types.String                      `tfsdk:"network_id"`
	Number                    types.String                      `tfsdk:"number"`
	ConcentratorNetworkId     types.String                      `tfsdk:"concentrator_network_id"`
	ConcentratorVlanId        types.Int64                       `tfsdk:"concentrator_vlan_id"`
	FailoverHeartbeatInterval types.Int64                       `tfsdk:"failover_heartbeat_interval"`
	FailoverIdleTimeout       types.Int64                       `tfsdk:"failover_idle_timeout"`
	FailoverRequestIp         types.String                      `tfsdk:"failover_request_ip"`
	SplitTunnelEnabled        types.Bool                        `tfsdk:"split_tunnel_enabled"`
	SplitTunnelRules          []WirelessSSIDVPNSplitTunnelRules `tfsdk:"split_tunnel_rules"`
}

type WirelessSSIDVPNSplitTunnelRules struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessSSIDVPN) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/vpn", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSSIDVPN) toBody(ctx context.Context, state WirelessSSIDVPN) string {
	body := ""
	if !data.ConcentratorNetworkId.IsNull() {
		body, _ = sjson.Set(body, "concentrator.networkId", data.ConcentratorNetworkId.ValueString())
	}
	if !data.ConcentratorVlanId.IsNull() {
		body, _ = sjson.Set(body, "concentrator.vlanId", data.ConcentratorVlanId.ValueInt64())
	}
	if !data.FailoverHeartbeatInterval.IsNull() {
		body, _ = sjson.Set(body, "failover.heartbeatInterval", data.FailoverHeartbeatInterval.ValueInt64())
	}
	if !data.FailoverIdleTimeout.IsNull() {
		body, _ = sjson.Set(body, "failover.idleTimeout", data.FailoverIdleTimeout.ValueInt64())
	}
	if !data.FailoverRequestIp.IsNull() {
		body, _ = sjson.Set(body, "failover.requestIp", data.FailoverRequestIp.ValueString())
	}
	if !data.SplitTunnelEnabled.IsNull() {
		body, _ = sjson.Set(body, "splitTunnel.enabled", data.SplitTunnelEnabled.ValueBool())
	}
	if len(data.SplitTunnelRules) > 0 {
		body, _ = sjson.Set(body, "splitTunnel.rules", []interface{}{})
		for _, item := range data.SplitTunnelRules {
			itemBody := ""
			if !item.Comment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "comment", item.Comment.ValueString())
			}
			if !item.DestCidr.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destCidr", item.DestCidr.ValueString())
			}
			if !item.DestPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destPort", item.DestPort.ValueString())
			}
			if !item.Policy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "policy", item.Policy.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			body, _ = sjson.SetRaw(body, "splitTunnel.rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSSIDVPN) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("concentrator.networkId"); value.Exists() && value.Value() != nil {
		data.ConcentratorNetworkId = types.StringValue(value.String())
	} else {
		data.ConcentratorNetworkId = types.StringNull()
	}
	if value := res.Get("concentrator.vlanId"); value.Exists() && value.Value() != nil {
		data.ConcentratorVlanId = types.Int64Value(value.Int())
	} else {
		data.ConcentratorVlanId = types.Int64Null()
	}
	if value := res.Get("failover.heartbeatInterval"); value.Exists() && value.Value() != nil {
		data.FailoverHeartbeatInterval = types.Int64Value(value.Int())
	} else {
		data.FailoverHeartbeatInterval = types.Int64Null()
	}
	if value := res.Get("failover.idleTimeout"); value.Exists() && value.Value() != nil {
		data.FailoverIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.FailoverIdleTimeout = types.Int64Null()
	}
	if value := res.Get("failover.requestIp"); value.Exists() && value.Value() != nil {
		data.FailoverRequestIp = types.StringValue(value.String())
	} else {
		data.FailoverRequestIp = types.StringNull()
	}
	if value := res.Get("splitTunnel.enabled"); value.Exists() && value.Value() != nil {
		data.SplitTunnelEnabled = types.BoolValue(value.Bool())
	} else {
		data.SplitTunnelEnabled = types.BoolNull()
	}
	if value := res.Get("splitTunnel.rules"); value.Exists() && value.Value() != nil {
		data.SplitTunnelRules = make([]WirelessSSIDVPNSplitTunnelRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDVPNSplitTunnelRules{}
			if value := res.Get("comment"); value.Exists() && value.Value() != nil {
				data.Comment = types.StringValue(value.String())
			} else {
				data.Comment = types.StringNull()
			}
			if value := res.Get("destCidr"); value.Exists() && value.Value() != nil {
				data.DestCidr = types.StringValue(value.String())
			} else {
				data.DestCidr = types.StringNull()
			}
			if value := res.Get("destPort"); value.Exists() && value.Value() != nil {
				data.DestPort = types.StringValue(value.String())
			} else {
				data.DestPort = types.StringNull()
			}
			if value := res.Get("policy"); value.Exists() && value.Value() != nil {
				data.Policy = types.StringValue(value.String())
			} else {
				data.Policy = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			(*parent).SplitTunnelRules = append((*parent).SplitTunnelRules, data)
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
func (data *WirelessSSIDVPN) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("concentrator.networkId"); value.Exists() && !data.ConcentratorNetworkId.IsNull() {
		data.ConcentratorNetworkId = types.StringValue(value.String())
	} else {
		data.ConcentratorNetworkId = types.StringNull()
	}
	if value := res.Get("concentrator.vlanId"); value.Exists() && !data.ConcentratorVlanId.IsNull() {
		data.ConcentratorVlanId = types.Int64Value(value.Int())
	} else {
		data.ConcentratorVlanId = types.Int64Null()
	}
	if value := res.Get("failover.heartbeatInterval"); value.Exists() && !data.FailoverHeartbeatInterval.IsNull() {
		data.FailoverHeartbeatInterval = types.Int64Value(value.Int())
	} else {
		data.FailoverHeartbeatInterval = types.Int64Null()
	}
	if value := res.Get("failover.idleTimeout"); value.Exists() && !data.FailoverIdleTimeout.IsNull() {
		data.FailoverIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.FailoverIdleTimeout = types.Int64Null()
	}
	if value := res.Get("failover.requestIp"); value.Exists() && !data.FailoverRequestIp.IsNull() {
		data.FailoverRequestIp = types.StringValue(value.String())
	} else {
		data.FailoverRequestIp = types.StringNull()
	}
	if value := res.Get("splitTunnel.enabled"); value.Exists() && !data.SplitTunnelEnabled.IsNull() {
		data.SplitTunnelEnabled = types.BoolValue(value.Bool())
	} else {
		data.SplitTunnelEnabled = types.BoolNull()
	}
	{
		l := len(res.Get("splitTunnel.rules").Array())
		tflog.Debug(ctx, fmt.Sprintf("splitTunnel.rules array resizing from %d to %d", len(data.SplitTunnelRules), l))
		if len(data.SplitTunnelRules) > l {
			data.SplitTunnelRules = data.SplitTunnelRules[:l]
		}
	}
	for i := range data.SplitTunnelRules {
		parent := &data
		data := (*parent).SplitTunnelRules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("splitTunnel.rules.%d", i))
		if value := res.Get("comment"); value.Exists() && !data.Comment.IsNull() {
			data.Comment = types.StringValue(value.String())
		} else {
			data.Comment = types.StringNull()
		}
		if value := res.Get("destCidr"); value.Exists() && !data.DestCidr.IsNull() {
			data.DestCidr = types.StringValue(value.String())
		} else {
			data.DestCidr = types.StringNull()
		}
		if value := res.Get("destPort"); value.Exists() && !data.DestPort.IsNull() {
			data.DestPort = types.StringValue(value.String())
		} else {
			data.DestPort = types.StringNull()
		}
		if value := res.Get("policy"); value.Exists() && !data.Policy.IsNull() {
			data.Policy = types.StringValue(value.String())
		} else {
			data.Policy = types.StringNull()
		}
		if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		(*parent).SplitTunnelRules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessSSIDVPN) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns
