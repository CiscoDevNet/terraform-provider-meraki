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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessMQTTSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessMQTTSettingsDataSource{}
)

func NewWirelessMQTTSettingsDataSource() datasource.DataSource {
	return &WirelessMQTTSettingsDataSource{}
}

type WirelessMQTTSettingsDataSource struct {
	client *meraki.Client
}

func (d *WirelessMQTTSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_mqtt_settings"
}

func (d *WirelessMQTTSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Wireless MQTT Settings` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
				Required:            true,
			},
			"ble_enabled": schema.BoolAttribute{
				MarkdownDescription: "BLE Enabled",
				Computed:            true,
			},
			"ble_type": schema.StringAttribute{
				MarkdownDescription: "BLE type of clients to publish telemetry. Valid types are: All, iBeacon, Eddystone, Unknown",
				Computed:            true,
			},
			"ble_allow_lists_macs": schema.ListAttribute{
				MarkdownDescription: "Allowed MAC List",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ble_allow_lists_uuids": schema.ListAttribute{
				MarkdownDescription: "Allowed UUID List",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"ble_flush_frequency": schema.Int64Attribute{
				MarkdownDescription: "BLE Flush frequency in seconds. Will be between 1 and 2147483647. Default is 60 seconds",
				Computed:            true,
			},
			"ble_hysteresis_enabled": schema.BoolAttribute{
				MarkdownDescription: "BLE Hysteresis Enabled",
				Computed:            true,
			},
			"ble_hysteresis_threshold": schema.Int64Attribute{
				MarkdownDescription: "BLE Threshold. Will be between 1 and 2147483647. Default is 1 second",
				Computed:            true,
			},
			"mqtt_enabled": schema.BoolAttribute{
				MarkdownDescription: "MQTT Enabled",
				Computed:            true,
			},
			"mqtt_topic": schema.StringAttribute{
				MarkdownDescription: "MQTT Topic",
				Computed:            true,
			},
			"mqtt_broker_name": schema.StringAttribute{
				MarkdownDescription: "Broker Config Name",
				Computed:            true,
			},
			"mqtt_publishing_frequency": schema.Int64Attribute{
				MarkdownDescription: "MQTT Publishing Frequency in seconds. Will be between 1 and 2147483647. Default is 1 second",
				Computed:            true,
			},
			"mqtt_publishing_qos": schema.Int64Attribute{
				MarkdownDescription: "MQTT Publishing QoS. Valid types are: 0, 1, 2",
				Computed:            true,
			},
			"mqtt_message_fields": schema.ListAttribute{
				MarkdownDescription: "Select fields to populate in MQTT messages. Valid types are: RSSI, AP MAC address, Client MAC address, Timestamp, Radio, Network ID, Beacon type, Raw payload, Client UUID, Client major value, Client minor value, Signal power, Band, Slot ID",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"wifi_enabled": schema.BoolAttribute{
				MarkdownDescription: "Wi-Fi Enabled",
				Computed:            true,
			},
			"wifi_type": schema.StringAttribute{
				MarkdownDescription: "Wi-Fi client type. Valid types are: Visible, Associated",
				Computed:            true,
			},
			"wifi_allow_lists_macs": schema.ListAttribute{
				MarkdownDescription: "Allowed MAC List",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"wifi_flush_frequency": schema.Int64Attribute{
				MarkdownDescription: "Wi-Fi Flush frequency in seconds. Will be between 1 and 2147483647. Default is 60 seconds",
				Computed:            true,
			},
			"wifi_hysteresis_enabled": schema.BoolAttribute{
				MarkdownDescription: "Wi-Fi Hysteresis Enabled",
				Computed:            true,
			},
			"wifi_hysteresis_threshold": schema.Int64Attribute{
				MarkdownDescription: "Wi-Fi Threshold. Will be between 1 and 2147483647. Default is 1 second",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessMQTTSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

func (d *WirelessMQTTSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessMQTTSettings

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res meraki.Res
	var err error

	if !res.Exists() {
		res, err = d.client.Get(config.getPath() + "?networkIds[]=" + url.QueryEscape(config.NetworkId.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
		res = meraki.Res{Result: res.Get("items.0")}
	}

	config.fromBody(ctx, res)
	config.Id = config.OrganizationId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
