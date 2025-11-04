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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessMQTTSettings struct {
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

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessMQTTSettings) getPath() string {
	return fmt.Sprintf("/organizations/%v/wireless/mqtt/settings", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessMQTTSettings) toBody(ctx context.Context, state WirelessMQTTSettings) string {
	body := ""
	if !data.BleEnabled.IsNull() {
		body, _ = sjson.Set(body, "ble.enabled", data.BleEnabled.ValueBool())
	}
	if !data.BleType.IsNull() {
		body, _ = sjson.Set(body, "ble.type", data.BleType.ValueString())
	}
	if !data.BleAllowListsMacs.IsNull() {
		var values []string
		data.BleAllowListsMacs.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ble.allowLists.macs", values)
	}
	if !data.BleAllowListsUuids.IsNull() {
		var values []string
		data.BleAllowListsUuids.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ble.allowLists.uuids", values)
	}
	if !data.BleFlushFrequency.IsNull() {
		body, _ = sjson.Set(body, "ble.flush.frequency", data.BleFlushFrequency.ValueInt64())
	}
	if !data.BleHysteresisEnabled.IsNull() {
		body, _ = sjson.Set(body, "ble.hysteresis.enabled", data.BleHysteresisEnabled.ValueBool())
	}
	if !data.BleHysteresisThreshold.IsNull() {
		body, _ = sjson.Set(body, "ble.hysteresis.threshold", data.BleHysteresisThreshold.ValueInt64())
	}
	if !data.MqttEnabled.IsNull() {
		body, _ = sjson.Set(body, "mqtt.enabled", data.MqttEnabled.ValueBool())
	}
	if !data.MqttTopic.IsNull() {
		body, _ = sjson.Set(body, "mqtt.topic", data.MqttTopic.ValueString())
	}
	if !data.MqttBrokerName.IsNull() {
		body, _ = sjson.Set(body, "mqtt.broker.name", data.MqttBrokerName.ValueString())
	}
	if !data.MqttPublishingFrequency.IsNull() {
		body, _ = sjson.Set(body, "mqtt.publishing.frequency", data.MqttPublishingFrequency.ValueInt64())
	}
	if !data.MqttPublishingQos.IsNull() {
		body, _ = sjson.Set(body, "mqtt.publishing.qos", data.MqttPublishingQos.ValueInt64())
	}
	if !data.MqttMessageFields.IsNull() {
		var values []string
		data.MqttMessageFields.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "mqtt.messageFields", values)
	}
	if !data.NetworkId.IsNull() {
		body, _ = sjson.Set(body, "network.id", data.NetworkId.ValueString())
	}
	if !data.WifiEnabled.IsNull() {
		body, _ = sjson.Set(body, "wifi.enabled", data.WifiEnabled.ValueBool())
	}
	if !data.WifiType.IsNull() {
		body, _ = sjson.Set(body, "wifi.type", data.WifiType.ValueString())
	}
	if !data.WifiAllowListsMacs.IsNull() {
		var values []string
		data.WifiAllowListsMacs.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "wifi.allowLists.macs", values)
	}
	if !data.WifiFlushFrequency.IsNull() {
		body, _ = sjson.Set(body, "wifi.flush.frequency", data.WifiFlushFrequency.ValueInt64())
	}
	if !data.WifiHysteresisEnabled.IsNull() {
		body, _ = sjson.Set(body, "wifi.hysteresis.enabled", data.WifiHysteresisEnabled.ValueBool())
	}
	if !data.WifiHysteresisThreshold.IsNull() {
		body, _ = sjson.Set(body, "wifi.hysteresis.threshold", data.WifiHysteresisThreshold.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessMQTTSettings) fromBody(ctx context.Context, res meraki.Res) {
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
	if value := res.Get("network.id"); value.Exists() && value.Value() != nil {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessMQTTSettings) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("ble.enabled"); value.Exists() && !data.BleEnabled.IsNull() {
		data.BleEnabled = types.BoolValue(value.Bool())
	} else {
		data.BleEnabled = types.BoolNull()
	}
	if value := res.Get("ble.type"); value.Exists() && !data.BleType.IsNull() {
		data.BleType = types.StringValue(value.String())
	} else {
		data.BleType = types.StringNull()
	}
	if value := res.Get("ble.allowLists.macs"); value.Exists() && !data.BleAllowListsMacs.IsNull() {
		data.BleAllowListsMacs = helpers.GetStringList(value.Array())
	} else {
		data.BleAllowListsMacs = types.ListNull(types.StringType)
	}
	if value := res.Get("ble.allowLists.uuids"); value.Exists() && !data.BleAllowListsUuids.IsNull() {
		data.BleAllowListsUuids = helpers.GetStringList(value.Array())
	} else {
		data.BleAllowListsUuids = types.ListNull(types.StringType)
	}
	if value := res.Get("ble.flush.frequency"); value.Exists() && !data.BleFlushFrequency.IsNull() {
		data.BleFlushFrequency = types.Int64Value(value.Int())
	} else {
		data.BleFlushFrequency = types.Int64Null()
	}
	if value := res.Get("ble.hysteresis.enabled"); value.Exists() && !data.BleHysteresisEnabled.IsNull() {
		data.BleHysteresisEnabled = types.BoolValue(value.Bool())
	} else {
		data.BleHysteresisEnabled = types.BoolNull()
	}
	if value := res.Get("ble.hysteresis.threshold"); value.Exists() && !data.BleHysteresisThreshold.IsNull() {
		data.BleHysteresisThreshold = types.Int64Value(value.Int())
	} else {
		data.BleHysteresisThreshold = types.Int64Null()
	}
	if value := res.Get("mqtt.enabled"); value.Exists() && !data.MqttEnabled.IsNull() {
		data.MqttEnabled = types.BoolValue(value.Bool())
	} else {
		data.MqttEnabled = types.BoolNull()
	}
	if value := res.Get("mqtt.topic"); value.Exists() && !data.MqttTopic.IsNull() {
		data.MqttTopic = types.StringValue(value.String())
	} else {
		data.MqttTopic = types.StringNull()
	}
	if value := res.Get("mqtt.broker.name"); value.Exists() && !data.MqttBrokerName.IsNull() {
		data.MqttBrokerName = types.StringValue(value.String())
	} else {
		data.MqttBrokerName = types.StringNull()
	}
	if value := res.Get("mqtt.publishing.frequency"); value.Exists() && !data.MqttPublishingFrequency.IsNull() {
		data.MqttPublishingFrequency = types.Int64Value(value.Int())
	} else {
		data.MqttPublishingFrequency = types.Int64Null()
	}
	if value := res.Get("mqtt.publishing.qos"); value.Exists() && !data.MqttPublishingQos.IsNull() {
		data.MqttPublishingQos = types.Int64Value(value.Int())
	} else {
		data.MqttPublishingQos = types.Int64Null()
	}
	if value := res.Get("mqtt.messageFields"); value.Exists() && !data.MqttMessageFields.IsNull() {
		data.MqttMessageFields = helpers.GetStringList(value.Array())
	} else {
		data.MqttMessageFields = types.ListNull(types.StringType)
	}
	if value := res.Get("network.id"); value.Exists() && !data.NetworkId.IsNull() {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
	}
	if value := res.Get("wifi.enabled"); value.Exists() && !data.WifiEnabled.IsNull() {
		data.WifiEnabled = types.BoolValue(value.Bool())
	} else {
		data.WifiEnabled = types.BoolNull()
	}
	if value := res.Get("wifi.type"); value.Exists() && !data.WifiType.IsNull() {
		data.WifiType = types.StringValue(value.String())
	} else {
		data.WifiType = types.StringNull()
	}
	if value := res.Get("wifi.allowLists.macs"); value.Exists() && !data.WifiAllowListsMacs.IsNull() {
		data.WifiAllowListsMacs = helpers.GetStringList(value.Array())
	} else {
		data.WifiAllowListsMacs = types.ListNull(types.StringType)
	}
	if value := res.Get("wifi.flush.frequency"); value.Exists() && !data.WifiFlushFrequency.IsNull() {
		data.WifiFlushFrequency = types.Int64Value(value.Int())
	} else {
		data.WifiFlushFrequency = types.Int64Null()
	}
	if value := res.Get("wifi.hysteresis.enabled"); value.Exists() && !data.WifiHysteresisEnabled.IsNull() {
		data.WifiHysteresisEnabled = types.BoolValue(value.Bool())
	} else {
		data.WifiHysteresisEnabled = types.BoolNull()
	}
	if value := res.Get("wifi.hysteresis.threshold"); value.Exists() && !data.WifiHysteresisThreshold.IsNull() {
		data.WifiHysteresisThreshold = types.Int64Value(value.Int())
	} else {
		data.WifiHysteresisThreshold = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessMQTTSettings) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data WirelessMQTTSettings) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
