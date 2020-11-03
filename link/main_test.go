package main

import (
	"fmt"
	"link/link"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRun(t *testing.T) {
	testFile := "html/ex4.html"
	links, err := run(testFile)
	if err != nil {
		t.Error("Failed to parse " + testFile)
	}
	expected := link.Link{"/dog-cat", "dog cat "}

	if len(links) != 1 {
		t.Error(fmt.Sprintf("Returned result has %d elements, expected 1", len(links)))
	}

	if !cmp.Equal(links[0], expected) {
		t.Error(fmt.Sprintf("Returned result %+v, expected %+v", links[0], expected))
	}
}
