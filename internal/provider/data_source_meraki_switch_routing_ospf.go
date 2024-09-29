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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SwitchRoutingOSPFDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchRoutingOSPFDataSource{}
)

func NewSwitchRoutingOSPFDataSource() datasource.DataSource {
	return &SwitchRoutingOSPFDataSource{}
}

type SwitchRoutingOSPFDataSource struct {
	client *meraki.Client
}

func (d *SwitchRoutingOSPFDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_routing_ospf"
}

func (d *SwitchRoutingOSPFDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Switch Routing OSPF` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"dead_timer_in_seconds": schema.Int64Attribute{
				MarkdownDescription: "Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default.",
				Computed:            true,
			},
			"hello_timer_in_seconds": schema.Int64Attribute{
				MarkdownDescription: "Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.",
				Computed:            true,
			},
			"md5_authentication_enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default.",
				Computed:            true,
			},
			"md5_authentication_key_id": schema.Int64Attribute{
				MarkdownDescription: "MD5 authentication key index. Key index must be between 1 to 255",
				Computed:            true,
			},
			"md5_authentication_key_passphrase": schema.StringAttribute{
				MarkdownDescription: "MD5 authentication passphrase",
				Computed:            true,
			},
			"v3_dead_timer_in_seconds": schema.Int64Attribute{
				MarkdownDescription: "Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535",
				Computed:            true,
			},
			"v3_enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean value to enable or disable V3 OSPF routing. OSPF V3 routing is disabled by default.",
				Computed:            true,
			},
			"v3_hello_timer_in_seconds": schema.Int64Attribute{
				MarkdownDescription: "Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.",
				Computed:            true,
			},
			"v3_areas": schema.ListNestedAttribute{
				MarkdownDescription: "OSPF v3 areas",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"area_id": schema.StringAttribute{
							MarkdownDescription: "OSPF area ID",
							Computed:            true,
						},
						"area_name": schema.StringAttribute{
							MarkdownDescription: "Name of the OSPF area",
							Computed:            true,
						},
						"area_type": schema.StringAttribute{
							MarkdownDescription: "Area types in OSPF. Must be one of: ['normal', 'stub', 'nssa']",
							Computed:            true,
						},
					},
				},
			},
			"areas": schema.ListNestedAttribute{
				MarkdownDescription: "OSPF areas",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"area_id": schema.StringAttribute{
							MarkdownDescription: "OSPF area ID",
							Computed:            true,
						},
						"area_name": schema.StringAttribute{
							MarkdownDescription: "Name of the OSPF area",
							Computed:            true,
						},
						"area_type": schema.StringAttribute{
							MarkdownDescription: "Area types in OSPF. Must be one of: ['normal', 'stub', 'nssa']",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SwitchRoutingOSPFDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchRoutingOSPFDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SwitchRoutingOSPF

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
