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

type CameraSense struct {
	Id                    types.String `tfsdk:"id"`
	Serial                types.String `tfsdk:"serial"`
	DetectionModelId      types.String `tfsdk:"detection_model_id"`
	MqttBrokerId          types.String `tfsdk:"mqtt_broker_id"`
	SenseEnabled          types.Bool   `tfsdk:"sense_enabled"`
	AudioDetectionEnabled types.Bool   `tfsdk:"audio_detection_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CameraSense) getPath() string {
	return fmt.Sprintf("/devices/%v/camera/sense", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CameraSense) toBody(ctx context.Context, state CameraSense) string {
	body := ""
	if !data.DetectionModelId.IsNull() {
		body, _ = sjson.Set(body, "detectionModelId", data.DetectionModelId.ValueString())
	}
	if !data.MqttBrokerId.IsNull() {
		body, _ = sjson.Set(body, "mqttBrokerId", data.MqttBrokerId.ValueString())
	}
	if !data.SenseEnabled.IsNull() {
		body, _ = sjson.Set(body, "senseEnabled", data.SenseEnabled.ValueBool())
	}
	if !data.AudioDetectionEnabled.IsNull() {
		body, _ = sjson.Set(body, "audioDetection.enabled", data.AudioDetectionEnabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CameraSense) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("detectionModelId"); value.Exists() && value.Value() != nil {
		data.DetectionModelId = types.StringValue(value.String())
	} else {
		data.DetectionModelId = types.StringNull()
	}
	if value := res.Get("mqttBrokerId"); value.Exists() && value.Value() != nil {
		data.MqttBrokerId = types.StringValue(value.String())
	} else {
		data.MqttBrokerId = types.StringNull()
	}
	if value := res.Get("senseEnabled"); value.Exists() && value.Value() != nil {
		data.SenseEnabled = types.BoolValue(value.Bool())
	} else {
		data.SenseEnabled = types.BoolNull()
	}
	if value := res.Get("audioDetection.enabled"); value.Exists() && value.Value() != nil {
		data.AudioDetectionEnabled = types.BoolValue(value.Bool())
	} else {
		data.AudioDetectionEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *CameraSense) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("detectionModelId"); value.Exists() && !data.DetectionModelId.IsNull() {
		data.DetectionModelId = types.StringValue(value.String())
	} else {
		data.DetectionModelId = types.StringNull()
	}
	if value := res.Get("mqttBrokerId"); value.Exists() && !data.MqttBrokerId.IsNull() {
		data.MqttBrokerId = types.StringValue(value.String())
	} else {
		data.MqttBrokerId = types.StringNull()
	}
	if value := res.Get("senseEnabled"); value.Exists() && !data.SenseEnabled.IsNull() {
		data.SenseEnabled = types.BoolValue(value.Bool())
	} else {
		data.SenseEnabled = types.BoolNull()
	}
	if value := res.Get("audioDetection.enabled"); value.Exists() && !data.AudioDetectionEnabled.IsNull() {
		data.AudioDetectionEnabled = types.BoolValue(value.Bool())
	} else {
		data.AudioDetectionEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CameraSense) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data CameraSense) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
