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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
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
	_ datasource.DataSource              = &ApplianceUplinksSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceUplinksSettingsDataSource{}
)

func NewApplianceUplinksSettingsDataSource() datasource.DataSource {
	return &ApplianceUplinksSettingsDataSource{}
}

type ApplianceUplinksSettingsDataSource struct {
	client *meraki.Client
}

func (d *ApplianceUplinksSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_uplinks_settings"
}

func (d *ApplianceUplinksSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Appliance Uplinks Settings` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: "Device serial",
				Required:            true,
			},
			"interfaces_wan1_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable the interface.",
				Computed:            true,
			},
			"interfaces_wan1_pppoe_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether PPPoE is enabled.",
				Computed:            true,
			},
			"interfaces_wan1_pppoe_authentication_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether PPPoE authentication is enabled.",
				Computed:            true,
			},
			"interfaces_wan1_pppoe_authentication_password": schema.StringAttribute{
				MarkdownDescription: "Password for PPPoE authentication. This parameter is not returned.",
				Computed:            true,
				Sensitive:           true,
			},
			"interfaces_wan1_pppoe_authentication_username": schema.StringAttribute{
				MarkdownDescription: "Username for PPPoE authentication.",
				Computed:            true,
			},
			"interfaces_wan1_svis_ipv4_address": schema.StringAttribute{
				MarkdownDescription: "IP address and subnet mask when in static mode.",
				Computed:            true,
			},
			"interfaces_wan1_svis_ipv4_assignment_mode": schema.StringAttribute{
				MarkdownDescription: "The assignment mode for this SVI. Applies only when PPPoE is disabled.",
				Computed:            true,
			},
			"interfaces_wan1_svis_ipv4_gateway": schema.StringAttribute{
				MarkdownDescription: "Gateway IP address when in static mode.",
				Computed:            true,
			},
			"interfaces_wan1_svis_ipv4_nameservers_addresses": schema.ListAttribute{
				MarkdownDescription: "Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"interfaces_wan1_svis_ipv6_address": schema.StringAttribute{
				MarkdownDescription: "Static address that will override the one(s) received by SLAAC.",
				Computed:            true,
			},
			"interfaces_wan1_svis_ipv6_assignment_mode": schema.StringAttribute{
				MarkdownDescription: "The assignment mode for this SVI. Applies only when PPPoE is disabled.",
				Computed:            true,
			},
			"interfaces_wan1_svis_ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: "Static gateway that will override the one received by autoconf.",
				Computed:            true,
			},
			"interfaces_wan1_svis_ipv6_nameservers_addresses": schema.ListAttribute{
				MarkdownDescription: "Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"interfaces_wan1_vlan_tagging_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether VLAN tagging is enabled.",
				Computed:            true,
			},
			"interfaces_wan1_vlan_tagging_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "The ID of the VLAN to use for VLAN tagging.",
				Computed:            true,
			},
			"interfaces_wan2_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable the interface.",
				Computed:            true,
			},
			"interfaces_wan2_pppoe_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether PPPoE is enabled.",
				Computed:            true,
			},
			"interfaces_wan2_pppoe_authentication_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether PPPoE authentication is enabled.",
				Computed:            true,
			},
			"interfaces_wan2_pppoe_authentication_password": schema.StringAttribute{
				MarkdownDescription: "Password for PPPoE authentication. This parameter is not returned.",
				Computed:            true,
				Sensitive:           true,
			},
			"interfaces_wan2_pppoe_authentication_username": schema.StringAttribute{
				MarkdownDescription: "Username for PPPoE authentication.",
				Computed:            true,
			},
			"interfaces_wan2_svis_ipv4_address": schema.StringAttribute{
				MarkdownDescription: "IP address and subnet mask when in static mode.",
				Computed:            true,
			},
			"interfaces_wan2_svis_ipv4_assignment_mode": schema.StringAttribute{
				MarkdownDescription: "The assignment mode for this SVI. Applies only when PPPoE is disabled.",
				Computed:            true,
			},
			"interfaces_wan2_svis_ipv4_gateway": schema.StringAttribute{
				MarkdownDescription: "Gateway IP address when in static mode.",
				Computed:            true,
			},
			"interfaces_wan2_svis_ipv4_nameservers_addresses": schema.ListAttribute{
				MarkdownDescription: "Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"interfaces_wan2_svis_ipv6_address": schema.StringAttribute{
				MarkdownDescription: "Static address that will override the one(s) received by SLAAC.",
				Computed:            true,
			},
			"interfaces_wan2_svis_ipv6_assignment_mode": schema.StringAttribute{
				MarkdownDescription: "The assignment mode for this SVI. Applies only when PPPoE is disabled.",
				Computed:            true,
			},
			"interfaces_wan2_svis_ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: "Static gateway that will override the one received by autoconf.",
				Computed:            true,
			},
			"interfaces_wan2_svis_ipv6_nameservers_addresses": schema.ListAttribute{
				MarkdownDescription: "Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"interfaces_wan2_vlan_tagging_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether VLAN tagging is enabled.",
				Computed:            true,
			},
			"interfaces_wan2_vlan_tagging_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "The ID of the VLAN to use for VLAN tagging.",
				Computed:            true,
			},
		},
	}
}

func (d *ApplianceUplinksSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceUplinksSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplianceUplinksSettings

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
	config.Id = config.Serial

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
