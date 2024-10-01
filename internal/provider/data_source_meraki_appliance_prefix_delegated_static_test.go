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

func TestAccDataSourceMerakiAppliancePrefixDelegatedStatic(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_prefix_delegated_static.test", "description", "Prefix on WAN 1 of Long Island Office network"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_prefix_delegated_static.test", "prefix", "2002:db8:3c4d:15::/64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_prefix_delegated_static.test", "origin_type", "internet"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_prefix_delegated_static.test", "origin_interfaces.0", "wan1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiAppliancePrefixDelegatedStaticPrerequisitesConfig + testAccDataSourceMerakiAppliancePrefixDelegatedStaticConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiAppliancePrefixDelegatedStaticPrerequisitesConfig = `
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

func testAccDataSourceMerakiAppliancePrefixDelegatedStaticConfig() string {
	config := `resource "meraki_appliance_prefix_delegated_static" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	description = "Prefix on WAN 1 of Long Island Office network"` + "\n"
	config += `	prefix = "2002:db8:3c4d:15::/64"` + "\n"
	config += `	origin_type = "internet"` + "\n"
	config += `	origin_interfaces = ["wan1"]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_appliance_prefix_delegated_static" "test" {
			id = meraki_appliance_prefix_delegated_static.test.id
			network_id = meraki_network.test.id
			depends_on = [meraki_appliance_prefix_delegated_static.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
