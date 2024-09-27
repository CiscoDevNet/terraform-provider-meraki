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
	_ datasource.DataSource              = &SwitchAlternateManagementInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &SwitchAlternateManagementInterfaceDataSource{}
)

func NewSwitchAlternateManagementInterfaceDataSource() datasource.DataSource {
	return &SwitchAlternateManagementInterfaceDataSource{}
}

type SwitchAlternateManagementInterfaceDataSource struct {
	client *meraki.Client
}

func (d *SwitchAlternateManagementInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_alternate_management_interface"
}

func (d *SwitchAlternateManagementInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Switch Alternate Management Interface` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set",
				Computed:            true,
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: "Alternate management VLAN, must be between 1 and 4094",
				Computed:            true,
			},
			"protocols": schema.SetAttribute{
				MarkdownDescription: "Can be one or more of the following values: `radius`, `snmp` or `syslog`",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"switches": schema.ListNestedAttribute{
				MarkdownDescription: "Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put `switches` in the body when updating template networks. Also, an empty `switches` array will remove all previous assignments",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"alternate_management_ip": schema.StringAttribute{
							MarkdownDescription: "Switch alternative management IP. To remove a prior IP setting, provide an empty string",
							Computed:            true,
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: "Switch gateway must be in IP format. Only and must be specified for Polaris switches",
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: "Switch serial number",
							Computed:            true,
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: "Switch subnet mask must be in IP format. Only and must be specified for Polaris switches",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SwitchAlternateManagementInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *SwitchAlternateManagementInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SwitchAlternateManagementInterface

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
	config.Id = config.NetworkId

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
