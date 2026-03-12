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

func TestAccMerakiOrganizationsWirelessDevicesProvisioningDeployments(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "meta_counts_items_remaining", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "meta_counts_items_total", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.completed_at", "2018-02-11T00:00:00.090210Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.created_at", "2018-02-11T00:00:00.090210Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.deployment_id", "1284392014819"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.last_updated_at", "2018-02-11T00:00:00.090210Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.requested_at", "2018-02-11T00:00:00.090210Z"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.status", "ready"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.type", "replace"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_new_mac", "00:11:22:33:44:55"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_new_model", "CW9166I"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_new_name", "My AP"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_new_serial", "Q234-ABCD-5678"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_new_rf_profile_id", "1284392014819"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_new_rf_profile_name", "RF Profile Name"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_new_tags.0", "tag1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_old_after_action", "unclaim"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_old_mac", "00:11:22:33:44:55"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_old_model", "MR34"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_old_name", "My AP"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_old_serial", "Q234-ABCD-5678"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_old_rf_profile_id", "1284392014819"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_old_rf_profile_name", "RF Profile Name"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.devices_old_tags.0", "tag1"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.network_id", "N_24329156"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.network_name", "Main Office"))
	checks = append(checks, resource.TestCheckResourceAttr("meraki_organizations_wireless_devices_provisioning_deployments.test", "items.0.errors.0", "error message1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccMerakiOrganizationsWirelessDevicesProvisioningDeploymentsPrerequisitesConfig + testAccMerakiOrganizationsWirelessDevicesProvisioningDeploymentsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccMerakiOrganizationsWirelessDevicesProvisioningDeploymentsPrerequisitesConfig + testAccMerakiOrganizationsWirelessDevicesProvisioningDeploymentsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:            "meraki_organizations_wireless_devices_provisioning_deployments.test",
		ImportState:             true,
		ImportStateVerify:       true,
		ImportStateIdFunc:       merakiOrganizationsWirelessDevicesProvisioningDeploymentsImportStateIdFunc("meraki_organizations_wireless_devices_provisioning_deployments.test"),
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

func merakiOrganizationsWirelessDevicesProvisioningDeploymentsImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		primary := s.RootModule().Resources[resourceName].Primary
		OrganizationId := primary.Attributes["organization_id"]

		return fmt.Sprintf("%s", OrganizationId), nil
	}
}

// End of section. //template:end importStateIdFunc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccMerakiOrganizationsWirelessDevicesProvisioningDeploymentsPrerequisitesConfig = `
variable "test_org" {}
data "meraki_organization" "test" {
  name = var.test_org
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccMerakiOrganizationsWirelessDevicesProvisioningDeploymentsConfig_minimum() string {
	config := `resource "meraki_organizations_wireless_devices_provisioning_deployments" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  items = [{` + "\n"
	config += `    status = "ready"` + "\n"
	config += `    type = "replace"` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccMerakiOrganizationsWirelessDevicesProvisioningDeploymentsConfig_all() string {
	config := `resource "meraki_organizations_wireless_devices_provisioning_deployments" "test" {` + "\n"
	config += `  organization_id = data.meraki_organization.test.id` + "\n"
	config += `  meta_counts_items_remaining = 0` + "\n"
	config += `  meta_counts_items_total = 20` + "\n"
	config += `  items = [{` + "\n"
	config += `    completed_at = "2018-02-11T00:00:00.090210Z"` + "\n"
	config += `    created_at = "2018-02-11T00:00:00.090210Z"` + "\n"
	config += `    deployment_id = "1284392014819"` + "\n"
	config += `    last_updated_at = "2018-02-11T00:00:00.090210Z"` + "\n"
	config += `    requested_at = "2018-02-11T00:00:00.090210Z"` + "\n"
	config += `    status = "ready"` + "\n"
	config += `    type = "replace"` + "\n"
	config += `    devices_new_mac = "00:11:22:33:44:55"` + "\n"
	config += `    devices_new_model = "CW9166I"` + "\n"
	config += `    devices_new_name = "My AP"` + "\n"
	config += `    devices_new_serial = "Q234-ABCD-5678"` + "\n"
	config += `    devices_new_rf_profile_id = "1284392014819"` + "\n"
	config += `    devices_new_rf_profile_name = "RF Profile Name"` + "\n"
	config += `    devices_new_tags = ["tag1"]` + "\n"
	config += `    devices_old_after_action = "unclaim"` + "\n"
	config += `    devices_old_mac = "00:11:22:33:44:55"` + "\n"
	config += `    devices_old_model = "MR34"` + "\n"
	config += `    devices_old_name = "My AP"` + "\n"
	config += `    devices_old_serial = "Q234-ABCD-5678"` + "\n"
	config += `    devices_old_rf_profile_id = "1284392014819"` + "\n"
	config += `    devices_old_rf_profile_name = "RF Profile Name"` + "\n"
	config += `    devices_old_tags = ["tag1"]` + "\n"
	config += `    network_id = "N_24329156"` + "\n"
	config += `    network_name = "Main Office"` + "\n"
	config += `    errors = ["error message1"]` + "\n"
	config += `  }]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAdditional

// End of section. //template:end testAccConfigAdditional
