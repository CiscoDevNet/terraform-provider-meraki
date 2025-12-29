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

type NetworkFirmwareUpgrades struct {
	Id                                                     types.String `tfsdk:"id"`
	NetworkId                                              types.String `tfsdk:"network_id"`
	Timezone                                               types.String `tfsdk:"timezone"`
	ProductsApplianceParticipateInNextBetaRelease          types.Bool   `tfsdk:"products_appliance_participate_in_next_beta_release"`
	ProductsApplianceNextUpgradeTime                       types.String `tfsdk:"products_appliance_next_upgrade_time"`
	ProductsApplianceNextUpgradeToVersionId                types.String `tfsdk:"products_appliance_next_upgrade_to_version_id"`
	ProductsCameraParticipateInNextBetaRelease             types.Bool   `tfsdk:"products_camera_participate_in_next_beta_release"`
	ProductsCameraNextUpgradeTime                          types.String `tfsdk:"products_camera_next_upgrade_time"`
	ProductsCameraNextUpgradeToVersionId                   types.String `tfsdk:"products_camera_next_upgrade_to_version_id"`
	ProductsCellularGatewayParticipateInNextBetaRelease    types.Bool   `tfsdk:"products_cellular_gateway_participate_in_next_beta_release"`
	ProductsCellularGatewayNextUpgradeTime                 types.String `tfsdk:"products_cellular_gateway_next_upgrade_time"`
	ProductsCellularGatewayNextUpgradeToVersionId          types.String `tfsdk:"products_cellular_gateway_next_upgrade_to_version_id"`
	ProductsSecureConnectParticipateInNextBetaRelease      types.Bool   `tfsdk:"products_secure_connect_participate_in_next_beta_release"`
	ProductsSecureConnectNextUpgradeTime                   types.String `tfsdk:"products_secure_connect_next_upgrade_time"`
	ProductsSecureConnectNextUpgradeToVersionId            types.String `tfsdk:"products_secure_connect_next_upgrade_to_version_id"`
	ProductsSensorParticipateInNextBetaRelease             types.Bool   `tfsdk:"products_sensor_participate_in_next_beta_release"`
	ProductsSensorNextUpgradeTime                          types.String `tfsdk:"products_sensor_next_upgrade_time"`
	ProductsSensorNextUpgradeToVersionId                   types.String `tfsdk:"products_sensor_next_upgrade_to_version_id"`
	ProductsSwitchParticipateInNextBetaRelease             types.Bool   `tfsdk:"products_switch_participate_in_next_beta_release"`
	ProductsSwitchNextUpgradeTime                          types.String `tfsdk:"products_switch_next_upgrade_time"`
	ProductsSwitchNextUpgradeToVersionId                   types.String `tfsdk:"products_switch_next_upgrade_to_version_id"`
	ProductsSwitchCatalystParticipateInNextBetaRelease     types.Bool   `tfsdk:"products_switch_catalyst_participate_in_next_beta_release"`
	ProductsSwitchCatalystNextUpgradeTime                  types.String `tfsdk:"products_switch_catalyst_next_upgrade_time"`
	ProductsSwitchCatalystNextUpgradeToVersionId           types.String `tfsdk:"products_switch_catalyst_next_upgrade_to_version_id"`
	ProductsWirelessParticipateInNextBetaRelease           types.Bool   `tfsdk:"products_wireless_participate_in_next_beta_release"`
	ProductsWirelessNextUpgradeTime                        types.String `tfsdk:"products_wireless_next_upgrade_time"`
	ProductsWirelessNextUpgradeToVersionId                 types.String `tfsdk:"products_wireless_next_upgrade_to_version_id"`
	ProductsWirelessControllerParticipateInNextBetaRelease types.Bool   `tfsdk:"products_wireless_controller_participate_in_next_beta_release"`
	ProductsWirelessControllerNextUpgradeTime              types.String `tfsdk:"products_wireless_controller_next_upgrade_time"`
	ProductsWirelessControllerNextUpgradeToVersionId       types.String `tfsdk:"products_wireless_controller_next_upgrade_to_version_id"`
	UpgradeWindowDayOfWeek                                 types.String `tfsdk:"upgrade_window_day_of_week"`
	UpgradeWindowHourOfDay                                 types.String `tfsdk:"upgrade_window_hour_of_day"`
}

type NetworkFirmwareUpgradesIdentity struct {
	NetworkId types.String `tfsdk:"network_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data NetworkFirmwareUpgrades) getPath() string {
	return fmt.Sprintf("/networks/%v/firmwareUpgrades", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data NetworkFirmwareUpgrades) toBody(ctx context.Context, state NetworkFirmwareUpgrades) string {
	body := ""
	if !data.Timezone.IsNull() {
		body, _ = sjson.Set(body, "timezone", data.Timezone.ValueString())
	}
	if !data.ProductsApplianceParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.appliance.participateInNextBetaRelease", data.ProductsApplianceParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsApplianceNextUpgradeTime.IsNull() {
		body, _ = sjson.Set(body, "products.appliance.nextUpgrade.time", data.ProductsApplianceNextUpgradeTime.ValueString())
	}
	if !data.ProductsApplianceNextUpgradeToVersionId.IsNull() {
		body, _ = sjson.Set(body, "products.appliance.nextUpgrade.toVersion.id", data.ProductsApplianceNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsCameraParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.camera.participateInNextBetaRelease", data.ProductsCameraParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsCameraNextUpgradeTime.IsNull() {
		body, _ = sjson.Set(body, "products.camera.nextUpgrade.time", data.ProductsCameraNextUpgradeTime.ValueString())
	}
	if !data.ProductsCameraNextUpgradeToVersionId.IsNull() {
		body, _ = sjson.Set(body, "products.camera.nextUpgrade.toVersion.id", data.ProductsCameraNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsCellularGatewayParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.cellularGateway.participateInNextBetaRelease", data.ProductsCellularGatewayParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsCellularGatewayNextUpgradeTime.IsNull() {
		body, _ = sjson.Set(body, "products.cellularGateway.nextUpgrade.time", data.ProductsCellularGatewayNextUpgradeTime.ValueString())
	}
	if !data.ProductsCellularGatewayNextUpgradeToVersionId.IsNull() {
		body, _ = sjson.Set(body, "products.cellularGateway.nextUpgrade.toVersion.id", data.ProductsCellularGatewayNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsSecureConnectParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.secureConnect.participateInNextBetaRelease", data.ProductsSecureConnectParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsSecureConnectNextUpgradeTime.IsNull() {
		body, _ = sjson.Set(body, "products.secureConnect.nextUpgrade.time", data.ProductsSecureConnectNextUpgradeTime.ValueString())
	}
	if !data.ProductsSecureConnectNextUpgradeToVersionId.IsNull() {
		body, _ = sjson.Set(body, "products.secureConnect.nextUpgrade.toVersion.id", data.ProductsSecureConnectNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsSensorParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.sensor.participateInNextBetaRelease", data.ProductsSensorParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsSensorNextUpgradeTime.IsNull() {
		body, _ = sjson.Set(body, "products.sensor.nextUpgrade.time", data.ProductsSensorNextUpgradeTime.ValueString())
	}
	if !data.ProductsSensorNextUpgradeToVersionId.IsNull() {
		body, _ = sjson.Set(body, "products.sensor.nextUpgrade.toVersion.id", data.ProductsSensorNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsSwitchParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.switch.participateInNextBetaRelease", data.ProductsSwitchParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsSwitchNextUpgradeTime.IsNull() {
		body, _ = sjson.Set(body, "products.switch.nextUpgrade.time", data.ProductsSwitchNextUpgradeTime.ValueString())
	}
	if !data.ProductsSwitchNextUpgradeToVersionId.IsNull() {
		body, _ = sjson.Set(body, "products.switch.nextUpgrade.toVersion.id", data.ProductsSwitchNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsSwitchCatalystParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.switchCatalyst.participateInNextBetaRelease", data.ProductsSwitchCatalystParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsSwitchCatalystNextUpgradeTime.IsNull() {
		body, _ = sjson.Set(body, "products.switchCatalyst.nextUpgrade.time", data.ProductsSwitchCatalystNextUpgradeTime.ValueString())
	}
	if !data.ProductsSwitchCatalystNextUpgradeToVersionId.IsNull() {
		body, _ = sjson.Set(body, "products.switchCatalyst.nextUpgrade.toVersion.id", data.ProductsSwitchCatalystNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsWirelessParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.wireless.participateInNextBetaRelease", data.ProductsWirelessParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsWirelessNextUpgradeTime.IsNull() {
		body, _ = sjson.Set(body, "products.wireless.nextUpgrade.time", data.ProductsWirelessNextUpgradeTime.ValueString())
	}
	if !data.ProductsWirelessNextUpgradeToVersionId.IsNull() {
		body, _ = sjson.Set(body, "products.wireless.nextUpgrade.toVersion.id", data.ProductsWirelessNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsWirelessControllerParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.wirelessController.participateInNextBetaRelease", data.ProductsWirelessControllerParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsWirelessControllerNextUpgradeTime.IsNull() {
		body, _ = sjson.Set(body, "products.wirelessController.nextUpgrade.time", data.ProductsWirelessControllerNextUpgradeTime.ValueString())
	}
	if !data.ProductsWirelessControllerNextUpgradeToVersionId.IsNull() {
		body, _ = sjson.Set(body, "products.wirelessController.nextUpgrade.toVersion.id", data.ProductsWirelessControllerNextUpgradeToVersionId.ValueString())
	}
	if !data.UpgradeWindowDayOfWeek.IsNull() {
		body, _ = sjson.Set(body, "upgradeWindow.dayOfWeek", data.UpgradeWindowDayOfWeek.ValueString())
	}
	if !data.UpgradeWindowHourOfDay.IsNull() {
		body, _ = sjson.Set(body, "upgradeWindow.hourOfDay", data.UpgradeWindowHourOfDay.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *NetworkFirmwareUpgrades) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("timezone"); value.Exists() && value.Value() != nil {
		data.Timezone = types.StringValue(value.String())
	} else {
		data.Timezone = types.StringNull()
	}
	if value := res.Get("products.appliance.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsApplianceParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.camera.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.camera.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsCameraNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.camera.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsCameraNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.secureConnect.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.sensor.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSensorNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSensorNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.switch.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switch.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.switch.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.wireless.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.wirelessController.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("upgradeWindow.dayOfWeek"); value.Exists() && value.Value() != nil {
		data.UpgradeWindowDayOfWeek = types.StringValue(value.String())
	} else {
		data.UpgradeWindowDayOfWeek = types.StringNull()
	}
	if value := res.Get("upgradeWindow.hourOfDay"); value.Exists() && value.Value() != nil {
		data.UpgradeWindowHourOfDay = types.StringValue(value.String())
	} else {
		data.UpgradeWindowHourOfDay = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *NetworkFirmwareUpgrades) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("timezone"); value.Exists() && !data.Timezone.IsNull() {
		data.Timezone = types.StringValue(value.String())
	} else {
		data.Timezone = types.StringNull()
	}
	if value := res.Get("products.appliance.participateInNextBetaRelease"); value.Exists() && !data.ProductsApplianceParticipateInNextBetaRelease.IsNull() {
		data.ProductsApplianceParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsApplianceParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.time"); value.Exists() && !data.ProductsApplianceNextUpgradeTime.IsNull() {
		data.ProductsApplianceNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsApplianceNextUpgradeToVersionId.IsNull() {
		data.ProductsApplianceNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.camera.participateInNextBetaRelease"); value.Exists() && !data.ProductsCameraParticipateInNextBetaRelease.IsNull() {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.camera.nextUpgrade.time"); value.Exists() && !data.ProductsCameraNextUpgradeTime.IsNull() {
		data.ProductsCameraNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.camera.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsCameraNextUpgradeToVersionId.IsNull() {
		data.ProductsCameraNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.participateInNextBetaRelease"); value.Exists() && !data.ProductsCellularGatewayParticipateInNextBetaRelease.IsNull() {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.time"); value.Exists() && !data.ProductsCellularGatewayNextUpgradeTime.IsNull() {
		data.ProductsCellularGatewayNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsCellularGatewayNextUpgradeToVersionId.IsNull() {
		data.ProductsCellularGatewayNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.secureConnect.participateInNextBetaRelease"); value.Exists() && !data.ProductsSecureConnectParticipateInNextBetaRelease.IsNull() {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.time"); value.Exists() && !data.ProductsSecureConnectNextUpgradeTime.IsNull() {
		data.ProductsSecureConnectNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsSecureConnectNextUpgradeToVersionId.IsNull() {
		data.ProductsSecureConnectNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.sensor.participateInNextBetaRelease"); value.Exists() && !data.ProductsSensorParticipateInNextBetaRelease.IsNull() {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.time"); value.Exists() && !data.ProductsSensorNextUpgradeTime.IsNull() {
		data.ProductsSensorNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsSensorNextUpgradeToVersionId.IsNull() {
		data.ProductsSensorNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.switch.participateInNextBetaRelease"); value.Exists() && !data.ProductsSwitchParticipateInNextBetaRelease.IsNull() {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switch.nextUpgrade.time"); value.Exists() && !data.ProductsSwitchNextUpgradeTime.IsNull() {
		data.ProductsSwitchNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.switch.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsSwitchNextUpgradeToVersionId.IsNull() {
		data.ProductsSwitchNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.participateInNextBetaRelease"); value.Exists() && !data.ProductsSwitchCatalystParticipateInNextBetaRelease.IsNull() {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.time"); value.Exists() && !data.ProductsSwitchCatalystNextUpgradeTime.IsNull() {
		data.ProductsSwitchCatalystNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsSwitchCatalystNextUpgradeToVersionId.IsNull() {
		data.ProductsSwitchCatalystNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.wireless.participateInNextBetaRelease"); value.Exists() && !data.ProductsWirelessParticipateInNextBetaRelease.IsNull() {
		data.ProductsWirelessParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.time"); value.Exists() && !data.ProductsWirelessNextUpgradeTime.IsNull() {
		data.ProductsWirelessNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsWirelessNextUpgradeToVersionId.IsNull() {
		data.ProductsWirelessNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.wirelessController.participateInNextBetaRelease"); value.Exists() && !data.ProductsWirelessControllerParticipateInNextBetaRelease.IsNull() {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.time"); value.Exists() && !data.ProductsWirelessControllerNextUpgradeTime.IsNull() {
		data.ProductsWirelessControllerNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsWirelessControllerNextUpgradeToVersionId.IsNull() {
		data.ProductsWirelessControllerNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("upgradeWindow.dayOfWeek"); value.Exists() && !data.UpgradeWindowDayOfWeek.IsNull() {
		data.UpgradeWindowDayOfWeek = types.StringValue(value.String())
	} else {
		data.UpgradeWindowDayOfWeek = types.StringNull()
	}
	if value := res.Get("upgradeWindow.hourOfDay"); value.Exists() && !data.UpgradeWindowHourOfDay.IsNull() {
		data.UpgradeWindowHourOfDay = types.StringValue(value.String())
	} else {
		data.UpgradeWindowHourOfDay = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *NetworkFirmwareUpgrades) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *NetworkFirmwareUpgradesIdentity) toIdentity(ctx context.Context, plan *NetworkFirmwareUpgrades) {
	data.NetworkId = plan.NetworkId
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *NetworkFirmwareUpgrades) fromIdentity(ctx context.Context, identity *NetworkFirmwareUpgradesIdentity) {
	data.NetworkId = identity.NetworkId
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data NetworkFirmwareUpgrades) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
