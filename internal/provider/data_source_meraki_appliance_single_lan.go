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
	_ datasource.DataSource              = &ApplianceSingleLANDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceSingleLANDataSource{}
)

func NewApplianceSingleLANDataSource() datasource.DataSource {
	return &ApplianceSingleLANDataSource{}
}

type ApplianceSingleLANDataSource struct {
	client *meraki.Client
}

func (d *ApplianceSingleLANDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_single_lan"
}

func (d *ApplianceSingleLANDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Appliance Single LAN` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"appliance_ip": schema.StringAttribute{
				MarkdownDescription: "The appliance IP address of the single LAN",
				Computed:            true,
			},
			"subnet": schema.StringAttribute{
				MarkdownDescription: "The subnet of the single LAN configuration",
				Computed:            true,
			},
			"ipv6_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable IPv6 on VLAN.",
				Computed:            true,
			},
			"ipv6_prefix_assignments": schema.ListNestedAttribute{
				MarkdownDescription: "Prefix assignments on the VLAN",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"autonomous": schema.BoolAttribute{
							MarkdownDescription: "Auto assign a /64 prefix from the origin to the VLAN",
							Computed:            true,
						},
						"disabled": schema.BoolAttribute{
							MarkdownDescription: "Disable this assignment",
							Computed:            true,
						},
						"static_appliance_ip6": schema.StringAttribute{
							MarkdownDescription: "Manual configuration of the IPv6 Appliance IP",
							Computed:            true,
						},
						"static_prefix": schema.StringAttribute{
							MarkdownDescription: "Manual configuration of a /64 prefix on the VLAN",
							Computed:            true,
						},
						"origin_type": schema.StringAttribute{
							MarkdownDescription: "Type of the origin",
							Computed:            true,
						},
						"origin_interfaces": schema.ListAttribute{
							MarkdownDescription: "Interfaces associated with the prefix",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
			"mandatory_dhcp_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable Mandatory DHCP on LAN.",
				Computed:            true,
			},
		},
	}
}

func (d *ApplianceSingleLANDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceSingleLANDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplianceSingleLAN

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
