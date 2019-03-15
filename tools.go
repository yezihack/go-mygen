package gomygen

import (
	"encoding/json"
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

//判断文件 或 目录是否存在
func (t *Tools) IsDirOrFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
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
