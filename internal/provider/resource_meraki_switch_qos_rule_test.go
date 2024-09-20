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

func TestAccMerakiSwitchQoSRule(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_qos_rule.test", "dscp", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_qos_rule.test", "dst_port_range", "3000-3100"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_qos_rule.test", "protocol", "TCP"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_qos_rule.test", "src_port", "2000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_qos_rule.test", "vlan", "100"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchQoSRulePrerequisitesConfig + testAccMerakiSwitchQoSRuleConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchQoSRulePrerequisitesConfig + testAccMerakiSwitchQoSRuleConfig_all(),
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

const testAccMerakiSwitchQoSRulePrerequisitesConfig = `
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

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiSwitchQoSRuleConfig_minimum() string {
	config := `resource "meraki_switch_qos_rule" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	vlan = 100` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchQoSRuleConfig_all() string {
	config := `resource "meraki_switch_qos_rule" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	dscp = 0` + "\n"
	config += `	dst_port_range = "3000-3100"` + "\n"
	config += `	protocol = "TCP"` + "\n"
	config += `	src_port = 2000` + "\n"
	config += `	vlan = 100` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
