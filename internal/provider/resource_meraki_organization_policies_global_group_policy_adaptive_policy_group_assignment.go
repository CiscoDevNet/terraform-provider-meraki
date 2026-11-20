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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.ResourceWithIdentity    = &OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource{}
	_ resource.ResourceWithImportState = &OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource{}
)

func NewOrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource() resource.Resource {
	return &OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource{}
}

type OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource struct {
	client *meraki.Client
}

func (r *OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_policies_global_group_policy_adaptive_policy_group_assignment"
}

func (r *OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Organization Policies Global Group Policy Adaptive Policy Group Assignment` configuration.").String,

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
			"policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the organization-wide policy").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"adaptive_policy_group_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the adaptive policy group assigned to the policy").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"organization_id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("Organization ID").String,
				RequiredForImport: true,
			},
			"id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("").String,
				RequiredForImport: true,
			},
		},
	}
}

func (r *OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Create assigns the adaptive policy group to the policy. The `assign` endpoint only returns
// {"success": true}, so the new assignment ID is then looked up in the assignments list.
func (r *OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignment
	var identity OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentIdentity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	body := plan.toAssignBody()
	res, err := r.client.Post(plan.getCreatePath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to assign adaptive policy group (POST), got error: %s, %s", err, res.String()))
		return
	}
	listPath := fmt.Sprintf("%s?policyIds[]=%s&adaptivePolicyGroupIds[]=%s", plan.getPath(), url.QueryEscape(plan.PolicyId.ValueString()), url.QueryEscape(plan.AdaptivePolicyGroupId.ValueString()))
	res, err = r.client.Get(listPath)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve assignments (GET), got error: %s, %s", err, res.String()))
		return
	}
	items := res.Get("items")
	if !items.Exists() {
		items = res.Result
	}
	items.ForEach(func(_, v gjson.Result) bool {
		if v.Get("policyId").String() == plan.PolicyId.ValueString() && v.Get("adaptivePolicyGroupId").String() == plan.AdaptivePolicyGroupId.ValueString() {
			plan.Id = types.StringValue(v.Get("assignmentId").String())
			return false
		}
		return true
	})
	if plan.Id.IsUnknown() || plan.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "The adaptive policy group was assigned, but the assignment was not found in the assignments list")
		return
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

func (r *OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignment
	var identity OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentIdentity

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
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	if len(res.Array()) > 0 {
		res.ForEach(func(k, v gjson.Result) bool {
			if state.Id.ValueString() == v.Get("assignmentId").String() {
				res = meraki.Res{Result: v}
				return false
			}
			return true
		})
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
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

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignment
	var identity OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentIdentity

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
	identity.toIdentity(ctx, &plan)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Delete removes the adaptive policy group from the policy with a POST to the `remove` endpoint, which
// also succeeds if the assignment no longer exists.
func (r *OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignment

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	body := state.toAssignBody()
	res, err := r.client.Post(state.getDeletePath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to remove adaptive policy group (POST), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != "" || req.Identity == nil || req.Identity.Raw.IsNull() {
		idParts := strings.Split(req.ID, ",")

		if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
			resp.Diagnostics.AddError(
				"Unexpected Import Identifier",
				fmt.Sprintf("Expected import identifier with format: <organization_id>,<id>. Got: %q", req.ID),
			)
			return
		}
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("id"), idParts[1])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
	} else {
		var identity OrganizationPoliciesGlobalGroupPolicyAdaptivePolicyGroupAssignmentIdentity
		diags := req.Identity.Get(ctx, &identity)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), identity.OrganizationId.ValueString())...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), identity.Id.ValueString())...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
