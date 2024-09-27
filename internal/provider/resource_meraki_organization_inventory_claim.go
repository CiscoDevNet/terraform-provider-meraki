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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource = &OrganizationInventoryClaimResource{}
)

func NewOrganizationInventoryClaimResource() resource.Resource {
	return &OrganizationInventoryClaimResource{}
}

type OrganizationInventoryClaimResource struct {
	client *meraki.Client
}

func (r *OrganizationInventoryClaimResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_inventory_claim"
}

func (r *OrganizationInventoryClaimResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource allows claiming and releasing serials from the organization inventory. It will not not touch any existing serials already claimed and not included in `serials`. Licenses and orders can only be claimed but not released.").String,

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
			"licenses": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The licenses that should be claimed").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The key of the license").String,
							Required:            true,
						},
						"mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Co-term licensing only: either `renew` or `addDevices`. `addDevices` will increase the license limit, while `renew` will extend the amount of time until expiration. Defaults to `addDevices`. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. Does not apply to organizations using per-device licensing model.").AddStringEnumDescription("addDevices", "renew").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("addDevices", "renew"),
							},
						},
					},
				},
			},
			"orders": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The numbers of the orders that should be claimed").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"serials": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The list of serials to be claimed to the organization").String,
				ElementType:         types.StringType,
				Required:            true,
			},
		},
	}
}

func (r *OrganizationInventoryClaimResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *OrganizationInventoryClaimResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OrganizationInventoryClaim

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, OrganizationInventoryClaim{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.OrganizationId

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

func (r *OrganizationInventoryClaimResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OrganizationInventoryClaim

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	res, err := r.client.Get(state.getDevicesPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve inventory (GET), got error: %s, %s", err, res.String()))
		return
	}

	var serials, resultSerials []string
	state.Serials.ElementsAs(ctx, &serials, false)
	for _, serial := range serials {
		for _, device := range res.Array() {
			rSerial := device.Get("serial").String()
			if serial == rSerial {
				resultSerials = append(resultSerials, serial)
			}
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Retrieved the following serials: %v", state.Id.ValueString(), resultSerials))
	v := make([]attr.Value, len(resultSerials))
	for r := range resultSerials {
		v[r] = types.StringValue(resultSerials[r])
	}
	state.Serials = types.SetValueMust(types.StringType, v)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *OrganizationInventoryClaimResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state OrganizationInventoryClaim

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

	var planSerials, stateSerials, claimSerials, releaseSerials []string
	plan.Serials.ElementsAs(ctx, &planSerials, false)
	state.Serials.ElementsAs(ctx, &stateSerials, false)

	for _, planSerial := range planSerials {
		found := false
		for _, stateSerial := range stateSerials {
			if planSerial == stateSerial {
				found = true
			}
		}
		if !found {
			claimSerials = append(claimSerials, planSerial)
		}
	}

	for _, stateSerial := range stateSerials {
		found := false
		for _, planSerial := range planSerials {
			if planSerial == stateSerial {
				found = true
			}
		}
		if !found {
			releaseSerials = append(releaseSerials, stateSerial)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Serials to be claimed: %v", plan.Id.ValueString(), claimSerials))
	if len(claimSerials) > 0 {
		body, _ := sjson.Set("", "serials", claimSerials)
		res, err := r.client.Post(plan.getPath(), body)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to claim inventory, got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Serials to be released: %v", plan.Id.ValueString(), releaseSerials))
	if len(releaseSerials) > 0 {
		body, _ := sjson.Set("", "serials", releaseSerials)
		res, err := r.client.Post(plan.getReleasePath(), body)
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to release inventory, got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationInventoryClaimResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationInventoryClaim

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	body := state.toBody(ctx, OrganizationInventoryClaim{})
	res, err := r.client.Post(state.getReleasePath(), body)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to release inventory, got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import
