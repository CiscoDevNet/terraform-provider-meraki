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
	_ resource.Resource                = &WirelessSSIDHotspot20Resource{}
	_ resource.ResourceWithImportState = &WirelessSSIDHotspot20Resource{}
)

func NewWirelessSSIDHotspot20Resource() resource.Resource {
	return &WirelessSSIDHotspot20Resource{}
}

type WirelessSSIDHotspot20Resource struct {
	client *meraki.Client
}

func (r *WirelessSSIDHotspot20Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssid_hotspot_20"
}

func (r *WirelessSSIDHotspot20Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage the `Wireless SSID Hotspot 2.0` configuration.").String,

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
			"number": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Wireless SSID number").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Whether or not Hotspot 2.0 for this SSID is enabled").String,
				Optional:            true,
			},
			"network_access_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The network type of this SSID (`Private network`, `Private network with guest access`, `Chargeable public network`, `Free public network`, `Personal device network`, `Emergency services only network`, `Test or experimental`, `Wildcard`)").AddStringEnumDescription("Chargeable public network", "Emergency services only network", "Free public network", "Personal device network", "Private network", "Private network with guest access", "Test or experimental", "Wildcard").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Chargeable public network", "Emergency services only network", "Free public network", "Personal device network", "Private network", "Private network with guest access", "Test or experimental", "Wildcard"),
				},
			},
			"operator_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Operator name").String,
				Optional:            true,
			},
			"venue_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Venue name").String,
				Optional:            true,
			},
			"venue_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Venue type (`Unspecified`, `Unspecified Assembly`, `Arena`, `Stadium`, `Passenger Terminal`, `Amphitheater`, `Amusement Park`, `Place of Worship`, `Convention Center`, `Library`, `Museum`, `Restaurant`, `Theater`, `Bar`, `Coffee Shop`, `Zoo or Aquarium`, `Emergency Coordination Center`, `Unspecified Business`, `Doctor or Dentist office`, `Bank`, `Fire Station`, `Police Station`, `Post Office`, `Professional Office`, `Research and Development Facility`, `Attorney Office`, `Unspecified Educational`, `School, Primary`, `School, Secondary`, `University or College`, `Unspecified Factory and Industrial`, `Factory`, `Unspecified Institutional`, `Hospital`, `Long-Term Care Facility`, `Alcohol and Drug Rehabilitation Center`, `Group Home`, `Prison or Jail`, `Unspecified Mercantile`, `Retail Store`, `Grocery Market`, `Automotive Service Station`, `Shopping Mall`, `Gas Station`, `Unspecified Residential`, `Private Residence`, `Hotel or Motel`, `Dormitory`, `Boarding House`, `Unspecified Storage`, `Unspecified Utility and Miscellaneous`, `Unspecified Vehicular`, `Automobile or Truck`, `Airplane`, `Bus`, `Ferry`, `Ship or Boat`, `Train`, `Motor Bike`, `Unspecified Outdoor`, `Muni-mesh Network`, `City Park`, `Rest Area`, `Traffic Control`, `Bus Stop`, `Kiosk`)").AddStringEnumDescription("Airplane", "Alcohol and Drug Rehabilitation Center", "Amphitheater", "Amusement Park", "Arena", "Attorney Office", "Automobile or Truck", "Automotive Service Station", "Bank", "Bar", "Boarding House", "Bus", "Bus Stop", "City Park", "Coffee Shop", "Convention Center", "Doctor or Dentist office", "Dormitory", "Emergency Coordination Center", "Factory", "Ferry", "Fire Station", "Gas Station", "Grocery Market", "Group Home", "Hospital", "Hotel or Motel", "Kiosk", "Library", "Long-Term Care Facility", "Motor Bike", "Muni-mesh Network", "Museum", "Passenger Terminal", "Place of Worship", "Police Station", "Post Office", "Prison or Jail", "Private Residence", "Professional Office", "Research and Development Facility", "Rest Area", "Restaurant", "Retail Store", "School, Primary", "School, Secondary", "Ship or Boat", "Shopping Mall", "Stadium", "Theater", "Traffic Control", "Train", "University or College", "Unspecified", "Unspecified Assembly", "Unspecified Business", "Unspecified Educational", "Unspecified Factory and Industrial", "Unspecified Institutional", "Unspecified Mercantile", "Unspecified Outdoor", "Unspecified Residential", "Unspecified Storage", "Unspecified Utility and Miscellaneous", "Unspecified Vehicular", "Zoo or Aquarium").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("Airplane", "Alcohol and Drug Rehabilitation Center", "Amphitheater", "Amusement Park", "Arena", "Attorney Office", "Automobile or Truck", "Automotive Service Station", "Bank", "Bar", "Boarding House", "Bus", "Bus Stop", "City Park", "Coffee Shop", "Convention Center", "Doctor or Dentist office", "Dormitory", "Emergency Coordination Center", "Factory", "Ferry", "Fire Station", "Gas Station", "Grocery Market", "Group Home", "Hospital", "Hotel or Motel", "Kiosk", "Library", "Long-Term Care Facility", "Motor Bike", "Muni-mesh Network", "Museum", "Passenger Terminal", "Place of Worship", "Police Station", "Post Office", "Prison or Jail", "Private Residence", "Professional Office", "Research and Development Facility", "Rest Area", "Restaurant", "Retail Store", "School, Primary", "School, Secondary", "Ship or Boat", "Shopping Mall", "Stadium", "Theater", "Traffic Control", "Train", "University or College", "Unspecified", "Unspecified Assembly", "Unspecified Business", "Unspecified Educational", "Unspecified Factory and Industrial", "Unspecified Institutional", "Unspecified Mercantile", "Unspecified Outdoor", "Unspecified Residential", "Unspecified Storage", "Unspecified Utility and Miscellaneous", "Unspecified Vehicular", "Zoo or Aquarium"),
				},
			},
			"domains": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("An array of domain names").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"mcc_mncs": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("An array of MCC/MNC pairs").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mcc": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MCC value").String,
							Required:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"mnc": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("MNC value").String,
							Required:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
					},
				},
			},
			"nai_realms": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("An array of NAI realms").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"format": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The format for the realm (`1` or `0`)").AddStringEnumDescription("0", "1").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("0", "1"),
							},
						},
						"realm": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("The name of the realm").String,
							Required:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"methods": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("An array of EAP methods for the realm.").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("ID of method").String,
										Required:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
									},
									"authentication_types_non_eap_inner_authentication": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("An array of autentication types. Possible values are `Reserved`, `PAP`, `CHAP`, `MSCHAP`, `MSCHAPV2`.").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"authentication_types_eap_inner_authentication": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("An array of autentication types. Possible values are `EAP-TLS`, `EAP-SIM`, `EAP-AKA`, `EAP-TTLS with MSCHAPv2`.").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"authentication_types_credentials": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("An array of autentication types. Possible values are `SIM`, `USIM`, `NFC Secure Element`, `Hardware Token`, `Softoken`, `Certificate`, `username/password`, `none`, `Reserved`, `Vendor Specific`.").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
									"authentication_types_tunneled_eap_method_credentials": schema.SetAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("An array of autentication types. Possible values are `SIM`, `USIM`, `NFC Secure Element`, `Hardware Token`, `Softoken`, `Certificate`, `username/password`, `Reserved`, `Anonymous`, `Vendor Specific`.").String,
										ElementType:         types.StringType,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"roam_consort_ois": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("An array of roaming consortium OIs (hexadecimal number 3-5 octets in length)").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *WirelessSSIDHotspot20Resource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create

func (r *WirelessSSIDHotspot20Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WirelessSSIDHotspot20

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, WirelessSSIDHotspot20{})
	res, err := r.client.Put(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST/PUT), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = plan.Number
	plan.fromBodyUnknowns(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (r *WirelessSSIDHotspot20Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WirelessSSIDHotspot20

	// Read state
	diags := req.State.Get(ctx, &state)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))
	res, err := r.client.Get(state.getPath())
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

func (r *WirelessSSIDHotspot20Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state WirelessSSIDHotspot20

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

func (r *WirelessSSIDHotspot20Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WirelessSSIDHotspot20

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
func (r *WirelessSSIDHotspot20Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: <network_id>,<number>. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import
