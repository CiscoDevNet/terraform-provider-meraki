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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchAccessControlLists struct {
	Id        types.String                    `tfsdk:"id"`
	NetworkId types.String                    `tfsdk:"network_id"`
	Rules     []SwitchAccessControlListsRules `tfsdk:"rules"`
}

type SwitchAccessControlListsRules struct {
	Comment   types.String `tfsdk:"comment"`
	DstCidr   types.String `tfsdk:"dst_cidr"`
	DstPort   types.String `tfsdk:"dst_port"`
	IpVersion types.String `tfsdk:"ip_version"`
	Policy    types.String `tfsdk:"policy"`
	Protocol  types.String `tfsdk:"protocol"`
	SrcCidr   types.String `tfsdk:"src_cidr"`
	SrcPort   types.String `tfsdk:"src_port"`
	Vlan      types.String `tfsdk:"vlan"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchAccessControlLists) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/accessControlLists", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchAccessControlLists) toBody(ctx context.Context, state SwitchAccessControlLists) string {
	body := ""
	{
		body, _ = sjson.Set(body, "rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.Comment.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "comment", item.Comment.ValueString())
			}
			if !item.DstCidr.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dstCidr", item.DstCidr.ValueString())
			}
			if !item.DstPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dstPort", item.DstPort.ValueString())
			}
			if !item.IpVersion.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipVersion", item.IpVersion.ValueString())
			}
			if !item.Policy.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "policy", item.Policy.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			if !item.SrcCidr.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "srcCidr", item.SrcCidr.ValueString())
			}
			if !item.SrcPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "srcPort", item.SrcPort.ValueString())
			}
			if !item.Vlan.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vlan", item.Vlan.ValueString())
			}
			body, _ = sjson.SetRaw(body, "rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchAccessControlLists) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("rules"); value.Exists() && value.Value() != nil {
		data.Rules = make([]SwitchAccessControlListsRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			if value := res.Get("comment"); value.String() == "Default rule" {
				return true
			}
			parent := &data
			data := SwitchAccessControlListsRules{}
			if value := res.Get("comment"); value.Exists() && value.Value() != nil {
				data.Comment = types.StringValue(value.String())
			} else {
				data.Comment = types.StringNull()
			}
			if value := res.Get("dstCidr"); value.Exists() && value.Value() != nil {
				data.DstCidr = types.StringValue(value.String())
			} else {
				data.DstCidr = types.StringNull()
			}
			if value := res.Get("dstPort"); value.Exists() && value.Value() != nil {
				data.DstPort = types.StringValue(value.String())
			} else {
				data.DstPort = types.StringNull()
			}
			if value := res.Get("ipVersion"); value.Exists() && value.Value() != nil {
				data.IpVersion = types.StringValue(value.String())
			} else {
				data.IpVersion = types.StringNull()
			}
			if value := res.Get("policy"); value.Exists() && value.Value() != nil {
				data.Policy = types.StringValue(value.String())
			} else {
				data.Policy = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("srcCidr"); value.Exists() && value.Value() != nil {
				data.SrcCidr = types.StringValue(value.String())
			} else {
				data.SrcCidr = types.StringNull()
			}
			if value := res.Get("srcPort"); value.Exists() && value.Value() != nil {
				data.SrcPort = types.StringValue(value.String())
			} else {
				data.SrcPort = types.StringNull()
			}
			if value := res.Get("vlan"); value.Exists() && value.Value() != nil {
				data.Vlan = types.StringValue(value.String())
			} else {
				data.Vlan = types.StringNull()
			}
			(*parent).Rules = append((*parent).Rules, data)
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
func (data *SwitchAccessControlLists) fromBodyPartial(ctx context.Context, res meraki.Res) {
	{
		l := len(res.Get("rules").Array())
		tflog.Debug(ctx, fmt.Sprintf("rules array resizing from %d to %d", len(data.Rules), l))
		if len(data.Rules) > l {
			data.Rules = data.Rules[:l]
		}
	}
	for i := range data.Rules {
		parent := &data
		data := (*parent).Rules[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("rules.%d", i))
		if value := res.Get("comment"); value.Exists() && !data.Comment.IsNull() {
			data.Comment = types.StringValue(value.String())
		} else {
			data.Comment = types.StringNull()
		}
		if value := res.Get("dstCidr"); value.Exists() && !data.DstCidr.IsNull() {
			data.DstCidr = types.StringValue(value.String())
		} else {
			data.DstCidr = types.StringNull()
		}
		if value := res.Get("dstPort"); value.Exists() && !data.DstPort.IsNull() {
			data.DstPort = types.StringValue(value.String())
		} else {
			data.DstPort = types.StringNull()
		}
		if value := res.Get("ipVersion"); value.Exists() && !data.IpVersion.IsNull() {
			data.IpVersion = types.StringValue(value.String())
		} else {
			data.IpVersion = types.StringNull()
		}
		if value := res.Get("policy"); value.Exists() && !data.Policy.IsNull() {
			data.Policy = types.StringValue(value.String())
		} else {
			data.Policy = types.StringNull()
		}
		if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		if value := res.Get("srcCidr"); value.Exists() && !data.SrcCidr.IsNull() {
			data.SrcCidr = types.StringValue(value.String())
		} else {
			data.SrcCidr = types.StringNull()
		}
		if value := res.Get("srcPort"); value.Exists() && !data.SrcPort.IsNull() {
			data.SrcPort = types.StringValue(value.String())
		} else {
			data.SrcPort = types.StringNull()
		}
		if value := res.Get("vlan"); value.Exists() && !data.Vlan.IsNull() {
			data.Vlan = types.StringValue(value.String())
		} else {
			data.Vlan = types.StringNull()
		}
		(*parent).Rules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SwitchAccessControlLists) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin addDeleteValues

func (data SwitchAccessControlLists) addDeleteValues(ctx context.Context, body string) string {
	return body
}

// End of section. //template:end addDeleteValues
