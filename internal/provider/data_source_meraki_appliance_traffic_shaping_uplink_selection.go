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
	_ datasource.DataSource              = &ApplianceTrafficShapingUplinkSelectionDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceTrafficShapingUplinkSelectionDataSource{}
)

func NewApplianceTrafficShapingUplinkSelectionDataSource() datasource.DataSource {
	return &ApplianceTrafficShapingUplinkSelectionDataSource{}
}

type ApplianceTrafficShapingUplinkSelectionDataSource struct {
	client *meraki.Client
}

func (d *ApplianceTrafficShapingUplinkSelectionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_traffic_shaping_uplink_selection"
}

func (d *ApplianceTrafficShapingUplinkSelectionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Appliance Traffic Shaping Uplink Selection` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"active_active_auto_vpn_enabled": schema.BoolAttribute{
				MarkdownDescription: "Toggle for enabling or disabling active-active AutoVPN",
				Computed:            true,
			},
			"default_uplink": schema.StringAttribute{
				MarkdownDescription: "The default uplink. Must be a WAN interface `wanX`",
				Computed:            true,
			},
			"load_balancing_enabled": schema.BoolAttribute{
				MarkdownDescription: "Toggle for enabling or disabling load balancing",
				Computed:            true,
			},
			"failover_and_failback_immediate_enabled": schema.BoolAttribute{
				MarkdownDescription: "Toggle for enabling or disabling immediate WAN failover and failback",
				Computed:            true,
			},
			"vpn_traffic_uplink_preferences": schema.ListNestedAttribute{
				MarkdownDescription: "Array of uplink preference rules for VPN traffic",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"fail_over_criterion": schema.StringAttribute{
							MarkdownDescription: "Fail over criterion for this uplink preference rule. Must be one of: `poorPerformance` or `uplinkDown`",
							Computed:            true,
						},
						"preferred_uplink": schema.StringAttribute{
							MarkdownDescription: "Preferred uplink for uplink preference rule. Must be one of: `wan1`, `wan2`, `bestForVoIP`, `loadBalancing` or `defaultUplink`, or any other valid uplink(`wanX`) if it applies to the network",
							Computed:            true,
						},
						"builtin_performance_class_name": schema.StringAttribute{
							MarkdownDescription: "Name of builtin performance class, must be present when performanceClass type is `builtin`, and value must be one of: `VoIP`",
							Computed:            true,
						},
						"custom_performance_class_id": schema.StringAttribute{
							MarkdownDescription: "ID of created custom performance class, must be present when performanceClass type is `custom`",
							Computed:            true,
						},
						"performance_class_type": schema.StringAttribute{
							MarkdownDescription: "Type of this performance class. Must be one of: `builtin` or `custom`",
							Computed:            true,
						},
						"traffic_filters": schema.ListNestedAttribute{
							MarkdownDescription: "Array of traffic filters for this uplink preference rule",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of this traffic filter. Must be one of: `applicationCategory`, `application` or `custom`",
										Computed:            true,
									},
									"id": schema.StringAttribute{
										MarkdownDescription: "ID of this applicationCategory or application type traffic filter. E.g.: 'meraki:layer7/category/1', 'meraki:layer7/application/4'",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Protocol of this custom type traffic filter. Must be one of: `tcp`, `udp`, `icmp`, `icmp6` or `any`",
										Computed:            true,
									},
									"destination_cidr": schema.StringAttribute{
										MarkdownDescription: "CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')",
										Computed:            true,
									},
									"destination_fqdn": schema.StringAttribute{
										MarkdownDescription: "FQDN format address. Currently only availabe in `destination` of `vpnTrafficUplinkPreference` object. E.g.: `www.google.com`",
										Computed:            true,
									},
									"destination_host": schema.Int64Attribute{
										MarkdownDescription: "Host ID in the VLAN, should be used along with `vlan`, and not exceed the vlan subnet capacity. Currently only available under a template network.",
										Computed:            true,
									},
									"destination_network": schema.StringAttribute{
										MarkdownDescription: "Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: 'L_12345678'.",
										Computed:            true,
									},
									"destination_port": schema.StringAttribute{
										MarkdownDescription: "E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'",
										Computed:            true,
									},
									"destination_vlan": schema.Int64Attribute{
										MarkdownDescription: "VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.",
										Computed:            true,
									},
									"source_cidr": schema.StringAttribute{
										MarkdownDescription: "CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')",
										Computed:            true,
									},
									"source_host": schema.Int64Attribute{
										MarkdownDescription: "Host ID in the VLAN, should be used along with `vlan`, and not exceed the vlan subnet capacity. Currently only available under a template network.",
										Computed:            true,
									},
									"source_network": schema.StringAttribute{
										MarkdownDescription: "Meraki network ID. Currently only available under a template network, and the value should be ID of either same template network, or another template network currently. E.g.: 'L_12345678'.",
										Computed:            true,
									},
									"source_port": schema.StringAttribute{
										MarkdownDescription: "E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'",
										Computed:            true,
									},
									"source_vlan": schema.Int64Attribute{
										MarkdownDescription: "VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"wan_traffic_uplink_preferences": schema.ListNestedAttribute{
				MarkdownDescription: "Array of uplink preference rules for WAN traffic. Note: these preferences are shared (merged) with meraki_appliance_sdwan_internet_policies resource. It is recommended to only use one resource for these preferences, not both at the same time: Deleting this resource clears preferences created in meraki_appliance_sdwan_internet_policies, which isn't detected as a change by the provider. The same happens the other way around, but the change is detected in uplink_selection on a subsequent apply.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"preferred_uplink": schema.StringAttribute{
							MarkdownDescription: "Preferred uplink for uplink preference rule. Must be one of: `wan1` or `wan2`, or any other valid uplink(`wanX`) if it applies to the network",
							Computed:            true,
						},
						"traffic_filters": schema.ListNestedAttribute{
							MarkdownDescription: "Array of traffic filters for this uplink preference rule",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										MarkdownDescription: "Type of this traffic filter. Must be one of: `custom`",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Protocol of this custom type traffic filter. Must be one of: `tcp`, `udp`, `icmp6` or `any`",
										Computed:            true,
									},
									"destination_cidr": schema.StringAttribute{
										MarkdownDescription: "CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')",
										Computed:            true,
									},
									"destination_port": schema.StringAttribute{
										MarkdownDescription: "E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'",
										Computed:            true,
									},
									"source_cidr": schema.StringAttribute{
										MarkdownDescription: "CIDR format address, or 'any'. E.g.: '192.168.10.0/24', '192.168.10.1' (same as '192.168.10.1/32'), '0.0.0.0/0' (same as 'any')",
										Computed:            true,
									},
									"source_host": schema.Int64Attribute{
										MarkdownDescription: "Host ID in the VLAN, should be used along with `vlan`, and not exceed the vlan subnet capacity. Currently only available under a template network.",
										Computed:            true,
									},
									"source_port": schema.StringAttribute{
										MarkdownDescription: "E.g.: 'any', '0' (also means 'any'), '8080', '1-1024'",
										Computed:            true,
									},
									"source_vlan": schema.Int64Attribute{
										MarkdownDescription: "VLAN ID of the configured VLAN in the Meraki network. Currently only available under a template network.",
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

func (d *ApplianceTrafficShapingUplinkSelectionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceTrafficShapingUplinkSelectionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplianceTrafficShapingUplinkSelection

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
	config.Id = config.NetworkId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
