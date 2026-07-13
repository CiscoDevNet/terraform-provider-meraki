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
// It emits a separate model - DataSourceCameraSense - used only by the data source, always including every
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

type DataSourceCameraSense struct {
	Id                    types.String `tfsdk:"id"`
	Serial                types.String `tfsdk:"serial"`
	DetectionModelId      types.String `tfsdk:"detection_model_id"`
	MqttBrokerId          types.String `tfsdk:"mqtt_broker_id"`
	SenseEnabled          types.Bool   `tfsdk:"sense_enabled"`
	AudioDetectionEnabled types.Bool   `tfsdk:"audio_detection_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceCameraSense) getPath() string {
	return fmt.Sprintf("/devices/%v/camera/sense", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceCameraSense) fromBody(ctx context.Context, res meraki.Res) {
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
