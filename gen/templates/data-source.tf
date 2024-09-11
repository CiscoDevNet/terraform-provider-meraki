{{- $name := .Name -}}
data "meraki_{{snakeCase .Name}}" "example" {
  id = "{{$id := false}}{{range .Attributes}}{{if .Id}}{{$id = true}}{{.Example}}{{end}}{{end}}{{if not $id}}12345678{{end}}"
  {{- range  .Attributes}}
  {{- if .Reference}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
  {{- end}}
  {{- end}}
}
