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

type WirelessSSIDSplashSettings struct {
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

func (data WirelessSSIDSplashSettings) getPath() string {
	return fmt.Sprintf("/networks/%v/wireless/ssids/%v/splash/settings", url.QueryEscape(data.NetworkId.ValueString()), url.QueryEscape(data.Number.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data WirelessSSIDSplashSettings) toBody(ctx context.Context, state WirelessSSIDSplashSettings) string {
	body := ""
	if !data.AllowSimultaneousLogins.IsNull() {
		body, _ = sjson.Set(body, "allowSimultaneousLogins", data.AllowSimultaneousLogins.ValueBool())
	}
	if !data.BlockAllTrafficBeforeSignOn.IsNull() {
		body, _ = sjson.Set(body, "blockAllTrafficBeforeSignOn", data.BlockAllTrafficBeforeSignOn.ValueBool())
	}
	if !data.ControllerDisconnectionBehavior.IsNull() {
		body, _ = sjson.Set(body, "controllerDisconnectionBehavior", data.ControllerDisconnectionBehavior.ValueString())
	}
	if !data.RedirectUrl.IsNull() {
		body, _ = sjson.Set(body, "redirectUrl", data.RedirectUrl.ValueString())
	}
	if !data.SplashTimeout.IsNull() {
		body, _ = sjson.Set(body, "splashTimeout", data.SplashTimeout.ValueInt64())
	}
	if !data.SplashUrl.IsNull() {
		body, _ = sjson.Set(body, "splashUrl", data.SplashUrl.ValueString())
	}
	if !data.ThemeId.IsNull() {
		body, _ = sjson.Set(body, "themeId", data.ThemeId.ValueString())
	}
	if !data.UseRedirectUrl.IsNull() {
		body, _ = sjson.Set(body, "useRedirectUrl", data.UseRedirectUrl.ValueBool())
	}
	if !data.UseSplashUrl.IsNull() {
		body, _ = sjson.Set(body, "useSplashUrl", data.UseSplashUrl.ValueBool())
	}
	if !data.WelcomeMessage.IsNull() {
		body, _ = sjson.Set(body, "welcomeMessage", data.WelcomeMessage.ValueString())
	}
	if !data.BillingPrepaidAccessFastLoginEnabled.IsNull() {
		body, _ = sjson.Set(body, "billing.prepaidAccessFastLoginEnabled", data.BillingPrepaidAccessFastLoginEnabled.ValueBool())
	}
	if !data.BillingReplyToEmailAddress.IsNull() {
		body, _ = sjson.Set(body, "billing.replyToEmailAddress", data.BillingReplyToEmailAddress.ValueString())
	}
	if !data.BillingFreeAccessDurationInMinutes.IsNull() {
		body, _ = sjson.Set(body, "billing.freeAccess.durationInMinutes", data.BillingFreeAccessDurationInMinutes.ValueInt64())
	}
	if !data.BillingFreeAccessEnabled.IsNull() {
		body, _ = sjson.Set(body, "billing.freeAccess.enabled", data.BillingFreeAccessEnabled.ValueBool())
	}
	if !data.GuestSponsorshipDurationInMinutes.IsNull() {
		body, _ = sjson.Set(body, "guestSponsorship.durationInMinutes", data.GuestSponsorshipDurationInMinutes.ValueInt64())
	}
	if !data.GuestSponsorshipGuestCanRequestTimeframe.IsNull() {
		body, _ = sjson.Set(body, "guestSponsorship.guestCanRequestTimeframe", data.GuestSponsorshipGuestCanRequestTimeframe.ValueBool())
	}
	if !data.SelfRegistrationAuthorizationType.IsNull() {
		body, _ = sjson.Set(body, "selfRegistration.authorizationType", data.SelfRegistrationAuthorizationType.ValueString())
	}
	if !data.SelfRegistrationEnabled.IsNull() {
		body, _ = sjson.Set(body, "selfRegistration.enabled", data.SelfRegistrationEnabled.ValueBool())
	}
	if !data.SentryEnrollmentStrength.IsNull() {
		body, _ = sjson.Set(body, "sentryEnrollment.strength", data.SentryEnrollmentStrength.ValueString())
	}
	if !data.SentryEnrollmentSystemsManagerNetworkId.IsNull() {
		body, _ = sjson.Set(body, "sentryEnrollment.systemsManagerNetwork.id", data.SentryEnrollmentSystemsManagerNetworkId.ValueString())
	}
	if !data.SentryEnrollmentEnforcedSystems.IsNull() {
		var values []string
		data.SentryEnrollmentEnforcedSystems.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "sentryEnrollment.enforcedSystems", values)
	}
	if !data.SplashImageExtension.IsNull() {
		body, _ = sjson.Set(body, "splashImage.extension", data.SplashImageExtension.ValueString())
	}
	if !data.SplashImageMd5.IsNull() {
		body, _ = sjson.Set(body, "splashImage.md5", data.SplashImageMd5.ValueString())
	}
	if !data.SplashImageImageContents.IsNull() {
		body, _ = sjson.Set(body, "splashImage.image.contents", data.SplashImageImageContents.ValueString())
	}
	if !data.SplashImageImageFormat.IsNull() {
		body, _ = sjson.Set(body, "splashImage.image.format", data.SplashImageImageFormat.ValueString())
	}
	if !data.SplashLogoExtension.IsNull() {
		body, _ = sjson.Set(body, "splashLogo.extension", data.SplashLogoExtension.ValueString())
	}
	if !data.SplashLogoMd5.IsNull() {
		body, _ = sjson.Set(body, "splashLogo.md5", data.SplashLogoMd5.ValueString())
	}
	if !data.SplashLogoImageContents.IsNull() {
		body, _ = sjson.Set(body, "splashLogo.image.contents", data.SplashLogoImageContents.ValueString())
	}
	if !data.SplashLogoImageFormat.IsNull() {
		body, _ = sjson.Set(body, "splashLogo.image.format", data.SplashLogoImageFormat.ValueString())
	}
	if !data.SplashPrepaidFrontExtension.IsNull() {
		body, _ = sjson.Set(body, "splashPrepaidFront.extension", data.SplashPrepaidFrontExtension.ValueString())
	}
	if !data.SplashPrepaidFrontMd5.IsNull() {
		body, _ = sjson.Set(body, "splashPrepaidFront.md5", data.SplashPrepaidFrontMd5.ValueString())
	}
	if !data.SplashPrepaidFrontImageContents.IsNull() {
		body, _ = sjson.Set(body, "splashPrepaidFront.image.contents", data.SplashPrepaidFrontImageContents.ValueString())
	}
	if !data.SplashPrepaidFrontImageFormat.IsNull() {
		body, _ = sjson.Set(body, "splashPrepaidFront.image.format", data.SplashPrepaidFrontImageFormat.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *WirelessSSIDSplashSettings) fromBody(ctx context.Context, res meraki.Res) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *WirelessSSIDSplashSettings) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("allowSimultaneousLogins"); value.Exists() && !data.AllowSimultaneousLogins.IsNull() {
		data.AllowSimultaneousLogins = types.BoolValue(value.Bool())
	} else {
		data.AllowSimultaneousLogins = types.BoolNull()
	}
	if value := res.Get("blockAllTrafficBeforeSignOn"); value.Exists() && !data.BlockAllTrafficBeforeSignOn.IsNull() {
		data.BlockAllTrafficBeforeSignOn = types.BoolValue(value.Bool())
	} else {
		data.BlockAllTrafficBeforeSignOn = types.BoolNull()
	}
	if value := res.Get("controllerDisconnectionBehavior"); value.Exists() && !data.ControllerDisconnectionBehavior.IsNull() {
		data.ControllerDisconnectionBehavior = types.StringValue(value.String())
	} else {
		data.ControllerDisconnectionBehavior = types.StringNull()
	}
	if value := res.Get("redirectUrl"); value.Exists() && !data.RedirectUrl.IsNull() {
		data.RedirectUrl = types.StringValue(value.String())
	} else {
		data.RedirectUrl = types.StringNull()
	}
	if value := res.Get("splashTimeout"); value.Exists() && !data.SplashTimeout.IsNull() {
		data.SplashTimeout = types.Int64Value(value.Int())
	} else {
		data.SplashTimeout = types.Int64Null()
	}
	if value := res.Get("splashUrl"); value.Exists() && !data.SplashUrl.IsNull() {
		data.SplashUrl = types.StringValue(value.String())
	} else {
		data.SplashUrl = types.StringNull()
	}
	if value := res.Get("themeId"); value.Exists() && !data.ThemeId.IsNull() {
		data.ThemeId = types.StringValue(value.String())
	} else {
		data.ThemeId = types.StringNull()
	}
	if value := res.Get("useRedirectUrl"); value.Exists() && !data.UseRedirectUrl.IsNull() {
		data.UseRedirectUrl = types.BoolValue(value.Bool())
	} else {
		data.UseRedirectUrl = types.BoolNull()
	}
	if value := res.Get("useSplashUrl"); value.Exists() && !data.UseSplashUrl.IsNull() {
		data.UseSplashUrl = types.BoolValue(value.Bool())
	} else {
		data.UseSplashUrl = types.BoolNull()
	}
	if value := res.Get("welcomeMessage"); value.Exists() && !data.WelcomeMessage.IsNull() {
		data.WelcomeMessage = types.StringValue(value.String())
	} else {
		data.WelcomeMessage = types.StringNull()
	}
	if value := res.Get("billing.prepaidAccessFastLoginEnabled"); value.Exists() && !data.BillingPrepaidAccessFastLoginEnabled.IsNull() {
		data.BillingPrepaidAccessFastLoginEnabled = types.BoolValue(value.Bool())
	} else {
		data.BillingPrepaidAccessFastLoginEnabled = types.BoolNull()
	}
	if value := res.Get("billing.replyToEmailAddress"); value.Exists() && !data.BillingReplyToEmailAddress.IsNull() {
		data.BillingReplyToEmailAddress = types.StringValue(value.String())
	} else {
		data.BillingReplyToEmailAddress = types.StringNull()
	}
	if value := res.Get("billing.freeAccess.durationInMinutes"); value.Exists() && !data.BillingFreeAccessDurationInMinutes.IsNull() {
		data.BillingFreeAccessDurationInMinutes = types.Int64Value(value.Int())
	} else {
		data.BillingFreeAccessDurationInMinutes = types.Int64Null()
	}
	if value := res.Get("billing.freeAccess.enabled"); value.Exists() && !data.BillingFreeAccessEnabled.IsNull() {
		data.BillingFreeAccessEnabled = types.BoolValue(value.Bool())
	} else {
		data.BillingFreeAccessEnabled = types.BoolNull()
	}
	if value := res.Get("guestSponsorship.durationInMinutes"); value.Exists() && !data.GuestSponsorshipDurationInMinutes.IsNull() {
		data.GuestSponsorshipDurationInMinutes = types.Int64Value(value.Int())
	} else {
		data.GuestSponsorshipDurationInMinutes = types.Int64Null()
	}
	if value := res.Get("guestSponsorship.guestCanRequestTimeframe"); value.Exists() && !data.GuestSponsorshipGuestCanRequestTimeframe.IsNull() {
		data.GuestSponsorshipGuestCanRequestTimeframe = types.BoolValue(value.Bool())
	} else {
		data.GuestSponsorshipGuestCanRequestTimeframe = types.BoolNull()
	}
	if value := res.Get("selfRegistration.authorizationType"); value.Exists() && !data.SelfRegistrationAuthorizationType.IsNull() {
		data.SelfRegistrationAuthorizationType = types.StringValue(value.String())
	} else {
		data.SelfRegistrationAuthorizationType = types.StringNull()
	}
	if value := res.Get("selfRegistration.enabled"); value.Exists() && !data.SelfRegistrationEnabled.IsNull() {
		data.SelfRegistrationEnabled = types.BoolValue(value.Bool())
	} else {
		data.SelfRegistrationEnabled = types.BoolNull()
	}
	if value := res.Get("sentryEnrollment.strength"); value.Exists() && !data.SentryEnrollmentStrength.IsNull() {
		data.SentryEnrollmentStrength = types.StringValue(value.String())
	} else {
		data.SentryEnrollmentStrength = types.StringNull()
	}
	if value := res.Get("sentryEnrollment.systemsManagerNetwork.id"); value.Exists() && !data.SentryEnrollmentSystemsManagerNetworkId.IsNull() {
		data.SentryEnrollmentSystemsManagerNetworkId = types.StringValue(value.String())
	} else {
		data.SentryEnrollmentSystemsManagerNetworkId = types.StringNull()
	}
	if value := res.Get("sentryEnrollment.enforcedSystems"); value.Exists() && !data.SentryEnrollmentEnforcedSystems.IsNull() {
		data.SentryEnrollmentEnforcedSystems = helpers.GetStringSet(value.Array())
	} else {
		data.SentryEnrollmentEnforcedSystems = types.SetNull(types.StringType)
	}
	if value := res.Get("splashImage.extension"); value.Exists() && !data.SplashImageExtension.IsNull() {
		data.SplashImageExtension = types.StringValue(value.String())
	} else {
		data.SplashImageExtension = types.StringNull()
	}
	if value := res.Get("splashImage.md5"); value.Exists() && !data.SplashImageMd5.IsNull() {
		data.SplashImageMd5 = types.StringValue(value.String())
	} else {
		data.SplashImageMd5 = types.StringNull()
	}
	if value := res.Get("splashImage.image.contents"); value.Exists() && !data.SplashImageImageContents.IsNull() {
		data.SplashImageImageContents = types.StringValue(value.String())
	} else {
		data.SplashImageImageContents = types.StringNull()
	}
	if value := res.Get("splashImage.image.format"); value.Exists() && !data.SplashImageImageFormat.IsNull() {
		data.SplashImageImageFormat = types.StringValue(value.String())
	} else {
		data.SplashImageImageFormat = types.StringNull()
	}
	if value := res.Get("splashLogo.extension"); value.Exists() && !data.SplashLogoExtension.IsNull() {
		data.SplashLogoExtension = types.StringValue(value.String())
	} else {
		data.SplashLogoExtension = types.StringNull()
	}
	if value := res.Get("splashLogo.md5"); value.Exists() && !data.SplashLogoMd5.IsNull() {
		data.SplashLogoMd5 = types.StringValue(value.String())
	} else {
		data.SplashLogoMd5 = types.StringNull()
	}
	if value := res.Get("splashLogo.image.contents"); value.Exists() && !data.SplashLogoImageContents.IsNull() {
		data.SplashLogoImageContents = types.StringValue(value.String())
	} else {
		data.SplashLogoImageContents = types.StringNull()
	}
	if value := res.Get("splashLogo.image.format"); value.Exists() && !data.SplashLogoImageFormat.IsNull() {
		data.SplashLogoImageFormat = types.StringValue(value.String())
	} else {
		data.SplashLogoImageFormat = types.StringNull()
	}
	if value := res.Get("splashPrepaidFront.extension"); value.Exists() && !data.SplashPrepaidFrontExtension.IsNull() {
		data.SplashPrepaidFrontExtension = types.StringValue(value.String())
	} else {
		data.SplashPrepaidFrontExtension = types.StringNull()
	}
	if value := res.Get("splashPrepaidFront.md5"); value.Exists() && !data.SplashPrepaidFrontMd5.IsNull() {
		data.SplashPrepaidFrontMd5 = types.StringValue(value.String())
	} else {
		data.SplashPrepaidFrontMd5 = types.StringNull()
	}
	if value := res.Get("splashPrepaidFront.image.contents"); value.Exists() && !data.SplashPrepaidFrontImageContents.IsNull() {
		data.SplashPrepaidFrontImageContents = types.StringValue(value.String())
	} else {
		data.SplashPrepaidFrontImageContents = types.StringNull()
	}
	if value := res.Get("splashPrepaidFront.image.format"); value.Exists() && !data.SplashPrepaidFrontImageFormat.IsNull() {
		data.SplashPrepaidFrontImageFormat = types.StringValue(value.String())
	} else {
		data.SplashPrepaidFrontImageFormat = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *WirelessSSIDSplashSettings) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data WirelessSSIDSplashSettings) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
