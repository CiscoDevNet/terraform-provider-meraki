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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiSwitchPort(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_switch_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_switch_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "port_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "access_policy_type", "Sticky MAC allow list"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "allowed_vlans", "1,3,5-10"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "dai_trusted", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "isolation_enabled", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "link_negotiation", "Auto negotiate"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "name", "My switch port"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "poe_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "rstp_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "sticky_mac_allow_list_limit", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "stp_guard", "disabled"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "type", "access"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "udld", "Alert only"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "vlan", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "voice_vlan", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_port.test", "profile_enabled", "false"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchPortPrerequisitesConfig + testAccMerakiSwitchPortConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchPortPrerequisitesConfig + testAccMerakiSwitchPortConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiSwitchPortPrerequisitesConfig = `
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

func testAccMerakiSwitchPortConfig_minimum() string {
	config := `resource "meraki_switch_port" "test" {` + "\n"
	config += `	serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `	port_id = "1"` + "\n"
	config += `	enabled = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchPortConfig_all() string {
	config := `resource "meraki_switch_port" "test" {` + "\n"
	config += `	serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `	port_id = "1"` + "\n"
	config += `	access_policy_type = "Sticky MAC allow list"` + "\n"
	config += `	allowed_vlans = "1,3,5-10"` + "\n"
	config += `	dai_trusted = false` + "\n"
	config += `	enabled = true` + "\n"
	config += `	isolation_enabled = false` + "\n"
	config += `	link_negotiation = "Auto negotiate"` + "\n"
	config += `	name = "My switch port"` + "\n"
	config += `	poe_enabled = true` + "\n"
	config += `	rstp_enabled = true` + "\n"
	config += `	sticky_mac_allow_list_limit = 5` + "\n"
	config += `	stp_guard = "disabled"` + "\n"
	config += `	type = "access"` + "\n"
	config += `	udld = "Alert only"` + "\n"
	config += `	vlan = 10` + "\n"
	config += `	voice_vlan = 20` + "\n"
	config += `	profile_enabled = false` + "\n"
	config += `	sticky_mac_allow_list = ["34:56:fe:ce:8e:b0"]` + "\n"
	config += `	tags = ["tag1"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
