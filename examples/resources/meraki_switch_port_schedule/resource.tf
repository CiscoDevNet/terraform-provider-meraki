resource "meraki_switch_port_schedule" "example" {
  network_id = "L_123456"
  name = "Weekdays schedule"
  port_schedule_friday_active = true
  port_schedule_friday_from = "09:00"
  port_schedule_friday_to = "17:00"
  port_schedule_monday_active = true
  port_schedule_monday_from = "09:00"
  port_schedule_monday_to = "17:00"
  port_schedule_saturday_active = false
  port_schedule_saturday_from = "00:00"
  port_schedule_saturday_to = "24:00"
  port_schedule_sunday_active = false
  port_schedule_sunday_from = "00:00"
  port_schedule_sunday_to = "24:00"
  port_schedule_thursday_active = true
  port_schedule_thursday_from = "09:00"
  port_schedule_thursday_to = "17:00"
  port_schedule_tuesday_active = true
  port_schedule_tuesday_from = "09:00"
  port_schedule_tuesday_to = "17:00"
  port_schedule_wednesday_active = true
  port_schedule_wednesday_from = "09:00"
  port_schedule_wednesday_to = "17:00"
}
