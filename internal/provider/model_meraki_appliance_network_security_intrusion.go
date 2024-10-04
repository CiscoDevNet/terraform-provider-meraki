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

type ApplianceNetworkSecurityIntrusion struct {
	Id                            types.String `tfsdk:"id"`
	NetworkId                     types.String `tfsdk:"network_id"`
	IdsRulesets                   types.String `tfsdk:"ids_rulesets"`
	Mode                          types.String `tfsdk:"mode"`
	ProtectedNetworksUseDefault   types.Bool   `tfsdk:"protected_networks_use_default"`
	ProtectedNetworksExcludedCidr types.List   `tfsdk:"protected_networks_excluded_cidr"`
	ProtectedNetworksIncludedCidr types.List   `tfsdk:"protected_networks_included_cidr"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceNetworkSecurityIntrusion) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/security/intrusion", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceNetworkSecurityIntrusion) toBody(ctx context.Context, state ApplianceNetworkSecurityIntrusion) string {
	body := ""
	if !data.IdsRulesets.IsNull() {
		body, _ = sjson.Set(body, "idsRulesets", data.IdsRulesets.ValueString())
	}
	if !data.Mode.IsNull() {
		body, _ = sjson.Set(body, "mode", data.Mode.ValueString())
	}
	if !data.ProtectedNetworksUseDefault.IsNull() {
		body, _ = sjson.Set(body, "protectedNetworks.useDefault", data.ProtectedNetworksUseDefault.ValueBool())
	}
	if !data.ProtectedNetworksExcludedCidr.IsNull() {
		var values []string
		data.ProtectedNetworksExcludedCidr.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "protectedNetworks.excludedCidr", values)
	}
	if !data.ProtectedNetworksIncludedCidr.IsNull() {
		var values []string
		data.ProtectedNetworksIncludedCidr.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "protectedNetworks.includedCidr", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceNetworkSecurityIntrusion) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("idsRulesets"); value.Exists() && value.Value() != nil {
		data.IdsRulesets = types.StringValue(value.String())
	} else {
		data.IdsRulesets = types.StringNull()
	}
	if value := res.Get("mode"); value.Exists() && value.Value() != nil {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("protectedNetworks.useDefault"); value.Exists() && value.Value() != nil {
		data.ProtectedNetworksUseDefault = types.BoolValue(value.Bool())
	} else {
		data.ProtectedNetworksUseDefault = types.BoolNull()
	}
	if value := res.Get("protectedNetworks.excludedCidr"); value.Exists() && value.Value() != nil {
		data.ProtectedNetworksExcludedCidr = helpers.GetStringList(value.Array())
	} else {
		data.ProtectedNetworksExcludedCidr = types.ListNull(types.StringType)
	}
	if value := res.Get("protectedNetworks.includedCidr"); value.Exists() && value.Value() != nil {
		data.ProtectedNetworksIncludedCidr = helpers.GetStringList(value.Array())
	} else {
		data.ProtectedNetworksIncludedCidr = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceNetworkSecurityIntrusion) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("idsRulesets"); value.Exists() && !data.IdsRulesets.IsNull() {
		data.IdsRulesets = types.StringValue(value.String())
	} else {
		data.IdsRulesets = types.StringNull()
	}
	if value := res.Get("mode"); value.Exists() && !data.Mode.IsNull() {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("protectedNetworks.useDefault"); value.Exists() && !data.ProtectedNetworksUseDefault.IsNull() {
		data.ProtectedNetworksUseDefault = types.BoolValue(value.Bool())
	} else {
		data.ProtectedNetworksUseDefault = types.BoolNull()
	}
	if value := res.Get("protectedNetworks.excludedCidr"); value.Exists() && !data.ProtectedNetworksExcludedCidr.IsNull() {
		data.ProtectedNetworksExcludedCidr = helpers.GetStringList(value.Array())
	} else {
		data.ProtectedNetworksExcludedCidr = types.ListNull(types.StringType)
	}
	if value := res.Get("protectedNetworks.includedCidr"); value.Exists() && !data.ProtectedNetworksIncludedCidr.IsNull() {
		data.ProtectedNetworksIncludedCidr = helpers.GetStringList(value.Array())
	} else {
		data.ProtectedNetworksIncludedCidr = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial
