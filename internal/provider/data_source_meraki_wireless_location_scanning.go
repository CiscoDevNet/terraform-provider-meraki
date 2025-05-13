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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessLocationScanningDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessLocationScanningDataSource{}
)

func NewWirelessLocationScanningDataSource() datasource.DataSource {
	return &WirelessLocationScanningDataSource{}
}

type WirelessLocationScanningDataSource struct {
	client *meraki.Client
}

func (d *WirelessLocationScanningDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_location_scanning"
}

func (d *WirelessLocationScanningDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Wireless Location Scanning` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Collect location and scanning analytics",
				Computed:            true,
			},
			"api_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable push API for scanning events, analytics must be enabled",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessLocationScanningDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

func (d *WirelessLocationScanningDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessLocationScanning

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

	getPath := fmt.Sprintf("/organizations/%v/wireless/location/scanning/byNetwork", orgId)
	res, err = d.client.Get(getPath)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve wireless location scanning (GET), got error: %s, %s", err, res.String()))
		return
	}

	if len(res.Get("items").Array()) > 0 {
		res.Get("items").ForEach(func(k, v gjson.Result) bool {
			if config.NetworkId.ValueString() == v.Get("networkId").String() {
				res = meraki.Res{Result: v}
				return false
			}
			return true
		})
	}

	config.fromBody(ctx, res)
	config.Id = config.NetworkId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
