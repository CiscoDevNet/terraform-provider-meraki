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

func TestAccMerakiWirelessSSIDL7FirewallRules(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_l7_firewall_rules.test", "rules.0.policy", "deny"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_l7_firewall_rules.test", "rules.0.type", "host"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_l7_firewall_rules.test", "rules.0.value", "google.com"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessSSIDL7FirewallRulesPrerequisitesConfig + testAccMerakiWirelessSSIDL7FirewallRulesConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_wireless_ssid_l7_firewall_rules.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiWirelessSSIDL7FirewallRulesImportStateIdFunc("meraki_wireless_ssid_l7_firewall_rules.test"),
		ImportStateVerifyIgnore: []string{},
		Check:                   resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessSSIDL7FirewallRulesPrerequisitesConfig + testAccWirelessSSIDL7FirewallRulesConfigAdditional0,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiWirelessSSIDL7FirewallRulesImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		Number := primary.Attributes["number"]

		return fmt.Sprintf("%s,%s", NetworkId, Number), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiWirelessSSIDL7FirewallRulesPrerequisitesConfig = `
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
resource "meraki_wireless_ssid" "test" {
  network_id = meraki_network.test.id
  number     = "0"
  name       = "My SSID"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiWirelessSSIDL7FirewallRulesConfig_minimum() string {
	config := `resource "meraki_wireless_ssid_l7_firewall_rules" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  number = meraki_wireless_ssid.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiWirelessSSIDL7FirewallRulesConfig_all() string {
	config := `resource "meraki_wireless_ssid_l7_firewall_rules" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  number = meraki_wireless_ssid.test.id` + "\n"
	config += `  rules = [{` + "\n"
	config += `    policy = "deny"` + "\n"
	config += `    type = "host"` + "\n"
	config += `    value = "google.com"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

const testAccWirelessSSIDL7FirewallRulesConfigAdditional0 = `
resource "meraki_wireless_ssid_l7_firewall_rules" "test" {
  network_id = meraki_network.test.id
  number     = meraki_wireless_ssid.test.id
  rules = [
    {
      policy = "deny"
      type   = "host"
      value  = "google.com"
    },
    {
      policy = "deny"
      type   = "port"
      value  = "9091"
    },
    {
      policy = "deny"
      type   = "ipRange"
      value  = "192.168.0.3/16:80"
    },
    {
      policy = "deny"
      type   = "application"
      value  = "meraki:layer7/application/171"
    },
    {
      policy = "deny"
      type   = "applicationCategory"
      value  = "meraki:layer7/category/17"
    }
  ]
}
`

// End of section. //template:end testAccConfigAdditional
