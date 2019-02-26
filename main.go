package main

import (
	"fmt"
	"os"

	"github.com/ThreeKing2018/k3log"
	kconf "github.com/ThreeKing2018/k3log/conf"
	"github.com/urfave/cli"
	"github.com/yezihack/gm2m/common"
	"github.com/yezihack/gm2m/conf"
	"github.com/yezihack/gm2m/logic"
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//初使日志
	k3log.SetLogger(kconf.WithFilename(common.GetRootDir()+"log/m2m.log"),
		kconf.WithIsStdOut(true),
		kconf.WithProjectName("gm2m"),
		kconf.WithLogType(kconf.LogNormalType))
	defer k3log.Sync()
}

func main() {
	Start()
	//lg := new(logic.Logic)
	//lg.CreateCRUD()
	//lg.CreateMarkdown()
	//test()
}

var lang string

func test() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "language for the greeting`FILE`",
			Destination: &lang,
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "Nefertiti"
		fmt.Println("count", c.NArg(), "fist", c.Args().First())
		if c.NArg() > 0 {
			name = c.Args().Get(1)
		}
		if c.String("lang") == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		fmt.Println("lang", lang)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		k3log.Error(err)
	}
}

func Start() {
	app := cli.NewApp()
	app.Name = "gm2m"                        //项目名称
	app.Author = "百里"                        //作者名称
	app.Version = "1.0"                      //版本号
	app.Copyright = "@Copyright~2019"        //版权保护
	app.Usage = "是生成数据库表结构和markdown表结构的命令工具" //说明
	cli.HelpFlag = cli.BoolFlag{             //修改系统默认
		Name:  "help, h",
		Usage: "显示命令帮助",
	}
	cli.VersionFlag = cli.BoolFlag{ //修改系统默认
		Name:  "version, v",
		Usage: "显示版本号",
	}
	app.Flags = []cli.Flag{ //标识参数切片,可填写多个标识参数.例: gm2m -config g.ini
		cli.StringFlag{ //如何获取? 添加一个Destination, 或者 cli.Action, c.String("config"
			Name:   "config, c",                  //名称,逗号隔开,相当于别名
			Value:  "default.ini",                //默认值
			Hidden: false,                        //是否显示这命令
			EnvVar: "GM2M_CONFIG",                //环境变量
			Usage:  "配置文件,默认当前目录下,否则全路径 `FILE` ", //说明
		},
	}
	app.Commands = []cli.Command{ //命令参数
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "生成配置文件",
			Action: func(c *cli.Context) error {
				if common.CreateIniFile() {
					k3log.Info("生成配置文件完毕!")
				} else {
					k3log.Info("生成配置文件失败!")
				}
				return nil
			},
		},
		{
			Name:    "structure",
			Aliases: []string{"s"},
			Usage:   "生成golang数据结构数据",
			Action: func(c *cli.Context) error {
				err := new(logic.Logic).CreateCRUD()
				if err != nil {
					k3log.Error("生成golang数据结构数据", err)
					return nil
				}
				fmt.Println("生成golang数据结构数据完毕")
				return nil
			},
		},
		{
			Name:    "markdown",
			Aliases: []string{"m"},
			Usage:   "生成表结构的markdown文档",
			Action: func(c *cli.Context) error {
				err := new(logic.Logic).CreateMarkdown()
				if err != nil {
					k3log.Error("生成表结构的markdown文档", err)
					return nil
				}
				fmt.Println("生成表结构的markdown文档完毕")
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			cli.ShowAppHelp(c)
			return nil
		}
		os.Setenv(conf.TMP_ENV_INI_FILE, c.String("config"))
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		k3log.Error(err)
	}
}

func StudyCli() {
	app := cli.NewApp()
	app.Name = "gm2m"                        //项目名称
	app.Author = "百里"                        //作者名称
	app.Version = "1.0"                      //版本号
	app.Copyright = "@Copyright~2019"        //版权保护
	app.Usage = "是生成数据库表结构和markdown表结构的命令工具" //说明
	app.Flags = []cli.Flag{                  //标识参数切片,可填写多个标识参数.例: gm2m -config g.ini
		cli.StringFlag{
			Name:   "config, c",          //名称,逗号隔开,相当于别名
			Value:  "default.ini",        //默认值
			EnvVar: "GM2M_CONFIG",        //环境变量
			Usage:  "配置文件,默认当前目录下,否则全路径", //说明
		},
	}
	app.Commands = []cli.Command{ //命令参数
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "生成配置文件",
			Action: func(c *cli.Context) error {
				if common.CreateIniFile() {
					k3log.Info("生成配置文件完毕!")
				} else {
					k3log.Info("生成配置文件失败!")
				}
				return nil
			},
			After: func(c *cli.Context) error {
				fmt.Println("after")
				return nil
			},
			Before: func(c *cli.Context) error {
				fmt.Println("first", c.Args().First())
				fmt.Println("Before")
				if !c.Bool("ginger-crouton") {
					return cli.NewExitError("正确使用命令", 100) //会提示命令如何使用
				}
				return nil
			},
		},
		{
			Name:    "structure",
			Aliases: []string{"s"},
			Usage:   "生成golang数据结构数据",
			Action: func(c *cli.Context) error {
				fmt.Println("生成golang数据结构数据")
				return nil
			},
		},
		{
			Name:    "markdown",
			Aliases: []string{"m"},
			Usage:   "生成表结构的markdown文档",
			Action: func(c *cli.Context) error {
				fmt.Println("生成表结构的markdown文档")
				return nil
			},
		},
		{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "显示命令帮助",
			Action: func(c *cli.Context) error {
				cli.ShowAppHelp(c)
				return nil
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			configPath := c.Args().Get(0)
			fmt.Println("configPath", configPath)
		} else {
			cli.ShowAppHelp(c)
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		k3log.Panic(err)
	}
}
