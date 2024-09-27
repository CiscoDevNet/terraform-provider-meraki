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
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type OrganizationLoginSecurity struct {
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
	NumDifferentPasswords                         types.Int64  `tfsdk:"num_different_passwords"`
	PasswordExpirationDays                        types.Int64  `tfsdk:"password_expiration_days"`
	ApiAuthenticationIpRestrictionsForKeysEnabled types.Bool   `tfsdk:"api_authentication_ip_restrictions_for_keys_enabled"`
	ApiAuthenticationIpRestrictionsForKeysRanges  types.Set    `tfsdk:"api_authentication_ip_restrictions_for_keys_ranges"`
	LoginIpRanges                                 types.Set    `tfsdk:"login_ip_ranges"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationLoginSecurity) getPath() string {
	return fmt.Sprintf("/organizations/%v/loginSecurity", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationLoginSecurity) toBody(ctx context.Context, state OrganizationLoginSecurity) string {
	body := ""
	if !data.AccountLockoutAttempts.IsNull() {
		body, _ = sjson.Set(body, "accountLockoutAttempts", data.AccountLockoutAttempts.ValueInt64())
	}
	if !data.EnforceAccountLockout.IsNull() {
		body, _ = sjson.Set(body, "enforceAccountLockout", data.EnforceAccountLockout.ValueBool())
	}
	if !data.EnforceDifferentPasswords.IsNull() {
		body, _ = sjson.Set(body, "enforceDifferentPasswords", data.EnforceDifferentPasswords.ValueBool())
	}
	if !data.EnforceIdleTimeout.IsNull() {
		body, _ = sjson.Set(body, "enforceIdleTimeout", data.EnforceIdleTimeout.ValueBool())
	}
	if !data.EnforceLoginIpRanges.IsNull() {
		body, _ = sjson.Set(body, "enforceLoginIpRanges", data.EnforceLoginIpRanges.ValueBool())
	}
	if !data.EnforcePasswordExpiration.IsNull() {
		body, _ = sjson.Set(body, "enforcePasswordExpiration", data.EnforcePasswordExpiration.ValueBool())
	}
	if !data.EnforceStrongPasswords.IsNull() {
		body, _ = sjson.Set(body, "enforceStrongPasswords", data.EnforceStrongPasswords.ValueBool())
	}
	if !data.EnforceTwoFactorAuth.IsNull() {
		body, _ = sjson.Set(body, "enforceTwoFactorAuth", data.EnforceTwoFactorAuth.ValueBool())
	}
	if !data.IdleTimeoutMinutes.IsNull() {
		body, _ = sjson.Set(body, "idleTimeoutMinutes", data.IdleTimeoutMinutes.ValueInt64())
	}
	if !data.NumDifferentPasswords.IsNull() {
		body, _ = sjson.Set(body, "numDifferentPasswords", data.NumDifferentPasswords.ValueInt64())
	}
	if !data.PasswordExpirationDays.IsNull() {
		body, _ = sjson.Set(body, "passwordExpirationDays", data.PasswordExpirationDays.ValueInt64())
	}
	if !data.ApiAuthenticationIpRestrictionsForKeysEnabled.IsNull() {
		body, _ = sjson.Set(body, "apiAuthentication.ipRestrictionsForKeys.enabled", data.ApiAuthenticationIpRestrictionsForKeysEnabled.ValueBool())
	}
	if !data.ApiAuthenticationIpRestrictionsForKeysRanges.IsNull() {
		var values []string
		data.ApiAuthenticationIpRestrictionsForKeysRanges.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "apiAuthentication.ipRestrictionsForKeys.ranges", values)
	}
	if !data.LoginIpRanges.IsNull() {
		var values []string
		data.LoginIpRanges.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "loginIpRanges", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationLoginSecurity) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("accountLockoutAttempts"); value.Exists() {
		data.AccountLockoutAttempts = types.Int64Value(value.Int())
	} else {
		data.AccountLockoutAttempts = types.Int64Null()
	}
	if value := res.Get("enforceAccountLockout"); value.Exists() {
		data.EnforceAccountLockout = types.BoolValue(value.Bool())
	} else {
		data.EnforceAccountLockout = types.BoolNull()
	}
	if value := res.Get("enforceDifferentPasswords"); value.Exists() {
		data.EnforceDifferentPasswords = types.BoolValue(value.Bool())
	} else {
		data.EnforceDifferentPasswords = types.BoolNull()
	}
	if value := res.Get("enforceIdleTimeout"); value.Exists() {
		data.EnforceIdleTimeout = types.BoolValue(value.Bool())
	} else {
		data.EnforceIdleTimeout = types.BoolNull()
	}
	if value := res.Get("enforceLoginIpRanges"); value.Exists() {
		data.EnforceLoginIpRanges = types.BoolValue(value.Bool())
	} else {
		data.EnforceLoginIpRanges = types.BoolNull()
	}
	if value := res.Get("enforcePasswordExpiration"); value.Exists() {
		data.EnforcePasswordExpiration = types.BoolValue(value.Bool())
	} else {
		data.EnforcePasswordExpiration = types.BoolNull()
	}
	if value := res.Get("enforceStrongPasswords"); value.Exists() {
		data.EnforceStrongPasswords = types.BoolValue(value.Bool())
	} else {
		data.EnforceStrongPasswords = types.BoolNull()
	}
	if value := res.Get("enforceTwoFactorAuth"); value.Exists() {
		data.EnforceTwoFactorAuth = types.BoolValue(value.Bool())
	} else {
		data.EnforceTwoFactorAuth = types.BoolNull()
	}
	if value := res.Get("idleTimeoutMinutes"); value.Exists() {
		data.IdleTimeoutMinutes = types.Int64Value(value.Int())
	} else {
		data.IdleTimeoutMinutes = types.Int64Null()
	}
	if value := res.Get("numDifferentPasswords"); value.Exists() {
		data.NumDifferentPasswords = types.Int64Value(value.Int())
	} else {
		data.NumDifferentPasswords = types.Int64Null()
	}
	if value := res.Get("passwordExpirationDays"); value.Exists() {
		data.PasswordExpirationDays = types.Int64Value(value.Int())
	} else {
		data.PasswordExpirationDays = types.Int64Null()
	}
	if value := res.Get("apiAuthentication.ipRestrictionsForKeys.enabled"); value.Exists() {
		data.ApiAuthenticationIpRestrictionsForKeysEnabled = types.BoolValue(value.Bool())
	} else {
		data.ApiAuthenticationIpRestrictionsForKeysEnabled = types.BoolNull()
	}
	if value := res.Get("apiAuthentication.ipRestrictionsForKeys.ranges"); value.Exists() {
		data.ApiAuthenticationIpRestrictionsForKeysRanges = helpers.GetStringSet(value.Array())
	} else {
		data.ApiAuthenticationIpRestrictionsForKeysRanges = types.SetNull(types.StringType)
	}
	if value := res.Get("loginIpRanges"); value.Exists() {
		data.LoginIpRanges = helpers.GetStringSet(value.Array())
	} else {
		data.LoginIpRanges = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationLoginSecurity) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("accountLockoutAttempts"); value.Exists() && !data.AccountLockoutAttempts.IsNull() {
		data.AccountLockoutAttempts = types.Int64Value(value.Int())
	} else {
		data.AccountLockoutAttempts = types.Int64Null()
	}
	if value := res.Get("enforceAccountLockout"); value.Exists() && !data.EnforceAccountLockout.IsNull() {
		data.EnforceAccountLockout = types.BoolValue(value.Bool())
	} else {
		data.EnforceAccountLockout = types.BoolNull()
	}
	if value := res.Get("enforceDifferentPasswords"); value.Exists() && !data.EnforceDifferentPasswords.IsNull() {
		data.EnforceDifferentPasswords = types.BoolValue(value.Bool())
	} else {
		data.EnforceDifferentPasswords = types.BoolNull()
	}
	if value := res.Get("enforceIdleTimeout"); value.Exists() && !data.EnforceIdleTimeout.IsNull() {
		data.EnforceIdleTimeout = types.BoolValue(value.Bool())
	} else {
		data.EnforceIdleTimeout = types.BoolNull()
	}
	if value := res.Get("enforceLoginIpRanges"); value.Exists() && !data.EnforceLoginIpRanges.IsNull() {
		data.EnforceLoginIpRanges = types.BoolValue(value.Bool())
	} else {
		data.EnforceLoginIpRanges = types.BoolNull()
	}
	if value := res.Get("enforcePasswordExpiration"); value.Exists() && !data.EnforcePasswordExpiration.IsNull() {
		data.EnforcePasswordExpiration = types.BoolValue(value.Bool())
	} else {
		data.EnforcePasswordExpiration = types.BoolNull()
	}
	if value := res.Get("enforceStrongPasswords"); value.Exists() && !data.EnforceStrongPasswords.IsNull() {
		data.EnforceStrongPasswords = types.BoolValue(value.Bool())
	} else {
		data.EnforceStrongPasswords = types.BoolNull()
	}
	if value := res.Get("enforceTwoFactorAuth"); value.Exists() && !data.EnforceTwoFactorAuth.IsNull() {
		data.EnforceTwoFactorAuth = types.BoolValue(value.Bool())
	} else {
		data.EnforceTwoFactorAuth = types.BoolNull()
	}
	if value := res.Get("idleTimeoutMinutes"); value.Exists() && !data.IdleTimeoutMinutes.IsNull() {
		data.IdleTimeoutMinutes = types.Int64Value(value.Int())
	} else {
		data.IdleTimeoutMinutes = types.Int64Null()
	}
	if value := res.Get("numDifferentPasswords"); value.Exists() && !data.NumDifferentPasswords.IsNull() {
		data.NumDifferentPasswords = types.Int64Value(value.Int())
	} else {
		data.NumDifferentPasswords = types.Int64Null()
	}
	if value := res.Get("passwordExpirationDays"); value.Exists() && !data.PasswordExpirationDays.IsNull() {
		data.PasswordExpirationDays = types.Int64Value(value.Int())
	} else {
		data.PasswordExpirationDays = types.Int64Null()
	}
	if value := res.Get("apiAuthentication.ipRestrictionsForKeys.enabled"); value.Exists() && !data.ApiAuthenticationIpRestrictionsForKeysEnabled.IsNull() {
		data.ApiAuthenticationIpRestrictionsForKeysEnabled = types.BoolValue(value.Bool())
	} else {
		data.ApiAuthenticationIpRestrictionsForKeysEnabled = types.BoolNull()
	}
	if value := res.Get("apiAuthentication.ipRestrictionsForKeys.ranges"); value.Exists() && !data.ApiAuthenticationIpRestrictionsForKeysRanges.IsNull() {
		data.ApiAuthenticationIpRestrictionsForKeysRanges = helpers.GetStringSet(value.Array())
	} else {
		data.ApiAuthenticationIpRestrictionsForKeysRanges = types.SetNull(types.StringType)
	}
	if value := res.Get("loginIpRanges"); value.Exists() && !data.LoginIpRanges.IsNull() {
		data.LoginIpRanges = helpers.GetStringSet(value.Array())
	} else {
		data.LoginIpRanges = types.SetNull(types.StringType)
	}
}

// End of section. //template:end fromBodyPartial
