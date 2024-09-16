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
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchDHCPServerPolicy struct {
	Id                   types.String `tfsdk:"id"`
	NetworkId            types.String `tfsdk:"network_id"`
	DefaultPolicy        types.String `tfsdk:"default_policy"`
	AlertsEmailEnabled   types.Bool   `tfsdk:"alerts_email_enabled"`
	ArpInspectionEnabled types.Bool   `tfsdk:"arp_inspection_enabled"`
	AllowedServers       types.List   `tfsdk:"allowed_servers"`
	BlockedServers       types.List   `tfsdk:"blocked_servers"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchDHCPServerPolicy) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/dhcpServerPolicy", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchDHCPServerPolicy) toBody(ctx context.Context, state SwitchDHCPServerPolicy) string {
	body := ""
	if !data.DefaultPolicy.IsNull() {
		body, _ = sjson.Set(body, "defaultPolicy", data.DefaultPolicy.ValueString())
	}
	if !data.AlertsEmailEnabled.IsNull() {
		body, _ = sjson.Set(body, "alerts.email.enabled", data.AlertsEmailEnabled.ValueBool())
	}
	if !data.ArpInspectionEnabled.IsNull() {
		body, _ = sjson.Set(body, "arpInspection.enabled", data.ArpInspectionEnabled.ValueBool())
	}
	if !data.AllowedServers.IsNull() {
		var values []string
		data.AllowedServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "allowedServers", values)
	}
	if !data.BlockedServers.IsNull() {
		var values []string
		data.BlockedServers.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "blockedServers", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchDHCPServerPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("defaultPolicy"); value.Exists() {
		data.DefaultPolicy = types.StringValue(value.String())
	} else {
		data.DefaultPolicy = types.StringNull()
	}
	if value := res.Get("alerts.email.enabled"); value.Exists() {
		data.AlertsEmailEnabled = types.BoolValue(value.Bool())
	} else {
		data.AlertsEmailEnabled = types.BoolNull()
	}
	if value := res.Get("arpInspection.enabled"); value.Exists() {
		data.ArpInspectionEnabled = types.BoolValue(value.Bool())
	} else {
		data.ArpInspectionEnabled = types.BoolNull()
	}
	if value := res.Get("allowedServers"); value.Exists() {
		data.AllowedServers = helpers.GetStringList(value.Array())
	} else {
		data.AllowedServers = types.ListNull(types.StringType)
	}
	if value := res.Get("blockedServers"); value.Exists() {
		data.BlockedServers = helpers.GetStringList(value.Array())
	} else {
		data.BlockedServers = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchDHCPServerPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("defaultPolicy"); value.Exists() && !data.DefaultPolicy.IsNull() {
		data.DefaultPolicy = types.StringValue(value.String())
	} else {
		data.DefaultPolicy = types.StringNull()
	}
	if value := res.Get("alerts.email.enabled"); value.Exists() && !data.AlertsEmailEnabled.IsNull() {
		data.AlertsEmailEnabled = types.BoolValue(value.Bool())
	} else {
		data.AlertsEmailEnabled = types.BoolNull()
	}
	if value := res.Get("arpInspection.enabled"); value.Exists() && !data.ArpInspectionEnabled.IsNull() {
		data.ArpInspectionEnabled = types.BoolValue(value.Bool())
	} else {
		data.ArpInspectionEnabled = types.BoolNull()
	}
	if value := res.Get("allowedServers"); value.Exists() && !data.AllowedServers.IsNull() {
		data.AllowedServers = helpers.GetStringList(value.Array())
	} else {
		data.AllowedServers = types.ListNull(types.StringType)
	}
	if value := res.Get("blockedServers"); value.Exists() && !data.BlockedServers.IsNull() {
		data.BlockedServers = helpers.GetStringList(value.Array())
	} else {
		data.BlockedServers = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial