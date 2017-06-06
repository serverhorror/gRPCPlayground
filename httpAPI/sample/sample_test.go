package sample

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestQueryFiles(t *testing.T) {
	var err error
	var tx *sqlx.Tx
	db, err := sqlx.Connect("postgres", "postgres://reader:reader@vie-bio-postgres/marcherm?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	tx, err = db.Beginx()
	if err != nil {
		t.Fatal(err)
	}

	var (
		f         files
		fileQuery = `
		SELECT
			sf.id AS "id",
			sf.sample_id AS "sample_id",
			sf.source_uri AS "source_uri",
			sf.cloud_uri AS "cloud_uri",
			sf.creator AS "creator",
			s.id AS "sample.id",
			s.prefix AS "sample.prefix",
			s.provider AS "sample.provider",
			s.site AS "sample.site",
			s.creator AS "sample.creator",
			s.created AS "sample.created"
		FROM
			samples s JOIN
			sample_files sf ON (s.id = sf.sample_id)
		WHERE
			s.id = 28523
		`
	)

	err = tx.Select(&f, fileQuery)
	if err != nil {
		t.Fatal(err)
	}

	t.Errorf("f: %v", spew.Sdump(f))
	t.Fatal()
}

func TestLoadByID(t *testing.T) {
	tests := []struct {
		name    string
		id      uint64
		want    *Sample
		wantErr bool
	}{
		{id: 0,
			wantErr: true},
		{id: 1,
			want: &Sample{
				ID: 1,
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
