---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_ssid Data Source - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This data source can read the Wireless SSID configuration.
---

# meraki_wireless_ssid (Data Source)

This data source can read the `Wireless SSID` configuration.

## Example Usage

```terraform
data "meraki_wireless_ssid" "example" {
  network_id = "L_123456"
  number     = "0"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `number` (String) Wireless SSID number

### Optional

- `id` (String) The id of the object
- `name` (String) The name of the SSID
- `named_vlans_radius_guest_vlan_name` (String) RADIUS guest VLAN name.

### Read-Only

- `active_directory_credentials_logon_name` (String) The logon name of the Active Directory account.
- `active_directory_credentials_password` (String) The password to the Active Directory user account.
- `active_directory_servers` (Attributes List) The Active Directory servers to be used for authentication. (see [below for nested schema](#nestedatt--active_directory_servers))
- `adult_content_filtering_enabled` (Boolean) Boolean indicating whether or not adult content will be blocked
- `ap_tags_and_vlan_ids` (Attributes List) The list of tags and VLAN IDs used for VLAN tagging. This param is only valid when the ipAssignmentMode is `Bridge mode` or `Layer 3 roaming` (see [below for nested schema](#nestedatt--ap_tags_and_vlan_ids))
- `auth_mode` (String) The association control method for the SSID (`open`, `open-enhanced`, `psk`, `open-with-radius`, `open-with-nac`, `8021x-meraki`, `8021x-nac`, `8021x-radius`, `8021x-google`, `8021x-entra`, `8021x-localradius`, `ipsk-with-radius`, `ipsk-without-radius` or `ipsk-with-nac`)
- `availability_tags` (List of String) Accepts a list of tags for this SSID. If availableOnAllAps is false, then the SSID will only be broadcast by APs with tags matching any of the tags in this list.
- `available_on_all_aps` (Boolean) Boolean indicating whether all APs should broadcast the SSID or if it should be restricted to APs matching any availability tags. Can only be false if the SSID has availability tags.
- `band_selection` (String) The client-serving radio frequencies of this SSID in the default indoor RF profile. (`Dual band operation`, `5 GHz band only` or `Dual band operation with Band Steering`)
- `concentrator_network_id` (String) The concentrator to use when the ipAssignmentMode is `Layer 3 roaming with a concentrator` or `VPN`.
- `default_vlan_id` (Number) The default VLAN ID used for `all other APs`. This param is only valid when the ipAssignmentMode is `Bridge mode` or `Layer 3 roaming`
- `disassociate_clients_on_vpn_failover` (Boolean) Disassociate clients when `VPN` concentrator failover occurs in order to trigger clients to re-associate and generate new DHCP requests. This param is only valid if ipAssignmentMode is `VPN`.
- `dns_rewrite_dns_custom_nameservers` (List of String) User specified DNS servers (up to two servers)
- `dns_rewrite_enabled` (Boolean) Boolean indicating whether or not DNS server rewrite is enabled. If disabled, upstream DNS will be used
- `dot11r_adaptive` (Boolean) (Optional) Whether 802.11r is adaptive or not.
- `dot11r_enabled` (Boolean) Whether 802.11r is enabled or not.
- `dot11w_enabled` (Boolean) Whether 802.11w is enabled or not.
- `dot11w_required` (Boolean) (Optional) Whether 802.11w is required or not.
- `enabled` (Boolean) Whether or not the SSID is enabled
- `encryption_mode` (String) The psk encryption mode for the SSID (`wep` or `wpa`). This param is only valid if the authMode is `psk`
- `enterprise_admin_access` (String) Whether or not an SSID is accessible by `enterprise` administrators (`access disabled` or `access enabled`)
- `gre_concentrator_host` (String) The EoGRE concentrator`s IP or FQDN. This param is required when ipAssignmentMode is `Ethernet over GRE`.
- `gre_key` (Number) Optional numerical identifier that will add the GRE key field to the GRE header. Used to identify an individual traffic flow within a tunnel.
- `ip_assignment_mode` (String) The client IP assignment mode (`NAT mode`, `Bridge mode`, `Layer 3 roaming`, `Ethernet over GRE`, `Layer 3 roaming with a concentrator` or `VPN`)
- `lan_isolation_enabled` (Boolean) Boolean indicating whether Layer 2 LAN isolation should be enabled or disabled. Only configurable when ipAssignmentMode is `Bridge mode`.
- `ldap_base_distinguished_name` (String) The base distinguished name of users on the LDAP server.
- `ldap_credentials_distinguished_name` (String) The distinguished name of the LDAP user account (example: cn=user,dc=meraki,dc=com).
- `ldap_credentials_password` (String) The password of the LDAP user account.
- `ldap_server_ca_certificate_contents` (String) The contents of the CA certificate. Must be in PEM or DER format.
- `ldap_servers` (Attributes List) The LDAP servers to be used for authentication. (see [below for nested schema](#nestedatt--ldap_servers))
- `local_radius_cache_timeout` (Number) The duration (in seconds) for which LDAP and OCSP lookups are cached.
- `local_radius_certificate_authentication_client_root_ca_certificate_contents` (String) The contents of the Client CA Certificate. Must be in PEM or DER format.
- `local_radius_certificate_authentication_enabled` (Boolean) Whether or not to use EAP-TLS certificate-based authentication to validate wireless clients.
- `local_radius_certificate_authentication_ocsp_responder_url` (String) (Optional) The URL of the OCSP responder to verify client certificate status.
- `local_radius_certificate_authentication_use_ldap` (Boolean) Whether or not to verify the certificate with LDAP.
- `local_radius_certificate_authentication_use_ocsp` (Boolean) Whether or not to verify the certificate with OCSP.
- `local_radius_password_authentication_enabled` (Boolean) Whether or not to use EAP-TTLS/PAP or PEAP-GTC password-based authentication via LDAP lookup.
- `mandatory_dhcp_enabled` (Boolean) If true, Mandatory DHCP will enforce that clients connecting to this SSID must use the IP address assigned by the DHCP server. Clients who use a static IP address won`t be able to associate.
- `min_bitrate` (Number) The minimum bitrate in Mbps of this SSID in the default indoor RF profile. (`1`, `2`, `5.5`, `6`, `9`, `11`, `12`, `18`, `24`, `36`, `48` or `54`)
- `named_vlans_radius_guest_vlan_enabled` (Boolean) Whether or not RADIUS guest named VLAN is enabled.
- `named_vlans_tagging_by_ap_tags` (Attributes List) The list of AP tags and VLAN names used for named VLAN tagging. If an AP has a tag matching one in the list, then traffic on this SSID will be directed to use the VLAN name associated to the tag. (see [below for nested schema](#nestedatt--named_vlans_tagging_by_ap_tags))
- `named_vlans_tagging_default_vlan_name` (String) The default VLAN name used to tag traffic in the absence of a matching AP tag.
- `named_vlans_tagging_enabled` (Boolean) Whether or not traffic should be directed to use specific VLAN names.
- `oauth_allowed_domains` (List of String) (Optional) The list of domains allowed access to the network.
- `per_client_bandwidth_limit_down` (Number) The download bandwidth limit in Kbps. (0 represents no limit.)
- `per_client_bandwidth_limit_up` (Number) The upload bandwidth limit in Kbps. (0 represents no limit.)
- `per_ssid_bandwidth_limit_down` (Number) The total download bandwidth limit in Kbps. (0 represents no limit.)
- `per_ssid_bandwidth_limit_up` (Number) The total upload bandwidth limit in Kbps. (0 represents no limit.)
- `psk` (String) The passkey for the SSID. This param is only valid if the authMode is `psk`
- `radius_accounting_enabled` (Boolean) Whether or not RADIUS accounting is enabled. This param is only valid if the authMode is `open-with-radius`, `8021x-radius` or `ipsk-with-radius`
- `radius_accounting_interim_interval` (Number) The interval (in seconds) in which accounting information is updated and sent to the RADIUS accounting server.
- `radius_accounting_servers` (Attributes List) The RADIUS accounting 802.1X servers to be used for authentication. This param is only valid if the authMode is `open-with-radius`, `8021x-radius` or `ipsk-with-radius` and radiusAccountingEnabled is `true` (see [below for nested schema](#nestedatt--radius_accounting_servers))
- `radius_attribute_for_group_policies` (String) Specify the RADIUS attribute used to look up group policies (`Filter-Id`, `Reply-Message`, `Airespace-ACL-Name` or `Aruba-User-Role`). Access points must receive this attribute in the RADIUS Access-Accept message
- `radius_authentication_nas_id` (String) The template of the NAS identifier to be used for RADIUS authentication (ex. $NODE_MAC$:$VAP_NUM$).
- `radius_called_station_id` (String) The template of the called station identifier to be used for RADIUS (ex. $NODE_MAC$:$VAP_NUM$).
- `radius_coa_enabled` (Boolean) If true, Meraki devices will act as a RADIUS Dynamic Authorization Server and will respond to RADIUS Change-of-Authorization and Disconnect messages sent by the RADIUS server.
- `radius_failover_policy` (String) This policy determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable (`Deny access` or `Allow access`)
- `radius_fallback_enabled` (Boolean) Whether or not higher priority RADIUS servers should be retried after 60 seconds.
- `radius_guest_vlan_enabled` (Boolean) Whether or not RADIUS Guest VLAN is enabled. This param is only valid if the authMode is `open-with-radius` and addressing mode is not set to `isolated` or `nat` mode
- `radius_guest_vlan_id` (Number) VLAN ID of the RADIUS Guest VLAN. This param is only valid if the authMode is `open-with-radius` and addressing mode is not set to `isolated` or `nat` mode
- `radius_load_balancing_policy` (String) This policy determines which RADIUS server will be contacted first in an authentication attempt and the ordering of any necessary retry attempts (`Strict priority order` or `Round robin`)
- `radius_override` (Boolean) If true, the RADIUS response can override VLAN tag. This is not valid when ipAssignmentMode is `NAT mode`.
- `radius_proxy_enabled` (Boolean) If true, Meraki devices will proxy RADIUS messages through the Meraki cloud to the configured RADIUS auth and accounting servers.
- `radius_server_attempts_limit` (Number) The maximum number of transmit attempts after which a RADIUS server is failed over (must be between 1-5).
- `radius_server_timeout` (Number) The amount of time for which a RADIUS client waits for a reply from the RADIUS server (must be between 1-10 seconds).
- `radius_servers` (Attributes List) The RADIUS 802.1X servers to be used for authentication. This param is only valid if the authMode is `open-with-radius`, `8021x-radius` or `ipsk-with-radius` (see [below for nested schema](#nestedatt--radius_servers))
- `radius_testing_enabled` (Boolean) If true, Meraki devices will periodically send Access-Request messages to configured RADIUS servers using identity `meraki_8021x_test` to ensure that the RADIUS servers are reachable.
- `secondary_concentrator_network_id` (String) The secondary concentrator to use when the ipAssignmentMode is `VPN`. If configured, the APs will switch to using this concentrator if the primary concentrator is unreachable. This param is optional. (`disabled` represents no secondary concentrator.)
- `speed_burst_enabled` (Boolean) Boolean indicating whether or not to allow users to temporarily exceed the bandwidth limit for short periods while still keeping them under the bandwidth limit over time.
- `splash_guest_sponsor_domains` (List of String) Array of valid sponsor email domains for sponsored guest splash type.
- `splash_page` (String) The type of splash page for the SSID (`None`, `Click-through splash page`, `Billing`, `Password-protected with Meraki RADIUS`, `Password-protected with custom RADIUS`, `Password-protected with Active Directory`, `Password-protected with LDAP`, `SMS authentication`, `Systems Manager Sentry`, `Facebook Wi-Fi`, `Google OAuth`, `Microsoft Entra ID`, `Sponsored guest`, `Cisco ISE` or `Google Apps domain`). This attribute is not supported for template children.
- `use_vlan_tagging` (Boolean) Whether or not traffic should be directed to use specific VLANs. This param is only valid if the ipAssignmentMode is `Bridge mode` or `Layer 3 roaming`
- `visible` (Boolean) Boolean indicating whether APs should advertise or hide this SSID. APs will only broadcast this SSID if set to true
- `vlan_id` (Number) The VLAN ID used for VLAN tagging. This param is only valid when the ipAssignmentMode is `Layer 3 roaming with a concentrator` or `VPN`
- `walled_garden_enabled` (Boolean) Allow access to a configurable list of IP ranges, which users may access prior to sign-on.
- `walled_garden_ranges` (List of String) Specify your walled garden by entering an array of addresses, ranges using CIDR notation, domain names, and domain wildcards (e.g. `192.168.1.1/24`, `192.168.37.10/32`, `www.yahoo.com`, `*.google.com`]). Meraki`s splash page is automatically included in your walled garden.
- `wpa_encryption_mode` (String) The types of WPA encryption. (`WPA1 only`, `WPA1 and WPA2`, `WPA2 only`, `WPA3 Transition Mode`, `WPA3 only` or `WPA3 192-bit Security`)

<a id="nestedatt--active_directory_servers"></a>
### Nested Schema for `active_directory_servers`

Read-Only:

- `host` (String) IP address (or FQDN) of your Active Directory server.
- `port` (Number) (Optional) UDP port the Active Directory server listens on. By default, uses port 3268.


<a id="nestedatt--ap_tags_and_vlan_ids"></a>
### Nested Schema for `ap_tags_and_vlan_ids`

Read-Only:

- `tags` (List of String) Array of AP tags
- `vlan_id` (Number) Numerical identifier that is assigned to the VLAN


<a id="nestedatt--ldap_servers"></a>
### Nested Schema for `ldap_servers`

Read-Only:

- `host` (String) IP address (or FQDN) of your LDAP server.
- `port` (Number) UDP port the LDAP server listens on.


<a id="nestedatt--named_vlans_tagging_by_ap_tags"></a>
### Nested Schema for `named_vlans_tagging_by_ap_tags`

Read-Only:

- `tags` (List of String) List of AP tags.
- `vlan_name` (String) VLAN name that will be used to tag traffic.


<a id="nestedatt--radius_accounting_servers"></a>
### Nested Schema for `radius_accounting_servers`

Read-Only:

- `ca_certificate` (String) Certificate used for authorization for the RADSEC Server
- `host` (String) IP address (or FQDN) to which the APs will send RADIUS accounting messages
- `port` (Number) Port on the RADIUS server that is listening for accounting messages
- `radsec_enabled` (Boolean) Use RADSEC (TLS over TCP) to connect to this RADIUS accounting server. Requires radiusProxyEnabled.
- `secret` (String) Shared key used to authenticate messages between the APs and RADIUS server


<a id="nestedatt--radius_servers"></a>
### Nested Schema for `radius_servers`

Read-Only:

- `ca_certificate` (String) Certificate used for authorization for the RADSEC Server
- `host` (String) IP address (or FQDN) of your RADIUS server
- `open_roaming_certificate_id` (Number) The ID of the Openroaming Certificate attached to radius server.
- `port` (Number) UDP port the RADIUS server listens on for Access-requests
- `radsec_enabled` (Boolean) Use RADSEC (TLS over TCP) to connect to this RADIUS server. Requires radiusProxyEnabled.
- `secret` (String) RADIUS client shared secret