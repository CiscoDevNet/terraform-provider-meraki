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
	_ datasource.DataSource              = &NetworkMerakiAuthUserDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkMerakiAuthUserDataSource{}
)

func NewNetworkMerakiAuthUserDataSource() datasource.DataSource {
	return &NetworkMerakiAuthUserDataSource{}
}

type NetworkMerakiAuthUserDataSource struct {
	client *meraki.Client
}

func (d *NetworkMerakiAuthUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_meraki_auth_user"
}

func (d *NetworkMerakiAuthUserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Network Meraki Auth User` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"account_type": schema.StringAttribute{
				MarkdownDescription: "Authorization type for user. Can be `Guest` or `802.1X` for wireless networks, or `Client VPN` for MX networks. Defaults to `802.1X`.",
				Computed:            true,
			},
			"email": schema.StringAttribute{
				MarkdownDescription: "Email address of the user",
				Computed:            true,
			},
			"email_password_to_user": schema.BoolAttribute{
				MarkdownDescription: "Whether or not Meraki should email the password to user. Default is false.",
				Computed:            true,
			},
			"is_admin": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the user is a Dashboard administrator.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the user. Only required If the user is not a Dashboard administrator.",
				Optional:            true,
				Computed:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "The password for this user account. Only required If the user is not a Dashboard administrator.",
				Computed:            true,
			},
			"authorizations": schema.ListNestedAttribute{
				MarkdownDescription: "Authorization zones and expiration dates for the user.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"expires_at": schema.StringAttribute{
							MarkdownDescription: "Date for authorization to expire. Set to `Never` for the authorization to not expire, which is the default.",
							Computed:            true,
						},
						"ssid_number": schema.Int64Attribute{
							MarkdownDescription: "Required for wireless networks. The SSID for which the user is being authorized, which must be configured for the user`s given accountType.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *NetworkMerakiAuthUserDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *NetworkMerakiAuthUserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *NetworkMerakiAuthUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkMerakiAuthUser

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res gjson.Result
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
					res = v
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
