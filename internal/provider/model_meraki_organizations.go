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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type Organizations struct {
	Items []OrganizationsItems `tfsdk:"items"`
}

type OrganizationsItems struct {
	Id                types.String                     `tfsdk:"id"`
	Name              types.String                     `tfsdk:"name"`
	ManagementDetails []OrganizationsManagementDetails `tfsdk:"management_details"`
}

type OrganizationsManagementDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data Organizations) getPath() string {
	return "/organizations"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *Organizations) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]OrganizationsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := OrganizationsItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("management.details"); value.Exists() && value.Value() != nil {
			data.ManagementDetails = make([]OrganizationsManagementDetails, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := OrganizationsManagementDetails{}
				if value := res.Get("name"); value.Exists() && value.Value() != nil {
					data.Name = types.StringValue(value.String())
				} else {
					data.Name = types.StringNull()
				}
				if value := res.Get("value"); value.Exists() && value.Value() != nil {
					data.Value = types.StringValue(value.String())
				} else {
					data.Value = types.StringNull()
				}
				(*parent).ManagementDetails = append((*parent).ManagementDetails, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
