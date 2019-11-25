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
		if c.NumFlags() > 0 {
			dbName := c.String("d")//数据库名称
			if dbName == "" {
				return cli.NewExitError("database is null, please use -d params", 9)
			}
			DbConn.DBName = dbName
			if DbConn.Pass == "" {
				fmt.Print("input password>")
				line, _, err := bufio.NewReader(os.Stdin).ReadLine()
				if err == nil {
					DbConn.Pass = string(line)
					Clean()//清屏
				}
			}
			if err := Commands(); err != nil {
				return cli.NewExitError(err, 1)
			}
		}
		return nil
	}
	// cli start run
	if err := app.Run(os.Args); err != nil {
		log.Fatal("RUN>>", err)
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
		{"sgfoot", "sgfoot2020@gmail.com"},
	}
	app.Version = Version               //版本号
	app.Copyright = "@Copyright 2019"   //版权保护
	app.Usage = "Quickly generate CURD and documentation for operating MYSQL.etc" //说明
	cli.HelpFlag = &cli.BoolFlag{       //修改系统默认
		Name:  "help",
		Usage: "show help",
	}
	cli.VersionFlag = &cli.BoolFlag{ //修改系统默认
		Name:  "version, v",
		Usage: "print the version",
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "h",Value:"127.0.0.1",Usage: "Database address", },
		&cli.IntFlag{Name: "P", Value: 3306, Usage: "port number"},
		&cli.StringFlag{Name: "u", Value: "root", Usage: "database username"},
		&cli.StringFlag{Name: "p", Value: "root", Usage: "database password"},
		&cli.StringFlag{Name: "c", Value: "utf8mb4", Usage: "database format"},
		&cli.StringFlag{Name: "d", Value: "database name"},
	}
}

//实现命令功能
func Commands() error {
	var err error
	Conn, err = InitDB(DbConn)
	if Conn == nil || err != nil {
		return errors.New("database connect failed>>" + err.Error())
	}
	log.Println("Database connected success")
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
