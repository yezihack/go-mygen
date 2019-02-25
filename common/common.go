package common

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ThreeKing2018/k3log"
	"github.com/robfig/config"
	"github.com/yezihack/m2m/conf"
)

// GetRootDir 获取执行路径
func GetRootDir() string {
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
	return Substr(path, 0, strings.LastIndex(path, string(os.PathSeparator)))
}

//截取字符串
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//生成配置文件
func CreateIniFile() bool {
	iniFile := GetRootDir() + conf.DefaultIniFileName
	t := new(Tools)
	if t.IsDirOrFileExist(iniFile) {
		return true
	}
	root := GetRootPath(GetRootDir())
	iniPath := root + conf.DS + conf.TPL_CONFIG
	data := t.ReadFile(iniPath)
	if data == "" {
		k3log.Error("config内容为空", "null", "file", iniPath)
		return false
	}
	_, err := t.WriteFile(iniFile, data)
	if err != nil {
		k3log.Error("CheckIniConfig", err)
		return false
	}
	k3log.Info("ini配置文件生成成功", iniFile)
	return true
}

//读取配置文件
func ReadDbConfig() (result *conf.DBConfig, err error) {
	result = new(conf.DBConfig)
	iniFile := GetRootDir() + conf.DefaultIniFileName
	t := new(Tools)
	if t.IsDirOrFileExist(iniFile) == false {
		CreateIniFile() //生成配置
		fmt.Println("aaa")
	}
	cfg, err := config.ReadDefault(iniFile)
	if err != nil {
		return
	}
	host, err := cfg.String("mysql", "host")
	if err != nil {
		return
	}
	result.Host = host
	port, err := cfg.Int("mysql", "port")
	if err != nil {
		return
	}
	result.Port = port
	dbName, err := cfg.String("mysql", "db_name")
	if err != nil {
		return
	}
	result.DbName = dbName
	prefix, err := cfg.String("mysql", "prefix")
	if err != nil {
		return
	}
	result.Prefix = prefix
	username, err := cfg.String("mysql", "username")
	if err != nil {
		return
	}
	result.UserName = username
	password, err := cfg.String("mysql", "password")
	if err != nil {
		return
	}
	result.Password = password
	charset, err := cfg.String("mysql", "charset")
	if err != nil {
		return
	}
	result.Charset = charset
	return
}

//获取表列表
func GetConfTables() (result []string, err error) {
	iniFile := GetRootDir() + conf.DefaultIniFileName
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
	result = strings.Split(lists, ",")
	return
}
