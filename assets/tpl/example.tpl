package models

import (
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
"testing"
"strings"
"net/url"
)

type DBConfigEntity struct {
Host    string //地址
Port    int //端口
Name    string //用户
Pass    string //密码
DBName  string //库名
Charset string //编码
Timezone string //时区
MaxIdle int //最大空间连接
MaxOpen int //最大连接数
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
var {{.Name}}Db *{{.Name}}Model

func initConnection() {
cfg := DBConfigEntity{
Host:"localhost",
Port:3308,
Name:"root",
Pass:"123456",
DBName:"kindled",
Charset:"utf8mb4",
MaxOpen:10,
MaxIdle: 5,
}
DB = InitDB(cfg)
{{.Name}}Db = New{{.Name}}(DB)
}

//查询所有的数据
func Test{{.Name}}FindWhere(t *testing.T) {
initConnection()
book := {{.Name}}{
//todo
}
result, err := {{.Name}}Db.Find(&book) //函数里需要拼接一下sql
if err != nil {
t.Error(err)
}
fmt.Println(result[0])
}
//获取最后一条数据
func Test{{.Name}}Last(t *testing.T) {
initConnection()
result, err := {{.Name}}Db.First(nil)
if err != nil {
t.Error(err)
}
fmt.Println(result)
}
//获取总数量
func Test{{.Name}}Count(t *testing.T) {
initConnection()
result, err := {{.Name}}Db.First(nil)
if err != nil {
t.Error(err)
}
fmt.Println(result)
}
//创建数据
func Test{{.Name}}Create(t *testing.T) {
initConnection()
dd := {{.Name}}{
//todo
}
result, err := {{.Name}}Db.Create(&dd)
if err != nil {
t.Error(err)
}
fmt.Println(result)
}
//创建数据
func TestExampleUpdate(t *testing.T) {
initConnection()
dd := {{.Name}}{
//todo
}
result, err := {{.Name}}Db.Update(&dd)
if err != nil {
t.Error(err)
}
fmt.Println(result)
}