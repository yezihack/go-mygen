package mysql

import (
	"fmt"

	"database/sql"

	_ "github.com/gohouse/gorose/driver/mysql"
	"github.com/yezihack/gm2m/common"
)

// model结构
type ModelS struct {
	DB *sql.DB
	T  *common.Tools
}

//数据库配置结构
type DBConfig struct {
	Host    string //地址
	Port    int    //端口
	Name    string //名称
	Pass    string //密码
	DBName  string //库名
	Charset string //编码
}

//连接数据库
func InitDB(cfg DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", cfg.Name, cfg.Pass, cfg.Host, cfg.Port, cfg.DBName, cfg.Charset)
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
