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
	"strings"

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
	_ datasource.DataSource              = &WirelessAirMarshalRuleDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessAirMarshalRuleDataSource{}
)

func NewWirelessAirMarshalRuleDataSource() datasource.DataSource {
	return &WirelessAirMarshalRuleDataSource{}
}

type WirelessAirMarshalRuleDataSource struct {
	client *meraki.Client
}

func (d *WirelessAirMarshalRuleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_air_marshal_rule"
}

func (d *WirelessAirMarshalRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Wireless Air Marshal Rule` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "Indicates if this rule will allow, block, or alert.",
				Computed:            true,
			},
			"match_string": schema.StringAttribute{
				MarkdownDescription: "The string used to match.",
				Computed:            true,
			},
			"match_type": schema.StringAttribute{
				MarkdownDescription: "The type of match.",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessAirMarshalRuleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

func (d *WirelessAirMarshalRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessAirMarshalRule

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	networkPath := fmt.Sprintf("/networks/%v", config.NetworkId.ValueString())
	res, err := d.client.Get(networkPath)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network (GET), got error: %s, %s", err, res.String()))
		return
	}
	orgId := res.Get("organizationId").String()

	rulesPath := fmt.Sprintf("/organizations/%v/wireless/airMarshal/rules?networkIds[]=%v", orgId, config.NetworkId.ValueString())
	res, err = d.client.Get(rulesPath)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve rules (GET), got error: %s, %s", err, res.String()))
		return
	}

	if len(res.Get("items").Array()) > 0 {
		res.Get("items").ForEach(func(k, v gjson.Result) bool {
			if config.Id.ValueString() == v.Get("ruleId").String() {
				res = v
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
