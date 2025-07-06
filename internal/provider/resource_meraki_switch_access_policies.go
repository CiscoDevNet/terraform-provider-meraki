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
	"slices"
	"strconv"
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
	_ resource.Resource                = &SwitchAccessPoliciesResource{}
	_ resource.ResourceWithImportState = &SwitchAccessPoliciesResource{}
	_ resource.ResourceWithModifyPlan  = &SwitchAccessPoliciesResource{}
)

func NewSwitchAccessPoliciesResource() resource.Resource {
	return &SwitchAccessPoliciesResource{}
}

type SwitchAccessPoliciesResource struct {
	client *meraki.Client
}

func (r *SwitchAccessPoliciesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_access_policies"
}

func (r *SwitchAccessPoliciesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Switch Access Policy` configuration.").AddBulkResourceIds("name").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "The organization ID",
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network ID").String,
				Required:            true,
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The list of items",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the item",
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"access_policy_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Access Type of the policy. Automatically `Hybrid authentication` when hostMode is `Multi-Domain`.").AddStringEnumDescription("802.1x", "Hybrid authentication", "MAC authentication bypass").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("802.1x", "Hybrid authentication", "MAC authentication bypass"),
							},
						},
						"guest_group_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Group policy Number for guest group policy").String,
							Optional:            true,
						},
						"guest_port_bouncing": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers").String,
							Optional:            true,
						},
						"guest_sgt_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Security Group Tag ID for guest group policy").String,
							Optional:            true,
						},
						"guest_vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("ID for the guest VLAN allow unauthorized devices access to limited network resources").String,
							Optional:            true,
						},
						"host_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choose the Host Mode for the access policy.").AddStringEnumDescription("Multi-Auth", "Multi-Domain", "Multi-Host", "Single-Host").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Multi-Auth", "Multi-Domain", "Multi-Host", "Single-Host"),
							},
						},
						"increase_access_speed": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is `Hybrid Authentication.").String,
							Optional:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of the access policy").String,
							Required:            true,
						},
						"radius_accounting_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients").String,
							Required:            true,
						},
						"radius_coa_support_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Change of authentication for RADIUS re-authentication and disconnection").String,
							Required:            true,
						},
						"radius_group_attribute": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Acceptable values are `''` for None, or `'11'` for Group Policies ACL").String,
							Optional:            true,
						},
						"radius_testing_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers").String,
							Required:            true,
						},
						"url_redirect_walled_garden_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication").String,
							Required:            true,
						},
						"voice_vlan_clients": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is `Multi-Domain`.").String,
							Optional:            true,
						},
						"dot1x_control_direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Supports either `both` or `inbound`. Set to `inbound` to allow unauthorized egress on the switchport. Set to `both` to control both traffic directions with authorization. Defaults to `both`").AddStringEnumDescription("both", "inbound").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("both", "inbound"),
							},
						},
						"radius_failed_auth_group_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Group policy Number for failed authentication group policy").String,
							Optional:            true,
						},
						"radius_failed_auth_sgt_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Security Group Tag ID for failed authentication group policy").String,
							Optional:            true,
						},
						"radius_failed_auth_vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth").String,
							Optional:            true,
						},
						"radius_pre_authentication_group_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Group policy Number for pre-authentication group policy").String,
							Optional:            true,
						},
						"radius_re_authentication_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Re-authentication period in seconds. Will be null if hostMode is Multi-Auth").String,
							Optional:            true,
						},
						"radius_authentication_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Authentication mode of the policy ( Open | Closed )").AddStringEnumDescription("Closed", "Open").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Closed", "Open"),
							},
						},
						"radius_cache_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable to cache authorization and authentication responses on the RADIUS server").String,
							Optional:            true,
						},
						"radius_cache_timeout": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication").String,
							Optional:            true,
						},
						"radius_critical_auth_data_group_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Group policy Number for data VLAN").String,
							Optional:            true,
						},
						"radius_critical_auth_data_sgt_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Security Group Tag ID for data VLAN").String,
							Optional:            true,
						},
						"radius_critical_auth_data_vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth").String,
							Optional:            true,
						},
						"radius_critical_auth_suspend_port_bounce": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable to suspend port bounce when RADIUS servers are unreachable").String,
							Optional:            true,
						},
						"radius_critical_auth_voice_group_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Group policy Number for voice VLAN").String,
							Optional:            true,
						},
						"radius_critical_auth_voice_sgt_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Security Group Tag ID for voice VLAN").String,
							Optional:            true,
						},
						"radius_critical_auth_voice_vlan_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth").String,
							Optional:            true,
						},
						"radius_accounting_servers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Public IP address of the RADIUS accounting server").String,
										Optional:            true,
									},
									"organization_radius_server_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored").String,
										Optional:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("UDP port that the RADIUS Accounting server listens on for access requests").String,
										Optional:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("RADIUS client shared secret").String,
										Optional:            true,
									},
								},
							},
						},
						"radius_servers": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of RADIUS servers to require connecting devices to authenticate against before granting network access").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Public IP address of the RADIUS server").String,
										Optional:            true,
									},
									"organization_radius_server_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored").String,
										Optional:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("UDP port that the RADIUS server listens on for access requests").String,
										Optional:            true,
									},
									"secret": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("RADIUS client shared secret").String,
										Optional:            true,
									},
								},
							},
						},
						"url_redirect_walled_garden_ranges": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *SwitchAccessPoliciesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *SwitchAccessPoliciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ResourceSwitchAccessPolicies

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	actions := make([]meraki.ActionModel, len(plan.Items))

	for i, item := range plan.Items {
		actions[i] = meraki.ActionModel{
			Operation: "create",
			Resource:  plan.getPath(),
			Body:      item.toBody(ctx, ResourceSwitchAccessPoliciesItems{}),
		}
	}
	res, err := r.client.Batch(plan.OrganizationId.ValueString(), actions)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.OrganizationId
	for i := range plan.Items {
		plan.Items[i].Id = types.StringValue(res.Get("status.createdResources." + strconv.Itoa(i) + ".id").String())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *SwitchAccessPoliciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ResourceSwitchAccessPolicies

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
		state.Id = state.OrganizationId
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

func (r *SwitchAccessPoliciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ResourceSwitchAccessPolicies

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
	var actions []meraki.ActionModel
	// If there are destroy values, we need to compare the plan and state to determine what to delete
	for _, itemState := range state.Items {
		found := false
		for _, item := range plan.Items {
			if item.Name.ValueString() != itemState.Name.ValueString() {
				continue
			}

			// If the item is present in both plan and state, we can skip it
			found = true
			break
		}
		if !found {
			// If the item is present in state, but not in plan, we need to delete it
			actions = append(actions, meraki.ActionModel{
				Operation: "destroy",
				Resource:  plan.getPath() + "/" + itemState.Id.ValueString(),
				Body:      "{}",
			})
		}
	}
	newIdIndexes := make([]int, 0)
	// Check for new and updated items
	for i := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			if plan.Items[i].Name.ValueString() != itemState.Name.ValueString() {
				continue
			}

			found = true
			// If the item is present in both plan and state, we need to check if it has changes
			hasChanges := plan.hasChanges(ctx, &state, plan.Items[i].Id.ValueString())
			if hasChanges {
				actions = append(actions, meraki.ActionModel{
					Operation: "update",
					Resource:  plan.getPath() + "/" + plan.Items[i].Id.ValueString(),
					Body:      plan.Items[i].toBody(ctx, itemState),
				})
			}
			break
		}
		if !found {
			// If the item is present in plan, but not in state, we need to create it
			actions = append(actions, meraki.ActionModel{
				Operation: "create",
				Resource:  plan.getPath(),
				Body:      plan.Items[i].toBody(ctx, ResourceSwitchAccessPoliciesItems{}),
			})
			newIdIndexes = append(newIdIndexes, i)
		}
	}

	res, err := r.client.Batch(plan.OrganizationId.ValueString(), actions)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
		return
	}
	responseIndex := 0
	for i := range plan.Items {
		if slices.Contains(newIdIndexes, i) {
			plan.Items[i].Id = types.StringValue(res.Get("status.createdResources." + strconv.Itoa(responseIndex) + ".id").String())
			responseIndex++
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *SwitchAccessPoliciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ResourceSwitchAccessPolicies

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	actions := make([]meraki.ActionModel, len(state.Items))

	for i, item := range state.Items {
		actions[i] = meraki.ActionModel{
			Operation: "destroy",
			Resource:  state.getPath() + "/" + item.Id.ValueString(),
			Body:      "{}",
		}
	}
	res, err := r.client.Batch(state.OrganizationId.ValueString(), actions)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *SwitchAccessPoliciesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <organization_id>,<network_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin modifyPlan
func (r *SwitchAccessPoliciesResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan, state ResourceSwitchAccessPolicies

	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning ModifyPlan", plan.Id.ValueString()))
	// Remove incorrectly set IDs in plan (https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#prior-state-under-lists-and-sets)
	for i, item := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			if item.Name.ValueString() != itemState.Name.ValueString() {
				continue
			}
			found = true
		}
		if !found {
			plan.Items[i].Id = types.StringUnknown()
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: ModifyPlan finished successfully", plan.Id.ValueString()))

	diags = resp.Plan.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end modifyPlan
