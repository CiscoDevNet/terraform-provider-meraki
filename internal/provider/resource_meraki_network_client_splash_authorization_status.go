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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &NetworkClientSplashAuthorizationStatusResource{}
	_ resource.ResourceWithImportState = &NetworkClientSplashAuthorizationStatusResource{}
)

func NewNetworkClientSplashAuthorizationStatusResource() resource.Resource {
	return &NetworkClientSplashAuthorizationStatusResource{}
}

type NetworkClientSplashAuthorizationStatusResource struct {
	client *meraki.Client
}

func (r *NetworkClientSplashAuthorizationStatusResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_client_splash_authorization_status"
}

func (r *NetworkClientSplashAuthorizationStatusResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Network Client Splash Authorization Status` configuration.").String,

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
			"client_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ssids_0_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_1_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_10_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_11_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_12_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_13_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_14_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_2_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_3_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_4_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_5_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_6_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_7_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_8_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
			"ssids_9_is_authorized": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("New authorization status for the SSID (true, false).").String,
				Optional:            true,
			},
		},
	}
}

func (r *NetworkClientSplashAuthorizationStatusResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *NetworkClientSplashAuthorizationStatusResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, initialState NetworkClientSplashAuthorizationStatus

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
	helpers.SetJsonInitialState(ctx, initialState.toBody(ctx, NetworkClientSplashAuthorizationStatus{}), resp.Private, &resp.Diagnostics)

	// Create object
	body := plan.toBody(ctx, NetworkClientSplashAuthorizationStatus{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.ClientId
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *NetworkClientSplashAuthorizationStatusResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkClientSplashAuthorizationStatus

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

func (r *NetworkClientSplashAuthorizationStatusResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state NetworkClientSplashAuthorizationStatus

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

func (r *NetworkClientSplashAuthorizationStatusResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkClientSplashAuthorizationStatus

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	// If the resource is a singleton, we need to restore the initial state
	jsonInitialState, diags := helpers.GetJsonInitialState(ctx, req)
	jsonInitialState = state.addDeleteValues(ctx, jsonInitialState)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.Put(state.getPath(), jsonInitialState)
	if err != nil {
		resp.Diagnostics.AddWarning("Failed to restore initial state", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *NetworkClientSplashAuthorizationStatusResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_id>,<client_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
