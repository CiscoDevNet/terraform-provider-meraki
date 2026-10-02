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
	"testing"

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
