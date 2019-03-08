package logic

import (
	"errors"
	"fmt"

	"github.com/yezihack/gm2m/common"
	"github.com/yezihack/gm2m/log"
)

//创建目录
func CreateDir(path string) string {
	t := new(common.Tools)
	if t.IsDirOrFileExist(path) == false {
		b := t.CreateDir(path)
		if !b {
			log.Errorf("创建目录:%s失败\n", path)
			return ""
		}
		log.Infof("创建目录:%s成功\n", path)
	}
	return path
}

//写文件
func WriteFile(path, data string) (err error) {
	if _, err := new(common.Tools).WriteFile(path, data); err == nil {
		log.Infof("创建并写文件: %s成功\n", path)
	} else {
		log.Errorf("创建并写文件: %s失败, err: %v\n", path, err)
		return errors.New(fmt.Sprintf("创建文件:%s失败", path))
	}
	return nil
}

//追加写文件
func WriteAppendFile(path, data string) (err error) {
	if _, err := new(common.Tools).WriteFileAppend(path, data); err == nil {
		log.Infof("创建并追加写文件: %s成功\n", path)
	} else {
		return errors.New(fmt.Sprintf("创建并追加写文件:%s失败", path))
	}
	return nil
}
