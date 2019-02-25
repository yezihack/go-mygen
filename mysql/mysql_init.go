package mysql

import (
	"fmt"

	"github.com/ThreeKing2018/goutil/golog"
	"github.com/gohouse/gorose"
	_ "github.com/gohouse/gorose/driver/mysql"
	"github.com/yezihack/m2m/common"
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
	Dsn:             "%s:%s@tcp(%s:%d)/%s?charset=%s",
}
var DB *gorose.Session

func init() {
	dbConf, err := common.ReadDbConfig() //获取数据库配置
	if err != nil {
		panic(err)
		return
	}
	Config1.Dsn = fmt.Sprintf(Config1.Dsn, dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DbName, dbConf.Charset)
	fmt.Println(Config1.Dsn)
	connection, err := gorose.Open(Config1)
	if err != nil {
		golog.Panic("Mysql连接失败,请检查数据库配置ini文件", err)
		return
	}
	DB = connection.NewSession()
}
