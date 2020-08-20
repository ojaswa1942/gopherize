package main

import (
	"testing"
	"os"
	"flag"
)

var urlFlag = flag.String("url", "https://www.google.com", "the url you want sitemap for")
var maxDepth = flag.Int("depth", 2, "maximum depth of links")

func TestMain(m *testing.M) {
	flag.Parse()
	exitResponse := m.Run()
	os.Exit(exitResponse)
}

func TestURLFilter(t *testing.T) {
	links := []string{
		"thisisastring",
		"thisisalso",
		"isnot",
		"thinope",
		"hisnoter",
	}

	t.Run("using default filter fn", func(t *testing.T) {
		prefixs := []struct{
			prefix string
			count int
		}{
			{ "this", 2 },
			{ "is", 1 },
			{ "th", 3 },
		}

		for _, prefix := range prefixs {
			base, expectedLen := prefix.prefix, prefix.count
			filteredLinks := filter(links, withPrefix(base))
			if len(filteredLinks) != expectedLen {
				t.Errorf("filtered count mismatch for `%s`. expected: %d, received: %d", base, expectedLen, len(filteredLinks))
			}
		}
	})
	// can create & add custom fn if required
}

func TestGetPageLinks(t *testing.T) {
	siteLinks := get(*urlFlag)
	expectedLinks := 16

	if len(siteLinks) != expectedLinks {
		t.Errorf("page links count mismatch for `%s`. expected: %d, received: %d", *urlFlag, expectedLinks, len(siteLinks))
	}
}

func TestTraversal(t *testing.T) {
	sites := []struct{
		url string
		count int
	}{
		{ "https://www.google.com", 100 },
		{ "https://ojaswa.com", 1 },
		{ "http://invalidveryvery.neverever", 1 },
	}

	for _, site := range sites {
		siteLinks := traverse(site.url, *maxDepth)
		expectedLinks := site.count
		if len(siteLinks) != expectedLinks {
			t.Errorf("total links count mismatch for `%s`. expected: %d, received: %d", site.url, expectedLinks, len(siteLinks))
		}
	}
}

func openFile(filename string, t *testing.T) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		t.Errorf("cannot open file: %s", err)
	}
	return file
}