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

type OrganizationBrandingPolicy struct {
	Id                                             types.String `tfsdk:"id"`
	OrganizationId                                 types.String `tfsdk:"organization_id"`
	Enabled                                        types.Bool   `tfsdk:"enabled"`
	Name                                           types.String `tfsdk:"name"`
	AdminSettingsAppliesTo                         types.String `tfsdk:"admin_settings_applies_to"`
	AdminSettingsValues                            types.List   `tfsdk:"admin_settings_values"`
	CustomLogoEnabled                              types.Bool   `tfsdk:"custom_logo_enabled"`
	CustomLogoImageContents                        types.String `tfsdk:"custom_logo_image_contents"`
	CustomLogoImageFormat                          types.String `tfsdk:"custom_logo_image_format"`
	HelpSettingsApiDocsSubtab                      types.String `tfsdk:"help_settings_api_docs_subtab"`
	HelpSettingsCasesSubtab                        types.String `tfsdk:"help_settings_cases_subtab"`
	HelpSettingsCiscoMerakiProductDocumentation    types.String `tfsdk:"help_settings_cisco_meraki_product_documentation"`
	HelpSettingsCommunitySubtab                    types.String `tfsdk:"help_settings_community_subtab"`
	HelpSettingsDataProtectionRequestsSubtab       types.String `tfsdk:"help_settings_data_protection_requests_subtab"`
	HelpSettingsFirewallInfoSubtab                 types.String `tfsdk:"help_settings_firewall_info_subtab"`
	HelpSettingsGetHelpSubtab                      types.String `tfsdk:"help_settings_get_help_subtab"`
	HelpSettingsGetHelpSubtabKnowledgeBaseSearch   types.String `tfsdk:"help_settings_get_help_subtab_knowledge_base_search"`
	HelpSettingsHardwareReplacementsSubtab         types.String `tfsdk:"help_settings_hardware_replacements_subtab"`
	HelpSettingsHelpTab                            types.String `tfsdk:"help_settings_help_tab"`
	HelpSettingsHelpWidget                         types.String `tfsdk:"help_settings_help_widget"`
	HelpSettingsNewFeaturesSubtab                  types.String `tfsdk:"help_settings_new_features_subtab"`
	HelpSettingsSmForums                           types.String `tfsdk:"help_settings_sm_forums"`
	HelpSettingsSupportContactInfo                 types.String `tfsdk:"help_settings_support_contact_info"`
	HelpSettingsUniversalSearchKnowledgeBaseSearch types.String `tfsdk:"help_settings_universal_search_knowledge_base_search"`
}

type OrganizationBrandingPolicyIdentity struct {
	OrganizationId types.String `tfsdk:"organization_id"`
	Id             types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data OrganizationBrandingPolicy) getPath() string {
	return fmt.Sprintf("/organizations/%v/brandingPolicies", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data OrganizationBrandingPolicy) toBody(ctx context.Context, state OrganizationBrandingPolicy) string {
	body := ""
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.AdminSettingsAppliesTo.IsNull() {
		body, _ = sjson.Set(body, "adminSettings.appliesTo", data.AdminSettingsAppliesTo.ValueString())
	}
	if !data.AdminSettingsValues.IsNull() {
		var values []string
		data.AdminSettingsValues.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "adminSettings.values", values)
	}
	if !data.CustomLogoEnabled.IsNull() {
		body, _ = sjson.Set(body, "customLogo.enabled", data.CustomLogoEnabled.ValueBool())
	}
	if !data.CustomLogoImageContents.IsNull() {
		body, _ = sjson.Set(body, "customLogo.image.contents", data.CustomLogoImageContents.ValueString())
	}
	if !data.CustomLogoImageFormat.IsNull() {
		body, _ = sjson.Set(body, "customLogo.image.format", data.CustomLogoImageFormat.ValueString())
	}
	if !data.HelpSettingsApiDocsSubtab.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.apiDocsSubtab", data.HelpSettingsApiDocsSubtab.ValueString())
	}
	if !data.HelpSettingsCasesSubtab.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.casesSubtab", data.HelpSettingsCasesSubtab.ValueString())
	}
	if !data.HelpSettingsCiscoMerakiProductDocumentation.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.ciscoMerakiProductDocumentation", data.HelpSettingsCiscoMerakiProductDocumentation.ValueString())
	}
	if !data.HelpSettingsCommunitySubtab.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.communitySubtab", data.HelpSettingsCommunitySubtab.ValueString())
	}
	if !data.HelpSettingsDataProtectionRequestsSubtab.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.dataProtectionRequestsSubtab", data.HelpSettingsDataProtectionRequestsSubtab.ValueString())
	}
	if !data.HelpSettingsFirewallInfoSubtab.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.firewallInfoSubtab", data.HelpSettingsFirewallInfoSubtab.ValueString())
	}
	if !data.HelpSettingsGetHelpSubtab.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.getHelpSubtab", data.HelpSettingsGetHelpSubtab.ValueString())
	}
	if !data.HelpSettingsGetHelpSubtabKnowledgeBaseSearch.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.getHelpSubtabKnowledgeBaseSearch", data.HelpSettingsGetHelpSubtabKnowledgeBaseSearch.ValueString())
	}
	if !data.HelpSettingsHardwareReplacementsSubtab.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.hardwareReplacementsSubtab", data.HelpSettingsHardwareReplacementsSubtab.ValueString())
	}
	if !data.HelpSettingsHelpTab.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.helpTab", data.HelpSettingsHelpTab.ValueString())
	}
	if !data.HelpSettingsHelpWidget.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.helpWidget", data.HelpSettingsHelpWidget.ValueString())
	}
	if !data.HelpSettingsNewFeaturesSubtab.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.newFeaturesSubtab", data.HelpSettingsNewFeaturesSubtab.ValueString())
	}
	if !data.HelpSettingsSmForums.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.smForums", data.HelpSettingsSmForums.ValueString())
	}
	if !data.HelpSettingsSupportContactInfo.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.supportContactInfo", data.HelpSettingsSupportContactInfo.ValueString())
	}
	if !data.HelpSettingsUniversalSearchKnowledgeBaseSearch.IsNull() {
		body, _ = sjson.Set(body, "helpSettings.universalSearchKnowledgeBaseSearch", data.HelpSettingsUniversalSearchKnowledgeBaseSearch.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *OrganizationBrandingPolicy) fromBody(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && value.Value() != nil {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() && value.Value() != nil {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("adminSettings.appliesTo"); value.Exists() && value.Value() != nil {
		data.AdminSettingsAppliesTo = types.StringValue(value.String())
	} else {
		data.AdminSettingsAppliesTo = types.StringNull()
	}
	if value := res.Get("adminSettings.values"); value.Exists() && value.Value() != nil {
		data.AdminSettingsValues = helpers.GetStringList(value.Array())
	} else {
		data.AdminSettingsValues = types.ListNull(types.StringType)
	}
	if value := res.Get("customLogo.enabled"); value.Exists() && value.Value() != nil {
		data.CustomLogoEnabled = types.BoolValue(value.Bool())
	} else {
		data.CustomLogoEnabled = types.BoolNull()
	}
	if value := res.Get("customLogo.image.contents"); value.Exists() && value.Value() != nil {
		data.CustomLogoImageContents = types.StringValue(value.String())
	} else {
		data.CustomLogoImageContents = types.StringNull()
	}
	if value := res.Get("customLogo.image.format"); value.Exists() && value.Value() != nil {
		data.CustomLogoImageFormat = types.StringValue(value.String())
	} else {
		data.CustomLogoImageFormat = types.StringNull()
	}
	if value := res.Get("helpSettings.apiDocsSubtab"); value.Exists() && value.Value() != nil {
		data.HelpSettingsApiDocsSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsApiDocsSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.casesSubtab"); value.Exists() && value.Value() != nil {
		data.HelpSettingsCasesSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsCasesSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.ciscoMerakiProductDocumentation"); value.Exists() && value.Value() != nil {
		data.HelpSettingsCiscoMerakiProductDocumentation = types.StringValue(value.String())
	} else {
		data.HelpSettingsCiscoMerakiProductDocumentation = types.StringNull()
	}
	if value := res.Get("helpSettings.communitySubtab"); value.Exists() && value.Value() != nil {
		data.HelpSettingsCommunitySubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsCommunitySubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.dataProtectionRequestsSubtab"); value.Exists() && value.Value() != nil {
		data.HelpSettingsDataProtectionRequestsSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsDataProtectionRequestsSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.firewallInfoSubtab"); value.Exists() && value.Value() != nil {
		data.HelpSettingsFirewallInfoSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsFirewallInfoSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.getHelpSubtab"); value.Exists() && value.Value() != nil {
		data.HelpSettingsGetHelpSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsGetHelpSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.getHelpSubtabKnowledgeBaseSearch"); value.Exists() && value.Value() != nil {
		data.HelpSettingsGetHelpSubtabKnowledgeBaseSearch = types.StringValue(value.String())
	} else {
		data.HelpSettingsGetHelpSubtabKnowledgeBaseSearch = types.StringNull()
	}
	if value := res.Get("helpSettings.hardwareReplacementsSubtab"); value.Exists() && value.Value() != nil {
		data.HelpSettingsHardwareReplacementsSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsHardwareReplacementsSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.helpTab"); value.Exists() && value.Value() != nil {
		data.HelpSettingsHelpTab = types.StringValue(value.String())
	} else {
		data.HelpSettingsHelpTab = types.StringNull()
	}
	if value := res.Get("helpSettings.helpWidget"); value.Exists() && value.Value() != nil {
		data.HelpSettingsHelpWidget = types.StringValue(value.String())
	} else {
		data.HelpSettingsHelpWidget = types.StringNull()
	}
	if value := res.Get("helpSettings.newFeaturesSubtab"); value.Exists() && value.Value() != nil {
		data.HelpSettingsNewFeaturesSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsNewFeaturesSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.smForums"); value.Exists() && value.Value() != nil {
		data.HelpSettingsSmForums = types.StringValue(value.String())
	} else {
		data.HelpSettingsSmForums = types.StringNull()
	}
	if value := res.Get("helpSettings.supportContactInfo"); value.Exists() && value.Value() != nil {
		data.HelpSettingsSupportContactInfo = types.StringValue(value.String())
	} else {
		data.HelpSettingsSupportContactInfo = types.StringNull()
	}
	if value := res.Get("helpSettings.universalSearchKnowledgeBaseSearch"); value.Exists() && value.Value() != nil {
		data.HelpSettingsUniversalSearchKnowledgeBaseSearch = types.StringValue(value.String())
	} else {
		data.HelpSettingsUniversalSearchKnowledgeBaseSearch = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *OrganizationBrandingPolicy) fromBodyPartial(ctx context.Context, res meraki.Res) {
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("adminSettings.appliesTo"); value.Exists() && !data.AdminSettingsAppliesTo.IsNull() {
		data.AdminSettingsAppliesTo = types.StringValue(value.String())
	} else {
		data.AdminSettingsAppliesTo = types.StringNull()
	}
	if value := res.Get("adminSettings.values"); value.Exists() && !data.AdminSettingsValues.IsNull() {
		data.AdminSettingsValues = helpers.GetStringList(value.Array())
	} else {
		data.AdminSettingsValues = types.ListNull(types.StringType)
	}
	if value := res.Get("customLogo.enabled"); value.Exists() && !data.CustomLogoEnabled.IsNull() {
		data.CustomLogoEnabled = types.BoolValue(value.Bool())
	} else {
		data.CustomLogoEnabled = types.BoolNull()
	}
	if value := res.Get("customLogo.image.contents"); value.Exists() && !data.CustomLogoImageContents.IsNull() {
		data.CustomLogoImageContents = types.StringValue(value.String())
	} else {
		data.CustomLogoImageContents = types.StringNull()
	}
	if value := res.Get("customLogo.image.format"); value.Exists() && !data.CustomLogoImageFormat.IsNull() {
		data.CustomLogoImageFormat = types.StringValue(value.String())
	} else {
		data.CustomLogoImageFormat = types.StringNull()
	}
	if value := res.Get("helpSettings.apiDocsSubtab"); value.Exists() && !data.HelpSettingsApiDocsSubtab.IsNull() {
		data.HelpSettingsApiDocsSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsApiDocsSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.casesSubtab"); value.Exists() && !data.HelpSettingsCasesSubtab.IsNull() {
		data.HelpSettingsCasesSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsCasesSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.ciscoMerakiProductDocumentation"); value.Exists() && !data.HelpSettingsCiscoMerakiProductDocumentation.IsNull() {
		data.HelpSettingsCiscoMerakiProductDocumentation = types.StringValue(value.String())
	} else {
		data.HelpSettingsCiscoMerakiProductDocumentation = types.StringNull()
	}
	if value := res.Get("helpSettings.communitySubtab"); value.Exists() && !data.HelpSettingsCommunitySubtab.IsNull() {
		data.HelpSettingsCommunitySubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsCommunitySubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.dataProtectionRequestsSubtab"); value.Exists() && !data.HelpSettingsDataProtectionRequestsSubtab.IsNull() {
		data.HelpSettingsDataProtectionRequestsSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsDataProtectionRequestsSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.firewallInfoSubtab"); value.Exists() && !data.HelpSettingsFirewallInfoSubtab.IsNull() {
		data.HelpSettingsFirewallInfoSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsFirewallInfoSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.getHelpSubtab"); value.Exists() && !data.HelpSettingsGetHelpSubtab.IsNull() {
		data.HelpSettingsGetHelpSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsGetHelpSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.getHelpSubtabKnowledgeBaseSearch"); value.Exists() && !data.HelpSettingsGetHelpSubtabKnowledgeBaseSearch.IsNull() {
		data.HelpSettingsGetHelpSubtabKnowledgeBaseSearch = types.StringValue(value.String())
	} else {
		data.HelpSettingsGetHelpSubtabKnowledgeBaseSearch = types.StringNull()
	}
	if value := res.Get("helpSettings.hardwareReplacementsSubtab"); value.Exists() && !data.HelpSettingsHardwareReplacementsSubtab.IsNull() {
		data.HelpSettingsHardwareReplacementsSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsHardwareReplacementsSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.helpTab"); value.Exists() && !data.HelpSettingsHelpTab.IsNull() {
		data.HelpSettingsHelpTab = types.StringValue(value.String())
	} else {
		data.HelpSettingsHelpTab = types.StringNull()
	}
	if value := res.Get("helpSettings.helpWidget"); value.Exists() && !data.HelpSettingsHelpWidget.IsNull() {
		data.HelpSettingsHelpWidget = types.StringValue(value.String())
	} else {
		data.HelpSettingsHelpWidget = types.StringNull()
	}
	if value := res.Get("helpSettings.newFeaturesSubtab"); value.Exists() && !data.HelpSettingsNewFeaturesSubtab.IsNull() {
		data.HelpSettingsNewFeaturesSubtab = types.StringValue(value.String())
	} else {
		data.HelpSettingsNewFeaturesSubtab = types.StringNull()
	}
	if value := res.Get("helpSettings.smForums"); value.Exists() && !data.HelpSettingsSmForums.IsNull() {
		data.HelpSettingsSmForums = types.StringValue(value.String())
	} else {
		data.HelpSettingsSmForums = types.StringNull()
	}
	if value := res.Get("helpSettings.supportContactInfo"); value.Exists() && !data.HelpSettingsSupportContactInfo.IsNull() {
		data.HelpSettingsSupportContactInfo = types.StringValue(value.String())
	} else {
		data.HelpSettingsSupportContactInfo = types.StringNull()
	}
	if value := res.Get("helpSettings.universalSearchKnowledgeBaseSearch"); value.Exists() && !data.HelpSettingsUniversalSearchKnowledgeBaseSearch.IsNull() {
		data.HelpSettingsUniversalSearchKnowledgeBaseSearch = types.StringValue(value.String())
	} else {
		data.HelpSettingsUniversalSearchKnowledgeBaseSearch = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *OrganizationBrandingPolicy) fromBodyUnknowns(ctx context.Context, res meraki.Res) {
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toIdentity

func (data *OrganizationBrandingPolicyIdentity) toIdentity(ctx context.Context, plan *OrganizationBrandingPolicy) {
	data.OrganizationId = plan.OrganizationId
	data.Id = plan.Id
}

// End of section. //template:end toIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin fromIdentity

func (data *OrganizationBrandingPolicy) fromIdentity(ctx context.Context, identity *OrganizationBrandingPolicyIdentity) {
	data.OrganizationId = identity.OrganizationId
	data.Id = identity.Id
}

// End of section. //template:end fromIdentity

// Section below is generated&owned by "gen/generator.go". //template:begin toDestroyBody

func (data OrganizationBrandingPolicy) toDestroyBody(ctx context.Context) string {
	body := ""
	return body
}

// End of section. //template:end toDestroyBody
