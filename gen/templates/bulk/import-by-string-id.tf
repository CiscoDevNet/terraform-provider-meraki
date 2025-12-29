import {
  to = meraki_{{snakeCase .BulkName}}.example
  id = "{{range $i, $e := (getBulkImportAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
}