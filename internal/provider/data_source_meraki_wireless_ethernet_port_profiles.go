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
	_ datasource.DataSource              = &WirelessEthernetPortProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessEthernetPortProfilesDataSource{}
)

func NewWirelessEthernetPortProfilesDataSource() datasource.DataSource {
	return &WirelessEthernetPortProfilesDataSource{}
}

type WirelessEthernetPortProfilesDataSource struct {
	client *meraki.Client
}

func (d *WirelessEthernetPortProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ethernet_port_profiles"
}

func (d *WirelessEthernetPortProfilesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Wireless Ethernet Port Profile` configuration.",

		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The list of items",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the object",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "AP port profile name",
							Computed:            true,
						},
						"ports": schema.ListNestedAttribute{
							MarkdownDescription: "AP ports configuration",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										MarkdownDescription: "AP port enabled",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "AP port name",
										Computed:            true,
									},
									"psk_group_id": schema.StringAttribute{
										MarkdownDescription: "AP port PSK Group ID",
										Computed:            true,
									},
									"ssid": schema.Int64Attribute{
										MarkdownDescription: "AP port ssid number",
										Computed:            true,
									},
								},
							},
						},
						"usb_ports": schema.ListNestedAttribute{
							MarkdownDescription: "AP usb ports configuration",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										MarkdownDescription: "AP usb port enabled",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "AP usb port name",
										Computed:            true,
									},
									"ssid": schema.Int64Attribute{
										MarkdownDescription: "AP usb port ssid number",
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

func (d *WirelessEthernetPortProfilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessEthernetPortProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessEthernetPortProfiles

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "WirelessEthernetPortProfilesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "WirelessEthernetPortProfilesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read