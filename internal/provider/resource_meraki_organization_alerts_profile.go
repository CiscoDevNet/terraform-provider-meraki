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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &OrganizationAlertsProfileResource{}
	_ resource.ResourceWithImportState = &OrganizationAlertsProfileResource{}
)

func NewOrganizationAlertsProfileResource() resource.Resource {
	return &OrganizationAlertsProfileResource{}
}

type OrganizationAlertsProfileResource struct {
	client *meraki.Client
}

func (r *OrganizationAlertsProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_alerts_profile"
}

func (r *OrganizationAlertsProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Organization Alerts Profile` configuration.").String,

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
			"description": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("User supplied description of the alert").String,
				Optional:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The alert type").AddStringEnumDescription("appOutage", "voipJitter", "voipMos", "voipPacketLoss", "wanLatency", "wanPacketLoss", "wanStatus", "wanUtilization").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("appOutage", "voipJitter", "voipMos", "voipPacketLoss", "wanLatency", "wanPacketLoss", "wanStatus", "wanUtilization"),
				},
			},
			"alert_condition_bit_rate_bps": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts.").String,
				Optional:            true,
			},
			"alert_condition_duration": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The total duration in seconds that the threshold should be crossed before alerting").String,
				Optional:            true,
			},
			"alert_condition_interface": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The uplink observed for the alert. interface must be one of the following: wan1, wan2, wan3, cellular").AddStringEnumDescription("cellular", "wan1", "wan2", "wan3").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("cellular", "wan1", "wan2", "wan3"),
				},
			},
			"alert_condition_jitter_ms": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The threshold the metric must cross to be valid for alerting. Used only for VoIP Jitter alerts.").String,
				Optional:            true,
			},
			"alert_condition_latency_ms": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The threshold the metric must cross to be valid for alerting. Used only for WAN Latency alerts.").String,
				Optional:            true,
			},
			"alert_condition_loss_ratio": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The threshold the metric must cross to be valid for alerting. Used only for Packet Loss alerts.").String,
				Optional:            true,
			},
			"alert_condition_mos": schema.Float64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The threshold the metric must drop below to be valid for alerting. Used only for VoIP MOS alerts.").String,
				Optional:            true,
			},
			"alert_condition_window": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The look back period in seconds for sensing the alert").String,
				Optional:            true,
			},
			"recipients_emails": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of emails that will receive information about the alert").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"recipients_http_server_ids": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list base64 encoded urls of webhook endpoints that will receive information about the alert").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"network_tags": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Networks with these tags will be monitored for the alert").String,
				ElementType:         types.StringType,
				Required:            true,
			},
		},
	}
}

func (r *OrganizationAlertsProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *OrganizationAlertsProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OrganizationAlertsProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, OrganizationAlertsProfile{})
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

func (r *OrganizationAlertsProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OrganizationAlertsProfile

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
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	if len(res.Array()) > 0 {
		res.ForEach(func(k, v gjson.Result) bool {
			if state.Id.ValueString() == v.Get("id").String() {
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

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *OrganizationAlertsProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state OrganizationAlertsProfile

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

func (r *OrganizationAlertsProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationAlertsProfile

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
func (r *OrganizationAlertsProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
