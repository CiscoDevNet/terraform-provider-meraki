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
	_ datasource.DataSource              = &WirelessRadioSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessRadioSettingsDataSource{}
)

func NewWirelessRadioSettingsDataSource() datasource.DataSource {
	return &WirelessRadioSettingsDataSource{}
}

type WirelessRadioSettingsDataSource struct {
	client *meraki.Client
}

func (d *WirelessRadioSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_radio_settings"
}

func (d *WirelessRadioSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Wireless Radio Settings` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: "Wireless AP serial",
				Required:            true,
			},
			"rf_profile_id": schema.StringAttribute{
				MarkdownDescription: "The ID of an RF profile to assign to the device. If the value of this parameter is null, the appropriate basic RF profile (indoor or outdoor) will be assigned to the device. Assigning an RF profile will clear ALL manually configured overrides on the device (channel width, channel, power).",
				Computed:            true,
			},
			"five_ghz_settings_channel": schema.Int64Attribute{
				MarkdownDescription: "Sets a manual channel for 5 GHz. Can be `36`, `40`, `44`, `48`, `52`, `56`, `60`, `64`, `100`, `104`, `108`, `112`, `116`, `120`, `124`, `128`, `132`, `136`, `140`, `144`, `149`, `153`, `157`, `161`, `165`, `169`, `173` or `177` or null for using auto channel.",
				Computed:            true,
			},
			"five_ghz_settings_channel_width": schema.Int64Attribute{
				MarkdownDescription: "Sets a manual channel for 5 GHz. Can be `0`, `20`, `40`, `80` or `160` or null for using auto channel width.",
				Computed:            true,
			},
			"five_ghz_settings_target_power": schema.Int64Attribute{
				MarkdownDescription: "Set a manual target power for 5 GHz (dBm). Enter null for using auto power range.",
				Computed:            true,
			},
			"two_four_ghz_settings_channel": schema.Int64Attribute{
				MarkdownDescription: "Sets a manual channel for 2.4 GHz. Can be `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `13` or `14` or null for using auto channel.",
				Computed:            true,
			},
			"two_four_ghz_settings_target_power": schema.Int64Attribute{
				MarkdownDescription: "Set a manual target power for 2.4 GHz (dBm). Enter null for using auto power range.",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessRadioSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessRadioSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessRadioSettings

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
	config.Id = config.Serial

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
