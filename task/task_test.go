package main

import (
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	exitResponse := m.Run()
	os.Exit(exitResponse)
}

func TestPathArray(t *testing.T) {

}
