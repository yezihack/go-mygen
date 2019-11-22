package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"strings"
	"sync"
	"time"
)

type Logic struct {
	T    *Tools
	DB   *ModelS
	Path string
	Once sync.Once
	l    sync.Mutex
}

//生成结构实体文件
func (l *Logic) CreateEntity(formatList []string) (err error) {
	//表结构文件路径
	path := l.Path + "db_entity/"
	if l.T.IsDirOrFileExist(path) == false {
		if !l.T.CreateDir(path) {
			return errors.New("Build directory failed, path>>" + path)
		}
	}
	path += GOFILE_ENTITY
	if l.T.IsDirOrFileExist(path) == false {
		if !l.T.CreateFile(path) {
			err = errors.New("Create file fails, path>>" + path)
			return
		}
	}
	//将表结构写入文件
	for idx, table := range l.DB.DoTables {
		idx++
		//查询表结构信息
		tableDesc, err := l.DB.GetTableDesc(table.Name)
		if err != nil {
			log.Fatal("CreateEntityErr>>", err)
			continue
		}
		req := new(EntityReq)
		req.Index = idx
		req.Path = path
		req.TableName = table.Name
		req.TableComment = table.Comment
		req.TableDesc = tableDesc
		req.FormatList = formatList
		req.Pkg = PkgEntity
		//生成基础信息
		err = l.GenerateDBEntity(req)
		if err != nil {
			log.Fatal("CreateEntityErr>>", err)
			continue
		}
	}
	return
}

//生成原生的crud查询数据库
func (l *Logic) CreateCURD(formatList []string) (err error) {
	tableNameList := make([]*TableList, 0)
	//表结构文件路径
	structPath := l.GetMysqlDir() + GOFILE_ENTITY
	//将表结构写入文件
	for _, table := range l.DB.DoTables {
		tableNameList = append(tableNameList, &TableList{
			UpperTableName: TablePrefix + l.T.ToUpper(table.Name),
			TableName:      table.Name,
			Comment:        table.Comment,
		})
		//查询表结构信息
		tableDesc, err := l.DB.GetTableDesc(table.Name)
		if err != nil {
			return err
		}
		req := new(EntityReq)
		req.TableName = table.Name
		req.TableComment = table.Comment
		req.TableDesc = tableDesc
		req.Path = structPath
		req.FormatList = formatList
		req.Pkg = PkgDbModels
		//生成基础信息
		err = l.GenerateDBEntity(req)
		if err != nil {
			return err
		}
		//生成增,删,改,查文件
		err = l.GenerateCURDFile(table.Name, table.Comment, tableDesc)
		if err != nil {
			return err
		}
	}

	//生成所有表的文件
	err = l.GenerateTableList(tableNameList)
	if err != nil {
		return err
	}
	fmt.Println("`CURD` files created finish!")
	return nil
}

//生成mysql markdown文档
func (l *Logic) CreateMarkdown() (err error) {
	data := new(MarkDownData)
	//将表结构写入文件
	i := 1
	for _, table := range l.DB.DoTables {
		fmt.Println("Doing table:" + table.Name)
		data.TableList = append(data.TableList, &TableList{
			Index:          i,
			UpperTableName: l.T.ToUpper(table.Name),
			TableName:      table.Name,
			Comment:        table.Comment,
		})
		//查询表结构信息
		desc := new(MarkDownDataChild)
		desc.List, err = l.DB.GetTableDesc(table.Name)
		if err != nil {
			log.Fatal("markdown>>", err)
			continue
		}
		desc.Index = i
		desc.TableName = table.Name
		desc.Comment = table.Comment
		data.DescList = append(data.DescList, desc)
		i++
	}
	//生成所有表的文件
	err = l.GenerateMarkdown(data)
	if err != nil {
		return
	}
	return
}

//创建和获取MYSQL目录
func (l *Logic) GetMysqlDir() string {
	return CreateDir(l.Path + GODIR_MODELS + DS)
}

//获取根目录地址
func (l *Logic) GetRoot() string {
	return GetRootPath(l.Path) + DS
}

//创建结构体
func (l *Logic) GenerateDBStructure(tableName, tableComment, path string, tableDesc []*TableDesc) (err error) {
	//加入package
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
	TableData := new(TableInfo)
	TableData.Table = l.T.Capitalize(tableName)
	TableData.NullTable = DbNullPrefix + TableData.Table
	TableData.TableComment = tableComment
	//判断表结构是否加载过
	if l.T.CheckFileContainsChar(path, "type "+TableData.Table+" struct") == true {
		return
	}
	//加载模板文件
	tplByte, err := Asset(TPL_STRUCTURE)
	if err != nil {
		return
	}
	tpl, err := template.New("structure").Parse(string(tplByte))
	if err != nil {
		return
	}
	//装载表字段信息
	fts := []string{"json"}
	if err != nil {
		return
	}
	//判断是否含json
	if !InArrayString("json", fts) {
		index0 := fts[0]
		fts[0] = "json"
		fts = append(fts, index0)
	}
	for _, val := range tableDesc {
		TableData.Fields = append(TableData.Fields, &FieldsInfo{
			Name:         l.T.Capitalize(val.ColumnName),
			Type:         val.GolangType,
			NullType:     val.MysqlNullType,
			DbOriField:   val.ColumnName,
			FormatFields: FormatField(val.ColumnName, fts),
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

//创建结构实体
func (l *Logic) GenerateDBEntity(req *EntityReq) (err error) {
	l.l.Lock()
	defer l.l.Unlock()
	var s string
	s = fmt.Sprintf(`//判断package是否加载过
package %s
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
`, req.Pkg)
	//判断import是否加载过
	check := "github.com/go-sql-driver/mysql"
	if l.T.CheckFileContainsChar(req.Path, check) == false {
		l.T.WriteFile(req.Path, s)
	}
	//声明表结构变量
	TableData := new(TableInfo)
	TableData.Table = l.T.Capitalize(req.TableName)
	TableData.NullTable = TableData.Table + DbNullPrefix
	TableData.TableComment = AddToComment(req.TableComment, "")
	TableData.TableCommentNull = AddToComment(req.TableComment, " Null Entity")
	//判断表结构是否加载过
	if l.T.CheckFileContainsChar(req.Path, "type "+TableData.Table+" struct") == true {
		log.Println(req.Path + "It already exists. Please delete it and regenerate it")
		return
	}
	//加载模板文件
	tplByte, err := Asset(TPL_ENTITY)
	if err != nil {
		return
	}
	tpl, err := template.New("entity").Parse(string(tplByte))
	if err != nil {
		return
	}
	//装载表字段信息
	for _, val := range req.TableDesc {
		TableData.Fields = append(TableData.Fields, &FieldsInfo{
			Name:         l.T.Capitalize(val.ColumnName),
			Type:         val.GolangType,
			NullType:     val.MysqlNullType,
			DbOriField:   val.ColumnName,
			FormatFields: FormatField(val.ColumnName, req.FormatList),
			Remark:       AddToComment(val.ColumnComment, ""),
		})
	}
	content := bytes.NewBuffer([]byte{})
	tpl.Execute(content, TableData)
	//表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = WriteAppendFile(req.Path, con)
	if err != nil {
		return
	}
	return
}

//生成C增,U删,R查,D改,的文件
func (l *Logic) GenerateCURDFile(tableName, tableComment string, tableDesc []*TableDesc) (err error) {
	var (
		allFields       = make([]string, 0)
		insertFields    = make([]string, 0)
		InsertInfo      = make([]*SqlFieldInfo, 0)
		fieldsList      = make([]*SqlFieldInfo, 0)
		nullFieldList   = make([]*NullSqlFieldInfo, 0)
		updateList      = make([]string, 0)
		updateListField = make([]string, 0)
		PrimaryKey      = ""
		primaryType     = ""
	)

	for _, item := range tableDesc {
		allFields = append(allFields, "`"+item.ColumnName+"`")
		if item.PrimaryKey == false && item.ColumnName != "updated_at" && item.ColumnName != "created_at" {
			insertFields = append(insertFields, item.ColumnName)
			InsertInfo = append(InsertInfo, &SqlFieldInfo{
				HumpName: l.T.Capitalize(item.ColumnName),
				Comment:  item.ColumnComment,
			})
			if item.ColumnName == "identify" {
				updateList = append(updateList, item.ColumnName+"="+item.ColumnName+"+1")
			} else {
				updateList = append(updateList, item.ColumnName+"=?")
				if item.PrimaryKey == false {
					updateListField = append(updateListField, "value."+l.T.Capitalize(item.ColumnName))
				}
			}
		}
		if item.PrimaryKey {
			PrimaryKey = item.ColumnName
			primaryType = item.GolangType
		}
		fieldsList = append(fieldsList, &SqlFieldInfo{
			HumpName: l.T.Capitalize(item.ColumnName),
			Comment:  item.ColumnComment,
		})
		nullFieldList = append(nullFieldList, &NullSqlFieldInfo{
			HumpName:     l.T.Capitalize(item.ColumnName),
			OriFieldType: item.OriMysqlType,
			GoType:       MysqlTypeToGoType[item.OriMysqlType],
			Comment:      item.ColumnComment,
		})
	}
	//主键ID,用于更新
	if PrimaryKey != "" {
		updateListField = append(updateListField, "value."+l.T.Capitalize(PrimaryKey))
	}
	//拼出SQL所需要结构数据
	InsertMark := strings.Repeat("?,", len(insertFields))
	if len(InsertMark) > 0 {
		InsertMark = InsertMark[:len(InsertMark)-1]
	}
	sqlInfo := &SqlInfo{
		TableName:           tableName,
		PrimaryKey:          PrimaryKey,
		PrimaryType:         primaryType,
		StructTableName:     l.T.Capitalize(tableName),
		NullStructTableName: l.T.Capitalize(tableName) + DbNullPrefix,
		UpperTableName:      TablePrefix + l.T.ToUpper(tableName),
		AllFieldList:        strings.Join(allFields, ","),
		InsertFieldList:     strings.Join(insertFields, ","),
		InsertMark:          InsertMark,
		UpdateFieldList:     strings.Join(updateList, ","),
		UpdateListField:     updateListField,
		FieldsInfo:          fieldsList,
		NullFieldsInfo:      nullFieldList,
		InsertInfo:          InsertInfo,
	}
	err = l.GenerateSQL(sqlInfo, tableComment)
	//添加一个实例
	l.Once.Do(func() {
		l.GenerateExample(sqlInfo.StructTableName)
	})
	if err != nil {
		return
	}
	return
}

//生成一个实例文件
func (l *Logic) GenerateExample(name string) {
	//写入表名
	file := l.GetMysqlDir() + GoFile_Example

	//解析模板
	tplByte, err := Asset(TPL_EXAMPLE)
	if err != nil {
		return
	}
	tpl, err := template.New("example").Parse(string(tplByte))
	if err != nil {
		return
	}
	type s struct {
		Name string
	}
	ss := s{
		Name: name,
	}
	//解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, ss)
	if err != nil {
		return
	}
	//表信息写入文件
	err = WriteFile(file, content.String())
	if err != nil {
		return
	}
	return
}

//生成表列表
func (l *Logic) GenerateTableList(list []*TableList) (err error) {
	//写入表名
	file := l.GetMysqlDir() + GoFile_TableList
	//判断package是否加载过
	checkStr := "package " + PkgDbModels
	if l.T.CheckFileContainsChar(file, checkStr) == false {
		l.T.WriteFile(file, checkStr+"\n")
	}
	checkStr = "const"
	if l.T.CheckFileContainsChar(file, checkStr) {
		log.Println(file + "It already exists. Please delete it and regenerate it")
		return
	}
	tplByte, err := Asset(TPL_TABLES)
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
	err = WriteAppendFile(file, content.String())
	if err != nil {
		return
	}
	return
}

//生成SQL文件
func (l *Logic) GenerateSQL(info *SqlInfo, tableComment string) (err error) {
	//写入表名
	goFile := l.GetMysqlDir() + info.TableName + ".go"
	s := fmt.Sprintf(`
//%s
package %s
import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
`, tableComment, PkgDbModels)
	//判断package是否加载过
	if l.T.CheckFileContainsChar(goFile, "database/sql") == false {
		l.T.WriteFile(goFile, s)
	}

	//解析模板
	tplByte, err := Asset(TPL_CURD)
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
func (l *Logic) GenerateMarkdown(data *MarkDownData) (err error) {
	//写入markdown
	file := l.Path + fmt.Sprintf("markdown%s.md", time.Now().Format("2006-01-02_150405"))
	tplByte, err := Asset(TPL_MARKDOWN)
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
