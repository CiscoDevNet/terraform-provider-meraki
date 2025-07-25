---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "meraki_sensor_mqtt_broker Resource - terraform-provider-meraki"
subcategory: "Sensors"
description: |-
  This resource can manage the Sensor MQTT Broker configuration.
---

# meraki_sensor_mqtt_broker (Resource)

This resource can manage the `Sensor MQTT Broker` configuration.

## Example Usage

```terraform
resource "meraki_sensor_mqtt_broker" "example" {
  network_id     = "L_123456"
  mqtt_broker_id = "123456"
  enabled        = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) Set to true to enable MQTT broker for sensor network
- `mqtt_broker_id` (String) MQTT Broker ID
- `network_id` (String) Network ID

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import meraki_sensor_mqtt_broker.example "<network_id>,<mqtt_broker_id>"
```
