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
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.ResourceWithIdentity = &ApplianceUmbrellaAccountResource{}
)

func NewApplianceUmbrellaAccountResource() resource.Resource {
	return &ApplianceUmbrellaAccountResource{}
}

type ApplianceUmbrellaAccountResource struct {
	client *meraki.Client
}

func (r *ApplianceUmbrellaAccountResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_umbrella_account"
}

func (r *ApplianceUmbrellaAccountResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance Umbrella Account` configuration.").String,

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
			"api_key": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("API key for the Umbrella account").String,
				Optional:            true,
				Sensitive:           true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"api_key_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Write-only attribute.").String,
				WriteOnly:           true,
				Optional:            true,
			},
			"api_key_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Version of api_key_wo.").String,
				Optional:            true,
			},
			"api_secret": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("API secret for the Umbrella account").String,
				Optional:            true,
				Sensitive:           true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"api_secret_wo": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Write-only attribute.").String,
				WriteOnly:           true,
				Optional:            true,
			},
			"api_secret_wo_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Version of api_secret_wo.").String,
				Optional:            true,
			},
			"umbrella_organization_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Umbrella organization ID returned when the account is connected").String,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ApplianceUmbrellaAccountResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"network_id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("Network ID").String,
				RequiredForImport: true,
			},
		},
	}
}

func (r *ApplianceUmbrellaAccountResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ApplianceUmbrellaAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplianceUmbrellaAccount
	var identity ApplianceUmbrellaAccountIdentity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ApplianceUmbrellaAccount{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.NetworkId
	plan.fromBodyUnknowns(ctx, res)
	identity.toIdentity(ctx, &plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *ApplianceUmbrellaAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplianceUmbrellaAccount
	var identity ApplianceUmbrellaAccountIdentity

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
	identity.toIdentity(ctx, &state)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// umbrellaAccountNotFound reports whether a disconnect failed only because no account is
// connected, which the API returns as a 405 with "Umbrella account not found." rather than a 404.
func umbrellaAccountNotFound(err error) bool {
	return strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "Umbrella account not found")
}

// The API has no PUT for the Umbrella connection. Update is only reached when a write-only
// credential version changes (the plain credentials force a replacement), so it reconnects:
// disconnect, then connect again with the new credentials. Write-only values are null in the
// plan, so the body is built from the configuration.
func (r *ApplianceUmbrellaAccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, config ApplianceUmbrellaAccount
	var identity ApplianceUmbrellaAccountIdentity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read config, the only place write-only values are available
	diags = req.Config.Get(ctx, &config)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	res, err := r.client.Post(fmt.Sprintf("/networks/%v/appliance/umbrella/account/disconnect", url.QueryEscape(plan.NetworkId.ValueString())), "{}")
	if err != nil && !umbrellaAccountNotFound(err) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to disconnect Umbrella account (POST), got error: %s, %s", err, res.String()))
		return
	}
	res, err = r.client.Post(plan.getPath(), config.toBody(ctx, ApplianceUmbrellaAccount{}))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to connect Umbrella account (POST), got error: %s, %s", err, res.String()))
		return
	}
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	identity.toIdentity(ctx, &plan)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)
}

// The API has no DELETE for the Umbrella connection: destroying the resource disconnects the
// account with a POST to the sibling `disconnect` endpoint.
func (r *ApplianceUmbrellaAccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplianceUmbrellaAccount

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	res, err := r.client.Post(fmt.Sprintf("/networks/%v/appliance/umbrella/account/disconnect", url.QueryEscape(state.NetworkId.ValueString())), "{}")
	if err != nil && !umbrellaAccountNotFound(err) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to disconnect Umbrella account (POST), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
// End of section. //template:end import
