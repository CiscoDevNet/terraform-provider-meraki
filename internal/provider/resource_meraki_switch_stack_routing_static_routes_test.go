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

func TestAccMerakiSwitchStackRoutingStaticRoutes(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_switch_1_serial") == "" || os.Getenv("TF_VAR_test_switch_2_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_switch_1_serial and TF_VAR_test_switch_2_serial")
	}

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchStackRoutingStaticRoutesPrerequisitesConfig + testAccMerakiSwitchStackRoutingStaticRoutesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchStackRoutingStaticRoutesPrerequisitesConfig + testAccMerakiSwitchStackRoutingStaticRoutesConfig_all(),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:      "meraki_switch_stack_routing_static_routes.test",
		ImportState:       true,
		ImportStateIdFunc: merakiSwitchStackRoutingStaticRoutesImportStateIdFunc("meraki_switch_stack_routing_static_routes.test"),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiSwitchStackRoutingStaticRoutesImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]
		NetworkId := primary.Attributes["network_id"]
		SwitchStackId := primary.Attributes["switch_stack_id"]

		return fmt.Sprintf("%s,%s,%s", OrganizationId, NetworkId, SwitchStackId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiSwitchStackRoutingStaticRoutesPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_switch_1_serial" {}
variable "test_switch_2_serial" {}
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
  serials    = [var.test_switch_1_serial, var.test_switch_2_serial]
}
resource "meraki_switch_stack" "test" {
  network_id = meraki_network.test.id
  name      = "A cool stack"
  serials   = meraki_network_device_claim.test.serials
}
resource "meraki_switch_stack_routing_interface" "test" {
  network_id = meraki_network.test.id
  switch_stack_id = meraki_switch_stack.test.id
  default_gateway = "192.168.1.1"
  interface_ip = "192.168.1.2"
  name = "L3 interface"
  subnet = "192.168.1.0/24"
  vlan_id = 100
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiSwitchStackRoutingStaticRoutesConfig_minimum() string {
	config := `resource "meraki_switch_stack_routing_static_routes" "test" {` + "\n"
	config += ` network_id = meraki_network.test.id` + "\n"
	config += ` switch_stack_id = meraki_switch_stack_routing_interface.test.switch_stack_id` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  next_hop_ip = "192.168.1.1"` + "\n"
	config += `  subnet = "192.168.2.0/24"` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchStackRoutingStaticRoutesConfig_all() string {
	config := `resource "meraki_switch_stack_routing_static_routes" "test" {` + "\n"
	config += ` network_id = meraki_network.test.id` + "\n"
	config += ` switch_stack_id = meraki_switch_stack_routing_interface.test.switch_stack_id` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  name = "My route"` + "\n"
	config += `  next_hop_ip = "192.168.1.1"` + "\n"
	config += `  subnet = "192.168.2.0/24"` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
