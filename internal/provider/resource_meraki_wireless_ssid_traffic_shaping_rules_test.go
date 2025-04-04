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

func TestAccMerakiWirelessSSIDTrafficShapingRules(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_traffic_shaping_rules.test", "default_rules_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_traffic_shaping_rules.test", "traffic_shaping_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_traffic_shaping_rules.test", "rules.0.dscp_tag_value", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_traffic_shaping_rules.test", "rules.0.pcp_tag_value", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_traffic_shaping_rules.test", "rules.0.per_client_bandwidth_limits_settings", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_traffic_shaping_rules.test", "rules.0.per_client_bandwidth_limits_bandwidth_limits_limit_down", "1000000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_traffic_shaping_rules.test", "rules.0.per_client_bandwidth_limits_bandwidth_limits_limit_up", "1000000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_traffic_shaping_rules.test", "rules.0.definitions.0.type", "host"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_wireless_ssid_traffic_shaping_rules.test", "rules.0.definitions.0.value", "google.com"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiWirelessSSIDTrafficShapingRulesPrerequisitesConfig + testAccMerakiWirelessSSIDTrafficShapingRulesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessSSIDTrafficShapingRulesPrerequisitesConfig + testAccMerakiWirelessSSIDTrafficShapingRulesConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_wireless_ssid_traffic_shaping_rules.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiWirelessSSIDTrafficShapingRulesImportStateIdFunc("meraki_wireless_ssid_traffic_shaping_rules.test"),
		ImportStateVerifyIgnore: []string{},
		Check:                   resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiWirelessSSIDTrafficShapingRulesPrerequisitesConfig + testAccWirelessSSIDTrafficShapingRulesConfigAdditional0,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiWirelessSSIDTrafficShapingRulesImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		Number := primary.Attributes["number"]

		return fmt.Sprintf("%s,%s", NetworkId, Number), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiWirelessSSIDTrafficShapingRulesPrerequisitesConfig = `
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

func testAccMerakiWirelessSSIDTrafficShapingRulesConfig_minimum() string {
	config := `resource "meraki_wireless_ssid_traffic_shaping_rules" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  number = meraki_wireless_ssid.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiWirelessSSIDTrafficShapingRulesConfig_all() string {
	config := `resource "meraki_wireless_ssid_traffic_shaping_rules" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  number = meraki_wireless_ssid.test.id` + "\n"
	config += `  default_rules_enabled = true` + "\n"
	config += `  traffic_shaping_enabled = true` + "\n"
	config += `  rules = [{` + "\n"
	config += `    dscp_tag_value = 0` + "\n"
	config += `    pcp_tag_value = 0` + "\n"
	config += `    per_client_bandwidth_limits_settings = "custom"` + "\n"
	config += `    per_client_bandwidth_limits_bandwidth_limits_limit_down = 1000000` + "\n"
	config += `    per_client_bandwidth_limits_bandwidth_limits_limit_up = 1000000` + "\n"
	config += `    definitions = [{` + "\n"
	config += `      type = "host"` + "\n"
	config += `      value = "google.com"` + "\n"
	config += `    }]` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

const testAccWirelessSSIDTrafficShapingRulesConfigAdditional0 = `
resource "meraki_wireless_ssid_traffic_shaping_rules" "test" {
  network_id = meraki_network.test.id
  number = meraki_wireless_ssid.test.id
  rules = [
    {
      definitions = [
        {
          type  = "application"
          value = "meraki:layer7/application/171"
        },
        {
          type  = "host"
          value = "google.com"
        },
        {
          type  = "application"
          value = "meraki:layer7/application/132"
        },
        {
          type  = "application"
          value = "meraki:layer7/application/133"
        },
        {
          type  = "host"
          value = "test.com"
        }
      ]
    }
  ]
}
`

// End of section. //template:end testAccConfigAdditional
