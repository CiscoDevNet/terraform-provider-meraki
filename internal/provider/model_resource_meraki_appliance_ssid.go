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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceSSID struct {
	Id                                  types.String                 `tfsdk:"id"`
	NetworkId                           types.String                 `tfsdk:"network_id"`
	Number                              types.String                 `tfsdk:"number"`
	AuthMode                            types.String                 `tfsdk:"auth_mode"`
	DefaultVlanId                       types.Int64                  `tfsdk:"default_vlan_id"`
	Enabled                             types.Bool                   `tfsdk:"enabled"`
	EncryptionMode                      types.String                 `tfsdk:"encryption_mode"`
	Name                                types.String                 `tfsdk:"name"`
	Psk                                 types.String                 `tfsdk:"psk"`
	PskWo                               types.String                 `tfsdk:"psk_wo"`
	PskWoVersion                        types.Int64                  `tfsdk:"psk_wo_version"`
	Visible                             types.Bool                   `tfsdk:"visible"`
	WpaEncryptionMode                   types.String                 `tfsdk:"wpa_encryption_mode"`
	DhcpEnforcedDeauthenticationEnabled types.Bool                   `tfsdk:"dhcp_enforced_deauthentication_enabled"`
	Dot11wEnabled                       types.Bool                   `tfsdk:"dot11w_enabled"`
	Dot11wRequired                      types.Bool                   `tfsdk:"dot11w_required"`
	RadiusServers                       []ApplianceSSIDRadiusServers `tfsdk:"radius_servers"`
}

type ApplianceSSIDRadiusServers struct {
	Host            types.String `tfsdk:"host"`
	Port            types.Int64  `tfsdk:"port"`
	Secret          types.String `tfsdk:"secret"`
	SecretWo        types.String `tfsdk:"secret_wo"`
	SecretWoVersion types.Int64  `tfsdk:"secret_wo_version"`
}

type ApplianceSSIDIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
	Number    types.String `tfsdk:"number"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceSSID) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/ssids/%v", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceSSID) toBody(ctx context.Context, state ApplianceSSID) string {
	body := ""
	if !data.AuthMode.IsNull() {
		body, _ = sjson.Set(body, "authMode", data.AuthMode.ValueString())
	}
	if !data.DefaultVlanId.IsNull() {
		body, _ = sjson.Set(body, "defaultVlanId", data.DefaultVlanId.ValueInt64())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.EncryptionMode.IsNull() {
		body, _ = sjson.Set(body, "encryptionMode", data.EncryptionMode.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.PskWo.IsNull() {
		body, _ = sjson.Set(body, "psk", data.PskWo.ValueString())
	} else if !data.Psk.IsNull() {
		body, _ = sjson.Set(body, "psk", data.Psk.ValueString())
	}
	if !data.Visible.IsNull() {
		body, _ = sjson.Set(body, "visible", data.Visible.ValueBool())
	}
	if !data.WpaEncryptionMode.IsNull() {
		body, _ = sjson.Set(body, "wpaEncryptionMode", data.WpaEncryptionMode.ValueString())
	}
	if !data.DhcpEnforcedDeauthenticationEnabled.IsNull() {
		body, _ = sjson.Set(body, "dhcpEnforcedDeauthentication.enabled", data.DhcpEnforcedDeauthenticationEnabled.ValueBool())
	}
	if !data.Dot11wEnabled.IsNull() {
		body, _ = sjson.Set(body, "dot11w.enabled", data.Dot11wEnabled.ValueBool())
	}
	if !data.Dot11wRequired.IsNull() {
		body, _ = sjson.Set(body, "dot11w.required", data.Dot11wRequired.ValueBool())
	}
	if len(data.RadiusServers) > 0 {
		body, _ = sjson.Set(body, "radiusServers", []interface{}{})
		for _, item := range data.RadiusServers {
			itemBody := ""
			if !item.Host.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "host", item.Host.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			if !item.SecretWo.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secret", item.SecretWo.ValueString())
			} else if !item.Secret.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secret", item.Secret.ValueString())
			}
			body, _ = sjson.SetRaw(body, "radiusServers.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPreservingNulls

// toBodyPreservingNulls walks the same writable-attribute schema as toBody but
// reads directly from the raw API response (gjson) instead of from the
// Terraform model. Unlike toBody, it preserves attributes that the API
// explicitly returned as `null` (emitting them as JSON `null` rather than
// dropping them). This is used by the singleton restoreOriginalStateOnDestroy
// path so that explicit-null fields captured during Create are restored on
// Delete. Keep this method in sync with toBody — both walk the same
// `.Attributes` schema and must agree on which fields are writable.
func (data ApplianceSSID) toBodyPreservingNulls(ctx context.Context, res meraki.Res) string {
	body := ""
	if value := res.Get("authMode"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "authMode", "null")
		} else {
			body, _ = sjson.Set(body, "authMode", value.String())
		}
	}
	if value := res.Get("defaultVlanId"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "defaultVlanId", "null")
		} else {
			body, _ = sjson.Set(body, "defaultVlanId", value.Int())
		}
	}
	if value := res.Get("enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "enabled", "null")
		} else {
			body, _ = sjson.Set(body, "enabled", value.Bool())
		}
	}
	if value := res.Get("encryptionMode"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "encryptionMode", "null")
		} else {
			body, _ = sjson.Set(body, "encryptionMode", value.String())
		}
	}
	if value := res.Get("name"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "name", "null")
		} else {
			body, _ = sjson.Set(body, "name", value.String())
		}
	}
	if value := res.Get("psk"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "psk", "null")
		} else {
			body, _ = sjson.Set(body, "psk", value.String())
		}
	}
	if value := res.Get("visible"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "visible", "null")
		} else {
			body, _ = sjson.Set(body, "visible", value.Bool())
		}
	}
	if value := res.Get("wpaEncryptionMode"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "wpaEncryptionMode", "null")
		} else {
			body, _ = sjson.Set(body, "wpaEncryptionMode", value.String())
		}
	}
	if value := res.Get("dhcpEnforcedDeauthentication.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "dhcpEnforcedDeauthentication.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "dhcpEnforcedDeauthentication.enabled", value.Bool())
		}
	}
	if value := res.Get("dot11w.enabled"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "dot11w.enabled", "null")
		} else {
			body, _ = sjson.Set(body, "dot11w.enabled", value.Bool())
		}
	}
	if value := res.Get("dot11w.required"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "dot11w.required", "null")
		} else {
			body, _ = sjson.Set(body, "dot11w.required", value.Bool())
		}
	}
	if value := res.Get("radiusServers"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "radiusServers", "null")
		} else {
			body, _ = sjson.Set(body, "radiusServers", []interface{}{})
			parent := &body
			value.ForEach(func(k, res gjson.Result) bool {
				body := ""
				if value := res.Get("host"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "host", "null")
					} else {
						body, _ = sjson.Set(body, "host", value.String())
					}
				}
				if value := res.Get("port"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "port", "null")
					} else {
						body, _ = sjson.Set(body, "port", value.Int())
					}
				}
				*parent, _ = sjson.SetRaw(*parent, "radiusServers.-1", body)
				return true
			})
		}
	}
	return body
}

// End of section. //template:end toBodyPreservingNulls

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceSSID) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("authMode"); value.Exists() && value.Value() != nil {
		data.AuthMode = types.StringValue(value.String())
	} else {
		data.AuthMode = types.StringNull()
	}
	if value := res.Get("defaultVlanId"); value.Exists() && value.Value() != nil {
		data.DefaultVlanId = types.Int64Value(value.Int())
	} else {
		data.DefaultVlanId = types.Int64Null()
	}
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("encryptionMode"); value.Exists() && value.Value() != nil {
		data.EncryptionMode = types.StringValue(value.String())
	} else {
		data.EncryptionMode = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("psk"); value.Exists() && value.Value() != nil {
		data.Psk = types.StringValue(value.String())
	} else {
		data.Psk = types.StringNull()
	}
	if value := res.Get("visible"); value.Exists() && value.Value() != nil {
		data.Visible = types.BoolValue(value.Bool())
	} else {
		data.Visible = types.BoolNull()
	}
	if value := res.Get("wpaEncryptionMode"); value.Exists() && value.Value() != nil {
		data.WpaEncryptionMode = types.StringValue(value.String())
	} else {
		data.WpaEncryptionMode = types.StringNull()
	}
	if value := res.Get("dhcpEnforcedDeauthentication.enabled"); value.Exists() && value.Value() != nil {
		data.DhcpEnforcedDeauthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.DhcpEnforcedDeauthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("dot11w.enabled"); value.Exists() && value.Value() != nil {
		data.Dot11wEnabled = types.BoolValue(value.Bool())
	} else {
		data.Dot11wEnabled = types.BoolNull()
	}
	if value := res.Get("dot11w.required"); value.Exists() && value.Value() != nil {
		data.Dot11wRequired = types.BoolValue(value.Bool())
	} else {
		data.Dot11wRequired = types.BoolNull()
	}
	if value := res.Get("radiusServers"); value.Exists() && value.Value() != nil {
		data.RadiusServers = make([]ApplianceSSIDRadiusServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceSSIDRadiusServers{}
			if value := res.Get("host"); value.Exists() && value.Value() != nil {
				data.Host = types.StringValue(value.String())
			} else {
				data.Host = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && value.Value() != nil {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Null()
			}
			(*parent).RadiusServers = append((*parent).RadiusServers, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceSSID) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("authMode"); value.Exists() && !data.AuthMode.IsNull() {
		data.AuthMode = types.StringValue(value.String())
	} else {
		data.AuthMode = types.StringNull()
	}
	if value := res.Get("defaultVlanId"); value.Exists() && !data.DefaultVlanId.IsNull() {
		data.DefaultVlanId = types.Int64Value(value.Int())
	} else {
		data.DefaultVlanId = types.Int64Null()
	}
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("encryptionMode"); value.Exists() && !data.EncryptionMode.IsNull() {
		data.EncryptionMode = types.StringValue(value.String())
	} else {
		data.EncryptionMode = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("psk"); value.Exists() && !data.Psk.IsNull() {
		data.Psk = types.StringValue(value.String())
	} else {
		data.Psk = types.StringNull()
	}
	if value := res.Get("visible"); value.Exists() && !data.Visible.IsNull() {
		data.Visible = types.BoolValue(value.Bool())
	} else {
		data.Visible = types.BoolNull()
	}
	if value := res.Get("wpaEncryptionMode"); value.Exists() && !data.WpaEncryptionMode.IsNull() {
		data.WpaEncryptionMode = types.StringValue(value.String())
	} else {
		data.WpaEncryptionMode = types.StringNull()
	}
	if value := res.Get("dhcpEnforcedDeauthentication.enabled"); value.Exists() && !data.DhcpEnforcedDeauthenticationEnabled.IsNull() {
		data.DhcpEnforcedDeauthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.DhcpEnforcedDeauthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("dot11w.enabled"); value.Exists() && !data.Dot11wEnabled.IsNull() {
		data.Dot11wEnabled = types.BoolValue(value.Bool())
	} else {
		data.Dot11wEnabled = types.BoolNull()
	}
	if value := res.Get("dot11w.required"); value.Exists() && !data.Dot11wRequired.IsNull() {
		data.Dot11wRequired = types.BoolValue(value.Bool())
	} else {
		data.Dot11wRequired = types.BoolNull()
	}
	{
		l := len(res.Get("radiusServers").Array())
		tflog.Debug(ctx, fmt.Sprintf("radiusServers array resizing from %d to %d", len(data.RadiusServers), l))
		if len(data.RadiusServers) > l {
			data.RadiusServers = data.RadiusServers[:l]
		}
	}
	for i := range data.RadiusServers {
		parent := &data
		data := (*parent).RadiusServers[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("radiusServers.%d", i))
		if value := res.Get("host"); value.Exists() && !data.Host.IsNull() {
			data.Host = types.StringValue(value.String())
		} else {
			data.Host = types.StringNull()
		}
		if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
			data.Port = types.Int64Value(value.Int())
		} else {
			data.Port = types.Int64Null()
		}
		(*parent).RadiusServers[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceSSID) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceSSIDIdentity) toIdentity(ctx context.Context, plan *ApplianceSSID) {
	data.NetworkId = plan.NetworkId
	data.Number = plan.Number
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceSSID) fromIdentity(ctx context.Context, identity *ApplianceSSIDIdentity) {
	data.NetworkId = identity.NetworkId
	data.Number = identity.Number
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceSSID) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
