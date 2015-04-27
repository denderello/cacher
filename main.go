package main

import (
	"log"
	"sync"

	"github.com/denderello/cacher/db"
	"github.com/denderello/cacher/handler"
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

	m := &sync.Mutex{}

	s := server.NewHttpServer()

	s.RegisterHandler("/{key}", "GET", handler.NewGetKeyHandler(db))
	s.RegisterHandler("/{key}/{value}", "POST", handler.NewSetKeyHandler(m, db))

	s.Start(*serverAddr, *serverPort)
}
