## 1.1.0

- Add `meraki_network_alerts_settings` resource and data source
- Fix issue with not handling paginated responses correctly, [link](https://github.com/CiscoDevNet/terraform-provider-meraki/issues/71)
- Add `force_delete` attribute to `meraki_network_group_policy` resource, [link](https://github.com/CiscoDevNet/terraform-provider-meraki/issues/72)
- Add `details_by_device` attribute to `meraki_network_device_claim` resource
- Add `mac_whitelist_limit` attribute to `meraki_switch_port` resource and data sources
- Add `adaptive_policy_group_id` attribute to `meraki_wireless_ssid` resource and data source
- Add `meraki_wireless_location_scanning` resource and data source
- Add `meraki_wireless_location_scanning_receiver` resource and data sources
- Fix idempotency issue with `syslog_default_rule` attribute of `meraki_appliance_vpn_firewall_rules` resource, [link](https://github.com/CiscoDevNet/terraform-provider-meraki/issues/73)
- Fix idempotency issue with `syslog_default_rule` attribute of `meraki_appliance_l3_firewall_rules` resource, [link](https://github.com/CiscoDevNet/terraform-provider-meraki/issues/73)
- Fix idempotency issue with `syslog_default_rule` attribute of `meraki_appliance_inbound_firewall_rules` resource, [link](https://github.com/CiscoDevNet/terraform-provider-meraki/issues/73)

## 1.0.0

- Add `meraki_network_device_claim_vmx` resource
- Add `meraki_appliance_dns_local_profile_assignments` resource
- Add `meraki_appliance_dns_local_record` resource and data sources
- Add `meraki_appliance_dns_split_profile` resource and data sources
- Add `meraki_appliance_dns_split_profile_assignments` resource
- Add `meraki_sm_admin_role` resource and data sources
- Add `meraki_sm_target_group` resource and data sources
- Add warning if name query option of data source is being used and multiple objects with the same name exist
- BREAKING CHANGE: Rename `switch_port_ids` attribute of `meraki_switch_organization_ports_profiles_automation` resource and data source to `port_ids`

## 0.1.12

- Add MV84X settings to `meraki_camera_quality_retention_profile` resource and data source
- Add `minimum_password_length` attribute to `meraki_organization_login_security` resource and data source
- Add `include_sensor_url` and `message` attributes to `meraki_sensor_alerts_profile` resource and data source
- Make `name` attribute of `meraki_switch_routing_interface` resource and data source required
- Add `mode` attribute to `meraki_wireless_network_electronic_shelf_label` resource and data source
- Change type of `floor_number` attribute of `meraki_network_floor_plan` resource and data source from `integer` to `float`
- Add `meraki_wireless_ssid_firewall_isolation_allowlist_entry` resource and data sources
- Remove `encryption_enabled` and `encryption_certificate_id` attributes from `meraki_network_syslog_servers` resource and data source
- Add `meraki_appliance_firewall_multicast_forwarding` resource and data source
- Add `meraki_appliance_dns_local_profile` resource and data sources

## 0.1.11

- Do not import default rule for `meraki_appliance_cellular_firewall_rules`, `meraki_appliance_inbound_cellular_firewall_rules`, `meraki_appliance_inbound_firewall_rules`, `meraki_appliance_l3_firewall_rules`, `meraki_appliance_vpn_firewall_rules` and `meraki_switch_access_control_lists` resource
- Do not import default and local LAN access rules for `meraki_wireless_ssid_l3_firewall_rules` resource

## 0.1.10

- Add `value_countries` attribute to `meraki_appliance_l7_firewall_rules` resource and data source

## 0.1.9

- BREAKING CHANGE: Remove `application_value` attribute from `meraki_wireless_ssid_traffic_shaping_rules` resource and data source
- Add `subnet_nat_is_allowed`, `nat_enabled` and `nat_remote_subnet` attributes to `meraki_appliance_site_to_site_vpn` resource and data source
- Add `video_settings_mv53_x_quality` and `video_settings_mv53_x_resolution` attributes to `meraki_camera_quality_retention_profile` resource and data sources
- Add `self_registration_authorization_type` and `self_registration_enabled` attributes to `meraki_wireless_ssid_splash_settings` resource and data source

## 0.1.8

- Add `application_value` attribute to `meraki_wireless_ssid_traffic_shaping_rules` resource and data source

## 0.1.7

- Add `meraki_organization_auth_radius_server` resource and data sources
- Add `meraki_camera_custom_analytics` resource and data source
- Add `meraki_camera_quality_retention` resource and data source
- Add `meraki_camera_quality_retention_profile` resource and data sources
- Add `meraki_camera_role` resource and data sources
- Add `meraki_camera_sense` resource and data source
- Add `meraki_camera_video_settings` resource and data source
- Add `meraki_camera_wireless_profile` resource and data sources
- Add `meraki_camera_device_wireless_profiles` resource and data source
- Add `meraki_insight_monitored_media_server` resource and data source

## 0.1.6

- Add `meraki_sensor_alerts_profile` resource and data source
- Add `meraki_sensor_mqtt_broker` resource and data sources
- Add `meraki_sensor_relationships` resource and data source
- Add `meraki_sensor_network_relationships` data source

## 0.1.5

- Delete all rules when destroying `meraki_appliance_cellular_firewall_rules`, `meraki_appliance_inbound_cellular_firewall_rules`, `meraki_appliance_inbound_firewall_rules`, `meraki_appliance_l3_firewall_rules`, `meraki_appliance_l7_firewall_rules`, `meraki_appliance_one_to_many_nat_rules`, `meraki_appliance_one_to_one_nat_rules`, `meraki_appliance_port_forwarding_rules`, `meraki_appliance_traffic_shaping_rules`, `meraki_appliance_vpn_firewall_rules`, `meraki_wireless_ssid_l3_firewall_rules`, `meraki_wireless_ssid_l7_firewall_rules`, `meraki_wireless_ssid_traffic_shaping_rules` resources
- Add `floor_number` attribute to `meraki_network_floor_plan` resource and data sources
- Add `local_status_page_authentication_username` attribute to `meraki_network_settings` resource and data source
- Add encryption attributes to `meraki_network_syslog_servers` resource and data source
- Add `log` and `tcp_established` attributes to `meraki_organization_adaptive_policy_acl` resource and data sources
- Add `stackwise_virtual_is_dual_active_detector` and `stackwise_virtual_is_stack_wise_virtual_link` attributes to `meraki_switch_ports` data source
- Add `radius_radsec_tls_tunnel_timeout` attribute to `meraki_wireless_ssid` resource and data source
- Add `meraki_cellular_gateway_connectivity_monitoring_destinations` resource and data sources
- Add `meraki_cellular_gateway_dhcp` resource and data sources
- Add `meraki_cellular_gateway_lan` resource and data sources
- Add `meraki_cellular_gateway_port_forwarding_rules` resource and data sources
- Add `meraki_cellular_gateway_subnet_pool` resource and data sources
- Add `meraki_cellular_gateway_uplink` resource and data sources
- Add provider configuration to define a list of HTTP error codes to retry on

## 0.1.4

- Add `ip_version` attribute to `meraki_wireless_ssid_l3_firewall_rules` resource and data source
- Add `meraki_appliance_vmx_authentication_token` resource
- Configure default settings when deleting wireless SSID using `meraki_wireless_ssid` resource
- Add `fixed_ip_assignments` attribute to `meraki_appliance_vlan` resource and data sources

## 0.1.3

- Add `public_hostname` attribute to `meraki_appliance_third_party_vpn_peers` resource and data source
- Add `treat_these_traffic_types_as_one_threshold` attribute to `meraki_switch_storm_control` resource and data source
- Add `dhcp_boot_filename`, `dhcp_boot_next_server`, `dns_nameservers`, `vpn_nat_subnet`, `dhcp_relay_server_ips` and `reserved_ip_ranges` attributes to `meraki_appliance_vlan` resource and data source
- Add `meraki_organization_early_access_features_opt_in` resource and data sources
- Add `meraki_switch_organization_ports_profile` resource and data sources
- Add `meraki_switch_organization_ports_profiles_automation` resource and data sources
- Add `meraki_appliance_vlan_dhcp` resource
- Add `meraki_organization_license` resource and data source
- Add `meraki_network_firmware_upgrades` resource and data source
- Add `meraki_organization_licenses` data source
- Add `meraki_appliance_ports` data source
- Add `meraki_appliance_ssids` data source
- Add `meraki_switch_ports` data source
- Add `meraki_wireless_ssids` data source
- Add `meraki_organization_early_access_features` data source
- Add `meraki_organization_branding_policy` resource and data sources
- Add `meraki_organization_branding_policies_priorities` resource and data source
- Add `meraki_network_client_splash_authorization_status` resource and data source
- Add `meraki_network_devices` data source
- Add `meraki_organization_devices` data source
- Add `meraki_organization_firmware_upgrades` data source
- Add `meraki_organization_inventory_devices` data source
- Add `meraki_network_policies_by_client` data source
- Add `meraki_network_vlan_profile_assignments_by_device` data source

## 0.1.2

- Add `meraki_appliance_firewalled_service` resource and data source
- Add `meraki_appliance_inbound_cellular_firewall_rules` resource and data source
- Add `meraki_appliance_inbound_firewall_rules` resource and data source
- Add `meraki_appliance_l3_firewall_rules` resource and data source
- Add `meraki_appliance_l7_firewall_rules` resource and data source
- Add `meraki_appliance_one_to_many_nat_rules` resource and data source
- Add `meraki_appliance_one_to_one_nat_rules` resource and data source
- Add `meraki_appliance_port_forwarding_rules` resource and data source
- Add `meraki_appliance_firewall_settings` resource and data source
- Add `meraki_appliance_port` resource and data source
- Add `meraki_appliance_vlans_settings` resource and data source
- Add `meraki_appliance_prefix_delegated_static` resource and data source
- Add `meraki_appliance_prefix_delegated_statics` data source
- Add `meraki_appliance_radio_settings` resource and data source
- Add `meraki_appliance_rf_profile` resource and data source
- Add `meraki_appliance_rf_profiles` data source
- BREAKING CHANGE: Rename `per_ssid_settingsXX_*` attributes of `meraki_wireless_rf_profile` resource and data sources to `per_ssid_settings_XX_*`
- Add `meraki_appliance_sdwan_internet_policies` resource
- Add `meraki_appliance_network_security_intrusion` resource and data source
- Add `meraki_appliance_organization_security_intrusion` resource and data source
- Add `meraki_appliance_security_malware` resource and data source
- Add `meraki_appliance_settings` resource and data source
- Add `meraki_appliance_single_lan` resource and data source
- Add `meraki_appliance_ssid` resource and data source
- Add `meraki_appliance_static_route` resource and data source
- Add `meraki_appliance_static_routes` data source
- Add `meraki_appliance_vlan` resource and data source
- Add `meraki_appliance_vlans` data source
- Add `meraki_appliance_traffic_shaping` resource and data source
- Add `meraki_appliance_traffic_shaping_custom_performance_class` resource and data source
- Add `meraki_appliance_traffic_shaping_custom_performance_classes` data source
- Add `meraki_appliance_traffic_shaping_rules` resource and data source
- Add `meraki_appliance_traffic_shaping_uplink_bandwidth` resource and data source
- Add `meraki_appliance_traffic_shaping_uplink_selection` resource and data source
- Add `meraki_appliance_traffic_shaping_vpn_exclusions` resource and data source
- Add `meraki_appliance_uplink_settings` resource and data source
- Add `meraki_appliance_site_to_site_vpn` resource and data source
- Add `meraki_appliance_third_party_vpn_peers` resource and data source
- Add `meraki_appliance_vpn_firewall_rules` resource and data source
- Add `meraki_appliance_warm_spare` resource and data source
- Fix idempotency issue with `meraki_switch_stp` resource, [link](https://github.com/CiscoDevNet/terraform-provider-meraki/pull/17)

## 0.1.1

- Add `meraki_network_floor_plans` data source
- Add `meraki_network_group_policies` data source
- Add `meraki_network_auth_users` data source
- Add `meraki_network_mqtt_brokers` data source
- Add `meraki_network_vlan_profiles` data source
- Add `meraki_network_webhook_http_servers` data source
- Add `meraki_network_webhook_payload_templates` data source
- Add `meraki_networks` data source
- Add `meraki_organization_adaptive_policies` data source
- Add `meraki_organization_adaptive_policy_acls` data source
- Add `meraki_organization_adaptive_policy_groups` data source
- Add `meraki_organization_admins` data source
- Add `meraki_organization_alerts_profiles` data source
- Add `meraki_organization_config_templates` data source
- Add `meraki_organization_policy_object_groups` data source
- Add `meraki_organization_policy_objects` data source
- Add `meraki_organization_saml_idps` data source
- Add `meraki_organization_saml_roles` data source
- Add `meraki_organizations` data source
- Add `meraki_switch_access_policies` data source
- Add `meraki_switch_dhcp_server_policy_arp_inspection_trusted_servers` data source
- Add `meraki_switch_link_aggregations` data source
- Add `meraki_switch_port_schedules` data source
- Add `meraki_switch_qos_rules` data source
- Add `meraki_switch_routing_interfaces` data source
- Add `meraki_switch_routing_multicast_rendezvous_points` data source
- Add `meraki_switch_routing_static_routes` data source
- Add `meraki_switch_stack_routing_interfaces` data source
- Add `meraki_switch_stack_routing_static_routes` data source
- Add `meraki_switch_stacks` data source
- Add `meraki_wireless_ethernet_port_profiles` data source
- Add `meraki_wireless_rf_profiles` data source
- Add `meraki_wireless_ssid_identity_psk_users` data source

## 0.1.0

- Initial release
