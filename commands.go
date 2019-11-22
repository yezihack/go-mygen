package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type commands struct {
	l *Logic
}

func NewCommands(logic *Logic) *commands {
	return &commands{
		l: logic,
	}
}

//映射相应的命令
func (c *commands) Handlers() map[string]func(args []string) int {
	return map[string]func(args []string) int{
		"0":     c.CustomDir,
		"1":     c.MarkDown,
		"2":     c.GenerateEntry,
		"3":     c.GenerateCURD,
		"4":     c.CustomFormat,
		"5":     c.ShowTableList,
		"7":     c.Clean,
		"clear": c.Clean,
		"c":     c.Clean,
		"8":     c.Help,
		"h":     c.Help,
		"help":  c.Help,
		"quit":  c.Quit,
		"q":     c.Quit,
	}
}

//生成数据库表的markdown文档
func (c *commands) MarkDown(args []string) int {
	fmt.Println("正在准备生成markdown文档...")
	//检查目录是否存在
	CreateDir(c.l.Path)
	err := c.l.CreateMarkdown()
	if err != nil {
		log.Println("MarkDown>>", err)
	}
	return 0
}

//help list
func (c *commands) Help(args []string) int {
	for _, row := range CmdHelp {
		s := fmt.Sprintf("%s %s\n", "NO:"+row.No, row.Msg)
		fmt.Print(s)
	}
	return 0
}

//生成golang表对应的结构实体
func (c *commands) GenerateEntry(args []string) int {
	fmt.Print("需要设置结构的格式字符串吗?(是:y,否:n)>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	switch string(line) {
	case "yes", "y":
		formats = c._setFormat()
	}
	err := c.l.CreateEntity(formats)
	if err != nil {
		log.Println("GenerateEntry>>", err.Error())
	}
	go Gofmt(GetExeRootDir())
	return 0
}

//还可以自定义结构体解析实体,如json,gorm,xml
func (c *commands) CustomFormat(args []string) int {
	formats = c._setFormat()
	return 0
}

//生成golang操作mysql的CRUD增删改查语句
func (c *commands) GenerateCURD(args []string) int {
	err := c.l.CreateCURD(formats)
	if err != nil {
		log.Println("GenerateCURD>>", err.Error())
	}
	go Gofmt(GetExeRootDir())
	return 0
}

//自定义生成目录
func (c *commands) CustomDir(args []string) int {
	fmt.Print("请指定生成目录>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if string(line) != "" {
		path, err := c.l.T.GenerateDir(string(line))
		if err == nil {
			c.l.Path = path
			fmt.Println("目录设置成功:", path)
		} else {
			log.Println("设置目录失败>>", err)
		}
	}
	return 0
}

//显示所有的表名
func (c *commands) ShowTableList(args []string) int {
	if len(c.l.DB.Tables) == 0 {
		fmt.Println("呜呜,一个表也没有!!!")
		return 0
	}
	c._showTableList(c.l.DB.Tables)
	fmt.Print("选择你需要的表序列号?(默认全部,逗号隔开,all代表全部)>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if !strings.EqualFold(string(line), "") {
		c.l.DB.DoTables = c._filterTables(string(line), c.l.DB.Tables)
	}
	return 0
}

//清屏
func (c *commands) Clean(args []string) int {
	Clean()
	return 0
}

//退出
func (c *commands) Quit(args []string) int {
	return 1
}

//过滤表名
func (c *commands) _filterTables(ids string, tables []TableNameAndComment) []TableNameAndComment {
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
func (c *commands) _showTableList(NameAndComment []TableNameAndComment) {
	for idx, table := range NameAndComment {
		idx++
		info := fmt.Sprintf("%s:%s", strconv.Itoa(idx), table.Name)
		if table.Comment != "" {
			info += fmt.Sprintf("(%s)", table.Comment)
		}
		fmt.Println(info)
	}
	fmt.Println("共" + strconv.Itoa(len(NameAndComment)) + "张表\n")
}

//set struct format
func (c *commands) _setFormat() []string {
	fmt.Print("设置结构体的映射名称,以逗号隔开(例:json,gorm)>")
	input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if string(input) != "" {
		formatList := CheckCharDoSpecialArr(string(input), ',', `[\w\,\-]+`)
		if len(formatList) > 0 {
			fmt.Printf("Set format success: %v\n", formatList)
			return formatList
		}
	}
	fmt.Println("设置失败")
	return nil
}
