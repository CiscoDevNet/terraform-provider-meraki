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
// It emits a separate model - DataSourceOrganizationIntegrationsDeployable - used only by the data source, always including every
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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceOrganizationIntegrationsDeployable struct {
	Id                       types.String                                        `tfsdk:"id"`
	OrganizationId           types.String                                        `tfsdk:"organization_id"`
	MetaCountsItemsRemaining types.Int64                                         `tfsdk:"meta_counts_items_remaining"`
	MetaCountsItemsTotal     types.Int64                                         `tfsdk:"meta_counts_items_total"`
	Items                    []DataSourceOrganizationIntegrationsDeployableItems `tfsdk:"items"`
}

type DataSourceOrganizationIntegrationsDeployableItems struct {
	IsCiscoProduct   types.Bool   `tfsdk:"is_cisco_product"`
	IsDeployable     types.Bool   `tfsdk:"is_deployable"`
	LogoUrl          types.String `tfsdk:"logo_url"`
	Name             types.String `tfsdk:"name"`
	Provider         types.String `tfsdk:"provider"`
	RedirectUrl      types.String `tfsdk:"redirect_url"`
	ReleaseType      types.String `tfsdk:"release_type"`
	ShortDescription types.String `tfsdk:"short_description"`
	Type             types.String `tfsdk:"type"`
	Tags             types.List   `tfsdk:"tags"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationIntegrationsDeployable) getPath() string {
	return fmt.Sprintf("/organizations/%v/integrations/deployable", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationIntegrationsDeployable) fromBody(ctx context.Context, res meraki.Res) {
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
		data.Items = make([]DataSourceOrganizationIntegrationsDeployableItems, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceOrganizationIntegrationsDeployableItems{}
			if value := res.Get("isCiscoProduct"); value.Exists() && value.Value() != nil {
				data.IsCiscoProduct = types.BoolValue(value.Bool())
			} else {
				data.IsCiscoProduct = types.BoolNull()
			}
			if value := res.Get("isDeployable"); value.Exists() && value.Value() != nil {
				data.IsDeployable = types.BoolValue(value.Bool())
			} else {
				data.IsDeployable = types.BoolNull()
			}
			if value := res.Get("logoUrl"); value.Exists() && value.Value() != nil {
				data.LogoUrl = types.StringValue(value.String())
			} else {
				data.LogoUrl = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() && value.Value() != nil {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("provider"); value.Exists() && value.Value() != nil {
				data.Provider = types.StringValue(value.String())
			} else {
				data.Provider = types.StringNull()
			}
			if value := res.Get("redirectUrl"); value.Exists() && value.Value() != nil {
				data.RedirectUrl = types.StringValue(value.String())
			} else {
				data.RedirectUrl = types.StringNull()
			}
			if value := res.Get("releaseType"); value.Exists() && value.Value() != nil {
				data.ReleaseType = types.StringValue(value.String())
			} else {
				data.ReleaseType = types.StringNull()
			}
			if value := res.Get("shortDescription"); value.Exists() && value.Value() != nil {
				data.ShortDescription = types.StringValue(value.String())
			} else {
				data.ShortDescription = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && value.Value() != nil {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("tags"); value.Exists() && value.Value() != nil {
				data.Tags = helpers.GetStringList(value.Array())
			} else {
				data.Tags = types.ListNull(types.StringType)
			}
			(*parent).Items = append((*parent).Items, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
