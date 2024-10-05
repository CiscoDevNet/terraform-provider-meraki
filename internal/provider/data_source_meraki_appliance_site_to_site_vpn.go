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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ApplianceSiteToSiteVPNDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceSiteToSiteVPNDataSource{}
)

func NewApplianceSiteToSiteVPNDataSource() datasource.DataSource {
	return &ApplianceSiteToSiteVPNDataSource{}
}

type ApplianceSiteToSiteVPNDataSource struct {
	client *meraki.Client
}

func (d *ApplianceSiteToSiteVPNDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_site_to_site_vpn"
}

func (d *ApplianceSiteToSiteVPNDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Appliance Site to Site VPN` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "The site-to-site VPN mode. Can be one of `none`, `spoke` or `hub`",
				Computed:            true,
			},
			"hubs": schema.ListNestedAttribute{
				MarkdownDescription: "The list of VPN hubs, in order of preference. In spoke mode, at least 1 hub is required.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hub_id": schema.StringAttribute{
							MarkdownDescription: "The network ID of the hub.",
							Computed:            true,
						},
						"use_default_route": schema.BoolAttribute{
							MarkdownDescription: "Only valid in `spoke` mode. Indicates whether default route traffic should be sent to this hub.",
							Computed:            true,
						},
					},
				},
			},
			"subnets": schema.ListNestedAttribute{
				MarkdownDescription: "The list of subnets and their VPN presence.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"local_subnet": schema.StringAttribute{
							MarkdownDescription: "The CIDR notation subnet used within the VPN",
							Computed:            true,
						},
						"use_vpn": schema.BoolAttribute{
							MarkdownDescription: "Indicates the presence of the subnet in the VPN",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ApplianceSiteToSiteVPNDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceSiteToSiteVPNDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplianceSiteToSiteVPN

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
	config.Id = config.NetworkId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
