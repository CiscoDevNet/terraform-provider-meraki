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

type CameraQualityRetention struct {
	Id                             types.String `tfsdk:"id"`
	Serial                         types.String `tfsdk:"serial"`
	AudioRecordingEnabled          types.Bool   `tfsdk:"audio_recording_enabled"`
	MotionBasedRetentionEnabled    types.Bool   `tfsdk:"motion_based_retention_enabled"`
	MotionDetectorVersion          types.Int64  `tfsdk:"motion_detector_version"`
	ProfileId                      types.String `tfsdk:"profile_id"`
	Quality                        types.String `tfsdk:"quality"`
	Resolution                     types.String `tfsdk:"resolution"`
	RestrictedBandwidthModeEnabled types.Bool   `tfsdk:"restricted_bandwidth_mode_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CameraQualityRetention) getPath() string {
	return fmt.Sprintf("/devices/%v/camera/qualityAndRetention", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CameraQualityRetention) toBody(ctx context.Context, state CameraQualityRetention) string {
	body := ""
	if !data.AudioRecordingEnabled.IsNull() {
		body, _ = sjson.Set(body, "audioRecordingEnabled", data.AudioRecordingEnabled.ValueBool())
	}
	if !data.MotionBasedRetentionEnabled.IsNull() {
		body, _ = sjson.Set(body, "motionBasedRetentionEnabled", data.MotionBasedRetentionEnabled.ValueBool())
	}
	if !data.MotionDetectorVersion.IsNull() {
		body, _ = sjson.Set(body, "motionDetectorVersion", data.MotionDetectorVersion.ValueInt64())
	}
	if !data.ProfileId.IsNull() {
		body, _ = sjson.Set(body, "profileId", data.ProfileId.ValueString())
	}
	if !data.Quality.IsNull() {
		body, _ = sjson.Set(body, "quality", data.Quality.ValueString())
	}
	if !data.Resolution.IsNull() {
		body, _ = sjson.Set(body, "resolution", data.Resolution.ValueString())
	}
	if !data.RestrictedBandwidthModeEnabled.IsNull() {
		body, _ = sjson.Set(body, "restrictedBandwidthModeEnabled", data.RestrictedBandwidthModeEnabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CameraQualityRetention) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("audioRecordingEnabled"); value.Exists() && value.Value() != nil {
		data.AudioRecordingEnabled = types.BoolValue(value.Bool())
	} else {
		data.AudioRecordingEnabled = types.BoolNull()
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
	if value := res.Get("profileId"); value.Exists() && value.Value() != nil {
		data.ProfileId = types.StringValue(value.String())
	} else {
		data.ProfileId = types.StringNull()
	}
	if value := res.Get("quality"); value.Exists() && value.Value() != nil {
		data.Quality = types.StringValue(value.String())
	} else {
		data.Quality = types.StringNull()
	}
	if value := res.Get("resolution"); value.Exists() && value.Value() != nil {
		data.Resolution = types.StringValue(value.String())
	} else {
		data.Resolution = types.StringNull()
	}
	if value := res.Get("restrictedBandwidthModeEnabled"); value.Exists() && value.Value() != nil {
		data.RestrictedBandwidthModeEnabled = types.BoolValue(value.Bool())
	} else {
		data.RestrictedBandwidthModeEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *CameraQualityRetention) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("audioRecordingEnabled"); value.Exists() && !data.AudioRecordingEnabled.IsNull() {
		data.AudioRecordingEnabled = types.BoolValue(value.Bool())
	} else {
		data.AudioRecordingEnabled = types.BoolNull()
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
	if value := res.Get("profileId"); value.Exists() && !data.ProfileId.IsNull() {
		data.ProfileId = types.StringValue(value.String())
	} else {
		data.ProfileId = types.StringNull()
	}
	if value := res.Get("quality"); value.Exists() && !data.Quality.IsNull() {
		data.Quality = types.StringValue(value.String())
	} else {
		data.Quality = types.StringNull()
	}
	if value := res.Get("resolution"); value.Exists() && !data.Resolution.IsNull() {
		data.Resolution = types.StringValue(value.String())
	} else {
		data.Resolution = types.StringNull()
	}
	if value := res.Get("restrictedBandwidthModeEnabled"); value.Exists() && !data.RestrictedBandwidthModeEnabled.IsNull() {
		data.RestrictedBandwidthModeEnabled = types.BoolValue(value.Bool())
	} else {
		data.RestrictedBandwidthModeEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CameraQualityRetention) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data CameraQualityRetention) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
