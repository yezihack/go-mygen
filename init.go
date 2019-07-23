package gomygen

import (
	"database/sql"
	"fmt"
	"github.com/ThreeKing2018/k3log"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//连接数据库
func InitDB(cfg DBConfig) (*sql.DB, error) {
	if strings.EqualFold(cfg.Timezone, "") {
		cfg.Timezone = "'Asia/Shanghai'"
	}
	if strings.EqualFold(cfg.Charset, "") {
		cfg.Charset = "utf8mb4"
	}
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local&time_zone=%s",
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local",
		cfg.Name,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
		//url.QueryEscape(cfg.Timezone),
	)
	connection, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = connection.Ping(); err != nil {
		k3log.Error("InitDB", err)
		return nil, err
	}
	connection.SetMaxIdleConns(cfg.MaxIdleConn)
	connection.SetMaxOpenConns(cfg.MaxOpenConn)
	return connection, nil
}

//实例一个数据库对象
func NewDB() *ModelS {
	return new(ModelS)
}
func (m *ModelS) Using(db *sql.DB) {
	m.DB = db
}
