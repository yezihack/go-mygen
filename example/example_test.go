package models

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
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

func init() {
	cfg := DBConfigEntity{
		Host:    "localhost",
		Port:    3308,
		Name:    "root",
		Pass:    "123456",
		DBName:  "100blog",
		Charset: "utf8mb4",
		MaxOpen: 10,
		MaxIdle: 5,
	}
	DB = InitDB(cfg)
}
func TestUsersModel_Find(t *testing.T) {
	result, err := NewUsers(DB).First(nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(result.Id)
}

//查询单行数据
func TestFirst(t *testing.T) {
	value := Article{
		//todo
	}
	result, err := NewArticle(DB).First(&value) //调用查询单行数据的First
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//查询多行数据
func TestFind(t *testing.T) {
	value := Article{
		//todo
	}
	result, err := NewArticle(DB).Find(&value) //函数里需要拼接一下sql
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

// 使用IN查找数据
func TestFindIn(t *testing.T) {
	ids := []int{1, 2, 3, 4}
	result, err := NewArticle(DB).FindIn(ids)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//获取最后一条数据
func TestLast(t *testing.T) {
	result, err := NewArticle(DB).Last(nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//获取总数量
func TestCount(t *testing.T) {
	result, err := NewArticle(DB).Count()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//创建数据
func TestCreate(t *testing.T) {
	value := Article{
		//todo
	}
	result, err := NewArticle(DB).Create(&value)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//创建数据 支持事务
func TestCreateTx(t *testing.T) {
	value := Article{
		//todo
	}
	//开启一个事务对象
	tx, err := DB.Begin()
	if err != nil {
		t.Error(err)
	}
	//使用defer 自动回滚或提交
	defer func() {
		var errTx error
		if err != nil {
			errTx = tx.Rollback()
		} else {
			errTx = tx.Commit()
		}
		if errTx != nil {
			return
		}
		return
	}()
	result, err := NewArticleTx(tx).CreateTx(&value)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//更新数据
func TestUpdate(t *testing.T) {
	value := Article{
		//todo
	}
	result, err := NewArticle(DB).Update(&value)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//更新数据 支持事务操作
func TestUpdateTx(t *testing.T) {
	value := Article{
		//todo
	}
	//开启一个事务对象
	tx, err := DB.Begin()
	if err != nil {
		t.Error(err)
	}
	//使用defer 自动回滚或提交
	defer func() {
		var errTx error
		if err != nil {
			errTx = tx.Rollback()
		} else {
			errTx = tx.Commit()
		}
		if errTx != nil {
			return
		}
		return
	}()
	//实例事务的结构对象,传递一个事务句柄
	result, err := NewArticleTx(tx).UpdateTx(&value)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//创建数据 by SQL
func TestSave(t *testing.T) {
	sqlTxt := "insert into tables values (?, ?, ?)"
	result, err := NewArticle(DB).Save(sqlTxt, 1, 2, 3)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

//更新数据 支持事务操作
func TestSaveTx(t *testing.T) {
	sqlTxt := "insert into tables values (?, ?, ?)"
	//开启一个事务对象
	tx, err := DB.Begin()
	if err != nil {
		t.Error(err)
	}
	//使用defer 自动回滚或提交
	defer func() {
		var errTx error
		if err != nil {
			errTx = tx.Rollback()
		} else {
			errTx = tx.Commit()
		}
		if errTx != nil {
			return
		}
		return
	}()
	//实例事务的结构对象,传递一个事务句柄
	result, err := NewArticleTx(tx).SaveTx(sqlTxt)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
