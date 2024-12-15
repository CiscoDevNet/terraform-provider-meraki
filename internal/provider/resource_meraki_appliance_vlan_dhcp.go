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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
	_ resource.Resource = &ApplianceVLANDHCPResource{}
)

func NewApplianceVLANDHCPResource() resource.Resource {
	return &ApplianceVLANDHCPResource{}
}

type ApplianceVLANDHCPResource struct {
	client *meraki.Client
}

func (r *ApplianceVLANDHCPResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_vlan_dhcp"
}

func (r *ApplianceVLANDHCPResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource is meant to be used in addition to the `meraki_appliance_vlan` resource to configure DHCP settings for a VLAN. It requires the VLAN to be already configured on the appliance.").String,

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
			"vlan_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The VLAN ID of the new VLAN (must be between 1 and 4094)").String,
				Required:            true,
			},
			"dhcp_boot_filename": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DHCP boot option for boot filename").String,
				Optional:            true,
			},
			"dhcp_boot_next_server": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("DHCP boot option to direct boot clients to the server to load the boot file from").String,
				Optional:            true,
			},
			"dhcp_boot_options_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Use DHCP boot options specified in other properties").String,
				Optional:            true,
			},
			"dhcp_handling": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The appliance`s handling of DHCP requests on this VLAN. One of: `Run a DHCP server`, `Relay DHCP to another server` or `Do not respond to DHCP requests`").AddStringEnumDescription("Do not respond to DHCP requests", "Relay DHCP to another server", "Run a DHCP server").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Do not respond to DHCP requests", "Relay DHCP to another server", "Run a DHCP server"),
				},
			},
			"dhcp_lease_time": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: `30 minutes`, `1 hour`, `4 hours`, `12 hours`, `1 day` or `1 week`").AddStringEnumDescription("1 day", "1 hour", "1 week", "12 hours", "30 minutes", "4 hours").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("1 day", "1 hour", "1 week", "12 hours", "30 minutes", "4 hours"),
				},
			},
			"dns_nameservers": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The DNS nameservers used for DHCP responses, either 'upstream_dns', 'google_dns', 'opendns', or a newline seperated string of IP addresses or domain names").String,
				Optional:            true,
			},
			"mandatory_dhcp_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable Mandatory DHCP on VLAN.").String,
				Optional:            true,
			},
			"dhcp_options": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The list of DHCP options that will be included in DHCP responses. Each object in the list should have 'code', 'type', and 'value' properties.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"code": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The code for the DHCP option. This should be an integer between 2 and 254.").String,
							Required:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type for the DHCP option. One of: `text`, `ip`, `hex` or `integer`").AddStringEnumDescription("hex", "integer", "ip", "text").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("hex", "integer", "ip", "text"),
							},
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The value for the DHCP option").String,
							Required:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
					},
				},
			},
			"dhcp_relay_server_ips": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The IPs of the DHCP servers that DHCP requests should be relayed to").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"reserved_ip_ranges": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The DHCP reserved IP ranges on the VLAN").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"comment": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("A text comment for the reserved range").String,
							Required:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"end": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The last IP in the reserved range").String,
							Required:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"start": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The first IP in the reserved range").String,
							Required:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
					},
				},
			},
		},
	}
}

func (r *ApplianceVLANDHCPResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

func (r *ApplianceVLANDHCPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplianceVLANDHCP
	var existing ApplianceVLAN

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Read existing vlan
	res, err := r.client.Get(plan.getPath() + "/" + url.QueryEscape(plan.VlanId.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}
	existing.fromBody(ctx, res)

	// Create object
	body := plan.toBody(ctx, existing)
	res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.VlanId.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ApplianceVLANDHCPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplianceVLANDHCP

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBodyPartial(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *ApplianceVLANDHCPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplianceVLANDHCP
	var existing ApplianceVLAN

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

	// Read existing vlan
	res, err := r.client.Get(plan.getPath() + "/" + url.QueryEscape(plan.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}
	existing.fromBody(ctx, res)

	body := plan.toBody(ctx, existing)
	res, err = r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// Section below is generated&owned by "gen/generator.go". //template:begin delete

func (r *ApplianceVLANDHCPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplianceVLANDHCP

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
// End of section. //template:end import
