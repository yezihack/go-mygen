package logic

import (
	"errors"
	"fmt"

	"github.com/ThreeKing2018/gocolor"
	"github.com/yezihack/colorlog"
	"github.com/yezihack/gm2m/common"
	"github.com/yezihack/gm2m/conf"
)

//创建目录
func CreateDir(path string) string {
	t := new(common.Tools)
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
	if _, err := new(common.Tools).WriteFile(path, data); err == nil {
		colorlog.Info("创建并写文件: %s成功", path)
		return nil
	} else {
		colorlog.Error("创建并写文件: %s失败, err: %v", path, err)
		return errors.New(fmt.Sprintf("创建文件:%s失败", path))
	}
}

//追加写文件
func WriteAppendFile(path, data string) (err error) {
	if _, err := new(common.Tools).WriteFileAppend(path, data); err == nil {
		colorlog.Info("创建并追加写文件: %s成功", path)
		return nil
	} else {
		return errors.New(fmt.Sprintf("创建并追加写文件:%s失败", path))
	}
}

//显示帮助
func ShowCmdHelp() {
	for _, row := range conf.CmdHelp {
		s := fmt.Sprintf("%s %s\n", gocolor.SYellow("序号:"+row.No), gocolor.SBlue(row.Msg))
		fmt.Print(s)
	}
}
