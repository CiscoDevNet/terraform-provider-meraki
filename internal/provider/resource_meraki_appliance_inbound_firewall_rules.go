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
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.ResourceWithIdentity    = &ApplianceInboundFirewallRulesResource{}
	_ resource.ResourceWithImportState = &ApplianceInboundFirewallRulesResource{}
)

func NewApplianceInboundFirewallRulesResource() resource.Resource {
	return &ApplianceInboundFirewallRulesResource{}
}

type ApplianceInboundFirewallRulesResource struct {
	client *meraki.Client
}

func (r *ApplianceInboundFirewallRulesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_inbound_firewall_rules"
}

func (r *ApplianceInboundFirewallRulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance Inbound Firewall Rules` configuration.").String,

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
			"syslog_default_rule": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Log the special default rule (boolean value - enable only if you`ve configured a syslog server) (optional)").String,
				Optional:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("An ordered array of the firewall rules (not including the default rule)").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"comment": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Description of the rule (optional)").String,
							Optional:            true,
						},
						"dest_cidr": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or `any`").String,
							Required:            true,
						},
						"dest_port": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Comma-separated list of destination port(s) (integer in the range 1-65535), or `any`").String,
							Optional:            true,
						},
						"policy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("`allow` or `deny` traffic specified by this rule").AddStringEnumDescription("allow", "deny").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("allow", "deny"),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of protocol (must be `tcp`, `udp`, `icmp`, `icmp6` or `any`)").AddStringEnumDescription("any", "icmp", "icmp6", "tcp", "udp").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("any", "icmp", "icmp6", "tcp", "udp"),
							},
						},
						"src_cidr": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Comma-separated list of source IP address(es) (in IP or CIDR notation), or `any` (note: FQDN not supported for source addresses)").String,
							Required:            true,
						},
						"src_port": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Comma-separated list of source port(s) (integer in the range 1-65535), or `any`").String,
							Optional:            true,
						},
						"syslog_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ApplianceInboundFirewallRulesResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"network_id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("Network ID").String,
				RequiredForImport: true,
			},
		},
	}
}

func (r *ApplianceInboundFirewallRulesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ApplianceInboundFirewallRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplianceInboundFirewallRules
	var identity ApplianceInboundFirewallRulesIdentity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ApplianceInboundFirewallRules{})
	res, err := r.client.Put(plan.getPath(), body)
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

func (r *ApplianceInboundFirewallRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplianceInboundFirewallRules
	var identity ApplianceInboundFirewallRulesIdentity

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read identity
	diags = req.Identity.Get(ctx, &identity)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	state.fromIdentity(ctx, &identity)

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

func (r *ApplianceInboundFirewallRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplianceInboundFirewallRules

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

func (r *ApplianceInboundFirewallRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplianceInboundFirewallRules

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	body := state.toDestroyBody(ctx)
	res, err := r.client.Put(state.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ApplianceInboundFirewallRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != "" {
		idParts := strings.Split(req.ID, ",")

		if len(idParts) != 1 || idParts[0] == "" {
			resp.Diagnostics.AddError(
				"Unexpected Import Identifier",
				fmt.Sprintf("Expected import identifier with format: <network_id>. Got: %q", req.ID),
			)
			return
		}
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
	} else {
		var identity ApplianceInboundFirewallRulesIdentity
		diags := req.Identity.Get(ctx, &identity)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), identity.NetworkId.ValueString())...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), identity.NetworkId.ValueString())...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
