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

func TestAccDataSourceMerakiDeviceCellularSIMs(t *testing.T) {
	if os.Getenv("DEVICE_CELLULAR_SIMS") == "" {
		t.Skip("skipping test, set environment variable DEVICE_CELLULAR_SIMS")
	}
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_switch_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_switch_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_cellular_sims.test", "sim_failover_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_cellular_sims.test", "sim_failover_timeout", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_cellular_sims.test", "sim_ordering.0", "sim1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_cellular_sims.test", "sims.0.is_primary", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_cellular_sims.test", "sims.0.sim_order", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_cellular_sims.test", "sims.0.slot", "sim1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_cellular_sims.test", "sims.0.apns.0.name", "internet"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_cellular_sims.test", "sims.0.apns.0.authentication_password", "secret"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_cellular_sims.test", "sims.0.apns.0.authentication_type", "pap"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_cellular_sims.test", "sims.0.apns.0.authentication_username", "milesmeraki"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_device_cellular_sims.test", "sims.0.apns.0.allowed_ip_types.0", "ipv4"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiDeviceCellularSIMsPrerequisitesConfig + testAccDataSourceMerakiDeviceCellularSIMsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiDeviceCellularSIMsPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_switch_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_switch_1_serial]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiDeviceCellularSIMsConfig() string {
	config := `resource "meraki_device_cellular_sims" "test" {` + "\n"
	config += `	serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `	sim_failover_enabled = true` + "\n"
	config += `	sim_failover_timeout = 300` + "\n"
	config += `	sim_ordering = ["sim1"]` + "\n"
	config += `	sims = [{` + "\n"
	config += `		is_primary = false` + "\n"
	config += `		sim_order = 3` + "\n"
	config += `		slot = "sim1"` + "\n"
	config += `		apns = [{` + "\n"
	config += `			name = "internet"` + "\n"
	config += `			authentication_password = "secret"` + "\n"
	config += `			authentication_type = "pap"` + "\n"
	config += `			authentication_username = "milesmeraki"` + "\n"
	config += `			allowed_ip_types = ["ipv4"]` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_device_cellular_sims" "test" {
			serial = tolist(meraki_network_device_claim.test.serials)[0]
			depends_on = [meraki_device_cellular_sims.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig