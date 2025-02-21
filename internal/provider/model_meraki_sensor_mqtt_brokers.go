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

type SensorMQTTBrokers struct {
	NetworkId types.String             `tfsdk:"network_id"`
	Items     []SensorMQTTBrokersItems `tfsdk:"items"`
}

type SensorMQTTBrokersItems struct {
	Id           types.String `tfsdk:"id"`
	Enabled      types.Bool   `tfsdk:"enabled"`
	MqttBrokerId types.String `tfsdk:"mqtt_broker_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SensorMQTTBrokers) getPath() string {
	return fmt.Sprintf("/networks/%v/sensor/mqttBrokers", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SensorMQTTBrokers) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]SensorMQTTBrokersItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := SensorMQTTBrokersItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
			data.Enabled = types.BoolValue(value.Bool())
		} else {
			data.Enabled = types.BoolNull()
		}
		if value := res.Get("mqttBrokerId"); value.Exists() && value.Value() != nil {
			data.MqttBrokerId = types.StringValue(value.String())
		} else {
			data.MqttBrokerId = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
