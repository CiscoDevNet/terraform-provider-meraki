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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	_ resource.Resource                = &DevicesResource{}
	_ resource.ResourceWithImportState = &DevicesResource{}
	_ resource.ResourceWithModifyPlan  = &DevicesResource{}
)

func NewDevicesResource() resource.Resource {
	return &DevicesResource{}
}

type DevicesResource struct {
	client *meraki.Client
}

func (r *DevicesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (r *DevicesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Device` configuration in bulk.").AddBulkResourceIds("serial").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "The organization ID",
				Required:            true,
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The list of items",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"serial": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Switch serial").String,
							Required:            true,
						},
						"address": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The address of a device").String,
							Optional:            true,
						},
						"floor_plan_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The floor plan to associate to this device. null disassociates the device from the floorplan.").String,
							Optional:            true,
						},
						"lat": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The latitude of a device").String,
							Optional:            true,
						},
						"lng": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The longitude of a device").String,
							Optional:            true,
						},
						"move_map_marker": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.").String,
							Optional:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of a device").String,
							Optional:            true,
						},
						"notes": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The notes for the device. String. Limited to 255 characters.").String,
							Optional:            true,
						},
						"switch_profile_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The ID of a switch template to bind to the device (for available switch templates, see the `Switch Templates` endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch template, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.").String,
							Optional:            true,
						},
						"tags": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The list of tags of a device").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *DevicesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DevicesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ResourceDevices

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	actions := make([]meraki.ActionModel, len(plan.Items))

	for i, item := range plan.Items {
		actions[i] = meraki.ActionModel{
			Operation: "update",
			Resource:  plan.getItemPath(item.Serial.ValueString()),
			Body:      item.toBody(ctx, ResourceDevicesItems{}),
		}
	}
	var res meraki.Res
	var err error
	if len(actions) > 0 {
		res, err = r.client.Batch(plan.OrganizationId.ValueString(), actions)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
			return
		}
	}
	plan.Id = plan.OrganizationId

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DevicesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ResourceDevices

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath())
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 400")) {
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *DevicesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ResourceDevices

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
	var actions []meraki.ActionModel
	// Check for new and updated items
	for i := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			if plan.Items[i].Serial.ValueString() != itemState.Serial.ValueString() {
				continue
			}

			found = true
			// If the item is present in both plan and state, we need to check if it has changes
			hasChanges := plan.hasChanges(ctx, &state, plan.Items[i].Serial.ValueString())
			if hasChanges {
				actions = append(actions, meraki.ActionModel{
					Operation: "update",
					Resource:  plan.getItemPath(plan.Items[i].Serial.ValueString()),
					Body:      plan.Items[i].toBody(ctx, itemState),
				})
			}
			break
		}
		if !found {
			// If the item is present in plan, but not in state, we need to create it
			actions = append(actions, meraki.ActionModel{
				Operation: "update",
				Resource:  plan.getItemPath(plan.Items[i].Serial.ValueString()),
				Body:      plan.Items[i].toBody(ctx, ResourceDevicesItems{}),
			})
		}
	}

	var res meraki.Res
	var err error
	if len(actions) > 0 {
		res, err = r.client.Batch(plan.OrganizationId.ValueString(), actions)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *DevicesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ResourceDevices

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
func (r *DevicesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	itemIdParts := make([]string, 0)
	if strings.Contains(req.ID, ",[") {
		itemIdParts = strings.Split(strings.Split(strings.Split(req.ID, ",[")[1], "]")[0], ",")
	}
	idParts := strings.Split(strings.Split(req.ID, ",[")[0], ",")

	if len(idParts) != 1 || idParts[0] == "" {
		expectedIdentifier := "Expected import identifier with format: <organization_id>"
		expectedIdentifier += " or <organization_id>,[<serial>,...]"
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("%s. Got: %q", expectedIdentifier, req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)

	if len(itemIdParts) > 0 {
		items := make([]ResourceDevicesItems, len(itemIdParts))
		for i, itemId := range itemIdParts {
			item := ResourceDevicesItems{}
			item.Serial = types.StringValue(itemId)
			item.Tags = types.SetNull(types.StringType)
			items[i] = item
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("items"), items)...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin modifyPlan
func (r *DevicesResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan, state ResourceDevices

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

	tflog.Debug(ctx, fmt.Sprintf("%s: ModifyPlan finished successfully", plan.Id.ValueString()))

	diags = resp.Plan.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end modifyPlan
