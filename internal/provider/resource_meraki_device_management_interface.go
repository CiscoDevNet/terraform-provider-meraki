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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var (
	_ resource.Resource                = &DeviceManagementInterfaceResource{}
	_ resource.ResourceWithImportState = &DeviceManagementInterfaceResource{}
)

func NewDeviceManagementInterfaceResource() resource.Resource {
	return &DeviceManagementInterfaceResource{}
}

type DeviceManagementInterfaceResource struct {
	client *meraki.Client
}

func (r *DeviceManagementInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_management_interface"
}

func (r *DeviceManagementInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Device Management Interface` configuration.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Device serial").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"wan1_static_gateway_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IP of the gateway on the WAN.").String,
				Optional:            true,
			},
			"wan1_static_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IP the device should use on the WAN.").String,
				Optional:            true,
			},
			"wan1_static_subnet_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The subnet mask for the WAN.").String,
				Optional:            true,
			},
			"wan1_using_static_ip": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the interface to have static IP settings or use DHCP.").String,
				Optional:            true,
			},
			"wan1_vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.").String,
				Optional:            true,
			},
			"wan1_wan_enabled": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable the interface (only for MX devices). Valid values are `enabled`, `disabled`, and `not configured`.").AddStringEnumDescription("disabled", "enabled", "not configured").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "enabled", "not configured"),
				},
			},
			"wan1_static_dns": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Up to two DNS IPs.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"wan2_static_gateway_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IP of the gateway on the WAN.").String,
				Optional:            true,
			},
			"wan2_static_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IP the device should use on the WAN.").String,
				Optional:            true,
			},
			"wan2_static_subnet_mask": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The subnet mask for the WAN.").String,
				Optional:            true,
			},
			"wan2_using_static_ip": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure the interface to have static IP settings or use DHCP.").String,
				Optional:            true,
			},
			"wan2_vlan": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.").String,
				Optional:            true,
			},
			"wan2_wan_enabled": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable or disable the interface (only for MX devices). Valid values are `enabled`, `disabled`, and `not configured`.").AddStringEnumDescription("disabled", "enabled", "not configured").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "enabled", "not configured"),
				},
			},
			"wan2_static_dns": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Up to two DNS IPs.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *DeviceManagementInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *DeviceManagementInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, initialState DeviceManagementInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))
	// If the resource is a singleton, we need to read and save the initial state
	gres, err := r.client.Get(plan.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, gres.String()))
		return
	}
	initialState.fromBody(ctx, gres)
	helpers.SetJsonInitialState(ctx, initialState.toBody(ctx, DeviceManagementInterface{}), resp.Private, &resp.Diagnostics)

	// Create object
	body := plan.toBody(ctx, DeviceManagementInterface{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.Serial
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *DeviceManagementInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceManagementInterface

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

func (r *DeviceManagementInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DeviceManagementInterface

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
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *DeviceManagementInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceManagementInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	// If the resource is a singleton, we need to restore the initial state
	jsonInitialState, diags := helpers.GetJsonInitialState(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.Put(state.getPath(), jsonInitialState)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *DeviceManagementInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 1 || idParts[0] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <serial>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[0])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
