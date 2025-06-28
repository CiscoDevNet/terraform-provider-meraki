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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &AppliancePortsResource{}
	_ resource.ResourceWithImportState = &AppliancePortsResource{}
)

func NewAppliancePortsResource() resource.Resource {
	return &AppliancePortsResource{}
}

type AppliancePortsResource struct {
	client *meraki.Client
}

func (r *AppliancePortsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_ports"
}

func (r *AppliancePortsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance Port` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: "The organization ID",
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network ID").String,
				Required:            true,
			},
			"items": schema.ListNestedAttribute{
				MarkdownDescription: "The list of items",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Port ID").String,
							Required:            true,
						},
						"access_policy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the policy. Only applicable to Access ports. Valid values are: `open`, `8021x-radius`, `mac-radius`, `hybris-radius` for MX64 or Z3 or any MX supporting the per port authentication feature. Otherwise, `open` is the only valid value and `open` is the default value if the field is missing.").String,
							Optional:            true,
						},
						"allowed_vlans": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Comma-delimited list of the VLAN ID`s allowed on the port, or `all` to permit all VLAN`s on the port.").String,
							Optional:            true,
						},
						"drop_untagged_traffic": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Trunk port can Drop all Untagged traffic. When true, no VLAN is required. Access ports cannot have dropUntaggedTraffic set to true.").String,
							Optional:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The status of the port").String,
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of the port: `access` or `trunk`.").String,
							Optional:            true,
						},
						"vlan": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *AppliancePortsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *AppliancePortsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ResourceAppliancePorts

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	actions := make([]meraki.ActionModel, len(plan.Items))

	for i, item := range plan.Items {
		actions[i] = meraki.ActionModel{
			Operation: "update",
			Resource:  plan.getItemPath(item.PortId.ValueString()),
			Body:      item.toBody(ctx, ResourceAppliancePortsItems{}),
		}
	}
	res, err := r.client.Batch(plan.OrganizationId.ValueString(), actions)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.OrganizationId

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *AppliancePortsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ResourceAppliancePorts

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath())
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.Id = state.OrganizationId
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *AppliancePortsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ResourceAppliancePorts

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read state
	diags = req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))
	var actions []meraki.ActionModel

	// Check for new and updated items
	for _, item := range plan.Items {
		found := false
		for _, itemState := range state.Items {
			if item.PortId.ValueString() == itemState.PortId.ValueString() {
				found = true
				// If the item is present in both plan and state, we need to check if it has changes
				hasChanges := plan.hasChanges(ctx, &state, item.PortId.ValueString())
				if hasChanges {
					actions = append(actions, meraki.ActionModel{
						Operation: "update",
						Resource:  plan.getItemPath(item.PortId.ValueString()),
						Body:      item.toBody(ctx, itemState),
					})
				}
				break
			}
		}
		if !found {
			// If the item is present in plan, but not in state, we need to create it
			actions = append(actions, meraki.ActionModel{
				Operation: "update",
				Resource:  plan.getItemPath(item.PortId.ValueString()),
				Body:      item.toBody(ctx, ResourceAppliancePortsItems{}),
			})
		}
	}

	res, err := r.client.Batch(plan.OrganizationId.ValueString(), actions)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure objects (Action Batch), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *AppliancePortsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ResourceAppliancePorts

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *AppliancePortsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <organization_id>,<network_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
