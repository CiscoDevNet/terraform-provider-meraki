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

type WirelessSSIDDeviceTypeGroupPolicies struct {
	Id                 types.String                                            `tfsdk:"id"`
	NetworkId          types.String                                            `tfsdk:"network_id"`
	Number             types.String                                            `tfsdk:"number"`
	Enabled            types.Bool                                              `tfsdk:"enabled"`
	DeviceTypePolicies []WirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies `tfsdk:"device_type_policies"`
}

type WirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies struct {
	DevicePolicy  types.String `tfsdk:"device_policy"`
	DeviceType    types.String `tfsdk:"device_type"`
	GroupPolicyId types.Int64  `tfsdk:"group_policy_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessSSIDDeviceTypeGroupPolicies) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/deviceTypeGroupPolicies", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSSIDDeviceTypeGroupPolicies) toBody(ctx context.Context, state WirelessSSIDDeviceTypeGroupPolicies) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if len(data.DeviceTypePolicies) > 0 {
		body, _ = sjson.Set(body, "deviceTypePolicies", []interface{}{})
		for _, item := range data.DeviceTypePolicies {
			itemBody := ""
			if !item.DevicePolicy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "devicePolicy", item.DevicePolicy.ValueString())
			}
			if !item.DeviceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceType", item.DeviceType.ValueString())
			}
			if !item.GroupPolicyId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "groupPolicyId", item.GroupPolicyId.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "deviceTypePolicies.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSSIDDeviceTypeGroupPolicies) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("deviceTypePolicies"); value.Exists() && value.Value() != nil {
		data.DeviceTypePolicies = make([]WirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := WirelessSSIDDeviceTypeGroupPoliciesDeviceTypePolicies{}
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessSSIDDeviceTypeGroupPolicies) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	for i := 0; i < len(data.DeviceTypePolicies); i++ {
		keys := [...]string{"devicePolicy", "deviceType"}
		keyValues := [...]string{data.DeviceTypePolicies[i].DevicePolicy.ValueString(), data.DeviceTypePolicies[i].DeviceType.ValueString()}

		parent := &data
		data := (*parent).DeviceTypePolicies[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("deviceTypePolicies").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DeviceTypePolicies[%d] = %+v",
				i,
				(*parent).DeviceTypePolicies[i],
			))
			(*parent).DeviceTypePolicies = slices.Delete((*parent).DeviceTypePolicies, i, i+1)
			i--

			continue
		}
		if value := res.Get("devicePolicy"); value.Exists() && !data.DevicePolicy.IsNull() {
			data.DevicePolicy = types.StringValue(value.String())
		} else {
			data.DevicePolicy = types.StringNull()
		}
		if value := res.Get("deviceType"); value.Exists() && !data.DeviceType.IsNull() {
			data.DeviceType = types.StringValue(value.String())
		} else {
			data.DeviceType = types.StringNull()
		}
		if value := res.Get("groupPolicyId"); value.Exists() && !data.GroupPolicyId.IsNull() {
			data.GroupPolicyId = types.Int64Value(value.Int())
		} else {
			data.GroupPolicyId = types.Int64Null()
		}
		(*parent).DeviceTypePolicies[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessSSIDDeviceTypeGroupPolicies) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data WirelessSSIDDeviceTypeGroupPolicies) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
