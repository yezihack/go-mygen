package main

import (
	"os"

	"strings"

	"fmt"

	"bufio"

	"github.com/urfave/cli"
	"github.com/yezihack/gm2m/common"
	"github.com/yezihack/gm2m/conf"
	"github.com/yezihack/gm2m/log"
	"github.com/yezihack/gm2m/logic"
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
	var DbConn mysql.DBConfig
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
		DbConn.Name = user
		//数据库密码
		pass := c.String("p")
		if pass == "" {
			pass = "123456"
		}
		DbConn.Pass = pass
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
		DbConn.DBName = dbname
		return nil
	}
	var err error
	err = app.Run(os.Args)
	if err != nil {
		log.Error(err)
	}
	if DbConn.DBName != "" {
		db, err := mysql.InitDB(DbConn)
		if err != nil {
			log.Error(err)
			return
		}
		masterDB := mysql.NewDB(db)
		cmdLine := []conf.CmdS{
			{"1", "生成表markdown文档"},
			{"2", "生成表结构数据"},
			{"3", "生成CURD增删改查"},
			{"h", "查看帮助"},
			{"e,exit", "退出"},
		}
		var showCmd = func() {
			for _, row := range cmdLine {
				fmt.Printf("序列号:%s, %s\n", row.No, row.Msg)
			}
		}
		logicModel := logic.Logic{
			DB: masterDB,
		}
		var number byte
		for {
			showCmd()
			var err error
			fmt.Print("请输入以上命令序号:")
			r := bufio.NewReader(os.Stdin)
			number, err = r.ReadByte()
			if err != nil {
				os.Exit(9)
			}
			switch string(number) {
			case "1":
				err = logicModel.CreateMarkdown()
				//gofmt.GoFmt()
			case "2":
				err = logicModel.CreateStructure()
				common.Gofmt(common.GetExeRootDir())
			case "3":
				err = logicModel.CreateCRUD()
				common.Gofmt(common.GetExeRootDir())
			case "h":
				showCmd()
			case "e", "exit":
				os.Exit(9)
			default:
				log.Warn("命令输入有错误!!!")
			}

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

//output/gm2m -h localhost -P 3308 -u root -p 123456 -d kindled
//v2 输出指定位置
