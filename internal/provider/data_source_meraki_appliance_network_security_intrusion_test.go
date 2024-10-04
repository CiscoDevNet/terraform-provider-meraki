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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceMerakiApplianceNetworkSecurityIntrusion(t *testing.T) {
	if os.Getenv("APPLIANCE_NETWORK_SECURITY_INTRUSION") == "" {
		t.Skip("skipping test, set environment variable APPLIANCE_NETWORK_SECURITY_INTRUSION")
	}
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_network_security_intrusion.test", "ids_rulesets", "balanced"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_network_security_intrusion.test", "mode", "prevention"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_network_security_intrusion.test", "protected_networks_use_default", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_network_security_intrusion.test", "protected_networks_excluded_cidr.0", "10.0.0.0/8"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_network_security_intrusion.test", "protected_networks_included_cidr.0", "10.0.0.0/8"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiApplianceNetworkSecurityIntrusionPrerequisitesConfig + testAccDataSourceMerakiApplianceNetworkSecurityIntrusionConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiApplianceNetworkSecurityIntrusionPrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiApplianceNetworkSecurityIntrusionConfig() string {
	config := `resource "meraki_appliance_network_security_intrusion" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	ids_rulesets = "balanced"` + "\n"
	config += `	mode = "prevention"` + "\n"
	config += `	protected_networks_use_default = false` + "\n"
	config += `	protected_networks_excluded_cidr = ["10.0.0.0/8"]` + "\n"
	config += `	protected_networks_included_cidr = ["10.0.0.0/8"]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_appliance_network_security_intrusion" "test" {
			network_id = meraki_network.test.id
			depends_on = [meraki_appliance_network_security_intrusion.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
