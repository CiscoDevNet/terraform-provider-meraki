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
	_ datasource.DataSource              = &OrganizationAdminsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationAdminsDataSource{}
)

func NewOrganizationAdminsDataSource() datasource.DataSource {
	return &OrganizationAdminsDataSource{}
}

type OrganizationAdminsDataSource struct {
	client *meraki.Client
}

func (d *OrganizationAdminsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_admins"
}

func (d *OrganizationAdminsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Organization Admin` configuration in bulk.").String,

		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
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
						"authentication_method": schema.StringAttribute{
							MarkdownDescription: "No longer used as of Cisco SecureX end-of-life. Can be one of `Email`. The default is Email authentication.",
							Computed:            true,
						},
						"email": schema.StringAttribute{
							MarkdownDescription: "The email of the dashboard administrator. This attribute can not be updated.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the dashboard administrator",
							Computed:            true,
						},
						"org_access": schema.StringAttribute{
							MarkdownDescription: "The privilege of the dashboard administrator on the organization. Can be one of `full`, `read-only`, `enterprise` or `none`",
							Computed:            true,
						},
						"networks": schema.ListNestedAttribute{
							MarkdownDescription: "The list of networks that the dashboard administrator has privileges on",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"access": schema.StringAttribute{
										MarkdownDescription: "The privilege of the dashboard administrator on the network. Can be one of `full`, `read-only`, `guest-ambassador` or `monitor-only`",
										Computed:            true,
									},
									"id": schema.StringAttribute{
										MarkdownDescription: "The network ID",
										Computed:            true,
									},
								},
							},
						},
						"tags": schema.ListNestedAttribute{
							MarkdownDescription: "The list of tags that the dashboard administrator has privileges on",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"access": schema.StringAttribute{
										MarkdownDescription: "The privilege of the dashboard administrator on the tag. Can be one of `full`, `read-only`, `guest-ambassador` or `monitor-only`",
										Computed:            true,
									},
									"tag": schema.StringAttribute{
										MarkdownDescription: "The name of the tag",
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

func (d *OrganizationAdminsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationAdminsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceOrganizationAdmins

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "OrganizationAdminsDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "OrganizationAdminsDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
