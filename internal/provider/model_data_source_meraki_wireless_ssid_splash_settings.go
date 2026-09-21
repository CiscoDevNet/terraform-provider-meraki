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
// It emits a separate model - DataSourceWirelessSSIDSplashSettings - used only by the data source, always including every
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

	"github.com/CiscoDevNet/terraform-provider-meraki/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DataSourceWirelessSSIDSplashSettings struct {
	Id                                       types.String `tfsdk:"id"`
	NetworkId                                types.String `tfsdk:"network_id"`
	Number                                   types.String `tfsdk:"number"`
	AllowSimultaneousLogins                  types.Bool   `tfsdk:"allow_simultaneous_logins"`
	BlockAllTrafficBeforeSignOn              types.Bool   `tfsdk:"block_all_traffic_before_sign_on"`
	ControllerDisconnectionBehavior          types.String `tfsdk:"controller_disconnection_behavior"`
	RedirectUrl                              types.String `tfsdk:"redirect_url"`
	SplashTimeout                            types.Int64  `tfsdk:"splash_timeout"`
	SplashUrl                                types.String `tfsdk:"splash_url"`
	ThemeId                                  types.String `tfsdk:"theme_id"`
	UseRedirectUrl                           types.Bool   `tfsdk:"use_redirect_url"`
	UseSplashUrl                             types.Bool   `tfsdk:"use_splash_url"`
	WelcomeMessage                           types.String `tfsdk:"welcome_message"`
	BillingPrepaidAccessFastLoginEnabled     types.Bool   `tfsdk:"billing_prepaid_access_fast_login_enabled"`
	BillingReplyToEmailAddress               types.String `tfsdk:"billing_reply_to_email_address"`
	BillingFreeAccessDurationInMinutes       types.Int64  `tfsdk:"billing_free_access_duration_in_minutes"`
	BillingFreeAccessEnabled                 types.Bool   `tfsdk:"billing_free_access_enabled"`
	GuestSponsorshipDurationInMinutes        types.Int64  `tfsdk:"guest_sponsorship_duration_in_minutes"`
	GuestSponsorshipGuestCanRequestTimeframe types.Bool   `tfsdk:"guest_sponsorship_guest_can_request_timeframe"`
	SelfRegistrationAuthorizationType        types.String `tfsdk:"self_registration_authorization_type"`
	SelfRegistrationEnabled                  types.Bool   `tfsdk:"self_registration_enabled"`
	SentryEnrollmentStrength                 types.String `tfsdk:"sentry_enrollment_strength"`
	SentryEnrollmentSystemsManagerNetworkId  types.String `tfsdk:"sentry_enrollment_systems_manager_network_id"`
	SentryEnrollmentEnforcedSystems          types.Set    `tfsdk:"sentry_enrollment_enforced_systems"`
	SplashImageExtension                     types.String `tfsdk:"splash_image_extension"`
	SplashImageMd5                           types.String `tfsdk:"splash_image_md5"`
	SplashImageImageContents                 types.String `tfsdk:"splash_image_image_contents"`
	SplashImageImageFormat                   types.String `tfsdk:"splash_image_image_format"`
	SplashLogoExtension                      types.String `tfsdk:"splash_logo_extension"`
	SplashLogoMd5                            types.String `tfsdk:"splash_logo_md5"`
	SplashLogoImageContents                  types.String `tfsdk:"splash_logo_image_contents"`
	SplashLogoImageFormat                    types.String `tfsdk:"splash_logo_image_format"`
	SplashPrepaidFrontExtension              types.String `tfsdk:"splash_prepaid_front_extension"`
	SplashPrepaidFrontMd5                    types.String `tfsdk:"splash_prepaid_front_md5"`
	SplashPrepaidFrontImageContents          types.String `tfsdk:"splash_prepaid_front_image_contents"`
	SplashPrepaidFrontImageFormat            types.String `tfsdk:"splash_prepaid_front_image_format"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceWirelessSSIDSplashSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/splash/settings", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceWirelessSSIDSplashSettings) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("allowSimultaneousLogins"); value.Exists() && value.Value() != nil {
		data.AllowSimultaneousLogins = types.BoolValue(value.Bool())
	} else {
		data.AllowSimultaneousLogins = types.BoolNull()
	}
	if value := res.Get("blockAllTrafficBeforeSignOn"); value.Exists() && value.Value() != nil {
		data.BlockAllTrafficBeforeSignOn = types.BoolValue(value.Bool())
	} else {
		data.BlockAllTrafficBeforeSignOn = types.BoolNull()
	}
	if value := res.Get("controllerDisconnectionBehavior"); value.Exists() && value.Value() != nil {
		data.ControllerDisconnectionBehavior = types.StringValue(value.String())
	} else {
		data.ControllerDisconnectionBehavior = types.StringNull()
	}
	if value := res.Get("redirectUrl"); value.Exists() && value.Value() != nil {
		data.RedirectUrl = types.StringValue(value.String())
	} else {
		data.RedirectUrl = types.StringNull()
	}
	if value := res.Get("splashTimeout"); value.Exists() && value.Value() != nil {
		data.SplashTimeout = types.Int64Value(value.Int())
	} else {
		data.SplashTimeout = types.Int64Null()
	}
	if value := res.Get("splashUrl"); value.Exists() && value.Value() != nil {
		data.SplashUrl = types.StringValue(value.String())
	} else {
		data.SplashUrl = types.StringNull()
	}
	if value := res.Get("themeId"); value.Exists() && value.Value() != nil {
		data.ThemeId = types.StringValue(value.String())
	} else {
		data.ThemeId = types.StringNull()
	}
	if value := res.Get("useRedirectUrl"); value.Exists() && value.Value() != nil {
		data.UseRedirectUrl = types.BoolValue(value.Bool())
	} else {
		data.UseRedirectUrl = types.BoolNull()
	}
	if value := res.Get("useSplashUrl"); value.Exists() && value.Value() != nil {
		data.UseSplashUrl = types.BoolValue(value.Bool())
	} else {
		data.UseSplashUrl = types.BoolNull()
	}
	if value := res.Get("welcomeMessage"); value.Exists() && value.Value() != nil {
		data.WelcomeMessage = types.StringValue(value.String())
	} else {
		data.WelcomeMessage = types.StringNull()
	}
	if value := res.Get("billing.prepaidAccessFastLoginEnabled"); value.Exists() && value.Value() != nil {
		data.BillingPrepaidAccessFastLoginEnabled = types.BoolValue(value.Bool())
	} else {
		data.BillingPrepaidAccessFastLoginEnabled = types.BoolNull()
	}
	if value := res.Get("billing.replyToEmailAddress"); value.Exists() && value.Value() != nil {
		data.BillingReplyToEmailAddress = types.StringValue(value.String())
	} else {
		data.BillingReplyToEmailAddress = types.StringNull()
	}
	if value := res.Get("billing.freeAccess.durationInMinutes"); value.Exists() && value.Value() != nil {
		data.BillingFreeAccessDurationInMinutes = types.Int64Value(value.Int())
	} else {
		data.BillingFreeAccessDurationInMinutes = types.Int64Null()
	}
	if value := res.Get("billing.freeAccess.enabled"); value.Exists() && value.Value() != nil {
		data.BillingFreeAccessEnabled = types.BoolValue(value.Bool())
	} else {
		data.BillingFreeAccessEnabled = types.BoolNull()
	}
	if value := res.Get("guestSponsorship.durationInMinutes"); value.Exists() && value.Value() != nil {
		data.GuestSponsorshipDurationInMinutes = types.Int64Value(value.Int())
	} else {
		data.GuestSponsorshipDurationInMinutes = types.Int64Null()
	}
	if value := res.Get("guestSponsorship.guestCanRequestTimeframe"); value.Exists() && value.Value() != nil {
		data.GuestSponsorshipGuestCanRequestTimeframe = types.BoolValue(value.Bool())
	} else {
		data.GuestSponsorshipGuestCanRequestTimeframe = types.BoolNull()
	}
	if value := res.Get("selfRegistration.authorizationType"); value.Exists() && value.Value() != nil {
		data.SelfRegistrationAuthorizationType = types.StringValue(value.String())
	} else {
		data.SelfRegistrationAuthorizationType = types.StringNull()
	}
	if value := res.Get("selfRegistration.enabled"); value.Exists() && value.Value() != nil {
		data.SelfRegistrationEnabled = types.BoolValue(value.Bool())
	} else {
		data.SelfRegistrationEnabled = types.BoolNull()
	}
	if value := res.Get("sentryEnrollment.strength"); value.Exists() && value.Value() != nil {
		data.SentryEnrollmentStrength = types.StringValue(value.String())
	} else {
		data.SentryEnrollmentStrength = types.StringNull()
	}
	if value := res.Get("sentryEnrollment.systemsManagerNetwork.id"); value.Exists() && value.Value() != nil {
		data.SentryEnrollmentSystemsManagerNetworkId = types.StringValue(value.String())
	} else {
		data.SentryEnrollmentSystemsManagerNetworkId = types.StringNull()
	}
	if value := res.Get("sentryEnrollment.enforcedSystems"); value.Exists() && value.Value() != nil {
		data.SentryEnrollmentEnforcedSystems = helpers.GetStringSet(value.Array())
	} else {
		data.SentryEnrollmentEnforcedSystems = types.SetNull(types.StringType)
	}
	if value := res.Get("splashImage.extension"); value.Exists() && value.Value() != nil {
		data.SplashImageExtension = types.StringValue(value.String())
	} else {
		data.SplashImageExtension = types.StringNull()
	}
	if value := res.Get("splashImage.md5"); value.Exists() && value.Value() != nil {
		data.SplashImageMd5 = types.StringValue(value.String())
	} else {
		data.SplashImageMd5 = types.StringNull()
	}
	if value := res.Get("splashImage.image.contents"); value.Exists() && value.Value() != nil {
		data.SplashImageImageContents = types.StringValue(value.String())
	} else {
		data.SplashImageImageContents = types.StringNull()
	}
	if value := res.Get("splashImage.image.format"); value.Exists() && value.Value() != nil {
		data.SplashImageImageFormat = types.StringValue(value.String())
	} else {
		data.SplashImageImageFormat = types.StringNull()
	}
	if value := res.Get("splashLogo.extension"); value.Exists() && value.Value() != nil {
		data.SplashLogoExtension = types.StringValue(value.String())
	} else {
		data.SplashLogoExtension = types.StringNull()
	}
	if value := res.Get("splashLogo.md5"); value.Exists() && value.Value() != nil {
		data.SplashLogoMd5 = types.StringValue(value.String())
	} else {
		data.SplashLogoMd5 = types.StringNull()
	}
	if value := res.Get("splashLogo.image.contents"); value.Exists() && value.Value() != nil {
		data.SplashLogoImageContents = types.StringValue(value.String())
	} else {
		data.SplashLogoImageContents = types.StringNull()
	}
	if value := res.Get("splashLogo.image.format"); value.Exists() && value.Value() != nil {
		data.SplashLogoImageFormat = types.StringValue(value.String())
	} else {
		data.SplashLogoImageFormat = types.StringNull()
	}
	if value := res.Get("splashPrepaidFront.extension"); value.Exists() && value.Value() != nil {
		data.SplashPrepaidFrontExtension = types.StringValue(value.String())
	} else {
		data.SplashPrepaidFrontExtension = types.StringNull()
	}
	if value := res.Get("splashPrepaidFront.md5"); value.Exists() && value.Value() != nil {
		data.SplashPrepaidFrontMd5 = types.StringValue(value.String())
	} else {
		data.SplashPrepaidFrontMd5 = types.StringNull()
	}
	if value := res.Get("splashPrepaidFront.image.contents"); value.Exists() && value.Value() != nil {
		data.SplashPrepaidFrontImageContents = types.StringValue(value.String())
	} else {
		data.SplashPrepaidFrontImageContents = types.StringNull()
	}
	if value := res.Get("splashPrepaidFront.image.format"); value.Exists() && value.Value() != nil {
		data.SplashPrepaidFrontImageFormat = types.StringValue(value.String())
	} else {
		data.SplashPrepaidFrontImageFormat = types.StringNull()
	}
}

// End of section. //template:end fromBody
