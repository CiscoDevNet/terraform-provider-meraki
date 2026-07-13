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
// It emits a separate model - DataSourceWirelessAlternateManagementInterface - used only by the data source, always including every
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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessAlternateManagementInterface struct {
	Id           types.String                                                 `tfsdk:"id"`
	NetworkId    types.String                                                 `tfsdk:"network_id"`
	Enabled      types.Bool                                                   `tfsdk:"enabled"`
	VlanId       types.Int64                                                  `tfsdk:"vlan_id"`
	AccessPoints []DataSourceWirelessAlternateManagementInterfaceAccessPoints `tfsdk:"access_points"`
	Protocols    types.Set                                                    `tfsdk:"protocols"`
}

type DataSourceWirelessAlternateManagementInterfaceAccessPoints struct {
	AlternateManagementIp types.String `tfsdk:"alternate_management_ip"`
	Dns1                  types.String `tfsdk:"dns1"`
	Dns2                  types.String `tfsdk:"dns2"`
	Gateway               types.String `tfsdk:"gateway"`
	Serial                types.String `tfsdk:"serial"`
	SubnetMask            types.String `tfsdk:"subnet_mask"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessAlternateManagementInterface) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/alternateManagementInterface", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessAlternateManagementInterface) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("vlanId"); value.Exists() && value.Value() != nil {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("accessPoints"); value.Exists() && value.Value() != nil {
		data.AccessPoints = make([]DataSourceWirelessAlternateManagementInterfaceAccessPoints, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessAlternateManagementInterfaceAccessPoints{}
			if value := res.Get("alternateManagementIp"); value.Exists() && value.Value() != nil {
				data.AlternateManagementIp = types.StringValue(value.String())
			} else {
				data.AlternateManagementIp = types.StringNull()
			}
			if value := res.Get("dns1"); value.Exists() && value.Value() != nil {
				data.Dns1 = types.StringValue(value.String())
			} else {
				data.Dns1 = types.StringNull()
			}
			if value := res.Get("dns2"); value.Exists() && value.Value() != nil {
				data.Dns2 = types.StringValue(value.String())
			} else {
				data.Dns2 = types.StringNull()
			}
			if value := res.Get("gateway"); value.Exists() && value.Value() != nil {
				data.Gateway = types.StringValue(value.String())
			} else {
				data.Gateway = types.StringNull()
			}
			if value := res.Get("serial"); value.Exists() && value.Value() != nil {
				data.Serial = types.StringValue(value.String())
			} else {
				data.Serial = types.StringNull()
			}
			if value := res.Get("subnetMask"); value.Exists() && value.Value() != nil {
				data.SubnetMask = types.StringValue(value.String())
			} else {
				data.SubnetMask = types.StringNull()
			}
			(*parent).AccessPoints = append((*parent).AccessPoints, data)
			return true
		})
	}
	if value := res.Get("protocols"); value.Exists() && value.Value() != nil {
		data.Protocols = helpers.GetStringSet(value.Array())
	} else {
		data.Protocols = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody
