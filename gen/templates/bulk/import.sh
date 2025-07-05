terraform import meraki_{{snakeCase .BulkName}}.example "{{range $i, $e := (getBulkImportAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
