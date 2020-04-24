type {{.StructTableName}}Model struct {
E
DB *sql.DB
Tx *sql.Tx
}

// not transaction
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
// transaction object
func New{{.StructTableName}}Tx(tx *sql.Tx) *{{.StructTableName}}Model {
	return &{{.StructTableName}}Model{
		Tx: tx,
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
if err != nil {
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
if err != nil {
err = m.E.Stack(err)
return
}
rowResult = &{{.PkgEntity}}{{.StructTableName}}{
{{range .NullFieldsInfo}}{{if eq .GoType "float64"}}{{.HumpName}}:row.{{.HumpName}}.Float64, // {{.Comment}}
{{else if eq .GoType "int64"}}{{.HumpName}}:row.{{.HumpName}}.Int64,// {{.Comment}}
{{else if eq .GoType "time.Time"}}{{.HumpName}}:row.{{.HumpName}}.Time,// {{.Comment}}
{{else if eq .GoType "int32"}}{{.HumpName}}:row.{{.HumpName}}.Int32,// {{.Comment}}
{{else}}{{.HumpName}}:row.{{.HumpName}}.String,// {{.Comment}}
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


// _更新数据
func (m *{{.StructTableName}}Model) SaveTx(sqlTxt string, value ...interface{}) (b bool, err error) {
stmt, err := m.Tx.Prepare(sqlTxt)
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

// 新增信息 tx
func (m *{{.StructTableName}}Model) CreateTx(value *{{.PkgEntity}}{{.StructTableName}}) (lastId int64, err error) {
const sqlText = "INSERT INTO " + {{.PkgTable}}{{.UpperTableName}} + " ({{.InsertFieldList}}) VALUES ({{.InsertMark}})"
stmt, err := m.Tx.Prepare(sqlText)
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
const sqlText = "UPDATE " + {{.PkgTable}}{{.UpperTableName}} + " SET {{.UpdateFieldList}} WHERE {{.PrimaryKey}} = ?"
params := make([]interface{}, 0)
{{range $i, $val := .UpdateListField}}params = append(params, {{$val}})
{{end}}
return m.Save(sqlText, params...)
}


// 更新数据 tx
func (m *{{.StructTableName}}Model) UpdateTx(value *{{.PkgEntity}}{{.StructTableName}}) (b bool, err error) {
const sqlText = "UPDATE " + {{.PkgTable}}{{.UpperTableName}} + " SET {{.UpdateFieldList}} WHERE {{.PrimaryKey}} = ?"
params := make([]interface{}, 0)
{{range $i, $val := .UpdateListField}}params = append(params, {{$val}})
{{end}}
return m.SaveTx(sqlText, params...)
}


// 查询多行数据
func (m *{{.StructTableName}}Model) Find() (resList []*{{.PkgEntity}}{{.StructTableName}}, err error) {
sqlText := "SELECT" + m.getColumns() + "FROM " + {{.PkgTable}}{{.UpperTableName}}
resList, err = m.getRows(sqlText)
return
}

// 获取单行数据
func (m *{{.StructTableName}}Model) First(id int64) (result *{{.PkgEntity}}{{.StructTableName}}, err error) {
sqlText := "SELECT" + m.getColumns() + "FROM " + {{.PkgTable}}{{.UpperTableName}} + " WHERE {{.PrimaryKey}} = ? LIMIT 1"
result, err = m.getRow(sqlText, id)
return
}

// 获取最后一行数据
func (m *{{.StructTableName}}Model) Last() (result *{{.PkgEntity}}{{.StructTableName}}, err error) {
sqlText := "SELECT" + m.getColumns() + "FROM " + {{.PkgTable}}{{.UpperTableName}} + " ORDER BY ID DESC LIMIT 1"
result, err = m.getRow(sqlText)
return
}

// 单列数据
func (m *{{.StructTableName}}Model) Pluck(id int64) (result map[int64]interface{}, err error) {
	const sqlText = "SELECT {{.PrimaryKey}}, {{.SecondField}} FROM " + {{.PkgTable}}{{.UpperTableName}} + " where {{.PrimaryKey}} = ?"
	rows, err := m.DB.Query(sqlText, id)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer rows.Close()
	result = make(map[int64]interface{})
	var (
	    _id int64
	    _val interface{}
	)
	for rows.Next() {
		err = rows.Scan(&_id, &_val)
		if err != nil {
		    err = m.E.Stack(err)
			return
		}
		result[_id] = _val
	}
	return
}

// 单列数据 by 支持切片传入
// Get column data
func (m *{{.StructTableName}}Model) Plucks(ids []int64) (result map[int64]interface{}, err error) {
    result = make(map[int64]interface{})
    if len(ids) == 0 {
        return
    }
	sqlText := "SELECT {{.PrimaryKey}}, {{.SecondField}} FROM " + {{.PkgTable}}{{.UpperTableName}} + " where " +
		"{{.PrimaryKey}} in (" + RepeatQuestionMark(len(ids)) + ")"
	params := make([]interface{}, len(ids))
	for idx, id := range ids {
		params[idx] = id
	}
	rows, err := m.DB.Query(sqlText, params...)
	if err != nil {
		err = m.E.Stack(err)
		return
	}
	defer rows.Close()
	var (
		_id int64
		_val interface{}
	)
	for rows.Next() {
		err = rows.Scan(&_id, &_val)
		if err != nil {
			err = m.E.Stack(err)
			return
		}
		result[_id] = _val
	}
	return
}

// 获取单个数据
// Get one data
func (m *{{.StructTableName}}Model) One(id int64) (result int64, err error) {
	sqlText := "SELECT {{.PrimaryKey}} FROM " + {{.PkgTable}}{{.UpperTableName}} + " where {{.PrimaryKey}}=?"
	err = m.DB.QueryRow(sqlText, id).Scan(&result)
	if err != nil && err != sql.ErrNoRows {
	    err = m.E.Stack(err)
		return
	}
	return
}

// 获取行数
// Get line count
func (m *{{.StructTableName}}Model) Count() (count int64, err error) {
    sqlText := "SELECT COUNT(*) FROM " + {{.PkgTable}}{{.UpperTableName}}
    err = m.DB.QueryRow(sqlText).Scan(&count)
    if err != nil && err != sql.ErrNoRows{
        err = m.E.Stack(err)
        return
    }
    return
}

// 判断数据是否存在
// Check the data is have?
func (m *{{.StructTableName}}Model) Has(id int64) (b bool, err error) {
	sqlText := "SELECT {{.PrimaryKey}} FROM " + {{.PkgTable}}{{.UpperTableName}} + " where {{.PrimaryKey}} = ?"
	var count int64
    err = m.DB.QueryRow(sqlText, id).Scan(&count)
    if err != nil && err != sql.ErrNoRows {
        err = m.E.Stack(err)
        return
    }
    return count > 0, nil
}