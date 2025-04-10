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
	_ datasource.DataSource              = &SwitchLinkAggregationDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchLinkAggregationDataSource{}
)

func NewSwitchLinkAggregationDataSource() datasource.DataSource {
	return &SwitchLinkAggregationDataSource{}
}

type SwitchLinkAggregationDataSource struct {
	client *meraki.Client
}

func (d *SwitchLinkAggregationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_link_aggregation"
}

func (d *SwitchLinkAggregationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Switch Link Aggregation` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"switch_ports": schema.ListNestedAttribute{
				MarkdownDescription: "Array of switch or stack ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_id": schema.StringAttribute{
							MarkdownDescription: "Port identifier of switch port. For modules, the identifier is 'SlotNumber_ModuleType_PortNumber' (Ex: '1_8X10G_1'), otherwise it is just the port number (Ex: '8').",
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: "Serial number of the switch.",
							Computed:            true,
						},
					},
				},
			},
			"switch_profile_ports": schema.ListNestedAttribute{
				MarkdownDescription: "Array of switch profile ports for creating aggregation group. Minimum 2 and maximum 8 ports are supported.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_id": schema.StringAttribute{
							MarkdownDescription: "Port identifier of switch port. For modules, the identifier is 'SlotNumber_ModuleType_PortNumber' (Ex: '1_8X10G_1'), otherwise it is just the port number (Ex: '8').",
							Computed:            true,
						},
						"profile": schema.StringAttribute{
							MarkdownDescription: "Profile identifier.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SwitchLinkAggregationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchLinkAggregationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SwitchLinkAggregation

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
			if config.Id.ValueString() == v.Get("id").String() {
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
