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

func TestAccDataSourceMerakiOrganization(t *testing.T) {
	if os.Getenv("ORGANIZATION") == "" {
		t.Skip("skipping test, set environment variable ORGANIZATION")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization.test", "name", "My organization"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization.test", "management_details.0.name", "MSP ID"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization.test", "management_details.0.value", "123456"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiOrganizationConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceMerakiOrganizationConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiOrganizationConfig() string {
	config := `resource "meraki_organization" "test" {` + "\n"
	config += `	name = "My organization"` + "\n"
	config += `	management_details = [{` + "\n"
	config += `		name = "MSP ID"` + "\n"
	config += `		value = "123456"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization" "test" {
			id = meraki_organization.test.id
		}
	`
	return config
}

func testAccNamedDataSourceMerakiOrganizationConfig() string {
	config := `resource "meraki_organization" "test" {` + "\n"
	config += `	name = "My organization"` + "\n"
	config += `	management_details = [{` + "\n"
	config += `		name = "MSP ID"` + "\n"
	config += `		value = "123456"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization" "test" {
			name = meraki_organization.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
