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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &OrganizationSAMLRoleDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationSAMLRoleDataSource{}
)

func NewOrganizationSAMLRoleDataSource() datasource.DataSource {
	return &OrganizationSAMLRoleDataSource{}
}

type OrganizationSAMLRoleDataSource struct {
	client *meraki.Client
}

func (d *OrganizationSAMLRoleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_saml_role"
}

func (d *OrganizationSAMLRoleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Organization SAML Role` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
				Required:            true,
			},
			"org_access": schema.StringAttribute{
				MarkdownDescription: "The privilege of the SAML administrator on the organization. Can be one of `none`, `read-only`, `full` or `enterprise`",
				Computed:            true,
			},
			"role": schema.StringAttribute{
				MarkdownDescription: "The role of the SAML administrator",
				Computed:            true,
			},
			"networks": schema.ListNestedAttribute{
				MarkdownDescription: "The list of networks that the SAML administrator has privileges on",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access": schema.StringAttribute{
							MarkdownDescription: "The privilege of the SAML administrator on the network. Can be one of `full`, `read-only`, `guest-ambassador`, `monitor-only` or `ssid-admin`",
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
				MarkdownDescription: "The list of tags that the SAML administrator has privileges on",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access": schema.StringAttribute{
							MarkdownDescription: "The privilege of the SAML administrator on the tag. Can be one of `full`, `read-only`, `guest-ambassador` or `monitor-only`",
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
	}
}

func (d *OrganizationSAMLRoleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationSAMLRoleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OrganizationSAMLRole

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
