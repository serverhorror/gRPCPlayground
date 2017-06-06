package sample

import (
	"encoding/xml"
	"time"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type file struct {
	XMLName   xml.Name  `xml:"sample-file" json:"-" db:"-"`
	ID        uint64    `xml:"id"  json:"id" db:"id"`
	SampleID  uint64    `xml:"sample_id" json:"sample_id" db:"sample_id"`
	SourceURI string    `xml:"source_uri" json:"source_uri" db:"source_uri"`
	CloudURI  string    `xml:"cloud_uri" json:"cloud_uri" db:"cloud_uri"`
	Creator   string    `xml:"creator" json:"creator" db:"creator"`
	Created   time.Time `xml:"created" json:"created" db:"created"`
	Provider  string    `xml:"provider" json:"provider" db:"provider"`
	Site      string    `xml:"site" json:"site" db:"site"`

	Sample `db:"sample"`

	isSaved      bool
	rawFile      string
	uploadOutput *s3manager.UploadOutput
}

type files []file
