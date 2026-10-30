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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.ResourceWithIdentity    = &ApplianceVPNSiteToSiteHubVRFsResource{}
	_ resource.ResourceWithImportState = &ApplianceVPNSiteToSiteHubVRFsResource{}
)

func NewApplianceVPNSiteToSiteHubVRFsResource() resource.Resource {
	return &ApplianceVPNSiteToSiteHubVRFsResource{}
}

type ApplianceVPNSiteToSiteHubVRFsResource struct {
	client                        *meraki.Client
	restoreOriginalStateOnDestroy bool
}

func (r *ApplianceVPNSiteToSiteHubVRFsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_vpn_site_to_site_hub_vrfs"
}

func (r *ApplianceVPNSiteToSiteHubVRFsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance VPN Site to Site Hub VRFs` configuration.").AddEarlyAccessDescription().String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Network ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"hub_network_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Hub network ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"vrfs": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The list of VRFs for this source and hub mapping.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The VRF ID. Use `0` for the default VRF.").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ApplianceVPNSiteToSiteHubVRFsResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"network_id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("Network ID").String,
				RequiredForImport: true,
			},
			"hub_network_id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("Hub network ID").String,
				RequiredForImport: true,
			},
		},
	}
}

func (r *ApplianceVPNSiteToSiteHubVRFsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
	r.restoreOriginalStateOnDestroy = req.ProviderData.(*MerakiProviderData).RestoreOriginalStateOnDestroy
}

// End of section. //template:end model

// Create configures the resource.
// Custom (not-generated) changes: if restoreOriginalStateOnDestroy is enabled,
// fetch the original state using the organization-level GET endpoint.
func (r *ApplianceVPNSiteToSiteHubVRFsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplianceVPNSiteToSiteHubVRFs
	var identity ApplianceVPNSiteToSiteHubVRFsIdentity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))
	// If the resource is a singleton, we need to read and save the initial state
	if r.restoreOriginalStateOnDestroy {
		var initialState ApplianceVPNSiteToSiteHubVRFs
		networkPath := fmt.Sprintf("/networks/%v", plan.NetworkId.ValueString())
		nres, err := r.client.Get(networkPath)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network (GET), got error: %s, %s", err, nres.String()))
			return
		}
		orgId := nres.Get("organizationId").String()
		getPath := fmt.Sprintf("/organizations/%v/appliance/vpn/siteToSite/hubs/vrfs?networkIds[]=%v", orgId, plan.NetworkId.ValueString())
		gres, err := r.client.Get(getPath)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve appliance VPN site-to-site hub VRFs (GET), got error: %s, %s", err, gres.String()))
			return
		}
		if len(gres.Get("items").Array()) > 0 {
			gres.Get("items").ForEach(func(k, v gjson.Result) bool {
				if plan.Id.ValueString() == v.Get("hub.network.id").String() {
					gres = meraki.Res{Result: v}
					return false
				}
				return true
			})
		}
		helpers.SetJsonInitialState(ctx, initialState.toBodyPreservingNulls(ctx, gres), resp.Private, &resp.Diagnostics)
	}

	// Create object
	body := plan.toBody(ctx, ApplianceVPNSiteToSiteHubVRFs{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.HubNetworkId
	plan.fromBodyUnknowns(ctx, res)
	identity.toIdentity(ctx, &plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Read fetches the resource's current configuration.
// Custom (not-generated) changes: use the organization-level GET endpoint.
func (r *ApplianceVPNSiteToSiteHubVRFsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplianceVPNSiteToSiteHubVRFs
	var identity ApplianceVPNSiteToSiteHubVRFsIdentity

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

	networkPath := fmt.Sprintf("/networks/%v", state.NetworkId.ValueString())
	res, err := r.client.Get(networkPath)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 400")) {
		identity.toIdentity(ctx, &state)
		diags = resp.Identity.Set(ctx, &identity)
		resp.Diagnostics.Append(diags...)
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve network (GET), got error: %s, %s", err, res.String()))
		return
	}
	orgId := res.Get("organizationId").String()

	getPath := fmt.Sprintf("/organizations/%v/appliance/vpn/siteToSite/hubs/vrfs?networkIds[]=%v", orgId, state.NetworkId.ValueString())
	res, err = r.client.Get(getPath)
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 400")) {
		identity.toIdentity(ctx, &state)
		diags = resp.Identity.Set(ctx, &identity)
		resp.Diagnostics.Append(diags...)
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve appliance VPN site-to-site hub VRFs (GET), got error: %s, %s", err, res.String()))
		return
	}

	if len(res.Get("items").Array()) > 0 {
		res.Get("items").ForEach(func(k, v gjson.Result) bool {
			if state.Id.ValueString() == v.Get("hub.network.id").String() {
				res = meraki.Res{Result: v}
				return false
			}
			return true
		})
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	// After `terraform import` we switch to a full read.
	if imp {
		state.fromBody(ctx, res)
	} else {
		state.fromBodyPartial(ctx, res)
	}
	identity.toIdentity(ctx, &state)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Section below is generated&owned by "gen/generator.go". //template:begin update

func (r *ApplianceVPNSiteToSiteHubVRFsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplianceVPNSiteToSiteHubVRFs
	var identity ApplianceVPNSiteToSiteHubVRFsIdentity

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

	body := plan.toBody(ctx, state)
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	identity.toIdentity(ctx, &plan)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *ApplianceVPNSiteToSiteHubVRFsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplianceVPNSiteToSiteHubVRFs

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	if r.restoreOriginalStateOnDestroy {
		// Restore the saved initial state on destroy
		jsonInitialState, diags := helpers.GetJsonInitialState(ctx, req)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}

		res, err := r.client.Put(state.getPath(), jsonInitialState)
		if err != nil {
			resp.Diagnostics.AddWarning("Failed to restore initial state", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ApplianceVPNSiteToSiteHubVRFsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != "" || req.Identity == nil || req.Identity.Raw.IsNull() {
		idParts := strings.Split(req.ID, ",")

		if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
			resp.Diagnostics.AddError(
				"Unexpected Import Identifier",
				fmt.Sprintf("Expected import identifier with format: <network_id>,<hub_network_id>. Got: %q", req.ID),
			)
			return
		}
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("hub_network_id"), idParts[1])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("hub_network_id"), idParts[1])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
	} else {
		var identity ApplianceVPNSiteToSiteHubVRFsIdentity
		diags := req.Identity.Get(ctx, &identity)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), identity.NetworkId.ValueString())...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("hub_network_id"), identity.HubNetworkId.ValueString())...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), identity.HubNetworkId.ValueString())...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
