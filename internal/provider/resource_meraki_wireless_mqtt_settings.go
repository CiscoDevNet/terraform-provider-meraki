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
	_ resource.Resource                = &WirelessMQTTSettingsResource{}
	_ resource.ResourceWithImportState = &WirelessMQTTSettingsResource{}
)

func NewWirelessMQTTSettingsResource() resource.Resource {
	return &WirelessMQTTSettingsResource{}
}

type WirelessMQTTSettingsResource struct {
	client *meraki.Client
}

func (r *WirelessMQTTSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_mqtt_settings"
}

func (r *WirelessMQTTSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Wireless MQTT Settings` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Organization ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ble_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BLE Enabled").String,
				Optional:            true,
			},
			"ble_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BLE type of clients to publish telemetry. Valid types are: All, iBeacon, Eddystone, Unknown").AddStringEnumDescription("All", "Eddystone", "iBeacon", "Unknown").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("All", "Eddystone", "iBeacon", "Unknown"),
				},
			},
			"ble_allow_lists_macs": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allowed MAC List").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ble_allow_lists_uuids": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allowed UUID List").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"ble_flush_frequency": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("BLE Flush frequency in seconds. Will be between 1 and 2147483647. Default is 60 seconds").String,
				Optional:            true,
			},
			"ble_hysteresis_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("BLE Hysteresis Enabled").String,
				Optional:            true,
			},
			"ble_hysteresis_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("BLE Threshold. Will be between 1 and 2147483647. Default is 1 second").String,
				Optional:            true,
			},
			"mqtt_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MQTT Enabled").String,
				Optional:            true,
			},
			"mqtt_topic": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("MQTT Topic").String,
				Optional:            true,
			},
			"mqtt_broker_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Broker Config Name").String,
				Optional:            true,
			},
			"mqtt_publishing_frequency": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("MQTT Publishing Frequency in seconds. Will be between 1 and 2147483647. Default is 1 second").String,
				Optional:            true,
			},
			"mqtt_publishing_qos": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("MQTT Publishing QoS. Valid types are: 0, 1, 2").String,
				Optional:            true,
			},
			"mqtt_message_fields": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Select fields to populate in MQTT messages. Valid types are: RSSI, AP MAC address, Client MAC address, Timestamp, Radio, Network ID, Beacon type, Raw payload, Client UUID, Client major value, Client minor value, Signal power, Band, Slot ID").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network ID").String,
				Optional:            true,
			},
			"wifi_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Wi-Fi Enabled").String,
				Optional:            true,
			},
			"wifi_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Wi-Fi client type. Valid types are: Visible, Associated").AddStringEnumDescription("Associated", "Visible").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Associated", "Visible"),
				},
			},
			"wifi_allow_lists_macs": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Allowed MAC List").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"wifi_flush_frequency": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Wi-Fi Flush frequency in seconds. Will be between 1 and 2147483647. Default is 60 seconds").String,
				Optional:            true,
			},
			"wifi_hysteresis_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Wi-Fi Hysteresis Enabled").String,
				Optional:            true,
			},
			"wifi_hysteresis_threshold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Wi-Fi Threshold. Will be between 1 and 2147483647. Default is 1 second").String,
				Optional:            true,
			},
		},
	}
}

func (r *WirelessMQTTSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *WirelessMQTTSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WirelessMQTTSettings

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, WirelessMQTTSettings{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.OrganizationId
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

func (r *WirelessMQTTSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WirelessMQTTSettings

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath())
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	res = meraki.Res{Result: res.Get("items.0")}

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

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *WirelessMQTTSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WirelessMQTTSettings

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
	res, err := r.client.Put(plan.getPath(), body)
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

func (r *WirelessMQTTSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WirelessMQTTSettings

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *WirelessMQTTSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <organization_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
