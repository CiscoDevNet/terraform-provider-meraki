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
// It emits a separate model - DataSourceOrganizationSaseNetworksEligible - used only by the data source, always including every
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

type DataSourceOrganizationSaseNetworksEligible struct {
	Id                       types.String                                      `tfsdk:"id"`
	OrganizationId           types.String                                      `tfsdk:"organization_id"`
	MetaCountsItemsRemaining types.Int64                                       `tfsdk:"meta_counts_items_remaining"`
	MetaCountsItemsTotal     types.Int64                                       `tfsdk:"meta_counts_items_total"`
	Items                    []DataSourceOrganizationSaseNetworksEligibleItems `tfsdk:"items"`
}

type DataSourceOrganizationSaseNetworksEligibleItems struct {
	Name                       types.String `tfsdk:"name"`
	NetworkId                  types.String `tfsdk:"network_id"`
	Type                       types.String `tfsdk:"type"`
	AddressStreet              types.String `tfsdk:"address_street"`
	DevicePrimaryModel         types.String `tfsdk:"device_primary_model"`
	RegionName                 types.String `tfsdk:"region_name"`
	RoutingDefaultRouteEnabled types.Bool   `tfsdk:"routing_default_route_enabled"`
	VpnType                    types.String `tfsdk:"vpn_type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationSaseNetworksEligible) getPath() string {
	return fmt.Sprintf("/organizations/%v/sase/networks/eligible", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationSaseNetworksEligible) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("meta.counts.items.remaining"); value.Exists() && value.Value() != nil {
		data.MetaCountsItemsRemaining = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsRemaining = types.Int64Null()
	}
	if value := res.Get("meta.counts.items.total"); value.Exists() && value.Value() != nil {
		data.MetaCountsItemsTotal = types.Int64Value(value.Int())
	} else {
		data.MetaCountsItemsTotal = types.Int64Null()
	}
	if value := res.Get("items"); value.Exists() && value.Value() != nil {
		data.Items = make([]DataSourceOrganizationSaseNetworksEligibleItems, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceOrganizationSaseNetworksEligibleItems{}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("networkId"); value.Exists() && value.Value() != nil {
				data.NetworkId = types.StringValue(value.String())
			} else {
				data.NetworkId = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && value.Value() != nil {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("address.street"); value.Exists() && value.Value() != nil {
				data.AddressStreet = types.StringValue(value.String())
			} else {
				data.AddressStreet = types.StringNull()
			}
			if value := res.Get("device.primary.model"); value.Exists() && value.Value() != nil {
				data.DevicePrimaryModel = types.StringValue(value.String())
			} else {
				data.DevicePrimaryModel = types.StringNull()
			}
			if value := res.Get("region.name"); value.Exists() && value.Value() != nil {
				data.RegionName = types.StringValue(value.String())
			} else {
				data.RegionName = types.StringNull()
			}
			if value := res.Get("routing.defaultRoute.enabled"); value.Exists() && value.Value() != nil {
				data.RoutingDefaultRouteEnabled = types.BoolValue(value.Bool())
			} else {
				data.RoutingDefaultRouteEnabled = types.BoolNull()
			}
			if value := res.Get("vpn.type"); value.Exists() && value.Value() != nil {
				data.VpnType = types.StringValue(value.String())
			} else {
				data.VpnType = types.StringNull()
			}
			(*parent).Items = append((*parent).Items, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
