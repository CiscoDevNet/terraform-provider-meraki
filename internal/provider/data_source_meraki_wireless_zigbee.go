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
	_ datasource.DataSource              = &WirelessZigbeeDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessZigbeeDataSource{}
)

func NewWirelessZigbeeDataSource() datasource.DataSource {
	return &WirelessZigbeeDataSource{}
}

type WirelessZigbeeDataSource struct {
	client *meraki.Client
}

func (d *WirelessZigbeeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_zigbee"
}

func (d *WirelessZigbeeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Wireless Zigbee` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "To enable/disable Zigbee on the network",
				Computed:            true,
			},
			"defaults_channel": schema.StringAttribute{
				MarkdownDescription: "Channel",
				Computed:            true,
			},
			"defaults_transmit_power_level": schema.Int64Attribute{
				MarkdownDescription: "Transmit Power Level",
				Computed:            true,
			},
			"iot_controller_serial": schema.StringAttribute{
				MarkdownDescription: "Device Serial number",
				Computed:            true,
			},
			"lock_management_address": schema.StringAttribute{
				MarkdownDescription: "Host Address",
				Computed:            true,
			},
			"lock_management_password": schema.StringAttribute{
				MarkdownDescription: "Password",
				Computed:            true,
			},
			"lock_management_username": schema.StringAttribute{
				MarkdownDescription: "Username",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessZigbeeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

func (d *WirelessZigbeeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessZigbee

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	networkPath := fmt.Sprintf("/networks/%v", config.NetworkId.ValueString())
	res, err := d.client.Get(networkPath)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network (GET), got error: %s, %s", err, res.String()))
		return
	}
	orgId := res.Get("organizationId").String()

	settingsPath := fmt.Sprintf("/organizations/%v/wireless/zigbee/byNetwork?networkIds[]=%v", orgId, config.NetworkId.ValueString())
	res, err = d.client.Get(settingsPath)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve zigbee settings (GET), got error: %s, %s", err, res.String()))
		return
	}

	res = meraki.Res{Result: res.Get("items.0")}

	config.fromBody(ctx, res)
	config.Id = config.NetworkId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
