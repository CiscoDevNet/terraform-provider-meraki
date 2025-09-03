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

type DataSourceOrganizationEarlyAccessFeatures struct {
	OrganizationId types.String                                     `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationEarlyAccessFeaturesItems `tfsdk:"items"`
}

type DataSourceOrganizationEarlyAccessFeaturesItems struct {
	Id                types.String `tfsdk:"id"`
	DocumentationLink types.String `tfsdk:"documentation_link"`
	IsOrgScopedOnly   types.Bool   `tfsdk:"is_org_scoped_only"`
	Name              types.String `tfsdk:"name"`
	PrivacyLink       types.String `tfsdk:"privacy_link"`
	ShortName         types.String `tfsdk:"short_name"`
	SupportLink       types.String `tfsdk:"support_link"`
	Topic             types.String `tfsdk:"topic"`
	DescriptionsLong  types.String `tfsdk:"descriptions_long"`
	DescriptionsShort types.String `tfsdk:"descriptions_short"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationEarlyAccessFeatures) getPath() string {
	return fmt.Sprintf("/organizations/%v/earlyAccess/features", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationEarlyAccessFeatures) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationEarlyAccessFeaturesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationEarlyAccessFeaturesItems{}
		data.Id = types.StringValue(res.Get("").String())
		if value := res.Get("documentationLink"); value.Exists() && value.Value() != nil {
			data.DocumentationLink = types.StringValue(value.String())
		} else {
			data.DocumentationLink = types.StringNull()
		}
		if value := res.Get("isOrgScopedOnly"); value.Exists() && value.Value() != nil {
			data.IsOrgScopedOnly = types.BoolValue(value.Bool())
		} else {
			data.IsOrgScopedOnly = types.BoolNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("privacyLink"); value.Exists() && value.Value() != nil {
			data.PrivacyLink = types.StringValue(value.String())
		} else {
			data.PrivacyLink = types.StringNull()
		}
		if value := res.Get("shortName"); value.Exists() && value.Value() != nil {
			data.ShortName = types.StringValue(value.String())
		} else {
			data.ShortName = types.StringNull()
		}
		if value := res.Get("supportLink"); value.Exists() && value.Value() != nil {
			data.SupportLink = types.StringValue(value.String())
		} else {
			data.SupportLink = types.StringNull()
		}
		if value := res.Get("topic"); value.Exists() && value.Value() != nil {
			data.Topic = types.StringValue(value.String())
		} else {
			data.Topic = types.StringNull()
		}
		if value := res.Get("descriptions.long"); value.Exists() && value.Value() != nil {
			data.DescriptionsLong = types.StringValue(value.String())
		} else {
			data.DescriptionsLong = types.StringNull()
		}
		if value := res.Get("descriptions.short"); value.Exists() && value.Value() != nil {
			data.DescriptionsShort = types.StringValue(value.String())
		} else {
			data.DescriptionsShort = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
