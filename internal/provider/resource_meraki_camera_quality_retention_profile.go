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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &CameraQualityRetentionProfileResource{}
	_ resource.ResourceWithImportState = &CameraQualityRetentionProfileResource{}
)

func NewCameraQualityRetentionProfileResource() resource.Resource {
	return &CameraQualityRetentionProfileResource{}
}

type CameraQualityRetentionProfileResource struct {
	client *meraki.Client
}

func (r *CameraQualityRetentionProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_camera_quality_retention_profile"
}

func (r *CameraQualityRetentionProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Camera Quality Retention Profile` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"audio_recording_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether or not to record audio. Can be either true or false. Defaults to false.").String,
				Optional:            true,
			},
			"cloud_archive_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false.").String,
				Optional:            true,
			},
			"max_retention_days": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The maximum number of days for which the data will be stored, or `null` to keep data until storage space runs out. If the former, it can be in the range of one to ninety days.").String,
				Optional:            true,
			},
			"motion_based_retention_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.").String,
				Optional:            true,
			},
			"motion_detector_version": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.").String,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the new profile. Must be unique. This parameter is required.").String,
				Required:            true,
			},
			"restricted_bandwidth_mode_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.").String,
				Optional:            true,
			},
			"schedule_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Schedule for which this camera will record video, or `null` to always record.").String,
				Optional:            true,
			},
			"smart_retention_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Boolean indicating if Smart Retention is enabled(true) or disabled(false).").String,
				Optional:            true,
			},
			"video_settings_mv12_mv22_mv72_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv12_mv22_mv72_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1280x720` or `1920x1080`.").AddStringEnumDescription("1280x720", "1920x1080").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1280x720", "1920x1080"),
				},
			},
			"video_settings_mv12_we_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv12_we_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1280x720` or `1920x1080`.").AddStringEnumDescription("1280x720", "1920x1080").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1280x720", "1920x1080"),
				},
			},
			"video_settings_mv13_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv13_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv13_m_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv13_m_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv21_mv71_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv21_mv71_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1280x720`.").AddStringEnumDescription("1280x720").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1280x720"),
				},
			},
			"video_settings_mv22_x_mv72_x_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv22_x_mv72_x_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1280x720`, `1920x1080` or `2688x1512`.").AddStringEnumDescription("1280x720", "1920x1080", "2688x1512").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1280x720", "1920x1080", "2688x1512"),
				},
			},
			"video_settings_mv23_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv23_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv23_m_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv23_m_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv23_x_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv23_x_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv32_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv32_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1080x1080` or `2112x2112`.").AddStringEnumDescription("1080x1080", "2112x2112").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1080x1080", "2112x2112"),
				},
			},
			"video_settings_mv33_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv33_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1080x1080`, `2112x2112` or `2880x2880`.").AddStringEnumDescription("1080x1080", "2112x2112", "2880x2880").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1080x1080", "2112x2112", "2880x2880"),
				},
			},
			"video_settings_mv33_m_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv33_m_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1080x1080`, `2112x2112` or `2880x2880`.").AddStringEnumDescription("1080x1080", "2112x2112", "2880x2880").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1080x1080", "2112x2112", "2880x2880"),
				},
			},
			"video_settings_mv52_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv52_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1280x720`, `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1280x720", "1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1280x720", "1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv53_x_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv53_x_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv63_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv63_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv63_m_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv63_m_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv63_x_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv63_x_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv73_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv73_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv73_m_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv73_m_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv73_x_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv73_x_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1920x1080`, `2688x1512` or `3840x2160`.").AddStringEnumDescription("1920x1080", "2688x1512", "3840x2160").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1920x1080", "2688x1512", "3840x2160"),
				},
			},
			"video_settings_mv93_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv93_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1080x1080`, `2112x2112` or `2880x2880`.").AddStringEnumDescription("1080x1080", "2112x2112", "2880x2880").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1080x1080", "2112x2112", "2880x2880"),
				},
			},
			"video_settings_mv93_m_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv93_m_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1080x1080`, `2112x2112` or `2880x2880`.").AddStringEnumDescription("1080x1080", "2112x2112", "2880x2880").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1080x1080", "2112x2112", "2880x2880"),
				},
			},
			"video_settings_mv93_x_quality": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Quality of the camera. Can be one of `Standard`, `Enhanced` or `High`.").AddStringEnumDescription("Enhanced", "High", "Standard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Enhanced", "High", "Standard"),
				},
			},
			"video_settings_mv93_x_resolution": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Resolution of the camera. Can be one of `1080x1080`, `2112x2112` or `2880x2880`.").AddStringEnumDescription("1080x1080", "2112x2112", "2880x2880").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1080x1080", "2112x2112", "2880x2880"),
				},
			},
		},
	}
}

func (r *CameraQualityRetentionProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *CameraQualityRetentionProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CameraQualityRetentionProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, CameraQualityRetentionProfile{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *CameraQualityRetentionProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CameraQualityRetentionProfile

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
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
		state.fromBody(ctx, res)
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

func (r *CameraQualityRetentionProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CameraQualityRetentionProfile

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

	body := plan.toBody(ctx, state)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *CameraQualityRetentionProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CameraQualityRetentionProfile

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *CameraQualityRetentionProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
