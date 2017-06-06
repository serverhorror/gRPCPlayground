package sample

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
)

type testInput struct {
	method     string
	requestURL string
}
type expectedResult struct {
	StatusCode int
	MediaType  string
	Body       []byte
}

type testCase struct {
	in       *testInput
	expected *expectedResult
}

var (
	errNotImplemted = errors.New("not implement")

	testCases []testCase
)

// TestByID ensures that loading a Sample by its ID will yield the expected result
func TestFilterByID(t *testing.T) {
	testCases = []testCase{
		{in: &testInput{
			method:     http.MethodGet,
			requestURL: "http://example.invalid/sample?id=1",
		}, expected: &expectedResult{
			StatusCode: http.StatusOK,
			MediaType:  "application/json",
			Body:       []byte(`{"id":1}`),
		}},
		{in: &testInput{
			method:     http.MethodGet,
			requestURL: "http://example.invalid/sample?id=2",
		}, expected: &expectedResult{
			StatusCode: http.StatusOK,
			MediaType:  "application/json",
			Body:       []byte(`{"id":2}`),
		}},
	}

	for _, test := range testCases {
		t.Fatal(errNotImplemted)
		t.Log(test)
	}
}

// TestTagFilter ensures that loading Samples with a tag filter will yield the expected result
func TestFilterByTag(t *testing.T) {

	testCases = []testCase{
		{in: &testInput{
			method:     http.MethodGet,
			requestURL: "http://example.invalid/sample?tag=t1",
		}, expected: &expectedResult{
			StatusCode: http.StatusOK,
			MediaType:  "application/json",
			Body:       []byte(`{}`),
		}},
		{in: &testInput{
			method:     http.MethodGet,
			requestURL: "http://example.invalid/sample?tag=t1+t2",
		}, expected: &expectedResult{
			StatusCode: http.StatusOK,
			MediaType:  "application/json",
			Body:       []byte(`{tags":["t1"]}`),
		}},
	}

	for _, test := range testCases {
		req, err := http.NewRequest(test.in.method, test.in.requestURL, nil)
		if err != nil {
			t.Fatal(errors.Cause(err))
		}
		w := httptest.NewRecorder()
		s := &Sample{}
		s.ServeHTTP(w, req)

		resp := w.Result()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(errors.Cause(err))
		}

		if resp.StatusCode != test.expected.StatusCode {
			t.Fatal(fmt.Errorf("got: %v, expected: %v", resp.StatusCode, test.expected.StatusCode))
		}

		mediaType, _, err := mime.ParseMediaType(resp.Header.Get("content-type"))
		if err != nil {
			t.Fatal(errors.Cause(err))
		}

		if mediaType != test.expected.MediaType {
			t.Fatal(fmt.Errorf("got: %q, expected: %q", mediaType, test.expected.MediaType))
		}

		if !bytes.Equal(body, test.expected.Body) {
			t.Fatal(fmt.Errorf("got: %q, expected: %q", body, test.expected.Body))
		}
	}
}
