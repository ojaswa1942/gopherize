package urlshortener

import (
	"testing"
	"os"
)

var pathsToUrlsMap = map[string]string{
	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}

var sampleYaml = []byte(`
- path: /gopherize
  url: https://github.com/ojaswa1942/gopherize
- path: /urlshort-final
  url: https://github.com/ojaswa1942/gopherize/urlshortener
`)

var sampleJson = []byte(`[
	{"Path": "/site", "URL": "https://ojaswa.com"},
	{"Path": "/git", "URL": "https://github.com/ojaswa1942"},
	{"Path": "/connect", "URL": "https://linkedin.com/in/ojaswa23"}
]`)

func TestMain(m *testing.M) {
	exitResponse := m.Run()
	os.Exit(exitResponse)
}

func TestJSONParser(t *testing.T) {
	parsedPaths, err := parseJSON(sampleJson)
	if err != nil {
		t.Error("error parsing JSON: ", err)
	}
	testPathArray(t, parsedPaths)
}

func TestYAMLParser(t *testing.T) {
	parsedPaths, err := parseYml(sampleYaml)
	if err != nil {
		t.Error("error parsing YAML: ", err)
	}
	testPathArray(t, parsedPaths)
}

func TestCreateMap(t *testing.T) {
	t.Run("test yml to map", func (t *testing.T) {
		parsedPaths, _ := parseYml(sampleYaml)
		mappedPaths := createMap(parsedPaths)
		for i, urlPath := range parsedPaths {
			if mappedPaths[urlPath.Path] != urlPath.URL {
				t.Errorf("cannot match url for path entry #%d. expected: `%s`, found: `%s`", i, urlPath.URL, mappedPaths[urlPath.Path])
			}
		}
	})

	t.Run("test json to map", func (t *testing.T) {
		parsedPaths, _ := parseJSON(sampleJson)
		mappedPaths := createMap(parsedPaths)
		for i, urlPath := range parsedPaths {
			if mappedPaths[urlPath.Path] != urlPath.URL {
				t.Errorf("cannot match url for path entry #%d. expected: `%s`, found: `%s`", i, urlPath.URL, mappedPaths[urlPath.Path])
			}
		}
	})
}

func testPathArray(t *testing.T, paths []path) {
	t.Log("check for empty entries")
	for i, urlPath := range paths {
		if urlPath.Path == "" {
			t.Errorf("empty path found for entry #%d.", i)
		}
		if urlPath.URL == "" {
			t.Errorf("empty URL found for entry #%d.", i)
		}
	}
}
