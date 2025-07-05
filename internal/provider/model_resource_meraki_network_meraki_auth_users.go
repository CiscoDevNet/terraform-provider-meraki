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
	"slices"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceNetworkMerakiAuthUsers struct {
	Id             types.String                          `tfsdk:"id"`
	OrganizationId types.String                          `tfsdk:"organization_id"`
	NetworkId      types.String                          `tfsdk:"network_id"`
	Items          []ResourceNetworkMerakiAuthUsersItems `tfsdk:"items"`
}

type ResourceNetworkMerakiAuthUsersItems struct {
	Id                  types.String                                   `tfsdk:"id"`
	AccountType         types.String                                   `tfsdk:"account_type"`
	Email               types.String                                   `tfsdk:"email"`
	EmailPasswordToUser types.Bool                                     `tfsdk:"email_password_to_user"`
	IsAdmin             types.Bool                                     `tfsdk:"is_admin"`
	Name                types.String                                   `tfsdk:"name"`
	Password            types.String                                   `tfsdk:"password"`
	Authorizations      []ResourceNetworkMerakiAuthUsersAuthorizations `tfsdk:"authorizations"`
}

type ResourceNetworkMerakiAuthUsersAuthorizations struct {
	ExpiresAt  types.String `tfsdk:"expires_at"`
	SsidNumber types.Int64  `tfsdk:"ssid_number"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceNetworkMerakiAuthUsers) getPath() string {
	return fmt.Sprintf("/networks/%v/merakiAuthUsers", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceNetworkMerakiAuthUsersItems) toBody(ctx context.Context, state ResourceNetworkMerakiAuthUsersItems) string {
	body := ""
	if !data.AccountType.IsNull() {
		body, _ = sjson.Set(body, "accountType", data.AccountType.ValueString())
	}
	if !data.Email.IsNull() {
		body, _ = sjson.Set(body, "email", data.Email.ValueString())
	}
	if !data.EmailPasswordToUser.IsNull() {
		body, _ = sjson.Set(body, "emailPasswordToUser", data.EmailPasswordToUser.ValueBool())
	}
	if !data.IsAdmin.IsNull() {
		body, _ = sjson.Set(body, "isAdmin", data.IsAdmin.ValueBool())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Password.IsNull() {
		body, _ = sjson.Set(body, "password", data.Password.ValueString())
	}
	{
		body, _ = sjson.Set(body, "authorizations", []interface{}{})
		for _, item := range data.Authorizations {
			itemBody := ""
			if !item.ExpiresAt.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "expiresAt", item.ExpiresAt.ValueString())
			}
			if !item.SsidNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ssidNumber", item.SsidNumber.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "authorizations.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceNetworkMerakiAuthUsers) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceNetworkMerakiAuthUsersItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceNetworkMerakiAuthUsersItems{}
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
			data.Authorizations = make([]ResourceNetworkMerakiAuthUsersAuthorizations, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ResourceNetworkMerakiAuthUsersAuthorizations{}
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceNetworkMerakiAuthUsers) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("id").String() == (*parent).Items[i].Id.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
		if value := res.Get("accountType"); value.Exists() && !data.AccountType.IsNull() {
			data.AccountType = types.StringValue(value.String())
		} else {
			data.AccountType = types.StringNull()
		}
		if value := res.Get("email"); value.Exists() && !data.Email.IsNull() {
			data.Email = types.StringValue(value.String())
		} else {
			data.Email = types.StringNull()
		}
		if value := res.Get("isAdmin"); value.Exists() && !data.IsAdmin.IsNull() {
			data.IsAdmin = types.BoolValue(value.Bool())
		} else {
			data.IsAdmin = types.BoolNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		for i := 0; i < len(data.Authorizations); i++ {
			keys := [...]string{"ssidNumber"}
			keyValues := [...]string{strconv.FormatInt(data.Authorizations[i].SsidNumber.ValueInt64(), 10)}

			parent := &data
			data := (*parent).Authorizations[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("authorizations").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						if v.Get(keys[ik]).String() != keyValues[ik] {
							found = false
							break
						}
						found = true
					}
					if found {
						res = v
						return false
					}
					return true
				},
			)
			if !res.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing Authorizations[%d] = %+v",
					i,
					(*parent).Authorizations[i],
				))
				(*parent).Authorizations = slices.Delete((*parent).Authorizations, i, i+1)
				i--

				continue
			}
			if value := res.Get("expiresAt"); value.Exists() && !data.ExpiresAt.IsNull() {
				data.ExpiresAt = types.StringValue(value.String())
			} else {
				data.ExpiresAt = types.StringNull()
			}
			if value := res.Get("ssidNumber"); value.Exists() && !data.SsidNumber.IsNull() {
				data.SsidNumber = types.Int64Value(value.Int())
			} else {
				data.SsidNumber = types.Int64Null()
			}
			(*parent).Authorizations[i] = data
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceNetworkMerakiAuthUsers) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceNetworkMerakiAuthUsers) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceNetworkMerakiAuthUsers) hasChanges(ctx context.Context, state *ResourceNetworkMerakiAuthUsers, id string) bool {
	hasChanges := false

	item := ResourceNetworkMerakiAuthUsersItems{}
	for i := range data.Items {
		if data.Items[i].Id.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceNetworkMerakiAuthUsersItems{}
	for i := range state.Items {
		if state.Items[i].Id.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.AccountType.Equal(stateItem.AccountType) {
		hasChanges = true
	}
	if !item.Email.Equal(stateItem.Email) {
		hasChanges = true
	}
	if !item.EmailPasswordToUser.Equal(stateItem.EmailPasswordToUser) {
		hasChanges = true
	}
	if !item.IsAdmin.Equal(stateItem.IsAdmin) {
		hasChanges = true
	}
	if !item.Name.Equal(stateItem.Name) {
		hasChanges = true
	}
	if !item.Password.Equal(stateItem.Password) {
		hasChanges = true
	}
	if len(item.Authorizations) != len(stateItem.Authorizations) {
		hasChanges = true
	} else {
		for i := range item.Authorizations {
			if !item.Authorizations[i].ExpiresAt.Equal(stateItem.Authorizations[i].ExpiresAt) {
				hasChanges = true
			}
			if !item.Authorizations[i].SsidNumber.Equal(stateItem.Authorizations[i].SsidNumber) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges
