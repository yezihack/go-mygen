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
	Conn    *sql.DB   //连接对象
	stop    chan bool //关闭信号
	app     *cli.App  //cli对象
	DbConn  DBConfig  //db config
	formats []string  //format
)

//命令行实现
func start() {
	usage()
	close()
	DbConn.MaxIdleConn = 5
	DbConn.MaxOpenConn = 10
	app.Action = func(c *cli.Context) error {
		DbConn.Host = c.String("h")//数据库地址
		DbConn.Name =  c.String("u")//数据库用户名称
		DbConn.Port = c.Int("P")	//端口号
		DbConn.Pass = c.String("p")//密码
		DbConn.Charset =  c.String("c")//编码格式
		if c.NArg() > 0 {
			dbName := c.String("d")//数据库名称
			if dbName == "" {
				return cli.NewExitError("数据库名称为空, 请使用 -d dbname", 9)
			}
			DbConn.DBName = dbName
			if DbConn.Pass == "" {
				fmt.Print("输入密码>")
				line, _, err := bufio.NewReader(os.Stdin).ReadLine()
				fmt.Println(string(line))
				if err == nil {
					DbConn.Pass = string(line)
					Clean()//清屏
				}
			}
			if err := Commands(); err != nil {
				return cli.NewExitError(err, 1)
			}
		} else {
			fmt.Println(cli.VersionFlag)
			fmt.Println(cli.HelpFlag)
		}
		return nil
	}
	// cli start run
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func close() {
	go func() {
		<-stop
		Conn.Close()
	}()
}

func usage() {
	app.Name = "go-mygen" //项目名称
	app.Authors = []*cli.Author{
		{"百里", "sgfoot2020@gmail.com"},
	}
	app.Version = Version               //版本号
	app.Copyright = "@Copyright 2019"   //版权保护
	app.Usage = "快速生成操作MYSQL的CURD和文档等等" //说明
	cli.HelpFlag = &cli.BoolFlag{       //修改系统默认
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
		&cli.StringFlag{Name: "p", Value: "root", Usage: "数据库密码"},
		&cli.StringFlag{Name: "c", Value: "utf8mb4", Usage: "编码格式"},
		&cli.StringFlag{Name: "d", Usage: "数据库名称"},
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
	//初使工作
	DbModel := NewDB()
	DbModel.Using(Conn)
	DbModel.DBName = DbConn.DBName

	logic := &Logic{
		DB:   DbModel,
		Path: GetExeRootDir() + DefaultSavePath + DS, //默认当前命令所在目录
	}
	err = logic.DB.GetTableNameAndComment()
	if err != nil {
		return err
	}

	commands := NewCommands(logic)
	commands.Help(nil)
	handlers := commands.Handlers()

	br := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("input command>")
		line, _, _ := br.ReadLine()
		if len(line) == 0 {
			continue
		}
		tokens := strings.Split(string(line), " ")
		if handler, ok := handlers[strings.ToLower(tokens[0])]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("Unknown command>>", tokens[0])
		}
	}
	stop <- true
	return nil
}
