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

type DataSourceSwitchRoutingMulticastRendezvousPoints struct {
	NetworkId types.String                                            `tfsdk:"network_id"`
	Items     []DataSourceSwitchRoutingMulticastRendezvousPointsItems `tfsdk:"items"`
}

type DataSourceSwitchRoutingMulticastRendezvousPointsItems struct {
	Id             types.String `tfsdk:"id"`
	InterfaceIp    types.String `tfsdk:"interface_ip"`
	MulticastGroup types.String `tfsdk:"multicast_group"`
	VrfName        types.String `tfsdk:"vrf_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceSwitchRoutingMulticastRendezvousPoints) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/routing/multicast/rendezvousPoints", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceSwitchRoutingMulticastRendezvousPoints) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceSwitchRoutingMulticastRendezvousPointsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceSwitchRoutingMulticastRendezvousPointsItems{}
		data.Id = types.StringValue(res.Get("rendezvousPointId").String())
		if value := res.Get("interfaceIp"); value.Exists() && value.Value() != nil {
			data.InterfaceIp = types.StringValue(value.String())
		} else {
			data.InterfaceIp = types.StringNull()
		}
		if value := res.Get("multicastGroup"); value.Exists() && value.Value() != nil {
			data.MulticastGroup = types.StringValue(value.String())
		} else {
			data.MulticastGroup = types.StringNull()
		}
		if value := res.Get("vrf.name"); value.Exists() && value.Value() != nil {
			data.VrfName = types.StringValue(value.String())
		} else {
			data.VrfName = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
