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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type WirelessZigbee struct {
	Id                         types.String `tfsdk:"id"`
	NetworkId                  types.String `tfsdk:"network_id"`
	Enabled                    types.Bool   `tfsdk:"enabled"`
	DefaultsChannel            types.String `tfsdk:"defaults_channel"`
	DefaultsTransmitPowerLevel types.Int64  `tfsdk:"defaults_transmit_power_level"`
	IotControllerSerial        types.String `tfsdk:"iot_controller_serial"`
	LockManagementAddress      types.String `tfsdk:"lock_management_address"`
	LockManagementPassword     types.String `tfsdk:"lock_management_password"`
	LockManagementUsername     types.String `tfsdk:"lock_management_username"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessZigbee) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/zigbee", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessZigbee) toBody(ctx context.Context, state WirelessZigbee) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.DefaultsChannel.IsNull() {
		body, _ = sjson.Set(body, "defaults.channel", data.DefaultsChannel.ValueString())
	}
	if !data.DefaultsTransmitPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "defaults.transmitPowerLevel", data.DefaultsTransmitPowerLevel.ValueInt64())
	}
	if !data.IotControllerSerial.IsNull() {
		body, _ = sjson.Set(body, "iotController.serial", data.IotControllerSerial.ValueString())
	}
	if !data.LockManagementAddress.IsNull() {
		body, _ = sjson.Set(body, "lockManagement.address", data.LockManagementAddress.ValueString())
	}
	if !data.LockManagementPassword.IsNull() {
		body, _ = sjson.Set(body, "lockManagement.password", data.LockManagementPassword.ValueString())
	}
	if !data.LockManagementUsername.IsNull() {
		body, _ = sjson.Set(body, "lockManagement.username", data.LockManagementUsername.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessZigbee) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("defaults.channel"); value.Exists() && value.Value() != nil {
		data.DefaultsChannel = types.StringValue(value.String())
	} else {
		data.DefaultsChannel = types.StringNull()
	}
	if value := res.Get("defaults.transmitPowerLevel"); value.Exists() && value.Value() != nil {
		data.DefaultsTransmitPowerLevel = types.Int64Value(value.Int())
	} else {
		data.DefaultsTransmitPowerLevel = types.Int64Null()
	}
	if value := res.Get("iotController.serial"); value.Exists() && value.Value() != nil {
		data.IotControllerSerial = types.StringValue(value.String())
	} else {
		data.IotControllerSerial = types.StringNull()
	}
	if value := res.Get("lockManagement.address"); value.Exists() && value.Value() != nil {
		data.LockManagementAddress = types.StringValue(value.String())
	} else {
		data.LockManagementAddress = types.StringNull()
	}
	if value := res.Get("lockManagement.password"); value.Exists() && value.Value() != nil {
		data.LockManagementPassword = types.StringValue(value.String())
	} else {
		data.LockManagementPassword = types.StringNull()
	}
	if value := res.Get("lockManagement.username"); value.Exists() && value.Value() != nil {
		data.LockManagementUsername = types.StringValue(value.String())
	} else {
		data.LockManagementUsername = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessZigbee) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("defaults.channel"); value.Exists() && !data.DefaultsChannel.IsNull() {
		data.DefaultsChannel = types.StringValue(value.String())
	} else {
		data.DefaultsChannel = types.StringNull()
	}
	if value := res.Get("defaults.transmitPowerLevel"); value.Exists() && !data.DefaultsTransmitPowerLevel.IsNull() {
		data.DefaultsTransmitPowerLevel = types.Int64Value(value.Int())
	} else {
		data.DefaultsTransmitPowerLevel = types.Int64Null()
	}
	if value := res.Get("iotController.serial"); value.Exists() && !data.IotControllerSerial.IsNull() {
		data.IotControllerSerial = types.StringValue(value.String())
	} else {
		data.IotControllerSerial = types.StringNull()
	}
	if value := res.Get("lockManagement.address"); value.Exists() && !data.LockManagementAddress.IsNull() {
		data.LockManagementAddress = types.StringValue(value.String())
	} else {
		data.LockManagementAddress = types.StringNull()
	}
	if value := res.Get("lockManagement.password"); value.Exists() && !data.LockManagementPassword.IsNull() {
		data.LockManagementPassword = types.StringValue(value.String())
	} else {
		data.LockManagementPassword = types.StringNull()
	}
	if value := res.Get("lockManagement.username"); value.Exists() && !data.LockManagementUsername.IsNull() {
		data.LockManagementUsername = types.StringValue(value.String())
	} else {
		data.LockManagementUsername = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessZigbee) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data WirelessZigbee) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
