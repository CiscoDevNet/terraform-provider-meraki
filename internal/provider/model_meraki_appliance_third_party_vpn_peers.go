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

type ApplianceThirdPartyVPNPeers struct {
	Id             types.String                       `tfsdk:"id"`
	OrganizationId types.String                       `tfsdk:"organization_id"`
	Peers          []ApplianceThirdPartyVPNPeersPeers `tfsdk:"peers"`
}

type ApplianceThirdPartyVPNPeersPeers struct {
	IkeVersion                         types.String `tfsdk:"ike_version"`
	IpsecPoliciesPreset                types.String `tfsdk:"ipsec_policies_preset"`
	IsRouteBased                       types.Bool   `tfsdk:"is_route_based"`
	LocalId                            types.String `tfsdk:"local_id"`
	Name                               types.String `tfsdk:"name"`
	PeerId                             types.String `tfsdk:"peer_id"`
	PriorityInGroup                    types.Int64  `tfsdk:"priority_in_group"`
	PublicHostname                     types.String `tfsdk:"public_hostname"`
	PublicIp                           types.String `tfsdk:"public_ip"`
	RemoteId                           types.String `tfsdk:"remote_id"`
	Secret                             types.String `tfsdk:"secret"`
	EbgpNeighborEbgpHoldTimer          types.Int64  `tfsdk:"ebgp_neighbor_ebgp_hold_timer"`
	EbgpNeighborEbgpMultihop           types.Int64  `tfsdk:"ebgp_neighbor_ebgp_multihop"`
	EbgpNeighborIpVersion              types.Int64  `tfsdk:"ebgp_neighbor_ip_version"`
	EbgpNeighborMultiExitDiscriminator types.Int64  `tfsdk:"ebgp_neighbor_multi_exit_discriminator"`
	EbgpNeighborNeighborIp             types.String `tfsdk:"ebgp_neighbor_neighbor_ip"`
	EbgpNeighborRemoteAsNumber         types.Int64  `tfsdk:"ebgp_neighbor_remote_as_number"`
	EbgpNeighborSourceIp               types.String `tfsdk:"ebgp_neighbor_source_ip"`
	EbgpNeighborWeight                 types.Int64  `tfsdk:"ebgp_neighbor_weight"`
	EbgpNeighborPathPrepend            types.List   `tfsdk:"ebgp_neighbor_path_prepend"`
	GroupActiveActiveTunnel            types.Bool   `tfsdk:"group_active_active_tunnel"`
	GroupNumber                        types.Int64  `tfsdk:"group_number"`
	GroupFailoverDirectToInternet      types.Bool   `tfsdk:"group_failover_direct_to_internet"`
	IpsecPoliciesChildLifetime         types.Int64  `tfsdk:"ipsec_policies_child_lifetime"`
	IpsecPoliciesIkeLifetime           types.Int64  `tfsdk:"ipsec_policies_ike_lifetime"`
	IpsecPoliciesChildAuthAlgo         types.List   `tfsdk:"ipsec_policies_child_auth_algo"`
	IpsecPoliciesChildCipherAlgo       types.List   `tfsdk:"ipsec_policies_child_cipher_algo"`
	IpsecPoliciesChildPfsGroup         types.List   `tfsdk:"ipsec_policies_child_pfs_group"`
	IpsecPoliciesIkeAuthAlgo           types.List   `tfsdk:"ipsec_policies_ike_auth_algo"`
	IpsecPoliciesIkeCipherAlgo         types.List   `tfsdk:"ipsec_policies_ike_cipher_algo"`
	IpsecPoliciesIkeDiffieHellmanGroup types.List   `tfsdk:"ipsec_policies_ike_diffie_hellman_group"`
	IpsecPoliciesIkePrfAlgo            types.List   `tfsdk:"ipsec_policies_ike_prf_algo"`
	NetworkIds                         types.List   `tfsdk:"network_ids"`
	SlaPolicyId                        types.String `tfsdk:"sla_policy_id"`
	NetworkTags                        types.List   `tfsdk:"network_tags"`
	PrivateSubnets                     types.List   `tfsdk:"private_subnets"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceThirdPartyVPNPeers) getPath() string {
	return fmt.Sprintf("/organizations/%v/appliance/vpn/thirdPartyVPNPeers", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceThirdPartyVPNPeers) toBody(ctx context.Context, state ApplianceThirdPartyVPNPeers) string {
	body := ""
	{
		body, _ = sjson.Set(body, "peers", []interface{}{})
		for _, item := range data.Peers {
			itemBody := ""
			if !item.IkeVersion.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ikeVersion", item.IkeVersion.ValueString())
			}
			if !item.IpsecPoliciesPreset.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipsecPoliciesPreset", item.IpsecPoliciesPreset.ValueString())
			}
			if !item.IsRouteBased.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isRouteBased", item.IsRouteBased.ValueBool())
			}
			if !item.LocalId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "localId", item.LocalId.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.PeerId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "peerId", item.PeerId.ValueString())
			}
			if !item.PriorityInGroup.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "priorityInGroup", item.PriorityInGroup.ValueInt64())
			}
			if !item.PublicHostname.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "publicHostname", item.PublicHostname.ValueString())
			}
			if !item.PublicIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "publicIp", item.PublicIp.ValueString())
			}
			if !item.RemoteId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "remoteId", item.RemoteId.ValueString())
			}
			if !item.Secret.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secret", item.Secret.ValueString())
			}
			if !item.EbgpNeighborEbgpHoldTimer.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ebgpNeighbor.ebgpHoldTimer", item.EbgpNeighborEbgpHoldTimer.ValueInt64())
			}
			if !item.EbgpNeighborEbgpMultihop.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ebgpNeighbor.ebgpMultihop", item.EbgpNeighborEbgpMultihop.ValueInt64())
			}
			if !item.EbgpNeighborIpVersion.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ebgpNeighbor.ipVersion", item.EbgpNeighborIpVersion.ValueInt64())
			}
			if !item.EbgpNeighborMultiExitDiscriminator.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ebgpNeighbor.multiExitDiscriminator", item.EbgpNeighborMultiExitDiscriminator.ValueInt64())
			}
			if !item.EbgpNeighborNeighborIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ebgpNeighbor.neighborIp", item.EbgpNeighborNeighborIp.ValueString())
			}
			if !item.EbgpNeighborRemoteAsNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ebgpNeighbor.remoteAsNumber", item.EbgpNeighborRemoteAsNumber.ValueInt64())
			}
			if !item.EbgpNeighborSourceIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ebgpNeighbor.sourceIp", item.EbgpNeighborSourceIp.ValueString())
			}
			if !item.EbgpNeighborWeight.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ebgpNeighbor.weight", item.EbgpNeighborWeight.ValueInt64())
			}
			if !item.EbgpNeighborPathPrepend.IsNull() {
				var values []int64
				item.EbgpNeighborPathPrepend.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ebgpNeighbor.pathPrepend", values)
			}
			if !item.GroupActiveActiveTunnel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "group.activeActiveTunnel", item.GroupActiveActiveTunnel.ValueBool())
			}
			if !item.GroupNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "group.number", item.GroupNumber.ValueInt64())
			}
			if !item.GroupFailoverDirectToInternet.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "group.failover.directToInternet", item.GroupFailoverDirectToInternet.ValueBool())
			}
			if !item.IpsecPoliciesChildLifetime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipsecPolicies.childLifetime", item.IpsecPoliciesChildLifetime.ValueInt64())
			}
			if !item.IpsecPoliciesIkeLifetime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipsecPolicies.ikeLifetime", item.IpsecPoliciesIkeLifetime.ValueInt64())
			}
			if !item.IpsecPoliciesChildAuthAlgo.IsNull() {
				var values []string
				item.IpsecPoliciesChildAuthAlgo.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipsecPolicies.childAuthAlgo", values)
			}
			if !item.IpsecPoliciesChildCipherAlgo.IsNull() {
				var values []string
				item.IpsecPoliciesChildCipherAlgo.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipsecPolicies.childCipherAlgo", values)
			}
			if !item.IpsecPoliciesChildPfsGroup.IsNull() {
				var values []string
				item.IpsecPoliciesChildPfsGroup.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipsecPolicies.childPfsGroup", values)
			}
			if !item.IpsecPoliciesIkeAuthAlgo.IsNull() {
				var values []string
				item.IpsecPoliciesIkeAuthAlgo.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipsecPolicies.ikeAuthAlgo", values)
			}
			if !item.IpsecPoliciesIkeCipherAlgo.IsNull() {
				var values []string
				item.IpsecPoliciesIkeCipherAlgo.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipsecPolicies.ikeCipherAlgo", values)
			}
			if !item.IpsecPoliciesIkeDiffieHellmanGroup.IsNull() {
				var values []string
				item.IpsecPoliciesIkeDiffieHellmanGroup.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipsecPolicies.ikeDiffieHellmanGroup", values)
			}
			if !item.IpsecPoliciesIkePrfAlgo.IsNull() {
				var values []string
				item.IpsecPoliciesIkePrfAlgo.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "ipsecPolicies.ikePrfAlgo", values)
			}
			if !item.NetworkIds.IsNull() {
				var values []string
				item.NetworkIds.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "network.ids", values)
			}
			if !item.SlaPolicyId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "slaPolicy.id", item.SlaPolicyId.ValueString())
			}
			if !item.NetworkTags.IsNull() {
				var values []string
				item.NetworkTags.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "networkTags", values)
			}
			if !item.PrivateSubnets.IsNull() {
				var values []string
				item.PrivateSubnets.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "privateSubnets", values)
			}
			body, _ = sjson.SetRaw(body, "peers.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceThirdPartyVPNPeers) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("peers"); value.Exists() && value.Value() != nil {
		data.Peers = make([]ApplianceThirdPartyVPNPeersPeers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceThirdPartyVPNPeersPeers{}
			if value := res.Get("ikeVersion"); value.Exists() && value.Value() != nil {
				data.IkeVersion = types.StringValue(value.String())
			} else {
				data.IkeVersion = types.StringNull()
			}
			if value := res.Get("ipsecPoliciesPreset"); value.Exists() && value.Value() != nil {
				data.IpsecPoliciesPreset = types.StringValue(value.String())
			} else {
				data.IpsecPoliciesPreset = types.StringNull()
			}
			if value := res.Get("isRouteBased"); value.Exists() && value.Value() != nil {
				data.IsRouteBased = types.BoolValue(value.Bool())
			} else {
				data.IsRouteBased = types.BoolNull()
			}
			if value := res.Get("localId"); value.Exists() && value.Value() != nil {
				data.LocalId = types.StringValue(value.String())
			} else {
				data.LocalId = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("peerId"); value.Exists() && value.Value() != nil {
				data.PeerId = types.StringValue(value.String())
			} else {
				data.PeerId = types.StringNull()
			}
			if value := res.Get("priorityInGroup"); value.Exists() && value.Value() != nil {
				data.PriorityInGroup = types.Int64Value(value.Int())
			} else {
				data.PriorityInGroup = types.Int64Null()
			}
			if value := res.Get("publicHostname"); value.Exists() && value.Value() != nil {
				data.PublicHostname = types.StringValue(value.String())
			} else {
				data.PublicHostname = types.StringNull()
			}
			if value := res.Get("publicIp"); value.Exists() && value.Value() != nil {
				data.PublicIp = types.StringValue(value.String())
			} else {
				data.PublicIp = types.StringNull()
			}
			if value := res.Get("remoteId"); value.Exists() && value.Value() != nil {
				data.RemoteId = types.StringValue(value.String())
			} else {
				data.RemoteId = types.StringNull()
			}
			if value := res.Get("secret"); value.Exists() && value.Value() != nil {
				data.Secret = types.StringValue(value.String())
			} else {
				data.Secret = types.StringNull()
			}
			if value := res.Get("ebgpNeighbor.ebgpHoldTimer"); value.Exists() && value.Value() != nil {
				data.EbgpNeighborEbgpHoldTimer = types.Int64Value(value.Int())
			} else {
				data.EbgpNeighborEbgpHoldTimer = types.Int64Null()
			}
			if value := res.Get("ebgpNeighbor.ebgpMultihop"); value.Exists() && value.Value() != nil {
				data.EbgpNeighborEbgpMultihop = types.Int64Value(value.Int())
			} else {
				data.EbgpNeighborEbgpMultihop = types.Int64Null()
			}
			if value := res.Get("ebgpNeighbor.ipVersion"); value.Exists() && value.Value() != nil {
				data.EbgpNeighborIpVersion = types.Int64Value(value.Int())
			} else {
				data.EbgpNeighborIpVersion = types.Int64Null()
			}
			if value := res.Get("ebgpNeighbor.multiExitDiscriminator"); value.Exists() && value.Value() != nil {
				data.EbgpNeighborMultiExitDiscriminator = types.Int64Value(value.Int())
			} else {
				data.EbgpNeighborMultiExitDiscriminator = types.Int64Null()
			}
			if value := res.Get("ebgpNeighbor.neighborIp"); value.Exists() && value.Value() != nil {
				data.EbgpNeighborNeighborIp = types.StringValue(value.String())
			} else {
				data.EbgpNeighborNeighborIp = types.StringNull()
			}
			if value := res.Get("ebgpNeighbor.remoteAsNumber"); value.Exists() && value.Value() != nil {
				data.EbgpNeighborRemoteAsNumber = types.Int64Value(value.Int())
			} else {
				data.EbgpNeighborRemoteAsNumber = types.Int64Null()
			}
			if value := res.Get("ebgpNeighbor.sourceIp"); value.Exists() && value.Value() != nil {
				data.EbgpNeighborSourceIp = types.StringValue(value.String())
			} else {
				data.EbgpNeighborSourceIp = types.StringNull()
			}
			if value := res.Get("ebgpNeighbor.weight"); value.Exists() && value.Value() != nil {
				data.EbgpNeighborWeight = types.Int64Value(value.Int())
			} else {
				data.EbgpNeighborWeight = types.Int64Null()
			}
			if value := res.Get("ebgpNeighbor.pathPrepend"); value.Exists() && value.Value() != nil {
				data.EbgpNeighborPathPrepend = helpers.GetInt64List(value.Array())
			} else {
				data.EbgpNeighborPathPrepend = types.ListNull(types.Int64Type)
			}
			if value := res.Get("group.activeActiveTunnel"); value.Exists() && value.Value() != nil {
				data.GroupActiveActiveTunnel = types.BoolValue(value.Bool())
			} else {
				data.GroupActiveActiveTunnel = types.BoolNull()
			}
			if value := res.Get("group.number"); value.Exists() && value.Value() != nil {
				data.GroupNumber = types.Int64Value(value.Int())
			} else {
				data.GroupNumber = types.Int64Null()
			}
			if value := res.Get("group.failover.directToInternet"); value.Exists() && value.Value() != nil {
				data.GroupFailoverDirectToInternet = types.BoolValue(value.Bool())
			} else {
				data.GroupFailoverDirectToInternet = types.BoolNull()
			}
			if value := res.Get("ipsecPolicies.childLifetime"); value.Exists() && value.Value() != nil {
				data.IpsecPoliciesChildLifetime = types.Int64Value(value.Int())
			} else {
				data.IpsecPoliciesChildLifetime = types.Int64Null()
			}
			if value := res.Get("ipsecPolicies.ikeLifetime"); value.Exists() && value.Value() != nil {
				data.IpsecPoliciesIkeLifetime = types.Int64Value(value.Int())
			} else {
				data.IpsecPoliciesIkeLifetime = types.Int64Null()
			}
			if value := res.Get("ipsecPolicies.childAuthAlgo"); value.Exists() && value.Value() != nil {
				data.IpsecPoliciesChildAuthAlgo = helpers.GetStringList(value.Array())
			} else {
				data.IpsecPoliciesChildAuthAlgo = types.ListNull(types.StringType)
			}
			if value := res.Get("ipsecPolicies.childCipherAlgo"); value.Exists() && value.Value() != nil {
				data.IpsecPoliciesChildCipherAlgo = helpers.GetStringList(value.Array())
			} else {
				data.IpsecPoliciesChildCipherAlgo = types.ListNull(types.StringType)
			}
			if value := res.Get("ipsecPolicies.childPfsGroup"); value.Exists() && value.Value() != nil {
				data.IpsecPoliciesChildPfsGroup = helpers.GetStringList(value.Array())
			} else {
				data.IpsecPoliciesChildPfsGroup = types.ListNull(types.StringType)
			}
			if value := res.Get("ipsecPolicies.ikeAuthAlgo"); value.Exists() && value.Value() != nil {
				data.IpsecPoliciesIkeAuthAlgo = helpers.GetStringList(value.Array())
			} else {
				data.IpsecPoliciesIkeAuthAlgo = types.ListNull(types.StringType)
			}
			if value := res.Get("ipsecPolicies.ikeCipherAlgo"); value.Exists() && value.Value() != nil {
				data.IpsecPoliciesIkeCipherAlgo = helpers.GetStringList(value.Array())
			} else {
				data.IpsecPoliciesIkeCipherAlgo = types.ListNull(types.StringType)
			}
			if value := res.Get("ipsecPolicies.ikeDiffieHellmanGroup"); value.Exists() && value.Value() != nil {
				data.IpsecPoliciesIkeDiffieHellmanGroup = helpers.GetStringList(value.Array())
			} else {
				data.IpsecPoliciesIkeDiffieHellmanGroup = types.ListNull(types.StringType)
			}
			if value := res.Get("ipsecPolicies.ikePrfAlgo"); value.Exists() && value.Value() != nil {
				data.IpsecPoliciesIkePrfAlgo = helpers.GetStringList(value.Array())
			} else {
				data.IpsecPoliciesIkePrfAlgo = types.ListNull(types.StringType)
			}
			if value := res.Get("network.ids"); value.Exists() && value.Value() != nil {
				data.NetworkIds = helpers.GetStringList(value.Array())
			} else {
				data.NetworkIds = types.ListNull(types.StringType)
			}
			if value := res.Get("slaPolicy.id"); value.Exists() && value.Value() != nil {
				data.SlaPolicyId = types.StringValue(value.String())
			} else {
				data.SlaPolicyId = types.StringNull()
			}
			if value := res.Get("networkTags"); value.Exists() && value.Value() != nil {
				data.NetworkTags = helpers.GetStringList(value.Array())
			} else {
				data.NetworkTags = types.ListNull(types.StringType)
			}
			if value := res.Get("privateSubnets"); value.Exists() && value.Value() != nil {
				data.PrivateSubnets = helpers.GetStringList(value.Array())
			} else {
				data.PrivateSubnets = types.ListNull(types.StringType)
			}
			(*parent).Peers = append((*parent).Peers, data)
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
func (data *ApplianceThirdPartyVPNPeers) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Peers); i++ {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Peers[i].Name.ValueString()}

		parent := &data
		data := (*parent).Peers[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("peers").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Peers[%d] = %+v",
				i,
				(*parent).Peers[i],
			))
			(*parent).Peers = slices.Delete((*parent).Peers, i, i+1)
			i--

			continue
		}
		if value := res.Get("ikeVersion"); value.Exists() && !data.IkeVersion.IsNull() {
			data.IkeVersion = types.StringValue(value.String())
		} else {
			data.IkeVersion = types.StringNull()
		}
		if value := res.Get("ipsecPoliciesPreset"); value.Exists() && !data.IpsecPoliciesPreset.IsNull() {
			data.IpsecPoliciesPreset = types.StringValue(value.String())
		} else {
			data.IpsecPoliciesPreset = types.StringNull()
		}
		if value := res.Get("isRouteBased"); value.Exists() && !data.IsRouteBased.IsNull() {
			data.IsRouteBased = types.BoolValue(value.Bool())
		} else {
			data.IsRouteBased = types.BoolNull()
		}
		if value := res.Get("localId"); value.Exists() && !data.LocalId.IsNull() {
			data.LocalId = types.StringValue(value.String())
		} else {
			data.LocalId = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("peerId"); value.Exists() && !data.PeerId.IsNull() {
			data.PeerId = types.StringValue(value.String())
		} else {
			data.PeerId = types.StringNull()
		}
		if value := res.Get("priorityInGroup"); value.Exists() && !data.PriorityInGroup.IsNull() {
			data.PriorityInGroup = types.Int64Value(value.Int())
		} else {
			data.PriorityInGroup = types.Int64Null()
		}
		if value := res.Get("publicHostname"); value.Exists() && !data.PublicHostname.IsNull() {
			data.PublicHostname = types.StringValue(value.String())
		} else {
			data.PublicHostname = types.StringNull()
		}
		if value := res.Get("publicIp"); value.Exists() && !data.PublicIp.IsNull() {
			data.PublicIp = types.StringValue(value.String())
		} else {
			data.PublicIp = types.StringNull()
		}
		if value := res.Get("remoteId"); value.Exists() && !data.RemoteId.IsNull() {
			data.RemoteId = types.StringValue(value.String())
		} else {
			data.RemoteId = types.StringNull()
		}
		if value := res.Get("secret"); value.Exists() && !data.Secret.IsNull() {
			data.Secret = types.StringValue(value.String())
		} else {
			data.Secret = types.StringNull()
		}
		if value := res.Get("ebgpNeighbor.ebgpHoldTimer"); value.Exists() && !data.EbgpNeighborEbgpHoldTimer.IsNull() {
			data.EbgpNeighborEbgpHoldTimer = types.Int64Value(value.Int())
		} else {
			data.EbgpNeighborEbgpHoldTimer = types.Int64Null()
		}
		if value := res.Get("ebgpNeighbor.ebgpMultihop"); value.Exists() && !data.EbgpNeighborEbgpMultihop.IsNull() {
			data.EbgpNeighborEbgpMultihop = types.Int64Value(value.Int())
		} else {
			data.EbgpNeighborEbgpMultihop = types.Int64Null()
		}
		if value := res.Get("ebgpNeighbor.ipVersion"); value.Exists() && !data.EbgpNeighborIpVersion.IsNull() {
			data.EbgpNeighborIpVersion = types.Int64Value(value.Int())
		} else {
			data.EbgpNeighborIpVersion = types.Int64Null()
		}
		if value := res.Get("ebgpNeighbor.multiExitDiscriminator"); value.Exists() && !data.EbgpNeighborMultiExitDiscriminator.IsNull() {
			data.EbgpNeighborMultiExitDiscriminator = types.Int64Value(value.Int())
		} else {
			data.EbgpNeighborMultiExitDiscriminator = types.Int64Null()
		}
		if value := res.Get("ebgpNeighbor.neighborIp"); value.Exists() && !data.EbgpNeighborNeighborIp.IsNull() {
			data.EbgpNeighborNeighborIp = types.StringValue(value.String())
		} else {
			data.EbgpNeighborNeighborIp = types.StringNull()
		}
		if value := res.Get("ebgpNeighbor.remoteAsNumber"); value.Exists() && !data.EbgpNeighborRemoteAsNumber.IsNull() {
			data.EbgpNeighborRemoteAsNumber = types.Int64Value(value.Int())
		} else {
			data.EbgpNeighborRemoteAsNumber = types.Int64Null()
		}
		if value := res.Get("ebgpNeighbor.sourceIp"); value.Exists() && !data.EbgpNeighborSourceIp.IsNull() {
			data.EbgpNeighborSourceIp = types.StringValue(value.String())
		} else {
			data.EbgpNeighborSourceIp = types.StringNull()
		}
		if value := res.Get("ebgpNeighbor.weight"); value.Exists() && !data.EbgpNeighborWeight.IsNull() {
			data.EbgpNeighborWeight = types.Int64Value(value.Int())
		} else {
			data.EbgpNeighborWeight = types.Int64Null()
		}
		if value := res.Get("ebgpNeighbor.pathPrepend"); value.Exists() && !data.EbgpNeighborPathPrepend.IsNull() {
			data.EbgpNeighborPathPrepend = helpers.GetInt64List(value.Array())
		} else {
			data.EbgpNeighborPathPrepend = types.ListNull(types.Int64Type)
		}
		if value := res.Get("group.activeActiveTunnel"); value.Exists() && !data.GroupActiveActiveTunnel.IsNull() {
			data.GroupActiveActiveTunnel = types.BoolValue(value.Bool())
		} else {
			data.GroupActiveActiveTunnel = types.BoolNull()
		}
		if value := res.Get("group.number"); value.Exists() && !data.GroupNumber.IsNull() {
			data.GroupNumber = types.Int64Value(value.Int())
		} else {
			data.GroupNumber = types.Int64Null()
		}
		if value := res.Get("group.failover.directToInternet"); value.Exists() && !data.GroupFailoverDirectToInternet.IsNull() {
			data.GroupFailoverDirectToInternet = types.BoolValue(value.Bool())
		} else {
			data.GroupFailoverDirectToInternet = types.BoolNull()
		}
		if value := res.Get("ipsecPolicies.childLifetime"); value.Exists() && !data.IpsecPoliciesChildLifetime.IsNull() {
			data.IpsecPoliciesChildLifetime = types.Int64Value(value.Int())
		} else {
			data.IpsecPoliciesChildLifetime = types.Int64Null()
		}
		if value := res.Get("ipsecPolicies.ikeLifetime"); value.Exists() && !data.IpsecPoliciesIkeLifetime.IsNull() {
			data.IpsecPoliciesIkeLifetime = types.Int64Value(value.Int())
		} else {
			data.IpsecPoliciesIkeLifetime = types.Int64Null()
		}
		if value := res.Get("ipsecPolicies.childAuthAlgo"); value.Exists() && !data.IpsecPoliciesChildAuthAlgo.IsNull() {
			data.IpsecPoliciesChildAuthAlgo = helpers.GetStringList(value.Array())
		} else {
			data.IpsecPoliciesChildAuthAlgo = types.ListNull(types.StringType)
		}
		if value := res.Get("ipsecPolicies.childCipherAlgo"); value.Exists() && !data.IpsecPoliciesChildCipherAlgo.IsNull() {
			data.IpsecPoliciesChildCipherAlgo = helpers.GetStringList(value.Array())
		} else {
			data.IpsecPoliciesChildCipherAlgo = types.ListNull(types.StringType)
		}
		if value := res.Get("ipsecPolicies.childPfsGroup"); value.Exists() && !data.IpsecPoliciesChildPfsGroup.IsNull() {
			data.IpsecPoliciesChildPfsGroup = helpers.GetStringList(value.Array())
		} else {
			data.IpsecPoliciesChildPfsGroup = types.ListNull(types.StringType)
		}
		if value := res.Get("ipsecPolicies.ikeAuthAlgo"); value.Exists() && !data.IpsecPoliciesIkeAuthAlgo.IsNull() {
			data.IpsecPoliciesIkeAuthAlgo = helpers.GetStringList(value.Array())
		} else {
			data.IpsecPoliciesIkeAuthAlgo = types.ListNull(types.StringType)
		}
		if value := res.Get("ipsecPolicies.ikeCipherAlgo"); value.Exists() && !data.IpsecPoliciesIkeCipherAlgo.IsNull() {
			data.IpsecPoliciesIkeCipherAlgo = helpers.GetStringList(value.Array())
		} else {
			data.IpsecPoliciesIkeCipherAlgo = types.ListNull(types.StringType)
		}
		if value := res.Get("ipsecPolicies.ikeDiffieHellmanGroup"); value.Exists() && !data.IpsecPoliciesIkeDiffieHellmanGroup.IsNull() {
			data.IpsecPoliciesIkeDiffieHellmanGroup = helpers.GetStringList(value.Array())
		} else {
			data.IpsecPoliciesIkeDiffieHellmanGroup = types.ListNull(types.StringType)
		}
		if value := res.Get("ipsecPolicies.ikePrfAlgo"); value.Exists() && !data.IpsecPoliciesIkePrfAlgo.IsNull() {
			data.IpsecPoliciesIkePrfAlgo = helpers.GetStringList(value.Array())
		} else {
			data.IpsecPoliciesIkePrfAlgo = types.ListNull(types.StringType)
		}
		if value := res.Get("network.ids"); value.Exists() && !data.NetworkIds.IsNull() {
			data.NetworkIds = helpers.GetStringList(value.Array())
		} else {
			data.NetworkIds = types.ListNull(types.StringType)
		}
		if value := res.Get("slaPolicy.id"); value.Exists() && !data.SlaPolicyId.IsNull() {
			data.SlaPolicyId = types.StringValue(value.String())
		} else {
			data.SlaPolicyId = types.StringNull()
		}
		if value := res.Get("networkTags"); value.Exists() && !data.NetworkTags.IsNull() {
			data.NetworkTags = helpers.GetStringList(value.Array())
		} else {
			data.NetworkTags = types.ListNull(types.StringType)
		}
		if value := res.Get("privateSubnets"); value.Exists() && !data.PrivateSubnets.IsNull() {
			data.PrivateSubnets = helpers.GetStringList(value.Array())
		} else {
			data.PrivateSubnets = types.ListNull(types.StringType)
		}
		(*parent).Peers[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceThirdPartyVPNPeers) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceThirdPartyVPNPeers) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
