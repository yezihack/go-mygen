package main

import "os"

const (
	Version   = "v3.0.0"
	UpdatedAt = " 2019.11.22"
)

const (
	DS           = string(os.PathSeparator) //通用/
	DbNullPrefix = "Null"                   //处理数据为空时结构的前缀定义
	TablePrefix  = "TABLE_"                 //表前缀
)

const (
	TPL_CURD      = "tpl/curd.tpl"      //生成CRUD2模板
	TPL_STRUCTURE = "tpl/structure.tpl" //结构体模板
	TPL_ENTITY    = "tpl/entity.tpl"    //结构实体模板
	TPL_TABLES    = "tpl/tables.tpl"    //表结构模板
	TPL_MARKDOWN  = "tpl/markdown.tpl"  //markdown模板
	TPL_EXAMPLE   = "tpl/example.tpl"   //example模板
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
	GOFILE_ENTITY    = "db_entity.go"    //生成的结构体实体 go文件名称
	GoFile_TableList = "table_list.go"   //表文件
	GoFile_Example   = "example_test.go" //表文件
)

const (
	PkgDbModels = "models" //db_models命名空间
	PkgEntity   = "entity" // entity实体命名空间
)

//帮助文档
var CmdHelp = []CmdEntity{
	{"0", "设置生成目录"},
	{"1", "生成表markdown文档"},
	{"2", "生成表结构实体"},
	{"3", "生成CURD增删改查"},
	{"4", "设置结构体映射名称"},
	{"5", "查找或设置表名"},
	{"6", "查找或设置库名-开发中.."},
	{"7, c, clear", "清屏"},
	{"8, h", "查看帮助"},
	{"9, e, q, exit", "退出"},
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

