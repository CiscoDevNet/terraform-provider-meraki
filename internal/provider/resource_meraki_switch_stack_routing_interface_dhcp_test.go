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

func TestAccMerakiSwitchStackRoutingInterfaceDHCP(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_switch_1_serial") == "" || os.Getenv("TF_VAR_test_switch_2_serial") == "" {
        t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_switch_1_serial and TF_VAR_test_switch_2_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "boot_file_name", "home_boot_file"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "boot_next_server", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "boot_options_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "dhcp_lease_time", "1 day"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "dhcp_mode", "dhcpServer"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "dns_nameservers_option", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "dhcp_options.0.code", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "dhcp_options.0.type", "text"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "dhcp_options.0.value", "five"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "dns_custom_nameservers.0", "8.8.8.8"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "fixed_ip_assignments.0.ip", "192.168.1.12"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "fixed_ip_assignments.0.mac", "22:33:44:55:66:77"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "fixed_ip_assignments.0.name", "Cisco Meraki valued client"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "reserved_ip_ranges.0.comment", "A reserved IP range"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "reserved_ip_ranges.0.end", "192.168.1.10"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_stack_routing_interface_dhcp.test", "reserved_ip_ranges.0.start", "192.168.1.1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchStackRoutingInterfaceDHCPPrerequisitesConfig+testAccMerakiSwitchStackRoutingInterfaceDHCPConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchStackRoutingInterfaceDHCPPrerequisitesConfig+testAccMerakiSwitchStackRoutingInterfaceDHCPConfig_all(),
		Check: resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "meraki_switch_stack_routing_interface_dhcp.test",
		ImportState: true,
		ImportStateVerify: true,
		ImportStateIdFunc: merakiSwitchStackRoutingInterfaceDHCPImportStateIdFunc("meraki_switch_stack_routing_interface_dhcp.test"),
		ImportStateVerifyIgnore: []string{  },
		Check: resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiSwitchStackRoutingInterfaceDHCPImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		SwitchStackId := primary.Attributes["switch_stack_id"]
		InterfaceId := primary.Attributes["interface_id"]

		return fmt.Sprintf("%s,%s,%s", NetworkId,SwitchStackId,InterfaceId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiSwitchStackRoutingInterfaceDHCPPrerequisitesConfig = `
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

func testAccMerakiSwitchStackRoutingInterfaceDHCPConfig_minimum() string {
	config := `resource "meraki_switch_stack_routing_interface_dhcp" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  switch_stack_id = meraki_switch_stack.test.id` + "\n"
	config += `  interface_id = meraki_switch_stack_routing_interface.test.id` + "\n"
	config += `  dhcp_mode = "dhcpDisabled"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchStackRoutingInterfaceDHCPConfig_all() string {
	config := `resource "meraki_switch_stack_routing_interface_dhcp" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  switch_stack_id = meraki_switch_stack.test.id` + "\n"
	config += `  interface_id = meraki_switch_stack_routing_interface.test.id` + "\n"
	config += `  boot_file_name = "home_boot_file"` + "\n"
	config += `  boot_next_server = "1.2.3.4"` + "\n"
	config += `  boot_options_enabled = true` + "\n"
	config += `  dhcp_lease_time = "1 day"` + "\n"
	config += `  dhcp_mode = "dhcpServer"` + "\n"
	config += `  dns_nameservers_option = "custom"` + "\n"
	config += `  dhcp_options = [{` + "\n"
	config += `    code = "5"` + "\n"
	config += `    type = "text"` + "\n"
	config += `    value = "five"` + "\n"
	config += `  }]` + "\n"
	config += `  dns_custom_nameservers = ["8.8.8.8"]` + "\n"
	config += `  fixed_ip_assignments = [{` + "\n"
	config += `    ip = "192.168.1.12"` + "\n"
	config += `    mac = "22:33:44:55:66:77"` + "\n"
	config += `    name = "Cisco Meraki valued client"` + "\n"
	config += `  }]` + "\n"
	config += `  reserved_ip_ranges = [{` + "\n"
	config += `    comment = "A reserved IP range"` + "\n"
	config += `    end = "192.168.1.10"` + "\n"
	config += `    start = "192.168.1.1"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional



// End of section. //template:end testAccConfigAdditional
