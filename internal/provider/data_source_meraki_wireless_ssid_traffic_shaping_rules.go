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
	_ datasource.DataSource              = &WirelessSSIDTrafficShapingRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessSSIDTrafficShapingRulesDataSource{}
)

func NewWirelessSSIDTrafficShapingRulesDataSource() datasource.DataSource {
	return &WirelessSSIDTrafficShapingRulesDataSource{}
}

type WirelessSSIDTrafficShapingRulesDataSource struct {
	client *meraki.Client
}

func (d *WirelessSSIDTrafficShapingRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssid_traffic_shaping_rules"
}

func (d *WirelessSSIDTrafficShapingRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Wireless SSID Traffic Shaping Rules` configuration.").String,

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
			"default_rules_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network`s traffic shaping page. Note that default rules count against the rule limit of 8.",
				Computed:            true,
			},
			"traffic_shaping_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether traffic shaping rules are applied to clients on your SSID.",
				Computed:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: "An array of traffic shaping rules. Rules are applied in the order that they are specified in. An empty list (or null) means no rules. Note that you are allowed a maximum of 8 rules.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dscp_tag_value": schema.Int64Attribute{
							MarkdownDescription: "The DSCP tag applied by your rule. null means `Do not change DSCP tag`. For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.",
							Computed:            true,
						},
						"pcp_tag_value": schema.Int64Attribute{
							MarkdownDescription: "The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority). null means `Do not set PCP tag`.",
							Computed:            true,
						},
						"per_client_bandwidth_limits_settings": schema.StringAttribute{
							MarkdownDescription: "How bandwidth limits are applied by your rule. Can be one of `network default`, `ignore` or `custom`.",
							Computed:            true,
						},
						"per_client_bandwidth_limits_bandwidth_limits_limit_down": schema.Int64Attribute{
							MarkdownDescription: "The maximum download limit (integer, in Kbps).",
							Computed:            true,
						},
						"per_client_bandwidth_limits_bandwidth_limits_limit_up": schema.Int64Attribute{
							MarkdownDescription: "The maximum upload limit (integer, in Kbps).",
							Computed:            true,
						},
						"definitions": schema.ListNestedAttribute{
							MarkdownDescription: "A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "The type of definition. Can be one of `application`, `applicationCategory`, `host`, `port`, `ipRange` or `localNet`.",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "If 'type' is `host`, `port`, `ipRange` or `localNet`, then 'value' must be a string, matching either a hostname (e.g. 'somesite.com'), a port (e.g. 8080), or an IP range ('192.1.0.0', '192.1.0.0/16', or '10.1.0.0/16:80'). `localNet` also supports CIDR notation, excluding custom ports.",
										Computed:            true,
									},
									"application_value": schema.StringAttribute{
										MarkdownDescription: "If 'type' is `application` or `applicationCategory`, then 'application_value' must be an application category or application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories endpoint).",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *WirelessSSIDTrafficShapingRulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessSSIDTrafficShapingRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessSSIDTrafficShapingRules

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
