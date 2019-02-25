## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
{{range .TableList}}| {{.Index}}|[ {{.TableName}}](#{{.Index}}{{.TableName}})       |{{.Comment}}| |
{{end}}

{{range $key, $item := .DescList}}
### {{$item.Index}}.{{$item.TableName}}
> {{$item.Comment}}

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 备注|
|:---- |:---- |:----|:---- |:--- |:---- |
{{range $item.List}}|{{.ColumnName}}  | {{.ColumnTypeNumber}} | {{.IsNull}}| {{.DefaultValue}} | {{.PrimaryKey}} | {{.ColumnComment}}|
{{end}}
[TOP](#表总榄)
{{end}}