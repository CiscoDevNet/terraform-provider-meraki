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
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiWirelessMQTTSettings(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "ble_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "ble_type", "iBeacon"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "ble_flush_frequency", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "ble_hysteresis_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "ble_hysteresis_threshold", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "mqtt_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "mqtt_topic", "Test Topic"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "mqtt_publishing_frequency", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "mqtt_publishing_qos", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "mqtt_message_fields.0", "RSSI"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "wifi_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "wifi_type", "Associated"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "wifi_flush_frequency", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "wifi_hysteresis_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_mqtt_settings.test", "wifi_hysteresis_threshold", "1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiWirelessMQTTSettingsPrerequisitesConfig + testAccMerakiWirelessMQTTSettingsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessMQTTSettingsPrerequisitesConfig + testAccMerakiWirelessMQTTSettingsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_wireless_mqtt_settings.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiWirelessMQTTSettingsImportStateIdFunc("meraki_wireless_mqtt_settings.test"),
		ImportStateVerifyIgnore: []string{"ble_allow_lists_macs", "ble_allow_lists_uuids", "wifi_allow_lists_macs"},
		Check:                   resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiWirelessMQTTSettingsImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s,%s", OrganizationId, NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiWirelessMQTTSettingsPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless"]
}
resource "meraki_network_mqtt_broker" "test" {
  network_id = meraki_network.test.id
  host       = "1.2.3.4"
  name       = "My Broker"
  port       = 443
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiWirelessMQTTSettingsConfig_minimum() string {
	config := `resource "meraki_wireless_mqtt_settings" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  ble_enabled = false` + "\n"
	config += `  mqtt_enabled = true` + "\n"
	config += `  mqtt_topic = "Test Topic"` + "\n"
	config += `  mqtt_broker_name = meraki_network_mqtt_broker.test.name` + "\n"
	config += `  mqtt_publishing_frequency = 1` + "\n"
	config += `  mqtt_publishing_qos = 1` + "\n"
	config += `  mqtt_message_fields = ["RSSI"]` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  wifi_enabled = false` + "\n"
	config += `  wifi_type = "Visible"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiWirelessMQTTSettingsConfig_all() string {
	config := `resource "meraki_wireless_mqtt_settings" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  ble_enabled = false` + "\n"
	config += `  ble_type = "iBeacon"` + "\n"
	config += `  ble_flush_frequency = 60` + "\n"
	config += `  ble_hysteresis_enabled = true` + "\n"
	config += `  ble_hysteresis_threshold = 1` + "\n"
	config += `  mqtt_enabled = true` + "\n"
	config += `  mqtt_topic = "Test Topic"` + "\n"
	config += `  mqtt_broker_name = meraki_network_mqtt_broker.test.name` + "\n"
	config += `  mqtt_publishing_frequency = 1` + "\n"
	config += `  mqtt_publishing_qos = 1` + "\n"
	config += `  mqtt_message_fields = ["RSSI"]` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  wifi_enabled = false` + "\n"
	config += `  wifi_type = "Associated"` + "\n"
	config += `  wifi_flush_frequency = 60` + "\n"
	config += `  wifi_hysteresis_enabled = false` + "\n"
	config += `  wifi_hysteresis_threshold = 1` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
