package main

import (
	"log"
	"os"
	"testing"

	"github.com/shreeyashnaik/go-testing/web-app/pkg/db"
)

var app application

// This function in (setup_test.go) gets executed before all the tests
func TestMain(m *testing.M) {
	app.Session = getSession()
	app.DSN = "host=localhost port=5433 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5"

	pathToTemplates = "./../../templates"

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = db.PostgresConn{DB: conn}

	os.Exit(m.Run())
}
