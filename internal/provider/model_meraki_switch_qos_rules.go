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

type SwitchQoSRules struct {
	NetworkId types.String          `tfsdk:"network_id"`
	Items     []SwitchQoSRulesItems `tfsdk:"items"`
}

type SwitchQoSRulesItems struct {
	Id           types.String `tfsdk:"id"`
	Dscp         types.Int64  `tfsdk:"dscp"`
	DstPort      types.Int64  `tfsdk:"dst_port"`
	DstPortRange types.String `tfsdk:"dst_port_range"`
	Protocol     types.String `tfsdk:"protocol"`
	SrcPort      types.Int64  `tfsdk:"src_port"`
	SrcPortRange types.String `tfsdk:"src_port_range"`
	Vlan         types.Int64  `tfsdk:"vlan"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchQoSRules) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/qosRules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchQoSRules) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]SwitchQoSRulesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := SwitchQoSRulesItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("dscp"); value.Exists() && value.Value() != nil {
			data.Dscp = types.Int64Value(value.Int())
		} else {
			data.Dscp = types.Int64Null()
		}
		if value := res.Get("dstPort"); value.Exists() && value.Value() != nil {
			data.DstPort = types.Int64Value(value.Int())
		} else {
			data.DstPort = types.Int64Null()
		}
		if value := res.Get("dstPortRange"); value.Exists() && value.Value() != nil {
			data.DstPortRange = types.StringValue(value.String())
		} else {
			data.DstPortRange = types.StringNull()
		}
		if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		if value := res.Get("srcPort"); value.Exists() && value.Value() != nil {
			data.SrcPort = types.Int64Value(value.Int())
		} else {
			data.SrcPort = types.Int64Null()
		}
		if value := res.Get("srcPortRange"); value.Exists() && value.Value() != nil {
			data.SrcPortRange = types.StringValue(value.String())
		} else {
			data.SrcPortRange = types.StringNull()
		}
		if value := res.Get("vlan"); value.Exists() && value.Value() != nil {
			data.Vlan = types.Int64Value(value.Int())
		} else {
			data.Vlan = types.Int64Null()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
