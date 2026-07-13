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
// It emits a separate model - DataSourceWirelessSSIDVPN - used only by the data source, always including every
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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessSSIDVPN struct {
	Id                        types.String                                `tfsdk:"id"`
	NetworkId                 types.String                                `tfsdk:"network_id"`
	Number                    types.String                                `tfsdk:"number"`
	ConcentratorNetworkId     types.String                                `tfsdk:"concentrator_network_id"`
	ConcentratorVlanId        types.Int64                                 `tfsdk:"concentrator_vlan_id"`
	FailoverHeartbeatInterval types.Int64                                 `tfsdk:"failover_heartbeat_interval"`
	FailoverIdleTimeout       types.Int64                                 `tfsdk:"failover_idle_timeout"`
	FailoverRequestIp         types.String                                `tfsdk:"failover_request_ip"`
	SplitTunnelEnabled        types.Bool                                  `tfsdk:"split_tunnel_enabled"`
	SplitTunnelRules          []DataSourceWirelessSSIDVPNSplitTunnelRules `tfsdk:"split_tunnel_rules"`
}

type DataSourceWirelessSSIDVPNSplitTunnelRules struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessSSIDVPN) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/vpn", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessSSIDVPN) fromBody(ctx context.Context, res meraki.Res) {
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
		data.SplitTunnelRules = make([]DataSourceWirelessSSIDVPNSplitTunnelRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDVPNSplitTunnelRules{}
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
