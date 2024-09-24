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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceConnectivityMonitoringDestinations struct {
	Id           types.String                                              `tfsdk:"id"`
	NetworkId    types.String                                              `tfsdk:"network_id"`
	Destinations []ApplianceConnectivityMonitoringDestinationsDestinations `tfsdk:"destinations"`
}

type ApplianceConnectivityMonitoringDestinationsDestinations struct {
	Default     types.Bool   `tfsdk:"default"`
	Description types.String `tfsdk:"description"`
	Ip          types.String `tfsdk:"ip"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceConnectivityMonitoringDestinations) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/connectivityMonitoringDestinations", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceConnectivityMonitoringDestinations) toBody(ctx context.Context, state ApplianceConnectivityMonitoringDestinations) string {
	body := ""
	{
		body, _ = sjson.Set(body, "destinations", []interface{}{})
		for _, item := range data.Destinations {
			itemBody := ""
			if !item.Default.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "default", item.Default.ValueBool())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.Ip.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ip", item.Ip.ValueString())
			}
			body, _ = sjson.SetRaw(body, "destinations.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceConnectivityMonitoringDestinations) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("destinations"); value.Exists() {
		data.Destinations = make([]ApplianceConnectivityMonitoringDestinationsDestinations, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ApplianceConnectivityMonitoringDestinationsDestinations{}
			if value := res.Get("default"); value.Exists() {
				data.Default = types.BoolValue(value.Bool())
			} else {
				data.Default = types.BoolNull()
			}
			if value := res.Get("description"); value.Exists() {
				data.Description = types.StringValue(value.String())
			} else {
				data.Description = types.StringNull()
			}
			if value := res.Get("ip"); value.Exists() {
				data.Ip = types.StringValue(value.String())
			} else {
				data.Ip = types.StringNull()
			}
			(*parent).Destinations = append((*parent).Destinations, data)
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
func (data *ApplianceConnectivityMonitoringDestinations) fromBodyPartial(ctx context.Context, res gjson.Result) {
	for i := 0; i < len(data.Destinations); i++ {
		keys := [...]string{"ip"}
		keyValues := [...]string{data.Destinations[i].Ip.ValueString()}

		parent := &data
		data := (*parent).Destinations[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinations").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Destinations[%d] = %+v",
				i,
				(*parent).Destinations[i],
			))
			(*parent).Destinations = slices.Delete((*parent).Destinations, i, i+1)
			i--

			continue
		}
		if value := res.Get("default"); value.Exists() && !data.Default.IsNull() {
			data.Default = types.BoolValue(value.Bool())
		} else {
			data.Default = types.BoolNull()
		}
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("ip"); value.Exists() && !data.Ip.IsNull() {
			data.Ip = types.StringValue(value.String())
		} else {
			data.Ip = types.StringNull()
		}
		(*parent).Destinations[i] = data
	}
}

// End of section. //template:end fromBodyPartial
