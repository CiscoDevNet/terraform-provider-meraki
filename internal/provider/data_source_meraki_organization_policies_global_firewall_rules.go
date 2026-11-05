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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
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
	_ datasource.DataSource              = &OrganizationPoliciesGlobalFirewallRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationPoliciesGlobalFirewallRulesDataSource{}
)

func NewOrganizationPoliciesGlobalFirewallRulesDataSource() datasource.DataSource {
	return &OrganizationPoliciesGlobalFirewallRulesDataSource{}
}

type OrganizationPoliciesGlobalFirewallRulesDataSource struct {
	client *meraki.Client
}

func (d *OrganizationPoliciesGlobalFirewallRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_policies_global_firewall_rules"
}

func (d *OrganizationPoliciesGlobalFirewallRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Organization Policies Global Firewall Rule` configuration in bulk.").String,

		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
				Required:            true,
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The list of items",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the object",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Description of the firewall rule",
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether the rule is enabled",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the firewall rule",
							Computed:            true,
						},
						"policy": schema.StringAttribute{
							MarkdownDescription: "Rule policy - allow or deny traffic",
							Computed:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: "Rule priority (lower numbers = higher priority)",
							Computed:            true,
						},
						"ruleset_id": schema.StringAttribute{
							MarkdownDescription: "Firewall ruleset ID to associate the rule with",
							Computed:            true,
						},
						"destinations_criteria_address_ranges": schema.ListAttribute{
							MarkdownDescription: "Address ranges or addresses",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destinations_criteria_appliance_vlans": schema.ListNestedAttribute{
							MarkdownDescription: "Appliance VLANs",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"interface_id": schema.StringAttribute{
										MarkdownDescription: "Interface ID",
										Computed:            true,
									},
								},
							},
						},
						"destinations_criteria_application_categories": schema.ListNestedAttribute{
							MarkdownDescription: "Application categories",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Category ID",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Category name",
										Computed:            true,
									},
									"applications": schema.ListNestedAttribute{
										MarkdownDescription: "Applications in this category",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: "Application ID",
													Computed:            true,
												},
												"name": schema.StringAttribute{
													MarkdownDescription: "Application name",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
						"destinations_criteria_applications": schema.ListNestedAttribute{
							MarkdownDescription: "Applications",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Application ID",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "Application name",
										Computed:            true,
									},
								},
							},
						},
						"destinations_criteria_policy_object_groups": schema.ListNestedAttribute{
							MarkdownDescription: "Policy object groups",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Policy object group ID",
										Computed:            true,
									},
								},
							},
						},
						"destinations_criteria_policy_objects": schema.ListNestedAttribute{
							MarkdownDescription: "Policy objects",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Policy object ID",
										Computed:            true,
									},
								},
							},
						},
						"destinations_criteria_ports": schema.ListAttribute{
							MarkdownDescription: "Port numbers or ranges",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"destinations_criteria_services": schema.ListNestedAttribute{
							MarkdownDescription: "Protocol and port services",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Protocol (tcp, udp, etc)",
										Computed:            true,
									},
									"ports": schema.ListAttribute{
										MarkdownDescription: "Port numbers or ranges",
										ElementType:         types.StringType,
										Computed:            true,
									},
								},
							},
						},
						"destinations_match_criteria": schema.SetAttribute{
							MarkdownDescription: "Destination match criteria types",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"sources_criteria_address_ranges": schema.ListAttribute{
							MarkdownDescription: "Address ranges or addresses",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"sources_criteria_appliance_vlans": schema.ListNestedAttribute{
							MarkdownDescription: "Appliance VLANs",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"interface_id": schema.StringAttribute{
										MarkdownDescription: "Interface ID",
										Computed:            true,
									},
								},
							},
						},
						"sources_criteria_policy_object_groups": schema.ListNestedAttribute{
							MarkdownDescription: "Policy object groups",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Policy object group ID",
										Computed:            true,
									},
								},
							},
						},
						"sources_criteria_policy_objects": schema.ListNestedAttribute{
							MarkdownDescription: "Policy objects",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Policy object ID",
										Computed:            true,
									},
								},
							},
						},
						"sources_criteria_ports": schema.ListAttribute{
							MarkdownDescription: "Port numbers or ranges",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"sources_match_criteria": schema.SetAttribute{
							MarkdownDescription: "Source match criteria types",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationPoliciesGlobalFirewallRulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationPoliciesGlobalFirewallRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceOrganizationPoliciesGlobalFirewallRules

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "OrganizationPoliciesGlobalFirewallRulesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "OrganizationPoliciesGlobalFirewallRulesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
