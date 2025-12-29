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
	"slices"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceSensorAlertsProfiles struct {
	Id             types.String                        `tfsdk:"id"`
	OrganizationId types.String                        `tfsdk:"organization_id"`
	NetworkId      types.String                        `tfsdk:"network_id"`
	Items          []ResourceSensorAlertsProfilesItems `tfsdk:"items"`
}

type ResourceSensorAlertsProfilesItems struct {
	Id                      types.String                             `tfsdk:"id"`
	IncludeSensorUrl        types.Bool                               `tfsdk:"include_sensor_url"`
	Message                 types.String                             `tfsdk:"message"`
	Name                    types.String                             `tfsdk:"name"`
	RecipientsEmails        types.List                               `tfsdk:"recipients_emails"`
	RecipientsHttpServerIds types.List                               `tfsdk:"recipients_http_server_ids"`
	RecipientsSmsNumbers    types.List                               `tfsdk:"recipients_sms_numbers"`
	ScheduleId              types.String                             `tfsdk:"schedule_id"`
	Conditions              []ResourceSensorAlertsProfilesConditions `tfsdk:"conditions"`
	Serials                 types.List                               `tfsdk:"serials"`
}

type ResourceSensorAlertsProfilesConditions struct {
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

type ResourceSensorAlertsProfilesIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	NetworkId      types.String `tfsdk:"network_id"`
	ItemIds        types.List   `tfsdk:"item_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceSensorAlertsProfiles) getPath() string {
	return fmt.Sprintf("/networks/%v/sensor/alerts/profiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceSensorAlertsProfilesItems) toBody(ctx context.Context, state ResourceSensorAlertsProfilesItems) string {
	body := ""
	if !data.IncludeSensorUrl.IsNull() {
		body, _ = sjson.Set(body, "includeSensorUrl", data.IncludeSensorUrl.ValueBool())
	}
	if !data.Message.IsNull() {
		body, _ = sjson.Set(body, "message", data.Message.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.RecipientsEmails.IsNull() {
		var values []string
		data.RecipientsEmails.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "recipients.emails", values)
	}
	if !data.RecipientsHttpServerIds.IsNull() {
		var values []string
		data.RecipientsHttpServerIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "recipients.httpServerIds", values)
	}
	if !data.RecipientsSmsNumbers.IsNull() {
		var values []string
		data.RecipientsSmsNumbers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "recipients.smsNumbers", values)
	}
	if !data.ScheduleId.IsNull() {
		body, _ = sjson.Set(body, "schedule.id", data.ScheduleId.ValueString())
	}
	{
		body, _ = sjson.Set(body, "conditions", []interface{}{})
		for _, item := range data.Conditions {
			itemBody := ""
			if !item.Direction.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "direction", item.Direction.ValueString())
			}
			if !item.Duration.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "duration", item.Duration.ValueInt64())
			}
			if !item.Metric.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "metric", item.Metric.ValueString())
			}
			if !item.ThresholdApparentPowerDraw.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.apparentPower.draw", item.ThresholdApparentPowerDraw.ValueFloat64())
			}
			if !item.ThresholdCo2Concentration.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.co2.concentration", item.ThresholdCo2Concentration.ValueInt64())
			}
			if !item.ThresholdCo2Quality.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.co2.quality", item.ThresholdCo2Quality.ValueString())
			}
			if !item.ThresholdCurrentDraw.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.current.draw", item.ThresholdCurrentDraw.ValueFloat64())
			}
			if !item.ThresholdDoorOpen.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.door.open", item.ThresholdDoorOpen.ValueBool())
			}
			if !item.ThresholdFrequencyLevel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.frequency.level", item.ThresholdFrequencyLevel.ValueFloat64())
			}
			if !item.ThresholdHumidityQuality.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.humidity.quality", item.ThresholdHumidityQuality.ValueString())
			}
			if !item.ThresholdHumidityRelativePercentage.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.humidity.relativePercentage", item.ThresholdHumidityRelativePercentage.ValueInt64())
			}
			if !item.ThresholdIndoorAirQualityQuality.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.indoorAirQuality.quality", item.ThresholdIndoorAirQualityQuality.ValueString())
			}
			if !item.ThresholdIndoorAirQualityScore.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.indoorAirQuality.score", item.ThresholdIndoorAirQualityScore.ValueInt64())
			}
			if !item.ThresholdNoiseAmbientLevel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.noise.ambient.level", item.ThresholdNoiseAmbientLevel.ValueInt64())
			}
			if !item.ThresholdNoiseAmbientQuality.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.noise.ambient.quality", item.ThresholdNoiseAmbientQuality.ValueString())
			}
			if !item.ThresholdPm25Concentration.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.pm25.concentration", item.ThresholdPm25Concentration.ValueInt64())
			}
			if !item.ThresholdPm25Quality.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.pm25.quality", item.ThresholdPm25Quality.ValueString())
			}
			if !item.ThresholdPowerFactorPercentage.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.powerFactor.percentage", item.ThresholdPowerFactorPercentage.ValueInt64())
			}
			if !item.ThresholdRealPowerDraw.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.realPower.draw", item.ThresholdRealPowerDraw.ValueFloat64())
			}
			if !item.ThresholdTemperatureCelsius.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.temperature.celsius", item.ThresholdTemperatureCelsius.ValueFloat64())
			}
			if !item.ThresholdTemperatureFahrenheit.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.temperature.fahrenheit", item.ThresholdTemperatureFahrenheit.ValueFloat64())
			}
			if !item.ThresholdTemperatureQuality.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.temperature.quality", item.ThresholdTemperatureQuality.ValueString())
			}
			if !item.ThresholdTvocConcentration.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.tvoc.concentration", item.ThresholdTvocConcentration.ValueInt64())
			}
			if !item.ThresholdTvocQuality.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.tvoc.quality", item.ThresholdTvocQuality.ValueString())
			}
			if !item.ThresholdUpstreamPowerOutageDetected.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.upstreamPower.outageDetected", item.ThresholdUpstreamPowerOutageDetected.ValueBool())
			}
			if !item.ThresholdVoltageLevel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.voltage.level", item.ThresholdVoltageLevel.ValueFloat64())
			}
			if !item.ThresholdWaterPresent.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold.water.present", item.ThresholdWaterPresent.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "conditions.-1", itemBody)
		}
	}
	if !data.Serials.IsNull() {
		var values []string
		data.Serials.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "serials", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceSensorAlertsProfiles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceSensorAlertsProfilesItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceSensorAlertsProfilesItems{}
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
			data.Conditions = make([]ResourceSensorAlertsProfilesConditions, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceSensorAlertsProfilesConditions{}
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
	index := 0
	res.ForEach(func(k, res gjson.Result) bool {
		data.Items[index].Id = types.StringValue(res.Get("profileId").String())
		index++
		return true
	})
	data.Id = data.OrganizationId
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceSensorAlertsProfiles) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("profileId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
		if value := res.Get("includeSensorUrl"); value.Exists() && !data.IncludeSensorUrl.IsNull() {
			data.IncludeSensorUrl = types.BoolValue(value.Bool())
		} else {
			data.IncludeSensorUrl = types.BoolNull()
		}
		if value := res.Get("message"); value.Exists() && !data.Message.IsNull() {
			data.Message = types.StringValue(value.String())
		} else {
			data.Message = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("recipients.emails"); value.Exists() && !data.RecipientsEmails.IsNull() {
			data.RecipientsEmails = helpers.GetStringList(value.Array())
		} else {
			data.RecipientsEmails = types.ListNull(types.StringType)
		}
		if value := res.Get("recipients.httpServerIds"); value.Exists() && !data.RecipientsHttpServerIds.IsNull() {
			data.RecipientsHttpServerIds = helpers.GetStringList(value.Array())
		} else {
			data.RecipientsHttpServerIds = types.ListNull(types.StringType)
		}
		if value := res.Get("recipients.smsNumbers"); value.Exists() && !data.RecipientsSmsNumbers.IsNull() {
			data.RecipientsSmsNumbers = helpers.GetStringList(value.Array())
		} else {
			data.RecipientsSmsNumbers = types.ListNull(types.StringType)
		}
		if value := res.Get("schedule.id"); value.Exists() && !data.ScheduleId.IsNull() {
			data.ScheduleId = types.StringValue(value.String())
		} else {
			data.ScheduleId = types.StringNull()
		}
		for i := 0; i < len(data.Conditions); i++ {
			keys := [...]string{"direction", "metric"}
			keyValues := [...]string{data.Conditions[i].Direction.ValueString(), data.Conditions[i].Metric.ValueString()}

			parent := &data
			data := (*parent).Conditions[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("conditions").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing Conditions[%d] = %+v",
					i,
					(*parent).Conditions[i],
				))
				(*parent).Conditions = slices.Delete((*parent).Conditions, i, i+1)
				i--

				continue
			}
			if value := res.Get("direction"); value.Exists() && !data.Direction.IsNull() {
				data.Direction = types.StringValue(value.String())
			} else {
				data.Direction = types.StringNull()
			}
			if value := res.Get("duration"); value.Exists() && !data.Duration.IsNull() {
				data.Duration = types.Int64Value(value.Int())
			} else {
				data.Duration = types.Int64Null()
			}
			if value := res.Get("metric"); value.Exists() && !data.Metric.IsNull() {
				data.Metric = types.StringValue(value.String())
			} else {
				data.Metric = types.StringNull()
			}
			if value := res.Get("threshold.apparentPower.draw"); value.Exists() && !data.ThresholdApparentPowerDraw.IsNull() {
				data.ThresholdApparentPowerDraw = types.Float64Value(value.Float())
			} else {
				data.ThresholdApparentPowerDraw = types.Float64Null()
			}
			if value := res.Get("threshold.co2.concentration"); value.Exists() && !data.ThresholdCo2Concentration.IsNull() {
				data.ThresholdCo2Concentration = types.Int64Value(value.Int())
			} else {
				data.ThresholdCo2Concentration = types.Int64Null()
			}
			if value := res.Get("threshold.co2.quality"); value.Exists() && !data.ThresholdCo2Quality.IsNull() {
				data.ThresholdCo2Quality = types.StringValue(value.String())
			} else {
				data.ThresholdCo2Quality = types.StringNull()
			}
			if value := res.Get("threshold.current.draw"); value.Exists() && !data.ThresholdCurrentDraw.IsNull() {
				data.ThresholdCurrentDraw = types.Float64Value(value.Float())
			} else {
				data.ThresholdCurrentDraw = types.Float64Null()
			}
			if value := res.Get("threshold.door.open"); value.Exists() && !data.ThresholdDoorOpen.IsNull() {
				data.ThresholdDoorOpen = types.BoolValue(value.Bool())
			} else {
				data.ThresholdDoorOpen = types.BoolNull()
			}
			if value := res.Get("threshold.frequency.level"); value.Exists() && !data.ThresholdFrequencyLevel.IsNull() {
				data.ThresholdFrequencyLevel = types.Float64Value(value.Float())
			} else {
				data.ThresholdFrequencyLevel = types.Float64Null()
			}
			if value := res.Get("threshold.humidity.quality"); value.Exists() && !data.ThresholdHumidityQuality.IsNull() {
				data.ThresholdHumidityQuality = types.StringValue(value.String())
			} else {
				data.ThresholdHumidityQuality = types.StringNull()
			}
			if value := res.Get("threshold.humidity.relativePercentage"); value.Exists() && !data.ThresholdHumidityRelativePercentage.IsNull() {
				data.ThresholdHumidityRelativePercentage = types.Int64Value(value.Int())
			} else {
				data.ThresholdHumidityRelativePercentage = types.Int64Null()
			}
			if value := res.Get("threshold.indoorAirQuality.quality"); value.Exists() && !data.ThresholdIndoorAirQualityQuality.IsNull() {
				data.ThresholdIndoorAirQualityQuality = types.StringValue(value.String())
			} else {
				data.ThresholdIndoorAirQualityQuality = types.StringNull()
			}
			if value := res.Get("threshold.indoorAirQuality.score"); value.Exists() && !data.ThresholdIndoorAirQualityScore.IsNull() {
				data.ThresholdIndoorAirQualityScore = types.Int64Value(value.Int())
			} else {
				data.ThresholdIndoorAirQualityScore = types.Int64Null()
			}
			if value := res.Get("threshold.noise.ambient.level"); value.Exists() && !data.ThresholdNoiseAmbientLevel.IsNull() {
				data.ThresholdNoiseAmbientLevel = types.Int64Value(value.Int())
			} else {
				data.ThresholdNoiseAmbientLevel = types.Int64Null()
			}
			if value := res.Get("threshold.noise.ambient.quality"); value.Exists() && !data.ThresholdNoiseAmbientQuality.IsNull() {
				data.ThresholdNoiseAmbientQuality = types.StringValue(value.String())
			} else {
				data.ThresholdNoiseAmbientQuality = types.StringNull()
			}
			if value := res.Get("threshold.pm25.concentration"); value.Exists() && !data.ThresholdPm25Concentration.IsNull() {
				data.ThresholdPm25Concentration = types.Int64Value(value.Int())
			} else {
				data.ThresholdPm25Concentration = types.Int64Null()
			}
			if value := res.Get("threshold.pm25.quality"); value.Exists() && !data.ThresholdPm25Quality.IsNull() {
				data.ThresholdPm25Quality = types.StringValue(value.String())
			} else {
				data.ThresholdPm25Quality = types.StringNull()
			}
			if value := res.Get("threshold.powerFactor.percentage"); value.Exists() && !data.ThresholdPowerFactorPercentage.IsNull() {
				data.ThresholdPowerFactorPercentage = types.Int64Value(value.Int())
			} else {
				data.ThresholdPowerFactorPercentage = types.Int64Null()
			}
			if value := res.Get("threshold.realPower.draw"); value.Exists() && !data.ThresholdRealPowerDraw.IsNull() {
				data.ThresholdRealPowerDraw = types.Float64Value(value.Float())
			} else {
				data.ThresholdRealPowerDraw = types.Float64Null()
			}
			if value := res.Get("threshold.temperature.celsius"); value.Exists() && !data.ThresholdTemperatureCelsius.IsNull() {
				data.ThresholdTemperatureCelsius = types.Float64Value(value.Float())
			} else {
				data.ThresholdTemperatureCelsius = types.Float64Null()
			}
			if value := res.Get("threshold.temperature.fahrenheit"); value.Exists() && !data.ThresholdTemperatureFahrenheit.IsNull() {
				data.ThresholdTemperatureFahrenheit = types.Float64Value(value.Float())
			} else {
				data.ThresholdTemperatureFahrenheit = types.Float64Null()
			}
			if value := res.Get("threshold.temperature.quality"); value.Exists() && !data.ThresholdTemperatureQuality.IsNull() {
				data.ThresholdTemperatureQuality = types.StringValue(value.String())
			} else {
				data.ThresholdTemperatureQuality = types.StringNull()
			}
			if value := res.Get("threshold.tvoc.concentration"); value.Exists() && !data.ThresholdTvocConcentration.IsNull() {
				data.ThresholdTvocConcentration = types.Int64Value(value.Int())
			} else {
				data.ThresholdTvocConcentration = types.Int64Null()
			}
			if value := res.Get("threshold.tvoc.quality"); value.Exists() && !data.ThresholdTvocQuality.IsNull() {
				data.ThresholdTvocQuality = types.StringValue(value.String())
			} else {
				data.ThresholdTvocQuality = types.StringNull()
			}
			if value := res.Get("threshold.upstreamPower.outageDetected"); value.Exists() && !data.ThresholdUpstreamPowerOutageDetected.IsNull() {
				data.ThresholdUpstreamPowerOutageDetected = types.BoolValue(value.Bool())
			} else {
				data.ThresholdUpstreamPowerOutageDetected = types.BoolNull()
			}
			if value := res.Get("threshold.voltage.level"); value.Exists() && !data.ThresholdVoltageLevel.IsNull() {
				data.ThresholdVoltageLevel = types.Float64Value(value.Float())
			} else {
				data.ThresholdVoltageLevel = types.Float64Null()
			}
			if value := res.Get("threshold.water.present"); value.Exists() && !data.ThresholdWaterPresent.IsNull() {
				data.ThresholdWaterPresent = types.BoolValue(value.Bool())
			} else {
				data.ThresholdWaterPresent = types.BoolNull()
			}
			(*parent).Conditions[i] = data
		}
		if value := res.Get("serials"); value.Exists() && !data.Serials.IsNull() {
			data.Serials = helpers.GetStringList(value.Array())
		} else {
			data.Serials = types.ListNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyPartial(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceSensorAlertsProfiles) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyImport

func (data *ResourceSensorAlertsProfiles) fromBodyImport(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	toBeDeleted := make([]int, 0)
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result
		found := false

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("profileId").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					found = true
					return false
				}
				return true
			},
		)
		if !found {
			toBeDeleted = append(toBeDeleted, i)
			continue
		}
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
		if value := res.Get("recipients.emails"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.RecipientsEmails = helpers.GetStringList(value.Array())
		} else {
			data.RecipientsEmails = types.ListNull(types.StringType)
		}
		if value := res.Get("recipients.httpServerIds"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.RecipientsHttpServerIds = helpers.GetStringList(value.Array())
		} else {
			data.RecipientsHttpServerIds = types.ListNull(types.StringType)
		}
		if value := res.Get("recipients.smsNumbers"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.RecipientsSmsNumbers = helpers.GetStringList(value.Array())
		} else {
			data.RecipientsSmsNumbers = types.ListNull(types.StringType)
		}
		if value := res.Get("schedule.id"); value.Exists() && value.Value() != nil {
			data.ScheduleId = types.StringValue(value.String())
		} else {
			data.ScheduleId = types.StringNull()
		}
		if value := res.Get("conditions"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.Conditions = make([]ResourceSensorAlertsProfilesConditions, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceSensorAlertsProfilesConditions{}
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
		if value := res.Get("serials"); value.Exists() && value.Value() != nil && len(value.Array()) > 0 {
			data.Serials = helpers.GetStringList(value.Array())
		} else {
			data.Serials = types.ListNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
	for i := len(toBeDeleted) - 1; i >= 0; i-- {
		tflog.Debug(ctx, fmt.Sprintf("fromBodyImport(), removing item with id: %s", data.Items[toBeDeleted[i]].Id.ValueString()))
		data.Items = slices.Delete(data.Items, toBeDeleted[i], toBeDeleted[i]+1)
	}
}

// End of section. //template:end fromBodyImport

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ResourceSensorAlertsProfilesIdentity) toIdentity(ctx context.Context, plan *ResourceSensorAlertsProfiles) {
	data.OrganizationId = plan.OrganizationId
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ResourceSensorAlertsProfiles) fromIdentity(ctx context.Context, identity *ResourceSensorAlertsProfilesIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceSensorAlertsProfiles) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceSensorAlertsProfiles) hasChanges(ctx context.Context, state *ResourceSensorAlertsProfiles, id string) bool {
	hasChanges := false

	item := ResourceSensorAlertsProfilesItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceSensorAlertsProfilesItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.IncludeSensorUrl.Equal(stateItem.IncludeSensorUrl) {
		hasChanges = true
	}
	if !item.Message.Equal(stateItem.Message) {
		hasChanges = true
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	if !item.RecipientsEmails.Equal(stateItem.RecipientsEmails) {
		hasChanges = true
	}
	if !item.RecipientsHttpServerIds.Equal(stateItem.RecipientsHttpServerIds) {
		hasChanges = true
	}
	if !item.RecipientsSmsNumbers.Equal(stateItem.RecipientsSmsNumbers) {
		hasChanges = true
	}
	if !item.ScheduleId.Equal(stateItem.ScheduleId) {
		hasChanges = true
	}
	if len(item.Conditions) != len(stateItem.Conditions) {
		hasChanges = true
	} else {
		for i := range item.Conditions {
			if !item.Conditions[i].Direction.Equal(stateItem.Conditions[i].Direction) {
				hasChanges = true
			}
			if !item.Conditions[i].Duration.Equal(stateItem.Conditions[i].Duration) {
				hasChanges = true
			}
			if !item.Conditions[i].Metric.Equal(stateItem.Conditions[i].Metric) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdApparentPowerDraw.Equal(stateItem.Conditions[i].ThresholdApparentPowerDraw) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdCo2Concentration.Equal(stateItem.Conditions[i].ThresholdCo2Concentration) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdCo2Quality.Equal(stateItem.Conditions[i].ThresholdCo2Quality) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdCurrentDraw.Equal(stateItem.Conditions[i].ThresholdCurrentDraw) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdDoorOpen.Equal(stateItem.Conditions[i].ThresholdDoorOpen) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdFrequencyLevel.Equal(stateItem.Conditions[i].ThresholdFrequencyLevel) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdHumidityQuality.Equal(stateItem.Conditions[i].ThresholdHumidityQuality) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdHumidityRelativePercentage.Equal(stateItem.Conditions[i].ThresholdHumidityRelativePercentage) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdIndoorAirQualityQuality.Equal(stateItem.Conditions[i].ThresholdIndoorAirQualityQuality) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdIndoorAirQualityScore.Equal(stateItem.Conditions[i].ThresholdIndoorAirQualityScore) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdNoiseAmbientLevel.Equal(stateItem.Conditions[i].ThresholdNoiseAmbientLevel) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdNoiseAmbientQuality.Equal(stateItem.Conditions[i].ThresholdNoiseAmbientQuality) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdPm25Concentration.Equal(stateItem.Conditions[i].ThresholdPm25Concentration) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdPm25Quality.Equal(stateItem.Conditions[i].ThresholdPm25Quality) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdPowerFactorPercentage.Equal(stateItem.Conditions[i].ThresholdPowerFactorPercentage) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdRealPowerDraw.Equal(stateItem.Conditions[i].ThresholdRealPowerDraw) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdTemperatureCelsius.Equal(stateItem.Conditions[i].ThresholdTemperatureCelsius) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdTemperatureFahrenheit.Equal(stateItem.Conditions[i].ThresholdTemperatureFahrenheit) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdTemperatureQuality.Equal(stateItem.Conditions[i].ThresholdTemperatureQuality) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdTvocConcentration.Equal(stateItem.Conditions[i].ThresholdTvocConcentration) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdTvocQuality.Equal(stateItem.Conditions[i].ThresholdTvocQuality) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdUpstreamPowerOutageDetected.Equal(stateItem.Conditions[i].ThresholdUpstreamPowerOutageDetected) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdVoltageLevel.Equal(stateItem.Conditions[i].ThresholdVoltageLevel) {
				hasChanges = true
			}
			if !item.Conditions[i].ThresholdWaterPresent.Equal(stateItem.Conditions[i].ThresholdWaterPresent) {
				hasChanges = true
			}
		}
	}
	if !item.Serials.Equal(stateItem.Serials) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
