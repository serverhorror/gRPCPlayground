package server

import (
	"fmt"
	"net/http"

	"github.com/serverhorror/gRPCPlayground/httpAPI/common"
)

// * save sample
// * save metadata

type sample struct{}

func (s *sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", common.Fn())
}

func Main() error {

	http.Handle("/sample", &sample{})

	return http.ListenAndServe("[::1]:8000", nil)
}
