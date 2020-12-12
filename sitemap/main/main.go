package main

import (
	"flag"

	"github.com/leonseng/gophercises/sitemap"
)

func main() {
	// get user input for site and max depth
	var siteUrl = flag.String("u", "https://calhoun.io", "URL to generate site map for")
	var maxDepth = flag.Int("d", 10, "Maximum number of links to follow from the top site")
	flag.Parse()

	urlSet := sitemap.Build(*siteUrl, *maxDepth)
	sitemap.Print(urlSet)
}
