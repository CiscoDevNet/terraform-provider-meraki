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
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
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
	_ resource.ResourceWithIdentity    = &OrganizationPoliciesGlobalFirewallRuleResource{}
	_ resource.ResourceWithImportState = &OrganizationPoliciesGlobalFirewallRuleResource{}
)

func NewOrganizationPoliciesGlobalFirewallRuleResource() resource.Resource {
	return &OrganizationPoliciesGlobalFirewallRuleResource{}
}

type OrganizationPoliciesGlobalFirewallRuleResource struct {
	client *meraki.Client
}

func (r *OrganizationPoliciesGlobalFirewallRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_policies_global_firewall_rule"
}

func (r *OrganizationPoliciesGlobalFirewallRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Organization Policies Global Firewall Rule` configuration.").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Description of the firewall rule").String,
				Optional:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether the rule is enabled").String,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the firewall rule").String,
				Required:            true,
			},
			"policy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rule policy - allow or deny traffic").AddStringEnumDescription("allow", "deny").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "deny"),
				},
			},
			"priority": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Rule priority (lower numbers = higher priority)").String,
				Required:            true,
			},
			"ruleset_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Firewall ruleset ID to associate the rule with").String,
				Required:            true,
			},
			"destinations_criteria_address_ranges": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Address ranges or addresses").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"destinations_criteria_appliance_vlans": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Appliance VLANs").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface ID").String,
							Required:            true,
						},
					},
				},
			},
			"destinations_criteria_application_categories": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Application categories").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Category ID").String,
							Required:            true,
						},
					},
				},
			},
			"destinations_criteria_applications": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Applications").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Application ID").String,
							Required:            true,
						},
					},
				},
			},
			"destinations_criteria_policy_object_groups": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Policy object groups").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Policy object group ID").String,
							Required:            true,
						},
					},
				},
			},
			"destinations_criteria_policy_objects": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Policy objects").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Policy object ID").String,
							Required:            true,
						},
					},
				},
			},
			"destinations_criteria_ports": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port numbers or ranges").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"destinations_criteria_services": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Protocol and port services").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Protocol (tcp, udp, etc)").String,
							Required:            true,
						},
						"ports": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Port numbers or ranges").String,
							ElementType:         types.StringType,
							Required:            true,
						},
					},
				},
			},
			"destinations_match_criteria": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Destination match criteria types").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"sources_criteria_address_ranges": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Address ranges or addresses").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"sources_criteria_appliance_vlans": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Appliance VLANs").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface ID").String,
							Required:            true,
						},
					},
				},
			},
			"sources_criteria_policy_object_groups": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Policy object groups").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Policy object group ID").String,
							Required:            true,
						},
					},
				},
			},
			"sources_criteria_policy_objects": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Policy objects").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Policy object ID").String,
							Required:            true,
						},
					},
				},
			},
			"sources_criteria_ports": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port numbers or ranges").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"sources_match_criteria": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Source match criteria types").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *OrganizationPoliciesGlobalFirewallRuleResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
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

func (r *OrganizationPoliciesGlobalFirewallRuleResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *OrganizationPoliciesGlobalFirewallRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OrganizationPoliciesGlobalFirewallRule
	var identity OrganizationPoliciesGlobalFirewallRuleIdentity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, OrganizationPoliciesGlobalFirewallRule{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("ruleId").String())
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

func (r *OrganizationPoliciesGlobalFirewallRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OrganizationPoliciesGlobalFirewallRule
	var identity OrganizationPoliciesGlobalFirewallRuleIdentity

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
			if state.Id.ValueString() == v.Get("ruleId").String() {
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

func (r *OrganizationPoliciesGlobalFirewallRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state OrganizationPoliciesGlobalFirewallRule
	var identity OrganizationPoliciesGlobalFirewallRuleIdentity

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
	identity.toIdentity(ctx, &plan)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *OrganizationPoliciesGlobalFirewallRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationPoliciesGlobalFirewallRule

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
func (r *OrganizationPoliciesGlobalFirewallRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
		var identity OrganizationPoliciesGlobalFirewallRuleIdentity
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
