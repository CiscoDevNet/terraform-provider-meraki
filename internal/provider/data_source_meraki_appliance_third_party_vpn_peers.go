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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ApplianceThirdPartyVPNPeersDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceThirdPartyVPNPeersDataSource{}
)

func NewApplianceThirdPartyVPNPeersDataSource() datasource.DataSource {
	return &ApplianceThirdPartyVPNPeersDataSource{}
}

type ApplianceThirdPartyVPNPeersDataSource struct {
	client *meraki.Client
}

func (d *ApplianceThirdPartyVPNPeersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_third_party_vpn_peers"
}

func (d *ApplianceThirdPartyVPNPeersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Appliance Third Party VPN Peers` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
				Required:            true,
			},
			"peers": schema.ListNestedAttribute{
				MarkdownDescription: "The list of VPN peers",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ike_version": schema.StringAttribute{
							MarkdownDescription: "[optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to `1` when omitted.",
							Computed:            true,
						},
						"ipsec_policies_preset": schema.StringAttribute{
							MarkdownDescription: "One of the following available presets: `default`, `aws`, `azure`, `umbrella`, `zscaler`. If this is provided, the `ipsecPolicies` parameter is ignored.",
							Computed:            true,
						},
						"local_id": schema.StringAttribute{
							MarkdownDescription: "[optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the VPN peer",
							Computed:            true,
						},
						"public_ip": schema.StringAttribute{
							MarkdownDescription: "[optional] The public IP of the VPN peer",
							Computed:            true,
						},
						"remote_id": schema.StringAttribute{
							MarkdownDescription: "[optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "The shared secret with the VPN peer",
							Computed:            true,
						},
						"ipsec_policies_child_lifetime": schema.Int64Attribute{
							MarkdownDescription: "The lifetime of the Phase 2 SA in seconds.",
							Computed:            true,
						},
						"ipsec_policies_ike_lifetime": schema.Int64Attribute{
							MarkdownDescription: "The lifetime of the Phase 1 SA in seconds.",
							Computed:            true,
						},
						"ipsec_policies_child_auth_algo": schema.ListAttribute{
							MarkdownDescription: "This is the authentication algorithms to be used in Phase 2. The value should be an array with one of the following algorithms: `sha256`, `sha1`, `md5`",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ipsec_policies_child_cipher_algo": schema.ListAttribute{
							MarkdownDescription: "This is the cipher algorithms to be used in Phase 2. The value should be an array with one or more of the following algorithms: `aes256`, `aes192`, `aes128`, `tripledes`, `des`, `null`",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ipsec_policies_child_pfs_group": schema.ListAttribute{
							MarkdownDescription: "This is the Diffie-Hellman group to be used for Perfect Forward Secrecy in Phase 2. The value should be an array with one of the following values: `disabled`,`group14`, `group5`, `group2`, `group1`",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ipsec_policies_ike_auth_algo": schema.ListAttribute{
							MarkdownDescription: "This is the authentication algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: `sha256`, `sha1`, `md5`",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ipsec_policies_ike_cipher_algo": schema.ListAttribute{
							MarkdownDescription: "This is the cipher algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: `aes256`, `aes192`, `aes128`, `tripledes`, `des`",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ipsec_policies_ike_diffie_hellman_group": schema.ListAttribute{
							MarkdownDescription: "This is the Diffie-Hellman group to be used in Phase 1. The value should be an array with one of the following algorithms: `group14`, `group5`, `group2`, `group1`",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"ipsec_policies_ike_prf_algo": schema.ListAttribute{
							MarkdownDescription: "[optional] This is the pseudo-random function to be used in IKE_SA. The value should be an array with one of the following algorithms: `prfsha256`, `prfsha1`, `prfmd5`, `default`. The `default` option can be used to default to the Authentication algorithm.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"network_tags": schema.ListAttribute{
							MarkdownDescription: "A list of network tags that will connect with this peer. Use [`all`] for all networks. Use [`none`] for no networks. If not included, the default is [`all`].",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"private_subnets": schema.ListAttribute{
							MarkdownDescription: "The list of the private subnets of the VPN peer",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ApplianceThirdPartyVPNPeersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceThirdPartyVPNPeersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplianceThirdPartyVPNPeers

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res meraki.Res
	var err error

	if !res.Exists() {
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)
	config.Id = config.OrganizationId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
