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

func TestAccMerakiApplianceOneToManyNATRules(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_one_to_many_nat_rules.test", "rules.0.public_ip", "146.11.11.13"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_one_to_many_nat_rules.test", "rules.0.uplink", "internet1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_one_to_many_nat_rules.test", "rules.0.port_rules.0.local_ip", "192.168.128.1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_one_to_many_nat_rules.test", "rules.0.port_rules.0.local_port", "443"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_one_to_many_nat_rules.test", "rules.0.port_rules.0.name", "Rule 1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_one_to_many_nat_rules.test", "rules.0.port_rules.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_one_to_many_nat_rules.test", "rules.0.port_rules.0.public_port", "9443"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_one_to_many_nat_rules.test", "rules.0.port_rules.0.allowed_ips.0", "any"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiApplianceOneToManyNATRulesPrerequisitesConfig + testAccMerakiApplianceOneToManyNATRulesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiApplianceOneToManyNATRulesPrerequisitesConfig + testAccMerakiApplianceOneToManyNATRulesConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_appliance_one_to_many_nat_rules.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiApplianceOneToManyNATRulesImportStateIdFunc("meraki_appliance_one_to_many_nat_rules.test"),
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

func merakiApplianceOneToManyNATRulesImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiApplianceOneToManyNATRulesPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiApplianceOneToManyNATRulesConfig_minimum() string {
	config := `resource "meraki_appliance_one_to_many_nat_rules" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	rules = [{` + "\n"
	config += `		public_ip = "146.11.11.13"` + "\n"
	config += `		uplink = "internet1"` + "\n"
	config += `		port_rules = [{` + "\n"
	config += `		local_ip = "192.168.128.1"` + "\n"
	config += `		local_port = "443"` + "\n"
	config += `		protocol = "tcp"` + "\n"
	config += `		public_port = "9443"` + "\n"
	config += `		allowed_ips = ["any"]` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiApplianceOneToManyNATRulesConfig_all() string {
	config := `resource "meraki_appliance_one_to_many_nat_rules" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	rules = [{` + "\n"
	config += `		public_ip = "146.11.11.13"` + "\n"
	config += `		uplink = "internet1"` + "\n"
	config += `		port_rules = [{` + "\n"
	config += `			local_ip = "192.168.128.1"` + "\n"
	config += `			local_port = "443"` + "\n"
	config += `			name = "Rule 1"` + "\n"
	config += `			protocol = "tcp"` + "\n"
	config += `			public_port = "9443"` + "\n"
	config += `			allowed_ips = ["any"]` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
