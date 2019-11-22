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
	l Logic
}

func GetCommandHandlers() map[string]func(args []string) int {
	cmds := new(commands)
	return map[string]func(args []string) int{
		"dir":cmds.CustomDir,
		"d":cmds.CustomDir,
		"mark":cmds.MarkDown,
		"m":cmds.MarkDown,
		"entry":cmds.GenerateEntry,
		"e":cmds.GenerateEntry,
		"format":cmds.CustomFormat,
		"f":cmds.CustomFormat,
		"curd":cmds.GenerateCURD,
		"clear":cmds.Clean,
		"c": cmds.Clean,
		"quit":cmds.Quit,
		"q":cmds.Quit,
	}
}

//生成数据库表的markdown文档
func (c *commands) MarkDown(args []string) int {
	fmt.Println("正在生成markdown文档...")
	err := c.l.CreateMarkdown()
	if err != nil {
		log.Println("MarkDown>", err.Error())
	}
	return 0
}
//生成golang表对应的结构实体
func  (c *commands) GenerateEntry(args []string) int {
	fmt.Print("需要设置结构的格式字符串吗?(是:y,否:n)>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	switch string(line) {
	case "yes", "y":
		formats = c._setFormat()
	}
	err := c.l.CreateEntity(formats)
	if err != nil {
		log.Println("GenerateEntry>", err.Error())
	}
	go Gofmt(GetExeRootDir())
	return 0
}
//还可以自定义结构体解析实体,如json,gorm,xml
func  (c *commands) CustomFormat(args []string) int {
	formats = c._setFormat()
	return 0
}
//生成golang操作mysql的CRUD增删改查语句
func  (c *commands) GenerateCURD(args []string) int {
	err := c.l.CreateCURD(formats)
	if err != nil {
		log.Println("GenerateCURD>>", err.Error())
	}
	go Gofmt(GetExeRootDir())
	return 0
}
//自定义生成目录
func  (c *commands) CustomDir(args []string) int {
	fmt.Print("请指定生成目录>")
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if string(line) != "" {
		path, err := c.l.T.GenerateDir(string(line))
		if err == nil {
			c.l.Path = path
			fmt.Println("目录设置成功>", path)
		} else {
			log.Println("设置目录失败>>", err)
		}
	}
	return 0
}
//显示所有的表名
func (c *commands) ShowTableList(args []string) int {
	if len(c.l.DB.Tables) == 0 {
		fmt.Println( "呜呜,一个表也没有!!!")
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
func  (c *commands) Clean(args []string) int {
	return 0
}
//退出
func  (c *commands) Quit(args []string) int {
	return 0
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
	fmt.Print("请输入结构体的映射名称,支持多个,以逗号隔开(例:json,gorm)>")
	input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	if string(input) != "" {
		formatList := CheckCharDoSpecialArr(string(input), ',', `[\w\,\-]+`)
		if len(formatList) > 0 {
			fmt.Printf("设置值: %v, 设置成功!!!", formatList)
			return formatList
		}
	}
	fmt.Print("设置失败")
	return nil
}
