//{{.TableComment}}
type {{.Table}} struct {
{{range $j, $item := .Fields}}{{$item.Name}}       {{$item.Type}}    `json:"{{$item.DbName}}"`           //{{$item.Remark}}
{{end}}
}
type {{.NullTable}} struct {
{{range $j, $row := .Fields}}{{$row.Name}}    {{$row.NullType}}         //{{$row.Remark}}
{{end}}
}