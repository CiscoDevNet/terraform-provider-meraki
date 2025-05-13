terraform import meraki_{{snakeCase .Name}}.example "{{range $i, $e := (importAttributes .)}}{{if $i}},{{end}}<{{.TfName}}>{{end}}"
