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
// It emits a separate model - DataSourceWirelessZigbee - used only by the data source, always including every
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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessZigbee struct {
	Id                              types.String `tfsdk:"id"`
	NetworkId                       types.String `tfsdk:"network_id"`
	Enabled                         types.Bool   `tfsdk:"enabled"`
	DefaultsChannel                 types.String `tfsdk:"defaults_channel"`
	DefaultsTransmitPowerLevel      types.Int64  `tfsdk:"defaults_transmit_power_level"`
	IotControllerSerial             types.String `tfsdk:"iot_controller_serial"`
	LockManagementAddress           types.String `tfsdk:"lock_management_address"`
	LockManagementPassword          types.String `tfsdk:"lock_management_password"`
	LockManagementPasswordWo        types.String `tfsdk:"lock_management_password_wo"`
	LockManagementPasswordWoVersion types.Int64  `tfsdk:"lock_management_password_wo_version"`
	LockManagementUsername          types.String `tfsdk:"lock_management_username"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessZigbee) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/zigbee", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessZigbee) fromBody(ctx context.Context, res meraki.Res) {
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
