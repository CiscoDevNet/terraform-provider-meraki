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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SwitchStackRoutingStaticRoutes struct {
	NetworkId     types.String                          `tfsdk:"network_id"`
	SwitchStackId types.String                          `tfsdk:"switch_stack_id"`
	Items         []SwitchStackRoutingStaticRoutesItems `tfsdk:"items"`
}

type SwitchStackRoutingStaticRoutesItems struct {
	Id                          types.String `tfsdk:"id"`
	AdvertiseViaOspfEnabled     types.Bool   `tfsdk:"advertise_via_ospf_enabled"`
	Name                        types.String `tfsdk:"name"`
	NextHopIp                   types.String `tfsdk:"next_hop_ip"`
	PreferOverOspfRoutesEnabled types.Bool   `tfsdk:"prefer_over_ospf_routes_enabled"`
	Subnet                      types.String `tfsdk:"subnet"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchStackRoutingStaticRoutes) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/stacks/%v/routing/staticRoutes", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.SwitchStackId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchStackRoutingStaticRoutes) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]SwitchStackRoutingStaticRoutesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := SwitchStackRoutingStaticRoutesItems{}
		data.Id = types.StringValue(res.Get("staticRouteId").String())
		if value := res.Get("advertiseViaOspfEnabled"); value.Exists() && value.Value() != nil {
			data.AdvertiseViaOspfEnabled = types.BoolValue(value.Bool())
		} else {
			data.AdvertiseViaOspfEnabled = types.BoolNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("nextHopIp"); value.Exists() && value.Value() != nil {
			data.NextHopIp = types.StringValue(value.String())
		} else {
			data.NextHopIp = types.StringNull()
		}
		if value := res.Get("preferOverOspfRoutesEnabled"); value.Exists() && value.Value() != nil {
			data.PreferOverOspfRoutesEnabled = types.BoolValue(value.Bool())
		} else {
			data.PreferOverOspfRoutesEnabled = types.BoolNull()
		}
		if value := res.Get("subnet"); value.Exists() && value.Value() != nil {
			data.Subnet = types.StringValue(value.String())
		} else {
			data.Subnet = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
