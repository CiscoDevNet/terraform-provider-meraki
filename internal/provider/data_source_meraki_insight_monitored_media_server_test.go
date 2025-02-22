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

func TestAccDataSourceMerakiInsightMonitoredMediaServer(t *testing.T) {
	if os.Getenv("INSIGHT") == "" {
		t.Skip("skipping test, set environment variable INSIGHT")
	}
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_insight_monitored_media_server.test", "address", "123.123.123.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_insight_monitored_media_server.test", "best_effort_monitoring_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_insight_monitored_media_server.test", "name", "Sample VoIP Provider"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiInsightMonitoredMediaServerPrerequisitesConfig + testAccDataSourceMerakiInsightMonitoredMediaServerConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiInsightMonitoredMediaServerPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiInsightMonitoredMediaServerConfig() string {
	config := `resource "meraki_insight_monitored_media_server" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  address = "123.123.123.1"` + "\n"
	config += `  best_effort_monitoring_enabled = true` + "\n"
	config += `  name = "Sample VoIP Provider"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_insight_monitored_media_server" "test" {
			id = meraki_insight_monitored_media_server.test.id
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_insight_monitored_media_server.test]
		}
	`
	return config
}

func testAccNamedDataSourceMerakiInsightMonitoredMediaServerConfig() string {
	config := `resource "meraki_insight_monitored_media_server" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  address = "123.123.123.1"` + "\n"
	config += `  best_effort_monitoring_enabled = true` + "\n"
	config += `  name = "Sample VoIP Provider"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_insight_monitored_media_server" "test" {
			name = meraki_insight_monitored_media_server.test.name
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_insight_monitored_media_server.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
