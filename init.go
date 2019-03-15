package gomygen

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"strings"
)

//连接数据库
func InitDB(cfg DBConfig) (*sql.DB, error) {
	if strings.EqualFold(cfg.Timezone, "") {
		cfg.Timezone = "'Asia/Shanghai'"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local&time_zone=%s",
		cfg.Name,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
		url.QueryEscape(cfg.Timezone),
	)

	connection, err := sql.Open("mysql", dsn)
	defer connection.Close()
	if err != nil {
		return nil, err
	}
	if err = connection.Ping(); err != nil {
		return nil, err
	}
	connection.SetMaxIdleConns(cfg.MaxIdleConn)
	connection.SetMaxOpenConns(cfg.MaxOpenConn)
	return connection, nil
}

//实例一个数据库对象
func NewDB(db *sql.DB) *ModelS {
	return &ModelS{
		DB: db,
	}
}
