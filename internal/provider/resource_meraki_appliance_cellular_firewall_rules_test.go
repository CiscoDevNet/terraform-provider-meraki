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

	goversion "github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccMerakiApplianceCellularFirewallRules(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_cellular_firewall_rules.test", "rules.0.comment", "Allow TCP traffic to subnet with HTTP servers."))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_cellular_firewall_rules.test", "rules.0.dest_cidr", "192.168.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_cellular_firewall_rules.test", "rules.0.dest_port", "443"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_cellular_firewall_rules.test", "rules.0.policy", "allow"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_cellular_firewall_rules.test", "rules.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_cellular_firewall_rules.test", "rules.0.src_cidr", "Any"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_cellular_firewall_rules.test", "rules.0.src_port", "Any"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_cellular_firewall_rules.test", "rules.0.syslog_enabled", "false"))

	var steps []resource.TestStep
	var tfVersion *goversion.Version
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiApplianceCellularFirewallRulesPrerequisitesConfig + testAccMerakiApplianceCellularFirewallRulesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiApplianceCellularFirewallRulesPrerequisitesConfig + testAccMerakiApplianceCellularFirewallRulesConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_appliance_cellular_firewall_rules.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiApplianceCellularFirewallRulesImportStateIdFunc("meraki_appliance_cellular_firewall_rules.test"),
		ImportStateVerifyIgnore: []string{},
		Check:                   resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiApplianceCellularFirewallRulesPrerequisitesConfig + testAccApplianceCellularFirewallRulesConfigAdditional0,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			terraformVersionCapture{Version: &tfVersion},
		},
		Steps: steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiApplianceCellularFirewallRulesImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiApplianceCellularFirewallRulesPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["appliance"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiApplianceCellularFirewallRulesConfig_minimum() string {
	config := `resource "meraki_appliance_cellular_firewall_rules" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  rules = [{` + "\n"
	config += `    dest_cidr = "192.168.1.0/24"` + "\n"
	config += `    dest_port = "443"` + "\n"
	config += `    policy = "allow"` + "\n"
	config += `    protocol = "tcp"` + "\n"
	config += `    src_cidr = "Any"` + "\n"
	config += `    src_port = "Any"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccMerakiApplianceCellularFirewallRulesConfig_all() string {
	config := `resource "meraki_appliance_cellular_firewall_rules" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  rules = [{` + "\n"
	config += `    comment = "Allow TCP traffic to subnet with HTTP servers."` + "\n"
	config += `    dest_cidr = "192.168.1.0/24"` + "\n"
	config += `    dest_port = "443"` + "\n"
	config += `    policy = "allow"` + "\n"
	config += `    protocol = "tcp"` + "\n"
	config += `    src_cidr = "Any"` + "\n"
	config += `    src_port = "Any"` + "\n"
	config += `    syslog_enabled = false` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

const testAccApplianceCellularFirewallRulesConfigAdditional0 = `
resource "meraki_organization_policy_object" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "test_policy_object"
  category        = "network"
  type            = "cidr"
  cidr            = "10.10.10.1"
}
resource "meraki_organization_policy_object_group" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "test_policy_group"
  category        = "NetworkObjectGroup"
  object_ids      = [meraki_organization_policy_object.test.id]
}
resource "meraki_appliance_cellular_firewall_rules" "test" {
  network_id = meraki_network.test.id
  rules = [{
    comment        = "Policy object based rule"
    dest_cidr      = "GRP(${meraki_organization_policy_object_group.test.id}),OBJ(${meraki_organization_policy_object.test.id})"
    dest_port      = "443"
    policy         = "allow"
    protocol       = "tcp"
    src_cidr       = "Any"
    src_port       = "Any"
    syslog_enabled = false
  }]
}
`

// End of section. //template:end testAccConfigAdditional
