package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"strings"
	"testing"
)

type DBConfigEntity struct {
	Host     string //地址
	Port     int    //端口
	Name     string //用户
	Pass     string //密码
	DBName   string //库名
	Charset  string //编码
	Timezone string //时区
	MaxIdle  int    //最大空间连接
	MaxOpen  int    //最大连接数
}

//连接数据库
func InitDB(cfg DBConfigEntity) *sql.DB {
	if strings.EqualFold(cfg.Timezone, "") {
		cfg.Timezone = "'Asia/Shanghai'"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local&time_zone=%s",
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
		panic(err)
	}
	connection.SetMaxOpenConns(cfg.MaxOpen)
	connection.SetMaxIdleConns(cfg.MaxIdle)
	return connection
}

var DB *sql.DB
var SgBlogTagsDb *SgBlogTagsModel

func initConnection() {
	cfg := DBConfigEntity{
		Host:    "localhost",
		Port:    3308,
		Name:    "root",
		Pass:    "123456",
		DBName:  "kindled",
		Charset: "utf8mb4",
		MaxOpen: 10,
		MaxIdle: 5,
	}
	DB = InitDB(cfg)
	SgBlogTagsDb = NewSgBlogTags(DB)
}

//查询所有的数据
func TestSgBlogTagsFindWhere(t *testing.T) {
	initConnection()
	book := SgBlogTags{
		//todo
	}
	result, err := SgBlogTagsDb.Find(&book) //函数里需要拼接一下sql
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result[0])
}

//获取最后一条数据
func TestSgBlogTagsLast(t *testing.T) {
	initConnection()
	result, err := SgBlogTagsDb.First(nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//获取总数量
func TestSgBlogTagsCount(t *testing.T) {
	initConnection()
	result, err := SgBlogTagsDb.First(nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//创建数据
func TestSgBlogTagsCreate(t *testing.T) {
	initConnection()
	dd := SgBlogTags{
		//todo
	}
	result, err := SgBlogTagsDb.Create(&dd)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//创建数据
func TestExampleUpdate(t *testing.T) {
	initConnection()
	dd := SgBlogTags{
		//todo
	}
	result, err := SgBlogTagsDb.Update(&dd)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
