package main

import (
	"github.com/urfave/cli"
	"log"
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
