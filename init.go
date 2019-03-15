package gm2m

import (
	"database/sql"
	"fmt"
	"net/url"
)

//连接数据库
func InitDB(cfg DBConfig) (*sql.DB, error) {
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
	if err != nil {
		return nil, err
	}
	if err = connection.Ping(); err != nil {
		return nil, err
	}
	return connection, nil
}

//实例一个数据库对象
func NewDB(db *sql.DB) *ModelS {
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	return &ModelS{
		DB: db,
	}
}
