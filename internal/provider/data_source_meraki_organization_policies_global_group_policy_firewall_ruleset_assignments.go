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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &OrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentsDataSource{}
)

func NewOrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentsDataSource() datasource.DataSource {
	return &OrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentsDataSource{}
}

type OrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentsDataSource struct {
	client *meraki.Client
}

func (d *OrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_policies_global_group_policy_firewall_ruleset_assignments"
}

func (d *OrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Organization Policies Global Group Policy Firewall Ruleset Assignment` configuration in bulk.").String,

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
						"policy_id": schema.StringAttribute{
							MarkdownDescription: "ID of the policy to assign the ruleset to",
							Computed:            true,
						},
						"priority": schema.Int64Attribute{
							MarkdownDescription: "Priority of the ruleset assignment (lower numbers = higher priority)",
							Computed:            true,
						},
						"ruleset_id": schema.StringAttribute{
							MarkdownDescription: "ID of the ruleset to assign",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceOrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignments

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "OrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentsDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "OrganizationPoliciesGlobalGroupPolicyFirewallRulesetAssignmentsDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
