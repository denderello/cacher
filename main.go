package main

import (
	"github.com/denderello/cacher/server"
	"github.com/spf13/pflag"
)

var (
	serverAddr = pflag.String("server-addr", "0.0.0.0", "Address the server listens on")
	serverPort = pflag.String("server-port", "8080", "Port the server listens on")
)

func main() {
	pflag.Parse()

	s := server.NewHttpServer()
	s.Start(*serverAddr, *serverPort)
}
