package mysql

import (
	"os"

	"fmt"

	"github.com/ThreeKing2018/k3log"
	"github.com/gohouse/gorose"
	_ "github.com/gohouse/gorose/driver/mysql"
	"github.com/yezihack/gm2m/common"
	"github.com/yezihack/gm2m/conf"
)

type DbTools struct {
	T *common.Tools
}

var Config1 = &gorose.DbConfigSingle{
	Driver:          "mysql",
	EnableQueryLog:  true,
	SetMaxOpenConns: 0,
	SetMaxIdleConns: 0,
	Prefix:          "",
	Dsn:             "%s:%s@tcp(%s:%s)/%s?charset=%s", // username:password@tcp(localhost:3306)/dbname?charset=utf8mb4
}
var __DB *gorose.Session

func init() {
	InitMysql()
}

func GetMasterDB() *gorose.Session {
	if __DB == nil {
		panic("数据库未连接,请检查配置文件(default.ini) 或 环境变量(GM2M_CONFIG)")
	}
	return __DB
}

func InitMysql() {
	iniFile := common.GetExeRootDir() + conf.DefaultIniFileName
	if new(common.Tools).IsDirOrFileExist(iniFile) {
		setDbConnByIniFile(iniFile)
	} else if dsn := os.Getenv("GM2M_CONFIG"); dsn != "" {
		setDbConnByDsnString(dsn)
	} else if ini := os.Getenv(conf.TMP_ENV_INI_FILE); ini != "" {
		setDbConnByIniFile(ini)
	} else {
		k3log.Error("请检查数据库配置default.ini文件或环境变量:GM2M_CONFIG")
	}
}
func setDbConnByIniFile(iniFile string) {
	dbConf, err := common.ReadDbConfig(iniFile) //获取数据库配置
	if err != nil {
		panic(err)
		return
	}
	Config1.Dsn = fmt.Sprintf(Config1.Dsn, dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DbName, dbConf.Charset)
	connection, err := gorose.Open(Config1)
	if err != nil {
		panic(err)
		return
	}
	__DB = connection.NewSession()
}

func setDbConnByDsnString(dsn string) {
	connection, err := gorose.Open(dsn)
	if err != nil {
		k3log.Error("Mysql连接失败,请检查环境变量: GM2M_CONFIG", err)
		return
	}
	__DB = connection.NewSession()
}
