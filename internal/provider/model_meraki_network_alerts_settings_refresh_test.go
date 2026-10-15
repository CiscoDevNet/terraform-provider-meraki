// Copyright © 2026 Cisco Systems, Inc. and its affiliates.
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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

func TestNetworkAlertsSettingsFromBodyPartialMatchesAlertsByType(t *testing.T) {
	ctx := context.Background()
	state := NetworkAlertsSettings{
		Alerts: []NetworkAlertsSettingsAlerts{
			alertState("applianceDown", true, 10, 100, 1000),
			alertState("usageAlert", false, 20, 200, 2000),
			alertState("snr", true, 30, 300, 3000),
		},
	}
	res := meraki.Res{Result: gjson.Parse(`{
		"alerts": [
			{"type":"snr","enabled":false,"alertDestinations":{"allAdmins":false},"filters":{"timeout":31,"period":301,"threshold":3001}},
			{"type":"nodeHardwareFailure","enabled":true,"filters":{"timeout":99}},
			{"type":"usageAlert","enabled":true,"alertDestinations":{"allAdmins":true},"filters":{"timeout":21,"period":201,"threshold":2001}},
			{"type":"applianceDown","enabled":false,"alertDestinations":{"allAdmins":true},"filters":{"timeout":11,"period":101,"threshold":1001}}
		]
	}`)}

	state.fromBodyPartial(ctx, res)

	assertAlerts(t, state.Alerts, []alertExpectation{
		{"applianceDown", false, true, 11, 101, 1001},
		{"usageAlert", true, true, 21, 201, 2001},
		{"snr", false, false, 31, 301, 3001},
	})
}

func TestNetworkAlertsSettingsFromBodyPartialRetainsAbsentAlerts(t *testing.T) {
	ctx := context.Background()
	state := NetworkAlertsSettings{
		Alerts: []NetworkAlertsSettingsAlerts{
			alertState("applianceDown", true, 10, 100, 1000),
			alertState("usageAlert", false, 20, 200, 2000),
			alertState("snr", true, 30, 300, 3000),
		},
	}
	res := meraki.Res{Result: gjson.Parse(`{
		"alerts": [
			{"type":"snr","enabled":false,"alertDestinations":{"allAdmins":false},"filters":{"timeout":31,"period":301,"threshold":3001}},
			{"type":"applianceDown","enabled":false,"alertDestinations":{"allAdmins":true},"filters":{"timeout":11,"period":101,"threshold":1001}}
		]
	}`)}

	state.fromBodyPartial(ctx, res)

	assertAlerts(t, state.Alerts, []alertExpectation{
		{"applianceDown", false, true, 11, 101, 1001},
		{"usageAlert", false, false, 20, 200, 2000},
		{"snr", false, false, 31, 301, 3001},
	})
}

type alertExpectation struct {
	typeName  string
	enabled   bool
	allAdmins bool
	timeout   int64
	period    int64
	threshold int64
}

func alertState(typeName string, enabled bool, timeout, period, threshold int64) NetworkAlertsSettingsAlerts {
	return NetworkAlertsSettingsAlerts{
		Type:                       types.StringValue(typeName),
		Enabled:                    types.BoolValue(enabled),
		AlertDestinationsAllAdmins: types.BoolValue(false),
		FiltersTimeout:             types.Int64Value(timeout),
		FiltersPeriod:              types.Int64Value(period),
		FiltersThreshold:           types.Int64Value(threshold),
	}
}

func assertAlerts(t *testing.T, actual []NetworkAlertsSettingsAlerts, expected []alertExpectation) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Fatalf("expected %d alerts, got %d", len(expected), len(actual))
	}
	for i, want := range expected {
		got := actual[i]
		if got.Type.ValueString() != want.typeName ||
			got.Enabled.ValueBool() != want.enabled ||
			got.AlertDestinationsAllAdmins.ValueBool() != want.allAdmins ||
			got.FiltersTimeout.ValueInt64() != want.timeout ||
			got.FiltersPeriod.ValueInt64() != want.period ||
			got.FiltersThreshold.ValueInt64() != want.threshold {
			t.Errorf("alerts[%d]: expected %+v, got type=%q enabled=%t all_admins=%t timeout=%d period=%d threshold=%d",
				i, want, got.Type.ValueString(), got.Enabled.ValueBool(), got.AlertDestinationsAllAdmins.ValueBool(),
				got.FiltersTimeout.ValueInt64(), got.FiltersPeriod.ValueInt64(), got.FiltersThreshold.ValueInt64())
		}
	}
}
