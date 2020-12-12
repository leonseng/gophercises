package sitemap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"

	"github.com/leonseng/gophercises/link"
	"golang.org/x/net/html"
)

type Location struct {
	Value string `xml:"loc"`
}

type UrlSet struct {
	XmlNs     string     `xml:"xmlns,attr"`
	XMLName   xml.Name   `xml:"urlset"`
	Locations []Location `xml:"url"`
}

func Print(urlSet UrlSet) {
	w := &bytes.Buffer{}
	w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"))

	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")

	if err := enc.Encode(urlSet); err != nil {
		panic(err)
	}

	fmt.Println(w.String())
}

func Build(siteUrl string, maxDepth int) (UrlSet) {
	siteLocations := make(map[string]bool) // using map to quickly check if a site has been parsed previously
	ParseSite(siteUrl, maxDepth, siteLocations)

	// convert siteLocations to a list of strings as specified in UrlSet
	siteLocationKeys := make([]Location, len(siteLocations))
	i := 0
	for k := range siteLocations {
		siteLocationKeys[i].Value = k
		i++
	}

	return UrlSet{
		XmlNs:     "http://www.sitemaps.org/schemas/sitemap/0.9",
		Locations: siteLocationKeys,
	}
}

func ParseSite(siteUrl string, maxDepth int, siteLocations map[string]bool) {
	// fmt.Printf("Checking %s\n", siteUrl)
	if _, ok := siteLocations[siteUrl]; ok { // site link parsed previously, skip
		return
	}

	siteLocations[siteUrl] = true
	// fmt.Printf("Added %s\n", siteUrl)

	// final level, stop processing
	if maxDepth <= 1 {
		return
	}

	// fetch site
	resp, err := http.Get(siteUrl)
	if err != nil {
		panic(err)
	}
	parentScheme, parentHost := resp.Request.URL.Scheme, resp.Request.URL.Host

	// get all links in site
	var links []link.Link
	doc, _ := html.Parse(resp.Body)
	link.GetHtmlLinks(doc, &links)

	for _, l := range links {
		hrefLink, err := url.Parse(l.Href)
		if err != nil {
			panic(err)
		}

		// ignore links in different domains (/just-the-path or https://domain.com/with-domain)
		if hrefLink.Host != "" && hrefLink.Host != parentHost {
			continue
		}

		// build full url for link
		linkUrl := fmt.Sprintf("%s://%s%s", parentScheme, parentHost, hrefLink.Path)
		ParseSite(linkUrl, maxDepth-1, siteLocations)
	}
}
