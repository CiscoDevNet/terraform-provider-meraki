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
	_ datasource.DataSource              = &AppliancePortDataSource{}
	_ datasource.DataSourceWithConfigure = &AppliancePortDataSource{}
)

func NewAppliancePortDataSource() datasource.DataSource {
	return &AppliancePortDataSource{}
}

type AppliancePortDataSource struct {
	client *meraki.Client
}

func (d *AppliancePortDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_port"
}

func (d *AppliancePortDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Appliance Port` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"port_id": schema.StringAttribute{
				MarkdownDescription: "Port ID",
				Required:            true,
			},
			"access_policy": schema.StringAttribute{
				MarkdownDescription: "The name of the policy. Only applicable to Access ports. Valid values are: `open`, `8021x-radius`, `mac-radius`, `hybris-radius` for MX64 or Z3 or any MX supporting the per port authentication feature. Otherwise, `open` is the only valid value and `open` is the default value if the field is missing.",
				Computed:            true,
			},
			"allowed_vlans": schema.StringAttribute{
				MarkdownDescription: "Comma-delimited list of the VLAN ID`s allowed on the port, or `all` to permit all VLAN`s on the port.",
				Computed:            true,
			},
			"drop_untagged_traffic": schema.BoolAttribute{
				MarkdownDescription: "Trunk port can Drop all Untagged traffic. When true, no VLAN is required. Access ports cannot have dropUntaggedTraffic set to true.",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "The status of the port",
				Computed:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The type of the port: `access` or `trunk`.",
				Computed:            true,
			},
			"vlan": schema.Int64Attribute{
				MarkdownDescription: "Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.",
				Computed:            true,
			},
		},
	}
}

func (d *AppliancePortDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *AppliancePortDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config AppliancePort

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
	config.Id = config.PortId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
