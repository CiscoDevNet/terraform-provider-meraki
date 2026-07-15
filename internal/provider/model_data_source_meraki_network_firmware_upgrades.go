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
// It emits a separate model - DataSourceNetworkFirmwareUpgrades - used only by the data source, always including every
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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceNetworkFirmwareUpgrades struct {
	Id                                                          types.String                                                                   `tfsdk:"id"`
	NetworkId                                                   types.String                                                                   `tfsdk:"network_id"`
	Timezone                                                    types.String                                                                   `tfsdk:"timezone"`
	ProductsApplianceIsUpgradeAvailable                         types.Bool                                                                     `tfsdk:"products_appliance_is_upgrade_available"`
	ProductsApplianceParticipateInNextBetaRelease               types.Bool                                                                     `tfsdk:"products_appliance_participate_in_next_beta_release"`
	ProductsApplianceCurrentVersionFirmware                     types.String                                                                   `tfsdk:"products_appliance_current_version_firmware"`
	ProductsApplianceCurrentVersionId                           types.String                                                                   `tfsdk:"products_appliance_current_version_id"`
	ProductsApplianceCurrentVersionReleaseDate                  types.String                                                                   `tfsdk:"products_appliance_current_version_release_date"`
	ProductsApplianceCurrentVersionReleaseType                  types.String                                                                   `tfsdk:"products_appliance_current_version_release_type"`
	ProductsApplianceCurrentVersionShortName                    types.String                                                                   `tfsdk:"products_appliance_current_version_short_name"`
	ProductsApplianceLastUpgradeTime                            types.String                                                                   `tfsdk:"products_appliance_last_upgrade_time"`
	ProductsApplianceLastUpgradeFromVersionFirmware             types.String                                                                   `tfsdk:"products_appliance_last_upgrade_from_version_firmware"`
	ProductsApplianceLastUpgradeFromVersionId                   types.String                                                                   `tfsdk:"products_appliance_last_upgrade_from_version_id"`
	ProductsApplianceLastUpgradeFromVersionReleaseDate          types.String                                                                   `tfsdk:"products_appliance_last_upgrade_from_version_release_date"`
	ProductsApplianceLastUpgradeFromVersionReleaseType          types.String                                                                   `tfsdk:"products_appliance_last_upgrade_from_version_release_type"`
	ProductsApplianceLastUpgradeFromVersionShortName            types.String                                                                   `tfsdk:"products_appliance_last_upgrade_from_version_short_name"`
	ProductsApplianceLastUpgradeToVersionFirmware               types.String                                                                   `tfsdk:"products_appliance_last_upgrade_to_version_firmware"`
	ProductsApplianceLastUpgradeToVersionId                     types.String                                                                   `tfsdk:"products_appliance_last_upgrade_to_version_id"`
	ProductsApplianceLastUpgradeToVersionReleaseDate            types.String                                                                   `tfsdk:"products_appliance_last_upgrade_to_version_release_date"`
	ProductsApplianceLastUpgradeToVersionReleaseType            types.String                                                                   `tfsdk:"products_appliance_last_upgrade_to_version_release_type"`
	ProductsApplianceLastUpgradeToVersionShortName              types.String                                                                   `tfsdk:"products_appliance_last_upgrade_to_version_short_name"`
	ProductsApplianceNextUpgradeTime                            types.String                                                                   `tfsdk:"products_appliance_next_upgrade_time"`
	ProductsApplianceNextUpgradeToVersionId                     types.String                                                                   `tfsdk:"products_appliance_next_upgrade_to_version_id"`
	ProductsApplianceNextUpgradeToVersionFirmware               types.String                                                                   `tfsdk:"products_appliance_next_upgrade_to_version_firmware"`
	ProductsApplianceNextUpgradeToVersionReleaseDate            types.String                                                                   `tfsdk:"products_appliance_next_upgrade_to_version_release_date"`
	ProductsApplianceNextUpgradeToVersionReleaseType            types.String                                                                   `tfsdk:"products_appliance_next_upgrade_to_version_release_type"`
	ProductsApplianceNextUpgradeToVersionShortName              types.String                                                                   `tfsdk:"products_appliance_next_upgrade_to_version_short_name"`
	ProductsApplianceAvailableVersions                          []DataSourceNetworkFirmwareUpgradesProductsApplianceAvailableVersions          `tfsdk:"products_appliance_available_versions"`
	ProductsCameraIsUpgradeAvailable                            types.Bool                                                                     `tfsdk:"products_camera_is_upgrade_available"`
	ProductsCameraParticipateInNextBetaRelease                  types.Bool                                                                     `tfsdk:"products_camera_participate_in_next_beta_release"`
	ProductsCameraCurrentVersionFirmware                        types.String                                                                   `tfsdk:"products_camera_current_version_firmware"`
	ProductsCameraCurrentVersionId                              types.String                                                                   `tfsdk:"products_camera_current_version_id"`
	ProductsCameraCurrentVersionReleaseDate                     types.String                                                                   `tfsdk:"products_camera_current_version_release_date"`
	ProductsCameraCurrentVersionReleaseType                     types.String                                                                   `tfsdk:"products_camera_current_version_release_type"`
	ProductsCameraCurrentVersionShortName                       types.String                                                                   `tfsdk:"products_camera_current_version_short_name"`
	ProductsCameraLastUpgradeTime                               types.String                                                                   `tfsdk:"products_camera_last_upgrade_time"`
	ProductsCameraLastUpgradeFromVersionFirmware                types.String                                                                   `tfsdk:"products_camera_last_upgrade_from_version_firmware"`
	ProductsCameraLastUpgradeFromVersionId                      types.String                                                                   `tfsdk:"products_camera_last_upgrade_from_version_id"`
	ProductsCameraLastUpgradeFromVersionReleaseDate             types.String                                                                   `tfsdk:"products_camera_last_upgrade_from_version_release_date"`
	ProductsCameraLastUpgradeFromVersionReleaseType             types.String                                                                   `tfsdk:"products_camera_last_upgrade_from_version_release_type"`
	ProductsCameraLastUpgradeFromVersionShortName               types.String                                                                   `tfsdk:"products_camera_last_upgrade_from_version_short_name"`
	ProductsCameraLastUpgradeToVersionFirmware                  types.String                                                                   `tfsdk:"products_camera_last_upgrade_to_version_firmware"`
	ProductsCameraLastUpgradeToVersionId                        types.String                                                                   `tfsdk:"products_camera_last_upgrade_to_version_id"`
	ProductsCameraLastUpgradeToVersionReleaseDate               types.String                                                                   `tfsdk:"products_camera_last_upgrade_to_version_release_date"`
	ProductsCameraLastUpgradeToVersionReleaseType               types.String                                                                   `tfsdk:"products_camera_last_upgrade_to_version_release_type"`
	ProductsCameraLastUpgradeToVersionShortName                 types.String                                                                   `tfsdk:"products_camera_last_upgrade_to_version_short_name"`
	ProductsCameraNextUpgradeTime                               types.String                                                                   `tfsdk:"products_camera_next_upgrade_time"`
	ProductsCameraNextUpgradeToVersionId                        types.String                                                                   `tfsdk:"products_camera_next_upgrade_to_version_id"`
	ProductsCameraNextUpgradeToVersionFirmware                  types.String                                                                   `tfsdk:"products_camera_next_upgrade_to_version_firmware"`
	ProductsCameraNextUpgradeToVersionReleaseDate               types.String                                                                   `tfsdk:"products_camera_next_upgrade_to_version_release_date"`
	ProductsCameraNextUpgradeToVersionReleaseType               types.String                                                                   `tfsdk:"products_camera_next_upgrade_to_version_release_type"`
	ProductsCameraNextUpgradeToVersionShortName                 types.String                                                                   `tfsdk:"products_camera_next_upgrade_to_version_short_name"`
	ProductsCameraAvailableVersions                             []DataSourceNetworkFirmwareUpgradesProductsCameraAvailableVersions             `tfsdk:"products_camera_available_versions"`
	ProductsCellularGatewayIsUpgradeAvailable                   types.Bool                                                                     `tfsdk:"products_cellular_gateway_is_upgrade_available"`
	ProductsCellularGatewayParticipateInNextBetaRelease         types.Bool                                                                     `tfsdk:"products_cellular_gateway_participate_in_next_beta_release"`
	ProductsCellularGatewayCurrentVersionFirmware               types.String                                                                   `tfsdk:"products_cellular_gateway_current_version_firmware"`
	ProductsCellularGatewayCurrentVersionId                     types.String                                                                   `tfsdk:"products_cellular_gateway_current_version_id"`
	ProductsCellularGatewayCurrentVersionReleaseDate            types.String                                                                   `tfsdk:"products_cellular_gateway_current_version_release_date"`
	ProductsCellularGatewayCurrentVersionReleaseType            types.String                                                                   `tfsdk:"products_cellular_gateway_current_version_release_type"`
	ProductsCellularGatewayCurrentVersionShortName              types.String                                                                   `tfsdk:"products_cellular_gateway_current_version_short_name"`
	ProductsCellularGatewayLastUpgradeTime                      types.String                                                                   `tfsdk:"products_cellular_gateway_last_upgrade_time"`
	ProductsCellularGatewayLastUpgradeFromVersionFirmware       types.String                                                                   `tfsdk:"products_cellular_gateway_last_upgrade_from_version_firmware"`
	ProductsCellularGatewayLastUpgradeFromVersionId             types.String                                                                   `tfsdk:"products_cellular_gateway_last_upgrade_from_version_id"`
	ProductsCellularGatewayLastUpgradeFromVersionReleaseDate    types.String                                                                   `tfsdk:"products_cellular_gateway_last_upgrade_from_version_release_date"`
	ProductsCellularGatewayLastUpgradeFromVersionReleaseType    types.String                                                                   `tfsdk:"products_cellular_gateway_last_upgrade_from_version_release_type"`
	ProductsCellularGatewayLastUpgradeFromVersionShortName      types.String                                                                   `tfsdk:"products_cellular_gateway_last_upgrade_from_version_short_name"`
	ProductsCellularGatewayLastUpgradeToVersionFirmware         types.String                                                                   `tfsdk:"products_cellular_gateway_last_upgrade_to_version_firmware"`
	ProductsCellularGatewayLastUpgradeToVersionId               types.String                                                                   `tfsdk:"products_cellular_gateway_last_upgrade_to_version_id"`
	ProductsCellularGatewayLastUpgradeToVersionReleaseDate      types.String                                                                   `tfsdk:"products_cellular_gateway_last_upgrade_to_version_release_date"`
	ProductsCellularGatewayLastUpgradeToVersionReleaseType      types.String                                                                   `tfsdk:"products_cellular_gateway_last_upgrade_to_version_release_type"`
	ProductsCellularGatewayLastUpgradeToVersionShortName        types.String                                                                   `tfsdk:"products_cellular_gateway_last_upgrade_to_version_short_name"`
	ProductsCellularGatewayNextUpgradeTime                      types.String                                                                   `tfsdk:"products_cellular_gateway_next_upgrade_time"`
	ProductsCellularGatewayNextUpgradeToVersionId               types.String                                                                   `tfsdk:"products_cellular_gateway_next_upgrade_to_version_id"`
	ProductsCellularGatewayNextUpgradeToVersionFirmware         types.String                                                                   `tfsdk:"products_cellular_gateway_next_upgrade_to_version_firmware"`
	ProductsCellularGatewayNextUpgradeToVersionReleaseDate      types.String                                                                   `tfsdk:"products_cellular_gateway_next_upgrade_to_version_release_date"`
	ProductsCellularGatewayNextUpgradeToVersionReleaseType      types.String                                                                   `tfsdk:"products_cellular_gateway_next_upgrade_to_version_release_type"`
	ProductsCellularGatewayNextUpgradeToVersionShortName        types.String                                                                   `tfsdk:"products_cellular_gateway_next_upgrade_to_version_short_name"`
	ProductsCellularGatewayAvailableVersions                    []DataSourceNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions    `tfsdk:"products_cellular_gateway_available_versions"`
	ProductsSecureConnectIsUpgradeAvailable                     types.Bool                                                                     `tfsdk:"products_secure_connect_is_upgrade_available"`
	ProductsSecureConnectParticipateInNextBetaRelease           types.Bool                                                                     `tfsdk:"products_secure_connect_participate_in_next_beta_release"`
	ProductsSecureConnectCurrentVersionFirmware                 types.String                                                                   `tfsdk:"products_secure_connect_current_version_firmware"`
	ProductsSecureConnectCurrentVersionId                       types.String                                                                   `tfsdk:"products_secure_connect_current_version_id"`
	ProductsSecureConnectCurrentVersionReleaseDate              types.String                                                                   `tfsdk:"products_secure_connect_current_version_release_date"`
	ProductsSecureConnectCurrentVersionReleaseType              types.String                                                                   `tfsdk:"products_secure_connect_current_version_release_type"`
	ProductsSecureConnectCurrentVersionShortName                types.String                                                                   `tfsdk:"products_secure_connect_current_version_short_name"`
	ProductsSecureConnectLastUpgradeTime                        types.String                                                                   `tfsdk:"products_secure_connect_last_upgrade_time"`
	ProductsSecureConnectLastUpgradeFromVersionFirmware         types.String                                                                   `tfsdk:"products_secure_connect_last_upgrade_from_version_firmware"`
	ProductsSecureConnectLastUpgradeFromVersionId               types.String                                                                   `tfsdk:"products_secure_connect_last_upgrade_from_version_id"`
	ProductsSecureConnectLastUpgradeFromVersionReleaseDate      types.String                                                                   `tfsdk:"products_secure_connect_last_upgrade_from_version_release_date"`
	ProductsSecureConnectLastUpgradeFromVersionReleaseType      types.String                                                                   `tfsdk:"products_secure_connect_last_upgrade_from_version_release_type"`
	ProductsSecureConnectLastUpgradeFromVersionShortName        types.String                                                                   `tfsdk:"products_secure_connect_last_upgrade_from_version_short_name"`
	ProductsSecureConnectLastUpgradeToVersionFirmware           types.String                                                                   `tfsdk:"products_secure_connect_last_upgrade_to_version_firmware"`
	ProductsSecureConnectLastUpgradeToVersionId                 types.String                                                                   `tfsdk:"products_secure_connect_last_upgrade_to_version_id"`
	ProductsSecureConnectLastUpgradeToVersionReleaseDate        types.String                                                                   `tfsdk:"products_secure_connect_last_upgrade_to_version_release_date"`
	ProductsSecureConnectLastUpgradeToVersionReleaseType        types.String                                                                   `tfsdk:"products_secure_connect_last_upgrade_to_version_release_type"`
	ProductsSecureConnectLastUpgradeToVersionShortName          types.String                                                                   `tfsdk:"products_secure_connect_last_upgrade_to_version_short_name"`
	ProductsSecureConnectNextUpgradeTime                        types.String                                                                   `tfsdk:"products_secure_connect_next_upgrade_time"`
	ProductsSecureConnectNextUpgradeToVersionId                 types.String                                                                   `tfsdk:"products_secure_connect_next_upgrade_to_version_id"`
	ProductsSecureConnectNextUpgradeToVersionFirmware           types.String                                                                   `tfsdk:"products_secure_connect_next_upgrade_to_version_firmware"`
	ProductsSecureConnectNextUpgradeToVersionReleaseDate        types.String                                                                   `tfsdk:"products_secure_connect_next_upgrade_to_version_release_date"`
	ProductsSecureConnectNextUpgradeToVersionReleaseType        types.String                                                                   `tfsdk:"products_secure_connect_next_upgrade_to_version_release_type"`
	ProductsSecureConnectNextUpgradeToVersionShortName          types.String                                                                   `tfsdk:"products_secure_connect_next_upgrade_to_version_short_name"`
	ProductsSecureConnectAvailableVersions                      []DataSourceNetworkFirmwareUpgradesProductsSecureConnectAvailableVersions      `tfsdk:"products_secure_connect_available_versions"`
	ProductsSensorIsUpgradeAvailable                            types.Bool                                                                     `tfsdk:"products_sensor_is_upgrade_available"`
	ProductsSensorParticipateInNextBetaRelease                  types.Bool                                                                     `tfsdk:"products_sensor_participate_in_next_beta_release"`
	ProductsSensorCurrentVersionFirmware                        types.String                                                                   `tfsdk:"products_sensor_current_version_firmware"`
	ProductsSensorCurrentVersionId                              types.String                                                                   `tfsdk:"products_sensor_current_version_id"`
	ProductsSensorCurrentVersionReleaseDate                     types.String                                                                   `tfsdk:"products_sensor_current_version_release_date"`
	ProductsSensorCurrentVersionReleaseType                     types.String                                                                   `tfsdk:"products_sensor_current_version_release_type"`
	ProductsSensorCurrentVersionShortName                       types.String                                                                   `tfsdk:"products_sensor_current_version_short_name"`
	ProductsSensorLastUpgradeTime                               types.String                                                                   `tfsdk:"products_sensor_last_upgrade_time"`
	ProductsSensorLastUpgradeFromVersionFirmware                types.String                                                                   `tfsdk:"products_sensor_last_upgrade_from_version_firmware"`
	ProductsSensorLastUpgradeFromVersionId                      types.String                                                                   `tfsdk:"products_sensor_last_upgrade_from_version_id"`
	ProductsSensorLastUpgradeFromVersionReleaseDate             types.String                                                                   `tfsdk:"products_sensor_last_upgrade_from_version_release_date"`
	ProductsSensorLastUpgradeFromVersionReleaseType             types.String                                                                   `tfsdk:"products_sensor_last_upgrade_from_version_release_type"`
	ProductsSensorLastUpgradeFromVersionShortName               types.String                                                                   `tfsdk:"products_sensor_last_upgrade_from_version_short_name"`
	ProductsSensorLastUpgradeToVersionFirmware                  types.String                                                                   `tfsdk:"products_sensor_last_upgrade_to_version_firmware"`
	ProductsSensorLastUpgradeToVersionId                        types.String                                                                   `tfsdk:"products_sensor_last_upgrade_to_version_id"`
	ProductsSensorLastUpgradeToVersionReleaseDate               types.String                                                                   `tfsdk:"products_sensor_last_upgrade_to_version_release_date"`
	ProductsSensorLastUpgradeToVersionReleaseType               types.String                                                                   `tfsdk:"products_sensor_last_upgrade_to_version_release_type"`
	ProductsSensorLastUpgradeToVersionShortName                 types.String                                                                   `tfsdk:"products_sensor_last_upgrade_to_version_short_name"`
	ProductsSensorNextUpgradeTime                               types.String                                                                   `tfsdk:"products_sensor_next_upgrade_time"`
	ProductsSensorNextUpgradeToVersionId                        types.String                                                                   `tfsdk:"products_sensor_next_upgrade_to_version_id"`
	ProductsSensorNextUpgradeToVersionFirmware                  types.String                                                                   `tfsdk:"products_sensor_next_upgrade_to_version_firmware"`
	ProductsSensorNextUpgradeToVersionReleaseDate               types.String                                                                   `tfsdk:"products_sensor_next_upgrade_to_version_release_date"`
	ProductsSensorNextUpgradeToVersionReleaseType               types.String                                                                   `tfsdk:"products_sensor_next_upgrade_to_version_release_type"`
	ProductsSensorNextUpgradeToVersionShortName                 types.String                                                                   `tfsdk:"products_sensor_next_upgrade_to_version_short_name"`
	ProductsSensorAvailableVersions                             []DataSourceNetworkFirmwareUpgradesProductsSensorAvailableVersions             `tfsdk:"products_sensor_available_versions"`
	ProductsSwitchIsUpgradeAvailable                            types.Bool                                                                     `tfsdk:"products_switch_is_upgrade_available"`
	ProductsSwitchParticipateInNextBetaRelease                  types.Bool                                                                     `tfsdk:"products_switch_participate_in_next_beta_release"`
	ProductsSwitchCurrentVersionFirmware                        types.String                                                                   `tfsdk:"products_switch_current_version_firmware"`
	ProductsSwitchCurrentVersionId                              types.String                                                                   `tfsdk:"products_switch_current_version_id"`
	ProductsSwitchCurrentVersionReleaseDate                     types.String                                                                   `tfsdk:"products_switch_current_version_release_date"`
	ProductsSwitchCurrentVersionReleaseType                     types.String                                                                   `tfsdk:"products_switch_current_version_release_type"`
	ProductsSwitchCurrentVersionShortName                       types.String                                                                   `tfsdk:"products_switch_current_version_short_name"`
	ProductsSwitchLastUpgradeTime                               types.String                                                                   `tfsdk:"products_switch_last_upgrade_time"`
	ProductsSwitchLastUpgradeFromVersionFirmware                types.String                                                                   `tfsdk:"products_switch_last_upgrade_from_version_firmware"`
	ProductsSwitchLastUpgradeFromVersionId                      types.String                                                                   `tfsdk:"products_switch_last_upgrade_from_version_id"`
	ProductsSwitchLastUpgradeFromVersionReleaseDate             types.String                                                                   `tfsdk:"products_switch_last_upgrade_from_version_release_date"`
	ProductsSwitchLastUpgradeFromVersionReleaseType             types.String                                                                   `tfsdk:"products_switch_last_upgrade_from_version_release_type"`
	ProductsSwitchLastUpgradeFromVersionShortName               types.String                                                                   `tfsdk:"products_switch_last_upgrade_from_version_short_name"`
	ProductsSwitchLastUpgradeToVersionFirmware                  types.String                                                                   `tfsdk:"products_switch_last_upgrade_to_version_firmware"`
	ProductsSwitchLastUpgradeToVersionId                        types.String                                                                   `tfsdk:"products_switch_last_upgrade_to_version_id"`
	ProductsSwitchLastUpgradeToVersionReleaseDate               types.String                                                                   `tfsdk:"products_switch_last_upgrade_to_version_release_date"`
	ProductsSwitchLastUpgradeToVersionReleaseType               types.String                                                                   `tfsdk:"products_switch_last_upgrade_to_version_release_type"`
	ProductsSwitchLastUpgradeToVersionShortName                 types.String                                                                   `tfsdk:"products_switch_last_upgrade_to_version_short_name"`
	ProductsSwitchNextUpgradeTime                               types.String                                                                   `tfsdk:"products_switch_next_upgrade_time"`
	ProductsSwitchNextUpgradeToVersionId                        types.String                                                                   `tfsdk:"products_switch_next_upgrade_to_version_id"`
	ProductsSwitchNextUpgradeToVersionFirmware                  types.String                                                                   `tfsdk:"products_switch_next_upgrade_to_version_firmware"`
	ProductsSwitchNextUpgradeToVersionReleaseDate               types.String                                                                   `tfsdk:"products_switch_next_upgrade_to_version_release_date"`
	ProductsSwitchNextUpgradeToVersionReleaseType               types.String                                                                   `tfsdk:"products_switch_next_upgrade_to_version_release_type"`
	ProductsSwitchNextUpgradeToVersionShortName                 types.String                                                                   `tfsdk:"products_switch_next_upgrade_to_version_short_name"`
	ProductsSwitchAvailableVersions                             []DataSourceNetworkFirmwareUpgradesProductsSwitchAvailableVersions             `tfsdk:"products_switch_available_versions"`
	ProductsSwitchCatalystIsUpgradeAvailable                    types.Bool                                                                     `tfsdk:"products_switch_catalyst_is_upgrade_available"`
	ProductsSwitchCatalystParticipateInNextBetaRelease          types.Bool                                                                     `tfsdk:"products_switch_catalyst_participate_in_next_beta_release"`
	ProductsSwitchCatalystCurrentVersionFirmware                types.String                                                                   `tfsdk:"products_switch_catalyst_current_version_firmware"`
	ProductsSwitchCatalystCurrentVersionId                      types.String                                                                   `tfsdk:"products_switch_catalyst_current_version_id"`
	ProductsSwitchCatalystCurrentVersionReleaseDate             types.String                                                                   `tfsdk:"products_switch_catalyst_current_version_release_date"`
	ProductsSwitchCatalystCurrentVersionReleaseType             types.String                                                                   `tfsdk:"products_switch_catalyst_current_version_release_type"`
	ProductsSwitchCatalystCurrentVersionShortName               types.String                                                                   `tfsdk:"products_switch_catalyst_current_version_short_name"`
	ProductsSwitchCatalystLastUpgradeTime                       types.String                                                                   `tfsdk:"products_switch_catalyst_last_upgrade_time"`
	ProductsSwitchCatalystLastUpgradeFromVersionFirmware        types.String                                                                   `tfsdk:"products_switch_catalyst_last_upgrade_from_version_firmware"`
	ProductsSwitchCatalystLastUpgradeFromVersionId              types.String                                                                   `tfsdk:"products_switch_catalyst_last_upgrade_from_version_id"`
	ProductsSwitchCatalystLastUpgradeFromVersionReleaseDate     types.String                                                                   `tfsdk:"products_switch_catalyst_last_upgrade_from_version_release_date"`
	ProductsSwitchCatalystLastUpgradeFromVersionReleaseType     types.String                                                                   `tfsdk:"products_switch_catalyst_last_upgrade_from_version_release_type"`
	ProductsSwitchCatalystLastUpgradeFromVersionShortName       types.String                                                                   `tfsdk:"products_switch_catalyst_last_upgrade_from_version_short_name"`
	ProductsSwitchCatalystLastUpgradeToVersionFirmware          types.String                                                                   `tfsdk:"products_switch_catalyst_last_upgrade_to_version_firmware"`
	ProductsSwitchCatalystLastUpgradeToVersionId                types.String                                                                   `tfsdk:"products_switch_catalyst_last_upgrade_to_version_id"`
	ProductsSwitchCatalystLastUpgradeToVersionReleaseDate       types.String                                                                   `tfsdk:"products_switch_catalyst_last_upgrade_to_version_release_date"`
	ProductsSwitchCatalystLastUpgradeToVersionReleaseType       types.String                                                                   `tfsdk:"products_switch_catalyst_last_upgrade_to_version_release_type"`
	ProductsSwitchCatalystLastUpgradeToVersionShortName         types.String                                                                   `tfsdk:"products_switch_catalyst_last_upgrade_to_version_short_name"`
	ProductsSwitchCatalystNextUpgradeTime                       types.String                                                                   `tfsdk:"products_switch_catalyst_next_upgrade_time"`
	ProductsSwitchCatalystNextUpgradeToVersionId                types.String                                                                   `tfsdk:"products_switch_catalyst_next_upgrade_to_version_id"`
	ProductsSwitchCatalystNextUpgradeToVersionFirmware          types.String                                                                   `tfsdk:"products_switch_catalyst_next_upgrade_to_version_firmware"`
	ProductsSwitchCatalystNextUpgradeToVersionReleaseDate       types.String                                                                   `tfsdk:"products_switch_catalyst_next_upgrade_to_version_release_date"`
	ProductsSwitchCatalystNextUpgradeToVersionReleaseType       types.String                                                                   `tfsdk:"products_switch_catalyst_next_upgrade_to_version_release_type"`
	ProductsSwitchCatalystNextUpgradeToVersionShortName         types.String                                                                   `tfsdk:"products_switch_catalyst_next_upgrade_to_version_short_name"`
	ProductsSwitchCatalystAvailableVersions                     []DataSourceNetworkFirmwareUpgradesProductsSwitchCatalystAvailableVersions     `tfsdk:"products_switch_catalyst_available_versions"`
	ProductsWirelessIsUpgradeAvailable                          types.Bool                                                                     `tfsdk:"products_wireless_is_upgrade_available"`
	ProductsWirelessParticipateInNextBetaRelease                types.Bool                                                                     `tfsdk:"products_wireless_participate_in_next_beta_release"`
	ProductsWirelessCurrentVersionFirmware                      types.String                                                                   `tfsdk:"products_wireless_current_version_firmware"`
	ProductsWirelessCurrentVersionId                            types.String                                                                   `tfsdk:"products_wireless_current_version_id"`
	ProductsWirelessCurrentVersionReleaseDate                   types.String                                                                   `tfsdk:"products_wireless_current_version_release_date"`
	ProductsWirelessCurrentVersionReleaseType                   types.String                                                                   `tfsdk:"products_wireless_current_version_release_type"`
	ProductsWirelessCurrentVersionShortName                     types.String                                                                   `tfsdk:"products_wireless_current_version_short_name"`
	ProductsWirelessLastUpgradeTime                             types.String                                                                   `tfsdk:"products_wireless_last_upgrade_time"`
	ProductsWirelessLastUpgradeFromVersionFirmware              types.String                                                                   `tfsdk:"products_wireless_last_upgrade_from_version_firmware"`
	ProductsWirelessLastUpgradeFromVersionId                    types.String                                                                   `tfsdk:"products_wireless_last_upgrade_from_version_id"`
	ProductsWirelessLastUpgradeFromVersionReleaseDate           types.String                                                                   `tfsdk:"products_wireless_last_upgrade_from_version_release_date"`
	ProductsWirelessLastUpgradeFromVersionReleaseType           types.String                                                                   `tfsdk:"products_wireless_last_upgrade_from_version_release_type"`
	ProductsWirelessLastUpgradeFromVersionShortName             types.String                                                                   `tfsdk:"products_wireless_last_upgrade_from_version_short_name"`
	ProductsWirelessLastUpgradeToVersionFirmware                types.String                                                                   `tfsdk:"products_wireless_last_upgrade_to_version_firmware"`
	ProductsWirelessLastUpgradeToVersionId                      types.String                                                                   `tfsdk:"products_wireless_last_upgrade_to_version_id"`
	ProductsWirelessLastUpgradeToVersionReleaseDate             types.String                                                                   `tfsdk:"products_wireless_last_upgrade_to_version_release_date"`
	ProductsWirelessLastUpgradeToVersionReleaseType             types.String                                                                   `tfsdk:"products_wireless_last_upgrade_to_version_release_type"`
	ProductsWirelessLastUpgradeToVersionShortName               types.String                                                                   `tfsdk:"products_wireless_last_upgrade_to_version_short_name"`
	ProductsWirelessNextUpgradeTime                             types.String                                                                   `tfsdk:"products_wireless_next_upgrade_time"`
	ProductsWirelessNextUpgradePredownloadEnabled               types.Bool                                                                     `tfsdk:"products_wireless_next_upgrade_predownload_enabled"`
	ProductsWirelessNextUpgradeToVersionId                      types.String                                                                   `tfsdk:"products_wireless_next_upgrade_to_version_id"`
	ProductsWirelessNextUpgradeToVersionFirmware                types.String                                                                   `tfsdk:"products_wireless_next_upgrade_to_version_firmware"`
	ProductsWirelessNextUpgradeToVersionReleaseDate             types.String                                                                   `tfsdk:"products_wireless_next_upgrade_to_version_release_date"`
	ProductsWirelessNextUpgradeToVersionReleaseType             types.String                                                                   `tfsdk:"products_wireless_next_upgrade_to_version_release_type"`
	ProductsWirelessNextUpgradeToVersionShortName               types.String                                                                   `tfsdk:"products_wireless_next_upgrade_to_version_short_name"`
	ProductsWirelessAvailableVersions                           []DataSourceNetworkFirmwareUpgradesProductsWirelessAvailableVersions           `tfsdk:"products_wireless_available_versions"`
	ProductsWirelessControllerIsUpgradeAvailable                types.Bool                                                                     `tfsdk:"products_wireless_controller_is_upgrade_available"`
	ProductsWirelessControllerParticipateInNextBetaRelease      types.Bool                                                                     `tfsdk:"products_wireless_controller_participate_in_next_beta_release"`
	ProductsWirelessControllerCurrentVersionFirmware            types.String                                                                   `tfsdk:"products_wireless_controller_current_version_firmware"`
	ProductsWirelessControllerCurrentVersionId                  types.String                                                                   `tfsdk:"products_wireless_controller_current_version_id"`
	ProductsWirelessControllerCurrentVersionReleaseDate         types.String                                                                   `tfsdk:"products_wireless_controller_current_version_release_date"`
	ProductsWirelessControllerCurrentVersionReleaseType         types.String                                                                   `tfsdk:"products_wireless_controller_current_version_release_type"`
	ProductsWirelessControllerCurrentVersionShortName           types.String                                                                   `tfsdk:"products_wireless_controller_current_version_short_name"`
	ProductsWirelessControllerLastUpgradeTime                   types.String                                                                   `tfsdk:"products_wireless_controller_last_upgrade_time"`
	ProductsWirelessControllerLastUpgradeFromVersionFirmware    types.String                                                                   `tfsdk:"products_wireless_controller_last_upgrade_from_version_firmware"`
	ProductsWirelessControllerLastUpgradeFromVersionId          types.String                                                                   `tfsdk:"products_wireless_controller_last_upgrade_from_version_id"`
	ProductsWirelessControllerLastUpgradeFromVersionReleaseDate types.String                                                                   `tfsdk:"products_wireless_controller_last_upgrade_from_version_release_date"`
	ProductsWirelessControllerLastUpgradeFromVersionReleaseType types.String                                                                   `tfsdk:"products_wireless_controller_last_upgrade_from_version_release_type"`
	ProductsWirelessControllerLastUpgradeFromVersionShortName   types.String                                                                   `tfsdk:"products_wireless_controller_last_upgrade_from_version_short_name"`
	ProductsWirelessControllerLastUpgradeToVersionFirmware      types.String                                                                   `tfsdk:"products_wireless_controller_last_upgrade_to_version_firmware"`
	ProductsWirelessControllerLastUpgradeToVersionId            types.String                                                                   `tfsdk:"products_wireless_controller_last_upgrade_to_version_id"`
	ProductsWirelessControllerLastUpgradeToVersionReleaseDate   types.String                                                                   `tfsdk:"products_wireless_controller_last_upgrade_to_version_release_date"`
	ProductsWirelessControllerLastUpgradeToVersionReleaseType   types.String                                                                   `tfsdk:"products_wireless_controller_last_upgrade_to_version_release_type"`
	ProductsWirelessControllerLastUpgradeToVersionShortName     types.String                                                                   `tfsdk:"products_wireless_controller_last_upgrade_to_version_short_name"`
	ProductsWirelessControllerNextUpgradeTime                   types.String                                                                   `tfsdk:"products_wireless_controller_next_upgrade_time"`
	ProductsWirelessControllerNextUpgradeToVersionId            types.String                                                                   `tfsdk:"products_wireless_controller_next_upgrade_to_version_id"`
	ProductsWirelessControllerNextUpgradeToVersionFirmware      types.String                                                                   `tfsdk:"products_wireless_controller_next_upgrade_to_version_firmware"`
	ProductsWirelessControllerNextUpgradeToVersionReleaseDate   types.String                                                                   `tfsdk:"products_wireless_controller_next_upgrade_to_version_release_date"`
	ProductsWirelessControllerNextUpgradeToVersionReleaseType   types.String                                                                   `tfsdk:"products_wireless_controller_next_upgrade_to_version_release_type"`
	ProductsWirelessControllerNextUpgradeToVersionShortName     types.String                                                                   `tfsdk:"products_wireless_controller_next_upgrade_to_version_short_name"`
	ProductsWirelessControllerAvailableVersions                 []DataSourceNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions `tfsdk:"products_wireless_controller_available_versions"`
	UpgradeWindowDayOfWeek                                      types.String                                                                   `tfsdk:"upgrade_window_day_of_week"`
	UpgradeWindowHourOfDay                                      types.String                                                                   `tfsdk:"upgrade_window_hour_of_day"`
}

type DataSourceNetworkFirmwareUpgradesProductsApplianceAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type DataSourceNetworkFirmwareUpgradesProductsCameraAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type DataSourceNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type DataSourceNetworkFirmwareUpgradesProductsSecureConnectAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type DataSourceNetworkFirmwareUpgradesProductsSensorAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type DataSourceNetworkFirmwareUpgradesProductsSwitchAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type DataSourceNetworkFirmwareUpgradesProductsSwitchCatalystAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type DataSourceNetworkFirmwareUpgradesProductsWirelessAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type DataSourceNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions struct {
	Firmware    types.String `tfsdk:"firmware"`
	Id          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworkFirmwareUpgrades) getPath() string {
	return fmt.Sprintf("/networks/%v/firmwareUpgrades", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkFirmwareUpgrades) fromBody(ctx context.Context, res meraki.Res) {
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
		data.ProductsApplianceAvailableVersions = make([]DataSourceNetworkFirmwareUpgradesProductsApplianceAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceNetworkFirmwareUpgradesProductsApplianceAvailableVersions{}
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
		data.ProductsCameraAvailableVersions = make([]DataSourceNetworkFirmwareUpgradesProductsCameraAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceNetworkFirmwareUpgradesProductsCameraAvailableVersions{}
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
		data.ProductsCellularGatewayAvailableVersions = make([]DataSourceNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceNetworkFirmwareUpgradesProductsCellularGatewayAvailableVersions{}
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
		data.ProductsSecureConnectAvailableVersions = make([]DataSourceNetworkFirmwareUpgradesProductsSecureConnectAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceNetworkFirmwareUpgradesProductsSecureConnectAvailableVersions{}
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
		data.ProductsSensorAvailableVersions = make([]DataSourceNetworkFirmwareUpgradesProductsSensorAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceNetworkFirmwareUpgradesProductsSensorAvailableVersions{}
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
		data.ProductsSwitchAvailableVersions = make([]DataSourceNetworkFirmwareUpgradesProductsSwitchAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceNetworkFirmwareUpgradesProductsSwitchAvailableVersions{}
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
	if value := res.Get("products.switchCatalyst.isUpgradeAvailable"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchCatalystIsUpgradeAvailable = types.BoolNull()
	}
	if value := res.Get("products.switchCatalyst.participateInNextBetaRelease"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolValue(value.Bool())
	} else {
		data.ProductsSwitchCatalystParticipateInNextBetaRelease = types.BoolNull()
	}
	if value := res.Get("products.switchCatalyst.currentVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystCurrentVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystCurrentVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.currentVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystCurrentVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystCurrentVersionId = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.currentVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystCurrentVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystCurrentVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.currentVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystCurrentVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystCurrentVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.currentVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystCurrentVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystCurrentVersionShortName = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.lastUpgrade.time"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystLastUpgradeTime = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystLastUpgradeTime = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.lastUpgrade.fromVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystLastUpgradeFromVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystLastUpgradeFromVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.lastUpgrade.fromVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystLastUpgradeFromVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystLastUpgradeFromVersionId = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.lastUpgrade.fromVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystLastUpgradeFromVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystLastUpgradeFromVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.lastUpgrade.fromVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystLastUpgradeFromVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystLastUpgradeFromVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.lastUpgrade.fromVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystLastUpgradeFromVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystLastUpgradeFromVersionShortName = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.lastUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystLastUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystLastUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.lastUpgrade.toVersion.id"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystLastUpgradeToVersionId = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystLastUpgradeToVersionId = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.lastUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystLastUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystLastUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.lastUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystLastUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystLastUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.lastUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystLastUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystLastUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.toVersion.firmware"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystNextUpgradeToVersionFirmware = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeToVersionFirmware = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.toVersion.releaseDate"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystNextUpgradeToVersionReleaseDate = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeToVersionReleaseDate = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.toVersion.releaseType"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystNextUpgradeToVersionReleaseType = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeToVersionReleaseType = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.nextUpgrade.toVersion.shortName"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystNextUpgradeToVersionShortName = types.StringValue(value.String())
	} else {
		data.ProductsSwitchCatalystNextUpgradeToVersionShortName = types.StringNull()
	}
	if value := res.Get("products.switchCatalyst.availableVersions"); value.Exists() && value.Value() != nil {
		data.ProductsSwitchCatalystAvailableVersions = make([]DataSourceNetworkFirmwareUpgradesProductsSwitchCatalystAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceNetworkFirmwareUpgradesProductsSwitchCatalystAvailableVersions{}
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
			(*parent).ProductsSwitchCatalystAvailableVersions = append((*parent).ProductsSwitchCatalystAvailableVersions, data)
			return true
		})
	}
	if value := res.Get("products.wireless.isUpgradeAvailable"); value.Exists() && value.Value() != nil {
		data.ProductsWirelessIsUpgradeAvailable = types.BoolValue(value.Bool())
	} else {
		data.ProductsWirelessIsUpgradeAvailable = types.BoolNull()
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
		data.ProductsWirelessAvailableVersions = make([]DataSourceNetworkFirmwareUpgradesProductsWirelessAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceNetworkFirmwareUpgradesProductsWirelessAvailableVersions{}
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
		data.ProductsWirelessControllerAvailableVersions = make([]DataSourceNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceNetworkFirmwareUpgradesProductsWirelessControllerAvailableVersions{}
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
