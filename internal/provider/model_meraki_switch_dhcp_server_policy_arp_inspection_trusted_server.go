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

type SwitchDHCPServerPolicyARPInspectionTrustedServer struct {
	Id          types.String `tfsdk:"id"`
	NetworkId   types.String `tfsdk:"network_id"`
	Mac         types.String `tfsdk:"mac"`
	Vlan        types.Int64  `tfsdk:"vlan"`
	Ipv4Address types.String `tfsdk:"ipv4_address"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchDHCPServerPolicyARPInspectionTrustedServer) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/dhcpServerPolicy/arpInspection/trustedServers", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchDHCPServerPolicyARPInspectionTrustedServer) toBody(ctx context.Context, state SwitchDHCPServerPolicyARPInspectionTrustedServer) string {
	body := ""
	if !data.Mac.IsNull() {
		body, _ = sjson.Set(body, "mac", data.Mac.ValueString())
	}
	if !data.Vlan.IsNull() {
		body, _ = sjson.Set(body, "vlan", data.Vlan.ValueInt64())
	}
	if !data.Ipv4Address.IsNull() {
		body, _ = sjson.Set(body, "ipv4.address", data.Ipv4Address.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchDHCPServerPolicyARPInspectionTrustedServer) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("mac"); value.Exists() && value.Value() != nil {
		data.Mac = types.StringValue(value.String())
	} else {
		data.Mac = types.StringNull()
	}
	if value := res.Get("vlan"); value.Exists() && value.Value() != nil {
		data.Vlan = types.Int64Value(value.Int())
	} else {
		data.Vlan = types.Int64Null()
	}
	if value := res.Get("ipv4.address"); value.Exists() && value.Value() != nil {
		data.Ipv4Address = types.StringValue(value.String())
	} else {
		data.Ipv4Address = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchDHCPServerPolicyARPInspectionTrustedServer) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("mac"); value.Exists() && !data.Mac.IsNull() {
		data.Mac = types.StringValue(value.String())
	} else {
		data.Mac = types.StringNull()
	}
	if value := res.Get("vlan"); value.Exists() && !data.Vlan.IsNull() {
		data.Vlan = types.Int64Value(value.Int())
	} else {
		data.Vlan = types.Int64Null()
	}
	if value := res.Get("ipv4.address"); value.Exists() && !data.Ipv4Address.IsNull() {
		data.Ipv4Address = types.StringValue(value.String())
	} else {
		data.Ipv4Address = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SwitchDHCPServerPolicyARPInspectionTrustedServer) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data SwitchDHCPServerPolicyARPInspectionTrustedServer) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
