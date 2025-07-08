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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceSensorMQTTBrokers struct {
	Id             types.String                     `tfsdk:"id"`
	OrganizationId types.String                     `tfsdk:"organization_id"`
	NetworkId      types.String                     `tfsdk:"network_id"`
	Items          []ResourceSensorMQTTBrokersItems `tfsdk:"items"`
}

type ResourceSensorMQTTBrokersItems struct {
	MqttBrokerId types.String `tfsdk:"mqtt_broker_id"`
	Enabled      types.Bool   `tfsdk:"enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceSensorMQTTBrokers) getPath() string {
	return fmt.Sprintf("/networks/%v/sensor/mqttBrokers", url.QueryEscape(data.NetworkId.ValueString()))
}

func (data ResourceSensorMQTTBrokers) getItemPath(id string) string {
	return fmt.Sprintf("/networks/%v/sensor/mqttBrokers/%v", url.QueryEscape(data.NetworkId.ValueString()), id)
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceSensorMQTTBrokersItems) toBody(ctx context.Context, state ResourceSensorMQTTBrokersItems) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceSensorMQTTBrokers) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceSensorMQTTBrokersItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceSensorMQTTBrokersItems{}
		if value := res.Get("mqttBrokerId"); value.Exists() && value.Value() != nil {
			data.MqttBrokerId = types.StringValue(value.String())
		} else {
			data.MqttBrokerId = types.StringNull()
		}
		if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	data.Id = data.OrganizationId
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceSensorMQTTBrokers) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("mqttBrokerId").String() == (*parent).Items[i].MqttBrokerId.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
		if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceSensorMQTTBrokers) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceSensorMQTTBrokers) toDestroyBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "enabled", false)
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceSensorMQTTBrokers) hasChanges(ctx context.Context, state *ResourceSensorMQTTBrokers, id string) bool {
	hasChanges := false

	item := ResourceSensorMQTTBrokersItems{}
	for i := range data.Items {
		if data.Items[i].MqttBrokerId.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceSensorMQTTBrokersItems{}
	for i := range state.Items {
		if state.Items[i].MqttBrokerId.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.MqttBrokerId.Equal(stateItem.MqttBrokerId) {
		hasChanges = true
	}
	if !item.Enabled.Equal(stateItem.Enabled) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
