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
	_ datasource.DataSource              = &WirelessSSIDFirewallIsolationAllowlistEntryDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessSSIDFirewallIsolationAllowlistEntryDataSource{}
)

func NewWirelessSSIDFirewallIsolationAllowlistEntryDataSource() datasource.DataSource {
	return &WirelessSSIDFirewallIsolationAllowlistEntryDataSource{}
}

type WirelessSSIDFirewallIsolationAllowlistEntryDataSource struct {
	client *meraki.Client
}

func (d *WirelessSSIDFirewallIsolationAllowlistEntryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssid_firewall_isolation_allowlist_entry"
}

func (d *WirelessSSIDFirewallIsolationAllowlistEntryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Wireless SSID Firewall Isolation Allowlist Entry` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of mac address",
				Computed:            true,
			},
			"client_mac": schema.StringAttribute{
				MarkdownDescription: "L2 Isolation mac address",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "The ID of network",
				Computed:            true,
			},
			"ssid_number": schema.Int64Attribute{
				MarkdownDescription: "The number of SSID",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessSSIDFirewallIsolationAllowlistEntryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessSSIDFirewallIsolationAllowlistEntryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessSSIDFirewallIsolationAllowlistEntry

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
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	if len(res.Array()) > 0 {
		res.ForEach(func(k, v gjson.Result) bool {
			if config.Id.ValueString() == v.Get("entryId").String() {
				res = meraki.Res{Result: v}
				return false
			}
			return true
		})
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
