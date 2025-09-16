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
	_ resource.Resource                = &SwitchStackRoutingInterfaceDHCPResource{}
	_ resource.ResourceWithImportState = &SwitchStackRoutingInterfaceDHCPResource{}
)

func NewSwitchStackRoutingInterfaceDHCPResource() resource.Resource {
	return &SwitchStackRoutingInterfaceDHCPResource{}
}

type SwitchStackRoutingInterfaceDHCPResource struct {
	client *meraki.Client
}

func (r *SwitchStackRoutingInterfaceDHCPResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_switch_stack_routing_interface_dhcp"
}

func (r *SwitchStackRoutingInterfaceDHCPResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Switch Stack Routing Interface DHCP` configuration.").String,

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
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Switch stack ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Interface ID").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"boot_file_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The PXE boot server file name for the DHCP server running on the switch stack interface").String,
				Optional:            true,
			},
			"boot_next_server": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The PXE boot server IP for the DHCP server running on the switch stack interface").String,
				Optional:            true,
			},
			"boot_options_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface").String,
				Optional:            true,
			},
			"dhcp_lease_time": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The DHCP lease time config for the dhcp server running on switch stack interface (`30 minutes`, `1 hour`, `4 hours`, `12 hours`, `1 day` or `1 week`)").AddStringEnumDescription("1 day", "1 hour", "1 week", "12 hours", "30 minutes", "4 hours").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1 day", "1 hour", "1 week", "12 hours", "30 minutes", "4 hours"),
				},
			},
			"dhcp_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The DHCP mode options for the switch stack interface (`dhcpDisabled`, `dhcpRelay` or `dhcpServer`)").AddStringEnumDescription("dhcpDisabled", "dhcpRelay", "dhcpServer").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("dhcpDisabled", "dhcpRelay", "dhcpServer"),
				},
			},
			"dns_nameservers_option": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The DHCP name server option for the dhcp server running on the switch stack interface (`googlePublicDns`, `openDns` or `custom`)").AddStringEnumDescription("custom", "googlePublicDns", "openDns").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("custom", "googlePublicDns", "openDns"),
				},
			},
			"dhcp_options": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"code": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The code for DHCP option which should be from 2 to 254").String,
							Required:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type of the DHCP option which should be one of (`text`, `ip`, `integer` or `hex`)").AddStringEnumDescription("hex", "integer", "ip", "text").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("hex", "integer", "ip", "text"),
							},
						},
						"value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The value of the DHCP option").String,
							Required:            true,
						},
					},
				},
			},
			"dhcp_relay_server_ips": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"dns_custom_nameservers": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The DHCP name server IPs when DHCP name server option is ` custom`").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"fixed_ip_assignments": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Array of DHCP fixed IP assignments for the DHCP server running on the switch stack interface").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The IP address of the client which has fixed IP address assigned to it").String,
							Required:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The MAC address of the client which has fixed IP address").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the client which has fixed IP address").String,
							Required:            true,
						},
					},
				},
			},
			"reserved_ip_ranges": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"comment": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The comment for the reserved IP range").String,
							Optional:            true,
						},
						"end": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The ending IP address of the reserved IP range").String,
							Required:            true,
						},
						"start": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The starting IP address of the reserved IP range").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *SwitchStackRoutingInterfaceDHCPResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *SwitchStackRoutingInterfaceDHCPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, initialState SwitchStackRoutingInterfaceDHCP

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
	helpers.SetJsonInitialState(ctx, initialState.toBody(ctx, SwitchStackRoutingInterfaceDHCP{}), resp.Private, &resp.Diagnostics)

	// Create object
	body := plan.toBody(ctx, SwitchStackRoutingInterfaceDHCP{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.InterfaceId
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *SwitchStackRoutingInterfaceDHCPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SwitchStackRoutingInterfaceDHCP

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

func (r *SwitchStackRoutingInterfaceDHCPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SwitchStackRoutingInterfaceDHCP

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

func (r *SwitchStackRoutingInterfaceDHCPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SwitchStackRoutingInterfaceDHCP

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	// If the resource is a singleton, we need to restore the initial state
	jsonInitialState, diags := helpers.GetJsonInitialState(ctx, req)
	jsonInitialState = state.addDeleteValues(ctx, jsonInitialState)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.Put(state.getPath(), jsonInitialState)
	if err != nil {
		resp.Diagnostics.AddWarning("Failed to restore initial state", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *SwitchStackRoutingInterfaceDHCPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_id>,<switch_stack_id>,<interface_id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("switch_stack_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("interface_id"), idParts[2])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[2])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
