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

type SwitchQoSRule struct {
	Id           types.String `tfsdk:"id"`
	NetworkId    types.String `tfsdk:"network_id"`
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

func (data SwitchQoSRule) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/qosRules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchQoSRule) toBody(ctx context.Context, state SwitchQoSRule) string {
	body := ""
	if !data.Dscp.IsNull() {
		body, _ = sjson.Set(body, "dscp", data.Dscp.ValueInt64())
	}
	if !data.DstPort.IsNull() {
		body, _ = sjson.Set(body, "dstPort", data.DstPort.ValueInt64())
	}
	if !data.DstPortRange.IsNull() {
		body, _ = sjson.Set(body, "dstPortRange", data.DstPortRange.ValueString())
	}
	if !data.Protocol.IsNull() {
		body, _ = sjson.Set(body, "protocol", data.Protocol.ValueString())
	}
	if !data.SrcPort.IsNull() {
		body, _ = sjson.Set(body, "srcPort", data.SrcPort.ValueInt64())
	}
	if !data.SrcPortRange.IsNull() {
		body, _ = sjson.Set(body, "srcPortRange", data.SrcPortRange.ValueString())
	}
	if !data.Vlan.IsNull() {
		body, _ = sjson.Set(body, "vlan", data.Vlan.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchQoSRule) fromBody(ctx context.Context, res meraki.Res) {
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchQoSRule) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("dscp"); value.Exists() && !data.Dscp.IsNull() {
		data.Dscp = types.Int64Value(value.Int())
	} else {
		data.Dscp = types.Int64Null()
	}
	if value := res.Get("dstPort"); value.Exists() && !data.DstPort.IsNull() {
		data.DstPort = types.Int64Value(value.Int())
	} else {
		data.DstPort = types.Int64Null()
	}
	if value := res.Get("dstPortRange"); value.Exists() && !data.DstPortRange.IsNull() {
		data.DstPortRange = types.StringValue(value.String())
	} else {
		data.DstPortRange = types.StringNull()
	}
	if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("srcPort"); value.Exists() && !data.SrcPort.IsNull() {
		data.SrcPort = types.Int64Value(value.Int())
	} else {
		data.SrcPort = types.Int64Null()
	}
	if value := res.Get("srcPortRange"); value.Exists() && !data.SrcPortRange.IsNull() {
		data.SrcPortRange = types.StringValue(value.String())
	} else {
		data.SrcPortRange = types.StringNull()
	}
	if value := res.Get("vlan"); value.Exists() && !data.Vlan.IsNull() {
		data.Vlan = types.Int64Value(value.Int())
	} else {
		data.Vlan = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial
