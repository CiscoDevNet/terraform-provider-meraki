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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SwitchOrganizationPortsProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchOrganizationPortsProfileDataSource{}
)

func NewSwitchOrganizationPortsProfileDataSource() datasource.DataSource {
	return &SwitchOrganizationPortsProfileDataSource{}
}

type SwitchOrganizationPortsProfileDataSource struct {
	client *meraki.Client
}

func (d *SwitchOrganizationPortsProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_organization_ports_profile"
}

func (d *SwitchOrganizationPortsProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Switch Organization Ports Profile` configuration.").AddEarlyAccessDescription().String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Text describing the profile.",
				Computed:            true,
			},
			"is_organization_wide": schema.BoolAttribute{
				MarkdownDescription: "The scope of the profile whether it is organization level or network level",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the profile.",
				Optional:            true,
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "The network identifier",
				Computed:            true,
			},
			"networks_type": schema.StringAttribute{
				MarkdownDescription: "Flag to identify if the networks networks are excluded or included",
				Computed:            true,
			},
			"networks_values": schema.ListNestedAttribute{
				MarkdownDescription: "The network object containing name and id",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The network identifier",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The network name",
							Computed:            true,
						},
					},
				},
			},
			"port_access_policy_number": schema.Int64Attribute{
				MarkdownDescription: "The number of a custom access policy to configure on the port profile. Only applicable when `accessPolicyType` is `Custom access policy`.",
				Computed:            true,
			},
			"port_access_policy_type": schema.StringAttribute{
				MarkdownDescription: "The type of the access policy of the port profile. Only applicable to access ports.",
				Computed:            true,
			},
			"port_adaptive_policy_group_id": schema.StringAttribute{
				MarkdownDescription: "The adaptive policy group ID that will be used to tag traffic through this port profile. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API.",
				Computed:            true,
			},
			"port_adaptive_policy_voice_group_id": schema.StringAttribute{
				MarkdownDescription: "The adaptive policy group ID that will be used to tag voice traffic through this port profile. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API.",
				Computed:            true,
			},
			"port_allowed_vlans": schema.StringAttribute{
				MarkdownDescription: "The VLANs allowed on the port profile. Only applicable to trunk ports.",
				Computed:            true,
			},
			"port_dai_trusted": schema.BoolAttribute{
				MarkdownDescription: "If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.",
				Computed:            true,
			},
			"port_isolation_enabled": schema.BoolAttribute{
				MarkdownDescription: "The isolation status of the port profile.",
				Computed:            true,
			},
			"port_peer_sgt_capable": schema.BoolAttribute{
				MarkdownDescription: "If true, Peer SGT is enabled for traffic through this port profile. Applicable to trunk ports only.",
				Computed:            true,
			},
			"port_poe_enabled": schema.BoolAttribute{
				MarkdownDescription: "The PoE status of the port profile.",
				Computed:            true,
			},
			"port_rstp_enabled": schema.BoolAttribute{
				MarkdownDescription: "The rapid spanning tree protocol status.",
				Computed:            true,
			},
			"port_sticky_mac_allow_list_limit": schema.Int64Attribute{
				MarkdownDescription: "The maximum number of MAC addresses for sticky MAC allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.",
				Computed:            true,
			},
			"port_storm_control_enabled": schema.BoolAttribute{
				MarkdownDescription: "The storm control status of the port profile.",
				Computed:            true,
			},
			"port_stp_guard": schema.StringAttribute{
				MarkdownDescription: "The state of the STP guard.",
				Computed:            true,
			},
			"port_type": schema.StringAttribute{
				MarkdownDescription: "The type of the port profile.",
				Computed:            true,
			},
			"port_udld": schema.StringAttribute{
				MarkdownDescription: "The action to take when Unidirectional Link is detected. LinkDefault configuration is Alert only.",
				Computed:            true,
			},
			"port_vlan": schema.Int64Attribute{
				MarkdownDescription: "The VLAN of the port profile. A null value will clear the value set for trunk ports.",
				Computed:            true,
			},
			"port_voice_vlan": schema.Int64Attribute{
				MarkdownDescription: "The voice VLAN of the port profile. Only applicable to access ports.",
				Computed:            true,
			},
			"port_mac_allow_list": schema.ListAttribute{
				MarkdownDescription: "Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when `accessPolicyType` is `MAC allow list`.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"port_sticky_mac_allow_list": schema.ListAttribute{
				MarkdownDescription: "The initial list of MAC addresses for sticky Mac allow list. Only applicable when `accessPolicyType` is `Sticky MAC allow list`.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: "Space-seperated list of tags",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}
func (d *SwitchOrganizationPortsProfileDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *SwitchOrganizationPortsProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchOrganizationPortsProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SwitchOrganizationPortsProfile

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res meraki.Res
	var err error
	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		found := false
		if len(res.Array()) > 0 {
			res.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("name").String() {
					if found {
						resp.Diagnostics.AddWarning("Multiple objects with same name", fmt.Sprintf("Found multiple objects with name: %s", config.Name.ValueString()))
						return false
					}
					config.Id = types.StringValue(v.Get("profileId").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					res = meraki.Res{Result: v}
					found = true
				}
				return true
			})
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}

	if !res.Exists() {
		res, err = d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
