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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSourceCameraQualityRetention - used only by the data source, always including every
// attribute (unlike model_resource.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model_resource.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model_resource.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceCameraQualityRetention struct {
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

func (data DataSourceCameraQualityRetention) getPath() string {
	return fmt.Sprintf("/devices/%v/camera/qualityAndRetention", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceCameraQualityRetention) fromBody(ctx context.Context, res meraki.Res) {
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
