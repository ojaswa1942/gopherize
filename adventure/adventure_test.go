package adventure

import (
	"testing"
	"os"
)


func TestMain(m *testing.M) {
	exitResponse := m.Run()
	os.Exit(exitResponse)
}

func TestCSV(t *testing.T) {

}
