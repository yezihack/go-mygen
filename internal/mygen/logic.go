package mygen

import (
	"bytes"
	"fmt"
	"github.com/yezihack/go-mygen/internal/config"
	"github.com/yezihack/go-mygen/internal/entity"
	"github.com/yezihack/go-mygen/internal/pkg"
	"github.com/yezihack/go-mygen/third_party"
	"html/template"
	"log"
	"strings"
	"sync"
	"time"
)

type Logic struct {
	T    *pkg.Tools
	DB   *entity.ModelS
	Path string
	Once sync.Once
	l    sync.Mutex
}

// 生成结构实体文件
func (l *Logic) CreateEntity(formatList []string) (err error) {
	// 表结构文件路径
	path := pkg.CreateDir(l.Path+config.GODIR_Entity) + config.DS + config.GOFILE_ENTITY
	// 将表结构写入文件
	for idx, table := range l.DB.DoTables {
		idx++
		// 查询表结构信息
		tableDesc, err := l.DB.GetTableDesc(table.Name)
		if err != nil {
			log.Fatal("CreateEntityErr>>", err)
			continue
		}
		req := new(entity.EntityReq)
		req.Index = idx
		req.EntityPath = path
		req.TableName = table.Name
		req.TableComment = table.Comment
		req.TableDesc = tableDesc
		req.FormatList = formatList
		req.EntityPkg = config.PkgEntity
		// 生成基础信息
		err = l.GenerateDBEntity(req)
		if err != nil {
			log.Fatal("CreateEntityErr>>", err)
			continue
		}
	}
	return
}

// 生成原生的crud查询数据库
func (l *Logic) CreateCURD(formatList []string) (err error) {
	tableNameList := make([]*entity.TableList, 0)
	// 表结构文件路径
	// 将表结构写入文件
	for _, table := range l.DB.DoTables {
		tableNameList = append(tableNameList, &entity.TableList{
			UpperTableName: config.TablePrefix + l.T.ToUpper(table.Name),
			TableName:      table.Name,
			Comment:        table.Comment,
		})
		// 查询表结构信息
		tableDesc, err := l.DB.GetTableDesc(table.Name)
		if err != nil {
			return err
		}
		req := new(entity.EntityReq)
		req.TableName = table.Name
		req.TableComment = table.Comment
		req.TableDesc = tableDesc
		req.EntityPath = l.GetEntityDir() + config.GOFILE_ENTITY
		req.FormatList = formatList
		req.Pkg = config.PkgDbModels
		req.EntityPkg = config.PkgEntity
		// 生成基础信息
		err = l.GenerateDBEntity(req)
		if err != nil {
			return err
		}
		// 生成增,删,改,查文件
		err = l.GenerateCURDFile(table.Name, table.Comment, tableDesc)
		if err != nil {
			return err
		}
	}

	// 生成所有表的文件
	if err = l.GenerateTableList(tableNameList); err != nil {
		return
	}
	//生成init文件
	if err = l.GenerateInit(); err != nil {
		return
	}
	//生成error文件
	if err = l.GenerateError(); err != nil {
		return
	}
	fmt.Println("`CURD` files created finish!")
	return nil
}

// 生成mysql markdown文档
func (l *Logic) CreateMarkdown() (err error) {
	data := new(entity.MarkDownData)
	// 将表结构写入文件
	i := 1
	for _, table := range l.DB.DoTables {
		fmt.Println("Doing table:" + table.Name)
		data.TableList = append(data.TableList, &entity.TableList{
			Index:          i,
			UpperTableName: l.T.ToUpper(table.Name),
			TableName:      table.Name,
			Comment:        table.Comment,
		})
		// 查询表结构信息
		desc := new(entity.MarkDownDataChild)
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
	// 生成所有表的文件
	err = l.GenerateMarkdown(data)
	if err != nil {
		return
	}
	return
}

// 创建和获取MYSQL目录
func (l *Logic) GetMysqlDir() string {
	return pkg.CreateDir(l.Path + config.GODIR_MODELS + config.DS)
}

// 创建和获取MYSQL目录
func (l *Logic) GetConfigDir() string {
	return pkg.CreateDir(l.Path + config.GODIR_MODELS + config.DS + config.GODIR_Config + config.DS)
}

// 创建和获取MYSQL目录
func (l *Logic) GetEntityDir() string {
	return pkg.CreateDir(l.Path + config.GODIR_MODELS + config.DS + config.GODIR_Entity + config.DS)
}

// 获取根目录地址
func (l *Logic) GetRoot() string {
	return pkg.GetRootPath(l.Path) + config.DS
}

// 创建结构体
func (l *Logic) GenerateDBStructure(tableName, tableComment, path string, tableDesc []*entity.TableDesc) (err error) {
	// 加入package
	packageStr := `// 数据库表内结构体信息
package mysql
` // 判断package是否加载过
	// 判断文件是否存在.

	if l.T.CheckFileContainsChar(path, packageStr) == false {
		l.T.WriteFile(path, packageStr)
	}
	// 判断import是否加载过
	importStr := `import "database/sql"`
	if l.T.CheckFileContainsChar(path, importStr) == false {
		l.T.WriteFileAppend(path, importStr)
	}
	// 声明表结构变量
	TableData := new(entity.TableInfo)
	TableData.Table = l.T.Capitalize(tableName)
	TableData.NullTable = config.DbNullPrefix + TableData.Table
	TableData.TableComment = tableComment
	// 判断表结构是否加载过
	if l.T.CheckFileContainsChar(path, "type "+TableData.Table+" struct") == true {
		return
	}
	// 加载模板文件
	tplByte, err := third_party.Asset(config.TPL_STRUCTURE)
	if err != nil {
		return
	}
	tpl, err := template.New("structure").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 装载表字段信息
	fts := []string{"json"}
	if err != nil {
		return
	}
	// 判断是否含json
	if !pkg.InArrayString("json", fts) {
		index0 := fts[0]
		fts[0] = "json"
		fts = append(fts, index0)
	}
	for _, val := range tableDesc {
		TableData.Fields = append(TableData.Fields, &entity.FieldsInfo{
			Name:         l.T.Capitalize(val.ColumnName),
			Type:         val.GolangType,
			NullType:     val.MysqlNullType,
			DbOriField:   val.ColumnName,
			FormatFields: pkg.FormatField(val.ColumnName, fts),
			Remark:       val.ColumnComment,
		})
	}
	content := bytes.NewBuffer([]byte{})
	tpl.Execute(content, TableData)
	// 表信息写入文件

	err = pkg.WriteAppendFile(path, content.String())
	if err != nil {
		return
	}
	return
}

// 创建结构实体
func (l *Logic) GenerateDBEntity(req *entity.EntityReq) (err error) {
	l.l.Lock()
	defer l.l.Unlock()
	var s string
	s = fmt.Sprintf(`// 判断package是否加载过
package %s
import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)
`, req.EntityPkg)
	// 判断import是否加载过
	check := "github.com/go-sql-driver/mysql"
	if l.T.CheckFileContainsChar(req.EntityPath, check) == false {
		l.T.WriteFile(req.EntityPath, s)
	}
	// 声明表结构变量
	TableData := new(entity.TableInfo)
	TableData.Table = l.T.Capitalize(req.TableName)
	TableData.NullTable = TableData.Table + config.DbNullPrefix
	TableData.TableComment = pkg.AddToComment(req.TableComment, "")
	TableData.TableCommentNull = pkg.AddToComment(req.TableComment, " Null Entity")
	// 判断表结构是否加载过
	if l.T.CheckFileContainsChar(req.EntityPath, "type "+TableData.Table+" struct") == true {
		log.Println(req.EntityPath + "It already exists. Please delete it and regenerate it")
		return
	}
	// 加载模板文件
	tplByte, err := third_party.Asset(config.TPL_ENTITY)
	if err != nil {
		return
	}
	tpl, err := template.New("entity").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 装载表字段信息
	for _, val := range req.TableDesc {
		TableData.Fields = append(TableData.Fields, &entity.FieldsInfo{
			Name:         l.T.Capitalize(val.ColumnName),
			Type:         val.GolangType,
			NullType:     val.MysqlNullType,
			DbOriField:   val.ColumnName,
			FormatFields: pkg.FormatField(val.ColumnName, req.FormatList),
			Remark:       pkg.AddToComment(val.ColumnComment, ""),
		})
	}
	content := bytes.NewBuffer([]byte{})
	tpl.Execute(content, TableData)
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = pkg.WriteAppendFile(req.EntityPath, con)
	if err != nil {
		return
	}
	return
}

// 生成C增,U删,R查,D改,的文件
func (l *Logic) GenerateCURDFile(tableName, tableComment string, tableDesc []*entity.TableDesc) (err error) {
	var (
		allFields       = make([]string, 0)
		insertFields    = make([]string, 0)
		InsertInfo      = make([]*entity.SqlFieldInfo, 0)
		fieldsList      = make([]*entity.SqlFieldInfo, 0)
		nullFieldList   = make([]*entity.NullSqlFieldInfo, 0)
		updateList      = make([]string, 0)
		updateListField = make([]string, 0)
		PrimaryKey      = ""
		primaryType     = ""
	)
	// 存放第个字段
	var secondField string
	for _, item := range tableDesc {
		allFields = append(allFields, pkg.AddQuote(item.ColumnName))
		if item.PrimaryKey == false && item.ColumnName != "updated_at" && item.ColumnName != "created_at" {
			insertFields = append(insertFields, pkg.AddQuote(item.ColumnName))
			InsertInfo = append(InsertInfo, &entity.SqlFieldInfo{
				HumpName: l.T.Capitalize(item.ColumnName),
				Comment:  item.ColumnComment,
			})
			if item.ColumnName == "identify" {
				updateList = append(updateList, pkg.AddQuote(item.ColumnName)+"="+item.ColumnName+"+1")
			} else {
				updateList = append(updateList, pkg.AddQuote(item.ColumnName)+"=?")
				if item.PrimaryKey == false {
					updateListField = append(updateListField, "value."+l.T.Capitalize(item.ColumnName))
				}
			}
		}
		if item.PrimaryKey {
			PrimaryKey = item.ColumnName
			primaryType = item.GolangType
		} else {
			// 除了主键外的任意一个字段即可。
			if secondField == "" {
				secondField = item.ColumnName
			}
		}
		fieldsList = append(fieldsList, &entity.SqlFieldInfo{
			HumpName: l.T.Capitalize(item.ColumnName),
			Comment:  item.ColumnComment,
		})
		nullFieldList = append(nullFieldList, &entity.NullSqlFieldInfo{
			HumpName:     l.T.Capitalize(item.ColumnName),
			OriFieldType: item.OriMysqlType,
			GoType:       config.MysqlTypeToGoType[item.OriMysqlType],
			Comment:      item.ColumnComment,
		})
	}
	// 主键ID,用于更新
	if PrimaryKey != "" {
		updateListField = append(updateListField, "value."+l.T.Capitalize(PrimaryKey))
	}
	// 拼出SQL所需要结构数据
	InsertMark := strings.Repeat("?,", len(insertFields))
	if len(InsertMark) > 0 {
		InsertMark = InsertMark[:len(InsertMark)-1]
	}
	sqlInfo := &entity.SqlInfo{
		TableName:           tableName,
		PrimaryKey:          pkg.AddQuote(PrimaryKey),
		PrimaryType:         primaryType,
		StructTableName:     l.T.Capitalize(tableName),
		NullStructTableName: l.T.Capitalize(tableName) + config.DbNullPrefix,
		PkgEntity:           config.PkgEntity + ".",
		PkgTable:            config.PkgTable + ".",
		UpperTableName:      config.TablePrefix + l.T.ToUpper(tableName),
		AllFieldList:        strings.Join(allFields, ","),
		InsertFieldList:     strings.Join(insertFields, ","),
		InsertMark:          InsertMark,
		UpdateFieldList:     strings.Join(updateList, ","),
		UpdateListField:     updateListField,
		FieldsInfo:          fieldsList,
		NullFieldsInfo:      nullFieldList,
		InsertInfo:          InsertInfo,
		SecondField:         pkg.AddQuote(secondField),
	}
	err = l.GenerateSQL(sqlInfo, tableComment)
	// 添加一个实例
	l.Once.Do(func() {
		l.GenerateExample(sqlInfo.StructTableName)
	})

	if err != nil {
		return
	}
	return
}

// 生成一个实例文件
func (l *Logic) GenerateExample(name string) {
	// 写入表名
	file := l.GetMysqlDir() + config.GoFile_Example

	// 解析模板
	tplByte, err := third_party.Asset(config.TPL_EXAMPLE)
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
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, ss)
	if err != nil {
		return
	}
	// 表信息写入文件
	err = pkg.WriteFile(file, content.String())
	if err != nil {
		return
	}
	return
}

// 生成表列表
func (l *Logic) GenerateTableList(list []*entity.TableList) (err error) {
	// 写入表名
	file := l.GetConfigDir() + config.GoFile_TableList
	// 判断package是否加载过
	checkStr := "package " + config.PkgTable
	if l.T.CheckFileContainsChar(file, checkStr) == false {
		l.T.WriteFile(file, checkStr+"\n")
	}
	checkStr = "const"
	if l.T.CheckFileContainsChar(file, checkStr) {
		log.Println(file + "It already exists. Please delete it and regenerate it")
		return
	}
	tplByte, err := third_party.Asset(config.TPL_TABLES)
	if err != nil {
		return
	}
	tpl, err := template.New("table_list").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, list)
	if err != nil {
		return
	}
	// 表信息写入文件
	err = pkg.WriteAppendFile(file, content.String())
	if err != nil {
		return
	}
	return
}

// 生成init
func (l *Logic) GenerateInit() (err error) {
	file := l.GetMysqlDir() + config.GoFile_Init
	// 判断package是否加载过
	checkStr := "package " + config.PkgDbModels
	if l.T.CheckFileContainsChar(file, checkStr) == false {
		l.T.WriteFile(file, checkStr+"\n")
	}
	checkStr = "DBConfig"
	if l.T.CheckFileContainsChar(file, checkStr) {
		log.Println(file + "It already exists. Please delete it and regenerate it")
		return
	}
	tplByte, err := third_party.Asset(config.TPL_INIT)
	if err != nil {
		return
	}
	tpl, err := template.New("init").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, nil)
	if err != nil {
		return
	}
	// 表信息写入文件
	err = pkg.WriteAppendFile(file, content.String())
	if err != nil {
		return
	}
	return
}

// 生成init
func (l *Logic) GenerateError() (err error) {
	file := l.GetMysqlDir() + config.GoFile_Error
	// 判断package是否加载过
	checkStr := "package " + config.PkgDbModels
	if l.T.CheckFileContainsChar(file, checkStr) == false {
		l.T.WriteFile(file, checkStr+"\n")
	}
	// 判断是否已经生成过此文件
	checkStr = "Stack"
	if l.T.CheckFileContainsChar(file, checkStr) {
		log.Println(file + "It already exists. Please delete it and regenerate it")
		return
	}
	tplByte, err := third_party.Asset(config.TPL_Error)
	if err != nil {
		return
	}
	tpl, err := template.New("error").Parse(string(tplByte))
	if err != nil {
		return
	}
	// analysis execute template
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, nil)
	if err != nil {
		return
	}
	// append write to file
	err = pkg.WriteAppendFile(file, content.String())
	if err != nil {
		return
	}
	return
}

// 生成SQL文件
func (l *Logic) GenerateSQL(info *entity.SqlInfo, tableComment string) (err error) {
	// 写入表名
	goFile := l.GetMysqlDir() + info.TableName + ".go"
	s := fmt.Sprintf(`
// %s
package %s
import(
	"database/sql"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)
`, tableComment, config.PkgDbModels)
	// 判断package是否加载过
	if l.T.CheckFileContainsChar(goFile, "database/sql") == false {
		l.T.WriteFile(goFile, s)
	}

	// 解析模板
	tplByte, err := third_party.Asset(config.TPL_CURD)
	if err != nil {
		return
	}
	tpl, err := template.New("CURD").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, info)
	if err != nil {
		return
	}
	// 表信息写入文件
	if l.T.CheckFileContainsChar(goFile, info.StructTableName) == false {
		err = pkg.WriteAppendFile(goFile, content.String())
		if err != nil {
			return
		}
	}
	return
}

// 生成表列表
func (l *Logic) GenerateMarkdown(data *entity.MarkDownData) (err error) {
	// 写入markdown
	file := l.Path + fmt.Sprintf("markdown%s.md", time.Now().Format("2006-01-02_150405"))
	tplByte, err := third_party.Asset(config.TPL_MARKDOWN)
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	tpl, err := template.New("markdown").Parse(string(tplByte))
	err = tpl.Execute(content, data)
	if err != nil {
		return
	}
	// 表信息写入文件
	err = pkg.WriteAppendFile(file, content.String())
	if err != nil {
		return
	}
	return
}
