package net

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestURL(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Test")
	}))
	defer ts.Close()

	res, _ := RequestURL(ts.URL)
	if res != "Test" {
		t.Fatalf("Actual - %s", res)
	}
}
