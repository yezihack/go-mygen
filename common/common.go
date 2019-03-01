package common

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"os/exec"

	"github.com/ThreeKing2018/k3log"
	"github.com/robfig/config"
	"github.com/yezihack/gm2m/conf"
)

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

//生成配置文件
func CreateIniFile() (path string, err error) {
	path = GetExeRootDir() + conf.DefaultIniFileName
	t := new(Tools)
	if t.IsDirOrFileExist(path) {
		return
	}
	root := GetRootPath(GetExeRootDir())
	pathTpl := root + conf.DS + conf.TPL_CONFIG
	data := t.ReadFile(pathTpl)
	if data == "" {
		k3log.Error("config内容为空", "null", "file", path)
		return
	}
	_, err = t.WriteFile(path, data)
	if err != nil {
		k3log.Error("CheckIniConfig", err)
		return
	}
	return
}

//读取配置文件
func ReadDbConfig(iniFile string) (result *conf.DBConfig, err error) {
	result = new(conf.DBConfig)
	cfg, err := config.ReadDefault(iniFile)
	if err != nil {
		return
	}
	host, err := cfg.String("mysql", "host")
	if err != nil {
		return
	}
	result.Host = strings.TrimSpace(host)
	port, err := cfg.String("mysql", "port")
	if err != nil {
		return
	}
	result.Port = port
	dbName, err := cfg.String("mysql", "db_name")
	if err != nil {
		return
	}
	result.DbName = strings.TrimSpace(dbName)

	username, err := cfg.String("mysql", "username")
	if err != nil {
		return
	}
	result.UserName = strings.TrimSpace(username)
	password, err := cfg.String("mysql", "password")
	if err != nil {
		return
	}
	result.Password = strings.TrimSpace(password)
	charset, err := cfg.String("mysql", "charset")
	if err != nil {
		return
	}
	result.Charset = strings.TrimSpace(charset)
	return
}

//获取表列表
func GetConfTables() (result []string, err error) {
	iniFile := GetExeRootDir() + conf.DefaultIniFileName
	t := new(Tools)
	if t.IsDirOrFileExist(iniFile) == false {
		err = errors.New("ini文件不存在")
		return
	}
	cfg, err := config.ReadDefault(iniFile)
	if err != nil {
		return
	}
	var lists string
	lists, err = cfg.String("tables", "list")
	if err != nil {
		return
	}
	if lists == "" {
		return
	}
	result = strings.Split(strings.TrimSpace(lists), ",")
	return
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
		cmd := exec.Command("goimports", "-l", "-w", path)
		_, err := cmd.Output()
		if err != nil {
			panic(err)
			return false
		}
		return true
	}
	return false
}
