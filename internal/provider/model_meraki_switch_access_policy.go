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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchAccessPolicy struct {
	Id                                  types.String                                `tfsdk:"id"`
	NetworkId                           types.String                                `tfsdk:"network_id"`
	AccessPolicyType                    types.String                                `tfsdk:"access_policy_type"`
	GuestPortBouncing                   types.Bool                                  `tfsdk:"guest_port_bouncing"`
	GuestVlanId                         types.Int64                                 `tfsdk:"guest_vlan_id"`
	HostMode                            types.String                                `tfsdk:"host_mode"`
	IncreaseAccessSpeed                 types.Bool                                  `tfsdk:"increase_access_speed"`
	Name                                types.String                                `tfsdk:"name"`
	RadiusAccountingEnabled             types.Bool                                  `tfsdk:"radius_accounting_enabled"`
	RadiusCoaSupportEnabled             types.Bool                                  `tfsdk:"radius_coa_support_enabled"`
	RadiusGroupAttribute                types.String                                `tfsdk:"radius_group_attribute"`
	RadiusTestingEnabled                types.Bool                                  `tfsdk:"radius_testing_enabled"`
	UrlRedirectWalledGardenEnabled      types.Bool                                  `tfsdk:"url_redirect_walled_garden_enabled"`
	VoiceVlanClients                    types.Bool                                  `tfsdk:"voice_vlan_clients"`
	Dot1xControlDirection               types.String                                `tfsdk:"dot1x_control_direction"`
	RadiusFailedAuthVlanId              types.Int64                                 `tfsdk:"radius_failed_auth_vlan_id"`
	RadiusReAuthenticationInterval      types.Int64                                 `tfsdk:"radius_re_authentication_interval"`
	RadiusCacheEnabled                  types.Bool                                  `tfsdk:"radius_cache_enabled"`
	RadiusCacheTimeout                  types.Int64                                 `tfsdk:"radius_cache_timeout"`
	RadiusCriticalAuthDataVlanId        types.Int64                                 `tfsdk:"radius_critical_auth_data_vlan_id"`
	RadiusCriticalAuthSuspendPortBounce types.Bool                                  `tfsdk:"radius_critical_auth_suspend_port_bounce"`
	RadiusCriticalAuthVoiceVlanId       types.Int64                                 `tfsdk:"radius_critical_auth_voice_vlan_id"`
	RadiusAccountingServers             []SwitchAccessPolicyRadiusAccountingServers `tfsdk:"radius_accounting_servers"`
	RadiusServers                       []SwitchAccessPolicyRadiusServers           `tfsdk:"radius_servers"`
	UrlRedirectWalledGardenRanges       types.Set                                   `tfsdk:"url_redirect_walled_garden_ranges"`
}

type SwitchAccessPolicyRadiusAccountingServers struct {
	Host                       types.String `tfsdk:"host"`
	OrganizationRadiusServerId types.String `tfsdk:"organization_radius_server_id"`
	Port                       types.Int64  `tfsdk:"port"`
	Secret                     types.String `tfsdk:"secret"`
}

type SwitchAccessPolicyRadiusServers struct {
	Host                       types.String `tfsdk:"host"`
	OrganizationRadiusServerId types.String `tfsdk:"organization_radius_server_id"`
	Port                       types.Int64  `tfsdk:"port"`
	Secret                     types.String `tfsdk:"secret"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchAccessPolicy) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/accessPolicies", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SwitchAccessPolicy) toBody(ctx context.Context, state SwitchAccessPolicy) string {
	body := ""
	if !data.AccessPolicyType.IsNull() {
		body, _ = sjson.Set(body, "accessPolicyType", data.AccessPolicyType.ValueString())
	}
	if !data.GuestPortBouncing.IsNull() {
		body, _ = sjson.Set(body, "guestPortBouncing", data.GuestPortBouncing.ValueBool())
	}
	if !data.GuestVlanId.IsNull() {
		body, _ = sjson.Set(body, "guestVlanId", data.GuestVlanId.ValueInt64())
	}
	if !data.HostMode.IsNull() {
		body, _ = sjson.Set(body, "hostMode", data.HostMode.ValueString())
	}
	if !data.IncreaseAccessSpeed.IsNull() {
		body, _ = sjson.Set(body, "increaseAccessSpeed", data.IncreaseAccessSpeed.ValueBool())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.RadiusAccountingEnabled.IsNull() {
		body, _ = sjson.Set(body, "radiusAccountingEnabled", data.RadiusAccountingEnabled.ValueBool())
	}
	if !data.RadiusCoaSupportEnabled.IsNull() {
		body, _ = sjson.Set(body, "radiusCoaSupportEnabled", data.RadiusCoaSupportEnabled.ValueBool())
	}
	if !data.RadiusGroupAttribute.IsNull() {
		body, _ = sjson.Set(body, "radiusGroupAttribute", data.RadiusGroupAttribute.ValueString())
	}
	if !data.RadiusTestingEnabled.IsNull() {
		body, _ = sjson.Set(body, "radiusTestingEnabled", data.RadiusTestingEnabled.ValueBool())
	}
	if !data.UrlRedirectWalledGardenEnabled.IsNull() {
		body, _ = sjson.Set(body, "urlRedirectWalledGardenEnabled", data.UrlRedirectWalledGardenEnabled.ValueBool())
	}
	if !data.VoiceVlanClients.IsNull() {
		body, _ = sjson.Set(body, "voiceVlanClients", data.VoiceVlanClients.ValueBool())
	}
	if !data.Dot1xControlDirection.IsNull() {
		body, _ = sjson.Set(body, "dot1x.controlDirection", data.Dot1xControlDirection.ValueString())
	}
	if !data.RadiusFailedAuthVlanId.IsNull() {
		body, _ = sjson.Set(body, "radius.failedAuthVlanId", data.RadiusFailedAuthVlanId.ValueInt64())
	}
	if !data.RadiusReAuthenticationInterval.IsNull() {
		body, _ = sjson.Set(body, "radius.reAuthenticationInterval", data.RadiusReAuthenticationInterval.ValueInt64())
	}
	if !data.RadiusCacheEnabled.IsNull() {
		body, _ = sjson.Set(body, "radius.cache.enabled", data.RadiusCacheEnabled.ValueBool())
	}
	if !data.RadiusCacheTimeout.IsNull() {
		body, _ = sjson.Set(body, "radius.cache.timeout", data.RadiusCacheTimeout.ValueInt64())
	}
	if !data.RadiusCriticalAuthDataVlanId.IsNull() {
		body, _ = sjson.Set(body, "radius.criticalAuth.dataVlanId", data.RadiusCriticalAuthDataVlanId.ValueInt64())
	}
	if !data.RadiusCriticalAuthSuspendPortBounce.IsNull() {
		body, _ = sjson.Set(body, "radius.criticalAuth.suspendPortBounce", data.RadiusCriticalAuthSuspendPortBounce.ValueBool())
	}
	if !data.RadiusCriticalAuthVoiceVlanId.IsNull() {
		body, _ = sjson.Set(body, "radius.criticalAuth.voiceVlanId", data.RadiusCriticalAuthVoiceVlanId.ValueInt64())
	}
	if len(data.RadiusAccountingServers) > 0 {
		body, _ = sjson.Set(body, "radiusAccountingServers", []interface{}{})
		for _, item := range data.RadiusAccountingServers {
			itemBody := ""
			if !item.Host.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "host", item.Host.ValueString())
			}
			if !item.OrganizationRadiusServerId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "organizationRadiusServerId", item.OrganizationRadiusServerId.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			if !item.Secret.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secret", item.Secret.ValueString())
			}
			body, _ = sjson.SetRaw(body, "radiusAccountingServers.-1", itemBody)
		}
	}
	{
		body, _ = sjson.Set(body, "radiusServers", []interface{}{})
		for _, item := range data.RadiusServers {
			itemBody := ""
			if !item.Host.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "host", item.Host.ValueString())
			}
			if !item.OrganizationRadiusServerId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "organizationRadiusServerId", item.OrganizationRadiusServerId.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			if !item.Secret.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secret", item.Secret.ValueString())
			}
			body, _ = sjson.SetRaw(body, "radiusServers.-1", itemBody)
		}
	}
	if !data.UrlRedirectWalledGardenRanges.IsNull() {
		var values []string
		data.UrlRedirectWalledGardenRanges.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "urlRedirectWalledGardenRanges", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchAccessPolicy) fromBody(ctx context.Context, res meraki.Res) {
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
		data.RadiusAccountingServers = make([]SwitchAccessPolicyRadiusAccountingServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchAccessPolicyRadiusAccountingServers{}
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
		data.RadiusServers = make([]SwitchAccessPolicyRadiusServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SwitchAccessPolicyRadiusServers{}
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SwitchAccessPolicy) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("accessPolicyType"); value.Exists() && !data.AccessPolicyType.IsNull() {
		data.AccessPolicyType = types.StringValue(value.String())
	} else {
		data.AccessPolicyType = types.StringNull()
	}
	if value := res.Get("guestPortBouncing"); value.Exists() && !data.GuestPortBouncing.IsNull() {
		data.GuestPortBouncing = types.BoolValue(value.Bool())
	} else {
		data.GuestPortBouncing = types.BoolNull()
	}
	if value := res.Get("guestVlanId"); value.Exists() && !data.GuestVlanId.IsNull() {
		data.GuestVlanId = types.Int64Value(value.Int())
	} else {
		data.GuestVlanId = types.Int64Null()
	}
	if value := res.Get("hostMode"); value.Exists() && !data.HostMode.IsNull() {
		data.HostMode = types.StringValue(value.String())
	} else {
		data.HostMode = types.StringNull()
	}
	if value := res.Get("increaseAccessSpeed"); value.Exists() && !data.IncreaseAccessSpeed.IsNull() {
		data.IncreaseAccessSpeed = types.BoolValue(value.Bool())
	} else {
		data.IncreaseAccessSpeed = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("radiusAccountingEnabled"); value.Exists() && !data.RadiusAccountingEnabled.IsNull() {
		data.RadiusAccountingEnabled = types.BoolValue(value.Bool())
	} else {
		data.RadiusAccountingEnabled = types.BoolNull()
	}
	if value := res.Get("radiusCoaSupportEnabled"); value.Exists() && !data.RadiusCoaSupportEnabled.IsNull() {
		data.RadiusCoaSupportEnabled = types.BoolValue(value.Bool())
	} else {
		data.RadiusCoaSupportEnabled = types.BoolNull()
	}
	if value := res.Get("radiusGroupAttribute"); value.Exists() && !data.RadiusGroupAttribute.IsNull() {
		data.RadiusGroupAttribute = types.StringValue(value.String())
	} else {
		data.RadiusGroupAttribute = types.StringNull()
	}
	if value := res.Get("radiusTestingEnabled"); value.Exists() && !data.RadiusTestingEnabled.IsNull() {
		data.RadiusTestingEnabled = types.BoolValue(value.Bool())
	} else {
		data.RadiusTestingEnabled = types.BoolNull()
	}
	if value := res.Get("urlRedirectWalledGardenEnabled"); value.Exists() && !data.UrlRedirectWalledGardenEnabled.IsNull() {
		data.UrlRedirectWalledGardenEnabled = types.BoolValue(value.Bool())
	} else {
		data.UrlRedirectWalledGardenEnabled = types.BoolNull()
	}
	if value := res.Get("voiceVlanClients"); value.Exists() && !data.VoiceVlanClients.IsNull() {
		data.VoiceVlanClients = types.BoolValue(value.Bool())
	} else {
		data.VoiceVlanClients = types.BoolNull()
	}
	if value := res.Get("dot1x.controlDirection"); value.Exists() && !data.Dot1xControlDirection.IsNull() {
		data.Dot1xControlDirection = types.StringValue(value.String())
	} else {
		data.Dot1xControlDirection = types.StringNull()
	}
	if value := res.Get("radius.failedAuthVlanId"); value.Exists() && !data.RadiusFailedAuthVlanId.IsNull() {
		data.RadiusFailedAuthVlanId = types.Int64Value(value.Int())
	} else {
		data.RadiusFailedAuthVlanId = types.Int64Null()
	}
	if value := res.Get("radius.reAuthenticationInterval"); value.Exists() && !data.RadiusReAuthenticationInterval.IsNull() {
		data.RadiusReAuthenticationInterval = types.Int64Value(value.Int())
	} else {
		data.RadiusReAuthenticationInterval = types.Int64Null()
	}
	if value := res.Get("radius.criticalAuth.dataVlanId"); value.Exists() && !data.RadiusCriticalAuthDataVlanId.IsNull() {
		data.RadiusCriticalAuthDataVlanId = types.Int64Value(value.Int())
	} else {
		data.RadiusCriticalAuthDataVlanId = types.Int64Null()
	}
	if value := res.Get("radius.criticalAuth.suspendPortBounce"); value.Exists() && !data.RadiusCriticalAuthSuspendPortBounce.IsNull() {
		data.RadiusCriticalAuthSuspendPortBounce = types.BoolValue(value.Bool())
	} else {
		data.RadiusCriticalAuthSuspendPortBounce = types.BoolNull()
	}
	if value := res.Get("radius.criticalAuth.voiceVlanId"); value.Exists() && !data.RadiusCriticalAuthVoiceVlanId.IsNull() {
		data.RadiusCriticalAuthVoiceVlanId = types.Int64Value(value.Int())
	} else {
		data.RadiusCriticalAuthVoiceVlanId = types.Int64Null()
	}
	{
		l := len(res.Get("radiusAccountingServers").Array())
		tflog.Debug(ctx, fmt.Sprintf("radiusAccountingServers array resizing from %d to %d", len(data.RadiusAccountingServers), l))
		if len(data.RadiusAccountingServers) > l {
			data.RadiusAccountingServers = data.RadiusAccountingServers[:l]
		}
	}
	for i := range data.RadiusAccountingServers {
		parent := &data
		data := (*parent).RadiusAccountingServers[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("radiusAccountingServers.%d", i))
		if value := res.Get("host"); value.Exists() && !data.Host.IsNull() {
			data.Host = types.StringValue(value.String())
		} else {
			data.Host = types.StringNull()
		}
		if value := res.Get("organizationRadiusServerId"); value.Exists() && !data.OrganizationRadiusServerId.IsNull() {
			data.OrganizationRadiusServerId = types.StringValue(value.String())
		} else {
			data.OrganizationRadiusServerId = types.StringNull()
		}
		if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
			data.Port = types.Int64Value(value.Int())
		} else {
			data.Port = types.Int64Null()
		}
		(*parent).RadiusAccountingServers[i] = data
	}
	{
		l := len(res.Get("radiusServers").Array())
		tflog.Debug(ctx, fmt.Sprintf("radiusServers array resizing from %d to %d", len(data.RadiusServers), l))
		if len(data.RadiusServers) > l {
			data.RadiusServers = data.RadiusServers[:l]
		}
	}
	for i := range data.RadiusServers {
		parent := &data
		data := (*parent).RadiusServers[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("radiusServers.%d", i))
		if value := res.Get("host"); value.Exists() && !data.Host.IsNull() {
			data.Host = types.StringValue(value.String())
		} else {
			data.Host = types.StringNull()
		}
		if value := res.Get("organizationRadiusServerId"); value.Exists() && !data.OrganizationRadiusServerId.IsNull() {
			data.OrganizationRadiusServerId = types.StringValue(value.String())
		} else {
			data.OrganizationRadiusServerId = types.StringNull()
		}
		if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
			data.Port = types.Int64Value(value.Int())
		} else {
			data.Port = types.Int64Null()
		}
		(*parent).RadiusServers[i] = data
	}
	if value := res.Get("urlRedirectWalledGardenRanges"); value.Exists() && !data.UrlRedirectWalledGardenRanges.IsNull() {
		data.UrlRedirectWalledGardenRanges = helpers.GetStringSet(value.Array())
	} else {
		data.UrlRedirectWalledGardenRanges = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SwitchAccessPolicy) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns
