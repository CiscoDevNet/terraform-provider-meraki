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

type SensorRelationships struct {
	Id                       types.String                                  `tfsdk:"id"`
	Serial                   types.String                                  `tfsdk:"serial"`
	LivestreamRelatedDevices []SensorRelationshipsLivestreamRelatedDevices `tfsdk:"livestream_related_devices"`
}

type SensorRelationshipsLivestreamRelatedDevices struct {
	Serial types.String `tfsdk:"serial"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SensorRelationships) getPath() string {
	return fmt.Sprintf("/devices/%v/sensor/relationships", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SensorRelationships) toBody(ctx context.Context, state SensorRelationships) string {
	body := ""
	if len(data.LivestreamRelatedDevices) > 0 {
		body, _ = sjson.Set(body, "livestream.relatedDevices", []interface{}{})
		for _, item := range data.LivestreamRelatedDevices {
			itemBody := ""
			if !item.Serial.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "serial", item.Serial.ValueString())
			}
			body, _ = sjson.SetRaw(body, "livestream.relatedDevices.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SensorRelationships) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("livestream.relatedDevices"); value.Exists() && value.Value() != nil {
		data.LivestreamRelatedDevices = make([]SensorRelationshipsLivestreamRelatedDevices, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SensorRelationshipsLivestreamRelatedDevices{}
			if value := res.Get("serial"); value.Exists() && value.Value() != nil {
				data.Serial = types.StringValue(value.String())
			} else {
				data.Serial = types.StringNull()
			}
			(*parent).LivestreamRelatedDevices = append((*parent).LivestreamRelatedDevices, data)
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
func (data *SensorRelationships) fromBodyPartial(ctx context.Context, res meraki.Res) {
	for i := 0; i < len(data.LivestreamRelatedDevices); i++ {
		keys := [...]string{"serial"}
		keyValues := [...]string{data.LivestreamRelatedDevices[i].Serial.ValueString()}

		parent := &data
		data := (*parent).LivestreamRelatedDevices[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("livestream.relatedDevices").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing LivestreamRelatedDevices[%d] = %+v",
				i,
				(*parent).LivestreamRelatedDevices[i],
			))
			(*parent).LivestreamRelatedDevices = slices.Delete((*parent).LivestreamRelatedDevices, i, i+1)
			i--

			continue
		}
		if value := res.Get("serial"); value.Exists() && !data.Serial.IsNull() {
			data.Serial = types.StringValue(value.String())
		} else {
			data.Serial = types.StringNull()
		}
		(*parent).LivestreamRelatedDevices[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SensorRelationships) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data SensorRelationships) toDestroyBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "livestream.relatedDevices", []interface{}{})
	return body
}

// End of section. //template:end toDestroyBody
