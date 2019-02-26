package logic

import (
	"errors"
	"fmt"

	"github.com/ThreeKing2018/k3log"
	"github.com/yezihack/gm2m/common"
)

//创建目录
func CreateDir(path string) string {
	t := new(common.Tools)
	if t.IsDirOrFileExist(path) == false {
		b := t.CreateDir(path)
		if !b {
			k3log.Info(fmt.Sprintf("创建目录:%s失败\n", path))
			return ""
		}
		k3log.Info(fmt.Sprintf("创建目录:%s成功\n", path))
	}
	return path
}

//写文件
func WriteFile(path, data string) (err error) {
	if _, err := new(common.Tools).WriteFile(path, data); err == nil {
		k3log.Info(fmt.Sprintf("创建并写文件: %s成功\n", path))
	} else {
		k3log.Info("创建并写文件: %s失败, err: %v\n", path, err)
		return errors.New(fmt.Sprintf("创建文件:%s失败", path))
	}
	return nil
}

//追加写文件
func WriteAppendFile(path, data string) (err error) {
	if _, err := new(common.Tools).WriteFileAppend(path, data); err == nil {
		k3log.Info(fmt.Sprintf("创建并追加写文件: %s成功\n", path))
	} else {
		k3log.Info(fmt.Sprintf("创建并追加写文件: %s失败, err: %v\n", path, err))
		return errors.New(fmt.Sprintf("创建并追加写文件:%s失败", path))
	}
	return nil
}
