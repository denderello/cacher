package main

import (
	"log"

	"github.com/denderello/cacher/db"
	"github.com/denderello/cacher/server"
	"github.com/spf13/pflag"
)

var (
	serverAddr   = pflag.String("server-addr", "0.0.0.0", "Address the server listens on")
	serverPort   = pflag.String("server-port", "8080", "Port the server listens on")
	databaseFile = pflag.String("database-file", "db.csv", "Databse file to use")
)

func main() {
	pflag.Parse()

	db, err := db.NewSlowFileDatabase(*databaseFile)
	if err != nil {
		log.Fatalf("Could not read databse file: %v", err)
	}

	if err := db.Open(); err != nil {
		log.Fatalf("Could not open database file: %v", err)
	}
	defer db.Close()

	s := server.NewHttpServer()
	s.Start(*serverAddr, *serverPort)
}
