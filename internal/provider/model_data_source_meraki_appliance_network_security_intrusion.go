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
// It emits a separate model - DataSourceApplianceNetworkSecurityIntrusion - used only by the data source, always including every
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

type DataSourceApplianceNetworkSecurityIntrusion struct {
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

func (data DataSourceApplianceNetworkSecurityIntrusion) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/security/intrusion", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceNetworkSecurityIntrusion) fromBody(ctx context.Context, res meraki.Res) {
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
