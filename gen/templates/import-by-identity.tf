import {
  to = meraki_{{snakeCase .Name}}.example
  identity = {
    {{- range (importAttributes .)}}
    "{{.TfName}}": "<{{.TfName}}>"
    {{- end}}
  }
}