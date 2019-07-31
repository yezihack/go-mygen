//判断package是否加载过
package models

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

//文章表
type Article struct {
	Id        int64
	Title     string    //文章标题
	Content   string    //文章内容
	ReadCount int64     //阅读数量
	State     int64     //状态,1使用中,0禁用
	CreatedAt time.Time //创建时间
	UpdatedAt time.Time //更新时间

}

//文章表 Null Entity
type ArticleNull struct {
	Id        sql.NullInt64
	Title     sql.NullString //文章标题
	Content   sql.NullString //文章内容
	ReadCount sql.NullInt64  //阅读数量
	State     sql.NullInt64  //状态,1使用中,0禁用
	CreatedAt mysql.NullTime //创建时间
	UpdatedAt mysql.NullTime //更新时间

}

//收藏文章表
type Cellect struct {
	Id        int64
	ArticleId int64     //文章ID
	UserId    int64     //用户ID
	State     int64     //状态,1公开,0私用
	CreatedAt time.Time //创建时间
	UpdatedAt time.Time //更新时间

}

//收藏文章表 Null Entity
type CellectNull struct {
	Id        sql.NullInt64
	ArticleId sql.NullInt64  //文章ID
	UserId    sql.NullInt64  //用户ID
	State     sql.NullInt64  //状态,1公开,0私用
	CreatedAt mysql.NullTime //创建时间
	UpdatedAt mysql.NullTime //更新时间

}

//用户表
type Users struct {
	Id        int64
	Username  string    //用户名称
	Email     string    //用户邮箱
	Password  string    //用户密码
	Openid    string    //微信OPENID
	State     int64     //用户状态,1使用中,0禁用
	CreatedAt time.Time //创建时间
	UpdatedAt time.Time //更新时间

}

//用户表 Null Entity
type UsersNull struct {
	Id        sql.NullInt64
	Username  sql.NullString //用户名称
	Email     sql.NullString //用户邮箱
	Password  sql.NullString //用户密码
	Openid    sql.NullString //微信OPENID
	State     sql.NullInt64  //用户状态,1使用中,0禁用
	CreatedAt mysql.NullTime //创建时间
	UpdatedAt mysql.NullTime //更新时间

}
