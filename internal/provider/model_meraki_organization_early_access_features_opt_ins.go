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

type OrganizationEarlyAccessFeaturesOptIns struct {
	OrganizationId types.String                                 `tfsdk:"organization_id"`
	Items          []OrganizationEarlyAccessFeaturesOptInsItems `tfsdk:"items"`
}

type OrganizationEarlyAccessFeaturesOptInsItems struct {
	Id                   types.String `tfsdk:"id"`
	ShortName            types.String `tfsdk:"short_name"`
	LimitScopeToNetworks types.List   `tfsdk:"limit_scope_to_networks"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationEarlyAccessFeaturesOptIns) getPath() string {
	return fmt.Sprintf("/organizations/%v/earlyAccess/features/optIns", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationEarlyAccessFeaturesOptIns) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]OrganizationEarlyAccessFeaturesOptInsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := OrganizationEarlyAccessFeaturesOptInsItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
			data.ShortName = types.StringValue(value.String())
		} else {
			data.ShortName = types.StringNull()
		}
		if value := res.Get("limitScopeToNetworks"); value.Exists() && value.Value() != nil {
			data.LimitScopeToNetworks = helpers.GetStringList(value.Array())
		} else {
			data.LimitScopeToNetworks = types.ListNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
