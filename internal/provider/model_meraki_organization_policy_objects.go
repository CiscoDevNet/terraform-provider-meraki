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

type OrganizationPolicyObjects struct {
	OrganizationId types.String                     `tfsdk:"organization_id"`
	Items          []OrganizationPolicyObjectsItems `tfsdk:"items"`
}

type OrganizationPolicyObjectsItems struct {
	Id       types.String `tfsdk:"id"`
	Category types.String `tfsdk:"category"`
	Cidr     types.String `tfsdk:"cidr"`
	Fqdn     types.String `tfsdk:"fqdn"`
	Ip       types.String `tfsdk:"ip"`
	Mask     types.String `tfsdk:"mask"`
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	GroupIds types.Set    `tfsdk:"group_ids"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationPolicyObjects) getPath() string {
	return fmt.Sprintf("/organizations/%v/policyObjects", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationPolicyObjects) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]OrganizationPolicyObjectsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := OrganizationPolicyObjectsItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("category"); value.Exists() && value.Value() != nil {
			data.Category = types.StringValue(value.String())
		} else {
			data.Category = types.StringNull()
		}
		if value := res.Get("cidr"); value.Exists() && value.Value() != nil {
			data.Cidr = types.StringValue(value.String())
		} else {
			data.Cidr = types.StringNull()
		}
		if value := res.Get("fqdn"); value.Exists() && value.Value() != nil {
			data.Fqdn = types.StringValue(value.String())
		} else {
			data.Fqdn = types.StringNull()
		}
		if value := res.Get("ip"); value.Exists() && value.Value() != nil {
			data.Ip = types.StringValue(value.String())
		} else {
			data.Ip = types.StringNull()
		}
		if value := res.Get("mask"); value.Exists() && value.Value() != nil {
			data.Mask = types.StringValue(value.String())
		} else {
			data.Mask = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && value.Value() != nil {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("groupIds"); value.Exists() && value.Value() != nil {
			data.GroupIds = helpers.GetStringSet(value.Array())
		} else {
			data.GroupIds = types.SetNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
