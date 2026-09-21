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

// This template renders for every definition with a data source (gated only by NoDataSource in gen/generator.go).
// It emits a separate model - DataSourceNetworkSNMP - used only by the data source, always including every
// attribute (unlike model_resource.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model_resource.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model_resource.go's resource-side type names.

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

type DataSourceNetworkSNMP struct {
	Id                       types.String                 `tfsdk:"id"`
	NetworkId                types.String                 `tfsdk:"network_id"`
	Access                   types.String                 `tfsdk:"access"`
	CommunityString          types.String                 `tfsdk:"community_string"`
	CommunityStringWo        types.String                 `tfsdk:"community_string_wo"`
	CommunityStringWoVersion types.Int64                  `tfsdk:"community_string_wo_version"`
	Users                    []DataSourceNetworkSNMPUsers `tfsdk:"users"`
}

type DataSourceNetworkSNMPUsers struct {
	Passphrase          types.String `tfsdk:"passphrase"`
	PassphraseWo        types.String `tfsdk:"passphrase_wo"`
	PassphraseWoVersion types.Int64  `tfsdk:"passphrase_wo_version"`
	Username            types.String `tfsdk:"username"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworkSNMP) getPath() string {
	return fmt.Sprintf("/networks/%v/snmp", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkSNMP) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("access"); value.Exists() && value.Value() != nil {
		data.Access = types.StringValue(value.String())
	} else {
		data.Access = types.StringNull()
	}
	if value := res.Get("communityString"); value.Exists() && value.Value() != nil {
		data.CommunityString = types.StringValue(value.String())
	} else {
		data.CommunityString = types.StringNull()
	}
	if value := res.Get("users"); value.Exists() && value.Value() != nil {
		data.Users = make([]DataSourceNetworkSNMPUsers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DataSourceNetworkSNMPUsers{}
			if value := res.Get("passphrase"); value.Exists() && value.Value() != nil {
				data.Passphrase = types.StringValue(value.String())
			} else {
				data.Passphrase = types.StringNull()
			}
			if value := res.Get("username"); value.Exists() && value.Value() != nil {
				data.Username = types.StringValue(value.String())
			} else {
				data.Username = types.StringNull()
			}
			(*parent).Users = append((*parent).Users, data)
			return true
		})
	}
}

// End of section. //template:end fromBody
