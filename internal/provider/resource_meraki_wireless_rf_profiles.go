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
	"net/url"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &WirelessRFProfilesResource{}
	_ resource.ResourceWithImportState = &WirelessRFProfilesResource{}
	_ resource.ResourceWithModifyPlan  = &WirelessRFProfilesResource{}
)

func NewWirelessRFProfilesResource() resource.Resource {
	return &WirelessRFProfilesResource{}
}

type WirelessRFProfilesResource struct {
	client *meraki.Client
}

func (r *WirelessRFProfilesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_rf_profiles"
}

func (r *WirelessRFProfilesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Wireless RF Profile` configuration in bulk.").AddBulkResourceIds("name", ).String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "The organization ID",
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network ID").String,
				Required:            true,
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The list of items",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the item",
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"band_selection_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Band selection can be set to either `ssid` or `ap`. This param is required on creation.").AddStringEnumDescription("ap", "ssid", ).String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("ap", "ssid", ),
							},
						},
						"client_balancing_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to best available access point. Can be either true or false. Defaults to true.").String,
							Optional:            true,
						},
						"min_bitrate_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Minimum bitrate can be set to either `band` or `ssid`. Defaults to band.").AddStringEnumDescription("band", "ssid", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("band", "ssid", ),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the new profile. Must be unique. This param is required on creation.").String,
							Required:            true,
						},
						"ap_band_settings_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`. Defaults to dual.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"ap_band_settings_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band. Can be either true or false. Defaults to true.").String,
							Optional:            true,
						},
						"ap_band_settings_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"five_ghz_settings_channel_width": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets channel width (MHz) for 5Ghz band. Can be one of `auto`, `20`, `40` or `80`. Defaults to auto.").String,
							Optional:            true,
						},
						"five_ghz_settings_max_power": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 30.").String,
							Optional:            true,
						},
						"five_ghz_settings_min_bitrate": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of 5Ghz band. Can be one of `6`, `9`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 12.").String,
							Optional:            true,
						},
						"five_ghz_settings_min_power": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 8.").String,
							Optional:            true,
						},
						"five_ghz_settings_rxsop": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.").String,
							Optional:            true,
						},
						"five_ghz_settings_valid_auto_channels": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets valid auto channels for 5Ghz band. Can be one of `36`, `40`, `44`, `48`, `52`, `56`, `60`, `64`, `100`, `104`, `108`, `112`, `116`, `120`, `124`, `128`, `132`, `136`, `140`, `144`, `149`, `153`, `157`, `161` or `165`.Defaults to [36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165].").String,
							ElementType:         types.Int64Type,
							Optional:            true,
						},
						"flex_radios_by_model": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Flex radios by model.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"model": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Model of the AP").String,
										Optional:            true,
									},
									"bands": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Band to use for each flex radio. For example, [`6`] will set the AP`s first flex radio to 6 GHz").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
								},
							},
						},
						"per_ssid_settings_0_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_0_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_0_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_0_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_1_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_1_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_1_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_1_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_10_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_10_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_10_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_10_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_11_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_11_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_11_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_11_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_12_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_12_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_12_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_12_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_13_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_13_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_13_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_13_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_14_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_14_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_14_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_14_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_2_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_2_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_2_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_2_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_3_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_3_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_3_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_3_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_4_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_4_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_4_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_4_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_5_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_5_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_5_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_5_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_6_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_6_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_6_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_6_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_7_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_7_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_7_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_7_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_8_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_8_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_8_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_8_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"per_ssid_settings_9_band_operation_mode": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.").AddStringEnumDescription("2.4ghz", "5ghz", "6ghz", "dual", "multi", ).String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("2.4ghz", "5ghz", "6ghz", "dual", "multi", ),
							},
						},
						"per_ssid_settings_9_band_steering_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.").String,
							Optional:            true,
						},
						"per_ssid_settings_9_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of this SSID. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`.").String,
							Optional:            true,
						},
						"per_ssid_settings_9_bands_enabled": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of enabled bands. Can include ['2.4', '5', '6', 'disabled']").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
						"six_ghz_settings_channel_width": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets channel width (MHz) for 6Ghz band. Can be one of `0`, `20`, `40`, `80` or `160`. Defaults to 0.").String,
							Optional:            true,
						},
						"six_ghz_settings_max_power": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 30.").String,
							Optional:            true,
						},
						"six_ghz_settings_min_bitrate": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of 6Ghz band. Can be one of `6`, `9`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 12.").String,
							Optional:            true,
						},
						"six_ghz_settings_min_power": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 8.").String,
							Optional:            true,
						},
						"six_ghz_settings_rxsop": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.").String,
							Optional:            true,
						},
						"six_ghz_settings_valid_auto_channels": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets valid auto channels for 6Ghz band. Can be one of `1`, `5`, `9`, `13`, `17`, `21`, `25`, `29`, `33`, `37`, `41`, `45`, `49`, `53`, `57`, `61`, `65`, `69`, `73`, `77`, `81`, `85`, `89`, `93`, `97`, `101`, `105`, `109`, `113`, `117`, `121`, `125`, `129`, `133`, `137`, `141`, `145`, `149`, `153`, `157`, `161`, `165`, `169`, `173`, `177`, `181`, `185`, `189`, `193`, `197`, `201`, `205`, `209`, `213`, `217`, `221`, `225`, `229` or `233`.Defaults to [1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233].").String,
							ElementType:         types.Int64Type,
							Optional:            true,
						},
						"transmission_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Toggle for radio transmission. When false, radios will not transmit at all.").String,
							Optional:            true,
						},
						"two_four_ghz_settings_ax_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.").String,
							Optional:            true,
						},
						"two_four_ghz_settings_max_power": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 30.").String,
							Optional:            true,
						},
						"two_four_ghz_settings_min_bitrate": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 11.").String,
							Optional:            true,
						},
						"two_four_ghz_settings_min_power": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 5.").String,
							Optional:            true,
						},
						"two_four_ghz_settings_rxsop": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.").String,
							Optional:            true,
						},
						"two_four_ghz_settings_valid_auto_channels": schema.SetAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Sets valid auto channels for 2.4Ghz band. Can be one of `1`, `6` or `11`. Defaults to [1, 6, 11].").String,
							ElementType:         types.Int64Type,
							Optional:            true,
						},
						"is_indoor_default": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set this profile as the default indoor rf profile. If the profile ID is one of `indoor` or `outdoor`, then a new profile will be created from the respective ID and set as the default").String,
							Optional:            true,
						},
						"is_outdoor_default": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set this profile as the default outdoor rf profile. If the profile ID is one of `indoor` or `outdoor`, then a new profile will be created from the respective ID and set as the default").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *WirelessRFProfilesResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *WirelessRFProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ResourceWirelessRFProfiles

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	actions := make([]meraki.ActionModel, len(plan.Items))

	for i, item := range plan.Items {
		actions[i] = meraki.ActionModel{
			Operation: "create",
			Resource:  plan.getPath(),
			Body:      item.toBody(ctx, ResourceWirelessRFProfilesItems{}),
		}
	}
	var res meraki.Res
	var err error
	if len(actions) > 0 {
		res, err = r.client.Batch(plan.OrganizationId.ValueString(), actions)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
			return
		}
	}
	plan.Id = plan.OrganizationId
	for i := range plan.Items {
		plan.Items[i].Id = types.StringValue(res.Get("status.createdResources." + strconv.Itoa(i) + ".id").String())
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *WirelessRFProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ResourceWirelessRFProfiles

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath())
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		if len(state.Items) > 0 {
			state.fromBodyImport(ctx, res)
		} else {
			state.fromBody(ctx, res)
		}
	} else {
		state.fromBodyPartial(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *WirelessRFProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ResourceWirelessRFProfiles

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))
	var actions []meraki.ActionModel
	// If there are destroy values, we need to compare the plan and state to determine what to delete
	for _, itemState := range state.Items {
		found := false
		for _, item := range plan.Items {
			if item.Name.ValueString() != itemState.Name.ValueString() {
				continue
			}

			// If the item is present in both plan and state, we can skip it
			found = true
			break
		}
		if !found {
			// If the item is present in state, but not in plan, we need to delete it
			actions = append(actions, meraki.ActionModel{
				Operation: "destroy",
				Resource:  plan.getPath() + "/" + itemState.Id.ValueString(),
				Body:      "{}",
			})
		}
	}
	newIdIndexes := make([]int, 0)
	// Check for new and updated items
	for i := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			if plan.Items[i].Name.ValueString() != itemState.Name.ValueString() {
				continue
			}

			found = true
			// If the item is present in both plan and state, we need to check if it has changes
			hasChanges := plan.hasChanges(ctx, &state, plan.Items[i].Id.ValueString())
			if hasChanges {
				actions = append(actions, meraki.ActionModel{
					Operation: "update",
					Resource:  plan.getPath() + "/" + plan.Items[i].Id.ValueString(),
					Body:      plan.Items[i].toBody(ctx, itemState),
				})					
			}
			break
		}
		if !found {
			// If the item is present in plan, but not in state, we need to create it
			actions = append(actions, meraki.ActionModel{
				Operation: "create",
				Resource:  plan.getPath(),
				Body:      plan.Items[i].toBody(ctx, ResourceWirelessRFProfilesItems{}),
			})
			newIdIndexes = append(newIdIndexes, i)
		}
	}

	var res meraki.Res
	var err error
	if len(actions) > 0 {
		res, err = r.client.Batch(plan.OrganizationId.ValueString(), actions)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
			return
		}
	}
	responseIndex := 0
	for i := range plan.Items {
		if slices.Contains(newIdIndexes, i) {
			plan.Items[i].Id = types.StringValue(res.Get("status.createdResources." + strconv.Itoa(responseIndex) + ".id").String())
			responseIndex++
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *WirelessRFProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ResourceWirelessRFProfiles

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	actions := make([]meraki.ActionModel, len(state.Items))

	for i, item := range state.Items {
		actions[i] = meraki.ActionModel{
			Operation: "destroy",
			Resource:  state.getPath() + "/" + item.Id.ValueString(),
			Body:      "{}",
		}
	}
	if len(actions) > 0 {
		res, err := r.client.Batch(state.OrganizationId.ValueString(), actions)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *WirelessRFProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	itemIdParts := make([]string, 0)
	if strings.Contains(req.ID, ",[") {
		itemIdParts = strings.Split(strings.Split(strings.Split(req.ID, ",[")[1], "]")[0], ",")
	}
	idParts := strings.Split(strings.Split(req.ID, ",[")[0], ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		expectedIdentifier := "Expected import identifier with format: <organization_id>,<network_id>"
		expectedIdentifier += " or <organization_id>,<network_id>,[<id>,...]"
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("%s. Got: %q",expectedIdentifier, req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[1])...)

	if len(itemIdParts) > 0 {
		items := make([]ResourceWirelessRFProfilesItems, len(itemIdParts))
		for i, itemId := range itemIdParts {
			item := ResourceWirelessRFProfilesItems{}
			item.Id = types.StringValue(itemId)
			item.ApBandSettingsBandsEnabled = types.SetNull(types.StringType)
			item.FiveGhzSettingsValidAutoChannels = types.SetNull(types.Int64Type)
			item.PerSsidSettings0BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings1BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings10BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings11BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings12BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings13BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings14BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings2BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings3BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings4BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings5BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings6BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings7BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings8BandsEnabled = types.SetNull(types.StringType)
			item.PerSsidSettings9BandsEnabled = types.SetNull(types.StringType)
			item.SixGhzSettingsValidAutoChannels = types.SetNull(types.Int64Type)
			item.TwoFourGhzSettingsValidAutoChannels = types.SetNull(types.Int64Type)
			items[i] = item
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("items"), items)...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
// End of section. //template:end import

// Section below is generated&owned by "gen/generator.go". //template:begin modifyPlan
func (r *WirelessRFProfilesResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan, state ResourceWirelessRFProfiles

	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning ModifyPlan", plan.Id.ValueString()))
	// Remove incorrectly set IDs in plan (https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification#prior-state-under-lists-and-sets)
	for i, item := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			if item.Name.ValueString() != itemState.Name.ValueString() {
				continue
			}
			found = true
		}
		if !found {
			plan.Items[i].Id = types.StringUnknown()
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: ModifyPlan finished successfully", plan.Id.ValueString()))

	diags = resp.Plan.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}
// End of section. //template:end modifyPlan
