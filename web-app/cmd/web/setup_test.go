package main

import (
	"os"
	"testing"
)

var app application

// This function in (setup_test.go) gets executed before all the tests
func TestMain(m *testing.M) {
	app.Session = getSession()

	pathToTemplates = "./../../templates"

	os.Exit(m.Run())
}
