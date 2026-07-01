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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceMerakiNetworkFirmwareUpgradesData(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "timezone", "America/Los_Angeles"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_is_upgrade_available", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_current_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_current_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_current_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_current_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_current_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_last_upgrade_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_last_upgrade_from_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_last_upgrade_from_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_last_upgrade_from_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_last_upgrade_from_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_last_upgrade_from_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_last_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_last_upgrade_to_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_last_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_last_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_last_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_next_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_next_upgrade_to_version_id", "1001"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_next_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_next_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_next_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_available_versions.0.firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_available_versions.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_available_versions.0.release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_available_versions.0.release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_appliance_available_versions.0.short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_is_upgrade_available", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_current_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_current_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_current_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_current_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_current_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_last_upgrade_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_last_upgrade_from_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_last_upgrade_from_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_last_upgrade_from_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_last_upgrade_from_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_last_upgrade_from_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_last_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_last_upgrade_to_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_last_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_last_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_last_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_next_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_next_upgrade_to_version_id", "1003"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_next_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_next_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_next_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_available_versions.0.firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_available_versions.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_available_versions.0.release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_available_versions.0.release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_camera_available_versions.0.short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_is_upgrade_available", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_current_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_current_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_current_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_current_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_current_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_last_upgrade_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_last_upgrade_from_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_last_upgrade_from_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_last_upgrade_from_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_last_upgrade_from_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_last_upgrade_from_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_last_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_last_upgrade_to_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_last_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_last_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_last_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_next_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_next_upgrade_to_version_id", "1004"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_next_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_next_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_next_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_available_versions.0.firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_available_versions.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_available_versions.0.release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_available_versions.0.release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_cellular_gateway_available_versions.0.short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_is_upgrade_available", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_current_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_current_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_current_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_current_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_current_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_last_upgrade_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_last_upgrade_from_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_last_upgrade_from_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_last_upgrade_from_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_last_upgrade_from_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_last_upgrade_from_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_last_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_last_upgrade_to_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_last_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_last_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_last_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_next_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_next_upgrade_to_version_id", "1008"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_next_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_next_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_next_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_available_versions.0.firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_available_versions.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_available_versions.0.release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_available_versions.0.release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_secure_connect_available_versions.0.short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_is_upgrade_available", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_current_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_current_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_current_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_current_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_current_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_last_upgrade_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_last_upgrade_from_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_last_upgrade_from_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_last_upgrade_from_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_last_upgrade_from_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_last_upgrade_from_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_last_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_last_upgrade_to_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_last_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_last_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_last_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_next_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_next_upgrade_to_version_id", "1005"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_next_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_next_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_next_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_available_versions.0.firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_available_versions.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_available_versions.0.release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_available_versions.0.release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_sensor_available_versions.0.short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_is_upgrade_available", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_current_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_current_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_current_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_current_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_current_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_last_upgrade_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_last_upgrade_from_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_last_upgrade_from_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_last_upgrade_from_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_last_upgrade_from_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_last_upgrade_from_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_last_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_last_upgrade_to_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_last_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_last_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_last_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_next_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_next_upgrade_to_version_id", "1002"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_next_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_next_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_next_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_available_versions.0.firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_available_versions.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_available_versions.0.release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_available_versions.0.release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_available_versions.0.short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_is_upgrade_available", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_catalyst_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_catalyst_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_switch_catalyst_next_upgrade_to_version_id", "1234"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_current_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_current_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_current_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_current_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_current_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_last_upgrade_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_last_upgrade_from_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_last_upgrade_from_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_last_upgrade_from_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_last_upgrade_from_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_last_upgrade_from_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_last_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_last_upgrade_to_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_last_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_last_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_last_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_next_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_next_upgrade_to_version_id", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_next_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_next_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_next_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_available_versions.0.firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_available_versions.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_available_versions.0.release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_available_versions.0.release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_available_versions.0.short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_is_upgrade_available", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_participate_in_next_beta_release", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_current_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_current_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_current_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_current_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_current_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_last_upgrade_time", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_last_upgrade_from_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_last_upgrade_from_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_last_upgrade_from_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_last_upgrade_from_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_last_upgrade_from_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_last_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_last_upgrade_to_version_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_last_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_last_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_last_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_next_upgrade_time", "2019-03-17T17:22:52Z"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_next_upgrade_to_version_firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_next_upgrade_to_version_id", "1006"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_next_upgrade_to_version_release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_next_upgrade_to_version_release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_next_upgrade_to_version_short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_available_versions.0.firmware", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_available_versions.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_available_versions.0.release_date", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_available_versions.0.release_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "products_wireless_controller_available_versions.0.short_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "upgrade_window_day_of_week", "sun"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_firmware_upgrades_data.test", "upgrade_window_hour_of_day", "4:00"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiNetworkFirmwareUpgradesDataPrerequisitesConfig + testAccDataSourceMerakiNetworkFirmwareUpgradesDataConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiNetworkFirmwareUpgradesDataPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance", "sensor", "camera"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiNetworkFirmwareUpgradesDataConfig() string {
	config := `resource "meraki_network_firmware_upgrades_data" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  timezone = "America/Los_Angeles"` + "\n"
	config += `  products_appliance_is_upgrade_available = ` + "\n"
	config += `  products_appliance_participate_in_next_beta_release = false` + "\n"
	config += `  products_appliance_current_version_firmware = ""` + "\n"
	config += `  products_appliance_current_version_id = ""` + "\n"
	config += `  products_appliance_current_version_release_date = ""` + "\n"
	config += `  products_appliance_current_version_release_type = ""` + "\n"
	config += `  products_appliance_current_version_short_name = ""` + "\n"
	config += `  products_appliance_last_upgrade_time = ""` + "\n"
	config += `  products_appliance_last_upgrade_from_version_firmware = ""` + "\n"
	config += `  products_appliance_last_upgrade_from_version_id = ""` + "\n"
	config += `  products_appliance_last_upgrade_from_version_release_date = ""` + "\n"
	config += `  products_appliance_last_upgrade_from_version_release_type = ""` + "\n"
	config += `  products_appliance_last_upgrade_from_version_short_name = ""` + "\n"
	config += `  products_appliance_last_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_appliance_last_upgrade_to_version_id = ""` + "\n"
	config += `  products_appliance_last_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_appliance_last_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_appliance_last_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_appliance_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_appliance_next_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_appliance_next_upgrade_to_version_id = "1001"` + "\n"
	config += `  products_appliance_next_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_appliance_next_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_appliance_next_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_appliance_available_versions = [{` + "\n"
	config += `    firmware = ""` + "\n"
	config += `    id = ""` + "\n"
	config += `    release_date = ""` + "\n"
	config += `    release_type = ""` + "\n"
	config += `    short_name = ""` + "\n"
	config += `  }]` + "\n"
	config += `  products_camera_is_upgrade_available = ` + "\n"
	config += `  products_camera_participate_in_next_beta_release = false` + "\n"
	config += `  products_camera_current_version_firmware = ""` + "\n"
	config += `  products_camera_current_version_id = ""` + "\n"
	config += `  products_camera_current_version_release_date = ""` + "\n"
	config += `  products_camera_current_version_release_type = ""` + "\n"
	config += `  products_camera_current_version_short_name = ""` + "\n"
	config += `  products_camera_last_upgrade_time = ""` + "\n"
	config += `  products_camera_last_upgrade_from_version_firmware = ""` + "\n"
	config += `  products_camera_last_upgrade_from_version_id = ""` + "\n"
	config += `  products_camera_last_upgrade_from_version_release_date = ""` + "\n"
	config += `  products_camera_last_upgrade_from_version_release_type = ""` + "\n"
	config += `  products_camera_last_upgrade_from_version_short_name = ""` + "\n"
	config += `  products_camera_last_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_camera_last_upgrade_to_version_id = ""` + "\n"
	config += `  products_camera_last_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_camera_last_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_camera_last_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_camera_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_camera_next_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_camera_next_upgrade_to_version_id = "1003"` + "\n"
	config += `  products_camera_next_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_camera_next_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_camera_next_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_camera_available_versions = [{` + "\n"
	config += `    firmware = ""` + "\n"
	config += `    id = ""` + "\n"
	config += `    release_date = ""` + "\n"
	config += `    release_type = ""` + "\n"
	config += `    short_name = ""` + "\n"
	config += `  }]` + "\n"
	config += `  products_cellular_gateway_is_upgrade_available = ` + "\n"
	config += `  products_cellular_gateway_participate_in_next_beta_release = false` + "\n"
	config += `  products_cellular_gateway_current_version_firmware = ""` + "\n"
	config += `  products_cellular_gateway_current_version_id = ""` + "\n"
	config += `  products_cellular_gateway_current_version_release_date = ""` + "\n"
	config += `  products_cellular_gateway_current_version_release_type = ""` + "\n"
	config += `  products_cellular_gateway_current_version_short_name = ""` + "\n"
	config += `  products_cellular_gateway_last_upgrade_time = ""` + "\n"
	config += `  products_cellular_gateway_last_upgrade_from_version_firmware = ""` + "\n"
	config += `  products_cellular_gateway_last_upgrade_from_version_id = ""` + "\n"
	config += `  products_cellular_gateway_last_upgrade_from_version_release_date = ""` + "\n"
	config += `  products_cellular_gateway_last_upgrade_from_version_release_type = ""` + "\n"
	config += `  products_cellular_gateway_last_upgrade_from_version_short_name = ""` + "\n"
	config += `  products_cellular_gateway_last_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_cellular_gateway_last_upgrade_to_version_id = ""` + "\n"
	config += `  products_cellular_gateway_last_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_cellular_gateway_last_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_cellular_gateway_last_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_cellular_gateway_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_cellular_gateway_next_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_cellular_gateway_next_upgrade_to_version_id = "1004"` + "\n"
	config += `  products_cellular_gateway_next_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_cellular_gateway_next_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_cellular_gateway_next_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_cellular_gateway_available_versions = [{` + "\n"
	config += `    firmware = ""` + "\n"
	config += `    id = ""` + "\n"
	config += `    release_date = ""` + "\n"
	config += `    release_type = ""` + "\n"
	config += `    short_name = ""` + "\n"
	config += `  }]` + "\n"
	config += `  products_secure_connect_is_upgrade_available = ` + "\n"
	config += `  products_secure_connect_participate_in_next_beta_release = false` + "\n"
	config += `  products_secure_connect_current_version_firmware = ""` + "\n"
	config += `  products_secure_connect_current_version_id = ""` + "\n"
	config += `  products_secure_connect_current_version_release_date = ""` + "\n"
	config += `  products_secure_connect_current_version_release_type = ""` + "\n"
	config += `  products_secure_connect_current_version_short_name = ""` + "\n"
	config += `  products_secure_connect_last_upgrade_time = ""` + "\n"
	config += `  products_secure_connect_last_upgrade_from_version_firmware = ""` + "\n"
	config += `  products_secure_connect_last_upgrade_from_version_id = ""` + "\n"
	config += `  products_secure_connect_last_upgrade_from_version_release_date = ""` + "\n"
	config += `  products_secure_connect_last_upgrade_from_version_release_type = ""` + "\n"
	config += `  products_secure_connect_last_upgrade_from_version_short_name = ""` + "\n"
	config += `  products_secure_connect_last_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_secure_connect_last_upgrade_to_version_id = ""` + "\n"
	config += `  products_secure_connect_last_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_secure_connect_last_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_secure_connect_last_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_secure_connect_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_secure_connect_next_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_secure_connect_next_upgrade_to_version_id = "1008"` + "\n"
	config += `  products_secure_connect_next_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_secure_connect_next_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_secure_connect_next_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_secure_connect_available_versions = [{` + "\n"
	config += `    firmware = ""` + "\n"
	config += `    id = ""` + "\n"
	config += `    release_date = ""` + "\n"
	config += `    release_type = ""` + "\n"
	config += `    short_name = ""` + "\n"
	config += `  }]` + "\n"
	config += `  products_sensor_is_upgrade_available = ` + "\n"
	config += `  products_sensor_participate_in_next_beta_release = false` + "\n"
	config += `  products_sensor_current_version_firmware = ""` + "\n"
	config += `  products_sensor_current_version_id = ""` + "\n"
	config += `  products_sensor_current_version_release_date = ""` + "\n"
	config += `  products_sensor_current_version_release_type = ""` + "\n"
	config += `  products_sensor_current_version_short_name = ""` + "\n"
	config += `  products_sensor_last_upgrade_time = ""` + "\n"
	config += `  products_sensor_last_upgrade_from_version_firmware = ""` + "\n"
	config += `  products_sensor_last_upgrade_from_version_id = ""` + "\n"
	config += `  products_sensor_last_upgrade_from_version_release_date = ""` + "\n"
	config += `  products_sensor_last_upgrade_from_version_release_type = ""` + "\n"
	config += `  products_sensor_last_upgrade_from_version_short_name = ""` + "\n"
	config += `  products_sensor_last_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_sensor_last_upgrade_to_version_id = ""` + "\n"
	config += `  products_sensor_last_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_sensor_last_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_sensor_last_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_sensor_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_sensor_next_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_sensor_next_upgrade_to_version_id = "1005"` + "\n"
	config += `  products_sensor_next_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_sensor_next_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_sensor_next_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_sensor_available_versions = [{` + "\n"
	config += `    firmware = ""` + "\n"
	config += `    id = ""` + "\n"
	config += `    release_date = ""` + "\n"
	config += `    release_type = ""` + "\n"
	config += `    short_name = ""` + "\n"
	config += `  }]` + "\n"
	config += `  products_switch_is_upgrade_available = ` + "\n"
	config += `  products_switch_participate_in_next_beta_release = false` + "\n"
	config += `  products_switch_current_version_firmware = ""` + "\n"
	config += `  products_switch_current_version_id = ""` + "\n"
	config += `  products_switch_current_version_release_date = ""` + "\n"
	config += `  products_switch_current_version_release_type = ""` + "\n"
	config += `  products_switch_current_version_short_name = ""` + "\n"
	config += `  products_switch_last_upgrade_time = ""` + "\n"
	config += `  products_switch_last_upgrade_from_version_firmware = ""` + "\n"
	config += `  products_switch_last_upgrade_from_version_id = ""` + "\n"
	config += `  products_switch_last_upgrade_from_version_release_date = ""` + "\n"
	config += `  products_switch_last_upgrade_from_version_release_type = ""` + "\n"
	config += `  products_switch_last_upgrade_from_version_short_name = ""` + "\n"
	config += `  products_switch_last_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_switch_last_upgrade_to_version_id = ""` + "\n"
	config += `  products_switch_last_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_switch_last_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_switch_last_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_switch_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_switch_next_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_switch_next_upgrade_to_version_id = "1002"` + "\n"
	config += `  products_switch_next_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_switch_next_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_switch_next_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_switch_available_versions = [{` + "\n"
	config += `    firmware = ""` + "\n"
	config += `    id = ""` + "\n"
	config += `    release_date = ""` + "\n"
	config += `    release_type = ""` + "\n"
	config += `    short_name = ""` + "\n"
	config += `  }]` + "\n"
	config += `  products_wireless_is_upgrade_available = ` + "\n"
	config += `  products_switch_catalyst_participate_in_next_beta_release = false` + "\n"
	config += `  products_switch_catalyst_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_switch_catalyst_next_upgrade_to_version_id = "1234"` + "\n"
	config += `  products_wireless_participate_in_next_beta_release = false` + "\n"
	config += `  products_wireless_current_version_firmware = ""` + "\n"
	config += `  products_wireless_current_version_id = ""` + "\n"
	config += `  products_wireless_current_version_release_date = ""` + "\n"
	config += `  products_wireless_current_version_release_type = ""` + "\n"
	config += `  products_wireless_current_version_short_name = ""` + "\n"
	config += `  products_wireless_last_upgrade_time = ""` + "\n"
	config += `  products_wireless_last_upgrade_from_version_firmware = ""` + "\n"
	config += `  products_wireless_last_upgrade_from_version_id = ""` + "\n"
	config += `  products_wireless_last_upgrade_from_version_release_date = ""` + "\n"
	config += `  products_wireless_last_upgrade_from_version_release_type = ""` + "\n"
	config += `  products_wireless_last_upgrade_from_version_short_name = ""` + "\n"
	config += `  products_wireless_last_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_wireless_last_upgrade_to_version_id = ""` + "\n"
	config += `  products_wireless_last_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_wireless_last_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_wireless_last_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_wireless_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_wireless_next_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_wireless_next_upgrade_to_version_id = "1000"` + "\n"
	config += `  products_wireless_next_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_wireless_next_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_wireless_next_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_wireless_available_versions = [{` + "\n"
	config += `    firmware = ""` + "\n"
	config += `    id = ""` + "\n"
	config += `    release_date = ""` + "\n"
	config += `    release_type = ""` + "\n"
	config += `    short_name = ""` + "\n"
	config += `  }]` + "\n"
	config += `  products_wireless_controller_is_upgrade_available = ` + "\n"
	config += `  products_wireless_controller_participate_in_next_beta_release = false` + "\n"
	config += `  products_wireless_controller_current_version_firmware = ""` + "\n"
	config += `  products_wireless_controller_current_version_id = ""` + "\n"
	config += `  products_wireless_controller_current_version_release_date = ""` + "\n"
	config += `  products_wireless_controller_current_version_release_type = ""` + "\n"
	config += `  products_wireless_controller_current_version_short_name = ""` + "\n"
	config += `  products_wireless_controller_last_upgrade_time = ""` + "\n"
	config += `  products_wireless_controller_last_upgrade_from_version_firmware = ""` + "\n"
	config += `  products_wireless_controller_last_upgrade_from_version_id = ""` + "\n"
	config += `  products_wireless_controller_last_upgrade_from_version_release_date = ""` + "\n"
	config += `  products_wireless_controller_last_upgrade_from_version_release_type = ""` + "\n"
	config += `  products_wireless_controller_last_upgrade_from_version_short_name = ""` + "\n"
	config += `  products_wireless_controller_last_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_wireless_controller_last_upgrade_to_version_id = ""` + "\n"
	config += `  products_wireless_controller_last_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_wireless_controller_last_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_wireless_controller_last_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_wireless_controller_next_upgrade_time = "2019-03-17T17:22:52Z"` + "\n"
	config += `  products_wireless_controller_next_upgrade_to_version_firmware = ""` + "\n"
	config += `  products_wireless_controller_next_upgrade_to_version_id = "1006"` + "\n"
	config += `  products_wireless_controller_next_upgrade_to_version_release_date = ""` + "\n"
	config += `  products_wireless_controller_next_upgrade_to_version_release_type = ""` + "\n"
	config += `  products_wireless_controller_next_upgrade_to_version_short_name = ""` + "\n"
	config += `  products_wireless_controller_available_versions = [{` + "\n"
	config += `    firmware = ""` + "\n"
	config += `    id = ""` + "\n"
	config += `    release_date = ""` + "\n"
	config += `    release_type = ""` + "\n"
	config += `    short_name = ""` + "\n"
	config += `  }]` + "\n"
	config += `  upgrade_window_day_of_week = "sun"` + "\n"
	config += `  upgrade_window_hour_of_day = "4:00"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_network_firmware_upgrades_data" "test" {
			network_id = meraki_network.test.id
			depends_on = [meraki_network_firmware_upgrades_data.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
