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

func TestAccMerakiSwitchStackRoutingStaticRoute(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_switch_1_serial") == "" || os.Getenv("TF_VAR_test_switch_2_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_switch_1_serial and TF_VAR_test_switch_2_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_static_route.test", "name", "My route"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_static_route.test", "next_hop_ip", "192.168.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_static_route.test", "subnet", "192.168.2.0/24"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchStackRoutingStaticRoutePrerequisitesConfig + testAccMerakiSwitchStackRoutingStaticRouteConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchStackRoutingStaticRoutePrerequisitesConfig + testAccMerakiSwitchStackRoutingStaticRouteConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_switch_stack_routing_static_route.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiSwitchStackRoutingStaticRouteImportStateIdFunc("meraki_switch_stack_routing_static_route.test"),
		ImportStateVerifyIgnore: []string{"advertise_via_ospf_enabled", "prefer_over_ospf_routes_enabled"},
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

func merakiSwitchStackRoutingStaticRouteImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		SwitchStackId := primary.Attributes["switch_stack_id"]
		Id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s,%s", NetworkId, SwitchStackId, Id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiSwitchStackRoutingStaticRoutePrerequisitesConfig = `
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

func testAccMerakiSwitchStackRoutingStaticRouteConfig_minimum() string {
	config := `resource "meraki_switch_stack_routing_static_route" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  switch_stack_id = meraki_switch_stack_routing_interface.test.switch_stack_id` + "\n"
	config += `  next_hop_ip = "192.168.1.1"` + "\n"
	config += `  subnet = "192.168.2.0/24"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchStackRoutingStaticRouteConfig_all() string {
	config := `resource "meraki_switch_stack_routing_static_route" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  switch_stack_id = meraki_switch_stack_routing_interface.test.switch_stack_id` + "\n"
	config += `  name = "My route"` + "\n"
	config += `  next_hop_ip = "192.168.1.1"` + "\n"
	config += `  subnet = "192.168.2.0/24"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
