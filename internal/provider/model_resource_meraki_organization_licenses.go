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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ResourceOrganizationLicenses struct {
	Id             types.String                        `tfsdk:"id"`
	OrganizationId types.String                        `tfsdk:"organization_id"`
	Items          []ResourceOrganizationLicensesItems `tfsdk:"items"`
}

type ResourceOrganizationLicensesItems struct {
	LicenseId    types.String `tfsdk:"license_id"`
	DeviceSerial types.String `tfsdk:"device_serial"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ResourceOrganizationLicenses) getPath() string {
	return fmt.Sprintf("/organizations/%v/licenses", url.QueryEscape(data.OrganizationId.ValueString()))
}

func (data ResourceOrganizationLicenses) getItemPath(id string) string {
	return fmt.Sprintf("/organizations/%v/licenses/%v", url.QueryEscape(data.OrganizationId.ValueString()), id)
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ResourceOrganizationLicensesItems) toBody(ctx context.Context, state ResourceOrganizationLicensesItems) string {
	body := ""
	if !data.DeviceSerial.IsNull() {
		body, _ = sjson.Set(body, "deviceSerial", data.DeviceSerial.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ResourceOrganizationLicenses) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]ResourceOrganizationLicensesItems, 0)
	if res.Get("items").Exists() {
		res = meraki.Res{Result: res.Get("items")}
	}
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := ResourceOrganizationLicensesItems{}
		if value := res.Get("id"); value.Exists() && value.Value() != nil {
			data.LicenseId = types.StringValue(value.String())
		} else {
			data.LicenseId = types.StringNull()
		}
		if value := res.Get("deviceSerial"); value.Exists() && value.Value() != nil {
			data.DeviceSerial = types.StringValue(value.String())
		} else {
			data.DeviceSerial = types.StringNull()
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
	data.Id = data.OrganizationId
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ResourceOrganizationLicenses) fromBodyPartial(ctx context.Context, res meraki.Res) {
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
				if v.Get("id").String() == (*parent).Items[i].LicenseId.ValueString() {
					res = v
					return false
				}
				return true
			},
		)
		if value := res.Get("deviceSerial"); value.Exists() && !data.DeviceSerial.IsNull() {
			data.DeviceSerial = types.StringValue(value.String())
		} else {
			data.DeviceSerial = types.StringNull()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ResourceOrganizationLicenses) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ResourceOrganizationLicenses) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges

func (data *ResourceOrganizationLicenses) hasChanges(ctx context.Context, state *ResourceOrganizationLicenses, id string) bool {
	hasChanges := false

	item := ResourceOrganizationLicensesItems{}
	for i := range data.Items {
		if data.Items[i].LicenseId.ValueString() == id {
			item = data.Items[i]
			break
		}
	}
	stateItem := ResourceOrganizationLicensesItems{}
	for i := range state.Items {
		if state.Items[i].LicenseId.ValueString() == id {
			stateItem = state.Items[i]
			break
		}
	}
	if !item.LicenseId.Equal(stateItem.LicenseId) {
		hasChanges = true
	}
	if !item.DeviceSerial.Equal(stateItem.DeviceSerial) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges
