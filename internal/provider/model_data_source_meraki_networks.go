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

type DataSourceNetworks struct {
	OrganizationId types.String              `tfsdk:"organization_id"`
	Items          []DataSourceNetworksItems `tfsdk:"items"`
}

type DataSourceNetworksItems struct {
	Id           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Notes        types.String `tfsdk:"notes"`
	TimeZone     types.String `tfsdk:"time_zone"`
	ProductTypes types.Set    `tfsdk:"product_types"`
	Tags         types.Set    `tfsdk:"tags"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworks) getPath() string {
	return fmt.Sprintf("/organizations/%v/networks", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworks) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceNetworksItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceNetworksItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("notes"); value.Exists() && value.Value() != nil {
			data.Notes = types.StringValue(value.String())
		} else {
			data.Notes = types.StringNull()
		}
		if value := res.Get("timeZone"); value.Exists() && value.Value() != nil {
			data.TimeZone = types.StringValue(value.String())
		} else {
			data.TimeZone = types.StringNull()
		}
		if value := res.Get("productTypes"); value.Exists() && value.Value() != nil {
			data.ProductTypes = helpers.GetStringSet(value.Array())
		} else {
			data.ProductTypes = types.SetNull(types.StringType)
		}
		if value := res.Get("tags"); value.Exists() && value.Value() != nil {
			data.Tags = helpers.GetStringSet(value.Array())
		} else {
			data.Tags = types.SetNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
