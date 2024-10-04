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

type ApplianceStaticRoutes struct {
	NetworkId types.String                 `tfsdk:"network_id"`
	Items     []ApplianceStaticRoutesItems `tfsdk:"items"`
}

type ApplianceStaticRoutesItems struct {
	Id            types.String `tfsdk:"id"`
	GatewayIp     types.String `tfsdk:"gateway_ip"`
	GatewayVlanId types.String `tfsdk:"gateway_vlan_id"`
	Name          types.String `tfsdk:"name"`
	Subnet        types.String `tfsdk:"subnet"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceStaticRoutes) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/staticRoutes", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceStaticRoutes) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ApplianceStaticRoutesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ApplianceStaticRoutesItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("gatewayIp"); value.Exists() && value.Value() != nil {
			data.GatewayIp = types.StringValue(value.String())
		} else {
			data.GatewayIp = types.StringNull()
		}
		if value := res.Get("gatewayVlanId"); value.Exists() && value.Value() != nil {
			data.GatewayVlanId = types.StringValue(value.String())
		} else {
			data.GatewayVlanId = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
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
