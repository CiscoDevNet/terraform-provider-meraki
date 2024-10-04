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
	Visible                             types.Bool                   `tfsdk:"visible"`
	WpaEncryptionMode                   types.String                 `tfsdk:"wpa_encryption_mode"`
	DhcpEnforcedDeauthenticationEnabled types.Bool                   `tfsdk:"dhcp_enforced_deauthentication_enabled"`
	Dot11wEnabled                       types.Bool                   `tfsdk:"dot11w_enabled"`
	Dot11wRequired                      types.Bool                   `tfsdk:"dot11w_required"`
	RadiusServers                       []ApplianceSSIDRadiusServers `tfsdk:"radius_servers"`
}

type ApplianceSSIDRadiusServers struct {
	Host   types.String `tfsdk:"host"`
	Port   types.Int64  `tfsdk:"port"`
	Secret types.String `tfsdk:"secret"`
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
	if !data.Psk.IsNull() {
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
			if !item.Secret.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secret", item.Secret.ValueString())
			}
			body, _ = sjson.SetRaw(body, "radiusServers.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

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
