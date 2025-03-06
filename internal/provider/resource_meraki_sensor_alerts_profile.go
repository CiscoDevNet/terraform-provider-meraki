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
	_ resource.Resource                = &SensorAlertsProfileResource{}
	_ resource.ResourceWithImportState = &SensorAlertsProfileResource{}
)

func NewSensorAlertsProfileResource() resource.Resource {
	return &SensorAlertsProfileResource{}
}

type SensorAlertsProfileResource struct {
	client *meraki.Client
}

func (r *SensorAlertsProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sensor_alerts_profile"
}

func (r *SensorAlertsProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Sensor Alerts Profile` configuration.").String,

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
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Name of the sensor alert profile.").String,
				Required:            true,
			},
			"recipients_emails": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of emails that will receive information about the alert.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"recipients_http_server_ids": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of webhook endpoint IDs that will receive information about the alert.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"recipients_sms_numbers": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of SMS numbers that will receive information about the alert.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"schedule_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.").String,
				Optional:            true,
			},
			"conditions": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of conditions that will cause the profile to send an alert.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("If `above`, an alert will be sent when a sensor reads above the threshold. If `below`, an alert will be sent when a sensor reads below the threshold. Only applicable for temperature, humidity, realPower, apparentPower, powerFactor, voltage, current, and frequency thresholds.").AddStringEnumDescription("above", "below").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("above", "below"),
							},
						},
						"duration": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.").String,
							Optional:            true,
						},
						"metric": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of sensor metric that will be monitored for changes.").AddStringEnumDescription("apparentPower", "co2", "current", "door", "frequency", "humidity", "indoorAirQuality", "noise", "pm25", "powerFactor", "realPower", "temperature", "tvoc", "upstreamPower", "voltage", "water").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("apparentPower", "co2", "current", "door", "frequency", "humidity", "indoorAirQuality", "noise", "pm25", "powerFactor", "realPower", "temperature", "tvoc", "upstreamPower", "voltage", "water"),
							},
						},
						"threshold_apparent_power_draw": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold in volt-amps. Must be between 0 and 3750.").String,
							Optional:            true,
						},
						"threshold_co2_concentration": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as CO2 parts per million.").String,
							Optional:            true,
						},
						"threshold_co2_quality": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as a qualitative CO2 level.").AddStringEnumDescription("fair", "good", "inadequate", "poor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("fair", "good", "inadequate", "poor"),
							},
						},
						"threshold_current_draw": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold in amps. Must be between 0 and 15.").String,
							Optional:            true,
						},
						"threshold_door_open": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold for a door open event. Must be set to true.").String,
							Optional:            true,
						},
						"threshold_frequency_level": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold in hertz. Must be between 0 and 60.").String,
							Optional:            true,
						},
						"threshold_humidity_quality": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as a qualitative humidity level.").AddStringEnumDescription("fair", "good", "inadequate", "poor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("fair", "good", "inadequate", "poor"),
							},
						},
						"threshold_humidity_relative_percentage": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold in %RH.").String,
							Optional:            true,
						},
						"threshold_indoor_air_quality_quality": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as a qualitative indoor air quality level.").AddStringEnumDescription("fair", "good", "inadequate", "poor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("fair", "good", "inadequate", "poor"),
							},
						},
						"threshold_indoor_air_quality_score": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as indoor air quality score.").String,
							Optional:            true,
						},
						"threshold_noise_ambient_level": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as adjusted decibels.").String,
							Optional:            true,
						},
						"threshold_noise_ambient_quality": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as a qualitative ambient noise level.").AddStringEnumDescription("fair", "good", "inadequate", "poor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("fair", "good", "inadequate", "poor"),
							},
						},
						"threshold_pm25_concentration": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as PM2.5 parts per million.").String,
							Optional:            true,
						},
						"threshold_pm25_quality": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as a qualitative PM2.5 level.").AddStringEnumDescription("fair", "good", "inadequate", "poor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("fair", "good", "inadequate", "poor"),
							},
						},
						"threshold_power_factor_percentage": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as the ratio of active power to apparent power. Must be between 0 and 100.").String,
							Optional:            true,
						},
						"threshold_real_power_draw": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold in watts. Must be between 0 and 3750.").String,
							Optional:            true,
						},
						"threshold_temperature_celsius": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold in degrees Celsius.").String,
							Optional:            true,
						},
						"threshold_temperature_fahrenheit": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold in degrees Fahrenheit.").String,
							Optional:            true,
						},
						"threshold_temperature_quality": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as a qualitative temperature level.").AddStringEnumDescription("fair", "good", "inadequate", "poor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("fair", "good", "inadequate", "poor"),
							},
						},
						"threshold_tvoc_concentration": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as TVOC micrograms per cubic meter.").String,
							Optional:            true,
						},
						"threshold_tvoc_quality": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold as a qualitative TVOC level.").AddStringEnumDescription("fair", "good", "inadequate", "poor").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("fair", "good", "inadequate", "poor"),
							},
						},
						"threshold_upstream_power_outage_detected": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold for an upstream power event. Must be set to true.").String,
							Optional:            true,
						},
						"threshold_voltage_level": schema.Float64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold in volts. Must be between 0 and 250.").String,
							Optional:            true,
						},
						"threshold_water_present": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Alerting threshold for a water detection event. Must be set to true.").String,
							Optional:            true,
						},
					},
				},
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of device serials assigned to this sensor alert profile.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *SensorAlertsProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *SensorAlertsProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SensorAlertsProfile

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, SensorAlertsProfile{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("profileId").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *SensorAlertsProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SensorAlertsProfile

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

func (r *SensorAlertsProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SensorAlertsProfile

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

func (r *SensorAlertsProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SensorAlertsProfile

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
func (r *SensorAlertsProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
