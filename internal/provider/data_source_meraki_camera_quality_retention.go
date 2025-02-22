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
	_ datasource.DataSource              = &CameraQualityRetentionDataSource{}
	_ datasource.DataSourceWithConfigure = &CameraQualityRetentionDataSource{}
)

func NewCameraQualityRetentionDataSource() datasource.DataSource {
	return &CameraQualityRetentionDataSource{}
}

type CameraQualityRetentionDataSource struct {
	client *meraki.Client
}

func (d *CameraQualityRetentionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_camera_quality_retention"
}

func (d *CameraQualityRetentionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Camera Quality Retention` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: "Device serial",
				Required:            true,
			},
			"audio_recording_enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating if audio recording is enabled(true) or disabled(false) on the camera",
				Computed:            true,
			},
			"motion_based_retention_enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating if motion-based retention is enabled(true) or disabled(false) on the camera.",
				Computed:            true,
			},
			"motion_detector_version": schema.Int64Attribute{
				MarkdownDescription: "The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.",
				Computed:            true,
			},
			"profile_id": schema.StringAttribute{
				MarkdownDescription: "The ID of a quality and retention profile to assign to the camera. The profile`s settings will override all of the per-camera quality and retention settings. If the value of this parameter is null, any existing profile will be unassigned from the camera.",
				Computed:            true,
			},
			"quality": schema.StringAttribute{
				MarkdownDescription: "Quality of the camera. Can be one of `Standard`, `High` or `Enhanced`. Not all qualities are supported by every camera model.",
				Computed:            true,
			},
			"resolution": schema.StringAttribute{
				MarkdownDescription: "Resolution of the camera. Can be one of `1280x720`, `1920x1080`, `1080x1080`, `2112x2112`, `2880x2880`, `2688x1512` or `3840x2160`.Not all resolutions are supported by every camera model.",
				Computed:            true,
			},
			"restricted_bandwidth_mode_enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating if restricted bandwidth is enabled(true) or disabled(false) on the camera. This setting does not apply to MV2 cameras.",
				Computed:            true,
			},
		},
	}
}

func (d *CameraQualityRetentionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *CameraQualityRetentionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CameraQualityRetention

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
	config.Id = config.Serial

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
