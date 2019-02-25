package conf

//数据库的配置
type DBConfig struct {
	Host     string //数据库地址
	Port     int    // 数据库端口
	DbName   string //数据库名
	Prefix   string //表前缀
	UserName string //用户名称
	Password string //用户密码
	Charset  string //编码
}
