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
	_ datasource.DataSource              = &NetworkFloorPlansDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkFloorPlansDataSource{}
)

func NewNetworkFloorPlansDataSource() datasource.DataSource {
	return &NetworkFloorPlansDataSource{}
}

type NetworkFloorPlansDataSource struct {
	client *meraki.Client
}

func (d *NetworkFloorPlansDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_floor_plans"
}

func (d *NetworkFloorPlansDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Network Floor Plan` configuration.",

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
						"image_contents": schema.StringAttribute{
							MarkdownDescription: "The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of your floor plan.",
							Computed:            true,
						},
						"bottom_left_corner_lat": schema.Float64Attribute{
							MarkdownDescription: "Latitude",
							Computed:            true,
						},
						"bottom_left_corner_lng": schema.Float64Attribute{
							MarkdownDescription: "Longitude",
							Computed:            true,
						},
						"bottom_right_corner_lat": schema.Float64Attribute{
							MarkdownDescription: "Latitude",
							Computed:            true,
						},
						"bottom_right_corner_lng": schema.Float64Attribute{
							MarkdownDescription: "Longitude",
							Computed:            true,
						},
						"center_lat": schema.Float64Attribute{
							MarkdownDescription: "Latitude",
							Computed:            true,
						},
						"center_lng": schema.Float64Attribute{
							MarkdownDescription: "Longitude",
							Computed:            true,
						},
						"top_left_corner_lat": schema.Float64Attribute{
							MarkdownDescription: "Latitude",
							Computed:            true,
						},
						"top_left_corner_lng": schema.Float64Attribute{
							MarkdownDescription: "Longitude",
							Computed:            true,
						},
						"top_right_corner_lat": schema.Float64Attribute{
							MarkdownDescription: "Latitude",
							Computed:            true,
						},
						"top_right_corner_lng": schema.Float64Attribute{
							MarkdownDescription: "Longitude",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworkFloorPlansDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *NetworkFloorPlansDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkFloorPlans

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "NetworkFloorPlansDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "NetworkFloorPlansDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read