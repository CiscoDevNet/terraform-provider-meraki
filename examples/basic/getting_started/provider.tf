terraform {
  required_providers {
    meraki = {
      source = "CiscoDevNet/meraki"
    }
  }
}

provider "meraki" {
  // By default uses env $MERAKI_API_KEY
}
