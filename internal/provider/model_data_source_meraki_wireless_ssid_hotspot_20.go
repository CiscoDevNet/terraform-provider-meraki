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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSourceWirelessSSIDHotspot20 - used only by the data source, always including every
// attribute (unlike model_resource.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model_resource.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model_resource.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessSSIDHotspot20 struct {
	Id                types.String                               `tfsdk:"id"`
	NetworkId         types.String                               `tfsdk:"network_id"`
	Number            types.String                               `tfsdk:"number"`
	Enabled           types.Bool                                 `tfsdk:"enabled"`
	NetworkAccessType types.String                               `tfsdk:"network_access_type"`
	OperatorName      types.String                               `tfsdk:"operator_name"`
	VenueName         types.String                               `tfsdk:"venue_name"`
	VenueType         types.String                               `tfsdk:"venue_type"`
	Domains           types.Set                                  `tfsdk:"domains"`
	MccMncs           []DataSourceWirelessSSIDHotspot20MccMncs   `tfsdk:"mcc_mncs"`
	NaiRealms         []DataSourceWirelessSSIDHotspot20NaiRealms `tfsdk:"nai_realms"`
	RoamConsortOis    types.Set                                  `tfsdk:"roam_consort_ois"`
}

type DataSourceWirelessSSIDHotspot20MccMncs struct {
	Mcc types.String `tfsdk:"mcc"`
	Mnc types.String `tfsdk:"mnc"`
}

type DataSourceWirelessSSIDHotspot20NaiRealms struct {
	Format  types.String                                      `tfsdk:"format"`
	Realm   types.String                                      `tfsdk:"realm"`
	Methods []DataSourceWirelessSSIDHotspot20NaiRealmsMethods `tfsdk:"methods"`
}

type DataSourceWirelessSSIDHotspot20NaiRealmsMethods struct {
	Id                                              types.String `tfsdk:"id"`
	AuthenticationTypesNonEapInnerAuthentication    types.Set    `tfsdk:"authentication_types_non_eap_inner_authentication"`
	AuthenticationTypesEapInnerAuthentication       types.Set    `tfsdk:"authentication_types_eap_inner_authentication"`
	AuthenticationTypesCredentials                  types.Set    `tfsdk:"authentication_types_credentials"`
	AuthenticationTypesTunneledEapMethodCredentials types.Set    `tfsdk:"authentication_types_tunneled_eap_method_credentials"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessSSIDHotspot20) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/hotspot20", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessSSIDHotspot20) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("networkAccessType"); value.Exists() && value.Value() != nil {
		data.NetworkAccessType = types.StringValue(value.String())
	} else {
		data.NetworkAccessType = types.StringNull()
	}
	if value := res.Get("operator.name"); value.Exists() && value.Value() != nil {
		data.OperatorName = types.StringValue(value.String())
	} else {
		data.OperatorName = types.StringNull()
	}
	if value := res.Get("venue.name"); value.Exists() && value.Value() != nil {
		data.VenueName = types.StringValue(value.String())
	} else {
		data.VenueName = types.StringNull()
	}
	if value := res.Get("venue.type"); value.Exists() && value.Value() != nil {
		data.VenueType = types.StringValue(value.String())
	} else {
		data.VenueType = types.StringNull()
	}
	if value := res.Get("domains"); value.Exists() && value.Value() != nil {
		data.Domains = helpers.GetStringSet(value.Array())
	} else {
		data.Domains = types.SetNull(types.StringType)
	}
	if value := res.Get("mccMncs"); value.Exists() && value.Value() != nil {
		data.MccMncs = make([]DataSourceWirelessSSIDHotspot20MccMncs, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDHotspot20MccMncs{}
			if value := res.Get("mcc"); value.Exists() && value.Value() != nil {
				data.Mcc = types.StringValue(value.String())
			} else {
				data.Mcc = types.StringNull()
			}
			if value := res.Get("mnc"); value.Exists() && value.Value() != nil {
				data.Mnc = types.StringValue(value.String())
			} else {
				data.Mnc = types.StringNull()
			}
			(*parent).MccMncs = append((*parent).MccMncs, data)
			return true
		})
	}
	if value := res.Get("naiRealms"); value.Exists() && value.Value() != nil {
		data.NaiRealms = make([]DataSourceWirelessSSIDHotspot20NaiRealms, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDHotspot20NaiRealms{}
			if value := res.Get("format"); value.Exists() && value.Value() != nil {
				data.Format = types.StringValue(value.String())
			} else {
				data.Format = types.StringNull()
			}
			if value := res.Get("realm"); value.Exists() && value.Value() != nil {
				data.Realm = types.StringValue(value.String())
			} else {
				data.Realm = types.StringNull()
			}
			if value := res.Get("methods"); value.Exists() && value.Value() != nil {
				data.Methods = make([]DataSourceWirelessSSIDHotspot20NaiRealmsMethods, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DataSourceWirelessSSIDHotspot20NaiRealmsMethods{}
					if value := res.Get("id"); value.Exists() && value.Value() != nil {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("authenticationTypes.nonEapInnerAuthentication"); value.Exists() && value.Value() != nil {
						data.AuthenticationTypesNonEapInnerAuthentication = helpers.GetStringSet(value.Array())
					} else {
						data.AuthenticationTypesNonEapInnerAuthentication = types.SetNull(types.StringType)
					}
					if value := res.Get("authenticationTypes.eapInnerAuthentication"); value.Exists() && value.Value() != nil {
						data.AuthenticationTypesEapInnerAuthentication = helpers.GetStringSet(value.Array())
					} else {
						data.AuthenticationTypesEapInnerAuthentication = types.SetNull(types.StringType)
					}
					if value := res.Get("authenticationTypes.credentials"); value.Exists() && value.Value() != nil {
						data.AuthenticationTypesCredentials = helpers.GetStringSet(value.Array())
					} else {
						data.AuthenticationTypesCredentials = types.SetNull(types.StringType)
					}
					if value := res.Get("authenticationTypes.tunneledEapMethodCredentials"); value.Exists() && value.Value() != nil {
						data.AuthenticationTypesTunneledEapMethodCredentials = helpers.GetStringSet(value.Array())
					} else {
						data.AuthenticationTypesTunneledEapMethodCredentials = types.SetNull(types.StringType)
					}
					(*parent).Methods = append((*parent).Methods, data)
					return true
				})
			}
			(*parent).NaiRealms = append((*parent).NaiRealms, data)
			return true
		})
	}
	if value := res.Get("roamConsortOis"); value.Exists() && value.Value() != nil {
		data.RoamConsortOis = helpers.GetStringSet(value.Array())
	} else {
		data.RoamConsortOis = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody
