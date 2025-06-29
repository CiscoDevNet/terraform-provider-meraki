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

type DataSourceNetworkMerakiAuthUsers struct {
	NetworkId types.String                            `tfsdk:"network_id"`
	Items     []DataSourceNetworkMerakiAuthUsersItems `tfsdk:"items"`
}

type DataSourceNetworkMerakiAuthUsersItems struct {
	Id                  types.String                                     `tfsdk:"id"`
	AccountType         types.String                                     `tfsdk:"account_type"`
	Email               types.String                                     `tfsdk:"email"`
	EmailPasswordToUser types.Bool                                       `tfsdk:"email_password_to_user"`
	IsAdmin             types.Bool                                       `tfsdk:"is_admin"`
	Name                types.String                                     `tfsdk:"name"`
	Password            types.String                                     `tfsdk:"password"`
	Authorizations      []DataSourceNetworkMerakiAuthUsersAuthorizations `tfsdk:"authorizations"`
}

type DataSourceNetworkMerakiAuthUsersAuthorizations struct {
	ExpiresAt  types.String `tfsdk:"expires_at"`
	SsidNumber types.Int64  `tfsdk:"ssid_number"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworkMerakiAuthUsers) getPath() string {
	return fmt.Sprintf("/networks/%v/merakiAuthUsers", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkMerakiAuthUsers) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceNetworkMerakiAuthUsersItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceNetworkMerakiAuthUsersItems{}
		data.Id = types.StringValue(res.Get("id").String())
		if value := res.Get("accountType"); value.Exists() && value.Value() != nil {
			data.AccountType = types.StringValue(value.String())
		} else {
			data.AccountType = types.StringNull()
		}
		if value := res.Get("email"); value.Exists() && value.Value() != nil {
			data.Email = types.StringValue(value.String())
		} else {
			data.Email = types.StringNull()
		}
		if value := res.Get("isAdmin"); value.Exists() && value.Value() != nil {
			data.IsAdmin = types.BoolValue(value.Bool())
		} else {
			data.IsAdmin = types.BoolNull()
		}
		if value := res.Get("name"); value.Exists() && value.Value() != nil {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("authorizations"); value.Exists() && value.Value() != nil {
			data.Authorizations = make([]DataSourceNetworkMerakiAuthUsersAuthorizations, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := DataSourceNetworkMerakiAuthUsersAuthorizations{}
				if value := res.Get("expiresAt"); value.Exists() && value.Value() != nil {
					data.ExpiresAt = types.StringValue(value.String())
				} else {
					data.ExpiresAt = types.StringNull()
				}
				if value := res.Get("ssidNumber"); value.Exists() && value.Value() != nil {
					data.SsidNumber = types.Int64Value(value.Int())
				} else {
					data.SsidNumber = types.Int64Null()
				}
				(*parent).Authorizations = append((*parent).Authorizations, data)
				return true
			})
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
