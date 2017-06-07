package server

import "net/http"

func Main(handlers map[string]http.Handler, handleFuncs map[string]http.HandlerFunc) error {

	for prefix, handle := range handlers {
		http.Handle(prefix, handle)
	}
	for prefix, handle := range handleFuncs {
		http.HandleFunc(prefix, handle)
	}

	return http.ListenAndServe("[::1]:8000", nil)
}
