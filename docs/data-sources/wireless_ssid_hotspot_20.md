---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_wireless_ssid_hotspot_20 Data Source - terraform-provider-meraki"
subcategory: "Wireless"
description: |-
  This data source can read the Wireless SSID Hotspot 2.0 configuration.
---

# meraki_wireless_ssid_hotspot_20 (Data Source)

This data source can read the `Wireless SSID Hotspot 2.0` configuration.

## Example Usage

```terraform
data "meraki_wireless_ssid_hotspot_20" "example" {
  network_id = "L_123456"
  number     = "0"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `network_id` (String) Network ID
- `number` (String) Wireless SSID number

### Read-Only

- `domains` (List of String) An array of domain names
- `enabled` (Boolean) Whether or not Hotspot 2.0 for this SSID is enabled
- `id` (String) The id of the object
- `mcc_mncs` (Attributes List) An array of MCC/MNC pairs (see [below for nested schema](#nestedatt--mcc_mncs))
- `nai_realms` (Attributes List) An array of NAI realms (see [below for nested schema](#nestedatt--nai_realms))
- `network_access_type` (String) The network type of this SSID (`Private network`, `Private network with guest access`, `Chargeable public network`, `Free public network`, `Personal device network`, `Emergency services only network`, `Test or experimental`, `Wildcard`)
- `operator_name` (String) Operator name
- `roam_consort_ois` (List of String) An array of roaming consortium OIs (hexadecimal number 3-5 octets in length)
- `venue_name` (String) Venue name
- `venue_type` (String) Venue type (`Unspecified`, `Unspecified Assembly`, `Arena`, `Stadium`, `Passenger Terminal`, `Amphitheater`, `Amusement Park`, `Place of Worship`, `Convention Center`, `Library`, `Museum`, `Restaurant`, `Theater`, `Bar`, `Coffee Shop`, `Zoo or Aquarium`, `Emergency Coordination Center`, `Unspecified Business`, `Doctor or Dentist office`, `Bank`, `Fire Station`, `Police Station`, `Post Office`, `Professional Office`, `Research and Development Facility`, `Attorney Office`, `Unspecified Educational`, `School, Primary`, `School, Secondary`, `University or College`, `Unspecified Factory and Industrial`, `Factory`, `Unspecified Institutional`, `Hospital`, `Long-Term Care Facility`, `Alcohol and Drug Rehabilitation Center`, `Group Home`, `Prison or Jail`, `Unspecified Mercantile`, `Retail Store`, `Grocery Market`, `Automotive Service Station`, `Shopping Mall`, `Gas Station`, `Unspecified Residential`, `Private Residence`, `Hotel or Motel`, `Dormitory`, `Boarding House`, `Unspecified Storage`, `Unspecified Utility and Miscellaneous`, `Unspecified Vehicular`, `Automobile or Truck`, `Airplane`, `Bus`, `Ferry`, `Ship or Boat`, `Train`, `Motor Bike`, `Unspecified Outdoor`, `Muni-mesh Network`, `City Park`, `Rest Area`, `Traffic Control`, `Bus Stop`, `Kiosk`)

<a id="nestedatt--mcc_mncs"></a>
### Nested Schema for `mcc_mncs`

Read-Only:

- `mcc` (String) MCC value
- `mnc` (String) MNC value


<a id="nestedatt--nai_realms"></a>
### Nested Schema for `nai_realms`

Read-Only:

- `format` (String) The format for the realm (`1` or `0`)
- `methods` (Attributes List) An array of EAP methods for the realm. (see [below for nested schema](#nestedatt--nai_realms--methods))
- `realm` (String) The name of the realm

<a id="nestedatt--nai_realms--methods"></a>
### Nested Schema for `nai_realms.methods`

Read-Only:

- `authentication_types_credentials` (List of String) An array of autentication types. Possible values are `SIM`, `USIM`, `NFC Secure Element`, `Hardware Token`, `Softoken`, `Certificate`, `username/password`, `none`, `Reserved`, `Vendor Specific`.
- `authentication_types_eap_inner_authentication` (List of String) An array of autentication types. Possible values are `EAP-TLS`, `EAP-SIM`, `EAP-AKA`, `EAP-TTLS with MSCHAPv2`.
- `authentication_types_non_eap_inner_authentication` (List of String) An array of autentication types. Possible values are `Reserved`, `PAP`, `CHAP`, `MSCHAP`, `MSCHAPV2`.
- `authentication_types_tunneled_eap_method_credentials` (List of String) An array of autentication types. Possible values are `SIM`, `USIM`, `NFC Secure Element`, `Hardware Token`, `Softoken`, `Certificate`, `username/password`, `Reserved`, `Anonymous`, `Vendor Specific`.
- `id` (String) ID of method