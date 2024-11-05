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

func TestAccDataSourceMerakiOrganizationBrandingPolicy(t *testing.T) {
	if os.Getenv("ORGANIZATION_BRANDING_POLICY") == "" {
		t.Skip("skipping test, set environment variable ORGANIZATION_BRANDING_POLICY")
	}
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "name", "My Branding Policy"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "admin_settings_applies_to", "All admins of networks..."))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "admin_settings_values.0", "N_1234"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "custom_logo_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "custom_logo_image_contents", "Hyperg26C8F4h8CvcoUqpA=="))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "custom_logo_image_format", "jpg"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_api_docs_subtab", "default or inherit"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_cases_subtab", "hide"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_cisco_meraki_product_documentation", "show"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_community_subtab", "show"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_data_protection_requests_subtab", "default or inherit"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_firewall_info_subtab", "hide"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_get_help_subtab", "default or inherit"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_get_help_subtab_knowledge_base_search", "<h1>Some custom HTML content</h1>"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_hardware_replacements_subtab", "hide"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_help_tab", "show"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_help_widget", "hide"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_new_features_subtab", "show"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_sm_forums", "hide"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_support_contact_info", "show"))
	checks = append(checks, resource.TestCheckResourceAttr("data.meraki_organization_branding_policy.test", "help_settings_universal_search_knowledge_base_search", "hide"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceMerakiOrganizationBrandingPolicyPrerequisitesConfig + testAccDataSourceMerakiOrganizationBrandingPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceMerakiOrganizationBrandingPolicyPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceMerakiOrganizationBrandingPolicyConfig() string {
	config := `resource "meraki_organization_branding_policy" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	enabled = true` + "\n"
	config += `	name = "My Branding Policy"` + "\n"
	config += `	admin_settings_applies_to = "All admins of networks..."` + "\n"
	config += `	admin_settings_values = ["N_1234"]` + "\n"
	config += `	custom_logo_enabled = true` + "\n"
	config += `	custom_logo_image_contents = "Hyperg26C8F4h8CvcoUqpA=="` + "\n"
	config += `	custom_logo_image_format = "jpg"` + "\n"
	config += `	help_settings_api_docs_subtab = "default or inherit"` + "\n"
	config += `	help_settings_cases_subtab = "hide"` + "\n"
	config += `	help_settings_cisco_meraki_product_documentation = "show"` + "\n"
	config += `	help_settings_community_subtab = "show"` + "\n"
	config += `	help_settings_data_protection_requests_subtab = "default or inherit"` + "\n"
	config += `	help_settings_firewall_info_subtab = "hide"` + "\n"
	config += `	help_settings_get_help_subtab = "default or inherit"` + "\n"
	config += `	help_settings_get_help_subtab_knowledge_base_search = "<h1>Some custom HTML content</h1>"` + "\n"
	config += `	help_settings_hardware_replacements_subtab = "hide"` + "\n"
	config += `	help_settings_help_tab = "show"` + "\n"
	config += `	help_settings_help_widget = "hide"` + "\n"
	config += `	help_settings_new_features_subtab = "show"` + "\n"
	config += `	help_settings_sm_forums = "hide"` + "\n"
	config += `	help_settings_support_contact_info = "show"` + "\n"
	config += `	help_settings_universal_search_knowledge_base_search = "hide"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_branding_policy" "test" {
			id = meraki_organization_branding_policy.test.id
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_organization_branding_policy.test]
		}
	`
	return config
}

func testAccNamedDataSourceMerakiOrganizationBrandingPolicyConfig() string {
	config := `resource "meraki_organization_branding_policy" "test" {` + "\n"
	config += `	organization_id = data.meraki_organization.test.id` + "\n"
	config += `	enabled = true` + "\n"
	config += `	name = "My Branding Policy"` + "\n"
	config += `	admin_settings_applies_to = "All admins of networks..."` + "\n"
	config += `	admin_settings_values = ["N_1234"]` + "\n"
	config += `	custom_logo_enabled = true` + "\n"
	config += `	custom_logo_image_contents = "Hyperg26C8F4h8CvcoUqpA=="` + "\n"
	config += `	custom_logo_image_format = "jpg"` + "\n"
	config += `	help_settings_api_docs_subtab = "default or inherit"` + "\n"
	config += `	help_settings_cases_subtab = "hide"` + "\n"
	config += `	help_settings_cisco_meraki_product_documentation = "show"` + "\n"
	config += `	help_settings_community_subtab = "show"` + "\n"
	config += `	help_settings_data_protection_requests_subtab = "default or inherit"` + "\n"
	config += `	help_settings_firewall_info_subtab = "hide"` + "\n"
	config += `	help_settings_get_help_subtab = "default or inherit"` + "\n"
	config += `	help_settings_get_help_subtab_knowledge_base_search = "<h1>Some custom HTML content</h1>"` + "\n"
	config += `	help_settings_hardware_replacements_subtab = "hide"` + "\n"
	config += `	help_settings_help_tab = "show"` + "\n"
	config += `	help_settings_help_widget = "hide"` + "\n"
	config += `	help_settings_new_features_subtab = "show"` + "\n"
	config += `	help_settings_sm_forums = "hide"` + "\n"
	config += `	help_settings_support_contact_info = "show"` + "\n"
	config += `	help_settings_universal_search_knowledge_base_search = "hide"` + "\n"
	config += `}` + "\n"

	config += `
		data "meraki_organization_branding_policy" "test" {
			name = meraki_organization_branding_policy.test.name
			organization_id = data.meraki_organization.test.id
			depends_on = [meraki_organization_branding_policy.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
