package server

import "net/http"

func Main() error {

	http.Handle("/sample", nil)

	return http.ListenAndServe("[::1]:8000", nil)
}
