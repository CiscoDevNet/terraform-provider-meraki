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
	_ datasource.DataSource              = &WirelessNetworkBluetoothSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessNetworkBluetoothSettingsDataSource{}
)

func NewWirelessNetworkBluetoothSettingsDataSource() datasource.DataSource {
	return &WirelessNetworkBluetoothSettingsDataSource{}
}

type WirelessNetworkBluetoothSettingsDataSource struct {
	client *meraki.Client
}

func (d *WirelessNetworkBluetoothSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_network_bluetooth_settings"
}

func (d *WirelessNetworkBluetoothSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Wireless Network Bluetooth Settings` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"advertising_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether APs will advertise beacons.",
				Computed:            true,
			},
			"major": schema.Int64Attribute{
				MarkdownDescription: "The major number to be used in the beacon identifier. Only valid in `Non-unique` mode.",
				Computed:            true,
			},
			"major_minor_assignment_mode": schema.StringAttribute{
				MarkdownDescription: "The way major and minor number should be assigned to nodes in the network. (`Unique`, `Non-unique`)",
				Computed:            true,
			},
			"minor": schema.Int64Attribute{
				MarkdownDescription: "The minor number to be used in the beacon identifier. Only valid in `Non-unique` mode.",
				Computed:            true,
			},
			"scanning_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether APs will scan for Bluetooth enabled clients.",
				Computed:            true,
			},
			"uuid": schema.StringAttribute{
				MarkdownDescription: "The UUID to be used in the beacon identifier.",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessNetworkBluetoothSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessNetworkBluetoothSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessNetworkBluetoothSettings

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