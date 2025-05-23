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

func TestAccMerakiWirelessAlternateManagementInterface(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_ap_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_ap_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_alternate_management_interface.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_alternate_management_interface.test", "vlan_id", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_alternate_management_interface.test", "access_points.0.alternate_management_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_alternate_management_interface.test", "access_points.0.dns1", "8.8.8.8"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_alternate_management_interface.test", "access_points.0.dns2", "8.8.4.4"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_alternate_management_interface.test", "access_points.0.gateway", "1.2.3.5"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_alternate_management_interface.test", "access_points.0.subnet_mask", "255.255.255.0"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiWirelessAlternateManagementInterfacePrerequisitesConfig + testAccMerakiWirelessAlternateManagementInterfaceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessAlternateManagementInterfacePrerequisitesConfig + testAccMerakiWirelessAlternateManagementInterfaceConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_wireless_alternate_management_interface.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiWirelessAlternateManagementInterfaceImportStateIdFunc("meraki_wireless_alternate_management_interface.test"),
		ImportStateVerifyIgnore: []string{},
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

func merakiWirelessAlternateManagementInterfaceImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiWirelessAlternateManagementInterfacePrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_ap_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_ap_1_serial]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiWirelessAlternateManagementInterfaceConfig_minimum() string {
	config := `resource "meraki_wireless_alternate_management_interface" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  enabled = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiWirelessAlternateManagementInterfaceConfig_all() string {
	config := `resource "meraki_wireless_alternate_management_interface" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  enabled = true` + "\n"
	config += `  vlan_id = 100` + "\n"
	config += `  access_points = [{` + "\n"
	config += `    alternate_management_ip = "1.2.3.4"` + "\n"
	config += `    dns1 = "8.8.8.8"` + "\n"
	config += `    dns2 = "8.8.4.4"` + "\n"
	config += `    gateway = "1.2.3.5"` + "\n"
	config += `    serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `    subnet_mask = "255.255.255.0"` + "\n"
	config += `  }]` + "\n"
	config += `  protocols = ["radius"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
