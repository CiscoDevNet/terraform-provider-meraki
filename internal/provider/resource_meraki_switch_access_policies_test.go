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

func TestAccMerakiSwitchAccessPolicies(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchAccessPoliciesPrerequisitesConfig + testAccMerakiSwitchAccessPoliciesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchAccessPoliciesPrerequisitesConfig + testAccMerakiSwitchAccessPoliciesConfig_all(),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:      "meraki_switch_access_policies.test",
		ImportState:       true,
		ImportStateIdFunc: merakiSwitchAccessPoliciesImportStateIdFunc("meraki_switch_access_policies.test"),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiSwitchAccessPoliciesImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s,%s", OrganizationId, NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiSwitchAccessPoliciesPrerequisitesConfig = `
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

func testAccMerakiSwitchAccessPoliciesConfig_minimum() string {
	config := `resource "meraki_switch_access_policies" "test" {` + "\n"
	config += ` network_id = meraki_network.test.id` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  host_mode = "Single-Host"` + "\n"
	config += `  name = "Access policy #1"` + "\n"
	config += `  radius_accounting_enabled = false` + "\n"
	config += `  radius_coa_support_enabled = false` + "\n"
	config += `  radius_testing_enabled = true` + "\n"
	config += `  url_redirect_walled_garden_enabled = false` + "\n"
	config += `  radius_servers = [{` + "\n"
	config += `    host = "1.2.3.4"` + "\n"
	config += `    port = 22` + "\n"
	config += `    secret = "secret"` + "\n"
	config += `  }]` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchAccessPoliciesConfig_all() string {
	config := `resource "meraki_switch_access_policies" "test" {` + "\n"
	config += ` network_id = meraki_network.test.id` + "\n"
	config += ` organization_id = data.meraki_organization.test.id` + "\n"
	config += ` items = [{` + "\n"
	config += `  access_policy_type = "Hybrid authentication"` + "\n"
	config += `  guest_port_bouncing = true` + "\n"
	config += `  guest_vlan_id = 100` + "\n"
	config += `  host_mode = "Single-Host"` + "\n"
	config += `  increase_access_speed = false` + "\n"
	config += `  name = "Access policy #1"` + "\n"
	config += `  radius_accounting_enabled = true` + "\n"
	config += `  radius_coa_support_enabled = false` + "\n"
	config += `  radius_group_attribute = "11"` + "\n"
	config += `  radius_testing_enabled = true` + "\n"
	config += `  url_redirect_walled_garden_enabled = true` + "\n"
	config += `  voice_vlan_clients = true` + "\n"
	config += `  dot1x_control_direction = "inbound"` + "\n"
	config += `  radius_failed_auth_vlan_id = 100` + "\n"
	config += `  radius_re_authentication_interval = 120` + "\n"
	config += `  radius_cache_enabled = false` + "\n"
	config += `  radius_cache_timeout = 24` + "\n"
	config += `  radius_critical_auth_data_vlan_id = 100` + "\n"
	config += `  radius_critical_auth_suspend_port_bounce = true` + "\n"
	config += `  radius_critical_auth_voice_vlan_id = 101` + "\n"
	config += `  radius_accounting_servers = [{` + "\n"
	config += `    host = "1.2.3.4"` + "\n"
	config += `    port = 22` + "\n"
	config += `    secret = "secret"` + "\n"
	config += `  }]` + "\n"
	config += `  radius_servers = [{` + "\n"
	config += `    host = "1.2.3.4"` + "\n"
	config += `    port = 22` + "\n"
	config += `    secret = "secret"` + "\n"
	config += `  }]` + "\n"
	config += `  url_redirect_walled_garden_ranges = ["192.168.1.0/24"]` + "\n"
	config += ` }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
