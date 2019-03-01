
//打印错误日志
func (t *{{.StructTableName}}) errWrite(err error, sql string, addParams ...interface{}) {
type ErrorType struct {
{{.StructTableName}}
AddParams []interface{} `json:"add_params"`
}
var infoList []interface{}
for k, _ := range addParams {
infoList = append(infoList, addParams[k])
errType := ErrorType{
{{.StructTableName}}: *t,
AddParams:       infoList,
}
paramsJson, _ := json.Marshal(errType)
k3log.Error("msg", "dbError", "tableName", {{.UpperTableName}}, "error", err.Error(), "sql", sql, "params", string(paramsJson))
}
}
//获取表的所有字段
func (t *{{.StructTableName}}) selectColumn() string {
return " {{.AllFieldList}} "
}
//查询数据,基础函数
func (t *{{.StructTableName}}) _selectBody(db *sql.DB, sqlText string, params []interface{}) (_bodyArr []*{{.StructTableName}}, err error) {
_bodyArr = make([]*{{.StructTableName}}, 0)
rows, err := db.Query(sqlText, params...)
if nil != err {
t.errWrite(err, sqlText, params)
return
}
defer rows.Close()
for rows.Next() {
_one := {{.NullStructTableName}}{}
err = rows.Scan(
{{range .NullFieldsInfo}}&_one.{{.HumpName}},//{{.Comment}}
{{end}})
if nil != err {
t.errWrite(err, sqlText, params)
continue
}
_bodyArr = append(_bodyArr, &{{.StructTableName}}{
{{range .NullFieldsInfo}}{{if eq .OriFieldType "int"}}{{.HumpName}}:int(_one.{{.HumpName}}.Int64),//{{.Comment}}
{{else if eq .OriFieldType "float64"}}{{.HumpName}}:int(_one.{{.HumpName}}.Float64),//{{.Comment}}
{{else if eq .OriFieldType "int64"}}{{.HumpName}}:_one.{{.HumpName}}.Int64,//{{.Comment}}
{{else}}{{.HumpName}}:_one.{{.HumpName}}.String,//{{.Comment}}
{{end}}{{end}}})
}
return
}
//更新数据基础函数
func (t *{{.StructTableName}}) _updateBody(db *sql.DB, sqlText string, params []interface{}) (b bool, err error) {
stmt, err := db.Prepare(sqlText)
if nil != err {
t.errWrite(err, sqlText, params)
return
}
defer stmt.Close()
res, err := stmt.Exec(
params...,
)
_count, err := res.RowsAffected()
if nil != err {
t.errWrite(err, sqlText, params)
return
}
return _count > 0, nil
}
//插入数据
func (t *{{.StructTableName}}) Insert(db *sql.DB) (b bool, err error) {
const sqlText = "INSERT INTO " + {{.UpperTableName}} + " ({{.InsertFieldList}}) VALUES ({{.InsertMark}})"
stmt, err := db.Prepare(sqlText)
if nil != err {
t.errWrite(err, sqlText)
return
}
defer stmt.Close()
res, err := stmt.Exec(
{{range .InsertInfo}}t.{{.HumpName}},//{{.Comment}}
{{end}})
if nil != err {
t.errWrite(err, sqlText)
return
}
n, err := res.RowsAffected()
if nil != err {
t.errWrite(err, sqlText)
return
}
return n > 0, nil
}
//查询单列数据
func (t *{{.StructTableName}}) First(db *sql.DB, id int64) (result *{{.StructTableName}}, err error) {
sqlText := "SELECT" + t.selectColumn() + "FROM " + {{.UpperTableName}} + " WHERE {{.PrimaryKey}} = ? limit 1"
list, err := t._selectBody(db, sqlText, []interface{}{id})
if err != nil {
return
}
if len(list) > 0 {
result = list[0]
}
return
}
//查询多条数据,传入id,id2,id3
func (t *{{.StructTableName}}) Find(db *sql.DB, ids ...{{.PrimaryType}}) (resultList []*{{.StructTableName}}, err error) {
resultList = make([]*{{.StructTableName}}, 0)
sqlText := "SELECT" + t.selectColumn() + "FROM " + {{.UpperTableName}} + " WHERE {{.PrimaryKey}} in ("+strings.TrimRight(strings.Repeat("?,", len(ids)), ",")+")"
var args []interface{}
for _, id := range ids {
args = append(args, id)
}
resultList, err = t._selectBody(db, sqlText, args)
if err != nil {
return
}
return
}
//主键更新数据
func (t *{{.StructTableName}}) Update(db *sql.DB, id {{.PrimaryType}}) (err error) {
sqlText := "UPDATE " + {{.UpperTableName}} + " SET {{.UpdateFieldList}} WHERE {{.PrimaryKey}} = ?"
params := make([]interface{}, 0)
{{range $i, $val := .UpdateListField}}params = append(params, {{$val}})
{{end}}
//主键{{.PrimaryKey}}
params = append(params, id)
_, err = t._updateBody(db, sqlText, params)
if err != nil {
return
}
return
}