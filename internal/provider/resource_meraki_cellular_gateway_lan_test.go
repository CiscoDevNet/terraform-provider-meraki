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

func TestAccMerakiCellularGatewayLAN(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" || os.Getenv("TF_VAR_test_cellular_1_serial") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network and TF_VAR_test_cellular_1_serial")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_cellular_gateway_lan.test", "fixed_ip_assignments.0.ip", "172.31.128.10"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_cellular_gateway_lan.test", "fixed_ip_assignments.0.mac", "0b:00:00:00:00:ac"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_cellular_gateway_lan.test", "fixed_ip_assignments.0.name", "server 1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_cellular_gateway_lan.test", "reserved_ip_ranges.0.comment", "A reserved IP range"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_cellular_gateway_lan.test", "reserved_ip_ranges.0.end", "172.31.128.1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_cellular_gateway_lan.test", "reserved_ip_ranges.0.start", "172.31.128.0"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiCellularGatewayLANPrerequisitesConfig + testAccMerakiCellularGatewayLANConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiCellularGatewayLANPrerequisitesConfig + testAccMerakiCellularGatewayLANConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_cellular_gateway_lan.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiCellularGatewayLANImportStateIdFunc("meraki_cellular_gateway_lan.test"),
		ImportStateVerifyIgnore: []string{},
		Check:                   resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin importStateIdFunc

func merakiCellularGatewayLANImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		Serial := primary.Attributes["serial"]

		return fmt.Sprintf("%s", Serial), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiCellularGatewayLANPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
variable "test_cellular_1_serial" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["cellularGateway"]
}
resource "meraki_network_device_claim" "test" {
  network_id = meraki_network.test.id
  serials    = [var.test_cellular_1_serial]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiCellularGatewayLANConfig_minimum() string {
	config := `resource "meraki_cellular_gateway_lan" "test" {` + "\n"
	config += `  serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  fixed_ip_assignments = [{` + "\n"
	config += `    ip = "172.31.128.10"` + "\n"
	config += `    mac = "0b:00:00:00:00:ac"` + "\n"
	config += `  }]` + "\n"
	config += `  reserved_ip_ranges = [{` + "\n"
	config += `    comment = "A reserved IP range"` + "\n"
	config += `    end = "172.31.128.1"` + "\n"
	config += `    start = "172.31.128.0"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiCellularGatewayLANConfig_all() string {
	config := `resource "meraki_cellular_gateway_lan" "test" {` + "\n"
	config += `  serial = tolist(meraki_network_device_claim.test.serials)[0]` + "\n"
	config += `  fixed_ip_assignments = [{` + "\n"
	config += `    ip = "172.31.128.10"` + "\n"
	config += `    mac = "0b:00:00:00:00:ac"` + "\n"
	config += `    name = "server 1"` + "\n"
	config += `  }]` + "\n"
	config += `  reserved_ip_ranges = [{` + "\n"
	config += `    comment = "A reserved IP range"` + "\n"
	config += `    end = "172.31.128.1"` + "\n"
	config += `    start = "172.31.128.0"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
