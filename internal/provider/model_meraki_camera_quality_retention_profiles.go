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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type CameraQualityRetentionProfiles struct {
	NetworkId types.String                          `tfsdk:"network_id"`
	Items     []CameraQualityRetentionProfilesItems `tfsdk:"items"`
}

type CameraQualityRetentionProfilesItems struct {
	Id                                  types.String `tfsdk:"id"`
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
	VideoSettingsMv93Quality            types.String `tfsdk:"video_settings_mv93_quality"`
	VideoSettingsMv93Resolution         types.String `tfsdk:"video_settings_mv93_resolution"`
	VideoSettingsMv93MQuality           types.String `tfsdk:"video_settings_mv93_m_quality"`
	VideoSettingsMv93MResolution        types.String `tfsdk:"video_settings_mv93_m_resolution"`
	VideoSettingsMv93XQuality           types.String `tfsdk:"video_settings_mv93_x_quality"`
	VideoSettingsMv93XResolution        types.String `tfsdk:"video_settings_mv93_x_resolution"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CameraQualityRetentionProfiles) getPath() string {
	return fmt.Sprintf("/networks/%v/camera/qualityRetentionProfiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CameraQualityRetentionProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]CameraQualityRetentionProfilesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := CameraQualityRetentionProfilesItems{}
		data.Id = types.StringValue(res.Get("id").String())
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
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
