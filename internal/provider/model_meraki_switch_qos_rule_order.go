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

type SwitchQoSRuleOrder struct {
	Id        types.String `tfsdk:"id"`
	NetworkId types.String `tfsdk:"network_id"`
	RuleIds   types.List   `tfsdk:"rule_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchQoSRuleOrder) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/qosRules/order", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchQoSRuleOrder) toBody(ctx context.Context, state SwitchQoSRuleOrder) string {
	body := ""
	if !data.RuleIds.IsNull() {
		var values []string
		data.RuleIds.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "ruleIds", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchQoSRuleOrder) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("ruleIds"); value.Exists() && value.Value() != nil {
		data.RuleIds = helpers.GetStringList(value.Array())
	} else {
		data.RuleIds = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchQoSRuleOrder) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("ruleIds"); value.Exists() && !data.RuleIds.IsNull() {
		data.RuleIds = helpers.GetStringList(value.Array())
	} else {
		data.RuleIds = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial
