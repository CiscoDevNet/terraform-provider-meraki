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
	_ datasource.DataSource              = &SwitchStackRoutingStaticRoutesDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchStackRoutingStaticRoutesDataSource{}
)

func NewSwitchStackRoutingStaticRoutesDataSource() datasource.DataSource {
	return &SwitchStackRoutingStaticRoutesDataSource{}
}

type SwitchStackRoutingStaticRoutesDataSource struct {
	client *meraki.Client
}

func (d *SwitchStackRoutingStaticRoutesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_stack_routing_static_routes"
}

func (d *SwitchStackRoutingStaticRoutesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Switch Stack Routing Static Route` configuration in bulk.").String,

		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: "Switch stack ID",
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
						"advertise_via_ospf_enabled": schema.BoolAttribute{
							MarkdownDescription: "Option to advertise static route via OSPF",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name or description for layer 3 static route",
							Computed:            true,
						},
						"next_hop_ip": schema.StringAttribute{
							MarkdownDescription: "IP address of the next hop device to which the device sends its traffic for the subnet",
							Computed:            true,
						},
						"prefer_over_ospf_routes_enabled": schema.BoolAttribute{
							MarkdownDescription: "Option to prefer static route over OSPF routes",
							Computed:            true,
						},
						"subnet": schema.StringAttribute{
							MarkdownDescription: "The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SwitchStackRoutingStaticRoutesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchStackRoutingStaticRoutesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceSwitchStackRoutingStaticRoutes

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "SwitchStackRoutingStaticRoutesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "SwitchStackRoutingStaticRoutesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
