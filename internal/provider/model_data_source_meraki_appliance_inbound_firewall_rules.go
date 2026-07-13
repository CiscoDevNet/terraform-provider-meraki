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
// It emits a separate model - DataSourceApplianceInboundFirewallRules - used only by the data source, always including every
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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceApplianceInboundFirewallRules struct {
	Id                types.String                                   `tfsdk:"id"`
	NetworkId         types.String                                   `tfsdk:"network_id"`
	SyslogDefaultRule types.Bool                                     `tfsdk:"syslog_default_rule"`
	Rules             []DataSourceApplianceInboundFirewallRulesRules `tfsdk:"rules"`
}

type DataSourceApplianceInboundFirewallRulesRules struct {
	Comment       types.String `tfsdk:"comment"`
	DestCidr      types.String `tfsdk:"dest_cidr"`
	DestPort      types.String `tfsdk:"dest_port"`
	Policy        types.String `tfsdk:"policy"`
	Protocol      types.String `tfsdk:"protocol"`
	SrcCidr       types.String `tfsdk:"src_cidr"`
	SrcPort       types.String `tfsdk:"src_port"`
	SyslogEnabled types.Bool   `tfsdk:"syslog_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceInboundFirewallRules) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/firewall/inboundFirewallRules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceInboundFirewallRules) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("rules"); value.Exists() && value.Value() != nil {
		data.Rules = make([]DataSourceApplianceInboundFirewallRulesRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			if value := res.Get("comment"); value.String() == "Default rule" {
				return true
			}
			parent := &data
			data := DataSourceApplianceInboundFirewallRulesRules{}
			if value := res.Get("comment"); value.Exists() && value.Value() != nil {
				data.Comment = types.StringValue(value.String())
			} else {
				data.Comment = types.StringNull()
			}
			if value := res.Get("destCidr"); value.Exists() && value.Value() != nil {
				data.DestCidr = types.StringValue(value.String())
			} else {
				data.DestCidr = types.StringNull()
			}
			if value := res.Get("destPort"); value.Exists() && value.Value() != nil {
				data.DestPort = types.StringValue(value.String())
			} else {
				data.DestPort = types.StringNull()
			}
			if value := res.Get("policy"); value.Exists() && value.Value() != nil {
				data.Policy = types.StringValue(value.String())
			} else {
				data.Policy = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("srcCidr"); value.Exists() && value.Value() != nil {
				data.SrcCidr = types.StringValue(value.String())
			} else {
				data.SrcCidr = types.StringNull()
			}
			if value := res.Get("srcPort"); value.Exists() && value.Value() != nil {
				data.SrcPort = types.StringValue(value.String())
			} else {
				data.SrcPort = types.StringNull()
			}
			if value := res.Get("syslogEnabled"); value.Exists() && value.Value() != nil {
				data.SyslogEnabled = types.BoolValue(value.Bool())
			} else {
				data.SyslogEnabled = types.BoolNull()
			}
			(*parent).Rules = append((*parent).Rules, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
