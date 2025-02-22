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
	_ datasource.DataSource              = &CameraRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &CameraRolesDataSource{}
)

func NewCameraRolesDataSource() datasource.DataSource {
	return &CameraRolesDataSource{}
}

type CameraRolesDataSource struct {
	client *meraki.Client
}

func (d *CameraRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_camera_roles"
}

func (d *CameraRolesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Camera Role` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
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
							MarkdownDescription: "The name of the new role. Must be unique. This parameter is required.",
							Computed:            true,
						},
						"applied_on_devices": schema.ListNestedAttribute{
							MarkdownDescription: "Device tag on which this specified permission is applied.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Device id.",
										Computed:            true,
									},
									"in_networks_with_id": schema.StringAttribute{
										MarkdownDescription: "Network id scope",
										Computed:            true,
									},
									"in_networks_with_tag": schema.StringAttribute{
										MarkdownDescription: "Network tag scope",
										Computed:            true,
									},
									"permission_scope_id": schema.StringAttribute{
										MarkdownDescription: "Permission scope id",
										Computed:            true,
									},
									"tag": schema.StringAttribute{
										MarkdownDescription: "Device tag.",
										Computed:            true,
									},
								},
							},
						},
						"applied_on_networks": schema.ListNestedAttribute{
							MarkdownDescription: "Network tag on which this specified permission is applied.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Network id.",
										Computed:            true,
									},
									"permission_scope_id": schema.StringAttribute{
										MarkdownDescription: "Permission scope id",
										Computed:            true,
									},
									"tag": schema.StringAttribute{
										MarkdownDescription: "Network tag",
										Computed:            true,
									},
								},
							},
						},
						"applied_org_wide": schema.ListNestedAttribute{
							MarkdownDescription: "Permissions to be applied org wide.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"permission_scope_id": schema.StringAttribute{
										MarkdownDescription: "Permission scope id",
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

func (d *CameraRolesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *CameraRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CameraRoles

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "CameraRolesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "CameraRolesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
