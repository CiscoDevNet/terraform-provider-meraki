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

func TestAccMerakiWirelessEthernetPortProfileAssign(t *testing.T) {
	if os.Getenv("WIRELESS_ETHERNET_PORT_PROFILE_ASSIGN") == "" {
		t.Skip("skipping test, set environment variable WIRELESS_ETHERNET_PORT_PROFILE_ASSIGN")
	}
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_ap_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_ap_1_serial")
	}
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiWirelessEthernetPortProfileAssignPrerequisitesConfig + testAccMerakiWirelessEthernetPortProfileAssignConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessEthernetPortProfileAssignPrerequisitesConfig + testAccMerakiWirelessEthernetPortProfileAssignConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiWirelessEthernetPortProfileAssignImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiWirelessEthernetPortProfileAssignPrerequisitesConfig = `
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
resource "meraki_wireless_ethernet_port_profile" "test" {
  network_id = meraki_network.test.id
  name       = "test"
  ports = [
    {
      enabled = true
      name    = "port"
    }
  ]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiWirelessEthernetPortProfileAssignConfig_minimum() string {
	config := `resource "meraki_wireless_ethernet_port_profile_assign" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  profile_id = meraki_wireless_ethernet_port_profile.test.id` + "\n"
	config += `  serials = meraki_network_device_claim.test.serials` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiWirelessEthernetPortProfileAssignConfig_all() string {
	config := `resource "meraki_wireless_ethernet_port_profile_assign" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  profile_id = meraki_wireless_ethernet_port_profile.test.id` + "\n"
	config += `  serials = meraki_network_device_claim.test.serials` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
