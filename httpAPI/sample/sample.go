package sample

import (
	"errors"
	"time"

	"github.com/serverhorror/gRPCPlayground/httpAPI/label"
	"github.com/serverhorror/gRPCPlayground/httpAPI/tag"
)

var (
	notImplemented = errors.New("not implemented")
)

type Sample struct {
	ID       uint64    `xml:"id" json:"id" db:"id"`
	Prefix   string    `xml:"prefix" json:"prefix" db:"prefix"`
	Provider string    `xml:"provider" json:"provider" db:"provider"`
	Site     string    `xml:"site" json:"site" db:"site"`
	Creator  string    `xml:"creator" json:"creator" db:"creator"`
	Created  time.Time `xml:"created" json:"created" db:"created"`

	tag.Tags
	label.Labels

	isDirty bool
}

func New(prefix, site string, labels label.Labels, tags tag.Tags) *Sample {
	return &Sample{
		Prefix:  prefix,
		Site:    site,
		Labels:  labels,
		Tags:    tags,
		Created: time.Now(),
		isDirty: true,
	}
}

func LoadByID(id uint64) (*Sample, error) {

	// var s *Sample
	return nil, notImplemented
}

func (s *Sample) Files() files {
	return nil
}

// func (s *Sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Content-Type: %v\n", r.Header.Get("content-type"))
// 	fmt.Fprintf(w, "%v\n", common.Fn())
// }
