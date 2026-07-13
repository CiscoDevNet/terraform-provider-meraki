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
// It emits a separate model - DataSourceDeviceCellularSIMs - used only by the data source, always including every
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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceDeviceCellularSIMs struct {
	Id                 types.String                       `tfsdk:"id"`
	Serial             types.String                       `tfsdk:"serial"`
	SimFailoverEnabled types.Bool                         `tfsdk:"sim_failover_enabled"`
	SimFailoverTimeout types.Int64                        `tfsdk:"sim_failover_timeout"`
	SimOrdering        types.List                         `tfsdk:"sim_ordering"`
	Sims               []DataSourceDeviceCellularSIMsSims `tfsdk:"sims"`
}

type DataSourceDeviceCellularSIMsSims struct {
	IsPrimary types.Bool                             `tfsdk:"is_primary"`
	SimOrder  types.Int64                            `tfsdk:"sim_order"`
	Slot      types.String                           `tfsdk:"slot"`
	Apns      []DataSourceDeviceCellularSIMsSimsApns `tfsdk:"apns"`
}

type DataSourceDeviceCellularSIMsSimsApns struct {
	Name                            types.String `tfsdk:"name"`
	AuthenticationPassword          types.String `tfsdk:"authentication_password"`
	AuthenticationPasswordWo        types.String `tfsdk:"authentication_password_wo"`
	AuthenticationPasswordWoVersion types.Int64  `tfsdk:"authentication_password_wo_version"`
	AuthenticationType              types.String `tfsdk:"authentication_type"`
	AuthenticationUsername          types.String `tfsdk:"authentication_username"`
	AllowedIpTypes                  types.Set    `tfsdk:"allowed_ip_types"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceDeviceCellularSIMs) getPath() string {
	return fmt.Sprintf("/devices/%v/cellular/sims", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceDeviceCellularSIMs) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("simFailover.enabled"); value.Exists() && value.Value() != nil {
		data.SimFailoverEnabled = types.BoolValue(value.Bool())
	} else {
		data.SimFailoverEnabled = types.BoolNull()
	}
	if value := res.Get("simFailover.timeout"); value.Exists() && value.Value() != nil {
		data.SimFailoverTimeout = types.Int64Value(value.Int())
	} else {
		data.SimFailoverTimeout = types.Int64Null()
	}
	if value := res.Get("simOrdering"); value.Exists() && value.Value() != nil {
		data.SimOrdering = helpers.GetStringList(value.Array())
	} else {
		data.SimOrdering = types.ListNull(types.StringType)
	}
	if value := res.Get("sims"); value.Exists() && value.Value() != nil {
		data.Sims = make([]DataSourceDeviceCellularSIMsSims, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceDeviceCellularSIMsSims{}
			if value := res.Get("isPrimary"); value.Exists() && value.Value() != nil {
				data.IsPrimary = types.BoolValue(value.Bool())
			} else {
				data.IsPrimary = types.BoolNull()
			}
			if value := res.Get("simOrder"); value.Exists() && value.Value() != nil {
				data.SimOrder = types.Int64Value(value.Int())
			} else {
				data.SimOrder = types.Int64Null()
			}
			if value := res.Get("slot"); value.Exists() && value.Value() != nil {
				data.Slot = types.StringValue(value.String())
			} else {
				data.Slot = types.StringNull()
			}
			if value := res.Get("apns"); value.Exists() && value.Value() != nil {
				data.Apns = make([]DataSourceDeviceCellularSIMsSimsApns, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DataSourceDeviceCellularSIMsSimsApns{}
					if value := res.Get("name"); value.Exists() && value.Value() != nil {
						data.Name = types.StringValue(value.String())
					} else {
						data.Name = types.StringNull()
					}
					if value := res.Get("authentication.password"); value.Exists() && value.Value() != nil {
						data.AuthenticationPassword = types.StringValue(value.String())
					} else {
						data.AuthenticationPassword = types.StringNull()
					}
					if value := res.Get("authentication.type"); value.Exists() && value.Value() != nil {
						data.AuthenticationType = types.StringValue(value.String())
					} else {
						data.AuthenticationType = types.StringNull()
					}
					if value := res.Get("authentication.username"); value.Exists() && value.Value() != nil {
						data.AuthenticationUsername = types.StringValue(value.String())
					} else {
						data.AuthenticationUsername = types.StringNull()
					}
					if value := res.Get("allowedIpTypes"); value.Exists() && value.Value() != nil {
						data.AllowedIpTypes = helpers.GetStringSet(value.Array())
					} else {
						data.AllowedIpTypes = types.SetNull(types.StringType)
					}
					(*parent).Apns = append((*parent).Apns, data)
					return true
				})
			}
			(*parent).Sims = append((*parent).Sims, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
