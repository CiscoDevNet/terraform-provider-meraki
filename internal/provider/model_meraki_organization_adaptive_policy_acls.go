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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationAdaptivePolicyACLs struct {
	OrganizationId types.String                          `tfsdk:"organization_id"`
	Items          []OrganizationAdaptivePolicyACLsItems `tfsdk:"items"`
}

type OrganizationAdaptivePolicyACLsItems struct {
	Id          types.String                          `tfsdk:"id"`
	Description types.String                          `tfsdk:"description"`
	IpVersion   types.String                          `tfsdk:"ip_version"`
	Name        types.String                          `tfsdk:"name"`
	Rules       []OrganizationAdaptivePolicyACLsRules `tfsdk:"rules"`
}

type OrganizationAdaptivePolicyACLsRules struct {
	DstPort  types.String `tfsdk:"dst_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
	SrcPort  types.String `tfsdk:"src_port"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationAdaptivePolicyACLs) getPath() string {
	return fmt.Sprintf("/organizations/%v/adaptivePolicy/acls", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationAdaptivePolicyACLs) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]OrganizationAdaptivePolicyACLsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := OrganizationAdaptivePolicyACLsItems{}
		data.Id = types.StringValue(res.Get("aclId").String())
		if value := res.Get("description"); value.Exists() && value.Value() != nil {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("ipVersion"); value.Exists() && value.Value() != nil {
			data.IpVersion = types.StringValue(value.String())
		} else {
			data.IpVersion = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("rules"); value.Exists() && value.Value() != nil {
			data.Rules = make([]OrganizationAdaptivePolicyACLsRules, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := OrganizationAdaptivePolicyACLsRules{}
				if value := res.Get("dstPort"); value.Exists() && value.Value() != nil {
					data.DstPort = types.StringValue(value.String())
				} else {
					data.DstPort = types.StringNull()
				}
				if value := res.Get("policy"); value.Exists() && value.Value() != nil {
					data.Policy = types.StringValue(value.String())
				} else {
					data.Policy = types.StringNull()
				}
				if value := res.Get("protocol"); value.Exists() && value.Value() != nil {
					data.Protocol = types.StringValue(value.String())
				} else {
					data.Protocol = types.StringNull()
				}
				if value := res.Get("srcPort"); value.Exists() && value.Value() != nil {
					data.SrcPort = types.StringValue(value.String())
				} else {
					data.SrcPort = types.StringNull()
				}
				(*parent).Rules = append((*parent).Rules, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
