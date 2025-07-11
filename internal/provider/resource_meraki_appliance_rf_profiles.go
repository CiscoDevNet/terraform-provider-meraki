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
	"slices"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &ApplianceRFProfilesResource{}
	_ resource.ResourceWithImportState = &ApplianceRFProfilesResource{}
	_ resource.ResourceWithModifyPlan  = &ApplianceRFProfilesResource{}
)

func NewApplianceRFProfilesResource() resource.Resource {
	return &ApplianceRFProfilesResource{}
}

type ApplianceRFProfilesResource struct {
	client *meraki.Client
}

func (r *ApplianceRFProfilesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_rf_profiles"
}

func (r *ApplianceRFProfilesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance RF Profile` configuration in bulk.").AddBulkResourceIds("name").String,

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
			"network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network ID").String,
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
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the new profile. Must be unique. This param is required on creation.").String,
							Required:            true,
						},
						"five_ghz_settings_ax_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Determines whether ax radio on 5Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.").String,
							Optional:            true,
						},
						"five_ghz_settings_min_bitrate": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of 5Ghz band. Can be one of `6`, `9`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 12.").String,
							Optional:            true,
						},
						"per_ssid_settings_1_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi"),
							},
						},
						"per_ssid_settings_1_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_2_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi"),
							},
						},
						"per_ssid_settings_2_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_3_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi"),
							},
						},
						"per_ssid_settings_3_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_4_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi"),
							},
						},
						"per_ssid_settings_4_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"two_four_ghz_settings_ax_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.").String,
							Optional:            true,
						},
						"two_four_ghz_settings_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 11.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ApplianceRFProfilesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ApplianceRFProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ResourceApplianceRFProfiles

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	actions := make([]meraki.ActionModel, len(plan.Items))

	for i, item := range plan.Items {
		actions[i] = meraki.ActionModel{
			Operation: "create",
			Resource:  plan.getPath(),
			Body:      item.toBody(ctx, ResourceApplianceRFProfilesItems{}),
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
	for i := range plan.Items {
		plan.Items[i].Id = types.StringValue(res.Get("status.createdResources." + strconv.Itoa(i) + ".id").String())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *ApplianceRFProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ResourceApplianceRFProfiles

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

func (r *ApplianceRFProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ResourceApplianceRFProfiles

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
	// If there are destroy values, we need to compare the plan and state to determine what to delete
	for _, itemState := range state.Items {
		found := false
		for _, item := range plan.Items {
			if item.Name.ValueString() != itemState.Name.ValueString() {
				continue
			}

			// If the item is present in both plan and state, we can skip it
			found = true
			break
		}
		if !found {
			// If the item is present in state, but not in plan, we need to delete it
			actions = append(actions, meraki.ActionModel{
				Operation: "destroy",
				Resource:  plan.getPath() + "/" + itemState.Id.ValueString(),
				Body:      "{}",
			})
		}
	}
	newIdIndexes := make([]int, 0)
	// Check for new and updated items
	for i := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			if plan.Items[i].Name.ValueString() != itemState.Name.ValueString() {
				continue
			}

			found = true
			// If the item is present in both plan and state, we need to check if it has changes
			hasChanges := plan.hasChanges(ctx, &state, plan.Items[i].Id.ValueString())
			if hasChanges {
				actions = append(actions, meraki.ActionModel{
					Operation: "update",
					Resource:  plan.getPath() + "/" + plan.Items[i].Id.ValueString(),
					Body:      plan.Items[i].toBody(ctx, itemState),
				})
			}
			break
		}
		if !found {
			// If the item is present in plan, but not in state, we need to create it
			actions = append(actions, meraki.ActionModel{
				Operation: "create",
				Resource:  plan.getPath(),
				Body:      plan.Items[i].toBody(ctx, ResourceApplianceRFProfilesItems{}),
			})
			newIdIndexes = append(newIdIndexes, i)
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
	responseIndex := 0
	for i := range plan.Items {
		if slices.Contains(newIdIndexes, i) {
			plan.Items[i].Id = types.StringValue(res.Get("status.createdResources." + strconv.Itoa(responseIndex) + ".id").String())
			responseIndex++
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *ApplianceRFProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ResourceApplianceRFProfiles

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	actions := make([]meraki.ActionModel, len(state.Items))

	for i, item := range state.Items {
		actions[i] = meraki.ActionModel{
			Operation: "destroy",
			Resource:  state.getPath() + "/" + item.Id.ValueString(),
			Body:      "{}",
		}
	}
	if len(actions) > 0 {
		res, err := r.client.Batch(state.OrganizationId.ValueString(), actions)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ApplianceRFProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	itemIdParts := make([]string, 0)
	if strings.Contains(req.ID, ",[") {
		itemIdParts = strings.Split(strings.Split(strings.Split(req.ID, ",[")[1], "]")[0], ",")
	}
	idParts := strings.Split(strings.Split(req.ID, ",[")[0], ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		expectedIdentifier := "Expected import identifier with format: <organization_id>,<network_id>"
		expectedIdentifier += " or <organization_id>,<network_id>,[<id>,...]"
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("%s. Got: %q", expectedIdentifier, req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[1])...)

	if len(itemIdParts) > 0 {
		items := make([]ResourceApplianceRFProfilesItems, len(itemIdParts))
		for i, itemId := range itemIdParts {
			item := ResourceApplianceRFProfilesItems{}
			item.Id = types.StringValue(itemId)
			items[i] = item
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("items"), items)...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin modifyPlan
func (r *ApplianceRFProfilesResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan, state ResourceApplianceRFProfiles

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
			if item.Name.ValueString() != itemState.Name.ValueString() {
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
