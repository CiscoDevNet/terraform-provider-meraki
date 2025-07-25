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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceOrganizationDevices struct {
	OrganizationId types.String                         `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationDevicesItems `tfsdk:"items"`
}

type DataSourceOrganizationDevicesItems struct {
	Id          types.String                           `tfsdk:"id"`
	Address     types.String                           `tfsdk:"address"`
	Firmware    types.String                           `tfsdk:"firmware"`
	Imei        types.String                           `tfsdk:"imei"`
	LanIp       types.String                           `tfsdk:"lan_ip"`
	Lat         types.Float64                          `tfsdk:"lat"`
	Lng         types.Float64                          `tfsdk:"lng"`
	Mac         types.String                           `tfsdk:"mac"`
	Model       types.String                           `tfsdk:"model"`
	Name        types.String                           `tfsdk:"name"`
	NetworkId   types.String                           `tfsdk:"network_id"`
	Notes       types.String                           `tfsdk:"notes"`
	ProductType types.String                           `tfsdk:"product_type"`
	Serial      types.String                           `tfsdk:"serial"`
	Details     []DataSourceOrganizationDevicesDetails `tfsdk:"details"`
	Tags        types.List                             `tfsdk:"tags"`
}

type DataSourceOrganizationDevicesDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationDevices) getPath() string {
	return fmt.Sprintf("/organizations/%v/devices", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationDevices) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationDevicesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationDevicesItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("address"); value.Exists() && value.Value() != nil {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
		}
		if value := res.Get("firmware"); value.Exists() && value.Value() != nil {
			data.Firmware = types.StringValue(value.String())
		} else {
			data.Firmware = types.StringNull()
		}
		if value := res.Get("imei"); value.Exists() && value.Value() != nil {
			data.Imei = types.StringValue(value.String())
		} else {
			data.Imei = types.StringNull()
		}
		if value := res.Get("lanIp"); value.Exists() && value.Value() != nil {
			data.LanIp = types.StringValue(value.String())
		} else {
			data.LanIp = types.StringNull()
		}
		if value := res.Get("lat"); value.Exists() && value.Value() != nil {
			data.Lat = types.Float64Value(value.Float())
		} else {
			data.Lat = types.Float64Null()
		}
		if value := res.Get("lng"); value.Exists() && value.Value() != nil {
			data.Lng = types.Float64Value(value.Float())
		} else {
			data.Lng = types.Float64Null()
		}
		if value := res.Get("mac"); value.Exists() && value.Value() != nil {
			data.Mac = types.StringValue(value.String())
		} else {
			data.Mac = types.StringNull()
		}
		if value := res.Get("model"); value.Exists() && value.Value() != nil {
			data.Model = types.StringValue(value.String())
		} else {
			data.Model = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("networkId"); value.Exists() && value.Value() != nil {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("notes"); value.Exists() && value.Value() != nil {
			data.Notes = types.StringValue(value.String())
		} else {
			data.Notes = types.StringNull()
		}
		if value := res.Get("productType"); value.Exists() && value.Value() != nil {
			data.ProductType = types.StringValue(value.String())
		} else {
			data.ProductType = types.StringNull()
		}
		if value := res.Get("serial"); value.Exists() && value.Value() != nil {
			data.Serial = types.StringValue(value.String())
		} else {
			data.Serial = types.StringNull()
		}
		if value := res.Get("details"); value.Exists() && value.Value() != nil {
			data.Details = make([]DataSourceOrganizationDevicesDetails, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationDevicesDetails{}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("value"); value.Exists() && value.Value() != nil {
					data.Value = types.StringValue(value.String())
				} else {
					data.Value = types.StringNull()
				}
				(*parent).Details = append((*parent).Details, data)
				return true
			})
		}
		if value := res.Get("tags"); value.Exists() && value.Value() != nil {
			data.Tags = helpers.GetStringList(value.Array())
		} else {
			data.Tags = types.ListNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
