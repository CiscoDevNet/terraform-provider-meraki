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
	_ datasource.DataSource              = &NetworkClientSplashAuthorizationStatusDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkClientSplashAuthorizationStatusDataSource{}
)

func NewNetworkClientSplashAuthorizationStatusDataSource() datasource.DataSource {
	return &NetworkClientSplashAuthorizationStatusDataSource{}
}

type NetworkClientSplashAuthorizationStatusDataSource struct {
	client *meraki.Client
}

func (d *NetworkClientSplashAuthorizationStatusDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_client_splash_authorization_status"
}

func (d *NetworkClientSplashAuthorizationStatusDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Network Client Splash Authorization Status` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: "Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.",
				Required:            true,
			},
			"ssids_0_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_1_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_10_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_11_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_12_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_13_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_14_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_2_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_3_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_4_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_5_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_6_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_7_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_8_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
			"ssids_9_is_authorized": schema.BoolAttribute{
				MarkdownDescription: "New authorization status for the SSID (true, false).",
				Computed:            true,
			},
		},
	}
}

func (d *NetworkClientSplashAuthorizationStatusDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *NetworkClientSplashAuthorizationStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkClientSplashAuthorizationStatus

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
	config.Id = config.ClientId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
