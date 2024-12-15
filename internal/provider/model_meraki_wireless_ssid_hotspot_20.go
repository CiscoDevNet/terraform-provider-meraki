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
	"slices"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessSSIDHotspot20 struct {
	Id                types.String                     `tfsdk:"id"`
	NetworkId         types.String                     `tfsdk:"network_id"`
	Number            types.String                     `tfsdk:"number"`
	Enabled           types.Bool                       `tfsdk:"enabled"`
	NetworkAccessType types.String                     `tfsdk:"network_access_type"`
	OperatorName      types.String                     `tfsdk:"operator_name"`
	VenueName         types.String                     `tfsdk:"venue_name"`
	VenueType         types.String                     `tfsdk:"venue_type"`
	Domains           types.Set                        `tfsdk:"domains"`
	MccMncs           []WirelessSSIDHotspot20MccMncs   `tfsdk:"mcc_mncs"`
	NaiRealms         []WirelessSSIDHotspot20NaiRealms `tfsdk:"nai_realms"`
	RoamConsortOis    types.Set                        `tfsdk:"roam_consort_ois"`
}

type WirelessSSIDHotspot20MccMncs struct {
	Mcc types.String `tfsdk:"mcc"`
	Mnc types.String `tfsdk:"mnc"`
}

type WirelessSSIDHotspot20NaiRealms struct {
	Format  types.String                            `tfsdk:"format"`
	Realm   types.String                            `tfsdk:"realm"`
	Methods []WirelessSSIDHotspot20NaiRealmsMethods `tfsdk:"methods"`
}

type WirelessSSIDHotspot20NaiRealmsMethods struct {
	Id                                              types.String `tfsdk:"id"`
	AuthenticationTypesNonEapInnerAuthentication    types.Set    `tfsdk:"authentication_types_non_eap_inner_authentication"`
	AuthenticationTypesEapInnerAuthentication       types.Set    `tfsdk:"authentication_types_eap_inner_authentication"`
	AuthenticationTypesCredentials                  types.Set    `tfsdk:"authentication_types_credentials"`
	AuthenticationTypesTunneledEapMethodCredentials types.Set    `tfsdk:"authentication_types_tunneled_eap_method_credentials"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessSSIDHotspot20) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/hotspot20", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSSIDHotspot20) toBody(ctx context.Context, state WirelessSSIDHotspot20) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.NetworkAccessType.IsNull() {
		body, _ = sjson.Set(body, "networkAccessType", data.NetworkAccessType.ValueString())
	}
	if !data.OperatorName.IsNull() {
		body, _ = sjson.Set(body, "operator.name", data.OperatorName.ValueString())
	}
	if !data.VenueName.IsNull() {
		body, _ = sjson.Set(body, "venue.name", data.VenueName.ValueString())
	}
	if !data.VenueType.IsNull() {
		body, _ = sjson.Set(body, "venue.type", data.VenueType.ValueString())
	}
	if !data.Domains.IsNull() {
		var values []string
		data.Domains.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "domains", values)
	}
	if len(data.MccMncs) > 0 {
		body, _ = sjson.Set(body, "mccMncs", []interface{}{})
		for _, item := range data.MccMncs {
			itemBody := ""
			if !item.Mcc.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "mcc", item.Mcc.ValueString())
			}
			if !item.Mnc.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "mnc", item.Mnc.ValueString())
			}
			body, _ = sjson.SetRaw(body, "mccMncs.-1", itemBody)
		}
	}
	if len(data.NaiRealms) > 0 {
		body, _ = sjson.Set(body, "naiRealms", []interface{}{})
		for _, item := range data.NaiRealms {
			itemBody := ""
			if !item.Format.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "format", item.Format.ValueString())
			}
			if !item.Realm.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "realm", item.Realm.ValueString())
			}
			if len(item.Methods) > 0 {
				itemBody, _ = sjson.Set(itemBody, "methods", []interface{}{})
				for _, childItem := range item.Methods {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.AuthenticationTypesNonEapInnerAuthentication.IsNull() {
						var values []string
						childItem.AuthenticationTypesNonEapInnerAuthentication.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "authenticationTypes.nonEapInnerAuthentication", values)
					}
					if !childItem.AuthenticationTypesEapInnerAuthentication.IsNull() {
						var values []string
						childItem.AuthenticationTypesEapInnerAuthentication.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "authenticationTypes.eapInnerAuthentication", values)
					}
					if !childItem.AuthenticationTypesCredentials.IsNull() {
						var values []string
						childItem.AuthenticationTypesCredentials.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "authenticationTypes.credentials", values)
					}
					if !childItem.AuthenticationTypesTunneledEapMethodCredentials.IsNull() {
						var values []string
						childItem.AuthenticationTypesTunneledEapMethodCredentials.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "authenticationTypes.tunneledEapMethodCredentials", values)
					}
					itemBody, _ = sjson.SetRaw(itemBody, "methods.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "naiRealms.-1", itemBody)
		}
	}
	if !data.RoamConsortOis.IsNull() {
		var values []string
		data.RoamConsortOis.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "roamConsortOis", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSSIDHotspot20) fromBody(ctx context.Context, res meraki.Res) {
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
		data.MccMncs = make([]WirelessSSIDHotspot20MccMncs, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDHotspot20MccMncs{}
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
		data.NaiRealms = make([]WirelessSSIDHotspot20NaiRealms, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDHotspot20NaiRealms{}
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
				data.Methods = make([]WirelessSSIDHotspot20NaiRealmsMethods, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := WirelessSSIDHotspot20NaiRealmsMethods{}
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessSSIDHotspot20) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("networkAccessType"); value.Exists() && !data.NetworkAccessType.IsNull() {
		data.NetworkAccessType = types.StringValue(value.String())
	} else {
		data.NetworkAccessType = types.StringNull()
	}
	if value := res.Get("operator.name"); value.Exists() && !data.OperatorName.IsNull() {
		data.OperatorName = types.StringValue(value.String())
	} else {
		data.OperatorName = types.StringNull()
	}
	if value := res.Get("venue.name"); value.Exists() && !data.VenueName.IsNull() {
		data.VenueName = types.StringValue(value.String())
	} else {
		data.VenueName = types.StringNull()
	}
	if value := res.Get("venue.type"); value.Exists() && !data.VenueType.IsNull() {
		data.VenueType = types.StringValue(value.String())
	} else {
		data.VenueType = types.StringNull()
	}
	if value := res.Get("domains"); value.Exists() && !data.Domains.IsNull() {
		data.Domains = helpers.GetStringSet(value.Array())
	} else {
		data.Domains = types.SetNull(types.StringType)
	}
	for i := 0; i < len(data.MccMncs); i++ {
		keys := [...]string{"mcc", "mnc"}
		keyValues := [...]string{data.MccMncs[i].Mcc.ValueString(), data.MccMncs[i].Mnc.ValueString()}

		parent := &data
		data := (*parent).MccMncs[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("mccMncs").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing MccMncs[%d] = %+v",
				i,
				(*parent).MccMncs[i],
			))
			(*parent).MccMncs = slices.Delete((*parent).MccMncs, i, i+1)
			i--

			continue
		}
		if value := res.Get("mcc"); value.Exists() && !data.Mcc.IsNull() {
			data.Mcc = types.StringValue(value.String())
		} else {
			data.Mcc = types.StringNull()
		}
		if value := res.Get("mnc"); value.Exists() && !data.Mnc.IsNull() {
			data.Mnc = types.StringValue(value.String())
		} else {
			data.Mnc = types.StringNull()
		}
		(*parent).MccMncs[i] = data
	}
	for i := 0; i < len(data.NaiRealms); i++ {
		keys := [...]string{"realm"}
		keyValues := [...]string{data.NaiRealms[i].Realm.ValueString()}

		parent := &data
		data := (*parent).NaiRealms[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("naiRealms").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing NaiRealms[%d] = %+v",
				i,
				(*parent).NaiRealms[i],
			))
			(*parent).NaiRealms = slices.Delete((*parent).NaiRealms, i, i+1)
			i--

			continue
		}
		if value := res.Get("format"); value.Exists() && !data.Format.IsNull() {
			data.Format = types.StringValue(value.String())
		} else {
			data.Format = types.StringNull()
		}
		if value := res.Get("realm"); value.Exists() && !data.Realm.IsNull() {
			data.Realm = types.StringValue(value.String())
		} else {
			data.Realm = types.StringNull()
		}
		for i := 0; i < len(data.Methods); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Methods[i].Id.ValueString()}

			parent := &data
			data := (*parent).Methods[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("methods").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing Methods[%d] = %+v",
					i,
					(*parent).Methods[i],
				))
				(*parent).Methods = slices.Delete((*parent).Methods, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("authenticationTypes.nonEapInnerAuthentication"); value.Exists() && !data.AuthenticationTypesNonEapInnerAuthentication.IsNull() {
				data.AuthenticationTypesNonEapInnerAuthentication = helpers.GetStringSet(value.Array())
			} else {
				data.AuthenticationTypesNonEapInnerAuthentication = types.SetNull(types.StringType)
			}
			if value := res.Get("authenticationTypes.eapInnerAuthentication"); value.Exists() && !data.AuthenticationTypesEapInnerAuthentication.IsNull() {
				data.AuthenticationTypesEapInnerAuthentication = helpers.GetStringSet(value.Array())
			} else {
				data.AuthenticationTypesEapInnerAuthentication = types.SetNull(types.StringType)
			}
			if value := res.Get("authenticationTypes.credentials"); value.Exists() && !data.AuthenticationTypesCredentials.IsNull() {
				data.AuthenticationTypesCredentials = helpers.GetStringSet(value.Array())
			} else {
				data.AuthenticationTypesCredentials = types.SetNull(types.StringType)
			}
			if value := res.Get("authenticationTypes.tunneledEapMethodCredentials"); value.Exists() && !data.AuthenticationTypesTunneledEapMethodCredentials.IsNull() {
				data.AuthenticationTypesTunneledEapMethodCredentials = helpers.GetStringSet(value.Array())
			} else {
				data.AuthenticationTypesTunneledEapMethodCredentials = types.SetNull(types.StringType)
			}
			(*parent).Methods[i] = data
		}
		(*parent).NaiRealms[i] = data
	}
	if value := res.Get("roamConsortOis"); value.Exists() && !data.RoamConsortOis.IsNull() {
		data.RoamConsortOis = helpers.GetStringSet(value.Array())
	} else {
		data.RoamConsortOis = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessSSIDHotspot20) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns
