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
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type CellularGatewayPortForwardingRules struct {
	Id     types.String                              `tfsdk:"id"`
	Serial types.String                              `tfsdk:"serial"`
	Rules  []CellularGatewayPortForwardingRulesRules `tfsdk:"rules"`
}

type CellularGatewayPortForwardingRulesRules struct {
	Access     types.String `tfsdk:"access"`
	LanIp      types.String `tfsdk:"lan_ip"`
	LocalPort  types.String `tfsdk:"local_port"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	PublicPort types.String `tfsdk:"public_port"`
	AllowedIps types.List   `tfsdk:"allowed_ips"`
}

type CellularGatewayPortForwardingRulesIdentity struct {
	Serial types.String `tfsdk:"serial"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CellularGatewayPortForwardingRules) getPath() string {
	return fmt.Sprintf("/devices/%v/cellularGateway/portForwardingRules", url.QueryEscape(data.Serial.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CellularGatewayPortForwardingRules) toBody(ctx context.Context, state CellularGatewayPortForwardingRules) string {
	body := ""
	{
		body, _ = sjson.Set(body, "rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.Access.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "access", item.Access.ValueString())
			}
			if !item.LanIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "lanIp", item.LanIp.ValueString())
			}
			if !item.LocalPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "localPort", item.LocalPort.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			if !item.PublicPort.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "publicPort", item.PublicPort.ValueString())
			}
			if !item.AllowedIps.IsNull() {
				var values []string
				item.AllowedIps.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "allowedIps", values)
			}
			body, _ = sjson.SetRaw(body, "rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CellularGatewayPortForwardingRules) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("rules"); value.Exists() && value.Value() != nil {
		data.Rules = make([]CellularGatewayPortForwardingRulesRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := CellularGatewayPortForwardingRulesRules{}
			if value := res.Get("access"); value.Exists() && value.Value() != nil {
				data.Access = types.StringValue(value.String())
			} else {
				data.Access = types.StringNull()
			}
			if value := res.Get("lanIp"); value.Exists() && value.Value() != nil {
				data.LanIp = types.StringValue(value.String())
			} else {
				data.LanIp = types.StringNull()
			}
			if value := res.Get("localPort"); value.Exists() && value.Value() != nil {
				data.LocalPort = types.StringValue(value.String())
			} else {
				data.LocalPort = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("publicPort"); value.Exists() && value.Value() != nil {
				data.PublicPort = types.StringValue(value.String())
			} else {
				data.PublicPort = types.StringNull()
			}
			if value := res.Get("allowedIps"); value.Exists() && value.Value() != nil {
				data.AllowedIps = helpers.GetStringList(value.Array())
			} else {
				data.AllowedIps = types.ListNull(types.StringType)
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
func (data *CellularGatewayPortForwardingRules) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
		if value := res.Get("access"); value.Exists() && !data.Access.IsNull() {
			data.Access = types.StringValue(value.String())
		} else {
			data.Access = types.StringNull()
		}
		if value := res.Get("lanIp"); value.Exists() && !data.LanIp.IsNull() {
			data.LanIp = types.StringValue(value.String())
		} else {
			data.LanIp = types.StringNull()
		}
		if value := res.Get("localPort"); value.Exists() && !data.LocalPort.IsNull() {
			data.LocalPort = types.StringValue(value.String())
		} else {
			data.LocalPort = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		if value := res.Get("publicPort"); value.Exists() && !data.PublicPort.IsNull() {
			data.PublicPort = types.StringValue(value.String())
		} else {
			data.PublicPort = types.StringNull()
		}
		if value := res.Get("allowedIps"); value.Exists() && !data.AllowedIps.IsNull() {
			data.AllowedIps = helpers.GetStringList(value.Array())
		} else {
			data.AllowedIps = types.ListNull(types.StringType)
		}
		(*parent).Rules[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CellularGatewayPortForwardingRules) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *CellularGatewayPortForwardingRulesIdentity) toIdentity(ctx context.Context, plan *CellularGatewayPortForwardingRules) {
	data.Serial = plan.Serial
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *CellularGatewayPortForwardingRules) fromIdentity(ctx context.Context, identity *CellularGatewayPortForwardingRulesIdentity) {
	data.Serial = identity.Serial
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data CellularGatewayPortForwardingRules) toDestroyBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "rules", []interface{}{})
	return body
}

// End of section. //template:end toDestroyBody
