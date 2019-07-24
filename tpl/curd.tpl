type {{.StructTableName}}Model struct {
DB *sql.DB
}

func New{{.StructTableName}}(db *sql.DB) *{{.StructTableName}}Model {
return &{{.StructTableName}}Model{
DB:db,
}
}

//获取所有的表字段
func (m *{{.StructTableName}}Model) getColumns() string {
return " {{.AllFieldList}} "
}

//获取多行数据.
func (m *{{.StructTableName}}Model) getRows(sqlTxt string, params ...interface{}) (rowsResult []*{{.StructTableName}}, err error) {
query, err := m.DB.Query(sqlTxt, params...)
defer query.Close()
if err != nil && err != sql.ErrNoRows {
return
}
for query.Next() {
row := {{.NullStructTableName}}{}
err = query.Scan(
{{range .NullFieldsInfo}}&row.{{.HumpName}},//{{.Comment}}
{{end}})
if nil != err && err != sql.ErrNoRows {
continue
}
rowsResult = append(rowsResult, &{{.StructTableName}}{
{{range .NullFieldsInfo}}{{if eq .GoType "float64"}}{{.HumpName}}:row.{{.HumpName}}.Float64,//{{.Comment}}
{{else if eq .GoType "int64"}}{{.HumpName}}:row.{{.HumpName}}.Int64,//{{.Comment}}
{{else if eq .GoType "time.Time"}}{{.HumpName}}:row.{{.HumpName}}.Time,//{{.Comment}}
{{else}}{{.HumpName}}:row.{{.HumpName}}.String,//{{.Comment}}
{{end}}{{end}}})
}
return
}

//获取单行数据
func (m *{{.StructTableName}}Model) getRow(sqlTxt string, params ...interface{}) (rowResult *{{.StructTableName}}, err error) {
query := m.DB.QueryRow(sqlTxt, params...)
row := {{.NullStructTableName}}{}
err = query.Scan(
{{range .NullFieldsInfo}}&row.{{.HumpName}},//{{.Comment}}
{{end}})
if nil != err && err != sql.ErrNoRows {
return
}
rowResult = &{{.StructTableName}}{
{{range .NullFieldsInfo}}{{if eq .GoType "float64"}}{{.HumpName}}:row.{{.HumpName}}.Float64 //{{.Comment}}
{{else if eq .GoType "int64"}}{{.HumpName}}:row.{{.HumpName}}.Int64,//{{.Comment}}
{{else if eq .GoType "time.Time"}}{{.HumpName}}:row.{{.HumpName}}.Time,//{{.Comment}}
{{else}}{{.HumpName}}:row.{{.HumpName}}.String,//{{.Comment}}
{{end}}{{end}}}

return
}

//_更新数据
func (m *{{.StructTableName}}Model) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
stmt, err := m.DB.Prepare(sqlTxt)
if err != nil {
return
}
defer stmt.Close()
result, err := stmt.Exec(value...)
if err != nil {
return
}
var affectCount int64
affectCount, err = result.RowsAffected()
if err != nil {
return
}
b = affectCount > 0
return
}

//新增信息
func (m *{{.StructTableName}}Model) Create(value *{{.StructTableName}}) (lastId int64, err error) {
sqlText := "INSERT INTO " + {{.UpperTableName}} + " ({{.InsertFieldList}}) VALUES ({{.InsertMark}})"
stmt, err := m.DB.Prepare(sqlText)
if err != nil {
return
}
defer stmt.Close()
result, err := stmt.Exec(
{{range .InsertInfo}}value.{{.HumpName}},//{{.Comment}}
{{end}})
if err != nil {
return
}
lastId, err = result.LastInsertId()
if err != nil {
return
}
return
}

//更新数据
func (m *{{.StructTableName}}Model) Update(value *{{.StructTableName}}) (b bool, err error) {
sqlText := "UPDATE " + {{.UpperTableName}} + " SET {{.UpdateFieldList}} WHERE {{.PrimaryKey}} = ?"
params := make([]interface{}, 0)
{{range $i, $val := .UpdateListField}}params = append(params, {{$val}})
{{end}}
return m.Save(sqlText, params...)
}

//查询多行数据
func (m *{{.StructTableName}}Model) Find(value *{{.StructTableName}}) (resList []*{{.StructTableName}}, err error) {
sqlText := "SELECT" + m.getColumns() + "FROM " + {{.UpperTableName}}
resList, err = m.getRows(sqlText)
return
}

//获取单行数据
func (m *{{.StructTableName}}Model) First(value *{{.StructTableName}}) (result *{{.StructTableName}}, err error) {
sqlText := "SELECT" + m.getColumns() + "FROM " + {{.UpperTableName}} + " LIMIT 1"
result, err = m.getRow(sqlText)
if err != nil {
return
}
return
}

//获取单行数据
func (m *{{.StructTableName}}Model) Last(value *{{.StructTableName}}) (result *{{.StructTableName}}, err error) {
sqlText := "SELECT" + m.getColumns() + "FROM " + {{.UpperTableName}} + " ORDER BY ID DESC LIMIT 1"
result, err = m.getRow(sqlText)
if err != nil {
return
}
return
}

//获取行数
func (m *{{.StructTableName}}Model) Count() (count int64, err error) {
sqlText := "SELECT COUNT(*) FROM " + {{.UpperTableName}}
query := m.DB.QueryRow(sqlText)
if err != nil {
return
}
err = query.Scan(&count)
if err != nil {
return
}
return
}
