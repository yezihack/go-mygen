type {{.StructTableName}}Model struct {
E
DB *sql.DB
}

func New{{.StructTableName}}(db ...*sql.DB) *{{.StructTableName}}Model {
if len(db) > 0 {
    return &{{.StructTableName}}Model{
        DB: db[0],
    }
}
return &{{.StructTableName}}Model{
    DB: masterDB,
}
}

// 获取所有的表字段
func (m *{{.StructTableName}}Model) getColumns() string {
return " {{.AllFieldList}} "
}

// 获取多行数据.
func (m *{{.StructTableName}}Model) getRows(sqlTxt string, params ...interface{}) (rowsResult []*{{.PkgEntity}}{{.StructTableName}}, err error) {
query, err := m.DB.Query(sqlTxt, params...)
if err != nil {
err = m.E.Stack(err)
return
}
defer query.Close()
for query.Next() {
row := {{.PkgEntity}}{{.NullStructTableName}}{}
err = query.Scan(
{{range .NullFieldsInfo}}&row.{{.HumpName}},// {{.Comment}}
{{end}})
if err != nil &&  err != sql.ErrNoRows {
    err = m.E.Stack(err)
    return
}
rowsResult = append(rowsResult, &{{.PkgEntity}}{{.StructTableName}}{
{{range .NullFieldsInfo}}{{if eq .GoType "float64"}}{{.HumpName}}:row.{{.HumpName}}.Float64,// {{.Comment}}
{{else if eq .GoType "int64"}}{{.HumpName}}:row.{{.HumpName}}.Int64,// {{.Comment}}
{{else if eq .GoType "time.Time"}}{{.HumpName}}:row.{{.HumpName}}.Time,// {{.Comment}}
{{else if eq .GoType "int32"}}{{.HumpName}}:row.{{.HumpName}}.Int32,// {{.Comment}}
{{else}}{{.HumpName}}:row.{{.HumpName}}.String,// {{.Comment}}
{{end}}{{end}}})
}
return
}

// 获取单行数据
func (m *{{.StructTableName}}Model) getRow(sqlText string, params ...interface{}) (rowResult *{{.PkgEntity}}{{.StructTableName}}, err error) {
query := m.DB.QueryRow(sqlText, params...)
row := {{.PkgEntity}}{{.NullStructTableName}}{}
err = query.Scan(
{{range .NullFieldsInfo}}&row.{{.HumpName}},// {{.Comment}}
{{end}})
if err != sql.ErrNoRows {
err = m.E.Stack(err)
return
}
rowResult = &{{.PkgEntity}}{{.StructTableName}}{
{{range .NullFieldsInfo}}{{if eq .GoType "float64"}}{{.HumpName}}:row.{{.HumpName}}.Float64, // {{.Comment}}
{{else if eq .GoType "int64"}}{{.HumpName}}:row.{{.HumpName}}.Int64,// {{.Comment}}
{{else if eq .GoType "time.Time"}}{{.HumpName}}:row.{{.HumpName}}.Time,// {{.Comment}}
{{else if eq .GoType "int32"}}{{.HumpName}}:row.{{.HumpName}}.Int32,// {{.Comment}}
{{else}}{{.HumpName}}:row.{{.HumpName}}.String,//{{.Comment}}
{{end}}{{end}}}
return
}

// _更新数据
func (m *{{.StructTableName}}Model) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
stmt, err := m.DB.Prepare(sqlTxt)
if err != nil {
err = m.E.Stack(err)
return
}
defer stmt.Close()
result, err := stmt.Exec(value...)
if err != nil {
err = m.E.Stack(err)
return
}
var affectCount int64
affectCount, err = result.RowsAffected()
if err != nil {
err = m.E.Stack(err)
return
}
b = affectCount > 0
return
}

// 新增信息
func (m *{{.StructTableName}}Model) Create(value *{{.PkgEntity}}{{.StructTableName}}) (lastId int64, err error) {
const sqlText = "INSERT INTO " + {{.PkgTable}}{{.UpperTableName}} + " ({{.InsertFieldList}}) VALUES ({{.InsertMark}})"
stmt, err := m.DB.Prepare(sqlText)
if err != nil {
err = m.E.Stack(err)
return
}
defer stmt.Close()
result, err := stmt.Exec(
{{range .InsertInfo}}value.{{.HumpName}},// {{.Comment}}
{{end}})
if err != nil {
err = m.E.Stack(err)
return
}
lastId, err = result.LastInsertId()
if err != nil {
err = m.E.Stack(err)
return
}
return
}

// 更新数据
func (m *{{.StructTableName}}Model) Update(value *{{.PkgEntity}}{{.StructTableName}}) (b bool, err error) {
 sqlText := "UPDATE " + {{.PkgTable}}{{.UpperTableName}} + " SET {{.UpdateFieldList}} WHERE {{.PrimaryKey}} = ?"
params := make([]interface{}, 0)
{{range $i, $val := .UpdateListField}}params = append(params, {{$val}})
{{end}}
return m.Save(sqlText, params...)
}

// 查询多行数据
func (m *{{.StructTableName}}Model) Find(value *{{.PkgEntity}}{{.StructTableName}}) (resList []*{{.PkgEntity}}{{.StructTableName}}, err error) {
 sqlText := "SELECT" + m.getColumns() + "FROM " + {{.PkgTable}}{{.UpperTableName}}
resList, err = m.getRows(sqlText)
return
}

// 获取单行数据
func (m *{{.StructTableName}}Model) First(value *{{.PkgEntity}}{{.StructTableName}}) (result *{{.PkgEntity}}{{.StructTableName}}, err error) {
 sqlText := "SELECT" + m.getColumns() + "FROM " + {{.PkgTable}}{{.UpperTableName}} + " LIMIT 1"
result, err = m.getRow(sqlText)
return
}

// 获取最后一行数据
func (m *{{.StructTableName}}Model) Last(value *{{.PkgEntity}}{{.StructTableName}}) (result *{{.PkgEntity}}{{.StructTableName}}, err error) {
 sqlText := "SELECT" + m.getColumns() + "FROM " + {{.PkgTable}}{{.UpperTableName}} + " ORDER BY ID DESC LIMIT 1"
result, err = m.getRow(sqlText)
return
}

// 获取单个数据
func (m *{{.StructTableName}}Model) One(userId int64) (result int64, err error) {
	sqlText := "SELECT id FROM " + {{.PkgTable}}{{.UpperTableName}} + " where id=?"
	rows := m.DB.QueryRow(sqlText, userId)
	if err = rows.Scan(&result); err != nil {
	    err = m.E.Stack(err)
		return
	}
	return
}

// 获取行数
func (m *{{.StructTableName}}Model) Count() (count int64, err error) {
 sqlText := "SELECT COUNT(*) FROM " + {{.PkgTable}}{{.UpperTableName}}
query := m.DB.QueryRow(sqlText)
err = query.Scan(&count)
if err != nil {
err = m.E.Stack(err)
return
}
return
}

// 判断是否存在
func (m *{{.StructTableName}}Model) Exists(id int64) (b bool, err error) {
	sqlText := "SELECT COUNT(*) FROM " + {{.PkgTable}}{{.UpperTableName}} + " where id = ?"
	query := m.DB.QueryRow(sqlText, id)
	var count int64
	err = query.Scan(&count)
	if err != nil {
	    err = m.E.Stack(err)
		return
	}
	if count > 0 {
		b = true
		return
	}
	return
}