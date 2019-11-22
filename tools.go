package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

type Tools struct {
	l sync.Mutex
}

//检查某字符是否存在文件里
func (t *Tools) CheckFileContainsChar(filename, s string) bool {
	data := t.ReadFile(filename)
	if len(data) > 0 {
		return strings.LastIndex(data, s) > 0
	}
	return false
}

//读取文件内容
func (t *Tools) ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

//写文件
func (t *Tools) WriteFile(filename string, data string) (count int, err error) {
	var f *os.File
	if t.IsDirOrFileExist(filename) == false {
		f, err = os.Create(filename)
		if err != nil {
			return
		}
	} else {
		f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	}
	defer f.Close()
	count, err = io.WriteString(f, data)
	if err != nil {
		return
	}
	return
}

//追加写文件
func (t *Tools) WriteFileAppend(filename string, data string) (count int, err error) {
	var f *os.File
	if t.IsDirOrFileExist(filename) == false {
		f, err = os.Create(filename)
		if err != nil {
			return
		}
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)
	}
	defer f.Close()
	count, err = io.WriteString(f, data)
	if err != nil {
		return
	}
	return
}

//创建文件
func (t *Tools) CreateFile(path string) bool {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		return false
	}
	return true
}

//创建目录
func (t *Tools) CreateDir(path string) bool {
	if t.IsDirOrFileExist(path) == false {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}

//生成目录,不存在则创建,存在则加/
func (t *Tools) GenerateDir(path string) (string, error) {
	if len(path) == 0 {
		return "", errors.New("目录为空")
	}
	last := path[len(path)-1:]
	if !strings.EqualFold(last, string(os.PathSeparator)) {
		path = path + string(os.PathSeparator)
	}
	if !t.IsDir(path) {
		if t.CreateDir(path) {
			return path, nil
		}
		return "", errors.New(path + "创建失败或权限不足")
	}
	return path, nil
}

//判断文件 或 目录是否存在
func (t *Tools) IsDirOrFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 判断给定文件名是否是一个目录
// 如果文件名存在并且为目录则返回 true。如果 filename 是一个相对路径，则按照当前工作目录检查其相对路径。
func (t *Tools) IsDir(filename string) bool {
	return t.isFileOrDir(filename, true)
}

// 判断给定文件名是否为一个正常的文件
// 如果文件存在且为正常的文件则返回 true
func (t *Tools) IsFile(filename string) bool {
	return t.isFileOrDir(filename, false)
}

// 判断是文件还是目录，根据decideDir为true表示判断是否为目录；否则判断是否为文件
func (t *Tools) isFileOrDir(filename string, decideDir bool) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	isDir := fileInfo.IsDir()
	if decideDir {
		return isDir
	}
	return !isDir
}

//将字符串转换成驼峰格式
// Capitalize 字符首字母大写
func (t *Tools) Capitalize(s string) string {
	var upperStr string
	chars := strings.Split(s, "_")
	for _, val := range chars {
		vv := []rune(val) // 后文有介绍
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
					vv[i] -= 32 // string的码表相差32位
					upperStr += string(vv[i])
				}
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

//将分隔_拆掉,全大写
func (t *Tools) ToUpper(s string) string {
	return strings.ToUpper(s)
}

//转json
func (t *Tools) ToJson(s interface{}) string {
	js, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(js)
}


//创建目录
func CreateDir(path string) string {
	t := new(Tools)
	if t.IsDirOrFileExist(path) == false {
		b := t.CreateDir(path)
		if !b {
			log.Fatalf("创建目录:%s失败", path)
			return ""
		}
		fmt.Println("创建目录:%s成功", path)
	}
	return path
}

//写文件
func WriteFile(path, data string) (err error) {
	if _, err := new(Tools).WriteFile(path, data); err == nil {
		fmt.Printf("创建并写文件: %s成功\n", path)
		return nil
	} else {
		return errors.New(fmt.Sprintf("创建文件:%s失败", path))
	}
}

//追加写文件
func WriteAppendFile(path, data string) (err error) {
	if _, err := new(Tools).WriteFileAppend(path, data); err == nil {
		fmt.Printf("创建并追加写文件: %s成功", path)
		return nil
	} else {
		return errors.New(fmt.Sprintf("创建并追加写文件:%s失败", path))
	}
}

//显示帮助
func ShowCmdHelp() {
	for _, row := range CmdHelp {
		s := fmt.Sprintf("%s %s\n", "序号:"+row.No, row.Msg)
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

//获取操作系统
func GetOs() int {
	switch runtime.GOOS {
	case "darwin":
		return Darwin
	case "windows":
		return Window
	case "linux":
		return Linux
	default:
		return Unknown
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

