import {
  to = meraki_{{snakeCase .BulkName}}.example
  identity = {
    {{- range (getBulkImportAttributes .)}}
    "{{.TfName}}": "<{{.TfName}}>"
    {{- end}}
  }
}