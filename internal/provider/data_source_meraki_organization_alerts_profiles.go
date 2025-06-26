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
	_ datasource.DataSource              = &OrganizationAlertsProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationAlertsProfilesDataSource{}
)

func NewOrganizationAlertsProfilesDataSource() datasource.DataSource {
	return &OrganizationAlertsProfilesDataSource{}
}

type OrganizationAlertsProfilesDataSource struct {
	client *meraki.Client
}

func (d *OrganizationAlertsProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_alerts_profiles"
}

func (d *OrganizationAlertsProfilesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This data source can read the `Organization Alerts Profile` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "Organization ID",
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
						"description": schema.StringAttribute{
							MarkdownDescription: "User supplied description of the alert",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "The alert type",
							Computed:            true,
						},
						"alert_condition_bit_rate_bps": schema.Int64Attribute{
							MarkdownDescription: "The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts.",
							Computed:            true,
						},
						"alert_condition_duration": schema.Int64Attribute{
							MarkdownDescription: "The total duration in seconds that the threshold should be crossed before alerting",
							Computed:            true,
						},
						"alert_condition_interface": schema.StringAttribute{
							MarkdownDescription: "The uplink observed for the alert. interface must be one of the following: wan1, wan2, wan3, cellular",
							Computed:            true,
						},
						"alert_condition_jitter_ms": schema.Int64Attribute{
							MarkdownDescription: "The threshold the metric must cross to be valid for alerting. Used only for VoIP Jitter alerts.",
							Computed:            true,
						},
						"alert_condition_latency_ms": schema.Int64Attribute{
							MarkdownDescription: "The threshold the metric must cross to be valid for alerting. Used only for WAN Latency alerts.",
							Computed:            true,
						},
						"alert_condition_loss_ratio": schema.Float64Attribute{
							MarkdownDescription: "The threshold the metric must cross to be valid for alerting. Used only for Packet Loss alerts.",
							Computed:            true,
						},
						"alert_condition_mos": schema.Float64Attribute{
							MarkdownDescription: "The threshold the metric must drop below to be valid for alerting. Used only for VoIP MOS alerts.",
							Computed:            true,
						},
						"alert_condition_window": schema.Int64Attribute{
							MarkdownDescription: "The look back period in seconds for sensing the alert",
							Computed:            true,
						},
						"recipients_emails": schema.SetAttribute{
							MarkdownDescription: "A list of emails that will receive information about the alert",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"recipients_http_server_ids": schema.SetAttribute{
							MarkdownDescription: "A list base64 encoded urls of webhook endpoints that will receive information about the alert",
							ElementType:         types.StringType,
							Computed:            true,
						},
						"network_tags": schema.SetAttribute{
							MarkdownDescription: "Networks with these tags will be monitored for the alert",
							ElementType:         types.StringType,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationAlertsProfilesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationAlertsProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config DataSourceOrganizationAlertsProfiles

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "OrganizationAlertsProfilesDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "OrganizationAlertsProfilesDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
