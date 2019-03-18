package gomygen

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Tools struct {
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
