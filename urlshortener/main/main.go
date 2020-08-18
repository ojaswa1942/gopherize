// Depicts sample usage
package main

import (
	"fmt"
	"net/http"
	"github.com/ojaswa1942/gopherize/urlshortener"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /gopherize
  url: https://github.com/ojaswa1942/gopherize
- path: /urlshort-final
  url: https://github.com/ojaswa1942/gopherize/urlshortener
`
	yamlHandler, err := urlshortener.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	// Build the JSONHandler using the YAMLHandler as the
	// fallback
	json := `[
		{"Path": "/site", "URL": "https://ojaswa.com"},
		{"Path": "/git", "URL": "https://github.com/ojaswa1942"},
		{"Path": "/connect", "URL": "https://linkedin.com/in/ojaswa23"}
	]`

	jsonHandler, errr := urlshortener.JSONHandler([]byte(json), yamlHandler)
	if errr != nil {
		panic(errr)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	// Multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}