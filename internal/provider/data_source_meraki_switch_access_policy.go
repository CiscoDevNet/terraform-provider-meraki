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
	_ datasource.DataSource              = &SwitchAccessPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchAccessPolicyDataSource{}
)

func NewSwitchAccessPolicyDataSource() datasource.DataSource {
	return &SwitchAccessPolicyDataSource{}
}

type SwitchAccessPolicyDataSource struct {
	client *meraki.Client
}

func (d *SwitchAccessPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_access_policy"
}

func (d *SwitchAccessPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Switch Access Policy` configuration.",

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
			"access_policy_type": schema.StringAttribute{
				MarkdownDescription: "Access Type of the policy. Automatically `Hybrid authentication` when hostMode is `Multi-Domain`.",
				Computed:            true,
			},
			"guest_port_bouncing": schema.BoolAttribute{
				MarkdownDescription: "If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers",
				Computed:            true,
			},
			"guest_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "ID for the guest VLAN allow unauthorized devices access to limited network resources",
				Computed:            true,
			},
			"host_mode": schema.StringAttribute{
				MarkdownDescription: "Choose the Host Mode for the access policy.",
				Computed:            true,
			},
			"increase_access_speed": schema.BoolAttribute{
				MarkdownDescription: "Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is `Hybrid Authentication.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the access policy",
				Optional:            true,
				Computed:            true,
			},
			"radius_accounting_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients",
				Computed:            true,
			},
			"radius_coa_support_enabled": schema.BoolAttribute{
				MarkdownDescription: "Change of authentication for RADIUS re-authentication and disconnection",
				Computed:            true,
			},
			"radius_group_attribute": schema.StringAttribute{
				MarkdownDescription: "Acceptable values are `''` for None, or `'11'` for Group Policies ACL",
				Computed:            true,
			},
			"radius_testing_enabled": schema.BoolAttribute{
				MarkdownDescription: "If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers",
				Computed:            true,
			},
			"url_redirect_walled_garden_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication",
				Computed:            true,
			},
			"voice_vlan_clients": schema.BoolAttribute{
				MarkdownDescription: "CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is `Multi-Domain`.",
				Computed:            true,
			},
			"dot1x_control_direction": schema.StringAttribute{
				MarkdownDescription: "Supports either `both` or `inbound`. Set to `inbound` to allow unauthorized egress on the switchport. Set to `both` to control both traffic directions with authorization. Defaults to `both`",
				Computed:            true,
			},
			"radius_failed_auth_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth",
				Computed:            true,
			},
			"radius_re_authentication_interval": schema.Int64Attribute{
				MarkdownDescription: "Re-authentication period in seconds. Will be null if hostMode is Multi-Auth",
				Computed:            true,
			},
			"radius_cache_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable to cache authorization and authentication responses on the RADIUS server",
				Computed:            true,
			},
			"radius_cache_timeout": schema.Int64Attribute{
				MarkdownDescription: "If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication",
				Computed:            true,
			},
			"radius_critical_auth_data_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth",
				Computed:            true,
			},
			"radius_critical_auth_suspend_port_bounce": schema.BoolAttribute{
				MarkdownDescription: "Enable to suspend port bounce when RADIUS servers are unreachable",
				Computed:            true,
			},
			"radius_critical_auth_voice_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth",
				Computed:            true,
			},
			"radius_accounting_servers": schema.ListNestedAttribute{
				MarkdownDescription: "List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host": schema.StringAttribute{
							MarkdownDescription: "Public IP address of the RADIUS accounting server",
							Computed:            true,
						},
						"organization_radius_server_id": schema.StringAttribute{
							MarkdownDescription: "Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "UDP port that the RADIUS Accounting server listens on for access requests",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "RADIUS client shared secret",
							Computed:            true,
						},
					},
				},
			},
			"radius_servers": schema.ListNestedAttribute{
				MarkdownDescription: "List of RADIUS servers to require connecting devices to authenticate against before granting network access",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host": schema.StringAttribute{
							MarkdownDescription: "Public IP address of the RADIUS server",
							Computed:            true,
						},
						"organization_radius_server_id": schema.StringAttribute{
							MarkdownDescription: "Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "UDP port that the RADIUS server listens on for access requests",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "RADIUS client shared secret",
							Computed:            true,
						},
					},
				},
			},
			"url_redirect_walled_garden_ranges": schema.SetAttribute{
				MarkdownDescription: "IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}
func (d *SwitchAccessPolicyDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *SwitchAccessPolicyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchAccessPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SwitchAccessPolicy

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
					config.Id = types.StringValue(v.Get("accessPolicyNumber").String())
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
