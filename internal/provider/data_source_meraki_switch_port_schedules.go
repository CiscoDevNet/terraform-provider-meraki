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
	_ datasource.DataSource              = &SwitchPortSchedulesDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchPortSchedulesDataSource{}
)

func NewSwitchPortSchedulesDataSource() datasource.DataSource {
	return &SwitchPortSchedulesDataSource{}
}

type SwitchPortSchedulesDataSource struct {
	client *meraki.Client
}

func (d *SwitchPortSchedulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_port_schedules"
}

func (d *SwitchPortSchedulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Switch Port Schedule` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The list of items",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the object",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name for your port schedule. Required",
							Computed:            true,
						},
						"port_schedule_friday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"port_schedule_friday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_friday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_monday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"port_schedule_monday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_monday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_saturday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"port_schedule_saturday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_saturday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_sunday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"port_schedule_sunday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_sunday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_thursday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"port_schedule_thursday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_thursday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_tuesday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"port_schedule_tuesday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_tuesday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_wednesday_active": schema.BoolAttribute{
							MarkdownDescription: "Whether the schedule is active (true) or inactive (false) during the time specified between `from` and `to`. Defaults to true.",
							Computed:            true,
						},
						"port_schedule_wednesday_from": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be less than the time specified in `to`. Defaults to `00:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
						"port_schedule_wednesday_to": schema.StringAttribute{
							MarkdownDescription: "The time, from `00:00` to `24:00`. Must be greater than the time specified in `from`. Defaults to `24:00`. Only 30 minute increments are allowed.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SwitchPortSchedulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchPortSchedulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SwitchPortSchedules

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "SwitchPortSchedulesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "SwitchPortSchedulesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
