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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceDataSource{}
)

func NewDeviceDataSource() datasource.DataSource {
	return &DeviceDataSource{}
}

type DeviceDataSource struct {
	client *meraki.Client
}

func (d *DeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (d *DeviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Device` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: "Switch serial",
				Required:            true,
			},
			"address": schema.StringAttribute{
				MarkdownDescription: "The address of a device",
				Computed:            true,
			},
			"floor_plan_id": schema.StringAttribute{
				MarkdownDescription: "The floor plan to associate to this device. null disassociates the device from the floorplan.",
				Computed:            true,
			},
			"lat": schema.Float64Attribute{
				MarkdownDescription: "The latitude of a device",
				Computed:            true,
			},
			"lng": schema.Float64Attribute{
				MarkdownDescription: "The longitude of a device",
				Computed:            true,
			},
			"move_map_marker": schema.BoolAttribute{
				MarkdownDescription: "Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of a device",
				Computed:            true,
			},
			"notes": schema.StringAttribute{
				MarkdownDescription: "The notes for the device. String. Limited to 255 characters.",
				Computed:            true,
			},
			"switch_profile_id": schema.StringAttribute{
				MarkdownDescription: "The ID of a switch template to bind to the device (for available switch templates, see the `Switch Templates` endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch template, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.",
				Computed:            true,
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: "The list of tags of a device",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *DeviceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Device

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res gjson.Result
	var err error

	if !res.Exists() {
		res, err = d.client.Get(config.getPath())
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
