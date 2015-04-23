package server

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/juju/errgo"
)

type HttpServer struct {
	router *mux.Router
}

func NewHttpServer() *HttpServer {
	hs := &HttpServer{
		router: mux.NewRouter(),
	}

	http.Handle("/", handlers.LoggingHandler(os.Stdout, hs.router))

	return hs
}

func (hs *HttpServer) RegisterHandler(route, method string, handler http.Handler) {
	hs.router.Handle(route, handler).
		Methods(method)
}

func (hs *HttpServer) Start(addr, port string) error {
	log.Printf("Starting http server at %s:%s", addr, port)
	return errgo.Mask(http.ListenAndServe(net.JoinHostPort(addr, port), nil))
}
