package main

import (
	"fmt"
	"net/http"

	"github.com/FreezeSnail/urlShortener/src/db"
	api "github.com/FreezeSnail/urlShortener/src/http/rest"
)

func main() {

	db, err := db.NewSQLite("urlshortener.db")
	if err != nil {
		panic(err)
	}

	s := api.URLShortnerServer{
		DB: db,
	}
	h := api.Handler(s)

	//TODO(SNAIL) Add background process to check token ttl

	fmt.Print("Hello world\n")
	http.ListenAndServe(":3000", h)
	fmt.Print("goodbye\n")
}
