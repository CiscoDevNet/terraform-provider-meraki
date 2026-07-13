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
// It emits a separate model - DataSourceOrganizationIntegrationsXDRNetworks - used only by the data source, always including every
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

type DataSourceOrganizationIntegrationsXDRNetworks struct {
	Id             types.String                                            `tfsdk:"id"`
	OrganizationId types.String                                            `tfsdk:"organization_id"`
	Networks       []DataSourceOrganizationIntegrationsXDRNetworksNetworks `tfsdk:"networks"`
}

type DataSourceOrganizationIntegrationsXDRNetworksNetworks struct {
	NetworkId    types.String `tfsdk:"network_id"`
	ProductTypes types.Set    `tfsdk:"product_types"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationIntegrationsXDRNetworks) getPath() string {
	return fmt.Sprintf("/organizations/%v/integrations/xdr/networks/enable", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// getNetworksPath is custom (kept in sync manually with the resource-side model.go's equivalent) - the data
// source reads the plain networks-list GET endpoint, distinct from getPath's enable/disable action endpoint.
func (data DataSourceOrganizationIntegrationsXDRNetworks) getNetworksPath() string {
	return fmt.Sprintf("/organizations/%v/integrations/xdr/networks", url.QueryEscape(data.OrganizationId.ValueString()))
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationIntegrationsXDRNetworks) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("networks"); value.Exists() && value.Value() != nil {
		data.Networks = make([]DataSourceOrganizationIntegrationsXDRNetworksNetworks, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceOrganizationIntegrationsXDRNetworksNetworks{}
			if value := res.Get("networkId"); value.Exists() && value.Value() != nil {
				data.NetworkId = types.StringValue(value.String())
			} else {
				data.NetworkId = types.StringNull()
			}
			if value := res.Get("productTypes"); value.Exists() && value.Value() != nil {
				data.ProductTypes = helpers.GetStringSet(value.Array())
			} else {
				data.ProductTypes = types.SetNull(types.StringType)
			}
			(*parent).Networks = append((*parent).Networks, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
