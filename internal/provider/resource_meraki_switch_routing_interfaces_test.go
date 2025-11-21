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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiSwitchRoutingInterfaces(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_switch_1_serial") == "" {
        t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_switch_1_serial")
	}

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchRoutingInterfacesPrerequisitesConfig+testAccMerakiSwitchRoutingInterfacesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchRoutingInterfacesPrerequisitesConfig+testAccMerakiSwitchRoutingInterfacesConfig_all(),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "meraki_switch_routing_interfaces.test",
		ImportState: true,
		ImportStateIdFunc: merakiSwitchRoutingInterfacesImportStateIdFunc("meraki_switch_routing_interfaces.test"),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiSwitchRoutingInterfacesImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]
		Serial := primary.Attributes["serial"]

		return fmt.Sprintf("%s,%s", OrganizationId,Serial), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiSwitchRoutingInterfacesPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_switch_1_serial" {}
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
  serials    = [var.test_switch_1_serial]
}

`
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiSwitchRoutingInterfacesConfig_minimum() string {
	config := `resource "meraki_switch_routing_interfaces" "test" {` + "\n"
	config += ` serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  default_gateway = "192.168.1.1"` + "\n"
	config += `  interface_ip = "192.168.1.2"` + "\n"
	config += `  name = "L3 interface"` + "\n"
	config += `  subnet = "192.168.1.0/24"` + "\n"
	config += `  vlan_id = 100` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchRoutingInterfacesConfig_all() string {
	config := `resource "meraki_switch_routing_interfaces" "test" {` + "\n"
	config += ` serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  default_gateway = "192.168.1.1"` + "\n"
	config += `  interface_ip = "192.168.1.2"` + "\n"
	config += `  mode = "vlan"` + "\n"
	config += `  multicast_routing = "disabled"` + "\n"
	config += `  name = "L3 interface"` + "\n"
	config += `  subnet = "192.168.1.0/24"` + "\n"
	config += `  vlan_id = 100` + "\n"
	config += `  ipv6_address = "1:2:3:4::1"` + "\n"
	config += `  ipv6_assignment_mode = "static"` + "\n"
	config += `  ipv6_gateway = "1:2:3:4::2"` + "\n"
	config += `  ipv6_prefix = "1:2:3:4::/64"` + "\n"
	config += `  ospf_settings_area = "ospfDisabled"` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
