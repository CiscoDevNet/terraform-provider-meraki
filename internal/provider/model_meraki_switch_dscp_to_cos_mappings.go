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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchDSCPToCoSMappings struct {
	Id        types.String                      `tfsdk:"id"`
	NetworkId types.String                      `tfsdk:"network_id"`
	Mappings  []SwitchDSCPToCoSMappingsMappings `tfsdk:"mappings"`
}

type SwitchDSCPToCoSMappingsMappings struct {
	Cos   types.Int64  `tfsdk:"cos"`
	Dscp  types.Int64  `tfsdk:"dscp"`
	Title types.String `tfsdk:"title"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchDSCPToCoSMappings) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/dscpToCosMappings", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchDSCPToCoSMappings) toBody(ctx context.Context, state SwitchDSCPToCoSMappings) string {
	body := ""
	{
		body, _ = sjson.Set(body, "mappings", []interface{}{})
		for _, item := range data.Mappings {
			itemBody := ""
			if !item.Cos.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "cos", item.Cos.ValueInt64())
			}
			if !item.Dscp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dscp", item.Dscp.ValueInt64())
			}
			if !item.Title.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "title", item.Title.ValueString())
			}
			body, _ = sjson.SetRaw(body, "mappings.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchDSCPToCoSMappings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("mappings"); value.Exists() && value.Value() != nil {
		data.Mappings = make([]SwitchDSCPToCoSMappingsMappings, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchDSCPToCoSMappingsMappings{}
			if value := res.Get("cos"); value.Exists() && value.Value() != nil {
				data.Cos = types.Int64Value(value.Int())
			} else {
				data.Cos = types.Int64Null()
			}
			if value := res.Get("dscp"); value.Exists() && value.Value() != nil {
				data.Dscp = types.Int64Value(value.Int())
			} else {
				data.Dscp = types.Int64Null()
			}
			if value := res.Get("title"); value.Exists() && value.Value() != nil {
				data.Title = types.StringValue(value.String())
			} else {
				data.Title = types.StringNull()
			}
			(*parent).Mappings = append((*parent).Mappings, data)
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
func (data *SwitchDSCPToCoSMappings) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.Mappings); i++ {
		keys := [...]string{"cos", "dscp", "title"}
		keyValues := [...]string{strconv.FormatInt(data.Mappings[i].Cos.ValueInt64(), 10), strconv.FormatInt(data.Mappings[i].Dscp.ValueInt64(), 10), data.Mappings[i].Title.ValueString()}

		parent := &data
		data := (*parent).Mappings[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("mappings").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Mappings[%d] = %+v",
				i,
				(*parent).Mappings[i],
			))
			(*parent).Mappings = slices.Delete((*parent).Mappings, i, i+1)
			i--

			continue
		}
		if value := res.Get("cos"); value.Exists() && !data.Cos.IsNull() {
			data.Cos = types.Int64Value(value.Int())
		} else {
			data.Cos = types.Int64Null()
		}
		if value := res.Get("dscp"); value.Exists() && !data.Dscp.IsNull() {
			data.Dscp = types.Int64Value(value.Int())
		} else {
			data.Dscp = types.Int64Null()
		}
		if value := res.Get("title"); value.Exists() && !data.Title.IsNull() {
			data.Title = types.StringValue(value.String())
		} else {
			data.Title = types.StringNull()
		}
		(*parent).Mappings[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SwitchDSCPToCoSMappings) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data SwitchDSCPToCoSMappings) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
