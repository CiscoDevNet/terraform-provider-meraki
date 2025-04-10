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
	_ datasource.DataSource              = &NetworkDevicesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkDevicesDataSource{}
)

func NewNetworkDevicesDataSource() datasource.DataSource {
	return &NetworkDevicesDataSource{}
}

type NetworkDevicesDataSource struct {
	client *meraki.Client
}

func (d *NetworkDevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_devices"
}

func (d *NetworkDevicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Network Devices` configuration.").String,

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
						"address": schema.StringAttribute{
							MarkdownDescription: "Physical address of the device",
							Computed:            true,
						},
						"firmware": schema.StringAttribute{
							MarkdownDescription: "Firmware version of the device",
							Computed:            true,
						},
						"floor_plan_id": schema.StringAttribute{
							MarkdownDescription: "The floor plan to associate to this device. null disassociates the device from the floorplan.",
							Computed:            true,
						},
						"lan_ip": schema.StringAttribute{
							MarkdownDescription: "LAN IP address of the device",
							Computed:            true,
						},
						"lat": schema.Float64Attribute{
							MarkdownDescription: "Latitude of the device",
							Computed:            true,
						},
						"lng": schema.Float64Attribute{
							MarkdownDescription: "Longitude of the device",
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: "MAC address of the device",
							Computed:            true,
						},
						"model": schema.StringAttribute{
							MarkdownDescription: "Model of the device",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the device",
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: "ID of the network the device belongs to",
							Computed:            true,
						},
						"notes": schema.StringAttribute{
							MarkdownDescription: "Notes for the device, limited to 255 characters",
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: "Serial number of the device",
							Computed:            true,
						},
						"beacon_id_params_major": schema.Int64Attribute{
							MarkdownDescription: "The major number to be used in the beacon identifier",
							Computed:            true,
						},
						"beacon_id_params_minor": schema.Int64Attribute{
							MarkdownDescription: "The minor number to be used in the beacon identifier",
							Computed:            true,
						},
						"beacon_id_params_uuid": schema.StringAttribute{
							MarkdownDescription: "The UUID to be used in the beacon identifier",
							Computed:            true,
						},
						"details": schema.ListNestedAttribute{
							MarkdownDescription: "Additional device information",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Additional property name",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "Additional property value",
										Computed:            true,
									},
								},
							},
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: "List of tags assigned to the device",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworkDevicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *NetworkDevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkDevices

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "NetworkDevicesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "NetworkDevicesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
