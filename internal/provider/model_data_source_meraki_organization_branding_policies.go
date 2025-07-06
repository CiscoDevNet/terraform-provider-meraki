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

type DataSourceOrganizationBrandingPolicies struct {
	OrganizationId types.String                                  `tfsdk:"organization_id"`
	Items          []DataSourceOrganizationBrandingPoliciesItems `tfsdk:"items"`
}

type DataSourceOrganizationBrandingPoliciesItems struct {
	Id                                             types.String `tfsdk:"id"`
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

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DataSourceOrganizationBrandingPolicies) getPath() string {
	return fmt.Sprintf("/organizations/%v/brandingPolicies", url.QueryEscape(data.OrganizationId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DataSourceOrganizationBrandingPolicies) fromBody(ctx context.Context, res meraki.Res) {
	data.Items = make([]DataSourceOrganizationBrandingPoliciesItems, 0)
	res.ForEach(func(k, res gjson.Result) bool {
		parent := &data
		data := DataSourceOrganizationBrandingPoliciesItems{}
		data.Id = types.StringValue(res.Get("id").String())
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
		(*parent).Items = append((*parent).Items, data)
		return true
	})
}

// End of section. //template:end fromBody
