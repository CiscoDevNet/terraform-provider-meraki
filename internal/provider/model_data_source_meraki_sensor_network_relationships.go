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

type DataSourceSensorNetworkRelationships struct {
	NetworkId types.String                                `tfsdk:"network_id"`
	Items     []DataSourceSensorNetworkRelationshipsItems `tfsdk:"items"`
}

type DataSourceSensorNetworkRelationshipsItems struct {
	Id                                    types.String                                                                `tfsdk:"id"`
	DeviceName                            types.String                                                                `tfsdk:"device_name"`
	DeviceProductType                     types.String                                                                `tfsdk:"device_product_type"`
	DeviceSerial                          types.String                                                                `tfsdk:"device_serial"`
	RelationshipsLivestreamRelatedDevices []DataSourceSensorNetworkRelationshipsRelationshipsLivestreamRelatedDevices `tfsdk:"relationships_livestream_related_devices"`
}

type DataSourceSensorNetworkRelationshipsRelationshipsLivestreamRelatedDevices struct {
	ProductType types.String `tfsdk:"product_type"`
	Serial      types.String `tfsdk:"serial"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceSensorNetworkRelationships) getPath() string {
	return fmt.Sprintf("/networks/%v/sensor/relationships", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceSensorNetworkRelationships) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceSensorNetworkRelationshipsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceSensorNetworkRelationshipsItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("device.name"); value.Exists() && value.Value() != nil {
			data.DeviceName = types.StringValue(value.String())
		} else {
			data.DeviceName = types.StringNull()
		}
		if value := res.Get("device.productType"); value.Exists() && value.Value() != nil {
			data.DeviceProductType = types.StringValue(value.String())
		} else {
			data.DeviceProductType = types.StringNull()
		}
		if value := res.Get("device.serial"); value.Exists() && value.Value() != nil {
			data.DeviceSerial = types.StringValue(value.String())
		} else {
			data.DeviceSerial = types.StringNull()
		}
		if value := res.Get("relationships.livestream.relatedDevices"); value.Exists() && value.Value() != nil {
			data.RelationshipsLivestreamRelatedDevices = make([]DataSourceSensorNetworkRelationshipsRelationshipsLivestreamRelatedDevices, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceSensorNetworkRelationshipsRelationshipsLivestreamRelatedDevices{}
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
				(*parent).RelationshipsLivestreamRelatedDevices = append((*parent).RelationshipsLivestreamRelatedDevices, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
