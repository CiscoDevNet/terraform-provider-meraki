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

func TestAccDataSourceMerakiApplianceVPNSiteToSiteHubVRFs(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_appliance_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_appliance_1_serial")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiApplianceVPNSiteToSiteHubVRFsPrerequisitesConfig + testAccDataSourceMerakiApplianceVPNSiteToSiteHubVRFsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiApplianceVPNSiteToSiteHubVRFsPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_appliance_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_organization_early_access_features_opt_in" "test_vrf" {
  organization_id = data.meraki_organization.test.id
  short_name      = "has_mx_vrf_26_1_early_access"
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless", "appliance", "sensor", "camera"]
}
resource "meraki_network" "test_spoke" {
  organization_id = data.meraki_organization.test.id
  name            = "spoke test network"
  product_types   = ["switch", "wireless", "appliance", "sensor", "camera"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_appliance_1_serial]
}
resource "meraki_appliance_routing_vrfs_settings" "test" {
  organization_id = data.meraki_organization.test.id
  enabled         = true
}
resource "meraki_appliance_site_to_site_vpn" "test" {
  network_id = meraki_network.test.id
  mode       = "hub"
}
resource "meraki_appliance_site_to_site_vpn" "test_spoke" {
  network_id = meraki_network.test_spoke.id
  mode       = "spoke"
  hubs       = [
    {
      hub_id = meraki_network.test.id
    }
  ]
}
resource "meraki_routing_vrf" "test" {
  organization_id   = data.meraki_organization.test.id
  name              = "test-vrf"
  appliance_enabled = true
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiApplianceVPNSiteToSiteHubVRFsConfig() string {
	config := `resource "meraki_appliance_vpn_site_to_site_hub_vrfs" "test" {` + "\n"
	config += `  network_id = meraki_network.test_spoke.id` + "\n"
	config += `  hub_network_id = meraki_network.test.id` + "\n"
	config += `  vrfs = [{` + "\n"
	config += `    id = meraki_routing_vrf.test.id` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_appliance_vpn_site_to_site_hub_vrfs" "test" {
			network_id = meraki_network.test_spoke.id
			hub_network_id = meraki_network.test.id
			depends_on = [meraki_appliance_vpn_site_to_site_hub_vrfs.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
