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

type WirelessSSIDFirewallIsolationAllowlistEntry struct {
	Id             types.String `tfsdk:"id"`
	OrganizationId types.String `tfsdk:"organization_id"`
	Description    types.String `tfsdk:"description"`
	ClientMac      types.String `tfsdk:"client_mac"`
	NetworkId      types.String `tfsdk:"network_id"`
	SsidNumber     types.Int64  `tfsdk:"ssid_number"`
}

type WirelessSSIDFirewallIsolationAllowlistEntryIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessSSIDFirewallIsolationAllowlistEntry) getPath() string {
	return fmt.Sprintf("/organizations/%v/wireless/ssids/firewall/isolation/allowlist/entries", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSSIDFirewallIsolationAllowlistEntry) toBody(ctx context.Context, state WirelessSSIDFirewallIsolationAllowlistEntry) string {
	body := ""
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.ClientMac.IsNull() {
		body, _ = sjson.Set(body, "client.mac", data.ClientMac.ValueString())
	}
	if !data.NetworkId.IsNull() {
		body, _ = sjson.Set(body, "network.id", data.NetworkId.ValueString())
	}
	if !data.SsidNumber.IsNull() {
		body, _ = sjson.Set(body, "ssid.number", data.SsidNumber.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSSIDFirewallIsolationAllowlistEntry) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && value.Value() != nil {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("client.mac"); value.Exists() && value.Value() != nil {
		data.ClientMac = types.StringValue(value.String())
	} else {
		data.ClientMac = types.StringNull()
	}
	if value := res.Get("network.id"); value.Exists() && value.Value() != nil {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
	}
	if value := res.Get("ssid.number"); value.Exists() && value.Value() != nil {
		data.SsidNumber = types.Int64Value(value.Int())
	} else {
		data.SsidNumber = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessSSIDFirewallIsolationAllowlistEntry) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("client.mac"); value.Exists() && !data.ClientMac.IsNull() {
		data.ClientMac = types.StringValue(value.String())
	} else {
		data.ClientMac = types.StringNull()
	}
	if value := res.Get("network.id"); value.Exists() && !data.NetworkId.IsNull() {
		data.NetworkId = types.StringValue(value.String())
	} else {
		data.NetworkId = types.StringNull()
	}
	if value := res.Get("ssid.number"); value.Exists() && !data.SsidNumber.IsNull() {
		data.SsidNumber = types.Int64Value(value.Int())
	} else {
		data.SsidNumber = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessSSIDFirewallIsolationAllowlistEntry) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *WirelessSSIDFirewallIsolationAllowlistEntryIdentity) toIdentity(ctx context.Context, plan *WirelessSSIDFirewallIsolationAllowlistEntry) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *WirelessSSIDFirewallIsolationAllowlistEntry) fromIdentity(ctx context.Context, identity *WirelessSSIDFirewallIsolationAllowlistEntryIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data WirelessSSIDFirewallIsolationAllowlistEntry) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
