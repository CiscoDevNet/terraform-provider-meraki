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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &OrganizationLoginSecurityDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationLoginSecurityDataSource{}
)

func NewOrganizationLoginSecurityDataSource() datasource.DataSource {
	return &OrganizationLoginSecurityDataSource{}
}

type OrganizationLoginSecurityDataSource struct {
	client *meraki.Client
}

func (d *OrganizationLoginSecurityDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_login_security"
}

func (d *OrganizationLoginSecurityDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Organization Login Security` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
				Required:            true,
			},
			"account_lockout_attempts": schema.Int64Attribute{
				MarkdownDescription: "Number of consecutive failed login attempts after which users` accounts will be locked.",
				Computed:            true,
			},
			"enforce_account_lockout": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether users` Dashboard accounts will be locked out after a specified number of consecutive failed login attempts.",
				Computed:            true,
			},
			"enforce_different_passwords": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether users, when setting a new password, are forced to choose a new password that is different from any past passwords.",
				Computed:            true,
			},
			"enforce_idle_timeout": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether users will be logged out after being idle for the specified number of minutes.",
				Computed:            true,
			},
			"enforce_login_ip_ranges": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether organization will restrict access to Dashboard (including the API) from certain IP addresses.",
				Computed:            true,
			},
			"enforce_password_expiration": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether users are forced to change their password every X number of days.",
				Computed:            true,
			},
			"enforce_strong_passwords": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether users will be forced to choose strong passwords for their accounts. Strong passwords are at least 8 characters that contain 3 of the following: number, uppercase letter, lowercase letter, and symbol",
				Computed:            true,
			},
			"enforce_two_factor_auth": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether users in this organization will be required to use an extra verification code when logging in to Dashboard. This code will be sent to their mobile phone via SMS, or can be generated by the authenticator application.",
				Computed:            true,
			},
			"idle_timeout_minutes": schema.Int64Attribute{
				MarkdownDescription: "Number of minutes users can remain idle before being logged out of their accounts.",
				Computed:            true,
			},
			"num_different_passwords": schema.Int64Attribute{
				MarkdownDescription: "Number of recent passwords that new password must be distinct from.",
				Computed:            true,
			},
			"password_expiration_days": schema.Int64Attribute{
				MarkdownDescription: "Number of days after which users will be forced to change their password.",
				Computed:            true,
			},
			"api_authentication_ip_restrictions_for_keys_enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.",
				Computed:            true,
			},
			"api_authentication_ip_restrictions_for_keys_ranges": schema.ListAttribute{
				MarkdownDescription: "List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"login_ip_ranges": schema.ListAttribute{
				MarkdownDescription: "List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *OrganizationLoginSecurityDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationLoginSecurityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OrganizationLoginSecurity

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res gjson.Result
	var err error

	if !res.Exists() {
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)
	config.Id = config.OrganizationId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read