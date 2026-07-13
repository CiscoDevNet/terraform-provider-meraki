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
// It emits a separate model - DataSourceSwitchRoutingOSPF - used only by the data source, always including every
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

type DataSourceSwitchRoutingOSPF struct {
	Id                                      types.String                         `tfsdk:"id"`
	NetworkId                               types.String                         `tfsdk:"network_id"`
	DeadTimerInSeconds                      types.Int64                          `tfsdk:"dead_timer_in_seconds"`
	Enabled                                 types.Bool                           `tfsdk:"enabled"`
	HelloTimerInSeconds                     types.Int64                          `tfsdk:"hello_timer_in_seconds"`
	Md5AuthenticationEnabled                types.Bool                           `tfsdk:"md5_authentication_enabled"`
	Md5AuthenticationKeyId                  types.Int64                          `tfsdk:"md5_authentication_key_id"`
	Md5AuthenticationKeyPassphrase          types.String                         `tfsdk:"md5_authentication_key_passphrase"`
	Md5AuthenticationKeyPassphraseWo        types.String                         `tfsdk:"md5_authentication_key_passphrase_wo"`
	Md5AuthenticationKeyPassphraseWoVersion types.Int64                          `tfsdk:"md5_authentication_key_passphrase_wo_version"`
	V3DeadTimerInSeconds                    types.Int64                          `tfsdk:"v3_dead_timer_in_seconds"`
	V3Enabled                               types.Bool                           `tfsdk:"v3_enabled"`
	V3HelloTimerInSeconds                   types.Int64                          `tfsdk:"v3_hello_timer_in_seconds"`
	V3Areas                                 []DataSourceSwitchRoutingOSPFV3Areas `tfsdk:"v3_areas"`
	Areas                                   []DataSourceSwitchRoutingOSPFAreas   `tfsdk:"areas"`
}

type DataSourceSwitchRoutingOSPFV3Areas struct {
	AreaId   types.String `tfsdk:"area_id"`
	AreaName types.String `tfsdk:"area_name"`
	AreaType types.String `tfsdk:"area_type"`
}

type DataSourceSwitchRoutingOSPFAreas struct {
	AreaId   types.String `tfsdk:"area_id"`
	AreaName types.String `tfsdk:"area_name"`
	AreaType types.String `tfsdk:"area_type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceSwitchRoutingOSPF) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/routing/ospf", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceSwitchRoutingOSPF) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("deadTimerInSeconds"); value.Exists() && value.Value() != nil {
		data.DeadTimerInSeconds = types.Int64Value(value.Int())
	} else {
		data.DeadTimerInSeconds = types.Int64Null()
	}
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("helloTimerInSeconds"); value.Exists() && value.Value() != nil {
		data.HelloTimerInSeconds = types.Int64Value(value.Int())
	} else {
		data.HelloTimerInSeconds = types.Int64Null()
	}
	if value := res.Get("md5AuthenticationEnabled"); value.Exists() && value.Value() != nil {
		data.Md5AuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.Md5AuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("md5AuthenticationKey.id"); value.Exists() && value.Value() != nil {
		data.Md5AuthenticationKeyId = types.Int64Value(value.Int())
	} else {
		data.Md5AuthenticationKeyId = types.Int64Null()
	}
	if value := res.Get("md5AuthenticationKey.passphrase"); value.Exists() && value.Value() != nil {
		data.Md5AuthenticationKeyPassphrase = types.StringValue(value.String())
	} else {
		data.Md5AuthenticationKeyPassphrase = types.StringNull()
	}
	if value := res.Get("v3.deadTimerInSeconds"); value.Exists() && value.Value() != nil {
		data.V3DeadTimerInSeconds = types.Int64Value(value.Int())
	} else {
		data.V3DeadTimerInSeconds = types.Int64Null()
	}
	if value := res.Get("v3.enabled"); value.Exists() && value.Value() != nil {
		data.V3Enabled = types.BoolValue(value.Bool())
	} else {
		data.V3Enabled = types.BoolNull()
	}
	if value := res.Get("v3.helloTimerInSeconds"); value.Exists() && value.Value() != nil {
		data.V3HelloTimerInSeconds = types.Int64Value(value.Int())
	} else {
		data.V3HelloTimerInSeconds = types.Int64Null()
	}
	if value := res.Get("v3.areas"); value.Exists() && value.Value() != nil {
		data.V3Areas = make([]DataSourceSwitchRoutingOSPFV3Areas, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceSwitchRoutingOSPFV3Areas{}
			if value := res.Get("areaId"); value.Exists() && value.Value() != nil {
				data.AreaId = types.StringValue(value.String())
			} else {
				data.AreaId = types.StringNull()
			}
			if value := res.Get("areaName"); value.Exists() && value.Value() != nil {
				data.AreaName = types.StringValue(value.String())
			} else {
				data.AreaName = types.StringNull()
			}
			if value := res.Get("areaType"); value.Exists() && value.Value() != nil {
				data.AreaType = types.StringValue(value.String())
			} else {
				data.AreaType = types.StringNull()
			}
			(*parent).V3Areas = append((*parent).V3Areas, data)
			return true
		})
	}
	if value := res.Get("areas"); value.Exists() && value.Value() != nil {
		data.Areas = make([]DataSourceSwitchRoutingOSPFAreas, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceSwitchRoutingOSPFAreas{}
			if value := res.Get("areaId"); value.Exists() && value.Value() != nil {
				data.AreaId = types.StringValue(value.String())
			} else {
				data.AreaId = types.StringNull()
			}
			if value := res.Get("areaName"); value.Exists() && value.Value() != nil {
				data.AreaName = types.StringValue(value.String())
			} else {
				data.AreaName = types.StringNull()
			}
			if value := res.Get("areaType"); value.Exists() && value.Value() != nil {
				data.AreaType = types.StringValue(value.String())
			} else {
				data.AreaType = types.StringNull()
			}
			(*parent).Areas = append((*parent).Areas, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
