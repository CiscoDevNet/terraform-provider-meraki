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
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &NetworkDeviceClaimResource{}
	_ resource.ResourceWithImportState = &NetworkDeviceClaimResource{}
)

func NewNetworkDeviceClaimResource() resource.Resource {
	return &NetworkDeviceClaimResource{}
}

type NetworkDeviceClaimResource struct {
	client *meraki.Client
}

func (r *NetworkDeviceClaimResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_device_claim"
}

func (r *NetworkDeviceClaimResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource allows claiming and removing serials to a network. It will not not touch any existing serials already claimed and not included in `serials`. Removing a serial from a network will return it to the organization inventory.").String,

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
			"details_by_device": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Optional details for claimed devices (currently only used for Catalyst devices)").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"serial": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The serial of the device these details relate to").String,
							Required:            true,
						},
						"details": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("An array of details. Supported list of details includes: 'device mode', 'username', 'password', 'enable password', 'ap mapping type' and 'ap network id'. For onboarding into hybrid mode, the value of the device mode detail must be 'monitored'").String,
							Required:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Name of device detail").String,
										Required:            true,
									},
									"value": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Value of device detail").String,
										Optional:            true,
										Sensitive:           true,
									},
								},
							},
						},
					},
				},
			},
			"serials": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of serials of devices to claim").String,
				ElementType:         types.StringType,
				Required:            true,
			},
		},
	}
}

func (r *NetworkDeviceClaimResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *NetworkDeviceClaimResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkDeviceClaim

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, NetworkDeviceClaim{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.NetworkId
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

func (r *NetworkDeviceClaimResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkDeviceClaim

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	res, err := r.client.Get(state.getDevicesPath())
	if err != nil && (strings.Contains(err.Error(), "StatusCode 404") || strings.Contains(err.Error(), "StatusCode 400")) {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve inventory (GET), got error: %s, %s", err, res.String()))
		return
	}

	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	var resultSerials []string
	if imp {
		for _, device := range res.Array() {
			rSerial := device.Get("serial").String()
			resultSerials = append(resultSerials, rSerial)
		}
	} else {
		var serials []string
		state.Serials.ElementsAs(ctx, &serials, false)
		for _, serial := range serials {
			for _, device := range res.Array() {
				rSerial := device.Get("serial").String()
				if serial == rSerial {
					resultSerials = append(resultSerials, serial)
				}
			}
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("%s: Retrieved the following serials: %v", state.Id.ValueString(), resultSerials))
	v := make([]attr.Value, len(resultSerials))
	for r := range resultSerials {
		v[r] = types.StringValue(resultSerials[r])
	}
	state.Serials = types.SetValueMust(types.StringType, v)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

func (r *NetworkDeviceClaimResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state NetworkDeviceClaim

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

	var planSerials, stateSerials, removeSerials []string
	plan.Serials.ElementsAs(ctx, &planSerials, false)
	state.Serials.ElementsAs(ctx, &stateSerials, false)

	claimBody := ""
	for _, planSerial := range planSerials {
		found := false
		for _, stateSerial := range stateSerials {
			if planSerial == stateSerial {
				found = true
			}
		}
		if !found {
			claimBody, _ = sjson.Set(claimBody, "serials.-1", planSerial)
			for _, details := range plan.DetailsByDevice {
				if details.Serial.ValueString() == planSerial {
					deviceDetailsBody, _ := sjson.Set("", "serial", details.Serial.ValueString())
					for _, detail := range details.Details {
						detailsBody, _ := sjson.Set("", "name", detail.Name.ValueString())
						detailsBody, _ = sjson.Set(detailsBody, "value", detail.Value.ValueString())
						deviceDetailsBody, _ = sjson.SetRaw(deviceDetailsBody, "details.-1", detailsBody)
					}
					claimBody, _ = sjson.SetRaw(claimBody, "detailsByDevice.-1", deviceDetailsBody)
				}
			}
		}
	}

	for _, stateSerial := range stateSerials {
		found := false
		for _, planSerial := range planSerials {
			if planSerial == stateSerial {
				found = true
			}
		}
		if !found {
			removeSerials = append(removeSerials, stateSerial)
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Serials to be claimed: %v", plan.Id.ValueString(), gjson.Get(claimBody, "serials").String()))
	if len(gjson.Get(claimBody, "serials").Array()) > 0 {
		res, err := r.client.Post(plan.getPath(), claimBody)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to claim devices to network, got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Serials to be removed: %v", plan.Id.ValueString(), removeSerials))
	for _, serial := range removeSerials {
		body, _ := sjson.Set("", "serial", serial)
		res, err := r.client.Post(plan.getRemovePath(), body)
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to remove device from network, got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworkDeviceClaimResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkDeviceClaim

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	var stateSerials []string
	state.Serials.ElementsAs(ctx, &stateSerials, false)

	for _, serial := range stateSerials {
		body, _ := sjson.Set("", "serial", serial)
		res, err := r.client.Post(state.getRemovePath(), body)
		if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to remove device from network, got error: %s, %s", err, res.String()))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *NetworkDeviceClaimResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
