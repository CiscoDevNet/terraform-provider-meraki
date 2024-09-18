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

func TestAccMerakiNetworkClientPolicy(t *testing.T) {
	if os.Getenv("NETWORK_CLIENT") == "" {
		t.Skip("skipping test, set environment variable NETWORK_CLIENT")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_client_policy.test", "device_policy", "Group policy"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_network_client_policy.test", "group_policy_id", "101"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiNetworkClientPolicyPrerequisitesConfig + testAccMerakiNetworkClientPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiNetworkClientPolicyPrerequisitesConfig + testAccMerakiNetworkClientPolicyConfig_all(),
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

const testAccMerakiNetworkClientPolicyPrerequisitesConfig = `
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

func testAccMerakiNetworkClientPolicyConfig_minimum() string {
	config := `resource "meraki_network_client_policy" "test" {` + "\n"
	config += `	network_id = meraki_network_vlan_profile.default.network_id` + "\n"
	config += `	client_id = 1.2.3.4` + "\n"
	config += `	device_policy = "Group policy"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiNetworkClientPolicyConfig_all() string {
	config := `resource "meraki_network_client_policy" "test" {` + "\n"
	config += `	network_id = meraki_network_vlan_profile.default.network_id` + "\n"
	config += `	client_id = 1.2.3.4` + "\n"
	config += `	device_policy = "Group policy"` + "\n"
	config += `	group_policy_id = "101"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll