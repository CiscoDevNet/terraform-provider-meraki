resource "meraki_{{snakeCase .BulkName}}" "example" {
{{- range getBulkParentAttributes .}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
  organization_id = "123456"
  items = [{
  {{- range getBulkItemAttributes .}}
  {{- if and (not .ExcludeExample) (not .ExcludeTest) (not .Value) (not .Computed)}}
  {{- if isNestedListSetMap .}}
    {{- if isNestedMap .}}
    {{.TfName}} = {
      "{{.MapKeyExample}}" = {
    {{- else}}
    {{.TfName}} = [
      {
    {{- end}}
        {{- range  .Attributes}}
        {{- if and (not .ExcludeExample) (not .ExcludeTest) (not .Value) (not .Computed)}}
        {{- if isNestedListSetMap .}}
          {{- if isNestedMap .}}
          {{.TfName}} = {
            "{{.MapKeyExample}}" = {
          {{- else}}
          {{.TfName}} = [
            {
          {{- end}}
            {{- range  .Attributes}}
            {{- if and (not .ExcludeExample) (not .ExcludeTest) (not .Value) (not .Computed)}}
            {{- if isNestedListSetMap .}}
              {{- if isNestedMap .}}
              {{.TfName}} = {
                "{{.MapKeyExample}}" = {
              {{- else}}
              {{.TfName}} = [
                {
              {{- end}}
                  {{- range  .Attributes}}
                  {{- if and (not .ExcludeExample) (not .ExcludeTest) (not .Value) (not .Computed)}}
                  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
                  {{- end}}
                  {{- end}}
                }
              {{- if isNestedMap .}}
              }
              {{- else}}
              ]
              {{- end}}
            {{- else}}
            {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
            {{- end}}
            {{- end}}
            {{- end}}
            }
          {{- if isNestedMap .}}
          }
          {{- else}}
          ]
          {{- end}}
        {{- else}}
        {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
        {{- end}}
        {{- end}}
        {{- end}}
      }
    {{- if isNestedMap .}}
    }
    {{- else}}
    ]
    {{- end}}
  {{- else}}
    {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
  {{- end}}
  {{- end}}
  {{- end}}
  }]
}
