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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ApplianceVLANsDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceVLANsDataSource{}
)

func NewApplianceVLANsDataSource() datasource.DataSource {
	return &ApplianceVLANsDataSource{}
}

type ApplianceVLANsDataSource struct {
	client *meraki.Client
}

func (d *ApplianceVLANsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_vlans"
}

func (d *ApplianceVLANsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Appliance VLAN` configuration in bulk.").String,

		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"items": schema.SetNestedAttribute{
				MarkdownDescription: "The list of items",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the object",
							Computed:            true,
						},
						"appliance_ip": schema.StringAttribute{
							MarkdownDescription: "The local IP of the appliance on the VLAN",
							Computed:            true,
						},
						"cidr": schema.StringAttribute{
							MarkdownDescription: "CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.",
							Computed:            true,
						},
						"dhcp_boot_options_enabled": schema.BoolAttribute{
							MarkdownDescription: "Use DHCP boot options specified in other properties",
							Computed:            true,
						},
						"dhcp_handling": schema.StringAttribute{
							MarkdownDescription: "The appliance`s handling of DHCP requests on this VLAN. One of: `Run a DHCP server`, `Relay DHCP to another server` or `Do not respond to DHCP requests`",
							Computed:            true,
						},
						"dhcp_lease_time": schema.StringAttribute{
							MarkdownDescription: "The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: `30 minutes`, `1 hour`, `4 hours`, `12 hours`, `1 day` or `1 week`",
							Computed:            true,
						},
						"group_policy_id": schema.StringAttribute{
							MarkdownDescription: "The id of the desired group policy to apply to the VLAN",
							Computed:            true,
						},
						"vlan_id": schema.StringAttribute{
							MarkdownDescription: "The VLAN ID of the new VLAN (must be between 1 and 4094)",
							Computed:            true,
						},
						"mask": schema.Int64Attribute{
							MarkdownDescription: "Mask used for the subnet of all bound to the template networks. Applicable only for template network.",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the new VLAN",
							Computed:            true,
						},
						"subnet": schema.StringAttribute{
							MarkdownDescription: "The subnet of the VLAN",
							Computed:            true,
						},
						"template_vlan_type": schema.StringAttribute{
							MarkdownDescription: "Type of subnetting of the VLAN. Applicable only for template network.",
							Computed:            true,
						},
						"ipv6_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable IPv6 on VLAN.",
							Computed:            true,
						},
						"ipv6_prefix_assignments": schema.ListNestedAttribute{
							MarkdownDescription: "Prefix assignments on the VLAN",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"autonomous": schema.BoolAttribute{
										MarkdownDescription: "Auto assign a /64 prefix from the origin to the VLAN",
										Computed:            true,
									},
									"static_appliance_ip6": schema.StringAttribute{
										MarkdownDescription: "Manual configuration of the IPv6 Appliance IP",
										Computed:            true,
									},
									"static_prefix": schema.StringAttribute{
										MarkdownDescription: "Manual configuration of a /64 prefix on the VLAN",
										Computed:            true,
									},
									"origin_type": schema.StringAttribute{
										MarkdownDescription: "Type of the origin",
										Computed:            true,
									},
									"origin_interfaces": schema.ListAttribute{
										MarkdownDescription: "Interfaces associated with the prefix",
										ElementType:         types.StringType,
										Computed:            true,
									},
								},
							},
						},
						"mandatory_dhcp_enabled": schema.BoolAttribute{
							MarkdownDescription: "Enable Mandatory DHCP on VLAN.",
							Computed:            true,
						},
						"dhcp_options": schema.ListNestedAttribute{
							MarkdownDescription: "The list of DHCP options that will be included in DHCP responses. Each object in the list should have 'code', 'type', and 'value' properties.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"code": schema.StringAttribute{
										MarkdownDescription: "The code for the DHCP option. This should be an integer between 2 and 254.",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "The type for the DHCP option. One of: `text`, `ip`, `hex` or `integer`",
										Computed:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: "The value for the DHCP option",
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

func (d *ApplianceVLANsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceVLANsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceApplianceVLANs

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "ApplianceVLANsDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "ApplianceVLANsDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
