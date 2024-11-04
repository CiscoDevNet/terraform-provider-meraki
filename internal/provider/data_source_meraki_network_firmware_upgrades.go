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
	_ datasource.DataSource              = &NetworkFirmwareUpgradesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkFirmwareUpgradesDataSource{}
)

func NewNetworkFirmwareUpgradesDataSource() datasource.DataSource {
	return &NetworkFirmwareUpgradesDataSource{}
}

type NetworkFirmwareUpgradesDataSource struct {
	client *meraki.Client
}

func (d *NetworkFirmwareUpgradesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_firmware_upgrades"
}

func (d *NetworkFirmwareUpgradesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Network Firmware Upgrades` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"timezone": schema.StringAttribute{
				MarkdownDescription: "The timezone for the network",
				Computed:            true,
			},
			"products_appliance_participate_in_next_beta_release": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the network wants beta firmware",
				Computed:            true,
			},
			"products_appliance_next_upgrade_time": schema.StringAttribute{
				MarkdownDescription: "The time of the last successful upgrade",
				Computed:            true,
			},
			"products_appliance_next_upgrade_to_version_id": schema.StringAttribute{
				MarkdownDescription: "The version ID",
				Computed:            true,
			},
			"products_camera_participate_in_next_beta_release": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the network wants beta firmware",
				Computed:            true,
			},
			"products_camera_next_upgrade_time": schema.StringAttribute{
				MarkdownDescription: "The time of the last successful upgrade",
				Computed:            true,
			},
			"products_camera_next_upgrade_to_version_id": schema.StringAttribute{
				MarkdownDescription: "The version ID",
				Computed:            true,
			},
			"products_cellular_gateway_participate_in_next_beta_release": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the network wants beta firmware",
				Computed:            true,
			},
			"products_cellular_gateway_next_upgrade_time": schema.StringAttribute{
				MarkdownDescription: "The time of the last successful upgrade",
				Computed:            true,
			},
			"products_cellular_gateway_next_upgrade_to_version_id": schema.StringAttribute{
				MarkdownDescription: "The version ID",
				Computed:            true,
			},
			"products_secure_connect_participate_in_next_beta_release": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the network wants beta firmware",
				Computed:            true,
			},
			"products_secure_connect_next_upgrade_time": schema.StringAttribute{
				MarkdownDescription: "The time of the last successful upgrade",
				Computed:            true,
			},
			"products_secure_connect_next_upgrade_to_version_id": schema.StringAttribute{
				MarkdownDescription: "The version ID",
				Computed:            true,
			},
			"products_sensor_participate_in_next_beta_release": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the network wants beta firmware",
				Computed:            true,
			},
			"products_sensor_next_upgrade_time": schema.StringAttribute{
				MarkdownDescription: "The time of the last successful upgrade",
				Computed:            true,
			},
			"products_sensor_next_upgrade_to_version_id": schema.StringAttribute{
				MarkdownDescription: "The version ID",
				Computed:            true,
			},
			"products_switch_participate_in_next_beta_release": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the network wants beta firmware",
				Computed:            true,
			},
			"products_switch_next_upgrade_time": schema.StringAttribute{
				MarkdownDescription: "The time of the last successful upgrade",
				Computed:            true,
			},
			"products_switch_next_upgrade_to_version_id": schema.StringAttribute{
				MarkdownDescription: "The version ID",
				Computed:            true,
			},
			"products_switch_catalyst_participate_in_next_beta_release": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the network wants beta firmware",
				Computed:            true,
			},
			"products_switch_catalyst_next_upgrade_time": schema.StringAttribute{
				MarkdownDescription: "The time of the last successful upgrade",
				Computed:            true,
			},
			"products_switch_catalyst_next_upgrade_to_version_id": schema.StringAttribute{
				MarkdownDescription: "The version ID",
				Computed:            true,
			},
			"products_wireless_participate_in_next_beta_release": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the network wants beta firmware",
				Computed:            true,
			},
			"products_wireless_next_upgrade_time": schema.StringAttribute{
				MarkdownDescription: "The time of the last successful upgrade",
				Computed:            true,
			},
			"products_wireless_next_upgrade_to_version_id": schema.StringAttribute{
				MarkdownDescription: "The version ID",
				Computed:            true,
			},
			"products_wireless_controller_participate_in_next_beta_release": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the network wants beta firmware",
				Computed:            true,
			},
			"products_wireless_controller_next_upgrade_time": schema.StringAttribute{
				MarkdownDescription: "The time of the last successful upgrade",
				Computed:            true,
			},
			"products_wireless_controller_next_upgrade_to_version_id": schema.StringAttribute{
				MarkdownDescription: "The version ID",
				Computed:            true,
			},
			"upgrade_window_day_of_week": schema.StringAttribute{
				MarkdownDescription: "Day of the week",
				Computed:            true,
			},
			"upgrade_window_hour_of_day": schema.StringAttribute{
				MarkdownDescription: "Hour of the day",
				Computed:            true,
			},
		},
	}
}

func (d *NetworkFirmwareUpgradesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *NetworkFirmwareUpgradesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkFirmwareUpgrades

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
