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

type SMAdminRoles struct {
	OrganizationId types.String        `tfsdk:"organization_id"`
	Items          []SMAdminRolesItems `tfsdk:"items"`
}

type SMAdminRolesItems struct {
	Id    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Scope types.String `tfsdk:"scope"`
	Tags  types.List   `tfsdk:"tags"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SMAdminRoles) getPath() string {
	return fmt.Sprintf("/organizations/%v/sm/admins/roles", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SMAdminRoles) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]SMAdminRolesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := SMAdminRolesItems{}
		data.Id = types.StringValue(res.Get("roleId").String())
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("scope"); value.Exists() && value.Value() != nil {
			data.Scope = types.StringValue(value.String())
		} else {
			data.Scope = types.StringNull()
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

// End of section. //template:end fromBody
