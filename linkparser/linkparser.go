package linkparser

import (
	"fmt"
	"io"
	"golang.org/x/net/html"
	"strings"
)

// Represents a link in a HTML
type Link struct {
	URL, Text string
}

// Take HTML document as input and return
// slice of parsed Link
func ParseLinks(r io.Reader) ([]Link, error) {
	rootNode, err := html.Parse(r)
	if err != nil {
		fmt.Println("unable to generate tree")
		return nil, err
	}

	linkNodes := searchElement(rootNode, "a")
	linkNodesAnswer := genLinksDetails(linkNodes)

	return linkNodesAnswer, nil
}

// A generic function to find some ElementNode
func searchElement(node *html.Node, elem string) []*html.Node {
	if node == nil {
		return []*html.Node{}
	}

	if node.Type == html.ElementNode && node.Data == elem {
		return []*html.Node{node}
	}

	// DFS for each child node
	answerNodes := make([]*html.Node, 0)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		childSlice := searchElement(child, elem)
		answerNodes = append(answerNodes, childSlice...)
	}
	
	return answerNodes
}

// converts slice of nodes into required output format
func genLinksDetails(linkNodes []*html.Node) []Link {
	answer := make([]Link, len(linkNodes))
	for i, node := range linkNodes {
		answer[i].URL = extractHref(node)
		answer[i].Text = strings.TrimSpace(extractText(node))
	}
	return answer
}

// given a link node, extracts href
func extractHref(node *html.Node) (string) {
	href := ""
	for _, attribute := range node.Attr {
		if attribute.Key == "href" {
			href = attribute.Val
		}
	}
	return href
}

// given a node, extracts text within it
func extractText(node *html.Node) string {
	if node == nil {
		return ""
	}

	if node.Type == html.TextNode {
		return node.Data
	}

	// DFS for each child node
	text := ""
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		childText := extractText(child)
		text += childText
	}
	
	return text
}