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
// It emits a separate model - DataSourceWirelessSettings - used only by the data source, always including every
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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessSettings struct {
	Id                                   types.String `tfsdk:"id"`
	NetworkId                            types.String `tfsdk:"network_id"`
	Ipv6BridgeEnabled                    types.Bool   `tfsdk:"ipv6_bridge_enabled"`
	LedLightsOn                          types.Bool   `tfsdk:"led_lights_on"`
	LocationAnalyticsEnabled             types.Bool   `tfsdk:"location_analytics_enabled"`
	MeshingEnabled                       types.Bool   `tfsdk:"meshing_enabled"`
	UpgradeStrategy                      types.String `tfsdk:"upgrade_strategy"`
	MulticastToUnicastConversionEnabled  types.Bool   `tfsdk:"multicast_to_unicast_conversion_enabled"`
	NamedVlansPoolDhcpMonitoringDuration types.Int64  `tfsdk:"named_vlans_pool_dhcp_monitoring_duration"`
	NamedVlansPoolDhcpMonitoringEnabled  types.Bool   `tfsdk:"named_vlans_pool_dhcp_monitoring_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/settings", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessSettings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("ipv6BridgeEnabled"); value.Exists() && value.Value() != nil {
		data.Ipv6BridgeEnabled = types.BoolValue(value.Bool())
	} else {
		data.Ipv6BridgeEnabled = types.BoolNull()
	}
	if value := res.Get("ledLightsOn"); value.Exists() && value.Value() != nil {
		data.LedLightsOn = types.BoolValue(value.Bool())
	} else {
		data.LedLightsOn = types.BoolNull()
	}
	if value := res.Get("locationAnalyticsEnabled"); value.Exists() && value.Value() != nil {
		data.LocationAnalyticsEnabled = types.BoolValue(value.Bool())
	} else {
		data.LocationAnalyticsEnabled = types.BoolNull()
	}
	if value := res.Get("meshingEnabled"); value.Exists() && value.Value() != nil {
		data.MeshingEnabled = types.BoolValue(value.Bool())
	} else {
		data.MeshingEnabled = types.BoolNull()
	}
	if value := res.Get("upgradeStrategy"); value.Exists() && value.Value() != nil {
		data.UpgradeStrategy = types.StringValue(value.String())
	} else {
		data.UpgradeStrategy = types.StringNull()
	}
	if value := res.Get("multicastToUnicastConversion.enabled"); value.Exists() && value.Value() != nil {
		data.MulticastToUnicastConversionEnabled = types.BoolValue(value.Bool())
	} else {
		data.MulticastToUnicastConversionEnabled = types.BoolNull()
	}
	if value := res.Get("namedVlans.poolDhcpMonitoring.duration"); value.Exists() && value.Value() != nil {
		data.NamedVlansPoolDhcpMonitoringDuration = types.Int64Value(value.Int())
	} else {
		data.NamedVlansPoolDhcpMonitoringDuration = types.Int64Null()
	}
	if value := res.Get("namedVlans.poolDhcpMonitoring.enabled"); value.Exists() && value.Value() != nil {
		data.NamedVlansPoolDhcpMonitoringEnabled = types.BoolValue(value.Bool())
	} else {
		data.NamedVlansPoolDhcpMonitoringEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody
