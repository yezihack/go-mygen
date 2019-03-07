package main

import (
	"os"

	"strings"

	"github.com/ThreeKing2018/k3log"
	"github.com/urfave/cli"
	"github.com/yezihack/gm2m/conf"
	"github.com/yezihack/gm2m/log"
	"github.com/yezihack/gm2m/mysql"
)

func main() {
	Conn()
}

type DbConnS struct {
	Host    string
	Port    int
	User    string
	Pass    string
	DbName  string
	Charset string
}

func Conn() {
	app := cli.NewApp()
	app.Name = "gm2m"                        //项目名称
	app.Author = "百里 github.com/yezihack"    //作者名称
	app.Version = "1.0"                      //版本号
	app.Copyright = "@Copyright~2019"        //版权保护
	app.Usage = "是生成数据库表结构和markdown表结构的命令工具" //说明
	cli.HelpFlag = cli.BoolFlag{             //修改系统默认
		Name:  "help",
		Usage: "显示命令帮助",
	}
	cli.VersionFlag = cli.BoolFlag{ //修改系统默认
		Name:  "version, v",
		Usage: "显示版本号",
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "h", Value: "localhost", Usage: "数据库地址"},
		cli.IntFlag{Name: "P", Value: 3306, Usage: "端口号"},
		cli.StringFlag{Name: "u", Value: "root", Usage: "数据库用户名称"},
		cli.StringFlag{Name: "p", Value: "123456", Usage: "数据库密码"},
		cli.StringFlag{Name: "c", Value: "utf8mb4", Usage: "编码格式"},
		cli.StringFlag{Name: "d", Value: "", Usage: "*数据库名称"},
	}
	var DbConn conf.DBConfig
	app.Action = func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			cli.ShowAppHelp(c)
			return nil
		}
		//数据库地址
		host := c.String("h")
		if strings.EqualFold(host, "") {
			host = "localhost"
		}
		DbConn.Host = host
		//端口号
		port := c.Int("P")
		if port == 0 {
			port = 3306
		}
		DbConn.Port = port
		//数据库用户名称
		user := c.String("u")
		if user == "" {
			user = "root"
		}
		DbConn.UserName = user
		//数据库密码
		pass := c.String("p")
		if pass == "" {
			pass = "123456"
		}
		DbConn.Password = pass
		//编码格式
		charset := c.String("c")
		if charset == "" {
			charset = "utf8mb4"
		}
		DbConn.Charset = charset
		//数据库名称
		dbname := c.String("d")
		if dbname == "" {
			return cli.NewExitError("数据库名称为空, 请使用 -d dbname", 9)
		}
		DbConn.DbName = dbname
		return nil
	}
	var err error
	err = app.Run(os.Args)
	if err != nil {
		k3log.Error(err)
	}
	if DbConn.DbName != "" {
		err = mysql.SetDbConn(DbConn)
		if err != nil {
			log.Error(err)
			return
		}

	}
}
