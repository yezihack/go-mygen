//判断package是否加载过
package entity

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SgBlogTags struct {
	Id        int64
	BlogId    int64 //文章ID,与blogs表相关
	TagId     int64 //标签ID,与tags表相关
	CreatedAt int64 //创建时间
	UpdatedAt int64 //更新时间

}

type SgBlogTagsNull struct {
	Id        sql.NullInt64
	BlogId    sql.NullInt64 //文章ID,与blogs表相关
	TagId     sql.NullInt64 //标签ID,与tags表相关
	CreatedAt sql.NullInt64 //创建时间
	UpdatedAt sql.NullInt64 //更新时间

}

type SgBlogs struct {
	Id         int64
	UserId     int64  //用户ID,与users表相关
	Title      string //标题
	Content    string
	Type       int64 //文章类型,1为原创,2为转载
	IsPush     int64 //是否推送，0未，1是
	Status     int64 //文章状态:1发布,2草稿
	FirstTagId int64 //首标签ID
	ReadCount  int64 //阅读次数
	StarCount  int64 //点赞次数
	CreatedAt  int64 //创建时间
	UpdatedAt  int64 //更新时间

}

type SgBlogsNull struct {
	Id         sql.NullInt64
	UserId     sql.NullInt64  //用户ID,与users表相关
	Title      sql.NullString //标题
	Content    sql.NullString
	Type       sql.NullInt64 //文章类型,1为原创,2为转载
	IsPush     sql.NullInt64 //是否推送，0未，1是
	Status     sql.NullInt64 //文章状态:1发布,2草稿
	FirstTagId sql.NullInt64 //首标签ID
	ReadCount  sql.NullInt64 //阅读次数
	StarCount  sql.NullInt64 //点赞次数
	CreatedAt  sql.NullInt64 //创建时间
	UpdatedAt  sql.NullInt64 //更新时间

}

type SgConfigs struct {
	Id         int64
	Name       string //名称
	Key        string //键
	Value      string //值
	ConfigType string //变量类型：system 系统内置/user 用户定义
	Remark     string
	CreatedAt  int64 //创建时间
	UpdatedAt  int64 //更新时间

}

type SgConfigsNull struct {
	Id         sql.NullInt64
	Name       sql.NullString //名称
	Key        sql.NullString //键
	Value      sql.NullString //值
	ConfigType sql.NullString //变量类型：system 系统内置/user 用户定义
	Remark     sql.NullString
	CreatedAt  sql.NullInt64 //创建时间
	UpdatedAt  sql.NullInt64 //更新时间

}

type SgMigrations struct {
	Id        int64
	Migration string
	Batch     int64
}

type SgMigrationsNull struct {
	Id        sql.NullInt64
	Migration sql.NullString
	Batch     sql.NullInt64
}

type SgTags struct {
	Id        int64
	Name      string //名称
	Remark    string //备注
	CreatedAt int64  //创建时间
	UpdatedAt int64  //更新时间

}

type SgTagsNull struct {
	Id        sql.NullInt64
	Name      sql.NullString //名称
	Remark    sql.NullString //备注
	CreatedAt sql.NullInt64  //创建时间
	UpdatedAt sql.NullInt64  //更新时间

}

type SgUsers struct {
	Id           int64
	Name         string //用户名
	Pass         string //密码
	Salt         string //盐
	Google2faKey string //谷歌两步验证密钥
	LoginCount   int64  //登陆次数
	Token        string //记住密码生成的随机数
	LoginIp      string //登陆IP
	CreatedAt    int64  //创建时间
	UpdatedAt    int64  //更新时间

}

type SgUsersNull struct {
	Id           sql.NullInt64
	Name         sql.NullString //用户名
	Pass         sql.NullString //密码
	Salt         sql.NullString //盐
	Google2faKey sql.NullString //谷歌两步验证密钥
	LoginCount   sql.NullInt64  //登陆次数
	Token        sql.NullString //记住密码生成的随机数
	LoginIp      sql.NullString //登陆IP
	CreatedAt    sql.NullInt64  //创建时间
	UpdatedAt    sql.NullInt64  //更新时间

}
