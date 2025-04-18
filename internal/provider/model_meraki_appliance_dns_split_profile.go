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
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ApplianceDNSSplitProfile struct {
	Id                   types.String `tfsdk:"id"`
	OrganizationId       types.String `tfsdk:"organization_id"`
	Name                 types.String `tfsdk:"name"`
	NameserversAddresses types.List   `tfsdk:"nameservers_addresses"`
	Hostnames            types.List   `tfsdk:"hostnames"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ApplianceDNSSplitProfile) getPath() string {
	return fmt.Sprintf("/organizations/%v/appliance/dns/split/profiles", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ApplianceDNSSplitProfile) toBody(ctx context.Context, state ApplianceDNSSplitProfile) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.NameserversAddresses.IsNull() {
		var values []string
		data.NameserversAddresses.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "nameservers.addresses", values)
	}
	if !data.Hostnames.IsNull() {
		var values []string
		data.Hostnames.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "hostnames", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ApplianceDNSSplitProfile) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("nameservers.addresses"); value.Exists() && value.Value() != nil {
		data.NameserversAddresses = helpers.GetStringList(value.Array())
	} else {
		data.NameserversAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get("hostnames"); value.Exists() && value.Value() != nil {
		data.Hostnames = helpers.GetStringList(value.Array())
	} else {
		data.Hostnames = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ApplianceDNSSplitProfile) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("nameservers.addresses"); value.Exists() && !data.NameserversAddresses.IsNull() {
		data.NameserversAddresses = helpers.GetStringList(value.Array())
	} else {
		data.NameserversAddresses = types.ListNull(types.StringType)
	}
	if value := res.Get("hostnames"); value.Exists() && !data.Hostnames.IsNull() {
		data.Hostnames = helpers.GetStringList(value.Array())
	} else {
		data.Hostnames = types.ListNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ApplianceDNSSplitProfile) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data ApplianceDNSSplitProfile) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
