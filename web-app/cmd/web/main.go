package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/shreeyashnaik/go-testing/web-app/pkg/db"
)

type application struct {
	Session *scs.SessionManager
	DB      db.PostgresConn
	DSN     string
}

func main() {
	// set up an app config
	app := application{}

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5433 user=postgres password=postgres dbname=users sslmode=disable timezone=IST connect_timeout=5", "Postgres connection")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = db.PostgresConn{DB: conn}

	// get a session manager
	app.Session = getSession()

	// get application routes
	mux := app.routes()

	// print out a message
	log.Println("Starting server on port 8080..")

	// start the server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
