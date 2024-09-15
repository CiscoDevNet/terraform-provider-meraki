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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &NetworkGroupPolicyResource{}
	_ resource.ResourceWithImportState = &NetworkGroupPolicyResource{}
)

func NewNetworkGroupPolicyResource() resource.Resource {
	return &NetworkGroupPolicyResource{}
}

type NetworkGroupPolicyResource struct {
	client *meraki.Client
}

func (r *NetworkGroupPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_group_policy"
}

func (r *NetworkGroupPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Network Group Policy` configuration.").String,

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
						"services": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A list of Bonjour services. At least one service must be specified. Available services are `All Services`, `AirPlay`, `AFP`, `BitTorrent`, `FTP`, `iChat`, `iTunes`, `Printers`, `Samba`, `Scanners` and `SSH`").String,
							ElementType:         types.StringType,
							Optional:            true,
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
			"content_filtering_allowed_url_patterns": schema.ListAttribute{
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
			"content_filtering_blocked_url_categories": schema.ListAttribute{
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
			"content_filtering_blocked_url_patterns": schema.ListAttribute{
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
							MarkdownDescription: helpers.NewAttributeDescription("The `value` of what you want to block. If `type` is `host`, `port` or `ipRange`, `value` must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If `type` is `application` or `applicationCategory`, then `value` must be an object with an ID for the application.").String,
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
							Optional:            true,
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
										MarkdownDescription: helpers.NewAttributeDescription("If 'type' is `host`, `port`, `ipRange` or `localNet`, then 'value' must be a string, matching either a hostname (e.g. 'somesite.com'), a port (e.g. 8080), or an IP range ('192.1.0.0', '192.1.0.0/16', or '10.1.0.0/16:80'). `localNet` also supports CIDR notation, excluding custom ports. If 'type' is `application` or `applicationCategory`, then 'value' must be an object with the structure { 'id': 'meraki:layer7/...' }, where 'id' is the application category or application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories endpoint).").String,
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
		},
	}
}

func (r *NetworkGroupPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *NetworkGroupPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkGroupPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, NetworkGroupPolicy{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("groupPolicyId").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *NetworkGroupPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkGroupPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
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

func (r *NetworkGroupPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state NetworkGroupPolicy

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

func (r *NetworkGroupPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkGroupPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *NetworkGroupPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
