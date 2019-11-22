package main

import (
	"github.com/urfave/cli"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app = cli.NewApp()
	stop = make(chan bool)
	start()
}
