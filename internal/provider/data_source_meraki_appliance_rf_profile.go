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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ApplianceRFProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceRFProfileDataSource{}
)

func NewApplianceRFProfileDataSource() datasource.DataSource {
	return &ApplianceRFProfileDataSource{}
}

type ApplianceRFProfileDataSource struct {
	client *meraki.Client
}

func (d *ApplianceRFProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_rf_profile"
}

func (d *ApplianceRFProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Appliance RF Profile` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the new profile. Must be unique. This param is required on creation.",
				Optional:            true,
				Computed:            true,
			},
			"five_ghz_settings_ax_enabled": schema.BoolAttribute{
				MarkdownDescription: "Determines whether ax radio on 5Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.",
				Computed:            true,
			},
			"five_ghz_settings_min_bitrate": schema.Int64Attribute{
				MarkdownDescription: "Sets min bitrate (Mbps) of 5Ghz band. Can be one of `6`, `9`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 12.",
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
			"per_ssid_settings_2_band_operation_mode": schema.StringAttribute{
				MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
				Computed:            true,
			},
			"per_ssid_settings_2_band_steering_enabled": schema.BoolAttribute{
				MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
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
			"per_ssid_settings_4_band_operation_mode": schema.StringAttribute{
				MarkdownDescription: "Choice between `dual`, `2.4ghz`, `5ghz`, `6ghz` or `multi`.",
				Computed:            true,
			},
			"per_ssid_settings_4_band_steering_enabled": schema.BoolAttribute{
				MarkdownDescription: "Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.",
				Computed:            true,
			},
			"two_four_ghz_settings_ax_enabled": schema.BoolAttribute{
				MarkdownDescription: "Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.",
				Computed:            true,
			},
			"two_four_ghz_settings_min_bitrate": schema.Float64Attribute{
				MarkdownDescription: "Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of `1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`. Defaults to 11.",
				Computed:            true,
			},
		},
	}
}
func (d *ApplianceRFProfileDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *ApplianceRFProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceRFProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplianceRFProfile

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res meraki.Res
	var err error
	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		if len(res.Array()) > 0 {
			res.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("name").String() {
					config.Id = types.StringValue(v.Get("id").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					res = meraki.Res{Result: v}
					return false
				}
				return true
			})
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}

	if !res.Exists() {
		res, err = d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
