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
	_ resource.Resource                = &WirelessSSIDSplashSettingsResource{}
	_ resource.ResourceWithImportState = &WirelessSSIDSplashSettingsResource{}
)

func NewWirelessSSIDSplashSettingsResource() resource.Resource {
	return &WirelessSSIDSplashSettingsResource{}
}

type WirelessSSIDSplashSettingsResource struct {
	client *meraki.Client
}

func (r *WirelessSSIDSplashSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssid_splash_settings"
}

func (r *WirelessSSIDSplashSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Wireless SSID Splash Settings` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Wireless SSID number").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"allow_simultaneous_logins": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether or not to allow simultaneous logins from different devices.").String,
				Optional:            true,
			},
			"block_all_traffic_before_sign_on": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.").String,
				Optional:            true,
			},
			"controller_disconnection_behavior": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("How login attempts should be handled when the controller is unreachable. Can be either `open`, `restricted`, or `default`.").AddStringEnumDescription("default", "open", "restricted").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("default", "open", "restricted"),
				},
			},
			"redirect_url": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The custom redirect URL where the users will go after the splash page.").String,
				Optional:            true,
			},
			"splash_timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Splash timeout in minutes. This will determine how often users will see the splash page.").String,
				Optional:            true,
			},
			"splash_url": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("[optional] The custom splash URL of the click-through splash page. Note that the URL can be configured without necessarily being used. In order to enable the custom URL, see `useSplashUrl`").String,
				Optional:            true,
			},
			"theme_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The id of the selected splash theme.").String,
				Optional:            true,
			},
			"use_redirect_url": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page. A custom redirect URL must be set if this is true.").String,
				Optional:            true,
			},
			"use_splash_url": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("[optional] Boolean indicating whether the users will be redirected to the custom splash url. A custom splash URL must be set if this is true. Note that depending on your SSID`s access control settings, it may not be possible to use the custom splash URL.").String,
				Optional:            true,
			},
			"welcome_message": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The welcome message for the users on the splash page.").String,
				Optional:            true,
			},
			"billing_prepaid_access_fast_login_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether or not billing uses the fast login prepaid access option.").String,
				Optional:            true,
			},
			"billing_reply_to_email_address": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The email address that receives replies from clients.").String,
				Optional:            true,
			},
			"billing_free_access_duration_in_minutes": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("How long a device can use a network for free.").String,
				Optional:            true,
			},
			"billing_free_access_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether or not free access is enabled.").String,
				Optional:            true,
			},
			"guest_sponsorship_duration_in_minutes": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Duration in minutes of sponsored guest authorization. Must be between 1 and 60480 (6 weeks)").String,
				Optional:            true,
			},
			"guest_sponsorship_guest_can_request_timeframe": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether or not guests can specify how much time they are requesting.").String,
				Optional:            true,
			},
			"self_registration_authorization_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("How created user accounts should be authorized. Must be included in: [admin, auto, self_email]").AddStringEnumDescription("admin", "auto", "self_email").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("admin", "auto", "self_email"),
				},
			},
			"self_registration_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether or not to allow users to create their own account on the network.").String,
				Optional:            true,
			},
			"sentry_enrollment_strength": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The strength of the enforcement of selected system types. Must be one of: `focused`, `click-through`, and `strict`.").AddStringEnumDescription("click-through", "focused", "strict").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("click-through", "focused", "strict"),
				},
			},
			"sentry_enrollment_systems_manager_network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The network ID of the Systems Manager network.").String,
				Optional:            true,
			},
			"sentry_enrollment_enforced_systems": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The system types that the Sentry enforces. Must be included in: `iOS, `Android`, `macOS`, and `Windows`.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"splash_image_extension": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The extension of the image file.").String,
				Optional:            true,
			},
			"splash_image_md5": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The MD5 value of the image file. Setting this to null will remove the image from the splash page.").String,
				Optional:            true,
			},
			"splash_image_image_contents": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The file contents (a base 64 encoded string) of your new image.").String,
				Optional:            true,
			},
			"splash_image_image_format": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The format of the encoded contents. Supported formats are `png`, `gif`, and jpg`.").AddStringEnumDescription("gif", "jpg", "png").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("gif", "jpg", "png"),
				},
			},
			"splash_logo_extension": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The extension of the logo file.").String,
				Optional:            true,
			},
			"splash_logo_md5": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The MD5 value of the logo file. Setting this to null will remove the logo from the splash page.").String,
				Optional:            true,
			},
			"splash_logo_image_contents": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The file contents (a base 64 encoded string) of your new logo.").String,
				Optional:            true,
			},
			"splash_logo_image_format": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The format of the encoded contents. Supported formats are `png`, `gif`, and jpg`.").AddStringEnumDescription("gif", "jpg", "png").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("gif", "jpg", "png"),
				},
			},
			"splash_prepaid_front_extension": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The extension of the prepaid front image file.").String,
				Optional:            true,
			},
			"splash_prepaid_front_md5": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The MD5 value of the prepaid front image file. Setting this to null will remove the prepaid front from the splash page.").String,
				Optional:            true,
			},
			"splash_prepaid_front_image_contents": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The file contents (a base 64 encoded string) of your new prepaid front.").String,
				Optional:            true,
			},
			"splash_prepaid_front_image_format": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The format of the encoded contents. Supported formats are `png`, `gif`, and jpg`.").AddStringEnumDescription("gif", "jpg", "png").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("gif", "jpg", "png"),
				},
			},
		},
	}
}

func (r *WirelessSSIDSplashSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *WirelessSSIDSplashSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WirelessSSIDSplashSettings

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, WirelessSSIDSplashSettings{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.Number
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *WirelessSSIDSplashSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WirelessSSIDSplashSettings

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

func (r *WirelessSSIDSplashSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WirelessSSIDSplashSettings

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

func (r *WirelessSSIDSplashSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WirelessSSIDSplashSettings

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
func (r *WirelessSSIDSplashSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_id>,<number>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
