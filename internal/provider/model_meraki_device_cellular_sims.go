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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceCellularSIMs struct {
	Id                 types.String             `tfsdk:"id"`
	Serial             types.String             `tfsdk:"serial"`
	SimFailoverEnabled types.Bool               `tfsdk:"sim_failover_enabled"`
	SimFailoverTimeout types.Int64              `tfsdk:"sim_failover_timeout"`
	SimOrdering        types.List               `tfsdk:"sim_ordering"`
	Sims               []DeviceCellularSIMsSims `tfsdk:"sims"`
}

type DeviceCellularSIMsSims struct {
	IsPrimary types.Bool                   `tfsdk:"is_primary"`
	SimOrder  types.Int64                  `tfsdk:"sim_order"`
	Slot      types.String                 `tfsdk:"slot"`
	Apns      []DeviceCellularSIMsSimsApns `tfsdk:"apns"`
}

type DeviceCellularSIMsSimsApns struct {
	Name                   types.String `tfsdk:"name"`
	AuthenticationPassword types.String `tfsdk:"authentication_password"`
	AuthenticationType     types.String `tfsdk:"authentication_type"`
	AuthenticationUsername types.String `tfsdk:"authentication_username"`
	AllowedIpTypes         types.Set    `tfsdk:"allowed_ip_types"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceCellularSIMs) getPath() string {
	return fmt.Sprintf("/devices/%v/cellular/sims", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceCellularSIMs) toBody(ctx context.Context, state DeviceCellularSIMs) string {
	body := ""
	if !data.SimFailoverEnabled.IsNull() {
		body, _ = sjson.Set(body, "simFailover.enabled", data.SimFailoverEnabled.ValueBool())
	}
	if !data.SimFailoverTimeout.IsNull() {
		body, _ = sjson.Set(body, "simFailover.timeout", data.SimFailoverTimeout.ValueInt64())
	}
	if !data.SimOrdering.IsNull() {
		var values []string
		data.SimOrdering.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "simOrdering", values)
	}
	if len(data.Sims) > 0 {
		body, _ = sjson.Set(body, "sims", []interface{}{})
		for _, item := range data.Sims {
			itemBody := ""
			if !item.IsPrimary.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isPrimary", item.IsPrimary.ValueBool())
			}
			if !item.SimOrder.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "simOrder", item.SimOrder.ValueInt64())
			}
			if !item.Slot.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "slot", item.Slot.ValueString())
			}
			if len(item.Apns) > 0 {
				itemBody, _ = sjson.Set(itemBody, "apns", []interface{}{})
				for _, childItem := range item.Apns {
					itemChildBody := ""
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.AuthenticationPassword.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "authentication.password", childItem.AuthenticationPassword.ValueString())
					}
					if !childItem.AuthenticationType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "authentication.type", childItem.AuthenticationType.ValueString())
					}
					if !childItem.AuthenticationUsername.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "authentication.username", childItem.AuthenticationUsername.ValueString())
					}
					if !childItem.AllowedIpTypes.IsNull() {
						var values []string
						childItem.AllowedIpTypes.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "allowedIpTypes", values)
					}
					itemBody, _ = sjson.SetRaw(itemBody, "apns.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "sims.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceCellularSIMs) fromBody(ctx context.Context, res gjson.Result) {
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
		data.Sims = make([]DeviceCellularSIMsSims, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceCellularSIMsSims{}
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
				data.Apns = make([]DeviceCellularSIMsSimsApns, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceCellularSIMsSimsApns{}
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceCellularSIMs) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("simFailover.enabled"); value.Exists() && !data.SimFailoverEnabled.IsNull() {
		data.SimFailoverEnabled = types.BoolValue(value.Bool())
	} else {
		data.SimFailoverEnabled = types.BoolNull()
	}
	if value := res.Get("simFailover.timeout"); value.Exists() && !data.SimFailoverTimeout.IsNull() {
		data.SimFailoverTimeout = types.Int64Value(value.Int())
	} else {
		data.SimFailoverTimeout = types.Int64Null()
	}
	if value := res.Get("simOrdering"); value.Exists() && !data.SimOrdering.IsNull() {
		data.SimOrdering = helpers.GetStringList(value.Array())
	} else {
		data.SimOrdering = types.ListNull(types.StringType)
	}
	for i := 0; i < len(data.Sims); i++ {
		keys := [...]string{"slot"}
		keyValues := [...]string{data.Sims[i].Slot.ValueString()}

		parent := &data
		data := (*parent).Sims[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("sims").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Sims[%d] = %+v",
				i,
				(*parent).Sims[i],
			))
			(*parent).Sims = slices.Delete((*parent).Sims, i, i+1)
			i--

			continue
		}
		if value := res.Get("isPrimary"); value.Exists() && !data.IsPrimary.IsNull() {
			data.IsPrimary = types.BoolValue(value.Bool())
		} else {
			data.IsPrimary = types.BoolNull()
		}
		if value := res.Get("simOrder"); value.Exists() && !data.SimOrder.IsNull() {
			data.SimOrder = types.Int64Value(value.Int())
		} else {
			data.SimOrder = types.Int64Null()
		}
		if value := res.Get("slot"); value.Exists() && !data.Slot.IsNull() {
			data.Slot = types.StringValue(value.String())
		} else {
			data.Slot = types.StringNull()
		}
		for i := 0; i < len(data.Apns); i++ {
			keys := [...]string{"name"}
			keyValues := [...]string{data.Apns[i].Name.ValueString()}

			parent := &data
			data := (*parent).Apns[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("apns").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Apns[%d] = %+v",
					i,
					(*parent).Apns[i],
				))
				(*parent).Apns = slices.Delete((*parent).Apns, i, i+1)
				i--

				continue
			}
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("authentication.password"); value.Exists() && !data.AuthenticationPassword.IsNull() {
				data.AuthenticationPassword = types.StringValue(value.String())
			} else {
				data.AuthenticationPassword = types.StringNull()
			}
			if value := res.Get("authentication.type"); value.Exists() && !data.AuthenticationType.IsNull() {
				data.AuthenticationType = types.StringValue(value.String())
			} else {
				data.AuthenticationType = types.StringNull()
			}
			if value := res.Get("authentication.username"); value.Exists() && !data.AuthenticationUsername.IsNull() {
				data.AuthenticationUsername = types.StringValue(value.String())
			} else {
				data.AuthenticationUsername = types.StringNull()
			}
			if value := res.Get("allowedIpTypes"); value.Exists() && !data.AllowedIpTypes.IsNull() {
				data.AllowedIpTypes = helpers.GetStringSet(value.Array())
			} else {
				data.AllowedIpTypes = types.SetNull(types.StringType)
			}
			(*parent).Apns[i] = data
		}
		(*parent).Sims[i] = data
	}
}

// End of section. //template:end fromBodyPartial
