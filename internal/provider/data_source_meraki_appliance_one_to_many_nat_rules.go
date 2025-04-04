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
	_ datasource.DataSource              = &ApplianceOneToManyNATRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceOneToManyNATRulesDataSource{}
)

func NewApplianceOneToManyNATRulesDataSource() datasource.DataSource {
	return &ApplianceOneToManyNATRulesDataSource{}
}

type ApplianceOneToManyNATRulesDataSource struct {
	client *meraki.Client
}

func (d *ApplianceOneToManyNATRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_one_to_many_nat_rules"
}

func (d *ApplianceOneToManyNATRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Appliance One To Many NAT Rules` configuration.").String,

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
				MarkdownDescription: "An array of 1:Many nat rules",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"public_ip": schema.StringAttribute{
							MarkdownDescription: "The IP address that will be used to access the internal resource from the WAN",
							Computed:            true,
						},
						"uplink": schema.StringAttribute{
							MarkdownDescription: "The physical WAN interface on which the traffic will arrive (`internet1` or, if available, `internet2`)",
							Computed:            true,
						},
						"port_rules": schema.ListNestedAttribute{
							MarkdownDescription: "An array of associated forwarding rules",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"local_ip": schema.StringAttribute{
										MarkdownDescription: "Local IP address to which traffic will be forwarded",
										Computed:            true,
									},
									"local_port": schema.StringAttribute{
										MarkdownDescription: "Destination port of the forwarded traffic that will be sent from the MX to the specified host on the LAN. If you simply wish to forward the traffic without translating the port, this should be the same as the Public port",
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: "A description of the rule",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "`tcp` or `udp`",
										Computed:            true,
									},
									"public_port": schema.StringAttribute{
										MarkdownDescription: "Destination port of the traffic that is arriving on the WAN",
										Computed:            true,
									},
									"allowed_ips": schema.ListAttribute{
										MarkdownDescription: "Remote IP addresses or ranges that are permitted to access the internal resource via this port forwarding rule, or `any`",
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

func (d *ApplianceOneToManyNATRulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceOneToManyNATRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplianceOneToManyNATRules

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
