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
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource = &OrganizationWirelessDevicesProvisioningDeploymentsResource{}
)

func NewOrganizationWirelessDevicesProvisioningDeploymentsResource() resource.Resource {
	return &OrganizationWirelessDevicesProvisioningDeploymentsResource{}
}

type OrganizationWirelessDevicesProvisioningDeploymentsResource struct {
	client *meraki.Client
}

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_wireless_devices_provisioning_deployments"
}

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Organization Wireless Devices Provisioning Deployments` configuration.").AddEarlyAccessDescription().String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Organization ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"meta_counts_items_remaining": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The number of items in the dataset that are available on subsequent pages").String,
				Optional:            true,
			},
			"meta_counts_items_total": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The total number of items in the dataset").String,
				Optional:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of zero touch deployments to create").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"completed_at": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Timestamp of when the zero touch deployment request was completed").String,
							Optional:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Timestamp of when the zero touch deployment request was created").String,
							Optional:            true,
						},
						"deployment_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("zero touch deployment request identifier").String,
							Optional:            true,
						},
						"last_updated_at": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Timestamp of when the zero touch deployment request was last updated").String,
							Optional:            true,
						},
						"requested_at": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Timestamp of when the zero touch deployment request was created").String,
							Optional:            true,
						},
						"status": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Status of the zero touch deployment request. enum = [ready, in progress, completed, failed]").String,
							Required:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the zero touch deployment request. enum = [deploy, replace]").String,
							Required:            true,
						},
						"devices_new_mac": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MAC address of the new device").String,
							Optional:            true,
						},
						"devices_new_model": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Model of the new device").String,
							Optional:            true,
						},
						"devices_new_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the new device or serial number if not named").String,
							Optional:            true,
						},
						"devices_new_serial": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Serial number of the new device").String,
							Optional:            true,
						},
						"devices_new_rf_profile_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of RfProfile for new device").String,
							Optional:            true,
						},
						"devices_new_rf_profile_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of RfProfile for new device").String,
							Optional:            true,
						},
						"devices_new_tags": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tag(s) of the new device").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"devices_old_after_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Action to be taken on the old device").String,
							Optional:            true,
						},
						"devices_old_mac": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MAC address of the old device").String,
							Optional:            true,
						},
						"devices_old_model": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Model of the old device").String,
							Optional:            true,
						},
						"devices_old_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the old device").String,
							Optional:            true,
						},
						"devices_old_serial": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Serial number of the old device").String,
							Optional:            true,
						},
						"devices_old_rf_profile_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the RF profile").String,
							Optional:            true,
						},
						"devices_old_rf_profile_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the RF profile").String,
							Optional:            true,
						},
						"devices_old_tags": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tag(s) of the old device").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the network the device is being added to").String,
							Optional:            true,
						},
						"network_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the network the device is being added to").String,
							Optional:            true,
						},
						"errors": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Array of error message(s) if any").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OrganizationWirelessDevicesProvisioningDeployments

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, OrganizationWirelessDevicesProvisioningDeployments{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.OrganizationId
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OrganizationWirelessDevicesProvisioningDeployments

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state OrganizationWirelessDevicesProvisioningDeployments

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationWirelessDevicesProvisioningDeployments

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import
