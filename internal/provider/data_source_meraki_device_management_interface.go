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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &DeviceManagementInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceManagementInterfaceDataSource{}
)

func NewDeviceManagementInterfaceDataSource() datasource.DataSource {
	return &DeviceManagementInterfaceDataSource{}
}

type DeviceManagementInterfaceDataSource struct {
	client *meraki.Client
}

func (d *DeviceManagementInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_management_interface"
}

func (d *DeviceManagementInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Device Management Interface` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: "Device serial",
				Required:            true,
			},
			"wan1_static_gateway_ip": schema.StringAttribute{
				MarkdownDescription: "The IP of the gateway on the WAN.",
				Computed:            true,
			},
			"wan1_static_ip": schema.StringAttribute{
				MarkdownDescription: "The IP the device should use on the WAN.",
				Computed:            true,
			},
			"wan1_static_subnet_mask": schema.StringAttribute{
				MarkdownDescription: "The subnet mask for the WAN.",
				Computed:            true,
			},
			"wan1_using_static_ip": schema.BoolAttribute{
				MarkdownDescription: "Configure the interface to have static IP settings or use DHCP.",
				Computed:            true,
			},
			"wan1_vlan": schema.Int64Attribute{
				MarkdownDescription: "The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.",
				Computed:            true,
			},
			"wan1_wan_enabled": schema.StringAttribute{
				MarkdownDescription: "Enable or disable the interface (only for MX devices). Valid values are `enabled`, `disabled`, and `not configured`.",
				Computed:            true,
			},
			"wan1_static_dns": schema.ListAttribute{
				MarkdownDescription: "Up to two DNS IPs.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"wan2_static_gateway_ip": schema.StringAttribute{
				MarkdownDescription: "The IP of the gateway on the WAN.",
				Computed:            true,
			},
			"wan2_static_ip": schema.StringAttribute{
				MarkdownDescription: "The IP the device should use on the WAN.",
				Computed:            true,
			},
			"wan2_static_subnet_mask": schema.StringAttribute{
				MarkdownDescription: "The subnet mask for the WAN.",
				Computed:            true,
			},
			"wan2_using_static_ip": schema.BoolAttribute{
				MarkdownDescription: "Configure the interface to have static IP settings or use DHCP.",
				Computed:            true,
			},
			"wan2_vlan": schema.Int64Attribute{
				MarkdownDescription: "The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.",
				Computed:            true,
			},
			"wan2_wan_enabled": schema.StringAttribute{
				MarkdownDescription: "Enable or disable the interface (only for MX devices). Valid values are `enabled`, `disabled`, and `not configured`.",
				Computed:            true,
			},
			"wan2_static_dns": schema.ListAttribute{
				MarkdownDescription: "Up to two DNS IPs.",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *DeviceManagementInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceManagementInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceManagementInterface

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
	config.Id = config.Serial

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
