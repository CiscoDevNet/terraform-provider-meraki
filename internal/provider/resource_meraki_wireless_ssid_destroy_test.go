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

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/tidwall/gjson"
)

// Regression test for https://github.com/netascode/terraform-meraki-nac-meraki/issues/204:
// destroying an SSID configured for named VLAN tagging failed because the destroy-time PUT
// reset ipAssignmentMode to "NAT mode" without also clearing the fields that are only valid
// in Bridge mode / Layer 3 roaming (useVlanTagging, namedVlans.tagging.enabled,
// namedVlans.radius.guestVlan.enabled), which the Meraki API then rejected.
func TestWirelessSSIDToDestroyBody(t *testing.T) {
	body := WirelessSSID{}.toDestroyBody(context.Background())
	assertDestroyBodyClearsNamedVlanDependencies(t, body)
}

func TestResourceWirelessSSIDsToDestroyBody(t *testing.T) {
	body := ResourceWirelessSSIDs{}.toDestroyBody(context.Background())
	assertDestroyBodyClearsNamedVlanDependencies(t, body)
}

func assertDestroyBodyClearsNamedVlanDependencies(t *testing.T, body string) {
	t.Helper()

	if got := gjson.Get(body, "ipAssignmentMode").String(); got != "NAT mode" {
		t.Errorf("ipAssignmentMode = %q, want %q", got, "NAT mode")
	}
	for _, path := range []string{
		"useVlanTagging",
		"namedVlans.tagging.enabled",
		"namedVlans.radius.guestVlan.enabled",
	} {
		result := gjson.Get(body, path)
		if !result.Exists() {
			t.Errorf("destroy body is missing %q; the Meraki API rejects ipAssignmentMode = \"NAT mode\" while this remains enabled", path)
			continue
		}
		if result.Bool() {
			t.Errorf("%s = true, want false alongside ipAssignmentMode = \"NAT mode\"", path)
		}
	}
}

// This acceptance test lives in a hand-maintained file rather than the generated
// resource_meraki_wireless_ssid_test.go, because that file's testAccConfigAdditional section is
// rewritten from the additional_tests YAML field on every "make gen" and cannot hold a custom
// TestCase/Steps function like this one. It reproduces the exact scenario from the issue
// (use_vlan_tagging left false, only named_vlans.tagging enabled) rather than combining both
// VLAN-tagging mechanisms, since the Meraki API rejects that combination outright. The
// use_vlan_tagging and named_vlans.radius.guestVlan.enabled twin bugs are covered by
// TestWirelessSSIDToDestroyBody instead.
func TestAccMerakiWirelessSSID_namedVlanTaggingDestroy(t *testing.T) {
	if os.Getenv("TF_VAR_test_org") == "" || os.Getenv("TF_VAR_test_network") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_test_org and TF_VAR_test_network")
	}

	// Note: the VLAN name "default" is stripped by the provider's network_vlan_profile
	// ignore_import_values handling on every read, not just import, so it must not be used
	// here or the first config step below would never converge to a stable plan.
	config := testAccMerakiWirelessSSIDPrerequisitesConfig + `
resource "meraki_network_vlan_profile" "named_vlan_test" {
  network_id = meraki_network.test.id
  iname      = "Default"
  name       = "Default Profile"
  vlan_names = [
    {
      name    = "native"
      vlan_id = "1"
    },
    {
      name    = "guest_66"
      vlan_id = "66"
    }
  ]
  vlan_groups = [
    {
      name     = "named-group-1"
      vlan_ids = "2,5-7"
    }
  ]
}

resource "meraki_wireless_ssid" "named_vlan_test" {
  network_id = meraki_network.test.id
  number = "1"
  name = "Guest-Open"
  enabled = true
  auth_mode = "open"
  splash_page = "None"
  ip_assignment_mode = "Bridge mode"
  use_vlan_tagging = false
  named_vlans_tagging_enabled = true
  named_vlans_tagging_default_vlan_name = "guest_66"

  depends_on = [meraki_network_vlan_profile.named_vlan_test]
}
`

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Create the SSID with named VLAN tagging enabled.
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("meraki_wireless_ssid.named_vlan_test", "named_vlans_tagging_enabled", "true"),
					resource.TestCheckResourceAttr("meraki_wireless_ssid.named_vlan_test", "use_vlan_tagging", "false"),
				),
			},
			{
				// Remove it from config, mirroring NaC decommissioning the SSID. Before the
				// destroy_value fix this step failed with:
				//   HTTP Request failed: StatusCode 400, JSON error: ["Named VLAN tagging
				//   requires Bridge Mode or Layer 3 Roaming IP assignment mode."]
				Config: testAccMerakiWirelessSSIDPrerequisitesConfig,
			},
		},
	})
}
