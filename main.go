package main

import (
	"log"

	"github.com/urfave/cli"
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
