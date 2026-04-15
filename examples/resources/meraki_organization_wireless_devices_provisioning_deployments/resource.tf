resource "meraki_organization_wireless_devices_provisioning_deployments" "example" {
  organization_id             = "123456"
  meta_counts_items_remaining = 0
  meta_counts_items_total     = 20
  items = [
    {
      completed_at                = "2018-02-11T00:00:00.090210Z"
      created_at                  = "2018-02-11T00:00:00.090210Z"
      deployment_id               = "1284392014819"
      last_updated_at             = "2018-02-11T00:00:00.090210Z"
      requested_at                = "2018-02-11T00:00:00.090210Z"
      status                      = "ready"
      type                        = "replace"
      devices_new_mac             = "00:11:22:33:44:55"
      devices_new_model           = "CW9166I"
      devices_new_name            = "My AP"
      devices_new_serial          = "Q234-ABCD-5678"
      devices_new_rf_profile_id   = "1284392014819"
      devices_new_rf_profile_name = "RF Profile Name"
      devices_new_tags            = ["tag1"]
      devices_old_after_action    = "unclaim"
      devices_old_mac             = "00:11:22:33:44:55"
      devices_old_model           = "MR34"
      devices_old_name            = "My AP"
      devices_old_serial          = "Q234-ABCD-5678"
      devices_old_rf_profile_id   = "1284392014819"
      devices_old_rf_profile_name = "RF Profile Name"
      devices_old_tags            = ["tag1"]
      network_id                  = "N_24329156"
      network_name                = "Main Office"
      errors                      = ["error message1"]
    }
  ]
}
