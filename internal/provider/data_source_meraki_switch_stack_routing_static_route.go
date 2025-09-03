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
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SwitchStackRoutingStaticRouteDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchStackRoutingStaticRouteDataSource{}
)

func NewSwitchStackRoutingStaticRouteDataSource() datasource.DataSource {
	return &SwitchStackRoutingStaticRouteDataSource{}
}

type SwitchStackRoutingStaticRouteDataSource struct {
	client *meraki.Client
}

func (d *SwitchStackRoutingStaticRouteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_stack_routing_static_route"
}

func (d *SwitchStackRoutingStaticRouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Switch Stack Routing Static Route` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Optional:            true,
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: "Switch stack ID",
				Required:            true,
			},
			"advertise_via_ospf_enabled": schema.BoolAttribute{
				MarkdownDescription: "Option to advertise static route via OSPF",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Name or description for layer 3 static route",
				Optional:            true,
				Computed:            true,
			},
			"next_hop_ip": schema.StringAttribute{
				MarkdownDescription: "IP address of the next hop device to which the device sends its traffic for the subnet",
				Computed:            true,
			},
			"prefer_over_ospf_routes_enabled": schema.BoolAttribute{
				MarkdownDescription: "Option to prefer static route over OSPF routes",
				Computed:            true,
			},
			"subnet": schema.StringAttribute{
				MarkdownDescription: "The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)",
				Computed:            true,
			},
			"vrf_leak_route_to_default_vrf": schema.BoolAttribute{
				MarkdownDescription: "Whether or not next-hop IP is reachable via default VRF",
				Computed:            true,
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: "The name of the VRF this static route belongs to",
				Optional:            true,
				Computed:            true,
			},
		},
	}
}
func (d *SwitchStackRoutingStaticRouteDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *SwitchStackRoutingStaticRouteDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchStackRoutingStaticRouteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SwitchStackRoutingStaticRoute

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res meraki.Res
	var err error
	if config.Id.IsNull() && !config.Name.IsNull() {
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve objects, got error: %s", err))
			return
		}
		found := false
		if len(res.Array()) > 0 {
			res.ForEach(func(k, v gjson.Result) bool {
				if config.Name.ValueString() == v.Get("name").String() {
					if found {
						resp.Diagnostics.AddWarning("Multiple objects with same name", fmt.Sprintf("Found multiple objects with name: %s", config.Name.ValueString()))
						return false
					}
					config.Id = types.StringValue(v.Get("staticRouteId").String())
					tflog.Debug(ctx, fmt.Sprintf("%s: Found object with name '%v', id: %v", config.Id.String(), config.Name.ValueString(), config.Id.String()))
					res = meraki.Res{Result: v}
					found = true
				}
				return true
			})
		}

		if config.Id.IsNull() {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to find object with name: %s", config.Name.ValueString()))
			return
		}
	}

	if !res.Exists() {
		res, err = d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
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
