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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchRoutingOSPF struct {
	Id                             types.String               `tfsdk:"id"`
	NetworkId                      types.String               `tfsdk:"network_id"`
	DeadTimerInSeconds             types.Int64                `tfsdk:"dead_timer_in_seconds"`
	Enabled                        types.Bool                 `tfsdk:"enabled"`
	HelloTimerInSeconds            types.Int64                `tfsdk:"hello_timer_in_seconds"`
	Md5AuthenticationEnabled       types.Bool                 `tfsdk:"md5_authentication_enabled"`
	Md5AuthenticationKeyId         types.Int64                `tfsdk:"md5_authentication_key_id"`
	Md5AuthenticationKeyPassphrase types.String               `tfsdk:"md5_authentication_key_passphrase"`
	V3DeadTimerInSeconds           types.Int64                `tfsdk:"v3_dead_timer_in_seconds"`
	V3Enabled                      types.Bool                 `tfsdk:"v3_enabled"`
	V3HelloTimerInSeconds          types.Int64                `tfsdk:"v3_hello_timer_in_seconds"`
	V3Areas                        []SwitchRoutingOSPFV3Areas `tfsdk:"v3_areas"`
	Areas                          []SwitchRoutingOSPFAreas   `tfsdk:"areas"`
}

type SwitchRoutingOSPFV3Areas struct {
	AreaId   types.String `tfsdk:"area_id"`
	AreaName types.String `tfsdk:"area_name"`
	AreaType types.String `tfsdk:"area_type"`
}

type SwitchRoutingOSPFAreas struct {
	AreaId   types.String `tfsdk:"area_id"`
	AreaName types.String `tfsdk:"area_name"`
	AreaType types.String `tfsdk:"area_type"`
}

type SwitchRoutingOSPFIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchRoutingOSPF) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/routing/ospf", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchRoutingOSPF) toBody(ctx context.Context, state SwitchRoutingOSPF) string {
	body := ""
	if !data.DeadTimerInSeconds.IsNull() {
		body, _ = sjson.Set(body, "deadTimerInSeconds", data.DeadTimerInSeconds.ValueInt64())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.HelloTimerInSeconds.IsNull() {
		body, _ = sjson.Set(body, "helloTimerInSeconds", data.HelloTimerInSeconds.ValueInt64())
	}
	if !data.Md5AuthenticationEnabled.IsNull() {
		body, _ = sjson.Set(body, "md5AuthenticationEnabled", data.Md5AuthenticationEnabled.ValueBool())
	}
	if !data.Md5AuthenticationKeyId.IsNull() {
		body, _ = sjson.Set(body, "md5AuthenticationKey.id", data.Md5AuthenticationKeyId.ValueInt64())
	}
	if !data.Md5AuthenticationKeyPassphrase.IsNull() {
		body, _ = sjson.Set(body, "md5AuthenticationKey.passphrase", data.Md5AuthenticationKeyPassphrase.ValueString())
	}
	if !data.V3DeadTimerInSeconds.IsNull() {
		body, _ = sjson.Set(body, "v3.deadTimerInSeconds", data.V3DeadTimerInSeconds.ValueInt64())
	}
	if !data.V3Enabled.IsNull() {
		body, _ = sjson.Set(body, "v3.enabled", data.V3Enabled.ValueBool())
	}
	if !data.V3HelloTimerInSeconds.IsNull() {
		body, _ = sjson.Set(body, "v3.helloTimerInSeconds", data.V3HelloTimerInSeconds.ValueInt64())
	}
	if data.V3Areas != nil {
		body, _ = sjson.Set(body, "v3.areas", []interface{}{})
		for _, item := range data.V3Areas {
			itemBody := ""
			if !item.AreaId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaId", item.AreaId.ValueString())
			}
			if !item.AreaName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaName", item.AreaName.ValueString())
			}
			if !item.AreaType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaType", item.AreaType.ValueString())
			}
			body, _ = sjson.SetRaw(body, "v3.areas.-1", itemBody)
		}
	}
	if data.Areas != nil {
		body, _ = sjson.Set(body, "areas", []interface{}{})
		for _, item := range data.Areas {
			itemBody := ""
			if !item.AreaId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaId", item.AreaId.ValueString())
			}
			if !item.AreaName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaName", item.AreaName.ValueString())
			}
			if !item.AreaType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaType", item.AreaType.ValueString())
			}
			body, _ = sjson.SetRaw(body, "areas.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchRoutingOSPF) fromBody(ctx context.Context, res meraki.Res) {
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
		data.V3Areas = make([]SwitchRoutingOSPFV3Areas, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchRoutingOSPFV3Areas{}
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
		data.Areas = make([]SwitchRoutingOSPFAreas, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchRoutingOSPFAreas{}
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchRoutingOSPF) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("deadTimerInSeconds"); value.Exists() && !data.DeadTimerInSeconds.IsNull() {
		data.DeadTimerInSeconds = types.Int64Value(value.Int())
	} else {
		data.DeadTimerInSeconds = types.Int64Null()
	}
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("helloTimerInSeconds"); value.Exists() && !data.HelloTimerInSeconds.IsNull() {
		data.HelloTimerInSeconds = types.Int64Value(value.Int())
	} else {
		data.HelloTimerInSeconds = types.Int64Null()
	}
	if value := res.Get("md5AuthenticationEnabled"); value.Exists() && !data.Md5AuthenticationEnabled.IsNull() {
		data.Md5AuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.Md5AuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("md5AuthenticationKey.id"); value.Exists() && !data.Md5AuthenticationKeyId.IsNull() {
		data.Md5AuthenticationKeyId = types.Int64Value(value.Int())
	} else {
		data.Md5AuthenticationKeyId = types.Int64Null()
	}
	if value := res.Get("md5AuthenticationKey.passphrase"); value.Exists() && !data.Md5AuthenticationKeyPassphrase.IsNull() {
		data.Md5AuthenticationKeyPassphrase = types.StringValue(value.String())
	} else {
		data.Md5AuthenticationKeyPassphrase = types.StringNull()
	}
	if value := res.Get("v3.deadTimerInSeconds"); value.Exists() && !data.V3DeadTimerInSeconds.IsNull() {
		data.V3DeadTimerInSeconds = types.Int64Value(value.Int())
	} else {
		data.V3DeadTimerInSeconds = types.Int64Null()
	}
	if value := res.Get("v3.enabled"); value.Exists() && !data.V3Enabled.IsNull() {
		data.V3Enabled = types.BoolValue(value.Bool())
	} else {
		data.V3Enabled = types.BoolNull()
	}
	if value := res.Get("v3.helloTimerInSeconds"); value.Exists() && !data.V3HelloTimerInSeconds.IsNull() {
		data.V3HelloTimerInSeconds = types.Int64Value(value.Int())
	} else {
		data.V3HelloTimerInSeconds = types.Int64Null()
	}
	for i := 0; i < len(data.V3Areas); i++ {
		keys := [...]string{"areaId"}
		keyValues := [...]string{data.V3Areas[i].AreaId.ValueString()}

		parent := &data
		data := (*parent).V3Areas[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("v3.areas").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing V3Areas[%d] = %+v",
				i,
				(*parent).V3Areas[i],
			))
			(*parent).V3Areas = slices.Delete((*parent).V3Areas, i, i+1)
			i--

			continue
		}
		if value := res.Get("areaId"); value.Exists() && !data.AreaId.IsNull() {
			data.AreaId = types.StringValue(value.String())
		} else {
			data.AreaId = types.StringNull()
		}
		if value := res.Get("areaName"); value.Exists() && !data.AreaName.IsNull() {
			data.AreaName = types.StringValue(value.String())
		} else {
			data.AreaName = types.StringNull()
		}
		if value := res.Get("areaType"); value.Exists() && !data.AreaType.IsNull() {
			data.AreaType = types.StringValue(value.String())
		} else {
			data.AreaType = types.StringNull()
		}
		(*parent).V3Areas[i] = data
	}
	for i := 0; i < len(data.Areas); i++ {
		keys := [...]string{"areaId"}
		keyValues := [...]string{data.Areas[i].AreaId.ValueString()}

		parent := &data
		data := (*parent).Areas[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("areas").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Areas[%d] = %+v",
				i,
				(*parent).Areas[i],
			))
			(*parent).Areas = slices.Delete((*parent).Areas, i, i+1)
			i--

			continue
		}
		if value := res.Get("areaId"); value.Exists() && !data.AreaId.IsNull() {
			data.AreaId = types.StringValue(value.String())
		} else {
			data.AreaId = types.StringNull()
		}
		if value := res.Get("areaName"); value.Exists() && !data.AreaName.IsNull() {
			data.AreaName = types.StringValue(value.String())
		} else {
			data.AreaName = types.StringNull()
		}
		if value := res.Get("areaType"); value.Exists() && !data.AreaType.IsNull() {
			data.AreaType = types.StringValue(value.String())
		} else {
			data.AreaType = types.StringNull()
		}
		(*parent).Areas[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SwitchRoutingOSPF) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *SwitchRoutingOSPFIdentity) toIdentity(ctx context.Context, plan *SwitchRoutingOSPF) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *SwitchRoutingOSPF) fromIdentity(ctx context.Context, identity *SwitchRoutingOSPFIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data SwitchRoutingOSPF) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
