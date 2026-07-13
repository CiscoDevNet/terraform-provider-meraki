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
// It emits a separate model - DataSourceNetworkSettings - used only by the data source, always including every
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
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceNetworkSettings struct {
	Id                                             types.String `tfsdk:"id"`
	NetworkId                                      types.String `tfsdk:"network_id"`
	LocalStatusPageEnabled                         types.Bool   `tfsdk:"local_status_page_enabled"`
	RemoteStatusPageEnabled                        types.Bool   `tfsdk:"remote_status_page_enabled"`
	LocalStatusPageAuthenticationEnabled           types.Bool   `tfsdk:"local_status_page_authentication_enabled"`
	LocalStatusPageAuthenticationPassword          types.String `tfsdk:"local_status_page_authentication_password"`
	LocalStatusPageAuthenticationPasswordWo        types.String `tfsdk:"local_status_page_authentication_password_wo"`
	LocalStatusPageAuthenticationPasswordWoVersion types.Int64  `tfsdk:"local_status_page_authentication_password_wo_version"`
	LocalStatusPageAuthenticationUsername          types.String `tfsdk:"local_status_page_authentication_username"`
	NamedVlansEnabled                              types.Bool   `tfsdk:"named_vlans_enabled"`
	SecurePortEnabled                              types.Bool   `tfsdk:"secure_port_enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceNetworkSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/settings", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceNetworkSettings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("localStatusPageEnabled"); value.Exists() && value.Value() != nil {
		data.LocalStatusPageEnabled = types.BoolValue(value.Bool())
	} else {
		data.LocalStatusPageEnabled = types.BoolNull()
	}
	if value := res.Get("remoteStatusPageEnabled"); value.Exists() && value.Value() != nil {
		data.RemoteStatusPageEnabled = types.BoolValue(value.Bool())
	} else {
		data.RemoteStatusPageEnabled = types.BoolNull()
	}
	if value := res.Get("localStatusPage.authentication.enabled"); value.Exists() && value.Value() != nil {
		data.LocalStatusPageAuthenticationEnabled = types.BoolValue(value.Bool())
	} else {
		data.LocalStatusPageAuthenticationEnabled = types.BoolNull()
	}
	if value := res.Get("localStatusPage.authentication.username"); value.Exists() && value.Value() != nil {
		data.LocalStatusPageAuthenticationUsername = types.StringValue(value.String())
	} else {
		data.LocalStatusPageAuthenticationUsername = types.StringNull()
	}
	if value := res.Get("namedVlans.enabled"); value.Exists() && value.Value() != nil {
		data.NamedVlansEnabled = types.BoolValue(value.Bool())
	} else {
		data.NamedVlansEnabled = types.BoolNull()
	}
	if value := res.Get("securePort.enabled"); value.Exists() && value.Value() != nil {
		data.SecurePortEnabled = types.BoolValue(value.Bool())
	} else {
		data.SecurePortEnabled = types.BoolNull()
	}
}

// End of section. //template:end fromBody
