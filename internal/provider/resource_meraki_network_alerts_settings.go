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
	_ resource.Resource                = &NetworkAlertsSettingsResource{}
	_ resource.ResourceWithImportState = &NetworkAlertsSettingsResource{}
)

func NewNetworkAlertsSettingsResource() resource.Resource {
	return &NetworkAlertsSettingsResource{}
}

type NetworkAlertsSettingsResource struct {
	client *meraki.Client
}

func (r *NetworkAlertsSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_alerts_settings"
}

func (r *NetworkAlertsSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Network Alerts Settings` configuration.").String,

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
			"default_destinations_all_admins": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If true, then all network admins will receive emails.").String,
				Optional:            true,
			},
			"default_destinations_snmp": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network.").String,
				Optional:            true,
			},
			"default_destinations_emails": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of emails that will receive the alert(s).").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"default_destinations_http_server_ids": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of HTTP server IDs to send a Webhook to").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"muting_by_port_schedules_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("If true, then wireless unreachable alerts will be muted when caused by a port schedule").String,
				Optional:            true,
			},
			"alerts": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A boolean depicting if the alert is turned on or off").String,
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of alert").String,
							Required:            true,
						},
						"alert_destinations_all_admins": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If true, then all network admins will receive emails for this alert").String,
							Optional:            true,
						},
						"alert_destinations_snmp": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network").String,
							Optional:            true,
						},
						"alert_destinations_emails": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A list of emails that will receive information about the alert").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"alert_destinations_http_server_ids": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A list of HTTP server IDs to send a Webhook to for this alert").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"alert_destinations_sms_numbers": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A list of phone numbers that will receive text messages about the alert. Only available for sensors status alerts.").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"filters_failure_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Failure Type").String,
							Optional:            true,
						},
						"filters_lookback_window": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Loopback Window (in sec)").String,
							Optional:            true,
						},
						"filters_min_duration": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Min Duration").String,
							Optional:            true,
						},
						"filters_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name").String,
							Optional:            true,
						},
						"filters_period": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Period").String,
							Optional:            true,
						},
						"filters_priority": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Priority").String,
							Optional:            true,
						},
						"filters_regex": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Regex").String,
							Optional:            true,
						},
						"filters_selector": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Selector").String,
							Optional:            true,
						},
						"filters_ssid_num": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("SSID Number").String,
							Optional:            true,
						},
						"filters_tag": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Tag").String,
							Optional:            true,
						},
						"filters_threshold": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Threshold").String,
							Optional:            true,
						},
						"filters_timeout": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Timeout").String,
							Optional:            true,
						},
						"filters_conditions": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Conditions").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"direction": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Direction").AddStringEnumDescription("+", "-").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("+", "-"),
										},
									},
									"duration": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Duration").String,
										Optional:            true,
									},
									"threshold": schema.Float64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Threshold").String,
										Optional:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of condition").String,
										Required:            true,
									},
									"unit": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Unit").String,
										Optional:            true,
									},
								},
							},
						},
						"filters_serials": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Serials").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *NetworkAlertsSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *NetworkAlertsSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, initialState NetworkAlertsSettings

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
	helpers.SetJsonInitialState(ctx, initialState.toBody(ctx, NetworkAlertsSettings{}), resp.Private, &resp.Diagnostics)

	// Create object
	body := plan.toBody(ctx, NetworkAlertsSettings{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.NetworkId
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *NetworkAlertsSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkAlertsSettings

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

func (r *NetworkAlertsSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state NetworkAlertsSettings

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

func (r *NetworkAlertsSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkAlertsSettings

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
func (r *NetworkAlertsSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
