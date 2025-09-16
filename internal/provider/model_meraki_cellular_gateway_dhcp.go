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
	"github.com/netascode/go-meraki"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type CellularGatewayDHCP struct {
	Id                   types.String `tfsdk:"id"`
	NetworkId            types.String `tfsdk:"network_id"`
	DhcpLeaseTime        types.String `tfsdk:"dhcp_lease_time"`
	DnsNameservers       types.String `tfsdk:"dns_nameservers"`
	DnsCustomNameservers types.List   `tfsdk:"dns_custom_nameservers"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CellularGatewayDHCP) getPath() string {
	return fmt.Sprintf("/networks/%v/cellularGateway/dhcp", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CellularGatewayDHCP) toBody(ctx context.Context, state CellularGatewayDHCP) string {
	body := ""
	if !data.DhcpLeaseTime.IsNull() {
		body, _ = sjson.Set(body, "dhcpLeaseTime", data.DhcpLeaseTime.ValueString())
	}
	if !data.DnsNameservers.IsNull() {
		body, _ = sjson.Set(body, "dnsNameservers", data.DnsNameservers.ValueString())
	}
	if !data.DnsCustomNameservers.IsNull() {
		var values []string
		data.DnsCustomNameservers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "dnsCustomNameservers", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CellularGatewayDHCP) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("dhcpLeaseTime"); value.Exists() && value.Value() != nil {
		data.DhcpLeaseTime = types.StringValue(value.String())
	} else {
		data.DhcpLeaseTime = types.StringNull()
	}
	if value := res.Get("dnsNameservers"); value.Exists() && value.Value() != nil {
		data.DnsNameservers = types.StringValue(value.String())
	} else {
		data.DnsNameservers = types.StringNull()
	}
	if value := res.Get("dnsCustomNameservers"); value.Exists() && value.Value() != nil {
		data.DnsCustomNameservers = helpers.GetStringList(value.Array())
	} else {
		data.DnsCustomNameservers = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *CellularGatewayDHCP) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("dhcpLeaseTime"); value.Exists() && !data.DhcpLeaseTime.IsNull() {
		data.DhcpLeaseTime = types.StringValue(value.String())
	} else {
		data.DhcpLeaseTime = types.StringNull()
	}
	if value := res.Get("dnsNameservers"); value.Exists() && !data.DnsNameservers.IsNull() {
		data.DnsNameservers = types.StringValue(value.String())
	} else {
		data.DnsNameservers = types.StringNull()
	}
	if value := res.Get("dnsCustomNameservers"); value.Exists() && !data.DnsCustomNameservers.IsNull() {
		data.DnsCustomNameservers = helpers.GetStringList(value.Array())
	} else {
		data.DnsCustomNameservers = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CellularGatewayDHCP) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data CellularGatewayDHCP) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
