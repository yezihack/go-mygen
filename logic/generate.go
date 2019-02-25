package logic

import (
	"bytes"
	"strings"

	"github.com/ThreeKing2018/k3log"
	"github.com/alecthomas/template"
	"github.com/yezihack/m2m/common"
	"github.com/yezihack/m2m/conf"
	"github.com/yezihack/m2m/mysql"
)

type Logic struct {
	T  *common.Tools
	db mysql.DbTools
}

//创建和获取MYSQL目录
func (l *Logic) GetMysqlDir() string {
	return CreateDir(common.GetRootDir() + "mysql/")
}

//获取根目录地址
func (l *Logic) GetRoot() string {
	return common.GetRootPath(common.GetRootDir()) + conf.DS
}

//创建结构体
func (l *Logic) GenerateDBStructure(tableName, tableComment string, tableDesc []*mysql.TableDesc) (err error) {
	//表结构文件路径
	tableInfoFile := l.GetMysqlDir() + conf.GOFILE_STRUCTURE
	//表结构模板文件路径
	tplFile := l.GetRoot() + conf.TPL_STRUCTURE
	//加入package mysql
	packageStr := `//数据库表内结构体信息
package mysql
` //判断package是否加载过
	if l.T.CheckFileContainsChar(tableInfoFile, packageStr) == false {
		l.T.WriteFile(tableInfoFile, packageStr)
	}
	//判断import是否加载过
	importStr := `import "database/sql"`
	if l.T.CheckFileContainsChar(tableInfoFile, importStr) == false {
		l.T.WriteFileAppend(tableInfoFile, importStr)
	}
	//声明表结构变量
	TableData := new(mysql.TableInfo)
	TableData.Table = l.T.Capitalize(tableName)
	TableData.NullTable = conf.DbNullPrefix + TableData.Table
	TableData.TableComment = tableComment
	//判断表结构是否加载过
	if l.T.CheckFileContainsChar(tableInfoFile, "type "+TableData.Table+" struct") == true {
		return
	}
	//加载模板文件
	tpl, err := template.ParseFiles(tplFile)
	if err != nil {
		k3log.Error("ParseFiles", err)
		return
	}
	//装载表信息
	for _, val := range tableDesc {
		TableData.Fields = append(TableData.Fields, &mysql.FieldsInfo{
			Name:     l.T.Capitalize(val.ColumnName),
			Type:     val.GolangType,
			NullType: val.MysqlNullType,
			DbName:   val.ColumnName,
			Remark:   val.ColumnComment,
		})
	}
	content := bytes.NewBuffer([]byte{})
	tpl.Execute(content, TableData)
	//表信息写入文件
	err = WriteAppendFile(tableInfoFile, content.String())
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
	tplFile := l.GetRoot() + conf.TPL_TABLES

	//判断package是否加载过
	checkStr := "package mysql"
	if l.T.CheckFileContainsChar(tableListFile, checkStr) == false {
		l.T.WriteFile(tableListFile, checkStr)
	}
	checkStr = "const"
	if l.T.CheckFileContainsChar(tableListFile, checkStr) {
		return
	}
	tpl, err := template.ParseFiles(tplFile)
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
	tplFile := l.GetRoot() + conf.TPL_CRUD
	//判断package是否加载过
	checkStr := "package mysql"
	if l.T.CheckFileContainsChar(goFile, checkStr) == false {
		l.T.WriteFile(goFile, checkStr)
	}

	//解析模板
	tpl, err := template.ParseFiles(tplFile)
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
	file := common.GetRootDir() + "markdown.md"
	tplFile := l.GetRoot() + conf.TPL_MARKDOWN

	tpl, err := template.ParseFiles(tplFile)
	if err != nil {
		return
	}
	//解析
	content := bytes.NewBuffer([]byte{})
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
