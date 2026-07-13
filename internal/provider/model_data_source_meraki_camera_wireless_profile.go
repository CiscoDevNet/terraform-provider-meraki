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
// It emits a separate model - DataSourceCameraWirelessProfile - used only by the data source, always including every
// attribute (unlike model.go's resource-side struct, which excludes `data_source_only` attributes), since
// terraform-plugin-framework requires a model struct's tfsdk-tagged fields to exactly match the schema it's decoded
// against. Keep this file's shape in sync with model.go's `types`/`getPath`/`fromBody` sections - the only
// differences are that nothing is skipped for `.DataSourceOnly` here, and every type is prefixed with `DataSource`
// to avoid colliding with model.go's resource-side type names.

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

type DataSourceCameraWirelessProfile struct {
	Id                        types.String `tfsdk:"id"`
	NetworkId                 types.String `tfsdk:"network_id"`
	Name                      types.String `tfsdk:"name"`
	IdentityPassword          types.String `tfsdk:"identity_password"`
	IdentityPasswordWo        types.String `tfsdk:"identity_password_wo"`
	IdentityPasswordWoVersion types.Int64  `tfsdk:"identity_password_wo_version"`
	IdentityUsername          types.String `tfsdk:"identity_username"`
	SsidAuthMode              types.String `tfsdk:"ssid_auth_mode"`
	SsidEncryptionMode        types.String `tfsdk:"ssid_encryption_mode"`
	SsidName                  types.String `tfsdk:"ssid_name"`
	SsidPsk                   types.String `tfsdk:"ssid_psk"`
	SsidPskWo                 types.String `tfsdk:"ssid_psk_wo"`
	SsidPskWoVersion          types.Int64  `tfsdk:"ssid_psk_wo_version"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceCameraWirelessProfile) getPath() string {
	return fmt.Sprintf("/networks/%v/camera/wirelessProfiles", url.QueryEscape(data.NetworkId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceCameraWirelessProfile) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("identity.password"); value.Exists() && value.Value() != nil {
		data.IdentityPassword = types.StringValue(value.String())
	} else {
		data.IdentityPassword = types.StringNull()
	}
	if value := res.Get("identity.username"); value.Exists() && value.Value() != nil {
		data.IdentityUsername = types.StringValue(value.String())
	} else {
		data.IdentityUsername = types.StringNull()
	}
	if value := res.Get("ssid.authMode"); value.Exists() && value.Value() != nil {
		data.SsidAuthMode = types.StringValue(value.String())
	} else {
		data.SsidAuthMode = types.StringNull()
	}
	if value := res.Get("ssid.encryptionMode"); value.Exists() && value.Value() != nil {
		data.SsidEncryptionMode = types.StringValue(value.String())
	} else {
		data.SsidEncryptionMode = types.StringNull()
	}
	if value := res.Get("ssid.name"); value.Exists() && value.Value() != nil {
		data.SsidName = types.StringValue(value.String())
	} else {
		data.SsidName = types.StringNull()
	}
	if value := res.Get("ssid.psk"); value.Exists() && value.Value() != nil {
		data.SsidPsk = types.StringValue(value.String())
	} else {
		data.SsidPsk = types.StringNull()
	}
}

// End of section. //template:end fromBody
