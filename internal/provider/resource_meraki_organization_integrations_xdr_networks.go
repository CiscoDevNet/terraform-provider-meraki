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
	"github.com/hashicorp/terraform-plugin-framework/resource/identityschema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.ResourceWithIdentity    = &OrganizationIntegrationsXDRNetworksResource{}
	_ resource.ResourceWithImportState = &OrganizationIntegrationsXDRNetworksResource{}
)

func NewOrganizationIntegrationsXDRNetworksResource() resource.Resource {
	return &OrganizationIntegrationsXDRNetworksResource{}
}

type OrganizationIntegrationsXDRNetworksResource struct {
	client *meraki.Client
}

func (r *OrganizationIntegrationsXDRNetworksResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization_integrations_xdr_networks"
}

func (r *OrganizationIntegrationsXDRNetworksResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Organization Integrations XDR Networks` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Organization ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"networks": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List containing the network ID and the product type to enable XDR on").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Network ID").String,
							Required:            true,
						},
						"product_types": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("List of products for which to enable XDR").String,
							ElementType:         types.StringType,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *OrganizationIntegrationsXDRNetworksResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"organization_id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("Organization ID").String,
				RequiredForImport: true,
			},
		},
	}
}

func (r *OrganizationIntegrationsXDRNetworksResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *OrganizationIntegrationsXDRNetworksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OrganizationIntegrationsXDRNetworks
	var identity OrganizationIntegrationsXDRNetworksIdentity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, OrganizationIntegrationsXDRNetworks{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.OrganizationId
	plan.fromBodyUnknowns(ctx, res)
	identity.toIdentity(ctx, &plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

func (r *OrganizationIntegrationsXDRNetworksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OrganizationIntegrationsXDRNetworks
	var identity OrganizationIntegrationsXDRNetworksIdentity

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// Read identity if available (requires Terraform >= 1.12.0)
	if req.Identity != nil && !req.Identity.Raw.IsNull() {
		diags = req.Identity.Get(ctx, &identity)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		state.fromIdentity(ctx, &identity)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	res, err := r.client.Get(state.getNetworksPath())
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

	items := res.Get("items")
	if imp {
		state.Networks = make([]OrganizationIntegrationsXDRNetworksNetworks, 0)
		items.ForEach(func(_, item gjson.Result) bool {
			n := OrganizationIntegrationsXDRNetworksNetworks{}
			if v := item.Get("networkId"); v.Exists() {
				n.NetworkId = types.StringValue(v.String())
			}
			if v := item.Get("productTypes"); v.Exists() {
				n.ProductTypes = helpers.GetStringList(v.Array())
			}
			state.Networks = append(state.Networks, n)
			return true
		})
	} else {
		stateNetworkIds := make(map[string]bool)
		for _, n := range state.Networks {
			stateNetworkIds[n.NetworkId.ValueString()] = true
		}
		var result []OrganizationIntegrationsXDRNetworksNetworks
		items.ForEach(func(_, item gjson.Result) bool {
			nid := item.Get("networkId").String()
			if stateNetworkIds[nid] {
				n := OrganizationIntegrationsXDRNetworksNetworks{}
				n.NetworkId = types.StringValue(nid)
				if v := item.Get("productTypes"); v.Exists() {
					n.ProductTypes = helpers.GetStringList(v.Array())
				}
				result = append(result, n)
			}
			return true
		})
		state.Networks = result
	}
	identity.toIdentity(ctx, &state)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *OrganizationIntegrationsXDRNetworksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state OrganizationIntegrationsXDRNetworks

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

	stateNetworkIds := make(map[string]bool)
	for _, n := range state.Networks {
		stateNetworkIds[n.NetworkId.ValueString()] = true
	}
	planNetworkIds := make(map[string]bool)
	for _, n := range plan.Networks {
		planNetworkIds[n.NetworkId.ValueString()] = true
	}

	// Enable networks in plan but not in state
	enableBody := plan.toBody(ctx, state)
	var enableNetworks []OrganizationIntegrationsXDRNetworksNetworks
	for _, n := range plan.Networks {
		if !stateNetworkIds[n.NetworkId.ValueString()] {
			enableNetworks = append(enableNetworks, n)
		}
	}
	if len(enableNetworks) > 0 {
		enableData := OrganizationIntegrationsXDRNetworks{
			OrganizationId: plan.OrganizationId,
			Networks:       enableNetworks,
		}
		enableBody = enableData.toBody(ctx, OrganizationIntegrationsXDRNetworks{})
		res, err := r.client.Post(plan.getPath(), enableBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to enable XDR networks, got error: %s, %s", err, res.String()))
			return
		}
	}

	// Disable networks in state but not in plan
	var disableNetworks []OrganizationIntegrationsXDRNetworksNetworks
	for _, n := range state.Networks {
		if !planNetworkIds[n.NetworkId.ValueString()] {
			disableNetworks = append(disableNetworks, n)
		}
	}
	if len(disableNetworks) > 0 {
		disableData := OrganizationIntegrationsXDRNetworks{
			OrganizationId: plan.OrganizationId,
			Networks:       disableNetworks,
		}
		disableBody := disableData.toBody(ctx, OrganizationIntegrationsXDRNetworks{})
		res, err := r.client.Post(plan.getDisablePath(), disableBody)
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to disable XDR networks, got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationIntegrationsXDRNetworksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationIntegrationsXDRNetworks

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	body := state.toBody(ctx, OrganizationIntegrationsXDRNetworks{})
	res, err := r.client.Post(state.getDisablePath(), body)
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to disable XDR networks, got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *OrganizationIntegrationsXDRNetworksResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != "" || req.Identity == nil || req.Identity.Raw.IsNull() {
		idParts := strings.Split(req.ID, ",")

		if len(idParts) != 1 || idParts[0] == "" {
			resp.Diagnostics.AddError(
				"Unexpected Import Identifier",
				fmt.Sprintf("Expected import identifier with format: <organization_id>. Got: %q", req.ID),
			)
			return
		}
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
	} else {
		var identity OrganizationIntegrationsXDRNetworksIdentity
		diags := req.Identity.Get(ctx, &identity)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), identity.OrganizationId.ValueString())...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), identity.OrganizationId.ValueString())...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
