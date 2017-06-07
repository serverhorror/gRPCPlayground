package main

import (
	"net/http"

	"github.com/serverhorror/gRPCPlayground/httpAPI/server"
)

func main() {

	var (
		handlers = map[string]http.Handler{
			"/swagger-ui/": http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./dist"))),
		}
		handleFuncs map[string]http.HandlerFunc
	)

	server.Main(handlers, handleFuncs)
}
