resource "meraki_organization_branding_policy" "example" {
  organization_id                                      = "123456"
  enabled                                              = true
  name                                                 = "My Branding Policy"
  admin_settings_applies_to                            = "All admins of networks..."
  admin_settings_values                                = ["N_1234"]
  custom_logo_enabled                                  = true
  custom_logo_image_contents                           = "Hyperg26C8F4h8CvcoUqpA=="
  custom_logo_image_format                             = "jpg"
  help_settings_api_docs_subtab                        = "default or inherit"
  help_settings_cases_subtab                           = "hide"
  help_settings_cisco_meraki_product_documentation     = "show"
  help_settings_community_subtab                       = "show"
  help_settings_data_protection_requests_subtab        = "default or inherit"
  help_settings_firewall_info_subtab                   = "hide"
  help_settings_get_help_subtab                        = "default or inherit"
  help_settings_get_help_subtab_knowledge_base_search  = "<h1>Some custom HTML content</h1>"
  help_settings_hardware_replacements_subtab           = "hide"
  help_settings_help_tab                               = "show"
  help_settings_help_widget                            = "hide"
  help_settings_new_features_subtab                    = "show"
  help_settings_sm_forums                              = "hide"
  help_settings_support_contact_info                   = "show"
  help_settings_universal_search_knowledge_base_search = "hide"
}
