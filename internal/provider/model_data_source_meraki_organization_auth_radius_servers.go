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

type DataSourceOrganizationAuthRADIUSServers struct {
	OrganizationId types.String                                   `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationAuthRADIUSServersItems `tfsdk:"items"`
}

type DataSourceOrganizationAuthRADIUSServersItems struct {
	Id      types.String                                   `tfsdk:"id"`
	Address types.String                                   `tfsdk:"address"`
	Name    types.String                                   `tfsdk:"name"`
	Secret  types.String                                   `tfsdk:"secret"`
	Modes   []DataSourceOrganizationAuthRADIUSServersModes `tfsdk:"modes"`
}

type DataSourceOrganizationAuthRADIUSServersModes struct {
	Mode types.String `tfsdk:"mode"`
	Port types.Int64  `tfsdk:"port"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationAuthRADIUSServers) getPath() string {
	return fmt.Sprintf("/organizations/%v/auth/radius/servers", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationAuthRADIUSServers) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationAuthRADIUSServersItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationAuthRADIUSServersItems{}
		data.Id = types.StringValue(res.Get("serverId").String())
		if value := res.Get("address"); value.Exists() && value.Value() != nil {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("modes"); value.Exists() && value.Value() != nil {
			data.Modes = make([]DataSourceOrganizationAuthRADIUSServersModes, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceOrganizationAuthRADIUSServersModes{}
				if value := res.Get("mode"); value.Exists() && value.Value() != nil {
					data.Mode = types.StringValue(value.String())
				} else {
					data.Mode = types.StringNull()
				}
				if value := res.Get("port"); value.Exists() && value.Value() != nil {
					data.Port = types.Int64Value(value.Int())
				} else {
					data.Port = types.Int64Null()
				}
				(*parent).Modes = append((*parent).Modes, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
