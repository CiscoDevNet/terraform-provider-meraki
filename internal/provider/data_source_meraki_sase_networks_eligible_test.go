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

func TestAccDataSourceMerakiSaseNetworksEligible(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sase_networks_eligible.test", "meta_counts_items_remaining", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sase_networks_eligible.test", "meta_counts_items_total", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sase_networks_eligible.test", "items.0.name", "London Office"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sase_networks_eligible.test", "items.0.network_id", "N_123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sase_networks_eligible.test", "items.0.type", "Meraki spoke"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sase_networks_eligible.test", "items.0.address_street", "123 Main St"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sase_networks_eligible.test", "items.0.device_primary_model", "MX95"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sase_networks_eligible.test", "items.0.region_name", "US East"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sase_networks_eligible.test", "items.0.routing_default_route_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_sase_networks_eligible.test", "items.0.vpn_type", "hub"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiSaseNetworksEligiblePrerequisitesConfig + testAccDataSourceMerakiSaseNetworksEligibleConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiSaseNetworksEligiblePrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiSaseNetworksEligibleConfig() string {
	config := `resource "meraki_sase_networks_eligible" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  meta_counts_items_remaining = 0` + "\n"
	config += `  meta_counts_items_total = 1` + "\n"
	config += `  items = [{` + "\n"
	config += `    name = "London Office"` + "\n"
	config += `    network_id = "N_123"` + "\n"
	config += `    type = "Meraki spoke"` + "\n"
	config += `    address_street = "123 Main St"` + "\n"
	config += `    device_primary_model = "MX95"` + "\n"
	config += `    region_name = "US East"` + "\n"
	config += `    routing_default_route_enabled = true` + "\n"
	config += `    vpn_type = "hub"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_sase_networks_eligible" "test" {
			id = meraki_sase_networks_eligible.test.id
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_sase_networks_eligible.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
