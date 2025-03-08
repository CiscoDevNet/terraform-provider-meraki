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

func TestAccMerakiApplianceSDWANInternetPolicies(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.fail_over_criterion", "poorPerformance"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.preferred_uplink", "wan1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.builtin_performance_class_name", "VoIP"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.custom_performance_class_id", "123456"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.performance_class_type", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.type", "custom"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.destination_cidr", "any"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.destination_port", "any"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.destination_applications.0.id", "meraki:layer7/application/3"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.destination_applications.0.name", "DNS"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.destination_applications.0.type", "major"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.source_cidr", "any"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.source_host", "254"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.source_port", "1-1024"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_appliance_sdwan_internet_policies.test", "wan_traffic_uplink_preferences.0.traffic_filters.0.source_vlan", "10"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiApplianceSDWANInternetPoliciesPrerequisitesConfig + testAccMerakiApplianceSDWANInternetPoliciesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiApplianceSDWANInternetPoliciesPrerequisitesConfig + testAccMerakiApplianceSDWANInternetPoliciesConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiApplianceSDWANInternetPoliciesImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		NetworkId := primary.Attributes["network_id"]

		return fmt.Sprintf("%s", NetworkId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiApplianceSDWANInternetPoliciesPrerequisitesConfig = `
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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiApplianceSDWANInternetPoliciesConfig_minimum() string {
	config := `resource "meraki_appliance_sdwan_internet_policies" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  wan_traffic_uplink_preferences = [{` + "\n"
	config += `    preferred_uplink = "wan1"` + "\n"
	config += `    traffic_filters = [{` + "\n"
	config += `    type = "custom"` + "\n"
	config += `    protocol = "any"` + "\n"
	config += `    destination_cidr = "any"` + "\n"
	config += `    destination_port = "any"` + "\n"
	config += `    source_cidr = "any"` + "\n"
	config += `    }]` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiApplianceSDWANInternetPoliciesConfig_all() string {
	config := `resource "meraki_appliance_sdwan_internet_policies" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  wan_traffic_uplink_preferences = [{` + "\n"
	config += `    fail_over_criterion = "poorPerformance"` + "\n"
	config += `    preferred_uplink = "wan1"` + "\n"
	config += `    builtin_performance_class_name = "VoIP"` + "\n"
	config += `    custom_performance_class_id = "123456"` + "\n"
	config += `    performance_class_type = "custom"` + "\n"
	config += `    traffic_filters = [{` + "\n"
	config += `      type = "custom"` + "\n"
	config += `      protocol = "tcp"` + "\n"
	config += `      destination_cidr = "any"` + "\n"
	config += `      destination_port = "any"` + "\n"
	config += `      destination_applications = [{` + "\n"
	config += `        id = "meraki:layer7/application/3"` + "\n"
	config += `        name = "DNS"` + "\n"
	config += `        type = "major"` + "\n"
	config += `      }]` + "\n"
	config += `      source_cidr = "any"` + "\n"
	config += `      source_host = 254` + "\n"
	config += `      source_port = "1-1024"` + "\n"
	config += `      source_vlan = 10` + "\n"
	config += `    }]` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
