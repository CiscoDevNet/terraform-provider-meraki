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

type ApplianceTrafficShapingVPNExclusions struct {
	Id                types.String                                            `tfsdk:"id"`
	NetworkId         types.String                                            `tfsdk:"network_id"`
	Custom            []ApplianceTrafficShapingVPNExclusionsCustom            `tfsdk:"custom"`
	MajorApplications []ApplianceTrafficShapingVPNExclusionsMajorApplications `tfsdk:"major_applications"`
}

type ApplianceTrafficShapingVPNExclusionsCustom struct {
	Destination types.String `tfsdk:"destination"`
	Port        types.String `tfsdk:"port"`
	Protocol    types.String `tfsdk:"protocol"`
}

type ApplianceTrafficShapingVPNExclusionsMajorApplications struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceTrafficShapingVPNExclusions) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/trafficShaping/vpnExclusions", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceTrafficShapingVPNExclusions) toBody(ctx context.Context, state ApplianceTrafficShapingVPNExclusions) string {
	body := ""
	if len(data.Custom) > 0 {
		body, _ = sjson.Set(body, "custom", []interface{}{})
		for _, item := range data.Custom {
			itemBody := ""
			if !item.Destination.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "destination", item.Destination.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			body, _ = sjson.SetRaw(body, "custom.-1", itemBody)
		}
	}
	if len(data.MajorApplications) > 0 {
		body, _ = sjson.Set(body, "majorApplications", []interface{}{})
		for _, item := range data.MajorApplications {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "majorApplications.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceTrafficShapingVPNExclusions) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("custom"); value.Exists() && value.Value() != nil {
		data.Custom = make([]ApplianceTrafficShapingVPNExclusionsCustom, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceTrafficShapingVPNExclusionsCustom{}
			if value := res.Get("destination"); value.Exists() && value.Value() != nil {
				data.Destination = types.StringValue(value.String())
			} else {
				data.Destination = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && value.Value() != nil {
				data.Port = types.StringValue(value.String())
			} else {
				data.Port = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			(*parent).Custom = append((*parent).Custom, data)
			return true
		})
	}
	if value := res.Get("majorApplications"); value.Exists() && value.Value() != nil {
		data.MajorApplications = make([]ApplianceTrafficShapingVPNExclusionsMajorApplications, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceTrafficShapingVPNExclusionsMajorApplications{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).MajorApplications = append((*parent).MajorApplications, data)
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
func (data *ApplianceTrafficShapingVPNExclusions) fromBodyPartial(ctx context.Context, res meraki.Res) {
	{
		l := len(res.Get("custom").Array())
		tflog.Debug(ctx, fmt.Sprintf("custom array resizing from %d to %d", len(data.Custom), l))
		if len(data.Custom) > l {
			data.Custom = data.Custom[:l]
		}
	}
	for i := range data.Custom {
		parent := &data
		data := (*parent).Custom[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("custom.%d", i))
		if value := res.Get("destination"); value.Exists() && !data.Destination.IsNull() {
			data.Destination = types.StringValue(value.String())
		} else {
			data.Destination = types.StringNull()
		}
		if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
			data.Port = types.StringValue(value.String())
		} else {
			data.Port = types.StringNull()
		}
		if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		(*parent).Custom[i] = data
	}
	for i := 0; i < len(data.MajorApplications); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.MajorApplications[i].Id.ValueString()}

		parent := &data
		data := (*parent).MajorApplications[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("majorApplications").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing MajorApplications[%d] = %+v",
				i,
				(*parent).MajorApplications[i],
			))
			(*parent).MajorApplications = slices.Delete((*parent).MajorApplications, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).MajorApplications[i] = data
	}
}

// End of section. //template:end fromBodyPartial
