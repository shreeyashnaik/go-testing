package main

import (
	"os"
	"testing"

	"github.com/shreeyashnaik/go-testing/web-app/pkg/repository/dbrepo"
)

var app application

// This function in (setup_test.go) gets executed before all the tests
func TestMain(m *testing.M) {
	app.Session = getSession()

	app.DB = &dbrepo.TestDBRepo{}

	os.Exit(m.Run())
}
