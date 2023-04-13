package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	database "github.com/FreezeSnail/urlShortener/src/db"
	api "github.com/FreezeSnail/urlShortener/src/http/rest"
)

func main() {

	db, err := sql.Open("sqlite3", "urlshortner.db")
	if err != nil {
		panic(err)
	}
	err = database.RunMigrations(db)
	if err != nil {
		panic(err)
	}

	s := api.URLShortnerServer{}
	h := api.Handler(s)

	fmt.Print("Hello world\n")
	http.ListenAndServe(":3000", h)
	fmt.Print("goodbye\n")
}
