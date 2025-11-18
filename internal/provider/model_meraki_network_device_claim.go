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
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkDeviceClaim struct {
	Id              types.String                        `tfsdk:"id"`
	NetworkId       types.String                        `tfsdk:"network_id"`
	DetailsByDevice []NetworkDeviceClaimDetailsByDevice `tfsdk:"details_by_device"`
	Serials         types.Set                           `tfsdk:"serials"`
}

type NetworkDeviceClaimDetailsByDevice struct {
	Serial  types.String                               `tfsdk:"serial"`
	Details []NetworkDeviceClaimDetailsByDeviceDetails `tfsdk:"details"`
}

type NetworkDeviceClaimDetailsByDeviceDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkDeviceClaim) getPath() string {
	return fmt.Sprintf("/networks/%v/devices/claim", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

func (data NetworkDeviceClaim) getRemovePath() string {
	return fmt.Sprintf("/networks/%v/devices/remove", url.QueryEscape(data.NetworkId.ValueString()))
}

func (data NetworkDeviceClaim) getDevicesPath() string {
	return fmt.Sprintf("/networks/%v/devices", url.QueryEscape(data.NetworkId.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkDeviceClaim) toBody(ctx context.Context, state NetworkDeviceClaim) string {
	body := ""
	if data.DetailsByDevice != nil {
		body, _ = sjson.Set(body, "detailsByDevice", []interface{}{})
		for _, item := range data.DetailsByDevice {
			itemBody := ""
			if !item.Serial.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serial", item.Serial.ValueString())
			}
			{
				itemBody, _ = sjson.Set(itemBody, "details", []interface{}{})
				for _, childItem := range item.Details {
					itemChildBody := ""
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "details.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "detailsByDevice.-1", itemBody)
		}
	}
	if !data.Serials.IsNull() {
		var values []string
		data.Serials.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "serials", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkDeviceClaim) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("detailsByDevice"); value.Exists() && value.Value() != nil {
		data.DetailsByDevice = make([]NetworkDeviceClaimDetailsByDevice, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkDeviceClaimDetailsByDevice{}
			if value := res.Get("serial"); value.Exists() && value.Value() != nil {
				data.Serial = types.StringValue(value.String())
			} else {
				data.Serial = types.StringNull()
			}
			if value := res.Get("details"); value.Exists() && value.Value() != nil {
				data.Details = make([]NetworkDeviceClaimDetailsByDeviceDetails, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := NetworkDeviceClaimDetailsByDeviceDetails{}
					if value := res.Get("name"); value.Exists() && value.Value() != nil {
						data.Name = types.StringValue(value.String())
					} else {
						data.Name = types.StringNull()
					}
					if value := res.Get("value"); value.Exists() && value.Value() != nil {
						data.Value = types.StringValue(value.String())
					} else {
						data.Value = types.StringNull()
					}
					(*parent).Details = append((*parent).Details, data)
					return true
				})
			}
			(*parent).DetailsByDevice = append((*parent).DetailsByDevice, data)
			return true
		})
	}
	if value := res.Get("serials"); value.Exists() && value.Value() != nil {
		data.Serials = helpers.GetStringSet(value.Array())
	} else {
		data.Serials = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkDeviceClaim) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.DetailsByDevice); i++ {
		keys := [...]string{"serial"}
		keyValues := [...]string{data.DetailsByDevice[i].Serial.ValueString()}

		parent := &data
		data := (*parent).DetailsByDevice[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("detailsByDevice").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DetailsByDevice[%d] = %+v",
				i,
				(*parent).DetailsByDevice[i],
			))
			(*parent).DetailsByDevice = slices.Delete((*parent).DetailsByDevice, i, i+1)
			i--

			continue
		}
		if value := res.Get("serial"); value.Exists() && !data.Serial.IsNull() {
			data.Serial = types.StringValue(value.String())
		} else {
			data.Serial = types.StringNull()
		}
		for i := 0; i < len(data.Details); i++ {
			keys := [...]string{"name", "value"}
			keyValues := [...]string{data.Details[i].Name.ValueString(), data.Details[i].Value.ValueString()}

			parent := &data
			data := (*parent).Details[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("details").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Details[%d] = %+v",
					i,
					(*parent).Details[i],
				))
				(*parent).Details = slices.Delete((*parent).Details, i, i+1)
				i--

				continue
			}
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).Details[i] = data
		}
		(*parent).DetailsByDevice[i] = data
	}
	if value := res.Get("serials"); value.Exists() && !data.Serials.IsNull() {
		data.Serials = helpers.GetStringSet(value.Array())
	} else {
		data.Serials = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkDeviceClaim) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkDeviceClaim) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
