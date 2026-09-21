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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSourceCellularGatewayDHCP - used only by the data source, always including every
// attribute (unlike model_resource.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model_resource.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model_resource.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceCellularGatewayDHCP struct {
	Id                   types.String `tfsdk:"id"`
	NetworkId            types.String `tfsdk:"network_id"`
	DhcpLeaseTime        types.String `tfsdk:"dhcp_lease_time"`
	DnsNameservers       types.String `tfsdk:"dns_nameservers"`
	DnsCustomNameservers types.List   `tfsdk:"dns_custom_nameservers"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceCellularGatewayDHCP) getPath() string {
	return fmt.Sprintf("/networks/%v/cellularGateway/dhcp", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceCellularGatewayDHCP) fromBody(ctx context.Context, res meraki.Res) {
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
