package sample

import (
	"errors"

	"github.com/serverhorror/gRPCPlayground/httpAPI/label"
	"github.com/serverhorror/gRPCPlayground/httpAPI/tag"
)

var (
	notImplemented = errors.New("not implemented")
)

type Option func(*[]Sample) error

func WithTags(t *tag.Tags) Option {
	return func(s *[]Sample) error {
		return notImplemented
	}
}

type Sample struct {
	ID     int64
	Labels *label.Labels
	Tags   *tag.Tags
}

type Samples []Sample

func New(labels *label.Labels, tags *tag.Tags) *Sample {
	s := &Sample{
		Labels: labels,
		Tags:   tags,
	}
	return s
}

func LoadByID(id int64) (Sample, error) {
	return nil, notImplemented
}

// func (s *Sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Content-Type: %v\n", r.Header.Get("content-type"))
// 	fmt.Fprintf(w, "%v\n", common.Fn())
// }
