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

func TestAccDataSourceMerakiNetworkTrafficAnalysis(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_traffic_analysis.test", "mode", "basic"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_traffic_analysis.test", "custom_pie_chart_items.0.name", "Item from hostname"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_traffic_analysis.test", "custom_pie_chart_items.0.type", "host"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_network_traffic_analysis.test", "custom_pie_chart_items.0.value", "example.com"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiNetworkTrafficAnalysisPrerequisitesConfig + testAccDataSourceMerakiNetworkTrafficAnalysisConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiNetworkTrafficAnalysisPrerequisitesConfig = `
variable "test_org" {}
variable "test_network" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_network" "test" {
  organization_id = data.meraki_organization.test.id
  name            = var.test_network
  product_types   = ["switch", "wireless"]
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiNetworkTrafficAnalysisConfig() string {
	config := `resource "meraki_network_traffic_analysis" "test" {` + "\n"
	config += `  network_id = meraki_network.test.id` + "\n"
	config += `  mode = "basic"` + "\n"
	config += `  custom_pie_chart_items = [{` + "\n"
	config += `    name = "Item from hostname"` + "\n"
	config += `    type = "host"` + "\n"
	config += `    value = "example.com"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_network_traffic_analysis" "test" {
			network_id = meraki_network.test.id
			depends_on = [meraki_network_traffic_analysis.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
