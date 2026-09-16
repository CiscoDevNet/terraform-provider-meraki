action "meraki_{{snakeCase .Name}}" "example" {
  config {
    {{- range .Attributes}}
    {{- if and (not .ExcludeExample) (not .Value) (not .Computed) (not .DataSourceOnly) (not (isNestedListSetMap .))}}
    {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
    {{- end}}
    {{- end}}
  }
}
