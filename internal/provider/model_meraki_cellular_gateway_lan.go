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

type CellularGatewayLAN struct {
	Id                 types.String                           `tfsdk:"id"`
	Serial             types.String                           `tfsdk:"serial"`
	FixedIpAssignments []CellularGatewayLANFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	ReservedIpRanges   []CellularGatewayLANReservedIpRanges   `tfsdk:"reserved_ip_ranges"`
}

type CellularGatewayLANFixedIpAssignments struct {
	Ip   types.String `tfsdk:"ip"`
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
}

type CellularGatewayLANReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CellularGatewayLAN) getPath() string {
	return fmt.Sprintf("/devices/%v/cellularGateway/lan", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CellularGatewayLAN) toBody(ctx context.Context, state CellularGatewayLAN) string {
	body := ""
	{
		body, _ = sjson.Set(body, "fixedIpAssignments", []interface{}{})
		for _, item := range data.FixedIpAssignments {
			itemBody := ""
			if !item.Ip.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ip", item.Ip.ValueString())
			}
			if !item.Mac.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "mac", item.Mac.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "fixedIpAssignments.-1", itemBody)
		}
	}
	{
		body, _ = sjson.Set(body, "reservedIpRanges", []interface{}{})
		for _, item := range data.ReservedIpRanges {
			itemBody := ""
			if !item.Comment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "comment", item.Comment.ValueString())
			}
			if !item.End.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "end", item.End.ValueString())
			}
			if !item.Start.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "start", item.Start.ValueString())
			}
			body, _ = sjson.SetRaw(body, "reservedIpRanges.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CellularGatewayLAN) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("fixedIpAssignments"); value.Exists() && value.Value() != nil {
		data.FixedIpAssignments = make([]CellularGatewayLANFixedIpAssignments, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := CellularGatewayLANFixedIpAssignments{}
			if value := res.Get("ip"); value.Exists() && value.Value() != nil {
				data.Ip = types.StringValue(value.String())
			} else {
				data.Ip = types.StringNull()
			}
			if value := res.Get("mac"); value.Exists() && value.Value() != nil {
				data.Mac = types.StringValue(value.String())
			} else {
				data.Mac = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).FixedIpAssignments = append((*parent).FixedIpAssignments, data)
			return true
		})
	}
	if value := res.Get("reservedIpRanges"); value.Exists() && value.Value() != nil {
		data.ReservedIpRanges = make([]CellularGatewayLANReservedIpRanges, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := CellularGatewayLANReservedIpRanges{}
			if value := res.Get("comment"); value.Exists() && value.Value() != nil {
				data.Comment = types.StringValue(value.String())
			} else {
				data.Comment = types.StringNull()
			}
			if value := res.Get("end"); value.Exists() && value.Value() != nil {
				data.End = types.StringValue(value.String())
			} else {
				data.End = types.StringNull()
			}
			if value := res.Get("start"); value.Exists() && value.Value() != nil {
				data.Start = types.StringValue(value.String())
			} else {
				data.Start = types.StringNull()
			}
			(*parent).ReservedIpRanges = append((*parent).ReservedIpRanges, data)
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
func (data *CellularGatewayLAN) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.FixedIpAssignments); i++ {
		keys := [...]string{"ip", "mac", "name"}
		keyValues := [...]string{data.FixedIpAssignments[i].Ip.ValueString(), data.FixedIpAssignments[i].Mac.ValueString(), data.FixedIpAssignments[i].Name.ValueString()}

		parent := &data
		data := (*parent).FixedIpAssignments[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("fixedIpAssignments").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing FixedIpAssignments[%d] = %+v",
				i,
				(*parent).FixedIpAssignments[i],
			))
			(*parent).FixedIpAssignments = slices.Delete((*parent).FixedIpAssignments, i, i+1)
			i--

			continue
		}
		if value := res.Get("ip"); value.Exists() && !data.Ip.IsNull() {
			data.Ip = types.StringValue(value.String())
		} else {
			data.Ip = types.StringNull()
		}
		if value := res.Get("mac"); value.Exists() && !data.Mac.IsNull() {
			data.Mac = types.StringValue(value.String())
		} else {
			data.Mac = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).FixedIpAssignments[i] = data
	}
	for i := 0; i < len(data.ReservedIpRanges); i++ {
		keys := [...]string{"comment", "end", "start"}
		keyValues := [...]string{data.ReservedIpRanges[i].Comment.ValueString(), data.ReservedIpRanges[i].End.ValueString(), data.ReservedIpRanges[i].Start.ValueString()}

		parent := &data
		data := (*parent).ReservedIpRanges[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("reservedIpRanges").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ReservedIpRanges[%d] = %+v",
				i,
				(*parent).ReservedIpRanges[i],
			))
			(*parent).ReservedIpRanges = slices.Delete((*parent).ReservedIpRanges, i, i+1)
			i--

			continue
		}
		if value := res.Get("comment"); value.Exists() && !data.Comment.IsNull() {
			data.Comment = types.StringValue(value.String())
		} else {
			data.Comment = types.StringNull()
		}
		if value := res.Get("end"); value.Exists() && !data.End.IsNull() {
			data.End = types.StringValue(value.String())
		} else {
			data.End = types.StringNull()
		}
		if value := res.Get("start"); value.Exists() && !data.Start.IsNull() {
			data.Start = types.StringValue(value.String())
		} else {
			data.Start = types.StringNull()
		}
		(*parent).ReservedIpRanges[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CellularGatewayLAN) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data CellularGatewayLAN) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
