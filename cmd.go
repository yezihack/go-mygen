package gomygen

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ThreeKing2018/gocolor"
	"github.com/ThreeKing2018/k3log"
	"github.com/urfave/cli"
	"github.com/yezihack/colorlog"
	"runtime/debug"
)

const (
	Version   = "v1.1.5"
	UpdatedAt = "  2019.08.19"
	LogFile   = "/tmp/gomygen.log"
)

//命令行实现
func Cmd() {
	k3log.NewProduction("go-mygen", LogFile)
	defer func() {
		if err := recover(); err != nil {
			k3log.Error("err", err, "stack", string(debug.Stack()))
			colorlog.Error("查看错误日志:"+LogFile+", %v", err)
		}
	}()
	err1 := WriteFile(LogFile, "")
	if err1 != nil {
		os.Exit(0)
	}
	app := cli.NewApp()
	app.Name = "go-mygen"               //项目名称
	app.Author = "百里"                   //作者名称
	app.Version = Version + UpdatedAt   //版本号
	app.Copyright = "@Copyright 2019"   //版权保护
	app.Usage = "快速生成操作MYSQL的CURD和文档等等" //说明
	cli.HelpFlag = cli.BoolFlag{        //修改系统默认
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
		cli.StringFlag{Name: "p", Value: "", Usage: "数据库密码"},
		cli.StringFlag{Name: "c", Value: "utf8mb4", Usage: "编码格式"},
		cli.StringFlag{Name: "d", Value: "", Usage: "*数据库名称"},
	}
	var DbConn DBConfig
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
			gocolor.Blue("输入密码:")
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
		colorlog.Error("%v", err)
	}
}

//实现命令功能
func Commands(DbConn DBConfig) error {
	DbConn.MaxIdleConn = 5
	DbConn.MaxOpenConn = 10
	db, err := InitDB(DbConn)
	if db == nil || err != nil {
		colorlog.Error("数据库连接失败")
		return err
	}
	defer db.Close()
	colorlog.Info("数据库连接成功")
	masterDB := NewDB()
	masterDB.Using(db)
	masterDB.DBName = DbConn.DBName
	lg := Logic{
		DB:   masterDB,
		Path: GetExeRootDir(), //默认当前命令所在目录
	}
	err = lg.DB.GetTableNameAndComment()
	if err != nil {
		return err
	}

	var input []byte
	ShowCmdHelp() //显示帮助
	formatList := make([]string, 0)
	for {
		var err error
		gocolor.Cyan("请输入以上命令序号:")
		input, _, err = bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			os.Exit(9)
		}
		switch string(input) {
		case "0": //设置指定生成目录
			gocolor.Blue("请指定生成目录:")
			input, _, err = bufio.NewReader(os.Stdin).ReadLine()
			if string(input) != "" {
				path, err := lg.T.GenerateDir(string(input))
				if err == nil {
					lg.Path = path
					colorlog.Info("目录设置成功:%s", path)
				} else {
					colorlog.Warn("设置目录失败,原因:%v", err)
				}
			}
		case "1": //生成表markdown文档
			fmt.Println("正在生成markdown文档...")
			err = lg.CreateMarkdown()
		case "2": //生成表结构数据
			gocolor.Blue("需要设置结构的格式字符串吗?(是:y,否:n):")
			input, _, err = bufio.NewReader(os.Stdin).ReadLine()
			switch string(input) {
			case "yes", "y":
				formatList = setFormat()
			}
			err = lg.CreateEntity(formatList)
			go Gofmt(GetExeRootDir())
		case "3": //生成CURD增删改查
			err = lg.CreateCURD(formatList)
			go Gofmt(GetExeRootDir())
		case "4": //设置结构体的映射名称,支持多个,以逗号隔开
			formatList = setFormat()
		case "5": //列出所有的表名
			if len(lg.DB.Tables) == 0 {
				gocolor.Green("%s\n", "呜呜,一个表也没有!!!")
				continue
			}
			ShowTableList(lg.DB.Tables)
			gocolor.Blue("选择你需要的表序列号?(默认全部,逗号隔开,all代表全部):")
			input, _, err = bufio.NewReader(os.Stdin).ReadLine()
			if !strings.EqualFold(string(input), "") {
				lg.DB.DoTables = filterTables(string(input), lg.DB.Tables)
			}
		case "7", "h", "", "ll": //显示帮助
			ShowCmdHelp()
		case "8", "c", "clear": //清屏
			Clean()
		case "9", "e", "q", "exit": //退出
			os.Exit(0)
		default:
			colorlog.Warn("命令输入有错误!!!")
		}
		if err != nil {
			colorlog.Error("出错了, 参考:", error.Error(err))
		}
	}
	return nil
}

//过滤表名
func filterTables(ids string, tables []TableNameAndComment) []TableNameAndComment {
	lst := strings.Split(ids, ",")
	result := make([]TableNameAndComment, 0)
	if strings.ToLower(ids) == "all" {
		return tables
	}
	for _, id := range lst {
		id = strings.TrimSpace(id)
		for _, t := range tables {
			if strconv.Itoa(t.Index) == id || id == t.Name {
				result = append(result, t)
			}
		}
	}
	return result
}

//显示所有名视图
func ShowTableList(NameAndComment []TableNameAndComment) {
	for idx, table := range NameAndComment {
		idx++
		info := fmt.Sprintf("%s:%s", gocolor.SBlueBG(strconv.Itoa(idx)), gocolor.SGreen(table.Name))
		if table.Comment != "" {
			info += fmt.Sprintf("(%s)", table.Comment)
		}
		gocolor.Blue(info)
		fmt.Println()
	}
	gocolor.Green("共" + strconv.Itoa(len(NameAndComment)) + "张表\n")
}

func setFormat() []string {
	gocolor.Blue("请输入结构体的映射名称,支持多个,以逗号隔开(例:json,gorm):")
	input, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		colorlog.Warn("%v", err)
		return nil
	}
	if string(input) != "" {
		formatList := CheckCharDoSpecialArr(string(input), ',', `[\w\,\-]+`)
		if len(formatList) > 0 {
			colorlog.Info("设置值: %v, 设置成功!!!", formatList)
			return formatList
		}
	}
	colorlog.Warn("设置失败")
	return nil
}
