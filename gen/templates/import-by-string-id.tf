import {
  to = meraki_{{snakeCase .Name}}.example
  id = "{{range $i, $e := (importAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
}