package linkparser

import (
	"testing"
	"os"
	"flag"
	"golang.org/x/net/html"
	"strings"
)

// var filename = flag.String("story", "stories/gopher.json", "json file containing cyoa story")
var filename = flag.String("html", "html_samples/test.html", "json file containing cyoa story")
var rootNode *html.Node

func TestMain(m *testing.M) {
	flag.Parse()
	exitResponse := m.Run()
	os.Exit(exitResponse)
}

func TestParseTree(t *testing.T) {
	htmlFile := openFile(*filename, t)
	var err error
	rootNode, err = html.Parse(htmlFile)
	if err != nil {
		t.Error("unable to generate tree")
	}

	if rootNode.Type != html.DocumentNode {
		t.Error("parsed root node is not DocumentElement")
	}
}

func TestSearchElement(t *testing.T) {
	// t.Log(rootNode)
	elementList := []struct{
		element string
		count int
	}{
		{"a", 3},
		{"h1", 1},
		{"img", 2},
		{"div", 9},
		{"body", 1},
	}

	for _, entry := range elementList {
		t.Log("searching", entry.element)
		if nodes := searchElement(rootNode, entry.element); len(nodes) != entry.count {
			t.Errorf("element count mismatch for element `%s`. expected: %d, received: %d\n", entry.element, entry.count, len(nodes))
		}
	}
}

func TestExtractHref(t *testing.T) {
	nodes := searchElement(rootNode, "a");
	for _, entry := range nodes {
		href := extractHref(entry);
		t.Logf("extracted href: %s", href)
		if href == "" {
			t.Logf("extracted href: %s", href)
			t.Errorf("invalid href")
		}
	}
}

func TestExtractText(t *testing.T) {
	elementList := []struct{
		element, firstText string
	}{
		{"a", "Login"},
		{"h1", "coding exercises for budding gophers"},
	}

	for _, entry := range elementList {
		t.Log("searching", entry.element)
		nodes := searchElement(rootNode, entry.element);
		if firstNodeText := strings.TrimSpace(extractText(nodes[0])); entry.firstText != firstNodeText {
			t.Errorf("element text mismatch for element `%s`. expected: `%s`, received: `%s`\n", entry.element, entry.firstText, firstNodeText)
		}
	}
}

func TestGenerateLinkDetails(t *testing.T) {
	nodes := searchElement(rootNode, "a");
	formattedLinks := genLinksDetails(nodes)
	if(len(nodes) != len(formattedLinks)) {
		t.Errorf("length mismatch. expected: %d, found: %d.", len(nodes), len(formattedLinks))
	}
}

func TestParseLinks(t *testing.T) {
	htmlFile := openFile(*filename, t)
	links, err := ParseLinks(htmlFile)
	if err != nil {
		t.Errorf("received error while parsing: %s", err)
	}
	expectedLinks := 3
	if len(links) != expectedLinks {
		t.Errorf("link count mismatch. expected: %d, received: %d", expectedLinks, len(links))
	}
}

func openFile(filename string, t *testing.T) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		t.Errorf("cannot open file: %s", err)
	}
	return file
}