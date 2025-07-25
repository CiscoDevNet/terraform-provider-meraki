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
	_ datasource.DataSource              = &OrganizationAdaptivePolicyACLsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationAdaptivePolicyACLsDataSource{}
)

func NewOrganizationAdaptivePolicyACLsDataSource() datasource.DataSource {
	return &OrganizationAdaptivePolicyACLsDataSource{}
}

type OrganizationAdaptivePolicyACLsDataSource struct {
	client *meraki.Client
}

func (d *OrganizationAdaptivePolicyACLsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_adaptive_policy_acls"
}

func (d *OrganizationAdaptivePolicyACLsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Organization Adaptive Policy ACL` configuration in bulk.").String,

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
							MarkdownDescription: "Description of the adaptive policy ACL",
							Computed:            true,
						},
						"ip_version": schema.StringAttribute{
							MarkdownDescription: "IP version of adpative policy ACL. One of: `any`, `ipv4` or `ipv6`",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the adaptive policy ACL",
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
				},
			},
		},
	}
}

func (d *OrganizationAdaptivePolicyACLsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationAdaptivePolicyACLsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceOrganizationAdaptivePolicyACLs

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "OrganizationAdaptivePolicyACLsDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "OrganizationAdaptivePolicyACLsDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
