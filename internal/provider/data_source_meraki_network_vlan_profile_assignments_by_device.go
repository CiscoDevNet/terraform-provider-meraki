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
	_ datasource.DataSource              = &NetworkVLANProfileAssignmentsByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkVLANProfileAssignmentsByDeviceDataSource{}
)

func NewNetworkVLANProfileAssignmentsByDeviceDataSource() datasource.DataSource {
	return &NetworkVLANProfileAssignmentsByDeviceDataSource{}
}

type NetworkVLANProfileAssignmentsByDeviceDataSource struct {
	client *meraki.Client
}

func (d *NetworkVLANProfileAssignmentsByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_vlan_profile_assignments_by_device"
}

func (d *NetworkVLANProfileAssignmentsByDeviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Network VLAN Profile Assignments By Device` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The list of items",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The id of the object",
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: "MAC address of the device",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the Device",
							Computed:            true,
						},
						"product_type": schema.StringAttribute{
							MarkdownDescription: "The product type",
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: "Serial of the Device",
							Computed:            true,
						},
						"stack_id": schema.StringAttribute{
							MarkdownDescription: "ID of the Switch Stack",
							Computed:            true,
						},
						"vlan_profile_iname": schema.StringAttribute{
							MarkdownDescription: "IName of the VLAN Profile",
							Computed:            true,
						},
						"vlan_profile_is_default": schema.BoolAttribute{
							MarkdownDescription: "Is this VLAN profile the default for the network?",
							Computed:            true,
						},
						"vlan_profile_name": schema.StringAttribute{
							MarkdownDescription: "Name of the VLAN Profile",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworkVLANProfileAssignmentsByDeviceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *NetworkVLANProfileAssignmentsByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkVLANProfileAssignmentsByDevice

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "NetworkVLANProfileAssignmentsByDeviceDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "NetworkVLANProfileAssignmentsByDeviceDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
