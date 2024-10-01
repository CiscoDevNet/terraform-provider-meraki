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

type SwitchLinkAggregations struct {
	NetworkId types.String                  `tfsdk:"network_id"`
	Items     []SwitchLinkAggregationsItems `tfsdk:"items"`
}

type SwitchLinkAggregationsItems struct {
	Id                 types.String                               `tfsdk:"id"`
	SwitchPorts        []SwitchLinkAggregationsSwitchPorts        `tfsdk:"switch_ports"`
	SwitchProfilePorts []SwitchLinkAggregationsSwitchProfilePorts `tfsdk:"switch_profile_ports"`
}

type SwitchLinkAggregationsSwitchPorts struct {
	PortId types.String `tfsdk:"port_id"`
	Serial types.String `tfsdk:"serial"`
}

type SwitchLinkAggregationsSwitchProfilePorts struct {
	PortId  types.String `tfsdk:"port_id"`
	Profile types.String `tfsdk:"profile"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SwitchLinkAggregations) getPath() string {
	return fmt.Sprintf("/networks/%v/switch/linkAggregations", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SwitchLinkAggregations) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]SwitchLinkAggregationsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := SwitchLinkAggregationsItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("switchPorts"); value.Exists() && value.Value() != nil {
			data.SwitchPorts = make([]SwitchLinkAggregationsSwitchPorts, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := SwitchLinkAggregationsSwitchPorts{}
				if value := res.Get("portId"); value.Exists() && value.Value() != nil {
					data.PortId = types.StringValue(value.String())
				} else {
					data.PortId = types.StringNull()
				}
				if value := res.Get("serial"); value.Exists() && value.Value() != nil {
					data.Serial = types.StringValue(value.String())
				} else {
					data.Serial = types.StringNull()
				}
				(*parent).SwitchPorts = append((*parent).SwitchPorts, data)
				return true
			})
		}
		if value := res.Get("switchProfilePorts"); value.Exists() && value.Value() != nil {
			data.SwitchProfilePorts = make([]SwitchLinkAggregationsSwitchProfilePorts, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := SwitchLinkAggregationsSwitchProfilePorts{}
				if value := res.Get("portId"); value.Exists() && value.Value() != nil {
					data.PortId = types.StringValue(value.String())
				} else {
					data.PortId = types.StringNull()
				}
				if value := res.Get("profile"); value.Exists() && value.Value() != nil {
					data.Profile = types.StringValue(value.String())
				} else {
					data.Profile = types.StringNull()
				}
				(*parent).SwitchProfilePorts = append((*parent).SwitchProfilePorts, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
