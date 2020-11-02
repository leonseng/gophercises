package main

import (
	"flag"
	"fmt"
	"link/link"
	"os"

	"golang.org/x/net/html"
)

func main() {
	var htmlFile = flag.String("f", "", "HTML file to parse")
	flag.Parse()

	links, err := run(*htmlFile)
	if err != nil {
		panic("something is not right")
	}

	fmt.Printf("%+v\n", links)
}

func run(html_file string) ([]link.Link, error) {
	file, err := os.Open(html_file)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(file)

	if err != nil {
		return nil, err
	}

	l := make([]link.Link, 0, 20)
	link.Parse(doc, &l)
	return l, nil
}
