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
	_ datasource.DataSource              = &WirelessRFProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessRFProfilesDataSource{}
)

func NewWirelessRFProfilesDataSource() datasource.DataSource {
	return &WirelessRFProfilesDataSource{}
}

type WirelessRFProfilesDataSource struct {
	client *meraki.Client
}

func (d *WirelessRFProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_rf_profiles"
}

func (d *WirelessRFProfilesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Wireless RF Profile` configuration in bulk.").String,

		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
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
						"band_selection_type": schema.StringAttribute{
							MarkdownDescription: "Band selection can be set to either `ssid` or `ap`. This param is required on creation.",
							Computed:            true,
						},
						"client_balancing_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to best available access point. Can be either true or false. Defaults to true.",
							Computed:            true,
						},
						"min_bitrate_type": schema.StringAttribute{
							MarkdownDescription: "Minimum bitrate can be set to either `band` or `ssid`. Defaults to band.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the new profile. Must be unique. This param is required on creation.",
							Computed:            true,
						},
						"ap_band_settings_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`. Defaults to dual.",
							Computed:            true,
						},
						"ap_band_settings_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band. Can be either true or false. Defaults to true.",
							Computed:            true,
						},
						"ap_band_settings_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"five_ghz_settings_channel_width": schema.StringAttribute{
							MarkdownDescription: "Sets channel width (MHz) for 5Ghz band. Can be one of `auto`, `20`, `40` or `80`. Defaults to auto.",
							Computed:            true,
						},
						"five_ghz_settings_max_power": schema.Int64Attribute{
							MarkdownDescription: "Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 30.",
							Computed:            true,
						},
						"five_ghz_settings_min_bitrate": schema.Int64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of 5Ghz band. Can be one of `6`, `9`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 12.",
							Computed:            true,
						},
						"five_ghz_settings_min_power": schema.Int64Attribute{
							MarkdownDescription: "Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 8.",
							Computed:            true,
						},
						"five_ghz_settings_rxsop": schema.Int64Attribute{
							MarkdownDescription: "The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.",
							Computed:            true,
						},
						"five_ghz_settings_valid_auto_channels": schema.SetAttribute{
							MarkdownDescription: "Sets valid auto channels for 5Ghz band. Can be one of `36`, `40`, `44`, `48`, `52`, `56`, `60`, `64`, `100`, `104`, `108`, `112`, `116`, `120`, `124`, `128`, `132`, `136`, `140`, `144`, `149`, `153`, `157`, `161` or `165`.Defaults to [36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165].",
							ElementType:         types.Int64Type,
							Computed:            true,
						},
						"flex_radios_by_model": schema.ListNestedAttribute{
							MarkdownDescription: "Flex radios by model.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"model": schema.StringAttribute{
										MarkdownDescription: "Model of the AP",
										Computed:            true,
									},
									"bands": schema.SetAttribute{
										MarkdownDescription: "Band to use for each flex radio. For example, [`6`] will set the AP`s first flex radio to 6 GHz",
										ElementType:         types.StringType,
										Computed:            true,
									},
								},
							},
						},
						"per_ssid_settings_0_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_0_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_0_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_0_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_1_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_1_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_1_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_1_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_10_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_10_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_10_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_10_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_11_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_11_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_11_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_11_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_12_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_12_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_12_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_12_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_13_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_13_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_13_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_13_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_14_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_14_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_14_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_14_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_2_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_2_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_2_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_2_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_3_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_3_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_3_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_3_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_4_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_4_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_4_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_4_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_5_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_5_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_5_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_5_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_6_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_6_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_6_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_6_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_7_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_7_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_7_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_7_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_8_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_8_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_8_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_8_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"per_ssid_settings_9_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
							Computed:            true,
						},
						"per_ssid_settings_9_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
							Computed:            true,
						},
						"per_ssid_settings_9_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.",
							Computed:            true,
						},
						"per_ssid_settings_9_bands_enabled": schema.SetAttribute{
							MarkdownDescription: "List of enabled bands. Can include ['2.4', '5', '6', 'disabled']",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"six_ghz_settings_channel_width": schema.StringAttribute{
							MarkdownDescription: "Sets channel width (MHz) for 6Ghz band. Can be one of `0`, `20`, `40`, `80` or `160`. Defaults to 0.",
							Computed:            true,
						},
						"six_ghz_settings_max_power": schema.Int64Attribute{
							MarkdownDescription: "Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 30.",
							Computed:            true,
						},
						"six_ghz_settings_min_bitrate": schema.Int64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of 6Ghz band. Can be one of `6`, `9`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 12.",
							Computed:            true,
						},
						"six_ghz_settings_min_power": schema.Int64Attribute{
							MarkdownDescription: "Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 8.",
							Computed:            true,
						},
						"six_ghz_settings_rxsop": schema.Int64Attribute{
							MarkdownDescription: "The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.",
							Computed:            true,
						},
						"six_ghz_settings_valid_auto_channels": schema.SetAttribute{
							MarkdownDescription: "Sets valid auto channels for 6Ghz band. Can be one of `1`, `5`, `9`, `13`, `17`, `21`, `25`, `29`, `33`, `37`, `41`, `45`, `49`, `53`, `57`, `61`, `65`, `69`, `73`, `77`, `81`, `85`, `89`, `93`, `97`, `101`, `105`, `109`, `113`, `117`, `121`, `125`, `129`, `133`, `137`, `141`, `145`, `149`, `153`, `157`, `161`, `165`, `169`, `173`, `177`, `181`, `185`, `189`, `193`, `197`, `201`, `205`, `209`, `213`, `217`, `221`, `225`, `229` or `233`.Defaults to [1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233].",
							ElementType:         types.Int64Type,
							Computed:            true,
						},
						"transmission_enabled": schema.BoolAttribute{
							MarkdownDescription: "Toggle for radio transmission. When false, radios will not transmit at all.",
							Computed:            true,
						},
						"two_four_ghz_settings_ax_enabled": schema.BoolAttribute{
							MarkdownDescription: "Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.",
							Computed:            true,
						},
						"two_four_ghz_settings_max_power": schema.Int64Attribute{
							MarkdownDescription: "Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 30.",
							Computed:            true,
						},
						"two_four_ghz_settings_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: "Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 11.",
							Computed:            true,
						},
						"two_four_ghz_settings_min_power": schema.Int64Attribute{
							MarkdownDescription: "Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 5.",
							Computed:            true,
						},
						"two_four_ghz_settings_rxsop": schema.Int64Attribute{
							MarkdownDescription: "The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.",
							Computed:            true,
						},
						"two_four_ghz_settings_valid_auto_channels": schema.SetAttribute{
							MarkdownDescription: "Sets valid auto channels for 2.4Ghz band. Can be one of `1`, `6` or `11`. Defaults to [1, 6, 11].",
							ElementType:         types.Int64Type,
							Computed:            true,
						},
						"is_indoor_default": schema.BoolAttribute{
							MarkdownDescription: "Set this profile as the default indoor rf profile. If the profile ID is one of `indoor` or `outdoor`, then a new profile will be created from the respective ID and set as the default",
							Computed:            true,
						},
						"is_outdoor_default": schema.BoolAttribute{
							MarkdownDescription: "Set this profile as the default outdoor rf profile. If the profile ID is one of `indoor` or `outdoor`, then a new profile will be created from the respective ID and set as the default",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WirelessRFProfilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessRFProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceWirelessRFProfiles

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "WirelessRFProfilesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "WirelessRFProfilesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
