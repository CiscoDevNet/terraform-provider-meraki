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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SwitchAccessControlListsDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchAccessControlListsDataSource{}
)

func NewSwitchAccessControlListsDataSource() datasource.DataSource {
	return &SwitchAccessControlListsDataSource{}
}

type SwitchAccessControlListsDataSource struct {
	client *meraki.Client
}

func (d *SwitchAccessControlListsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_access_control_lists"
}

func (d *SwitchAccessControlListsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Switch Access Control Lists` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: "An ordered array of the access control list rules (not including the default rule). An empty array will clear the rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"comment": schema.StringAttribute{
							MarkdownDescription: "Description of the rule (optional).",
							Computed:            true,
						},
						"dst_cidr": schema.StringAttribute{
							MarkdownDescription: "Destination IP address (in IP or CIDR notation) or `any`.",
							Computed:            true,
						},
						"dst_port": schema.StringAttribute{
							MarkdownDescription: "Destination port. Must be in the range of 1-65535 or `any`. Default is `any`.",
							Computed:            true,
						},
						"ip_version": schema.StringAttribute{
							MarkdownDescription: "IP address version (must be `any`, `ipv4` or `ipv6`). Applicable only if network supports IPv6. Default value is `ipv4`.",
							Computed:            true,
						},
						"policy": schema.StringAttribute{
							MarkdownDescription: "`allow` or `deny` traffic specified by this rule.",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "The type of protocol (must be `tcp`, `udp`, or `any`).",
							Computed:            true,
						},
						"src_cidr": schema.StringAttribute{
							MarkdownDescription: "Source IP address (in IP or CIDR notation) or `any`.",
							Computed:            true,
						},
						"src_port": schema.StringAttribute{
							MarkdownDescription: "Source port. Must be in the range of 1-65535 or `any`. Default is `any`.",
							Computed:            true,
						},
						"vlan": schema.StringAttribute{
							MarkdownDescription: "Incoming traffic VLAN. Must be in the range of 1-4095 or `any`. Default is `any`.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SwitchAccessControlListsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchAccessControlListsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SwitchAccessControlLists

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
	config.Id = config.NetworkId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
