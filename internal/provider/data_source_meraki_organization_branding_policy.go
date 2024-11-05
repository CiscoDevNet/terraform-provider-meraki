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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &OrganizationBrandingPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationBrandingPolicyDataSource{}
)

func NewOrganizationBrandingPolicyDataSource() datasource.DataSource {
	return &OrganizationBrandingPolicyDataSource{}
}

type OrganizationBrandingPolicyDataSource struct {
	client *meraki.Client
}

func (d *OrganizationBrandingPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_branding_policy"
}

func (d *OrganizationBrandingPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Organization Branding Policy` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether this policy is enabled.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the Dashboard branding policy.",
				Optional:            true,
				Computed:            true,
			},
			"admin_settings_applies_to": schema.StringAttribute{
				MarkdownDescription: "Which kinds of admins this policy applies to. Can be one of `All organization admins`, `All enterprise admins`, `All network admins`, `All admins of networks...`, `All admins of networks tagged...`, `Specific admins...`, `All admins` or `All SAML admins`.",
				Computed:            true,
			},
			"admin_settings_values": schema.ListAttribute{
				MarkdownDescription: "If `appliesTo` is set to one of `Specific admins...`, `All admins of networks...` or `All admins of networks tagged...`, then you must specify this `values` property to provide the set of entities to apply the branding policy to. For `Specific admins...`, specify an array of admin IDs. For `All admins of networks...`, specify an array of network IDs and/or configuration template IDs. For `All admins of networks tagged...`, specify an array of tag names.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"custom_logo_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not there is a custom logo enabled.",
				Computed:            true,
			},
			"custom_logo_image_contents": schema.StringAttribute{
				MarkdownDescription: "The file contents (a base 64 encoded string) of your new logo.",
				Computed:            true,
			},
			"custom_logo_image_format": schema.StringAttribute{
				MarkdownDescription: "The format of the encoded contents. Supported formats are `png`, `gif`, and jpg`.",
				Computed:            true,
			},
			"help_settings_api_docs_subtab": schema.StringAttribute{
				MarkdownDescription: "The `Help -> API docs` subtab where a detailed description of the Dashboard API is listed. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
			"help_settings_cases_subtab": schema.StringAttribute{
				MarkdownDescription: "The `Help -> Cases` Dashboard subtab on which Cisco Meraki support cases for this organization can be managed. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
			"help_settings_cisco_meraki_product_documentation": schema.StringAttribute{
				MarkdownDescription: "The `Product Manuals` section of the `Help -> Get Help` subtab. Can be one of `default or inherit`, `hide`, `show`, or a replacement custom HTML string.",
				Computed:            true,
			},
			"help_settings_community_subtab": schema.StringAttribute{
				MarkdownDescription: "The `Help -> Community` subtab which provides a link to Meraki Community. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
			"help_settings_data_protection_requests_subtab": schema.StringAttribute{
				MarkdownDescription: "The `Help -> Data protection requests` Dashboard subtab on which requests to delete, restrict, or export end-user data can be audited. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
			"help_settings_firewall_info_subtab": schema.StringAttribute{
				MarkdownDescription: "The `Help -> Firewall info` subtab where necessary upstream firewall rules for communication to the Cisco Meraki cloud are listed. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
			"help_settings_get_help_subtab": schema.StringAttribute{
				MarkdownDescription: "The `Help -> Get Help` subtab on which Cisco Meraki KB, Product Manuals, and Support/Case Information are displayed. Note that if this subtab is hidden, branding customizations for the KB on `Get help`, Cisco Meraki product documentation, and support contact info will not be visible. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
			"help_settings_get_help_subtab_knowledge_base_search": schema.StringAttribute{
				MarkdownDescription: "The KB search box which appears on the Help page. Can be one of `default or inherit`, `hide`, `show`, or a replacement custom HTML string.",
				Computed:            true,
			},
			"help_settings_hardware_replacements_subtab": schema.StringAttribute{
				MarkdownDescription: "The `Help -> Replacement info` subtab where important information regarding device replacements is detailed. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
			"help_settings_help_tab": schema.StringAttribute{
				MarkdownDescription: "The Help tab, under which all support information resides. If this tab is hidden, no other `Help` branding customizations will be visible. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
			"help_settings_help_widget": schema.StringAttribute{
				MarkdownDescription: "The `Help Widget` is a support widget which provides access to live chat, documentation links, Sales contact info, and other contact avenues to reach Meraki Support. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
			"help_settings_new_features_subtab": schema.StringAttribute{
				MarkdownDescription: "The `Help -> New features` subtab where new Dashboard features are detailed. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
			"help_settings_sm_forums": schema.StringAttribute{
				MarkdownDescription: "The `SM Forums` subtab which links to community-based support for Cisco Meraki Systems Manager. Only configurable for organizations that contain Systems Manager networks. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
			"help_settings_support_contact_info": schema.StringAttribute{
				MarkdownDescription: "The `Contact Meraki Support` section of the `Help -> Get Help` subtab. Can be one of `default or inherit`, `hide`, `show`, or a replacement custom HTML string.",
				Computed:            true,
			},
			"help_settings_universal_search_knowledge_base_search": schema.StringAttribute{
				MarkdownDescription: "The universal search box always visible on Dashboard will, by default, present results from the Meraki KB. This configures whether these Meraki KB results should be returned. Can be one of `default or inherit`, `hide` or `show`.",
				Computed:            true,
			},
		},
	}
}
func (d *OrganizationBrandingPolicyDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *OrganizationBrandingPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationBrandingPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OrganizationBrandingPolicy

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res meraki.Res
	var err error
	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		if len(res.Array()) > 0 {
			res.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("name").String() {
					config.Id = types.StringValue(v.Get("id").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					res = meraki.Res{Result: v}
					return false
				}
				return true
			})
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}

	if !res.Exists() {
		res, err = d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
