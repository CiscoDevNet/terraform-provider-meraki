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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type NetworkFirmwareUpgrades struct {
	Id                                                          types.String                                                         `tfsdk:"id"`
	NetworkId                                                   types.String                                                         `tfsdk:"network_id"`
	Timezone                                                    types.String                                                         `tfsdk:"timezone"`
	ProductsApplianceIsUpgradeAvailable                         types.Bool                                                           `tfsdk:"products_appliance_is_upgrade_available"`
	ProductsApplianceParticipateInNextBetaRelease               types.Bool                                                           `tfsdk:"products_appliance_participate_in_next_beta_release"`
	ProductsApplianceCurrentVersionFirmware                     types.String                                                         `tfsdk:"products_appliance_current_version_firmware"`
	ProductsApplianceCurrentVersionId                           types.String                                                         `tfsdk:"products_appliance_current_version_id"`
	ProductsApplianceCurrentVersionReleaseDate                  types.String                                                         `tfsdk:"products_appliance_current_version_release_date"`
	ProductsApplianceCurrentVersionReleaseType                  types.String                                                         `tfsdk:"products_appliance_current_version_release_type"`
	ProductsApplianceCurrentVersionShortName                    types.String                                                         `tfsdk:"products_appliance_current_version_short_name"`
	ProductsApplianceLastUpgradeTime                            types.String                                                         `tfsdk:"products_appliance_last_upgrade_time"`
	ProductsApplianceLastUpgradeFromVersionFirmware             types.String                                                         `tfsdk:"products_appliance_last_upgrade_from_version_firmware"`
	ProductsApplianceLastUpgradeFromVersionId                   types.String                                                         `tfsdk:"products_appliance_last_upgrade_from_version_id"`
	ProductsApplianceLastUpgradeFromVersionReleaseDate          types.String                                                         `tfsdk:"products_appliance_last_upgrade_from_version_release_date"`
	ProductsApplianceLastUpgradeFromVersionReleaseType          types.String                                                         `tfsdk:"products_appliance_last_upgrade_from_version_release_type"`
	ProductsApplianceLastUpgradeFromVersionShortName            types.String                                                         `tfsdk:"products_appliance_last_upgrade_from_version_short_name"`
	ProductsApplianceLastUpgradeToVersionFirmware               types.String                                                         `tfsdk:"products_appliance_last_upgrade_to_version_firmware"`
	ProductsApplianceLastUpgradeToVersionId                     types.String                                                         `tfsdk:"products_appliance_last_upgrade_to_version_id"`
	ProductsApplianceLastUpgradeToVersionReleaseDate            types.String                                                         `tfsdk:"products_appliance_last_upgrade_to_version_release_date"`
	ProductsApplianceLastUpgradeToVersionReleaseType            types.String                                                         `tfsdk:"products_appliance_last_upgrade_to_version_release_type"`
	ProductsApplianceLastUpgradeToVersionShortName              types.String                                                         `tfsdk:"products_appliance_last_upgrade_to_version_short_name"`
	ProductsApplianceNextUpgradeTime                            types.String                                                         `tfsdk:"products_appliance_next_upgrade_time"`
	ProductsApplianceNextUpgradeToVersionFirmware               types.String                                                         `tfsdk:"products_appliance_next_upgrade_to_version_firmware"`
	ProductsApplianceNextUpgradeToVersionId                     types.String                                                         `tfsdk:"products_appliance_next_upgrade_to_version_id"`
	ProductsApplianceNextUpgradeToVersionReleaseDate            types.String                                                         `tfsdk:"products_appliance_next_upgrade_to_version_release_date"`
	ProductsApplianceNextUpgradeToVersionReleaseType            types.String                                                         `tfsdk:"products_appliance_next_upgrade_to_version_release_type"`
	ProductsApplianceNextUpgradeToVersionShortName              types.String                                                         `tfsdk:"products_appliance_next_upgrade_to_version_short_name"`
	ProductsApplianceAvailableVersions                          []NetworkFirmwareUpgradesProductsApplianceAvailableVersions          `tfsdk:"products_appliance_available_versions"`
	ProductsCameraIsUpgradeAvailable                            types.Bool                                                           `tfsdk:"products_camera_is_upgrade_available"`
	ProductsCameraParticipateInNextBetaRelease                  types.Bool                                                           `tfsdk:"products_camera_participate_in_next_beta_release"`
	ProductsCameraCurrentVersionFirmware                        types.String                                                         `tfsdk:"products_camera_current_version_firmware"`
	ProductsCameraCurrentVersionId                              types.String                                                         `tfsdk:"products_camera_current_version_id"`
	ProductsCameraCurrentVersionReleaseDate                     types.String                                                         `tfsdk:"products_camera_current_version_release_date"`
	ProductsCameraCurrentVersionReleaseType                     types.String                                                         `tfsdk:"products_camera_current_version_release_type"`
	ProductsCameraCurrentVersionShortName                       types.String                                                         `tfsdk:"products_camera_current_version_short_name"`
	ProductsCameraLastUpgradeTime                               types.String                                                         `tfsdk:"products_camera_last_upgrade_time"`
	ProductsCameraLastUpgradeFromVersionFirmware                types.String                                                         `tfsdk:"products_camera_last_upgrade_from_version_firmware"`
	ProductsCameraLastUpgradeFromVersionId                      types.String                                                         `tfsdk:"products_camera_last_upgrade_from_version_id"`
	ProductsCameraLastUpgradeFromVersionReleaseDate             types.String                                                         `tfsdk:"products_camera_last_upgrade_from_version_release_date"`
	ProductsCameraLastUpgradeFromVersionReleaseType             types.String                                                         `tfsdk:"products_camera_last_upgrade_from_version_release_type"`
	ProductsCameraLastUpgradeFromVersionShortName               types.String                                                         `tfsdk:"products_camera_last_upgrade_from_version_short_name"`
	ProductsCameraLastUpgradeToVersionFirmware                  types.String                                                         `tfsdk:"products_camera_last_upgrade_to_version_firmware"`
	ProductsCameraLastUpgradeToVersionId                        types.String                                                         `tfsdk:"products_camera_last_upgrade_to_version_id"`
	ProductsCameraLastUpgradeToVersionReleaseDate               types.String                                                         `tfsdk:"products_camera_last_upgrade_to_version_release_date"`
	ProductsCameraLastUpgradeToVersionReleaseType               types.String                                                         `tfsdk:"products_camera_last_upgrade_to_version_release_type"`
	ProductsCameraLastUpgradeToVersionShortName                 types.String                                                         `tfsdk:"products_camera_last_upgrade_to_version_short_name"`
	ProductsCameraNextUpgradeTime                               types.String                                                         `tfsdk:"products_camera_next_upgrade_time"`
	ProductsCameraNextUpgradeToVersionFirmware                  types.String                                                         `tfsdk:"products_camera_next_upgrade_to_version_firmware"`
	ProductsCameraNextUpgradeToVersionId                        types.String                                                         `tfsdk:"products_camera_next_upgrade_to_version_id"`
	ProductsCameraNextUpgradeToVersionReleaseDate               types.String                                                         `tfsdk:"products_camera_next_upgrade_to_version_release_date"`
	ProductsCameraNextUpgradeToVersionReleaseType               types.String                                                         `tfsdk:"products_camera_next_upgrade_to_version_release_type"`
	ProductsCameraNextUpgradeToVersionShortName                 types.String                                                         `tfsdk:"products_camera_next_upgrade_to_version_short_name"`
	ProductsCameraAvailableVersions                             []NetworkFirmwareUpgradesProductsCameraAvailableVersions             `tfsdk:"products_camera_available_versions"`
	ProductsCellularGatewayIsUpgradeAvailable                   types.Bool                                                           `tfsdk:"products_cellular_gateway_is_upgrade_available"`
	ProductsCellularGatewayParticipateInNextBetaRelease         types.Bool                                                           `tfsdk:"products_cellular_gateway_participate_in_next_beta_release"`
	ProductsCellularGatewayCurrentVersionFirmware               types.String                                                         `tfsdk:"products_cellular_gateway_current_version_firmware"`
	ProductsCellularGatewayCurrentVersionId                     types.String                                                         `tfsdk:"products_cellular_gateway_current_version_id"`
	ProductsCellularGatewayCurrentVersionReleaseDate            types.String                                                         `tfsdk:"products_cellular_gateway_current_version_release_date"`
	ProductsCellularGatewayCurrentVersionReleaseType            types.String                                                         `tfsdk:"products_cellular_gateway_current_version_release_type"`
	ProductsCellularGatewayCurrentVersionShortName              types.String                                                         `tfsdk:"products_cellular_gateway_current_version_short_name"`
	ProductsCellularGatewayLastUpgradeTime                      types.String                                                         `tfsdk:"products_cellular_gateway_last_upgrade_time"`
	ProductsCellularGatewayLastUpgradeFromVersionFirmware       types.String                                                         `tfsdk:"products_cellular_gateway_last_upgrade_from_version_firmware"`
	ProductsCellularGatewayLastUpgradeFromVersionId             types.String                                                         `tfsdk:"products_cellular_gateway_last_upgrade_from_version_id"`
	ProductsCellularGatewayLastUpgradeFromVersionReleaseDate    types.String                                                         `tfsdk:"products_cellular_gateway_last_upgrade_from_version_release_date"`
	ProductsCellularGatewayLastUpgradeFromVersionReleaseType    types.String                                                         `tfsdk:"products_cellular_gateway_last_upgrade_from_version_release_type"`
	ProductsCellularGatewayLastUpgradeFromVersionShortName      types.String                                                         `tfsdk:"products_cellular_gateway_last_upgrade_from_version_short_name"`
	ProductsCellularGatewayLastUpgradeToVersionFirmware         types.String                                                         `tfsdk:"products_cellular_gateway_last_upgrade_to_version_firmware"`
	ProductsCellularGatewayLastUpgradeToVersionId               types.String                                                         `tfsdk:"products_cellular_gateway_last_upgrade_to_version_id"`
	ProductsCellularGatewayLastUpgradeToVersionReleaseDate      types.String                                                         `tfsdk:"products_cellular_gateway_last_upgrade_to_version_release_date"`
	ProductsCellularGatewayLastUpgradeToVersionReleaseType      types.String                                                         `tfsdk:"products_cellular_gateway_last_upgrade_to_version_release_type"`
	ProductsCellularGatewayLastUpgradeToVersionShortName        types.String                                                         `tfsdk:"products_cellular_gateway_last_upgrade_to_version_short_name"`
	ProductsCellularGatewayNextUpgradeTime                      types.String                                                         `tfsdk:"products_cellular_gateway_next_upgrade_time"`
	ProductsCellularGatewayNextUpgradeToVersionFirmware         types.String                                                         `tfsdk:"products_cellular_gateway_next_upgrade_to_version_firmware"`
	ProductsCellularGatewayNextUpgradeToVersionId               types.String                                                         `tfsdk:"products_cellular_gateway_next_upgrade_to_version_id"`
	ProductsCellularGatewayNextUpgradeToVersionReleaseDate      types.String                                                         `tfsdk:"products_cellular_gateway_next_upgrade_to_version_release_date"`
	ProductsCellularGatewayNextUpgradeToVersionReleaseType      types.String                                                         `tfsdk:"products_cellular_gateway_next_upgrade_to_version_release_type"`
	ProductsCellularGatewayNextUpgradeToVersionShortName        types.String                                                         `tfsdk:"products_cellular_gateway_next_upgrade_to_version_short_name"`
	ProductsCellularGatewayAvailableVersions                    []NetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions    `tfsdk:"products_cellular_gateway_available_versions"`
	ProductsSecureConnectIsUpgradeAvailable                     types.Bool                                                           `tfsdk:"products_secure_connect_is_upgrade_available"`
	ProductsSecureConnectParticipateInNextBetaRelease           types.Bool                                                           `tfsdk:"products_secure_connect_participate_in_next_beta_release"`
	ProductsSecureConnectCurrentVersionFirmware                 types.String                                                         `tfsdk:"products_secure_connect_current_version_firmware"`
	ProductsSecureConnectCurrentVersionId                       types.String                                                         `tfsdk:"products_secure_connect_current_version_id"`
	ProductsSecureConnectCurrentVersionReleaseDate              types.String                                                         `tfsdk:"products_secure_connect_current_version_release_date"`
	ProductsSecureConnectCurrentVersionReleaseType              types.String                                                         `tfsdk:"products_secure_connect_current_version_release_type"`
	ProductsSecureConnectCurrentVersionShortName                types.String                                                         `tfsdk:"products_secure_connect_current_version_short_name"`
	ProductsSecureConnectLastUpgradeTime                        types.String                                                         `tfsdk:"products_secure_connect_last_upgrade_time"`
	ProductsSecureConnectLastUpgradeFromVersionFirmware         types.String                                                         `tfsdk:"products_secure_connect_last_upgrade_from_version_firmware"`
	ProductsSecureConnectLastUpgradeFromVersionId               types.String                                                         `tfsdk:"products_secure_connect_last_upgrade_from_version_id"`
	ProductsSecureConnectLastUpgradeFromVersionReleaseDate      types.String                                                         `tfsdk:"products_secure_connect_last_upgrade_from_version_release_date"`
	ProductsSecureConnectLastUpgradeFromVersionReleaseType      types.String                                                         `tfsdk:"products_secure_connect_last_upgrade_from_version_release_type"`
	ProductsSecureConnectLastUpgradeFromVersionShortName        types.String                                                         `tfsdk:"products_secure_connect_last_upgrade_from_version_short_name"`
	ProductsSecureConnectLastUpgradeToVersionFirmware           types.String                                                         `tfsdk:"products_secure_connect_last_upgrade_to_version_firmware"`
	ProductsSecureConnectLastUpgradeToVersionId                 types.String                                                         `tfsdk:"products_secure_connect_last_upgrade_to_version_id"`
	ProductsSecureConnectLastUpgradeToVersionReleaseDate        types.String                                                         `tfsdk:"products_secure_connect_last_upgrade_to_version_release_date"`
	ProductsSecureConnectLastUpgradeToVersionReleaseType        types.String                                                         `tfsdk:"products_secure_connect_last_upgrade_to_version_release_type"`
	ProductsSecureConnectLastUpgradeToVersionShortName          types.String                                                         `tfsdk:"products_secure_connect_last_upgrade_to_version_short_name"`
	ProductsSecureConnectNextUpgradeTime                        types.String                                                         `tfsdk:"products_secure_connect_next_upgrade_time"`
	ProductsSecureConnectNextUpgradeToVersionFirmware           types.String                                                         `tfsdk:"products_secure_connect_next_upgrade_to_version_firmware"`
	ProductsSecureConnectNextUpgradeToVersionId                 types.String                                                         `tfsdk:"products_secure_connect_next_upgrade_to_version_id"`
	ProductsSecureConnectNextUpgradeToVersionReleaseDate        types.String                                                         `tfsdk:"products_secure_connect_next_upgrade_to_version_release_date"`
	ProductsSecureConnectNextUpgradeToVersionReleaseType        types.String                                                         `tfsdk:"products_secure_connect_next_upgrade_to_version_release_type"`
	ProductsSecureConnectNextUpgradeToVersionShortName          types.String                                                         `tfsdk:"products_secure_connect_next_upgrade_to_version_short_name"`
	ProductsSecureConnectAvailableVersions                      []NetworkFirmwareUpgradesProductsSecureConnectAvailableVersions      `tfsdk:"products_secure_connect_available_versions"`
	ProductsSensorIsUpgradeAvailable                            types.Bool                                                           `tfsdk:"products_sensor_is_upgrade_available"`
	ProductsSensorParticipateInNextBetaRelease                  types.Bool                                                           `tfsdk:"products_sensor_participate_in_next_beta_release"`
	ProductsSensorCurrentVersionFirmware                        types.String                                                         `tfsdk:"products_sensor_current_version_firmware"`
	ProductsSensorCurrentVersionId                              types.String                                                         `tfsdk:"products_sensor_current_version_id"`
	ProductsSensorCurrentVersionReleaseDate                     types.String                                                         `tfsdk:"products_sensor_current_version_release_date"`
	ProductsSensorCurrentVersionReleaseType                     types.String                                                         `tfsdk:"products_sensor_current_version_release_type"`
	ProductsSensorCurrentVersionShortName                       types.String                                                         `tfsdk:"products_sensor_current_version_short_name"`
	ProductsSensorLastUpgradeTime                               types.String                                                         `tfsdk:"products_sensor_last_upgrade_time"`
	ProductsSensorLastUpgradeFromVersionFirmware                types.String                                                         `tfsdk:"products_sensor_last_upgrade_from_version_firmware"`
	ProductsSensorLastUpgradeFromVersionId                      types.String                                                         `tfsdk:"products_sensor_last_upgrade_from_version_id"`
	ProductsSensorLastUpgradeFromVersionReleaseDate             types.String                                                         `tfsdk:"products_sensor_last_upgrade_from_version_release_date"`
	ProductsSensorLastUpgradeFromVersionReleaseType             types.String                                                         `tfsdk:"products_sensor_last_upgrade_from_version_release_type"`
	ProductsSensorLastUpgradeFromVersionShortName               types.String                                                         `tfsdk:"products_sensor_last_upgrade_from_version_short_name"`
	ProductsSensorLastUpgradeToVersionFirmware                  types.String                                                         `tfsdk:"products_sensor_last_upgrade_to_version_firmware"`
	ProductsSensorLastUpgradeToVersionId                        types.String                                                         `tfsdk:"products_sensor_last_upgrade_to_version_id"`
	ProductsSensorLastUpgradeToVersionReleaseDate               types.String                                                         `tfsdk:"products_sensor_last_upgrade_to_version_release_date"`
	ProductsSensorLastUpgradeToVersionReleaseType               types.String                                                         `tfsdk:"products_sensor_last_upgrade_to_version_release_type"`
	ProductsSensorLastUpgradeToVersionShortName                 types.String                                                         `tfsdk:"products_sensor_last_upgrade_to_version_short_name"`
	ProductsSensorNextUpgradeTime                               types.String                                                         `tfsdk:"products_sensor_next_upgrade_time"`
	ProductsSensorNextUpgradeToVersionFirmware                  types.String                                                         `tfsdk:"products_sensor_next_upgrade_to_version_firmware"`
	ProductsSensorNextUpgradeToVersionId                        types.String                                                         `tfsdk:"products_sensor_next_upgrade_to_version_id"`
	ProductsSensorNextUpgradeToVersionReleaseDate               types.String                                                         `tfsdk:"products_sensor_next_upgrade_to_version_release_date"`
	ProductsSensorNextUpgradeToVersionReleaseType               types.String                                                         `tfsdk:"products_sensor_next_upgrade_to_version_release_type"`
	ProductsSensorNextUpgradeToVersionShortName                 types.String                                                         `tfsdk:"products_sensor_next_upgrade_to_version_short_name"`
	ProductsSensorAvailableVersions                             []NetworkFirmwareUpgradesProductsSensorAvailableVersions             `tfsdk:"products_sensor_available_versions"`
	ProductsSwitchIsUpgradeAvailable                            types.Bool                                                           `tfsdk:"products_switch_is_upgrade_available"`
	ProductsSwitchParticipateInNextBetaRelease                  types.Bool                                                           `tfsdk:"products_switch_participate_in_next_beta_release"`
	ProductsSwitchCurrentVersionFirmware                        types.String                                                         `tfsdk:"products_switch_current_version_firmware"`
	ProductsSwitchCurrentVersionId                              types.String                                                         `tfsdk:"products_switch_current_version_id"`
	ProductsSwitchCurrentVersionReleaseDate                     types.String                                                         `tfsdk:"products_switch_current_version_release_date"`
	ProductsSwitchCurrentVersionReleaseType                     types.String                                                         `tfsdk:"products_switch_current_version_release_type"`
	ProductsSwitchCurrentVersionShortName                       types.String                                                         `tfsdk:"products_switch_current_version_short_name"`
	ProductsSwitchLastUpgradeTime                               types.String                                                         `tfsdk:"products_switch_last_upgrade_time"`
	ProductsSwitchLastUpgradeFromVersionFirmware                types.String                                                         `tfsdk:"products_switch_last_upgrade_from_version_firmware"`
	ProductsSwitchLastUpgradeFromVersionId                      types.String                                                         `tfsdk:"products_switch_last_upgrade_from_version_id"`
	ProductsSwitchLastUpgradeFromVersionReleaseDate             types.String                                                         `tfsdk:"products_switch_last_upgrade_from_version_release_date"`
	ProductsSwitchLastUpgradeFromVersionReleaseType             types.String                                                         `tfsdk:"products_switch_last_upgrade_from_version_release_type"`
	ProductsSwitchLastUpgradeFromVersionShortName               types.String                                                         `tfsdk:"products_switch_last_upgrade_from_version_short_name"`
	ProductsSwitchLastUpgradeToVersionFirmware                  types.String                                                         `tfsdk:"products_switch_last_upgrade_to_version_firmware"`
	ProductsSwitchLastUpgradeToVersionId                        types.String                                                         `tfsdk:"products_switch_last_upgrade_to_version_id"`
	ProductsSwitchLastUpgradeToVersionReleaseDate               types.String                                                         `tfsdk:"products_switch_last_upgrade_to_version_release_date"`
	ProductsSwitchLastUpgradeToVersionReleaseType               types.String                                                         `tfsdk:"products_switch_last_upgrade_to_version_release_type"`
	ProductsSwitchLastUpgradeToVersionShortName                 types.String                                                         `tfsdk:"products_switch_last_upgrade_to_version_short_name"`
	ProductsSwitchNextUpgradeTime                               types.String                                                         `tfsdk:"products_switch_next_upgrade_time"`
	ProductsSwitchNextUpgradeToVersionFirmware                  types.String                                                         `tfsdk:"products_switch_next_upgrade_to_version_firmware"`
	ProductsSwitchNextUpgradeToVersionId                        types.String                                                         `tfsdk:"products_switch_next_upgrade_to_version_id"`
	ProductsSwitchNextUpgradeToVersionReleaseDate               types.String                                                         `tfsdk:"products_switch_next_upgrade_to_version_release_date"`
	ProductsSwitchNextUpgradeToVersionReleaseType               types.String                                                         `tfsdk:"products_switch_next_upgrade_to_version_release_type"`
	ProductsSwitchNextUpgradeToVersionShortName                 types.String                                                         `tfsdk:"products_switch_next_upgrade_to_version_short_name"`
	ProductsSwitchAvailableVersions                             []NetworkFirmwareUpgradesProductsSwitchAvailableVersions             `tfsdk:"products_switch_available_versions"`
	ProductsWirelessIsUpgradeAvailable                          types.Bool                                                           `tfsdk:"products_wireless_is_upgrade_available"`
	ProductsSwitchCatalystParticipateInNextBetaRelease          types.Bool                                                           `tfsdk:"products_switch_catalyst_participate_in_next_beta_release"`
	ProductsSwitchCatalystNextUpgradeTime                       types.String                                                         `tfsdk:"products_switch_catalyst_next_upgrade_time"`
	ProductsSwitchCatalystNextUpgradeToVersionId                types.String                                                         `tfsdk:"products_switch_catalyst_next_upgrade_to_version_id"`
	ProductsWirelessParticipateInNextBetaRelease                types.Bool                                                           `tfsdk:"products_wireless_participate_in_next_beta_release"`
	ProductsWirelessCurrentVersionFirmware                      types.String                                                         `tfsdk:"products_wireless_current_version_firmware"`
	ProductsWirelessCurrentVersionId                            types.String                                                         `tfsdk:"products_wireless_current_version_id"`
	ProductsWirelessCurrentVersionReleaseDate                   types.String                                                         `tfsdk:"products_wireless_current_version_release_date"`
	ProductsWirelessCurrentVersionReleaseType                   types.String                                                         `tfsdk:"products_wireless_current_version_release_type"`
	ProductsWirelessCurrentVersionShortName                     types.String                                                         `tfsdk:"products_wireless_current_version_short_name"`
	ProductsWirelessLastUpgradeTime                             types.String                                                         `tfsdk:"products_wireless_last_upgrade_time"`
	ProductsWirelessLastUpgradeFromVersionFirmware              types.String                                                         `tfsdk:"products_wireless_last_upgrade_from_version_firmware"`
	ProductsWirelessLastUpgradeFromVersionId                    types.String                                                         `tfsdk:"products_wireless_last_upgrade_from_version_id"`
	ProductsWirelessLastUpgradeFromVersionReleaseDate           types.String                                                         `tfsdk:"products_wireless_last_upgrade_from_version_release_date"`
	ProductsWirelessLastUpgradeFromVersionReleaseType           types.String                                                         `tfsdk:"products_wireless_last_upgrade_from_version_release_type"`
	ProductsWirelessLastUpgradeFromVersionShortName             types.String                                                         `tfsdk:"products_wireless_last_upgrade_from_version_short_name"`
	ProductsWirelessLastUpgradeToVersionFirmware                types.String                                                         `tfsdk:"products_wireless_last_upgrade_to_version_firmware"`
	ProductsWirelessLastUpgradeToVersionId                      types.String                                                         `tfsdk:"products_wireless_last_upgrade_to_version_id"`
	ProductsWirelessLastUpgradeToVersionReleaseDate             types.String                                                         `tfsdk:"products_wireless_last_upgrade_to_version_release_date"`
	ProductsWirelessLastUpgradeToVersionReleaseType             types.String                                                         `tfsdk:"products_wireless_last_upgrade_to_version_release_type"`
	ProductsWirelessLastUpgradeToVersionShortName               types.String                                                         `tfsdk:"products_wireless_last_upgrade_to_version_short_name"`
	ProductsWirelessNextUpgradeTime                             types.String                                                         `tfsdk:"products_wireless_next_upgrade_time"`
	ProductsWirelessNextUpgradeToVersionFirmware                types.String                                                         `tfsdk:"products_wireless_next_upgrade_to_version_firmware"`
	ProductsWirelessNextUpgradeToVersionId                      types.String                                                         `tfsdk:"products_wireless_next_upgrade_to_version_id"`
	ProductsWirelessNextUpgradeToVersionReleaseDate             types.String                                                         `tfsdk:"products_wireless_next_upgrade_to_version_release_date"`
	ProductsWirelessNextUpgradeToVersionReleaseType             types.String                                                         `tfsdk:"products_wireless_next_upgrade_to_version_release_type"`
	ProductsWirelessNextUpgradeToVersionShortName               types.String                                                         `tfsdk:"products_wireless_next_upgrade_to_version_short_name"`
	ProductsWirelessAvailableVersions                           []NetworkFirmwareUpgradesProductsWirelessAvailableVersions           `tfsdk:"products_wireless_available_versions"`
	ProductsWirelessControllerIsUpgradeAvailable                types.Bool                                                           `tfsdk:"products_wireless_controller_is_upgrade_available"`
	ProductsWirelessControllerParticipateInNextBetaRelease      types.Bool                                                           `tfsdk:"products_wireless_controller_participate_in_next_beta_release"`
	ProductsWirelessControllerCurrentVersionFirmware            types.String                                                         `tfsdk:"products_wireless_controller_current_version_firmware"`
	ProductsWirelessControllerCurrentVersionId                  types.String                                                         `tfsdk:"products_wireless_controller_current_version_id"`
	ProductsWirelessControllerCurrentVersionReleaseDate         types.String                                                         `tfsdk:"products_wireless_controller_current_version_release_date"`
	ProductsWirelessControllerCurrentVersionReleaseType         types.String                                                         `tfsdk:"products_wireless_controller_current_version_release_type"`
	ProductsWirelessControllerCurrentVersionShortName           types.String                                                         `tfsdk:"products_wireless_controller_current_version_short_name"`
	ProductsWirelessControllerLastUpgradeTime                   types.String                                                         `tfsdk:"products_wireless_controller_last_upgrade_time"`
	ProductsWirelessControllerLastUpgradeFromVersionFirmware    types.String                                                         `tfsdk:"products_wireless_controller_last_upgrade_from_version_firmware"`
	ProductsWirelessControllerLastUpgradeFromVersionId          types.String                                                         `tfsdk:"products_wireless_controller_last_upgrade_from_version_id"`
	ProductsWirelessControllerLastUpgradeFromVersionReleaseDate types.String                                                         `tfsdk:"products_wireless_controller_last_upgrade_from_version_release_date"`
	ProductsWirelessControllerLastUpgradeFromVersionReleaseType types.String                                                         `tfsdk:"products_wireless_controller_last_upgrade_from_version_release_type"`
	ProductsWirelessControllerLastUpgradeFromVersionShortName   types.String                                                         `tfsdk:"products_wireless_controller_last_upgrade_from_version_short_name"`
	ProductsWirelessControllerLastUpgradeToVersionFirmware      types.String                                                         `tfsdk:"products_wireless_controller_last_upgrade_to_version_firmware"`
	ProductsWirelessControllerLastUpgradeToVersionId            types.String                                                         `tfsdk:"products_wireless_controller_last_upgrade_to_version_id"`
	ProductsWirelessControllerLastUpgradeToVersionReleaseDate   types.String                                                         `tfsdk:"products_wireless_controller_last_upgrade_to_version_release_date"`
	ProductsWirelessControllerLastUpgradeToVersionReleaseType   types.String                                                         `tfsdk:"products_wireless_controller_last_upgrade_to_version_release_type"`
	ProductsWirelessControllerLastUpgradeToVersionShortName     types.String                                                         `tfsdk:"products_wireless_controller_last_upgrade_to_version_short_name"`
	ProductsWirelessControllerNextUpgradeTime                   types.String                                                         `tfsdk:"products_wireless_controller_next_upgrade_time"`
	ProductsWirelessControllerNextUpgradeToVersionFirmware      types.String                                                         `tfsdk:"products_wireless_controller_next_upgrade_to_version_firmware"`
	ProductsWirelessControllerNextUpgradeToVersionId            types.String                                                         `tfsdk:"products_wireless_controller_next_upgrade_to_version_id"`
	ProductsWirelessControllerNextUpgradeToVersionReleaseDate   types.String                                                         `tfsdk:"products_wireless_controller_next_upgrade_to_version_release_date"`
	ProductsWirelessControllerNextUpgradeToVersionReleaseType   types.String                                                         `tfsdk:"products_wireless_controller_next_upgrade_to_version_release_type"`
	ProductsWirelessControllerNextUpgradeToVersionShortName     types.String                                                         `tfsdk:"products_wireless_controller_next_upgrade_to_version_short_name"`
	ProductsWirelessControllerAvailableVersions                 []NetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions `tfsdk:"products_wireless_controller_available_versions"`
	UpgradeWindowDayOfWeek                                      types.String                                                         `tfsdk:"upgrade_window_day_of_week"`
	UpgradeWindowHourOfDay                                      types.String                                                         `tfsdk:"upgrade_window_hour_of_day"`
}

type NetworkFirmwareUpgradesProductsApplianceAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type NetworkFirmwareUpgradesProductsCameraAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type NetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type NetworkFirmwareUpgradesProductsSecureConnectAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type NetworkFirmwareUpgradesProductsSensorAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type NetworkFirmwareUpgradesProductsSwitchAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type NetworkFirmwareUpgradesProductsWirelessAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type NetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
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
	if value := res.Get("products.appliance.isUpgradeAvailable"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsApplianceIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.appliance.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsApplianceParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.appliance.currentVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsApplianceCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.appliance.currentVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsApplianceCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.appliance.currentVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsApplianceCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.appliance.currentVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsApplianceCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.appliance.currentVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsApplianceCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.fromVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.fromVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.fromVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.fromVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.fromVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.appliance.availableVersions"); value.Exists() && value.Value() != nil {
		data.ProductsApplianceAvailableVersions = make([]NetworkFirmwareUpgradesProductsApplianceAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesProductsApplianceAvailableVersions{}
			if value := res.Get("firmware"); value.Exists() && value.Value() != nil {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("releaseDate"); value.Exists() && value.Value() != nil {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
			if value := res.Get("releaseType"); value.Exists() && value.Value() != nil {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
			if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
			(*parent).ProductsApplianceAvailableVersions = append((*parent).ProductsApplianceAvailableVersions, data)
			return true
		})
	}
	if value := res.Get("products.camera.isUpgradeAvailable"); value.Exists() && value.Value() != nil {
		data.ProductsCameraIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsCameraIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.camera.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.camera.currentVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsCameraCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCameraCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.camera.currentVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsCameraCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCameraCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.camera.currentVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsCameraCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCameraCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.camera.currentVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsCameraCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCameraCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.camera.currentVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsCameraCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCameraCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsCameraLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.fromVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsCameraLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.fromVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsCameraLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.fromVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsCameraLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.fromVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsCameraLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.fromVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsCameraLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsCameraLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsCameraLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsCameraLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsCameraLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsCameraLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.camera.nextUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsCameraNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.camera.nextUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsCameraNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.camera.nextUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsCameraNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.camera.nextUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsCameraNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.camera.availableVersions"); value.Exists() && value.Value() != nil {
		data.ProductsCameraAvailableVersions = make([]NetworkFirmwareUpgradesProductsCameraAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesProductsCameraAvailableVersions{}
			if value := res.Get("firmware"); value.Exists() && value.Value() != nil {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("releaseDate"); value.Exists() && value.Value() != nil {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
			if value := res.Get("releaseType"); value.Exists() && value.Value() != nil {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
			if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
			(*parent).ProductsCameraAvailableVersions = append((*parent).ProductsCameraAvailableVersions, data)
			return true
		})
	}
	if value := res.Get("products.cellularGateway.isUpgradeAvailable"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsCellularGatewayIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.cellularGateway.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.cellularGateway.currentVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.currentVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.currentVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.currentVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.currentVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.availableVersions"); value.Exists() && value.Value() != nil {
		data.ProductsCellularGatewayAvailableVersions = make([]NetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions{}
			if value := res.Get("firmware"); value.Exists() && value.Value() != nil {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("releaseDate"); value.Exists() && value.Value() != nil {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
			if value := res.Get("releaseType"); value.Exists() && value.Value() != nil {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
			if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
			(*parent).ProductsCellularGatewayAvailableVersions = append((*parent).ProductsCellularGatewayAvailableVersions, data)
			return true
		})
	}
	if value := res.Get("products.secureConnect.isUpgradeAvailable"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsSecureConnectIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.secureConnect.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.secureConnect.currentVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.secureConnect.currentVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.secureConnect.currentVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.secureConnect.currentVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.secureConnect.currentVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.secureConnect.availableVersions"); value.Exists() && value.Value() != nil {
		data.ProductsSecureConnectAvailableVersions = make([]NetworkFirmwareUpgradesProductsSecureConnectAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesProductsSecureConnectAvailableVersions{}
			if value := res.Get("firmware"); value.Exists() && value.Value() != nil {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("releaseDate"); value.Exists() && value.Value() != nil {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
			if value := res.Get("releaseType"); value.Exists() && value.Value() != nil {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
			if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
			(*parent).ProductsSecureConnectAvailableVersions = append((*parent).ProductsSecureConnectAvailableVersions, data)
			return true
		})
	}
	if value := res.Get("products.sensor.isUpgradeAvailable"); value.Exists() && value.Value() != nil {
		data.ProductsSensorIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsSensorIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.sensor.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.sensor.currentVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSensorCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSensorCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.sensor.currentVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSensorCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSensorCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.sensor.currentVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSensorCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSensorCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.sensor.currentVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSensorCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSensorCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.sensor.currentVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSensorCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSensorCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSensorLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.fromVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSensorLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.fromVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSensorLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.fromVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSensorLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.fromVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSensorLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.fromVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSensorLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSensorLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSensorLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSensorLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSensorLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSensorLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSensorNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSensorNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSensorNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSensorNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.sensor.availableVersions"); value.Exists() && value.Value() != nil {
		data.ProductsSensorAvailableVersions = make([]NetworkFirmwareUpgradesProductsSensorAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesProductsSensorAvailableVersions{}
			if value := res.Get("firmware"); value.Exists() && value.Value() != nil {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("releaseDate"); value.Exists() && value.Value() != nil {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
			if value := res.Get("releaseType"); value.Exists() && value.Value() != nil {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
			if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
			(*parent).ProductsSensorAvailableVersions = append((*parent).ProductsSensorAvailableVersions, data)
			return true
		})
	}
	if value := res.Get("products.switch.isUpgradeAvailable"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.switch.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switch.currentVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switch.currentVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.switch.currentVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switch.currentVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switch.currentVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.fromVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.fromVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.fromVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.fromVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.fromVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.switch.nextUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switch.nextUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switch.nextUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switch.nextUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.switch.availableVersions"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchAvailableVersions = make([]NetworkFirmwareUpgradesProductsSwitchAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesProductsSwitchAvailableVersions{}
			if value := res.Get("firmware"); value.Exists() && value.Value() != nil {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("releaseDate"); value.Exists() && value.Value() != nil {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
			if value := res.Get("releaseType"); value.Exists() && value.Value() != nil {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
			if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
			(*parent).ProductsSwitchAvailableVersions = append((*parent).ProductsSwitchAvailableVersions, data)
			return true
		})
	}
	if value := res.Get("products.wireless.isUpgradeAvailable"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessIsUpgradeAvailable = types.BoolNull()
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
	if value := res.Get("products.wireless.currentVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wireless.currentVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.wireless.currentVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wireless.currentVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wireless.currentVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.fromVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.fromVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.fromVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.fromVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.fromVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wireless.availableVersions"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessAvailableVersions = make([]NetworkFirmwareUpgradesProductsWirelessAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesProductsWirelessAvailableVersions{}
			if value := res.Get("firmware"); value.Exists() && value.Value() != nil {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("releaseDate"); value.Exists() && value.Value() != nil {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
			if value := res.Get("releaseType"); value.Exists() && value.Value() != nil {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
			if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
			(*parent).ProductsWirelessAvailableVersions = append((*parent).ProductsWirelessAvailableVersions, data)
			return true
		})
	}
	if value := res.Get("products.wirelessController.isUpgradeAvailable"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessControllerIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.wirelessController.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wirelessController.currentVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wirelessController.currentVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.wirelessController.currentVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wirelessController.currentVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wirelessController.currentVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wirelessController.availableVersions"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessControllerAvailableVersions = make([]NetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := NetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions{}
			if value := res.Get("firmware"); value.Exists() && value.Value() != nil {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("releaseDate"); value.Exists() && value.Value() != nil {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
			if value := res.Get("releaseType"); value.Exists() && value.Value() != nil {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
			if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
			(*parent).ProductsWirelessControllerAvailableVersions = append((*parent).ProductsWirelessControllerAvailableVersions, data)
			return true
		})
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
	if value := res.Get("products.appliance.isUpgradeAvailable"); value.Exists() && !data.ProductsApplianceIsUpgradeAvailable.IsNull() {
		data.ProductsApplianceIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsApplianceIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.appliance.participateInNextBetaRelease"); value.Exists() && !data.ProductsApplianceParticipateInNextBetaRelease.IsNull() {
		data.ProductsApplianceParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsApplianceParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.appliance.currentVersion.firmware"); value.Exists() && !data.ProductsApplianceCurrentVersionFirmware.IsNull() {
		data.ProductsApplianceCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsApplianceCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.appliance.currentVersion.id"); value.Exists() && !data.ProductsApplianceCurrentVersionId.IsNull() {
		data.ProductsApplianceCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsApplianceCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.appliance.currentVersion.releaseDate"); value.Exists() && !data.ProductsApplianceCurrentVersionReleaseDate.IsNull() {
		data.ProductsApplianceCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsApplianceCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.appliance.currentVersion.releaseType"); value.Exists() && !data.ProductsApplianceCurrentVersionReleaseType.IsNull() {
		data.ProductsApplianceCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsApplianceCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.appliance.currentVersion.shortName"); value.Exists() && !data.ProductsApplianceCurrentVersionShortName.IsNull() {
		data.ProductsApplianceCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsApplianceCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.time"); value.Exists() && !data.ProductsApplianceLastUpgradeTime.IsNull() {
		data.ProductsApplianceLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsApplianceLastUpgradeFromVersionFirmware.IsNull() {
		data.ProductsApplianceLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsApplianceLastUpgradeFromVersionId.IsNull() {
		data.ProductsApplianceLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsApplianceLastUpgradeFromVersionReleaseDate.IsNull() {
		data.ProductsApplianceLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsApplianceLastUpgradeFromVersionReleaseType.IsNull() {
		data.ProductsApplianceLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsApplianceLastUpgradeFromVersionShortName.IsNull() {
		data.ProductsApplianceLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsApplianceLastUpgradeToVersionFirmware.IsNull() {
		data.ProductsApplianceLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsApplianceLastUpgradeToVersionId.IsNull() {
		data.ProductsApplianceLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsApplianceLastUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsApplianceLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsApplianceLastUpgradeToVersionReleaseType.IsNull() {
		data.ProductsApplianceLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.appliance.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsApplianceLastUpgradeToVersionShortName.IsNull() {
		data.ProductsApplianceLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsApplianceLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsApplianceNextUpgradeToVersionFirmware.IsNull() {
		data.ProductsApplianceNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsApplianceNextUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsApplianceNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsApplianceNextUpgradeToVersionReleaseType.IsNull() {
		data.ProductsApplianceNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.appliance.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsApplianceNextUpgradeToVersionShortName.IsNull() {
		data.ProductsApplianceNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsApplianceNextUpgradeToVersionShortName = types.StringNull()
	}
	for i := 0; i < len(data.ProductsApplianceAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsApplianceAvailableVersions[i].Firmware.ValueString(), data.ProductsApplianceAvailableVersions[i].Id.ValueString(), data.ProductsApplianceAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsApplianceAvailableVersions[i].ReleaseType.ValueString(), data.ProductsApplianceAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsApplianceAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.appliance.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsApplianceAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsApplianceAvailableVersions[i],
			))
			(*parent).ProductsApplianceAvailableVersions = slices.Delete((*parent).ProductsApplianceAvailableVersions, i, i+1)
			i--

			continue
		}
		if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
			data.Firmware = types.StringValue(value.String())
		} else {
			data.Firmware = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
			data.ReleaseDate = types.StringValue(value.String())
		} else {
			data.ReleaseDate = types.StringNull()
		}
		if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
			data.ReleaseType = types.StringValue(value.String())
		} else {
			data.ReleaseType = types.StringNull()
		}
		if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
			data.ShortName = types.StringValue(value.String())
		} else {
			data.ShortName = types.StringNull()
		}
		(*parent).ProductsApplianceAvailableVersions[i] = data
	}
	if value := res.Get("products.camera.isUpgradeAvailable"); value.Exists() && !data.ProductsCameraIsUpgradeAvailable.IsNull() {
		data.ProductsCameraIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsCameraIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.camera.participateInNextBetaRelease"); value.Exists() && !data.ProductsCameraParticipateInNextBetaRelease.IsNull() {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCameraParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.camera.currentVersion.firmware"); value.Exists() && !data.ProductsCameraCurrentVersionFirmware.IsNull() {
		data.ProductsCameraCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCameraCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.camera.currentVersion.id"); value.Exists() && !data.ProductsCameraCurrentVersionId.IsNull() {
		data.ProductsCameraCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCameraCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.camera.currentVersion.releaseDate"); value.Exists() && !data.ProductsCameraCurrentVersionReleaseDate.IsNull() {
		data.ProductsCameraCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCameraCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.camera.currentVersion.releaseType"); value.Exists() && !data.ProductsCameraCurrentVersionReleaseType.IsNull() {
		data.ProductsCameraCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCameraCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.camera.currentVersion.shortName"); value.Exists() && !data.ProductsCameraCurrentVersionShortName.IsNull() {
		data.ProductsCameraCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCameraCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.time"); value.Exists() && !data.ProductsCameraLastUpgradeTime.IsNull() {
		data.ProductsCameraLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsCameraLastUpgradeFromVersionFirmware.IsNull() {
		data.ProductsCameraLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsCameraLastUpgradeFromVersionId.IsNull() {
		data.ProductsCameraLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsCameraLastUpgradeFromVersionReleaseDate.IsNull() {
		data.ProductsCameraLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsCameraLastUpgradeFromVersionReleaseType.IsNull() {
		data.ProductsCameraLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsCameraLastUpgradeFromVersionShortName.IsNull() {
		data.ProductsCameraLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsCameraLastUpgradeToVersionFirmware.IsNull() {
		data.ProductsCameraLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsCameraLastUpgradeToVersionId.IsNull() {
		data.ProductsCameraLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsCameraLastUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsCameraLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsCameraLastUpgradeToVersionReleaseType.IsNull() {
		data.ProductsCameraLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.camera.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsCameraLastUpgradeToVersionShortName.IsNull() {
		data.ProductsCameraLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCameraLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.camera.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsCameraNextUpgradeToVersionFirmware.IsNull() {
		data.ProductsCameraNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.camera.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsCameraNextUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsCameraNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.camera.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsCameraNextUpgradeToVersionReleaseType.IsNull() {
		data.ProductsCameraNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.camera.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsCameraNextUpgradeToVersionShortName.IsNull() {
		data.ProductsCameraNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCameraNextUpgradeToVersionShortName = types.StringNull()
	}
	for i := 0; i < len(data.ProductsCameraAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsCameraAvailableVersions[i].Firmware.ValueString(), data.ProductsCameraAvailableVersions[i].Id.ValueString(), data.ProductsCameraAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsCameraAvailableVersions[i].ReleaseType.ValueString(), data.ProductsCameraAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsCameraAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.camera.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsCameraAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsCameraAvailableVersions[i],
			))
			(*parent).ProductsCameraAvailableVersions = slices.Delete((*parent).ProductsCameraAvailableVersions, i, i+1)
			i--

			continue
		}
		if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
			data.Firmware = types.StringValue(value.String())
		} else {
			data.Firmware = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
			data.ReleaseDate = types.StringValue(value.String())
		} else {
			data.ReleaseDate = types.StringNull()
		}
		if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
			data.ReleaseType = types.StringValue(value.String())
		} else {
			data.ReleaseType = types.StringNull()
		}
		if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
			data.ShortName = types.StringValue(value.String())
		} else {
			data.ShortName = types.StringNull()
		}
		(*parent).ProductsCameraAvailableVersions[i] = data
	}
	if value := res.Get("products.cellularGateway.isUpgradeAvailable"); value.Exists() && !data.ProductsCellularGatewayIsUpgradeAvailable.IsNull() {
		data.ProductsCellularGatewayIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsCellularGatewayIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.cellularGateway.participateInNextBetaRelease"); value.Exists() && !data.ProductsCellularGatewayParticipateInNextBetaRelease.IsNull() {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsCellularGatewayParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.cellularGateway.currentVersion.firmware"); value.Exists() && !data.ProductsCellularGatewayCurrentVersionFirmware.IsNull() {
		data.ProductsCellularGatewayCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.currentVersion.id"); value.Exists() && !data.ProductsCellularGatewayCurrentVersionId.IsNull() {
		data.ProductsCellularGatewayCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.currentVersion.releaseDate"); value.Exists() && !data.ProductsCellularGatewayCurrentVersionReleaseDate.IsNull() {
		data.ProductsCellularGatewayCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.currentVersion.releaseType"); value.Exists() && !data.ProductsCellularGatewayCurrentVersionReleaseType.IsNull() {
		data.ProductsCellularGatewayCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.currentVersion.shortName"); value.Exists() && !data.ProductsCellularGatewayCurrentVersionShortName.IsNull() {
		data.ProductsCellularGatewayCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.time"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeTime.IsNull() {
		data.ProductsCellularGatewayLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeFromVersionFirmware.IsNull() {
		data.ProductsCellularGatewayLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeFromVersionId.IsNull() {
		data.ProductsCellularGatewayLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeFromVersionReleaseDate.IsNull() {
		data.ProductsCellularGatewayLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeFromVersionReleaseType.IsNull() {
		data.ProductsCellularGatewayLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeFromVersionShortName.IsNull() {
		data.ProductsCellularGatewayLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeToVersionFirmware.IsNull() {
		data.ProductsCellularGatewayLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeToVersionId.IsNull() {
		data.ProductsCellularGatewayLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsCellularGatewayLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeToVersionReleaseType.IsNull() {
		data.ProductsCellularGatewayLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeToVersionShortName.IsNull() {
		data.ProductsCellularGatewayLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsCellularGatewayNextUpgradeToVersionFirmware.IsNull() {
		data.ProductsCellularGatewayNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsCellularGatewayNextUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsCellularGatewayNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsCellularGatewayNextUpgradeToVersionReleaseType.IsNull() {
		data.ProductsCellularGatewayNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsCellularGatewayNextUpgradeToVersionShortName.IsNull() {
		data.ProductsCellularGatewayNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsCellularGatewayNextUpgradeToVersionShortName = types.StringNull()
	}
	for i := 0; i < len(data.ProductsCellularGatewayAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsCellularGatewayAvailableVersions[i].Firmware.ValueString(), data.ProductsCellularGatewayAvailableVersions[i].Id.ValueString(), data.ProductsCellularGatewayAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsCellularGatewayAvailableVersions[i].ReleaseType.ValueString(), data.ProductsCellularGatewayAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsCellularGatewayAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.cellularGateway.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsCellularGatewayAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsCellularGatewayAvailableVersions[i],
			))
			(*parent).ProductsCellularGatewayAvailableVersions = slices.Delete((*parent).ProductsCellularGatewayAvailableVersions, i, i+1)
			i--

			continue
		}
		if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
			data.Firmware = types.StringValue(value.String())
		} else {
			data.Firmware = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
			data.ReleaseDate = types.StringValue(value.String())
		} else {
			data.ReleaseDate = types.StringNull()
		}
		if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
			data.ReleaseType = types.StringValue(value.String())
		} else {
			data.ReleaseType = types.StringNull()
		}
		if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
			data.ShortName = types.StringValue(value.String())
		} else {
			data.ShortName = types.StringNull()
		}
		(*parent).ProductsCellularGatewayAvailableVersions[i] = data
	}
	if value := res.Get("products.secureConnect.isUpgradeAvailable"); value.Exists() && !data.ProductsSecureConnectIsUpgradeAvailable.IsNull() {
		data.ProductsSecureConnectIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsSecureConnectIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.secureConnect.participateInNextBetaRelease"); value.Exists() && !data.ProductsSecureConnectParticipateInNextBetaRelease.IsNull() {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSecureConnectParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.secureConnect.currentVersion.firmware"); value.Exists() && !data.ProductsSecureConnectCurrentVersionFirmware.IsNull() {
		data.ProductsSecureConnectCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.secureConnect.currentVersion.id"); value.Exists() && !data.ProductsSecureConnectCurrentVersionId.IsNull() {
		data.ProductsSecureConnectCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.secureConnect.currentVersion.releaseDate"); value.Exists() && !data.ProductsSecureConnectCurrentVersionReleaseDate.IsNull() {
		data.ProductsSecureConnectCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.secureConnect.currentVersion.releaseType"); value.Exists() && !data.ProductsSecureConnectCurrentVersionReleaseType.IsNull() {
		data.ProductsSecureConnectCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.secureConnect.currentVersion.shortName"); value.Exists() && !data.ProductsSecureConnectCurrentVersionShortName.IsNull() {
		data.ProductsSecureConnectCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.time"); value.Exists() && !data.ProductsSecureConnectLastUpgradeTime.IsNull() {
		data.ProductsSecureConnectLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsSecureConnectLastUpgradeFromVersionFirmware.IsNull() {
		data.ProductsSecureConnectLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsSecureConnectLastUpgradeFromVersionId.IsNull() {
		data.ProductsSecureConnectLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsSecureConnectLastUpgradeFromVersionReleaseDate.IsNull() {
		data.ProductsSecureConnectLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsSecureConnectLastUpgradeFromVersionReleaseType.IsNull() {
		data.ProductsSecureConnectLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsSecureConnectLastUpgradeFromVersionShortName.IsNull() {
		data.ProductsSecureConnectLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSecureConnectLastUpgradeToVersionFirmware.IsNull() {
		data.ProductsSecureConnectLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsSecureConnectLastUpgradeToVersionId.IsNull() {
		data.ProductsSecureConnectLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSecureConnectLastUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsSecureConnectLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSecureConnectLastUpgradeToVersionReleaseType.IsNull() {
		data.ProductsSecureConnectLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.secureConnect.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSecureConnectLastUpgradeToVersionShortName.IsNull() {
		data.ProductsSecureConnectLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSecureConnectNextUpgradeToVersionFirmware.IsNull() {
		data.ProductsSecureConnectNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSecureConnectNextUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsSecureConnectNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSecureConnectNextUpgradeToVersionReleaseType.IsNull() {
		data.ProductsSecureConnectNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.secureConnect.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSecureConnectNextUpgradeToVersionShortName.IsNull() {
		data.ProductsSecureConnectNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSecureConnectNextUpgradeToVersionShortName = types.StringNull()
	}
	for i := 0; i < len(data.ProductsSecureConnectAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsSecureConnectAvailableVersions[i].Firmware.ValueString(), data.ProductsSecureConnectAvailableVersions[i].Id.ValueString(), data.ProductsSecureConnectAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsSecureConnectAvailableVersions[i].ReleaseType.ValueString(), data.ProductsSecureConnectAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsSecureConnectAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.secureConnect.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsSecureConnectAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsSecureConnectAvailableVersions[i],
			))
			(*parent).ProductsSecureConnectAvailableVersions = slices.Delete((*parent).ProductsSecureConnectAvailableVersions, i, i+1)
			i--

			continue
		}
		if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
			data.Firmware = types.StringValue(value.String())
		} else {
			data.Firmware = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
			data.ReleaseDate = types.StringValue(value.String())
		} else {
			data.ReleaseDate = types.StringNull()
		}
		if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
			data.ReleaseType = types.StringValue(value.String())
		} else {
			data.ReleaseType = types.StringNull()
		}
		if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
			data.ShortName = types.StringValue(value.String())
		} else {
			data.ShortName = types.StringNull()
		}
		(*parent).ProductsSecureConnectAvailableVersions[i] = data
	}
	if value := res.Get("products.sensor.isUpgradeAvailable"); value.Exists() && !data.ProductsSensorIsUpgradeAvailable.IsNull() {
		data.ProductsSensorIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsSensorIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.sensor.participateInNextBetaRelease"); value.Exists() && !data.ProductsSensorParticipateInNextBetaRelease.IsNull() {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSensorParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.sensor.currentVersion.firmware"); value.Exists() && !data.ProductsSensorCurrentVersionFirmware.IsNull() {
		data.ProductsSensorCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSensorCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.sensor.currentVersion.id"); value.Exists() && !data.ProductsSensorCurrentVersionId.IsNull() {
		data.ProductsSensorCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSensorCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.sensor.currentVersion.releaseDate"); value.Exists() && !data.ProductsSensorCurrentVersionReleaseDate.IsNull() {
		data.ProductsSensorCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSensorCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.sensor.currentVersion.releaseType"); value.Exists() && !data.ProductsSensorCurrentVersionReleaseType.IsNull() {
		data.ProductsSensorCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSensorCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.sensor.currentVersion.shortName"); value.Exists() && !data.ProductsSensorCurrentVersionShortName.IsNull() {
		data.ProductsSensorCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSensorCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.time"); value.Exists() && !data.ProductsSensorLastUpgradeTime.IsNull() {
		data.ProductsSensorLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsSensorLastUpgradeFromVersionFirmware.IsNull() {
		data.ProductsSensorLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsSensorLastUpgradeFromVersionId.IsNull() {
		data.ProductsSensorLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsSensorLastUpgradeFromVersionReleaseDate.IsNull() {
		data.ProductsSensorLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsSensorLastUpgradeFromVersionReleaseType.IsNull() {
		data.ProductsSensorLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsSensorLastUpgradeFromVersionShortName.IsNull() {
		data.ProductsSensorLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSensorLastUpgradeToVersionFirmware.IsNull() {
		data.ProductsSensorLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsSensorLastUpgradeToVersionId.IsNull() {
		data.ProductsSensorLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSensorLastUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsSensorLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSensorLastUpgradeToVersionReleaseType.IsNull() {
		data.ProductsSensorLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.sensor.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSensorLastUpgradeToVersionShortName.IsNull() {
		data.ProductsSensorLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSensorLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSensorNextUpgradeToVersionFirmware.IsNull() {
		data.ProductsSensorNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSensorNextUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsSensorNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSensorNextUpgradeToVersionReleaseType.IsNull() {
		data.ProductsSensorNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.sensor.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSensorNextUpgradeToVersionShortName.IsNull() {
		data.ProductsSensorNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSensorNextUpgradeToVersionShortName = types.StringNull()
	}
	for i := 0; i < len(data.ProductsSensorAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsSensorAvailableVersions[i].Firmware.ValueString(), data.ProductsSensorAvailableVersions[i].Id.ValueString(), data.ProductsSensorAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsSensorAvailableVersions[i].ReleaseType.ValueString(), data.ProductsSensorAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsSensorAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.sensor.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsSensorAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsSensorAvailableVersions[i],
			))
			(*parent).ProductsSensorAvailableVersions = slices.Delete((*parent).ProductsSensorAvailableVersions, i, i+1)
			i--

			continue
		}
		if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
			data.Firmware = types.StringValue(value.String())
		} else {
			data.Firmware = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
			data.ReleaseDate = types.StringValue(value.String())
		} else {
			data.ReleaseDate = types.StringNull()
		}
		if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
			data.ReleaseType = types.StringValue(value.String())
		} else {
			data.ReleaseType = types.StringNull()
		}
		if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
			data.ShortName = types.StringValue(value.String())
		} else {
			data.ShortName = types.StringNull()
		}
		(*parent).ProductsSensorAvailableVersions[i] = data
	}
	if value := res.Get("products.switch.isUpgradeAvailable"); value.Exists() && !data.ProductsSwitchIsUpgradeAvailable.IsNull() {
		data.ProductsSwitchIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.switch.participateInNextBetaRelease"); value.Exists() && !data.ProductsSwitchParticipateInNextBetaRelease.IsNull() {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switch.currentVersion.firmware"); value.Exists() && !data.ProductsSwitchCurrentVersionFirmware.IsNull() {
		data.ProductsSwitchCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switch.currentVersion.id"); value.Exists() && !data.ProductsSwitchCurrentVersionId.IsNull() {
		data.ProductsSwitchCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.switch.currentVersion.releaseDate"); value.Exists() && !data.ProductsSwitchCurrentVersionReleaseDate.IsNull() {
		data.ProductsSwitchCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switch.currentVersion.releaseType"); value.Exists() && !data.ProductsSwitchCurrentVersionReleaseType.IsNull() {
		data.ProductsSwitchCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switch.currentVersion.shortName"); value.Exists() && !data.ProductsSwitchCurrentVersionShortName.IsNull() {
		data.ProductsSwitchCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.time"); value.Exists() && !data.ProductsSwitchLastUpgradeTime.IsNull() {
		data.ProductsSwitchLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsSwitchLastUpgradeFromVersionFirmware.IsNull() {
		data.ProductsSwitchLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsSwitchLastUpgradeFromVersionId.IsNull() {
		data.ProductsSwitchLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsSwitchLastUpgradeFromVersionReleaseDate.IsNull() {
		data.ProductsSwitchLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsSwitchLastUpgradeFromVersionReleaseType.IsNull() {
		data.ProductsSwitchLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsSwitchLastUpgradeFromVersionShortName.IsNull() {
		data.ProductsSwitchLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSwitchLastUpgradeToVersionFirmware.IsNull() {
		data.ProductsSwitchLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsSwitchLastUpgradeToVersionId.IsNull() {
		data.ProductsSwitchLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSwitchLastUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsSwitchLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSwitchLastUpgradeToVersionReleaseType.IsNull() {
		data.ProductsSwitchLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switch.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSwitchLastUpgradeToVersionShortName.IsNull() {
		data.ProductsSwitchLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.switch.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSwitchNextUpgradeToVersionFirmware.IsNull() {
		data.ProductsSwitchNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switch.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSwitchNextUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsSwitchNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switch.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSwitchNextUpgradeToVersionReleaseType.IsNull() {
		data.ProductsSwitchNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switch.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSwitchNextUpgradeToVersionShortName.IsNull() {
		data.ProductsSwitchNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchNextUpgradeToVersionShortName = types.StringNull()
	}
	for i := 0; i < len(data.ProductsSwitchAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsSwitchAvailableVersions[i].Firmware.ValueString(), data.ProductsSwitchAvailableVersions[i].Id.ValueString(), data.ProductsSwitchAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsSwitchAvailableVersions[i].ReleaseType.ValueString(), data.ProductsSwitchAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsSwitchAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.switch.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsSwitchAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsSwitchAvailableVersions[i],
			))
			(*parent).ProductsSwitchAvailableVersions = slices.Delete((*parent).ProductsSwitchAvailableVersions, i, i+1)
			i--

			continue
		}
		if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
			data.Firmware = types.StringValue(value.String())
		} else {
			data.Firmware = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
			data.ReleaseDate = types.StringValue(value.String())
		} else {
			data.ReleaseDate = types.StringNull()
		}
		if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
			data.ReleaseType = types.StringValue(value.String())
		} else {
			data.ReleaseType = types.StringNull()
		}
		if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
			data.ShortName = types.StringValue(value.String())
		} else {
			data.ShortName = types.StringNull()
		}
		(*parent).ProductsSwitchAvailableVersions[i] = data
	}
	if value := res.Get("products.wireless.isUpgradeAvailable"); value.Exists() && !data.ProductsWirelessIsUpgradeAvailable.IsNull() {
		data.ProductsWirelessIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessIsUpgradeAvailable = types.BoolNull()
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
	if value := res.Get("products.wireless.currentVersion.firmware"); value.Exists() && !data.ProductsWirelessCurrentVersionFirmware.IsNull() {
		data.ProductsWirelessCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wireless.currentVersion.id"); value.Exists() && !data.ProductsWirelessCurrentVersionId.IsNull() {
		data.ProductsWirelessCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.wireless.currentVersion.releaseDate"); value.Exists() && !data.ProductsWirelessCurrentVersionReleaseDate.IsNull() {
		data.ProductsWirelessCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wireless.currentVersion.releaseType"); value.Exists() && !data.ProductsWirelessCurrentVersionReleaseType.IsNull() {
		data.ProductsWirelessCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wireless.currentVersion.shortName"); value.Exists() && !data.ProductsWirelessCurrentVersionShortName.IsNull() {
		data.ProductsWirelessCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.time"); value.Exists() && !data.ProductsWirelessLastUpgradeTime.IsNull() {
		data.ProductsWirelessLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsWirelessLastUpgradeFromVersionFirmware.IsNull() {
		data.ProductsWirelessLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsWirelessLastUpgradeFromVersionId.IsNull() {
		data.ProductsWirelessLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsWirelessLastUpgradeFromVersionReleaseDate.IsNull() {
		data.ProductsWirelessLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsWirelessLastUpgradeFromVersionReleaseType.IsNull() {
		data.ProductsWirelessLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsWirelessLastUpgradeFromVersionShortName.IsNull() {
		data.ProductsWirelessLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsWirelessLastUpgradeToVersionFirmware.IsNull() {
		data.ProductsWirelessLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsWirelessLastUpgradeToVersionId.IsNull() {
		data.ProductsWirelessLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsWirelessLastUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsWirelessLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsWirelessLastUpgradeToVersionReleaseType.IsNull() {
		data.ProductsWirelessLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wireless.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsWirelessLastUpgradeToVersionShortName.IsNull() {
		data.ProductsWirelessLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsWirelessNextUpgradeToVersionFirmware.IsNull() {
		data.ProductsWirelessNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsWirelessNextUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsWirelessNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsWirelessNextUpgradeToVersionReleaseType.IsNull() {
		data.ProductsWirelessNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wireless.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsWirelessNextUpgradeToVersionShortName.IsNull() {
		data.ProductsWirelessNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessNextUpgradeToVersionShortName = types.StringNull()
	}
	for i := 0; i < len(data.ProductsWirelessAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsWirelessAvailableVersions[i].Firmware.ValueString(), data.ProductsWirelessAvailableVersions[i].Id.ValueString(), data.ProductsWirelessAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsWirelessAvailableVersions[i].ReleaseType.ValueString(), data.ProductsWirelessAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsWirelessAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.wireless.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsWirelessAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsWirelessAvailableVersions[i],
			))
			(*parent).ProductsWirelessAvailableVersions = slices.Delete((*parent).ProductsWirelessAvailableVersions, i, i+1)
			i--

			continue
		}
		if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
			data.Firmware = types.StringValue(value.String())
		} else {
			data.Firmware = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
			data.ReleaseDate = types.StringValue(value.String())
		} else {
			data.ReleaseDate = types.StringNull()
		}
		if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
			data.ReleaseType = types.StringValue(value.String())
		} else {
			data.ReleaseType = types.StringNull()
		}
		if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
			data.ShortName = types.StringValue(value.String())
		} else {
			data.ShortName = types.StringNull()
		}
		(*parent).ProductsWirelessAvailableVersions[i] = data
	}
	if value := res.Get("products.wirelessController.isUpgradeAvailable"); value.Exists() && !data.ProductsWirelessControllerIsUpgradeAvailable.IsNull() {
		data.ProductsWirelessControllerIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessControllerIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.wirelessController.participateInNextBetaRelease"); value.Exists() && !data.ProductsWirelessControllerParticipateInNextBetaRelease.IsNull() {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessControllerParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.wirelessController.currentVersion.firmware"); value.Exists() && !data.ProductsWirelessControllerCurrentVersionFirmware.IsNull() {
		data.ProductsWirelessControllerCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wirelessController.currentVersion.id"); value.Exists() && !data.ProductsWirelessControllerCurrentVersionId.IsNull() {
		data.ProductsWirelessControllerCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.wirelessController.currentVersion.releaseDate"); value.Exists() && !data.ProductsWirelessControllerCurrentVersionReleaseDate.IsNull() {
		data.ProductsWirelessControllerCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wirelessController.currentVersion.releaseType"); value.Exists() && !data.ProductsWirelessControllerCurrentVersionReleaseType.IsNull() {
		data.ProductsWirelessControllerCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wirelessController.currentVersion.shortName"); value.Exists() && !data.ProductsWirelessControllerCurrentVersionShortName.IsNull() {
		data.ProductsWirelessControllerCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.time"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeTime.IsNull() {
		data.ProductsWirelessControllerLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeFromVersionFirmware.IsNull() {
		data.ProductsWirelessControllerLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeFromVersionId.IsNull() {
		data.ProductsWirelessControllerLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeFromVersionReleaseDate.IsNull() {
		data.ProductsWirelessControllerLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeFromVersionReleaseType.IsNull() {
		data.ProductsWirelessControllerLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeFromVersionShortName.IsNull() {
		data.ProductsWirelessControllerLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeToVersionFirmware.IsNull() {
		data.ProductsWirelessControllerLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeToVersionId.IsNull() {
		data.ProductsWirelessControllerLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsWirelessControllerLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeToVersionReleaseType.IsNull() {
		data.ProductsWirelessControllerLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wirelessController.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeToVersionShortName.IsNull() {
		data.ProductsWirelessControllerLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsWirelessControllerNextUpgradeToVersionFirmware.IsNull() {
		data.ProductsWirelessControllerNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsWirelessControllerNextUpgradeToVersionReleaseDate.IsNull() {
		data.ProductsWirelessControllerNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsWirelessControllerNextUpgradeToVersionReleaseType.IsNull() {
		data.ProductsWirelessControllerNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.wirelessController.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsWirelessControllerNextUpgradeToVersionShortName.IsNull() {
		data.ProductsWirelessControllerNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsWirelessControllerNextUpgradeToVersionShortName = types.StringNull()
	}
	for i := 0; i < len(data.ProductsWirelessControllerAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsWirelessControllerAvailableVersions[i].Firmware.ValueString(), data.ProductsWirelessControllerAvailableVersions[i].Id.ValueString(), data.ProductsWirelessControllerAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsWirelessControllerAvailableVersions[i].ReleaseType.ValueString(), data.ProductsWirelessControllerAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsWirelessControllerAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.wirelessController.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsWirelessControllerAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsWirelessControllerAvailableVersions[i],
			))
			(*parent).ProductsWirelessControllerAvailableVersions = slices.Delete((*parent).ProductsWirelessControllerAvailableVersions, i, i+1)
			i--

			continue
		}
		if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
			data.Firmware = types.StringValue(value.String())
		} else {
			data.Firmware = types.StringNull()
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
			data.ReleaseDate = types.StringValue(value.String())
		} else {
			data.ReleaseDate = types.StringNull()
		}
		if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
			data.ReleaseType = types.StringValue(value.String())
		} else {
			data.ReleaseType = types.StringNull()
		}
		if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
			data.ShortName = types.StringValue(value.String())
		} else {
			data.ShortName = types.StringNull()
		}
		(*parent).ProductsWirelessControllerAvailableVersions[i] = data
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
	if data.ProductsApplianceIsUpgradeAvailable.IsUnknown() {
		if value := res.Get("products.appliance.isUpgradeAvailable"); value.Exists() && !data.ProductsApplianceIsUpgradeAvailable.IsNull() {
			data.ProductsApplianceIsUpgradeAvailable = types.BoolValue(value.Bool())
		} else {
			data.ProductsApplianceIsUpgradeAvailable = types.BoolNull()
		}
	}
	if data.ProductsApplianceCurrentVersionFirmware.IsUnknown() {
		if value := res.Get("products.appliance.currentVersion.firmware"); value.Exists() && !data.ProductsApplianceCurrentVersionFirmware.IsNull() {
			data.ProductsApplianceCurrentVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsApplianceCurrentVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsApplianceCurrentVersionId.IsUnknown() {
		if value := res.Get("products.appliance.currentVersion.id"); value.Exists() && !data.ProductsApplianceCurrentVersionId.IsNull() {
			data.ProductsApplianceCurrentVersionId = types.StringValue(value.String())
		} else {
			data.ProductsApplianceCurrentVersionId = types.StringNull()
		}
	}
	if data.ProductsApplianceCurrentVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.appliance.currentVersion.releaseDate"); value.Exists() && !data.ProductsApplianceCurrentVersionReleaseDate.IsNull() {
			data.ProductsApplianceCurrentVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsApplianceCurrentVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsApplianceCurrentVersionReleaseType.IsUnknown() {
		if value := res.Get("products.appliance.currentVersion.releaseType"); value.Exists() && !data.ProductsApplianceCurrentVersionReleaseType.IsNull() {
			data.ProductsApplianceCurrentVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsApplianceCurrentVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsApplianceCurrentVersionShortName.IsUnknown() {
		if value := res.Get("products.appliance.currentVersion.shortName"); value.Exists() && !data.ProductsApplianceCurrentVersionShortName.IsNull() {
			data.ProductsApplianceCurrentVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsApplianceCurrentVersionShortName = types.StringNull()
		}
	}
	if data.ProductsApplianceLastUpgradeTime.IsUnknown() {
		if value := res.Get("products.appliance.lastUpgrade.time"); value.Exists() && !data.ProductsApplianceLastUpgradeTime.IsNull() {
			data.ProductsApplianceLastUpgradeTime = types.StringValue(value.String())
		} else {
			data.ProductsApplianceLastUpgradeTime = types.StringNull()
		}
	}
	if data.ProductsApplianceLastUpgradeFromVersionFirmware.IsUnknown() {
		if value := res.Get("products.appliance.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsApplianceLastUpgradeFromVersionFirmware.IsNull() {
			data.ProductsApplianceLastUpgradeFromVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsApplianceLastUpgradeFromVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsApplianceLastUpgradeFromVersionId.IsUnknown() {
		if value := res.Get("products.appliance.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsApplianceLastUpgradeFromVersionId.IsNull() {
			data.ProductsApplianceLastUpgradeFromVersionId = types.StringValue(value.String())
		} else {
			data.ProductsApplianceLastUpgradeFromVersionId = types.StringNull()
		}
	}
	if data.ProductsApplianceLastUpgradeFromVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.appliance.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsApplianceLastUpgradeFromVersionReleaseDate.IsNull() {
			data.ProductsApplianceLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsApplianceLastUpgradeFromVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsApplianceLastUpgradeFromVersionReleaseType.IsUnknown() {
		if value := res.Get("products.appliance.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsApplianceLastUpgradeFromVersionReleaseType.IsNull() {
			data.ProductsApplianceLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsApplianceLastUpgradeFromVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsApplianceLastUpgradeFromVersionShortName.IsUnknown() {
		if value := res.Get("products.appliance.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsApplianceLastUpgradeFromVersionShortName.IsNull() {
			data.ProductsApplianceLastUpgradeFromVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsApplianceLastUpgradeFromVersionShortName = types.StringNull()
		}
	}
	if data.ProductsApplianceLastUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.appliance.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsApplianceLastUpgradeToVersionFirmware.IsNull() {
			data.ProductsApplianceLastUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsApplianceLastUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsApplianceLastUpgradeToVersionId.IsUnknown() {
		if value := res.Get("products.appliance.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsApplianceLastUpgradeToVersionId.IsNull() {
			data.ProductsApplianceLastUpgradeToVersionId = types.StringValue(value.String())
		} else {
			data.ProductsApplianceLastUpgradeToVersionId = types.StringNull()
		}
	}
	if data.ProductsApplianceLastUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.appliance.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsApplianceLastUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsApplianceLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsApplianceLastUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsApplianceLastUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.appliance.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsApplianceLastUpgradeToVersionReleaseType.IsNull() {
			data.ProductsApplianceLastUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsApplianceLastUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsApplianceLastUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.appliance.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsApplianceLastUpgradeToVersionShortName.IsNull() {
			data.ProductsApplianceLastUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsApplianceLastUpgradeToVersionShortName = types.StringNull()
		}
	}
	if data.ProductsApplianceNextUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.appliance.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsApplianceNextUpgradeToVersionFirmware.IsNull() {
			data.ProductsApplianceNextUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsApplianceNextUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsApplianceNextUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.appliance.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsApplianceNextUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsApplianceNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsApplianceNextUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsApplianceNextUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.appliance.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsApplianceNextUpgradeToVersionReleaseType.IsNull() {
			data.ProductsApplianceNextUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsApplianceNextUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsApplianceNextUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.appliance.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsApplianceNextUpgradeToVersionShortName.IsNull() {
			data.ProductsApplianceNextUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsApplianceNextUpgradeToVersionShortName = types.StringNull()
		}
	}
	for i := 0; i < len(data.ProductsApplianceAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsApplianceAvailableVersions[i].Firmware.ValueString(), data.ProductsApplianceAvailableVersions[i].Id.ValueString(), data.ProductsApplianceAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsApplianceAvailableVersions[i].ReleaseType.ValueString(), data.ProductsApplianceAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsApplianceAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.appliance.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsApplianceAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsApplianceAvailableVersions[i],
			))
			(*parent).ProductsApplianceAvailableVersions = slices.Delete((*parent).ProductsApplianceAvailableVersions, i, i+1)
			i--

			continue
		}
		if data.Firmware.IsUnknown() {
			if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
		}
		if data.Id.IsUnknown() {
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
		}
		if data.ReleaseDate.IsUnknown() {
			if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
		}
		if data.ReleaseType.IsUnknown() {
			if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
		}
		if data.ShortName.IsUnknown() {
			if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
		}
		(*parent).ProductsApplianceAvailableVersions[i] = data
	}
	if data.ProductsCameraIsUpgradeAvailable.IsUnknown() {
		if value := res.Get("products.camera.isUpgradeAvailable"); value.Exists() && !data.ProductsCameraIsUpgradeAvailable.IsNull() {
			data.ProductsCameraIsUpgradeAvailable = types.BoolValue(value.Bool())
		} else {
			data.ProductsCameraIsUpgradeAvailable = types.BoolNull()
		}
	}
	if data.ProductsCameraCurrentVersionFirmware.IsUnknown() {
		if value := res.Get("products.camera.currentVersion.firmware"); value.Exists() && !data.ProductsCameraCurrentVersionFirmware.IsNull() {
			data.ProductsCameraCurrentVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsCameraCurrentVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsCameraCurrentVersionId.IsUnknown() {
		if value := res.Get("products.camera.currentVersion.id"); value.Exists() && !data.ProductsCameraCurrentVersionId.IsNull() {
			data.ProductsCameraCurrentVersionId = types.StringValue(value.String())
		} else {
			data.ProductsCameraCurrentVersionId = types.StringNull()
		}
	}
	if data.ProductsCameraCurrentVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.camera.currentVersion.releaseDate"); value.Exists() && !data.ProductsCameraCurrentVersionReleaseDate.IsNull() {
			data.ProductsCameraCurrentVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsCameraCurrentVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsCameraCurrentVersionReleaseType.IsUnknown() {
		if value := res.Get("products.camera.currentVersion.releaseType"); value.Exists() && !data.ProductsCameraCurrentVersionReleaseType.IsNull() {
			data.ProductsCameraCurrentVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsCameraCurrentVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsCameraCurrentVersionShortName.IsUnknown() {
		if value := res.Get("products.camera.currentVersion.shortName"); value.Exists() && !data.ProductsCameraCurrentVersionShortName.IsNull() {
			data.ProductsCameraCurrentVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsCameraCurrentVersionShortName = types.StringNull()
		}
	}
	if data.ProductsCameraLastUpgradeTime.IsUnknown() {
		if value := res.Get("products.camera.lastUpgrade.time"); value.Exists() && !data.ProductsCameraLastUpgradeTime.IsNull() {
			data.ProductsCameraLastUpgradeTime = types.StringValue(value.String())
		} else {
			data.ProductsCameraLastUpgradeTime = types.StringNull()
		}
	}
	if data.ProductsCameraLastUpgradeFromVersionFirmware.IsUnknown() {
		if value := res.Get("products.camera.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsCameraLastUpgradeFromVersionFirmware.IsNull() {
			data.ProductsCameraLastUpgradeFromVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsCameraLastUpgradeFromVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsCameraLastUpgradeFromVersionId.IsUnknown() {
		if value := res.Get("products.camera.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsCameraLastUpgradeFromVersionId.IsNull() {
			data.ProductsCameraLastUpgradeFromVersionId = types.StringValue(value.String())
		} else {
			data.ProductsCameraLastUpgradeFromVersionId = types.StringNull()
		}
	}
	if data.ProductsCameraLastUpgradeFromVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.camera.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsCameraLastUpgradeFromVersionReleaseDate.IsNull() {
			data.ProductsCameraLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsCameraLastUpgradeFromVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsCameraLastUpgradeFromVersionReleaseType.IsUnknown() {
		if value := res.Get("products.camera.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsCameraLastUpgradeFromVersionReleaseType.IsNull() {
			data.ProductsCameraLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsCameraLastUpgradeFromVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsCameraLastUpgradeFromVersionShortName.IsUnknown() {
		if value := res.Get("products.camera.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsCameraLastUpgradeFromVersionShortName.IsNull() {
			data.ProductsCameraLastUpgradeFromVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsCameraLastUpgradeFromVersionShortName = types.StringNull()
		}
	}
	if data.ProductsCameraLastUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.camera.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsCameraLastUpgradeToVersionFirmware.IsNull() {
			data.ProductsCameraLastUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsCameraLastUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsCameraLastUpgradeToVersionId.IsUnknown() {
		if value := res.Get("products.camera.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsCameraLastUpgradeToVersionId.IsNull() {
			data.ProductsCameraLastUpgradeToVersionId = types.StringValue(value.String())
		} else {
			data.ProductsCameraLastUpgradeToVersionId = types.StringNull()
		}
	}
	if data.ProductsCameraLastUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.camera.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsCameraLastUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsCameraLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsCameraLastUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsCameraLastUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.camera.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsCameraLastUpgradeToVersionReleaseType.IsNull() {
			data.ProductsCameraLastUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsCameraLastUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsCameraLastUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.camera.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsCameraLastUpgradeToVersionShortName.IsNull() {
			data.ProductsCameraLastUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsCameraLastUpgradeToVersionShortName = types.StringNull()
		}
	}
	if data.ProductsCameraNextUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.camera.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsCameraNextUpgradeToVersionFirmware.IsNull() {
			data.ProductsCameraNextUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsCameraNextUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsCameraNextUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.camera.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsCameraNextUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsCameraNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsCameraNextUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsCameraNextUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.camera.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsCameraNextUpgradeToVersionReleaseType.IsNull() {
			data.ProductsCameraNextUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsCameraNextUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsCameraNextUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.camera.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsCameraNextUpgradeToVersionShortName.IsNull() {
			data.ProductsCameraNextUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsCameraNextUpgradeToVersionShortName = types.StringNull()
		}
	}
	for i := 0; i < len(data.ProductsCameraAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsCameraAvailableVersions[i].Firmware.ValueString(), data.ProductsCameraAvailableVersions[i].Id.ValueString(), data.ProductsCameraAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsCameraAvailableVersions[i].ReleaseType.ValueString(), data.ProductsCameraAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsCameraAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.camera.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsCameraAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsCameraAvailableVersions[i],
			))
			(*parent).ProductsCameraAvailableVersions = slices.Delete((*parent).ProductsCameraAvailableVersions, i, i+1)
			i--

			continue
		}
		if data.Firmware.IsUnknown() {
			if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
		}
		if data.Id.IsUnknown() {
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
		}
		if data.ReleaseDate.IsUnknown() {
			if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
		}
		if data.ReleaseType.IsUnknown() {
			if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
		}
		if data.ShortName.IsUnknown() {
			if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
		}
		(*parent).ProductsCameraAvailableVersions[i] = data
	}
	if data.ProductsCellularGatewayIsUpgradeAvailable.IsUnknown() {
		if value := res.Get("products.cellularGateway.isUpgradeAvailable"); value.Exists() && !data.ProductsCellularGatewayIsUpgradeAvailable.IsNull() {
			data.ProductsCellularGatewayIsUpgradeAvailable = types.BoolValue(value.Bool())
		} else {
			data.ProductsCellularGatewayIsUpgradeAvailable = types.BoolNull()
		}
	}
	if data.ProductsCellularGatewayCurrentVersionFirmware.IsUnknown() {
		if value := res.Get("products.cellularGateway.currentVersion.firmware"); value.Exists() && !data.ProductsCellularGatewayCurrentVersionFirmware.IsNull() {
			data.ProductsCellularGatewayCurrentVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayCurrentVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayCurrentVersionId.IsUnknown() {
		if value := res.Get("products.cellularGateway.currentVersion.id"); value.Exists() && !data.ProductsCellularGatewayCurrentVersionId.IsNull() {
			data.ProductsCellularGatewayCurrentVersionId = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayCurrentVersionId = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayCurrentVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.cellularGateway.currentVersion.releaseDate"); value.Exists() && !data.ProductsCellularGatewayCurrentVersionReleaseDate.IsNull() {
			data.ProductsCellularGatewayCurrentVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayCurrentVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayCurrentVersionReleaseType.IsUnknown() {
		if value := res.Get("products.cellularGateway.currentVersion.releaseType"); value.Exists() && !data.ProductsCellularGatewayCurrentVersionReleaseType.IsNull() {
			data.ProductsCellularGatewayCurrentVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayCurrentVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayCurrentVersionShortName.IsUnknown() {
		if value := res.Get("products.cellularGateway.currentVersion.shortName"); value.Exists() && !data.ProductsCellularGatewayCurrentVersionShortName.IsNull() {
			data.ProductsCellularGatewayCurrentVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayCurrentVersionShortName = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayLastUpgradeTime.IsUnknown() {
		if value := res.Get("products.cellularGateway.lastUpgrade.time"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeTime.IsNull() {
			data.ProductsCellularGatewayLastUpgradeTime = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayLastUpgradeTime = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayLastUpgradeFromVersionFirmware.IsUnknown() {
		if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeFromVersionFirmware.IsNull() {
			data.ProductsCellularGatewayLastUpgradeFromVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayLastUpgradeFromVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayLastUpgradeFromVersionId.IsUnknown() {
		if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeFromVersionId.IsNull() {
			data.ProductsCellularGatewayLastUpgradeFromVersionId = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayLastUpgradeFromVersionId = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayLastUpgradeFromVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeFromVersionReleaseDate.IsNull() {
			data.ProductsCellularGatewayLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayLastUpgradeFromVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayLastUpgradeFromVersionReleaseType.IsUnknown() {
		if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeFromVersionReleaseType.IsNull() {
			data.ProductsCellularGatewayLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayLastUpgradeFromVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayLastUpgradeFromVersionShortName.IsUnknown() {
		if value := res.Get("products.cellularGateway.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeFromVersionShortName.IsNull() {
			data.ProductsCellularGatewayLastUpgradeFromVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayLastUpgradeFromVersionShortName = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayLastUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeToVersionFirmware.IsNull() {
			data.ProductsCellularGatewayLastUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayLastUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayLastUpgradeToVersionId.IsUnknown() {
		if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeToVersionId.IsNull() {
			data.ProductsCellularGatewayLastUpgradeToVersionId = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayLastUpgradeToVersionId = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayLastUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsCellularGatewayLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayLastUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayLastUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeToVersionReleaseType.IsNull() {
			data.ProductsCellularGatewayLastUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayLastUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayLastUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.cellularGateway.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsCellularGatewayLastUpgradeToVersionShortName.IsNull() {
			data.ProductsCellularGatewayLastUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayLastUpgradeToVersionShortName = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayNextUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsCellularGatewayNextUpgradeToVersionFirmware.IsNull() {
			data.ProductsCellularGatewayNextUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayNextUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayNextUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsCellularGatewayNextUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsCellularGatewayNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayNextUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayNextUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsCellularGatewayNextUpgradeToVersionReleaseType.IsNull() {
			data.ProductsCellularGatewayNextUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayNextUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsCellularGatewayNextUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.cellularGateway.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsCellularGatewayNextUpgradeToVersionShortName.IsNull() {
			data.ProductsCellularGatewayNextUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsCellularGatewayNextUpgradeToVersionShortName = types.StringNull()
		}
	}
	for i := 0; i < len(data.ProductsCellularGatewayAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsCellularGatewayAvailableVersions[i].Firmware.ValueString(), data.ProductsCellularGatewayAvailableVersions[i].Id.ValueString(), data.ProductsCellularGatewayAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsCellularGatewayAvailableVersions[i].ReleaseType.ValueString(), data.ProductsCellularGatewayAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsCellularGatewayAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.cellularGateway.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsCellularGatewayAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsCellularGatewayAvailableVersions[i],
			))
			(*parent).ProductsCellularGatewayAvailableVersions = slices.Delete((*parent).ProductsCellularGatewayAvailableVersions, i, i+1)
			i--

			continue
		}
		if data.Firmware.IsUnknown() {
			if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
		}
		if data.Id.IsUnknown() {
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
		}
		if data.ReleaseDate.IsUnknown() {
			if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
		}
		if data.ReleaseType.IsUnknown() {
			if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
		}
		if data.ShortName.IsUnknown() {
			if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
		}
		(*parent).ProductsCellularGatewayAvailableVersions[i] = data
	}
	if data.ProductsSecureConnectIsUpgradeAvailable.IsUnknown() {
		if value := res.Get("products.secureConnect.isUpgradeAvailable"); value.Exists() && !data.ProductsSecureConnectIsUpgradeAvailable.IsNull() {
			data.ProductsSecureConnectIsUpgradeAvailable = types.BoolValue(value.Bool())
		} else {
			data.ProductsSecureConnectIsUpgradeAvailable = types.BoolNull()
		}
	}
	if data.ProductsSecureConnectCurrentVersionFirmware.IsUnknown() {
		if value := res.Get("products.secureConnect.currentVersion.firmware"); value.Exists() && !data.ProductsSecureConnectCurrentVersionFirmware.IsNull() {
			data.ProductsSecureConnectCurrentVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectCurrentVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSecureConnectCurrentVersionId.IsUnknown() {
		if value := res.Get("products.secureConnect.currentVersion.id"); value.Exists() && !data.ProductsSecureConnectCurrentVersionId.IsNull() {
			data.ProductsSecureConnectCurrentVersionId = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectCurrentVersionId = types.StringNull()
		}
	}
	if data.ProductsSecureConnectCurrentVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.secureConnect.currentVersion.releaseDate"); value.Exists() && !data.ProductsSecureConnectCurrentVersionReleaseDate.IsNull() {
			data.ProductsSecureConnectCurrentVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectCurrentVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSecureConnectCurrentVersionReleaseType.IsUnknown() {
		if value := res.Get("products.secureConnect.currentVersion.releaseType"); value.Exists() && !data.ProductsSecureConnectCurrentVersionReleaseType.IsNull() {
			data.ProductsSecureConnectCurrentVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectCurrentVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSecureConnectCurrentVersionShortName.IsUnknown() {
		if value := res.Get("products.secureConnect.currentVersion.shortName"); value.Exists() && !data.ProductsSecureConnectCurrentVersionShortName.IsNull() {
			data.ProductsSecureConnectCurrentVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectCurrentVersionShortName = types.StringNull()
		}
	}
	if data.ProductsSecureConnectLastUpgradeTime.IsUnknown() {
		if value := res.Get("products.secureConnect.lastUpgrade.time"); value.Exists() && !data.ProductsSecureConnectLastUpgradeTime.IsNull() {
			data.ProductsSecureConnectLastUpgradeTime = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectLastUpgradeTime = types.StringNull()
		}
	}
	if data.ProductsSecureConnectLastUpgradeFromVersionFirmware.IsUnknown() {
		if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsSecureConnectLastUpgradeFromVersionFirmware.IsNull() {
			data.ProductsSecureConnectLastUpgradeFromVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectLastUpgradeFromVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSecureConnectLastUpgradeFromVersionId.IsUnknown() {
		if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsSecureConnectLastUpgradeFromVersionId.IsNull() {
			data.ProductsSecureConnectLastUpgradeFromVersionId = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectLastUpgradeFromVersionId = types.StringNull()
		}
	}
	if data.ProductsSecureConnectLastUpgradeFromVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsSecureConnectLastUpgradeFromVersionReleaseDate.IsNull() {
			data.ProductsSecureConnectLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectLastUpgradeFromVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSecureConnectLastUpgradeFromVersionReleaseType.IsUnknown() {
		if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsSecureConnectLastUpgradeFromVersionReleaseType.IsNull() {
			data.ProductsSecureConnectLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectLastUpgradeFromVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSecureConnectLastUpgradeFromVersionShortName.IsUnknown() {
		if value := res.Get("products.secureConnect.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsSecureConnectLastUpgradeFromVersionShortName.IsNull() {
			data.ProductsSecureConnectLastUpgradeFromVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectLastUpgradeFromVersionShortName = types.StringNull()
		}
	}
	if data.ProductsSecureConnectLastUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.secureConnect.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSecureConnectLastUpgradeToVersionFirmware.IsNull() {
			data.ProductsSecureConnectLastUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectLastUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSecureConnectLastUpgradeToVersionId.IsUnknown() {
		if value := res.Get("products.secureConnect.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsSecureConnectLastUpgradeToVersionId.IsNull() {
			data.ProductsSecureConnectLastUpgradeToVersionId = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectLastUpgradeToVersionId = types.StringNull()
		}
	}
	if data.ProductsSecureConnectLastUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.secureConnect.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSecureConnectLastUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsSecureConnectLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectLastUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSecureConnectLastUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.secureConnect.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSecureConnectLastUpgradeToVersionReleaseType.IsNull() {
			data.ProductsSecureConnectLastUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectLastUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSecureConnectLastUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.secureConnect.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSecureConnectLastUpgradeToVersionShortName.IsNull() {
			data.ProductsSecureConnectLastUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectLastUpgradeToVersionShortName = types.StringNull()
		}
	}
	if data.ProductsSecureConnectNextUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.secureConnect.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSecureConnectNextUpgradeToVersionFirmware.IsNull() {
			data.ProductsSecureConnectNextUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectNextUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSecureConnectNextUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.secureConnect.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSecureConnectNextUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsSecureConnectNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectNextUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSecureConnectNextUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.secureConnect.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSecureConnectNextUpgradeToVersionReleaseType.IsNull() {
			data.ProductsSecureConnectNextUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectNextUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSecureConnectNextUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.secureConnect.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSecureConnectNextUpgradeToVersionShortName.IsNull() {
			data.ProductsSecureConnectNextUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSecureConnectNextUpgradeToVersionShortName = types.StringNull()
		}
	}
	for i := 0; i < len(data.ProductsSecureConnectAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsSecureConnectAvailableVersions[i].Firmware.ValueString(), data.ProductsSecureConnectAvailableVersions[i].Id.ValueString(), data.ProductsSecureConnectAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsSecureConnectAvailableVersions[i].ReleaseType.ValueString(), data.ProductsSecureConnectAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsSecureConnectAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.secureConnect.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsSecureConnectAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsSecureConnectAvailableVersions[i],
			))
			(*parent).ProductsSecureConnectAvailableVersions = slices.Delete((*parent).ProductsSecureConnectAvailableVersions, i, i+1)
			i--

			continue
		}
		if data.Firmware.IsUnknown() {
			if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
		}
		if data.Id.IsUnknown() {
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
		}
		if data.ReleaseDate.IsUnknown() {
			if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
		}
		if data.ReleaseType.IsUnknown() {
			if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
		}
		if data.ShortName.IsUnknown() {
			if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
		}
		(*parent).ProductsSecureConnectAvailableVersions[i] = data
	}
	if data.ProductsSensorIsUpgradeAvailable.IsUnknown() {
		if value := res.Get("products.sensor.isUpgradeAvailable"); value.Exists() && !data.ProductsSensorIsUpgradeAvailable.IsNull() {
			data.ProductsSensorIsUpgradeAvailable = types.BoolValue(value.Bool())
		} else {
			data.ProductsSensorIsUpgradeAvailable = types.BoolNull()
		}
	}
	if data.ProductsSensorCurrentVersionFirmware.IsUnknown() {
		if value := res.Get("products.sensor.currentVersion.firmware"); value.Exists() && !data.ProductsSensorCurrentVersionFirmware.IsNull() {
			data.ProductsSensorCurrentVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSensorCurrentVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSensorCurrentVersionId.IsUnknown() {
		if value := res.Get("products.sensor.currentVersion.id"); value.Exists() && !data.ProductsSensorCurrentVersionId.IsNull() {
			data.ProductsSensorCurrentVersionId = types.StringValue(value.String())
		} else {
			data.ProductsSensorCurrentVersionId = types.StringNull()
		}
	}
	if data.ProductsSensorCurrentVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.sensor.currentVersion.releaseDate"); value.Exists() && !data.ProductsSensorCurrentVersionReleaseDate.IsNull() {
			data.ProductsSensorCurrentVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSensorCurrentVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSensorCurrentVersionReleaseType.IsUnknown() {
		if value := res.Get("products.sensor.currentVersion.releaseType"); value.Exists() && !data.ProductsSensorCurrentVersionReleaseType.IsNull() {
			data.ProductsSensorCurrentVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSensorCurrentVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSensorCurrentVersionShortName.IsUnknown() {
		if value := res.Get("products.sensor.currentVersion.shortName"); value.Exists() && !data.ProductsSensorCurrentVersionShortName.IsNull() {
			data.ProductsSensorCurrentVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSensorCurrentVersionShortName = types.StringNull()
		}
	}
	if data.ProductsSensorLastUpgradeTime.IsUnknown() {
		if value := res.Get("products.sensor.lastUpgrade.time"); value.Exists() && !data.ProductsSensorLastUpgradeTime.IsNull() {
			data.ProductsSensorLastUpgradeTime = types.StringValue(value.String())
		} else {
			data.ProductsSensorLastUpgradeTime = types.StringNull()
		}
	}
	if data.ProductsSensorLastUpgradeFromVersionFirmware.IsUnknown() {
		if value := res.Get("products.sensor.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsSensorLastUpgradeFromVersionFirmware.IsNull() {
			data.ProductsSensorLastUpgradeFromVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSensorLastUpgradeFromVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSensorLastUpgradeFromVersionId.IsUnknown() {
		if value := res.Get("products.sensor.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsSensorLastUpgradeFromVersionId.IsNull() {
			data.ProductsSensorLastUpgradeFromVersionId = types.StringValue(value.String())
		} else {
			data.ProductsSensorLastUpgradeFromVersionId = types.StringNull()
		}
	}
	if data.ProductsSensorLastUpgradeFromVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.sensor.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsSensorLastUpgradeFromVersionReleaseDate.IsNull() {
			data.ProductsSensorLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSensorLastUpgradeFromVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSensorLastUpgradeFromVersionReleaseType.IsUnknown() {
		if value := res.Get("products.sensor.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsSensorLastUpgradeFromVersionReleaseType.IsNull() {
			data.ProductsSensorLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSensorLastUpgradeFromVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSensorLastUpgradeFromVersionShortName.IsUnknown() {
		if value := res.Get("products.sensor.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsSensorLastUpgradeFromVersionShortName.IsNull() {
			data.ProductsSensorLastUpgradeFromVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSensorLastUpgradeFromVersionShortName = types.StringNull()
		}
	}
	if data.ProductsSensorLastUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.sensor.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSensorLastUpgradeToVersionFirmware.IsNull() {
			data.ProductsSensorLastUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSensorLastUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSensorLastUpgradeToVersionId.IsUnknown() {
		if value := res.Get("products.sensor.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsSensorLastUpgradeToVersionId.IsNull() {
			data.ProductsSensorLastUpgradeToVersionId = types.StringValue(value.String())
		} else {
			data.ProductsSensorLastUpgradeToVersionId = types.StringNull()
		}
	}
	if data.ProductsSensorLastUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.sensor.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSensorLastUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsSensorLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSensorLastUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSensorLastUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.sensor.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSensorLastUpgradeToVersionReleaseType.IsNull() {
			data.ProductsSensorLastUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSensorLastUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSensorLastUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.sensor.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSensorLastUpgradeToVersionShortName.IsNull() {
			data.ProductsSensorLastUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSensorLastUpgradeToVersionShortName = types.StringNull()
		}
	}
	if data.ProductsSensorNextUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.sensor.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSensorNextUpgradeToVersionFirmware.IsNull() {
			data.ProductsSensorNextUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSensorNextUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSensorNextUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.sensor.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSensorNextUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsSensorNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSensorNextUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSensorNextUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.sensor.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSensorNextUpgradeToVersionReleaseType.IsNull() {
			data.ProductsSensorNextUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSensorNextUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSensorNextUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.sensor.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSensorNextUpgradeToVersionShortName.IsNull() {
			data.ProductsSensorNextUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSensorNextUpgradeToVersionShortName = types.StringNull()
		}
	}
	for i := 0; i < len(data.ProductsSensorAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsSensorAvailableVersions[i].Firmware.ValueString(), data.ProductsSensorAvailableVersions[i].Id.ValueString(), data.ProductsSensorAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsSensorAvailableVersions[i].ReleaseType.ValueString(), data.ProductsSensorAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsSensorAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.sensor.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsSensorAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsSensorAvailableVersions[i],
			))
			(*parent).ProductsSensorAvailableVersions = slices.Delete((*parent).ProductsSensorAvailableVersions, i, i+1)
			i--

			continue
		}
		if data.Firmware.IsUnknown() {
			if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
		}
		if data.Id.IsUnknown() {
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
		}
		if data.ReleaseDate.IsUnknown() {
			if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
		}
		if data.ReleaseType.IsUnknown() {
			if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
		}
		if data.ShortName.IsUnknown() {
			if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
		}
		(*parent).ProductsSensorAvailableVersions[i] = data
	}
	if data.ProductsSwitchIsUpgradeAvailable.IsUnknown() {
		if value := res.Get("products.switch.isUpgradeAvailable"); value.Exists() && !data.ProductsSwitchIsUpgradeAvailable.IsNull() {
			data.ProductsSwitchIsUpgradeAvailable = types.BoolValue(value.Bool())
		} else {
			data.ProductsSwitchIsUpgradeAvailable = types.BoolNull()
		}
	}
	if data.ProductsSwitchCurrentVersionFirmware.IsUnknown() {
		if value := res.Get("products.switch.currentVersion.firmware"); value.Exists() && !data.ProductsSwitchCurrentVersionFirmware.IsNull() {
			data.ProductsSwitchCurrentVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSwitchCurrentVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSwitchCurrentVersionId.IsUnknown() {
		if value := res.Get("products.switch.currentVersion.id"); value.Exists() && !data.ProductsSwitchCurrentVersionId.IsNull() {
			data.ProductsSwitchCurrentVersionId = types.StringValue(value.String())
		} else {
			data.ProductsSwitchCurrentVersionId = types.StringNull()
		}
	}
	if data.ProductsSwitchCurrentVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.switch.currentVersion.releaseDate"); value.Exists() && !data.ProductsSwitchCurrentVersionReleaseDate.IsNull() {
			data.ProductsSwitchCurrentVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSwitchCurrentVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSwitchCurrentVersionReleaseType.IsUnknown() {
		if value := res.Get("products.switch.currentVersion.releaseType"); value.Exists() && !data.ProductsSwitchCurrentVersionReleaseType.IsNull() {
			data.ProductsSwitchCurrentVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSwitchCurrentVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSwitchCurrentVersionShortName.IsUnknown() {
		if value := res.Get("products.switch.currentVersion.shortName"); value.Exists() && !data.ProductsSwitchCurrentVersionShortName.IsNull() {
			data.ProductsSwitchCurrentVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSwitchCurrentVersionShortName = types.StringNull()
		}
	}
	if data.ProductsSwitchLastUpgradeTime.IsUnknown() {
		if value := res.Get("products.switch.lastUpgrade.time"); value.Exists() && !data.ProductsSwitchLastUpgradeTime.IsNull() {
			data.ProductsSwitchLastUpgradeTime = types.StringValue(value.String())
		} else {
			data.ProductsSwitchLastUpgradeTime = types.StringNull()
		}
	}
	if data.ProductsSwitchLastUpgradeFromVersionFirmware.IsUnknown() {
		if value := res.Get("products.switch.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsSwitchLastUpgradeFromVersionFirmware.IsNull() {
			data.ProductsSwitchLastUpgradeFromVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSwitchLastUpgradeFromVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSwitchLastUpgradeFromVersionId.IsUnknown() {
		if value := res.Get("products.switch.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsSwitchLastUpgradeFromVersionId.IsNull() {
			data.ProductsSwitchLastUpgradeFromVersionId = types.StringValue(value.String())
		} else {
			data.ProductsSwitchLastUpgradeFromVersionId = types.StringNull()
		}
	}
	if data.ProductsSwitchLastUpgradeFromVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.switch.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsSwitchLastUpgradeFromVersionReleaseDate.IsNull() {
			data.ProductsSwitchLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSwitchLastUpgradeFromVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSwitchLastUpgradeFromVersionReleaseType.IsUnknown() {
		if value := res.Get("products.switch.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsSwitchLastUpgradeFromVersionReleaseType.IsNull() {
			data.ProductsSwitchLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSwitchLastUpgradeFromVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSwitchLastUpgradeFromVersionShortName.IsUnknown() {
		if value := res.Get("products.switch.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsSwitchLastUpgradeFromVersionShortName.IsNull() {
			data.ProductsSwitchLastUpgradeFromVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSwitchLastUpgradeFromVersionShortName = types.StringNull()
		}
	}
	if data.ProductsSwitchLastUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.switch.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSwitchLastUpgradeToVersionFirmware.IsNull() {
			data.ProductsSwitchLastUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSwitchLastUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSwitchLastUpgradeToVersionId.IsUnknown() {
		if value := res.Get("products.switch.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsSwitchLastUpgradeToVersionId.IsNull() {
			data.ProductsSwitchLastUpgradeToVersionId = types.StringValue(value.String())
		} else {
			data.ProductsSwitchLastUpgradeToVersionId = types.StringNull()
		}
	}
	if data.ProductsSwitchLastUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.switch.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSwitchLastUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsSwitchLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSwitchLastUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSwitchLastUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.switch.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSwitchLastUpgradeToVersionReleaseType.IsNull() {
			data.ProductsSwitchLastUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSwitchLastUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSwitchLastUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.switch.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSwitchLastUpgradeToVersionShortName.IsNull() {
			data.ProductsSwitchLastUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSwitchLastUpgradeToVersionShortName = types.StringNull()
		}
	}
	if data.ProductsSwitchNextUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.switch.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsSwitchNextUpgradeToVersionFirmware.IsNull() {
			data.ProductsSwitchNextUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsSwitchNextUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsSwitchNextUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.switch.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsSwitchNextUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsSwitchNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsSwitchNextUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsSwitchNextUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.switch.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsSwitchNextUpgradeToVersionReleaseType.IsNull() {
			data.ProductsSwitchNextUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsSwitchNextUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsSwitchNextUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.switch.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsSwitchNextUpgradeToVersionShortName.IsNull() {
			data.ProductsSwitchNextUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsSwitchNextUpgradeToVersionShortName = types.StringNull()
		}
	}
	for i := 0; i < len(data.ProductsSwitchAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsSwitchAvailableVersions[i].Firmware.ValueString(), data.ProductsSwitchAvailableVersions[i].Id.ValueString(), data.ProductsSwitchAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsSwitchAvailableVersions[i].ReleaseType.ValueString(), data.ProductsSwitchAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsSwitchAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.switch.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsSwitchAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsSwitchAvailableVersions[i],
			))
			(*parent).ProductsSwitchAvailableVersions = slices.Delete((*parent).ProductsSwitchAvailableVersions, i, i+1)
			i--

			continue
		}
		if data.Firmware.IsUnknown() {
			if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
		}
		if data.Id.IsUnknown() {
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
		}
		if data.ReleaseDate.IsUnknown() {
			if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
		}
		if data.ReleaseType.IsUnknown() {
			if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
		}
		if data.ShortName.IsUnknown() {
			if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
		}
		(*parent).ProductsSwitchAvailableVersions[i] = data
	}
	if data.ProductsWirelessIsUpgradeAvailable.IsUnknown() {
		if value := res.Get("products.wireless.isUpgradeAvailable"); value.Exists() && !data.ProductsWirelessIsUpgradeAvailable.IsNull() {
			data.ProductsWirelessIsUpgradeAvailable = types.BoolValue(value.Bool())
		} else {
			data.ProductsWirelessIsUpgradeAvailable = types.BoolNull()
		}
	}
	if data.ProductsWirelessCurrentVersionFirmware.IsUnknown() {
		if value := res.Get("products.wireless.currentVersion.firmware"); value.Exists() && !data.ProductsWirelessCurrentVersionFirmware.IsNull() {
			data.ProductsWirelessCurrentVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsWirelessCurrentVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsWirelessCurrentVersionId.IsUnknown() {
		if value := res.Get("products.wireless.currentVersion.id"); value.Exists() && !data.ProductsWirelessCurrentVersionId.IsNull() {
			data.ProductsWirelessCurrentVersionId = types.StringValue(value.String())
		} else {
			data.ProductsWirelessCurrentVersionId = types.StringNull()
		}
	}
	if data.ProductsWirelessCurrentVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.wireless.currentVersion.releaseDate"); value.Exists() && !data.ProductsWirelessCurrentVersionReleaseDate.IsNull() {
			data.ProductsWirelessCurrentVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsWirelessCurrentVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsWirelessCurrentVersionReleaseType.IsUnknown() {
		if value := res.Get("products.wireless.currentVersion.releaseType"); value.Exists() && !data.ProductsWirelessCurrentVersionReleaseType.IsNull() {
			data.ProductsWirelessCurrentVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsWirelessCurrentVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsWirelessCurrentVersionShortName.IsUnknown() {
		if value := res.Get("products.wireless.currentVersion.shortName"); value.Exists() && !data.ProductsWirelessCurrentVersionShortName.IsNull() {
			data.ProductsWirelessCurrentVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsWirelessCurrentVersionShortName = types.StringNull()
		}
	}
	if data.ProductsWirelessLastUpgradeTime.IsUnknown() {
		if value := res.Get("products.wireless.lastUpgrade.time"); value.Exists() && !data.ProductsWirelessLastUpgradeTime.IsNull() {
			data.ProductsWirelessLastUpgradeTime = types.StringValue(value.String())
		} else {
			data.ProductsWirelessLastUpgradeTime = types.StringNull()
		}
	}
	if data.ProductsWirelessLastUpgradeFromVersionFirmware.IsUnknown() {
		if value := res.Get("products.wireless.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsWirelessLastUpgradeFromVersionFirmware.IsNull() {
			data.ProductsWirelessLastUpgradeFromVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsWirelessLastUpgradeFromVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsWirelessLastUpgradeFromVersionId.IsUnknown() {
		if value := res.Get("products.wireless.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsWirelessLastUpgradeFromVersionId.IsNull() {
			data.ProductsWirelessLastUpgradeFromVersionId = types.StringValue(value.String())
		} else {
			data.ProductsWirelessLastUpgradeFromVersionId = types.StringNull()
		}
	}
	if data.ProductsWirelessLastUpgradeFromVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.wireless.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsWirelessLastUpgradeFromVersionReleaseDate.IsNull() {
			data.ProductsWirelessLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsWirelessLastUpgradeFromVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsWirelessLastUpgradeFromVersionReleaseType.IsUnknown() {
		if value := res.Get("products.wireless.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsWirelessLastUpgradeFromVersionReleaseType.IsNull() {
			data.ProductsWirelessLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsWirelessLastUpgradeFromVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsWirelessLastUpgradeFromVersionShortName.IsUnknown() {
		if value := res.Get("products.wireless.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsWirelessLastUpgradeFromVersionShortName.IsNull() {
			data.ProductsWirelessLastUpgradeFromVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsWirelessLastUpgradeFromVersionShortName = types.StringNull()
		}
	}
	if data.ProductsWirelessLastUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.wireless.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsWirelessLastUpgradeToVersionFirmware.IsNull() {
			data.ProductsWirelessLastUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsWirelessLastUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsWirelessLastUpgradeToVersionId.IsUnknown() {
		if value := res.Get("products.wireless.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsWirelessLastUpgradeToVersionId.IsNull() {
			data.ProductsWirelessLastUpgradeToVersionId = types.StringValue(value.String())
		} else {
			data.ProductsWirelessLastUpgradeToVersionId = types.StringNull()
		}
	}
	if data.ProductsWirelessLastUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.wireless.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsWirelessLastUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsWirelessLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsWirelessLastUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsWirelessLastUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.wireless.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsWirelessLastUpgradeToVersionReleaseType.IsNull() {
			data.ProductsWirelessLastUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsWirelessLastUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsWirelessLastUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.wireless.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsWirelessLastUpgradeToVersionShortName.IsNull() {
			data.ProductsWirelessLastUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsWirelessLastUpgradeToVersionShortName = types.StringNull()
		}
	}
	if data.ProductsWirelessNextUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.wireless.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsWirelessNextUpgradeToVersionFirmware.IsNull() {
			data.ProductsWirelessNextUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsWirelessNextUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsWirelessNextUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.wireless.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsWirelessNextUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsWirelessNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsWirelessNextUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsWirelessNextUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.wireless.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsWirelessNextUpgradeToVersionReleaseType.IsNull() {
			data.ProductsWirelessNextUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsWirelessNextUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsWirelessNextUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.wireless.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsWirelessNextUpgradeToVersionShortName.IsNull() {
			data.ProductsWirelessNextUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsWirelessNextUpgradeToVersionShortName = types.StringNull()
		}
	}
	for i := 0; i < len(data.ProductsWirelessAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsWirelessAvailableVersions[i].Firmware.ValueString(), data.ProductsWirelessAvailableVersions[i].Id.ValueString(), data.ProductsWirelessAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsWirelessAvailableVersions[i].ReleaseType.ValueString(), data.ProductsWirelessAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsWirelessAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.wireless.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsWirelessAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsWirelessAvailableVersions[i],
			))
			(*parent).ProductsWirelessAvailableVersions = slices.Delete((*parent).ProductsWirelessAvailableVersions, i, i+1)
			i--

			continue
		}
		if data.Firmware.IsUnknown() {
			if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
		}
		if data.Id.IsUnknown() {
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
		}
		if data.ReleaseDate.IsUnknown() {
			if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
		}
		if data.ReleaseType.IsUnknown() {
			if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
		}
		if data.ShortName.IsUnknown() {
			if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
		}
		(*parent).ProductsWirelessAvailableVersions[i] = data
	}
	if data.ProductsWirelessControllerIsUpgradeAvailable.IsUnknown() {
		if value := res.Get("products.wirelessController.isUpgradeAvailable"); value.Exists() && !data.ProductsWirelessControllerIsUpgradeAvailable.IsNull() {
			data.ProductsWirelessControllerIsUpgradeAvailable = types.BoolValue(value.Bool())
		} else {
			data.ProductsWirelessControllerIsUpgradeAvailable = types.BoolNull()
		}
	}
	if data.ProductsWirelessControllerCurrentVersionFirmware.IsUnknown() {
		if value := res.Get("products.wirelessController.currentVersion.firmware"); value.Exists() && !data.ProductsWirelessControllerCurrentVersionFirmware.IsNull() {
			data.ProductsWirelessControllerCurrentVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerCurrentVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerCurrentVersionId.IsUnknown() {
		if value := res.Get("products.wirelessController.currentVersion.id"); value.Exists() && !data.ProductsWirelessControllerCurrentVersionId.IsNull() {
			data.ProductsWirelessControllerCurrentVersionId = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerCurrentVersionId = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerCurrentVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.wirelessController.currentVersion.releaseDate"); value.Exists() && !data.ProductsWirelessControllerCurrentVersionReleaseDate.IsNull() {
			data.ProductsWirelessControllerCurrentVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerCurrentVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerCurrentVersionReleaseType.IsUnknown() {
		if value := res.Get("products.wirelessController.currentVersion.releaseType"); value.Exists() && !data.ProductsWirelessControllerCurrentVersionReleaseType.IsNull() {
			data.ProductsWirelessControllerCurrentVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerCurrentVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerCurrentVersionShortName.IsUnknown() {
		if value := res.Get("products.wirelessController.currentVersion.shortName"); value.Exists() && !data.ProductsWirelessControllerCurrentVersionShortName.IsNull() {
			data.ProductsWirelessControllerCurrentVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerCurrentVersionShortName = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerLastUpgradeTime.IsUnknown() {
		if value := res.Get("products.wirelessController.lastUpgrade.time"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeTime.IsNull() {
			data.ProductsWirelessControllerLastUpgradeTime = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerLastUpgradeTime = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerLastUpgradeFromVersionFirmware.IsUnknown() {
		if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.firmware"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeFromVersionFirmware.IsNull() {
			data.ProductsWirelessControllerLastUpgradeFromVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerLastUpgradeFromVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerLastUpgradeFromVersionId.IsUnknown() {
		if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.id"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeFromVersionId.IsNull() {
			data.ProductsWirelessControllerLastUpgradeFromVersionId = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerLastUpgradeFromVersionId = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerLastUpgradeFromVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.releaseDate"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeFromVersionReleaseDate.IsNull() {
			data.ProductsWirelessControllerLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerLastUpgradeFromVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerLastUpgradeFromVersionReleaseType.IsUnknown() {
		if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.releaseType"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeFromVersionReleaseType.IsNull() {
			data.ProductsWirelessControllerLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerLastUpgradeFromVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerLastUpgradeFromVersionShortName.IsUnknown() {
		if value := res.Get("products.wirelessController.lastUpgrade.fromVersion.shortName"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeFromVersionShortName.IsNull() {
			data.ProductsWirelessControllerLastUpgradeFromVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerLastUpgradeFromVersionShortName = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerLastUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.wirelessController.lastUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeToVersionFirmware.IsNull() {
			data.ProductsWirelessControllerLastUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerLastUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerLastUpgradeToVersionId.IsUnknown() {
		if value := res.Get("products.wirelessController.lastUpgrade.toVersion.id"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeToVersionId.IsNull() {
			data.ProductsWirelessControllerLastUpgradeToVersionId = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerLastUpgradeToVersionId = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerLastUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.wirelessController.lastUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsWirelessControllerLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerLastUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerLastUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.wirelessController.lastUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeToVersionReleaseType.IsNull() {
			data.ProductsWirelessControllerLastUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerLastUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerLastUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.wirelessController.lastUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsWirelessControllerLastUpgradeToVersionShortName.IsNull() {
			data.ProductsWirelessControllerLastUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerLastUpgradeToVersionShortName = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerNextUpgradeToVersionFirmware.IsUnknown() {
		if value := res.Get("products.wirelessController.nextUpgrade.toVersion.firmware"); value.Exists() && !data.ProductsWirelessControllerNextUpgradeToVersionFirmware.IsNull() {
			data.ProductsWirelessControllerNextUpgradeToVersionFirmware = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerNextUpgradeToVersionFirmware = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerNextUpgradeToVersionReleaseDate.IsUnknown() {
		if value := res.Get("products.wirelessController.nextUpgrade.toVersion.releaseDate"); value.Exists() && !data.ProductsWirelessControllerNextUpgradeToVersionReleaseDate.IsNull() {
			data.ProductsWirelessControllerNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerNextUpgradeToVersionReleaseDate = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerNextUpgradeToVersionReleaseType.IsUnknown() {
		if value := res.Get("products.wirelessController.nextUpgrade.toVersion.releaseType"); value.Exists() && !data.ProductsWirelessControllerNextUpgradeToVersionReleaseType.IsNull() {
			data.ProductsWirelessControllerNextUpgradeToVersionReleaseType = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerNextUpgradeToVersionReleaseType = types.StringNull()
		}
	}
	if data.ProductsWirelessControllerNextUpgradeToVersionShortName.IsUnknown() {
		if value := res.Get("products.wirelessController.nextUpgrade.toVersion.shortName"); value.Exists() && !data.ProductsWirelessControllerNextUpgradeToVersionShortName.IsNull() {
			data.ProductsWirelessControllerNextUpgradeToVersionShortName = types.StringValue(value.String())
		} else {
			data.ProductsWirelessControllerNextUpgradeToVersionShortName = types.StringNull()
		}
	}
	for i := 0; i < len(data.ProductsWirelessControllerAvailableVersions); i++ {
		keys := [...]string{"firmware", "id", "releaseDate", "releaseType", "shortName"}
		keyValues := [...]string{data.ProductsWirelessControllerAvailableVersions[i].Firmware.ValueString(), data.ProductsWirelessControllerAvailableVersions[i].Id.ValueString(), data.ProductsWirelessControllerAvailableVersions[i].ReleaseDate.ValueString(), data.ProductsWirelessControllerAvailableVersions[i].ReleaseType.ValueString(), data.ProductsWirelessControllerAvailableVersions[i].ShortName.ValueString()}

		parent := &data
		data := (*parent).ProductsWirelessControllerAvailableVersions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("products.wirelessController.availableVersions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ProductsWirelessControllerAvailableVersions[%d] = %+v",
				i,
				(*parent).ProductsWirelessControllerAvailableVersions[i],
			))
			(*parent).ProductsWirelessControllerAvailableVersions = slices.Delete((*parent).ProductsWirelessControllerAvailableVersions, i, i+1)
			i--

			continue
		}
		if data.Firmware.IsUnknown() {
			if value := res.Get("firmware"); value.Exists() && !data.Firmware.IsNull() {
				data.Firmware = types.StringValue(value.String())
			} else {
				data.Firmware = types.StringNull()
			}
		}
		if data.Id.IsUnknown() {
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
		}
		if data.ReleaseDate.IsUnknown() {
			if value := res.Get("releaseDate"); value.Exists() && !data.ReleaseDate.IsNull() {
				data.ReleaseDate = types.StringValue(value.String())
			} else {
				data.ReleaseDate = types.StringNull()
			}
		}
		if data.ReleaseType.IsUnknown() {
			if value := res.Get("releaseType"); value.Exists() && !data.ReleaseType.IsNull() {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
		}
		if data.ShortName.IsUnknown() {
			if value := res.Get("shortName"); value.Exists() && !data.ShortName.IsNull() {
				data.ShortName = types.StringValue(value.String())
			} else {
				data.ShortName = types.StringNull()
			}
		}
		(*parent).ProductsWirelessControllerAvailableVersions[i] = data
	}
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
