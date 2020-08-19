package linkparser

import (
	"testing"
	"os"
	// "net/http"
	// "flag"
)

// var filename = flag.String("story", "stories/gopher.json", "json file containing cyoa story")

func TestMain(m *testing.M) {
	// flag.Parse()
	exitResponse := m.Run()
	os.Exit(exitResponse)
}

func TestDefaultHandler(t *testing.T) {

}
