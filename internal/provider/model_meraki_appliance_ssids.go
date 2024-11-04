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
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceSSIDs struct {
	NetworkId types.String          `tfsdk:"network_id"`
	Items     []ApplianceSSIDsItems `tfsdk:"items"`
}

type ApplianceSSIDsItems struct {
	Id                types.String                  `tfsdk:"id"`
	AuthMode          types.String                  `tfsdk:"auth_mode"`
	DefaultVlanId     types.Int64                   `tfsdk:"default_vlan_id"`
	Enabled           types.Bool                    `tfsdk:"enabled"`
	EncryptionMode    types.String                  `tfsdk:"encryption_mode"`
	Name              types.String                  `tfsdk:"name"`
	Number            types.Int64                   `tfsdk:"number"`
	Visible           types.Bool                    `tfsdk:"visible"`
	WpaEncryptionMode types.String                  `tfsdk:"wpa_encryption_mode"`
	RadiusServers     []ApplianceSSIDsRadiusServers `tfsdk:"radius_servers"`
}

type ApplianceSSIDsRadiusServers struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceSSIDs) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/ssids", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceSSIDs) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ApplianceSSIDsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ApplianceSSIDsItems{}
		data.Id = types.StringValue(res.Get("").String())
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
		if value := res.Get("number"); value.Exists() && value.Value() != nil {
			data.Number = types.Int64Value(value.Int())
		} else {
			data.Number = types.Int64Null()
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
		if value := res.Get("radiusServers"); value.Exists() && value.Value() != nil {
			data.RadiusServers = make([]ApplianceSSIDsRadiusServers, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ApplianceSSIDsRadiusServers{}
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
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
