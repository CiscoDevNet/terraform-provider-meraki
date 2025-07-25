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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ApplianceSSIDsDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceSSIDsDataSource{}
)

func NewApplianceSSIDsDataSource() datasource.DataSource {
	return &ApplianceSSIDsDataSource{}
}

type ApplianceSSIDsDataSource struct {
	client *meraki.Client
}

func (d *ApplianceSSIDsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_ssids"
}

func (d *ApplianceSSIDsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Appliance SSIDs` configuration in bulk.").String,

		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The list of items",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the object",
							Computed:            true,
						},
						"auth_mode": schema.StringAttribute{
							MarkdownDescription: "The association control method for the SSID.",
							Computed:            true,
						},
						"default_vlan_id": schema.Int64Attribute{
							MarkdownDescription: "The VLAN ID of the VLAN associated to this SSID.",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether or not the SSID is enabled.",
							Computed:            true,
						},
						"encryption_mode": schema.StringAttribute{
							MarkdownDescription: "The psk encryption mode for the SSID.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the SSID.",
							Computed:            true,
						},
						"number": schema.Int64Attribute{
							MarkdownDescription: "The number of the SSID.",
							Computed:            true,
						},
						"visible": schema.BoolAttribute{
							MarkdownDescription: "Boolean indicating whether the MX should advertise or hide this SSID.",
							Computed:            true,
						},
						"wpa_encryption_mode": schema.StringAttribute{
							MarkdownDescription: "WPA encryption mode for the SSID.",
							Computed:            true,
						},
						"radius_servers": schema.ListNestedAttribute{
							MarkdownDescription: "The RADIUS 802.1x servers to be used for authentication.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										MarkdownDescription: "The IP address of your RADIUS server.",
										Computed:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: "The UDP port your RADIUS servers listens on for Access-requests.",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *ApplianceSSIDsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceSSIDsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceApplianceSSIDs

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "ApplianceSSIDsDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "ApplianceSSIDsDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
