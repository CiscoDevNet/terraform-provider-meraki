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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.ResourceWithIdentity    = &OrganizationWirelessDevicesProvisioningDeploymentsResource{}
	_ resource.ResourceWithImportState = &OrganizationWirelessDevicesProvisioningDeploymentsResource{}
	_ resource.ResourceWithModifyPlan  = &OrganizationWirelessDevicesProvisioningDeploymentsResource{}
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
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Organization Wireless Devices Provisioning Deployment` configuration in bulk.").AddBulkResourceIds("devices_new_serial").String,

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
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The list of items",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the item",
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the zero touch deployment request").AddStringEnumDescription("deploy", "replace").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("deploy", "replace"),
							},
						},
						"status": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Status of the zero touch deployment request").AddStringEnumDescription("completed", "failed", "in progress", "ready").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("completed", "failed", "in progress", "ready"),
							},
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the network the device is being added to").String,
							Optional:            true,
						},
						"devices_new_serial": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Serial number of the new device").String,
							Required:            true,
						},
						"devices_new_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the new device or serial number if not named").String,
							Optional:            true,
						},
						"devices_new_tags": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tag(s) of the new device").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"devices_new_rf_profile_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of RfProfile for new device").String,
							Optional:            true,
						},
						"devices_old_serial": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Serial number of the old device, only for `replace` deployments").String,
							Optional:            true,
						},
						"devices_old_after_action": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Action to be taken on the old device, only for `replace` deployments").String,
							Optional:            true,
						},
						"devices_old_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the old device").String,
							Optional:            true,
						},
						"devices_old_tags": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tag(s) of the old device").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"devices_old_rf_profile_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID of the RF profile of the old device").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"organization_id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("Organization ID").String,
				RequiredForImport: true,
			},
			"item_ids": identityschema.ListAttribute{
				Description:       "List of item IDs",
				ElementType:       types.StringType,
				OptionalForImport: true,
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

// The deployments API has no action batch support; it takes batches natively instead. POST and
// PUT on the collection path take and return `{"items": [...]}`, and only DELETE is per
// deployment. Items are matched by the serial of the new device, which the API echoes back.

// deploymentsBatchBody wraps item bodies as the `{"items": [...]}` payload the API expects.
func deploymentsBatchBody(items []string) string {
	body := "{}"
	for _, item := range items {
		body, _ = sjson.SetRaw(body, "items.-1", item)
	}
	return body
}

// deploymentIdsBySerial maps the new device serial of each returned deployment to its ID.
func deploymentIdsBySerial(res meraki.Res) map[string]string {
	ids := make(map[string]string)
	res.Get("items").ForEach(func(_, v gjson.Result) bool {
		ids[v.Get("devices.new.serial").String()] = v.Get("deploymentId").String()
		return true
	})
	return ids
}

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ResourceOrganizationWirelessDevicesProvisioningDeployments
	var identity ResourceOrganizationWirelessDevicesProvisioningDeploymentsIdentity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	items := make([]string, len(plan.Items))
	for i, item := range plan.Items {
		items[i] = item.toBody(ctx, ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems{})
	}
	res, err := r.client.Post(plan.getPath(), deploymentsBatchBody(items))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create deployments (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.OrganizationId
	ids := deploymentIdsBySerial(res)
	for i := range plan.Items {
		plan.Items[i].Id = types.StringValue(ids[plan.Items[i].DevicesNewSerial.ValueString()])
	}
	identity.toIdentity(ctx, &plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ResourceOrganizationWirelessDevicesProvisioningDeployments
	var identity ResourceOrganizationWirelessDevicesProvisioningDeploymentsIdentity

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read identity if available (requires Terraform >= 1.12.0)
	if req.Identity != nil && !req.Identity.Raw.IsNull() {
		diags = req.Identity.Get(ctx, &identity)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		state.fromIdentity(ctx, &identity)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath())
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 400")) {
		identity.toIdentity(ctx, &state)
		diags = resp.Identity.Set(ctx, &identity)
		resp.Diagnostics.Append(diags...)
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		if len(state.Items) > 0 {
			state.fromBodyImport(ctx, res)
		} else {
			state.fromBody(ctx, res)
		}
	} else {
		state.fromBodyPartial(ctx, res)
	}
	identity.toIdentity(ctx, &state)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ResourceOrganizationWirelessDevicesProvisioningDeployments

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

	// Delete deployments that are in state but no longer in plan
	for _, itemState := range state.Items {
		found := false
		for _, item := range plan.Items {
			if item.DevicesNewSerial.ValueString() == itemState.DevicesNewSerial.ValueString() {
				found = true
				break
			}
		}
		if !found {
			res, err := r.client.Delete(plan.getPath() + "/" + url.QueryEscape(itemState.Id.ValueString()))
			if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete deployment (DELETE), got error: %s, %s", err, res.String()))
				return
			}
		}
	}

	// Split the remaining plan items into changed and new ones
	var changed, created []string
	for i := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			if plan.Items[i].DevicesNewSerial.ValueString() != itemState.DevicesNewSerial.ValueString() {
				continue
			}
			found = true
			plan.Items[i].Id = itemState.Id
			if plan.Items[i].toBody(ctx, ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems{}) != itemState.toBody(ctx, ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems{}) {
				item, _ := sjson.Set(plan.Items[i].toBody(ctx, itemState), "deploymentId", itemState.Id.ValueString())
				changed = append(changed, item)
			}
			break
		}
		if !found {
			created = append(created, plan.Items[i].toBody(ctx, ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems{}))
		}
	}

	if len(changed) > 0 {
		res, err := r.client.Put(plan.getPath(), deploymentsBatchBody(changed))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update deployments (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}
	if len(created) > 0 {
		res, err := r.client.Post(plan.getPath(), deploymentsBatchBody(created))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to create deployments (POST), got error: %s, %s", err, res.String()))
			return
		}
		ids := deploymentIdsBySerial(res)
		for i := range plan.Items {
			if id, ok := ids[plan.Items[i].DevicesNewSerial.ValueString()]; ok {
				plan.Items[i].Id = types.StringValue(id)
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	var identity ResourceOrganizationWirelessDevicesProvisioningDeploymentsIdentity
	identity.toIdentity(ctx, &plan)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ResourceOrganizationWirelessDevicesProvisioningDeployments

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	for _, item := range state.Items {
		res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(item.Id.ValueString()))
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete deployment (DELETE), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != "" || req.Identity == nil || req.Identity.Raw.IsNull() {
		itemIdParts := make([]string, 0)
		if strings.Contains(req.ID, ",[") {
			itemIdParts = strings.Split(strings.Split(strings.Split(req.ID, ",[")[1], "]")[0], ",")
		}
		idParts := strings.Split(strings.Split(req.ID, ",[")[0], ",")

		if len(idParts) != 1 || idParts[0] == "" {
			expectedIdentifier := "Expected import identifier with format: <organization_id>"
			expectedIdentifier += " or <organization_id>,[<id>,...]"
			resp.Diagnostics.AddError(
				"Unexpected Import Identifier",
				fmt.Sprintf("%s. Got: %q", expectedIdentifier, req.ID),
			)
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)

		if len(itemIdParts) > 0 {
			items := make([]ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems, len(itemIdParts))
			for i, itemId := range itemIdParts {
				item := ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems{}
				item.Id = types.StringValue(itemId)
				item.DevicesNewTags = types.SetNull(types.StringType)
				item.DevicesOldTags = types.SetNull(types.StringType)
				items[i] = item
			}
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("items"), items)...)
		}
	} else {
		var identity ResourceOrganizationWirelessDevicesProvisioningDeploymentsIdentity
		diags := req.Identity.Get(ctx, &identity)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), identity.OrganizationId.ValueString())...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), identity.OrganizationId.ValueString())...)

		if len(identity.ItemIds.Elements()) > 0 {
			items := make([]ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems, len(identity.ItemIds.Elements()))
			var values []string
			identity.ItemIds.ElementsAs(ctx, &values, false)
			for i, itemId := range values {
				item := ResourceOrganizationWirelessDevicesProvisioningDeploymentsItems{}
				item.Id = types.StringValue(itemId)
				item.DevicesNewTags = types.SetNull(types.StringType)
				item.DevicesOldTags = types.SetNull(types.StringType)
				items[i] = item
			}
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("items"), items)...)
		}
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin modifyPlan
func (r *OrganizationWirelessDevicesProvisioningDeploymentsResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan, state ResourceOrganizationWirelessDevicesProvisioningDeployments

	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning ModifyPlan", plan.Id.ValueString()))
	// Remove incorrectly set IDs in plan (https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#prior-state-under-lists-and-sets)
	for i, item := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			if item.DevicesNewSerial.ValueString() != itemState.DevicesNewSerial.ValueString() {
				continue
			}
			found = true
		}
		if !found {
			plan.Items[i].Id = types.StringUnknown()
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: ModifyPlan finished successfully", plan.Id.ValueString()))

	diags = resp.Plan.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end modifyPlan
