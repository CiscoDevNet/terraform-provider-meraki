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
	_ datasource.DataSource              = &NetworkWirelessRadioRRMDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkWirelessRadioRRMDataSource{}
)

func NewNetworkWirelessRadioRRMDataSource() datasource.DataSource {
	return &NetworkWirelessRadioRRMDataSource{}
}

type NetworkWirelessRadioRRMDataSource struct {
	client *meraki.Client
}

func (d *NetworkWirelessRadioRRMDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_wireless_radio_rrm"
}

func (d *NetworkWirelessRadioRRMDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Network Wireless Radio RRM` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"ai_enabled": schema.BoolAttribute{
				MarkdownDescription: "Toggle for enabling or disabling AI in a network",
				Computed:            true,
			},
			"busy_hour_minimize_changes_enabled": schema.BoolAttribute{
				MarkdownDescription: "Toggle for enabling or disabling Busy Hour in a network",
				Computed:            true,
			},
			"busy_hour_schedule_mode": schema.StringAttribute{
				MarkdownDescription: "The Busy Hour mode applied to the network when minimizeChanges is enabled. Must be one of `automatic` or `manual`. Automatic busy hour is only available on firmware versions >= MR 27.0",
				Computed:            true,
			},
			"busy_hour_schedule_manual_end": schema.StringAttribute{
				MarkdownDescription: "The hour that Manual Busy Hour ends each day, in the network time zone",
				Computed:            true,
			},
			"busy_hour_schedule_manual_start": schema.StringAttribute{
				MarkdownDescription: "The hour that Manual Busy Hour starts each day, in the network time zone",
				Computed:            true,
			},
			"channel_avoidance_enabled": schema.BoolAttribute{
				MarkdownDescription: "Toggle for enabling or disabling channel avoidance in a network",
				Computed:            true,
			},
			"fra_enabled": schema.BoolAttribute{
				MarkdownDescription: "Toggle to activate or deactivate FRA in a network, contingent on AI-RRM being enabled.",
				Computed:            true,
			},
		},
	}
}

func (d *NetworkWirelessRadioRRMDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

func (d *NetworkWirelessRadioRRMDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceNetworkWirelessRadioRRM

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	networkPath := fmt.Sprintf("/networks/%v", config.NetworkId.ValueString())
	res, err := d.client.Get(networkPath)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network (GET), got error: %s, %s", err, res.String()))
		return
	}
	orgId := res.Get("organizationId").String()

	settingsPath := fmt.Sprintf("/organizations/%v/wireless/radio/rrm/byNetwork?networkIds[]=%v", orgId, config.NetworkId.ValueString())
	res, err = d.client.Get(settingsPath)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve RRM settings (GET), got error: %s, %s", err, res.String()))
		return
	}

	res = meraki.Res{Result: res.Get("items.0")}

	config.fromBody(ctx, res)
	config.Id = config.NetworkId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
