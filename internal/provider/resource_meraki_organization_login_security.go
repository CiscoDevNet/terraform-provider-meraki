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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ resource.Resource                = &OrganizationLoginSecurityResource{}
	_ resource.ResourceWithImportState = &OrganizationLoginSecurityResource{}
)

func NewOrganizationLoginSecurityResource() resource.Resource {
	return &OrganizationLoginSecurityResource{}
}

type OrganizationLoginSecurityResource struct {
	client *meraki.Client
}

func (r *OrganizationLoginSecurityResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_login_security"
}

func (r *OrganizationLoginSecurityResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Organization Login Security` configuration.").String,

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
			"account_lockout_attempts": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of consecutive failed login attempts after which users` accounts will be locked.").String,
				Optional:            true,
			},
			"enforce_account_lockout": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether users` Dashboard accounts will be locked out after a specified number of consecutive failed login attempts.").String,
				Optional:            true,
			},
			"enforce_different_passwords": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether users, when setting a new password, are forced to choose a new password that is different from any past passwords.").String,
				Optional:            true,
			},
			"enforce_idle_timeout": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether users will be logged out after being idle for the specified number of minutes.").String,
				Optional:            true,
			},
			"enforce_login_ip_ranges": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether organization will restrict access to Dashboard (including the API) from certain IP addresses.").String,
				Optional:            true,
			},
			"enforce_password_expiration": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether users are forced to change their password every X number of days.").String,
				Optional:            true,
			},
			"enforce_strong_passwords": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether users will be forced to choose strong passwords for their accounts. Strong passwords are at least 8 characters that contain 3 of the following: number, uppercase letter, lowercase letter, and symbol").String,
				Optional:            true,
			},
			"enforce_two_factor_auth": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether users in this organization will be required to use an extra verification code when logging in to Dashboard. This code will be sent to their mobile phone via SMS, or can be generated by the authenticator application.").String,
				Optional:            true,
			},
			"idle_timeout_minutes": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of minutes users can remain idle before being logged out of their accounts.").String,
				Optional:            true,
			},
			"minimum_password_length": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Minimum number of characters required in admins` passwords.").AddIntegerRangeDescription(8, 16).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(8, 16),
				},
			},
			"num_different_passwords": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of recent passwords that new password must be distinct from.").String,
				Optional:            true,
			},
			"password_expiration_days": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of days after which users will be forced to change their password.").String,
				Optional:            true,
			},
			"api_authentication_ip_restrictions_for_keys_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.").String,
				Optional:            true,
			},
			"api_authentication_ip_restrictions_for_keys_ranges": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"login_ip_ranges": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *OrganizationLoginSecurityResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *OrganizationLoginSecurityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, initialState OrganizationLoginSecurity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))
	// If the resource is a singleton, we need to read and save the initial state
	gres, err := r.client.Get(plan.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, gres.String()))
		return
	}
	initialState.fromBody(ctx, gres)
	helpers.SetJsonInitialState(ctx, initialState.toBody(ctx, OrganizationLoginSecurity{}), resp.Private, &resp.Diagnostics)

	// Create object
	body := plan.toBody(ctx, OrganizationLoginSecurity{})
	res, err := r.client.Put(plan.getPath(), body)
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

func (r *OrganizationLoginSecurityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OrganizationLoginSecurity

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
		state.fromBody(ctx, res)
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

func (r *OrganizationLoginSecurityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state OrganizationLoginSecurity

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

	body := plan.toBody(ctx, state)
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *OrganizationLoginSecurityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationLoginSecurity

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	// If the resource is a singleton, we need to restore the initial state
	jsonInitialState, diags := helpers.GetJsonInitialState(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.Put(state.getPath(), jsonInitialState)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *OrganizationLoginSecurityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <organization_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
