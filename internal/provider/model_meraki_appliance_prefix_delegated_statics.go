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
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type AppliancePrefixDelegatedStatics struct {
	NetworkId types.String                           `tfsdk:"network_id"`
	Items     []AppliancePrefixDelegatedStaticsItems `tfsdk:"items"`
}

type AppliancePrefixDelegatedStaticsItems struct {
	Id               types.String `tfsdk:"id"`
	Description      types.String `tfsdk:"description"`
	Prefix           types.String `tfsdk:"prefix"`
	OriginType       types.String `tfsdk:"origin_type"`
	OriginInterfaces types.List   `tfsdk:"origin_interfaces"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data AppliancePrefixDelegatedStatics) getPath() string {
	return fmt.Sprintf("/networks/%v/appliance/prefixes/delegated/statics", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *AppliancePrefixDelegatedStatics) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]AppliancePrefixDelegatedStaticsItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := AppliancePrefixDelegatedStaticsItems{}
		data.Id = types.StringValue(res.Get("staticDelegatedPrefixId").String())
		if value := res.Get("description"); value.Exists() && value.Value() != nil {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("prefix"); value.Exists() && value.Value() != nil {
			data.Prefix = types.StringValue(value.String())
		} else {
			data.Prefix = types.StringNull()
		}
		if value := res.Get("origin.type"); value.Exists() && value.Value() != nil {
			data.OriginType = types.StringValue(value.String())
		} else {
			data.OriginType = types.StringNull()
		}
		if value := res.Get("origin.interfaces"); value.Exists() && value.Value() != nil {
			data.OriginInterfaces = helpers.GetStringList(value.Array())
		} else {
			data.OriginInterfaces = types.ListNull(types.StringType)
		}
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
