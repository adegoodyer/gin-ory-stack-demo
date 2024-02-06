package main

import (
	"flag"
	"fmt"
	"gin-ory-stack-demo/config"
	"gin-ory-stack-demo/server"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	config.Init(*environment)
	server.Init()
}
