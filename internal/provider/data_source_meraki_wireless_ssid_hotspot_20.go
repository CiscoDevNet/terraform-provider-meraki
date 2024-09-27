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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessSSIDHotspot20DataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessSSIDHotspot20DataSource{}
)

func NewWirelessSSIDHotspot20DataSource() datasource.DataSource {
	return &WirelessSSIDHotspot20DataSource{}
}

type WirelessSSIDHotspot20DataSource struct {
	client *meraki.Client
}

func (d *WirelessSSIDHotspot20DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssid_hotspot_20"
}

func (d *WirelessSSIDHotspot20DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Wireless SSID Hotspot 2.0` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: "Wireless SSID number",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not Hotspot 2.0 for this SSID is enabled",
				Computed:            true,
			},
			"network_access_type": schema.StringAttribute{
				MarkdownDescription: "The network type of this SSID (`Private network`, `Private network with guest access`, `Chargeable public network`, `Free public network`, `Personal device network`, `Emergency services only network`, `Test or experimental`, `Wildcard`)",
				Computed:            true,
			},
			"operator_name": schema.StringAttribute{
				MarkdownDescription: "Operator name",
				Computed:            true,
			},
			"venue_name": schema.StringAttribute{
				MarkdownDescription: "Venue name",
				Computed:            true,
			},
			"venue_type": schema.StringAttribute{
				MarkdownDescription: "Venue type (`Unspecified`, `Unspecified Assembly`, `Arena`, `Stadium`, `Passenger Terminal`, `Amphitheater`, `Amusement Park`, `Place of Worship`, `Convention Center`, `Library`, `Museum`, `Restaurant`, `Theater`, `Bar`, `Coffee Shop`, `Zoo or Aquarium`, `Emergency Coordination Center`, `Unspecified Business`, `Doctor or Dentist office`, `Bank`, `Fire Station`, `Police Station`, `Post Office`, `Professional Office`, `Research and Development Facility`, `Attorney Office`, `Unspecified Educational`, `School, Primary`, `School, Secondary`, `University or College`, `Unspecified Factory and Industrial`, `Factory`, `Unspecified Institutional`, `Hospital`, `Long-Term Care Facility`, `Alcohol and Drug Rehabilitation Center`, `Group Home`, `Prison or Jail`, `Unspecified Mercantile`, `Retail Store`, `Grocery Market`, `Automotive Service Station`, `Shopping Mall`, `Gas Station`, `Unspecified Residential`, `Private Residence`, `Hotel or Motel`, `Dormitory`, `Boarding House`, `Unspecified Storage`, `Unspecified Utility and Miscellaneous`, `Unspecified Vehicular`, `Automobile or Truck`, `Airplane`, `Bus`, `Ferry`, `Ship or Boat`, `Train`, `Motor Bike`, `Unspecified Outdoor`, `Muni-mesh Network`, `City Park`, `Rest Area`, `Traffic Control`, `Bus Stop`, `Kiosk`)",
				Computed:            true,
			},
			"domains": schema.SetAttribute{
				MarkdownDescription: "An array of domain names",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"mcc_mncs": schema.ListNestedAttribute{
				MarkdownDescription: "An array of MCC/MNC pairs",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mcc": schema.StringAttribute{
							MarkdownDescription: "MCC value",
							Computed:            true,
						},
						"mnc": schema.StringAttribute{
							MarkdownDescription: "MNC value",
							Computed:            true,
						},
					},
				},
			},
			"nai_realms": schema.ListNestedAttribute{
				MarkdownDescription: "An array of NAI realms",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"format": schema.StringAttribute{
							MarkdownDescription: "The format for the realm (`1` or `0`)",
							Computed:            true,
						},
						"realm": schema.StringAttribute{
							MarkdownDescription: "The name of the realm",
							Computed:            true,
						},
						"methods": schema.ListNestedAttribute{
							MarkdownDescription: "An array of EAP methods for the realm.",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "ID of method",
										Computed:            true,
									},
									"authentication_types_non_eap_inner_authentication": schema.SetAttribute{
										MarkdownDescription: "An array of autentication types. Possible values are `Reserved`, `PAP`, `CHAP`, `MSCHAP`, `MSCHAPV2`.",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"authentication_types_eap_inner_authentication": schema.SetAttribute{
										MarkdownDescription: "An array of autentication types. Possible values are `EAP-TLS`, `EAP-SIM`, `EAP-AKA`, `EAP-TTLS with MSCHAPv2`.",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"authentication_types_credentials": schema.SetAttribute{
										MarkdownDescription: "An array of autentication types. Possible values are `SIM`, `USIM`, `NFC Secure Element`, `Hardware Token`, `Softoken`, `Certificate`, `username/password`, `none`, `Reserved`, `Vendor Specific`.",
										ElementType:         types.StringType,
										Computed:            true,
									},
									"authentication_types_tunneled_eap_method_credentials": schema.SetAttribute{
										MarkdownDescription: "An array of autentication types. Possible values are `SIM`, `USIM`, `NFC Secure Element`, `Hardware Token`, `Softoken`, `Certificate`, `username/password`, `Reserved`, `Anonymous`, `Vendor Specific`.",
										ElementType:         types.StringType,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
			"roam_consort_ois": schema.SetAttribute{
				MarkdownDescription: "An array of roaming consortium OIs (hexadecimal number 3-5 octets in length)",
				ElementType:         types.StringType,
				Computed:            true,
			},
		},
	}
}

func (d *WirelessSSIDHotspot20DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessSSIDHotspot20DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessSSIDHotspot20

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res gjson.Result
	var err error

	if !res.Exists() {
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)
	config.Id = config.Number

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
