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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type AppliancePort struct {
	Id                  types.String `tfsdk:"id"`
	NetworkId           types.String `tfsdk:"network_id"`
	PortId              types.String `tfsdk:"port_id"`
	AccessPolicy        types.String `tfsdk:"access_policy"`
	AllowedVlans        types.String `tfsdk:"allowed_vlans"`
	DropUntaggedTraffic types.Bool   `tfsdk:"drop_untagged_traffic"`
	Enabled             types.Bool   `tfsdk:"enabled"`
	Type                types.String `tfsdk:"type"`
	Vlan                types.Int64  `tfsdk:"vlan"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data AppliancePort) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/ports/%v", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.PortId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data AppliancePort) toBody(ctx context.Context, state AppliancePort) string {
	body := ""
	if !data.AccessPolicy.IsNull() {
		body, _ = sjson.Set(body, "accessPolicy", data.AccessPolicy.ValueString())
	}
	if !data.AllowedVlans.IsNull() {
		body, _ = sjson.Set(body, "allowedVlans", data.AllowedVlans.ValueString())
	}
	if !data.DropUntaggedTraffic.IsNull() {
		body, _ = sjson.Set(body, "dropUntaggedTraffic", data.DropUntaggedTraffic.ValueBool())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.Type.IsNull() {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	if !data.Vlan.IsNull() {
		body, _ = sjson.Set(body, "vlan", data.Vlan.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *AppliancePort) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("accessPolicy"); value.Exists() && value.Value() != nil {
		data.AccessPolicy = types.StringValue(value.String())
	} else {
		data.AccessPolicy = types.StringNull()
	}
	if value := res.Get("allowedVlans"); value.Exists() && value.Value() != nil {
		data.AllowedVlans = types.StringValue(value.String())
	} else {
		data.AllowedVlans = types.StringNull()
	}
	if value := res.Get("dropUntaggedTraffic"); value.Exists() && value.Value() != nil {
		data.DropUntaggedTraffic = types.BoolValue(value.Bool())
	} else {
		data.DropUntaggedTraffic = types.BoolNull()
	}
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("type"); value.Exists() && value.Value() != nil {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("vlan"); value.Exists() && value.Value() != nil {
		data.Vlan = types.Int64Value(value.Int())
	} else {
		data.Vlan = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *AppliancePort) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("accessPolicy"); value.Exists() && !data.AccessPolicy.IsNull() {
		data.AccessPolicy = types.StringValue(value.String())
	} else {
		data.AccessPolicy = types.StringNull()
	}
	if value := res.Get("allowedVlans"); value.Exists() && !data.AllowedVlans.IsNull() {
		data.AllowedVlans = types.StringValue(value.String())
	} else {
		data.AllowedVlans = types.StringNull()
	}
	if value := res.Get("dropUntaggedTraffic"); value.Exists() && !data.DropUntaggedTraffic.IsNull() {
		data.DropUntaggedTraffic = types.BoolValue(value.Bool())
	} else {
		data.DropUntaggedTraffic = types.BoolNull()
	}
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("vlan"); value.Exists() && !data.Vlan.IsNull() {
		data.Vlan = types.Int64Value(value.Int())
	} else {
		data.Vlan = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *AppliancePort) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data AppliancePort) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
