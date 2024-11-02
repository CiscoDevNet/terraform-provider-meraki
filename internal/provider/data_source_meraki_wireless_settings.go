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
	_ datasource.DataSource              = &WirelessSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessSettingsDataSource{}
)

func NewWirelessSettingsDataSource() datasource.DataSource {
	return &WirelessSettingsDataSource{}
}

type WirelessSettingsDataSource struct {
	client *meraki.Client
}

func (d *WirelessSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_settings"
}

func (d *WirelessSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Wireless settings` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"ipv6_bridge_enabled": schema.BoolAttribute{
				MarkdownDescription: "Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)",
				Computed:            true,
			},
			"led_lights_on": schema.BoolAttribute{
				MarkdownDescription: "Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)",
				Computed:            true,
			},
			"location_analytics_enabled": schema.BoolAttribute{
				MarkdownDescription: "Toggle for enabling or disabling location analytics for your network",
				Computed:            true,
			},
			"meshing_enabled": schema.BoolAttribute{
				MarkdownDescription: "Toggle for enabling or disabling meshing in a network",
				Computed:            true,
			},
			"upgrade_strategy": schema.StringAttribute{
				MarkdownDescription: "The default strategy that network devices will use to perform an upgrade. Requires firmware version MR 26.8 or higher.",
				Computed:            true,
			},
			"named_vlans_pool_dhcp_monitoring_duration": schema.Int64Attribute{
				MarkdownDescription: "The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.",
				Computed:            true,
			},
			"named_vlans_pool_dhcp_monitoring_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP.",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessSettings

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
