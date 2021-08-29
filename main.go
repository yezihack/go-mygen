package main

import (
	"github.com/urfave/cli/v2"
	"github.com/yezihack/go-mygen/internal/mygen"
)

var (
	stop chan bool // 关闭信号
	app  *cli.App  // cli对象
)

func main() {
	app = cli.NewApp()
	stop = make(chan bool)
	mygen.start()
}
