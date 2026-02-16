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

type ApplianceVPNBGP struct {
	Id            types.String               `tfsdk:"id"`
	NetworkId     types.String               `tfsdk:"network_id"`
	AsNumber      types.Int64                `tfsdk:"as_number"`
	Enabled       types.Bool                 `tfsdk:"enabled"`
	IbgpHoldTimer types.Int64                `tfsdk:"ibgp_hold_timer"`
	Neighbors     []ApplianceVPNBGPNeighbors `tfsdk:"neighbors"`
}

type ApplianceVPNBGPNeighbors struct {
	AllowTransit           types.Bool   `tfsdk:"allow_transit"`
	EbgpHoldTimer          types.Int64  `tfsdk:"ebgp_hold_timer"`
	EbgpMultihop           types.Int64  `tfsdk:"ebgp_multihop"`
	Ip                     types.String `tfsdk:"ip"`
	MultiExitDiscriminator types.Int64  `tfsdk:"multi_exit_discriminator"`
	NextHopIp              types.String `tfsdk:"next_hop_ip"`
	ReceiveLimit           types.Int64  `tfsdk:"receive_limit"`
	RemoteAsNumber         types.Int64  `tfsdk:"remote_as_number"`
	SourceInterface        types.String `tfsdk:"source_interface"`
	Weight                 types.Int64  `tfsdk:"weight"`
	AuthenticationPassword types.String `tfsdk:"authentication_password"`
	Ipv6Address            types.String `tfsdk:"ipv6_address"`
	TtlSecurityEnabled     types.Bool   `tfsdk:"ttl_security_enabled"`
	PathPrepend            types.List   `tfsdk:"path_prepend"`
}

type ApplianceVPNBGPIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceVPNBGP) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/vpn/bgp", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceVPNBGP) toBody(ctx context.Context, state ApplianceVPNBGP) string {
	body := ""
	if !data.AsNumber.IsNull() {
		body, _ = sjson.Set(body, "asNumber", data.AsNumber.ValueInt64())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.IbgpHoldTimer.IsNull() {
		body, _ = sjson.Set(body, "ibgpHoldTimer", data.IbgpHoldTimer.ValueInt64())
	}
	if data.Neighbors != nil {
		body, _ = sjson.Set(body, "neighbors", []interface{}{})
		for _, item := range data.Neighbors {
			itemBody := ""
			if !item.AllowTransit.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "allowTransit", item.AllowTransit.ValueBool())
			}
			if !item.EbgpHoldTimer.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ebgpHoldTimer", item.EbgpHoldTimer.ValueInt64())
			}
			if !item.EbgpMultihop.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ebgpMultihop", item.EbgpMultihop.ValueInt64())
			}
			if !item.Ip.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ip", item.Ip.ValueString())
			}
			if !item.MultiExitDiscriminator.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "multiExitDiscriminator", item.MultiExitDiscriminator.ValueInt64())
			}
			if !item.NextHopIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nextHopIp", item.NextHopIp.ValueString())
			}
			if !item.ReceiveLimit.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "receiveLimit", item.ReceiveLimit.ValueInt64())
			}
			if !item.RemoteAsNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "remoteAsNumber", item.RemoteAsNumber.ValueInt64())
			}
			if !item.SourceInterface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sourceInterface", item.SourceInterface.ValueString())
			}
			if !item.Weight.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "weight", item.Weight.ValueInt64())
			}
			if !item.AuthenticationPassword.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authentication.password", item.AuthenticationPassword.ValueString())
			}
			if !item.Ipv6Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6.address", item.Ipv6Address.ValueString())
			}
			if !item.TtlSecurityEnabled.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ttlSecurity.enabled", item.TtlSecurityEnabled.ValueBool())
			}
			if !item.PathPrepend.IsNull() {
				var values []int64
				item.PathPrepend.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "pathPrepend", values)
			}
			body, _ = sjson.SetRaw(body, "neighbors.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceVPNBGP) fromBody(ctx context.Context, res meraki.Res) {
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
		data.Neighbors = make([]ApplianceVPNBGPNeighbors, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceVPNBGPNeighbors{}
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceVPNBGP) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("asNumber"); value.Exists() && !data.AsNumber.IsNull() {
		data.AsNumber = types.Int64Value(value.Int())
	} else {
		data.AsNumber = types.Int64Null()
	}
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("ibgpHoldTimer"); value.Exists() && !data.IbgpHoldTimer.IsNull() {
		data.IbgpHoldTimer = types.Int64Value(value.Int())
	} else {
		data.IbgpHoldTimer = types.Int64Null()
	}
	for i := 0; i < len(data.Neighbors); i++ {
		keys := [...]string{"ip"}
		keyValues := [...]string{data.Neighbors[i].Ip.ValueString()}

		parent := &data
		data := (*parent).Neighbors[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("neighbors").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Neighbors[%d] = %+v",
				i,
				(*parent).Neighbors[i],
			))
			(*parent).Neighbors = slices.Delete((*parent).Neighbors, i, i+1)
			i--

			continue
		}
		if value := res.Get("allowTransit"); value.Exists() && !data.AllowTransit.IsNull() {
			data.AllowTransit = types.BoolValue(value.Bool())
		} else {
			data.AllowTransit = types.BoolNull()
		}
		if value := res.Get("ebgpHoldTimer"); value.Exists() && !data.EbgpHoldTimer.IsNull() {
			data.EbgpHoldTimer = types.Int64Value(value.Int())
		} else {
			data.EbgpHoldTimer = types.Int64Null()
		}
		if value := res.Get("ebgpMultihop"); value.Exists() && !data.EbgpMultihop.IsNull() {
			data.EbgpMultihop = types.Int64Value(value.Int())
		} else {
			data.EbgpMultihop = types.Int64Null()
		}
		if value := res.Get("ip"); value.Exists() && !data.Ip.IsNull() {
			data.Ip = types.StringValue(value.String())
		} else {
			data.Ip = types.StringNull()
		}
		if value := res.Get("multiExitDiscriminator"); value.Exists() && !data.MultiExitDiscriminator.IsNull() {
			data.MultiExitDiscriminator = types.Int64Value(value.Int())
		} else {
			data.MultiExitDiscriminator = types.Int64Null()
		}
		if value := res.Get("nextHopIp"); value.Exists() && !data.NextHopIp.IsNull() {
			data.NextHopIp = types.StringValue(value.String())
		} else {
			data.NextHopIp = types.StringNull()
		}
		if value := res.Get("receiveLimit"); value.Exists() && !data.ReceiveLimit.IsNull() {
			data.ReceiveLimit = types.Int64Value(value.Int())
		} else {
			data.ReceiveLimit = types.Int64Null()
		}
		if value := res.Get("remoteAsNumber"); value.Exists() && !data.RemoteAsNumber.IsNull() {
			data.RemoteAsNumber = types.Int64Value(value.Int())
		} else {
			data.RemoteAsNumber = types.Int64Null()
		}
		if value := res.Get("sourceInterface"); value.Exists() && !data.SourceInterface.IsNull() {
			data.SourceInterface = types.StringValue(value.String())
		} else {
			data.SourceInterface = types.StringNull()
		}
		if value := res.Get("weight"); value.Exists() && !data.Weight.IsNull() {
			data.Weight = types.Int64Value(value.Int())
		} else {
			data.Weight = types.Int64Null()
		}
		if value := res.Get("authentication.password"); value.Exists() && !data.AuthenticationPassword.IsNull() {
			data.AuthenticationPassword = types.StringValue(value.String())
		} else {
			data.AuthenticationPassword = types.StringNull()
		}
		if value := res.Get("ipv6.address"); value.Exists() && !data.Ipv6Address.IsNull() {
			data.Ipv6Address = types.StringValue(value.String())
		} else {
			data.Ipv6Address = types.StringNull()
		}
		if value := res.Get("ttlSecurity.enabled"); value.Exists() && !data.TtlSecurityEnabled.IsNull() {
			data.TtlSecurityEnabled = types.BoolValue(value.Bool())
		} else {
			data.TtlSecurityEnabled = types.BoolNull()
		}
		if value := res.Get("pathPrepend"); value.Exists() && !data.PathPrepend.IsNull() {
			data.PathPrepend = helpers.GetInt64List(value.Array())
		} else {
			data.PathPrepend = types.ListNull(types.Int64Type)
		}
		(*parent).Neighbors[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceVPNBGP) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceVPNBGPIdentity) toIdentity(ctx context.Context, plan *ApplianceVPNBGP) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceVPNBGP) fromIdentity(ctx context.Context, identity *ApplianceVPNBGPIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceVPNBGP) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
