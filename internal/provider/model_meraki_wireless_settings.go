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

type WirelessSettings struct {
	Id                                   types.String `tfsdk:"id"`
	NetworkId                            types.String `tfsdk:"network_id"`
	Ipv6BridgeEnabled                    types.Bool   `tfsdk:"ipv6_bridge_enabled"`
	LedLightsOn                          types.Bool   `tfsdk:"led_lights_on"`
	LocationAnalyticsEnabled             types.Bool   `tfsdk:"location_analytics_enabled"`
	MeshingEnabled                       types.Bool   `tfsdk:"meshing_enabled"`
	UpgradeStrategy                      types.String `tfsdk:"upgrade_strategy"`
	NamedVlansPoolDhcpMonitoringDuration types.Int64  `tfsdk:"named_vlans_pool_dhcp_monitoring_duration"`
	NamedVlansPoolDhcpMonitoringEnabled  types.Bool   `tfsdk:"named_vlans_pool_dhcp_monitoring_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data WirelessSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/settings", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSettings) toBody(ctx context.Context, state WirelessSettings) string {
	body := ""
	if !data.Ipv6BridgeEnabled.IsNull() {
		body, _ = sjson.Set(body, "ipv6BridgeEnabled", data.Ipv6BridgeEnabled.ValueBool())
	}
	if !data.LedLightsOn.IsNull() {
		body, _ = sjson.Set(body, "ledLightsOn", data.LedLightsOn.ValueBool())
	}
	if !data.LocationAnalyticsEnabled.IsNull() {
		body, _ = sjson.Set(body, "locationAnalyticsEnabled", data.LocationAnalyticsEnabled.ValueBool())
	}
	if !data.MeshingEnabled.IsNull() {
		body, _ = sjson.Set(body, "meshingEnabled", data.MeshingEnabled.ValueBool())
	}
	if !data.UpgradeStrategy.IsNull() {
		body, _ = sjson.Set(body, "upgradeStrategy", data.UpgradeStrategy.ValueString())
	}
	if !data.NamedVlansPoolDhcpMonitoringDuration.IsNull() {
		body, _ = sjson.Set(body, "namedVlans.poolDhcpMonitoring.duration", data.NamedVlansPoolDhcpMonitoringDuration.ValueInt64())
	}
	if !data.NamedVlansPoolDhcpMonitoringEnabled.IsNull() {
		body, _ = sjson.Set(body, "namedVlans.poolDhcpMonitoring.enabled", data.NamedVlansPoolDhcpMonitoringEnabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSettings) fromBody(ctx context.Context, res meraki.Res) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessSettings) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("ipv6BridgeEnabled"); value.Exists() && !data.Ipv6BridgeEnabled.IsNull() {
		data.Ipv6BridgeEnabled = types.BoolValue(value.Bool())
	} else {
		data.Ipv6BridgeEnabled = types.BoolNull()
	}
	if value := res.Get("ledLightsOn"); value.Exists() && !data.LedLightsOn.IsNull() {
		data.LedLightsOn = types.BoolValue(value.Bool())
	} else {
		data.LedLightsOn = types.BoolNull()
	}
	if value := res.Get("locationAnalyticsEnabled"); value.Exists() && !data.LocationAnalyticsEnabled.IsNull() {
		data.LocationAnalyticsEnabled = types.BoolValue(value.Bool())
	} else {
		data.LocationAnalyticsEnabled = types.BoolNull()
	}
	if value := res.Get("meshingEnabled"); value.Exists() && !data.MeshingEnabled.IsNull() {
		data.MeshingEnabled = types.BoolValue(value.Bool())
	} else {
		data.MeshingEnabled = types.BoolNull()
	}
	if value := res.Get("upgradeStrategy"); value.Exists() && !data.UpgradeStrategy.IsNull() {
		data.UpgradeStrategy = types.StringValue(value.String())
	} else {
		data.UpgradeStrategy = types.StringNull()
	}
	if value := res.Get("namedVlans.poolDhcpMonitoring.duration"); value.Exists() && !data.NamedVlansPoolDhcpMonitoringDuration.IsNull() {
		data.NamedVlansPoolDhcpMonitoringDuration = types.Int64Value(value.Int())
	} else {
		data.NamedVlansPoolDhcpMonitoringDuration = types.Int64Null()
	}
	if value := res.Get("namedVlans.poolDhcpMonitoring.enabled"); value.Exists() && !data.NamedVlansPoolDhcpMonitoringEnabled.IsNull() {
		data.NamedVlansPoolDhcpMonitoringEnabled = types.BoolValue(value.Bool())
	} else {
		data.NamedVlansPoolDhcpMonitoringEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessSettings) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data WirelessSettings) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
