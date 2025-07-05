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
	_ datasource.DataSource              = &CameraQualityRetentionProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &CameraQualityRetentionProfilesDataSource{}
)

func NewCameraQualityRetentionProfilesDataSource() datasource.DataSource {
	return &CameraQualityRetentionProfilesDataSource{}
}

type CameraQualityRetentionProfilesDataSource struct {
	client *meraki.Client
}

func (d *CameraQualityRetentionProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_camera_quality_retention_profiles"
}

func (d *CameraQualityRetentionProfilesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Camera Quality Retention Profile` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The list of items",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the object",
							Computed:            true,
						},
						"audio_recording_enabled": schema.BoolAttribute{
							MarkdownDescription: "Whether or not to record audio. Can be either true or false. Defaults to false.",
							Computed:            true,
						},
						"cloud_archive_enabled": schema.BoolAttribute{
							MarkdownDescription: "Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false.",
							Computed:            true,
						},
						"max_retention_days": schema.Int64Attribute{
							MarkdownDescription: "The maximum number of days for which the data will be stored, or `null` to keep data until storage space runs out. If the former, it can be in the range of one to ninety days.",
							Computed:            true,
						},
						"motion_based_retention_enabled": schema.BoolAttribute{
							MarkdownDescription: "Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.",
							Computed:            true,
						},
						"motion_detector_version": schema.Int64Attribute{
							MarkdownDescription: "The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the new profile. Must be unique. This parameter is required.",
							Computed:            true,
						},
						"restricted_bandwidth_mode_enabled": schema.BoolAttribute{
							MarkdownDescription: "Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.",
							Computed:            true,
						},
						"schedule_id": schema.StringAttribute{
							MarkdownDescription: "Schedule for which this camera will record video, or `null` to always record.",
							Computed:            true,
						},
						"smart_retention_enabled": schema.BoolAttribute{
							MarkdownDescription: "Boolean indicating if Smart Retention is enabled(true) or disabled(false).",
							Computed:            true,
						},
						"video_settings_mv12_mv22_mv72_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv12_mv22_mv72_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1280x720` or `1920x1080`.",
							Computed:            true,
						},
						"video_settings_mv12_we_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv12_we_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1280x720` or `1920x1080`.",
							Computed:            true,
						},
						"video_settings_mv13_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv13_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv13_m_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv13_m_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv21_mv71_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv21_mv71_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1280x720`.",
							Computed:            true,
						},
						"video_settings_mv22_x_mv72_x_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv22_x_mv72_x_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1280x720`, `1920x1080` or `2688x1512`.",
							Computed:            true,
						},
						"video_settings_mv23_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv23_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv23_m_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv23_m_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv23_x_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv23_x_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv32_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv32_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1080x1080` or `2112x2112`.",
							Computed:            true,
						},
						"video_settings_mv33_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv33_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1080x1080`, `2112x2112` or `2880x2880`.",
							Computed:            true,
						},
						"video_settings_mv33_m_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv33_m_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1080x1080`, `2112x2112` or `2880x2880`.",
							Computed:            true,
						},
						"video_settings_mv52_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv52_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1280x720`, `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv53_x_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv53_x_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv63_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv63_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv63_m_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv63_m_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv63_x_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv63_x_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv73_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv73_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv73_m_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv73_m_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv73_x_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv73_x_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.",
							Computed:            true,
						},
						"video_settings_mv84_x_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard` or `Enhanced`.",
							Computed:            true,
						},
						"video_settings_mv84_x_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1440x1080` or `2560x1920`.",
							Computed:            true,
						},
						"video_settings_mv93_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv93_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1080x1080`, `2112x2112` or `2880x2880`.",
							Computed:            true,
						},
						"video_settings_mv93_m_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv93_m_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1080x1080`, `2112x2112` or `2880x2880`.",
							Computed:            true,
						},
						"video_settings_mv93_x_quality": schema.StringAttribute{
							MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.",
							Computed:            true,
						},
						"video_settings_mv93_x_resolution": schema.StringAttribute{
							MarkdownDescription: "Resolution of the camera. Can be one of `1080x1080`, `2112x2112` or `2880x2880`.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *CameraQualityRetentionProfilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *CameraQualityRetentionProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceCameraQualityRetentionProfiles

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "CameraQualityRetentionProfilesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "CameraQualityRetentionProfilesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
