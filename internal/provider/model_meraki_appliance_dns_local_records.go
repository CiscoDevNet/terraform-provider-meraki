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

type ApplianceDNSLocalRecords struct {
	OrganizationId types.String                    `tfsdk:"organization_id"`
	Items          []ApplianceDNSLocalRecordsItems `tfsdk:"items"`
}

type ApplianceDNSLocalRecordsItems struct {
	Id        types.String `tfsdk:"id"`
	Address   types.String `tfsdk:"address"`
	Hostname  types.String `tfsdk:"hostname"`
	ProfileId types.String `tfsdk:"profile_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceDNSLocalRecords) getPath() string {
	return fmt.Sprintf("/organizations/%v/appliance/dns/local/records", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceDNSLocalRecords) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ApplianceDNSLocalRecordsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ApplianceDNSLocalRecordsItems{}
		data.Id = types.StringValue(res.Get("recordId").String())
		if value := res.Get("address"); value.Exists() && value.Value() != nil {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
		}
		if value := res.Get("hostname"); value.Exists() && value.Value() != nil {
			data.Hostname = types.StringValue(value.String())
		} else {
			data.Hostname = types.StringNull()
		}
		if value := res.Get("profile.id"); value.Exists() && value.Value() != nil {
			data.ProfileId = types.StringValue(value.String())
		} else {
			data.ProfileId = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
