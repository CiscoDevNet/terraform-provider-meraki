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
	_ datasource.DataSource              = &CellularGatewayPortForwardingRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &CellularGatewayPortForwardingRulesDataSource{}
)

func NewCellularGatewayPortForwardingRulesDataSource() datasource.DataSource {
	return &CellularGatewayPortForwardingRulesDataSource{}
}

type CellularGatewayPortForwardingRulesDataSource struct {
	client *meraki.Client
}

func (d *CellularGatewayPortForwardingRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cellular_gateway_port_forwarding_rules"
}

func (d *CellularGatewayPortForwardingRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Cellular Gateway Port Forwarding Rules` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: "Device serial",
				Required:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: "An array of port forwarding params",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"access": schema.StringAttribute{
							MarkdownDescription: "`any` or `restricted`. Specify the right to make inbound connections on the specified ports or port ranges. If `restricted`, a list of allowed IPs is mandatory.",
							Computed:            true,
						},
						"lan_ip": schema.StringAttribute{
							MarkdownDescription: "The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN",
							Computed:            true,
						},
						"local_port": schema.StringAttribute{
							MarkdownDescription: "A port or port ranges that will receive the forwarded traffic from the WAN",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "A descriptive name for the rule",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "TCP or UDP",
							Computed:            true,
						},
						"public_port": schema.StringAttribute{
							MarkdownDescription: "A port or port ranges that will be forwarded to the host on the LAN",
							Computed:            true,
						},
						"allowed_ips": schema.ListAttribute{
							MarkdownDescription: "An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges.",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *CellularGatewayPortForwardingRulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *CellularGatewayPortForwardingRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CellularGatewayPortForwardingRules

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
	config.Id = config.Serial

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
