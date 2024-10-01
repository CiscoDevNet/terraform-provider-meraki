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
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchAccessPolicies struct {
	NetworkId types.String                `tfsdk:"network_id"`
	Items     []SwitchAccessPoliciesItems `tfsdk:"items"`
}

type SwitchAccessPoliciesItems struct {
	Id                                  types.String                                  `tfsdk:"id"`
	AccessPolicyType                    types.String                                  `tfsdk:"access_policy_type"`
	GuestPortBouncing                   types.Bool                                    `tfsdk:"guest_port_bouncing"`
	GuestVlanId                         types.Int64                                   `tfsdk:"guest_vlan_id"`
	HostMode                            types.String                                  `tfsdk:"host_mode"`
	IncreaseAccessSpeed                 types.Bool                                    `tfsdk:"increase_access_speed"`
	Name                                types.String                                  `tfsdk:"name"`
	RadiusAccountingEnabled             types.Bool                                    `tfsdk:"radius_accounting_enabled"`
	RadiusCoaSupportEnabled             types.Bool                                    `tfsdk:"radius_coa_support_enabled"`
	RadiusGroupAttribute                types.String                                  `tfsdk:"radius_group_attribute"`
	RadiusTestingEnabled                types.Bool                                    `tfsdk:"radius_testing_enabled"`
	UrlRedirectWalledGardenEnabled      types.Bool                                    `tfsdk:"url_redirect_walled_garden_enabled"`
	VoiceVlanClients                    types.Bool                                    `tfsdk:"voice_vlan_clients"`
	Dot1xControlDirection               types.String                                  `tfsdk:"dot1x_control_direction"`
	RadiusFailedAuthVlanId              types.Int64                                   `tfsdk:"radius_failed_auth_vlan_id"`
	RadiusReAuthenticationInterval      types.Int64                                   `tfsdk:"radius_re_authentication_interval"`
	RadiusCacheEnabled                  types.Bool                                    `tfsdk:"radius_cache_enabled"`
	RadiusCacheTimeout                  types.Int64                                   `tfsdk:"radius_cache_timeout"`
	RadiusCriticalAuthDataVlanId        types.Int64                                   `tfsdk:"radius_critical_auth_data_vlan_id"`
	RadiusCriticalAuthSuspendPortBounce types.Bool                                    `tfsdk:"radius_critical_auth_suspend_port_bounce"`
	RadiusCriticalAuthVoiceVlanId       types.Int64                                   `tfsdk:"radius_critical_auth_voice_vlan_id"`
	RadiusAccountingServers             []SwitchAccessPoliciesRadiusAccountingServers `tfsdk:"radius_accounting_servers"`
	RadiusServers                       []SwitchAccessPoliciesRadiusServers           `tfsdk:"radius_servers"`
	UrlRedirectWalledGardenRanges       types.Set                                     `tfsdk:"url_redirect_walled_garden_ranges"`
}

type SwitchAccessPoliciesRadiusAccountingServers struct {
	Host                       types.String `tfsdk:"host"`
	OrganizationRadiusServerId types.String `tfsdk:"organization_radius_server_id"`
	Port                       types.Int64  `tfsdk:"port"`
	Secret                     types.String `tfsdk:"secret"`
}

type SwitchAccessPoliciesRadiusServers struct {
	Host                       types.String `tfsdk:"host"`
	OrganizationRadiusServerId types.String `tfsdk:"organization_radius_server_id"`
	Port                       types.Int64  `tfsdk:"port"`
	Secret                     types.String `tfsdk:"secret"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchAccessPolicies) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/accessPolicies", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchAccessPolicies) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]SwitchAccessPoliciesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := SwitchAccessPoliciesItems{}
		data.Id = types.StringValue(res.Get("accessPolicyNumber").String())
		if value := res.Get("accessPolicyType"); value.Exists() && value.Value() != nil {
			data.AccessPolicyType = types.StringValue(value.String())
		} else {
			data.AccessPolicyType = types.StringNull()
		}
		if value := res.Get("guestPortBouncing"); value.Exists() && value.Value() != nil {
			data.GuestPortBouncing = types.BoolValue(value.Bool())
		} else {
			data.GuestPortBouncing = types.BoolNull()
		}
		if value := res.Get("guestVlanId"); value.Exists() && value.Value() != nil {
			data.GuestVlanId = types.Int64Value(value.Int())
		} else {
			data.GuestVlanId = types.Int64Null()
		}
		if value := res.Get("hostMode"); value.Exists() && value.Value() != nil {
			data.HostMode = types.StringValue(value.String())
		} else {
			data.HostMode = types.StringNull()
		}
		if value := res.Get("increaseAccessSpeed"); value.Exists() && value.Value() != nil {
			data.IncreaseAccessSpeed = types.BoolValue(value.Bool())
		} else {
			data.IncreaseAccessSpeed = types.BoolNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("radiusAccountingEnabled"); value.Exists() && value.Value() != nil {
			data.RadiusAccountingEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusAccountingEnabled = types.BoolNull()
		}
		if value := res.Get("radiusCoaSupportEnabled"); value.Exists() && value.Value() != nil {
			data.RadiusCoaSupportEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusCoaSupportEnabled = types.BoolNull()
		}
		if value := res.Get("radiusGroupAttribute"); value.Exists() && value.Value() != nil {
			data.RadiusGroupAttribute = types.StringValue(value.String())
		} else {
			data.RadiusGroupAttribute = types.StringNull()
		}
		if value := res.Get("radiusTestingEnabled"); value.Exists() && value.Value() != nil {
			data.RadiusTestingEnabled = types.BoolValue(value.Bool())
		} else {
			data.RadiusTestingEnabled = types.BoolNull()
		}
		if value := res.Get("urlRedirectWalledGardenEnabled"); value.Exists() && value.Value() != nil {
			data.UrlRedirectWalledGardenEnabled = types.BoolValue(value.Bool())
		} else {
			data.UrlRedirectWalledGardenEnabled = types.BoolNull()
		}
		if value := res.Get("voiceVlanClients"); value.Exists() && value.Value() != nil {
			data.VoiceVlanClients = types.BoolValue(value.Bool())
		} else {
			data.VoiceVlanClients = types.BoolNull()
		}
		if value := res.Get("dot1x.controlDirection"); value.Exists() && value.Value() != nil {
			data.Dot1xControlDirection = types.StringValue(value.String())
		} else {
			data.Dot1xControlDirection = types.StringNull()
		}
		if value := res.Get("radius.failedAuthVlanId"); value.Exists() && value.Value() != nil {
			data.RadiusFailedAuthVlanId = types.Int64Value(value.Int())
		} else {
			data.RadiusFailedAuthVlanId = types.Int64Null()
		}
		if value := res.Get("radius.reAuthenticationInterval"); value.Exists() && value.Value() != nil {
			data.RadiusReAuthenticationInterval = types.Int64Value(value.Int())
		} else {
			data.RadiusReAuthenticationInterval = types.Int64Null()
		}
		if value := res.Get("radius.criticalAuth.dataVlanId"); value.Exists() && value.Value() != nil {
			data.RadiusCriticalAuthDataVlanId = types.Int64Value(value.Int())
		} else {
			data.RadiusCriticalAuthDataVlanId = types.Int64Null()
		}
		if value := res.Get("radius.criticalAuth.suspendPortBounce"); value.Exists() && value.Value() != nil {
			data.RadiusCriticalAuthSuspendPortBounce = types.BoolValue(value.Bool())
		} else {
			data.RadiusCriticalAuthSuspendPortBounce = types.BoolNull()
		}
		if value := res.Get("radius.criticalAuth.voiceVlanId"); value.Exists() && value.Value() != nil {
			data.RadiusCriticalAuthVoiceVlanId = types.Int64Value(value.Int())
		} else {
			data.RadiusCriticalAuthVoiceVlanId = types.Int64Null()
		}
		if value := res.Get("radiusAccountingServers"); value.Exists() && value.Value() != nil {
			data.RadiusAccountingServers = make([]SwitchAccessPoliciesRadiusAccountingServers, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := SwitchAccessPoliciesRadiusAccountingServers{}
				if value := res.Get("host"); value.Exists() && value.Value() != nil {
					data.Host = types.StringValue(value.String())
				} else {
					data.Host = types.StringNull()
				}
				if value := res.Get("organizationRadiusServerId"); value.Exists() && value.Value() != nil {
					data.OrganizationRadiusServerId = types.StringValue(value.String())
				} else {
					data.OrganizationRadiusServerId = types.StringNull()
				}
				if value := res.Get("port"); value.Exists() && value.Value() != nil {
					data.Port = types.Int64Value(value.Int())
				} else {
					data.Port = types.Int64Null()
				}
				(*parent).RadiusAccountingServers = append((*parent).RadiusAccountingServers, data)
				return true
			})
		}
		if value := res.Get("radiusServers"); value.Exists() && value.Value() != nil {
			data.RadiusServers = make([]SwitchAccessPoliciesRadiusServers, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := SwitchAccessPoliciesRadiusServers{}
				if value := res.Get("host"); value.Exists() && value.Value() != nil {
					data.Host = types.StringValue(value.String())
				} else {
					data.Host = types.StringNull()
				}
				if value := res.Get("organizationRadiusServerId"); value.Exists() && value.Value() != nil {
					data.OrganizationRadiusServerId = types.StringValue(value.String())
				} else {
					data.OrganizationRadiusServerId = types.StringNull()
				}
				if value := res.Get("port"); value.Exists() && value.Value() != nil {
					data.Port = types.Int64Value(value.Int())
				} else {
					data.Port = types.Int64Null()
				}
				(*parent).RadiusServers = append((*parent).RadiusServers, data)
				return true
			})
		}
		if value := res.Get("urlRedirectWalledGardenRanges"); value.Exists() && value.Value() != nil {
			data.UrlRedirectWalledGardenRanges = helpers.GetStringSet(value.Array())
		} else {
			data.UrlRedirectWalledGardenRanges = types.SetNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody