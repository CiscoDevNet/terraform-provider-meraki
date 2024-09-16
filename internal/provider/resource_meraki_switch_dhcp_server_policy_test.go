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

func TestAccMerakiSwitchDHCPServerPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_dhcp_server_policy.test", "default_policy", "block"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_dhcp_server_policy.test", "alerts_email_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_dhcp_server_policy.test", "arp_inspection_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_dhcp_server_policy.test", "allowed_servers.0", "00:50:56:00:00:01"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_switch_dhcp_server_policy.test", "blocked_servers.0", "00:50:56:00:00:03"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiSwitchDHCPServerPolicyPrerequisitesConfig + testAccMerakiSwitchDHCPServerPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiSwitchDHCPServerPolicyPrerequisitesConfig + testAccMerakiSwitchDHCPServerPolicyConfig_all(),
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

const testAccMerakiSwitchDHCPServerPolicyPrerequisitesConfig = `
data "meraki_organization" "test" {
  name = "Dev"
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = "Network1"
  product_types   = ["switch", "wireless"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiSwitchDHCPServerPolicyConfig_minimum() string {
	config := `resource "meraki_switch_dhcp_server_policy" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	default_policy = "allow"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiSwitchDHCPServerPolicyConfig_all() string {
	config := `resource "meraki_switch_dhcp_server_policy" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	default_policy = "block"` + "\n"
	config += `	alerts_email_enabled = true` + "\n"
	config += `	arp_inspection_enabled = true` + "\n"
	config += `	allowed_servers = ["00:50:56:00:00:01"]` + "\n"
	config += `	blocked_servers = ["00:50:56:00:00:03"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll