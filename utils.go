package gomygen

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/ThreeKing2018/gocolor"
	"github.com/yezihack/colorlog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

//创建目录
func CreateDir(path string) string {
	t := new(Tools)
	if t.IsDirOrFileExist(path) == false {
		b := t.CreateDir(path)
		if !b {
			colorlog.Error("创建目录:%s失败", path)
			return ""
		}
		colorlog.Info("创建目录:%s成功", path)
	}
	return path
}

//写文件
func WriteFile(path, data string) (err error) {
	if _, err := new(Tools).WriteFile(path, data); err == nil {
		colorlog.Info("创建并写文件: %s成功", path)
		return nil
	} else {
		colorlog.Error("创建并写文件: %s失败, err: %v", path, err)
		return errors.New(fmt.Sprintf("创建文件:%s失败", path))
	}
}

//追加写文件
func WriteAppendFile(path, data string) (err error) {
	if _, err := new(Tools).WriteFileAppend(path, data); err == nil {
		colorlog.Info("创建并追加写文件: %s成功", path)
		return nil
	} else {
		return errors.New(fmt.Sprintf("创建并追加写文件:%s失败", path))
	}
}

//显示帮助
func ShowCmdHelp() {
	for _, row := range CmdHelp {
		s := fmt.Sprintf("%s %s\n", gocolor.SYellow("序号:"+row.No), gocolor.SBlue(row.Msg))
		fmt.Print(s)
	}
}

//GetRootDir 获取执行路径
func GetExeRootDir() string {
	// 文件不存在获取执行路径
	file, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file = fmt.Sprintf("%s%s", file, string(os.PathSeparator))
	}
	return file
}

//获取根目录
func GetRootPath(path string) string {
	if path[len(path)-1:] == string(os.PathSeparator) {
		path = path[:len(path)-1]
	}
	return SubStr(path, 0, strings.LastIndex(path, string(os.PathSeparator)))
}

//截取字符串
func SubStr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func ErrMsg(msg string, err error) interface{} {
	m := make(map[string]interface{}, 2)
	m["msg"] = msg
	m["err"] = err
	return m
}

//FMT 格式代码
func Gofmt(path string) bool {
	if new(Tools).IsDirOrFileExist(path) {
		if !ExecCommand("goimports", "-l", "-w", path) {
			if !ExecCommand("gofmt", "-l", "-w", path) {
				return ExecCommand("go", "fmt", path)
			}
		}
		return true
	}
	return false
}

//清屏
func Clean() {
	switch GetOs() {
	case Darwin, Linux:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case Window:
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func GetOs() int {
	switch runtime.GOOS {
	case "darwin":
		return Darwin
	case "windows":
		return Window
	case "linux":
		return Linux
	default:
		return Unknow
	}
}

func ExecCommand(name string, args ...string) bool {
	cmd := exec.Command(name, args...)
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	return true
}

//拼接特殊字符串
func FormatField(field string, formats []string) string {
	if len(formats) == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	for key := range formats {
		buf.WriteString(fmt.Sprintf(`%s:"%s" `, formats[key], field))
	}
	return "`" + strings.TrimRight(string(buf.String()), " ") + "`"
}

//添加注释 //
func AddToComment(s string, suff string) string {
	if strings.EqualFold(s, "") {
		return ""
	}
	return "//" + s + suff
}

//判断是否包存在某字符
func InArrayString(str string, arr []string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}

//检查字符串,去掉特殊字符
func CheckCharDoSpecial(s string, char byte, regs string) string {
	reg := regexp.MustCompile(regs)
	var result []string
	if arr := reg.FindAllString(s, -1); len(arr) > 0 {
		buf := bytes.Buffer{}
		for key, val := range arr {
			if val != string(char) {
				buf.WriteString(val)
			}
			if val == string(char) && buf.Len() > 0 {
				result = append(result, buf.String())
				buf.Reset()
			}
			//处理最后一批数据
			if buf.Len() > 0 && key == len(arr)-1 {
				result = append(result, buf.String())
			}
		}
	}
	return strings.Join(result, string(char))
}
func CheckCharDoSpecialArr(s string, char byte, reg string) []string {
	s = CheckCharDoSpecial(s, char, reg)
	return strings.Split(s, string(char))
}
