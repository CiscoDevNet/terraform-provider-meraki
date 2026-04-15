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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &OrganizationWirelessDevicesProvisioningDeploymentsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationWirelessDevicesProvisioningDeploymentsDataSource{}
)

func NewOrganizationWirelessDevicesProvisioningDeploymentsDataSource() datasource.DataSource {
	return &OrganizationWirelessDevicesProvisioningDeploymentsDataSource{}
}

type OrganizationWirelessDevicesProvisioningDeploymentsDataSource struct {
	client *meraki.Client
}

func (d *OrganizationWirelessDevicesProvisioningDeploymentsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_wireless_devices_provisioning_deployments"
}

func (d *OrganizationWirelessDevicesProvisioningDeploymentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Organization Wireless Devices Provisioning Deployments` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
				Required:            true,
			},
			"meta_counts_items_remaining": schema.Int64Attribute{
				MarkdownDescription: "The number of items in the dataset that are available on subsequent pages",
				Computed:            true,
			},
			"meta_counts_items_total": schema.Int64Attribute{
				MarkdownDescription: "The total number of items in the dataset",
				Computed:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "List of zero touch deployments to create",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"completed_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp of when the zero touch deployment request was completed",
							Computed:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp of when the zero touch deployment request was created",
							Computed:            true,
						},
						"deployment_id": schema.StringAttribute{
							MarkdownDescription: "zero touch deployment request identifier",
							Computed:            true,
						},
						"last_updated_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp of when the zero touch deployment request was last updated",
							Computed:            true,
						},
						"requested_at": schema.StringAttribute{
							MarkdownDescription: "Timestamp of when the zero touch deployment request was created",
							Computed:            true,
						},
						"status": schema.StringAttribute{
							MarkdownDescription: "Status of the zero touch deployment request. enum = [ready, in progress, completed, failed]",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Type of the zero touch deployment request. enum = [deploy, replace]",
							Computed:            true,
						},
						"devices_new_mac": schema.StringAttribute{
							MarkdownDescription: "MAC address of the new device",
							Computed:            true,
						},
						"devices_new_model": schema.StringAttribute{
							MarkdownDescription: "Model of the new device",
							Computed:            true,
						},
						"devices_new_name": schema.StringAttribute{
							MarkdownDescription: "Name of the new device or serial number if not named",
							Computed:            true,
						},
						"devices_new_serial": schema.StringAttribute{
							MarkdownDescription: "Serial number of the new device",
							Computed:            true,
						},
						"devices_new_rf_profile_id": schema.StringAttribute{
							MarkdownDescription: "ID of RfProfile for new device",
							Computed:            true,
						},
						"devices_new_rf_profile_name": schema.StringAttribute{
							MarkdownDescription: "Name of RfProfile for new device",
							Computed:            true,
						},
						"devices_new_tags": schema.ListAttribute{
							MarkdownDescription: "Tag(s) of the new device",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"devices_old_after_action": schema.StringAttribute{
							MarkdownDescription: "Action to be taken on the old device",
							Computed:            true,
						},
						"devices_old_mac": schema.StringAttribute{
							MarkdownDescription: "MAC address of the old device",
							Computed:            true,
						},
						"devices_old_model": schema.StringAttribute{
							MarkdownDescription: "Model of the old device",
							Computed:            true,
						},
						"devices_old_name": schema.StringAttribute{
							MarkdownDescription: "Name of the old device",
							Computed:            true,
						},
						"devices_old_serial": schema.StringAttribute{
							MarkdownDescription: "Serial number of the old device",
							Computed:            true,
						},
						"devices_old_rf_profile_id": schema.StringAttribute{
							MarkdownDescription: "ID of the RF profile",
							Computed:            true,
						},
						"devices_old_rf_profile_name": schema.StringAttribute{
							MarkdownDescription: "Name of the RF profile",
							Computed:            true,
						},
						"devices_old_tags": schema.ListAttribute{
							MarkdownDescription: "Tag(s) of the old device",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: "ID of the network the device is being added to",
							Computed:            true,
						},
						"network_name": schema.StringAttribute{
							MarkdownDescription: "Name of the network the device is being added to",
							Computed:            true,
						},
						"errors": schema.ListAttribute{
							MarkdownDescription: "Array of error message(s) if any",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationWirelessDevicesProvisioningDeploymentsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationWirelessDevicesProvisioningDeploymentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OrganizationWirelessDevicesProvisioningDeployments

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
	config.Id = config.OrganizationId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
