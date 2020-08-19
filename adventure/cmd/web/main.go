package main

import (
	"github.com/ojaswa1942/gopherize/adventure"
	"fmt"
	"flag"
	"os"
	"log"
	"net/http"
	"html/template"
	"strings"
)

func main() {
	port := flag.Int("port", 3000, "port to start application on")
	filename := flag.String("story", "stories/gopher.json", "json file containing cyoa story")
	flag.Parse()

	encodedStory := openFile(*filename)
	story, err := adventure.DecodeJSONStory(encodedStory)
	if err != nil {
		exit(fmt.Sprintf("error parsing json: %s", err), true)
	}
	customTemplate := template.Must(template.New("").Parse("Hello yo !"))
	_ = customTemplate
	_ = customPathFn
	// h := adventure.NewHandler(story, adventure.WithTemplate(customTemplate))
	// h := adventure.NewHandler(story, adventure.WithPathFn(customPathFn), adventure.WithURLPrefix("/story"))
	h := adventure.NewHandler(story)
	fmt.Printf("Starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		exit(fmt.Sprintf("cannot open file: %s", err), true)
	}
	return file
}

func customPathFn(r *http.Request) (string) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "story/intro"
	}
	path = path[len("story/"):]
	return path
}


func exit(msg string, error bool) {
	fmt.Println(msg)
	if error {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}