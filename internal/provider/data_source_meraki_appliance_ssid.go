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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ApplianceSSIDDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceSSIDDataSource{}
)

func NewApplianceSSIDDataSource() datasource.DataSource {
	return &ApplianceSSIDDataSource{}
}

type ApplianceSSIDDataSource struct {
	client *meraki.Client
}

func (d *ApplianceSSIDDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_ssid"
}

func (d *ApplianceSSIDDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Appliance SSID` configuration.",

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
			"auth_mode": schema.StringAttribute{
				MarkdownDescription: "The association control method for the SSID (`open`, `psk`, `8021x-meraki` or `8021x-radius`).",
				Computed:            true,
			},
			"default_vlan_id": schema.Int64Attribute{
				MarkdownDescription: "The VLAN ID of the VLAN associated to this SSID. This parameter is only valid if the network is in routed mode.",
				Computed:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not the SSID is enabled.",
				Computed:            true,
			},
			"encryption_mode": schema.StringAttribute{
				MarkdownDescription: "The psk encryption mode for the SSID (`wep` or `wpa`). This param is only valid if the authMode is `psk`.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the SSID.",
				Computed:            true,
			},
			"psk": schema.StringAttribute{
				MarkdownDescription: "The passkey for the SSID. This param is only valid if the authMode is `psk`.",
				Computed:            true,
			},
			"visible": schema.BoolAttribute{
				MarkdownDescription: "Boolean indicating whether the MX should advertise or hide this SSID.",
				Computed:            true,
			},
			"wpa_encryption_mode": schema.StringAttribute{
				MarkdownDescription: "The types of WPA encryption. (`WPA1 and WPA2`, `WPA2 only`, `WPA3 Transition Mode` or `WPA3 only`). This param is only valid if (1) the authMode is `psk` & the encryptionMode is `wpa` OR (2) the authMode is `8021x-meraki` OR (3) the authMode is `8021x-radius`",
				Computed:            true,
			},
			"dhcp_enforced_deauthentication_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable DCHP Enforced Deauthentication on the SSID.",
				Computed:            true,
			},
			"dot11w_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether 802.11w is enabled or not.",
				Computed:            true,
			},
			"dot11w_required": schema.BoolAttribute{
				MarkdownDescription: "(Optional) Whether 802.11w is required or not.",
				Computed:            true,
			},
			"radius_servers": schema.ListNestedAttribute{
				MarkdownDescription: "The RADIUS 802.1x servers to be used for authentication. This param is only valid if the authMode is `8021x-radius`.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host": schema.StringAttribute{
							MarkdownDescription: "The IP address of your RADIUS server.",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "The UDP port your RADIUS servers listens on for Access-requests.",
							Computed:            true,
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: "The RADIUS client shared secret.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ApplianceSSIDDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceSSIDDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplianceSSID

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
