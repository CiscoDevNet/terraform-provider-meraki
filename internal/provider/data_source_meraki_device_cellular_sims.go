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
	_ datasource.DataSource              = &DeviceCellularSIMsDataSource{}
	_ datasource.DataSourceWithConfigure = &DeviceCellularSIMsDataSource{}
)

func NewDeviceCellularSIMsDataSource() datasource.DataSource {
	return &DeviceCellularSIMsDataSource{}
}

type DeviceCellularSIMsDataSource struct {
	client *meraki.Client
}

func (d *DeviceCellularSIMsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_cellular_sims"
}

func (d *DeviceCellularSIMsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Device Cellular SIMs` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: "Device serial",
				Required:            true,
			},
			"sim_failover_enabled": schema.BoolAttribute{
				MarkdownDescription: "Failover to secondary SIM (optional)",
				Computed:            true,
			},
			"sim_failover_timeout": schema.Int64Attribute{
				MarkdownDescription: "Failover timeout in seconds (optional)",
				Computed:            true,
			},
			"sim_ordering": schema.ListAttribute{
				MarkdownDescription: "Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It`s required for devices with 3 or more SIMs and can be used in place of `isPrimary` for dual-SIM devices. To indicate eSIM, use `sim3`. Sim failover will occur only between primary and secondary sim slots.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"sims": schema.ListNestedAttribute{
				MarkdownDescription: "List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged.",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"is_primary": schema.BoolAttribute{
							MarkdownDescription: "If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using `simOrdering`.",
							Computed:            true,
						},
						"sim_order": schema.Int64Attribute{
							MarkdownDescription: "Priority of SIM slot being configured. Use a value between 1 and total number of SIMs available. The value must be unique for each SIM.",
							Computed:            true,
						},
						"slot": schema.StringAttribute{
							MarkdownDescription: "SIM slot being configured. Must be `sim1` on single-sim devices. Use `sim3` for eSIM.",
							Computed:            true,
						},
						"apns": schema.ListNestedAttribute{
							MarkdownDescription: "APN configurations. If empty, the default APN will be used.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "APN name.",
										Computed:            true,
									},
									"authentication_password": schema.StringAttribute{
										MarkdownDescription: "APN password, if type is set (if APN password is not supplied, the password is left unchanged).",
										Computed:            true,
									},
									"authentication_type": schema.StringAttribute{
										MarkdownDescription: "APN auth type.",
										Computed:            true,
									},
									"authentication_username": schema.StringAttribute{
										MarkdownDescription: "APN username, if type is set.",
										Computed:            true,
									},
									"allowed_ip_types": schema.ListAttribute{
										MarkdownDescription: "IP versions to support (permitted values include `ipv4`, `ipv6`).",
										ElementType:         types.StringType,
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

func (d *DeviceCellularSIMsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *DeviceCellularSIMsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DeviceCellularSIMs

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
	config.Id = config.Serial

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
