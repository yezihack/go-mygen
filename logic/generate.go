package logic

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/yezihack/gm2m/common"
	"github.com/yezihack/gm2m/conf"
	"github.com/yezihack/gm2m/log"
	"github.com/yezihack/gm2m/mysql"
	tpldata "github.com/yezihack/gm2m/tpl"
)

type Logic struct {
	T  *common.Tools
	DB *mysql.ModelS
}

//创建和获取MYSQL目录
func (l *Logic) GetMysqlDir() string {
	return CreateDir(common.GetExeRootDir() + "mysql/")
}

//获取根目录地址
func (l *Logic) GetRoot() string {
	return common.GetRootPath(common.GetExeRootDir()) + conf.DS
}

//创建结构体
func (l *Logic) GenerateDBStructure(tableName, tableComment, path string, tableDesc []*mysql.TableDesc) (err error) {
	//加入package mysql
	packageStr := `//数据库表内结构体信息
package mysql
` //判断package是否加载过
	//判断文件是否存在.

	if l.T.CheckFileContainsChar(path, packageStr) == false {
		l.T.WriteFile(path, packageStr)
	}
	//判断import是否加载过
	importStr := `import "database/sql"`
	if l.T.CheckFileContainsChar(path, importStr) == false {
		l.T.WriteFileAppend(path, importStr)
	}
	//声明表结构变量
	TableData := new(mysql.TableInfo)
	TableData.Table = l.T.Capitalize(tableName)
	TableData.NullTable = conf.DbNullPrefix + TableData.Table
	TableData.TableComment = tableComment
	//判断表结构是否加载过
	if l.T.CheckFileContainsChar(path, "type "+TableData.Table+" struct") == true {
		return
	}
	//加载模板文件
	tplByte, err := tpldata.Asset(conf.TPL_STRUCTURE)
	if err != nil {
		return
	}
	tpl, err := template.New("structure").Parse(string(tplByte))
	if err != nil {
		log.Error("ParseFiles", err)
		return
	}
	//装载表字段信息
	fts := []string{"json"}
	if err != nil {
		log.Error("GetConfFormat", err)
		return
	}
	//判断是否含json
	if !common.InArrayString("json", fts) {
		index0 := fts[0]
		fts[0] = "json"
		fts = append(fts, index0)
	}
	for _, val := range tableDesc {
		TableData.Fields = append(TableData.Fields, &mysql.FieldsInfo{
			Name:         l.T.Capitalize(val.ColumnName),
			Type:         val.GolangType,
			NullType:     val.MysqlNullType,
			DbOriField:   val.ColumnName,
			FormatFields: common.FormatField(val.ColumnName, fts),
			Remark:       val.ColumnComment,
		})
	}
	content := bytes.NewBuffer([]byte{})
	tpl.Execute(content, TableData)
	//表信息写入文件
	err = WriteAppendFile(path, content.String())
	if err != nil {
		return
	}
	return
}

//生成C增,R查,U删,D改,的文件
func (l *Logic) GenerateCURDFile(tableName string, tableDesc []*mysql.TableDesc) (err error) {
	allFields := make([]string, 0)
	insertFields := make([]string, 0)
	InsertInfo := make([]*mysql.SqlFieldInfo, 0)
	fieldsList := make([]*mysql.SqlFieldInfo, 0)
	nullFieldList := make([]*mysql.NullSqlFieldInfo, 0)
	updateList := make([]string, 0)
	updateListField := make([]string, 0)
	PrimaryKey, primaryType := "", ""
	for _, item := range tableDesc {
		allFields = append(allFields, "`"+item.ColumnName+"`")
		if item.PrimaryKey == false && item.ColumnName != "updated_at" && item.ColumnName != "created_at" {
			insertFields = append(insertFields, item.ColumnName)
			InsertInfo = append(InsertInfo, &mysql.SqlFieldInfo{
				HumpName: l.T.Capitalize(item.ColumnName),
				Comment:  item.ColumnComment,
			})
			if item.ColumnName == "identify" {
				updateList = append(updateList, item.ColumnName+"="+item.ColumnName+"+1")
			} else {
				updateList = append(updateList, item.ColumnName+"=?")
				updateListField = append(updateListField, "t."+l.T.Capitalize(item.ColumnName))
			}
		}
		if item.PrimaryKey {
			PrimaryKey = item.ColumnName
			primaryType = item.GolangType
		}
		fieldsList = append(fieldsList, &mysql.SqlFieldInfo{
			HumpName: l.T.Capitalize(item.ColumnName),
			Comment:  item.ColumnComment,
		})
		nullFieldList = append(nullFieldList, &mysql.NullSqlFieldInfo{
			HumpName:     l.T.Capitalize(item.ColumnName),
			OriFieldType: conf.MysqlTypeToGoType[item.OriMysqlType],
			Comment:      item.ColumnComment,
		})
	}
	//拼出SQL所需要结构数据
	InsertMark := strings.Repeat("?,", len(insertFields))
	sqlInfo := &mysql.SqlInfo{
		TableName:           tableName,
		PrimaryKey:          PrimaryKey,
		PrimaryType:         primaryType,
		StructTableName:     l.T.Capitalize(tableName),
		NullStructTableName: conf.DbNullPrefix + l.T.Capitalize(tableName),
		UpperTableName:      l.T.ToUpper(tableName),
		AllFieldList:        strings.Join(allFields, ","),
		InsertFieldList:     strings.Join(insertFields, ","),
		InsertMark:          InsertMark[:len(InsertMark)-1],
		UpdateFieldList:     strings.Join(updateList, ","),
		UpdateListField:     updateListField,
		FieldsInfo:          fieldsList,
		NullFieldsInfo:      nullFieldList,
		InsertInfo:          InsertInfo,
	}
	err = l.GenerateSQL(sqlInfo)
	if err != nil {
		return
	}
	return
}

//生成表列表
func (l *Logic) GenerateTableList(list []*mysql.TableList) (err error) {
	//写入表名
	tableListFile := l.GetMysqlDir() + "table_list.go"
	//判断package是否加载过
	checkStr := "package mysql"
	if l.T.CheckFileContainsChar(tableListFile, checkStr) == false {
		l.T.WriteFile(tableListFile, checkStr)
	}
	checkStr = "const"
	if l.T.CheckFileContainsChar(tableListFile, checkStr) {
		return
	}

	tplByte, err := tpldata.Asset(conf.TPL_TABLES)
	if err != nil {
		return
	}
	tpl, err := template.New("table_list").Parse(string(tplByte))
	if err != nil {
		return
	}
	//解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, list)
	if err != nil {
		return
	}
	//表信息写入文件
	err = WriteAppendFile(tableListFile, content.String())
	if err != nil {
		return
	}
	return
}

//生成SQL文件
func (l *Logic) GenerateSQL(info *mysql.SqlInfo) (err error) {
	//写入表名
	goFile := l.GetMysqlDir() + info.TableName + ".go"
	//判断package是否加载过
	checkStr := "package mysql"
	if l.T.CheckFileContainsChar(goFile, checkStr) == false {
		l.T.WriteFile(goFile, checkStr)
	}

	//解析模板
	tplByte, err := tpldata.Asset(conf.TPL_CRUD)
	if err != nil {
		return
	}
	tpl, err := template.New("CURD").Parse(string(tplByte))
	if err != nil {
		return
	}
	//解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, info)
	if err != nil {
		return
	}
	//表信息写入文件
	if l.T.CheckFileContainsChar(goFile, info.StructTableName) == false {
		err = WriteAppendFile(goFile, content.String())
		if err != nil {
			return
		}
	}
	return
}

//生成表列表
func (l *Logic) GenerateMarkdown(data *mysql.MarkDownData) (err error) {
	//写入markdown
	file := common.GetExeRootDir() + "markdown.md"
	tplByte, err := tpldata.Asset(conf.TPL_MARKDOWN)
	if err != nil {
		return
	}
	//解析
	content := bytes.NewBuffer([]byte{})
	tpl, err := template.New("markdown").Parse(string(tplByte))
	err = tpl.Execute(content, data)
	if err != nil {
		return
	}
	//表信息写入文件
	err = WriteAppendFile(file, content.String())
	if err != nil {
		return
	}
	return
}
