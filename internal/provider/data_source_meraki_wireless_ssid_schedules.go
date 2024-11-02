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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessSSIDSchedulesDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessSSIDSchedulesDataSource{}
)

func NewWirelessSSIDSchedulesDataSource() datasource.DataSource {
	return &WirelessSSIDSchedulesDataSource{}
}

type WirelessSSIDSchedulesDataSource struct {
	client *meraki.Client
}

func (d *WirelessSSIDSchedulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssid_schedules"
}

func (d *WirelessSSIDSchedulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Wireless SSID Schedules` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: "Wireless SSID number",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "If true, the SSID outage schedule is enabled.",
				Computed:            true,
			},
			"ranges": schema.ListNestedAttribute{
				MarkdownDescription: "List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"end_day": schema.StringAttribute{
							MarkdownDescription: "Day of when the outage ends. Can be either full day name, or three letter abbreviation",
							Computed:            true,
						},
						"end_time": schema.StringAttribute{
							MarkdownDescription: "24 hour time when the outage ends.",
							Computed:            true,
						},
						"start_day": schema.StringAttribute{
							MarkdownDescription: "Day of when the outage starts. Can be either full day name, or three letter abbreviation.",
							Computed:            true,
						},
						"start_time": schema.StringAttribute{
							MarkdownDescription: "24 hour time when the outage starts.",
							Computed:            true,
						},
					},
				},
			},
			"ranges_in_seconds": schema.ListNestedAttribute{
				MarkdownDescription: "List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"end": schema.Int64Attribute{
							MarkdownDescription: "Seconds since Sunday at midnight when that outage range ends.",
							Computed:            true,
						},
						"start": schema.Int64Attribute{
							MarkdownDescription: "Seconds since Sunday at midnight when the outage range starts.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *WirelessSSIDSchedulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessSSIDSchedulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessSSIDSchedules

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
	config.Id = config.Number

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
