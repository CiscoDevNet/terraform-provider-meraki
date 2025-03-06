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
	_ resource.Resource                = &ApplianceVLANResource{}
	_ resource.ResourceWithImportState = &ApplianceVLANResource{}
)

func NewApplianceVLANResource() resource.Resource {
	return &ApplianceVLANResource{}
}

type ApplianceVLANResource struct {
	client *meraki.Client
}

func (r *ApplianceVLANResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appliance_vlan"
}

func (r *ApplianceVLANResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Appliance VLAN` configuration.").String,

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
			"appliance_ip": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The local IP of the appliance on the VLAN").String,
				Optional:            true,
			},
			"cidr": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.").String,
				Optional:            true,
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
			"group_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The id of the desired group policy to apply to the VLAN").String,
				Optional:            true,
			},
			"vlan_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The VLAN ID of the new VLAN (must be between 1 and 4094)").String,
				Required:            true,
			},
			"mask": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Mask used for the subnet of all bound to the template networks. Applicable only for template network.").String,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The name of the new VLAN").String,
				Required:            true,
			},
			"subnet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The subnet of the VLAN").String,
				Optional:            true,
			},
			"template_vlan_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of subnetting of the VLAN. Applicable only for template network.").AddStringEnumDescription("same", "unique").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("same", "unique"),
				},
			},
			"vpn_nat_subnet": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN").String,
				Optional:            true,
			},
			"fixed_ip_assignments": schema.MapNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The DHCP fixed IP assignments on the VLAN. Thekey of this map is a MAC address.").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The IP address to assign to the client").String,
							Required:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the client").String,
							Required:            true,
						},
					},
				},
			},
			"ipv6_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable IPv6 on VLAN.").String,
				Optional:            true,
			},
			"ipv6_prefix_assignments": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prefix assignments on the VLAN").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"autonomous": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Auto assign a /64 prefix from the origin to the VLAN").String,
							Optional:            true,
						},
						"static_appliance_ip6": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Manual configuration of the IPv6 Appliance IP").String,
							Optional:            true,
						},
						"static_prefix": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Manual configuration of a /64 prefix on the VLAN").String,
							Optional:            true,
						},
						"origin_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Type of the origin").AddStringEnumDescription("independent", "internet").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("independent", "internet"),
							},
						},
						"origin_interfaces": schema.ListAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Interfaces associated with the prefix").String,
							ElementType:         types.StringType,
							Optional:            true,
						},
					},
				},
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
						},
						"type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The type for the DHCP option. One of: `text`, `ip`, `hex` or `integer`").AddStringEnumDescription("hex", "integer", "ip", "text").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("hex", "integer", "ip", "text"),
							},
						},
						"value": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The value for the DHCP option").String,
							Required:            true,
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
						},
						"end": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The last IP in the reserved range").String,
							Required:            true,
						},
						"start": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The first IP in the reserved range").String,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *ApplianceVLANResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *ApplianceVLANResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplianceVLAN

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, ApplianceVLAN{})
	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("id").String())
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *ApplianceVLANResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ApplianceVLAN

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

func (r *ApplianceVLANResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ApplianceVLAN

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
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
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

func (r *ApplianceVLANResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplianceVLAN

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && !strings.Contains(err.Error(), "StatusCode 404") {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ApplianceVLANResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_id>,<id>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
