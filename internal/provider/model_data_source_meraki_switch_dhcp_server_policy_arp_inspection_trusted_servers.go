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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceSwitchDHCPServerPolicyARPInspectionTrustedServers struct {
	NetworkId types.String                                                       `tfsdk:"network_id"`
	Items     []DataSourceSwitchDHCPServerPolicyARPInspectionTrustedServersItems `tfsdk:"items"`
}

type DataSourceSwitchDHCPServerPolicyARPInspectionTrustedServersItems struct {
	Id          types.String `tfsdk:"id"`
	Mac         types.String `tfsdk:"mac"`
	Vlan        types.Int64  `tfsdk:"vlan"`
	Ipv4Address types.String `tfsdk:"ipv4_address"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceSwitchDHCPServerPolicyARPInspectionTrustedServers) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/dhcpServerPolicy/arpInspection/trustedServers", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceSwitchDHCPServerPolicyARPInspectionTrustedServers) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceSwitchDHCPServerPolicyARPInspectionTrustedServersItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceSwitchDHCPServerPolicyARPInspectionTrustedServersItems{}
		data.Id = types.StringValue(res.Get("trustedServerId").String())
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
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
