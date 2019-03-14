
//{{.TableComment}}
type {{.Table}} struct {
{{range $j, $item := .Fields}}{{$item.Name}}       {{$item.Type}}    `{{$item.FormatFields}}`           //{{$item.Remark}}
{{end}}
}
//{{.TableComment}} that may be null.
type {{.NullTable}} struct {
{{range $j, $row := .Fields}}{{$row.Name}}    {{$row.NullType}}         //{{$row.Remark}}
{{end}}
}