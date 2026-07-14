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
// It emits a separate model - DataSourceWirelessNetworkBluetoothSettings - used only by the data source, always including every
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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessNetworkBluetoothSettings struct {
	Id                       types.String `tfsdk:"id"`
	NetworkId                types.String `tfsdk:"network_id"`
	AdvertisingEnabled       types.Bool   `tfsdk:"advertising_enabled"`
	Major                    types.Int64  `tfsdk:"major"`
	MajorMinorAssignmentMode types.String `tfsdk:"major_minor_assignment_mode"`
	Minor                    types.Int64  `tfsdk:"minor"`
	ScanningEnabled          types.Bool   `tfsdk:"scanning_enabled"`
	Uuid                     types.String `tfsdk:"uuid"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessNetworkBluetoothSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/bluetooth/settings", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessNetworkBluetoothSettings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("advertisingEnabled"); value.Exists() && value.Value() != nil {
		data.AdvertisingEnabled = types.BoolValue(value.Bool())
	} else {
		data.AdvertisingEnabled = types.BoolNull()
	}
	if value := res.Get("major"); value.Exists() && value.Value() != nil {
		data.Major = types.Int64Value(value.Int())
	} else {
		data.Major = types.Int64Null()
	}
	if value := res.Get("majorMinorAssignmentMode"); value.Exists() && value.Value() != nil {
		data.MajorMinorAssignmentMode = types.StringValue(value.String())
	} else {
		data.MajorMinorAssignmentMode = types.StringNull()
	}
	if value := res.Get("minor"); value.Exists() && value.Value() != nil {
		data.Minor = types.Int64Value(value.Int())
	} else {
		data.Minor = types.Int64Null()
	}
	if value := res.Get("scanningEnabled"); value.Exists() && value.Value() != nil {
		data.ScanningEnabled = types.BoolValue(value.Bool())
	} else {
		data.ScanningEnabled = types.BoolNull()
	}
	if value := res.Get("uuid"); value.Exists() && value.Value() != nil {
		data.Uuid = types.StringValue(value.String())
	} else {
		data.Uuid = types.StringNull()
	}
}

// End of section. //template:end fromBody
