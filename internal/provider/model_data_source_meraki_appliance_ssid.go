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
// It emits a separate model - DataSourceApplianceSSID - used only by the data source, always including every
// attribute (unlike model.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceApplianceSSID struct {
	Id                                  types.String                           `tfsdk:"id"`
	NetworkId                           types.String                           `tfsdk:"network_id"`
	Number                              types.String                           `tfsdk:"number"`
	AuthMode                            types.String                           `tfsdk:"auth_mode"`
	DefaultVlanId                       types.Int64                            `tfsdk:"default_vlan_id"`
	Enabled                             types.Bool                             `tfsdk:"enabled"`
	EncryptionMode                      types.String                           `tfsdk:"encryption_mode"`
	Name                                types.String                           `tfsdk:"name"`
	Psk                                 types.String                           `tfsdk:"psk"`
	PskWo                               types.String                           `tfsdk:"psk_wo"`
	PskWoVersion                        types.Int64                            `tfsdk:"psk_wo_version"`
	Visible                             types.Bool                             `tfsdk:"visible"`
	WpaEncryptionMode                   types.String                           `tfsdk:"wpa_encryption_mode"`
	DhcpEnforcedDeauthenticationEnabled types.Bool                             `tfsdk:"dhcp_enforced_deauthentication_enabled"`
	Dot11wEnabled                       types.Bool                             `tfsdk:"dot11w_enabled"`
	Dot11wRequired                      types.Bool                             `tfsdk:"dot11w_required"`
	RadiusServers                       []DataSourceApplianceSSIDRadiusServers `tfsdk:"radius_servers"`
}

type DataSourceApplianceSSIDRadiusServers struct {
	Host            types.String `tfsdk:"host"`
	Port            types.Int64  `tfsdk:"port"`
	Secret          types.String `tfsdk:"secret"`
	SecretWo        types.String `tfsdk:"secret_wo"`
	SecretWoVersion types.Int64  `tfsdk:"secret_wo_version"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceSSID) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/ssids/%v", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceSSID) fromBody(ctx context.Context, res meraki.Res) {
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
		data.RadiusServers = make([]DataSourceApplianceSSIDRadiusServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceSSIDRadiusServers{}
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
