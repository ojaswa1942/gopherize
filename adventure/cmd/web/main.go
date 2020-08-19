package main

import (
	"fmt"
	"flag"
	"os"
	"log"
	"net/http"
	"../.."
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
	h := adventure.NewHandler(story)
	fmt.Printf("Starting server on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
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