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
// It emits a separate model - DataSourceOrganizationLoginSecurity - used only by the data source, always including every
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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceOrganizationLoginSecurity struct {
	Id                                            types.String `tfsdk:"id"`
	OrganizationId                                types.String `tfsdk:"organization_id"`
	AccountLockoutAttempts                        types.Int64  `tfsdk:"account_lockout_attempts"`
	EnforceAccountLockout                         types.Bool   `tfsdk:"enforce_account_lockout"`
	EnforceDifferentPasswords                     types.Bool   `tfsdk:"enforce_different_passwords"`
	EnforceIdleTimeout                            types.Bool   `tfsdk:"enforce_idle_timeout"`
	EnforceLoginIpRanges                          types.Bool   `tfsdk:"enforce_login_ip_ranges"`
	EnforcePasswordExpiration                     types.Bool   `tfsdk:"enforce_password_expiration"`
	EnforceStrongPasswords                        types.Bool   `tfsdk:"enforce_strong_passwords"`
	EnforceTwoFactorAuth                          types.Bool   `tfsdk:"enforce_two_factor_auth"`
	IdleTimeoutMinutes                            types.Int64  `tfsdk:"idle_timeout_minutes"`
	MinimumPasswordLength                         types.Int64  `tfsdk:"minimum_password_length"`
	NumDifferentPasswords                         types.Int64  `tfsdk:"num_different_passwords"`
	PasswordExpirationDays                        types.Int64  `tfsdk:"password_expiration_days"`
	ApiAuthenticationIpRestrictionsForKeysEnabled types.Bool   `tfsdk:"api_authentication_ip_restrictions_for_keys_enabled"`
	ApiAuthenticationIpRestrictionsForKeysRanges  types.Set    `tfsdk:"api_authentication_ip_restrictions_for_keys_ranges"`
	LoginIpRanges                                 types.Set    `tfsdk:"login_ip_ranges"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationLoginSecurity) getPath() string {
	return fmt.Sprintf("/organizations/%v/loginSecurity", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationLoginSecurity) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("accountLockoutAttempts"); value.Exists() && value.Value() != nil {
		data.AccountLockoutAttempts = types.Int64Value(value.Int())
	} else {
		data.AccountLockoutAttempts = types.Int64Null()
	}
	if value := res.Get("enforceAccountLockout"); value.Exists() && value.Value() != nil {
		data.EnforceAccountLockout = types.BoolValue(value.Bool())
	} else {
		data.EnforceAccountLockout = types.BoolNull()
	}
	if value := res.Get("enforceDifferentPasswords"); value.Exists() && value.Value() != nil {
		data.EnforceDifferentPasswords = types.BoolValue(value.Bool())
	} else {
		data.EnforceDifferentPasswords = types.BoolNull()
	}
	if value := res.Get("enforceIdleTimeout"); value.Exists() && value.Value() != nil {
		data.EnforceIdleTimeout = types.BoolValue(value.Bool())
	} else {
		data.EnforceIdleTimeout = types.BoolNull()
	}
	if value := res.Get("enforceLoginIpRanges"); value.Exists() && value.Value() != nil {
		data.EnforceLoginIpRanges = types.BoolValue(value.Bool())
	} else {
		data.EnforceLoginIpRanges = types.BoolNull()
	}
	if value := res.Get("enforcePasswordExpiration"); value.Exists() && value.Value() != nil {
		data.EnforcePasswordExpiration = types.BoolValue(value.Bool())
	} else {
		data.EnforcePasswordExpiration = types.BoolNull()
	}
	if value := res.Get("enforceStrongPasswords"); value.Exists() && value.Value() != nil {
		data.EnforceStrongPasswords = types.BoolValue(value.Bool())
	} else {
		data.EnforceStrongPasswords = types.BoolNull()
	}
	if value := res.Get("enforceTwoFactorAuth"); value.Exists() && value.Value() != nil {
		data.EnforceTwoFactorAuth = types.BoolValue(value.Bool())
	} else {
		data.EnforceTwoFactorAuth = types.BoolNull()
	}
	if value := res.Get("idleTimeoutMinutes"); value.Exists() && value.Value() != nil {
		data.IdleTimeoutMinutes = types.Int64Value(value.Int())
	} else {
		data.IdleTimeoutMinutes = types.Int64Null()
	}
	if value := res.Get("minimumPasswordLength"); value.Exists() && value.Value() != nil {
		data.MinimumPasswordLength = types.Int64Value(value.Int())
	} else {
		data.MinimumPasswordLength = types.Int64Null()
	}
	if value := res.Get("numDifferentPasswords"); value.Exists() && value.Value() != nil {
		data.NumDifferentPasswords = types.Int64Value(value.Int())
	} else {
		data.NumDifferentPasswords = types.Int64Null()
	}
	if value := res.Get("passwordExpirationDays"); value.Exists() && value.Value() != nil {
		data.PasswordExpirationDays = types.Int64Value(value.Int())
	} else {
		data.PasswordExpirationDays = types.Int64Null()
	}
	if value := res.Get("apiAuthentication.ipRestrictionsForKeys.enabled"); value.Exists() && value.Value() != nil {
		data.ApiAuthenticationIpRestrictionsForKeysEnabled = types.BoolValue(value.Bool())
	} else {
		data.ApiAuthenticationIpRestrictionsForKeysEnabled = types.BoolNull()
	}
	if value := res.Get("apiAuthentication.ipRestrictionsForKeys.ranges"); value.Exists() && value.Value() != nil {
		data.ApiAuthenticationIpRestrictionsForKeysRanges = helpers.GetStringSet(value.Array())
	} else {
		data.ApiAuthenticationIpRestrictionsForKeysRanges = types.SetNull(types.StringType)
	}
	if value := res.Get("loginIpRanges"); value.Exists() && value.Value() != nil {
		data.LoginIpRanges = helpers.GetStringSet(value.Array())
	} else {
		data.LoginIpRanges = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody
