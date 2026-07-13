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
// It emits a separate model - DataSourceWirelessSSIDDeviceTypeGroupPolicies - used only by the data source, always including every
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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessSSIDDeviceTypeGroupPolicies struct {
	Id                 types.String                                                      `tfsdk:"id"`
	NetworkId          types.String                                                      `tfsdk:"network_id"`
	Number             types.String                                                      `tfsdk:"number"`
	Enabled            types.Bool                                                        `tfsdk:"enabled"`
	DeviceTypePolicies []DataSourceWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies `tfsdk:"device_type_policies"`
}

type DataSourceWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	DeviceType    types.String `tfsdk:"device_type"`
	GroupPolicyId types.Int64  `tfsdk:"group_policy_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessSSIDDeviceTypeGroupPolicies) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/deviceTypeGroupPolicies", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessSSIDDeviceTypeGroupPolicies) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("deviceTypePolicies"); value.Exists() && value.Value() != nil {
		data.DeviceTypePolicies = make([]DataSourceWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceWirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies{}
			if value := res.Get("devicePolicy"); value.Exists() && value.Value() != nil {
				data.DevicePolicy = types.StringValue(value.String())
			} else {
				data.DevicePolicy = types.StringNull()
			}
			if value := res.Get("deviceType"); value.Exists() && value.Value() != nil {
				data.DeviceType = types.StringValue(value.String())
			} else {
				data.DeviceType = types.StringNull()
			}
			if value := res.Get("groupPolicyId"); value.Exists() && value.Value() != nil {
				data.GroupPolicyId = types.Int64Value(value.Int())
			} else {
				data.GroupPolicyId = types.Int64Null()
			}
			(*parent).DeviceTypePolicies = append((*parent).DeviceTypePolicies, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
