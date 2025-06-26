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
	_ datasource.DataSource              = &OrganizationLicensesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationLicensesDataSource{}
)

func NewOrganizationLicensesDataSource() datasource.DataSource {
	return &OrganizationLicensesDataSource{}
}

type OrganizationLicensesDataSource struct {
	client *meraki.Client
}

func (d *OrganizationLicensesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_licenses"
}

func (d *OrganizationLicensesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Organization Licenses` configuration.").String,

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
						"activation_date": schema.StringAttribute{
							MarkdownDescription: "The date the license started burning",
							Computed:            true,
						},
						"claim_date": schema.StringAttribute{
							MarkdownDescription: "The date the license was claimed into the organization",
							Computed:            true,
						},
						"device_serial": schema.StringAttribute{
							MarkdownDescription: "Serial number of the device the license is assigned to",
							Computed:            true,
						},
						"duration_in_days": schema.Int64Attribute{
							MarkdownDescription: "The duration of the individual license",
							Computed:            true,
						},
						"expiration_date": schema.StringAttribute{
							MarkdownDescription: "The date the license will expire",
							Computed:            true,
						},
						"head_license_id": schema.StringAttribute{
							MarkdownDescription: "The id of the head license this license is queued behind. If there is no head license, it returns nil.",
							Computed:            true,
						},
						"license_key": schema.StringAttribute{
							MarkdownDescription: "License key",
							Computed:            true,
						},
						"license_type": schema.StringAttribute{
							MarkdownDescription: "License type",
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: "ID of the network the license is assigned to",
							Computed:            true,
						},
						"order_number": schema.StringAttribute{
							MarkdownDescription: "Order number",
							Computed:            true,
						},
						"seat_count": schema.Int64Attribute{
							MarkdownDescription: "The number of seats of the license. Only applicable to SM licenses.",
							Computed:            true,
						},
						"state": schema.StringAttribute{
							MarkdownDescription: "The state of the license. All queued licenses have a status of `recentlyQueued`.",
							Computed:            true,
						},
						"total_duration_in_days": schema.Int64Attribute{
							MarkdownDescription: "The duration of the license plus all permanently queued licenses associated with it",
							Computed:            true,
						},
						"permanently_queued_licenses": schema.ListNestedAttribute{
							MarkdownDescription: "DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial= to retrieved queued licenses for a given device.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"duration_in_days": schema.Int64Attribute{
										MarkdownDescription: "The duration of the individual license",
										Computed:            true,
									},
									"id": schema.StringAttribute{
										MarkdownDescription: "Permanently queued license ID",
										Computed:            true,
									},
									"license_key": schema.StringAttribute{
										MarkdownDescription: "License key",
										Computed:            true,
									},
									"license_type": schema.StringAttribute{
										MarkdownDescription: "License type",
										Computed:            true,
									},
									"order_number": schema.StringAttribute{
										MarkdownDescription: "Order number",
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

func (d *OrganizationLicensesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationLicensesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceOrganizationLicenses

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "OrganizationLicensesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "OrganizationLicensesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
