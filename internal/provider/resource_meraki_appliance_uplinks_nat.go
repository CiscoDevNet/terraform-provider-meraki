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
	"net/url"
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
	_ resource.ResourceWithIdentity    = &ApplianceUplinksNATResource{}
	_ resource.ResourceWithImportState = &ApplianceUplinksNATResource{}
)

func NewApplianceUplinksNATResource() resource.Resource {
	return &ApplianceUplinksNATResource{}
}

type ApplianceUplinksNATResource struct {
	client                        *meraki.Client
	restoreOriginalStateOnDestroy bool
}

func (r *ApplianceUplinksNATResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_uplinks_nat"
}

func (r *ApplianceUplinksNATResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource allows configuring the NAT exceptions of the appliance uplinks of a network. The API requires an entry in `uplinks` for each WAN uplink of the network, for example both `wan1` and `wan2`.").String,

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
			"uplinks": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Per-uplink NAT exception configuration on the network.").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interface name of the uplink").String,
							Required:            true,
						},
						"nat_enabled": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Whether NAT is enabled on the uplink").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ApplianceUplinksNATResource) IdentitySchema(ctx context.Context, req resource.IdentitySchemaRequest, resp *resource.IdentitySchemaResponse) {
	resp.IdentitySchema = identityschema.Schema{
		Attributes: map[string]identityschema.Attribute{
			"network_id": identityschema.StringAttribute{
				Description:       helpers.NewAttributeDescription("Network ID").String,
				RequiredForImport: true,
			},
		},
	}
}

func (r *ApplianceUplinksNATResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
	r.restoreOriginalStateOnDestroy = req.ProviderData.(*MerakiProviderData).RestoreOriginalStateOnDestroy
}

// End of section. //template:end model

// getUplinksNAT resolves the organization of the network and returns the entry of the
// org level byNetwork endpoint that belongs to it, as there is no network level GET.
func (r *ApplianceUplinksNATResource) getUplinksNAT(networkId string) (meraki.Res, error) {
	nres, err := r.client.Get(fmt.Sprintf("/networks/%v", networkId))
	if err != nil {
		return nres, err
	}
	orgId := nres.Get("organizationId").String()
	res, err := r.client.Get(fmt.Sprintf("/organizations/%v/appliance/uplinks/nat/byNetwork?networkIds[]=%v", orgId, url.QueryEscape(networkId)))
	if err != nil {
		return res, err
	}
	// The API wraps the list in `items`, although the spec documents a bare array
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	result := meraki.Res{Result: gjson.Parse("{}")}
	res.ForEach(func(k, v gjson.Result) bool {
		if v.Get("networkId").String() == networkId {
			result = meraki.Res{Result: v}
			return false
		}
		return true
	})
	return result, nil
}

// Create configures the uplinks NAT settings and, when restoring on destroy, saves the
// initial state read from the org level endpoint.
func (r *ApplianceUplinksNATResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplianceUplinksNAT
	var identity ApplianceUplinksNATIdentity

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))
	// If the resource is a singleton, we need to read and save the initial state
	if r.restoreOriginalStateOnDestroy {
		var initialState ApplianceUplinksNAT
		gres, err := r.getUplinksNAT(plan.NetworkId.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve uplinks NAT settings (GET), got error: %s, %s", err, gres.String()))
			return
		}
		helpers.SetJsonInitialState(ctx, initialState.toBodyPreservingNulls(ctx, gres), resp.Private, &resp.Diagnostics)
	}

	// Create object
	body := plan.toBody(ctx, ApplianceUplinksNAT{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.NetworkId
	plan.fromBodyUnknowns(ctx, res)
	identity.toIdentity(ctx, &plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = resp.Identity.Set(ctx, &identity)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// Read refreshes the state from the org level byNetwork endpoint.
func (r *ApplianceUplinksNATResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplianceUplinksNAT
	var identity ApplianceUplinksNATIdentity

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
	res, err := r.getUplinksNAT(state.NetworkId.ValueString())
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 400")) {
		identity.toIdentity(ctx, &state)
		diags = resp.Identity.Set(ctx, &identity)
		resp.Diagnostics.Append(diags...)
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve uplinks NAT settings (GET), got error: %s, %s", err, res.String()))
		return
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

func (r *ApplianceUplinksNATResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplianceUplinksNAT
	var identity ApplianceUplinksNATIdentity

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

func (r *ApplianceUplinksNATResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplianceUplinksNAT

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
func (r *ApplianceUplinksNATResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if req.ID != "" || req.Identity == nil || req.Identity.Raw.IsNull() {
		idParts := strings.Split(req.ID, ",")

		if len(idParts) != 1 || idParts[0] == "" {
			resp.Diagnostics.AddError(
				"Unexpected Import Identifier",
				fmt.Sprintf("Expected import identifier with format: <network_id>. Got: %q", req.ID),
			)
			return
		}
		resp.Diagnostics.Append(resp.Identity.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)
	} else {
		var identity ApplianceUplinksNATIdentity
		diags := req.Identity.Get(ctx, &identity)
		if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), identity.NetworkId.ValueString())...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), identity.NetworkId.ValueString())...)
	}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
