package config

import (
	"github.com/yezihack/go-mygen/internal/entity"
	"os"
)

const (
	ProjectName = "go-mygen"
	Version     = "v3.3.9"
	Copyright   = "2020.12"
	Author      = "百里"
	AuthorEmail = "freeit@126.com"
)

const (
	DS              = string(os.PathSeparator) // 通用/
	DbNullPrefix    = "Null"                   // 处理数据为空时结构的前缀定义
	TablePrefix     = "TABLE_"                 // 表前缀
	DefaultSavePath = "output"                 // 默认生成目录名称
)

const (
	TPL_CURD      = "assets/tpl/curd.tpl"      // 生成CRUD2模板
	TPL_STRUCTURE = "assets/tpl/structure.tpl" // 结构体模板
	TPL_ENTITY    = "assets/tpl/entity.tpl"    // 结构实体模板
	TPL_TABLES    = "assets/tpl/tables.tpl"    // 表结构模板
	TPL_INIT      = "assets/tpl/init.tpl"      // init模板
	TPL_MARKDOWN  = "assets/tpl/markdown.tpl"  // markdown模板
	TPL_EXAMPLE   = "assets/tpl/example.tpl"   // example模板
	TPL_Error     = "assets/tpl/e.tpl"         // error模板
)

const (
	Unknown = iota
	Darwin
	Window
	Linux
)

// generate file name
const (
	GODIR_MODELS     = "db_models"       // model file
	GODIR_Config     = "config"          // config file
	GODIR_Entity     = "entity"          // entity file
	GOFILE_ENTITY    = "db_entity.go"    // entity table file
	GoFile_TableList = "table_list.go"   // table file
	GoFile_Init      = "init.go"         // init file
	GoFile_Error     = "e.go"            // error file
	GoFile_Example   = "example_test.go" // example file
)

const (
	PkgDbModels = "mysql"  // db_models package name
	PkgEntity   = "entity" // entity package name
	PkgTable    = "config" // table package name
)

// help list
var CmdHelp = []entity.CmdEntity{
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
	"tinyint":    "int32",
	"smallint":   "int32",
	"mediumint":  "int32",
	"int":        "int32",
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
	"tinyint":    "sql.NullInt32",
	"smallint":   "sql.NullInt32",
	"mediumint":  "sql.NullInt32",
	"int":        "sql.NullInt32",
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
