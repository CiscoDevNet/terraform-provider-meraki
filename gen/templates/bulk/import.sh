terraform import meraki_{{snakeCase .Name}}.example "{{range $i, $e := (getBulkImportAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
