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

func TestAccMerakiNetworkGroupPolicy(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "name", "No video streaming"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "splash_auth_settings", "bypass"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "bandwidth_settings", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "bandwidth_limit_down", "1000000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "bandwidth_limit_up", "1000000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "bonjour_forwarding_settings", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "bonjour_forwarding_rules.0.description", "A simple bonjour rule"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "bonjour_forwarding_rules.0.vlan_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "firewall_and_traffic_shaping_settings", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "l3_firewall_rules.0.comment", "Allow TCP traffic to subnet with HTTP servers."))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "l3_firewall_rules.0.dest_cidr", "192.168.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "l3_firewall_rules.0.dest_port", "443"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "l3_firewall_rules.0.policy", "allow"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "l3_firewall_rules.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "l7_firewall_rules.0.policy", "deny"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "l7_firewall_rules.0.type", "host"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "l7_firewall_rules.0.value", "google.com"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "traffic_shaping_rules.0.dscp_tag_value", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "traffic_shaping_rules.0.pcp_tag_value", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "traffic_shaping_rules.0.per_client_bandwidth_limits_settings", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "traffic_shaping_rules.0.per_client_bandwidth_limits_bandwidth_limits_limit_down", "1000000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "traffic_shaping_rules.0.per_client_bandwidth_limits_bandwidth_limits_limit_up", "1000000"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "traffic_shaping_rules.0.definitions.0.type", "host"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "traffic_shaping_rules.0.definitions.0.value", "google.com"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_friday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_friday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_friday_to", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_monday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_monday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_monday_to", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_saturday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_saturday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_saturday_to", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_sunday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_sunday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_sunday_to", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_thursday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_thursday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_thursday_to", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_tuesday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_tuesday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_tuesday_to", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_wednesday_active", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_wednesday_from", "09:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "scheduling_wednesday_to", "17:00"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "vlan_tagging_settings", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "vlan_tagging_vlan_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_group_policy.test", "force_delete", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiNetworkGroupPolicyPrerequisitesConfig + testAccMerakiNetworkGroupPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkGroupPolicyPrerequisitesConfig + testAccMerakiNetworkGroupPolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_network_group_policy.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiNetworkGroupPolicyImportStateIdFunc("meraki_network_group_policy.test"),
		ImportStateVerifyIgnore: []string{"traffic_shaping_rules.0.priority"},
		Check:                   resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkGroupPolicyPrerequisitesConfig + testAccNetworkGroupPolicyConfigAdditional0,
	})
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkGroupPolicyPrerequisitesConfig + testAccNetworkGroupPolicyConfigAdditional1,
	})
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkGroupPolicyPrerequisitesConfig + testAccNetworkGroupPolicyConfigAdditional2,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiNetworkGroupPolicyImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]
		ForceDelete := primary.Attributes["force_delete"]
		Id := primary.Attributes["id"]

		return fmt.Sprintf("%s,%s,%s", NetworkId, ForceDelete, Id), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiNetworkGroupPolicyPrerequisitesConfig = `
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

func testAccMerakiNetworkGroupPolicyConfig_minimum() string {
	config := `resource "meraki_network_group_policy" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  name = "No video streaming"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiNetworkGroupPolicyConfig_all() string {
	config := `resource "meraki_network_group_policy" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  name = "No video streaming"` + "\n"
	config += `  splash_auth_settings = "bypass"` + "\n"
	config += `  bandwidth_settings = "custom"` + "\n"
	config += `  bandwidth_limit_down = 1000000` + "\n"
	config += `  bandwidth_limit_up = 1000000` + "\n"
	config += `  bonjour_forwarding_settings = "custom"` + "\n"
	config += `  bonjour_forwarding_rules = [{` + "\n"
	config += `    description = "A simple bonjour rule"` + "\n"
	config += `    vlan_id = "1"` + "\n"
	config += `    services = ["All Services"]` + "\n"
	config += `  }]` + "\n"
	config += `  firewall_and_traffic_shaping_settings = "custom"` + "\n"
	config += `  l3_firewall_rules = [{` + "\n"
	config += `    comment = "Allow TCP traffic to subnet with HTTP servers."` + "\n"
	config += `    dest_cidr = "192.168.1.0/24"` + "\n"
	config += `    dest_port = "443"` + "\n"
	config += `    policy = "allow"` + "\n"
	config += `    protocol = "tcp"` + "\n"
	config += `  }]` + "\n"
	config += `  l7_firewall_rules = [{` + "\n"
	config += `    policy = "deny"` + "\n"
	config += `    type = "host"` + "\n"
	config += `    value = "google.com"` + "\n"
	config += `  }]` + "\n"
	config += `  traffic_shaping_rules = [{` + "\n"
	config += `    dscp_tag_value = 0` + "\n"
	config += `    pcp_tag_value = 0` + "\n"
	config += `    priority = "normal"` + "\n"
	config += `    per_client_bandwidth_limits_settings = "custom"` + "\n"
	config += `    per_client_bandwidth_limits_bandwidth_limits_limit_down = 1000000` + "\n"
	config += `    per_client_bandwidth_limits_bandwidth_limits_limit_up = 1000000` + "\n"
	config += `    definitions = [{` + "\n"
	config += `      type = "host"` + "\n"
	config += `      value = "google.com"` + "\n"
	config += `    }]` + "\n"
	config += `  }]` + "\n"
	config += `  scheduling_enabled = true` + "\n"
	config += `  scheduling_friday_active = true` + "\n"
	config += `  scheduling_friday_from = "09:00"` + "\n"
	config += `  scheduling_friday_to = "17:00"` + "\n"
	config += `  scheduling_monday_active = true` + "\n"
	config += `  scheduling_monday_from = "09:00"` + "\n"
	config += `  scheduling_monday_to = "17:00"` + "\n"
	config += `  scheduling_saturday_active = true` + "\n"
	config += `  scheduling_saturday_from = "09:00"` + "\n"
	config += `  scheduling_saturday_to = "17:00"` + "\n"
	config += `  scheduling_sunday_active = true` + "\n"
	config += `  scheduling_sunday_from = "09:00"` + "\n"
	config += `  scheduling_sunday_to = "17:00"` + "\n"
	config += `  scheduling_thursday_active = true` + "\n"
	config += `  scheduling_thursday_from = "09:00"` + "\n"
	config += `  scheduling_thursday_to = "17:00"` + "\n"
	config += `  scheduling_tuesday_active = true` + "\n"
	config += `  scheduling_tuesday_from = "09:00"` + "\n"
	config += `  scheduling_tuesday_to = "17:00"` + "\n"
	config += `  scheduling_wednesday_active = true` + "\n"
	config += `  scheduling_wednesday_from = "09:00"` + "\n"
	config += `  scheduling_wednesday_to = "17:00"` + "\n"
	config += `  vlan_tagging_settings = "custom"` + "\n"
	config += `  vlan_tagging_vlan_id = "1"` + "\n"
	config += `  force_delete = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

const testAccNetworkGroupPolicyConfigAdditional0 = `
resource "meraki_network_group_policy" "test" {
  network_id   = meraki_network.test.id
  name         = "No video streaming"
  force_delete = true
}
`

const testAccNetworkGroupPolicyConfigAdditional1 = `
resource "meraki_network_group_policy" "test" {
  network_id = meraki_network.test.id
  name       = "No video streaming"
  l7_firewall_rules = [
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

const testAccNetworkGroupPolicyConfigAdditional2 = `
resource "meraki_network_group_policy" "test" {
  network_id = meraki_network.test.id
  name       = "No video streaming"
  traffic_shaping_rules = [
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
          type  = "port"
          value = "9091"
        },
        {
          type  = "applicationCategory"
          value = "meraki:layer7/category/27"
        },
        {
          type  = "application"
          value = "meraki:layer7/application/133"
        },
        {
          type  = "ipRange"
          value = "192.168.0.3/16:80"
        },
        {
          type  = "localNet"
          value = "192.168.2.4/16"
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
