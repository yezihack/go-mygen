package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)
var (
	Conn *sql.DB //连接对象
	stop chan bool //关闭信号
	app *cli.App //cli对象
	DbConn DBConfig //db config
	formats []string //format
)

//命令行实现
func Start() {
	app = cli.NewApp()
	stop = make(chan bool)
	Usage()
	close()
	DbConn.MaxIdleConn = 5
	DbConn.MaxOpenConn = 10
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
			fmt.Print("输入密码>")
			line, _, err := bufio.NewReader(os.Stdin).ReadLine()
			fmt.Println(string(line))
			if err == nil {
				pass = string(line)
				//清屏
				Clean()
			}
		}
		DbConn.Pass = pass
		//编码格式
		charset := c.String("c")
		if charset == "" {
			charset = "utf8mb4"
		}
		DbConn.Charset = charset
		//数据库名称
		dbName := c.String("d")
		if dbName == "" {
			return cli.NewExitError("数据库名称为空, 请使用 -d dbname", 9)
		}
		DbConn.DBName = dbName
		err := Commands()
		if err != nil {
			return cli.NewExitError(err, 9)
		}
		return nil
	}
	var err error
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func close() {
	go func() {
		<-stop
		Conn.Close()
	}()
}

func Usage() {
	app.Name = "go-mygen"               //项目名称
	app.Authors = []*cli.Author{{Name:"百里", Email:"sgfoot2020@gmail.com"}}                  //作者名称
	app.Version = Version   //版本号
	app.Copyright = "@Copyright 2019"   //版权保护
	app.Usage = "快速生成操作MYSQL的CURD和文档等等" //说明
	cli.HelpFlag = &cli.BoolFlag{        //修改系统默认
		Name:  "help, h",
		Usage: "显示命令帮助",
	}
	cli.VersionFlag = &cli.BoolFlag{ //修改系统默认
		Name:  "version, v",
		Usage: "显示版本号",
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "h", Value: "localhost", Usage: "数据库地址"},
		&cli.IntFlag{Name: "P", Value: 3306, Usage: "端口号"},
		&cli.StringFlag{Name: "u", Value: "root", Usage: "数据库用户名称"},
		&cli.StringFlag{Name: "p", Value: "", Usage: "数据库密码"},
		&cli.StringFlag{Name: "c", Value: "utf8mb4", Usage: "编码格式"},
		&cli.StringFlag{Name: "d", Value: "", Usage: "*数据库名称"},
	}
}
//实现命令功能
func Commands() error {
	var err error
	Conn, err = InitDB(DbConn)
	if Conn == nil || err != nil {
		return errors.New("数据库连接失败:" + err.Error())
	}
	log.Println("数据库连接成功")
	masterDB := NewDB()
	masterDB.Using(Conn)
	masterDB.DBName = DbConn.DBName
	lg := Logic{
		DB:   masterDB,
		Path: GetExeRootDir(), //默认当前命令所在目录
	}
	err = lg.DB.GetTableNameAndComment()
	if err != nil {
		return err
	}
	ShowCmdHelp() //显示帮助
	br := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandlers()
	for {
		fmt.Print("input command>")
		line, _, _ := br.ReadLine()
		tokens := strings.Split(string(line), " ")
		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("Unknown command>>", tokens[0])
		}
	}
	return nil
}
