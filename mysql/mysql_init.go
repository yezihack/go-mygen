package mysql

import (
	"fmt"

	_ "github.com/gohouse/gorose/driver/mysql"
	"github.com/pkg/errors"
	"github.com/yezihack/gm2m/common"
	"github.com/yezihack/gm2m/conf"
	"database/sql"
)

type DbTools struct {
	T *common.Tools
}
//conf.DBConfig
func SetDbConn(dbConf conf.DBConfig) (db *sql.DB, err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DbName, dbConf.Charset)
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.New("数据库连接失败, err" + err.Error())
	}
	db = connection
	return db, nil
}
