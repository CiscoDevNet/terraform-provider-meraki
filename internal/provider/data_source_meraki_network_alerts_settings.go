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
	_ datasource.DataSource              = &NetworkAlertsSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkAlertsSettingsDataSource{}
)

func NewNetworkAlertsSettingsDataSource() datasource.DataSource {
	return &NetworkAlertsSettingsDataSource{}
}

type NetworkAlertsSettingsDataSource struct {
	client *meraki.Client
}

func (d *NetworkAlertsSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_alerts_settings"
}

func (d *NetworkAlertsSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Network Alerts Settings` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"default_destinations_all_admins": schema.BoolAttribute{
				MarkdownDescription: "If true, then all network admins will receive emails.",
				Computed:            true,
			},
			"default_destinations_snmp": schema.BoolAttribute{
				MarkdownDescription: "If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network.",
				Computed:            true,
			},
			"default_destinations_emails": schema.ListAttribute{
				MarkdownDescription: "A list of emails that will receive the alert(s).",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"default_destinations_http_server_ids": schema.ListAttribute{
				MarkdownDescription: "A list of HTTP server IDs to send a Webhook to",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"muting_by_port_schedules_enabled": schema.BoolAttribute{
				MarkdownDescription: "If true, then wireless unreachable alerts will be muted when caused by a port schedule",
				Computed:            true,
			},
			"alerts": schema.ListNestedAttribute{
				MarkdownDescription: "Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							MarkdownDescription: "A boolean depicting if the alert is turned on or off",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "The type of alert",
							Computed:            true,
						},
						"alert_destinations_all_admins": schema.BoolAttribute{
							MarkdownDescription: "If true, then all network admins will receive emails for this alert",
							Computed:            true,
						},
						"alert_destinations_snmp": schema.BoolAttribute{
							MarkdownDescription: "If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network",
							Computed:            true,
						},
						"alert_destinations_emails": schema.ListAttribute{
							MarkdownDescription: "A list of emails that will receive information about the alert",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"alert_destinations_http_server_ids": schema.ListAttribute{
							MarkdownDescription: "A list of HTTP server IDs to send a Webhook to for this alert",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"alert_destinations_sms_numbers": schema.ListAttribute{
							MarkdownDescription: "A list of phone numbers that will receive text messages about the alert. Only available for sensors status alerts.",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"filters_failure_type": schema.StringAttribute{
							MarkdownDescription: "Failure Type",
							Computed:            true,
						},
						"filters_lookback_window": schema.Int64Attribute{
							MarkdownDescription: "Loopback Window (in sec)",
							Computed:            true,
						},
						"filters_min_duration": schema.Int64Attribute{
							MarkdownDescription: "Min Duration",
							Computed:            true,
						},
						"filters_name": schema.StringAttribute{
							MarkdownDescription: "Name",
							Computed:            true,
						},
						"filters_period": schema.Int64Attribute{
							MarkdownDescription: "Period",
							Computed:            true,
						},
						"filters_priority": schema.StringAttribute{
							MarkdownDescription: "Priority",
							Computed:            true,
						},
						"filters_regex": schema.StringAttribute{
							MarkdownDescription: "Regex",
							Computed:            true,
						},
						"filters_selector": schema.StringAttribute{
							MarkdownDescription: "Selector",
							Computed:            true,
						},
						"filters_ssid_num": schema.Int64Attribute{
							MarkdownDescription: "SSID Number",
							Computed:            true,
						},
						"filters_tag": schema.StringAttribute{
							MarkdownDescription: "Tag",
							Computed:            true,
						},
						"filters_threshold": schema.Int64Attribute{
							MarkdownDescription: "Threshold",
							Computed:            true,
						},
						"filters_timeout": schema.Int64Attribute{
							MarkdownDescription: "Timeout",
							Computed:            true,
						},
						"filters_conditions": schema.ListNestedAttribute{
							MarkdownDescription: "Conditions",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"direction": schema.StringAttribute{
										MarkdownDescription: "Direction",
										Computed:            true,
									},
									"duration": schema.Int64Attribute{
										MarkdownDescription: "Duration",
										Computed:            true,
									},
									"threshold": schema.Float64Attribute{
										MarkdownDescription: "Threshold",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of condition",
										Computed:            true,
									},
									"unit": schema.StringAttribute{
										MarkdownDescription: "Unit",
										Computed:            true,
									},
								},
							},
						},
						"filters_serials": schema.SetAttribute{
							MarkdownDescription: "Serials",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworkAlertsSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *NetworkAlertsSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkAlertsSettings

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
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)
	config.Id = config.NetworkId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
