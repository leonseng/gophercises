package link

import (
	// "fmt"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(n *html.Node, l *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
	OuterLoop:
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				// fmt.Printf("%s\n", attr.Val)

				// find TextNode
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.TextNode {
						// fmt.Printf("%s\n", c.Data)
						*l = append(*l, Link{attr.Val, c.Data})
						break OuterLoop
					}
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Parse(c, l)
	}
}
