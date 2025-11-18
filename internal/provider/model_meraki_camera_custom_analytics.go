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

type CameraCustomAnalytics struct {
	Id         types.String                      `tfsdk:"id"`
	Serial     types.String                      `tfsdk:"serial"`
	ArtifactId types.String                      `tfsdk:"artifact_id"`
	Enabled    types.Bool                        `tfsdk:"enabled"`
	Parameters []CameraCustomAnalyticsParameters `tfsdk:"parameters"`
}

type CameraCustomAnalyticsParameters struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CameraCustomAnalytics) getPath() string {
	return fmt.Sprintf("/devices/%v/camera/customAnalytics", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CameraCustomAnalytics) toBody(ctx context.Context, state CameraCustomAnalytics) string {
	body := ""
	if !data.ArtifactId.IsNull() {
		body, _ = sjson.Set(body, "artifactId", data.ArtifactId.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if data.Parameters != nil {
		body, _ = sjson.Set(body, "parameters", []interface{}{})
		for _, item := range data.Parameters {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Value.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Value.ValueString())
			}
			body, _ = sjson.SetRaw(body, "parameters.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CameraCustomAnalytics) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("artifactId"); value.Exists() && value.Value() != nil {
		data.ArtifactId = types.StringValue(value.String())
	} else {
		data.ArtifactId = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("parameters"); value.Exists() && value.Value() != nil {
		data.Parameters = make([]CameraCustomAnalyticsParameters, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := CameraCustomAnalyticsParameters{}
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
			(*parent).Parameters = append((*parent).Parameters, data)
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
func (data *CameraCustomAnalytics) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("artifactId"); value.Exists() && !data.ArtifactId.IsNull() {
		data.ArtifactId = types.StringValue(value.String())
	} else {
		data.ArtifactId = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	for i := 0; i < len(data.Parameters); i++ {
		keys := [...]string{"name", "value"}
		keyValues := [...]string{data.Parameters[i].Name.ValueString(), data.Parameters[i].Value.ValueString()}

		parent := &data
		data := (*parent).Parameters[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("parameters").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Parameters[%d] = %+v",
				i,
				(*parent).Parameters[i],
			))
			(*parent).Parameters = slices.Delete((*parent).Parameters, i, i+1)
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
		(*parent).Parameters[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CameraCustomAnalytics) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data CameraCustomAnalytics) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
