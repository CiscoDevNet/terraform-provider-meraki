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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &ApplianceThirdPartyVPNPeersResource{}
	_ resource.ResourceWithImportState = &ApplianceThirdPartyVPNPeersResource{}
)

func NewApplianceThirdPartyVPNPeersResource() resource.Resource {
	return &ApplianceThirdPartyVPNPeersResource{}
}

type ApplianceThirdPartyVPNPeersResource struct {
	client *meraki.Client
}

func (r *ApplianceThirdPartyVPNPeersResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_third_party_vpn_peers"
}

func (r *ApplianceThirdPartyVPNPeersResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance Third Party VPN Peers` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Organization ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"peers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The list of VPN peers").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ike_version": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] The IKE version to be used for the IPsec VPN peer configuration. Defaults to `1` when omitted.").AddStringEnumDescription("1", "2").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("1", "2"),
							},
						},
						"ipsec_policies_preset": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("One of the following available presets: `default`, `aws`, `azure`, `umbrella`, `zscaler`. If this is provided, the `ipsecPolicies` parameter is ignored.").String,
							Optional:            true,
						},
						"is_route_based": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] If true, the VPN peer is route-based. If not included, the default is false. Supported only for MX 19.1 and above.").String,
							Optional:            true,
						},
						"local_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] The local ID is used to identify the MX to the peer. This will apply to all MXs this peer applies to.").String,
							Optional:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the VPN peer").String,
							Required:            true,
						},
						"peer_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The ID of the IPsec peer").String,
							Optional:            true,
						},
						"priority_in_group": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] Represents the order of peer inside a group. If you submit a request with the numbers [1, 9, 999], these numbers will be automatically adjusted to a sequential order starting from 1. So, they will be changed to [1, 2, 3] to reflect their positions in the sequence.").String,
							Optional:            true,
						},
						"public_hostname": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] The public hostname of the VPN peer").String,
							Optional:            true,
						},
						"public_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] The public IP of the VPN peer").String,
							Optional:            true,
						},
						"remote_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] The remote ID is used to identify the connecting VPN peer. This can either be a valid IPv4 Address, FQDN or User FQDN.").String,
							Optional:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The shared secret with the VPN peer").String,
							Required:            true,
						},
						"ebgp_neighbor_ebgp_hold_timer": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The eBGP hold timer in seconds for each neighbor. The eBGP hold timer must be an integer between 12 and 240.").String,
							Optional:            true,
						},
						"ebgp_neighbor_ebgp_multihop": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure this if the neighbor is not adjacent. The eBGP multi-hop must be an integer between 1 and 255.").String,
							Optional:            true,
						},
						"ebgp_neighbor_ip_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The IP version of the neighbor").String,
							Optional:            true,
						},
						"ebgp_neighbor_multi_exit_discriminator": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configures the local metric associated with routes received from the remote peer. Routes from peers with lower metrics are will be preferred. Must be an integer between 0 and 4294967295. MED is 6th in the decision tree when identical routes from multiple peers exist.").String,
							Optional:            true,
						},
						"ebgp_neighbor_neighbor_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IPv4/IPv6 address of the neighbor").String,
							Optional:            true,
						},
						"ebgp_neighbor_remote_as_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Remote ASN of the neighbor. The remote ASN must be an integer between 1 and 4294967295.").String,
							Optional:            true,
						},
						"ebgp_neighbor_source_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Source IP of eBGP neighbor").String,
							Optional:            true,
						},
						"ebgp_neighbor_weight": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configures the local metric associated with routes received from the remote peer. Routes from peers with lower metrics are will be preferred. Must be an integer between 0 and 4294967295. MED is 6th in the decision tree when identical routes from multiple peers exist.").String,
							Optional:            true,
						},
						"ebgp_neighbor_path_prepend": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Prepends the AS_PATH BGP Attribute associated with routes received from the remote peer. Configurable value of ASNs to prepend. Length of the array may not exceed 10, and each ASN in the array must be an integer between 1 and 4294967295. AS_PATH is 4th in the decision tree when identical routes from multiple peers exist.").String,
							ElementType:         types.Int64Type,
							Optional:            true,
						},
						"group_active_active_tunnel": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] Both primary and backup tunnels are active.").String,
							Optional:            true,
						},
						"group_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] Represents the ordering of primary and backup tunnels group. primary and backup tunnels are grouped by this number. If you submit a request with the numbers [1, 9, 999], these numbers will be automatically adjusted to a sequential order starting from 1. So, they will be changed to [1, 2, 3] to reflect their positions in the sequence.").String,
							Optional:            true,
						},
						"group_failover_direct_to_internet": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] When both primary and backup tunnels are down, direct traffic to the internet. Traffic will be routed via the WAN").String,
							Optional:            true,
						},
						"ipsec_policies_child_lifetime": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The lifetime of the Phase 2 SA in seconds.").String,
							Optional:            true,
						},
						"ipsec_policies_ike_lifetime": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The lifetime of the Phase 1 SA in seconds.").String,
							Optional:            true,
						},
						"ipsec_policies_child_auth_algo": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("This is the authentication algorithms to be used in Phase 2. The value should be an array with one of the following algorithms: `sha256`, `sha1`, `md5`").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"ipsec_policies_child_cipher_algo": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("This is the cipher algorithms to be used in Phase 2. The value should be an array with one or more of the following algorithms: `aes256`, `aes192`, `aes128`, `tripledes`, `des`, `null`").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"ipsec_policies_child_pfs_group": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("This is the Diffie-Hellman group to be used for Perfect Forward Secrecy in Phase 2. The value should be an array with one of the following values: `disabled`,`group14`, `group5`, `group2`, `group1`").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"ipsec_policies_ike_auth_algo": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("This is the authentication algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: `sha256`, `sha1`, `md5`").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"ipsec_policies_ike_cipher_algo": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("This is the cipher algorithm to be used in Phase 1. The value should be an array with one of the following algorithms: `aes256`, `aes192`, `aes128`, `tripledes`, `des`").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"ipsec_policies_ike_diffie_hellman_group": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("This is the Diffie-Hellman group to be used in Phase 1. The value should be an array with one of the following algorithms: `group14`, `group5`, `group2`, `group1`").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"ipsec_policies_ike_prf_algo": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] This is the pseudo-random function to be used in IKE_SA. The value should be an array with one of the following algorithms: `prfsha256`, `prfsha1`, `prfmd5`, `default`. The `default` option can be used to default to the Authentication algorithm.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"network_ids": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("[optional] A list of network IDs.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"sla_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The ID of the SLA policy").String,
							Optional:            true,
						},
						"network_tags": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A list of network tags that will connect with this peer. Use [`all`] for all networks. Use [`none`] for no networks. If not included, the default is [`all`].").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"private_subnets": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The list of the private subnets of the VPN peer").String,
							ElementType:         types.StringType,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ApplianceThirdPartyVPNPeersResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ApplianceThirdPartyVPNPeersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplianceThirdPartyVPNPeers

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ApplianceThirdPartyVPNPeers{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.OrganizationId
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *ApplianceThirdPartyVPNPeersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplianceThirdPartyVPNPeers

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath())
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *ApplianceThirdPartyVPNPeersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplianceThirdPartyVPNPeers

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx, state)
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *ApplianceThirdPartyVPNPeersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplianceThirdPartyVPNPeers

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	body := state.toDestroyBody(ctx)
	res, err := r.client.Put(state.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ApplianceThirdPartyVPNPeersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <organization_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
