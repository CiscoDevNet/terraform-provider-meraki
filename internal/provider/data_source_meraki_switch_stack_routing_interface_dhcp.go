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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SwitchStackRoutingInterfaceDHCPDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchStackRoutingInterfaceDHCPDataSource{}
)

func NewSwitchStackRoutingInterfaceDHCPDataSource() datasource.DataSource {
	return &SwitchStackRoutingInterfaceDHCPDataSource{}
}

type SwitchStackRoutingInterfaceDHCPDataSource struct {
	client *meraki.Client
}

func (d *SwitchStackRoutingInterfaceDHCPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_stack_routing_interface_dhcp"
}

func (d *SwitchStackRoutingInterfaceDHCPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Switch Stack Routing Interface DHCP` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
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
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "Interface ID",
				Required:            true,
			},
			"boot_file_name": schema.StringAttribute{
				MarkdownDescription: "The PXE boot server file name for the DHCP server running on the switch stack interface",
				Computed:            true,
			},
			"boot_next_server": schema.StringAttribute{
				MarkdownDescription: "The PXE boot server IP for the DHCP server running on the switch stack interface",
				Computed:            true,
			},
			"boot_options_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface",
				Computed:            true,
			},
			"dhcp_lease_time": schema.StringAttribute{
				MarkdownDescription: "The DHCP lease time config for the dhcp server running on switch stack interface (`30 minutes`, `1 hour`, `4 hours`, `12 hours`, `1 day` or `1 week`)",
				Computed:            true,
			},
			"dhcp_mode": schema.StringAttribute{
				MarkdownDescription: "The DHCP mode options for the switch stack interface (`dhcpDisabled`, `dhcpRelay` or `dhcpServer`)",
				Computed:            true,
			},
			"dns_nameservers_option": schema.StringAttribute{
				MarkdownDescription: "The DHCP name server option for the dhcp server running on the switch stack interface (`googlePublicDns`, `openDns` or `custom`)",
				Computed:            true,
			},
			"dhcp_options": schema.ListNestedAttribute{
				MarkdownDescription: "Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"code": schema.StringAttribute{
							MarkdownDescription: "The code for DHCP option which should be from 2 to 254",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "The type of the DHCP option which should be one of (`text`, `ip`, `integer` or `hex`)",
							Computed:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "The value of the DHCP option",
							Computed:            true,
						},
					},
				},
			},
			"dhcp_relay_server_ips": schema.SetAttribute{
				MarkdownDescription: "The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"dns_custom_nameservers": schema.ListAttribute{
				MarkdownDescription: "The DHCP name server IPs when DHCP name server option is ` custom`",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"fixed_ip_assignments": schema.ListNestedAttribute{
				MarkdownDescription: "Array of DHCP fixed IP assignments for the DHCP server running on the switch stack interface",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							MarkdownDescription: "The IP address of the client which has fixed IP address assigned to it",
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: "The MAC address of the client which has fixed IP address",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the client which has fixed IP address",
							Computed:            true,
						},
					},
				},
			},
			"reserved_ip_ranges": schema.ListNestedAttribute{
				MarkdownDescription: "Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"comment": schema.StringAttribute{
							MarkdownDescription: "The comment for the reserved IP range",
							Computed:            true,
						},
						"end": schema.StringAttribute{
							MarkdownDescription: "The ending IP address of the reserved IP range",
							Computed:            true,
						},
						"start": schema.StringAttribute{
							MarkdownDescription: "The starting IP address of the reserved IP range",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SwitchStackRoutingInterfaceDHCPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchStackRoutingInterfaceDHCPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SwitchStackRoutingInterfaceDHCP

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res gjson.Result
	var err error

	if !res.Exists() {
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)
	config.Id = config.InterfaceId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
