package conf

import "os"

const (
	DefaultIniFileName = "default.ini"            //默认文件配置名称
	DS                 = string(os.PathSeparator) //通用/
	DbNullPrefix       = "null"                   //处理数据为空时结构的前缀定义
)
const (
	TMP_ENV_INI_FILE = "GM2M_ENV_INI_FILE" //临时的环境变量
)

const (
	TPL_CONFIG    = "tpl/conf.tpl"      //配置模板
	TPL_CRUD      = "tpl/crud.tpl"      //生成CRUD模板
	TPL_STRUCTURE = "tpl/structure.tpl" //结构体模板
	TPL_TABLES    = "tpl/tables.tpl"    //表结构模板
	TPL_MARKDOWN  = "tpl/markdown.tpl"  //markdown模板
)

const (
	Unknow = iota
	Darwin
	Window
	Linux
)

const (
	GOFILE_STRUCTURE = "db_structure.go" //生成的结构体 go文件名称
)

//帮助文档
var CmdHelp = []CmdS{
	{"1", "生成表markdown文档"},
	{"2", "生成表结构数据"},
	{"3", "生成CURD增删改查"},
	{"7, c, clear", "清屏"},
	{"8, h", "查看帮助"},
	{"9, e, exit", "退出"},
}

//mysql类型 <=> golang类型
var MysqlTypeToGoType = map[string]string{
	"tinyint":    "int",
	"smallint":   "int",
	"mediumint":  "int",
	"int":        "int",
	"integer":    "int",
	"bigint":     "int64",
	"float":      "float64",
	"double":     "float64",
	"decimal":    "float64",
	"date":       "string",
	"time":       "string",
	"year":       "string",
	"datetime":   "string",
	"timestamp":  "string",
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
	"datetime":   "sql.NullString",
	"timestamp":  "sql.NullString",
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
