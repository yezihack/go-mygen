package main

import "os"

const (
	ProjectName = "go-mygen"
	Version     = "v3.2.0"
	Copyright   = "2020.03"
	Author      = "百里"
	AuthorEmail = "barry300@126.com"
)

const (
	DS              = string(os.PathSeparator) //通用/
	DbNullPrefix    = "Null"                   //处理数据为空时结构的前缀定义
	TablePrefix     = "TABLE_"                 //表前缀
	DefaultSavePath = "output"                 //默认生成目录名称
)

const (
	TPL_CURD      = "assets/tpl/curd.tpl"      //生成CRUD2模板
	TPL_STRUCTURE = "assets/tpl/structure.tpl" //结构体模板
	TPL_ENTITY    = "assets/tpl/entity.tpl"    //结构实体模板
	TPL_TABLES    = "assets/tpl/tables.tpl"    //表结构模板
	TPL_INIT      = "assets/tpl/init.tpl"      //init模板
	TPL_MARKDOWN  = "assets/tpl/markdown.tpl"  //markdown模板
	TPL_EXAMPLE   = "assets/tpl/example.tpl"   //example模板
)

const (
	Unknown = iota
	Darwin
	Window
	Linux
)

//生成的go文件
const (
	GODIR_MODELS     = "db_models"       //生成的结构体 go文件名称
	GODIR_Config     = "config"          //生成的结构体 go文件名称
	GODIR_Entity     = "entity"          //生成的结构体 go文件名称
	GOFILE_ENTITY    = "db_entity.go"    //生成的结构体实体 go文件名称
	GoFile_TableList = "table_list.go"   //表文件
	GoFile_Init      = "init.go"         //表文件
	GoFile_Example   = "example_test.go" //表文件
)

const (
	PkgDbModels = "mysql"  //db_models命名空间
	PkgEntity   = "entity" // entity实体命名空间
	PkgTable    = "config" //表的空间名称
)

//帮助文档
var CmdHelp = []CmdEntity{
	{"0", "Set build directory"},
	{"1", "Generate the table markdown document"},
	{"2", "Generate table structure entities"},
	{"3", "Generate CURD insert, delete, update and select"},
	{"4", "Sets the struct mapping name"},
	{"5", "Find or set the table name"},
	{"7, c, clear", "Clear the screen"},
	{"8, h, help", "Show help list"},
	{"9, q, quit", "Quit"},
}

//mysql类型 <=> golang类型
var MysqlTypeToGoType = map[string]string{
	"tinyint":    "int64",
	"smallint":   "int64",
	"mediumint":  "int64",
	"int":        "int64",
	"integer":    "int64",
	"bigint":     "int64",
	"float":      "float64",
	"double":     "float64",
	"decimal":    "float64",
	"date":       "string",
	"time":       "string",
	"year":       "string",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"char":       "string",
	"varchar":    "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
}

//MYSQL => golang mysql NULL TYPE
var MysqlTypeToGoNullType = map[string]string{
	"tinyint":    "sql.NullInt64",
	"smallint":   "sql.NullInt64",
	"mediumint":  "sql.NullInt64",
	"int":        "sql.NullInt64",
	"integer":    "sql.NullInt64",
	"bigint":     "sql.NullInt64",
	"float":      "sql.NullFloat64",
	"double":     "sql.NullFloat64",
	"decimal":    "sql.NullFloat64",
	"date":       "sql.NullString",
	"time":       "sql.NullString",
	"year":       "sql.NullString",
	"datetime":   "mysql.NullTime",
	"timestamp":  "mysql.NullTime",
	"char":       "sql.NullString",
	"varchar":    "sql.NullString",
	"tinyblob":   "sql.NullString",
	"tinytext":   "sql.NullString",
	"blob":       "sql.NullString",
	"text":       "sql.NullString",
	"mediumblob": "sql.NullString",
	"mediumtext": "sql.NullString",
	"longblob":   "sql.NullString",
	"longtext":   "sql.NullString",
}
