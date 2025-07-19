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
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SensorAlertsProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &SensorAlertsProfileDataSource{}
)

func NewSensorAlertsProfileDataSource() datasource.DataSource {
	return &SensorAlertsProfileDataSource{}
}

type SensorAlertsProfileDataSource struct {
	client *meraki.Client
}

func (d *SensorAlertsProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sensor_alerts_profile"
}

func (d *SensorAlertsProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Sensor Alerts Profile` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"include_sensor_url": schema.BoolAttribute{
				MarkdownDescription: "Include dashboard link to sensor in messages (default: true).",
				Computed:            true,
			},
			"message": schema.StringAttribute{
				MarkdownDescription: "A custom message that will appear in email and text message alerts.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the sensor alert profile.",
				Optional:            true,
				Computed:            true,
			},
			"recipients_emails": schema.ListAttribute{
				MarkdownDescription: "A list of emails that will receive information about the alert.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"recipients_http_server_ids": schema.ListAttribute{
				MarkdownDescription: "A list of webhook endpoint IDs that will receive information about the alert.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"recipients_sms_numbers": schema.ListAttribute{
				MarkdownDescription: "A list of SMS numbers that will receive information about the alert.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"schedule_id": schema.StringAttribute{
				MarkdownDescription: "ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.",
				Computed:            true,
			},
			"conditions": schema.ListNestedAttribute{
				MarkdownDescription: "List of conditions that will cause the profile to send an alert.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: "If `above`, an alert will be sent when a sensor reads above the threshold. If `below`, an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.",
							Computed:            true,
						},
						"duration": schema.Int64Attribute{
							MarkdownDescription: "Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 1 hour and 30 minutes, 2 hours, 4 hours, and 8 hours. Default is 0.",
							Computed:            true,
						},
						"metric": schema.StringAttribute{
							MarkdownDescription: "The type of sensor metric that will be monitored for changes.",
							Computed:            true,
						},
						"threshold_apparent_power_draw": schema.Float64Attribute{
							MarkdownDescription: "Alerting threshold in volt-amps. Must be between 0 and 3750.",
							Computed:            true,
						},
						"threshold_co2_concentration": schema.Int64Attribute{
							MarkdownDescription: "Alerting threshold as CO2 parts per million.",
							Computed:            true,
						},
						"threshold_co2_quality": schema.StringAttribute{
							MarkdownDescription: "Alerting threshold as a qualitative CO2 level.",
							Computed:            true,
						},
						"threshold_current_draw": schema.Float64Attribute{
							MarkdownDescription: "Alerting threshold in amps. Must be between 0 and 15.",
							Computed:            true,
						},
						"threshold_door_open": schema.BoolAttribute{
							MarkdownDescription: "Alerting threshold for a door open event. Must be set to true.",
							Computed:            true,
						},
						"threshold_frequency_level": schema.Float64Attribute{
							MarkdownDescription: "Alerting threshold in hertz. Must be between 0 and 60.",
							Computed:            true,
						},
						"threshold_humidity_quality": schema.StringAttribute{
							MarkdownDescription: "Alerting threshold as a qualitative humidity level.",
							Computed:            true,
						},
						"threshold_humidity_relative_percentage": schema.Int64Attribute{
							MarkdownDescription: "Alerting threshold in %RH.",
							Computed:            true,
						},
						"threshold_indoor_air_quality_quality": schema.StringAttribute{
							MarkdownDescription: "Alerting threshold as a qualitative indoor air quality level.",
							Computed:            true,
						},
						"threshold_indoor_air_quality_score": schema.Int64Attribute{
							MarkdownDescription: "Alerting threshold as indoor air quality score.",
							Computed:            true,
						},
						"threshold_noise_ambient_level": schema.Int64Attribute{
							MarkdownDescription: "Alerting threshold as adjusted decibels.",
							Computed:            true,
						},
						"threshold_noise_ambient_quality": schema.StringAttribute{
							MarkdownDescription: "Alerting threshold as a qualitative ambient noise level.",
							Computed:            true,
						},
						"threshold_pm25_concentration": schema.Int64Attribute{
							MarkdownDescription: "Alerting threshold as PM2.5 parts per million.",
							Computed:            true,
						},
						"threshold_pm25_quality": schema.StringAttribute{
							MarkdownDescription: "Alerting threshold as a qualitative PM2.5 level.",
							Computed:            true,
						},
						"threshold_power_factor_percentage": schema.Int64Attribute{
							MarkdownDescription: "Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.",
							Computed:            true,
						},
						"threshold_real_power_draw": schema.Float64Attribute{
							MarkdownDescription: "Alerting threshold in watts. Must be between 0 and 3750.",
							Computed:            true,
						},
						"threshold_temperature_celsius": schema.Float64Attribute{
							MarkdownDescription: "Alerting threshold in degrees Celsius.",
							Computed:            true,
						},
						"threshold_temperature_fahrenheit": schema.Float64Attribute{
							MarkdownDescription: "Alerting threshold in degrees Fahrenheit.",
							Computed:            true,
						},
						"threshold_temperature_quality": schema.StringAttribute{
							MarkdownDescription: "Alerting threshold as a qualitative temperature level.",
							Computed:            true,
						},
						"threshold_tvoc_concentration": schema.Int64Attribute{
							MarkdownDescription: "Alerting threshold as TVOC micrograms per cubic meter.",
							Computed:            true,
						},
						"threshold_tvoc_quality": schema.StringAttribute{
							MarkdownDescription: "Alerting threshold as a qualitative TVOC level.",
							Computed:            true,
						},
						"threshold_upstream_power_outage_detected": schema.BoolAttribute{
							MarkdownDescription: "Alerting threshold for an upstream power event. Must be set to true.",
							Computed:            true,
						},
						"threshold_voltage_level": schema.Float64Attribute{
							MarkdownDescription: "Alerting threshold in volts. Must be between 0 and 250.",
							Computed:            true,
						},
						"threshold_water_present": schema.BoolAttribute{
							MarkdownDescription: "Alerting threshold for a water detection event. Must be set to true.",
							Computed:            true,
						},
					},
				},
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: "List of device serials assigned to this sensor alert profile.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}
func (d *SensorAlertsProfileDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *SensorAlertsProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SensorAlertsProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SensorAlertsProfile

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res meraki.Res
	var err error
	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		found := false
		if len(res.Array()) > 0 {
			res.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("name").String() {
					if found {
						resp.Diagnostics.AddWarning("Multiple objects with same name", fmt.Sprintf("Found multiple objects with name: %s", config.Name.ValueString()))
						return false
					}
					config.Id = types.StringValue(v.Get("profileId").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					res = meraki.Res{Result: v}
					found = true
				}
				return true
			})
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}

	if !res.Exists() {
		res, err = d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
