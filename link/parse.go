package link

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func GetHtmlLinks(n *html.Node, l *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				linkText := getLinkText(n)
				*l = append(*l, Link{attr.Val, linkText})
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		GetHtmlLinks(c, l)
	}
}

func getLinkText(n *html.Node) (linkText string) {
	if n.Type == html.TextNode {
		return n.Data
	}

	// process other nodes like <strong></strong> which contains textNode within
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		linkText += getLinkText(c)
	}

	// return a string with additional whitespaces stripped
	return strings.Join(strings.Fields(linkText), " ")
}
