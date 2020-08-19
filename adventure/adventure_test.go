package adventure

import (
	"testing"
	"os"
	"net/http"
	"flag"
)

var filename = flag.String("story", "stories/gopher.json", "json file containing cyoa story")

func TestMain(m *testing.M) {
	flag.Parse()
	exitResponse := m.Run()
	os.Exit(exitResponse)
}

func TestDefaultPathFn(t *testing.T) {
	urlMapping := []struct{
		url, parsedPath string
	}{
		{ "https://ojaswa.com", "intro" },
		{ "https://ojaswa.com/", "intro" },
		{ "https://ojaswa.com/intro", "intro" },
		{ "https://ojaswa.com/lalala", "lalala" },
		{ "https://ojaswa.com/yes/cmoncmon", "yes/cmoncmon" },
		{ "https://ojaswa.com/  hola ", "hola" },
	}
	t.Logf("testing against %d urls", len(urlMapping))
	for i, urlObj := range urlMapping {
		dummyRequest, err := http.NewRequest(http.MethodGet, urlObj.url, nil)
		if err != nil {
			t.Error("cannot create request for url: ", urlObj.url)
		}
		// t.Log(dummyRequest)
		if parsed := defaultPathFn(dummyRequest); parsed != urlObj.parsedPath {
			t.Errorf("parsed url #%d does not match. expected: `%s`, got: `%s`", i, urlObj.parsedPath, parsed)
		}
	}
}

func TestDecodeJSON(t *testing.T) {
	t.Log("opening file")
	encodedStory := openFile(*filename, t)
	t.Log("parsing file")
	story, err := DecodeJSONStory(encodedStory)
	if err != nil {
		t.Errorf("error parsing json: %s", err)
	}

	if testChapter, okay := story["custom-test-key"]; okay {
		t.Log("identified default/custom storybook")

		t.Run("assert number of chapters", func (t *testing.T) {
			totalChapters := 8
			if len(story) != totalChapters {
				t.Errorf("map entry count mismatch. expected: %d, found: %d", totalChapters, len(story))
			}
		})

		t.Run("assert custom key", func (t *testing.T) {
			if testChapter.Title != "Test Me Senpai" {
				t.Errorf("incorrect custom test title. expected: %s, found: %s", "Test Me Senpai", testChapter.Title)
			}
		})
	}
}

func TestDefaultHandler(t *testing.T) {
	// Will need to mock a request and verify the handler properties using
	// response write data
	// We cannot directly assert for `handler` values as http.Handler interface
	// does not expose anything except ServeHTTP function

	// encodedStory := openFile(*filename, t)
	// story, _ := DecodeJSONStory(encodedStory)
	// defaultHandler := handler(NewHandler(story))
	// _, ok := NewHandler(story).(handler)
	// t.Log(ok)
	// if defaultHandler.urlPrefix != "/" {
	// 	t.Errorf("incorrect default urlPrefix. expected: `/`, found: %s", defaultHandler.urlPrefix)
	// }
}

func openFile(filename string, t *testing.T) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		t.Errorf("cannot open file: %s", err)
	}
	return file
}