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
// It emits a separate model - DataSourceWirelessMQTTSettings - used only by the data source, always including every
// attribute (unlike model_resource.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model_resource.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model_resource.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessMQTTSettings struct {
	Id                      types.String `tfsdk:"id"`
	OrganizationId          types.String `tfsdk:"organization_id"`
	BleEnabled              types.Bool   `tfsdk:"ble_enabled"`
	BleType                 types.String `tfsdk:"ble_type"`
	BleAllowListsMacs       types.List   `tfsdk:"ble_allow_lists_macs"`
	BleAllowListsUuids      types.List   `tfsdk:"ble_allow_lists_uuids"`
	BleFlushFrequency       types.Int64  `tfsdk:"ble_flush_frequency"`
	BleHysteresisEnabled    types.Bool   `tfsdk:"ble_hysteresis_enabled"`
	BleHysteresisThreshold  types.Int64  `tfsdk:"ble_hysteresis_threshold"`
	MqttEnabled             types.Bool   `tfsdk:"mqtt_enabled"`
	MqttTopic               types.String `tfsdk:"mqtt_topic"`
	MqttBrokerName          types.String `tfsdk:"mqtt_broker_name"`
	MqttPublishingFrequency types.Int64  `tfsdk:"mqtt_publishing_frequency"`
	MqttPublishingQos       types.Int64  `tfsdk:"mqtt_publishing_qos"`
	MqttMessageFields       types.List   `tfsdk:"mqtt_message_fields"`
	NetworkId               types.String `tfsdk:"network_id"`
	WifiEnabled             types.Bool   `tfsdk:"wifi_enabled"`
	WifiType                types.String `tfsdk:"wifi_type"`
	WifiAllowListsMacs      types.List   `tfsdk:"wifi_allow_lists_macs"`
	WifiFlushFrequency      types.Int64  `tfsdk:"wifi_flush_frequency"`
	WifiHysteresisEnabled   types.Bool   `tfsdk:"wifi_hysteresis_enabled"`
	WifiHysteresisThreshold types.Int64  `tfsdk:"wifi_hysteresis_threshold"`
}

// End of section. //template:end types

// getPath is de-generated (kept in sync manually with the resource-side model_resource.go's custom getPath) because
// `network_id` is a second `reference: true` attribute used only for post-fetch filtering against the org-wide
// GET response, not part of the URL - the naive template would wrongly pass it as an extra fmt.Sprintf arg.
func (data DataSourceWirelessMQTTSettings) getPath() string {
	return fmt.Sprintf("/organizations/%v/wireless/mqtt/settings", url.QueryEscape(data.OrganizationId.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessMQTTSettings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("ble.enabled"); value.Exists() && value.Value() != nil {
		data.BleEnabled = types.BoolValue(value.Bool())
	} else {
		data.BleEnabled = types.BoolNull()
	}
	if value := res.Get("ble.type"); value.Exists() && value.Value() != nil {
		data.BleType = types.StringValue(value.String())
	} else {
		data.BleType = types.StringNull()
	}
	if value := res.Get("ble.allowLists.macs"); value.Exists() && value.Value() != nil {
		data.BleAllowListsMacs = helpers.GetStringList(value.Array())
	} else {
		data.BleAllowListsMacs = types.ListNull(types.StringType)
	}
	if value := res.Get("ble.allowLists.uuids"); value.Exists() && value.Value() != nil {
		data.BleAllowListsUuids = helpers.GetStringList(value.Array())
	} else {
		data.BleAllowListsUuids = types.ListNull(types.StringType)
	}
	if value := res.Get("ble.flush.frequency"); value.Exists() && value.Value() != nil {
		data.BleFlushFrequency = types.Int64Value(value.Int())
	} else {
		data.BleFlushFrequency = types.Int64Null()
	}
	if value := res.Get("ble.hysteresis.enabled"); value.Exists() && value.Value() != nil {
		data.BleHysteresisEnabled = types.BoolValue(value.Bool())
	} else {
		data.BleHysteresisEnabled = types.BoolNull()
	}
	if value := res.Get("ble.hysteresis.threshold"); value.Exists() && value.Value() != nil {
		data.BleHysteresisThreshold = types.Int64Value(value.Int())
	} else {
		data.BleHysteresisThreshold = types.Int64Null()
	}
	if value := res.Get("mqtt.enabled"); value.Exists() && value.Value() != nil {
		data.MqttEnabled = types.BoolValue(value.Bool())
	} else {
		data.MqttEnabled = types.BoolNull()
	}
	if value := res.Get("mqtt.topic"); value.Exists() && value.Value() != nil {
		data.MqttTopic = types.StringValue(value.String())
	} else {
		data.MqttTopic = types.StringNull()
	}
	if value := res.Get("mqtt.broker.name"); value.Exists() && value.Value() != nil {
		data.MqttBrokerName = types.StringValue(value.String())
	} else {
		data.MqttBrokerName = types.StringNull()
	}
	if value := res.Get("mqtt.publishing.frequency"); value.Exists() && value.Value() != nil {
		data.MqttPublishingFrequency = types.Int64Value(value.Int())
	} else {
		data.MqttPublishingFrequency = types.Int64Null()
	}
	if value := res.Get("mqtt.publishing.qos"); value.Exists() && value.Value() != nil {
		data.MqttPublishingQos = types.Int64Value(value.Int())
	} else {
		data.MqttPublishingQos = types.Int64Null()
	}
	if value := res.Get("mqtt.messageFields"); value.Exists() && value.Value() != nil {
		data.MqttMessageFields = helpers.GetStringList(value.Array())
	} else {
		data.MqttMessageFields = types.ListNull(types.StringType)
	}
	if value := res.Get("wifi.enabled"); value.Exists() && value.Value() != nil {
		data.WifiEnabled = types.BoolValue(value.Bool())
	} else {
		data.WifiEnabled = types.BoolNull()
	}
	if value := res.Get("wifi.type"); value.Exists() && value.Value() != nil {
		data.WifiType = types.StringValue(value.String())
	} else {
		data.WifiType = types.StringNull()
	}
	if value := res.Get("wifi.allowLists.macs"); value.Exists() && value.Value() != nil {
		data.WifiAllowListsMacs = helpers.GetStringList(value.Array())
	} else {
		data.WifiAllowListsMacs = types.ListNull(types.StringType)
	}
	if value := res.Get("wifi.flush.frequency"); value.Exists() && value.Value() != nil {
		data.WifiFlushFrequency = types.Int64Value(value.Int())
	} else {
		data.WifiFlushFrequency = types.Int64Null()
	}
	if value := res.Get("wifi.hysteresis.enabled"); value.Exists() && value.Value() != nil {
		data.WifiHysteresisEnabled = types.BoolValue(value.Bool())
	} else {
		data.WifiHysteresisEnabled = types.BoolNull()
	}
	if value := res.Get("wifi.hysteresis.threshold"); value.Exists() && value.Value() != nil {
		data.WifiHysteresisThreshold = types.Int64Value(value.Int())
	} else {
		data.WifiHysteresisThreshold = types.Int64Null()
	}
}

// End of section. //template:end fromBody
