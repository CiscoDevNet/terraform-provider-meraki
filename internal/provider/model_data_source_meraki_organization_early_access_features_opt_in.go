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
// It emits a separate model - DataSourceOrganizationEarlyAccessFeaturesOptIn - used only by the data source, always including every
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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceOrganizationEarlyAccessFeaturesOptIn struct {
	Id                   types.String `tfsdk:"id"`
	OrganizationId       types.String `tfsdk:"organization_id"`
	ShortName            types.String `tfsdk:"short_name"`
	LimitScopeToNetworks types.List   `tfsdk:"limit_scope_to_networks"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationEarlyAccessFeaturesOptIn) getPath() string {
	return fmt.Sprintf("/organizations/%v/earlyAccess/features/optIns", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// fromBody is de-generated (kept in sync manually with the resource-side model_resource.go's custom fromBody) because the
// API returns limitScopeToNetworks as a list of {"id": "..."} objects, not a plain list of strings.
func (data *DataSourceOrganizationEarlyAccessFeaturesOptIn) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
		data.ShortName = types.StringValue(value.String())
	} else {
		data.ShortName = types.StringNull()
	}
	if value := res.Get("limitScopeToNetworks"); value.Exists() && value.Value() != nil {
		data.LimitScopeToNetworks = helpers.GetStringListFromMapList(value.Array(), "id")
	} else {
		data.LimitScopeToNetworks = types.ListNull(types.StringType)
	}
}
