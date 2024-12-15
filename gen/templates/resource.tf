resource "meraki_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if and (not .ExcludeExample) (not .ExcludeTest) (not .Value) (not .Computed)}}
{{- if isNestedListSet .}}
  {{.TfName}} = [
    {
      {{- range  .Attributes}}
      {{- if and (not .ExcludeExample) (not .ExcludeTest) (not .Value) (not .Computed)}}
      {{- if isNestedListSet .}}
        {{.TfName}} = [
          {
          {{- range  .Attributes}}
          {{- if and (not .ExcludeExample) (not .ExcludeTest) (not .Value) (not .Computed)}}
          {{- if isNestedListSet .}}
            {{.TfName}} = [
              {
                {{- range  .Attributes}}
                {{- if and (not .ExcludeExample) (not .ExcludeTest) (not .Value) (not .Computed)}}
                {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
                {{- end}}
                {{- end}}
              }
            ]
          {{- else}}
          {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
          {{- end}}
          {{- end}}
          {{- end}}
          }
        ]
      {{- else}}
      {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
      {{- end}}
      {{- end}}
      {{- end}}
    }
  {{- if isNestedListSet .}}
  ]
  {{- else}}
  }
  {{- end}}
{{- else}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
{{- end}}
}
