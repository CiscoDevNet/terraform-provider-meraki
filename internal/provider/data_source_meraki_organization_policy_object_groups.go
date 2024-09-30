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
	_ datasource.DataSource              = &OrganizationPolicyObjectGroupsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationPolicyObjectGroupsDataSource{}
)

func NewOrganizationPolicyObjectGroupsDataSource() datasource.DataSource {
	return &OrganizationPolicyObjectGroupsDataSource{}
}

type OrganizationPolicyObjectGroupsDataSource struct {
	client *meraki.Client
}

func (d *OrganizationPolicyObjectGroupsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_policy_object_groups"
}

func (d *OrganizationPolicyObjectGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Organization Policy Object Group` configuration.",

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
						"category": schema.StringAttribute{
							MarkdownDescription: "Category of a policy object group (one of: NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup)",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "A name for the group of network addresses, unique within the organization (alphanumeric, space, dash, or underscore characters only)",
							Computed:            true,
						},
						"object_ids": schema.SetAttribute{
							MarkdownDescription: "A list of Policy Object ID`s that this NetworkObjectGroup should be associated to (note: these ID`s will replace the existing associated Policy Objects)",
							ElementType:         types.Int64Type,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationPolicyObjectGroupsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *OrganizationPolicyObjectGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config OrganizationPolicyObjectGroups

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", "OrganizationPolicyObjectGroupsDataSource"))

	res, err := d.client.Get(config.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", "OrganizationPolicyObjectGroupsDataSource"))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
