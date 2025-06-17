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
	_ resource.Resource                = &OrganizationBrandingPolicyResource{}
	_ resource.ResourceWithImportState = &OrganizationBrandingPolicyResource{}
)

func NewOrganizationBrandingPolicyResource() resource.Resource {
	return &OrganizationBrandingPolicyResource{}
}

type OrganizationBrandingPolicyResource struct {
	client *meraki.Client
}

func (r *OrganizationBrandingPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_branding_policy"
}

func (r *OrganizationBrandingPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Organization Branding Policy` configuration.").String,

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
			"enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating whether this policy is enabled.").String,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the Dashboard branding policy.").String,
				Required:            true,
			},
			"admin_settings_applies_to": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Which kinds of admins this policy applies to. Can be one of `All organization admins`, `All enterprise admins`, `All network admins`, `All admins of networks...`, `All admins of networks tagged...`, `Specific admins...`, `All admins` or `All SAML admins`.").AddStringEnumDescription("All SAML admins", "All admins", "All admins of networks tagged...", "All admins of networks...", "All enterprise admins", "All network admins", "All organization admins", "Specific admins...").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("All SAML admins", "All admins", "All admins of networks tagged...", "All admins of networks...", "All enterprise admins", "All network admins", "All organization admins", "Specific admins..."),
				},
			},
			"admin_settings_values": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If `appliesTo` is set to one of `Specific admins...`, `All admins of networks...` or `All admins of networks tagged...`, then you must specify this `values` property to provide the set of entities to apply the branding policy to. For `Specific admins...`, specify an array of admin IDs. For `All admins of networks...`, specify an array of network IDs and/or configuration template IDs. For `All admins of networks tagged...`, specify an array of tag names.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"custom_logo_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether or not there is a custom logo enabled.").String,
				Optional:            true,
			},
			"custom_logo_image_contents": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The file contents (a base 64 encoded string) of your new logo.").String,
				Optional:            true,
			},
			"custom_logo_image_format": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The format of the encoded contents. Supported formats are `png`, `gif`, and jpg`.").AddStringEnumDescription("gif", "jpg", "png").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("gif", "jpg", "png"),
				},
			},
			"help_settings_api_docs_subtab": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `Help -> API docs` subtab where a detailed description of the Dashboard API is listed. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
			"help_settings_cases_subtab": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `Help -> Cases` Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
			"help_settings_cisco_meraki_product_documentation": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `Product Manuals` section of the `Help -> Get Help` subtab. Can be one of `default or inherit`, `hide`, `show`, or a replacement custom HTML string.").String,
				Optional:            true,
			},
			"help_settings_community_subtab": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `Help -> Community` subtab which provides a link to Meraki Community. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
			"help_settings_data_protection_requests_subtab": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `Help -> Data protection requests` Dashboard subtab on which requests to delete, restrict, or export end-user data can be audited. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
			"help_settings_firewall_info_subtab": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `Help -> Firewall info` subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are listed. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
			"help_settings_get_help_subtab": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `Help -> Get Help` subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note that if this subtab is hidden, branding customizations for the KB on `Get help`, Cisco Meraki product documentation, and support contact info will not be visible. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
			"help_settings_get_help_subtab_knowledge_base_search": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The KB search box which appears on the Help page. Can be one of `default or inherit`, `hide`, `show`, or a replacement custom HTML string.").String,
				Optional:            true,
			},
			"help_settings_hardware_replacements_subtab": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `Help -> Replacement info` subtab where important information regarding device replacements is detailed. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
			"help_settings_help_tab": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The Help tab, under which all support information resides. If this tab is hidden, no other `Help` branding customizations will be visible. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
			"help_settings_help_widget": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `Help Widget` is a support widget which provides access to live chat, documentation links, Sales contact info, and other contact avenues to reach Meraki Support. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
			"help_settings_new_features_subtab": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `Help -> New features` subtab where new Dashboard features are detailed. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
			"help_settings_sm_forums": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `SM Forums` subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for organizations that contain Systems Manager networks. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
			"help_settings_support_contact_info": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The `Contact Meraki Support` section of the `Help -> Get Help` subtab. Can be one of `default or inherit`, `hide`, `show`, or a replacement custom HTML string.").String,
				Optional:            true,
			},
			"help_settings_universal_search_knowledge_base_search": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures whether these Meraki KB results should be returned. Can be one of `default or inherit`, `hide` or `show`.").AddStringEnumDescription("default or inherit", "hide", "show").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default or inherit", "hide", "show"),
				},
			},
		},
	}
}

func (r *OrganizationBrandingPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *OrganizationBrandingPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OrganizationBrandingPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, OrganizationBrandingPolicy{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *OrganizationBrandingPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OrganizationBrandingPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
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

func (r *OrganizationBrandingPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state OrganizationBrandingPolicy

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
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
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

func (r *OrganizationBrandingPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationBrandingPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *OrganizationBrandingPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <organization_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
