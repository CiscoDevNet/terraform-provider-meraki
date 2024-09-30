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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &NetworkGroupPoliciesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkGroupPoliciesDataSource{}
)

func NewNetworkGroupPoliciesDataSource() datasource.DataSource {
	return &NetworkGroupPoliciesDataSource{}
}

type NetworkGroupPoliciesDataSource struct {
	client *meraki.Client
}

func (d *NetworkGroupPoliciesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_group_policies"
}

func (d *NetworkGroupPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Network Group Policy` configuration.",

		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The list of items",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the object",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name for your group policy. Required.",
							Computed:            true,
						},
						"splash_auth_settings": schema.StringAttribute{
							MarkdownDescription: "Whether clients bound to your policy will bypass splash authorization or behave according to the network`s rules. Can be one of `network default` or `bypass`. Only available if your network has a wireless configuration.",
							Computed:            true,
						},
						"bandwidth_settings": schema.StringAttribute{
							MarkdownDescription: "How bandwidth limits are enforced. Can be `network default`, `ignore` or `custom`.",
							Computed:            true,
						},
						"bandwidth_limit_down": schema.Int64Attribute{
							MarkdownDescription: "The maximum download limit (integer, in Kbps). null indicates no limit",
							Computed:            true,
						},
						"bandwidth_limit_up": schema.Int64Attribute{
							MarkdownDescription: "The maximum upload limit (integer, in Kbps). null indicates no limit",
							Computed:            true,
						},
						"bonjour_forwarding_settings": schema.StringAttribute{
							MarkdownDescription: "How Bonjour rules are applied. Can be `network default`, `ignore` or `custom`.",
							Computed:            true,
						},
						"bonjour_forwarding_rules": schema.ListNestedAttribute{
							MarkdownDescription: "A list of the Bonjour forwarding rules for your group policy. If `settings` is set to `custom`, at least one rule must be specified.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"description": schema.StringAttribute{
										MarkdownDescription: "A description for your Bonjour forwarding rule. Optional.",
										Computed:            true,
									},
									"vlan_id": schema.StringAttribute{
										MarkdownDescription: "The ID of the service VLAN. Required.",
										Computed:            true,
									},
									"services": schema.SetAttribute{
										MarkdownDescription: "A list of Bonjour services. At least one service must be specified. Available services are `All Services`, `AirPlay`, `AFP`, `BitTorrent`, `FTP`, `iChat`, `iTunes`, `Printers`, `Samba`, `Scanners` and `SSH`",
										ElementType:         types.StringType,
										Computed:            true,
									},
								},
							},
						},
						"content_filtering_allowed_url_patterns_settings": schema.StringAttribute{
							MarkdownDescription: "How URL patterns are applied. Can be `network default`, `append` or `override`.",
							Computed:            true,
						},
						"content_filtering_allowed_url_patterns": schema.SetAttribute{
							MarkdownDescription: "A list of URL patterns that are allowed",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"content_filtering_blocked_url_categories_settings": schema.StringAttribute{
							MarkdownDescription: "How URL categories are applied. Can be `network default`, `append` or `override`.",
							Computed:            true,
						},
						"content_filtering_blocked_url_categories": schema.SetAttribute{
							MarkdownDescription: "A list of URL categories to block",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"content_filtering_blocked_url_patterns_settings": schema.StringAttribute{
							MarkdownDescription: "How URL patterns are applied. Can be `network default`, `append` or `override`.",
							Computed:            true,
						},
						"content_filtering_blocked_url_patterns": schema.SetAttribute{
							MarkdownDescription: "A list of URL patterns that are blocked",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"firewall_and_traffic_shaping_settings": schema.StringAttribute{
							MarkdownDescription: "How firewall and traffic shaping rules are enforced. Can be `network default`, `ignore` or `custom`.",
							Computed:            true,
						},
						"l3_firewall_rules": schema.ListNestedAttribute{
							MarkdownDescription: "An ordered array of the L3 firewall rules",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"comment": schema.StringAttribute{
										MarkdownDescription: "Description of the rule (optional)",
										Computed:            true,
									},
									"dest_cidr": schema.StringAttribute{
										MarkdownDescription: "Destination IP address (in IP or CIDR notation), a fully-qualified domain name (FQDN, if your network supports it) or `any`.",
										Computed:            true,
									},
									"dest_port": schema.StringAttribute{
										MarkdownDescription: "Destination port (integer in the range 1-65535), a port range (e.g. 8080-9090), or `any`",
										Computed:            true,
									},
									"policy": schema.StringAttribute{
										MarkdownDescription: "`allow` or `deny` traffic specified by this rule",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "The type of protocol (must be `tcp`, `udp`, `icmp`, `icmp6` or `any`)",
										Computed:            true,
									},
								},
							},
						},
						"l7_firewall_rules": schema.ListNestedAttribute{
							MarkdownDescription: "An ordered array of L7 firewall rules",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"policy": schema.StringAttribute{
										MarkdownDescription: "The policy applied to matching traffic. Must be `deny`.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of the L7 Rule. Must be `application`, `applicationCategory`, `host`, `port` or `ipRange`",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "The `value` of what you want to block. If `type` is `host`, `port` or `ipRange`, `value` must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If `type` is `application` or `applicationCategory`, then `value` must be an object with an ID for the application.",
										Computed:            true,
									},
								},
							},
						},
						"traffic_shaping_rules": schema.ListNestedAttribute{
							MarkdownDescription: "An array of traffic shaping rules. Rules are applied in the order that they are specified in. An empty list (or null) means no rules. Note that you are allowed a maximum of 8 rules.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"dscp_tag_value": schema.Int64Attribute{
										MarkdownDescription: "The DSCP tag applied by your rule. null means `Do not change DSCP tag`. For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.",
										Computed:            true,
									},
									"pcp_tag_value": schema.Int64Attribute{
										MarkdownDescription: "The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority). null means `Do not set PCP tag`.",
										Computed:            true,
									},
									"priority": schema.StringAttribute{
										MarkdownDescription: "A string, indicating the priority level for packets bound to your rule. Can be `low`, `normal` or `high`.",
										Computed:            true,
									},
									"per_client_bandwidth_limits_settings": schema.StringAttribute{
										MarkdownDescription: "How bandwidth limits are applied by your rule. Can be one of `network default`, `ignore` or `custom`.",
										Computed:            true,
									},
									"per_client_bandwidth_limits_bandwidth_limits_limit_down": schema.Int64Attribute{
										MarkdownDescription: "The maximum download limit (integer, in Kbps).",
										Computed:            true,
									},
									"per_client_bandwidth_limits_bandwidth_limits_limit_up": schema.Int64Attribute{
										MarkdownDescription: "The maximum upload limit (integer, in Kbps).",
										Computed:            true,
									},
									"definitions": schema.ListNestedAttribute{
										MarkdownDescription: "A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													MarkdownDescription: "The type of definition. Can be one of `application`, `applicationCategory`, `host`, `port`, `ipRange` or `localNet`.",
													Computed:            true,
												},
												"value": schema.StringAttribute{
													MarkdownDescription: "If 'type' is `host`, `port`, `ipRange` or `localNet`, then 'value' must be a string, matching either a hostname (e.g. 'somesite.com'), a port (e.g. 8080), or an IP range ('192.1.0.0', '192.1.0.0/16', or '10.1.0.0/16:80'). `localNet` also supports CIDR notation, excluding custom ports. If 'type' is `application` or `applicationCategory`, then 'value' must be an object with the structure { 'id': 'meraki:layer7/...' }, where 'id' is the application category or application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories endpoint).",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
						"scheduling_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed.",
							Computed:            true,
						},
						"scheduling_friday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"scheduling_friday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_friday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_monday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"scheduling_monday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_monday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_saturday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"scheduling_saturday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_saturday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_sunday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"scheduling_sunday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_sunday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_thursday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"scheduling_thursday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_thursday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_tuesday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"scheduling_tuesday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_tuesday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_wednesday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"scheduling_wednesday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"scheduling_wednesday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"vlan_tagging_settings": schema.StringAttribute{
							MarkdownDescription: "How VLAN tagging is applied. Can be `network default`, `ignore` or `custom`.",
							Computed:            true,
						},
						"vlan_tagging_vlan_id": schema.StringAttribute{
							MarkdownDescription: "The ID of the vlan you want to tag. This only applies if `settings` is set to `custom`.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworkGroupPoliciesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *NetworkGroupPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkGroupPolicies

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "NetworkGroupPoliciesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "NetworkGroupPoliciesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
