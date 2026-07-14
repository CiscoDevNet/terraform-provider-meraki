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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSourceSwitchDHCPServerPolicy - used only by the data source, always including every
// attribute (unlike model_resource.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model_resource.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model_resource.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceSwitchDHCPServerPolicy struct {
	Id                   types.String `tfsdk:"id"`
	NetworkId            types.String `tfsdk:"network_id"`
	DefaultPolicy        types.String `tfsdk:"default_policy"`
	AlertsEmailEnabled   types.Bool   `tfsdk:"alerts_email_enabled"`
	ArpInspectionEnabled types.Bool   `tfsdk:"arp_inspection_enabled"`
	AllowedServers       types.Set    `tfsdk:"allowed_servers"`
	BlockedServers       types.Set    `tfsdk:"blocked_servers"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceSwitchDHCPServerPolicy) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/dhcpServerPolicy", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceSwitchDHCPServerPolicy) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("defaultPolicy"); value.Exists() && value.Value() != nil {
		data.DefaultPolicy = types.StringValue(value.String())
	} else {
		data.DefaultPolicy = types.StringNull()
	}
	if value := res.Get("alerts.email.enabled"); value.Exists() && value.Value() != nil {
		data.AlertsEmailEnabled = types.BoolValue(value.Bool())
	} else {
		data.AlertsEmailEnabled = types.BoolNull()
	}
	if value := res.Get("arpInspection.enabled"); value.Exists() && value.Value() != nil {
		data.ArpInspectionEnabled = types.BoolValue(value.Bool())
	} else {
		data.ArpInspectionEnabled = types.BoolNull()
	}
	if value := res.Get("allowedServers"); value.Exists() && value.Value() != nil {
		data.AllowedServers = helpers.GetStringSet(value.Array())
	} else {
		data.AllowedServers = types.SetNull(types.StringType)
	}
	if value := res.Get("blockedServers"); value.Exists() && value.Value() != nil {
		data.BlockedServers = helpers.GetStringSet(value.Array())
	} else {
		data.BlockedServers = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody
