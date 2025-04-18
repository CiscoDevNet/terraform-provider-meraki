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

func TestAccDataSourceMerakiApplianceDNSLocalRecord(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_dns_local_record.test", "address", "10.1.1.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_appliance_dns_local_record.test", "hostname", "www.test.com"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiApplianceDNSLocalRecordPrerequisitesConfig + testAccDataSourceMerakiApplianceDNSLocalRecordConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiApplianceDNSLocalRecordPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}
resource "meraki_appliance_dns_local_profile" "test" {
  organization_id = data.meraki_organization.test.id
  name = "Default Profile"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiApplianceDNSLocalRecordConfig() string {
	config := `resource "meraki_appliance_dns_local_record" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  address = "10.1.1.0"` + "\n"
	config += `  hostname = "www.test.com"` + "\n"
	config += `  profile_id = meraki_appliance_dns_local_profile.test.id` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_appliance_dns_local_record" "test" {
			id = meraki_appliance_dns_local_record.test.id
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_appliance_dns_local_record.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
