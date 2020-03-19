
const (
{{range $i, $item := .}}
{{$item.UpperTableName}}    = "{{$item.TableName}}" // {{$item.Comment}}{{end}}
)