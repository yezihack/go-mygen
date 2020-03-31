package main

import (
	"log"

	"github.com/urfave/cli/v2"
)

var (
	stop chan bool //关闭信号
	app  *cli.App  //cli对象
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app = cli.NewApp()
	stop = make(chan bool)
	start()
}
