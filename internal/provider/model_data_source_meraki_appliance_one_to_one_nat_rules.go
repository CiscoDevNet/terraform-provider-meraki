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
// It emits a separate model - DataSourceApplianceOneToOneNATRules - used only by the data source, always including every
// attribute (unlike model.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model.go's resource-side type names.

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceApplianceOneToOneNATRules struct {
	Id        types.String                               `tfsdk:"id"`
	NetworkId types.String                               `tfsdk:"network_id"`
	Rules     []DataSourceApplianceOneToOneNATRulesRules `tfsdk:"rules"`
}

type DataSourceApplianceOneToOneNATRulesRules struct {
	LanIp          types.String                                             `tfsdk:"lan_ip"`
	Name           types.String                                             `tfsdk:"name"`
	PublicIp       types.String                                             `tfsdk:"public_ip"`
	Uplink         types.String                                             `tfsdk:"uplink"`
	AllowedInbound []DataSourceApplianceOneToOneNATRulesRulesAllowedInbound `tfsdk:"allowed_inbound"`
}

type DataSourceApplianceOneToOneNATRulesRulesAllowedInbound struct {
	Protocol         types.String `tfsdk:"protocol"`
	AllowedIps       types.List   `tfsdk:"allowed_ips"`
	DestinationPorts types.List   `tfsdk:"destination_ports"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceOneToOneNATRules) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/firewall/oneToOneNatRules", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceOneToOneNATRules) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("rules"); value.Exists() && value.Value() != nil {
		data.Rules = make([]DataSourceApplianceOneToOneNATRulesRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceOneToOneNATRulesRules{}
			if value := res.Get("lanIp"); value.Exists() && value.Value() != nil {
				data.LanIp = types.StringValue(value.String())
			} else {
				data.LanIp = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("publicIp"); value.Exists() && value.Value() != nil {
				data.PublicIp = types.StringValue(value.String())
			} else {
				data.PublicIp = types.StringNull()
			}
			if value := res.Get("uplink"); value.Exists() && value.Value() != nil {
				data.Uplink = types.StringValue(value.String())
			} else {
				data.Uplink = types.StringNull()
			}
			if value := res.Get("allowedInbound"); value.Exists() && value.Value() != nil {
				data.AllowedInbound = make([]DataSourceApplianceOneToOneNATRulesRulesAllowedInbound, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DataSourceApplianceOneToOneNATRulesRulesAllowedInbound{}
					if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
						data.Protocol = types.StringValue(value.String())
					} else {
						data.Protocol = types.StringNull()
					}
					if value := res.Get("allowedIps"); value.Exists() && value.Value() != nil {
						data.AllowedIps = helpers.GetStringList(value.Array())
					} else {
						data.AllowedIps = types.ListNull(types.StringType)
					}
					if value := res.Get("destinationPorts"); value.Exists() && value.Value() != nil {
						data.DestinationPorts = helpers.GetStringList(value.Array())
					} else {
						data.DestinationPorts = types.ListNull(types.StringType)
					}
					(*parent).AllowedInbound = append((*parent).AllowedInbound, data)
					return true
				})
			}
			(*parent).Rules = append((*parent).Rules, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
