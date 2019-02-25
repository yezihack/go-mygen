package main

import (
	"fmt"

	"github.com/ThreeKing2018/k3log"
	"github.com/ThreeKing2018/k3log/conf"
	"github.com/yezihack/m2m/common"
	"github.com/yezihack/m2m/logic"
)

func init() {
	//初使日志
	k3log.SetLogger(conf.WithFilename("./k3log/m2m.log"),
		conf.WithIsStdOut(true),
		conf.WithProjectName("goM2M"),
		conf.WithLogType(conf.LogNormalType))
	defer k3log.Sync()
}

func main() {
	fmt.Println(common.GetRootDir())
	lg := new(logic.Logic)
	lg.CreateCRUD()
	lg.CreateMarkdown()
}
