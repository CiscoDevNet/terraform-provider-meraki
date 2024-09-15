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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchMTU struct {
	Id             types.String         `tfsdk:"id"`
	NetworkId      types.String         `tfsdk:"network_id"`
	DefaultMtuSize types.Int64          `tfsdk:"default_mtu_size"`
	Overrides      []SwitchMTUOverrides `tfsdk:"overrides"`
}

type SwitchMTUOverrides struct {
	MtuSize        types.Int64 `tfsdk:"mtu_size"`
	SwitchProfiles types.List  `tfsdk:"switch_profiles"`
	Switches       types.List  `tfsdk:"switches"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchMTU) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/mtu", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchMTU) toBody(ctx context.Context, state SwitchMTU) string {
	body := ""
	if !data.DefaultMtuSize.IsNull() {
		body, _ = sjson.Set(body, "defaultMtuSize", data.DefaultMtuSize.ValueInt64())
	}
	if len(data.Overrides) > 0 {
		body, _ = sjson.Set(body, "overrides", []interface{}{})
		for _, item := range data.Overrides {
			itemBody := ""
			if !item.MtuSize.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "mtuSize", item.MtuSize.ValueInt64())
			}
			if !item.SwitchProfiles.IsNull() {
				var values []string
				item.SwitchProfiles.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "switchProfiles", values)
			}
			if !item.Switches.IsNull() {
				var values []string
				item.Switches.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "switches", values)
			}
			body, _ = sjson.SetRaw(body, "overrides.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchMTU) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("defaultMtuSize"); value.Exists() {
		data.DefaultMtuSize = types.Int64Value(value.Int())
	} else {
		data.DefaultMtuSize = types.Int64Null()
	}
	if value := res.Get("overrides"); value.Exists() {
		data.Overrides = make([]SwitchMTUOverrides, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchMTUOverrides{}
			if value := res.Get("mtuSize"); value.Exists() {
				data.MtuSize = types.Int64Value(value.Int())
			} else {
				data.MtuSize = types.Int64Null()
			}
			if value := res.Get("switchProfiles"); value.Exists() {
				data.SwitchProfiles = helpers.GetStringList(value.Array())
			} else {
				data.SwitchProfiles = types.ListNull(types.StringType)
			}
			if value := res.Get("switches"); value.Exists() {
				data.Switches = helpers.GetStringList(value.Array())
			} else {
				data.Switches = types.ListNull(types.StringType)
			}
			(*parent).Overrides = append((*parent).Overrides, data)
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
func (data *SwitchMTU) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("defaultMtuSize"); value.Exists() && !data.DefaultMtuSize.IsNull() {
		data.DefaultMtuSize = types.Int64Value(value.Int())
	} else {
		data.DefaultMtuSize = types.Int64Null()
	}
	{
		l := len(res.Get("overrides").Array())
		tflog.Debug(ctx, fmt.Sprintf("overrides array resizing from %d to %d", len(data.Overrides), l))
		if len(data.Overrides) > l {
			data.Overrides = data.Overrides[:l]
		}
	}
	for i := range data.Overrides {
		parent := &data
		data := (*parent).Overrides[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("overrides.%d", i))
		if value := res.Get("mtuSize"); value.Exists() && !data.MtuSize.IsNull() {
			data.MtuSize = types.Int64Value(value.Int())
		} else {
			data.MtuSize = types.Int64Null()
		}
		if value := res.Get("switchProfiles"); value.Exists() && !data.SwitchProfiles.IsNull() {
			data.SwitchProfiles = helpers.GetStringList(value.Array())
		} else {
			data.SwitchProfiles = types.ListNull(types.StringType)
		}
		if value := res.Get("switches"); value.Exists() && !data.Switches.IsNull() {
			data.Switches = helpers.GetStringList(value.Array())
		} else {
			data.Switches = types.ListNull(types.StringType)
		}
		(*parent).Overrides[i] = data
	}
}

// End of section. //template:end fromBodyPartial
