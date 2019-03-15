package main

import (
	"os"

	"strings"

	"bufio"

	"github.com/ThreeKing2018/gocolor"
	"github.com/urfave/cli"
	"github.com/yezihack/colorlog"
	"github.com/yezihack/gm2m/common"
	"github.com/yezihack/gm2m/logic"
	"github.com/yezihack/gm2m/mysql"
)

func main() {
	Conn()
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
		err := Commands(DbConn)
		if err != nil {
			return cli.NewExitError(err, 9)
		}
		return nil
	}
	var err error
	err = app.Run(os.Args)
	if err != nil {
		colorlog.Error("Err", err)
	}
}

func Commands(DbConn mysql.DBConfig) error {
	db, err := mysql.InitDB(DbConn)
	if db == nil || err != nil {
		return err
	}
	colorlog.Info("数据库连接成功")
	masterDB := mysql.NewDB(db)
	logicModel := logic.Logic{
		DB: masterDB,
	}
	var input []byte
	logic.ShowCmdHelp() //显示帮助
	formatList := make([]string, 0)
	for {
		var err error
		gocolor.Cyan("请输入以上命令序号:")
		input, _, err = bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			os.Exit(9)
		}
		switch string(input) {
		case "1": //生成表markdown文档
			err = logicModel.CreateMarkdown()
		case "2": //生成表结构数据
			err = logicModel.CreateEntity(formatList)
			common.Gofmt(common.GetExeRootDir())
		case "3": //生成CURD增删改查
			err = logicModel.CreateCURD(formatList)
			common.Gofmt(common.GetExeRootDir())
		case "4": //设置结构体的映射名称,支持多个,以逗号隔开
			gocolor.Blue("请输入结构体的映射名称(例:json):")
			input, _, err = bufio.NewReader(os.Stdin).ReadLine()
			if string(input) != "" {
				formatList = common.CheckCharDoSpecialArr(string(input), ',', `[\w\,\-]+`)
				if len(formatList) > 0 {
					colorlog.Info("设置值: %v, 设置成功!!!", formatList)
				}
			}
		case "7", "h", "": //显示帮助
			logic.ShowCmdHelp()
		case "8", "c", "clear": //清屏
			common.Clean()
		case "9", "e", "exit": //退出
			os.Exit(9)
		default:
			colorlog.Warn("命令输入有错误!!!")
		}
		if err != nil {
			return err
		}
	}
	return nil
}
//output/gm2m -h localhost -P 3308 -u root -p 123456 -d kindled
//v2 输出指定位置
