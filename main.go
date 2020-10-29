package main

import (
	"flag"
	"restserverfd/server"
)

func main() {
	port := flag.String("p", "50001", "localhost port")
	flag.Parse()
	server.Run(*port)
}
