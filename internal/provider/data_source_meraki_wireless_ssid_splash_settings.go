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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-meraki"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessSSIDSplashSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessSSIDSplashSettingsDataSource{}
)

func NewWirelessSSIDSplashSettingsDataSource() datasource.DataSource {
	return &WirelessSSIDSplashSettingsDataSource{}
}

type WirelessSSIDSplashSettingsDataSource struct {
	client *meraki.Client
}

func (d *WirelessSSIDSplashSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_ssid_splash_settings"
}

func (d *WirelessSSIDSplashSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the `Wireless SSID Splash Settings` configuration.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: "Wireless SSID number",
				Required:            true,
			},
			"allow_simultaneous_logins": schema.BoolAttribute{
				MarkdownDescription: "Whether or not to allow simultaneous logins from different devices.",
				Computed:            true,
			},
			"block_all_traffic_before_sign_on": schema.BoolAttribute{
				MarkdownDescription: "How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.",
				Computed:            true,
			},
			"controller_disconnection_behavior": schema.StringAttribute{
				MarkdownDescription: "How login attempts should be handled when the controller is unreachable. Can be either `open`, `restricted`, or `default`.",
				Computed:            true,
			},
			"redirect_url": schema.StringAttribute{
				MarkdownDescription: "The custom redirect URL where the users will go after the splash page.",
				Computed:            true,
			},
			"splash_timeout": schema.Int64Attribute{
				MarkdownDescription: "Splash timeout in minutes. This will determine how often users will see the splash page.",
				Computed:            true,
			},
			"splash_url": schema.StringAttribute{
				MarkdownDescription: "[optional] The custom splash URL of the click-through splash page. Note that the URL can be configured without necessarily being used. In order to enable the custom URL, see `useSplashUrl`",
				Computed:            true,
			},
			"theme_id": schema.StringAttribute{
				MarkdownDescription: "The id of the selected splash theme.",
				Computed:            true,
			},
			"use_redirect_url": schema.BoolAttribute{
				MarkdownDescription: "The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page. A custom redirect URL must be set if this is true.",
				Computed:            true,
			},
			"use_splash_url": schema.BoolAttribute{
				MarkdownDescription: "[optional] Boolean indicating whether the users will be redirected to the custom splash url. A custom splash URL must be set if this is true. Note that depending on your SSID`s access control settings, it may not be possible to use the custom splash URL.",
				Computed:            true,
			},
			"welcome_message": schema.StringAttribute{
				MarkdownDescription: "The welcome message for the users on the splash page.",
				Computed:            true,
			},
			"billing_prepaid_access_fast_login_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not billing uses the fast login prepaid access option.",
				Computed:            true,
			},
			"billing_reply_to_email_address": schema.StringAttribute{
				MarkdownDescription: "The email address that receives replies from clients.",
				Computed:            true,
			},
			"billing_free_access_duration_in_minutes": schema.Int64Attribute{
				MarkdownDescription: "How long a device can use a network for free.",
				Computed:            true,
			},
			"billing_free_access_enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether or not free access is enabled.",
				Computed:            true,
			},
			"guest_sponsorship_duration_in_minutes": schema.Int64Attribute{
				MarkdownDescription: "Duration in minutes of sponsored guest authorization. Must be between 1 and 60480 (6 weeks)",
				Computed:            true,
			},
			"guest_sponsorship_guest_can_request_timeframe": schema.BoolAttribute{
				MarkdownDescription: "Whether or not guests can specify how much time they are requesting.",
				Computed:            true,
			},
			"sentry_enrollment_strength": schema.StringAttribute{
				MarkdownDescription: "The strength of the enforcement of selected system types. Must be one of: `focused`, `click-through`, and `strict`.",
				Computed:            true,
			},
			"sentry_enrollment_systems_manager_network_id": schema.StringAttribute{
				MarkdownDescription: "The network ID of the Systems Manager network.",
				Computed:            true,
			},
			"sentry_enrollment_enforced_systems": schema.SetAttribute{
				MarkdownDescription: "The system types that the Sentry enforces. Must be included in: `iOS, `Android`, `macOS`, and `Windows`.",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"splash_image_extension": schema.StringAttribute{
				MarkdownDescription: "The extension of the image file.",
				Computed:            true,
			},
			"splash_image_md5": schema.StringAttribute{
				MarkdownDescription: "The MD5 value of the image file. Setting this to null will remove the image from the splash page.",
				Computed:            true,
			},
			"splash_image_image_contents": schema.StringAttribute{
				MarkdownDescription: "The file contents (a base 64 encoded string) of your new image.",
				Computed:            true,
			},
			"splash_image_image_format": schema.StringAttribute{
				MarkdownDescription: "The format of the encoded contents. Supported formats are `png`, `gif`, and jpg`.",
				Computed:            true,
			},
			"splash_logo_extension": schema.StringAttribute{
				MarkdownDescription: "The extension of the logo file.",
				Computed:            true,
			},
			"splash_logo_md5": schema.StringAttribute{
				MarkdownDescription: "The MD5 value of the logo file. Setting this to null will remove the logo from the splash page.",
				Computed:            true,
			},
			"splash_logo_image_contents": schema.StringAttribute{
				MarkdownDescription: "The file contents (a base 64 encoded string) of your new logo.",
				Computed:            true,
			},
			"splash_logo_image_format": schema.StringAttribute{
				MarkdownDescription: "The format of the encoded contents. Supported formats are `png`, `gif`, and jpg`.",
				Computed:            true,
			},
			"splash_prepaid_front_extension": schema.StringAttribute{
				MarkdownDescription: "The extension of the prepaid front image file.",
				Computed:            true,
			},
			"splash_prepaid_front_md5": schema.StringAttribute{
				MarkdownDescription: "The MD5 value of the prepaid front image file. Setting this to null will remove the prepaid front from the splash page.",
				Computed:            true,
			},
			"splash_prepaid_front_image_contents": schema.StringAttribute{
				MarkdownDescription: "The file contents (a base 64 encoded string) of your new prepaid front.",
				Computed:            true,
			},
			"splash_prepaid_front_image_format": schema.StringAttribute{
				MarkdownDescription: "The format of the encoded contents. Supported formats are `png`, `gif`, and jpg`.",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessSSIDSplashSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*MerakiProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read

func (d *WirelessSSIDSplashSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessSSIDSplashSettings

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	var res meraki.Res
	var err error

	if !res.Exists() {
		res, err = d.client.Get(config.getPath())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
			return
		}
	}

	config.fromBody(ctx, res)
	config.Id = config.Number

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
