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
// It emits a separate model - DataSourceApplianceSiteToSiteVPN - used only by the data source, always including every
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

type DataSourceApplianceSiteToSiteVPN struct {
	Id                 types.String                              `tfsdk:"id"`
	NetworkId          types.String                              `tfsdk:"network_id"`
	Mode               types.String                              `tfsdk:"mode"`
	SubnetNatIsAllowed types.Bool                                `tfsdk:"subnet_nat_is_allowed"`
	Hubs               []DataSourceApplianceSiteToSiteVPNHubs    `tfsdk:"hubs"`
	Subnets            []DataSourceApplianceSiteToSiteVPNSubnets `tfsdk:"subnets"`
}

type DataSourceApplianceSiteToSiteVPNHubs struct {
	HubId           types.String `tfsdk:"hub_id"`
	UseDefaultRoute types.Bool   `tfsdk:"use_default_route"`
}

type DataSourceApplianceSiteToSiteVPNSubnets struct {
	LocalSubnet     types.String `tfsdk:"local_subnet"`
	UseVpn          types.Bool   `tfsdk:"use_vpn"`
	NatEnabled      types.Bool   `tfsdk:"nat_enabled"`
	NatRemoteSubnet types.String `tfsdk:"nat_remote_subnet"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceSiteToSiteVPN) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/vpn/siteToSiteVpn", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceSiteToSiteVPN) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("mode"); value.Exists() && value.Value() != nil {
		data.Mode = types.StringValue(value.String())
	} else {
		data.Mode = types.StringNull()
	}
	if value := res.Get("subnet.nat.isAllowed"); value.Exists() && value.Value() != nil {
		data.SubnetNatIsAllowed = types.BoolValue(value.Bool())
	} else {
		data.SubnetNatIsAllowed = types.BoolNull()
	}
	if value := res.Get("hubs"); value.Exists() && value.Value() != nil {
		data.Hubs = make([]DataSourceApplianceSiteToSiteVPNHubs, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceSiteToSiteVPNHubs{}
			if value := res.Get("hubId"); value.Exists() && value.Value() != nil {
				data.HubId = types.StringValue(value.String())
			} else {
				data.HubId = types.StringNull()
			}
			if value := res.Get("useDefaultRoute"); value.Exists() && value.Value() != nil {
				data.UseDefaultRoute = types.BoolValue(value.Bool())
			} else {
				data.UseDefaultRoute = types.BoolNull()
			}
			(*parent).Hubs = append((*parent).Hubs, data)
			return true
		})
	}
	if value := res.Get("subnets"); value.Exists() && value.Value() != nil {
		data.Subnets = make([]DataSourceApplianceSiteToSiteVPNSubnets, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceSiteToSiteVPNSubnets{}
			if value := res.Get("localSubnet"); value.Exists() && value.Value() != nil {
				data.LocalSubnet = types.StringValue(value.String())
			} else {
				data.LocalSubnet = types.StringNull()
			}
			if value := res.Get("useVpn"); value.Exists() && value.Value() != nil {
				data.UseVpn = types.BoolValue(value.Bool())
			} else {
				data.UseVpn = types.BoolNull()
			}
			if value := res.Get("nat.enabled"); value.Exists() && value.Value() != nil {
				data.NatEnabled = types.BoolValue(value.Bool())
			} else {
				data.NatEnabled = types.BoolNull()
			}
			if value := res.Get("nat.remoteSubnet"); value.Exists() && value.Value() != nil {
				data.NatRemoteSubnet = types.StringValue(value.String())
			} else {
				data.NatRemoteSubnet = types.StringNull()
			}
			(*parent).Subnets = append((*parent).Subnets, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
