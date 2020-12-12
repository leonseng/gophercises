package link

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/html"

	"github.com/google/go-cmp/cmp"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		testFile string
		expected []Link
	}{
		{
			"tests/html/ex1.html",
			[]Link{
				Link{Href: "/other-page", Text: "A link to another page"},
			},
		},
		{
			"tests/html/ex2.html",
			[]Link{
				Link{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
				Link{Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"},
			},
		},
		{
			"tests/html/ex3.html",
			[]Link{
				Link{Href: "#", Text: "Login"},
				Link{Href: "/lost", Text: "Lost? Need help?"},
				Link{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
			},
		},
		{
			"tests/html/ex4.html",
			[]Link{
				Link{Href: "/dog-cat", Text: "dog cat"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testFile, func(t *testing.T) {
			file, err := os.Open(tc.testFile)
			if err != nil {
				t.Error(err)
			}

			doc, err := html.Parse(file)
			if err != nil {
				t.Error(err)
			}

			links := make([]Link, 0, 20)
			GetHtmlLinks(doc, &links)
			if err != nil {
				t.Error(err)
			}

			if len(links) != len(tc.expected) {
				t.Error(fmt.Sprintf("Returned result has %d elements, expected %d", len(links), len(tc.expected)))
			}

			for i, testV := range tc.expected {
				if !cmp.Equal(links[i], testV) {
					t.Error(fmt.Sprintf("Returned result %+v, expected %+v", links[i], testV))
				}
			}
		})
	}
}
