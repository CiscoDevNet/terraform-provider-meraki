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
	_ datasource.DataSource              = &ApplianceOneToOneNATRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceOneToOneNATRulesDataSource{}
)

func NewApplianceOneToOneNATRulesDataSource() datasource.DataSource {
	return &ApplianceOneToOneNATRulesDataSource{}
}

type ApplianceOneToOneNATRulesDataSource struct {
	client *meraki.Client
}

func (d *ApplianceOneToOneNATRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_one_to_one_nat_rules"
}

func (d *ApplianceOneToOneNATRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Appliance One To One NAT Rules` configuration.",

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
				MarkdownDescription: "An array of 1:1 nat rules",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"lan_ip": schema.StringAttribute{
							MarkdownDescription: "The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "A descriptive name for the rule",
							Computed:            true,
						},
						"public_ip": schema.StringAttribute{
							MarkdownDescription: "The IP address that will be used to access the internal resource from the WAN",
							Computed:            true,
						},
						"uplink": schema.StringAttribute{
							MarkdownDescription: "The physical WAN interface on which the traffic will arrive (`internet1` or, if available, `internet2`)",
							Computed:            true,
						},
						"allowed_inbound": schema.ListNestedAttribute{
							MarkdownDescription: "The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Either of the following: `tcp`, `udp`, `icmp-ping` or `any`",
										Computed:            true,
									},
									"allowed_ips": schema.ListAttribute{
										MarkdownDescription: "An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges, or `any`",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"destination_ports": schema.ListAttribute{
										MarkdownDescription: "An array of ports or port ranges that will be forwarded to the host on the LAN",
										ElementType:         types.StringType,
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

func (d *ApplianceOneToOneNATRulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceOneToOneNATRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplianceOneToOneNATRules

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res meraki.Res
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
