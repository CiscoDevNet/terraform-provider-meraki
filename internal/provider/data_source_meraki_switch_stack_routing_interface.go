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
	_ datasource.DataSource              = &SwitchStackRoutingInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchStackRoutingInterfaceDataSource{}
)

func NewSwitchStackRoutingInterfaceDataSource() datasource.DataSource {
	return &SwitchStackRoutingInterfaceDataSource{}
}

type SwitchStackRoutingInterfaceDataSource struct {
	client *meraki.Client
}

func (d *SwitchStackRoutingInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_stack_routing_interface"
}

func (d *SwitchStackRoutingInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Switch Stack Routing Interface` configuration.").String,

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
			"default_gateway": schema.StringAttribute{
				MarkdownDescription: "The next hop for any traffic that isn`t going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface.",
				Computed:            true,
			},
			"interface_ip": schema.StringAttribute{
				MarkdownDescription: "The IP address this switch stack will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch`s management IP.",
				Computed:            true,
			},
			"multicast_routing": schema.StringAttribute{
				MarkdownDescription: "Enable multicast support if, multicast routing between VLANs is required. Options are, `disabled`, `enabled` or `IGMP snooping querier`. Default is `disabled`.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "A friendly name or description for the interface or VLAN.",
				Optional:            true,
				Computed:            true,
			},
			"subnet": schema.StringAttribute{
				MarkdownDescription: "The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).",
				Computed:            true,
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: "The VLAN this routed interface is on. VLAN must be between 1 and 4094.",
				Computed:            true,
			},
			"ipv6_address": schema.StringAttribute{
				MarkdownDescription: "The IPv6 address of the interface. Required if assignmentMode is `static`. Must not be included if assignmentMode is `eui-64`.",
				Computed:            true,
			},
			"ipv6_assignment_mode": schema.StringAttribute{
				MarkdownDescription: "The IPv6 assignment mode for the interface. Can be either `eui-64` or `static`.",
				Computed:            true,
			},
			"ipv6_gateway": schema.StringAttribute{
				MarkdownDescription: "The IPv6 default gateway of the interface. Required if prefix is defined and this is the first interface with IPv6 configured for the stack.",
				Computed:            true,
			},
			"ipv6_prefix": schema.StringAttribute{
				MarkdownDescription: "The IPv6 prefix of the interface. Required if IPv6 object is included.",
				Computed:            true,
			},
			"ospf_settings_area": schema.StringAttribute{
				MarkdownDescription: "The OSPF area to which this interface should belong. Can be either `ospfDisabled` or the identifier of an existing OSPF area. Defaults to `ospfDisabled`.",
				Computed:            true,
			},
			"ospf_settings_cost": schema.Int64Attribute{
				MarkdownDescription: "The path cost for this interface. Defaults to 1, but can be increased up to 65535 to give lower priority.",
				Computed:            true,
			},
			"ospf_settings_is_passive_enabled": schema.BoolAttribute{
				MarkdownDescription: "When enabled, OSPF will not run on the interface, but the subnet will still be advertised.",
				Computed:            true,
			},
		},
	}
}
func (d *SwitchStackRoutingInterfaceDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
			path.MatchRoot("name"),
		),
	}
}

func (d *SwitchStackRoutingInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchStackRoutingInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SwitchStackRoutingInterface

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
					}
					config.Id = types.StringValue(v.Get("interfaceId").String())
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
