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
// It emits a separate model - DataSourceApplianceThirdPartyVPNPeers - used only by the data source, always including every
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

type DataSourceApplianceThirdPartyVPNPeers struct {
	Id             types.String                                 `tfsdk:"id"`
	OrganizationId types.String                                 `tfsdk:"organization_id"`
	Peers          []DataSourceApplianceThirdPartyVPNPeersPeers `tfsdk:"peers"`
}

type DataSourceApplianceThirdPartyVPNPeersPeers struct {
	IkeVersion                         types.String `tfsdk:"ike_version"`
	IpsecPoliciesPreset                types.String `tfsdk:"ipsec_policies_preset"`
	IpVersion                          types.Int64  `tfsdk:"ip_version"`
	IsRouteBased                       types.Bool   `tfsdk:"is_route_based"`
	LocalId                            types.String `tfsdk:"local_id"`
	Name                               types.String `tfsdk:"name"`
	PeerId                             types.String `tfsdk:"peer_id"`
	PriorityInGroup                    types.Int64  `tfsdk:"priority_in_group"`
	PublicHostname                     types.String `tfsdk:"public_hostname"`
	PublicIp                           types.String `tfsdk:"public_ip"`
	RemoteId                           types.String `tfsdk:"remote_id"`
	Secret                             types.String `tfsdk:"secret"`
	SecretWo                           types.String `tfsdk:"secret_wo"`
	SecretWoVersion                    types.Int64  `tfsdk:"secret_wo_version"`
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

func (data DataSourceApplianceThirdPartyVPNPeers) getPath() string {
	return fmt.Sprintf("/organizations/%v/appliance/vpn/thirdPartyVPNPeers", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceThirdPartyVPNPeers) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("peers"); value.Exists() && value.Value() != nil {
		data.Peers = make([]DataSourceApplianceThirdPartyVPNPeersPeers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceThirdPartyVPNPeersPeers{}
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
			if value := res.Get("ipVersion"); value.Exists() && value.Value() != nil {
				data.IpVersion = types.Int64Value(value.Int())
			} else {
				data.IpVersion = types.Int64Null()
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
