package main

import (
	"flag"

	"github.com/tylerd-lloyd/yamlserver/server"
)

func main() {
	port := flag.String("p", "50001", "localhost port")
	flag.Parse()
	server.Run(*port)
}
