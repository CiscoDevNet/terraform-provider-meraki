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
	_ resource.ResourceWithIdentity    = &ApplianceTrafficShapingRulesResource{}
	_ resource.ResourceWithImportState = &ApplianceTrafficShapingRulesResource{}
)

func NewApplianceTrafficShapingRulesResource() resource.Resource {
	return &ApplianceTrafficShapingRulesResource{}
}

type ApplianceTrafficShapingRulesResource struct {
	client *meraki.Client
}

func (r *ApplianceTrafficShapingRulesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_traffic_shaping_rules"
}

func (r *ApplianceTrafficShapingRulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance Traffic Shaping Rules` configuration.").String,

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
			"default_rules_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network`s traffic shaping page. Note that default rules count against the rule limit of 8.").String,
				Optional:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("An array of traffic shaping rules. Rules are applied in the order that they are specified in. An empty list (or null) means no rules. Note that you are allowed a maximum of 8 rules.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dscp_tag_value": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The DSCP tag applied by your rule. null means `Do not change DSCP tag`. For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.").String,
							Optional:            true,
						},
						"priority": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A string, indicating the priority level for packets bound to your rule. Can be `low`, `normal` or `high`.").String,
							Optional:            true,
						},
						"per_client_bandwidth_limit_settings": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("How bandwidth limits are applied by your rule. Can be one of `network default`, `ignore` or `custom`.").String,
							Optional:            true,
						},
						"per_client_bandwidth_limit_down": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The maximum download limit (integer, in Kbps).").String,
							Optional:            true,
						},
						"per_client_bandwidth_limit_up": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The maximum upload limit (integer, in Kbps).").String,
							Optional:            true,
						},
						"definitions": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The type of definition. Can be one of `application`, `applicationCategory`, `host`, `port`, `ipRange` or `localNet`.").AddStringEnumDescription("application", "applicationCategory", "host", "ipRange", "localNet", "port").String,
										Required:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("application", "applicationCategory", "host", "ipRange", "localNet", "port"),
										},
									},
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("If 'type' is `host`, `port`, `ipRange` or `localNet`, then 'value' must be a string, matching either a hostname (e.g. 'somesite.com'), a port (e.g. 8080), or an IP range ('192.1.0.0', '192.1.0.0/16', or '10.1.0.0/16:80'). `localNet` also supports CIDR notation, excluding custom ports. If 'type' is `application` or `applicationCategory`, then 'value' must be an application category or application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories endpoint).").String,
										Required:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *ApplianceTrafficShapingRulesResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"network_id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("Network ID").String,
				RequiredForImport: true,
			},
		},
	}
}

func (r *ApplianceTrafficShapingRulesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ApplianceTrafficShapingRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplianceTrafficShapingRules
	var identity ApplianceTrafficShapingRulesIdentity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ApplianceTrafficShapingRules{})
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

func (r *ApplianceTrafficShapingRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplianceTrafficShapingRules
	var identity ApplianceTrafficShapingRulesIdentity

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

func (r *ApplianceTrafficShapingRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplianceTrafficShapingRules

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

func (r *ApplianceTrafficShapingRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplianceTrafficShapingRules

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
func (r *ApplianceTrafficShapingRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
		var identity ApplianceTrafficShapingRulesIdentity
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
