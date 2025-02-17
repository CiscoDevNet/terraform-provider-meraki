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
	_ datasource.DataSource              = &OrganizationAdaptivePolicyACLDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationAdaptivePolicyACLDataSource{}
)

func NewOrganizationAdaptivePolicyACLDataSource() datasource.DataSource {
	return &OrganizationAdaptivePolicyACLDataSource{}
}

type OrganizationAdaptivePolicyACLDataSource struct {
	client *meraki.Client
}

func (d *OrganizationAdaptivePolicyACLDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_adaptive_policy_acl"
}

func (d *OrganizationAdaptivePolicyACLDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Organization Adaptive Policy ACL` configuration.").String,

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
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the adaptive policy ACL",
				Computed:            true,
			},
			"ip_version": schema.StringAttribute{
				MarkdownDescription: "IP version of adpative policy ACL. One of: `any`, `ipv4` or `ipv6`",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the adaptive policy ACL",
				Optional:            true,
				Computed:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: "An ordered array of the adaptive policy ACL rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dst_port": schema.StringAttribute{
							MarkdownDescription: "Destination port. Must be in the format of single port: `1`, port list: `1,2` or port range: `1-10`, and in the range of 1-65535, or `any`. Default is `any`.",
							Computed:            true,
						},
						"log": schema.BoolAttribute{
							MarkdownDescription: "If enabled, when this rule is hit an entry will be logged to the event log",
							Computed:            true,
						},
						"policy": schema.StringAttribute{
							MarkdownDescription: "`allow` or `deny` traffic specified by this rule.",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "The type of protocol (must be `tcp`, `udp`, `icmp` or `any`).",
							Computed:            true,
						},
						"src_port": schema.StringAttribute{
							MarkdownDescription: "Source port. Must be in the format of single port: `1`, port list: `1,2` or port range: `1-10`, and in the range of 1-65535, or `any`. Default is `any`.",
							Computed:            true,
						},
						"tcp_established": schema.BoolAttribute{
							MarkdownDescription: "If enabled, means TCP connection with this node must be established.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}
func (d *OrganizationAdaptivePolicyACLDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *OrganizationAdaptivePolicyACLDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationAdaptivePolicyACLDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OrganizationAdaptivePolicyACL

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
					config.Id = types.StringValue(v.Get("aclId").String())
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
