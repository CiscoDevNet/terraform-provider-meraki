data "meraki_{{snakeCase .BulkName}}" "example" {
  {{- range  .Attributes}}
  {{- if and .Reference (not .Id)}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
  {{- end}}
  {{- end}}
}
