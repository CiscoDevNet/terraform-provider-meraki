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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type CameraQualityRetentionProfile struct {
	Id                                  types.String `tfsdk:"id"`
	NetworkId                           types.String `tfsdk:"network_id"`
	AudioRecordingEnabled               types.Bool   `tfsdk:"audio_recording_enabled"`
	CloudArchiveEnabled                 types.Bool   `tfsdk:"cloud_archive_enabled"`
	MaxRetentionDays                    types.Int64  `tfsdk:"max_retention_days"`
	MotionBasedRetentionEnabled         types.Bool   `tfsdk:"motion_based_retention_enabled"`
	MotionDetectorVersion               types.Int64  `tfsdk:"motion_detector_version"`
	Name                                types.String `tfsdk:"name"`
	RestrictedBandwidthModeEnabled      types.Bool   `tfsdk:"restricted_bandwidth_mode_enabled"`
	ScheduleId                          types.String `tfsdk:"schedule_id"`
	SmartRetentionEnabled               types.Bool   `tfsdk:"smart_retention_enabled"`
	VideoSettingsMv12Mv22Mv72Quality    types.String `tfsdk:"video_settings_mv12_mv22_mv72_quality"`
	VideoSettingsMv12Mv22Mv72Resolution types.String `tfsdk:"video_settings_mv12_mv22_mv72_resolution"`
	VideoSettingsMv12WeQuality          types.String `tfsdk:"video_settings_mv12_we_quality"`
	VideoSettingsMv12WeResolution       types.String `tfsdk:"video_settings_mv12_we_resolution"`
	VideoSettingsMv13Quality            types.String `tfsdk:"video_settings_mv13_quality"`
	VideoSettingsMv13Resolution         types.String `tfsdk:"video_settings_mv13_resolution"`
	VideoSettingsMv13MQuality           types.String `tfsdk:"video_settings_mv13_m_quality"`
	VideoSettingsMv13MResolution        types.String `tfsdk:"video_settings_mv13_m_resolution"`
	VideoSettingsMv21Mv71Quality        types.String `tfsdk:"video_settings_mv21_mv71_quality"`
	VideoSettingsMv21Mv71Resolution     types.String `tfsdk:"video_settings_mv21_mv71_resolution"`
	VideoSettingsMv22XMv72XQuality      types.String `tfsdk:"video_settings_mv22_x_mv72_x_quality"`
	VideoSettingsMv22XMv72XResolution   types.String `tfsdk:"video_settings_mv22_x_mv72_x_resolution"`
	VideoSettingsMv23Quality            types.String `tfsdk:"video_settings_mv23_quality"`
	VideoSettingsMv23Resolution         types.String `tfsdk:"video_settings_mv23_resolution"`
	VideoSettingsMv23MQuality           types.String `tfsdk:"video_settings_mv23_m_quality"`
	VideoSettingsMv23MResolution        types.String `tfsdk:"video_settings_mv23_m_resolution"`
	VideoSettingsMv23XQuality           types.String `tfsdk:"video_settings_mv23_x_quality"`
	VideoSettingsMv23XResolution        types.String `tfsdk:"video_settings_mv23_x_resolution"`
	VideoSettingsMv32Quality            types.String `tfsdk:"video_settings_mv32_quality"`
	VideoSettingsMv32Resolution         types.String `tfsdk:"video_settings_mv32_resolution"`
	VideoSettingsMv33Quality            types.String `tfsdk:"video_settings_mv33_quality"`
	VideoSettingsMv33Resolution         types.String `tfsdk:"video_settings_mv33_resolution"`
	VideoSettingsMv33MQuality           types.String `tfsdk:"video_settings_mv33_m_quality"`
	VideoSettingsMv33MResolution        types.String `tfsdk:"video_settings_mv33_m_resolution"`
	VideoSettingsMv52Quality            types.String `tfsdk:"video_settings_mv52_quality"`
	VideoSettingsMv52Resolution         types.String `tfsdk:"video_settings_mv52_resolution"`
	VideoSettingsMv53XQuality           types.String `tfsdk:"video_settings_mv53_x_quality"`
	VideoSettingsMv53XResolution        types.String `tfsdk:"video_settings_mv53_x_resolution"`
	VideoSettingsMv63Quality            types.String `tfsdk:"video_settings_mv63_quality"`
	VideoSettingsMv63Resolution         types.String `tfsdk:"video_settings_mv63_resolution"`
	VideoSettingsMv63MQuality           types.String `tfsdk:"video_settings_mv63_m_quality"`
	VideoSettingsMv63MResolution        types.String `tfsdk:"video_settings_mv63_m_resolution"`
	VideoSettingsMv63XQuality           types.String `tfsdk:"video_settings_mv63_x_quality"`
	VideoSettingsMv63XResolution        types.String `tfsdk:"video_settings_mv63_x_resolution"`
	VideoSettingsMv73Quality            types.String `tfsdk:"video_settings_mv73_quality"`
	VideoSettingsMv73Resolution         types.String `tfsdk:"video_settings_mv73_resolution"`
	VideoSettingsMv73MQuality           types.String `tfsdk:"video_settings_mv73_m_quality"`
	VideoSettingsMv73MResolution        types.String `tfsdk:"video_settings_mv73_m_resolution"`
	VideoSettingsMv73XQuality           types.String `tfsdk:"video_settings_mv73_x_quality"`
	VideoSettingsMv73XResolution        types.String `tfsdk:"video_settings_mv73_x_resolution"`
	VideoSettingsMv84XQuality           types.String `tfsdk:"video_settings_mv84_x_quality"`
	VideoSettingsMv84XResolution        types.String `tfsdk:"video_settings_mv84_x_resolution"`
	VideoSettingsMv93Quality            types.String `tfsdk:"video_settings_mv93_quality"`
	VideoSettingsMv93Resolution         types.String `tfsdk:"video_settings_mv93_resolution"`
	VideoSettingsMv93MQuality           types.String `tfsdk:"video_settings_mv93_m_quality"`
	VideoSettingsMv93MResolution        types.String `tfsdk:"video_settings_mv93_m_resolution"`
	VideoSettingsMv93XQuality           types.String `tfsdk:"video_settings_mv93_x_quality"`
	VideoSettingsMv93XResolution        types.String `tfsdk:"video_settings_mv93_x_resolution"`
}

type CameraQualityRetentionProfileIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
	Id        types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CameraQualityRetentionProfile) getPath() string {
	return fmt.Sprintf("/networks/%v/camera/qualityRetentionProfiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CameraQualityRetentionProfile) toBody(ctx context.Context, state CameraQualityRetentionProfile) string {
	body := ""
	if !data.AudioRecordingEnabled.IsNull() {
		body, _ = sjson.Set(body, "audioRecordingEnabled", data.AudioRecordingEnabled.ValueBool())
	}
	if !data.CloudArchiveEnabled.IsNull() {
		body, _ = sjson.Set(body, "cloudArchiveEnabled", data.CloudArchiveEnabled.ValueBool())
	}
	if !data.MaxRetentionDays.IsNull() {
		body, _ = sjson.Set(body, "maxRetentionDays", data.MaxRetentionDays.ValueInt64())
	}
	if !data.MotionBasedRetentionEnabled.IsNull() {
		body, _ = sjson.Set(body, "motionBasedRetentionEnabled", data.MotionBasedRetentionEnabled.ValueBool())
	}
	if !data.MotionDetectorVersion.IsNull() {
		body, _ = sjson.Set(body, "motionDetectorVersion", data.MotionDetectorVersion.ValueInt64())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.RestrictedBandwidthModeEnabled.IsNull() {
		body, _ = sjson.Set(body, "restrictedBandwidthModeEnabled", data.RestrictedBandwidthModeEnabled.ValueBool())
	}
	if !data.ScheduleId.IsNull() {
		body, _ = sjson.Set(body, "scheduleId", data.ScheduleId.ValueString())
	}
	if !data.SmartRetentionEnabled.IsNull() {
		body, _ = sjson.Set(body, "smartRetention.enabled", data.SmartRetentionEnabled.ValueBool())
	}
	if !data.VideoSettingsMv12Mv22Mv72Quality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV12/MV22/MV72.quality", data.VideoSettingsMv12Mv22Mv72Quality.ValueString())
	}
	if !data.VideoSettingsMv12Mv22Mv72Resolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV12/MV22/MV72.resolution", data.VideoSettingsMv12Mv22Mv72Resolution.ValueString())
	}
	if !data.VideoSettingsMv12WeQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV12WE.quality", data.VideoSettingsMv12WeQuality.ValueString())
	}
	if !data.VideoSettingsMv12WeResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV12WE.resolution", data.VideoSettingsMv12WeResolution.ValueString())
	}
	if !data.VideoSettingsMv13Quality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV13.quality", data.VideoSettingsMv13Quality.ValueString())
	}
	if !data.VideoSettingsMv13Resolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV13.resolution", data.VideoSettingsMv13Resolution.ValueString())
	}
	if !data.VideoSettingsMv13MQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV13M.quality", data.VideoSettingsMv13MQuality.ValueString())
	}
	if !data.VideoSettingsMv13MResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV13M.resolution", data.VideoSettingsMv13MResolution.ValueString())
	}
	if !data.VideoSettingsMv21Mv71Quality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV21/MV71.quality", data.VideoSettingsMv21Mv71Quality.ValueString())
	}
	if !data.VideoSettingsMv21Mv71Resolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV21/MV71.resolution", data.VideoSettingsMv21Mv71Resolution.ValueString())
	}
	if !data.VideoSettingsMv22XMv72XQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV22X/MV72X.quality", data.VideoSettingsMv22XMv72XQuality.ValueString())
	}
	if !data.VideoSettingsMv22XMv72XResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV22X/MV72X.resolution", data.VideoSettingsMv22XMv72XResolution.ValueString())
	}
	if !data.VideoSettingsMv23Quality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV23.quality", data.VideoSettingsMv23Quality.ValueString())
	}
	if !data.VideoSettingsMv23Resolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV23.resolution", data.VideoSettingsMv23Resolution.ValueString())
	}
	if !data.VideoSettingsMv23MQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV23M.quality", data.VideoSettingsMv23MQuality.ValueString())
	}
	if !data.VideoSettingsMv23MResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV23M.resolution", data.VideoSettingsMv23MResolution.ValueString())
	}
	if !data.VideoSettingsMv23XQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV23X.quality", data.VideoSettingsMv23XQuality.ValueString())
	}
	if !data.VideoSettingsMv23XResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV23X.resolution", data.VideoSettingsMv23XResolution.ValueString())
	}
	if !data.VideoSettingsMv32Quality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV32.quality", data.VideoSettingsMv32Quality.ValueString())
	}
	if !data.VideoSettingsMv32Resolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV32.resolution", data.VideoSettingsMv32Resolution.ValueString())
	}
	if !data.VideoSettingsMv33Quality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV33.quality", data.VideoSettingsMv33Quality.ValueString())
	}
	if !data.VideoSettingsMv33Resolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV33.resolution", data.VideoSettingsMv33Resolution.ValueString())
	}
	if !data.VideoSettingsMv33MQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV33M.quality", data.VideoSettingsMv33MQuality.ValueString())
	}
	if !data.VideoSettingsMv33MResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV33M.resolution", data.VideoSettingsMv33MResolution.ValueString())
	}
	if !data.VideoSettingsMv52Quality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV52.quality", data.VideoSettingsMv52Quality.ValueString())
	}
	if !data.VideoSettingsMv52Resolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV52.resolution", data.VideoSettingsMv52Resolution.ValueString())
	}
	if !data.VideoSettingsMv53XQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV53X.quality", data.VideoSettingsMv53XQuality.ValueString())
	}
	if !data.VideoSettingsMv53XResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV53X.resolution", data.VideoSettingsMv53XResolution.ValueString())
	}
	if !data.VideoSettingsMv63Quality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV63.quality", data.VideoSettingsMv63Quality.ValueString())
	}
	if !data.VideoSettingsMv63Resolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV63.resolution", data.VideoSettingsMv63Resolution.ValueString())
	}
	if !data.VideoSettingsMv63MQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV63M.quality", data.VideoSettingsMv63MQuality.ValueString())
	}
	if !data.VideoSettingsMv63MResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV63M.resolution", data.VideoSettingsMv63MResolution.ValueString())
	}
	if !data.VideoSettingsMv63XQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV63X.quality", data.VideoSettingsMv63XQuality.ValueString())
	}
	if !data.VideoSettingsMv63XResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV63X.resolution", data.VideoSettingsMv63XResolution.ValueString())
	}
	if !data.VideoSettingsMv73Quality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV73.quality", data.VideoSettingsMv73Quality.ValueString())
	}
	if !data.VideoSettingsMv73Resolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV73.resolution", data.VideoSettingsMv73Resolution.ValueString())
	}
	if !data.VideoSettingsMv73MQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV73M.quality", data.VideoSettingsMv73MQuality.ValueString())
	}
	if !data.VideoSettingsMv73MResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV73M.resolution", data.VideoSettingsMv73MResolution.ValueString())
	}
	if !data.VideoSettingsMv73XQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV73X.quality", data.VideoSettingsMv73XQuality.ValueString())
	}
	if !data.VideoSettingsMv73XResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV73X.resolution", data.VideoSettingsMv73XResolution.ValueString())
	}
	if !data.VideoSettingsMv84XQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV84X.quality", data.VideoSettingsMv84XQuality.ValueString())
	}
	if !data.VideoSettingsMv84XResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV84X.resolution", data.VideoSettingsMv84XResolution.ValueString())
	}
	if !data.VideoSettingsMv93Quality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV93.quality", data.VideoSettingsMv93Quality.ValueString())
	}
	if !data.VideoSettingsMv93Resolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV93.resolution", data.VideoSettingsMv93Resolution.ValueString())
	}
	if !data.VideoSettingsMv93MQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV93M.quality", data.VideoSettingsMv93MQuality.ValueString())
	}
	if !data.VideoSettingsMv93MResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV93M.resolution", data.VideoSettingsMv93MResolution.ValueString())
	}
	if !data.VideoSettingsMv93XQuality.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV93X.quality", data.VideoSettingsMv93XQuality.ValueString())
	}
	if !data.VideoSettingsMv93XResolution.IsNull() {
		body, _ = sjson.Set(body, "videoSettings.MV93X.resolution", data.VideoSettingsMv93XResolution.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CameraQualityRetentionProfile) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("audioRecordingEnabled"); value.Exists() && value.Value() != nil {
		data.AudioRecordingEnabled = types.BoolValue(value.Bool())
	} else {
		data.AudioRecordingEnabled = types.BoolNull()
	}
	if value := res.Get("cloudArchiveEnabled"); value.Exists() && value.Value() != nil {
		data.CloudArchiveEnabled = types.BoolValue(value.Bool())
	} else {
		data.CloudArchiveEnabled = types.BoolNull()
	}
	if value := res.Get("maxRetentionDays"); value.Exists() && value.Value() != nil {
		data.MaxRetentionDays = types.Int64Value(value.Int())
	} else {
		data.MaxRetentionDays = types.Int64Null()
	}
	if value := res.Get("motionBasedRetentionEnabled"); value.Exists() && value.Value() != nil {
		data.MotionBasedRetentionEnabled = types.BoolValue(value.Bool())
	} else {
		data.MotionBasedRetentionEnabled = types.BoolNull()
	}
	if value := res.Get("motionDetectorVersion"); value.Exists() && value.Value() != nil {
		data.MotionDetectorVersion = types.Int64Value(value.Int())
	} else {
		data.MotionDetectorVersion = types.Int64Null()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("restrictedBandwidthModeEnabled"); value.Exists() && value.Value() != nil {
		data.RestrictedBandwidthModeEnabled = types.BoolValue(value.Bool())
	} else {
		data.RestrictedBandwidthModeEnabled = types.BoolNull()
	}
	if value := res.Get("scheduleId"); value.Exists() && value.Value() != nil {
		data.ScheduleId = types.StringValue(value.String())
	} else {
		data.ScheduleId = types.StringNull()
	}
	if value := res.Get("smartRetention.enabled"); value.Exists() && value.Value() != nil {
		data.SmartRetentionEnabled = types.BoolValue(value.Bool())
	} else {
		data.SmartRetentionEnabled = types.BoolNull()
	}
	if value := res.Get("videoSettings.MV12/MV22/MV72.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv12Mv22Mv72Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv12Mv22Mv72Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV12/MV22/MV72.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv12Mv22Mv72Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv12Mv22Mv72Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV12WE.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv12WeQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv12WeQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV12WE.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv12WeResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv12WeResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV13.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv13Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv13Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV13.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv13Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv13Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV13M.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv13MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv13MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV13M.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv13MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv13MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV21/MV71.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv21Mv71Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv21Mv71Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV21/MV71.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv21Mv71Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv21Mv71Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV22X/MV72X.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv22XMv72XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv22XMv72XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV22X/MV72X.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv22XMv72XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv22XMv72XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv23Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv23Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23M.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv23MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23M.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv23MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23X.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv23XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23X.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv23XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV32.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv32Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv32Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV32.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv32Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv32Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV33.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv33Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv33Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV33.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv33Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv33Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV33M.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv33MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv33MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV33M.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv33MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv33MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV52.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv52Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv52Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV52.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv52Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv52Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV53X.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv53XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv53XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV53X.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv53XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv53XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv63Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv63Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63M.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv63MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63M.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv63MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63X.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv63XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63X.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv63XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv73Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv73Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73M.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv73MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73M.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv73MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73X.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv73XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73X.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv73XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV84X.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv84XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv84XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV84X.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv84XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv84XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv93Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv93Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93M.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv93MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93M.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv93MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93X.quality"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv93XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93X.resolution"); value.Exists() && value.Value() != nil {
		data.VideoSettingsMv93XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93XResolution = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *CameraQualityRetentionProfile) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("audioRecordingEnabled"); value.Exists() && !data.AudioRecordingEnabled.IsNull() {
		data.AudioRecordingEnabled = types.BoolValue(value.Bool())
	} else {
		data.AudioRecordingEnabled = types.BoolNull()
	}
	if value := res.Get("cloudArchiveEnabled"); value.Exists() && !data.CloudArchiveEnabled.IsNull() {
		data.CloudArchiveEnabled = types.BoolValue(value.Bool())
	} else {
		data.CloudArchiveEnabled = types.BoolNull()
	}
	if value := res.Get("maxRetentionDays"); value.Exists() && !data.MaxRetentionDays.IsNull() {
		data.MaxRetentionDays = types.Int64Value(value.Int())
	} else {
		data.MaxRetentionDays = types.Int64Null()
	}
	if value := res.Get("motionBasedRetentionEnabled"); value.Exists() && !data.MotionBasedRetentionEnabled.IsNull() {
		data.MotionBasedRetentionEnabled = types.BoolValue(value.Bool())
	} else {
		data.MotionBasedRetentionEnabled = types.BoolNull()
	}
	if value := res.Get("motionDetectorVersion"); value.Exists() && !data.MotionDetectorVersion.IsNull() {
		data.MotionDetectorVersion = types.Int64Value(value.Int())
	} else {
		data.MotionDetectorVersion = types.Int64Null()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("restrictedBandwidthModeEnabled"); value.Exists() && !data.RestrictedBandwidthModeEnabled.IsNull() {
		data.RestrictedBandwidthModeEnabled = types.BoolValue(value.Bool())
	} else {
		data.RestrictedBandwidthModeEnabled = types.BoolNull()
	}
	if value := res.Get("scheduleId"); value.Exists() && !data.ScheduleId.IsNull() {
		data.ScheduleId = types.StringValue(value.String())
	} else {
		data.ScheduleId = types.StringNull()
	}
	if value := res.Get("smartRetention.enabled"); value.Exists() && !data.SmartRetentionEnabled.IsNull() {
		data.SmartRetentionEnabled = types.BoolValue(value.Bool())
	} else {
		data.SmartRetentionEnabled = types.BoolNull()
	}
	if value := res.Get("videoSettings.MV12/MV22/MV72.quality"); value.Exists() && !data.VideoSettingsMv12Mv22Mv72Quality.IsNull() {
		data.VideoSettingsMv12Mv22Mv72Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv12Mv22Mv72Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV12/MV22/MV72.resolution"); value.Exists() && !data.VideoSettingsMv12Mv22Mv72Resolution.IsNull() {
		data.VideoSettingsMv12Mv22Mv72Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv12Mv22Mv72Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV12WE.quality"); value.Exists() && !data.VideoSettingsMv12WeQuality.IsNull() {
		data.VideoSettingsMv12WeQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv12WeQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV12WE.resolution"); value.Exists() && !data.VideoSettingsMv12WeResolution.IsNull() {
		data.VideoSettingsMv12WeResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv12WeResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV13.quality"); value.Exists() && !data.VideoSettingsMv13Quality.IsNull() {
		data.VideoSettingsMv13Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv13Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV13.resolution"); value.Exists() && !data.VideoSettingsMv13Resolution.IsNull() {
		data.VideoSettingsMv13Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv13Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV13M.quality"); value.Exists() && !data.VideoSettingsMv13MQuality.IsNull() {
		data.VideoSettingsMv13MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv13MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV13M.resolution"); value.Exists() && !data.VideoSettingsMv13MResolution.IsNull() {
		data.VideoSettingsMv13MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv13MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV21/MV71.quality"); value.Exists() && !data.VideoSettingsMv21Mv71Quality.IsNull() {
		data.VideoSettingsMv21Mv71Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv21Mv71Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV21/MV71.resolution"); value.Exists() && !data.VideoSettingsMv21Mv71Resolution.IsNull() {
		data.VideoSettingsMv21Mv71Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv21Mv71Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV22X/MV72X.quality"); value.Exists() && !data.VideoSettingsMv22XMv72XQuality.IsNull() {
		data.VideoSettingsMv22XMv72XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv22XMv72XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV22X/MV72X.resolution"); value.Exists() && !data.VideoSettingsMv22XMv72XResolution.IsNull() {
		data.VideoSettingsMv22XMv72XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv22XMv72XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23.quality"); value.Exists() && !data.VideoSettingsMv23Quality.IsNull() {
		data.VideoSettingsMv23Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23.resolution"); value.Exists() && !data.VideoSettingsMv23Resolution.IsNull() {
		data.VideoSettingsMv23Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23M.quality"); value.Exists() && !data.VideoSettingsMv23MQuality.IsNull() {
		data.VideoSettingsMv23MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23M.resolution"); value.Exists() && !data.VideoSettingsMv23MResolution.IsNull() {
		data.VideoSettingsMv23MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23X.quality"); value.Exists() && !data.VideoSettingsMv23XQuality.IsNull() {
		data.VideoSettingsMv23XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV23X.resolution"); value.Exists() && !data.VideoSettingsMv23XResolution.IsNull() {
		data.VideoSettingsMv23XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv23XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV32.quality"); value.Exists() && !data.VideoSettingsMv32Quality.IsNull() {
		data.VideoSettingsMv32Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv32Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV32.resolution"); value.Exists() && !data.VideoSettingsMv32Resolution.IsNull() {
		data.VideoSettingsMv32Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv32Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV33.quality"); value.Exists() && !data.VideoSettingsMv33Quality.IsNull() {
		data.VideoSettingsMv33Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv33Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV33.resolution"); value.Exists() && !data.VideoSettingsMv33Resolution.IsNull() {
		data.VideoSettingsMv33Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv33Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV33M.quality"); value.Exists() && !data.VideoSettingsMv33MQuality.IsNull() {
		data.VideoSettingsMv33MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv33MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV33M.resolution"); value.Exists() && !data.VideoSettingsMv33MResolution.IsNull() {
		data.VideoSettingsMv33MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv33MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV52.quality"); value.Exists() && !data.VideoSettingsMv52Quality.IsNull() {
		data.VideoSettingsMv52Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv52Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV52.resolution"); value.Exists() && !data.VideoSettingsMv52Resolution.IsNull() {
		data.VideoSettingsMv52Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv52Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV53X.quality"); value.Exists() && !data.VideoSettingsMv53XQuality.IsNull() {
		data.VideoSettingsMv53XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv53XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV53X.resolution"); value.Exists() && !data.VideoSettingsMv53XResolution.IsNull() {
		data.VideoSettingsMv53XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv53XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63.quality"); value.Exists() && !data.VideoSettingsMv63Quality.IsNull() {
		data.VideoSettingsMv63Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63.resolution"); value.Exists() && !data.VideoSettingsMv63Resolution.IsNull() {
		data.VideoSettingsMv63Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63M.quality"); value.Exists() && !data.VideoSettingsMv63MQuality.IsNull() {
		data.VideoSettingsMv63MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63M.resolution"); value.Exists() && !data.VideoSettingsMv63MResolution.IsNull() {
		data.VideoSettingsMv63MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63X.quality"); value.Exists() && !data.VideoSettingsMv63XQuality.IsNull() {
		data.VideoSettingsMv63XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV63X.resolution"); value.Exists() && !data.VideoSettingsMv63XResolution.IsNull() {
		data.VideoSettingsMv63XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv63XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73.quality"); value.Exists() && !data.VideoSettingsMv73Quality.IsNull() {
		data.VideoSettingsMv73Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73.resolution"); value.Exists() && !data.VideoSettingsMv73Resolution.IsNull() {
		data.VideoSettingsMv73Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73M.quality"); value.Exists() && !data.VideoSettingsMv73MQuality.IsNull() {
		data.VideoSettingsMv73MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73M.resolution"); value.Exists() && !data.VideoSettingsMv73MResolution.IsNull() {
		data.VideoSettingsMv73MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73X.quality"); value.Exists() && !data.VideoSettingsMv73XQuality.IsNull() {
		data.VideoSettingsMv73XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV73X.resolution"); value.Exists() && !data.VideoSettingsMv73XResolution.IsNull() {
		data.VideoSettingsMv73XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv73XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV84X.quality"); value.Exists() && !data.VideoSettingsMv84XQuality.IsNull() {
		data.VideoSettingsMv84XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv84XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV84X.resolution"); value.Exists() && !data.VideoSettingsMv84XResolution.IsNull() {
		data.VideoSettingsMv84XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv84XResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93.quality"); value.Exists() && !data.VideoSettingsMv93Quality.IsNull() {
		data.VideoSettingsMv93Quality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93Quality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93.resolution"); value.Exists() && !data.VideoSettingsMv93Resolution.IsNull() {
		data.VideoSettingsMv93Resolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93Resolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93M.quality"); value.Exists() && !data.VideoSettingsMv93MQuality.IsNull() {
		data.VideoSettingsMv93MQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93MQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93M.resolution"); value.Exists() && !data.VideoSettingsMv93MResolution.IsNull() {
		data.VideoSettingsMv93MResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93MResolution = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93X.quality"); value.Exists() && !data.VideoSettingsMv93XQuality.IsNull() {
		data.VideoSettingsMv93XQuality = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93XQuality = types.StringNull()
	}
	if value := res.Get("videoSettings.MV93X.resolution"); value.Exists() && !data.VideoSettingsMv93XResolution.IsNull() {
		data.VideoSettingsMv93XResolution = types.StringValue(value.String())
	} else {
		data.VideoSettingsMv93XResolution = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CameraQualityRetentionProfile) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *CameraQualityRetentionProfileIdentity) toIdentity(ctx context.Context, plan *CameraQualityRetentionProfile) {
	data.NetworkId = plan.NetworkId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *CameraQualityRetentionProfile) fromIdentity(ctx context.Context, identity *CameraQualityRetentionProfileIdentity) {
	data.NetworkId = identity.NetworkId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data CameraQualityRetentionProfile) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
