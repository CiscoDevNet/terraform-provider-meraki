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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SensorAlertsProfiles struct {
	NetworkId types.String                `tfsdk:"network_id"`
	Items     []SensorAlertsProfilesItems `tfsdk:"items"`
}

type SensorAlertsProfilesItems struct {
	Id                      types.String                     `tfsdk:"id"`
	IncludeSensorUrl        types.Bool                       `tfsdk:"include_sensor_url"`
	Message                 types.String                     `tfsdk:"message"`
	Name                    types.String                     `tfsdk:"name"`
	RecipientsEmails        types.List                       `tfsdk:"recipients_emails"`
	RecipientsHttpServerIds types.List                       `tfsdk:"recipients_http_server_ids"`
	RecipientsSmsNumbers    types.List                       `tfsdk:"recipients_sms_numbers"`
	ScheduleId              types.String                     `tfsdk:"schedule_id"`
	Conditions              []SensorAlertsProfilesConditions `tfsdk:"conditions"`
	Serials                 types.List                       `tfsdk:"serials"`
}

type SensorAlertsProfilesConditions struct {
	Direction                            types.String  `tfsdk:"direction"`
	Duration                             types.Int64   `tfsdk:"duration"`
	Metric                               types.String  `tfsdk:"metric"`
	ThresholdApparentPowerDraw           types.Float64 `tfsdk:"threshold_apparent_power_draw"`
	ThresholdCo2Concentration            types.Int64   `tfsdk:"threshold_co2_concentration"`
	ThresholdCo2Quality                  types.String  `tfsdk:"threshold_co2_quality"`
	ThresholdCurrentDraw                 types.Float64 `tfsdk:"threshold_current_draw"`
	ThresholdDoorOpen                    types.Bool    `tfsdk:"threshold_door_open"`
	ThresholdFrequencyLevel              types.Float64 `tfsdk:"threshold_frequency_level"`
	ThresholdHumidityQuality             types.String  `tfsdk:"threshold_humidity_quality"`
	ThresholdHumidityRelativePercentage  types.Int64   `tfsdk:"threshold_humidity_relative_percentage"`
	ThresholdIndoorAirQualityQuality     types.String  `tfsdk:"threshold_indoor_air_quality_quality"`
	ThresholdIndoorAirQualityScore       types.Int64   `tfsdk:"threshold_indoor_air_quality_score"`
	ThresholdNoiseAmbientLevel           types.Int64   `tfsdk:"threshold_noise_ambient_level"`
	ThresholdNoiseAmbientQuality         types.String  `tfsdk:"threshold_noise_ambient_quality"`
	ThresholdPm25Concentration           types.Int64   `tfsdk:"threshold_pm25_concentration"`
	ThresholdPm25Quality                 types.String  `tfsdk:"threshold_pm25_quality"`
	ThresholdPowerFactorPercentage       types.Int64   `tfsdk:"threshold_power_factor_percentage"`
	ThresholdRealPowerDraw               types.Float64 `tfsdk:"threshold_real_power_draw"`
	ThresholdTemperatureCelsius          types.Float64 `tfsdk:"threshold_temperature_celsius"`
	ThresholdTemperatureFahrenheit       types.Float64 `tfsdk:"threshold_temperature_fahrenheit"`
	ThresholdTemperatureQuality          types.String  `tfsdk:"threshold_temperature_quality"`
	ThresholdTvocConcentration           types.Int64   `tfsdk:"threshold_tvoc_concentration"`
	ThresholdTvocQuality                 types.String  `tfsdk:"threshold_tvoc_quality"`
	ThresholdUpstreamPowerOutageDetected types.Bool    `tfsdk:"threshold_upstream_power_outage_detected"`
	ThresholdVoltageLevel                types.Float64 `tfsdk:"threshold_voltage_level"`
	ThresholdWaterPresent                types.Bool    `tfsdk:"threshold_water_present"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SensorAlertsProfiles) getPath() string {
	return fmt.Sprintf("/networks/%v/sensor/alerts/profiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SensorAlertsProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]SensorAlertsProfilesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := SensorAlertsProfilesItems{}
		data.Id = types.StringValue(res.Get("profileId").String())
		if value := res.Get("includeSensorUrl"); value.Exists() && value.Value() != nil {
			data.IncludeSensorUrl = types.BoolValue(value.Bool())
		} else {
			data.IncludeSensorUrl = types.BoolNull()
		}
		if value := res.Get("message"); value.Exists() && value.Value() != nil {
			data.Message = types.StringValue(value.String())
		} else {
			data.Message = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("recipients.emails"); value.Exists() && value.Value() != nil {
			data.RecipientsEmails = helpers.GetStringList(value.Array())
		} else {
			data.RecipientsEmails = types.ListNull(types.StringType)
		}
		if value := res.Get("recipients.httpServerIds"); value.Exists() && value.Value() != nil {
			data.RecipientsHttpServerIds = helpers.GetStringList(value.Array())
		} else {
			data.RecipientsHttpServerIds = types.ListNull(types.StringType)
		}
		if value := res.Get("recipients.smsNumbers"); value.Exists() && value.Value() != nil {
			data.RecipientsSmsNumbers = helpers.GetStringList(value.Array())
		} else {
			data.RecipientsSmsNumbers = types.ListNull(types.StringType)
		}
		if value := res.Get("schedule.id"); value.Exists() && value.Value() != nil {
			data.ScheduleId = types.StringValue(value.String())
		} else {
			data.ScheduleId = types.StringNull()
		}
		if value := res.Get("conditions"); value.Exists() && value.Value() != nil {
			data.Conditions = make([]SensorAlertsProfilesConditions, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := SensorAlertsProfilesConditions{}
				if value := res.Get("direction"); value.Exists() && value.Value() != nil {
					data.Direction = types.StringValue(value.String())
				} else {
					data.Direction = types.StringNull()
				}
				if value := res.Get("duration"); value.Exists() && value.Value() != nil {
					data.Duration = types.Int64Value(value.Int())
				} else {
					data.Duration = types.Int64Null()
				}
				if value := res.Get("metric"); value.Exists() && value.Value() != nil {
					data.Metric = types.StringValue(value.String())
				} else {
					data.Metric = types.StringNull()
				}
				if value := res.Get("threshold.apparentPower.draw"); value.Exists() && value.Value() != nil {
					data.ThresholdApparentPowerDraw = types.Float64Value(value.Float())
				} else {
					data.ThresholdApparentPowerDraw = types.Float64Null()
				}
				if value := res.Get("threshold.co2.concentration"); value.Exists() && value.Value() != nil {
					data.ThresholdCo2Concentration = types.Int64Value(value.Int())
				} else {
					data.ThresholdCo2Concentration = types.Int64Null()
				}
				if value := res.Get("threshold.co2.quality"); value.Exists() && value.Value() != nil {
					data.ThresholdCo2Quality = types.StringValue(value.String())
				} else {
					data.ThresholdCo2Quality = types.StringNull()
				}
				if value := res.Get("threshold.current.draw"); value.Exists() && value.Value() != nil {
					data.ThresholdCurrentDraw = types.Float64Value(value.Float())
				} else {
					data.ThresholdCurrentDraw = types.Float64Null()
				}
				if value := res.Get("threshold.door.open"); value.Exists() && value.Value() != nil {
					data.ThresholdDoorOpen = types.BoolValue(value.Bool())
				} else {
					data.ThresholdDoorOpen = types.BoolNull()
				}
				if value := res.Get("threshold.frequency.level"); value.Exists() && value.Value() != nil {
					data.ThresholdFrequencyLevel = types.Float64Value(value.Float())
				} else {
					data.ThresholdFrequencyLevel = types.Float64Null()
				}
				if value := res.Get("threshold.humidity.quality"); value.Exists() && value.Value() != nil {
					data.ThresholdHumidityQuality = types.StringValue(value.String())
				} else {
					data.ThresholdHumidityQuality = types.StringNull()
				}
				if value := res.Get("threshold.humidity.relativePercentage"); value.Exists() && value.Value() != nil {
					data.ThresholdHumidityRelativePercentage = types.Int64Value(value.Int())
				} else {
					data.ThresholdHumidityRelativePercentage = types.Int64Null()
				}
				if value := res.Get("threshold.indoorAirQuality.quality"); value.Exists() && value.Value() != nil {
					data.ThresholdIndoorAirQualityQuality = types.StringValue(value.String())
				} else {
					data.ThresholdIndoorAirQualityQuality = types.StringNull()
				}
				if value := res.Get("threshold.indoorAirQuality.score"); value.Exists() && value.Value() != nil {
					data.ThresholdIndoorAirQualityScore = types.Int64Value(value.Int())
				} else {
					data.ThresholdIndoorAirQualityScore = types.Int64Null()
				}
				if value := res.Get("threshold.noise.ambient.level"); value.Exists() && value.Value() != nil {
					data.ThresholdNoiseAmbientLevel = types.Int64Value(value.Int())
				} else {
					data.ThresholdNoiseAmbientLevel = types.Int64Null()
				}
				if value := res.Get("threshold.noise.ambient.quality"); value.Exists() && value.Value() != nil {
					data.ThresholdNoiseAmbientQuality = types.StringValue(value.String())
				} else {
					data.ThresholdNoiseAmbientQuality = types.StringNull()
				}
				if value := res.Get("threshold.pm25.concentration"); value.Exists() && value.Value() != nil {
					data.ThresholdPm25Concentration = types.Int64Value(value.Int())
				} else {
					data.ThresholdPm25Concentration = types.Int64Null()
				}
				if value := res.Get("threshold.pm25.quality"); value.Exists() && value.Value() != nil {
					data.ThresholdPm25Quality = types.StringValue(value.String())
				} else {
					data.ThresholdPm25Quality = types.StringNull()
				}
				if value := res.Get("threshold.powerFactor.percentage"); value.Exists() && value.Value() != nil {
					data.ThresholdPowerFactorPercentage = types.Int64Value(value.Int())
				} else {
					data.ThresholdPowerFactorPercentage = types.Int64Null()
				}
				if value := res.Get("threshold.realPower.draw"); value.Exists() && value.Value() != nil {
					data.ThresholdRealPowerDraw = types.Float64Value(value.Float())
				} else {
					data.ThresholdRealPowerDraw = types.Float64Null()
				}
				if value := res.Get("threshold.temperature.celsius"); value.Exists() && value.Value() != nil {
					data.ThresholdTemperatureCelsius = types.Float64Value(value.Float())
				} else {
					data.ThresholdTemperatureCelsius = types.Float64Null()
				}
				if value := res.Get("threshold.temperature.fahrenheit"); value.Exists() && value.Value() != nil {
					data.ThresholdTemperatureFahrenheit = types.Float64Value(value.Float())
				} else {
					data.ThresholdTemperatureFahrenheit = types.Float64Null()
				}
				if value := res.Get("threshold.temperature.quality"); value.Exists() && value.Value() != nil {
					data.ThresholdTemperatureQuality = types.StringValue(value.String())
				} else {
					data.ThresholdTemperatureQuality = types.StringNull()
				}
				if value := res.Get("threshold.tvoc.concentration"); value.Exists() && value.Value() != nil {
					data.ThresholdTvocConcentration = types.Int64Value(value.Int())
				} else {
					data.ThresholdTvocConcentration = types.Int64Null()
				}
				if value := res.Get("threshold.tvoc.quality"); value.Exists() && value.Value() != nil {
					data.ThresholdTvocQuality = types.StringValue(value.String())
				} else {
					data.ThresholdTvocQuality = types.StringNull()
				}
				if value := res.Get("threshold.upstreamPower.outageDetected"); value.Exists() && value.Value() != nil {
					data.ThresholdUpstreamPowerOutageDetected = types.BoolValue(value.Bool())
				} else {
					data.ThresholdUpstreamPowerOutageDetected = types.BoolNull()
				}
				if value := res.Get("threshold.voltage.level"); value.Exists() && value.Value() != nil {
					data.ThresholdVoltageLevel = types.Float64Value(value.Float())
				} else {
					data.ThresholdVoltageLevel = types.Float64Null()
				}
				if value := res.Get("threshold.water.present"); value.Exists() && value.Value() != nil {
					data.ThresholdWaterPresent = types.BoolValue(value.Bool())
				} else {
					data.ThresholdWaterPresent = types.BoolNull()
				}
				(*parent).Conditions = append((*parent).Conditions, data)
				return true
			})
		}
		if value := res.Get("serials"); value.Exists() && value.Value() != nil {
			data.Serials = helpers.GetStringList(value.Array())
		} else {
			data.Serials = types.ListNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
