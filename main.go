package main

import (
	"flag"
	"gopkg.in/yamlserver/server"
)

func main() {
	port := flag.String("p", "50001", "localhost port")
	flag.Parse()
	server.Run(*port)
}
