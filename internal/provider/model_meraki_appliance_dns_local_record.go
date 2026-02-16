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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceDNSLocalRecord struct {
	Id             types.String `tfsdk:"id"`
	OrganizationId types.String `tfsdk:"organization_id"`
	Address        types.String `tfsdk:"address"`
	Hostname       types.String `tfsdk:"hostname"`
	ProfileId      types.String `tfsdk:"profile_id"`
}

type ApplianceDNSLocalRecordIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceDNSLocalRecord) getPath() string {
	return fmt.Sprintf("/organizations/%v/appliance/dns/local/records", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceDNSLocalRecord) toBody(ctx context.Context, state ApplianceDNSLocalRecord) string {
	body := ""
	if !data.Address.IsNull() {
		body, _ = sjson.Set(body, "address", data.Address.ValueString())
	}
	if !data.Hostname.IsNull() {
		body, _ = sjson.Set(body, "hostname", data.Hostname.ValueString())
	}
	if !data.ProfileId.IsNull() {
		body, _ = sjson.Set(body, "profile.id", data.ProfileId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceDNSLocalRecord) fromBody(ctx context.Context, res meraki.Res) {
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
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceDNSLocalRecord) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("address"); value.Exists() && !data.Address.IsNull() {
		data.Address = types.StringValue(value.String())
	} else {
		data.Address = types.StringNull()
	}
	if value := res.Get("hostname"); value.Exists() && !data.Hostname.IsNull() {
		data.Hostname = types.StringValue(value.String())
	} else {
		data.Hostname = types.StringNull()
	}
	if value := res.Get("profile.id"); value.Exists() && !data.ProfileId.IsNull() {
		data.ProfileId = types.StringValue(value.String())
	} else {
		data.ProfileId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceDNSLocalRecord) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *ApplianceDNSLocalRecordIdentity) toIdentity(ctx context.Context, plan *ApplianceDNSLocalRecord) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *ApplianceDNSLocalRecord) fromIdentity(ctx context.Context, identity *ApplianceDNSLocalRecordIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceDNSLocalRecord) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
