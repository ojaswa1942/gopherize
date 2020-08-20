package main

import (
	"github.com/ojaswa1942/gopherize/linkparser"
	"net/http"
	"net/url"
	"fmt"
	"flag"
	"os"
	"io"
	"strings"
	"encoding/xml"
)

var xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type urlset struct {
	Urls []loc `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}
type loc struct {
	Value string `xml:"loc"`
}

func main() {
	urlFlag := flag.String("url", "https://ojaswa.com", "the url you want sitemap for")
	maxDepth := flag.Int("depth", 40, "maximum depth of links")
	flag.Parse()

	pages := traverse(*urlFlag, *maxDepth)

	generateXML(pages)

}

func generateXML(pages []string) {
	toXML := urlset{
		Urls: make([]loc, 0, len(pages)),
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXML.Urls = append(toXML.Urls, loc{page})
	}

	fmt.Printf(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXML); err != nil {
		exit(fmt.Sprintf("error while encoding: %s", err), true)
	}
	fmt.Println()
}

// For sorta of set
type empty struct{}

func traverse(urlStr string, maxDepth int) ([]string) {
	seen := make(map[string]empty)
	var q map[string]empty
	nextq := map[string]empty{
		urlStr: empty{},
	}

	for i := 0; i <= maxDepth; i++ {
		q, nextq = nextq, make(map[string]empty)
		if len(q) == 0 {
			break
		}
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = empty{}
			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					nextq[link] = empty{}
				}
			}
		}
	}

	answer := make([]string, 0, len(seen))
	for url, _ := range seen {
		answer = append(answer, url)
	}
	return answer
}

func get(urlString string) []string {
	resp, err := http.Get(urlString)
	if err != nil {
		fmt.Printf("error while parsing html: %s", err)
		return []string{}
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host: reqUrl.Host,
	}
	base := baseUrl.String()
	links := getHrefs(resp.Body, baseUrl)

	return filter(links, withPrefix(base))
}

func getHrefs(htmlReader io.Reader, baseUrl *url.URL) []string {
	base := baseUrl.String()
	links, _ := linkparser.ParseLinks(htmlReader)
	var hrefs []string
	
	for _, link := range links {
		switch {
			case strings.HasPrefix(link.URL, "//"):
				hrefs = append(hrefs, baseUrl.Scheme + link.URL[1:])
			case strings.HasPrefix(link.URL, "/"):
				hrefs = append(hrefs, base + link.URL)
			case strings.HasPrefix(link.URL, "http"):
				hrefs = append(hrefs, link.URL)
		}
	}
	return hrefs
}

func filter(links []string, keepFn func(string) bool) []string {
	var filtered []string
	for _, link := range links {
		if keepFn(link) {
			filtered = append(filtered, link)
		}
	}
	return filtered
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

func exit(msg string, error bool) {
	fmt.Println(msg)
	if error {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
