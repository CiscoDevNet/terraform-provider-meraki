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
	_ resource.Resource                = &NetworkGroupPoliciesResource{}
	_ resource.ResourceWithImportState = &NetworkGroupPoliciesResource{}
	_ resource.ResourceWithModifyPlan  = &NetworkGroupPoliciesResource{}
)

func NewNetworkGroupPoliciesResource() resource.Resource {
	return &NetworkGroupPoliciesResource{}
}

type NetworkGroupPoliciesResource struct {
	client *meraki.Client
}

func (r *NetworkGroupPoliciesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_group_policies"
}

func (r *NetworkGroupPoliciesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Network Group Policy` configuration in bulk.").AddBulkResourceIds("name").String,

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
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name for your group policy. Required.").String,
							Required:            true,
						},
						"splash_auth_settings": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether clients bound to your policy will bypass splash authorization or behave according to the network`s rules. Can be one of `network default` or `bypass`. Only available if your network has a wireless configuration.").AddStringEnumDescription("bypass", "network default").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("bypass", "network default"),
							},
						},
						"bandwidth_settings": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("How bandwidth limits are enforced. Can be `network default`, `ignore` or `custom`.").AddStringEnumDescription("custom", "ignore", "network default").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("custom", "ignore", "network default"),
							},
						},
						"bandwidth_limit_down": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The maximum download limit (integer, in Kbps). null indicates no limit").String,
							Optional:            true,
						},
						"bandwidth_limit_up": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The maximum upload limit (integer, in Kbps). null indicates no limit").String,
							Optional:            true,
						},
						"bonjour_forwarding_settings": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("How Bonjour rules are applied. Can be `network default`, `ignore` or `custom`.").AddStringEnumDescription("custom", "ignore", "network default").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("custom", "ignore", "network default"),
							},
						},
						"bonjour_forwarding_rules": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A list of the Bonjour forwarding rules for your group policy. If `settings` is set to `custom`, at least one rule must be specified.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"description": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("A description for your Bonjour forwarding rule. Optional.").String,
										Optional:            true,
									},
									"vlan_id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The ID of the service VLAN. Required.").String,
										Required:            true,
									},
									"services": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("A list of Bonjour services. At least one service must be specified. Available services are `All Services`, `AirPlay`, `AFP`, `BitTorrent`, `FTP`, `iChat`, `iTunes`, `Printers`, `Samba`, `Scanners` and `SSH`").String,
										ElementType:         types.StringType,
										Required:            true,
									},
								},
							},
						},
						"content_filtering_allowed_url_patterns_settings": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("How URL patterns are applied. Can be `network default`, `append` or `override`.").AddStringEnumDescription("append", "network default", "override").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("append", "network default", "override"),
							},
						},
						"content_filtering_allowed_url_patterns": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A list of URL patterns that are allowed").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"content_filtering_blocked_url_categories_settings": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("How URL categories are applied. Can be `network default`, `append` or `override`.").AddStringEnumDescription("append", "network default", "override").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("append", "network default", "override"),
							},
						},
						"content_filtering_blocked_url_categories": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A list of URL categories to block").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"content_filtering_blocked_url_patterns_settings": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("How URL patterns are applied. Can be `network default`, `append` or `override`.").AddStringEnumDescription("append", "network default", "override").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("append", "network default", "override"),
							},
						},
						"content_filtering_blocked_url_patterns": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A list of URL patterns that are blocked").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"firewall_and_traffic_shaping_settings": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("How firewall and traffic shaping rules are enforced. Can be `network default`, `ignore` or `custom`.").AddStringEnumDescription("custom", "ignore", "network default").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("custom", "ignore", "network default"),
							},
						},
						"l3_firewall_rules": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("An ordered array of the L3 firewall rules").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"comment": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Description of the rule (optional)").String,
										Optional:            true,
									},
									"dest_cidr": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or `any`.").String,
										Required:            true,
									},
									"dest_port": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or `any`").String,
										Optional:            true,
									},
									"policy": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("`allow` or `deny` traffic specified by this rule").String,
										Required:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The type of protocol (must be `tcp`, `udp`, `icmp`, `icmp6` or `any`)").String,
										Required:            true,
									},
								},
							},
						},
						"l7_firewall_rules": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("An ordered array of L7 firewall rules").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"policy": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The policy applied to matching traffic. Must be `deny`.").AddStringEnumDescription("deny").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("deny"),
										},
									},
									"type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Type of the L7 Rule. Must be `application`, `applicationCategory`, `host`, `port` or `ipRange`").AddStringEnumDescription("application", "applicationCategory", "host", "ipRange", "port").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("application", "applicationCategory", "host", "ipRange", "port"),
										},
									},
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("The `value` of what you want to block. If `type` is `host`, `port` or `ipRange`, `value` must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If `type` is `application` or `applicationCategory`, then `value` must be an ID for the application.").String,
										Optional:            true,
									},
								},
							},
						},
						"traffic_shaping_rules": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("An array of traffic shaping rules. Rules are applied in the order that they are specified in. An empty list (or null) means no rules. Note that you are allowed a maximum of 8 rules.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"dscp_tag_value": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("The DSCP tag applied by your rule. null means `Do not change DSCP tag`. For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.").String,
										Optional:            true,
									},
									"pcp_tag_value": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority). null means `Do not set PCP tag`.").String,
										Optional:            true,
									},
									"priority": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("A string, indicating the priority level for packets bound to your rule. Can be `low`, `normal` or `high`.").String,
										Optional:            true,
									},
									"per_client_bandwidth_limits_settings": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("How bandwidth limits are applied by your rule. Can be one of `network default`, `ignore` or `custom`.").String,
										Optional:            true,
									},
									"per_client_bandwidth_limits_bandwidth_limits_limit_down": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("The maximum download limit (integer, in Kbps).").String,
										Optional:            true,
									},
									"per_client_bandwidth_limits_bandwidth_limits_limit_up": schema.Int64Attribute{
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
						"scheduling_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed.").String,
							Optional:            true,
						},
						"scheduling_friday_active": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.").String,
							Optional:            true,
						},
						"scheduling_friday_from": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_friday_to": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_monday_active": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.").String,
							Optional:            true,
						},
						"scheduling_monday_from": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_monday_to": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_saturday_active": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.").String,
							Optional:            true,
						},
						"scheduling_saturday_from": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_saturday_to": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_sunday_active": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.").String,
							Optional:            true,
						},
						"scheduling_sunday_from": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_sunday_to": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_thursday_active": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.").String,
							Optional:            true,
						},
						"scheduling_thursday_from": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_thursday_to": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_tuesday_active": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.").String,
							Optional:            true,
						},
						"scheduling_tuesday_from": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_tuesday_to": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_wednesday_active": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.").String,
							Optional:            true,
						},
						"scheduling_wednesday_from": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"scheduling_wednesday_to": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.").String,
							Optional:            true,
						},
						"vlan_tagging_settings": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("How VLAN tagging is applied. Can be `network default`, `ignore` or `custom`.").AddStringEnumDescription("custom", "ignore", "network default").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("custom", "ignore", "network default"),
							},
						},
						"vlan_tagging_vlan_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The ID of the vlan you want to tag. This only applies if `settings` is set to `custom`.").String,
							Optional:            true,
						},
						"force_delete": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If true, the system deletes the GP even if there are active clients using the GP. After deletion, active clients that were assigned to that Group Policy will be left without any policy applied. Default is false.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *NetworkGroupPoliciesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *NetworkGroupPoliciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ResourceNetworkGroupPolicies

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
			Body:      item.toBody(ctx, ResourceNetworkGroupPoliciesItems{}),
		}
	}
	var res meraki.Res
	var err error
	if len(actions) > 0 {
		res, err = r.client.Batch(plan.OrganizationId.ValueString(), actions)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
			return
		}
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

func (r *NetworkGroupPoliciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ResourceNetworkGroupPolicies

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
	forceDelete, diags := helpers.IsFlag(ctx, req, "force_delete")
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		if len(state.Items) > 0 {
			state.fromBodyImport(ctx, res)
		} else {
			state.fromBody(ctx, res)
		}
		if forceDelete {
			for i := range state.Items {
				state.Items[i].ForceDelete = types.BoolValue(true)
			}
		}
	} else {
		state.fromBodyPartial(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *NetworkGroupPoliciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ResourceNetworkGroupPolicies

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
			path := plan.getPath() + "/" + itemState.Id.ValueString()
			if itemState.ForceDelete.ValueBool() {
				path += "?force=true"
			}
			// If the item is present in state, but not in plan, we need to delete it
			actions = append(actions, meraki.ActionModel{
				Operation: "destroy",
				Resource:  path,
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
				Body:      plan.Items[i].toBody(ctx, ResourceNetworkGroupPoliciesItems{}),
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

func (r *NetworkGroupPoliciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ResourceNetworkGroupPolicies

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	actions := make([]meraki.ActionModel, len(state.Items))

	for i, item := range state.Items {
		path := state.getPath() + "/" + item.Id.ValueString()
		if item.ForceDelete.ValueBool() {
			path += "?force=true"
		}
		actions[i] = meraki.ActionModel{
			Operation: "destroy",
			Resource:  path,
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

func (r *NetworkGroupPoliciesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	itemIdParts := make([]string, 0)
	if strings.Contains(req.ID, ",[") {
		itemIdParts = strings.Split(strings.Split(strings.Split(req.ID, ",[")[1], "]")[0], ",")
	}
	idParts := strings.Split(strings.Split(req.ID, ",[")[0], ",")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		expectedIdentifier := "Expected import identifier with format: <organization_id>,<network_id>,<force_delete>"
		expectedIdentifier += " or <organization_id>,<network_id>,<force_delete>,[<id>,...]"
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("%s. Got: %q", expectedIdentifier, req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[1])...)
	helpers.SetFlag(ctx, "force_delete", helpers.Must(strconv.ParseBool(idParts[2])), resp.Private, &resp.Diagnostics)

	if len(itemIdParts) > 0 {
		items := make([]ResourceNetworkGroupPoliciesItems, len(itemIdParts))
		for i, itemId := range itemIdParts {
			item := ResourceNetworkGroupPoliciesItems{}
			item.Id = types.StringValue(itemId)
			item.ContentFilteringAllowedUrlPatterns = types.SetNull(types.StringType)
			item.ContentFilteringBlockedUrlCategories = types.SetNull(types.StringType)
			item.ContentFilteringBlockedUrlPatterns = types.SetNull(types.StringType)
			items[i] = item
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("items"), items)...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin modifyPlan
func (r *NetworkGroupPoliciesResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan, state ResourceNetworkGroupPolicies

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
