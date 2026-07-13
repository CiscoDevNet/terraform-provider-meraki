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
// It emits a separate model - DataSourceApplianceVPNBGP - used only by the data source, always including every
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

type DataSourceApplianceVPNBGP struct {
	Id            types.String                         `tfsdk:"id"`
	NetworkId     types.String                         `tfsdk:"network_id"`
	AsNumber      types.Int64                          `tfsdk:"as_number"`
	Enabled       types.Bool                           `tfsdk:"enabled"`
	IbgpHoldTimer types.Int64                          `tfsdk:"ibgp_hold_timer"`
	Neighbors     []DataSourceApplianceVPNBGPNeighbors `tfsdk:"neighbors"`
}

type DataSourceApplianceVPNBGPNeighbors struct {
	AllowTransit                    types.Bool   `tfsdk:"allow_transit"`
	EbgpHoldTimer                   types.Int64  `tfsdk:"ebgp_hold_timer"`
	EbgpMultihop                    types.Int64  `tfsdk:"ebgp_multihop"`
	Ip                              types.String `tfsdk:"ip"`
	MultiExitDiscriminator          types.Int64  `tfsdk:"multi_exit_discriminator"`
	NextHopIp                       types.String `tfsdk:"next_hop_ip"`
	ReceiveLimit                    types.Int64  `tfsdk:"receive_limit"`
	RemoteAsNumber                  types.Int64  `tfsdk:"remote_as_number"`
	SourceInterface                 types.String `tfsdk:"source_interface"`
	Weight                          types.Int64  `tfsdk:"weight"`
	AuthenticationPassword          types.String `tfsdk:"authentication_password"`
	AuthenticationPasswordWo        types.String `tfsdk:"authentication_password_wo"`
	AuthenticationPasswordWoVersion types.Int64  `tfsdk:"authentication_password_wo_version"`
	Ipv6Address                     types.String `tfsdk:"ipv6_address"`
	TtlSecurityEnabled              types.Bool   `tfsdk:"ttl_security_enabled"`
	PathPrepend                     types.List   `tfsdk:"path_prepend"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceVPNBGP) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/vpn/bgp", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceVPNBGP) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("asNumber"); value.Exists() && value.Value() != nil {
		data.AsNumber = types.Int64Value(value.Int())
	} else {
		data.AsNumber = types.Int64Null()
	}
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("ibgpHoldTimer"); value.Exists() && value.Value() != nil {
		data.IbgpHoldTimer = types.Int64Value(value.Int())
	} else {
		data.IbgpHoldTimer = types.Int64Null()
	}
	if value := res.Get("neighbors"); value.Exists() && value.Value() != nil {
		data.Neighbors = make([]DataSourceApplianceVPNBGPNeighbors, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceVPNBGPNeighbors{}
			if value := res.Get("allowTransit"); value.Exists() && value.Value() != nil {
				data.AllowTransit = types.BoolValue(value.Bool())
			} else {
				data.AllowTransit = types.BoolNull()
			}
			if value := res.Get("ebgpHoldTimer"); value.Exists() && value.Value() != nil {
				data.EbgpHoldTimer = types.Int64Value(value.Int())
			} else {
				data.EbgpHoldTimer = types.Int64Null()
			}
			if value := res.Get("ebgpMultihop"); value.Exists() && value.Value() != nil {
				data.EbgpMultihop = types.Int64Value(value.Int())
			} else {
				data.EbgpMultihop = types.Int64Null()
			}
			if value := res.Get("ip"); value.Exists() && value.Value() != nil {
				data.Ip = types.StringValue(value.String())
			} else {
				data.Ip = types.StringNull()
			}
			if value := res.Get("multiExitDiscriminator"); value.Exists() && value.Value() != nil {
				data.MultiExitDiscriminator = types.Int64Value(value.Int())
			} else {
				data.MultiExitDiscriminator = types.Int64Null()
			}
			if value := res.Get("nextHopIp"); value.Exists() && value.Value() != nil {
				data.NextHopIp = types.StringValue(value.String())
			} else {
				data.NextHopIp = types.StringNull()
			}
			if value := res.Get("receiveLimit"); value.Exists() && value.Value() != nil {
				data.ReceiveLimit = types.Int64Value(value.Int())
			} else {
				data.ReceiveLimit = types.Int64Null()
			}
			if value := res.Get("remoteAsNumber"); value.Exists() && value.Value() != nil {
				data.RemoteAsNumber = types.Int64Value(value.Int())
			} else {
				data.RemoteAsNumber = types.Int64Null()
			}
			if value := res.Get("sourceInterface"); value.Exists() && value.Value() != nil {
				data.SourceInterface = types.StringValue(value.String())
			} else {
				data.SourceInterface = types.StringNull()
			}
			if value := res.Get("weight"); value.Exists() && value.Value() != nil {
				data.Weight = types.Int64Value(value.Int())
			} else {
				data.Weight = types.Int64Null()
			}
			if value := res.Get("authentication.password"); value.Exists() && value.Value() != nil {
				data.AuthenticationPassword = types.StringValue(value.String())
			} else {
				data.AuthenticationPassword = types.StringNull()
			}
			if value := res.Get("ipv6.address"); value.Exists() && value.Value() != nil {
				data.Ipv6Address = types.StringValue(value.String())
			} else {
				data.Ipv6Address = types.StringNull()
			}
			if value := res.Get("ttlSecurity.enabled"); value.Exists() && value.Value() != nil {
				data.TtlSecurityEnabled = types.BoolValue(value.Bool())
			} else {
				data.TtlSecurityEnabled = types.BoolNull()
			}
			if value := res.Get("pathPrepend"); value.Exists() && value.Value() != nil {
				data.PathPrepend = helpers.GetInt64List(value.Array())
			} else {
				data.PathPrepend = types.ListNull(types.Int64Type)
			}
			(*parent).Neighbors = append((*parent).Neighbors, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
