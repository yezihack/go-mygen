package mysql

import (
	"fmt"

	"os"

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
	Dsn:             "%s:%s@tcp(%s:%d)/%s?charset=%s", // username:password@tcp(localhost:3306)/dbname?charset=utf8mb4
}
var DB *gorose.Session

func init() {
	iniFile := common.GetRootDir() + conf.DefaultIniFileName
	if new(common.Tools).IsDirOrFileExist(iniFile) {
		setDbConnByIniFile(iniFile)
	} else if dsn := os.Getenv("GM2M_CONFIG"); dsn != "" {
		setDbConnByDsnString(dsn)
	} else if ini := os.Getenv(conf.TMP_ENV_INI_FILE); ini != "" {
		setDbConnByIniFile(ini)
	}
}
func setDbConnByIniFile(iniFile string) {
	dbConf, err := common.ReadDbConfig(iniFile) //获取数据库配置
	if err != nil {
		k3log.Error("请检查数据库配置default.ini文件", err)
		return
	}
	Config1.Dsn = fmt.Sprintf(Config1.Dsn, dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DbName, dbConf.Charset)
	connection, err := gorose.Open(Config1)
	if err != nil {
		k3log.Error("Mysql连接失败,请检查数据库配置default.ini文件", err)
		return
	}
	DB = connection.NewSession()
}
func setDbConnByDsnString(dsn string) {
	connection, err := gorose.Open(dsn)
	if err != nil {
		k3log.Error("Mysql连接失败,请检查环境变量: GM2M_CONFIG", err)
		return
	}
	DB = connection.NewSession()
}
