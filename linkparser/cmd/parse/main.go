package main

import (
	"github.com/ojaswa1942/gopherize/linkparser"
	"fmt"
	"flag"
	"os"
)

func main() {
	filename := flag.String("html", "html_samples/ex1.html", "json file containing cyoa story")
	flag.Parse()

	htmlFile := openFile(*filename)
	links, err := linkparser.ParseLinks(htmlFile)
	if err != nil {
		exit(fmt.Sprintf("error parsing json: %s", err), true)
	}

	fmt.Printf("%-10s%-50s%s\n", "Number", "URL", "Text")
	for i, link := range links {
		fmt.Printf("%-10d%-50s%s\n", i+1, link.URL, link.Text)
	}
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		exit(fmt.Sprintf("cannot open file: %s", err), true)
	}
	return file
}

func exit(msg string, error bool) {
	fmt.Println(msg)
	if error {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}