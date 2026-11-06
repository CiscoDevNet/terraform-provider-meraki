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
	"slices"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkFirmwareUpgrades struct {
	Id                                                     types.String                                         `tfsdk:"id"`
	NetworkId                                              types.String                                         `tfsdk:"network_id"`
	Timezone                                               types.String                                         `tfsdk:"timezone"`
	ProductsApplianceParticipateInNextBetaRelease          types.Bool                                           `tfsdk:"products_appliance_participate_in_next_beta_release"`
	ProductsApplianceNextUpgradeTime                       types.String                                         `tfsdk:"products_appliance_next_upgrade_time"`
	ProductsApplianceNextUpgradeToVersionId                types.String                                         `tfsdk:"products_appliance_next_upgrade_to_version_id"`
	ProductsCameraParticipateInNextBetaRelease             types.Bool                                           `tfsdk:"products_camera_participate_in_next_beta_release"`
	ProductsCameraNextUpgradeTime                          types.String                                         `tfsdk:"products_camera_next_upgrade_time"`
	ProductsCameraNextUpgradeToVersionId                   types.String                                         `tfsdk:"products_camera_next_upgrade_to_version_id"`
	ProductsCampusGatewayParticipateInNextBetaRelease      types.Bool                                           `tfsdk:"products_campus_gateway_participate_in_next_beta_release"`
	ProductsCampusGatewayNextUpgradeStrategy               types.String                                         `tfsdk:"products_campus_gateway_next_upgrade_strategy"`
	ProductsCampusGatewayNextUpgradeTime                   types.String                                         `tfsdk:"products_campus_gateway_next_upgrade_time"`
	ProductsCampusGatewayNextUpgradeToVersionId            types.String                                         `tfsdk:"products_campus_gateway_next_upgrade_to_version_id"`
	ProductsCellularGatewayParticipateInNextBetaRelease    types.Bool                                           `tfsdk:"products_cellular_gateway_participate_in_next_beta_release"`
	ProductsCellularGatewayNextUpgradeTime                 types.String                                         `tfsdk:"products_cellular_gateway_next_upgrade_time"`
	ProductsCellularGatewayNextUpgradeToVersionId          types.String                                         `tfsdk:"products_cellular_gateway_next_upgrade_to_version_id"`
	ProductsSecureConnectParticipateInNextBetaRelease      types.Bool                                           `tfsdk:"products_secure_connect_participate_in_next_beta_release"`
	ProductsSecureConnectNextUpgradeTime                   types.String                                         `tfsdk:"products_secure_connect_next_upgrade_time"`
	ProductsSecureConnectNextUpgradeToVersionId            types.String                                         `tfsdk:"products_secure_connect_next_upgrade_to_version_id"`
	ProductsSensorParticipateInNextBetaRelease             types.Bool                                           `tfsdk:"products_sensor_participate_in_next_beta_release"`
	ProductsSensorNextUpgradeTime                          types.String                                         `tfsdk:"products_sensor_next_upgrade_time"`
	ProductsSensorNextUpgradeToVersionId                   types.String                                         `tfsdk:"products_sensor_next_upgrade_to_version_id"`
	ProductsSwitchParticipateInNextBetaRelease             types.Bool                                           `tfsdk:"products_switch_participate_in_next_beta_release"`
	ProductsSwitchNextUpgradeTime                          types.String                                         `tfsdk:"products_switch_next_upgrade_time"`
	ProductsSwitchNextUpgradeToVersionId                   types.String                                         `tfsdk:"products_switch_next_upgrade_to_version_id"`
	ProductsSwitchCatalystParticipateInNextBetaRelease     types.Bool                                           `tfsdk:"products_switch_catalyst_participate_in_next_beta_release"`
	ProductsSwitchCatalystNextUpgradeTime                  types.String                                         `tfsdk:"products_switch_catalyst_next_upgrade_time"`
	ProductsSwitchCatalystNextUpgradeToVersionId           types.String                                         `tfsdk:"products_switch_catalyst_next_upgrade_to_version_id"`
	ProductsWirelessParticipateInNextBetaRelease           types.Bool                                           `tfsdk:"products_wireless_participate_in_next_beta_release"`
	ProductsWirelessNextUpgradeStrategy                    types.String                                         `tfsdk:"products_wireless_next_upgrade_strategy"`
	ProductsWirelessNextUpgradeTime                        types.String                                         `tfsdk:"products_wireless_next_upgrade_time"`
	ProductsWirelessNextUpgradePredownloadEnabled          types.Bool                                           `tfsdk:"products_wireless_next_upgrade_predownload_enabled"`
	ProductsWirelessNextUpgradeToVersionId                 types.String                                         `tfsdk:"products_wireless_next_upgrade_to_version_id"`
	ProductsWirelessControllerParticipateInNextBetaRelease types.Bool                                           `tfsdk:"products_wireless_controller_participate_in_next_beta_release"`
	ProductsWirelessControllerNextUpgradeTime              types.String                                         `tfsdk:"products_wireless_controller_next_upgrade_time"`
	ProductsWirelessControllerNextUpgradeToVersionId       types.String                                         `tfsdk:"products_wireless_controller_next_upgrade_to_version_id"`
	UpgradeWindowDayOfWeek                                 types.String                                         `tfsdk:"upgrade_window_day_of_week"`
	UpgradeWindowHourOfDay                                 types.String                                         `tfsdk:"upgrade_window_hour_of_day"`
	FeatureLossAcknowledgements                            []NetworkFirmwareUpgradesFeatureLossAcknowledgements `tfsdk:"feature_loss_acknowledgements"`
}

type NetworkFirmwareUpgradesFeatureLossAcknowledgements struct {
	NetworkId                types.String `tfsdk:"network_id"`
	ProductType              types.String `tfsdk:"product_type"`
	RemovedFeatureCheckNames types.List   `tfsdk:"removed_feature_check_names"`
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

func (data NetworkFirmwareUpgrades) toBody(ctx context.Context, state NetworkFirmwareUpgrades) string {
	body := ""
	if !data.Timezone.IsNull() {
		body, _ = sjson.Set(body, "timezone", data.Timezone.ValueString())
	}
	if !data.ProductsApplianceParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.appliance.participateInNextBetaRelease", data.ProductsApplianceParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsApplianceNextUpgradeTime.IsNull() && (data.ProductsApplianceNextUpgradeTime != state.ProductsApplianceNextUpgradeTime || data.ProductsApplianceNextUpgradeToVersionId != state.ProductsApplianceNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.appliance.nextUpgrade.time", data.ProductsApplianceNextUpgradeTime.ValueString())
	}
	if !data.ProductsApplianceNextUpgradeToVersionId.IsNull() && (data.ProductsApplianceNextUpgradeTime != state.ProductsApplianceNextUpgradeTime || data.ProductsApplianceNextUpgradeToVersionId != state.ProductsApplianceNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.appliance.nextUpgrade.toVersion.id", data.ProductsApplianceNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsCameraParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.camera.participateInNextBetaRelease", data.ProductsCameraParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsCameraNextUpgradeTime.IsNull() && (data.ProductsCameraNextUpgradeTime != state.ProductsCameraNextUpgradeTime || data.ProductsCameraNextUpgradeToVersionId != state.ProductsCameraNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.camera.nextUpgrade.time", data.ProductsCameraNextUpgradeTime.ValueString())
	}
	if !data.ProductsCameraNextUpgradeToVersionId.IsNull() && (data.ProductsCameraNextUpgradeTime != state.ProductsCameraNextUpgradeTime || data.ProductsCameraNextUpgradeToVersionId != state.ProductsCameraNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.camera.nextUpgrade.toVersion.id", data.ProductsCameraNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsCellularGatewayParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.cellularGateway.participateInNextBetaRelease", data.ProductsCellularGatewayParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsCellularGatewayNextUpgradeTime.IsNull() && (data.ProductsCellularGatewayNextUpgradeTime != state.ProductsCellularGatewayNextUpgradeTime || data.ProductsCellularGatewayNextUpgradeToVersionId != state.ProductsCellularGatewayNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.cellularGateway.nextUpgrade.time", data.ProductsCellularGatewayNextUpgradeTime.ValueString())
	}
	if !data.ProductsCellularGatewayNextUpgradeToVersionId.IsNull() && (data.ProductsCellularGatewayNextUpgradeTime != state.ProductsCellularGatewayNextUpgradeTime || data.ProductsCellularGatewayNextUpgradeToVersionId != state.ProductsCellularGatewayNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.cellularGateway.nextUpgrade.toVersion.id", data.ProductsCellularGatewayNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsSecureConnectParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.secureConnect.participateInNextBetaRelease", data.ProductsSecureConnectParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsSecureConnectNextUpgradeTime.IsNull() && (data.ProductsSecureConnectNextUpgradeTime != state.ProductsSecureConnectNextUpgradeTime || data.ProductsSecureConnectNextUpgradeToVersionId != state.ProductsSecureConnectNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.secureConnect.nextUpgrade.time", data.ProductsSecureConnectNextUpgradeTime.ValueString())
	}
	if !data.ProductsSecureConnectNextUpgradeToVersionId.IsNull() && (data.ProductsSecureConnectNextUpgradeTime != state.ProductsSecureConnectNextUpgradeTime || data.ProductsSecureConnectNextUpgradeToVersionId != state.ProductsSecureConnectNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.secureConnect.nextUpgrade.toVersion.id", data.ProductsSecureConnectNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsSensorParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.sensor.participateInNextBetaRelease", data.ProductsSensorParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsSensorNextUpgradeTime.IsNull() && (data.ProductsSensorNextUpgradeTime != state.ProductsSensorNextUpgradeTime || data.ProductsSensorNextUpgradeToVersionId != state.ProductsSensorNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.sensor.nextUpgrade.time", data.ProductsSensorNextUpgradeTime.ValueString())
	}
	if !data.ProductsSensorNextUpgradeToVersionId.IsNull() && (data.ProductsSensorNextUpgradeTime != state.ProductsSensorNextUpgradeTime || data.ProductsSensorNextUpgradeToVersionId != state.ProductsSensorNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.sensor.nextUpgrade.toVersion.id", data.ProductsSensorNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsSwitchParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.switch.participateInNextBetaRelease", data.ProductsSwitchParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsSwitchNextUpgradeTime.IsNull() && (data.ProductsSwitchNextUpgradeTime != state.ProductsSwitchNextUpgradeTime || data.ProductsSwitchNextUpgradeToVersionId != state.ProductsSwitchNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.switch.nextUpgrade.time", data.ProductsSwitchNextUpgradeTime.ValueString())
	}
	if !data.ProductsSwitchNextUpgradeToVersionId.IsNull() && (data.ProductsSwitchNextUpgradeTime != state.ProductsSwitchNextUpgradeTime || data.ProductsSwitchNextUpgradeToVersionId != state.ProductsSwitchNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.switch.nextUpgrade.toVersion.id", data.ProductsSwitchNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsSwitchCatalystParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.switchCatalyst.participateInNextBetaRelease", data.ProductsSwitchCatalystParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsSwitchCatalystNextUpgradeTime.IsNull() && (data.ProductsSwitchCatalystNextUpgradeTime != state.ProductsSwitchCatalystNextUpgradeTime || data.ProductsSwitchCatalystNextUpgradeToVersionId != state.ProductsSwitchCatalystNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.switchCatalyst.nextUpgrade.time", data.ProductsSwitchCatalystNextUpgradeTime.ValueString())
	}
	if !data.ProductsSwitchCatalystNextUpgradeToVersionId.IsNull() && (data.ProductsSwitchCatalystNextUpgradeTime != state.ProductsSwitchCatalystNextUpgradeTime || data.ProductsSwitchCatalystNextUpgradeToVersionId != state.ProductsSwitchCatalystNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.switchCatalyst.nextUpgrade.toVersion.id", data.ProductsSwitchCatalystNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsWirelessParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.wireless.participateInNextBetaRelease", data.ProductsWirelessParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsWirelessNextUpgradeTime.IsNull() && (data.ProductsWirelessNextUpgradeTime != state.ProductsWirelessNextUpgradeTime || data.ProductsWirelessNextUpgradePredownloadEnabled != state.ProductsWirelessNextUpgradePredownloadEnabled || data.ProductsWirelessNextUpgradeToVersionId != state.ProductsWirelessNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.wireless.nextUpgrade.time", data.ProductsWirelessNextUpgradeTime.ValueString())
	}
	if !data.ProductsWirelessNextUpgradePredownloadEnabled.IsNull() && (data.ProductsWirelessNextUpgradeTime != state.ProductsWirelessNextUpgradeTime || data.ProductsWirelessNextUpgradePredownloadEnabled != state.ProductsWirelessNextUpgradePredownloadEnabled || data.ProductsWirelessNextUpgradeToVersionId != state.ProductsWirelessNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.wireless.nextUpgrade.predownload.enabled", data.ProductsWirelessNextUpgradePredownloadEnabled.ValueBool())
	}
	if !data.ProductsWirelessNextUpgradeToVersionId.IsNull() && (data.ProductsWirelessNextUpgradeTime != state.ProductsWirelessNextUpgradeTime || data.ProductsWirelessNextUpgradePredownloadEnabled != state.ProductsWirelessNextUpgradePredownloadEnabled || data.ProductsWirelessNextUpgradeToVersionId != state.ProductsWirelessNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.wireless.nextUpgrade.toVersion.id", data.ProductsWirelessNextUpgradeToVersionId.ValueString())
	}
	if !data.ProductsWirelessControllerParticipateInNextBetaRelease.IsNull() {
		body, _ = sjson.Set(body, "products.wirelessController.participateInNextBetaRelease", data.ProductsWirelessControllerParticipateInNextBetaRelease.ValueBool())
	}
	if !data.ProductsWirelessControllerNextUpgradeTime.IsNull() && (data.ProductsWirelessControllerNextUpgradeTime != state.ProductsWirelessControllerNextUpgradeTime || data.ProductsWirelessControllerNextUpgradeToVersionId != state.ProductsWirelessControllerNextUpgradeToVersionId) {
		body, _ = sjson.Set(body, "products.wirelessController.nextUpgrade.time", data.ProductsWirelessControllerNextUpgradeTime.ValueString())
	}
	if !data.ProductsWirelessControllerNextUpgradeToVersionId.IsNull() && (data.ProductsWirelessControllerNextUpgradeTime != state.ProductsWirelessControllerNextUpgradeTime || data.ProductsWirelessControllerNextUpgradeToVersionId != state.ProductsWirelessControllerNextUpgradeToVersionId) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPreservingNulls

// toBodyPreservingNulls walks the same writable-attribute schema as toBody but
// reads directly from the raw API response (gjson) instead of from the
// Terraform model. Unlike toBody, it preserves attributes that the API
// explicitly returned as `null` (emitting them as JSON `null` rather than
// dropping them). This is used by the singleton restoreOriginalStateOnDestroy
// path so that explicit-null fields captured during Create are restored on
// Delete. Keep this method in sync with toBody — both walk the same
// `.Attributes` schema and must agree on which fields are writable.
func (data NetworkFirmwareUpgrades) toBodyPreservingNulls(ctx context.Context, res meraki.Res) string {
	body := ""
	if value := res.Get("timezone"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "timezone", "null")
		} else {
			body, _ = sjson.Set(body, "timezone", value.String())
		}
	}
	if value := res.Get("products.appliance.participateInNextBetaRelease"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.appliance.participateInNextBetaRelease", "null")
		} else {
			body, _ = sjson.Set(body, "products.appliance.participateInNextBetaRelease", value.Bool())
		}
	}
	if value := res.Get("products.camera.participateInNextBetaRelease"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.camera.participateInNextBetaRelease", "null")
		} else {
			body, _ = sjson.Set(body, "products.camera.participateInNextBetaRelease", value.Bool())
		}
	}
	if value := res.Get("products.campusGateway.participateInNextBetaRelease"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.campusGateway.participateInNextBetaRelease", "null")
		} else {
			body, _ = sjson.Set(body, "products.campusGateway.participateInNextBetaRelease", value.Bool())
		}
	}
	if value := res.Get("products.campusGateway.nextUpgrade.strategy"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.campusGateway.nextUpgrade.strategy", "null")
		} else {
			body, _ = sjson.Set(body, "products.campusGateway.nextUpgrade.strategy", value.String())
		}
	}
	if value := res.Get("products.campusGateway.nextUpgrade.time"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.campusGateway.nextUpgrade.time", "null")
		} else {
			body, _ = sjson.Set(body, "products.campusGateway.nextUpgrade.time", value.String())
		}
	}
	if value := res.Get("products.campusGateway.nextUpgrade.toVersion.id"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.campusGateway.nextUpgrade.toVersion.id", "null")
		} else {
			body, _ = sjson.Set(body, "products.campusGateway.nextUpgrade.toVersion.id", value.String())
		}
	}
	if value := res.Get("products.cellularGateway.participateInNextBetaRelease"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.cellularGateway.participateInNextBetaRelease", "null")
		} else {
			body, _ = sjson.Set(body, "products.cellularGateway.participateInNextBetaRelease", value.Bool())
		}
	}
	if value := res.Get("products.secureConnect.participateInNextBetaRelease"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.secureConnect.participateInNextBetaRelease", "null")
		} else {
			body, _ = sjson.Set(body, "products.secureConnect.participateInNextBetaRelease", value.Bool())
		}
	}
	if value := res.Get("products.sensor.participateInNextBetaRelease"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.sensor.participateInNextBetaRelease", "null")
		} else {
			body, _ = sjson.Set(body, "products.sensor.participateInNextBetaRelease", value.Bool())
		}
	}
	if value := res.Get("products.switch.participateInNextBetaRelease"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.switch.participateInNextBetaRelease", "null")
		} else {
			body, _ = sjson.Set(body, "products.switch.participateInNextBetaRelease", value.Bool())
		}
	}
	if value := res.Get("products.switchCatalyst.participateInNextBetaRelease"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.switchCatalyst.participateInNextBetaRelease", "null")
		} else {
			body, _ = sjson.Set(body, "products.switchCatalyst.participateInNextBetaRelease", value.Bool())
		}
	}
	if value := res.Get("products.wireless.participateInNextBetaRelease"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.wireless.participateInNextBetaRelease", "null")
		} else {
			body, _ = sjson.Set(body, "products.wireless.participateInNextBetaRelease", value.Bool())
		}
	}
	if value := res.Get("products.wireless.nextUpgrade.strategy"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.wireless.nextUpgrade.strategy", "null")
		} else {
			body, _ = sjson.Set(body, "products.wireless.nextUpgrade.strategy", value.String())
		}
	}
	if value := res.Get("products.wirelessController.participateInNextBetaRelease"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "products.wirelessController.participateInNextBetaRelease", "null")
		} else {
			body, _ = sjson.Set(body, "products.wirelessController.participateInNextBetaRelease", value.Bool())
		}
	}
	if value := res.Get("upgradeWindow.dayOfWeek"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "upgradeWindow.dayOfWeek", "null")
		} else {
			body, _ = sjson.Set(body, "upgradeWindow.dayOfWeek", value.String())
		}
	}
	if value := res.Get("upgradeWindow.hourOfDay"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "upgradeWindow.hourOfDay", "null")
		} else {
			body, _ = sjson.Set(body, "upgradeWindow.hourOfDay", value.String())
		}
	}
	if value := res.Get("featureLossAcknowledgements"); value.Exists() {
		if value.Value() == nil {
			body, _ = sjson.SetRaw(body, "featureLossAcknowledgements", "null")
		} else {
			body, _ = sjson.Set(body, "featureLossAcknowledgements", []interface{}{})
			parent := &body
			value.ForEach(func(k, res gjson.Result) bool {
				body := ""
				if value := res.Get("networkId"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "networkId", "null")
					} else {
						body, _ = sjson.Set(body, "networkId", value.String())
					}
				}
				if value := res.Get("productType"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "productType", "null")
					} else {
						body, _ = sjson.Set(body, "productType", value.String())
					}
				}
				if value := res.Get("removedFeatureCheckNames"); value.Exists() {
					if value.Value() == nil {
						body, _ = sjson.SetRaw(body, "removedFeatureCheckNames", "null")
					} else {
						var values []string
						for _, v := range value.Array() {
							values = append(values, v.String())
						}
						body, _ = sjson.Set(body, "removedFeatureCheckNames", values)
					}
				}
				*parent, _ = sjson.SetRaw(*parent, "featureLossAcknowledgements.-1", body)
				return true
			})
		}
	}
	return body
}

// End of section. //template:end toBodyPreservingNulls

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
	if value := res.Get("products.camera.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.campusGateway.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsCampusGatewayParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCampusGatewayParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.campusGateway.nextUpgrade.strategy"); value.Exists() && value.Value() != nil {
		data.ProductsCampusGatewayNextUpgradeStrategy = types.StringValue(value.String())
	} else {
		data.ProductsCampusGatewayNextUpgradeStrategy = types.StringNull()
	}
	if value := res.Get("products.campusGateway.nextUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsCampusGatewayNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCampusGatewayNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.campusGateway.nextUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsCampusGatewayNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCampusGatewayNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.secureConnect.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.sensor.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switch.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switchCatalyst.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wireless.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.strategy"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessNextUpgradeStrategy = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeStrategy = types.StringNull()
	}
	if value := res.Get("products.wirelessController.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolNull()
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
	if value := res.Get("featureLossAcknowledgements"); value.Exists() && value.Value() != nil {
		data.FeatureLossAcknowledgements = make([]NetworkFirmwareUpgradesFeatureLossAcknowledgements, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesFeatureLossAcknowledgements{}
			if value := res.Get("networkId"); value.Exists() && value.Value() != nil {
				data.NetworkId = types.StringValue(value.String())
			} else {
				data.NetworkId = types.StringNull()
			}
			if value := res.Get("productType"); value.Exists() && value.Value() != nil {
				data.ProductType = types.StringValue(value.String())
			} else {
				data.ProductType = types.StringNull()
			}
			if value := res.Get("removedFeatureCheckNames"); value.Exists() && value.Value() != nil {
				data.RemovedFeatureCheckNames = helpers.GetStringList(value.Array())
			} else {
				data.RemovedFeatureCheckNames = types.ListNull(types.StringType)
			}
			(*parent).FeatureLossAcknowledgements = append((*parent).FeatureLossAcknowledgements, data)
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
	if value := res.Get("products.camera.participateInNextBetaRelease"); value.Exists() && !data.ProductsCameraParticipateInNextBetaRelease.IsNull() {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.campusGateway.participateInNextBetaRelease"); value.Exists() && !data.ProductsCampusGatewayParticipateInNextBetaRelease.IsNull() {
		data.ProductsCampusGatewayParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCampusGatewayParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.campusGateway.nextUpgrade.strategy"); value.Exists() && !data.ProductsCampusGatewayNextUpgradeStrategy.IsNull() {
		data.ProductsCampusGatewayNextUpgradeStrategy = types.StringValue(value.String())
	} else {
		data.ProductsCampusGatewayNextUpgradeStrategy = types.StringNull()
	}
	if value := res.Get("products.campusGateway.nextUpgrade.time"); value.Exists() && !data.ProductsCampusGatewayNextUpgradeTime.IsNull() {
		data.ProductsCampusGatewayNextUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCampusGatewayNextUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.campusGateway.nextUpgrade.toVersion.id"); value.Exists() && !data.ProductsCampusGatewayNextUpgradeToVersionId.IsNull() {
		data.ProductsCampusGatewayNextUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCampusGatewayNextUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.participateInNextBetaRelease"); value.Exists() && !data.ProductsCellularGatewayParticipateInNextBetaRelease.IsNull() {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.secureConnect.participateInNextBetaRelease"); value.Exists() && !data.ProductsSecureConnectParticipateInNextBetaRelease.IsNull() {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.sensor.participateInNextBetaRelease"); value.Exists() && !data.ProductsSensorParticipateInNextBetaRelease.IsNull() {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switch.participateInNextBetaRelease"); value.Exists() && !data.ProductsSwitchParticipateInNextBetaRelease.IsNull() {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switchCatalyst.participateInNextBetaRelease"); value.Exists() && !data.ProductsSwitchCatalystParticipateInNextBetaRelease.IsNull() {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wireless.participateInNextBetaRelease"); value.Exists() && !data.ProductsWirelessParticipateInNextBetaRelease.IsNull() {
		data.ProductsWirelessParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.strategy"); value.Exists() && !data.ProductsWirelessNextUpgradeStrategy.IsNull() {
		data.ProductsWirelessNextUpgradeStrategy = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeStrategy = types.StringNull()
	}
	if value := res.Get("products.wirelessController.participateInNextBetaRelease"); value.Exists() && !data.ProductsWirelessControllerParticipateInNextBetaRelease.IsNull() {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolNull()
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
	for i := 0; i < len(data.FeatureLossAcknowledgements); i++ {
		keys := [...]string{"networkId", "productType"}
		keyValues := [...]string{data.FeatureLossAcknowledgements[i].NetworkId.ValueString(), data.FeatureLossAcknowledgements[i].ProductType.ValueString()}

		parent := &data
		data := (*parent).FeatureLossAcknowledgements[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("featureLossAcknowledgements").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing FeatureLossAcknowledgements[%d] = %+v",
				i,
				(*parent).FeatureLossAcknowledgements[i],
			))
			(*parent).FeatureLossAcknowledgements = slices.Delete((*parent).FeatureLossAcknowledgements, i, i+1)
			i--

			continue
		}
		if value := res.Get("networkId"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("productType"); value.Exists() && !data.ProductType.IsNull() {
			data.ProductType = types.StringValue(value.String())
		} else {
			data.ProductType = types.StringNull()
		}
		if value := res.Get("removedFeatureCheckNames"); value.Exists() && !data.RemovedFeatureCheckNames.IsNull() {
			data.RemovedFeatureCheckNames = helpers.GetStringList(value.Array())
		} else {
			data.RemovedFeatureCheckNames = types.ListNull(types.StringType)
		}
		(*parent).FeatureLossAcknowledgements[i] = data
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
