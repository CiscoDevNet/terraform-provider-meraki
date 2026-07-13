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
// It emits a separate model - DataSourceApplianceVPNSiteToSiteIPsecPeersSLAs - used only by the data source, always including every
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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceApplianceVPNSiteToSiteIPsecPeersSLAs struct {
	Id             types.String                                          `tfsdk:"id"`
	OrganizationId types.String                                          `tfsdk:"organization_id"`
	Items          []DataSourceApplianceVPNSiteToSiteIPsecPeersSLAsItems `tfsdk:"items"`
}

type DataSourceApplianceVPNSiteToSiteIPsecPeersSLAsItems struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Uri  types.String `tfsdk:"uri"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceApplianceVPNSiteToSiteIPsecPeersSLAs) getPath() string {
	return fmt.Sprintf("/organizations/%v/appliance/vpn/siteToSite/ipsec/peers/slas", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceApplianceVPNSiteToSiteIPsecPeersSLAs) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("items"); value.Exists() && value.Value() != nil {
		data.Items = make([]DataSourceApplianceVPNSiteToSiteIPsecPeersSLAsItems, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceApplianceVPNSiteToSiteIPsecPeersSLAsItems{}
			if value := res.Get("id"); value.Exists() && value.Value() != nil {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("uri"); value.Exists() && value.Value() != nil {
				data.Uri = types.StringValue(value.String())
			} else {
				data.Uri = types.StringNull()
			}
			(*parent).Items = append((*parent).Items, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
