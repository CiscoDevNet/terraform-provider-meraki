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

func TestAccMerakiApplianceConnectivityMonitoringDestinations(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_connectivity_monitoring_destinations.test", "destinations.0.default", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_connectivity_monitoring_destinations.test", "destinations.0.description", "Google"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_connectivity_monitoring_destinations.test", "destinations.0.ip", "1.2.3.4"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiApplianceConnectivityMonitoringDestinationsPrerequisitesConfig + testAccMerakiApplianceConnectivityMonitoringDestinationsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiApplianceConnectivityMonitoringDestinationsPrerequisitesConfig + testAccMerakiApplianceConnectivityMonitoringDestinationsConfig_all(),
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

const testAccMerakiApplianceConnectivityMonitoringDestinationsPrerequisitesConfig = `
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

func testAccMerakiApplianceConnectivityMonitoringDestinationsConfig_minimum() string {
	config := `resource "meraki_appliance_connectivity_monitoring_destinations" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	destinations = [{` + "\n"
	config += `		default = true` + "\n"
	config += `		ip = "1.2.3.4"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiApplianceConnectivityMonitoringDestinationsConfig_all() string {
	config := `resource "meraki_appliance_connectivity_monitoring_destinations" "test" {` + "\n"
	config += `	network_id = meraki_network.test.id` + "\n"
	config += `	destinations = [{` + "\n"
	config += `		default = true` + "\n"
	config += `		description = "Google"` + "\n"
	config += `		ip = "1.2.3.4"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll