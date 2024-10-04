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
	_ datasource.DataSource              = &ApplianceNetworkSecurityIntrusionDataSource{}
	_ datasource.DataSourceWithConfigure = &ApplianceNetworkSecurityIntrusionDataSource{}
)

func NewApplianceNetworkSecurityIntrusionDataSource() datasource.DataSource {
	return &ApplianceNetworkSecurityIntrusionDataSource{}
}

type ApplianceNetworkSecurityIntrusionDataSource struct {
	client *meraki.Client
}

func (d *ApplianceNetworkSecurityIntrusionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_network_security_intrusion"
}

func (d *ApplianceNetworkSecurityIntrusionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Appliance Network Security Intrusion` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"ids_rulesets": schema.StringAttribute{
				MarkdownDescription: "Set the detection ruleset `connectivity`/`balanced`/`security` (optional - omitting will leave current config unchanged). Default value is `balanced` if none currently saved",
				Computed:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "Set mode to `disabled`/`detection`/`prevention` (optional - omitting will leave current config unchanged)",
				Computed:            true,
			},
			"protected_networks_use_default": schema.BoolAttribute{
				MarkdownDescription: "true/false whether to use special IPv4 addresses: https://tools.ietf.org/html/rfc5735 (required). Default value is true if none currently saved",
				Computed:            true,
			},
			"protected_networks_excluded_cidr": schema.ListAttribute{
				MarkdownDescription: "list of IP addresses or subnets being excluded from protection (required if `useDefault` is false)",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"protected_networks_included_cidr": schema.ListAttribute{
				MarkdownDescription: "list of IP addresses or subnets being protected (required if `useDefault` is false)",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *ApplianceNetworkSecurityIntrusionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *ApplianceNetworkSecurityIntrusionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ApplianceNetworkSecurityIntrusion

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
